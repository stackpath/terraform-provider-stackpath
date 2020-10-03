package stackpath

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"reflect"

	"github.com/terraform-providers/terraform-provider-stackpath/stackpath/api/ipam/ipam_models"
	"github.com/terraform-providers/terraform-provider-stackpath/stackpath/api/storage/storage_models"
	"github.com/terraform-providers/terraform-provider-stackpath/stackpath/api/workload/workload_models"
	"golang.org/x/oauth2"
)

// InvalidClientSecretError models when a StackPath API OAuth 2 client ID is
// invalid, either due to an invalid format or because the client ID does not
// exist at StackPath.
type InvalidClientIDError struct{}

// NewInvalidClientIDError wraps an existing error as an invalid client ID error.
func NewInvalidClientIDError() *InvalidClientIDError {
	return &InvalidClientIDError{}
}

// Error returns a human-readable invalid client ID error message.
func (e *InvalidClientIDError) Error() string {
	return "invalid or unknown StackPath client ID"
}

// InvalidClientSecretError models when a StackPath API OAuth 2 client ID is
// correct, but the client secret is incorrect.
type InvalidClientSecretError struct{}

// NewInvalidClientSecretError wraps an existing error as an invalid client
// secret error.
func NewInvalidClientSecretError() *InvalidClientSecretError {
	return &InvalidClientSecretError{}
}

// Error returns a human-readable invalid client secret error message.
func (e *InvalidClientSecretError) Error() string {
	return "invalid StackPath client secret"
}

// APIError models an error received from the StackPath API.
type APIError struct {
	statusCode      int
	message         string
	requestID       string
	fieldViolations []fieldViolation
}

// Error satisfies the error interface for APIError.
func (e *APIError) Error() string {
	message := fmt.Sprintf(
		"a %d error was returned from StackPath: \"%s\"",
		e.statusCode,
		e.message,
	)

	if len(e.fieldViolations) > 0 {
		message = fmt.Sprintf("%s. The following fields have errors:", message)

		for i, violation := range e.fieldViolations {
			if i != 0 {
				message = fmt.Sprintf("%s,", message)
			}

			message = fmt.Sprintf("%s %s: %s", message, violation.field, violation.description)
		}
	}

	if e.requestID != "" {
		message = fmt.Sprintf("%s (request ID %s)", message, e.requestID)
	}

	return message
}

// fieldViolation models a StackPath API 400 error field violation in a single
// struct to ease type checking logic when sending errors to the user.
type fieldViolation struct {
	description string
	field       string
}

// HTTPStatusFromCode converts a gRPC error code into the corresponding HTTP response status.
// See: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
//
// This method is adapted from https://github.com/grpc-ecosystem/grpc-gateway/blob/master/runtime/errors.go
// to prevent having to import the entire grpc-gateway package.
func HTTPStatusFromCode(code int32) int {
	switch code {
	// OK
	case 0:
		return http.StatusOK
	// Canceled
	case 1:
		return http.StatusRequestTimeout
	// Unknown
	case 2:
		return http.StatusInternalServerError
	// InvalidArgument
	case 3:
		return http.StatusBadRequest
	// DeadlineExceeded
	case 4:
		return http.StatusGatewayTimeout
	// NotFound
	case 5:
		return http.StatusNotFound
	// AlreadyExists
	case 6:
		return http.StatusConflict
	// PermissionDenied
	case 7:
		return http.StatusForbidden
	// Unauthenticated
	case 16:
		return http.StatusUnauthorized
	// ResourceExhausted
	case 8:
		return http.StatusTooManyRequests
	// FailedPrecondition
	case 9:
		// Note, this deliberately doesn't translate to the similarly named '412 Precondition Failed' HTTP response status.
		return http.StatusBadRequest
	// Aborted
	case 10:
		return http.StatusConflict
	// OutOfRange
	case 11:
		return http.StatusBadRequest
	// Unimplemented
	case 12:
		return http.StatusNotImplemented
	// Internal
	case 13:
		return http.StatusInternalServerError
	// Unavailable
	case 14:
		return http.StatusServiceUnavailable
	// DataLoss
	case 15:
		return http.StatusInternalServerError
	}

	return http.StatusInternalServerError
}

// NewStackPathError factories common StackPath error scenarios into their own
// error types, or returns the original error.
func NewStackPathError(err error) error {
	switch rootErr := err.(type) {
	// Look for errors performing underlying OAuth 2 authentication.
	case *url.Error:
		switch typedErr := rootErr.Err.(type) {
		case *oauth2.RetrieveError:
			switch typedErr.Response.StatusCode {
			// A 401 Unauthorized error means the client ID was valid, but the
			// corresponding secret wasn't.
			case 401:
				return NewInvalidClientSecretError()

			// A 404 Not Found error means the client ID didn't exist.
			case 404:
				return NewInvalidClientIDError()
			}
		}
	}

	// Determine if this is a StackPath API error. StackPath API error messages
	// should have at least an HTTP status code and a message.
	var statusCode int
	var payload interface{}
	var fieldViolations []fieldViolation
	var message string
	var requestID string

	// There are a lot of generated API error structs, so inspect them with
	// reflection to get the underlying HTTP status code and Payload object.
	// Payload objects are easier to work with as each StackPath API service
	// has a single Payload struct.
	value := reflect.ValueOf(err)

	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	for i := 0; i < value.NumField(); i++ {
		f := value.Field(i)

		switch value.Type().Field(i).Name {
		case "_statusCode":
			statusCode = int(f.Int())
		case "Payload":
			if f.Kind() == reflect.Ptr {
				f = f.Elem()
			}

			payload = f.Interface()
		}
	}

	// Look through the payload for things like the error message, StackPath
	// request ID, and other error details.
	switch status := payload.(type) {
	case ipam_models.APIStatus:
		statusCode = HTTPStatusFromCode(status.Code)
		message = status.Message

		for _, d := range status.Details() {
			switch detail := d.(type) {
			case *ipam_models.StackpathRPCRequestInfo:
				requestID = detail.RequestID
			case *ipam_models.StackpathRPCBadRequest:
				for _, violation := range detail.FieldViolations {
					fieldViolations = append(fieldViolations, fieldViolation{
						description: violation.Description,
						field:       violation.Field,
					})
				}
			default:
				log.Printf("Received a %T detail from a StackPath API IPAM service error: %v", detail, detail)
			}
		}

	case workload_models.StackpathapiStatus:
		statusCode = HTTPStatusFromCode(status.Code)
		message = status.Message

		for _, d := range status.Details() {
			switch detail := d.(type) {
			case *workload_models.StackpathRPCRequestInfo:
				requestID = detail.RequestID
			case *workload_models.StackpathRPCBadRequest:
				for _, violation := range detail.FieldViolations {
					fieldViolations = append(fieldViolations, fieldViolation{
						description: violation.Description,
						field:       violation.Field,
					})
				}
			default:
				log.Printf("Received a %T detail from a StackPath API workload service error: %v", detail, detail)
			}
		}

	case storage_models.StackpathapiStatus:
		statusCode = HTTPStatusFromCode(status.Code)
		message = status.Message

		for _, d := range status.Details() {
			switch detail := d.(type) {
			case *storage_models.StackpathRPCRequestInfo:
				requestID = detail.RequestID
			case *storage_models.StackpathRPCBadRequest:
				for _, violation := range detail.FieldViolations {
					fieldViolations = append(fieldViolations, fieldViolation{
						description: violation.Description,
						field:       violation.Field,
					})
				}
			default:
				log.Printf("Received a %T detail from a StackPath API storage service error: %v", detail, detail)
			}
		}
	}

	// This wasn't a StackPath API error if there's no associated error message
	// and HTTP status code.
	if message == "" || statusCode == 0 {
		return err
	}

	// Log the underlying error in case the user is interested.
	log.Printf("Error received from the StackPath API. Request ID: \"%s\". Error: %s", requestID, err)

	return &APIError{
		statusCode:      statusCode,
		message:         message,
		requestID:       requestID,
		fieldViolations: fieldViolations,
	}
}

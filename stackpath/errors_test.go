package stackpath

import (
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/stackpath/terraform-provider-stackpath/stackpath/api/ipam/ipam_client/network_policies"
	"github.com/stackpath/terraform-provider-stackpath/stackpath/api/ipam/ipam_models"
	"golang.org/x/oauth2"
)

func TestBuildAnInvalidClientIDError(t *testing.T) {
	err := NewStackPathError(&url.Error{
		Err: &oauth2.RetrieveError{
			Response: &http.Response{
				StatusCode: 404,
			},
		},
	})

	switch err.(type) {
	case *InvalidClientIDError:
		if err.Error() != "invalid or unknown StackPath client ID" {
			t.Errorf("An invalid client ID error has the unexpected error message \"%s\"", err.Error())
		}
	default:
		t.Errorf("NewStackPathError() built an incorrect error type. Expected InvalidClientIDError but got %t", err)
	}
}

func TestBuildAnInvalidClientSecretError(t *testing.T) {
	err := NewStackPathError(&url.Error{
		Err: &oauth2.RetrieveError{
			Response: &http.Response{
				StatusCode: 401,
			},
		},
	})

	switch err.(type) {
	case *InvalidClientSecretError:
		if err.Error() != "invalid StackPath client secret" {
			t.Errorf("An invalid client ID error has the unexpected error message \"%s\"", err.Error())
		}
	default:
		t.Errorf("NewStackPathError() built an incorrect error type. Expected InvalidClientSecretError but got %t", err)
	}
}

func TestBuildANonStackPathError(t *testing.T) {
	type TestError struct {
		error
	}

	err := NewStackPathError(&TestError{errors.New("foo")})

	switch err.(type) {
	case *TestError:
	default:
		t.Errorf("NewStackPathError built an incorrect error type. Expected TestError but got %t", err)
	}
}

func TestBuildAnAPIError(t *testing.T) {
	expectedStatusCode := 500
	expectedMessage := "Internal Server Error"
	expectedRequestID := "deadbeef"
	expectedFullError := "a 500 error was returned from StackPath: \"Internal Server Error\" (request ID deadbeef)"

	result := network_policies.NewCreateNetworkPolicyInternalServerError()
	result.Payload = &ipam_models.APIStatus{
		Code:    13,
		Message: expectedMessage,
	}

	result.Payload.SetDetails([]ipam_models.APIStatusDetail{
		&ipam_models.StackpathRPCRequestInfo{
			RequestID: expectedRequestID,
		},
	})

	err := NewStackPathError(result)

	switch err.(type) {
	case *APIError:
	default:
		t.Errorf(
			"NewStackPathError built an incorrect error type. Expected APIError but got %T",
			err,
		)
	}

	if err.(*APIError).statusCode != expectedStatusCode {
		t.Errorf(
			"NewStackPathError built an APIError with an incorrect statusCode. Expected %d but got %d",
			expectedStatusCode,
			err.(*APIError).statusCode,
		)
	}

	if err.(*APIError).message != expectedMessage {
		t.Errorf(
			"NewStackPathError built an APIError with an incorrect message. Expected \"%s\" but got \"%s\"",
			expectedMessage,
			err.(*APIError).message,
		)
	}

	if err.(*APIError).requestID != expectedRequestID {
		t.Errorf(
			"NewStackPathError built an APIError with an incorrect requestID. Expected \"%s\" but got \"%s\"",
			expectedRequestID,
			err.(*APIError).requestID,
		)
	}

	if err.(*APIError).Error() != expectedFullError {
		t.Errorf(
			"NewStackPathError built an APIError that rendered an incorrect error message. Expected \"%s\" but got \"%s\"",
			expectedFullError,
			err.(*APIError).Error(),
		)
	}
}

func TestBuildAnAPIErrorWithNoRequestID(t *testing.T) {
	expectedFullError := "a 500 error was returned from StackPath: \"Internal Server Error\""

	result := network_policies.NewCreateNetworkPolicyInternalServerError()
	result.Payload = &ipam_models.APIStatus{
		Code:    13,
		Message: "Internal Server Error",
	}

	err := NewStackPathError(result)

	if err.(*APIError).Error() != expectedFullError {
		t.Errorf(
			"NewStackPathError built an APIError that rendered an incorrect error message. Expected \"%s\" but got \"%s\"",
			expectedFullError,
			err.(*APIError).Error(),
		)
	}
}

func TestBuildAnAPIErrorWithFieldViolations(t *testing.T) {
	expectedFullError := "a 400 error was returned from StackPath: \"network policy with priority 20000 already exists on stack\". The following fields have errors: foo: is required, bar: must be less than three (request ID deadbeef)"

	result := network_policies.NewCreateNetworkPolicyDefault(400)
	result.Payload = &ipam_models.APIStatus{
		Code:    3,
		Message: "network policy with priority 20000 already exists on stack",
	}

	result.Payload.SetDetails([]ipam_models.APIStatusDetail{
		&ipam_models.StackpathRPCRequestInfo{
			RequestID: "deadbeef",
		},
		&ipam_models.StackpathRPCBadRequest{
			FieldViolations: []*ipam_models.StackpathRPCBadRequestFieldViolation{
				{
					Field:       "foo",
					Description: "is required",
				},
				{
					Field:       "bar",
					Description: "must be less than three",
				},
			},
		},
	})

	err := NewStackPathError(result)

	if err.(*APIError).Error() != expectedFullError {
		t.Errorf(
			"NewStackPathError built an APIError that rendered an incorrect error message. Expected \"%s\" but got \"%s\"",
			expectedFullError,
			err.(*APIError).Error(),
		)
	}
}

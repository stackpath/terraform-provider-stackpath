// Code generated by go-swagger; DO NOT EDIT.

package user_credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/stackpath/terraform-provider-stackpath/v2/stackpath/api/storage/storage_models"
)

// DeleteCredentialReader is a Reader for the DeleteCredential structure.
type DeleteCredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteCredentialNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteCredentialUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteCredentialInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteCredentialDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteCredentialNoContent creates a DeleteCredentialNoContent with default headers values
func NewDeleteCredentialNoContent() *DeleteCredentialNoContent {
	return &DeleteCredentialNoContent{}
}

/* DeleteCredentialNoContent describes a response with status code 204, with default header values.

No content
*/
type DeleteCredentialNoContent struct {
}

func (o *DeleteCredentialNoContent) Error() string {
	return fmt.Sprintf("[DELETE /storage/v1/stacks/{stack_id}/users/{user_id}/credentials/{access_key}][%d] deleteCredentialNoContent ", 204)
}

func (o *DeleteCredentialNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCredentialUnauthorized creates a DeleteCredentialUnauthorized with default headers values
func NewDeleteCredentialUnauthorized() *DeleteCredentialUnauthorized {
	return &DeleteCredentialUnauthorized{}
}

/* DeleteCredentialUnauthorized describes a response with status code 401, with default header values.

Returned when an unauthorized request is attempted.
*/
type DeleteCredentialUnauthorized struct {
	Payload *storage_models.StackpathapiStatus
}

func (o *DeleteCredentialUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /storage/v1/stacks/{stack_id}/users/{user_id}/credentials/{access_key}][%d] deleteCredentialUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteCredentialUnauthorized) GetPayload() *storage_models.StackpathapiStatus {
	return o.Payload
}

func (o *DeleteCredentialUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(storage_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCredentialInternalServerError creates a DeleteCredentialInternalServerError with default headers values
func NewDeleteCredentialInternalServerError() *DeleteCredentialInternalServerError {
	return &DeleteCredentialInternalServerError{}
}

/* DeleteCredentialInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type DeleteCredentialInternalServerError struct {
	Payload *storage_models.StackpathapiStatus
}

func (o *DeleteCredentialInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /storage/v1/stacks/{stack_id}/users/{user_id}/credentials/{access_key}][%d] deleteCredentialInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteCredentialInternalServerError) GetPayload() *storage_models.StackpathapiStatus {
	return o.Payload
}

func (o *DeleteCredentialInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(storage_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCredentialDefault creates a DeleteCredentialDefault with default headers values
func NewDeleteCredentialDefault(code int) *DeleteCredentialDefault {
	return &DeleteCredentialDefault{
		_statusCode: code,
	}
}

/* DeleteCredentialDefault describes a response with status code -1, with default header values.

Default error structure.
*/
type DeleteCredentialDefault struct {
	_statusCode int

	Payload *storage_models.StackpathapiStatus
}

// Code gets the status code for the delete credential default response
func (o *DeleteCredentialDefault) Code() int {
	return o._statusCode
}

func (o *DeleteCredentialDefault) Error() string {
	return fmt.Sprintf("[DELETE /storage/v1/stacks/{stack_id}/users/{user_id}/credentials/{access_key}][%d] DeleteCredential default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteCredentialDefault) GetPayload() *storage_models.StackpathapiStatus {
	return o.Payload
}

func (o *DeleteCredentialDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(storage_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

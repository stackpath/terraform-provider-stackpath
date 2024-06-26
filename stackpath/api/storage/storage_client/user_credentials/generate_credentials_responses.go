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

// GenerateCredentialsReader is a Reader for the GenerateCredentials structure.
type GenerateCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GenerateCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGenerateCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGenerateCredentialsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGenerateCredentialsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGenerateCredentialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGenerateCredentialsOK creates a GenerateCredentialsOK with default headers values
func NewGenerateCredentialsOK() *GenerateCredentialsOK {
	return &GenerateCredentialsOK{}
}

/* GenerateCredentialsOK describes a response with status code 200, with default header values.

GenerateCredentialsOK generate credentials o k
*/
type GenerateCredentialsOK struct {
	Payload *storage_models.StorageGenerateCredentialsResponse
}

func (o *GenerateCredentialsOK) Error() string {
	return fmt.Sprintf("[POST /storage/v1/stacks/{stack_id}/users/{user_id}/credentials/generate][%d] generateCredentialsOK  %+v", 200, o.Payload)
}
func (o *GenerateCredentialsOK) GetPayload() *storage_models.StorageGenerateCredentialsResponse {
	return o.Payload
}

func (o *GenerateCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(storage_models.StorageGenerateCredentialsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateCredentialsUnauthorized creates a GenerateCredentialsUnauthorized with default headers values
func NewGenerateCredentialsUnauthorized() *GenerateCredentialsUnauthorized {
	return &GenerateCredentialsUnauthorized{}
}

/* GenerateCredentialsUnauthorized describes a response with status code 401, with default header values.

Returned when an unauthorized request is attempted.
*/
type GenerateCredentialsUnauthorized struct {
	Payload *storage_models.StackpathapiStatus
}

func (o *GenerateCredentialsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /storage/v1/stacks/{stack_id}/users/{user_id}/credentials/generate][%d] generateCredentialsUnauthorized  %+v", 401, o.Payload)
}
func (o *GenerateCredentialsUnauthorized) GetPayload() *storage_models.StackpathapiStatus {
	return o.Payload
}

func (o *GenerateCredentialsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(storage_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateCredentialsInternalServerError creates a GenerateCredentialsInternalServerError with default headers values
func NewGenerateCredentialsInternalServerError() *GenerateCredentialsInternalServerError {
	return &GenerateCredentialsInternalServerError{}
}

/* GenerateCredentialsInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type GenerateCredentialsInternalServerError struct {
	Payload *storage_models.StackpathapiStatus
}

func (o *GenerateCredentialsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /storage/v1/stacks/{stack_id}/users/{user_id}/credentials/generate][%d] generateCredentialsInternalServerError  %+v", 500, o.Payload)
}
func (o *GenerateCredentialsInternalServerError) GetPayload() *storage_models.StackpathapiStatus {
	return o.Payload
}

func (o *GenerateCredentialsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(storage_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateCredentialsDefault creates a GenerateCredentialsDefault with default headers values
func NewGenerateCredentialsDefault(code int) *GenerateCredentialsDefault {
	return &GenerateCredentialsDefault{
		_statusCode: code,
	}
}

/* GenerateCredentialsDefault describes a response with status code -1, with default header values.

Default error structure.
*/
type GenerateCredentialsDefault struct {
	_statusCode int

	Payload *storage_models.StackpathapiStatus
}

// Code gets the status code for the generate credentials default response
func (o *GenerateCredentialsDefault) Code() int {
	return o._statusCode
}

func (o *GenerateCredentialsDefault) Error() string {
	return fmt.Sprintf("[POST /storage/v1/stacks/{stack_id}/users/{user_id}/credentials/generate][%d] GenerateCredentials default  %+v", o._statusCode, o.Payload)
}
func (o *GenerateCredentialsDefault) GetPayload() *storage_models.StackpathapiStatus {
	return o.Payload
}

func (o *GenerateCredentialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(storage_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

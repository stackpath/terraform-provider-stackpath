// Code generated by go-swagger; DO NOT EDIT.

package buckets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/stackpath/terraform-provider-stackpath/v2/stackpath/api/storage/storage_models"
)

// DeleteBucketReader is a Reader for the DeleteBucket structure.
type DeleteBucketReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteBucketReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteBucketNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteBucketUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteBucketInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteBucketDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteBucketNoContent creates a DeleteBucketNoContent with default headers values
func NewDeleteBucketNoContent() *DeleteBucketNoContent {
	return &DeleteBucketNoContent{}
}

/*
DeleteBucketNoContent describes a response with status code 204, with default header values.

No content
*/
type DeleteBucketNoContent struct {
}

func (o *DeleteBucketNoContent) Error() string {
	return fmt.Sprintf("[DELETE /storage/v1/stacks/{stack_id}/buckets/{bucket_id}][%d] deleteBucketNoContent ", 204)
}

func (o *DeleteBucketNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteBucketUnauthorized creates a DeleteBucketUnauthorized with default headers values
func NewDeleteBucketUnauthorized() *DeleteBucketUnauthorized {
	return &DeleteBucketUnauthorized{}
}

/*
DeleteBucketUnauthorized describes a response with status code 401, with default header values.

Returned when an unauthorized request is attempted.
*/
type DeleteBucketUnauthorized struct {
	Payload *storage_models.StackpathapiStatus
}

func (o *DeleteBucketUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /storage/v1/stacks/{stack_id}/buckets/{bucket_id}][%d] deleteBucketUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteBucketUnauthorized) GetPayload() *storage_models.StackpathapiStatus {
	return o.Payload
}

func (o *DeleteBucketUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(storage_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteBucketInternalServerError creates a DeleteBucketInternalServerError with default headers values
func NewDeleteBucketInternalServerError() *DeleteBucketInternalServerError {
	return &DeleteBucketInternalServerError{}
}

/*
DeleteBucketInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type DeleteBucketInternalServerError struct {
	Payload *storage_models.StackpathapiStatus
}

func (o *DeleteBucketInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /storage/v1/stacks/{stack_id}/buckets/{bucket_id}][%d] deleteBucketInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteBucketInternalServerError) GetPayload() *storage_models.StackpathapiStatus {
	return o.Payload
}

func (o *DeleteBucketInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(storage_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteBucketDefault creates a DeleteBucketDefault with default headers values
func NewDeleteBucketDefault(code int) *DeleteBucketDefault {
	return &DeleteBucketDefault{
		_statusCode: code,
	}
}

/*
DeleteBucketDefault describes a response with status code -1, with default header values.

Default error structure.
*/
type DeleteBucketDefault struct {
	_statusCode int

	Payload *storage_models.StackpathapiStatus
}

// Code gets the status code for the delete bucket default response
func (o *DeleteBucketDefault) Code() int {
	return o._statusCode
}

func (o *DeleteBucketDefault) Error() string {
	return fmt.Sprintf("[DELETE /storage/v1/stacks/{stack_id}/buckets/{bucket_id}][%d] DeleteBucket default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteBucketDefault) GetPayload() *storage_models.StackpathapiStatus {
	return o.Payload
}

func (o *DeleteBucketDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(storage_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

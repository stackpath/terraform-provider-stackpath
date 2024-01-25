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

// GetBucketsReader is a Reader for the GetBuckets structure.
type GetBucketsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBucketsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBucketsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetBucketsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetBucketsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetBucketsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetBucketsOK creates a GetBucketsOK with default headers values
func NewGetBucketsOK() *GetBucketsOK {
	return &GetBucketsOK{}
}

/*
	GetBucketsOK describes a response with status code 200, with default header values.

GetBucketsOK get buckets o k
*/
type GetBucketsOK struct {
	Payload *storage_models.StorageGetBucketsResponse
}

func (o *GetBucketsOK) Error() string {
	return fmt.Sprintf("[GET /storage/v1/stacks/{stack_id}/buckets][%d] getBucketsOK  %+v", 200, o.Payload)
}
func (o *GetBucketsOK) GetPayload() *storage_models.StorageGetBucketsResponse {
	return o.Payload
}

func (o *GetBucketsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(storage_models.StorageGetBucketsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBucketsUnauthorized creates a GetBucketsUnauthorized with default headers values
func NewGetBucketsUnauthorized() *GetBucketsUnauthorized {
	return &GetBucketsUnauthorized{}
}

/*
	GetBucketsUnauthorized describes a response with status code 401, with default header values.

Returned when an unauthorized request is attempted.
*/
type GetBucketsUnauthorized struct {
	Payload *storage_models.StackpathapiStatus
}

func (o *GetBucketsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /storage/v1/stacks/{stack_id}/buckets][%d] getBucketsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetBucketsUnauthorized) GetPayload() *storage_models.StackpathapiStatus {
	return o.Payload
}

func (o *GetBucketsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(storage_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBucketsInternalServerError creates a GetBucketsInternalServerError with default headers values
func NewGetBucketsInternalServerError() *GetBucketsInternalServerError {
	return &GetBucketsInternalServerError{}
}

/*
	GetBucketsInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type GetBucketsInternalServerError struct {
	Payload *storage_models.StackpathapiStatus
}

func (o *GetBucketsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /storage/v1/stacks/{stack_id}/buckets][%d] getBucketsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetBucketsInternalServerError) GetPayload() *storage_models.StackpathapiStatus {
	return o.Payload
}

func (o *GetBucketsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(storage_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBucketsDefault creates a GetBucketsDefault with default headers values
func NewGetBucketsDefault(code int) *GetBucketsDefault {
	return &GetBucketsDefault{
		_statusCode: code,
	}
}

/*
	GetBucketsDefault describes a response with status code -1, with default header values.

Default error structure.
*/
type GetBucketsDefault struct {
	_statusCode int

	Payload *storage_models.StackpathapiStatus
}

// Code gets the status code for the get buckets default response
func (o *GetBucketsDefault) Code() int {
	return o._statusCode
}

func (o *GetBucketsDefault) Error() string {
	return fmt.Sprintf("[GET /storage/v1/stacks/{stack_id}/buckets][%d] GetBuckets default  %+v", o._statusCode, o.Payload)
}
func (o *GetBucketsDefault) GetPayload() *storage_models.StackpathapiStatus {
	return o.Payload
}

func (o *GetBucketsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(storage_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

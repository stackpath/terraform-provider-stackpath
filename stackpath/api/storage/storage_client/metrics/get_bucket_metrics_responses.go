// Code generated by go-swagger; DO NOT EDIT.

package metrics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/stackpath/terraform-provider-stackpath/v2/stackpath/api/storage/storage_models"
)

// GetBucketMetricsReader is a Reader for the GetBucketMetrics structure.
type GetBucketMetricsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBucketMetricsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBucketMetricsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetBucketMetricsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetBucketMetricsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetBucketMetricsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetBucketMetricsOK creates a GetBucketMetricsOK with default headers values
func NewGetBucketMetricsOK() *GetBucketMetricsOK {
	return &GetBucketMetricsOK{}
}

/*
	GetBucketMetricsOK describes a response with status code 200, with default header values.

GetBucketMetricsOK get bucket metrics o k
*/
type GetBucketMetricsOK struct {
	Payload *storage_models.PrometheusMetrics
}

func (o *GetBucketMetricsOK) Error() string {
	return fmt.Sprintf("[GET /storage/v1/stacks/{stack_id}/buckets/{bucket_id}/metrics][%d] getBucketMetricsOK  %+v", 200, o.Payload)
}
func (o *GetBucketMetricsOK) GetPayload() *storage_models.PrometheusMetrics {
	return o.Payload
}

func (o *GetBucketMetricsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(storage_models.PrometheusMetrics)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBucketMetricsUnauthorized creates a GetBucketMetricsUnauthorized with default headers values
func NewGetBucketMetricsUnauthorized() *GetBucketMetricsUnauthorized {
	return &GetBucketMetricsUnauthorized{}
}

/*
	GetBucketMetricsUnauthorized describes a response with status code 401, with default header values.

Returned when an unauthorized request is attempted.
*/
type GetBucketMetricsUnauthorized struct {
	Payload *storage_models.StackpathapiStatus
}

func (o *GetBucketMetricsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /storage/v1/stacks/{stack_id}/buckets/{bucket_id}/metrics][%d] getBucketMetricsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetBucketMetricsUnauthorized) GetPayload() *storage_models.StackpathapiStatus {
	return o.Payload
}

func (o *GetBucketMetricsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(storage_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBucketMetricsInternalServerError creates a GetBucketMetricsInternalServerError with default headers values
func NewGetBucketMetricsInternalServerError() *GetBucketMetricsInternalServerError {
	return &GetBucketMetricsInternalServerError{}
}

/*
	GetBucketMetricsInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type GetBucketMetricsInternalServerError struct {
	Payload *storage_models.StackpathapiStatus
}

func (o *GetBucketMetricsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /storage/v1/stacks/{stack_id}/buckets/{bucket_id}/metrics][%d] getBucketMetricsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetBucketMetricsInternalServerError) GetPayload() *storage_models.StackpathapiStatus {
	return o.Payload
}

func (o *GetBucketMetricsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(storage_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBucketMetricsDefault creates a GetBucketMetricsDefault with default headers values
func NewGetBucketMetricsDefault(code int) *GetBucketMetricsDefault {
	return &GetBucketMetricsDefault{
		_statusCode: code,
	}
}

/*
	GetBucketMetricsDefault describes a response with status code -1, with default header values.

Default error structure.
*/
type GetBucketMetricsDefault struct {
	_statusCode int

	Payload *storage_models.StackpathapiStatus
}

// Code gets the status code for the get bucket metrics default response
func (o *GetBucketMetricsDefault) Code() int {
	return o._statusCode
}

func (o *GetBucketMetricsDefault) Error() string {
	return fmt.Sprintf("[GET /storage/v1/stacks/{stack_id}/buckets/{bucket_id}/metrics][%d] GetBucketMetrics default  %+v", o._statusCode, o.Payload)
}
func (o *GetBucketMetricsDefault) GetPayload() *storage_models.StackpathapiStatus {
	return o.Payload
}

func (o *GetBucketMetricsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(storage_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// GetStackMetricsReader is a Reader for the GetStackMetrics structure.
type GetStackMetricsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStackMetricsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStackMetricsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetStackMetricsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetStackMetricsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetStackMetricsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetStackMetricsOK creates a GetStackMetricsOK with default headers values
func NewGetStackMetricsOK() *GetStackMetricsOK {
	return &GetStackMetricsOK{}
}

/*
	GetStackMetricsOK describes a response with status code 200, with default header values.

GetStackMetricsOK get stack metrics o k
*/
type GetStackMetricsOK struct {
	Payload *storage_models.PrometheusMetrics
}

func (o *GetStackMetricsOK) Error() string {
	return fmt.Sprintf("[GET /storage/v1/stacks/{stack_id}/metrics][%d] getStackMetricsOK  %+v", 200, o.Payload)
}
func (o *GetStackMetricsOK) GetPayload() *storage_models.PrometheusMetrics {
	return o.Payload
}

func (o *GetStackMetricsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(storage_models.PrometheusMetrics)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStackMetricsUnauthorized creates a GetStackMetricsUnauthorized with default headers values
func NewGetStackMetricsUnauthorized() *GetStackMetricsUnauthorized {
	return &GetStackMetricsUnauthorized{}
}

/*
	GetStackMetricsUnauthorized describes a response with status code 401, with default header values.

Returned when an unauthorized request is attempted.
*/
type GetStackMetricsUnauthorized struct {
	Payload *storage_models.StackpathapiStatus
}

func (o *GetStackMetricsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /storage/v1/stacks/{stack_id}/metrics][%d] getStackMetricsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetStackMetricsUnauthorized) GetPayload() *storage_models.StackpathapiStatus {
	return o.Payload
}

func (o *GetStackMetricsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(storage_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStackMetricsInternalServerError creates a GetStackMetricsInternalServerError with default headers values
func NewGetStackMetricsInternalServerError() *GetStackMetricsInternalServerError {
	return &GetStackMetricsInternalServerError{}
}

/*
	GetStackMetricsInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type GetStackMetricsInternalServerError struct {
	Payload *storage_models.StackpathapiStatus
}

func (o *GetStackMetricsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /storage/v1/stacks/{stack_id}/metrics][%d] getStackMetricsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetStackMetricsInternalServerError) GetPayload() *storage_models.StackpathapiStatus {
	return o.Payload
}

func (o *GetStackMetricsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(storage_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStackMetricsDefault creates a GetStackMetricsDefault with default headers values
func NewGetStackMetricsDefault(code int) *GetStackMetricsDefault {
	return &GetStackMetricsDefault{
		_statusCode: code,
	}
}

/*
	GetStackMetricsDefault describes a response with status code -1, with default header values.

Default error structure.
*/
type GetStackMetricsDefault struct {
	_statusCode int

	Payload *storage_models.StackpathapiStatus
}

// Code gets the status code for the get stack metrics default response
func (o *GetStackMetricsDefault) Code() int {
	return o._statusCode
}

func (o *GetStackMetricsDefault) Error() string {
	return fmt.Sprintf("[GET /storage/v1/stacks/{stack_id}/metrics][%d] GetStackMetrics default  %+v", o._statusCode, o.Payload)
}
func (o *GetStackMetricsDefault) GetPayload() *storage_models.StackpathapiStatus {
	return o.Payload
}

func (o *GetStackMetricsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(storage_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

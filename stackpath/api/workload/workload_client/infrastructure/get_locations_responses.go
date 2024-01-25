// Code generated by go-swagger; DO NOT EDIT.

package infrastructure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/stackpath/terraform-provider-stackpath/v2/stackpath/api/workload/workload_models"
)

// GetLocationsReader is a Reader for the GetLocations structure.
type GetLocationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLocationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLocationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetLocationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLocationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetLocationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLocationsOK creates a GetLocationsOK with default headers values
func NewGetLocationsOK() *GetLocationsOK {
	return &GetLocationsOK{}
}

/*
GetLocationsOK describes a response with status code 200, with default header values.

GetLocationsOK get locations o k
*/
type GetLocationsOK struct {
	Payload *workload_models.V1GetLocationsResponse
}

// IsSuccess returns true when this get locations o k response has a 2xx status code
func (o *GetLocationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get locations o k response has a 3xx status code
func (o *GetLocationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get locations o k response has a 4xx status code
func (o *GetLocationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get locations o k response has a 5xx status code
func (o *GetLocationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get locations o k response a status code equal to that given
func (o *GetLocationsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get locations o k response
func (o *GetLocationsOK) Code() int {
	return 200
}

func (o *GetLocationsOK) Error() string {
	return fmt.Sprintf("[GET /workload/v1/locations][%d] getLocationsOK  %+v", 200, o.Payload)
}

func (o *GetLocationsOK) String() string {
	return fmt.Sprintf("[GET /workload/v1/locations][%d] getLocationsOK  %+v", 200, o.Payload)
}

func (o *GetLocationsOK) GetPayload() *workload_models.V1GetLocationsResponse {
	return o.Payload
}

func (o *GetLocationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(workload_models.V1GetLocationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationsUnauthorized creates a GetLocationsUnauthorized with default headers values
func NewGetLocationsUnauthorized() *GetLocationsUnauthorized {
	return &GetLocationsUnauthorized{}
}

/*
GetLocationsUnauthorized describes a response with status code 401, with default header values.

Returned when an unauthorized request is attempted.
*/
type GetLocationsUnauthorized struct {
	Payload *workload_models.StackpathapiStatus
}

// IsSuccess returns true when this get locations unauthorized response has a 2xx status code
func (o *GetLocationsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get locations unauthorized response has a 3xx status code
func (o *GetLocationsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get locations unauthorized response has a 4xx status code
func (o *GetLocationsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get locations unauthorized response has a 5xx status code
func (o *GetLocationsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get locations unauthorized response a status code equal to that given
func (o *GetLocationsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get locations unauthorized response
func (o *GetLocationsUnauthorized) Code() int {
	return 401
}

func (o *GetLocationsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /workload/v1/locations][%d] getLocationsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLocationsUnauthorized) String() string {
	return fmt.Sprintf("[GET /workload/v1/locations][%d] getLocationsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLocationsUnauthorized) GetPayload() *workload_models.StackpathapiStatus {
	return o.Payload
}

func (o *GetLocationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(workload_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationsInternalServerError creates a GetLocationsInternalServerError with default headers values
func NewGetLocationsInternalServerError() *GetLocationsInternalServerError {
	return &GetLocationsInternalServerError{}
}

/*
GetLocationsInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type GetLocationsInternalServerError struct {
	Payload *workload_models.StackpathapiStatus
}

// IsSuccess returns true when this get locations internal server error response has a 2xx status code
func (o *GetLocationsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get locations internal server error response has a 3xx status code
func (o *GetLocationsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get locations internal server error response has a 4xx status code
func (o *GetLocationsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get locations internal server error response has a 5xx status code
func (o *GetLocationsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get locations internal server error response a status code equal to that given
func (o *GetLocationsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get locations internal server error response
func (o *GetLocationsInternalServerError) Code() int {
	return 500
}

func (o *GetLocationsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /workload/v1/locations][%d] getLocationsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLocationsInternalServerError) String() string {
	return fmt.Sprintf("[GET /workload/v1/locations][%d] getLocationsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLocationsInternalServerError) GetPayload() *workload_models.StackpathapiStatus {
	return o.Payload
}

func (o *GetLocationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(workload_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationsDefault creates a GetLocationsDefault with default headers values
func NewGetLocationsDefault(code int) *GetLocationsDefault {
	return &GetLocationsDefault{
		_statusCode: code,
	}
}

/*
GetLocationsDefault describes a response with status code -1, with default header values.

Default error structure.
*/
type GetLocationsDefault struct {
	_statusCode int

	Payload *workload_models.StackpathapiStatus
}

// IsSuccess returns true when this get locations default response has a 2xx status code
func (o *GetLocationsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get locations default response has a 3xx status code
func (o *GetLocationsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get locations default response has a 4xx status code
func (o *GetLocationsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get locations default response has a 5xx status code
func (o *GetLocationsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get locations default response a status code equal to that given
func (o *GetLocationsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get locations default response
func (o *GetLocationsDefault) Code() int {
	return o._statusCode
}

func (o *GetLocationsDefault) Error() string {
	return fmt.Sprintf("[GET /workload/v1/locations][%d] GetLocations default  %+v", o._statusCode, o.Payload)
}

func (o *GetLocationsDefault) String() string {
	return fmt.Sprintf("[GET /workload/v1/locations][%d] GetLocations default  %+v", o._statusCode, o.Payload)
}

func (o *GetLocationsDefault) GetPayload() *workload_models.StackpathapiStatus {
	return o.Payload
}

func (o *GetLocationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(workload_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

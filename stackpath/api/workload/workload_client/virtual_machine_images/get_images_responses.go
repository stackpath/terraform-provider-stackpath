// Code generated by go-swagger; DO NOT EDIT.

package virtual_machine_images

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/stackpath/terraform-provider-stackpath/v2/stackpath/api/workload/workload_models"
)

// GetImagesReader is a Reader for the GetImages structure.
type GetImagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetImagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetImagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetImagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetImagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetImagesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetImagesOK creates a GetImagesOK with default headers values
func NewGetImagesOK() *GetImagesOK {
	return &GetImagesOK{}
}

/*
GetImagesOK describes a response with status code 200, with default header values.

GetImagesOK get images o k
*/
type GetImagesOK struct {
	Payload *workload_models.V1GetImagesResponse
}

// IsSuccess returns true when this get images o k response has a 2xx status code
func (o *GetImagesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get images o k response has a 3xx status code
func (o *GetImagesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get images o k response has a 4xx status code
func (o *GetImagesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get images o k response has a 5xx status code
func (o *GetImagesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get images o k response a status code equal to that given
func (o *GetImagesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get images o k response
func (o *GetImagesOK) Code() int {
	return 200
}

func (o *GetImagesOK) Error() string {
	return fmt.Sprintf("[GET /workload/v1/stacks/{stack_id}/images][%d] getImagesOK  %+v", 200, o.Payload)
}

func (o *GetImagesOK) String() string {
	return fmt.Sprintf("[GET /workload/v1/stacks/{stack_id}/images][%d] getImagesOK  %+v", 200, o.Payload)
}

func (o *GetImagesOK) GetPayload() *workload_models.V1GetImagesResponse {
	return o.Payload
}

func (o *GetImagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(workload_models.V1GetImagesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetImagesUnauthorized creates a GetImagesUnauthorized with default headers values
func NewGetImagesUnauthorized() *GetImagesUnauthorized {
	return &GetImagesUnauthorized{}
}

/*
GetImagesUnauthorized describes a response with status code 401, with default header values.

Returned when an unauthorized request is attempted.
*/
type GetImagesUnauthorized struct {
	Payload *workload_models.StackpathapiStatus
}

// IsSuccess returns true when this get images unauthorized response has a 2xx status code
func (o *GetImagesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get images unauthorized response has a 3xx status code
func (o *GetImagesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get images unauthorized response has a 4xx status code
func (o *GetImagesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get images unauthorized response has a 5xx status code
func (o *GetImagesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get images unauthorized response a status code equal to that given
func (o *GetImagesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get images unauthorized response
func (o *GetImagesUnauthorized) Code() int {
	return 401
}

func (o *GetImagesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /workload/v1/stacks/{stack_id}/images][%d] getImagesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetImagesUnauthorized) String() string {
	return fmt.Sprintf("[GET /workload/v1/stacks/{stack_id}/images][%d] getImagesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetImagesUnauthorized) GetPayload() *workload_models.StackpathapiStatus {
	return o.Payload
}

func (o *GetImagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(workload_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetImagesInternalServerError creates a GetImagesInternalServerError with default headers values
func NewGetImagesInternalServerError() *GetImagesInternalServerError {
	return &GetImagesInternalServerError{}
}

/*
GetImagesInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type GetImagesInternalServerError struct {
	Payload *workload_models.StackpathapiStatus
}

// IsSuccess returns true when this get images internal server error response has a 2xx status code
func (o *GetImagesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get images internal server error response has a 3xx status code
func (o *GetImagesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get images internal server error response has a 4xx status code
func (o *GetImagesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get images internal server error response has a 5xx status code
func (o *GetImagesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get images internal server error response a status code equal to that given
func (o *GetImagesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get images internal server error response
func (o *GetImagesInternalServerError) Code() int {
	return 500
}

func (o *GetImagesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /workload/v1/stacks/{stack_id}/images][%d] getImagesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetImagesInternalServerError) String() string {
	return fmt.Sprintf("[GET /workload/v1/stacks/{stack_id}/images][%d] getImagesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetImagesInternalServerError) GetPayload() *workload_models.StackpathapiStatus {
	return o.Payload
}

func (o *GetImagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(workload_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetImagesDefault creates a GetImagesDefault with default headers values
func NewGetImagesDefault(code int) *GetImagesDefault {
	return &GetImagesDefault{
		_statusCode: code,
	}
}

/*
GetImagesDefault describes a response with status code -1, with default header values.

Default error structure.
*/
type GetImagesDefault struct {
	_statusCode int

	Payload *workload_models.StackpathapiStatus
}

// IsSuccess returns true when this get images default response has a 2xx status code
func (o *GetImagesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get images default response has a 3xx status code
func (o *GetImagesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get images default response has a 4xx status code
func (o *GetImagesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get images default response has a 5xx status code
func (o *GetImagesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get images default response a status code equal to that given
func (o *GetImagesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get images default response
func (o *GetImagesDefault) Code() int {
	return o._statusCode
}

func (o *GetImagesDefault) Error() string {
	return fmt.Sprintf("[GET /workload/v1/stacks/{stack_id}/images][%d] GetImages default  %+v", o._statusCode, o.Payload)
}

func (o *GetImagesDefault) String() string {
	return fmt.Sprintf("[GET /workload/v1/stacks/{stack_id}/images][%d] GetImages default  %+v", o._statusCode, o.Payload)
}

func (o *GetImagesDefault) GetPayload() *workload_models.StackpathapiStatus {
	return o.Payload
}

func (o *GetImagesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(workload_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

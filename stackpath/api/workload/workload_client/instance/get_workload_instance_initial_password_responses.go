// Code generated by go-swagger; DO NOT EDIT.

package instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/stackpath/terraform-provider-stackpath/v2/stackpath/api/workload/workload_models"
)

// GetWorkloadInstanceInitialPasswordReader is a Reader for the GetWorkloadInstanceInitialPassword structure.
type GetWorkloadInstanceInitialPasswordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkloadInstanceInitialPasswordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkloadInstanceInitialPasswordOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetWorkloadInstanceInitialPasswordUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkloadInstanceInitialPasswordInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetWorkloadInstanceInitialPasswordDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWorkloadInstanceInitialPasswordOK creates a GetWorkloadInstanceInitialPasswordOK with default headers values
func NewGetWorkloadInstanceInitialPasswordOK() *GetWorkloadInstanceInitialPasswordOK {
	return &GetWorkloadInstanceInitialPasswordOK{}
}

/* GetWorkloadInstanceInitialPasswordOK describes a response with status code 200, with default header values.

GetWorkloadInstanceInitialPasswordOK get workload instance initial password o k
*/
type GetWorkloadInstanceInitialPasswordOK struct {
	Payload *workload_models.V1GetWorkloadInstanceInitialPasswordResponse
}

func (o *GetWorkloadInstanceInitialPasswordOK) Error() string {
	return fmt.Sprintf("[GET /workload/v1/stacks/{stack_id}/workloads/{workload_id}/instances/{instance_name}/passwords/initial][%d] getWorkloadInstanceInitialPasswordOK  %+v", 200, o.Payload)
}
func (o *GetWorkloadInstanceInitialPasswordOK) GetPayload() *workload_models.V1GetWorkloadInstanceInitialPasswordResponse {
	return o.Payload
}

func (o *GetWorkloadInstanceInitialPasswordOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(workload_models.V1GetWorkloadInstanceInitialPasswordResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkloadInstanceInitialPasswordUnauthorized creates a GetWorkloadInstanceInitialPasswordUnauthorized with default headers values
func NewGetWorkloadInstanceInitialPasswordUnauthorized() *GetWorkloadInstanceInitialPasswordUnauthorized {
	return &GetWorkloadInstanceInitialPasswordUnauthorized{}
}

/* GetWorkloadInstanceInitialPasswordUnauthorized describes a response with status code 401, with default header values.

Returned when an unauthorized request is attempted.
*/
type GetWorkloadInstanceInitialPasswordUnauthorized struct {
	Payload *workload_models.StackpathapiStatus
}

func (o *GetWorkloadInstanceInitialPasswordUnauthorized) Error() string {
	return fmt.Sprintf("[GET /workload/v1/stacks/{stack_id}/workloads/{workload_id}/instances/{instance_name}/passwords/initial][%d] getWorkloadInstanceInitialPasswordUnauthorized  %+v", 401, o.Payload)
}
func (o *GetWorkloadInstanceInitialPasswordUnauthorized) GetPayload() *workload_models.StackpathapiStatus {
	return o.Payload
}

func (o *GetWorkloadInstanceInitialPasswordUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(workload_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkloadInstanceInitialPasswordInternalServerError creates a GetWorkloadInstanceInitialPasswordInternalServerError with default headers values
func NewGetWorkloadInstanceInitialPasswordInternalServerError() *GetWorkloadInstanceInitialPasswordInternalServerError {
	return &GetWorkloadInstanceInitialPasswordInternalServerError{}
}

/* GetWorkloadInstanceInitialPasswordInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type GetWorkloadInstanceInitialPasswordInternalServerError struct {
	Payload *workload_models.StackpathapiStatus
}

func (o *GetWorkloadInstanceInitialPasswordInternalServerError) Error() string {
	return fmt.Sprintf("[GET /workload/v1/stacks/{stack_id}/workloads/{workload_id}/instances/{instance_name}/passwords/initial][%d] getWorkloadInstanceInitialPasswordInternalServerError  %+v", 500, o.Payload)
}
func (o *GetWorkloadInstanceInitialPasswordInternalServerError) GetPayload() *workload_models.StackpathapiStatus {
	return o.Payload
}

func (o *GetWorkloadInstanceInitialPasswordInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(workload_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkloadInstanceInitialPasswordDefault creates a GetWorkloadInstanceInitialPasswordDefault with default headers values
func NewGetWorkloadInstanceInitialPasswordDefault(code int) *GetWorkloadInstanceInitialPasswordDefault {
	return &GetWorkloadInstanceInitialPasswordDefault{
		_statusCode: code,
	}
}

/* GetWorkloadInstanceInitialPasswordDefault describes a response with status code -1, with default header values.

Default error structure.
*/
type GetWorkloadInstanceInitialPasswordDefault struct {
	_statusCode int

	Payload *workload_models.StackpathapiStatus
}

// Code gets the status code for the get workload instance initial password default response
func (o *GetWorkloadInstanceInitialPasswordDefault) Code() int {
	return o._statusCode
}

func (o *GetWorkloadInstanceInitialPasswordDefault) Error() string {
	return fmt.Sprintf("[GET /workload/v1/stacks/{stack_id}/workloads/{workload_id}/instances/{instance_name}/passwords/initial][%d] GetWorkloadInstanceInitialPassword default  %+v", o._statusCode, o.Payload)
}
func (o *GetWorkloadInstanceInitialPasswordDefault) GetPayload() *workload_models.StackpathapiStatus {
	return o.Payload
}

func (o *GetWorkloadInstanceInitialPasswordDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(workload_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/terraform-providers/terraform-provider-stackpath/stackpath/api/workload/workload_models"
)

// RestartInstanceReader is a Reader for the RestartInstance structure.
type RestartInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RestartInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRestartInstanceNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRestartInstanceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRestartInstanceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewRestartInstanceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRestartInstanceNoContent creates a RestartInstanceNoContent with default headers values
func NewRestartInstanceNoContent() *RestartInstanceNoContent {
	return &RestartInstanceNoContent{}
}

/* RestartInstanceNoContent describes a response with status code 204, with default header values.

No content
*/
type RestartInstanceNoContent struct {
}

func (o *RestartInstanceNoContent) Error() string {
	return fmt.Sprintf("[POST /workload/v1/stacks/{stack_id}/workloads/{workload_id}/instances/{instance_name}/power/restart][%d] restartInstanceNoContent ", 204)
}

func (o *RestartInstanceNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRestartInstanceUnauthorized creates a RestartInstanceUnauthorized with default headers values
func NewRestartInstanceUnauthorized() *RestartInstanceUnauthorized {
	return &RestartInstanceUnauthorized{}
}

/* RestartInstanceUnauthorized describes a response with status code 401, with default header values.

Returned when an unauthorized request is attempted.
*/
type RestartInstanceUnauthorized struct {
	Payload *workload_models.StackpathapiStatus
}

func (o *RestartInstanceUnauthorized) Error() string {
	return fmt.Sprintf("[POST /workload/v1/stacks/{stack_id}/workloads/{workload_id}/instances/{instance_name}/power/restart][%d] restartInstanceUnauthorized  %+v", 401, o.Payload)
}
func (o *RestartInstanceUnauthorized) GetPayload() *workload_models.StackpathapiStatus {
	return o.Payload
}

func (o *RestartInstanceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(workload_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestartInstanceInternalServerError creates a RestartInstanceInternalServerError with default headers values
func NewRestartInstanceInternalServerError() *RestartInstanceInternalServerError {
	return &RestartInstanceInternalServerError{}
}

/* RestartInstanceInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type RestartInstanceInternalServerError struct {
	Payload *workload_models.StackpathapiStatus
}

func (o *RestartInstanceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /workload/v1/stacks/{stack_id}/workloads/{workload_id}/instances/{instance_name}/power/restart][%d] restartInstanceInternalServerError  %+v", 500, o.Payload)
}
func (o *RestartInstanceInternalServerError) GetPayload() *workload_models.StackpathapiStatus {
	return o.Payload
}

func (o *RestartInstanceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(workload_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestartInstanceDefault creates a RestartInstanceDefault with default headers values
func NewRestartInstanceDefault(code int) *RestartInstanceDefault {
	return &RestartInstanceDefault{
		_statusCode: code,
	}
}

/* RestartInstanceDefault describes a response with status code -1, with default header values.

Default error structure.
*/
type RestartInstanceDefault struct {
	_statusCode int

	Payload *workload_models.StackpathapiStatus
}

// Code gets the status code for the restart instance default response
func (o *RestartInstanceDefault) Code() int {
	return o._statusCode
}

func (o *RestartInstanceDefault) Error() string {
	return fmt.Sprintf("[POST /workload/v1/stacks/{stack_id}/workloads/{workload_id}/instances/{instance_name}/power/restart][%d] RestartInstance default  %+v", o._statusCode, o.Payload)
}
func (o *RestartInstanceDefault) GetPayload() *workload_models.StackpathapiStatus {
	return o.Payload
}

func (o *RestartInstanceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(workload_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

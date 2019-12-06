// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/terraform-providers/terraform-provider-stackpath/stackpath/internal/models"
)

// GetWorkloadInstancesReader is a Reader for the GetWorkloadInstances structure.
type GetWorkloadInstancesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkloadInstancesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetWorkloadInstancesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetWorkloadInstancesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetWorkloadInstancesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetWorkloadInstancesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWorkloadInstancesOK creates a GetWorkloadInstancesOK with default headers values
func NewGetWorkloadInstancesOK() *GetWorkloadInstancesOK {
	return &GetWorkloadInstancesOK{}
}

/*GetWorkloadInstancesOK handles this case with default header values.

GetWorkloadInstancesOK get workload instances Ok
*/
type GetWorkloadInstancesOK struct {
	Payload *models.V1GetWorkloadInstancesResponse
}

func (o *GetWorkloadInstancesOK) Error() string {
	return fmt.Sprintf("[GET /workload/v1/stacks/{stack_id}/workloads/{workload_id}/instances][%d] getWorkloadInstancesOk  %+v", 200, o.Payload)
}

func (o *GetWorkloadInstancesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1GetWorkloadInstancesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkloadInstancesUnauthorized creates a GetWorkloadInstancesUnauthorized with default headers values
func NewGetWorkloadInstancesUnauthorized() *GetWorkloadInstancesUnauthorized {
	return &GetWorkloadInstancesUnauthorized{}
}

/*GetWorkloadInstancesUnauthorized handles this case with default header values.

Returned when an unauthorized request is attempted.
*/
type GetWorkloadInstancesUnauthorized struct {
	Payload *models.StackpathapiStatus
}

func (o *GetWorkloadInstancesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /workload/v1/stacks/{stack_id}/workloads/{workload_id}/instances][%d] getWorkloadInstancesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkloadInstancesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkloadInstancesInternalServerError creates a GetWorkloadInstancesInternalServerError with default headers values
func NewGetWorkloadInstancesInternalServerError() *GetWorkloadInstancesInternalServerError {
	return &GetWorkloadInstancesInternalServerError{}
}

/*GetWorkloadInstancesInternalServerError handles this case with default header values.

Internal server error.
*/
type GetWorkloadInstancesInternalServerError struct {
	Payload *models.StackpathapiStatus
}

func (o *GetWorkloadInstancesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /workload/v1/stacks/{stack_id}/workloads/{workload_id}/instances][%d] getWorkloadInstancesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkloadInstancesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkloadInstancesDefault creates a GetWorkloadInstancesDefault with default headers values
func NewGetWorkloadInstancesDefault(code int) *GetWorkloadInstancesDefault {
	return &GetWorkloadInstancesDefault{
		_statusCode: code,
	}
}

/*GetWorkloadInstancesDefault handles this case with default header values.

Default error structure.
*/
type GetWorkloadInstancesDefault struct {
	_statusCode int

	Payload *models.StackpathapiStatus
}

// Code gets the status code for the get workload instances default response
func (o *GetWorkloadInstancesDefault) Code() int {
	return o._statusCode
}

func (o *GetWorkloadInstancesDefault) Error() string {
	return fmt.Sprintf("[GET /workload/v1/stacks/{stack_id}/workloads/{workload_id}/instances][%d] GetWorkloadInstances default  %+v", o._statusCode, o.Payload)
}

func (o *GetWorkloadInstancesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

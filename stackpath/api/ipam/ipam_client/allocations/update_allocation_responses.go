// Code generated by go-swagger; DO NOT EDIT.

package allocations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/stackpath/terraform-provider-stackpath/v2/stackpath/api/ipam/ipam_models"
)

// UpdateAllocationReader is a Reader for the UpdateAllocation structure.
type UpdateAllocationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAllocationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAllocationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateAllocationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateAllocationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateAllocationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateAllocationOK creates a UpdateAllocationOK with default headers values
func NewUpdateAllocationOK() *UpdateAllocationOK {
	return &UpdateAllocationOK{}
}

/*
UpdateAllocationOK describes a response with status code 200, with default header values.

UpdateAllocationOK update allocation o k
*/
type UpdateAllocationOK struct {
	Payload *ipam_models.V1Operation
}

func (o *UpdateAllocationOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/v1/stacks/{stack_id}/allocations/{allocation_slug}][%d] updateAllocationOK  %+v", 200, o.Payload)
}
func (o *UpdateAllocationOK) GetPayload() *ipam_models.V1Operation {
	return o.Payload
}

func (o *UpdateAllocationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.V1Operation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAllocationUnauthorized creates a UpdateAllocationUnauthorized with default headers values
func NewUpdateAllocationUnauthorized() *UpdateAllocationUnauthorized {
	return &UpdateAllocationUnauthorized{}
}

/*
UpdateAllocationUnauthorized describes a response with status code 401, with default header values.

Returned when an unauthorized request is attempted.
*/
type UpdateAllocationUnauthorized struct {
	Payload *ipam_models.APIStatus
}

func (o *UpdateAllocationUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /ipam/v1/stacks/{stack_id}/allocations/{allocation_slug}][%d] updateAllocationUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateAllocationUnauthorized) GetPayload() *ipam_models.APIStatus {
	return o.Payload
}

func (o *UpdateAllocationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAllocationInternalServerError creates a UpdateAllocationInternalServerError with default headers values
func NewUpdateAllocationInternalServerError() *UpdateAllocationInternalServerError {
	return &UpdateAllocationInternalServerError{}
}

/*
UpdateAllocationInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type UpdateAllocationInternalServerError struct {
	Payload *ipam_models.APIStatus
}

func (o *UpdateAllocationInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /ipam/v1/stacks/{stack_id}/allocations/{allocation_slug}][%d] updateAllocationInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateAllocationInternalServerError) GetPayload() *ipam_models.APIStatus {
	return o.Payload
}

func (o *UpdateAllocationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAllocationDefault creates a UpdateAllocationDefault with default headers values
func NewUpdateAllocationDefault(code int) *UpdateAllocationDefault {
	return &UpdateAllocationDefault{
		_statusCode: code,
	}
}

/*
UpdateAllocationDefault describes a response with status code -1, with default header values.

Default error structure.
*/
type UpdateAllocationDefault struct {
	_statusCode int

	Payload *ipam_models.APIStatus
}

// Code gets the status code for the update allocation default response
func (o *UpdateAllocationDefault) Code() int {
	return o._statusCode
}

func (o *UpdateAllocationDefault) Error() string {
	return fmt.Sprintf("[PATCH /ipam/v1/stacks/{stack_id}/allocations/{allocation_slug}][%d] UpdateAllocation default  %+v", o._statusCode, o.Payload)
}
func (o *UpdateAllocationDefault) GetPayload() *ipam_models.APIStatus {
	return o.Payload
}

func (o *UpdateAllocationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
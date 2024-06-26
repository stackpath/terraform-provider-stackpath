// Code generated by go-swagger; DO NOT EDIT.

package virtual_private_cloud

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/stackpath/terraform-provider-stackpath/v2/stackpath/api/ipam/ipam_models"
)

// UpdateRouteReader is a Reader for the UpdateRoute structure.
type UpdateRouteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRouteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRouteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateRouteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateRouteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateRouteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateRouteOK creates a UpdateRouteOK with default headers values
func NewUpdateRouteOK() *UpdateRouteOK {
	return &UpdateRouteOK{}
}

/* UpdateRouteOK describes a response with status code 200, with default header values.

UpdateRouteOK update route o k
*/
type UpdateRouteOK struct {
	Payload *ipam_models.NetworkUpdateRouteResponse
}

func (o *UpdateRouteOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/v1alpha/stacks/{stack_id}/networks/{network_id}/routes/{route_id}][%d] updateRouteOK  %+v", 200, o.Payload)
}
func (o *UpdateRouteOK) GetPayload() *ipam_models.NetworkUpdateRouteResponse {
	return o.Payload
}

func (o *UpdateRouteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.NetworkUpdateRouteResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRouteUnauthorized creates a UpdateRouteUnauthorized with default headers values
func NewUpdateRouteUnauthorized() *UpdateRouteUnauthorized {
	return &UpdateRouteUnauthorized{}
}

/* UpdateRouteUnauthorized describes a response with status code 401, with default header values.

Returned when an unauthorized request is attempted.
*/
type UpdateRouteUnauthorized struct {
	Payload *ipam_models.APIStatus
}

func (o *UpdateRouteUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /ipam/v1alpha/stacks/{stack_id}/networks/{network_id}/routes/{route_id}][%d] updateRouteUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateRouteUnauthorized) GetPayload() *ipam_models.APIStatus {
	return o.Payload
}

func (o *UpdateRouteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRouteInternalServerError creates a UpdateRouteInternalServerError with default headers values
func NewUpdateRouteInternalServerError() *UpdateRouteInternalServerError {
	return &UpdateRouteInternalServerError{}
}

/* UpdateRouteInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type UpdateRouteInternalServerError struct {
	Payload *ipam_models.APIStatus
}

func (o *UpdateRouteInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /ipam/v1alpha/stacks/{stack_id}/networks/{network_id}/routes/{route_id}][%d] updateRouteInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateRouteInternalServerError) GetPayload() *ipam_models.APIStatus {
	return o.Payload
}

func (o *UpdateRouteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRouteDefault creates a UpdateRouteDefault with default headers values
func NewUpdateRouteDefault(code int) *UpdateRouteDefault {
	return &UpdateRouteDefault{
		_statusCode: code,
	}
}

/* UpdateRouteDefault describes a response with status code -1, with default header values.

Default error structure.
*/
type UpdateRouteDefault struct {
	_statusCode int

	Payload *ipam_models.APIStatus
}

// Code gets the status code for the update route default response
func (o *UpdateRouteDefault) Code() int {
	return o._statusCode
}

func (o *UpdateRouteDefault) Error() string {
	return fmt.Sprintf("[PUT /ipam/v1alpha/stacks/{stack_id}/networks/{network_id}/routes/{route_id}][%d] UpdateRoute default  %+v", o._statusCode, o.Payload)
}
func (o *UpdateRouteDefault) GetPayload() *ipam_models.APIStatus {
	return o.Payload
}

func (o *UpdateRouteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

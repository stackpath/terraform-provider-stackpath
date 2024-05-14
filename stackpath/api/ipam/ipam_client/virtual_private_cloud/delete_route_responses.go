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

// DeleteRouteReader is a Reader for the DeleteRoute structure.
type DeleteRouteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRouteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteRouteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteRouteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteRouteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteRouteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteRouteNoContent creates a DeleteRouteNoContent with default headers values
func NewDeleteRouteNoContent() *DeleteRouteNoContent {
	return &DeleteRouteNoContent{}
}

/*
DeleteRouteNoContent describes a response with status code 204, with default header values.

No content
*/
type DeleteRouteNoContent struct {
}

func (o *DeleteRouteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/v1alpha/stacks/{stack_id}/networks/{network_id}/routes/{route_id}][%d] deleteRouteNoContent ", 204)
}

func (o *DeleteRouteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRouteUnauthorized creates a DeleteRouteUnauthorized with default headers values
func NewDeleteRouteUnauthorized() *DeleteRouteUnauthorized {
	return &DeleteRouteUnauthorized{}
}

/*
DeleteRouteUnauthorized describes a response with status code 401, with default header values.

Returned when an unauthorized request is attempted.
*/
type DeleteRouteUnauthorized struct {
	Payload *ipam_models.APIStatus
}

func (o *DeleteRouteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /ipam/v1alpha/stacks/{stack_id}/networks/{network_id}/routes/{route_id}][%d] deleteRouteUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteRouteUnauthorized) GetPayload() *ipam_models.APIStatus {
	return o.Payload
}

func (o *DeleteRouteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRouteInternalServerError creates a DeleteRouteInternalServerError with default headers values
func NewDeleteRouteInternalServerError() *DeleteRouteInternalServerError {
	return &DeleteRouteInternalServerError{}
}

/*
DeleteRouteInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type DeleteRouteInternalServerError struct {
	Payload *ipam_models.APIStatus
}

func (o *DeleteRouteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /ipam/v1alpha/stacks/{stack_id}/networks/{network_id}/routes/{route_id}][%d] deleteRouteInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteRouteInternalServerError) GetPayload() *ipam_models.APIStatus {
	return o.Payload
}

func (o *DeleteRouteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRouteDefault creates a DeleteRouteDefault with default headers values
func NewDeleteRouteDefault(code int) *DeleteRouteDefault {
	return &DeleteRouteDefault{
		_statusCode: code,
	}
}

/*
DeleteRouteDefault describes a response with status code -1, with default header values.

Default error structure.
*/
type DeleteRouteDefault struct {
	_statusCode int

	Payload *ipam_models.APIStatus
}

// Code gets the status code for the delete route default response
func (o *DeleteRouteDefault) Code() int {
	return o._statusCode
}

func (o *DeleteRouteDefault) Error() string {
	return fmt.Sprintf("[DELETE /ipam/v1alpha/stacks/{stack_id}/networks/{network_id}/routes/{route_id}][%d] DeleteRoute default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteRouteDefault) GetPayload() *ipam_models.APIStatus {
	return o.Payload
}

func (o *DeleteRouteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

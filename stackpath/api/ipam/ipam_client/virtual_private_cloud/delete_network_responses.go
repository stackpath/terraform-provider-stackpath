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

// DeleteNetworkReader is a Reader for the DeleteNetwork structure.
type DeleteNetworkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNetworkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteNetworkNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteNetworkUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteNetworkInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteNetworkDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteNetworkNoContent creates a DeleteNetworkNoContent with default headers values
func NewDeleteNetworkNoContent() *DeleteNetworkNoContent {
	return &DeleteNetworkNoContent{}
}

/* DeleteNetworkNoContent describes a response with status code 204, with default header values.

No content
*/
type DeleteNetworkNoContent struct {
}

func (o *DeleteNetworkNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/v1/stacks/{stack_id}/networks/{network_id}][%d] deleteNetworkNoContent ", 204)
}

func (o *DeleteNetworkNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNetworkUnauthorized creates a DeleteNetworkUnauthorized with default headers values
func NewDeleteNetworkUnauthorized() *DeleteNetworkUnauthorized {
	return &DeleteNetworkUnauthorized{}
}

/* DeleteNetworkUnauthorized describes a response with status code 401, with default header values.

Returned when an unauthorized request is attempted.
*/
type DeleteNetworkUnauthorized struct {
	Payload *ipam_models.APIStatus
}

func (o *DeleteNetworkUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /ipam/v1/stacks/{stack_id}/networks/{network_id}][%d] deleteNetworkUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteNetworkUnauthorized) GetPayload() *ipam_models.APIStatus {
	return o.Payload
}

func (o *DeleteNetworkUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNetworkInternalServerError creates a DeleteNetworkInternalServerError with default headers values
func NewDeleteNetworkInternalServerError() *DeleteNetworkInternalServerError {
	return &DeleteNetworkInternalServerError{}
}

/* DeleteNetworkInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type DeleteNetworkInternalServerError struct {
	Payload *ipam_models.APIStatus
}

func (o *DeleteNetworkInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /ipam/v1/stacks/{stack_id}/networks/{network_id}][%d] deleteNetworkInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteNetworkInternalServerError) GetPayload() *ipam_models.APIStatus {
	return o.Payload
}

func (o *DeleteNetworkInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNetworkDefault creates a DeleteNetworkDefault with default headers values
func NewDeleteNetworkDefault(code int) *DeleteNetworkDefault {
	return &DeleteNetworkDefault{
		_statusCode: code,
	}
}

/* DeleteNetworkDefault describes a response with status code -1, with default header values.

Default error structure.
*/
type DeleteNetworkDefault struct {
	_statusCode int

	Payload *ipam_models.APIStatus
}

// Code gets the status code for the delete network default response
func (o *DeleteNetworkDefault) Code() int {
	return o._statusCode
}

func (o *DeleteNetworkDefault) Error() string {
	return fmt.Sprintf("[DELETE /ipam/v1/stacks/{stack_id}/networks/{network_id}][%d] DeleteNetwork default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteNetworkDefault) GetPayload() *ipam_models.APIStatus {
	return o.Payload
}

func (o *DeleteNetworkDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

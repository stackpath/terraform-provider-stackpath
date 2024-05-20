// Code generated by go-swagger; DO NOT EDIT.

package network_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/stackpath/terraform-provider-stackpath/v2/stackpath/api/ipam/ipam_models"
)

// DeleteNetworkPolicyReader is a Reader for the DeleteNetworkPolicy structure.
type DeleteNetworkPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNetworkPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteNetworkPolicyNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteNetworkPolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteNetworkPolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteNetworkPolicyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteNetworkPolicyNoContent creates a DeleteNetworkPolicyNoContent with default headers values
func NewDeleteNetworkPolicyNoContent() *DeleteNetworkPolicyNoContent {
	return &DeleteNetworkPolicyNoContent{}
}

/* DeleteNetworkPolicyNoContent describes a response with status code 204, with default header values.

No content
*/
type DeleteNetworkPolicyNoContent struct {
}

func (o *DeleteNetworkPolicyNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/v1/stacks/{stack_id}/network_policies/{network_policy_id}][%d] deleteNetworkPolicyNoContent ", 204)
}

func (o *DeleteNetworkPolicyNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNetworkPolicyUnauthorized creates a DeleteNetworkPolicyUnauthorized with default headers values
func NewDeleteNetworkPolicyUnauthorized() *DeleteNetworkPolicyUnauthorized {
	return &DeleteNetworkPolicyUnauthorized{}
}

/* DeleteNetworkPolicyUnauthorized describes a response with status code 401, with default header values.

Returned when an unauthorized request is attempted.
*/
type DeleteNetworkPolicyUnauthorized struct {
	Payload *ipam_models.APIStatus
}

func (o *DeleteNetworkPolicyUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /ipam/v1/stacks/{stack_id}/network_policies/{network_policy_id}][%d] deleteNetworkPolicyUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteNetworkPolicyUnauthorized) GetPayload() *ipam_models.APIStatus {
	return o.Payload
}

func (o *DeleteNetworkPolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNetworkPolicyInternalServerError creates a DeleteNetworkPolicyInternalServerError with default headers values
func NewDeleteNetworkPolicyInternalServerError() *DeleteNetworkPolicyInternalServerError {
	return &DeleteNetworkPolicyInternalServerError{}
}

/* DeleteNetworkPolicyInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type DeleteNetworkPolicyInternalServerError struct {
	Payload *ipam_models.APIStatus
}

func (o *DeleteNetworkPolicyInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /ipam/v1/stacks/{stack_id}/network_policies/{network_policy_id}][%d] deleteNetworkPolicyInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteNetworkPolicyInternalServerError) GetPayload() *ipam_models.APIStatus {
	return o.Payload
}

func (o *DeleteNetworkPolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNetworkPolicyDefault creates a DeleteNetworkPolicyDefault with default headers values
func NewDeleteNetworkPolicyDefault(code int) *DeleteNetworkPolicyDefault {
	return &DeleteNetworkPolicyDefault{
		_statusCode: code,
	}
}

/* DeleteNetworkPolicyDefault describes a response with status code -1, with default header values.

Default error structure.
*/
type DeleteNetworkPolicyDefault struct {
	_statusCode int

	Payload *ipam_models.APIStatus
}

// Code gets the status code for the delete network policy default response
func (o *DeleteNetworkPolicyDefault) Code() int {
	return o._statusCode
}

func (o *DeleteNetworkPolicyDefault) Error() string {
	return fmt.Sprintf("[DELETE /ipam/v1/stacks/{stack_id}/network_policies/{network_policy_id}][%d] DeleteNetworkPolicy default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteNetworkPolicyDefault) GetPayload() *ipam_models.APIStatus {
	return o.Payload
}

func (o *DeleteNetworkPolicyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

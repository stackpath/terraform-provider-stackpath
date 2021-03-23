// Code generated by go-swagger; DO NOT EDIT.

package network_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/terraform-providers/terraform-provider-stackpath/stackpath/api/ipam/ipam_models"
)

// GetNetworkPolicyReader is a Reader for the GetNetworkPolicy structure.
type GetNetworkPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetNetworkPolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetNetworkPolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetNetworkPolicyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNetworkPolicyOK creates a GetNetworkPolicyOK with default headers values
func NewGetNetworkPolicyOK() *GetNetworkPolicyOK {
	return &GetNetworkPolicyOK{}
}

/* GetNetworkPolicyOK describes a response with status code 200, with default header values.

GetNetworkPolicyOK get network policy o k
*/
type GetNetworkPolicyOK struct {
	Payload *ipam_models.V1GetNetworkPolicyResponse
}

func (o *GetNetworkPolicyOK) Error() string {
	return fmt.Sprintf("[GET /ipam/v1/stacks/{stack_id}/network_policies/{network_policy_id}][%d] getNetworkPolicyOK  %+v", 200, o.Payload)
}
func (o *GetNetworkPolicyOK) GetPayload() *ipam_models.V1GetNetworkPolicyResponse {
	return o.Payload
}

func (o *GetNetworkPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.V1GetNetworkPolicyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworkPolicyUnauthorized creates a GetNetworkPolicyUnauthorized with default headers values
func NewGetNetworkPolicyUnauthorized() *GetNetworkPolicyUnauthorized {
	return &GetNetworkPolicyUnauthorized{}
}

/* GetNetworkPolicyUnauthorized describes a response with status code 401, with default header values.

Returned when an unauthorized request is attempted.
*/
type GetNetworkPolicyUnauthorized struct {
	Payload *ipam_models.APIStatus
}

func (o *GetNetworkPolicyUnauthorized) Error() string {
	return fmt.Sprintf("[GET /ipam/v1/stacks/{stack_id}/network_policies/{network_policy_id}][%d] getNetworkPolicyUnauthorized  %+v", 401, o.Payload)
}
func (o *GetNetworkPolicyUnauthorized) GetPayload() *ipam_models.APIStatus {
	return o.Payload
}

func (o *GetNetworkPolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworkPolicyInternalServerError creates a GetNetworkPolicyInternalServerError with default headers values
func NewGetNetworkPolicyInternalServerError() *GetNetworkPolicyInternalServerError {
	return &GetNetworkPolicyInternalServerError{}
}

/* GetNetworkPolicyInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type GetNetworkPolicyInternalServerError struct {
	Payload *ipam_models.APIStatus
}

func (o *GetNetworkPolicyInternalServerError) Error() string {
	return fmt.Sprintf("[GET /ipam/v1/stacks/{stack_id}/network_policies/{network_policy_id}][%d] getNetworkPolicyInternalServerError  %+v", 500, o.Payload)
}
func (o *GetNetworkPolicyInternalServerError) GetPayload() *ipam_models.APIStatus {
	return o.Payload
}

func (o *GetNetworkPolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworkPolicyDefault creates a GetNetworkPolicyDefault with default headers values
func NewGetNetworkPolicyDefault(code int) *GetNetworkPolicyDefault {
	return &GetNetworkPolicyDefault{
		_statusCode: code,
	}
}

/* GetNetworkPolicyDefault describes a response with status code -1, with default header values.

Default error structure.
*/
type GetNetworkPolicyDefault struct {
	_statusCode int

	Payload *ipam_models.APIStatus
}

// Code gets the status code for the get network policy default response
func (o *GetNetworkPolicyDefault) Code() int {
	return o._statusCode
}

func (o *GetNetworkPolicyDefault) Error() string {
	return fmt.Sprintf("[GET /ipam/v1/stacks/{stack_id}/network_policies/{network_policy_id}][%d] GetNetworkPolicy default  %+v", o._statusCode, o.Payload)
}
func (o *GetNetworkPolicyDefault) GetPayload() *ipam_models.APIStatus {
	return o.Payload
}

func (o *GetNetworkPolicyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// GetNetworkPoliciesReader is a Reader for the GetNetworkPolicies structure.
type GetNetworkPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetNetworkPoliciesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetNetworkPoliciesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetNetworkPoliciesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNetworkPoliciesOK creates a GetNetworkPoliciesOK with default headers values
func NewGetNetworkPoliciesOK() *GetNetworkPoliciesOK {
	return &GetNetworkPoliciesOK{}
}

/* GetNetworkPoliciesOK describes a response with status code 200, with default header values.

GetNetworkPoliciesOK get network policies o k
*/
type GetNetworkPoliciesOK struct {
	Payload *ipam_models.V1GetNetworkPoliciesResponse
}

func (o *GetNetworkPoliciesOK) Error() string {
	return fmt.Sprintf("[GET /ipam/v1/stacks/{stack_id}/network_policies][%d] getNetworkPoliciesOK  %+v", 200, o.Payload)
}
func (o *GetNetworkPoliciesOK) GetPayload() *ipam_models.V1GetNetworkPoliciesResponse {
	return o.Payload
}

func (o *GetNetworkPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.V1GetNetworkPoliciesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworkPoliciesUnauthorized creates a GetNetworkPoliciesUnauthorized with default headers values
func NewGetNetworkPoliciesUnauthorized() *GetNetworkPoliciesUnauthorized {
	return &GetNetworkPoliciesUnauthorized{}
}

/* GetNetworkPoliciesUnauthorized describes a response with status code 401, with default header values.

Returned when an unauthorized request is attempted.
*/
type GetNetworkPoliciesUnauthorized struct {
	Payload *ipam_models.APIStatus
}

func (o *GetNetworkPoliciesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /ipam/v1/stacks/{stack_id}/network_policies][%d] getNetworkPoliciesUnauthorized  %+v", 401, o.Payload)
}
func (o *GetNetworkPoliciesUnauthorized) GetPayload() *ipam_models.APIStatus {
	return o.Payload
}

func (o *GetNetworkPoliciesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworkPoliciesInternalServerError creates a GetNetworkPoliciesInternalServerError with default headers values
func NewGetNetworkPoliciesInternalServerError() *GetNetworkPoliciesInternalServerError {
	return &GetNetworkPoliciesInternalServerError{}
}

/* GetNetworkPoliciesInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type GetNetworkPoliciesInternalServerError struct {
	Payload *ipam_models.APIStatus
}

func (o *GetNetworkPoliciesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /ipam/v1/stacks/{stack_id}/network_policies][%d] getNetworkPoliciesInternalServerError  %+v", 500, o.Payload)
}
func (o *GetNetworkPoliciesInternalServerError) GetPayload() *ipam_models.APIStatus {
	return o.Payload
}

func (o *GetNetworkPoliciesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworkPoliciesDefault creates a GetNetworkPoliciesDefault with default headers values
func NewGetNetworkPoliciesDefault(code int) *GetNetworkPoliciesDefault {
	return &GetNetworkPoliciesDefault{
		_statusCode: code,
	}
}

/* GetNetworkPoliciesDefault describes a response with status code -1, with default header values.

Default error structure.
*/
type GetNetworkPoliciesDefault struct {
	_statusCode int

	Payload *ipam_models.APIStatus
}

// Code gets the status code for the get network policies default response
func (o *GetNetworkPoliciesDefault) Code() int {
	return o._statusCode
}

func (o *GetNetworkPoliciesDefault) Error() string {
	return fmt.Sprintf("[GET /ipam/v1/stacks/{stack_id}/network_policies][%d] GetNetworkPolicies default  %+v", o._statusCode, o.Payload)
}
func (o *GetNetworkPoliciesDefault) GetPayload() *ipam_models.APIStatus {
	return o.Payload
}

func (o *GetNetworkPoliciesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

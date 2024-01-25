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

// CreateNetworkSubnetReader is a Reader for the CreateNetworkSubnet structure.
type CreateNetworkSubnetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNetworkSubnetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateNetworkSubnetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateNetworkSubnetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateNetworkSubnetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateNetworkSubnetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateNetworkSubnetOK creates a CreateNetworkSubnetOK with default headers values
func NewCreateNetworkSubnetOK() *CreateNetworkSubnetOK {
	return &CreateNetworkSubnetOK{}
}

/*
	CreateNetworkSubnetOK describes a response with status code 200, with default header values.

CreateNetworkSubnetOK create network subnet o k
*/
type CreateNetworkSubnetOK struct {
	Payload *ipam_models.NetworkCreateNetworkSubnetResponse
}

func (o *CreateNetworkSubnetOK) Error() string {
	return fmt.Sprintf("[POST /ipam/v1alpha/stacks/{stack_id}/networks/{network_id}/subnets][%d] createNetworkSubnetOK  %+v", 200, o.Payload)
}
func (o *CreateNetworkSubnetOK) GetPayload() *ipam_models.NetworkCreateNetworkSubnetResponse {
	return o.Payload
}

func (o *CreateNetworkSubnetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.NetworkCreateNetworkSubnetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNetworkSubnetUnauthorized creates a CreateNetworkSubnetUnauthorized with default headers values
func NewCreateNetworkSubnetUnauthorized() *CreateNetworkSubnetUnauthorized {
	return &CreateNetworkSubnetUnauthorized{}
}

/*
	CreateNetworkSubnetUnauthorized describes a response with status code 401, with default header values.

Returned when an unauthorized request is attempted.
*/
type CreateNetworkSubnetUnauthorized struct {
	Payload *ipam_models.APIStatus
}

func (o *CreateNetworkSubnetUnauthorized) Error() string {
	return fmt.Sprintf("[POST /ipam/v1alpha/stacks/{stack_id}/networks/{network_id}/subnets][%d] createNetworkSubnetUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateNetworkSubnetUnauthorized) GetPayload() *ipam_models.APIStatus {
	return o.Payload
}

func (o *CreateNetworkSubnetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNetworkSubnetInternalServerError creates a CreateNetworkSubnetInternalServerError with default headers values
func NewCreateNetworkSubnetInternalServerError() *CreateNetworkSubnetInternalServerError {
	return &CreateNetworkSubnetInternalServerError{}
}

/*
	CreateNetworkSubnetInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type CreateNetworkSubnetInternalServerError struct {
	Payload *ipam_models.APIStatus
}

func (o *CreateNetworkSubnetInternalServerError) Error() string {
	return fmt.Sprintf("[POST /ipam/v1alpha/stacks/{stack_id}/networks/{network_id}/subnets][%d] createNetworkSubnetInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateNetworkSubnetInternalServerError) GetPayload() *ipam_models.APIStatus {
	return o.Payload
}

func (o *CreateNetworkSubnetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNetworkSubnetDefault creates a CreateNetworkSubnetDefault with default headers values
func NewCreateNetworkSubnetDefault(code int) *CreateNetworkSubnetDefault {
	return &CreateNetworkSubnetDefault{
		_statusCode: code,
	}
}

/*
	CreateNetworkSubnetDefault describes a response with status code -1, with default header values.

Default error structure.
*/
type CreateNetworkSubnetDefault struct {
	_statusCode int

	Payload *ipam_models.APIStatus
}

// Code gets the status code for the create network subnet default response
func (o *CreateNetworkSubnetDefault) Code() int {
	return o._statusCode
}

func (o *CreateNetworkSubnetDefault) Error() string {
	return fmt.Sprintf("[POST /ipam/v1alpha/stacks/{stack_id}/networks/{network_id}/subnets][%d] CreateNetworkSubnet default  %+v", o._statusCode, o.Payload)
}
func (o *CreateNetworkSubnetDefault) GetPayload() *ipam_models.APIStatus {
	return o.Payload
}

func (o *CreateNetworkSubnetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// DeleteAllocationClaimReader is a Reader for the DeleteAllocationClaim structure.
type DeleteAllocationClaimReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAllocationClaimReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAllocationClaimOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteAllocationClaimUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteAllocationClaimInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteAllocationClaimDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteAllocationClaimOK creates a DeleteAllocationClaimOK with default headers values
func NewDeleteAllocationClaimOK() *DeleteAllocationClaimOK {
	return &DeleteAllocationClaimOK{}
}

/*
DeleteAllocationClaimOK describes a response with status code 200, with default header values.

DeleteAllocationClaimOK delete allocation claim o k
*/
type DeleteAllocationClaimOK struct {
	Payload *ipam_models.V1Operation
}

func (o *DeleteAllocationClaimOK) Error() string {
	return fmt.Sprintf("[DELETE /ipam/v1/stacks/{stack_id}/allocation_claims/{allocation_claim_slug}][%d] deleteAllocationClaimOK  %+v", 200, o.Payload)
}
func (o *DeleteAllocationClaimOK) GetPayload() *ipam_models.V1Operation {
	return o.Payload
}

func (o *DeleteAllocationClaimOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.V1Operation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAllocationClaimUnauthorized creates a DeleteAllocationClaimUnauthorized with default headers values
func NewDeleteAllocationClaimUnauthorized() *DeleteAllocationClaimUnauthorized {
	return &DeleteAllocationClaimUnauthorized{}
}

/*
DeleteAllocationClaimUnauthorized describes a response with status code 401, with default header values.

Returned when an unauthorized request is attempted.
*/
type DeleteAllocationClaimUnauthorized struct {
	Payload *ipam_models.APIStatus
}

func (o *DeleteAllocationClaimUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /ipam/v1/stacks/{stack_id}/allocation_claims/{allocation_claim_slug}][%d] deleteAllocationClaimUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteAllocationClaimUnauthorized) GetPayload() *ipam_models.APIStatus {
	return o.Payload
}

func (o *DeleteAllocationClaimUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAllocationClaimInternalServerError creates a DeleteAllocationClaimInternalServerError with default headers values
func NewDeleteAllocationClaimInternalServerError() *DeleteAllocationClaimInternalServerError {
	return &DeleteAllocationClaimInternalServerError{}
}

/*
DeleteAllocationClaimInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type DeleteAllocationClaimInternalServerError struct {
	Payload *ipam_models.APIStatus
}

func (o *DeleteAllocationClaimInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /ipam/v1/stacks/{stack_id}/allocation_claims/{allocation_claim_slug}][%d] deleteAllocationClaimInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteAllocationClaimInternalServerError) GetPayload() *ipam_models.APIStatus {
	return o.Payload
}

func (o *DeleteAllocationClaimInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAllocationClaimDefault creates a DeleteAllocationClaimDefault with default headers values
func NewDeleteAllocationClaimDefault(code int) *DeleteAllocationClaimDefault {
	return &DeleteAllocationClaimDefault{
		_statusCode: code,
	}
}

/*
DeleteAllocationClaimDefault describes a response with status code -1, with default header values.

Default error structure.
*/
type DeleteAllocationClaimDefault struct {
	_statusCode int

	Payload *ipam_models.APIStatus
}

// Code gets the status code for the delete allocation claim default response
func (o *DeleteAllocationClaimDefault) Code() int {
	return o._statusCode
}

func (o *DeleteAllocationClaimDefault) Error() string {
	return fmt.Sprintf("[DELETE /ipam/v1/stacks/{stack_id}/allocation_claims/{allocation_claim_slug}][%d] DeleteAllocationClaim default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteAllocationClaimDefault) GetPayload() *ipam_models.APIStatus {
	return o.Payload
}

func (o *DeleteAllocationClaimDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

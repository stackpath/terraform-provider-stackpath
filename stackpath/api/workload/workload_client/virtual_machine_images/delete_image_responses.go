// Code generated by go-swagger; DO NOT EDIT.

package virtual_machine_images

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/stackpath/terraform-provider-stackpath/v2/stackpath/api/workload/workload_models"
)

// DeleteImageReader is a Reader for the DeleteImage structure.
type DeleteImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteImageNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteImageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteImageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteImageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteImageNoContent creates a DeleteImageNoContent with default headers values
func NewDeleteImageNoContent() *DeleteImageNoContent {
	return &DeleteImageNoContent{}
}

/*
DeleteImageNoContent describes a response with status code 204, with default header values.

No content
*/
type DeleteImageNoContent struct {
}

// IsSuccess returns true when this delete image no content response has a 2xx status code
func (o *DeleteImageNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete image no content response has a 3xx status code
func (o *DeleteImageNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete image no content response has a 4xx status code
func (o *DeleteImageNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete image no content response has a 5xx status code
func (o *DeleteImageNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete image no content response a status code equal to that given
func (o *DeleteImageNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete image no content response
func (o *DeleteImageNoContent) Code() int {
	return 204
}

func (o *DeleteImageNoContent) Error() string {
	return fmt.Sprintf("[DELETE /workload/v1/stacks/{stack_id}/images/{image_family}/{image_tag}][%d] deleteImageNoContent ", 204)
}

func (o *DeleteImageNoContent) String() string {
	return fmt.Sprintf("[DELETE /workload/v1/stacks/{stack_id}/images/{image_family}/{image_tag}][%d] deleteImageNoContent ", 204)
}

func (o *DeleteImageNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteImageUnauthorized creates a DeleteImageUnauthorized with default headers values
func NewDeleteImageUnauthorized() *DeleteImageUnauthorized {
	return &DeleteImageUnauthorized{}
}

/*
DeleteImageUnauthorized describes a response with status code 401, with default header values.

Returned when an unauthorized request is attempted.
*/
type DeleteImageUnauthorized struct {
	Payload *workload_models.StackpathapiStatus
}

// IsSuccess returns true when this delete image unauthorized response has a 2xx status code
func (o *DeleteImageUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete image unauthorized response has a 3xx status code
func (o *DeleteImageUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete image unauthorized response has a 4xx status code
func (o *DeleteImageUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete image unauthorized response has a 5xx status code
func (o *DeleteImageUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete image unauthorized response a status code equal to that given
func (o *DeleteImageUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete image unauthorized response
func (o *DeleteImageUnauthorized) Code() int {
	return 401
}

func (o *DeleteImageUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /workload/v1/stacks/{stack_id}/images/{image_family}/{image_tag}][%d] deleteImageUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteImageUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /workload/v1/stacks/{stack_id}/images/{image_family}/{image_tag}][%d] deleteImageUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteImageUnauthorized) GetPayload() *workload_models.StackpathapiStatus {
	return o.Payload
}

func (o *DeleteImageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(workload_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteImageInternalServerError creates a DeleteImageInternalServerError with default headers values
func NewDeleteImageInternalServerError() *DeleteImageInternalServerError {
	return &DeleteImageInternalServerError{}
}

/*
DeleteImageInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type DeleteImageInternalServerError struct {
	Payload *workload_models.StackpathapiStatus
}

// IsSuccess returns true when this delete image internal server error response has a 2xx status code
func (o *DeleteImageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete image internal server error response has a 3xx status code
func (o *DeleteImageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete image internal server error response has a 4xx status code
func (o *DeleteImageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete image internal server error response has a 5xx status code
func (o *DeleteImageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete image internal server error response a status code equal to that given
func (o *DeleteImageInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete image internal server error response
func (o *DeleteImageInternalServerError) Code() int {
	return 500
}

func (o *DeleteImageInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /workload/v1/stacks/{stack_id}/images/{image_family}/{image_tag}][%d] deleteImageInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteImageInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /workload/v1/stacks/{stack_id}/images/{image_family}/{image_tag}][%d] deleteImageInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteImageInternalServerError) GetPayload() *workload_models.StackpathapiStatus {
	return o.Payload
}

func (o *DeleteImageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(workload_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteImageDefault creates a DeleteImageDefault with default headers values
func NewDeleteImageDefault(code int) *DeleteImageDefault {
	return &DeleteImageDefault{
		_statusCode: code,
	}
}

/*
DeleteImageDefault describes a response with status code -1, with default header values.

Default error structure.
*/
type DeleteImageDefault struct {
	_statusCode int

	Payload *workload_models.StackpathapiStatus
}

// IsSuccess returns true when this delete image default response has a 2xx status code
func (o *DeleteImageDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete image default response has a 3xx status code
func (o *DeleteImageDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete image default response has a 4xx status code
func (o *DeleteImageDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete image default response has a 5xx status code
func (o *DeleteImageDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete image default response a status code equal to that given
func (o *DeleteImageDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete image default response
func (o *DeleteImageDefault) Code() int {
	return o._statusCode
}

func (o *DeleteImageDefault) Error() string {
	return fmt.Sprintf("[DELETE /workload/v1/stacks/{stack_id}/images/{image_family}/{image_tag}][%d] DeleteImage default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteImageDefault) String() string {
	return fmt.Sprintf("[DELETE /workload/v1/stacks/{stack_id}/images/{image_family}/{image_tag}][%d] DeleteImage default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteImageDefault) GetPayload() *workload_models.StackpathapiStatus {
	return o.Payload
}

func (o *DeleteImageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(workload_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
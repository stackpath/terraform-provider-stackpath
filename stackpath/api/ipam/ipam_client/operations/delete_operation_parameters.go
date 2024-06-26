// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDeleteOperationParams creates a new DeleteOperationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteOperationParams() *DeleteOperationParams {
	return &DeleteOperationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOperationParamsWithTimeout creates a new DeleteOperationParams object
// with the ability to set a timeout on a request.
func NewDeleteOperationParamsWithTimeout(timeout time.Duration) *DeleteOperationParams {
	return &DeleteOperationParams{
		timeout: timeout,
	}
}

// NewDeleteOperationParamsWithContext creates a new DeleteOperationParams object
// with the ability to set a context for a request.
func NewDeleteOperationParamsWithContext(ctx context.Context) *DeleteOperationParams {
	return &DeleteOperationParams{
		Context: ctx,
	}
}

// NewDeleteOperationParamsWithHTTPClient creates a new DeleteOperationParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteOperationParamsWithHTTPClient(client *http.Client) *DeleteOperationParams {
	return &DeleteOperationParams{
		HTTPClient: client,
	}
}

/*
DeleteOperationParams contains all the parameters to send to the API endpoint

	for the delete operation operation.

	Typically these are written to a http.Request.
*/
type DeleteOperationParams struct {

	// OperationName.
	OperationName string

	/* StackID.

	   A stack ID or slug
	*/
	StackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete operation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOperationParams) WithDefaults() *DeleteOperationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete operation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOperationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete operation params
func (o *DeleteOperationParams) WithTimeout(timeout time.Duration) *DeleteOperationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete operation params
func (o *DeleteOperationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete operation params
func (o *DeleteOperationParams) WithContext(ctx context.Context) *DeleteOperationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete operation params
func (o *DeleteOperationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete operation params
func (o *DeleteOperationParams) WithHTTPClient(client *http.Client) *DeleteOperationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete operation params
func (o *DeleteOperationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOperationName adds the operationName to the delete operation params
func (o *DeleteOperationParams) WithOperationName(operationName string) *DeleteOperationParams {
	o.SetOperationName(operationName)
	return o
}

// SetOperationName adds the operationName to the delete operation params
func (o *DeleteOperationParams) SetOperationName(operationName string) {
	o.OperationName = operationName
}

// WithStackID adds the stackID to the delete operation params
func (o *DeleteOperationParams) WithStackID(stackID string) *DeleteOperationParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the delete operation params
func (o *DeleteOperationParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOperationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param operation_name
	if err := r.SetPathParam("operation_name", o.OperationName); err != nil {
		return err
	}

	// path param stack_id
	if err := r.SetPathParam("stack_id", o.StackID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

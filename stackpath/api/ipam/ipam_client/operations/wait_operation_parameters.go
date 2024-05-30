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

// NewWaitOperationParams creates a new WaitOperationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaitOperationParams() *WaitOperationParams {
	return &WaitOperationParams{
		requestTimeout: cr.DefaultTimeout,
	}
}

// NewWaitOperationParamsWithTimeout creates a new WaitOperationParams object
// with the ability to set a timeout on a request.
func NewWaitOperationParamsWithTimeout(timeout time.Duration) *WaitOperationParams {
	return &WaitOperationParams{
		requestTimeout: timeout,
	}
}

// NewWaitOperationParamsWithContext creates a new WaitOperationParams object
// with the ability to set a context for a request.
func NewWaitOperationParamsWithContext(ctx context.Context) *WaitOperationParams {
	return &WaitOperationParams{
		Context: ctx,
	}
}

// NewWaitOperationParamsWithHTTPClient creates a new WaitOperationParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaitOperationParamsWithHTTPClient(client *http.Client) *WaitOperationParams {
	return &WaitOperationParams{
		HTTPClient: client,
	}
}

/*
WaitOperationParams contains all the parameters to send to the API endpoint

	for the wait operation operation.

	Typically these are written to a http.Request.
*/
type WaitOperationParams struct {

	// OperationName.
	OperationName string

	/* StackID.

	   A stack ID or slug
	*/
	StackID string

	/* Timeout.

	   Time to wait for an operation. Max value is limited to 30s.
	*/
	Timeout *string

	requestTimeout time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithDefaults hydrates default values in the wait operation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaitOperationParams) WithDefaults() *WaitOperationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the wait operation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaitOperationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithRequestTimeout adds the timeout to the wait operation params
func (o *WaitOperationParams) WithRequestTimeout(timeout time.Duration) *WaitOperationParams {
	o.SetRequestTimeout(timeout)
	return o
}

// SetRequestTimeout adds the timeout to the wait operation params
func (o *WaitOperationParams) SetRequestTimeout(timeout time.Duration) {
	o.requestTimeout = timeout
}

// WithContext adds the context to the wait operation params
func (o *WaitOperationParams) WithContext(ctx context.Context) *WaitOperationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the wait operation params
func (o *WaitOperationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the wait operation params
func (o *WaitOperationParams) WithHTTPClient(client *http.Client) *WaitOperationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the wait operation params
func (o *WaitOperationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOperationName adds the operationName to the wait operation params
func (o *WaitOperationParams) WithOperationName(operationName string) *WaitOperationParams {
	o.SetOperationName(operationName)
	return o
}

// SetOperationName adds the operationName to the wait operation params
func (o *WaitOperationParams) SetOperationName(operationName string) {
	o.OperationName = operationName
}

// WithStackID adds the stackID to the wait operation params
func (o *WaitOperationParams) WithStackID(stackID string) *WaitOperationParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the wait operation params
func (o *WaitOperationParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WithTimeout adds the timeout to the wait operation params
func (o *WaitOperationParams) WithTimeout(timeout *string) *WaitOperationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the wait operation params
func (o *WaitOperationParams) SetTimeout(timeout *string) {
	o.Timeout = timeout
}

// WriteToRequest writes these params to a swagger request
func (o *WaitOperationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.requestTimeout); err != nil {
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

	if o.Timeout != nil {

		// query param timeout
		var qrTimeout string

		if o.Timeout != nil {
			qrTimeout = *o.Timeout
		}
		qTimeout := qrTimeout
		if qTimeout != "" {

			if err := r.SetQueryParam("timeout", qTimeout); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

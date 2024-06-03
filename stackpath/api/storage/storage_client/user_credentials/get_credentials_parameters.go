// Code generated by go-swagger; DO NOT EDIT.

package user_credentials

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

// NewGetCredentialsParams creates a new GetCredentialsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCredentialsParams() *GetCredentialsParams {
	return &GetCredentialsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCredentialsParamsWithTimeout creates a new GetCredentialsParams object
// with the ability to set a timeout on a request.
func NewGetCredentialsParamsWithTimeout(timeout time.Duration) *GetCredentialsParams {
	return &GetCredentialsParams{
		timeout: timeout,
	}
}

// NewGetCredentialsParamsWithContext creates a new GetCredentialsParams object
// with the ability to set a context for a request.
func NewGetCredentialsParamsWithContext(ctx context.Context) *GetCredentialsParams {
	return &GetCredentialsParams{
		Context: ctx,
	}
}

// NewGetCredentialsParamsWithHTTPClient creates a new GetCredentialsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCredentialsParamsWithHTTPClient(client *http.Client) *GetCredentialsParams {
	return &GetCredentialsParams{
		HTTPClient: client,
	}
}

/* GetCredentialsParams contains all the parameters to send to the API endpoint
   for the get credentials operation.

   Typically these are written to a http.Request.
*/
type GetCredentialsParams struct {

	/* StackID.

	   The stack's ID for which the user belongs to
	*/
	StackID string

	/* UserID.

	   The user's ID for which the credentials belong to
	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCredentialsParams) WithDefaults() *GetCredentialsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCredentialsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get credentials params
func (o *GetCredentialsParams) WithTimeout(timeout time.Duration) *GetCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get credentials params
func (o *GetCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get credentials params
func (o *GetCredentialsParams) WithContext(ctx context.Context) *GetCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get credentials params
func (o *GetCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get credentials params
func (o *GetCredentialsParams) WithHTTPClient(client *http.Client) *GetCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get credentials params
func (o *GetCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithStackID adds the stackID to the get credentials params
func (o *GetCredentialsParams) WithStackID(stackID string) *GetCredentialsParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the get credentials params
func (o *GetCredentialsParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WithUserID adds the userID to the get credentials params
func (o *GetCredentialsParams) WithUserID(userID string) *GetCredentialsParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get credentials params
func (o *GetCredentialsParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param stack_id
	if err := r.SetPathParam("stack_id", o.StackID); err != nil {
		return err
	}

	// path param user_id
	if err := r.SetPathParam("user_id", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

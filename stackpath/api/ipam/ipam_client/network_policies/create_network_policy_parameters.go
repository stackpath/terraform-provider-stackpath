// Code generated by go-swagger; DO NOT EDIT.

package network_policies

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

	"github.com/stackpath/terraform-provider-stackpath/v2/stackpath/api/ipam/ipam_models"
)

// NewCreateNetworkPolicyParams creates a new CreateNetworkPolicyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateNetworkPolicyParams() *CreateNetworkPolicyParams {
	return &CreateNetworkPolicyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateNetworkPolicyParamsWithTimeout creates a new CreateNetworkPolicyParams object
// with the ability to set a timeout on a request.
func NewCreateNetworkPolicyParamsWithTimeout(timeout time.Duration) *CreateNetworkPolicyParams {
	return &CreateNetworkPolicyParams{
		timeout: timeout,
	}
}

// NewCreateNetworkPolicyParamsWithContext creates a new CreateNetworkPolicyParams object
// with the ability to set a context for a request.
func NewCreateNetworkPolicyParamsWithContext(ctx context.Context) *CreateNetworkPolicyParams {
	return &CreateNetworkPolicyParams{
		Context: ctx,
	}
}

// NewCreateNetworkPolicyParamsWithHTTPClient creates a new CreateNetworkPolicyParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateNetworkPolicyParamsWithHTTPClient(client *http.Client) *CreateNetworkPolicyParams {
	return &CreateNetworkPolicyParams{
		HTTPClient: client,
	}
}

/* CreateNetworkPolicyParams contains all the parameters to send to the API endpoint
   for the create network policy operation.

   Typically these are written to a http.Request.
*/
type CreateNetworkPolicyParams struct {

	// Body.
	Body *ipam_models.V1CreateNetworkPolicyRequest

	/* StackID.

	   A stack ID or slug
	*/
	StackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create network policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNetworkPolicyParams) WithDefaults() *CreateNetworkPolicyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create network policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNetworkPolicyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create network policy params
func (o *CreateNetworkPolicyParams) WithTimeout(timeout time.Duration) *CreateNetworkPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create network policy params
func (o *CreateNetworkPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create network policy params
func (o *CreateNetworkPolicyParams) WithContext(ctx context.Context) *CreateNetworkPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create network policy params
func (o *CreateNetworkPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create network policy params
func (o *CreateNetworkPolicyParams) WithHTTPClient(client *http.Client) *CreateNetworkPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create network policy params
func (o *CreateNetworkPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create network policy params
func (o *CreateNetworkPolicyParams) WithBody(body *ipam_models.V1CreateNetworkPolicyRequest) *CreateNetworkPolicyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create network policy params
func (o *CreateNetworkPolicyParams) SetBody(body *ipam_models.V1CreateNetworkPolicyRequest) {
	o.Body = body
}

// WithStackID adds the stackID to the create network policy params
func (o *CreateNetworkPolicyParams) WithStackID(stackID string) *CreateNetworkPolicyParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the create network policy params
func (o *CreateNetworkPolicyParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateNetworkPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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

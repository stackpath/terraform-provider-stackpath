// Code generated by go-swagger; DO NOT EDIT.

package virtual_private_cloud

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

// NewUpdateNetworkParams creates a new UpdateNetworkParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNetworkParams() *UpdateNetworkParams {
	return &UpdateNetworkParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkParamsWithTimeout creates a new UpdateNetworkParams object
// with the ability to set a timeout on a request.
func NewUpdateNetworkParamsWithTimeout(timeout time.Duration) *UpdateNetworkParams {
	return &UpdateNetworkParams{
		timeout: timeout,
	}
}

// NewUpdateNetworkParamsWithContext creates a new UpdateNetworkParams object
// with the ability to set a context for a request.
func NewUpdateNetworkParamsWithContext(ctx context.Context) *UpdateNetworkParams {
	return &UpdateNetworkParams{
		Context: ctx,
	}
}

// NewUpdateNetworkParamsWithHTTPClient creates a new UpdateNetworkParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNetworkParamsWithHTTPClient(client *http.Client) *UpdateNetworkParams {
	return &UpdateNetworkParams{
		HTTPClient: client,
	}
}

/*
UpdateNetworkParams contains all the parameters to send to the API endpoint

	for the update network operation.

	Typically these are written to a http.Request.
*/
type UpdateNetworkParams struct {

	// Body.
	Body *ipam_models.NetworkUpdateNetworkRequest

	/* NetworkID.

	   A VPC network ID or slug
	*/
	NetworkID string

	/* StackID.

	   A stack ID or slug
	*/
	StackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update network params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkParams) WithDefaults() *UpdateNetworkParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update network params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update network params
func (o *UpdateNetworkParams) WithTimeout(timeout time.Duration) *UpdateNetworkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network params
func (o *UpdateNetworkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network params
func (o *UpdateNetworkParams) WithContext(ctx context.Context) *UpdateNetworkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network params
func (o *UpdateNetworkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network params
func (o *UpdateNetworkParams) WithHTTPClient(client *http.Client) *UpdateNetworkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network params
func (o *UpdateNetworkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update network params
func (o *UpdateNetworkParams) WithBody(body *ipam_models.NetworkUpdateNetworkRequest) *UpdateNetworkParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update network params
func (o *UpdateNetworkParams) SetBody(body *ipam_models.NetworkUpdateNetworkRequest) {
	o.Body = body
}

// WithNetworkID adds the networkID to the update network params
func (o *UpdateNetworkParams) WithNetworkID(networkID string) *UpdateNetworkParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network params
func (o *UpdateNetworkParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithStackID adds the stackID to the update network params
func (o *UpdateNetworkParams) WithStackID(stackID string) *UpdateNetworkParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the update network params
func (o *UpdateNetworkParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
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

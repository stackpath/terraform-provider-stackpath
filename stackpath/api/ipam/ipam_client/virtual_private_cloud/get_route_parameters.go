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
)

// NewGetRouteParams creates a new GetRouteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRouteParams() *GetRouteParams {
	return &GetRouteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRouteParamsWithTimeout creates a new GetRouteParams object
// with the ability to set a timeout on a request.
func NewGetRouteParamsWithTimeout(timeout time.Duration) *GetRouteParams {
	return &GetRouteParams{
		timeout: timeout,
	}
}

// NewGetRouteParamsWithContext creates a new GetRouteParams object
// with the ability to set a context for a request.
func NewGetRouteParamsWithContext(ctx context.Context) *GetRouteParams {
	return &GetRouteParams{
		Context: ctx,
	}
}

// NewGetRouteParamsWithHTTPClient creates a new GetRouteParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRouteParamsWithHTTPClient(client *http.Client) *GetRouteParams {
	return &GetRouteParams{
		HTTPClient: client,
	}
}

/* GetRouteParams contains all the parameters to send to the API endpoint
   for the get route operation.

   Typically these are written to a http.Request.
*/
type GetRouteParams struct {

	/* NetworkID.

	   A VPC network ID or slug
	*/
	NetworkID string

	/* RouteID.

	   A VPC route ID
	*/
	RouteID string

	/* StackID.

	   A stack ID or slug
	*/
	StackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get route params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRouteParams) WithDefaults() *GetRouteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get route params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRouteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get route params
func (o *GetRouteParams) WithTimeout(timeout time.Duration) *GetRouteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get route params
func (o *GetRouteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get route params
func (o *GetRouteParams) WithContext(ctx context.Context) *GetRouteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get route params
func (o *GetRouteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get route params
func (o *GetRouteParams) WithHTTPClient(client *http.Client) *GetRouteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get route params
func (o *GetRouteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get route params
func (o *GetRouteParams) WithNetworkID(networkID string) *GetRouteParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get route params
func (o *GetRouteParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithRouteID adds the routeID to the get route params
func (o *GetRouteParams) WithRouteID(routeID string) *GetRouteParams {
	o.SetRouteID(routeID)
	return o
}

// SetRouteID adds the routeId to the get route params
func (o *GetRouteParams) SetRouteID(routeID string) {
	o.RouteID = routeID
}

// WithStackID adds the stackID to the get route params
func (o *GetRouteParams) WithStackID(stackID string) *GetRouteParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the get route params
func (o *GetRouteParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRouteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	// path param route_id
	if err := r.SetPathParam("route_id", o.RouteID); err != nil {
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

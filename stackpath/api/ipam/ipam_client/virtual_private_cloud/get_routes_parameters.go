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

// NewGetRoutesParams creates a new GetRoutesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRoutesParams() *GetRoutesParams {
	return &GetRoutesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRoutesParamsWithTimeout creates a new GetRoutesParams object
// with the ability to set a timeout on a request.
func NewGetRoutesParamsWithTimeout(timeout time.Duration) *GetRoutesParams {
	return &GetRoutesParams{
		timeout: timeout,
	}
}

// NewGetRoutesParamsWithContext creates a new GetRoutesParams object
// with the ability to set a context for a request.
func NewGetRoutesParamsWithContext(ctx context.Context) *GetRoutesParams {
	return &GetRoutesParams{
		Context: ctx,
	}
}

// NewGetRoutesParamsWithHTTPClient creates a new GetRoutesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRoutesParamsWithHTTPClient(client *http.Client) *GetRoutesParams {
	return &GetRoutesParams{
		HTTPClient: client,
	}
}

/* GetRoutesParams contains all the parameters to send to the API endpoint
   for the get routes operation.

   Typically these are written to a http.Request.
*/
type GetRoutesParams struct {

	/* NetworkID.

	   A VPC network ID or slug
	*/
	NetworkID string

	/* PageRequestAfter.

	   The cursor value after which data will be returned.
	*/
	PageRequestAfter *string

	/* PageRequestFilter.

	   SQL-style constraint filters.
	*/
	PageRequestFilter *string

	/* PageRequestFirst.

	   The number of items desired.
	*/
	PageRequestFirst *string

	/* PageRequestSortBy.

	   Sort the response by the given field.
	*/
	PageRequestSortBy *string

	/* StackID.

	   A stack ID or slug
	*/
	StackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get routes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRoutesParams) WithDefaults() *GetRoutesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get routes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRoutesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get routes params
func (o *GetRoutesParams) WithTimeout(timeout time.Duration) *GetRoutesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get routes params
func (o *GetRoutesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get routes params
func (o *GetRoutesParams) WithContext(ctx context.Context) *GetRoutesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get routes params
func (o *GetRoutesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get routes params
func (o *GetRoutesParams) WithHTTPClient(client *http.Client) *GetRoutesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get routes params
func (o *GetRoutesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get routes params
func (o *GetRoutesParams) WithNetworkID(networkID string) *GetRoutesParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get routes params
func (o *GetRoutesParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithPageRequestAfter adds the pageRequestAfter to the get routes params
func (o *GetRoutesParams) WithPageRequestAfter(pageRequestAfter *string) *GetRoutesParams {
	o.SetPageRequestAfter(pageRequestAfter)
	return o
}

// SetPageRequestAfter adds the pageRequestAfter to the get routes params
func (o *GetRoutesParams) SetPageRequestAfter(pageRequestAfter *string) {
	o.PageRequestAfter = pageRequestAfter
}

// WithPageRequestFilter adds the pageRequestFilter to the get routes params
func (o *GetRoutesParams) WithPageRequestFilter(pageRequestFilter *string) *GetRoutesParams {
	o.SetPageRequestFilter(pageRequestFilter)
	return o
}

// SetPageRequestFilter adds the pageRequestFilter to the get routes params
func (o *GetRoutesParams) SetPageRequestFilter(pageRequestFilter *string) {
	o.PageRequestFilter = pageRequestFilter
}

// WithPageRequestFirst adds the pageRequestFirst to the get routes params
func (o *GetRoutesParams) WithPageRequestFirst(pageRequestFirst *string) *GetRoutesParams {
	o.SetPageRequestFirst(pageRequestFirst)
	return o
}

// SetPageRequestFirst adds the pageRequestFirst to the get routes params
func (o *GetRoutesParams) SetPageRequestFirst(pageRequestFirst *string) {
	o.PageRequestFirst = pageRequestFirst
}

// WithPageRequestSortBy adds the pageRequestSortBy to the get routes params
func (o *GetRoutesParams) WithPageRequestSortBy(pageRequestSortBy *string) *GetRoutesParams {
	o.SetPageRequestSortBy(pageRequestSortBy)
	return o
}

// SetPageRequestSortBy adds the pageRequestSortBy to the get routes params
func (o *GetRoutesParams) SetPageRequestSortBy(pageRequestSortBy *string) {
	o.PageRequestSortBy = pageRequestSortBy
}

// WithStackID adds the stackID to the get routes params
func (o *GetRoutesParams) WithStackID(stackID string) *GetRoutesParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the get routes params
func (o *GetRoutesParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRoutesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if o.PageRequestAfter != nil {

		// query param page_request.after
		var qrPageRequestAfter string

		if o.PageRequestAfter != nil {
			qrPageRequestAfter = *o.PageRequestAfter
		}
		qPageRequestAfter := qrPageRequestAfter
		if qPageRequestAfter != "" {

			if err := r.SetQueryParam("page_request.after", qPageRequestAfter); err != nil {
				return err
			}
		}
	}

	if o.PageRequestFilter != nil {

		// query param page_request.filter
		var qrPageRequestFilter string

		if o.PageRequestFilter != nil {
			qrPageRequestFilter = *o.PageRequestFilter
		}
		qPageRequestFilter := qrPageRequestFilter
		if qPageRequestFilter != "" {

			if err := r.SetQueryParam("page_request.filter", qPageRequestFilter); err != nil {
				return err
			}
		}
	}

	if o.PageRequestFirst != nil {

		// query param page_request.first
		var qrPageRequestFirst string

		if o.PageRequestFirst != nil {
			qrPageRequestFirst = *o.PageRequestFirst
		}
		qPageRequestFirst := qrPageRequestFirst
		if qPageRequestFirst != "" {

			if err := r.SetQueryParam("page_request.first", qPageRequestFirst); err != nil {
				return err
			}
		}
	}

	if o.PageRequestSortBy != nil {

		// query param page_request.sort_by
		var qrPageRequestSortBy string

		if o.PageRequestSortBy != nil {
			qrPageRequestSortBy = *o.PageRequestSortBy
		}
		qPageRequestSortBy := qrPageRequestSortBy
		if qPageRequestSortBy != "" {

			if err := r.SetQueryParam("page_request.sort_by", qPageRequestSortBy); err != nil {
				return err
			}
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

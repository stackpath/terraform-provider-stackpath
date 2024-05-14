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

// NewGetNetworksParams creates a new GetNetworksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworksParams() *GetNetworksParams {
	return &GetNetworksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworksParamsWithTimeout creates a new GetNetworksParams object
// with the ability to set a timeout on a request.
func NewGetNetworksParamsWithTimeout(timeout time.Duration) *GetNetworksParams {
	return &GetNetworksParams{
		timeout: timeout,
	}
}

// NewGetNetworksParamsWithContext creates a new GetNetworksParams object
// with the ability to set a context for a request.
func NewGetNetworksParamsWithContext(ctx context.Context) *GetNetworksParams {
	return &GetNetworksParams{
		Context: ctx,
	}
}

// NewGetNetworksParamsWithHTTPClient creates a new GetNetworksParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworksParamsWithHTTPClient(client *http.Client) *GetNetworksParams {
	return &GetNetworksParams{
		HTTPClient: client,
	}
}

/*
GetNetworksParams contains all the parameters to send to the API endpoint

	for the get networks operation.

	Typically these are written to a http.Request.
*/
type GetNetworksParams struct {

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

// WithDefaults hydrates default values in the get networks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworksParams) WithDefaults() *GetNetworksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get networks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get networks params
func (o *GetNetworksParams) WithTimeout(timeout time.Duration) *GetNetworksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get networks params
func (o *GetNetworksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get networks params
func (o *GetNetworksParams) WithContext(ctx context.Context) *GetNetworksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get networks params
func (o *GetNetworksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get networks params
func (o *GetNetworksParams) WithHTTPClient(client *http.Client) *GetNetworksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get networks params
func (o *GetNetworksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageRequestAfter adds the pageRequestAfter to the get networks params
func (o *GetNetworksParams) WithPageRequestAfter(pageRequestAfter *string) *GetNetworksParams {
	o.SetPageRequestAfter(pageRequestAfter)
	return o
}

// SetPageRequestAfter adds the pageRequestAfter to the get networks params
func (o *GetNetworksParams) SetPageRequestAfter(pageRequestAfter *string) {
	o.PageRequestAfter = pageRequestAfter
}

// WithPageRequestFilter adds the pageRequestFilter to the get networks params
func (o *GetNetworksParams) WithPageRequestFilter(pageRequestFilter *string) *GetNetworksParams {
	o.SetPageRequestFilter(pageRequestFilter)
	return o
}

// SetPageRequestFilter adds the pageRequestFilter to the get networks params
func (o *GetNetworksParams) SetPageRequestFilter(pageRequestFilter *string) {
	o.PageRequestFilter = pageRequestFilter
}

// WithPageRequestFirst adds the pageRequestFirst to the get networks params
func (o *GetNetworksParams) WithPageRequestFirst(pageRequestFirst *string) *GetNetworksParams {
	o.SetPageRequestFirst(pageRequestFirst)
	return o
}

// SetPageRequestFirst adds the pageRequestFirst to the get networks params
func (o *GetNetworksParams) SetPageRequestFirst(pageRequestFirst *string) {
	o.PageRequestFirst = pageRequestFirst
}

// WithPageRequestSortBy adds the pageRequestSortBy to the get networks params
func (o *GetNetworksParams) WithPageRequestSortBy(pageRequestSortBy *string) *GetNetworksParams {
	o.SetPageRequestSortBy(pageRequestSortBy)
	return o
}

// SetPageRequestSortBy adds the pageRequestSortBy to the get networks params
func (o *GetNetworksParams) SetPageRequestSortBy(pageRequestSortBy *string) {
	o.PageRequestSortBy = pageRequestSortBy
}

// WithStackID adds the stackID to the get networks params
func (o *GetNetworksParams) WithStackID(stackID string) *GetNetworksParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the get networks params
func (o *GetNetworksParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

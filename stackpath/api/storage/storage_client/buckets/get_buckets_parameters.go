// Code generated by go-swagger; DO NOT EDIT.

package buckets

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

// NewGetBucketsParams creates a new GetBucketsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBucketsParams() *GetBucketsParams {
	return &GetBucketsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBucketsParamsWithTimeout creates a new GetBucketsParams object
// with the ability to set a timeout on a request.
func NewGetBucketsParamsWithTimeout(timeout time.Duration) *GetBucketsParams {
	return &GetBucketsParams{
		timeout: timeout,
	}
}

// NewGetBucketsParamsWithContext creates a new GetBucketsParams object
// with the ability to set a context for a request.
func NewGetBucketsParamsWithContext(ctx context.Context) *GetBucketsParams {
	return &GetBucketsParams{
		Context: ctx,
	}
}

// NewGetBucketsParamsWithHTTPClient creates a new GetBucketsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBucketsParamsWithHTTPClient(client *http.Client) *GetBucketsParams {
	return &GetBucketsParams{
		HTTPClient: client,
	}
}

/* GetBucketsParams contains all the parameters to send to the API endpoint
   for the get buckets operation.

   Typically these are written to a http.Request.
*/
type GetBucketsParams struct {

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

	   The ID for the stack for which the buckets will be retrieved
	*/
	StackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get buckets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBucketsParams) WithDefaults() *GetBucketsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get buckets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBucketsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get buckets params
func (o *GetBucketsParams) WithTimeout(timeout time.Duration) *GetBucketsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get buckets params
func (o *GetBucketsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get buckets params
func (o *GetBucketsParams) WithContext(ctx context.Context) *GetBucketsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get buckets params
func (o *GetBucketsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get buckets params
func (o *GetBucketsParams) WithHTTPClient(client *http.Client) *GetBucketsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get buckets params
func (o *GetBucketsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageRequestAfter adds the pageRequestAfter to the get buckets params
func (o *GetBucketsParams) WithPageRequestAfter(pageRequestAfter *string) *GetBucketsParams {
	o.SetPageRequestAfter(pageRequestAfter)
	return o
}

// SetPageRequestAfter adds the pageRequestAfter to the get buckets params
func (o *GetBucketsParams) SetPageRequestAfter(pageRequestAfter *string) {
	o.PageRequestAfter = pageRequestAfter
}

// WithPageRequestFilter adds the pageRequestFilter to the get buckets params
func (o *GetBucketsParams) WithPageRequestFilter(pageRequestFilter *string) *GetBucketsParams {
	o.SetPageRequestFilter(pageRequestFilter)
	return o
}

// SetPageRequestFilter adds the pageRequestFilter to the get buckets params
func (o *GetBucketsParams) SetPageRequestFilter(pageRequestFilter *string) {
	o.PageRequestFilter = pageRequestFilter
}

// WithPageRequestFirst adds the pageRequestFirst to the get buckets params
func (o *GetBucketsParams) WithPageRequestFirst(pageRequestFirst *string) *GetBucketsParams {
	o.SetPageRequestFirst(pageRequestFirst)
	return o
}

// SetPageRequestFirst adds the pageRequestFirst to the get buckets params
func (o *GetBucketsParams) SetPageRequestFirst(pageRequestFirst *string) {
	o.PageRequestFirst = pageRequestFirst
}

// WithPageRequestSortBy adds the pageRequestSortBy to the get buckets params
func (o *GetBucketsParams) WithPageRequestSortBy(pageRequestSortBy *string) *GetBucketsParams {
	o.SetPageRequestSortBy(pageRequestSortBy)
	return o
}

// SetPageRequestSortBy adds the pageRequestSortBy to the get buckets params
func (o *GetBucketsParams) SetPageRequestSortBy(pageRequestSortBy *string) {
	o.PageRequestSortBy = pageRequestSortBy
}

// WithStackID adds the stackID to the get buckets params
func (o *GetBucketsParams) WithStackID(stackID string) *GetBucketsParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the get buckets params
func (o *GetBucketsParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WriteToRequest writes these params to a swagger request
func (o *GetBucketsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

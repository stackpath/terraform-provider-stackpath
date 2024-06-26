// Code generated by go-swagger; DO NOT EDIT.

package instances

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

// NewGetWorkloadInstancesParams creates a new GetWorkloadInstancesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWorkloadInstancesParams() *GetWorkloadInstancesParams {
	return &GetWorkloadInstancesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkloadInstancesParamsWithTimeout creates a new GetWorkloadInstancesParams object
// with the ability to set a timeout on a request.
func NewGetWorkloadInstancesParamsWithTimeout(timeout time.Duration) *GetWorkloadInstancesParams {
	return &GetWorkloadInstancesParams{
		timeout: timeout,
	}
}

// NewGetWorkloadInstancesParamsWithContext creates a new GetWorkloadInstancesParams object
// with the ability to set a context for a request.
func NewGetWorkloadInstancesParamsWithContext(ctx context.Context) *GetWorkloadInstancesParams {
	return &GetWorkloadInstancesParams{
		Context: ctx,
	}
}

// NewGetWorkloadInstancesParamsWithHTTPClient creates a new GetWorkloadInstancesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWorkloadInstancesParamsWithHTTPClient(client *http.Client) *GetWorkloadInstancesParams {
	return &GetWorkloadInstancesParams{
		HTTPClient: client,
	}
}

/* GetWorkloadInstancesParams contains all the parameters to send to the API endpoint
   for the get workload instances operation.

   Typically these are written to a http.Request.
*/
type GetWorkloadInstancesParams struct {

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

	/* WorkloadID.

	   An EdgeCompute workload ID
	*/
	WorkloadID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get workload instances params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkloadInstancesParams) WithDefaults() *GetWorkloadInstancesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get workload instances params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkloadInstancesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get workload instances params
func (o *GetWorkloadInstancesParams) WithTimeout(timeout time.Duration) *GetWorkloadInstancesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workload instances params
func (o *GetWorkloadInstancesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workload instances params
func (o *GetWorkloadInstancesParams) WithContext(ctx context.Context) *GetWorkloadInstancesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workload instances params
func (o *GetWorkloadInstancesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workload instances params
func (o *GetWorkloadInstancesParams) WithHTTPClient(client *http.Client) *GetWorkloadInstancesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workload instances params
func (o *GetWorkloadInstancesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageRequestAfter adds the pageRequestAfter to the get workload instances params
func (o *GetWorkloadInstancesParams) WithPageRequestAfter(pageRequestAfter *string) *GetWorkloadInstancesParams {
	o.SetPageRequestAfter(pageRequestAfter)
	return o
}

// SetPageRequestAfter adds the pageRequestAfter to the get workload instances params
func (o *GetWorkloadInstancesParams) SetPageRequestAfter(pageRequestAfter *string) {
	o.PageRequestAfter = pageRequestAfter
}

// WithPageRequestFilter adds the pageRequestFilter to the get workload instances params
func (o *GetWorkloadInstancesParams) WithPageRequestFilter(pageRequestFilter *string) *GetWorkloadInstancesParams {
	o.SetPageRequestFilter(pageRequestFilter)
	return o
}

// SetPageRequestFilter adds the pageRequestFilter to the get workload instances params
func (o *GetWorkloadInstancesParams) SetPageRequestFilter(pageRequestFilter *string) {
	o.PageRequestFilter = pageRequestFilter
}

// WithPageRequestFirst adds the pageRequestFirst to the get workload instances params
func (o *GetWorkloadInstancesParams) WithPageRequestFirst(pageRequestFirst *string) *GetWorkloadInstancesParams {
	o.SetPageRequestFirst(pageRequestFirst)
	return o
}

// SetPageRequestFirst adds the pageRequestFirst to the get workload instances params
func (o *GetWorkloadInstancesParams) SetPageRequestFirst(pageRequestFirst *string) {
	o.PageRequestFirst = pageRequestFirst
}

// WithPageRequestSortBy adds the pageRequestSortBy to the get workload instances params
func (o *GetWorkloadInstancesParams) WithPageRequestSortBy(pageRequestSortBy *string) *GetWorkloadInstancesParams {
	o.SetPageRequestSortBy(pageRequestSortBy)
	return o
}

// SetPageRequestSortBy adds the pageRequestSortBy to the get workload instances params
func (o *GetWorkloadInstancesParams) SetPageRequestSortBy(pageRequestSortBy *string) {
	o.PageRequestSortBy = pageRequestSortBy
}

// WithStackID adds the stackID to the get workload instances params
func (o *GetWorkloadInstancesParams) WithStackID(stackID string) *GetWorkloadInstancesParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the get workload instances params
func (o *GetWorkloadInstancesParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WithWorkloadID adds the workloadID to the get workload instances params
func (o *GetWorkloadInstancesParams) WithWorkloadID(workloadID string) *GetWorkloadInstancesParams {
	o.SetWorkloadID(workloadID)
	return o
}

// SetWorkloadID adds the workloadId to the get workload instances params
func (o *GetWorkloadInstancesParams) SetWorkloadID(workloadID string) {
	o.WorkloadID = workloadID
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkloadInstancesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param workload_id
	if err := r.SetPathParam("workload_id", o.WorkloadID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

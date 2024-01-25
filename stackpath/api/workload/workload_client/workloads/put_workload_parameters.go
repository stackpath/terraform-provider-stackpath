// Code generated by go-swagger; DO NOT EDIT.

package workloads

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

	"github.com/stackpath/terraform-provider-stackpath/v2/stackpath/api/workload/workload_models"
)

// NewPutWorkloadParams creates a new PutWorkloadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutWorkloadParams() *PutWorkloadParams {
	return &PutWorkloadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutWorkloadParamsWithTimeout creates a new PutWorkloadParams object
// with the ability to set a timeout on a request.
func NewPutWorkloadParamsWithTimeout(timeout time.Duration) *PutWorkloadParams {
	return &PutWorkloadParams{
		timeout: timeout,
	}
}

// NewPutWorkloadParamsWithContext creates a new PutWorkloadParams object
// with the ability to set a context for a request.
func NewPutWorkloadParamsWithContext(ctx context.Context) *PutWorkloadParams {
	return &PutWorkloadParams{
		Context: ctx,
	}
}

// NewPutWorkloadParamsWithHTTPClient creates a new PutWorkloadParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutWorkloadParamsWithHTTPClient(client *http.Client) *PutWorkloadParams {
	return &PutWorkloadParams{
		HTTPClient: client,
	}
}

/*
PutWorkloadParams contains all the parameters to send to the API endpoint

	for the put workload operation.

	Typically these are written to a http.Request.
*/
type PutWorkloadParams struct {

	// Body.
	Body *workload_models.V1PutWorkloadRequest

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

// WithDefaults hydrates default values in the put workload params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutWorkloadParams) WithDefaults() *PutWorkloadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put workload params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutWorkloadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put workload params
func (o *PutWorkloadParams) WithTimeout(timeout time.Duration) *PutWorkloadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put workload params
func (o *PutWorkloadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put workload params
func (o *PutWorkloadParams) WithContext(ctx context.Context) *PutWorkloadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put workload params
func (o *PutWorkloadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put workload params
func (o *PutWorkloadParams) WithHTTPClient(client *http.Client) *PutWorkloadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put workload params
func (o *PutWorkloadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put workload params
func (o *PutWorkloadParams) WithBody(body *workload_models.V1PutWorkloadRequest) *PutWorkloadParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put workload params
func (o *PutWorkloadParams) SetBody(body *workload_models.V1PutWorkloadRequest) {
	o.Body = body
}

// WithStackID adds the stackID to the put workload params
func (o *PutWorkloadParams) WithStackID(stackID string) *PutWorkloadParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the put workload params
func (o *PutWorkloadParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WithWorkloadID adds the workloadID to the put workload params
func (o *PutWorkloadParams) WithWorkloadID(workloadID string) *PutWorkloadParams {
	o.SetWorkloadID(workloadID)
	return o
}

// SetWorkloadID adds the workloadId to the put workload params
func (o *PutWorkloadParams) SetWorkloadID(workloadID string) {
	o.WorkloadID = workloadID
}

// WriteToRequest writes these params to a swagger request
func (o *PutWorkloadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param workload_id
	if err := r.SetPathParam("workload_id", o.WorkloadID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
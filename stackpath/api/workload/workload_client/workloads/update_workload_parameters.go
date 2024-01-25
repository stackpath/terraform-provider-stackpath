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

// NewUpdateWorkloadParams creates a new UpdateWorkloadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateWorkloadParams() *UpdateWorkloadParams {
	return &UpdateWorkloadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateWorkloadParamsWithTimeout creates a new UpdateWorkloadParams object
// with the ability to set a timeout on a request.
func NewUpdateWorkloadParamsWithTimeout(timeout time.Duration) *UpdateWorkloadParams {
	return &UpdateWorkloadParams{
		timeout: timeout,
	}
}

// NewUpdateWorkloadParamsWithContext creates a new UpdateWorkloadParams object
// with the ability to set a context for a request.
func NewUpdateWorkloadParamsWithContext(ctx context.Context) *UpdateWorkloadParams {
	return &UpdateWorkloadParams{
		Context: ctx,
	}
}

// NewUpdateWorkloadParamsWithHTTPClient creates a new UpdateWorkloadParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateWorkloadParamsWithHTTPClient(client *http.Client) *UpdateWorkloadParams {
	return &UpdateWorkloadParams{
		HTTPClient: client,
	}
}

/*
UpdateWorkloadParams contains all the parameters to send to the API endpoint

	for the update workload operation.

	Typically these are written to a http.Request.
*/
type UpdateWorkloadParams struct {

	// Body.
	Body *workload_models.V1UpdateWorkloadRequest

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

// WithDefaults hydrates default values in the update workload params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateWorkloadParams) WithDefaults() *UpdateWorkloadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update workload params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateWorkloadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update workload params
func (o *UpdateWorkloadParams) WithTimeout(timeout time.Duration) *UpdateWorkloadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update workload params
func (o *UpdateWorkloadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update workload params
func (o *UpdateWorkloadParams) WithContext(ctx context.Context) *UpdateWorkloadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update workload params
func (o *UpdateWorkloadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update workload params
func (o *UpdateWorkloadParams) WithHTTPClient(client *http.Client) *UpdateWorkloadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update workload params
func (o *UpdateWorkloadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update workload params
func (o *UpdateWorkloadParams) WithBody(body *workload_models.V1UpdateWorkloadRequest) *UpdateWorkloadParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update workload params
func (o *UpdateWorkloadParams) SetBody(body *workload_models.V1UpdateWorkloadRequest) {
	o.Body = body
}

// WithStackID adds the stackID to the update workload params
func (o *UpdateWorkloadParams) WithStackID(stackID string) *UpdateWorkloadParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the update workload params
func (o *UpdateWorkloadParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WithWorkloadID adds the workloadID to the update workload params
func (o *UpdateWorkloadParams) WithWorkloadID(workloadID string) *UpdateWorkloadParams {
	o.SetWorkloadID(workloadID)
	return o
}

// SetWorkloadID adds the workloadId to the update workload params
func (o *UpdateWorkloadParams) SetWorkloadID(workloadID string) {
	o.WorkloadID = workloadID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateWorkloadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

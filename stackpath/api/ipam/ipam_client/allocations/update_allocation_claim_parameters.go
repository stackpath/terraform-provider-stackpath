// Code generated by go-swagger; DO NOT EDIT.

package allocations

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

// NewUpdateAllocationClaimParams creates a new UpdateAllocationClaimParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateAllocationClaimParams() *UpdateAllocationClaimParams {
	return &UpdateAllocationClaimParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAllocationClaimParamsWithTimeout creates a new UpdateAllocationClaimParams object
// with the ability to set a timeout on a request.
func NewUpdateAllocationClaimParamsWithTimeout(timeout time.Duration) *UpdateAllocationClaimParams {
	return &UpdateAllocationClaimParams{
		timeout: timeout,
	}
}

// NewUpdateAllocationClaimParamsWithContext creates a new UpdateAllocationClaimParams object
// with the ability to set a context for a request.
func NewUpdateAllocationClaimParamsWithContext(ctx context.Context) *UpdateAllocationClaimParams {
	return &UpdateAllocationClaimParams{
		Context: ctx,
	}
}

// NewUpdateAllocationClaimParamsWithHTTPClient creates a new UpdateAllocationClaimParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateAllocationClaimParamsWithHTTPClient(client *http.Client) *UpdateAllocationClaimParams {
	return &UpdateAllocationClaimParams{
		HTTPClient: client,
	}
}

/*
UpdateAllocationClaimParams contains all the parameters to send to the API endpoint

	for the update allocation claim operation.

	Typically these are written to a http.Request.
*/
type UpdateAllocationClaimParams struct {

	// AllocationClaimSlug.
	AllocationClaimSlug string

	// Body.
	Body *ipam_models.V1UpdateAllocationClaimRequest

	// StackID.
	StackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update allocation claim params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAllocationClaimParams) WithDefaults() *UpdateAllocationClaimParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update allocation claim params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAllocationClaimParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update allocation claim params
func (o *UpdateAllocationClaimParams) WithTimeout(timeout time.Duration) *UpdateAllocationClaimParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update allocation claim params
func (o *UpdateAllocationClaimParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update allocation claim params
func (o *UpdateAllocationClaimParams) WithContext(ctx context.Context) *UpdateAllocationClaimParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update allocation claim params
func (o *UpdateAllocationClaimParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update allocation claim params
func (o *UpdateAllocationClaimParams) WithHTTPClient(client *http.Client) *UpdateAllocationClaimParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update allocation claim params
func (o *UpdateAllocationClaimParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllocationClaimSlug adds the allocationClaimSlug to the update allocation claim params
func (o *UpdateAllocationClaimParams) WithAllocationClaimSlug(allocationClaimSlug string) *UpdateAllocationClaimParams {
	o.SetAllocationClaimSlug(allocationClaimSlug)
	return o
}

// SetAllocationClaimSlug adds the allocationClaimSlug to the update allocation claim params
func (o *UpdateAllocationClaimParams) SetAllocationClaimSlug(allocationClaimSlug string) {
	o.AllocationClaimSlug = allocationClaimSlug
}

// WithBody adds the body to the update allocation claim params
func (o *UpdateAllocationClaimParams) WithBody(body *ipam_models.V1UpdateAllocationClaimRequest) *UpdateAllocationClaimParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update allocation claim params
func (o *UpdateAllocationClaimParams) SetBody(body *ipam_models.V1UpdateAllocationClaimRequest) {
	o.Body = body
}

// WithStackID adds the stackID to the update allocation claim params
func (o *UpdateAllocationClaimParams) WithStackID(stackID string) *UpdateAllocationClaimParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the update allocation claim params
func (o *UpdateAllocationClaimParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAllocationClaimParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param allocation_claim_slug
	if err := r.SetPathParam("allocation_claim_slug", o.AllocationClaimSlug); err != nil {
		return err
	}
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

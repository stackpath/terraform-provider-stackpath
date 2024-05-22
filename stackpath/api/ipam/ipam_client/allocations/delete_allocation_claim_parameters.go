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
)

// NewDeleteAllocationClaimParams creates a new DeleteAllocationClaimParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAllocationClaimParams() *DeleteAllocationClaimParams {
	return &DeleteAllocationClaimParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAllocationClaimParamsWithTimeout creates a new DeleteAllocationClaimParams object
// with the ability to set a timeout on a request.
func NewDeleteAllocationClaimParamsWithTimeout(timeout time.Duration) *DeleteAllocationClaimParams {
	return &DeleteAllocationClaimParams{
		timeout: timeout,
	}
}

// NewDeleteAllocationClaimParamsWithContext creates a new DeleteAllocationClaimParams object
// with the ability to set a context for a request.
func NewDeleteAllocationClaimParamsWithContext(ctx context.Context) *DeleteAllocationClaimParams {
	return &DeleteAllocationClaimParams{
		Context: ctx,
	}
}

// NewDeleteAllocationClaimParamsWithHTTPClient creates a new DeleteAllocationClaimParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAllocationClaimParamsWithHTTPClient(client *http.Client) *DeleteAllocationClaimParams {
	return &DeleteAllocationClaimParams{
		HTTPClient: client,
	}
}

/*
DeleteAllocationClaimParams contains all the parameters to send to the API endpoint

	for the delete allocation claim operation.

	Typically these are written to a http.Request.
*/
type DeleteAllocationClaimParams struct {

	// AllocationClaimSlug.
	AllocationClaimSlug string

	// StackID.
	StackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete allocation claim params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAllocationClaimParams) WithDefaults() *DeleteAllocationClaimParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete allocation claim params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAllocationClaimParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete allocation claim params
func (o *DeleteAllocationClaimParams) WithTimeout(timeout time.Duration) *DeleteAllocationClaimParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete allocation claim params
func (o *DeleteAllocationClaimParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete allocation claim params
func (o *DeleteAllocationClaimParams) WithContext(ctx context.Context) *DeleteAllocationClaimParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete allocation claim params
func (o *DeleteAllocationClaimParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete allocation claim params
func (o *DeleteAllocationClaimParams) WithHTTPClient(client *http.Client) *DeleteAllocationClaimParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete allocation claim params
func (o *DeleteAllocationClaimParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllocationClaimSlug adds the allocationClaimSlug to the delete allocation claim params
func (o *DeleteAllocationClaimParams) WithAllocationClaimSlug(allocationClaimSlug string) *DeleteAllocationClaimParams {
	o.SetAllocationClaimSlug(allocationClaimSlug)
	return o
}

// SetAllocationClaimSlug adds the allocationClaimSlug to the delete allocation claim params
func (o *DeleteAllocationClaimParams) SetAllocationClaimSlug(allocationClaimSlug string) {
	o.AllocationClaimSlug = allocationClaimSlug
}

// WithStackID adds the stackID to the delete allocation claim params
func (o *DeleteAllocationClaimParams) WithStackID(stackID string) *DeleteAllocationClaimParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the delete allocation claim params
func (o *DeleteAllocationClaimParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAllocationClaimParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param allocation_claim_slug
	if err := r.SetPathParam("allocation_claim_slug", o.AllocationClaimSlug); err != nil {
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
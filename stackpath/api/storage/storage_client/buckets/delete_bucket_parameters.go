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
	"github.com/go-openapi/swag"
)

// NewDeleteBucketParams creates a new DeleteBucketParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteBucketParams() *DeleteBucketParams {
	return &DeleteBucketParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteBucketParamsWithTimeout creates a new DeleteBucketParams object
// with the ability to set a timeout on a request.
func NewDeleteBucketParamsWithTimeout(timeout time.Duration) *DeleteBucketParams {
	return &DeleteBucketParams{
		timeout: timeout,
	}
}

// NewDeleteBucketParamsWithContext creates a new DeleteBucketParams object
// with the ability to set a context for a request.
func NewDeleteBucketParamsWithContext(ctx context.Context) *DeleteBucketParams {
	return &DeleteBucketParams{
		Context: ctx,
	}
}

// NewDeleteBucketParamsWithHTTPClient creates a new DeleteBucketParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteBucketParamsWithHTTPClient(client *http.Client) *DeleteBucketParams {
	return &DeleteBucketParams{
		HTTPClient: client,
	}
}

/* DeleteBucketParams contains all the parameters to send to the API endpoint
   for the delete bucket operation.

   Typically these are written to a http.Request.
*/
type DeleteBucketParams struct {

	/* BucketID.

	   The ID for the bucket to delete
	*/
	BucketID string

	/* ForceDelete.

	   Force bucket deletion even if there is contents inside it.

	   Format: boolean
	*/
	ForceDelete *bool

	/* StackID.

	   The ID for the stack in which the bucket belongs
	*/
	StackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete bucket params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBucketParams) WithDefaults() *DeleteBucketParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete bucket params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBucketParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete bucket params
func (o *DeleteBucketParams) WithTimeout(timeout time.Duration) *DeleteBucketParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete bucket params
func (o *DeleteBucketParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete bucket params
func (o *DeleteBucketParams) WithContext(ctx context.Context) *DeleteBucketParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete bucket params
func (o *DeleteBucketParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete bucket params
func (o *DeleteBucketParams) WithHTTPClient(client *http.Client) *DeleteBucketParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete bucket params
func (o *DeleteBucketParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBucketID adds the bucketID to the delete bucket params
func (o *DeleteBucketParams) WithBucketID(bucketID string) *DeleteBucketParams {
	o.SetBucketID(bucketID)
	return o
}

// SetBucketID adds the bucketId to the delete bucket params
func (o *DeleteBucketParams) SetBucketID(bucketID string) {
	o.BucketID = bucketID
}

// WithForceDelete adds the forceDelete to the delete bucket params
func (o *DeleteBucketParams) WithForceDelete(forceDelete *bool) *DeleteBucketParams {
	o.SetForceDelete(forceDelete)
	return o
}

// SetForceDelete adds the forceDelete to the delete bucket params
func (o *DeleteBucketParams) SetForceDelete(forceDelete *bool) {
	o.ForceDelete = forceDelete
}

// WithStackID adds the stackID to the delete bucket params
func (o *DeleteBucketParams) WithStackID(stackID string) *DeleteBucketParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the delete bucket params
func (o *DeleteBucketParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteBucketParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param bucket_id
	if err := r.SetPathParam("bucket_id", o.BucketID); err != nil {
		return err
	}

	if o.ForceDelete != nil {

		// query param force_delete
		var qrForceDelete bool

		if o.ForceDelete != nil {
			qrForceDelete = *o.ForceDelete
		}
		qForceDelete := swag.FormatBool(qrForceDelete)
		if qForceDelete != "" {

			if err := r.SetQueryParam("force_delete", qForceDelete); err != nil {
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

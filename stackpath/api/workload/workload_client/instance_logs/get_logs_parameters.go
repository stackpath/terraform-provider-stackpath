// Code generated by go-swagger; DO NOT EDIT.

package instance_logs

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

// NewGetLogsParams creates a new GetLogsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLogsParams() *GetLogsParams {
	return &GetLogsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLogsParamsWithTimeout creates a new GetLogsParams object
// with the ability to set a timeout on a request.
func NewGetLogsParamsWithTimeout(timeout time.Duration) *GetLogsParams {
	return &GetLogsParams{
		timeout: timeout,
	}
}

// NewGetLogsParamsWithContext creates a new GetLogsParams object
// with the ability to set a context for a request.
func NewGetLogsParamsWithContext(ctx context.Context) *GetLogsParams {
	return &GetLogsParams{
		Context: ctx,
	}
}

// NewGetLogsParamsWithHTTPClient creates a new GetLogsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLogsParamsWithHTTPClient(client *http.Client) *GetLogsParams {
	return &GetLogsParams{
		HTTPClient: client,
	}
}

/* GetLogsParams contains all the parameters to send to the API endpoint
   for the get logs operation.

   Typically these are written to a http.Request.
*/
type GetLogsParams struct {

	/* ContainerName.

	   The name of the container to obtain logs for. This defaults to first container in the instance.
	*/
	ContainerName *string

	/* Follow.

	   Whether or not to follow the instance's log stream. This defaults to false.

	   Format: boolean
	*/
	Follow *bool

	/* InstanceName.

	   An EdgeCompute workload instance name
	*/
	InstanceName string

	/* LimitBytes.

	   The number of bytes to read from the server before terminating log output. This may not display a complete final line of logging, and may return slightly more or slightly less than the specified limit.
	*/
	LimitBytes *string

	/* Previous.

	   Whether or not to return log entries made by previously terminated containers. This defaults to false.

	   Format: boolean
	*/
	Previous *bool

	/* SinceSeconds.

	     A relative time in seconds before the current time from which to show logs. If this value precedes the time an instance was started, only logs since the instance's start time will be returned. If this value is in the future, no logs will be returned.

	Only one of since_seconds or since_time may be specified
	*/
	SinceSeconds *string

	/* SinceTime.

	     An RFC3339 timestamp from which to show logs. If this value precedes the time an instance was started, only logs since the instance's start time will be returned. If this value is in the future, no logs will be returned.

	Only one of since_seconds or since_time may be specified

	     Format: date-time
	*/
	SinceTime *strfmt.DateTime

	/* StackID.

	   A stack ID or slug
	*/
	StackID string

	/* TailLines.

	   The number of lines from the end of the logs to show. If not specified, logs are shown from the creation of the container or since_seconds or since_time.
	*/
	TailLines *string

	/* Timestamps.

	   Whether or not to add an RFC3339 or RFC3339Nano timestamp at the beginning of every line of log output. This defaults to false.

	   Format: boolean
	*/
	Timestamps *bool

	/* WorkloadID.

	   An EdgeCompute workload ID
	*/
	WorkloadID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get logs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLogsParams) WithDefaults() *GetLogsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get logs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLogsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get logs params
func (o *GetLogsParams) WithTimeout(timeout time.Duration) *GetLogsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get logs params
func (o *GetLogsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get logs params
func (o *GetLogsParams) WithContext(ctx context.Context) *GetLogsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get logs params
func (o *GetLogsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get logs params
func (o *GetLogsParams) WithHTTPClient(client *http.Client) *GetLogsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get logs params
func (o *GetLogsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContainerName adds the containerName to the get logs params
func (o *GetLogsParams) WithContainerName(containerName *string) *GetLogsParams {
	o.SetContainerName(containerName)
	return o
}

// SetContainerName adds the containerName to the get logs params
func (o *GetLogsParams) SetContainerName(containerName *string) {
	o.ContainerName = containerName
}

// WithFollow adds the follow to the get logs params
func (o *GetLogsParams) WithFollow(follow *bool) *GetLogsParams {
	o.SetFollow(follow)
	return o
}

// SetFollow adds the follow to the get logs params
func (o *GetLogsParams) SetFollow(follow *bool) {
	o.Follow = follow
}

// WithInstanceName adds the instanceName to the get logs params
func (o *GetLogsParams) WithInstanceName(instanceName string) *GetLogsParams {
	o.SetInstanceName(instanceName)
	return o
}

// SetInstanceName adds the instanceName to the get logs params
func (o *GetLogsParams) SetInstanceName(instanceName string) {
	o.InstanceName = instanceName
}

// WithLimitBytes adds the limitBytes to the get logs params
func (o *GetLogsParams) WithLimitBytes(limitBytes *string) *GetLogsParams {
	o.SetLimitBytes(limitBytes)
	return o
}

// SetLimitBytes adds the limitBytes to the get logs params
func (o *GetLogsParams) SetLimitBytes(limitBytes *string) {
	o.LimitBytes = limitBytes
}

// WithPrevious adds the previous to the get logs params
func (o *GetLogsParams) WithPrevious(previous *bool) *GetLogsParams {
	o.SetPrevious(previous)
	return o
}

// SetPrevious adds the previous to the get logs params
func (o *GetLogsParams) SetPrevious(previous *bool) {
	o.Previous = previous
}

// WithSinceSeconds adds the sinceSeconds to the get logs params
func (o *GetLogsParams) WithSinceSeconds(sinceSeconds *string) *GetLogsParams {
	o.SetSinceSeconds(sinceSeconds)
	return o
}

// SetSinceSeconds adds the sinceSeconds to the get logs params
func (o *GetLogsParams) SetSinceSeconds(sinceSeconds *string) {
	o.SinceSeconds = sinceSeconds
}

// WithSinceTime adds the sinceTime to the get logs params
func (o *GetLogsParams) WithSinceTime(sinceTime *strfmt.DateTime) *GetLogsParams {
	o.SetSinceTime(sinceTime)
	return o
}

// SetSinceTime adds the sinceTime to the get logs params
func (o *GetLogsParams) SetSinceTime(sinceTime *strfmt.DateTime) {
	o.SinceTime = sinceTime
}

// WithStackID adds the stackID to the get logs params
func (o *GetLogsParams) WithStackID(stackID string) *GetLogsParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the get logs params
func (o *GetLogsParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WithTailLines adds the tailLines to the get logs params
func (o *GetLogsParams) WithTailLines(tailLines *string) *GetLogsParams {
	o.SetTailLines(tailLines)
	return o
}

// SetTailLines adds the tailLines to the get logs params
func (o *GetLogsParams) SetTailLines(tailLines *string) {
	o.TailLines = tailLines
}

// WithTimestamps adds the timestamps to the get logs params
func (o *GetLogsParams) WithTimestamps(timestamps *bool) *GetLogsParams {
	o.SetTimestamps(timestamps)
	return o
}

// SetTimestamps adds the timestamps to the get logs params
func (o *GetLogsParams) SetTimestamps(timestamps *bool) {
	o.Timestamps = timestamps
}

// WithWorkloadID adds the workloadID to the get logs params
func (o *GetLogsParams) WithWorkloadID(workloadID string) *GetLogsParams {
	o.SetWorkloadID(workloadID)
	return o
}

// SetWorkloadID adds the workloadId to the get logs params
func (o *GetLogsParams) SetWorkloadID(workloadID string) {
	o.WorkloadID = workloadID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ContainerName != nil {

		// query param container_name
		var qrContainerName string

		if o.ContainerName != nil {
			qrContainerName = *o.ContainerName
		}
		qContainerName := qrContainerName
		if qContainerName != "" {

			if err := r.SetQueryParam("container_name", qContainerName); err != nil {
				return err
			}
		}
	}

	if o.Follow != nil {

		// query param follow
		var qrFollow bool

		if o.Follow != nil {
			qrFollow = *o.Follow
		}
		qFollow := swag.FormatBool(qrFollow)
		if qFollow != "" {

			if err := r.SetQueryParam("follow", qFollow); err != nil {
				return err
			}
		}
	}

	// path param instance_name
	if err := r.SetPathParam("instance_name", o.InstanceName); err != nil {
		return err
	}

	if o.LimitBytes != nil {

		// query param limit_bytes
		var qrLimitBytes string

		if o.LimitBytes != nil {
			qrLimitBytes = *o.LimitBytes
		}
		qLimitBytes := qrLimitBytes
		if qLimitBytes != "" {

			if err := r.SetQueryParam("limit_bytes", qLimitBytes); err != nil {
				return err
			}
		}
	}

	if o.Previous != nil {

		// query param previous
		var qrPrevious bool

		if o.Previous != nil {
			qrPrevious = *o.Previous
		}
		qPrevious := swag.FormatBool(qrPrevious)
		if qPrevious != "" {

			if err := r.SetQueryParam("previous", qPrevious); err != nil {
				return err
			}
		}
	}

	if o.SinceSeconds != nil {

		// query param since_seconds
		var qrSinceSeconds string

		if o.SinceSeconds != nil {
			qrSinceSeconds = *o.SinceSeconds
		}
		qSinceSeconds := qrSinceSeconds
		if qSinceSeconds != "" {

			if err := r.SetQueryParam("since_seconds", qSinceSeconds); err != nil {
				return err
			}
		}
	}

	if o.SinceTime != nil {

		// query param since_time
		var qrSinceTime strfmt.DateTime

		if o.SinceTime != nil {
			qrSinceTime = *o.SinceTime
		}
		qSinceTime := qrSinceTime.String()
		if qSinceTime != "" {

			if err := r.SetQueryParam("since_time", qSinceTime); err != nil {
				return err
			}
		}
	}

	// path param stack_id
	if err := r.SetPathParam("stack_id", o.StackID); err != nil {
		return err
	}

	if o.TailLines != nil {

		// query param tail_lines
		var qrTailLines string

		if o.TailLines != nil {
			qrTailLines = *o.TailLines
		}
		qTailLines := qrTailLines
		if qTailLines != "" {

			if err := r.SetQueryParam("tail_lines", qTailLines); err != nil {
				return err
			}
		}
	}

	if o.Timestamps != nil {

		// query param timestamps
		var qrTimestamps bool

		if o.Timestamps != nil {
			qrTimestamps = *o.Timestamps
		}
		qTimestamps := swag.FormatBool(qrTimestamps)
		if qTimestamps != "" {

			if err := r.SetQueryParam("timestamps", qTimestamps); err != nil {
				return err
			}
		}
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

// Code generated by go-swagger; DO NOT EDIT.

package metrics

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

// NewGetMetricsParams creates a new GetMetricsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMetricsParams() *GetMetricsParams {
	return &GetMetricsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMetricsParamsWithTimeout creates a new GetMetricsParams object
// with the ability to set a timeout on a request.
func NewGetMetricsParamsWithTimeout(timeout time.Duration) *GetMetricsParams {
	return &GetMetricsParams{
		timeout: timeout,
	}
}

// NewGetMetricsParamsWithContext creates a new GetMetricsParams object
// with the ability to set a context for a request.
func NewGetMetricsParamsWithContext(ctx context.Context) *GetMetricsParams {
	return &GetMetricsParams{
		Context: ctx,
	}
}

// NewGetMetricsParamsWithHTTPClient creates a new GetMetricsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMetricsParamsWithHTTPClient(client *http.Client) *GetMetricsParams {
	return &GetMetricsParams{
		HTTPClient: client,
	}
}

/*
GetMetricsParams contains all the parameters to send to the API endpoint

	for the get metrics operation.

	Typically these are written to a http.Request.
*/
type GetMetricsParams struct {

	/* EndDate.

	   An upper bound date to search metrics for.

	   Format: date-time
	*/
	EndDate *strfmt.DateTime

	// Granularity.
	//
	// Default: "DEFAULT"
	Granularity *string

	// GroupBy.
	//
	// Default: "NONE"
	GroupBy *string

	/* Grouping.

	   List of fields to group by.
	*/
	Grouping []string

	/* InstanceName.

	   An EdgeCompute workload instance name
	*/
	InstanceName *string

	/* Pop.

	   A StackPath POP to filter traffic metrics for. This field does not apply when retrieving INSTANCE type metrics
	*/
	Pop *string

	// Region.
	//
	// Default: "ALL"
	Region *string

	/* StackID.

	   A stack ID or slug
	*/
	StackID string

	/* StartDate.

	   A lower bound date to search metrics for.

	   Format: date-time
	*/
	StartDate *strfmt.DateTime

	// Type.
	//
	// Default: "BANDWIDTH"
	Type *string

	/* WorkloadID.

	   An EdgeCompute workload ID
	*/
	WorkloadID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get metrics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMetricsParams) WithDefaults() *GetMetricsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get metrics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMetricsParams) SetDefaults() {
	var (
		granularityDefault = string("DEFAULT")

		groupByDefault = string("NONE")

		regionDefault = string("ALL")

		typeVarDefault = string("BANDWIDTH")
	)

	val := GetMetricsParams{
		Granularity: &granularityDefault,
		GroupBy:     &groupByDefault,
		Region:      &regionDefault,
		Type:        &typeVarDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get metrics params
func (o *GetMetricsParams) WithTimeout(timeout time.Duration) *GetMetricsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get metrics params
func (o *GetMetricsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get metrics params
func (o *GetMetricsParams) WithContext(ctx context.Context) *GetMetricsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get metrics params
func (o *GetMetricsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get metrics params
func (o *GetMetricsParams) WithHTTPClient(client *http.Client) *GetMetricsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get metrics params
func (o *GetMetricsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndDate adds the endDate to the get metrics params
func (o *GetMetricsParams) WithEndDate(endDate *strfmt.DateTime) *GetMetricsParams {
	o.SetEndDate(endDate)
	return o
}

// SetEndDate adds the endDate to the get metrics params
func (o *GetMetricsParams) SetEndDate(endDate *strfmt.DateTime) {
	o.EndDate = endDate
}

// WithGranularity adds the granularity to the get metrics params
func (o *GetMetricsParams) WithGranularity(granularity *string) *GetMetricsParams {
	o.SetGranularity(granularity)
	return o
}

// SetGranularity adds the granularity to the get metrics params
func (o *GetMetricsParams) SetGranularity(granularity *string) {
	o.Granularity = granularity
}

// WithGroupBy adds the groupBy to the get metrics params
func (o *GetMetricsParams) WithGroupBy(groupBy *string) *GetMetricsParams {
	o.SetGroupBy(groupBy)
	return o
}

// SetGroupBy adds the groupBy to the get metrics params
func (o *GetMetricsParams) SetGroupBy(groupBy *string) {
	o.GroupBy = groupBy
}

// WithGrouping adds the grouping to the get metrics params
func (o *GetMetricsParams) WithGrouping(grouping []string) *GetMetricsParams {
	o.SetGrouping(grouping)
	return o
}

// SetGrouping adds the grouping to the get metrics params
func (o *GetMetricsParams) SetGrouping(grouping []string) {
	o.Grouping = grouping
}

// WithInstanceName adds the instanceName to the get metrics params
func (o *GetMetricsParams) WithInstanceName(instanceName *string) *GetMetricsParams {
	o.SetInstanceName(instanceName)
	return o
}

// SetInstanceName adds the instanceName to the get metrics params
func (o *GetMetricsParams) SetInstanceName(instanceName *string) {
	o.InstanceName = instanceName
}

// WithPop adds the pop to the get metrics params
func (o *GetMetricsParams) WithPop(pop *string) *GetMetricsParams {
	o.SetPop(pop)
	return o
}

// SetPop adds the pop to the get metrics params
func (o *GetMetricsParams) SetPop(pop *string) {
	o.Pop = pop
}

// WithRegion adds the region to the get metrics params
func (o *GetMetricsParams) WithRegion(region *string) *GetMetricsParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the get metrics params
func (o *GetMetricsParams) SetRegion(region *string) {
	o.Region = region
}

// WithStackID adds the stackID to the get metrics params
func (o *GetMetricsParams) WithStackID(stackID string) *GetMetricsParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the get metrics params
func (o *GetMetricsParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WithStartDate adds the startDate to the get metrics params
func (o *GetMetricsParams) WithStartDate(startDate *strfmt.DateTime) *GetMetricsParams {
	o.SetStartDate(startDate)
	return o
}

// SetStartDate adds the startDate to the get metrics params
func (o *GetMetricsParams) SetStartDate(startDate *strfmt.DateTime) {
	o.StartDate = startDate
}

// WithType adds the typeVar to the get metrics params
func (o *GetMetricsParams) WithType(typeVar *string) *GetMetricsParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get metrics params
func (o *GetMetricsParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithWorkloadID adds the workloadID to the get metrics params
func (o *GetMetricsParams) WithWorkloadID(workloadID *string) *GetMetricsParams {
	o.SetWorkloadID(workloadID)
	return o
}

// SetWorkloadID adds the workloadId to the get metrics params
func (o *GetMetricsParams) SetWorkloadID(workloadID *string) {
	o.WorkloadID = workloadID
}

// WriteToRequest writes these params to a swagger request
func (o *GetMetricsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EndDate != nil {

		// query param end_date
		var qrEndDate strfmt.DateTime

		if o.EndDate != nil {
			qrEndDate = *o.EndDate
		}
		qEndDate := qrEndDate.String()
		if qEndDate != "" {

			if err := r.SetQueryParam("end_date", qEndDate); err != nil {
				return err
			}
		}
	}

	if o.Granularity != nil {

		// query param granularity
		var qrGranularity string

		if o.Granularity != nil {
			qrGranularity = *o.Granularity
		}
		qGranularity := qrGranularity
		if qGranularity != "" {

			if err := r.SetQueryParam("granularity", qGranularity); err != nil {
				return err
			}
		}
	}

	if o.GroupBy != nil {

		// query param group_by
		var qrGroupBy string

		if o.GroupBy != nil {
			qrGroupBy = *o.GroupBy
		}
		qGroupBy := qrGroupBy
		if qGroupBy != "" {

			if err := r.SetQueryParam("group_by", qGroupBy); err != nil {
				return err
			}
		}
	}

	if o.Grouping != nil {

		// binding items for grouping
		joinedGrouping := o.bindParamGrouping(reg)

		// query array param grouping
		if err := r.SetQueryParam("grouping", joinedGrouping...); err != nil {
			return err
		}
	}

	if o.InstanceName != nil {

		// query param instance_name
		var qrInstanceName string

		if o.InstanceName != nil {
			qrInstanceName = *o.InstanceName
		}
		qInstanceName := qrInstanceName
		if qInstanceName != "" {

			if err := r.SetQueryParam("instance_name", qInstanceName); err != nil {
				return err
			}
		}
	}

	if o.Pop != nil {

		// query param pop
		var qrPop string

		if o.Pop != nil {
			qrPop = *o.Pop
		}
		qPop := qrPop
		if qPop != "" {

			if err := r.SetQueryParam("pop", qPop); err != nil {
				return err
			}
		}
	}

	if o.Region != nil {

		// query param region
		var qrRegion string

		if o.Region != nil {
			qrRegion = *o.Region
		}
		qRegion := qrRegion
		if qRegion != "" {

			if err := r.SetQueryParam("region", qRegion); err != nil {
				return err
			}
		}
	}

	// path param stack_id
	if err := r.SetPathParam("stack_id", o.StackID); err != nil {
		return err
	}

	if o.StartDate != nil {

		// query param start_date
		var qrStartDate strfmt.DateTime

		if o.StartDate != nil {
			qrStartDate = *o.StartDate
		}
		qStartDate := qrStartDate.String()
		if qStartDate != "" {

			if err := r.SetQueryParam("start_date", qStartDate); err != nil {
				return err
			}
		}
	}

	if o.Type != nil {

		// query param type
		var qrType string

		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}
	}

	if o.WorkloadID != nil {

		// query param workload_id
		var qrWorkloadID string

		if o.WorkloadID != nil {
			qrWorkloadID = *o.WorkloadID
		}
		qWorkloadID := qrWorkloadID
		if qWorkloadID != "" {

			if err := r.SetQueryParam("workload_id", qWorkloadID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetMetrics binds the parameter grouping
func (o *GetMetricsParams) bindParamGrouping(formats strfmt.Registry) []string {
	groupingIR := o.Grouping

	var groupingIC []string
	for _, groupingIIR := range groupingIR { // explode []string

		groupingIIV := groupingIIR // string as string
		groupingIC = append(groupingIC, groupingIIV)
	}

	// items.CollectionFormat: ""
	groupingIS := swag.JoinByFormat(groupingIC, "")

	return groupingIS
}

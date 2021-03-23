// Code generated by go-swagger; DO NOT EDIT.

package workload_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// V1WorkloadStatus Which status a workload is currently in
//
// - ACTIVE: The workload is active
//  - SUSPENDED: The workload is suspended
//  - BILLING_SUSPENDED: The workload is suspended due to non-payment
//  - INACTIVE: The workload is inactive or has been deleted
//
// swagger:model v1WorkloadStatus
type V1WorkloadStatus string

const (

	// V1WorkloadStatusACTIVE captures enum value "ACTIVE"
	V1WorkloadStatusACTIVE V1WorkloadStatus = "ACTIVE"

	// V1WorkloadStatusSUSPENDED captures enum value "SUSPENDED"
	V1WorkloadStatusSUSPENDED V1WorkloadStatus = "SUSPENDED"

	// V1WorkloadStatusBILLINGSUSPENDED captures enum value "BILLING_SUSPENDED"
	V1WorkloadStatusBILLINGSUSPENDED V1WorkloadStatus = "BILLING_SUSPENDED"

	// V1WorkloadStatusINACTIVE captures enum value "INACTIVE"
	V1WorkloadStatusINACTIVE V1WorkloadStatus = "INACTIVE"
)

// for schema
var v1WorkloadStatusEnum []interface{}

func init() {
	var res []V1WorkloadStatus
	if err := json.Unmarshal([]byte(`["ACTIVE","SUSPENDED","BILLING_SUSPENDED","INACTIVE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1WorkloadStatusEnum = append(v1WorkloadStatusEnum, v)
	}
}

func (m V1WorkloadStatus) validateV1WorkloadStatusEnum(path, location string, value V1WorkloadStatus) error {
	if err := validate.EnumCase(path, location, value, v1WorkloadStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 workload status
func (m V1WorkloadStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1WorkloadStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this v1 workload status based on context it is used
func (m V1WorkloadStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

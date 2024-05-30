// Code generated by go-swagger; DO NOT EDIT.

package ipam_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// V1ConditionStatus Valid condition statuses.
//
//   - CONDITION_STATUS_UNSPECIFIED: The condition status has not been set.
//   - CONDITION_STATUS_TRUE: The condition is true.
//   - CONDITION_STATUS_FALSE: The condition is false.
//   - CONDITION_STATUS_UNKNOWN: The condition status is unknown.
//
// swagger:model v1ConditionStatus
type V1ConditionStatus string

func NewV1ConditionStatus(value V1ConditionStatus) *V1ConditionStatus {
	v := value
	return &v
}

const (

	// V1ConditionStatusCONDITIONSTATUSUNSPECIFIED captures enum value "CONDITION_STATUS_UNSPECIFIED"
	V1ConditionStatusCONDITIONSTATUSUNSPECIFIED V1ConditionStatus = "CONDITION_STATUS_UNSPECIFIED"

	// V1ConditionStatusCONDITIONSTATUSTRUE captures enum value "CONDITION_STATUS_TRUE"
	V1ConditionStatusCONDITIONSTATUSTRUE V1ConditionStatus = "CONDITION_STATUS_TRUE"

	// V1ConditionStatusCONDITIONSTATUSFALSE captures enum value "CONDITION_STATUS_FALSE"
	V1ConditionStatusCONDITIONSTATUSFALSE V1ConditionStatus = "CONDITION_STATUS_FALSE"

	// V1ConditionStatusCONDITIONSTATUSUNKNOWN captures enum value "CONDITION_STATUS_UNKNOWN"
	V1ConditionStatusCONDITIONSTATUSUNKNOWN V1ConditionStatus = "CONDITION_STATUS_UNKNOWN"
)

// for schema
var v1ConditionStatusEnum []interface{}

func init() {
	var res []V1ConditionStatus
	if err := json.Unmarshal([]byte(`["CONDITION_STATUS_UNSPECIFIED","CONDITION_STATUS_TRUE","CONDITION_STATUS_FALSE","CONDITION_STATUS_UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1ConditionStatusEnum = append(v1ConditionStatusEnum, v)
	}
}

func (m V1ConditionStatus) validateV1ConditionStatusEnum(path, location string, value V1ConditionStatus) error {
	if err := validate.EnumCase(path, location, value, v1ConditionStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 condition status
func (m V1ConditionStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1ConditionStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this v1 condition status based on context it is used
func (m V1ConditionStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

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

// Apimetav1ConditionStatus apimetav1 condition status
//
// swagger:model apimetav1ConditionStatus
type Apimetav1ConditionStatus string

func NewApimetav1ConditionStatus(value Apimetav1ConditionStatus) *Apimetav1ConditionStatus {
	v := value
	return &v
}

const (

	// Apimetav1ConditionStatusCONDITIONSTATUSUNSPECIFIED captures enum value "CONDITION_STATUS_UNSPECIFIED"
	Apimetav1ConditionStatusCONDITIONSTATUSUNSPECIFIED Apimetav1ConditionStatus = "CONDITION_STATUS_UNSPECIFIED"

	// Apimetav1ConditionStatusCONDITIONSTATUSTRUE captures enum value "CONDITION_STATUS_TRUE"
	Apimetav1ConditionStatusCONDITIONSTATUSTRUE Apimetav1ConditionStatus = "CONDITION_STATUS_TRUE"

	// Apimetav1ConditionStatusCONDITIONSTATUSFALSE captures enum value "CONDITION_STATUS_FALSE"
	Apimetav1ConditionStatusCONDITIONSTATUSFALSE Apimetav1ConditionStatus = "CONDITION_STATUS_FALSE"

	// Apimetav1ConditionStatusCONDITIONSTATUSUNKNOWN captures enum value "CONDITION_STATUS_UNKNOWN"
	Apimetav1ConditionStatusCONDITIONSTATUSUNKNOWN Apimetav1ConditionStatus = "CONDITION_STATUS_UNKNOWN"
)

// for schema
var apimetav1ConditionStatusEnum []interface{}

func init() {
	var res []Apimetav1ConditionStatus
	if err := json.Unmarshal([]byte(`["CONDITION_STATUS_UNSPECIFIED","CONDITION_STATUS_TRUE","CONDITION_STATUS_FALSE","CONDITION_STATUS_UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		apimetav1ConditionStatusEnum = append(apimetav1ConditionStatusEnum, v)
	}
}

func (m Apimetav1ConditionStatus) validateApimetav1ConditionStatusEnum(path, location string, value Apimetav1ConditionStatus) error {
	if err := validate.EnumCase(path, location, value, apimetav1ConditionStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this apimetav1 condition status
func (m Apimetav1ConditionStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateApimetav1ConditionStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this apimetav1 condition status based on context it is used
func (m Apimetav1ConditionStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

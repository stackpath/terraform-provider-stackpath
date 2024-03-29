// Code generated by go-swagger; DO NOT EDIT.

package workload_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DataVectorResult Time series containing a single sample for each time series, all sharing the same timestamp
//
// swagger:model DataVectorResult
type DataVectorResult struct {

	// The data points' labels
	Metric map[string]string `json:"metric,omitempty"`

	// value
	Value *DataValue `json:"value,omitempty"`
}

// Validate validates this data vector result
func (m *DataVectorResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataVectorResult) validateValue(formats strfmt.Registry) error {
	if swag.IsZero(m.Value) { // not required
		return nil
	}

	if m.Value != nil {
		if err := m.Value.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("value")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this data vector result based on the context it is used
func (m *DataVectorResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataVectorResult) contextValidateValue(ctx context.Context, formats strfmt.Registry) error {

	if m.Value != nil {

		if swag.IsZero(m.Value) { // not required
			return nil
		}

		if err := m.Value.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("value")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DataVectorResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataVectorResult) UnmarshalBinary(b []byte) error {
	var res DataVectorResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

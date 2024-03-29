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

// MetricsData The data points in a metrics collection
//
// swagger:model MetricsData
type MetricsData struct {

	// matrix
	Matrix *DataMatrix `json:"matrix,omitempty"`

	// vector
	Vector *DataVector `json:"vector,omitempty"`
}

// Validate validates this metrics data
func (m *MetricsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMatrix(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVector(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetricsData) validateMatrix(formats strfmt.Registry) error {
	if swag.IsZero(m.Matrix) { // not required
		return nil
	}

	if m.Matrix != nil {
		if err := m.Matrix.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("matrix")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce
			}
			return err
		}
	}

	return nil
}

func (m *MetricsData) validateVector(formats strfmt.Registry) error {
	if swag.IsZero(m.Vector) { // not required
		return nil
	}

	if m.Vector != nil {
		if err := m.Vector.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vector")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this metrics data based on the context it is used
func (m *MetricsData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMatrix(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVector(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetricsData) contextValidateMatrix(ctx context.Context, formats strfmt.Registry) error {

	if m.Matrix != nil {

		if swag.IsZero(m.Matrix) { // not required
			return nil
		}

		if err := m.Matrix.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("matrix")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce
			}
			return err
		}
	}

	return nil
}

func (m *MetricsData) contextValidateVector(ctx context.Context, formats strfmt.Registry) error {

	if m.Vector != nil {

		if swag.IsZero(m.Vector) { // not required
			return nil
		}

		if err := m.Vector.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vector")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MetricsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetricsData) UnmarshalBinary(b []byte) error {
	var res MetricsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

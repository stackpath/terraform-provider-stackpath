// Code generated by go-swagger; DO NOT EDIT.

package workload_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1Condition v1 condition
//
// swagger:model v1Condition
type V1Condition struct {

	// last transition time
	// Format: date-time
	LastTransitionTime strfmt.DateTime `json:"lastTransitionTime,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// observed version
	ObservedVersion string `json:"observedVersion,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`

	// status
	Status *Apimetav1ConditionStatus `json:"status,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this v1 condition
func (m *V1Condition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastTransitionTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Condition) validateLastTransitionTime(formats strfmt.Registry) error {
	if swag.IsZero(m.LastTransitionTime) { // not required
		return nil
	}

	if err := validate.FormatOf("lastTransitionTime", "body", "date-time", m.LastTransitionTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1Condition) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 condition based on the context it is used
func (m *V1Condition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Condition) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {
		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Condition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Condition) UnmarshalBinary(b []byte) error {
	var res V1Condition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

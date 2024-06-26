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

// V1ImageCondition Further information about an image's status
//
// swagger:model v1ImageCondition
type V1ImageCondition struct {

	// The last time the condition was checked
	// Format: date-time
	CheckedAt strfmt.DateTime `json:"checkedAt,omitempty"`

	// A human readable message with details regarding the condition
	Message string `json:"message,omitempty"`

	// A stable identifier for the reason the condition is in its current state
	Reason string `json:"reason,omitempty"`

	// status
	Status *V1ImageConditionStatus `json:"status,omitempty"`

	// The last time the condition transitioned status
	// Format: date-time
	TransitionedAt strfmt.DateTime `json:"transitionedAt,omitempty"`

	// Type of the condition
	Type string `json:"type,omitempty"`
}

// Validate validates this v1 image condition
func (m *V1ImageCondition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCheckedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransitionedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ImageCondition) validateCheckedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CheckedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("checkedAt", "body", "date-time", m.CheckedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1ImageCondition) validateStatus(formats strfmt.Registry) error {
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

func (m *V1ImageCondition) validateTransitionedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.TransitionedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("transitionedAt", "body", "date-time", m.TransitionedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 image condition based on the context it is used
func (m *V1ImageCondition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ImageCondition) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

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
func (m *V1ImageCondition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ImageCondition) UnmarshalBinary(b []byte) error {
	var res V1ImageCondition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

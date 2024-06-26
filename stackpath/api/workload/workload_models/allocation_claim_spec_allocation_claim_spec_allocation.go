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

// AllocationClaimSpecAllocationClaimSpecAllocation allocation claim spec allocation claim spec allocation
//
// swagger:model AllocationClaimSpecAllocationClaimSpecAllocation
type AllocationClaimSpecAllocationClaimSpecAllocation struct {

	// name
	Name string `json:"name,omitempty"`

	// selector
	Selector *AllocationClaimSpecAllocationClaimSpecAllocationSelector `json:"selector,omitempty"`

	// template
	Template *V1Allocation `json:"template,omitempty"`
}

// Validate validates this allocation claim spec allocation claim spec allocation
func (m *AllocationClaimSpecAllocationClaimSpecAllocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelector(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AllocationClaimSpecAllocationClaimSpecAllocation) validateSelector(formats strfmt.Registry) error {
	if swag.IsZero(m.Selector) { // not required
		return nil
	}

	if m.Selector != nil {
		if err := m.Selector.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("selector")
			}
			return err
		}
	}

	return nil
}

func (m *AllocationClaimSpecAllocationClaimSpecAllocation) validateTemplate(formats strfmt.Registry) error {
	if swag.IsZero(m.Template) { // not required
		return nil
	}

	if m.Template != nil {
		if err := m.Template.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("template")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this allocation claim spec allocation claim spec allocation based on the context it is used
func (m *AllocationClaimSpecAllocationClaimSpecAllocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelector(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTemplate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AllocationClaimSpecAllocationClaimSpecAllocation) contextValidateSelector(ctx context.Context, formats strfmt.Registry) error {

	if m.Selector != nil {
		if err := m.Selector.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("selector")
			}
			return err
		}
	}

	return nil
}

func (m *AllocationClaimSpecAllocationClaimSpecAllocation) contextValidateTemplate(ctx context.Context, formats strfmt.Registry) error {

	if m.Template != nil {
		if err := m.Template.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("template")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AllocationClaimSpecAllocationClaimSpecAllocation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AllocationClaimSpecAllocationClaimSpecAllocation) UnmarshalBinary(b []byte) error {
	var res AllocationClaimSpecAllocationClaimSpecAllocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

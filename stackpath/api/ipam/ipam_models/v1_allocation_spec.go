// Code generated by go-swagger; DO NOT EDIT.

package ipam_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1AllocationSpec v1 allocation spec
//
// swagger:model v1AllocationSpec
type V1AllocationSpec struct {

	// allocation class
	AllocationClass string `json:"allocationClass,omitempty"`

	// ip family
	IPFamily *StackpathschemanetworkIPFamily `json:"ipFamily,omitempty"`

	// prefix length
	PrefixLength int32 `json:"prefixLength,omitempty"`

	// reclaim policy
	ReclaimPolicy *V1ReclaimPolicy `json:"reclaimPolicy,omitempty"`

	// selectors
	Selectors []*Metav1MatchExpression `json:"selectors"`
}

// Validate validates this v1 allocation spec
func (m *V1AllocationSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPFamily(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReclaimPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelectors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AllocationSpec) validateIPFamily(formats strfmt.Registry) error {
	if swag.IsZero(m.IPFamily) { // not required
		return nil
	}

	if m.IPFamily != nil {
		if err := m.IPFamily.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipFamily")
			}
			return err
		}
	}

	return nil
}

func (m *V1AllocationSpec) validateReclaimPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.ReclaimPolicy) { // not required
		return nil
	}

	if m.ReclaimPolicy != nil {
		if err := m.ReclaimPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reclaimPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *V1AllocationSpec) validateSelectors(formats strfmt.Registry) error {
	if swag.IsZero(m.Selectors) { // not required
		return nil
	}

	for i := 0; i < len(m.Selectors); i++ {
		if swag.IsZero(m.Selectors[i]) { // not required
			continue
		}

		if m.Selectors[i] != nil {
			if err := m.Selectors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("selectors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1 allocation spec based on the context it is used
func (m *V1AllocationSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIPFamily(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReclaimPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelectors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AllocationSpec) contextValidateIPFamily(ctx context.Context, formats strfmt.Registry) error {

	if m.IPFamily != nil {
		if err := m.IPFamily.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipFamily")
			}
			return err
		}
	}

	return nil
}

func (m *V1AllocationSpec) contextValidateReclaimPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.ReclaimPolicy != nil {
		if err := m.ReclaimPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reclaimPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *V1AllocationSpec) contextValidateSelectors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Selectors); i++ {

		if m.Selectors[i] != nil {
			if err := m.Selectors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("selectors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1AllocationSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AllocationSpec) UnmarshalBinary(b []byte) error {
	var res V1AllocationSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

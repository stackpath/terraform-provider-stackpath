// Code generated by go-swagger; DO NOT EDIT.

package ipam_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1AllocationClaimSpec v1 allocation claim spec
//
// swagger:model v1AllocationClaimSpec
type V1AllocationClaimSpec struct {

	// allocation
	Allocation *AllocationClaimSpecAllocationClaimSpecAllocation `json:"allocation,omitempty"`

	// ip family
	IPFamily *StackpathschemanetworkIPFamily `json:"ipFamily,omitempty"`

	// prefix length
	PrefixLength int32 `json:"prefixLength,omitempty"`

	// reclaim policy
	ReclaimPolicy *V1ReclaimPolicy `json:"reclaimPolicy,omitempty"`

	// resource binding
	ResourceBinding *V1TypedResourceReference `json:"resourceBinding,omitempty"`
}

// Validate validates this v1 allocation claim spec
func (m *V1AllocationClaimSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPFamily(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReclaimPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceBinding(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AllocationClaimSpec) validateAllocation(formats strfmt.Registry) error {
	if swag.IsZero(m.Allocation) { // not required
		return nil
	}

	if m.Allocation != nil {
		if err := m.Allocation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("allocation")
			}
			return err
		}
	}

	return nil
}

func (m *V1AllocationClaimSpec) validateIPFamily(formats strfmt.Registry) error {
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

func (m *V1AllocationClaimSpec) validateReclaimPolicy(formats strfmt.Registry) error {
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

func (m *V1AllocationClaimSpec) validateResourceBinding(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourceBinding) { // not required
		return nil
	}

	if m.ResourceBinding != nil {
		if err := m.ResourceBinding.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourceBinding")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 allocation claim spec based on the context it is used
func (m *V1AllocationClaimSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAllocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIPFamily(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReclaimPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResourceBinding(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AllocationClaimSpec) contextValidateAllocation(ctx context.Context, formats strfmt.Registry) error {

	if m.Allocation != nil {
		if err := m.Allocation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("allocation")
			}
			return err
		}
	}

	return nil
}

func (m *V1AllocationClaimSpec) contextValidateIPFamily(ctx context.Context, formats strfmt.Registry) error {

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

func (m *V1AllocationClaimSpec) contextValidateReclaimPolicy(ctx context.Context, formats strfmt.Registry) error {

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

func (m *V1AllocationClaimSpec) contextValidateResourceBinding(ctx context.Context, formats strfmt.Registry) error {

	if m.ResourceBinding != nil {
		if err := m.ResourceBinding.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourceBinding")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1AllocationClaimSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AllocationClaimSpec) UnmarshalBinary(b []byte) error {
	var res V1AllocationClaimSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

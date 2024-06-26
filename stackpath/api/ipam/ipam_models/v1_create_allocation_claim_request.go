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

// V1CreateAllocationClaimRequest v1 create allocation claim request
//
// swagger:model v1CreateAllocationClaimRequest
type V1CreateAllocationClaimRequest struct {

	// allocation claim
	AllocationClaim *V1AllocationClaim `json:"allocationClaim,omitempty"`
}

// Validate validates this v1 create allocation claim request
func (m *V1CreateAllocationClaimRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllocationClaim(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1CreateAllocationClaimRequest) validateAllocationClaim(formats strfmt.Registry) error {
	if swag.IsZero(m.AllocationClaim) { // not required
		return nil
	}

	if m.AllocationClaim != nil {
		if err := m.AllocationClaim.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("allocationClaim")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 create allocation claim request based on the context it is used
func (m *V1CreateAllocationClaimRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAllocationClaim(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1CreateAllocationClaimRequest) contextValidateAllocationClaim(ctx context.Context, formats strfmt.Registry) error {

	if m.AllocationClaim != nil {
		if err := m.AllocationClaim.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("allocationClaim")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1CreateAllocationClaimRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CreateAllocationClaimRequest) UnmarshalBinary(b []byte) error {
	var res V1CreateAllocationClaimRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

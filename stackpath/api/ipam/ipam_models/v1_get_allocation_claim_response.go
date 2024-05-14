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

// V1GetAllocationClaimResponse v1 get allocation claim response
//
// swagger:model v1GetAllocationClaimResponse
type V1GetAllocationClaimResponse struct {

	// allocation claim
	AllocationClaim *V1AllocationClaim `json:"allocationClaim,omitempty"`
}

// Validate validates this v1 get allocation claim response
func (m *V1GetAllocationClaimResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllocationClaim(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1GetAllocationClaimResponse) validateAllocationClaim(formats strfmt.Registry) error {
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

// ContextValidate validate this v1 get allocation claim response based on the context it is used
func (m *V1GetAllocationClaimResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAllocationClaim(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1GetAllocationClaimResponse) contextValidateAllocationClaim(ctx context.Context, formats strfmt.Registry) error {

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
func (m *V1GetAllocationClaimResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1GetAllocationClaimResponse) UnmarshalBinary(b []byte) error {
	var res V1GetAllocationClaimResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

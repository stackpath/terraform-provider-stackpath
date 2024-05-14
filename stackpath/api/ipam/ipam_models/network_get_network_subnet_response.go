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

// NetworkGetNetworkSubnetResponse A response from a request to retrieve a specific a subnet
//
// swagger:model networkGetNetworkSubnetResponse
type NetworkGetNetworkSubnetResponse struct {

	// subnet
	Subnet *NetworkNetworkSubnet `json:"subnet,omitempty"`
}

// Validate validates this network get network subnet response
func (m *NetworkGetNetworkSubnetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSubnet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkGetNetworkSubnetResponse) validateSubnet(formats strfmt.Registry) error {
	if swag.IsZero(m.Subnet) { // not required
		return nil
	}

	if m.Subnet != nil {
		if err := m.Subnet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subnet")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this network get network subnet response based on the context it is used
func (m *NetworkGetNetworkSubnetResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSubnet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkGetNetworkSubnetResponse) contextValidateSubnet(ctx context.Context, formats strfmt.Registry) error {

	if m.Subnet != nil {
		if err := m.Subnet.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subnet")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkGetNetworkSubnetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkGetNetworkSubnetResponse) UnmarshalBinary(b []byte) error {
	var res NetworkGetNetworkSubnetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

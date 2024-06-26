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

// NetworkGetNetworkResponse A response from a request to retrieve an individual network
//
// swagger:model networkGetNetworkResponse
type NetworkGetNetworkResponse struct {

	// network
	Network *NetworkNetwork `json:"network,omitempty"`
}

// Validate validates this network get network response
func (m *NetworkGetNetworkResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNetwork(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkGetNetworkResponse) validateNetwork(formats strfmt.Registry) error {
	if swag.IsZero(m.Network) { // not required
		return nil
	}

	if m.Network != nil {
		if err := m.Network.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this network get network response based on the context it is used
func (m *NetworkGetNetworkResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNetwork(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkGetNetworkResponse) contextValidateNetwork(ctx context.Context, formats strfmt.Registry) error {

	if m.Network != nil {
		if err := m.Network.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkGetNetworkResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkGetNetworkResponse) UnmarshalBinary(b []byte) error {
	var res NetworkGetNetworkResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

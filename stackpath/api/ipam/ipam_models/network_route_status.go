// Code generated by go-swagger; DO NOT EDIT.

package ipam_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NetworkRouteStatus network route status
//
// swagger:model networkRouteStatus
type NetworkRouteStatus struct {

	// A list of route gateway IPs in the region
	// Read Only: true
	GatewayIps []string `json:"gatewayIps"`

	// The region in which a route's status resides
	//
	// Regions take the form `<pop name>-<cluster name>`, where `pop` is a StackPath EdgeCompute POP and `cluster` is a networking cluster within that POP.
	// Read Only: true
	Region string `json:"region,omitempty"`

	// state
	// Read Only: true
	State *RouteStatusState `json:"state,omitempty"`
}

// Validate validates this network route status
func (m *NetworkRouteStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkRouteStatus) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	if m.State != nil {
		if err := m.State.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this network route status based on the context it is used
func (m *NetworkRouteStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGatewayIps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRegion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkRouteStatus) contextValidateGatewayIps(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "gatewayIps", "body", []string(m.GatewayIps)); err != nil {
		return err
	}

	return nil
}

func (m *NetworkRouteStatus) contextValidateRegion(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "region", "body", string(m.Region)); err != nil {
		return err
	}

	return nil
}

func (m *NetworkRouteStatus) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if m.State != nil {
		if err := m.State.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkRouteStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkRouteStatus) UnmarshalBinary(b []byte) error {
	var res NetworkRouteStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
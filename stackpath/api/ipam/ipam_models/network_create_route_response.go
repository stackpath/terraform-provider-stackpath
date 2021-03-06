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

// NetworkCreateRouteResponse A response from a request to create a route on a network
//
// swagger:model networkCreateRouteResponse
type NetworkCreateRouteResponse struct {

	// route
	Route *NetworkRoute `json:"route,omitempty"`
}

// Validate validates this network create route response
func (m *NetworkCreateRouteResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRoute(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkCreateRouteResponse) validateRoute(formats strfmt.Registry) error {
	if swag.IsZero(m.Route) { // not required
		return nil
	}

	if m.Route != nil {
		if err := m.Route.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("route")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this network create route response based on the context it is used
func (m *NetworkCreateRouteResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRoute(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkCreateRouteResponse) contextValidateRoute(ctx context.Context, formats strfmt.Registry) error {

	if m.Route != nil {
		if err := m.Route.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("route")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkCreateRouteResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkCreateRouteResponse) UnmarshalBinary(b []byte) error {
	var res NetworkCreateRouteResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

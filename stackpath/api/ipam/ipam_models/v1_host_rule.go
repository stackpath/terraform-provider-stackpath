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

// V1HostRule Defines IPs, instances, or networks
//
// swagger:model v1HostRule
type V1HostRule struct {

	// List of instance selectors
	InstanceSelectors []*V1MatchExpression `json:"instanceSelectors"`

	// List of ip blocks
	IPBlock []*V1IPBlock `json:"ipBlock"`

	// List of network selectors
	NetworkSelectors []*V1MatchExpression `json:"networkSelectors"`
}

// Validate validates this v1 host rule
func (m *V1HostRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstanceSelectors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPBlock(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkSelectors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1HostRule) validateInstanceSelectors(formats strfmt.Registry) error {
	if swag.IsZero(m.InstanceSelectors) { // not required
		return nil
	}

	for i := 0; i < len(m.InstanceSelectors); i++ {
		if swag.IsZero(m.InstanceSelectors[i]) { // not required
			continue
		}

		if m.InstanceSelectors[i] != nil {
			if err := m.InstanceSelectors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("instanceSelectors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1HostRule) validateIPBlock(formats strfmt.Registry) error {
	if swag.IsZero(m.IPBlock) { // not required
		return nil
	}

	for i := 0; i < len(m.IPBlock); i++ {
		if swag.IsZero(m.IPBlock[i]) { // not required
			continue
		}

		if m.IPBlock[i] != nil {
			if err := m.IPBlock[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ipBlock" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1HostRule) validateNetworkSelectors(formats strfmt.Registry) error {
	if swag.IsZero(m.NetworkSelectors) { // not required
		return nil
	}

	for i := 0; i < len(m.NetworkSelectors); i++ {
		if swag.IsZero(m.NetworkSelectors[i]) { // not required
			continue
		}

		if m.NetworkSelectors[i] != nil {
			if err := m.NetworkSelectors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("networkSelectors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1 host rule based on the context it is used
func (m *V1HostRule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInstanceSelectors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIPBlock(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetworkSelectors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1HostRule) contextValidateInstanceSelectors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.InstanceSelectors); i++ {

		if m.InstanceSelectors[i] != nil {
			if err := m.InstanceSelectors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("instanceSelectors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1HostRule) contextValidateIPBlock(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IPBlock); i++ {

		if m.IPBlock[i] != nil {
			if err := m.IPBlock[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ipBlock" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1HostRule) contextValidateNetworkSelectors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NetworkSelectors); i++ {

		if m.NetworkSelectors[i] != nil {
			if err := m.NetworkSelectors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("networkSelectors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1HostRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1HostRule) UnmarshalBinary(b []byte) error {
	var res V1HostRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

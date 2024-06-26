// Code generated by go-swagger; DO NOT EDIT.

package workload_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1NetworkInterface Network interfaces that will be created on instances in the workload
//
// swagger:model v1NetworkInterface
type V1NetworkInterface struct {

	// assignments
	Assignments []*V1Assignment `json:"assignments"`

	// Whether to provide [one-to-one NAT](https://en.wikipedia.org/wiki/Network_address_translation#Basic_NAT) for this network interface
	//
	// This is an optional property used to enable or disable the NAT'ing the network interface. NAT is enabled by default on the first/primary interface and disabled on secondary/multi interfaces.
	//
	// Mark this property `false` to disable NAT on the first interface. Mark this property `true` to enable NAT on secondary interfaces.
	//
	// The `ExcludeNAT` workload annotation supercedes this property.
	EnableOneToOneNat bool `json:"enableOneToOneNat"`

	// A list of IP families to use for interface ip assignments
	//
	// This is an optional property and supports ['IPv4'] or ['IPv4', 'IPv6'] list
	IPFamilies []*V1IPFamily `json:"ipFamilies"`

	// An IPv6 subnet interface's slug. This is an optional property used to attach a specific network interface to a ipv6 subnet.
	IPV6Subnet string `json:"ipv6Subnet,omitempty"`

	// A network interface's slug
	Network string `json:"network,omitempty"`

	// An IPv4 subnet interface's slug. This is an optional property used to attach a specific network interface to a IPv4 subnet.
	Subnet string `json:"subnet,omitempty"`
}

// Validate validates this v1 network interface
func (m *V1NetworkInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssignments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPFamilies(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1NetworkInterface) validateAssignments(formats strfmt.Registry) error {
	if swag.IsZero(m.Assignments) { // not required
		return nil
	}

	for i := 0; i < len(m.Assignments); i++ {
		if swag.IsZero(m.Assignments[i]) { // not required
			continue
		}

		if m.Assignments[i] != nil {
			if err := m.Assignments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("assignments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1NetworkInterface) validateIPFamilies(formats strfmt.Registry) error {
	if swag.IsZero(m.IPFamilies) { // not required
		return nil
	}

	for i := 0; i < len(m.IPFamilies); i++ {
		if swag.IsZero(m.IPFamilies[i]) { // not required
			continue
		}

		if m.IPFamilies[i] != nil {
			if err := m.IPFamilies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ipFamilies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1 network interface based on the context it is used
func (m *V1NetworkInterface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAssignments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIPFamilies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1NetworkInterface) contextValidateAssignments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Assignments); i++ {

		if m.Assignments[i] != nil {
			if err := m.Assignments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("assignments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1NetworkInterface) contextValidateIPFamilies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IPFamilies); i++ {

		if m.IPFamilies[i] != nil {
			if err := m.IPFamilies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ipFamilies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1NetworkInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1NetworkInterface) UnmarshalBinary(b []byte) error {
	var res V1NetworkInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

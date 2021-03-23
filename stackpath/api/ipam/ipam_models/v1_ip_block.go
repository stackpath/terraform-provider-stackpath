// Code generated by go-swagger; DO NOT EDIT.

package ipam_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1IPBlock Defines an IP block
//
// swagger:model v1IpBlock
type V1IPBlock struct {

	// A subnet that will define all the IPs allowed by a rule
	Cidr string `json:"cidr,omitempty"`

	// A list of subnets that will be excluded from the above subnet
	//
	// This allows a convenient way to allow multiple ip ranges in a single expression
	Except []string `json:"except"`
}

// Validate validates this v1 Ip block
func (m *V1IPBlock) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 Ip block based on context it is used
func (m *V1IPBlock) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1IPBlock) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1IPBlock) UnmarshalBinary(b []byte) error {
	var res V1IPBlock
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

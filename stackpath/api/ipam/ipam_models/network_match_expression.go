// Code generated by go-swagger; DO NOT EDIT.

package ipam_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NetworkMatchExpression An expression to match selectors against a set of values
//
// swagger:model networkMatchExpression
type NetworkMatchExpression struct {

	// The name of the selector to perform a match against
	//
	// Provide the key `workload.platform.stackpath.net/workload-slug` to select EdgeCompute workloads by their slug. All network interfaces in the workload are used as the route's gateway.
	Key string `json:"key,omitempty"`

	// The operation to perform to match a selector
	//
	// Valid values are "in" with support for more possible in the future
	Operator string `json:"operator,omitempty"`

	// The values to match in the selector
	Values []string `json:"values"`
}

// Validate validates this network match expression
func (m *NetworkMatchExpression) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this network match expression based on context it is used
func (m *NetworkMatchExpression) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NetworkMatchExpression) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkMatchExpression) UnmarshalBinary(b []byte) error {
	var res NetworkMatchExpression
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

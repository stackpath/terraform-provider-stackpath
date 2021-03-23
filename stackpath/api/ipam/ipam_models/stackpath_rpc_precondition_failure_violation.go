// Code generated by go-swagger; DO NOT EDIT.

package ipam_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StackpathRPCPreconditionFailureViolation stackpath rpc precondition failure violation
//
// swagger:model stackpath.rpc.PreconditionFailure.Violation
type StackpathRPCPreconditionFailureViolation struct {

	// description
	Description string `json:"description,omitempty"`

	// subject
	Subject string `json:"subject,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this stackpath rpc precondition failure violation
func (m *StackpathRPCPreconditionFailureViolation) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this stackpath rpc precondition failure violation based on context it is used
func (m *StackpathRPCPreconditionFailureViolation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StackpathRPCPreconditionFailureViolation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StackpathRPCPreconditionFailureViolation) UnmarshalBinary(b []byte) error {
	var res StackpathRPCPreconditionFailureViolation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

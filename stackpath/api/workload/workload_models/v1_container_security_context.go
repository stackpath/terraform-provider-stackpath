// Code generated by go-swagger; DO NOT EDIT.

package workload_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ContainerSecurityContext Security configuration that will be applied to a container
//
// swagger:model v1ContainerSecurityContext
type V1ContainerSecurityContext struct {

	// Controls whether a process can gain more privileges than its parent process
	AllowPrivilegeEscalation bool `json:"allowPrivilegeEscalation,omitempty"`

	// capabilities
	Capabilities *V1ContainerCapabilities `json:"capabilities,omitempty"`

	// Indicates whether this container has a read-only root filesystem
	ReadOnlyRootFilesystem bool `json:"readOnlyRootFilesystem,omitempty"`

	// The user GID to run the entry point of the container process
	RunAsGroup string `json:"runAsGroup,omitempty"`

	// Indicates that the container must run as a non-root user
	RunAsNonRoot bool `json:"runAsNonRoot,omitempty"`

	// The user UID to run the entry point of the container process
	RunAsUser string `json:"runAsUser,omitempty"`
}

// Validate validates this v1 container security context
func (m *V1ContainerSecurityContext) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCapabilities(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ContainerSecurityContext) validateCapabilities(formats strfmt.Registry) error {
	if swag.IsZero(m.Capabilities) { // not required
		return nil
	}

	if m.Capabilities != nil {
		if err := m.Capabilities.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("capabilities")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 container security context based on the context it is used
func (m *V1ContainerSecurityContext) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCapabilities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ContainerSecurityContext) contextValidateCapabilities(ctx context.Context, formats strfmt.Registry) error {

	if m.Capabilities != nil {

		if swag.IsZero(m.Capabilities) { // not required
			return nil
		}

		if err := m.Capabilities.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("capabilities")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ContainerSecurityContext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ContainerSecurityContext) UnmarshalBinary(b []byte) error {
	var res V1ContainerSecurityContext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

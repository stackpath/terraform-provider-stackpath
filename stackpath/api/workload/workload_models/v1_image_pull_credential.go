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

// V1ImagePullCredential The credentials that should be used to pull the container image
//
// swagger:model v1ImagePullCredential
type V1ImagePullCredential struct {

	// docker registry
	DockerRegistry *V1DockerRegistryCredentials `json:"dockerRegistry,omitempty"`
}

// Validate validates this v1 image pull credential
func (m *V1ImagePullCredential) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDockerRegistry(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ImagePullCredential) validateDockerRegistry(formats strfmt.Registry) error {
	if swag.IsZero(m.DockerRegistry) { // not required
		return nil
	}

	if m.DockerRegistry != nil {
		if err := m.DockerRegistry.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dockerRegistry")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 image pull credential based on the context it is used
func (m *V1ImagePullCredential) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDockerRegistry(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ImagePullCredential) contextValidateDockerRegistry(ctx context.Context, formats strfmt.Registry) error {

	if m.DockerRegistry != nil {

		if swag.IsZero(m.DockerRegistry) { // not required
			return nil
		}

		if err := m.DockerRegistry.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dockerRegistry")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ImagePullCredential) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ImagePullCredential) UnmarshalBinary(b []byte) error {
	var res V1ImagePullCredential
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

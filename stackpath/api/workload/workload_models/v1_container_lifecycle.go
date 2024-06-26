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

// V1ContainerLifecycle Defines actions that the management system should take in response to container lifecycle events
//
// swagger:model v1ContainerLifecycle
type V1ContainerLifecycle struct {

	// post start
	PostStart *V1ContainerLifecycleHandler `json:"postStart,omitempty"`

	// pre stop
	PreStop *V1ContainerLifecycleHandler `json:"preStop,omitempty"`
}

// Validate validates this v1 container lifecycle
func (m *V1ContainerLifecycle) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePostStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreStop(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ContainerLifecycle) validatePostStart(formats strfmt.Registry) error {
	if swag.IsZero(m.PostStart) { // not required
		return nil
	}

	if m.PostStart != nil {
		if err := m.PostStart.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postStart")
			}
			return err
		}
	}

	return nil
}

func (m *V1ContainerLifecycle) validatePreStop(formats strfmt.Registry) error {
	if swag.IsZero(m.PreStop) { // not required
		return nil
	}

	if m.PreStop != nil {
		if err := m.PreStop.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("preStop")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 container lifecycle based on the context it is used
func (m *V1ContainerLifecycle) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePostStart(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePreStop(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ContainerLifecycle) contextValidatePostStart(ctx context.Context, formats strfmt.Registry) error {

	if m.PostStart != nil {
		if err := m.PostStart.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postStart")
			}
			return err
		}
	}

	return nil
}

func (m *V1ContainerLifecycle) contextValidatePreStop(ctx context.Context, formats strfmt.Registry) error {

	if m.PreStop != nil {
		if err := m.PreStop.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("preStop")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ContainerLifecycle) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ContainerLifecycle) UnmarshalBinary(b []byte) error {
	var res V1ContainerLifecycle
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

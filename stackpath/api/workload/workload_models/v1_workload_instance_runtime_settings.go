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

// V1WorkloadInstanceRuntimeSettings The specification for workload instance level settings
//
// swagger:model v1WorkloadInstanceRuntimeSettings
type V1WorkloadInstanceRuntimeSettings struct {

	// containers
	Containers *V1WorkloadInstanceContainerRuntimeSettings `json:"containers,omitempty"`

	// virtual machines
	VirtualMachines *V1WorkloadInstanceVMRuntimeSettings `json:"virtualMachines,omitempty"`
}

// Validate validates this v1 workload instance runtime settings
func (m *V1WorkloadInstanceRuntimeSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContainers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualMachines(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1WorkloadInstanceRuntimeSettings) validateContainers(formats strfmt.Registry) error {
	if swag.IsZero(m.Containers) { // not required
		return nil
	}

	if m.Containers != nil {
		if err := m.Containers.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("containers")
			}
			return err
		}
	}

	return nil
}

func (m *V1WorkloadInstanceRuntimeSettings) validateVirtualMachines(formats strfmt.Registry) error {
	if swag.IsZero(m.VirtualMachines) { // not required
		return nil
	}

	if m.VirtualMachines != nil {
		if err := m.VirtualMachines.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtualMachines")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 workload instance runtime settings based on the context it is used
func (m *V1WorkloadInstanceRuntimeSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContainers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVirtualMachines(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1WorkloadInstanceRuntimeSettings) contextValidateContainers(ctx context.Context, formats strfmt.Registry) error {

	if m.Containers != nil {
		if err := m.Containers.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("containers")
			}
			return err
		}
	}

	return nil
}

func (m *V1WorkloadInstanceRuntimeSettings) contextValidateVirtualMachines(ctx context.Context, formats strfmt.Registry) error {

	if m.VirtualMachines != nil {
		if err := m.VirtualMachines.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtualMachines")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1WorkloadInstanceRuntimeSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1WorkloadInstanceRuntimeSettings) UnmarshalBinary(b []byte) error {
	var res V1WorkloadInstanceRuntimeSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

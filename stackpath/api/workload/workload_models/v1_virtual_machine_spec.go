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

// V1VirtualMachineSpec The specification for the desired state of a virtual machine in a workload
//
// swagger:model v1VirtualMachineSpec
type V1VirtualMachineSpec struct {

	// The image to use for the virtual machine
	//
	// This is in the format of <stack-slug>/<image-family>[:<image-tag>]. If the image tag portion is omitted, 'default' is assumed which is the most recently created, ready, and non-deprecated image of that slug. A set of common images is present on the 'stackpath-edge' stack.
	Image string `json:"image,omitempty"`

	// liveness probe
	LivenessProbe *V1Probe `json:"livenessProbe,omitempty"`

	// ports
	Ports V1InstancePortMapEntry `json:"ports,omitempty"`

	// readiness probe
	ReadinessProbe *V1Probe `json:"readinessProbe,omitempty"`

	// resources
	Resources *V1ResourceRequirements `json:"resources,omitempty"`

	// Base64 encoded cloud-init compatible user-data
	UserData string `json:"userData,omitempty"`

	// Volumes to mount in the virtual machine
	VolumeMounts []*V1InstanceVolumeMount `json:"volumeMounts"`
}

// Validate validates this v1 virtual machine spec
func (m *V1VirtualMachineSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLivenessProbe(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePorts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReadinessProbe(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumeMounts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VirtualMachineSpec) validateLivenessProbe(formats strfmt.Registry) error {
	if swag.IsZero(m.LivenessProbe) { // not required
		return nil
	}

	if m.LivenessProbe != nil {
		if err := m.LivenessProbe.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("livenessProbe")
			}
			return err
		}
	}

	return nil
}

func (m *V1VirtualMachineSpec) validatePorts(formats strfmt.Registry) error {
	if swag.IsZero(m.Ports) { // not required
		return nil
	}

	if m.Ports != nil {
		if err := m.Ports.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ports")
			}
			return err
		}
	}

	return nil
}

func (m *V1VirtualMachineSpec) validateReadinessProbe(formats strfmt.Registry) error {
	if swag.IsZero(m.ReadinessProbe) { // not required
		return nil
	}

	if m.ReadinessProbe != nil {
		if err := m.ReadinessProbe.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("readinessProbe")
			}
			return err
		}
	}

	return nil
}

func (m *V1VirtualMachineSpec) validateResources(formats strfmt.Registry) error {
	if swag.IsZero(m.Resources) { // not required
		return nil
	}

	if m.Resources != nil {
		if err := m.Resources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resources")
			}
			return err
		}
	}

	return nil
}

func (m *V1VirtualMachineSpec) validateVolumeMounts(formats strfmt.Registry) error {
	if swag.IsZero(m.VolumeMounts) { // not required
		return nil
	}

	for i := 0; i < len(m.VolumeMounts); i++ {
		if swag.IsZero(m.VolumeMounts[i]) { // not required
			continue
		}

		if m.VolumeMounts[i] != nil {
			if err := m.VolumeMounts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumeMounts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1 virtual machine spec based on the context it is used
func (m *V1VirtualMachineSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLivenessProbe(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePorts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReadinessProbe(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVolumeMounts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VirtualMachineSpec) contextValidateLivenessProbe(ctx context.Context, formats strfmt.Registry) error {

	if m.LivenessProbe != nil {
		if err := m.LivenessProbe.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("livenessProbe")
			}
			return err
		}
	}

	return nil
}

func (m *V1VirtualMachineSpec) contextValidatePorts(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Ports.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ports")
		}
		return err
	}

	return nil
}

func (m *V1VirtualMachineSpec) contextValidateReadinessProbe(ctx context.Context, formats strfmt.Registry) error {

	if m.ReadinessProbe != nil {
		if err := m.ReadinessProbe.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("readinessProbe")
			}
			return err
		}
	}

	return nil
}

func (m *V1VirtualMachineSpec) contextValidateResources(ctx context.Context, formats strfmt.Registry) error {

	if m.Resources != nil {
		if err := m.Resources.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resources")
			}
			return err
		}
	}

	return nil
}

func (m *V1VirtualMachineSpec) contextValidateVolumeMounts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VolumeMounts); i++ {

		if m.VolumeMounts[i] != nil {
			if err := m.VolumeMounts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumeMounts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1VirtualMachineSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VirtualMachineSpec) UnmarshalBinary(b []byte) error {
	var res V1VirtualMachineSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

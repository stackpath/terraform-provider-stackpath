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

// V1ContainerSpec The specification for the desired state of a container in a workload
//
// swagger:model v1ContainerSpec
type V1ContainerSpec struct {

	// The commands that start a container
	Command []string `json:"command"`

	// env
	Env V1EnvironmentVariableMapEntry `json:"env,omitempty"`

	// The location of a Docker image to run as a container
	Image string `json:"image,omitempty"`

	// liveness probe
	LivenessProbe *V1Probe `json:"livenessProbe,omitempty"`

	// ports
	Ports V1InstancePortMapEntry `json:"ports,omitempty"`

	// readiness probe
	ReadinessProbe *V1Probe `json:"readinessProbe,omitempty"`

	// resources
	Resources *V1ResourceRequirements `json:"resources,omitempty"`

	// Volumes to mount in the container
	VolumeMounts []*V1InstanceVolumeMount `json:"volumeMounts"`
}

// Validate validates this v1 container spec
func (m *V1ContainerSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnv(formats); err != nil {
		res = append(res, err)
	}

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

func (m *V1ContainerSpec) validateEnv(formats strfmt.Registry) error {
	if swag.IsZero(m.Env) { // not required
		return nil
	}

	if m.Env != nil {
		if err := m.Env.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("env")
			}
			return err
		}
	}

	return nil
}

func (m *V1ContainerSpec) validateLivenessProbe(formats strfmt.Registry) error {
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

func (m *V1ContainerSpec) validatePorts(formats strfmt.Registry) error {
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

func (m *V1ContainerSpec) validateReadinessProbe(formats strfmt.Registry) error {
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

func (m *V1ContainerSpec) validateResources(formats strfmt.Registry) error {
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

func (m *V1ContainerSpec) validateVolumeMounts(formats strfmt.Registry) error {
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

// ContextValidate validate this v1 container spec based on the context it is used
func (m *V1ContainerSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEnv(ctx, formats); err != nil {
		res = append(res, err)
	}

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

func (m *V1ContainerSpec) contextValidateEnv(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Env.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("env")
		}
		return err
	}

	return nil
}

func (m *V1ContainerSpec) contextValidateLivenessProbe(ctx context.Context, formats strfmt.Registry) error {

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

func (m *V1ContainerSpec) contextValidatePorts(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Ports.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ports")
		}
		return err
	}

	return nil
}

func (m *V1ContainerSpec) contextValidateReadinessProbe(ctx context.Context, formats strfmt.Registry) error {

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

func (m *V1ContainerSpec) contextValidateResources(ctx context.Context, formats strfmt.Registry) error {

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

func (m *V1ContainerSpec) contextValidateVolumeMounts(ctx context.Context, formats strfmt.Registry) error {

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
func (m *V1ContainerSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ContainerSpec) UnmarshalBinary(b []byte) error {
	var res V1ContainerSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

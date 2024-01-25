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

// V1ContainerLifecycleHandler Defines a specific lifecycle management action that should be taken
//
// swagger:model v1ContainerLifecycleHandler
type V1ContainerLifecycleHandler struct {

	// exec
	Exec *V1ExecAction `json:"exec,omitempty"`

	// http get
	HTTPGet *V1HTTPGetAction `json:"httpGet,omitempty"`

	// tcp socket
	TCPSocket *V1TCPSocketAction `json:"tcpSocket,omitempty"`
}

// Validate validates this v1 container lifecycle handler
func (m *V1ContainerLifecycleHandler) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPGet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTCPSocket(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ContainerLifecycleHandler) validateExec(formats strfmt.Registry) error {
	if swag.IsZero(m.Exec) { // not required
		return nil
	}

	if m.Exec != nil {
		if err := m.Exec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("exec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce
			}
			return err
		}
	}

	return nil
}

func (m *V1ContainerLifecycleHandler) validateHTTPGet(formats strfmt.Registry) error {
	if swag.IsZero(m.HTTPGet) { // not required
		return nil
	}

	if m.HTTPGet != nil {
		if err := m.HTTPGet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("httpGet")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce
			}
			return err
		}
	}

	return nil
}

func (m *V1ContainerLifecycleHandler) validateTCPSocket(formats strfmt.Registry) error {
	if swag.IsZero(m.TCPSocket) { // not required
		return nil
	}

	if m.TCPSocket != nil {
		if err := m.TCPSocket.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tcpSocket")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 container lifecycle handler based on the context it is used
func (m *V1ContainerLifecycleHandler) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHTTPGet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTCPSocket(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ContainerLifecycleHandler) contextValidateExec(ctx context.Context, formats strfmt.Registry) error {

	if m.Exec != nil {

		if swag.IsZero(m.Exec) { // not required
			return nil
		}

		if err := m.Exec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("exec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce
			}
			return err
		}
	}

	return nil
}

func (m *V1ContainerLifecycleHandler) contextValidateHTTPGet(ctx context.Context, formats strfmt.Registry) error {

	if m.HTTPGet != nil {

		if swag.IsZero(m.HTTPGet) { // not required
			return nil
		}

		if err := m.HTTPGet.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("httpGet")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce
			}
			return err
		}
	}

	return nil
}

func (m *V1ContainerLifecycleHandler) contextValidateTCPSocket(ctx context.Context, formats strfmt.Registry) error {

	if m.TCPSocket != nil {

		if swag.IsZero(m.TCPSocket) { // not required
			return nil
		}

		if err := m.TCPSocket.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tcpSocket")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ContainerLifecycleHandler) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ContainerLifecycleHandler) UnmarshalBinary(b []byte) error {
	var res V1ContainerLifecycleHandler
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
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

// V1PutWorkloadRequest v1 put workload request
//
// swagger:model v1PutWorkloadRequest
type V1PutWorkloadRequest struct {

	// workload
	Workload *V1Workload `json:"workload,omitempty"`
}

// Validate validates this v1 put workload request
func (m *V1PutWorkloadRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWorkload(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PutWorkloadRequest) validateWorkload(formats strfmt.Registry) error {
	if swag.IsZero(m.Workload) { // not required
		return nil
	}

	if m.Workload != nil {
		if err := m.Workload.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workload")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 put workload request based on the context it is used
func (m *V1PutWorkloadRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateWorkload(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PutWorkloadRequest) contextValidateWorkload(ctx context.Context, formats strfmt.Registry) error {

	if m.Workload != nil {
		if err := m.Workload.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workload")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1PutWorkloadRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PutWorkloadRequest) UnmarshalBinary(b []byte) error {
	var res V1PutWorkloadRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

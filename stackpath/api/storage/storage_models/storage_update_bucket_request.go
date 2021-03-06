// Code generated by go-swagger; DO NOT EDIT.

package storage_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StorageUpdateBucketRequest storage update bucket request
//
// swagger:model storageUpdateBucketRequest
type StorageUpdateBucketRequest struct {

	// visibility
	Visibility *StorageBucketVisibility `json:"visibility,omitempty"`
}

// Validate validates this storage update bucket request
func (m *StorageUpdateBucketRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVisibility(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StorageUpdateBucketRequest) validateVisibility(formats strfmt.Registry) error {
	if swag.IsZero(m.Visibility) { // not required
		return nil
	}

	if m.Visibility != nil {
		if err := m.Visibility.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("visibility")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this storage update bucket request based on the context it is used
func (m *StorageUpdateBucketRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVisibility(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StorageUpdateBucketRequest) contextValidateVisibility(ctx context.Context, formats strfmt.Registry) error {

	if m.Visibility != nil {
		if err := m.Visibility.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("visibility")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StorageUpdateBucketRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageUpdateBucketRequest) UnmarshalBinary(b []byte) error {
	var res StorageUpdateBucketRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

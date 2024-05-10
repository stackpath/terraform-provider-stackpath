// Code generated by go-swagger; DO NOT EDIT.

package workload_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1Metadata Metadata associated with an entity
//
// swagger:model v1Metadata
type V1Metadata struct {

	// annotations
	Annotations V1StringMapEntry `json:"annotations,omitempty"`

	// The date that a metadata entry was created
	// Read Only: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"createdAt,omitempty"`

	// The date an entity was requested to be deleted
	// Read Only: true
	// Format: date-time
	DeleteRequestedAt *strfmt.DateTime `json:"deleteRequestedAt,omitempty"`

	// labels
	Labels V1StringMapEntry `json:"labels,omitempty"`

	// The date that a metadata entry was last updated
	// Read Only: true
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updatedAt,omitempty"`

	// A metadata entry's version number
	//
	// Metadata versions start at 1 when they are created and increment by 1 every time they are updated.
	Version string `json:"version,omitempty"`
}

// Validate validates this v1 metadata
func (m *V1Metadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAnnotations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeleteRequestedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Metadata) validateAnnotations(formats strfmt.Registry) error {
	if swag.IsZero(m.Annotations) { // not required
		return nil
	}

	if m.Annotations != nil {
		if err := m.Annotations.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("annotations")
			}
			return err
		}
	}

	return nil
}

func (m *V1Metadata) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1Metadata) validateDeleteRequestedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.DeleteRequestedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("deleteRequestedAt", "body", "date-time", m.DeleteRequestedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1Metadata) validateLabels(formats strfmt.Registry) error {
	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	if m.Labels != nil {
		if err := m.Labels.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("labels")
			}
			return err
		}
	}

	return nil
}

func (m *V1Metadata) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 metadata based on the context it is used
func (m *V1Metadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAnnotations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeleteRequestedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Metadata) contextValidateAnnotations(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Annotations.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("annotations")
		}
		return err
	}

	return nil
}

func (m *V1Metadata) contextValidateCreatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "createdAt", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *V1Metadata) contextValidateDeleteRequestedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "deleteRequestedAt", "body", m.DeleteRequestedAt); err != nil {
		return err
	}

	return nil
}

func (m *V1Metadata) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Labels.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("labels")
		}
		return err
	}

	return nil
}

func (m *V1Metadata) contextValidateUpdatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "updatedAt", "body", m.UpdatedAt); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Metadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Metadata) UnmarshalBinary(b []byte) error {
	var res V1Metadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

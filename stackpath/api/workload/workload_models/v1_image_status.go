// Code generated by go-swagger; DO NOT EDIT.

package workload_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// V1ImageStatus Which capture status an image is currently in
//
// - IMAGE_STATUS_UNKNOWN: The image status is unknown
//  - PENDING: The image is pending creation
//  - PROCESSING: The image is processing
//  - READY: The image is ready
//  - FAILED: The image failed to be created
//
// swagger:model v1ImageStatus
type V1ImageStatus string

func NewV1ImageStatus(value V1ImageStatus) *V1ImageStatus {
	v := value
	return &v
}

const (

	// V1ImageStatusIMAGESTATUSUNKNOWN captures enum value "IMAGE_STATUS_UNKNOWN"
	V1ImageStatusIMAGESTATUSUNKNOWN V1ImageStatus = "IMAGE_STATUS_UNKNOWN"

	// V1ImageStatusPENDING captures enum value "PENDING"
	V1ImageStatusPENDING V1ImageStatus = "PENDING"

	// V1ImageStatusPROCESSING captures enum value "PROCESSING"
	V1ImageStatusPROCESSING V1ImageStatus = "PROCESSING"

	// V1ImageStatusREADY captures enum value "READY"
	V1ImageStatusREADY V1ImageStatus = "READY"

	// V1ImageStatusFAILED captures enum value "FAILED"
	V1ImageStatusFAILED V1ImageStatus = "FAILED"
)

// for schema
var v1ImageStatusEnum []interface{}

func init() {
	var res []V1ImageStatus
	if err := json.Unmarshal([]byte(`["IMAGE_STATUS_UNKNOWN","PENDING","PROCESSING","READY","FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1ImageStatusEnum = append(v1ImageStatusEnum, v)
	}
}

func (m V1ImageStatus) validateV1ImageStatusEnum(path, location string, value V1ImageStatus) error {
	if err := validate.EnumCase(path, location, value, v1ImageStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 image status
func (m V1ImageStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1ImageStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this v1 image status based on context it is used
func (m V1ImageStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

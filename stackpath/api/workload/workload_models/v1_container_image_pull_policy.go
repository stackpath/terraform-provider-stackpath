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

// V1ContainerImagePullPolicy The image pull policy for the container image. The value should be one of ALWAYS or IF_NOT_PRESENT
//
// swagger:model v1ContainerImagePullPolicy
type V1ContainerImagePullPolicy string

func NewV1ContainerImagePullPolicy(value V1ContainerImagePullPolicy) *V1ContainerImagePullPolicy {
	return &value
}

// Pointer returns a pointer to a freshly-allocated V1ContainerImagePullPolicy.
func (m V1ContainerImagePullPolicy) Pointer() *V1ContainerImagePullPolicy {
	return &m
}

const (

	// V1ContainerImagePullPolicyCONTAINERIMAGEPULLPOLICYUNSPECIFIED captures enum value "CONTAINER_IMAGE_PULL_POLICY_UNSPECIFIED"
	V1ContainerImagePullPolicyCONTAINERIMAGEPULLPOLICYUNSPECIFIED V1ContainerImagePullPolicy = "CONTAINER_IMAGE_PULL_POLICY_UNSPECIFIED"

	// V1ContainerImagePullPolicyALWAYS captures enum value "ALWAYS"
	V1ContainerImagePullPolicyALWAYS V1ContainerImagePullPolicy = "ALWAYS"

	// V1ContainerImagePullPolicyIFNOTPRESENT captures enum value "IF_NOT_PRESENT"
	V1ContainerImagePullPolicyIFNOTPRESENT V1ContainerImagePullPolicy = "IF_NOT_PRESENT"
)

// for schema
var v1ContainerImagePullPolicyEnum []interface{}

func init() {
	var res []V1ContainerImagePullPolicy
	if err := json.Unmarshal([]byte(`["CONTAINER_IMAGE_PULL_POLICY_UNSPECIFIED","ALWAYS","IF_NOT_PRESENT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1ContainerImagePullPolicyEnum = append(v1ContainerImagePullPolicyEnum, v)
	}
}

func (m V1ContainerImagePullPolicy) validateV1ContainerImagePullPolicyEnum(path, location string, value V1ContainerImagePullPolicy) error {
	if err := validate.EnumCase(path, location, value, v1ContainerImagePullPolicyEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 container image pull policy
func (m V1ContainerImagePullPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1ContainerImagePullPolicyEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this v1 container image pull policy based on context it is used
func (m V1ContainerImagePullPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
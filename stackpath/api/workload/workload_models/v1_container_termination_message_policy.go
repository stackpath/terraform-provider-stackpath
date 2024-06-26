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

// V1ContainerTerminationMessagePolicy The policy indicating how the termination message should be populated
//
// swagger:model v1ContainerTerminationMessagePolicy
type V1ContainerTerminationMessagePolicy string

func NewV1ContainerTerminationMessagePolicy(value V1ContainerTerminationMessagePolicy) *V1ContainerTerminationMessagePolicy {
	v := value
	return &v
}

const (

	// V1ContainerTerminationMessagePolicyCONTAINERTERMINATIONMESSAGEPOLICYUNSPECIFIED captures enum value "CONTAINER_TERMINATION_MESSAGE_POLICY_UNSPECIFIED"
	V1ContainerTerminationMessagePolicyCONTAINERTERMINATIONMESSAGEPOLICYUNSPECIFIED V1ContainerTerminationMessagePolicy = "CONTAINER_TERMINATION_MESSAGE_POLICY_UNSPECIFIED"

	// V1ContainerTerminationMessagePolicyFILE captures enum value "FILE"
	V1ContainerTerminationMessagePolicyFILE V1ContainerTerminationMessagePolicy = "FILE"

	// V1ContainerTerminationMessagePolicyFALLBACKTOLOGSONERROR captures enum value "FALLBACK_TO_LOGS_ON_ERROR"
	V1ContainerTerminationMessagePolicyFALLBACKTOLOGSONERROR V1ContainerTerminationMessagePolicy = "FALLBACK_TO_LOGS_ON_ERROR"
)

// for schema
var v1ContainerTerminationMessagePolicyEnum []interface{}

func init() {
	var res []V1ContainerTerminationMessagePolicy
	if err := json.Unmarshal([]byte(`["CONTAINER_TERMINATION_MESSAGE_POLICY_UNSPECIFIED","FILE","FALLBACK_TO_LOGS_ON_ERROR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1ContainerTerminationMessagePolicyEnum = append(v1ContainerTerminationMessagePolicyEnum, v)
	}
}

func (m V1ContainerTerminationMessagePolicy) validateV1ContainerTerminationMessagePolicyEnum(path, location string, value V1ContainerTerminationMessagePolicy) error {
	if err := validate.EnumCase(path, location, value, v1ContainerTerminationMessagePolicyEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 container termination message policy
func (m V1ContainerTerminationMessagePolicy) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1ContainerTerminationMessagePolicyEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this v1 container termination message policy based on context it is used
func (m V1ContainerTerminationMessagePolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

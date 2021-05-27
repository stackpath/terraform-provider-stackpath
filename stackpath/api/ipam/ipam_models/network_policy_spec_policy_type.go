// Code generated by go-swagger; DO NOT EDIT.

package ipam_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NetworkPolicySpecPolicyType network policy spec policy type
//
// swagger:model NetworkPolicySpecPolicyType
type NetworkPolicySpecPolicyType string

func NewNetworkPolicySpecPolicyType(value NetworkPolicySpecPolicyType) *NetworkPolicySpecPolicyType {
	v := value
	return &v
}

const (

	// NetworkPolicySpecPolicyTypePOLICYTYPENOTSPECIFIED captures enum value "POLICY_TYPE_NOT_SPECIFIED"
	NetworkPolicySpecPolicyTypePOLICYTYPENOTSPECIFIED NetworkPolicySpecPolicyType = "POLICY_TYPE_NOT_SPECIFIED"

	// NetworkPolicySpecPolicyTypeINGRESS captures enum value "INGRESS"
	NetworkPolicySpecPolicyTypeINGRESS NetworkPolicySpecPolicyType = "INGRESS"

	// NetworkPolicySpecPolicyTypeEGRESS captures enum value "EGRESS"
	NetworkPolicySpecPolicyTypeEGRESS NetworkPolicySpecPolicyType = "EGRESS"
)

// for schema
var networkPolicySpecPolicyTypeEnum []interface{}

func init() {
	var res []NetworkPolicySpecPolicyType
	if err := json.Unmarshal([]byte(`["POLICY_TYPE_NOT_SPECIFIED","INGRESS","EGRESS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		networkPolicySpecPolicyTypeEnum = append(networkPolicySpecPolicyTypeEnum, v)
	}
}

func (m NetworkPolicySpecPolicyType) validateNetworkPolicySpecPolicyTypeEnum(path, location string, value NetworkPolicySpecPolicyType) error {
	if err := validate.EnumCase(path, location, value, networkPolicySpecPolicyTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this network policy spec policy type
func (m NetworkPolicySpecPolicyType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateNetworkPolicySpecPolicyTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this network policy spec policy type based on context it is used
func (m NetworkPolicySpecPolicyType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

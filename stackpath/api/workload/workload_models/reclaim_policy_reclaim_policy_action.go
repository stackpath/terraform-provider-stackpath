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

// ReclaimPolicyReclaimPolicyAction reclaim policy reclaim policy action
//
// swagger:model ReclaimPolicyReclaimPolicyAction
type ReclaimPolicyReclaimPolicyAction string

func NewReclaimPolicyReclaimPolicyAction(value ReclaimPolicyReclaimPolicyAction) *ReclaimPolicyReclaimPolicyAction {
	v := value
	return &v
}

const (

	// ReclaimPolicyReclaimPolicyActionRECLAIMPOLICYACTIONUNSPECIFIED captures enum value "RECLAIM_POLICY_ACTION_UNSPECIFIED"
	ReclaimPolicyReclaimPolicyActionRECLAIMPOLICYACTIONUNSPECIFIED ReclaimPolicyReclaimPolicyAction = "RECLAIM_POLICY_ACTION_UNSPECIFIED"

	// ReclaimPolicyReclaimPolicyActionDELETE captures enum value "DELETE"
	ReclaimPolicyReclaimPolicyActionDELETE ReclaimPolicyReclaimPolicyAction = "DELETE"

	// ReclaimPolicyReclaimPolicyActionRETAIN captures enum value "RETAIN"
	ReclaimPolicyReclaimPolicyActionRETAIN ReclaimPolicyReclaimPolicyAction = "RETAIN"

	// ReclaimPolicyReclaimPolicyActionRETAINFORIDLEPERIOD captures enum value "RETAIN_FOR_IDLE_PERIOD"
	ReclaimPolicyReclaimPolicyActionRETAINFORIDLEPERIOD ReclaimPolicyReclaimPolicyAction = "RETAIN_FOR_IDLE_PERIOD"
)

// for schema
var reclaimPolicyReclaimPolicyActionEnum []interface{}

func init() {
	var res []ReclaimPolicyReclaimPolicyAction
	if err := json.Unmarshal([]byte(`["RECLAIM_POLICY_ACTION_UNSPECIFIED","DELETE","RETAIN","RETAIN_FOR_IDLE_PERIOD"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		reclaimPolicyReclaimPolicyActionEnum = append(reclaimPolicyReclaimPolicyActionEnum, v)
	}
}

func (m ReclaimPolicyReclaimPolicyAction) validateReclaimPolicyReclaimPolicyActionEnum(path, location string, value ReclaimPolicyReclaimPolicyAction) error {
	if err := validate.EnumCase(path, location, value, reclaimPolicyReclaimPolicyActionEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this reclaim policy reclaim policy action
func (m ReclaimPolicyReclaimPolicyAction) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateReclaimPolicyReclaimPolicyActionEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this reclaim policy reclaim policy action based on context it is used
func (m ReclaimPolicyReclaimPolicyAction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

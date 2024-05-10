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

// RouteStatusState - ROUTE_STATUS_UNSPECIFIED: Route in this region is in an undefined state
//   - RUNNING: Route has 1 or more assigned gateways and is correctly configured in this region
//   - NO_GATEWAY: Route does not have any assigned gateways but is configured in this region
//   - DELETING: Route is being deleted from the region
//
// swagger:model RouteStatusState
type RouteStatusState string

func NewRouteStatusState(value RouteStatusState) *RouteStatusState {
	v := value
	return &v
}

const (

	// RouteStatusStateROUTESTATUSUNSPECIFIED captures enum value "ROUTE_STATUS_UNSPECIFIED"
	RouteStatusStateROUTESTATUSUNSPECIFIED RouteStatusState = "ROUTE_STATUS_UNSPECIFIED"

	// RouteStatusStateRUNNING captures enum value "RUNNING"
	RouteStatusStateRUNNING RouteStatusState = "RUNNING"

	// RouteStatusStateNOGATEWAY captures enum value "NO_GATEWAY"
	RouteStatusStateNOGATEWAY RouteStatusState = "NO_GATEWAY"

	// RouteStatusStateDELETING captures enum value "DELETING"
	RouteStatusStateDELETING RouteStatusState = "DELETING"
)

// for schema
var routeStatusStateEnum []interface{}

func init() {
	var res []RouteStatusState
	if err := json.Unmarshal([]byte(`["ROUTE_STATUS_UNSPECIFIED","RUNNING","NO_GATEWAY","DELETING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		routeStatusStateEnum = append(routeStatusStateEnum, v)
	}
}

func (m RouteStatusState) validateRouteStatusStateEnum(path, location string, value RouteStatusState) error {
	if err := validate.EnumCase(path, location, value, routeStatusStateEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this route status state
func (m RouteStatusState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateRouteStatusStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this route status state based on context it is used
func (m RouteStatusState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

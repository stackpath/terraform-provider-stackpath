// Code generated by go-swagger; DO NOT EDIT.

package infrastructure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new infrastructure API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for infrastructure API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetLocations(params *GetLocationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetLocationsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetLocations gets compute locations

  Retrieve information about the StackPath edge network that can host a compute workload
*/
func (a *Client) GetLocations(params *GetLocationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetLocationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLocationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetLocations",
		Method:             "GET",
		PathPattern:        "/workload/v1/locations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLocationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLocationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetLocationsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

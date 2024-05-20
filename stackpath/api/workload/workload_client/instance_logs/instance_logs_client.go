// Code generated by go-swagger; DO NOT EDIT.

package instance_logs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new instance logs API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for instance logs API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetLogs(params *GetLogsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetLogsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetLogs gets log stream

  Retrieve a stream of logs generated by a workload instance's containers. Logs are generated by the containers and are not modified by StackPath. Log contents vary by the application running in the container, though many containerized applications are configured to log to STDOUT and STDERR.
*/
func (a *Client) GetLogs(params *GetLogsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetLogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLogsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetLogs",
		Method:             "GET",
		PathPattern:        "/workload/v1/stacks/{stack_id}/workloads/{workload_id}/instances/{instance_name}/logs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLogsReader{formats: a.formats},
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
	success, ok := result.(*GetLogsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetLogsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

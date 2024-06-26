// Code generated by go-swagger; DO NOT EDIT.

package workloads

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new workloads API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for workloads API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateWorkload(params *CreateWorkloadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateWorkloadOK, error)

	DeleteWorkload(params *DeleteWorkloadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteWorkloadNoContent, error)

	GetWorkload(params *GetWorkloadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetWorkloadOK, error)

	GetWorkloads(params *GetWorkloadsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetWorkloadsOK, error)

	PutWorkload(params *PutWorkloadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutWorkloadOK, error)

	UpdateWorkload(params *UpdateWorkloadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateWorkloadOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateWorkload creates a workload
*/
func (a *Client) CreateWorkload(params *CreateWorkloadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateWorkloadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateWorkloadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateWorkload",
		Method:             "POST",
		PathPattern:        "/workload/v1/stacks/{stack_id}/workloads",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateWorkloadReader{formats: a.formats},
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
	success, ok := result.(*CreateWorkloadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateWorkloadDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteWorkload deletes a workload
*/
func (a *Client) DeleteWorkload(params *DeleteWorkloadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteWorkloadNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteWorkloadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteWorkload",
		Method:             "DELETE",
		PathPattern:        "/workload/v1/stacks/{stack_id}/workloads/{workload_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteWorkloadReader{formats: a.formats},
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
	success, ok := result.(*DeleteWorkloadNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteWorkloadDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetWorkload gets a workload
*/
func (a *Client) GetWorkload(params *GetWorkloadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetWorkloadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWorkloadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetWorkload",
		Method:             "GET",
		PathPattern:        "/workload/v1/stacks/{stack_id}/workloads/{workload_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetWorkloadReader{formats: a.formats},
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
	success, ok := result.(*GetWorkloadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetWorkloadDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetWorkloads gets all workloads
*/
func (a *Client) GetWorkloads(params *GetWorkloadsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetWorkloadsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWorkloadsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetWorkloads",
		Method:             "GET",
		PathPattern:        "/workload/v1/stacks/{stack_id}/workloads",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetWorkloadsReader{formats: a.formats},
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
	success, ok := result.(*GetWorkloadsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetWorkloadsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PutWorkload replaces a workload
*/
func (a *Client) PutWorkload(params *PutWorkloadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutWorkloadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutWorkloadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutWorkload",
		Method:             "PUT",
		PathPattern:        "/workload/v1/stacks/{stack_id}/workloads/{workload_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutWorkloadReader{formats: a.formats},
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
	success, ok := result.(*PutWorkloadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutWorkloadDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdateWorkload updates a workload
*/
func (a *Client) UpdateWorkload(params *UpdateWorkloadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateWorkloadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateWorkloadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateWorkload",
		Method:             "PATCH",
		PathPattern:        "/workload/v1/stacks/{stack_id}/workloads/{workload_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateWorkloadReader{formats: a.formats},
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
	success, ok := result.(*UpdateWorkloadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateWorkloadDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

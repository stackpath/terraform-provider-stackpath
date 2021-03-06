// Code generated by go-swagger; DO NOT EDIT.

package buckets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new buckets API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for buckets API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateBucket(params *CreateBucketParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateBucketOK, error)

	DeleteBucket(params *DeleteBucketParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteBucketNoContent, error)

	GetBucket(params *GetBucketParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBucketOK, error)

	GetBuckets(params *GetBucketsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBucketsOK, error)

	UpdateBucket(params *UpdateBucketParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateBucketOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateBucket creates a bucket under a stack
*/
func (a *Client) CreateBucket(params *CreateBucketParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateBucketOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateBucketParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateBucket",
		Method:             "POST",
		PathPattern:        "/storage/v1/stacks/{stack_id}/buckets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateBucketReader{formats: a.formats},
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
	success, ok := result.(*CreateBucketOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateBucketDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteBucket deletes a given bucket
*/
func (a *Client) DeleteBucket(params *DeleteBucketParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteBucketNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteBucketParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteBucket",
		Method:             "DELETE",
		PathPattern:        "/storage/v1/stacks/{stack_id}/buckets/{bucket_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteBucketReader{formats: a.formats},
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
	success, ok := result.(*DeleteBucketNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteBucketDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetBucket retrieves a bucket in the storage provider for a given stack
*/
func (a *Client) GetBucket(params *GetBucketParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBucketOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBucketParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetBucket",
		Method:             "GET",
		PathPattern:        "/storage/v1/stacks/{stack_id}/buckets/{bucket_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBucketReader{formats: a.formats},
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
	success, ok := result.(*GetBucketOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetBucketDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetBuckets retrieves all buckets in the storage provider for a given stack
*/
func (a *Client) GetBuckets(params *GetBucketsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBucketsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBucketsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetBuckets",
		Method:             "GET",
		PathPattern:        "/storage/v1/stacks/{stack_id}/buckets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBucketsReader{formats: a.formats},
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
	success, ok := result.(*GetBucketsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetBucketsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdateBucket updates the name of a bucket
*/
func (a *Client) UpdateBucket(params *UpdateBucketParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateBucketOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateBucketParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateBucket",
		Method:             "PUT",
		PathPattern:        "/storage/v1/stacks/{stack_id}/buckets/{bucket_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateBucketReader{formats: a.formats},
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
	success, ok := result.(*UpdateBucketOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateBucketDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

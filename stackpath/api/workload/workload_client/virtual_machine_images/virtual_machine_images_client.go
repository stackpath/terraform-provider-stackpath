// Code generated by go-swagger; DO NOT EDIT.

package virtual_machine_images

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new virtual machine images API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for virtual machine images API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateImage(params *CreateImageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateImageOK, error)

	DeleteImage(params *DeleteImageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteImageNoContent, error)

	DeleteImagesForFamily(params *DeleteImagesForFamilyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteImagesForFamilyNoContent, error)

	GetImage(params *GetImageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetImageOK, error)

	GetImages(params *GetImagesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetImagesOK, error)

	GetImagesForFamily(params *GetImagesForFamilyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetImagesForFamilyOK, error)

	UpdateImage(params *UpdateImageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateImageOK, error)

	UpdateImageDeprecation(params *UpdateImageDeprecationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateImageDeprecationOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateImage creates an image
*/
func (a *Client) CreateImage(params *CreateImageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateImageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateImage",
		Method:             "POST",
		PathPattern:        "/workload/v1/stacks/{stack_id}/images/{image_family}/{image_tag}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateImageReader{formats: a.formats},
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
	success, ok := result.(*CreateImageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateImageDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteImage deletes an image
*/
func (a *Client) DeleteImage(params *DeleteImageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteImageNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteImageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteImage",
		Method:             "DELETE",
		PathPattern:        "/workload/v1/stacks/{stack_id}/images/{image_family}/{image_tag}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteImageReader{formats: a.formats},
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
	success, ok := result.(*DeleteImageNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteImageDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteImagesForFamily deletes a family s images
*/
func (a *Client) DeleteImagesForFamily(params *DeleteImagesForFamilyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteImagesForFamilyNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteImagesForFamilyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteImagesForFamily",
		Method:             "DELETE",
		PathPattern:        "/workload/v1/stacks/{stack_id}/images/{image_family}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteImagesForFamilyReader{formats: a.formats},
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
	success, ok := result.(*DeleteImagesForFamilyNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteImagesForFamilyDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetImage gets an image
*/
func (a *Client) GetImage(params *GetImageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetImageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetImage",
		Method:             "GET",
		PathPattern:        "/workload/v1/stacks/{stack_id}/images/{image_family}/{image_tag}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetImageReader{formats: a.formats},
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
	success, ok := result.(*GetImageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetImageDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetImages gets all images

  Only non-deprecated images are returned by default
*/
func (a *Client) GetImages(params *GetImagesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetImagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetImagesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetImages",
		Method:             "GET",
		PathPattern:        "/workload/v1/stacks/{stack_id}/images",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetImagesReader{formats: a.formats},
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
	success, ok := result.(*GetImagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetImagesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetImagesForFamily gets a family s images

  Only non-deprecated images are returned by default. This will not error but instead return an empty list if the family is not found.
*/
func (a *Client) GetImagesForFamily(params *GetImagesForFamilyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetImagesForFamilyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetImagesForFamilyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetImagesForFamily",
		Method:             "GET",
		PathPattern:        "/workload/v1/stacks/{stack_id}/images/{image_family}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetImagesForFamilyReader{formats: a.formats},
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
	success, ok := result.(*GetImagesForFamilyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetImagesForFamilyDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdateImage updates an image

  Only metadata and description can be updated. The metadata, if set, replaces the entire existing metadata set. The tag cannot be "default".
*/
func (a *Client) UpdateImage(params *UpdateImageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateImageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateImage",
		Method:             "PATCH",
		PathPattern:        "/workload/v1/stacks/{stack_id}/images/{image_family}/{image_tag}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateImageReader{formats: a.formats},
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
	success, ok := result.(*UpdateImageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateImageDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdateImageDeprecation updates deprecation settings

  This replaces an image's deprecation settings, so it can also undeprecate an image.
*/
func (a *Client) UpdateImageDeprecation(params *UpdateImageDeprecationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateImageDeprecationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateImageDeprecationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateImageDeprecation",
		Method:             "PUT",
		PathPattern:        "/workload/v1/stacks/{stack_id}/images/{image_family}/{image_tag}/deprecation",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateImageDeprecationReader{formats: a.formats},
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
	success, ok := result.(*UpdateImageDeprecationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateImageDeprecationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

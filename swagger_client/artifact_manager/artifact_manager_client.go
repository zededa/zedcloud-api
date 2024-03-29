// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package artifact_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new artifact manager API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for artifact manager API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ArtifactManagerCreateArtifact(params *ArtifactManagerCreateArtifactParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ArtifactManagerCreateArtifactOK, error)

	ArtifactManagerDeleteArtifact(params *ArtifactManagerDeleteArtifactParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ArtifactManagerDeleteArtifactOK, error)

	ArtifactManagerGetArtifactSignedURL(params *ArtifactManagerGetArtifactSignedURLParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ArtifactManagerGetArtifactSignedURLOK, error)

	ArtifactManagerGetArtifactStream(params *ArtifactManagerGetArtifactStreamParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ArtifactManagerGetArtifactStreamOK, *ArtifactManagerGetArtifactStreamPartialContent, error)

	ArtifactManagerQueryArtifacts(params *ArtifactManagerQueryArtifactsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ArtifactManagerQueryArtifactsOK, error)

	ArtifactManagerUploadArtifact(params *ArtifactManagerUploadArtifactParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ArtifactManagerUploadArtifactOK, *ArtifactManagerUploadArtifactAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ArtifactManagerCreateArtifact creates artifact file in the file storage

Create the artifact file in the file storage in AWS S3 or Azure BlobStorage
*/
func (a *Client) ArtifactManagerCreateArtifact(params *ArtifactManagerCreateArtifactParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ArtifactManagerCreateArtifactOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewArtifactManagerCreateArtifactParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ArtifactManager_CreateArtifact",
		Method:             "POST",
		PathPattern:        "/v1/artifacts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ArtifactManagerCreateArtifactReader{formats: a.formats},
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
	success, ok := result.(*ArtifactManagerCreateArtifactOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ArtifactManagerCreateArtifactDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ArtifactManagerDeleteArtifact deletes artifact file form the datastore

Drops the artifact file from the file storage in AWS S3 or Azure BlobStorage
*/
func (a *Client) ArtifactManagerDeleteArtifact(params *ArtifactManagerDeleteArtifactParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ArtifactManagerDeleteArtifactOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewArtifactManagerDeleteArtifactParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ArtifactManager_DeleteArtifact",
		Method:             "DELETE",
		PathPattern:        "/v1/artifacts/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ArtifactManagerDeleteArtifactReader{formats: a.formats},
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
	success, ok := result.(*ArtifactManagerDeleteArtifactOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ArtifactManagerDeleteArtifactDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ArtifactManagerGetArtifactSignedURL generates the signed Url for accessing the resource

Generate the URL which can be used to access the resource from datastore like s3, Azure etc for specified amount of time.
*/
func (a *Client) ArtifactManagerGetArtifactSignedURL(params *ArtifactManagerGetArtifactSignedURLParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ArtifactManagerGetArtifactSignedURLOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewArtifactManagerGetArtifactSignedURLParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ArtifactManager_GetArtifactSignedUrl",
		Method:             "GET",
		PathPattern:        "/v1/artifacts/id/{id}/url",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ArtifactManagerGetArtifactSignedURLReader{formats: a.formats},
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
	success, ok := result.(*ArtifactManagerGetArtifactSignedURLOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ArtifactManagerGetArtifactSignedURLDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ArtifactManagerGetArtifactStream downloads artifact file chunk by chunk

Downloads the artifact file in a stream from the file storage
*/
func (a *Client) ArtifactManagerGetArtifactStream(params *ArtifactManagerGetArtifactStreamParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ArtifactManagerGetArtifactStreamOK, *ArtifactManagerGetArtifactStreamPartialContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewArtifactManagerGetArtifactStreamParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ArtifactManager_GetArtifactStream",
		Method:             "GET",
		PathPattern:        "/v1/artifacts/id/{id}",
		ProducesMediaTypes: []string{"application/octet-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ArtifactManagerGetArtifactStreamReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ArtifactManagerGetArtifactStreamOK:
		return value, nil, nil
	case *ArtifactManagerGetArtifactStreamPartialContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ArtifactManagerGetArtifactStreamDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ArtifactManagerQueryArtifacts queries artifact files

Query the artifact file records.
*/
func (a *Client) ArtifactManagerQueryArtifacts(params *ArtifactManagerQueryArtifactsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ArtifactManagerQueryArtifactsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewArtifactManagerQueryArtifactsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ArtifactManager_QueryArtifacts",
		Method:             "GET",
		PathPattern:        "/v1/artifacts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ArtifactManagerQueryArtifactsReader{formats: a.formats},
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
	success, ok := result.(*ArtifactManagerQueryArtifactsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ArtifactManagerQueryArtifactsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ArtifactManagerUploadArtifact uploads the given chunk into the specified multiple part file

Uploads the given chunk into the specified file
*/
func (a *Client) ArtifactManagerUploadArtifact(params *ArtifactManagerUploadArtifactParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ArtifactManagerUploadArtifactOK, *ArtifactManagerUploadArtifactAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewArtifactManagerUploadArtifactParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ArtifactManager_UploadArtifact",
		Method:             "PUT",
		PathPattern:        "/v1/artifacts/id/{id}/upload/chunked",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ArtifactManagerUploadArtifactReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ArtifactManagerUploadArtifactOK:
		return value, nil, nil
	case *ArtifactManagerUploadArtifactAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ArtifactManagerUploadArtifactDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package image_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new image configuration API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for image configuration API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ImageConfigurationCreateImage(params *ImageConfigurationCreateImageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImageConfigurationCreateImageOK, error)

	ImageConfigurationDeleteImage(params *ImageConfigurationDeleteImageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImageConfigurationDeleteImageOK, error)

	ImageConfigurationDeleteLatestImage(params *ImageConfigurationDeleteLatestImageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImageConfigurationDeleteLatestImageOK, error)

	ImageConfigurationGetImage(params *ImageConfigurationGetImageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImageConfigurationGetImageOK, error)

	ImageConfigurationGetImageByName(params *ImageConfigurationGetImageByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImageConfigurationGetImageByNameOK, error)

	ImageConfigurationGetLatestImageVersion(params *ImageConfigurationGetLatestImageVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImageConfigurationGetLatestImageVersionOK, error)

	ImageConfigurationMarkEveImageLatest(params *ImageConfigurationMarkEveImageLatestParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImageConfigurationMarkEveImageLatestOK, error)

	ImageConfigurationMarkEveImageLatest2(params *ImageConfigurationMarkEveImageLatest2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImageConfigurationMarkEveImageLatest2OK, error)

	ImageConfigurationQueryImageProjectList(params *ImageConfigurationQueryImageProjectListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImageConfigurationQueryImageProjectListOK, error)

	ImageConfigurationQueryImages(params *ImageConfigurationQueryImagesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImageConfigurationQueryImagesOK, error)

	ImageConfigurationQueryLatestImageVersions(params *ImageConfigurationQueryLatestImageVersionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImageConfigurationQueryLatestImageVersionsOK, error)

	ImageConfigurationUpdateImage(params *ImageConfigurationUpdateImageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImageConfigurationUpdateImageOK, error)

	ImageConfigurationUplinkImage(params *ImageConfigurationUplinkImageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImageConfigurationUplinkImageOK, *ImageConfigurationUplinkImageAccepted, error)

	ImageConfigurationUploadImageChunked(params *ImageConfigurationUploadImageChunkedParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImageConfigurationUploadImageChunkedOK, *ImageConfigurationUploadImageChunkedAccepted, error)

	ImageConfigurationUploadImageFile(params *ImageConfigurationUploadImageFileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImageConfigurationUploadImageFileOK, *ImageConfigurationUploadImageFileAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ImageConfigurationCreateImage creates edge application image

Create an edge application image record.
*/
func (a *Client) ImageConfigurationCreateImage(params *ImageConfigurationCreateImageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImageConfigurationCreateImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImageConfigurationCreateImageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ImageConfiguration_CreateImage",
		Method:             "POST",
		PathPattern:        "/v1/apps/images",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImageConfigurationCreateImageReader{formats: a.formats},
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
	success, ok := result.(*ImageConfigurationCreateImageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ImageConfigurationCreateImageDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ImageConfigurationDeleteImage deletes edge application image

Delete an edge application image record.
*/
func (a *Client) ImageConfigurationDeleteImage(params *ImageConfigurationDeleteImageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImageConfigurationDeleteImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImageConfigurationDeleteImageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ImageConfiguration_DeleteImage",
		Method:             "DELETE",
		PathPattern:        "/v1/apps/images/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImageConfigurationDeleteImageReader{formats: a.formats},
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
	success, ok := result.(*ImageConfigurationDeleteImageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ImageConfigurationDeleteImageDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ImageConfigurationDeleteLatestImage deletes edge application image

Delete an edge application image record.
*/
func (a *Client) ImageConfigurationDeleteLatestImage(params *ImageConfigurationDeleteLatestImageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImageConfigurationDeleteLatestImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImageConfigurationDeleteLatestImageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ImageConfiguration_DeleteLatestImage",
		Method:             "DELETE",
		PathPattern:        "/v1/apps/images/baseos/latest",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImageConfigurationDeleteLatestImageReader{formats: a.formats},
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
	success, ok := result.(*ImageConfigurationDeleteLatestImageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ImageConfigurationDeleteLatestImageDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ImageConfigurationGetImage gets edge application image

Get the configuration (without security details) of an edge application image record.
*/
func (a *Client) ImageConfigurationGetImage(params *ImageConfigurationGetImageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImageConfigurationGetImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImageConfigurationGetImageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ImageConfiguration_GetImage",
		Method:             "GET",
		PathPattern:        "/v1/apps/images/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImageConfigurationGetImageReader{formats: a.formats},
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
	success, ok := result.(*ImageConfigurationGetImageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ImageConfigurationGetImageDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ImageConfigurationGetImageByName gets edge application image

Get the configuration (without security details) of an edge application image record.
*/
func (a *Client) ImageConfigurationGetImageByName(params *ImageConfigurationGetImageByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImageConfigurationGetImageByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImageConfigurationGetImageByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ImageConfiguration_GetImageByName",
		Method:             "GET",
		PathPattern:        "/v1/apps/images/name/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImageConfigurationGetImageByNameReader{formats: a.formats},
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
	success, ok := result.(*ImageConfigurationGetImageByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ImageConfigurationGetImageByNameDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ImageConfigurationGetLatestImageVersion gets latest version of e v e image

Query the latest version of EVE image for given hardware architecture.
*/
func (a *Client) ImageConfigurationGetLatestImageVersion(params *ImageConfigurationGetLatestImageVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImageConfigurationGetLatestImageVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImageConfigurationGetLatestImageVersionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ImageConfiguration_GetLatestImageVersion",
		Method:             "GET",
		PathPattern:        "/v1/apps/images/baseos/latest/hwclass/{imageArch}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImageConfigurationGetLatestImageVersionReader{formats: a.formats},
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
	success, ok := result.(*ImageConfigurationGetLatestImageVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ImageConfigurationGetLatestImageVersionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ImageConfigurationMarkEveImageLatest marks eve image as latest

Mark Eve image as latest for a given hardware architecture.
*/
func (a *Client) ImageConfigurationMarkEveImageLatest(params *ImageConfigurationMarkEveImageLatestParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImageConfigurationMarkEveImageLatestOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImageConfigurationMarkEveImageLatestParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ImageConfiguration_MarkEveImageLatest",
		Method:             "PUT",
		PathPattern:        "/v1/apps/images/baseos/latest",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImageConfigurationMarkEveImageLatestReader{formats: a.formats},
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
	success, ok := result.(*ImageConfigurationMarkEveImageLatestOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ImageConfigurationMarkEveImageLatestDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ImageConfigurationMarkEveImageLatest2 marks eve image as latest

Mark Eve image as latest for a given hardware architecture.
*/
func (a *Client) ImageConfigurationMarkEveImageLatest2(params *ImageConfigurationMarkEveImageLatest2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImageConfigurationMarkEveImageLatest2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImageConfigurationMarkEveImageLatest2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "ImageConfiguration_MarkEveImageLatest2",
		Method:             "PUT",
		PathPattern:        "/v1/apps/images/baseos/latest/hwclass/{imageArch}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImageConfigurationMarkEveImageLatest2Reader{formats: a.formats},
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
	success, ok := result.(*ImageConfigurationMarkEveImageLatest2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ImageConfigurationMarkEveImageLatest2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ImageConfigurationQueryImageProjectList queries common project access list among a list of images

Query common project access list among a list of images
*/
func (a *Client) ImageConfigurationQueryImageProjectList(params *ImageConfigurationQueryImageProjectListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImageConfigurationQueryImageProjectListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImageConfigurationQueryImageProjectListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ImageConfiguration_QueryImageProjectList",
		Method:             "GET",
		PathPattern:        "/v1/apps/images/projects",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImageConfigurationQueryImageProjectListReader{formats: a.formats},
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
	success, ok := result.(*ImageConfigurationQueryImageProjectListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ImageConfigurationQueryImageProjectListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ImageConfigurationQueryImages queries edge application images

Query the edge application image records.
*/
func (a *Client) ImageConfigurationQueryImages(params *ImageConfigurationQueryImagesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImageConfigurationQueryImagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImageConfigurationQueryImagesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ImageConfiguration_QueryImages",
		Method:             "GET",
		PathPattern:        "/v1/apps/images",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImageConfigurationQueryImagesReader{formats: a.formats},
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
	success, ok := result.(*ImageConfigurationQueryImagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ImageConfigurationQueryImagesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ImageConfigurationQueryLatestImageVersions queries latest version of e v e image for each hardware architecture

Query the latest version of EVE image for each hardware architecture.
*/
func (a *Client) ImageConfigurationQueryLatestImageVersions(params *ImageConfigurationQueryLatestImageVersionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImageConfigurationQueryLatestImageVersionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImageConfigurationQueryLatestImageVersionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ImageConfiguration_QueryLatestImageVersions",
		Method:             "GET",
		PathPattern:        "/v1/apps/images/baseos/latest",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImageConfigurationQueryLatestImageVersionsReader{formats: a.formats},
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
	success, ok := result.(*ImageConfigurationQueryLatestImageVersionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ImageConfigurationQueryLatestImageVersionsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ImageConfigurationUpdateImage updates edge application image

Update an edge application image. The usual pattern to update an edge application image record is to retrieve the record and update with the modified values in a new body to update the edge application image record.
*/
func (a *Client) ImageConfigurationUpdateImage(params *ImageConfigurationUpdateImageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImageConfigurationUpdateImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImageConfigurationUpdateImageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ImageConfiguration_UpdateImage",
		Method:             "PUT",
		PathPattern:        "/v1/apps/images/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImageConfigurationUpdateImageReader{formats: a.formats},
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
	success, ok := result.(*ImageConfigurationUpdateImageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ImageConfigurationUpdateImageDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ImageConfigurationUplinkImage uplinks edge application image

Uplinks the edge application image record to an existing binry file in the datastore.
*/
func (a *Client) ImageConfigurationUplinkImage(params *ImageConfigurationUplinkImageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImageConfigurationUplinkImageOK, *ImageConfigurationUplinkImageAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImageConfigurationUplinkImageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ImageConfiguration_UplinkImage",
		Method:             "PUT",
		PathPattern:        "/v1/apps/images/name/{name}/uplink",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImageConfigurationUplinkImageReader{formats: a.formats},
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
	case *ImageConfigurationUplinkImageOK:
		return value, nil, nil
	case *ImageConfigurationUplinkImageAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ImageConfigurationUplinkImageDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ImageConfigurationUploadImageChunked uploads edge application image binary file

Uploads the edge application image binary file in the datastore. This method uses multiple HTTP requests to upload image binary file in smaller chunks. Recommended for larger file size.
*/
func (a *Client) ImageConfigurationUploadImageChunked(params *ImageConfigurationUploadImageChunkedParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImageConfigurationUploadImageChunkedOK, *ImageConfigurationUploadImageChunkedAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImageConfigurationUploadImageChunkedParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ImageConfiguration_UploadImageChunked",
		Method:             "PUT",
		PathPattern:        "/v1/apps/images/name/{name}/upload/chunked",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImageConfigurationUploadImageChunkedReader{formats: a.formats},
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
	case *ImageConfigurationUploadImageChunkedOK:
		return value, nil, nil
	case *ImageConfigurationUploadImageChunkedAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ImageConfigurationUploadImageChunkedDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ImageConfigurationUploadImageFile uploads edge application image binary file

Uploads the edge application image binary file in the datastore. This method uses single HTTP request to upload the entire image binary file. Recommended for smaller file size.
*/
func (a *Client) ImageConfigurationUploadImageFile(params *ImageConfigurationUploadImageFileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImageConfigurationUploadImageFileOK, *ImageConfigurationUploadImageFileAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImageConfigurationUploadImageFileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ImageConfiguration_UploadImageFile",
		Method:             "PUT",
		PathPattern:        "/v1/apps/images/name/{name}/upload/file",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImageConfigurationUploadImageFileReader{formats: a.formats},
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
	case *ImageConfigurationUploadImageFileOK:
		return value, nil, nil
	case *ImageConfigurationUploadImageFileAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ImageConfigurationUploadImageFileDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

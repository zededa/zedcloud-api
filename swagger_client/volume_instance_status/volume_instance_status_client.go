// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package volume_instance_status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new volume instance status API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for volume instance status API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetVolumeInstanceEvents(params *GetVolumeInstanceEventsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVolumeInstanceEventsOK, error)

	GetVolumeInstanceEventsByName(params *GetVolumeInstanceEventsByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVolumeInstanceEventsByNameOK, error)

	GetVolumeInstanceStatus(params *GetVolumeInstanceStatusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVolumeInstanceStatusOK, error)

	GetVolumeInstanceStatusByName(params *GetVolumeInstanceStatusByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVolumeInstanceStatusByNameOK, error)

	QueryVolumeInstanceStatus(params *QueryVolumeInstanceStatusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*QueryVolumeInstanceStatusOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetVolumeInstanceEvents gets edge volume instance events by id

  Get configuration and status events of an edge volume by id.
*/
func (a *Client) GetVolumeInstanceEvents(params *GetVolumeInstanceEventsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVolumeInstanceEventsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVolumeInstanceEventsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetVolumeInstanceEvents",
		Method:             "GET",
		PathPattern:        "/v1/volumes/instances/id/{objid}/events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVolumeInstanceEventsReader{formats: a.formats},
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
	success, ok := result.(*GetVolumeInstanceEventsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetVolumeInstanceEvents: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetVolumeInstanceEventsByName gets edge volume instance events by name

  Get configuration and status events of an edge volume by name.
*/
func (a *Client) GetVolumeInstanceEventsByName(params *GetVolumeInstanceEventsByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVolumeInstanceEventsByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVolumeInstanceEventsByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetVolumeInstanceEventsByName",
		Method:             "GET",
		PathPattern:        "/v1/volumes/instances/name/{objname}/events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVolumeInstanceEventsByNameReader{formats: a.formats},
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
	success, ok := result.(*GetVolumeInstanceEventsByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetVolumeInstanceEventsByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetVolumeInstanceStatus gets edge volume instance status

  Get the status of an edge volume instance record.
*/
func (a *Client) GetVolumeInstanceStatus(params *GetVolumeInstanceStatusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVolumeInstanceStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVolumeInstanceStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetVolumeInstanceStatus",
		Method:             "GET",
		PathPattern:        "/v1/volumes/instances/id/{id}/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVolumeInstanceStatusReader{formats: a.formats},
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
	success, ok := result.(*GetVolumeInstanceStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetVolumeInstanceStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetVolumeInstanceStatusByName gets edge volume instance status

  Get the status of an edge volume instance record.
*/
func (a *Client) GetVolumeInstanceStatusByName(params *GetVolumeInstanceStatusByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVolumeInstanceStatusByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVolumeInstanceStatusByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetVolumeInstanceStatusByName",
		Method:             "GET",
		PathPattern:        "/v1/volumes/instances/name/{name}/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVolumeInstanceStatusByNameReader{formats: a.formats},
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
	success, ok := result.(*GetVolumeInstanceStatusByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetVolumeInstanceStatusByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  QueryVolumeInstanceStatus queries status of edge volume instances

  Query the status of edge volume instance records.
*/
func (a *Client) QueryVolumeInstanceStatus(params *QueryVolumeInstanceStatusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*QueryVolumeInstanceStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryVolumeInstanceStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "QueryVolumeInstanceStatus",
		Method:             "GET",
		PathPattern:        "/v1/volumes/instances/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryVolumeInstanceStatusReader{formats: a.formats},
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
	success, ok := result.(*QueryVolumeInstanceStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for QueryVolumeInstanceStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

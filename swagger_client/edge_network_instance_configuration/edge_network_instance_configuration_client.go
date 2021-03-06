// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package edge_network_instance_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new edge network instance configuration API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for edge network instance configuration API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateEdgeNetworkInstance(params *CreateEdgeNetworkInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateEdgeNetworkInstanceOK, error)

	DeleteEdgeNetworkInstance(params *DeleteEdgeNetworkInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteEdgeNetworkInstanceOK, error)

	GetEdgeNetworkInstance(params *GetEdgeNetworkInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEdgeNetworkInstanceOK, error)

	GetEdgeNetworkInstanceByName(params *GetEdgeNetworkInstanceByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEdgeNetworkInstanceByNameOK, error)

	QueryEdgeNetworkInstances(params *QueryEdgeNetworkInstancesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*QueryEdgeNetworkInstancesOK, error)

	UpdateEdgeNetworkInstance(params *UpdateEdgeNetworkInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateEdgeNetworkInstanceOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateEdgeNetworkInstance creates edge network instance

  Create an edge network instance record.
*/
func (a *Client) CreateEdgeNetworkInstance(params *CreateEdgeNetworkInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateEdgeNetworkInstanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateEdgeNetworkInstanceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateEdgeNetworkInstance",
		Method:             "POST",
		PathPattern:        "/v1/netinsts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateEdgeNetworkInstanceReader{formats: a.formats},
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
	success, ok := result.(*CreateEdgeNetworkInstanceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateEdgeNetworkInstance: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteEdgeNetworkInstance deletes edge network instance

  Delete an edge network instance record.
*/
func (a *Client) DeleteEdgeNetworkInstance(params *DeleteEdgeNetworkInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteEdgeNetworkInstanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteEdgeNetworkInstanceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteEdgeNetworkInstance",
		Method:             "DELETE",
		PathPattern:        "/v1/netinsts/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteEdgeNetworkInstanceReader{formats: a.formats},
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
	success, ok := result.(*DeleteEdgeNetworkInstanceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteEdgeNetworkInstance: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetEdgeNetworkInstance gets edge network instance

  Get the configuration (without security details) of an edge network instance record.
*/
func (a *Client) GetEdgeNetworkInstance(params *GetEdgeNetworkInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEdgeNetworkInstanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEdgeNetworkInstanceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetEdgeNetworkInstance",
		Method:             "GET",
		PathPattern:        "/v1/netinsts/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEdgeNetworkInstanceReader{formats: a.formats},
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
	success, ok := result.(*GetEdgeNetworkInstanceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetEdgeNetworkInstance: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetEdgeNetworkInstanceByName gets edge network instance

  Get the configuration (without security details) of an edge network instance record.
*/
func (a *Client) GetEdgeNetworkInstanceByName(params *GetEdgeNetworkInstanceByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEdgeNetworkInstanceByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEdgeNetworkInstanceByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetEdgeNetworkInstanceByName",
		Method:             "GET",
		PathPattern:        "/v1/netinsts/name/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEdgeNetworkInstanceByNameReader{formats: a.formats},
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
	success, ok := result.(*GetEdgeNetworkInstanceByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetEdgeNetworkInstanceByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  QueryEdgeNetworkInstances queries edge network instances

  Query the edge network instance records.
*/
func (a *Client) QueryEdgeNetworkInstances(params *QueryEdgeNetworkInstancesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*QueryEdgeNetworkInstancesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryEdgeNetworkInstancesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "QueryEdgeNetworkInstances",
		Method:             "GET",
		PathPattern:        "/v1/netinsts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryEdgeNetworkInstancesReader{formats: a.formats},
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
	success, ok := result.(*QueryEdgeNetworkInstancesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for QueryEdgeNetworkInstances: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateEdgeNetworkInstance updates edge network instance

  Update an edge network instance. The usual pattern to update an edge network instance record is to retrieve the record and update with the modified values in a new body to update the edge network instance record.
*/
func (a *Client) UpdateEdgeNetworkInstance(params *UpdateEdgeNetworkInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateEdgeNetworkInstanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateEdgeNetworkInstanceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateEdgeNetworkInstance",
		Method:             "PUT",
		PathPattern:        "/v1/netinsts/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateEdgeNetworkInstanceReader{formats: a.formats},
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
	success, ok := result.(*UpdateEdgeNetworkInstanceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateEdgeNetworkInstance: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

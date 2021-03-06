// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package datastore_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new datastore configuration API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for datastore configuration API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateDatastore(params *CreateDatastoreParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateDatastoreOK, error)

	DeleteDatastore(params *DeleteDatastoreParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteDatastoreOK, error)

	GetDatastore(params *GetDatastoreParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDatastoreOK, error)

	GetDatastoreByName(params *GetDatastoreByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDatastoreByNameOK, error)

	QueryDatastores(params *QueryDatastoresParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*QueryDatastoresOK, error)

	UpdateDatastore(params *UpdateDatastoreParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateDatastoreOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateDatastore creates datastore

  Create a Datastore record.
*/
func (a *Client) CreateDatastore(params *CreateDatastoreParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateDatastoreOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDatastoreParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateDatastore",
		Method:             "POST",
		PathPattern:        "/v1/datastores",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateDatastoreReader{formats: a.formats},
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
	success, ok := result.(*CreateDatastoreOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateDatastore: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteDatastore deletes datastore

  Delete a datastore record.
*/
func (a *Client) DeleteDatastore(params *DeleteDatastoreParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteDatastoreOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDatastoreParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteDatastore",
		Method:             "DELETE",
		PathPattern:        "/v1/datastores/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteDatastoreReader{formats: a.formats},
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
	success, ok := result.(*DeleteDatastoreOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteDatastore: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDatastore gets datastore

  Get the configuration (without security details) of a datastore record.
*/
func (a *Client) GetDatastore(params *GetDatastoreParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDatastoreOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDatastoreParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetDatastore",
		Method:             "GET",
		PathPattern:        "/v1/datastores/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDatastoreReader{formats: a.formats},
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
	success, ok := result.(*GetDatastoreOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetDatastore: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDatastoreByName gets datastore

  Get the configuration (without security details) of a datastore record.
*/
func (a *Client) GetDatastoreByName(params *GetDatastoreByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDatastoreByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDatastoreByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetDatastoreByName",
		Method:             "GET",
		PathPattern:        "/v1/datastores/name/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDatastoreByNameReader{formats: a.formats},
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
	success, ok := result.(*GetDatastoreByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetDatastoreByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  QueryDatastores queries datastores

  Query the datastore records.
*/
func (a *Client) QueryDatastores(params *QueryDatastoresParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*QueryDatastoresOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryDatastoresParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "QueryDatastores",
		Method:             "GET",
		PathPattern:        "/v1/datastores",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryDatastoresReader{formats: a.formats},
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
	success, ok := result.(*QueryDatastoresOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for QueryDatastores: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateDatastore updates datastore

  Update a datastore record. The usual pattern to update a datastore record is to retrieve the record and update with the modified values in a new body to update the datastore record.
*/
func (a *Client) UpdateDatastore(params *UpdateDatastoreParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateDatastoreOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDatastoreParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateDatastore",
		Method:             "PUT",
		PathPattern:        "/v1/datastores/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateDatastoreReader{formats: a.formats},
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
	success, ok := result.(*UpdateDatastoreOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateDatastore: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

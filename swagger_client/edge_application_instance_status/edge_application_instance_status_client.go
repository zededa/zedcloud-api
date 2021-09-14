// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package edge_application_instance_status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new edge application instance status API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for edge application instance status API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetEdgeApplicationInstanceEvents(params *GetEdgeApplicationInstanceEventsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEdgeApplicationInstanceEventsOK, error)

	GetEdgeApplicationInstanceEventsByName(params *GetEdgeApplicationInstanceEventsByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEdgeApplicationInstanceEventsByNameOK, error)

	GetEdgeApplicationInstanceLogs(params *GetEdgeApplicationInstanceLogsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEdgeApplicationInstanceLogsOK, error)

	GetEdgeApplicationInstanceResourceMetricsByID(params *GetEdgeApplicationInstanceResourceMetricsByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEdgeApplicationInstanceResourceMetricsByIDOK, error)

	GetEdgeApplicationInstanceResourceMetricsByName(params *GetEdgeApplicationInstanceResourceMetricsByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEdgeApplicationInstanceResourceMetricsByNameOK, error)

	GetEdgeApplicationInstanceStatus(params *GetEdgeApplicationInstanceStatusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEdgeApplicationInstanceStatusOK, error)

	GetEdgeApplicationInstanceStatusByName(params *GetEdgeApplicationInstanceStatusByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEdgeApplicationInstanceStatusByNameOK, error)

	GetEdgeApplicationInstanceTopTalkers(params *GetEdgeApplicationInstanceTopTalkersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEdgeApplicationInstanceTopTalkersOK, error)

	GetEdgeApplicationInstanceTopTalkers2(params *GetEdgeApplicationInstanceTopTalkers2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEdgeApplicationInstanceTopTalkers2OK, error)

	GetEdgeApplicationInstanceTrafficFlows(params *GetEdgeApplicationInstanceTrafficFlowsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEdgeApplicationInstanceTrafficFlowsOK, error)

	GetEdgeApplicationInstanceTrafficFlows2(params *GetEdgeApplicationInstanceTrafficFlows2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEdgeApplicationInstanceTrafficFlows2OK, error)

	QueryEdgeApplicationInstanceStatus(params *QueryEdgeApplicationInstanceStatusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*QueryEdgeApplicationInstanceStatusOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetEdgeApplicationInstanceEvents gets edge application instance events by id

  Get configuration and status events of an edge application instance by id.
*/
func (a *Client) GetEdgeApplicationInstanceEvents(params *GetEdgeApplicationInstanceEventsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEdgeApplicationInstanceEventsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEdgeApplicationInstanceEventsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetEdgeApplicationInstanceEvents",
		Method:             "GET",
		PathPattern:        "/v1/apps/instances/id/{objid}/events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEdgeApplicationInstanceEventsReader{formats: a.formats},
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
	success, ok := result.(*GetEdgeApplicationInstanceEventsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetEdgeApplicationInstanceEvents: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetEdgeApplicationInstanceEventsByName gets edge application instance events by name

  Get configuration and status events of an edge application instance by name.
*/
func (a *Client) GetEdgeApplicationInstanceEventsByName(params *GetEdgeApplicationInstanceEventsByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEdgeApplicationInstanceEventsByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEdgeApplicationInstanceEventsByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetEdgeApplicationInstanceEventsByName",
		Method:             "GET",
		PathPattern:        "/v1/apps/instances/name/{objname}/events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEdgeApplicationInstanceEventsByNameReader{formats: a.formats},
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
	success, ok := result.(*GetEdgeApplicationInstanceEventsByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetEdgeApplicationInstanceEventsByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetEdgeApplicationInstanceLogs gets edge application instance log

  Get the log of an edge application instance as reported by the edge node where the edge application instance has been deployed.
*/
func (a *Client) GetEdgeApplicationInstanceLogs(params *GetEdgeApplicationInstanceLogsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEdgeApplicationInstanceLogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEdgeApplicationInstanceLogsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetEdgeApplicationInstanceLogs",
		Method:             "GET",
		PathPattern:        "/v1/apps/instances/id/{id}/logs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEdgeApplicationInstanceLogsReader{formats: a.formats},
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
	success, ok := result.(*GetEdgeApplicationInstanceLogsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetEdgeApplicationInstanceLogs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetEdgeApplicationInstanceResourceMetricsByID gets edge application instance resource usage timeline

  Get the resource usage timeline of an edge application instance as reported by the edge node where the edge application instance has been deployed.
*/
func (a *Client) GetEdgeApplicationInstanceResourceMetricsByID(params *GetEdgeApplicationInstanceResourceMetricsByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEdgeApplicationInstanceResourceMetricsByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEdgeApplicationInstanceResourceMetricsByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetEdgeApplicationInstanceResourceMetricsById",
		Method:             "GET",
		PathPattern:        "/v1/apps/instances/id/{objid}/timeSeries/{mType}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEdgeApplicationInstanceResourceMetricsByIDReader{formats: a.formats},
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
	success, ok := result.(*GetEdgeApplicationInstanceResourceMetricsByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetEdgeApplicationInstanceResourceMetricsById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetEdgeApplicationInstanceResourceMetricsByName gets edge application instance resource usage timeline

  Get the resource usage timeline of an edge application instance as reported by the edge node where the edge application instance has been deployed.
*/
func (a *Client) GetEdgeApplicationInstanceResourceMetricsByName(params *GetEdgeApplicationInstanceResourceMetricsByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEdgeApplicationInstanceResourceMetricsByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEdgeApplicationInstanceResourceMetricsByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetEdgeApplicationInstanceResourceMetricsByName",
		Method:             "GET",
		PathPattern:        "/v1/apps/instances/name/{objname}/timeSeries/{mType}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEdgeApplicationInstanceResourceMetricsByNameReader{formats: a.formats},
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
	success, ok := result.(*GetEdgeApplicationInstanceResourceMetricsByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetEdgeApplicationInstanceResourceMetricsByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetEdgeApplicationInstanceStatus gets edge application instance status

  Get the status of an edge application instance as reported by the edge node where the edge application instance has been deployed.
*/
func (a *Client) GetEdgeApplicationInstanceStatus(params *GetEdgeApplicationInstanceStatusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEdgeApplicationInstanceStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEdgeApplicationInstanceStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetEdgeApplicationInstanceStatus",
		Method:             "GET",
		PathPattern:        "/v1/apps/instances/id/{id}/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEdgeApplicationInstanceStatusReader{formats: a.formats},
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
	success, ok := result.(*GetEdgeApplicationInstanceStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetEdgeApplicationInstanceStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetEdgeApplicationInstanceStatusByName gets edge application instance status

  Get the status of an edge application instance as reported by the edge node where the edge application instance has been deployed.
*/
func (a *Client) GetEdgeApplicationInstanceStatusByName(params *GetEdgeApplicationInstanceStatusByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEdgeApplicationInstanceStatusByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEdgeApplicationInstanceStatusByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetEdgeApplicationInstanceStatusByName",
		Method:             "GET",
		PathPattern:        "/v1/apps/instances/name/{name}/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEdgeApplicationInstanceStatusByNameReader{formats: a.formats},
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
	success, ok := result.(*GetEdgeApplicationInstanceStatusByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetEdgeApplicationInstanceStatusByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetEdgeApplicationInstanceTopTalkers gets edge application instance top talkers of the network traffic flows

  Get the top talkers of the network traffic flows of an edge application instance as reported by the edge node where the edge application instance has been deployed.
*/
func (a *Client) GetEdgeApplicationInstanceTopTalkers(params *GetEdgeApplicationInstanceTopTalkersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEdgeApplicationInstanceTopTalkersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEdgeApplicationInstanceTopTalkersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetEdgeApplicationInstanceTopTalkers",
		Method:             "GET",
		PathPattern:        "/v1/apps/instances/id/{id}/flowlog/toptalkers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEdgeApplicationInstanceTopTalkersReader{formats: a.formats},
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
	success, ok := result.(*GetEdgeApplicationInstanceTopTalkersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetEdgeApplicationInstanceTopTalkers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetEdgeApplicationInstanceTopTalkers2 gets edge application instance top talkers of the network traffic flows

  Get the top talkers of the network traffic flows of an edge application instance as reported by the edge node where the edge application instance has been deployed.
*/
func (a *Client) GetEdgeApplicationInstanceTopTalkers2(params *GetEdgeApplicationInstanceTopTalkers2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEdgeApplicationInstanceTopTalkers2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEdgeApplicationInstanceTopTalkers2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetEdgeApplicationInstanceTopTalkers2",
		Method:             "GET",
		PathPattern:        "/v1/apps/instances/name/{name}/flowlog/toptalkers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEdgeApplicationInstanceTopTalkers2Reader{formats: a.formats},
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
	success, ok := result.(*GetEdgeApplicationInstanceTopTalkers2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetEdgeApplicationInstanceTopTalkers2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetEdgeApplicationInstanceTrafficFlows gets edge application instance network traffic flow log

  Get the network traffic flow log of an edge application instance as reported by the edge node where the edge application instance has been deployed.
*/
func (a *Client) GetEdgeApplicationInstanceTrafficFlows(params *GetEdgeApplicationInstanceTrafficFlowsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEdgeApplicationInstanceTrafficFlowsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEdgeApplicationInstanceTrafficFlowsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetEdgeApplicationInstanceTrafficFlows",
		Method:             "GET",
		PathPattern:        "/v1/apps/instances/id/{id}/flowlog/classification",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEdgeApplicationInstanceTrafficFlowsReader{formats: a.formats},
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
	success, ok := result.(*GetEdgeApplicationInstanceTrafficFlowsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetEdgeApplicationInstanceTrafficFlows: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetEdgeApplicationInstanceTrafficFlows2 gets edge application instance network traffic flow log

  Get the network traffic flow log of an edge application instance as reported by the edge node where the edge application instance has been deployed.
*/
func (a *Client) GetEdgeApplicationInstanceTrafficFlows2(params *GetEdgeApplicationInstanceTrafficFlows2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEdgeApplicationInstanceTrafficFlows2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEdgeApplicationInstanceTrafficFlows2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetEdgeApplicationInstanceTrafficFlows2",
		Method:             "GET",
		PathPattern:        "/v1/apps/instances/name/{name}/flowlog/classification",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEdgeApplicationInstanceTrafficFlows2Reader{formats: a.formats},
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
	success, ok := result.(*GetEdgeApplicationInstanceTrafficFlows2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetEdgeApplicationInstanceTrafficFlows2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  QueryEdgeApplicationInstanceStatus queries status of edge application instances

  Query the status of edge application instances as reported by the edge nodes where the edge application instances have been deployed.
*/
func (a *Client) QueryEdgeApplicationInstanceStatus(params *QueryEdgeApplicationInstanceStatusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*QueryEdgeApplicationInstanceStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryEdgeApplicationInstanceStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "QueryEdgeApplicationInstanceStatus",
		Method:             "GET",
		PathPattern:        "/v1/apps/instances/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryEdgeApplicationInstanceStatusReader{formats: a.formats},
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
	success, ok := result.(*QueryEdgeApplicationInstanceStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for QueryEdgeApplicationInstanceStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

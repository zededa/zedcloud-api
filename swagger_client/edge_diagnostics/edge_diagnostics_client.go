// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package edge_diagnostics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new edge diagnostics API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for edge diagnostics API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	EdgeDiagnosticsGetDeviceTwinBootstrapConfig(params *EdgeDiagnosticsGetDeviceTwinBootstrapConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeDiagnosticsGetDeviceTwinBootstrapConfigOK, error)

	EdgeDiagnosticsGetDeviceTwinBootstrapConfigByName(params *EdgeDiagnosticsGetDeviceTwinBootstrapConfigByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeDiagnosticsGetDeviceTwinBootstrapConfigByNameOK, error)

	EdgeDiagnosticsGetDeviceTwinConfig(params *EdgeDiagnosticsGetDeviceTwinConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeDiagnosticsGetDeviceTwinConfigOK, error)

	EdgeDiagnosticsGetDeviceTwinConfigByName(params *EdgeDiagnosticsGetDeviceTwinConfigByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeDiagnosticsGetDeviceTwinConfigByNameOK, error)

	EdgeDiagnosticsGetDeviceTwinNextConfig(params *EdgeDiagnosticsGetDeviceTwinNextConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeDiagnosticsGetDeviceTwinNextConfigOK, error)

	EdgeDiagnosticsGetDeviceTwinNextConfigByName(params *EdgeDiagnosticsGetDeviceTwinNextConfigByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeDiagnosticsGetDeviceTwinNextConfigByNameOK, error)

	EdgeDiagnosticsGetDeviceTwinOfflineConfigByName(params *EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameOK, error)

	EdgeDiagnosticsGetDeviceTwinOfflineNextConfig(params *EdgeDiagnosticsGetDeviceTwinOfflineNextConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeDiagnosticsGetDeviceTwinOfflineNextConfigOK, error)

	EdgeDiagnosticsGetEventsTimeline(params *EdgeDiagnosticsGetEventsTimelineParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeDiagnosticsGetEventsTimelineOK, error)

	EdgeDiagnosticsGetResourceMetricsTimeline(params *EdgeDiagnosticsGetResourceMetricsTimelineParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeDiagnosticsGetResourceMetricsTimelineOK, error)

	EdgeDiagnosticsGetResourceMetricsTimeline2(params *EdgeDiagnosticsGetResourceMetricsTimeline2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeDiagnosticsGetResourceMetricsTimeline2OK, error)

	EdgeDiagnosticsGetTopUsers2(params *EdgeDiagnosticsGetTopUsers2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeDiagnosticsGetTopUsers2OK, error)

	EdgeDiagnosticsRegenDeviceConfig(params *EdgeDiagnosticsRegenDeviceConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeDiagnosticsRegenDeviceConfigOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
EdgeDiagnosticsGetDeviceTwinBootstrapConfig gets bootstrap device twin configuration

Get bootstrap Device twin configuration for the edge node. Doesn't change the existing edge node configuration.
*/
func (a *Client) EdgeDiagnosticsGetDeviceTwinBootstrapConfig(params *EdgeDiagnosticsGetDeviceTwinBootstrapConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeDiagnosticsGetDeviceTwinBootstrapConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeDiagnosticsGetDeviceTwinBootstrapConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeDiagnostics_GetDeviceTwinBootstrapConfig",
		Method:             "GET",
		PathPattern:        "/v1/devices/id/{id}/config/bootstrap",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EdgeDiagnosticsGetDeviceTwinBootstrapConfigReader{formats: a.formats},
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
	success, ok := result.(*EdgeDiagnosticsGetDeviceTwinBootstrapConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeDiagnosticsGetDeviceTwinBootstrapConfigDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
EdgeDiagnosticsGetDeviceTwinBootstrapConfigByName gets bootstrap device twin configuration

Get bootstrap Device twin configuration for the edge node. Doesn't change the existing edge node configuration.
*/
func (a *Client) EdgeDiagnosticsGetDeviceTwinBootstrapConfigByName(params *EdgeDiagnosticsGetDeviceTwinBootstrapConfigByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeDiagnosticsGetDeviceTwinBootstrapConfigByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeDiagnosticsGetDeviceTwinBootstrapConfigByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeDiagnostics_GetDeviceTwinBootstrapConfigByName",
		Method:             "GET",
		PathPattern:        "/v1/devices/name/{name}/config/bootstrap",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EdgeDiagnosticsGetDeviceTwinBootstrapConfigByNameReader{formats: a.formats},
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
	success, ok := result.(*EdgeDiagnosticsGetDeviceTwinBootstrapConfigByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeDiagnosticsGetDeviceTwinBootstrapConfigByNameDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
EdgeDiagnosticsGetDeviceTwinConfig gets current device twin configuration

Get currentnext Device twin configuration for the edge node. Edge node has read this configuration when it queried Cloud controller last time.
*/
func (a *Client) EdgeDiagnosticsGetDeviceTwinConfig(params *EdgeDiagnosticsGetDeviceTwinConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeDiagnosticsGetDeviceTwinConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeDiagnosticsGetDeviceTwinConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeDiagnostics_GetDeviceTwinConfig",
		Method:             "GET",
		PathPattern:        "/v1/devices/id/{id}/config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EdgeDiagnosticsGetDeviceTwinConfigReader{formats: a.formats},
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
	success, ok := result.(*EdgeDiagnosticsGetDeviceTwinConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeDiagnosticsGetDeviceTwinConfigDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
EdgeDiagnosticsGetDeviceTwinConfigByName gets current device twin configuration

Get currentnext Device twin configuration for the edge node. Edge node has read this configuration when it queried Cloud controller last time.
*/
func (a *Client) EdgeDiagnosticsGetDeviceTwinConfigByName(params *EdgeDiagnosticsGetDeviceTwinConfigByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeDiagnosticsGetDeviceTwinConfigByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeDiagnosticsGetDeviceTwinConfigByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeDiagnostics_GetDeviceTwinConfigByName",
		Method:             "GET",
		PathPattern:        "/v1/devices/name/{name}/config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EdgeDiagnosticsGetDeviceTwinConfigByNameReader{formats: a.formats},
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
	success, ok := result.(*EdgeDiagnosticsGetDeviceTwinConfigByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeDiagnosticsGetDeviceTwinConfigByNameDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
EdgeDiagnosticsGetDeviceTwinNextConfig gets next device twin configuration

Get next Device twin configuration for the edge node. Edge node will get this configuration when it queries Cloud controller next time.
*/
func (a *Client) EdgeDiagnosticsGetDeviceTwinNextConfig(params *EdgeDiagnosticsGetDeviceTwinNextConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeDiagnosticsGetDeviceTwinNextConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeDiagnosticsGetDeviceTwinNextConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeDiagnostics_GetDeviceTwinNextConfig",
		Method:             "GET",
		PathPattern:        "/v1/devices/id/{id}/config/next",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EdgeDiagnosticsGetDeviceTwinNextConfigReader{formats: a.formats},
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
	success, ok := result.(*EdgeDiagnosticsGetDeviceTwinNextConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeDiagnosticsGetDeviceTwinNextConfigDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
EdgeDiagnosticsGetDeviceTwinNextConfigByName gets next device twin configuration

Get next Device twin configuration for the edge node. Edge node will get this configuration when it queries Cloud controller next time.
*/
func (a *Client) EdgeDiagnosticsGetDeviceTwinNextConfigByName(params *EdgeDiagnosticsGetDeviceTwinNextConfigByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeDiagnosticsGetDeviceTwinNextConfigByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeDiagnosticsGetDeviceTwinNextConfigByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeDiagnostics_GetDeviceTwinNextConfigByName",
		Method:             "GET",
		PathPattern:        "/v1/devices/name/{name}/config/next",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EdgeDiagnosticsGetDeviceTwinNextConfigByNameReader{formats: a.formats},
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
	success, ok := result.(*EdgeDiagnosticsGetDeviceTwinNextConfigByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeDiagnosticsGetDeviceTwinNextConfigByNameDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
EdgeDiagnosticsGetDeviceTwinOfflineConfigByName gets offline device twin configuration

Get offline Device twin configuration for the edge node. Edge node will get this configuration when it queries Cloud controller next time.
*/
func (a *Client) EdgeDiagnosticsGetDeviceTwinOfflineConfigByName(params *EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeDiagnosticsGetDeviceTwinOfflineConfigByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeDiagnostics_GetDeviceTwinOfflineConfigByName",
		Method:             "GET",
		PathPattern:        "/v1/devices/name/{name}/config/offline",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameReader{formats: a.formats},
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
	success, ok := result.(*EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
EdgeDiagnosticsGetDeviceTwinOfflineNextConfig gets offline device twin configuration

Get offline Device twin configuration for the edge node. Edge node will get this configuration when it queries Cloud controller next time.
*/
func (a *Client) EdgeDiagnosticsGetDeviceTwinOfflineNextConfig(params *EdgeDiagnosticsGetDeviceTwinOfflineNextConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeDiagnosticsGetDeviceTwinOfflineNextConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeDiagnosticsGetDeviceTwinOfflineNextConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeDiagnostics_GetDeviceTwinOfflineNextConfig",
		Method:             "GET",
		PathPattern:        "/v1/devices/id/{id}/config/offline",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EdgeDiagnosticsGetDeviceTwinOfflineNextConfigReader{formats: a.formats},
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
	success, ok := result.(*EdgeDiagnosticsGetDeviceTwinOfflineNextConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeDiagnosticsGetDeviceTwinOfflineNextConfigDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
EdgeDiagnosticsGetEventsTimeline gets events timeline

Get aggregated events timeline
*/
func (a *Client) EdgeDiagnosticsGetEventsTimeline(params *EdgeDiagnosticsGetEventsTimelineParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeDiagnosticsGetEventsTimelineOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeDiagnosticsGetEventsTimelineParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeDiagnostics_GetEventsTimeline",
		Method:             "GET",
		PathPattern:        "/v1/events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EdgeDiagnosticsGetEventsTimelineReader{formats: a.formats},
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
	success, ok := result.(*EdgeDiagnosticsGetEventsTimelineOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeDiagnosticsGetEventsTimelineDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
EdgeDiagnosticsGetResourceMetricsTimeline gets resource usage timeline

Get the aggregated resource usage timeline as reported by the edge nodes and edge application instances.
*/
func (a *Client) EdgeDiagnosticsGetResourceMetricsTimeline(params *EdgeDiagnosticsGetResourceMetricsTimelineParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeDiagnosticsGetResourceMetricsTimelineOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeDiagnosticsGetResourceMetricsTimelineParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeDiagnostics_GetResourceMetricsTimeline",
		Method:             "GET",
		PathPattern:        "/v1/events/timeSeries/{mType}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EdgeDiagnosticsGetResourceMetricsTimelineReader{formats: a.formats},
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
	success, ok := result.(*EdgeDiagnosticsGetResourceMetricsTimelineOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeDiagnosticsGetResourceMetricsTimelineDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
EdgeDiagnosticsGetResourceMetricsTimeline2 gets resource usage timeline

Get the aggregated resource usage timeline as reported by the edge nodes and edge application instances.
*/
func (a *Client) EdgeDiagnosticsGetResourceMetricsTimeline2(params *EdgeDiagnosticsGetResourceMetricsTimeline2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeDiagnosticsGetResourceMetricsTimeline2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeDiagnosticsGetResourceMetricsTimeline2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeDiagnostics_GetResourceMetricsTimeline2",
		Method:             "GET",
		PathPattern:        "/v1/timeSeries/{mType}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EdgeDiagnosticsGetResourceMetricsTimeline2Reader{formats: a.formats},
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
	success, ok := result.(*EdgeDiagnosticsGetResourceMetricsTimeline2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeDiagnosticsGetResourceMetricsTimeline2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
EdgeDiagnosticsGetTopUsers2 gets top users

Get top users.
*/
func (a *Client) EdgeDiagnosticsGetTopUsers2(params *EdgeDiagnosticsGetTopUsers2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeDiagnosticsGetTopUsers2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeDiagnosticsGetTopUsers2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeDiagnostics_GetTopUsers2",
		Method:             "GET",
		PathPattern:        "/v1/events/topUsers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EdgeDiagnosticsGetTopUsers2Reader{formats: a.formats},
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
	success, ok := result.(*EdgeDiagnosticsGetTopUsers2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeDiagnosticsGetTopUsers2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
EdgeDiagnosticsRegenDeviceConfig res generate device configuration

Re-generate the device configuration. Edge node will get this configuration when it queries Cloud controller next time.
*/
func (a *Client) EdgeDiagnosticsRegenDeviceConfig(params *EdgeDiagnosticsRegenDeviceConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeDiagnosticsRegenDeviceConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeDiagnosticsRegenDeviceConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeDiagnostics_RegenDeviceConfig",
		Method:             "PUT",
		PathPattern:        "/v1/devices/id/{id}/config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EdgeDiagnosticsRegenDeviceConfigReader{formats: a.formats},
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
	success, ok := result.(*EdgeDiagnosticsRegenDeviceConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeDiagnosticsRegenDeviceConfigDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

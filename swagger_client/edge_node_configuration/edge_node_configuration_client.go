// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package edge_node_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new edge node configuration API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for edge node configuration API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	EdgeNodeConfigurationActivateEdgeNode(params *EdgeNodeConfigurationActivateEdgeNodeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNodeConfigurationActivateEdgeNodeOK, error)

	EdgeNodeConfigurationCreateEdgeNode(params *EdgeNodeConfigurationCreateEdgeNodeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNodeConfigurationCreateEdgeNodeOK, error)

	EdgeNodeConfigurationDeActivateEdgeNode(params *EdgeNodeConfigurationDeActivateEdgeNodeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNodeConfigurationDeActivateEdgeNodeOK, error)

	EdgeNodeConfigurationDeleteEdgeNode(params *EdgeNodeConfigurationDeleteEdgeNodeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNodeConfigurationDeleteEdgeNodeOK, error)

	EdgeNodeConfigurationGetEdgeNode(params *EdgeNodeConfigurationGetEdgeNodeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNodeConfigurationGetEdgeNodeOK, error)

	EdgeNodeConfigurationGetEdgeNodeAttestation(params *EdgeNodeConfigurationGetEdgeNodeAttestationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNodeConfigurationGetEdgeNodeAttestationOK, error)

	EdgeNodeConfigurationGetEdgeNodeByName(params *EdgeNodeConfigurationGetEdgeNodeByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNodeConfigurationGetEdgeNodeByNameOK, error)

	EdgeNodeConfigurationGetEdgeNodeBySerial(params *EdgeNodeConfigurationGetEdgeNodeBySerialParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNodeConfigurationGetEdgeNodeBySerialOK, error)

	EdgeNodeConfigurationGetEdgeNodeOnboarding(params *EdgeNodeConfigurationGetEdgeNodeOnboardingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNodeConfigurationGetEdgeNodeOnboardingOK, error)

	EdgeNodeConfigurationOffboard(params *EdgeNodeConfigurationOffboardParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNodeConfigurationOffboardOK, error)

	EdgeNodeConfigurationPreparePowerOff(params *EdgeNodeConfigurationPreparePowerOffParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNodeConfigurationPreparePowerOffOK, error)

	EdgeNodeConfigurationQueryEdgeNodes(params *EdgeNodeConfigurationQueryEdgeNodesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNodeConfigurationQueryEdgeNodesOK, error)

	EdgeNodeConfigurationReboot(params *EdgeNodeConfigurationRebootParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNodeConfigurationRebootOK, error)

	EdgeNodeConfigurationStartDebugEdgeNode(params *EdgeNodeConfigurationStartDebugEdgeNodeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNodeConfigurationStartDebugEdgeNodeOK, error)

	EdgeNodeConfigurationStopDebugEdgeNode(params *EdgeNodeConfigurationStopDebugEdgeNodeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNodeConfigurationStopDebugEdgeNodeOK, error)

	EdgeNodeConfigurationUpdateEdgeNode(params *EdgeNodeConfigurationUpdateEdgeNodeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNodeConfigurationUpdateEdgeNodeOK, error)

	EdgeNodeConfigurationUpdateEdgeNodeBaseOS(params *EdgeNodeConfigurationUpdateEdgeNodeBaseOSParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNodeConfigurationUpdateEdgeNodeBaseOSOK, error)

	EdgeNodeConfigurationUpdateEdgeNodeBaseOS2(params *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNodeConfigurationUpdateEdgeNodeBaseOS2OK, error)

	EdgeNodeConfigurationUpdateEdgeNodeBaseOS3(params *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNodeConfigurationUpdateEdgeNodeBaseOS3OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  EdgeNodeConfigurationActivateEdgeNode activates edge node

  Activate an edge node. If already in active state no action is taken.
*/
func (a *Client) EdgeNodeConfigurationActivateEdgeNode(params *EdgeNodeConfigurationActivateEdgeNodeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNodeConfigurationActivateEdgeNodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeNodeConfigurationActivateEdgeNodeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeNodeConfiguration_ActivateEdgeNode",
		Method:             "PUT",
		PathPattern:        "/v1/devices/id/{id}/activate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EdgeNodeConfigurationActivateEdgeNodeReader{formats: a.formats},
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
	success, ok := result.(*EdgeNodeConfigurationActivateEdgeNodeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeNodeConfigurationActivateEdgeNodeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  EdgeNodeConfigurationCreateEdgeNode creates edge node

  Create an edge node record.
*/
func (a *Client) EdgeNodeConfigurationCreateEdgeNode(params *EdgeNodeConfigurationCreateEdgeNodeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNodeConfigurationCreateEdgeNodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeNodeConfigurationCreateEdgeNodeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeNodeConfiguration_CreateEdgeNode",
		Method:             "POST",
		PathPattern:        "/v1/devices",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EdgeNodeConfigurationCreateEdgeNodeReader{formats: a.formats},
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
	success, ok := result.(*EdgeNodeConfigurationCreateEdgeNodeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeNodeConfigurationCreateEdgeNodeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  EdgeNodeConfigurationDeActivateEdgeNode deactivates edge node

  Deactivate an edge node. If already in inactive state no action is taken.
*/
func (a *Client) EdgeNodeConfigurationDeActivateEdgeNode(params *EdgeNodeConfigurationDeActivateEdgeNodeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNodeConfigurationDeActivateEdgeNodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeNodeConfigurationDeActivateEdgeNodeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeNodeConfiguration_DeActivateEdgeNode",
		Method:             "PUT",
		PathPattern:        "/v1/devices/id/{id}/deactivate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EdgeNodeConfigurationDeActivateEdgeNodeReader{formats: a.formats},
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
	success, ok := result.(*EdgeNodeConfigurationDeActivateEdgeNodeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeNodeConfigurationDeActivateEdgeNodeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  EdgeNodeConfigurationDeleteEdgeNode deletes edge node

  Delete an edge node record.
*/
func (a *Client) EdgeNodeConfigurationDeleteEdgeNode(params *EdgeNodeConfigurationDeleteEdgeNodeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNodeConfigurationDeleteEdgeNodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeNodeConfigurationDeleteEdgeNodeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeNodeConfiguration_DeleteEdgeNode",
		Method:             "DELETE",
		PathPattern:        "/v1/devices/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EdgeNodeConfigurationDeleteEdgeNodeReader{formats: a.formats},
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
	success, ok := result.(*EdgeNodeConfigurationDeleteEdgeNodeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeNodeConfigurationDeleteEdgeNodeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  EdgeNodeConfigurationGetEdgeNode gets edge node

  Get the configuration (without security details) of an edge node record.
*/
func (a *Client) EdgeNodeConfigurationGetEdgeNode(params *EdgeNodeConfigurationGetEdgeNodeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNodeConfigurationGetEdgeNodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeNodeConfigurationGetEdgeNodeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeNodeConfiguration_GetEdgeNode",
		Method:             "GET",
		PathPattern:        "/v1/devices/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EdgeNodeConfigurationGetEdgeNodeReader{formats: a.formats},
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
	success, ok := result.(*EdgeNodeConfigurationGetEdgeNodeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeNodeConfigurationGetEdgeNodeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  EdgeNodeConfigurationGetEdgeNodeAttestation gets edge node p c r attestation

  Get the PCR attestation of an edge node record if present.
*/
func (a *Client) EdgeNodeConfigurationGetEdgeNodeAttestation(params *EdgeNodeConfigurationGetEdgeNodeAttestationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNodeConfigurationGetEdgeNodeAttestationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeNodeConfigurationGetEdgeNodeAttestationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeNodeConfiguration_GetEdgeNodeAttestation",
		Method:             "GET",
		PathPattern:        "/v1/devices/id/{id}/attestation",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EdgeNodeConfigurationGetEdgeNodeAttestationReader{formats: a.formats},
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
	success, ok := result.(*EdgeNodeConfigurationGetEdgeNodeAttestationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeNodeConfigurationGetEdgeNodeAttestationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  EdgeNodeConfigurationGetEdgeNodeByName gets edge node by name

  Get the configuration (without security details) of an edge node record.
*/
func (a *Client) EdgeNodeConfigurationGetEdgeNodeByName(params *EdgeNodeConfigurationGetEdgeNodeByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNodeConfigurationGetEdgeNodeByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeNodeConfigurationGetEdgeNodeByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeNodeConfiguration_GetEdgeNodeByName",
		Method:             "GET",
		PathPattern:        "/v1/devices/name/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EdgeNodeConfigurationGetEdgeNodeByNameReader{formats: a.formats},
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
	success, ok := result.(*EdgeNodeConfigurationGetEdgeNodeByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeNodeConfigurationGetEdgeNodeByNameDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  EdgeNodeConfigurationGetEdgeNodeBySerial gets edge node by serial number

  Get the configuration (without security details) of an edge node record.
*/
func (a *Client) EdgeNodeConfigurationGetEdgeNodeBySerial(params *EdgeNodeConfigurationGetEdgeNodeBySerialParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNodeConfigurationGetEdgeNodeBySerialOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeNodeConfigurationGetEdgeNodeBySerialParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeNodeConfiguration_GetEdgeNodeBySerial",
		Method:             "GET",
		PathPattern:        "/v1/devices/serial/{serialno}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EdgeNodeConfigurationGetEdgeNodeBySerialReader{formats: a.formats},
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
	success, ok := result.(*EdgeNodeConfigurationGetEdgeNodeBySerialOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeNodeConfigurationGetEdgeNodeBySerialDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  EdgeNodeConfigurationGetEdgeNodeOnboarding gets edge node onboarding certificate

  Get the onboarding certificate of an edge node record if present, only valid for edge nodes that have not been onboarded.
*/
func (a *Client) EdgeNodeConfigurationGetEdgeNodeOnboarding(params *EdgeNodeConfigurationGetEdgeNodeOnboardingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNodeConfigurationGetEdgeNodeOnboardingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeNodeConfigurationGetEdgeNodeOnboardingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeNodeConfiguration_GetEdgeNodeOnboarding",
		Method:             "GET",
		PathPattern:        "/v1/devices/id/{id}/onboarding",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EdgeNodeConfigurationGetEdgeNodeOnboardingReader{formats: a.formats},
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
	success, ok := result.(*EdgeNodeConfigurationGetEdgeNodeOnboardingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeNodeConfigurationGetEdgeNodeOnboardingDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  EdgeNodeConfigurationOffboard offoards the edge ndoe

  The API is used for offbaording the device from the enterprise
*/
func (a *Client) EdgeNodeConfigurationOffboard(params *EdgeNodeConfigurationOffboardParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNodeConfigurationOffboardOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeNodeConfigurationOffboardParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeNodeConfiguration_Offboard",
		Method:             "PUT",
		PathPattern:        "/v1/devices/id/{id}/Offboard",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EdgeNodeConfigurationOffboardReader{formats: a.formats},
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
	success, ok := result.(*EdgeNodeConfigurationOffboardOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeNodeConfigurationOffboardDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  EdgeNodeConfigurationPreparePowerOff shutdowns applications on edge node

  Shutdown application instance on edge node in preparation for power off and equipment movement.
*/
func (a *Client) EdgeNodeConfigurationPreparePowerOff(params *EdgeNodeConfigurationPreparePowerOffParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNodeConfigurationPreparePowerOffOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeNodeConfigurationPreparePowerOffParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeNodeConfiguration_PreparePowerOff",
		Method:             "PUT",
		PathPattern:        "/v1/devices/id/{id}/preparepoweroff",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EdgeNodeConfigurationPreparePowerOffReader{formats: a.formats},
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
	success, ok := result.(*EdgeNodeConfigurationPreparePowerOffOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeNodeConfigurationPreparePowerOffDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  EdgeNodeConfigurationQueryEdgeNodes queries edge nodes

  Query the edge node records.
*/
func (a *Client) EdgeNodeConfigurationQueryEdgeNodes(params *EdgeNodeConfigurationQueryEdgeNodesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNodeConfigurationQueryEdgeNodesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeNodeConfigurationQueryEdgeNodesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeNodeConfiguration_QueryEdgeNodes",
		Method:             "GET",
		PathPattern:        "/v1/devices",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EdgeNodeConfigurationQueryEdgeNodesReader{formats: a.formats},
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
	success, ok := result.(*EdgeNodeConfigurationQueryEdgeNodesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeNodeConfigurationQueryEdgeNodesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  EdgeNodeConfigurationReboot reboots edge node

  Reboot an edge node remotely. This operation triggers reboot event. Please check bla-bla for device reboot status.
*/
func (a *Client) EdgeNodeConfigurationReboot(params *EdgeNodeConfigurationRebootParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNodeConfigurationRebootOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeNodeConfigurationRebootParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeNodeConfiguration_Reboot",
		Method:             "PUT",
		PathPattern:        "/v1/devices/id/{id}/reboot",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EdgeNodeConfigurationRebootReader{formats: a.formats},
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
	success, ok := result.(*EdgeNodeConfigurationRebootOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeNodeConfigurationRebootDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  EdgeNodeConfigurationStartDebugEdgeNode starts debug mode on the edge node

  Start debug mode on the edge node.
*/
func (a *Client) EdgeNodeConfigurationStartDebugEdgeNode(params *EdgeNodeConfigurationStartDebugEdgeNodeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNodeConfigurationStartDebugEdgeNodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeNodeConfigurationStartDebugEdgeNodeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeNodeConfiguration_StartDebugEdgeNode",
		Method:             "PUT",
		PathPattern:        "/v1/devices/id/{id}/debug/enable",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EdgeNodeConfigurationStartDebugEdgeNodeReader{formats: a.formats},
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
	success, ok := result.(*EdgeNodeConfigurationStartDebugEdgeNodeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeNodeConfigurationStartDebugEdgeNodeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  EdgeNodeConfigurationStopDebugEdgeNode stops debug mode on the edge node

  Stop debug mode on the edge node.
*/
func (a *Client) EdgeNodeConfigurationStopDebugEdgeNode(params *EdgeNodeConfigurationStopDebugEdgeNodeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNodeConfigurationStopDebugEdgeNodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeNodeConfigurationStopDebugEdgeNodeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeNodeConfiguration_StopDebugEdgeNode",
		Method:             "PUT",
		PathPattern:        "/v1/devices/id/{id}/debug/disable",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EdgeNodeConfigurationStopDebugEdgeNodeReader{formats: a.formats},
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
	success, ok := result.(*EdgeNodeConfigurationStopDebugEdgeNodeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeNodeConfigurationStopDebugEdgeNodeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  EdgeNodeConfigurationUpdateEdgeNode updates edge node

  Update an edge node record. The usual pattern to update an edge node record is to retrieve the record and update with the modified values in a new body to update the edge node record.
*/
func (a *Client) EdgeNodeConfigurationUpdateEdgeNode(params *EdgeNodeConfigurationUpdateEdgeNodeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNodeConfigurationUpdateEdgeNodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeNodeConfigurationUpdateEdgeNodeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeNodeConfiguration_UpdateEdgeNode",
		Method:             "PUT",
		PathPattern:        "/v1/devices/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EdgeNodeConfigurationUpdateEdgeNodeReader{formats: a.formats},
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
	success, ok := result.(*EdgeNodeConfigurationUpdateEdgeNodeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeNodeConfigurationUpdateEdgeNodeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  EdgeNodeConfigurationUpdateEdgeNodeBaseOS updates e v e image on edge node

  Update EVE image on edge node, if edge node is already running the latest EVE image no action is taken.
*/
func (a *Client) EdgeNodeConfigurationUpdateEdgeNodeBaseOS(params *EdgeNodeConfigurationUpdateEdgeNodeBaseOSParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNodeConfigurationUpdateEdgeNodeBaseOSOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeNodeConfigurationUpdateEdgeNodeBaseOSParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeNodeConfiguration_UpdateEdgeNodeBaseOS",
		Method:             "PUT",
		PathPattern:        "/v1/devices/id/{id}/apply",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EdgeNodeConfigurationUpdateEdgeNodeBaseOSReader{formats: a.formats},
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
	success, ok := result.(*EdgeNodeConfigurationUpdateEdgeNodeBaseOSOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeNodeConfigurationUpdateEdgeNodeBaseOSDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  EdgeNodeConfigurationUpdateEdgeNodeBaseOS2 updates e v e image on edge node

  Update EVE image on edge node, if edge node is already running the latest EVE image no action is taken.
*/
func (a *Client) EdgeNodeConfigurationUpdateEdgeNodeBaseOS2(params *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNodeConfigurationUpdateEdgeNodeBaseOS2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeNodeConfiguration_UpdateEdgeNodeBaseOS2",
		Method:             "PUT",
		PathPattern:        "/v1/devices/id/{id}/publish",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Reader{formats: a.formats},
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
	success, ok := result.(*EdgeNodeConfigurationUpdateEdgeNodeBaseOS2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  EdgeNodeConfigurationUpdateEdgeNodeBaseOS3 updates e v e image on edge node

  Update EVE image on edge node, if edge node is already running the latest EVE image no action is taken.
*/
func (a *Client) EdgeNodeConfigurationUpdateEdgeNodeBaseOS3(params *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNodeConfigurationUpdateEdgeNodeBaseOS3OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS3Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeNodeConfiguration_UpdateEdgeNodeBaseOS3",
		Method:             "PUT",
		PathPattern:        "/v1/devices/id/{id}/unpublish",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Reader{formats: a.formats},
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
	success, ok := result.(*EdgeNodeConfigurationUpdateEdgeNodeBaseOS3OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

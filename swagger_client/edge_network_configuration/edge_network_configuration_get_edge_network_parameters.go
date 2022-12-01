// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package edge_network_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewEdgeNetworkConfigurationGetEdgeNetworkParams creates a new EdgeNetworkConfigurationGetEdgeNetworkParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEdgeNetworkConfigurationGetEdgeNetworkParams() *EdgeNetworkConfigurationGetEdgeNetworkParams {
	return &EdgeNetworkConfigurationGetEdgeNetworkParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeNetworkConfigurationGetEdgeNetworkParamsWithTimeout creates a new EdgeNetworkConfigurationGetEdgeNetworkParams object
// with the ability to set a timeout on a request.
func NewEdgeNetworkConfigurationGetEdgeNetworkParamsWithTimeout(timeout time.Duration) *EdgeNetworkConfigurationGetEdgeNetworkParams {
	return &EdgeNetworkConfigurationGetEdgeNetworkParams{
		timeout: timeout,
	}
}

// NewEdgeNetworkConfigurationGetEdgeNetworkParamsWithContext creates a new EdgeNetworkConfigurationGetEdgeNetworkParams object
// with the ability to set a context for a request.
func NewEdgeNetworkConfigurationGetEdgeNetworkParamsWithContext(ctx context.Context) *EdgeNetworkConfigurationGetEdgeNetworkParams {
	return &EdgeNetworkConfigurationGetEdgeNetworkParams{
		Context: ctx,
	}
}

// NewEdgeNetworkConfigurationGetEdgeNetworkParamsWithHTTPClient creates a new EdgeNetworkConfigurationGetEdgeNetworkParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeNetworkConfigurationGetEdgeNetworkParamsWithHTTPClient(client *http.Client) *EdgeNetworkConfigurationGetEdgeNetworkParams {
	return &EdgeNetworkConfigurationGetEdgeNetworkParams{
		HTTPClient: client,
	}
}

/*
EdgeNetworkConfigurationGetEdgeNetworkParams contains all the parameters to send to the API endpoint

	for the edge network configuration get edge network operation.

	Typically these are written to a http.Request.
*/
type EdgeNetworkConfigurationGetEdgeNetworkParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* ID.

	   System defined universally unique Id of the network
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edge network configuration get edge network params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNetworkConfigurationGetEdgeNetworkParams) WithDefaults() *EdgeNetworkConfigurationGetEdgeNetworkParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge network configuration get edge network params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNetworkConfigurationGetEdgeNetworkParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge network configuration get edge network params
func (o *EdgeNetworkConfigurationGetEdgeNetworkParams) WithTimeout(timeout time.Duration) *EdgeNetworkConfigurationGetEdgeNetworkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge network configuration get edge network params
func (o *EdgeNetworkConfigurationGetEdgeNetworkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge network configuration get edge network params
func (o *EdgeNetworkConfigurationGetEdgeNetworkParams) WithContext(ctx context.Context) *EdgeNetworkConfigurationGetEdgeNetworkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge network configuration get edge network params
func (o *EdgeNetworkConfigurationGetEdgeNetworkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge network configuration get edge network params
func (o *EdgeNetworkConfigurationGetEdgeNetworkParams) WithHTTPClient(client *http.Client) *EdgeNetworkConfigurationGetEdgeNetworkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge network configuration get edge network params
func (o *EdgeNetworkConfigurationGetEdgeNetworkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the edge network configuration get edge network params
func (o *EdgeNetworkConfigurationGetEdgeNetworkParams) WithXRequestID(xRequestID *string) *EdgeNetworkConfigurationGetEdgeNetworkParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the edge network configuration get edge network params
func (o *EdgeNetworkConfigurationGetEdgeNetworkParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the edge network configuration get edge network params
func (o *EdgeNetworkConfigurationGetEdgeNetworkParams) WithID(id string) *EdgeNetworkConfigurationGetEdgeNetworkParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the edge network configuration get edge network params
func (o *EdgeNetworkConfigurationGetEdgeNetworkParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeNetworkConfigurationGetEdgeNetworkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

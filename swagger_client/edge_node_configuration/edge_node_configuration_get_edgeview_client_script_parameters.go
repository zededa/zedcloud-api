// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package edge_node_configuration

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

// NewEdgeNodeConfigurationGetEdgeviewClientScriptParams creates a new EdgeNodeConfigurationGetEdgeviewClientScriptParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEdgeNodeConfigurationGetEdgeviewClientScriptParams() *EdgeNodeConfigurationGetEdgeviewClientScriptParams {
	return &EdgeNodeConfigurationGetEdgeviewClientScriptParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeNodeConfigurationGetEdgeviewClientScriptParamsWithTimeout creates a new EdgeNodeConfigurationGetEdgeviewClientScriptParams object
// with the ability to set a timeout on a request.
func NewEdgeNodeConfigurationGetEdgeviewClientScriptParamsWithTimeout(timeout time.Duration) *EdgeNodeConfigurationGetEdgeviewClientScriptParams {
	return &EdgeNodeConfigurationGetEdgeviewClientScriptParams{
		timeout: timeout,
	}
}

// NewEdgeNodeConfigurationGetEdgeviewClientScriptParamsWithContext creates a new EdgeNodeConfigurationGetEdgeviewClientScriptParams object
// with the ability to set a context for a request.
func NewEdgeNodeConfigurationGetEdgeviewClientScriptParamsWithContext(ctx context.Context) *EdgeNodeConfigurationGetEdgeviewClientScriptParams {
	return &EdgeNodeConfigurationGetEdgeviewClientScriptParams{
		Context: ctx,
	}
}

// NewEdgeNodeConfigurationGetEdgeviewClientScriptParamsWithHTTPClient creates a new EdgeNodeConfigurationGetEdgeviewClientScriptParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeNodeConfigurationGetEdgeviewClientScriptParamsWithHTTPClient(client *http.Client) *EdgeNodeConfigurationGetEdgeviewClientScriptParams {
	return &EdgeNodeConfigurationGetEdgeviewClientScriptParams{
		HTTPClient: client,
	}
}

/*
EdgeNodeConfigurationGetEdgeviewClientScriptParams contains all the parameters to send to the API endpoint

	for the edge node configuration get edgeview client script operation.

	Typically these are written to a http.Request.
*/
type EdgeNodeConfigurationGetEdgeviewClientScriptParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* ID.

	   system generated unique id for a device
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edge node configuration get edgeview client script params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNodeConfigurationGetEdgeviewClientScriptParams) WithDefaults() *EdgeNodeConfigurationGetEdgeviewClientScriptParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge node configuration get edgeview client script params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNodeConfigurationGetEdgeviewClientScriptParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge node configuration get edgeview client script params
func (o *EdgeNodeConfigurationGetEdgeviewClientScriptParams) WithTimeout(timeout time.Duration) *EdgeNodeConfigurationGetEdgeviewClientScriptParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge node configuration get edgeview client script params
func (o *EdgeNodeConfigurationGetEdgeviewClientScriptParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge node configuration get edgeview client script params
func (o *EdgeNodeConfigurationGetEdgeviewClientScriptParams) WithContext(ctx context.Context) *EdgeNodeConfigurationGetEdgeviewClientScriptParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge node configuration get edgeview client script params
func (o *EdgeNodeConfigurationGetEdgeviewClientScriptParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge node configuration get edgeview client script params
func (o *EdgeNodeConfigurationGetEdgeviewClientScriptParams) WithHTTPClient(client *http.Client) *EdgeNodeConfigurationGetEdgeviewClientScriptParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge node configuration get edgeview client script params
func (o *EdgeNodeConfigurationGetEdgeviewClientScriptParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the edge node configuration get edgeview client script params
func (o *EdgeNodeConfigurationGetEdgeviewClientScriptParams) WithXRequestID(xRequestID *string) *EdgeNodeConfigurationGetEdgeviewClientScriptParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the edge node configuration get edgeview client script params
func (o *EdgeNodeConfigurationGetEdgeviewClientScriptParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the edge node configuration get edgeview client script params
func (o *EdgeNodeConfigurationGetEdgeviewClientScriptParams) WithID(id string) *EdgeNodeConfigurationGetEdgeviewClientScriptParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the edge node configuration get edgeview client script params
func (o *EdgeNodeConfigurationGetEdgeviewClientScriptParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeNodeConfigurationGetEdgeviewClientScriptParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

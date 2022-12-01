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

// NewEdgeNodeConfigurationStartDebugEdgeNodeParams creates a new EdgeNodeConfigurationStartDebugEdgeNodeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEdgeNodeConfigurationStartDebugEdgeNodeParams() *EdgeNodeConfigurationStartDebugEdgeNodeParams {
	return &EdgeNodeConfigurationStartDebugEdgeNodeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeNodeConfigurationStartDebugEdgeNodeParamsWithTimeout creates a new EdgeNodeConfigurationStartDebugEdgeNodeParams object
// with the ability to set a timeout on a request.
func NewEdgeNodeConfigurationStartDebugEdgeNodeParamsWithTimeout(timeout time.Duration) *EdgeNodeConfigurationStartDebugEdgeNodeParams {
	return &EdgeNodeConfigurationStartDebugEdgeNodeParams{
		timeout: timeout,
	}
}

// NewEdgeNodeConfigurationStartDebugEdgeNodeParamsWithContext creates a new EdgeNodeConfigurationStartDebugEdgeNodeParams object
// with the ability to set a context for a request.
func NewEdgeNodeConfigurationStartDebugEdgeNodeParamsWithContext(ctx context.Context) *EdgeNodeConfigurationStartDebugEdgeNodeParams {
	return &EdgeNodeConfigurationStartDebugEdgeNodeParams{
		Context: ctx,
	}
}

// NewEdgeNodeConfigurationStartDebugEdgeNodeParamsWithHTTPClient creates a new EdgeNodeConfigurationStartDebugEdgeNodeParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeNodeConfigurationStartDebugEdgeNodeParamsWithHTTPClient(client *http.Client) *EdgeNodeConfigurationStartDebugEdgeNodeParams {
	return &EdgeNodeConfigurationStartDebugEdgeNodeParams{
		HTTPClient: client,
	}
}

/*
EdgeNodeConfigurationStartDebugEdgeNodeParams contains all the parameters to send to the API endpoint

	for the edge node configuration start debug edge node operation.

	Typically these are written to a http.Request.
*/
type EdgeNodeConfigurationStartDebugEdgeNodeParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body EdgeNodeConfigurationStartDebugEdgeNodeBody

	/* ID.

	   system generated unique id for a device
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edge node configuration start debug edge node params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNodeConfigurationStartDebugEdgeNodeParams) WithDefaults() *EdgeNodeConfigurationStartDebugEdgeNodeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge node configuration start debug edge node params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNodeConfigurationStartDebugEdgeNodeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge node configuration start debug edge node params
func (o *EdgeNodeConfigurationStartDebugEdgeNodeParams) WithTimeout(timeout time.Duration) *EdgeNodeConfigurationStartDebugEdgeNodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge node configuration start debug edge node params
func (o *EdgeNodeConfigurationStartDebugEdgeNodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge node configuration start debug edge node params
func (o *EdgeNodeConfigurationStartDebugEdgeNodeParams) WithContext(ctx context.Context) *EdgeNodeConfigurationStartDebugEdgeNodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge node configuration start debug edge node params
func (o *EdgeNodeConfigurationStartDebugEdgeNodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge node configuration start debug edge node params
func (o *EdgeNodeConfigurationStartDebugEdgeNodeParams) WithHTTPClient(client *http.Client) *EdgeNodeConfigurationStartDebugEdgeNodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge node configuration start debug edge node params
func (o *EdgeNodeConfigurationStartDebugEdgeNodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the edge node configuration start debug edge node params
func (o *EdgeNodeConfigurationStartDebugEdgeNodeParams) WithXRequestID(xRequestID *string) *EdgeNodeConfigurationStartDebugEdgeNodeParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the edge node configuration start debug edge node params
func (o *EdgeNodeConfigurationStartDebugEdgeNodeParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the edge node configuration start debug edge node params
func (o *EdgeNodeConfigurationStartDebugEdgeNodeParams) WithBody(body EdgeNodeConfigurationStartDebugEdgeNodeBody) *EdgeNodeConfigurationStartDebugEdgeNodeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the edge node configuration start debug edge node params
func (o *EdgeNodeConfigurationStartDebugEdgeNodeParams) SetBody(body EdgeNodeConfigurationStartDebugEdgeNodeBody) {
	o.Body = body
}

// WithID adds the id to the edge node configuration start debug edge node params
func (o *EdgeNodeConfigurationStartDebugEdgeNodeParams) WithID(id string) *EdgeNodeConfigurationStartDebugEdgeNodeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the edge node configuration start debug edge node params
func (o *EdgeNodeConfigurationStartDebugEdgeNodeParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeNodeConfigurationStartDebugEdgeNodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
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

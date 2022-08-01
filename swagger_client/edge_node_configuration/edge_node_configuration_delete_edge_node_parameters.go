// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

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

// NewEdgeNodeConfigurationDeleteEdgeNodeParams creates a new EdgeNodeConfigurationDeleteEdgeNodeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEdgeNodeConfigurationDeleteEdgeNodeParams() *EdgeNodeConfigurationDeleteEdgeNodeParams {
	return &EdgeNodeConfigurationDeleteEdgeNodeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeNodeConfigurationDeleteEdgeNodeParamsWithTimeout creates a new EdgeNodeConfigurationDeleteEdgeNodeParams object
// with the ability to set a timeout on a request.
func NewEdgeNodeConfigurationDeleteEdgeNodeParamsWithTimeout(timeout time.Duration) *EdgeNodeConfigurationDeleteEdgeNodeParams {
	return &EdgeNodeConfigurationDeleteEdgeNodeParams{
		timeout: timeout,
	}
}

// NewEdgeNodeConfigurationDeleteEdgeNodeParamsWithContext creates a new EdgeNodeConfigurationDeleteEdgeNodeParams object
// with the ability to set a context for a request.
func NewEdgeNodeConfigurationDeleteEdgeNodeParamsWithContext(ctx context.Context) *EdgeNodeConfigurationDeleteEdgeNodeParams {
	return &EdgeNodeConfigurationDeleteEdgeNodeParams{
		Context: ctx,
	}
}

// NewEdgeNodeConfigurationDeleteEdgeNodeParamsWithHTTPClient creates a new EdgeNodeConfigurationDeleteEdgeNodeParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeNodeConfigurationDeleteEdgeNodeParamsWithHTTPClient(client *http.Client) *EdgeNodeConfigurationDeleteEdgeNodeParams {
	return &EdgeNodeConfigurationDeleteEdgeNodeParams{
		HTTPClient: client,
	}
}

/* EdgeNodeConfigurationDeleteEdgeNodeParams contains all the parameters to send to the API endpoint
   for the edge node configuration delete edge node operation.

   Typically these are written to a http.Request.
*/
type EdgeNodeConfigurationDeleteEdgeNodeParams struct {

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

// WithDefaults hydrates default values in the edge node configuration delete edge node params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNodeConfigurationDeleteEdgeNodeParams) WithDefaults() *EdgeNodeConfigurationDeleteEdgeNodeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge node configuration delete edge node params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNodeConfigurationDeleteEdgeNodeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge node configuration delete edge node params
func (o *EdgeNodeConfigurationDeleteEdgeNodeParams) WithTimeout(timeout time.Duration) *EdgeNodeConfigurationDeleteEdgeNodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge node configuration delete edge node params
func (o *EdgeNodeConfigurationDeleteEdgeNodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge node configuration delete edge node params
func (o *EdgeNodeConfigurationDeleteEdgeNodeParams) WithContext(ctx context.Context) *EdgeNodeConfigurationDeleteEdgeNodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge node configuration delete edge node params
func (o *EdgeNodeConfigurationDeleteEdgeNodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge node configuration delete edge node params
func (o *EdgeNodeConfigurationDeleteEdgeNodeParams) WithHTTPClient(client *http.Client) *EdgeNodeConfigurationDeleteEdgeNodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge node configuration delete edge node params
func (o *EdgeNodeConfigurationDeleteEdgeNodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the edge node configuration delete edge node params
func (o *EdgeNodeConfigurationDeleteEdgeNodeParams) WithXRequestID(xRequestID *string) *EdgeNodeConfigurationDeleteEdgeNodeParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the edge node configuration delete edge node params
func (o *EdgeNodeConfigurationDeleteEdgeNodeParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the edge node configuration delete edge node params
func (o *EdgeNodeConfigurationDeleteEdgeNodeParams) WithID(id string) *EdgeNodeConfigurationDeleteEdgeNodeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the edge node configuration delete edge node params
func (o *EdgeNodeConfigurationDeleteEdgeNodeParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeNodeConfigurationDeleteEdgeNodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
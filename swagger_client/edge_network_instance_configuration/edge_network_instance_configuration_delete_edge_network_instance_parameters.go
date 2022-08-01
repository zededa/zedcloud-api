// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package edge_network_instance_configuration

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

// NewEdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceParams creates a new EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceParams() *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceParams {
	return &EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceParamsWithTimeout creates a new EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceParams object
// with the ability to set a timeout on a request.
func NewEdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceParamsWithTimeout(timeout time.Duration) *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceParams {
	return &EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceParams{
		timeout: timeout,
	}
}

// NewEdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceParamsWithContext creates a new EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceParams object
// with the ability to set a context for a request.
func NewEdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceParamsWithContext(ctx context.Context) *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceParams {
	return &EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceParams{
		Context: ctx,
	}
}

// NewEdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceParamsWithHTTPClient creates a new EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceParamsWithHTTPClient(client *http.Client) *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceParams {
	return &EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceParams{
		HTTPClient: client,
	}
}

/* EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceParams contains all the parameters to send to the API endpoint
   for the edge network instance configuration delete edge network instance operation.

   Typically these are written to a http.Request.
*/
type EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* ID.

	   System defined universally unique Id of the network instance
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edge network instance configuration delete edge network instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceParams) WithDefaults() *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge network instance configuration delete edge network instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge network instance configuration delete edge network instance params
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceParams) WithTimeout(timeout time.Duration) *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge network instance configuration delete edge network instance params
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge network instance configuration delete edge network instance params
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceParams) WithContext(ctx context.Context) *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge network instance configuration delete edge network instance params
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge network instance configuration delete edge network instance params
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceParams) WithHTTPClient(client *http.Client) *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge network instance configuration delete edge network instance params
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the edge network instance configuration delete edge network instance params
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceParams) WithXRequestID(xRequestID *string) *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the edge network instance configuration delete edge network instance params
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the edge network instance configuration delete edge network instance params
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceParams) WithID(id string) *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the edge network instance configuration delete edge network instance params
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeNetworkInstanceConfigurationDeleteEdgeNetworkInstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package edge_network_instance_status

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

// NewEdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusParams creates a new EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusParams() *EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusParams {
	return &EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusParamsWithTimeout creates a new EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusParams object
// with the ability to set a timeout on a request.
func NewEdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusParamsWithTimeout(timeout time.Duration) *EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusParams {
	return &EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusParams{
		timeout: timeout,
	}
}

// NewEdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusParamsWithContext creates a new EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusParams object
// with the ability to set a context for a request.
func NewEdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusParamsWithContext(ctx context.Context) *EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusParams {
	return &EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusParams{
		Context: ctx,
	}
}

// NewEdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusParamsWithHTTPClient creates a new EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusParamsWithHTTPClient(client *http.Client) *EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusParams {
	return &EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusParams{
		HTTPClient: client,
	}
}

/*
EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusParams contains all the parameters to send to the API endpoint

	for the edge network instance status get edge network instance status operation.

	Typically these are written to a http.Request.
*/
type EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edge network instance status get edge network instance status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusParams) WithDefaults() *EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge network instance status get edge network instance status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge network instance status get edge network instance status params
func (o *EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusParams) WithTimeout(timeout time.Duration) *EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge network instance status get edge network instance status params
func (o *EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge network instance status get edge network instance status params
func (o *EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusParams) WithContext(ctx context.Context) *EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge network instance status get edge network instance status params
func (o *EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge network instance status get edge network instance status params
func (o *EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusParams) WithHTTPClient(client *http.Client) *EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge network instance status get edge network instance status params
func (o *EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the edge network instance status get edge network instance status params
func (o *EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusParams) WithXRequestID(xRequestID *string) *EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the edge network instance status get edge network instance status params
func (o *EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the edge network instance status get edge network instance status params
func (o *EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusParams) WithID(id string) *EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the edge network instance status get edge network instance status params
func (o *EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

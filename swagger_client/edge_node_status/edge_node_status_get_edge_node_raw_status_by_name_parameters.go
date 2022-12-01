// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package edge_node_status

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

// NewEdgeNodeStatusGetEdgeNodeRawStatusByNameParams creates a new EdgeNodeStatusGetEdgeNodeRawStatusByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEdgeNodeStatusGetEdgeNodeRawStatusByNameParams() *EdgeNodeStatusGetEdgeNodeRawStatusByNameParams {
	return &EdgeNodeStatusGetEdgeNodeRawStatusByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeNodeStatusGetEdgeNodeRawStatusByNameParamsWithTimeout creates a new EdgeNodeStatusGetEdgeNodeRawStatusByNameParams object
// with the ability to set a timeout on a request.
func NewEdgeNodeStatusGetEdgeNodeRawStatusByNameParamsWithTimeout(timeout time.Duration) *EdgeNodeStatusGetEdgeNodeRawStatusByNameParams {
	return &EdgeNodeStatusGetEdgeNodeRawStatusByNameParams{
		timeout: timeout,
	}
}

// NewEdgeNodeStatusGetEdgeNodeRawStatusByNameParamsWithContext creates a new EdgeNodeStatusGetEdgeNodeRawStatusByNameParams object
// with the ability to set a context for a request.
func NewEdgeNodeStatusGetEdgeNodeRawStatusByNameParamsWithContext(ctx context.Context) *EdgeNodeStatusGetEdgeNodeRawStatusByNameParams {
	return &EdgeNodeStatusGetEdgeNodeRawStatusByNameParams{
		Context: ctx,
	}
}

// NewEdgeNodeStatusGetEdgeNodeRawStatusByNameParamsWithHTTPClient creates a new EdgeNodeStatusGetEdgeNodeRawStatusByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeNodeStatusGetEdgeNodeRawStatusByNameParamsWithHTTPClient(client *http.Client) *EdgeNodeStatusGetEdgeNodeRawStatusByNameParams {
	return &EdgeNodeStatusGetEdgeNodeRawStatusByNameParams{
		HTTPClient: client,
	}
}

/*
EdgeNodeStatusGetEdgeNodeRawStatusByNameParams contains all the parameters to send to the API endpoint

	for the edge node status get edge node raw status by name operation.

	Typically these are written to a http.Request.
*/
type EdgeNodeStatusGetEdgeNodeRawStatusByNameParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Name.
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edge node status get edge node raw status by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameParams) WithDefaults() *EdgeNodeStatusGetEdgeNodeRawStatusByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge node status get edge node raw status by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge node status get edge node raw status by name params
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameParams) WithTimeout(timeout time.Duration) *EdgeNodeStatusGetEdgeNodeRawStatusByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge node status get edge node raw status by name params
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge node status get edge node raw status by name params
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameParams) WithContext(ctx context.Context) *EdgeNodeStatusGetEdgeNodeRawStatusByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge node status get edge node raw status by name params
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge node status get edge node raw status by name params
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameParams) WithHTTPClient(client *http.Client) *EdgeNodeStatusGetEdgeNodeRawStatusByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge node status get edge node raw status by name params
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the edge node status get edge node raw status by name params
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameParams) WithXRequestID(xRequestID *string) *EdgeNodeStatusGetEdgeNodeRawStatusByNameParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the edge node status get edge node raw status by name params
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithName adds the name to the edge node status get edge node raw status by name params
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameParams) WithName(name string) *EdgeNodeStatusGetEdgeNodeRawStatusByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the edge node status get edge node raw status by name params
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeNodeStatusGetEdgeNodeRawStatusByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

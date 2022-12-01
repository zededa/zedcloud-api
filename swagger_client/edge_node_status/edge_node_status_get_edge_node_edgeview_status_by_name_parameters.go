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

// NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameParams creates a new EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameParams() *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameParams {
	return &EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameParamsWithTimeout creates a new EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameParams object
// with the ability to set a timeout on a request.
func NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameParamsWithTimeout(timeout time.Duration) *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameParams {
	return &EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameParams{
		timeout: timeout,
	}
}

// NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameParamsWithContext creates a new EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameParams object
// with the ability to set a context for a request.
func NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameParamsWithContext(ctx context.Context) *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameParams {
	return &EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameParams{
		Context: ctx,
	}
}

// NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameParamsWithHTTPClient creates a new EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameParamsWithHTTPClient(client *http.Client) *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameParams {
	return &EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameParams{
		HTTPClient: client,
	}
}

/*
EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameParams contains all the parameters to send to the API endpoint

	for the edge node status get edge node edgeview status by name operation.

	Typically these are written to a http.Request.
*/
type EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameParams struct {

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

// WithDefaults hydrates default values in the edge node status get edge node edgeview status by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameParams) WithDefaults() *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge node status get edge node edgeview status by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge node status get edge node edgeview status by name params
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameParams) WithTimeout(timeout time.Duration) *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge node status get edge node edgeview status by name params
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge node status get edge node edgeview status by name params
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameParams) WithContext(ctx context.Context) *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge node status get edge node edgeview status by name params
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge node status get edge node edgeview status by name params
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameParams) WithHTTPClient(client *http.Client) *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge node status get edge node edgeview status by name params
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the edge node status get edge node edgeview status by name params
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameParams) WithXRequestID(xRequestID *string) *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the edge node status get edge node edgeview status by name params
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithName adds the name to the edge node status get edge node edgeview status by name params
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameParams) WithName(name string) *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the edge node status get edge node edgeview status by name params
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
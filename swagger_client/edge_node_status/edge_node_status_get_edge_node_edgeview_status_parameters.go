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

// NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusParams creates a new EdgeNodeStatusGetEdgeNodeEdgeviewStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusParams() *EdgeNodeStatusGetEdgeNodeEdgeviewStatusParams {
	return &EdgeNodeStatusGetEdgeNodeEdgeviewStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusParamsWithTimeout creates a new EdgeNodeStatusGetEdgeNodeEdgeviewStatusParams object
// with the ability to set a timeout on a request.
func NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusParamsWithTimeout(timeout time.Duration) *EdgeNodeStatusGetEdgeNodeEdgeviewStatusParams {
	return &EdgeNodeStatusGetEdgeNodeEdgeviewStatusParams{
		timeout: timeout,
	}
}

// NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusParamsWithContext creates a new EdgeNodeStatusGetEdgeNodeEdgeviewStatusParams object
// with the ability to set a context for a request.
func NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusParamsWithContext(ctx context.Context) *EdgeNodeStatusGetEdgeNodeEdgeviewStatusParams {
	return &EdgeNodeStatusGetEdgeNodeEdgeviewStatusParams{
		Context: ctx,
	}
}

// NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusParamsWithHTTPClient creates a new EdgeNodeStatusGetEdgeNodeEdgeviewStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusParamsWithHTTPClient(client *http.Client) *EdgeNodeStatusGetEdgeNodeEdgeviewStatusParams {
	return &EdgeNodeStatusGetEdgeNodeEdgeviewStatusParams{
		HTTPClient: client,
	}
}

/*
EdgeNodeStatusGetEdgeNodeEdgeviewStatusParams contains all the parameters to send to the API endpoint

	for the edge node status get edge node edgeview status operation.

	Typically these are written to a http.Request.
*/
type EdgeNodeStatusGetEdgeNodeEdgeviewStatusParams struct {

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

// WithDefaults hydrates default values in the edge node status get edge node edgeview status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusParams) WithDefaults() *EdgeNodeStatusGetEdgeNodeEdgeviewStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge node status get edge node edgeview status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge node status get edge node edgeview status params
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusParams) WithTimeout(timeout time.Duration) *EdgeNodeStatusGetEdgeNodeEdgeviewStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge node status get edge node edgeview status params
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge node status get edge node edgeview status params
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusParams) WithContext(ctx context.Context) *EdgeNodeStatusGetEdgeNodeEdgeviewStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge node status get edge node edgeview status params
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge node status get edge node edgeview status params
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusParams) WithHTTPClient(client *http.Client) *EdgeNodeStatusGetEdgeNodeEdgeviewStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge node status get edge node edgeview status params
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the edge node status get edge node edgeview status params
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusParams) WithXRequestID(xRequestID *string) *EdgeNodeStatusGetEdgeNodeEdgeviewStatusParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the edge node status get edge node edgeview status params
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the edge node status get edge node edgeview status params
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusParams) WithID(id string) *EdgeNodeStatusGetEdgeNodeEdgeviewStatusParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the edge node status get edge node edgeview status params
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
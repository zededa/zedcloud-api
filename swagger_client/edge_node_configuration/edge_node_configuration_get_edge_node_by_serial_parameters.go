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

// NewEdgeNodeConfigurationGetEdgeNodeBySerialParams creates a new EdgeNodeConfigurationGetEdgeNodeBySerialParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEdgeNodeConfigurationGetEdgeNodeBySerialParams() *EdgeNodeConfigurationGetEdgeNodeBySerialParams {
	return &EdgeNodeConfigurationGetEdgeNodeBySerialParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeNodeConfigurationGetEdgeNodeBySerialParamsWithTimeout creates a new EdgeNodeConfigurationGetEdgeNodeBySerialParams object
// with the ability to set a timeout on a request.
func NewEdgeNodeConfigurationGetEdgeNodeBySerialParamsWithTimeout(timeout time.Duration) *EdgeNodeConfigurationGetEdgeNodeBySerialParams {
	return &EdgeNodeConfigurationGetEdgeNodeBySerialParams{
		timeout: timeout,
	}
}

// NewEdgeNodeConfigurationGetEdgeNodeBySerialParamsWithContext creates a new EdgeNodeConfigurationGetEdgeNodeBySerialParams object
// with the ability to set a context for a request.
func NewEdgeNodeConfigurationGetEdgeNodeBySerialParamsWithContext(ctx context.Context) *EdgeNodeConfigurationGetEdgeNodeBySerialParams {
	return &EdgeNodeConfigurationGetEdgeNodeBySerialParams{
		Context: ctx,
	}
}

// NewEdgeNodeConfigurationGetEdgeNodeBySerialParamsWithHTTPClient creates a new EdgeNodeConfigurationGetEdgeNodeBySerialParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeNodeConfigurationGetEdgeNodeBySerialParamsWithHTTPClient(client *http.Client) *EdgeNodeConfigurationGetEdgeNodeBySerialParams {
	return &EdgeNodeConfigurationGetEdgeNodeBySerialParams{
		HTTPClient: client,
	}
}

/*
EdgeNodeConfigurationGetEdgeNodeBySerialParams contains all the parameters to send to the API endpoint

	for the edge node configuration get edge node by serial operation.

	Typically these are written to a http.Request.
*/
type EdgeNodeConfigurationGetEdgeNodeBySerialParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* Serialno.

	   unique identity number of the device
	*/
	Serialno string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edge node configuration get edge node by serial params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialParams) WithDefaults() *EdgeNodeConfigurationGetEdgeNodeBySerialParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge node configuration get edge node by serial params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge node configuration get edge node by serial params
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialParams) WithTimeout(timeout time.Duration) *EdgeNodeConfigurationGetEdgeNodeBySerialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge node configuration get edge node by serial params
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge node configuration get edge node by serial params
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialParams) WithContext(ctx context.Context) *EdgeNodeConfigurationGetEdgeNodeBySerialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge node configuration get edge node by serial params
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge node configuration get edge node by serial params
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialParams) WithHTTPClient(client *http.Client) *EdgeNodeConfigurationGetEdgeNodeBySerialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge node configuration get edge node by serial params
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the edge node configuration get edge node by serial params
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialParams) WithXRequestID(xRequestID *string) *EdgeNodeConfigurationGetEdgeNodeBySerialParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the edge node configuration get edge node by serial params
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithSerialno adds the serialno to the edge node configuration get edge node by serial params
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialParams) WithSerialno(serialno string) *EdgeNodeConfigurationGetEdgeNodeBySerialParams {
	o.SetSerialno(serialno)
	return o
}

// SetSerialno adds the serialno to the edge node configuration get edge node by serial params
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialParams) SetSerialno(serialno string) {
	o.Serialno = serialno
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param serialno
	if err := r.SetPathParam("serialno", o.Serialno); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

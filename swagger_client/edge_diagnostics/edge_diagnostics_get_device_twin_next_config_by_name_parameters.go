// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package edge_diagnostics

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

// NewEdgeDiagnosticsGetDeviceTwinNextConfigByNameParams creates a new EdgeDiagnosticsGetDeviceTwinNextConfigByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEdgeDiagnosticsGetDeviceTwinNextConfigByNameParams() *EdgeDiagnosticsGetDeviceTwinNextConfigByNameParams {
	return &EdgeDiagnosticsGetDeviceTwinNextConfigByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeDiagnosticsGetDeviceTwinNextConfigByNameParamsWithTimeout creates a new EdgeDiagnosticsGetDeviceTwinNextConfigByNameParams object
// with the ability to set a timeout on a request.
func NewEdgeDiagnosticsGetDeviceTwinNextConfigByNameParamsWithTimeout(timeout time.Duration) *EdgeDiagnosticsGetDeviceTwinNextConfigByNameParams {
	return &EdgeDiagnosticsGetDeviceTwinNextConfigByNameParams{
		timeout: timeout,
	}
}

// NewEdgeDiagnosticsGetDeviceTwinNextConfigByNameParamsWithContext creates a new EdgeDiagnosticsGetDeviceTwinNextConfigByNameParams object
// with the ability to set a context for a request.
func NewEdgeDiagnosticsGetDeviceTwinNextConfigByNameParamsWithContext(ctx context.Context) *EdgeDiagnosticsGetDeviceTwinNextConfigByNameParams {
	return &EdgeDiagnosticsGetDeviceTwinNextConfigByNameParams{
		Context: ctx,
	}
}

// NewEdgeDiagnosticsGetDeviceTwinNextConfigByNameParamsWithHTTPClient creates a new EdgeDiagnosticsGetDeviceTwinNextConfigByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeDiagnosticsGetDeviceTwinNextConfigByNameParamsWithHTTPClient(client *http.Client) *EdgeDiagnosticsGetDeviceTwinNextConfigByNameParams {
	return &EdgeDiagnosticsGetDeviceTwinNextConfigByNameParams{
		HTTPClient: client,
	}
}

/*
EdgeDiagnosticsGetDeviceTwinNextConfigByNameParams contains all the parameters to send to the API endpoint

	for the edge diagnostics get device twin next config by name operation.

	Typically these are written to a http.Request.
*/
type EdgeDiagnosticsGetDeviceTwinNextConfigByNameParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* Name.

	   user defined device name for a device
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edge diagnostics get device twin next config by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameParams) WithDefaults() *EdgeDiagnosticsGetDeviceTwinNextConfigByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge diagnostics get device twin next config by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge diagnostics get device twin next config by name params
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameParams) WithTimeout(timeout time.Duration) *EdgeDiagnosticsGetDeviceTwinNextConfigByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge diagnostics get device twin next config by name params
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge diagnostics get device twin next config by name params
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameParams) WithContext(ctx context.Context) *EdgeDiagnosticsGetDeviceTwinNextConfigByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge diagnostics get device twin next config by name params
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge diagnostics get device twin next config by name params
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameParams) WithHTTPClient(client *http.Client) *EdgeDiagnosticsGetDeviceTwinNextConfigByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge diagnostics get device twin next config by name params
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the edge diagnostics get device twin next config by name params
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameParams) WithXRequestID(xRequestID *string) *EdgeDiagnosticsGetDeviceTwinNextConfigByNameParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the edge diagnostics get device twin next config by name params
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithName adds the name to the edge diagnostics get device twin next config by name params
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameParams) WithName(name string) *EdgeDiagnosticsGetDeviceTwinNextConfigByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the edge diagnostics get device twin next config by name params
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

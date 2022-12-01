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

// NewEdgeDiagnosticsGetTopUsers2Params creates a new EdgeDiagnosticsGetTopUsers2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEdgeDiagnosticsGetTopUsers2Params() *EdgeDiagnosticsGetTopUsers2Params {
	return &EdgeDiagnosticsGetTopUsers2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeDiagnosticsGetTopUsers2ParamsWithTimeout creates a new EdgeDiagnosticsGetTopUsers2Params object
// with the ability to set a timeout on a request.
func NewEdgeDiagnosticsGetTopUsers2ParamsWithTimeout(timeout time.Duration) *EdgeDiagnosticsGetTopUsers2Params {
	return &EdgeDiagnosticsGetTopUsers2Params{
		timeout: timeout,
	}
}

// NewEdgeDiagnosticsGetTopUsers2ParamsWithContext creates a new EdgeDiagnosticsGetTopUsers2Params object
// with the ability to set a context for a request.
func NewEdgeDiagnosticsGetTopUsers2ParamsWithContext(ctx context.Context) *EdgeDiagnosticsGetTopUsers2Params {
	return &EdgeDiagnosticsGetTopUsers2Params{
		Context: ctx,
	}
}

// NewEdgeDiagnosticsGetTopUsers2ParamsWithHTTPClient creates a new EdgeDiagnosticsGetTopUsers2Params object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeDiagnosticsGetTopUsers2ParamsWithHTTPClient(client *http.Client) *EdgeDiagnosticsGetTopUsers2Params {
	return &EdgeDiagnosticsGetTopUsers2Params{
		HTTPClient: client,
	}
}

/*
EdgeDiagnosticsGetTopUsers2Params contains all the parameters to send to the API endpoint

	for the edge diagnostics get top users2 operation.

	Typically these are written to a http.Request.
*/
type EdgeDiagnosticsGetTopUsers2Params struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edge diagnostics get top users2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeDiagnosticsGetTopUsers2Params) WithDefaults() *EdgeDiagnosticsGetTopUsers2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge diagnostics get top users2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeDiagnosticsGetTopUsers2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge diagnostics get top users2 params
func (o *EdgeDiagnosticsGetTopUsers2Params) WithTimeout(timeout time.Duration) *EdgeDiagnosticsGetTopUsers2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge diagnostics get top users2 params
func (o *EdgeDiagnosticsGetTopUsers2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge diagnostics get top users2 params
func (o *EdgeDiagnosticsGetTopUsers2Params) WithContext(ctx context.Context) *EdgeDiagnosticsGetTopUsers2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge diagnostics get top users2 params
func (o *EdgeDiagnosticsGetTopUsers2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge diagnostics get top users2 params
func (o *EdgeDiagnosticsGetTopUsers2Params) WithHTTPClient(client *http.Client) *EdgeDiagnosticsGetTopUsers2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge diagnostics get top users2 params
func (o *EdgeDiagnosticsGetTopUsers2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the edge diagnostics get top users2 params
func (o *EdgeDiagnosticsGetTopUsers2Params) WithXRequestID(xRequestID *string) *EdgeDiagnosticsGetTopUsers2Params {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the edge diagnostics get top users2 params
func (o *EdgeDiagnosticsGetTopUsers2Params) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeDiagnosticsGetTopUsers2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
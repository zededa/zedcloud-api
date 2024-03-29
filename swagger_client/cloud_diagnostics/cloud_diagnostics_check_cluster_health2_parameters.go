// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package cloud_diagnostics

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

// NewCloudDiagnosticsCheckClusterHealth2Params creates a new CloudDiagnosticsCheckClusterHealth2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCloudDiagnosticsCheckClusterHealth2Params() *CloudDiagnosticsCheckClusterHealth2Params {
	return &CloudDiagnosticsCheckClusterHealth2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCloudDiagnosticsCheckClusterHealth2ParamsWithTimeout creates a new CloudDiagnosticsCheckClusterHealth2Params object
// with the ability to set a timeout on a request.
func NewCloudDiagnosticsCheckClusterHealth2ParamsWithTimeout(timeout time.Duration) *CloudDiagnosticsCheckClusterHealth2Params {
	return &CloudDiagnosticsCheckClusterHealth2Params{
		timeout: timeout,
	}
}

// NewCloudDiagnosticsCheckClusterHealth2ParamsWithContext creates a new CloudDiagnosticsCheckClusterHealth2Params object
// with the ability to set a context for a request.
func NewCloudDiagnosticsCheckClusterHealth2ParamsWithContext(ctx context.Context) *CloudDiagnosticsCheckClusterHealth2Params {
	return &CloudDiagnosticsCheckClusterHealth2Params{
		Context: ctx,
	}
}

// NewCloudDiagnosticsCheckClusterHealth2ParamsWithHTTPClient creates a new CloudDiagnosticsCheckClusterHealth2Params object
// with the ability to set a custom HTTPClient for a request.
func NewCloudDiagnosticsCheckClusterHealth2ParamsWithHTTPClient(client *http.Client) *CloudDiagnosticsCheckClusterHealth2Params {
	return &CloudDiagnosticsCheckClusterHealth2Params{
		HTTPClient: client,
	}
}

/*
CloudDiagnosticsCheckClusterHealth2Params contains all the parameters to send to the API endpoint

	for the cloud diagnostics check cluster health2 operation.

	Typically these are written to a http.Request.
*/
type CloudDiagnosticsCheckClusterHealth2Params struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// PingID.
	//
	// Format: int64
	PingID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cloud diagnostics check cluster health2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CloudDiagnosticsCheckClusterHealth2Params) WithDefaults() *CloudDiagnosticsCheckClusterHealth2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cloud diagnostics check cluster health2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CloudDiagnosticsCheckClusterHealth2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cloud diagnostics check cluster health2 params
func (o *CloudDiagnosticsCheckClusterHealth2Params) WithTimeout(timeout time.Duration) *CloudDiagnosticsCheckClusterHealth2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cloud diagnostics check cluster health2 params
func (o *CloudDiagnosticsCheckClusterHealth2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cloud diagnostics check cluster health2 params
func (o *CloudDiagnosticsCheckClusterHealth2Params) WithContext(ctx context.Context) *CloudDiagnosticsCheckClusterHealth2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cloud diagnostics check cluster health2 params
func (o *CloudDiagnosticsCheckClusterHealth2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cloud diagnostics check cluster health2 params
func (o *CloudDiagnosticsCheckClusterHealth2Params) WithHTTPClient(client *http.Client) *CloudDiagnosticsCheckClusterHealth2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cloud diagnostics check cluster health2 params
func (o *CloudDiagnosticsCheckClusterHealth2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the cloud diagnostics check cluster health2 params
func (o *CloudDiagnosticsCheckClusterHealth2Params) WithXRequestID(xRequestID *string) *CloudDiagnosticsCheckClusterHealth2Params {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the cloud diagnostics check cluster health2 params
func (o *CloudDiagnosticsCheckClusterHealth2Params) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithPingID adds the pingID to the cloud diagnostics check cluster health2 params
func (o *CloudDiagnosticsCheckClusterHealth2Params) WithPingID(pingID string) *CloudDiagnosticsCheckClusterHealth2Params {
	o.SetPingID(pingID)
	return o
}

// SetPingID adds the pingId to the cloud diagnostics check cluster health2 params
func (o *CloudDiagnosticsCheckClusterHealth2Params) SetPingID(pingID string) {
	o.PingID = pingID
}

// WriteToRequest writes these params to a swagger request
func (o *CloudDiagnosticsCheckClusterHealth2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param pingId
	if err := r.SetPathParam("pingId", o.PingID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

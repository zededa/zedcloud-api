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

// NewCloudDiagnosticsMarkCloudPolicyLatestParams creates a new CloudDiagnosticsMarkCloudPolicyLatestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCloudDiagnosticsMarkCloudPolicyLatestParams() *CloudDiagnosticsMarkCloudPolicyLatestParams {
	return &CloudDiagnosticsMarkCloudPolicyLatestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCloudDiagnosticsMarkCloudPolicyLatestParamsWithTimeout creates a new CloudDiagnosticsMarkCloudPolicyLatestParams object
// with the ability to set a timeout on a request.
func NewCloudDiagnosticsMarkCloudPolicyLatestParamsWithTimeout(timeout time.Duration) *CloudDiagnosticsMarkCloudPolicyLatestParams {
	return &CloudDiagnosticsMarkCloudPolicyLatestParams{
		timeout: timeout,
	}
}

// NewCloudDiagnosticsMarkCloudPolicyLatestParamsWithContext creates a new CloudDiagnosticsMarkCloudPolicyLatestParams object
// with the ability to set a context for a request.
func NewCloudDiagnosticsMarkCloudPolicyLatestParamsWithContext(ctx context.Context) *CloudDiagnosticsMarkCloudPolicyLatestParams {
	return &CloudDiagnosticsMarkCloudPolicyLatestParams{
		Context: ctx,
	}
}

// NewCloudDiagnosticsMarkCloudPolicyLatestParamsWithHTTPClient creates a new CloudDiagnosticsMarkCloudPolicyLatestParams object
// with the ability to set a custom HTTPClient for a request.
func NewCloudDiagnosticsMarkCloudPolicyLatestParamsWithHTTPClient(client *http.Client) *CloudDiagnosticsMarkCloudPolicyLatestParams {
	return &CloudDiagnosticsMarkCloudPolicyLatestParams{
		HTTPClient: client,
	}
}

/*
CloudDiagnosticsMarkCloudPolicyLatestParams contains all the parameters to send to the API endpoint

	for the cloud diagnostics mark cloud policy latest operation.

	Typically these are written to a http.Request.
*/
type CloudDiagnosticsMarkCloudPolicyLatestParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cloud diagnostics mark cloud policy latest params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CloudDiagnosticsMarkCloudPolicyLatestParams) WithDefaults() *CloudDiagnosticsMarkCloudPolicyLatestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cloud diagnostics mark cloud policy latest params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CloudDiagnosticsMarkCloudPolicyLatestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cloud diagnostics mark cloud policy latest params
func (o *CloudDiagnosticsMarkCloudPolicyLatestParams) WithTimeout(timeout time.Duration) *CloudDiagnosticsMarkCloudPolicyLatestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cloud diagnostics mark cloud policy latest params
func (o *CloudDiagnosticsMarkCloudPolicyLatestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cloud diagnostics mark cloud policy latest params
func (o *CloudDiagnosticsMarkCloudPolicyLatestParams) WithContext(ctx context.Context) *CloudDiagnosticsMarkCloudPolicyLatestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cloud diagnostics mark cloud policy latest params
func (o *CloudDiagnosticsMarkCloudPolicyLatestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cloud diagnostics mark cloud policy latest params
func (o *CloudDiagnosticsMarkCloudPolicyLatestParams) WithHTTPClient(client *http.Client) *CloudDiagnosticsMarkCloudPolicyLatestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cloud diagnostics mark cloud policy latest params
func (o *CloudDiagnosticsMarkCloudPolicyLatestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the cloud diagnostics mark cloud policy latest params
func (o *CloudDiagnosticsMarkCloudPolicyLatestParams) WithXRequestID(xRequestID *string) *CloudDiagnosticsMarkCloudPolicyLatestParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the cloud diagnostics mark cloud policy latest params
func (o *CloudDiagnosticsMarkCloudPolicyLatestParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WriteToRequest writes these params to a swagger request
func (o *CloudDiagnosticsMarkCloudPolicyLatestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package enterprise_entitlements_report

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

// NewEnterpriseEntitlementsReportGetUserReportParams creates a new EnterpriseEntitlementsReportGetUserReportParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEnterpriseEntitlementsReportGetUserReportParams() *EnterpriseEntitlementsReportGetUserReportParams {
	return &EnterpriseEntitlementsReportGetUserReportParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEnterpriseEntitlementsReportGetUserReportParamsWithTimeout creates a new EnterpriseEntitlementsReportGetUserReportParams object
// with the ability to set a timeout on a request.
func NewEnterpriseEntitlementsReportGetUserReportParamsWithTimeout(timeout time.Duration) *EnterpriseEntitlementsReportGetUserReportParams {
	return &EnterpriseEntitlementsReportGetUserReportParams{
		timeout: timeout,
	}
}

// NewEnterpriseEntitlementsReportGetUserReportParamsWithContext creates a new EnterpriseEntitlementsReportGetUserReportParams object
// with the ability to set a context for a request.
func NewEnterpriseEntitlementsReportGetUserReportParamsWithContext(ctx context.Context) *EnterpriseEntitlementsReportGetUserReportParams {
	return &EnterpriseEntitlementsReportGetUserReportParams{
		Context: ctx,
	}
}

// NewEnterpriseEntitlementsReportGetUserReportParamsWithHTTPClient creates a new EnterpriseEntitlementsReportGetUserReportParams object
// with the ability to set a custom HTTPClient for a request.
func NewEnterpriseEntitlementsReportGetUserReportParamsWithHTTPClient(client *http.Client) *EnterpriseEntitlementsReportGetUserReportParams {
	return &EnterpriseEntitlementsReportGetUserReportParams{
		HTTPClient: client,
	}
}

/*
EnterpriseEntitlementsReportGetUserReportParams contains all the parameters to send to the API endpoint

	for the enterprise entitlements report get user report operation.

	Typically these are written to a http.Request.
*/
type EnterpriseEntitlementsReportGetUserReportParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* TenantID.

	   Enterprise id for which we want to get summary report for all objects
	*/
	TenantID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the enterprise entitlements report get user report params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnterpriseEntitlementsReportGetUserReportParams) WithDefaults() *EnterpriseEntitlementsReportGetUserReportParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the enterprise entitlements report get user report params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnterpriseEntitlementsReportGetUserReportParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the enterprise entitlements report get user report params
func (o *EnterpriseEntitlementsReportGetUserReportParams) WithTimeout(timeout time.Duration) *EnterpriseEntitlementsReportGetUserReportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the enterprise entitlements report get user report params
func (o *EnterpriseEntitlementsReportGetUserReportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the enterprise entitlements report get user report params
func (o *EnterpriseEntitlementsReportGetUserReportParams) WithContext(ctx context.Context) *EnterpriseEntitlementsReportGetUserReportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the enterprise entitlements report get user report params
func (o *EnterpriseEntitlementsReportGetUserReportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the enterprise entitlements report get user report params
func (o *EnterpriseEntitlementsReportGetUserReportParams) WithHTTPClient(client *http.Client) *EnterpriseEntitlementsReportGetUserReportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the enterprise entitlements report get user report params
func (o *EnterpriseEntitlementsReportGetUserReportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the enterprise entitlements report get user report params
func (o *EnterpriseEntitlementsReportGetUserReportParams) WithXRequestID(xRequestID *string) *EnterpriseEntitlementsReportGetUserReportParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the enterprise entitlements report get user report params
func (o *EnterpriseEntitlementsReportGetUserReportParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithTenantID adds the tenantID to the enterprise entitlements report get user report params
func (o *EnterpriseEntitlementsReportGetUserReportParams) WithTenantID(tenantID *string) *EnterpriseEntitlementsReportGetUserReportParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the enterprise entitlements report get user report params
func (o *EnterpriseEntitlementsReportGetUserReportParams) SetTenantID(tenantID *string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *EnterpriseEntitlementsReportGetUserReportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.TenantID != nil {

		// query param tenantId
		var qrTenantID string

		if o.TenantID != nil {
			qrTenantID = *o.TenantID
		}
		qTenantID := qrTenantID
		if qTenantID != "" {

			if err := r.SetQueryParam("tenantId", qTenantID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

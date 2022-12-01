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

// NewEnterpriseEntitlementsReportGetProjectReportParams creates a new EnterpriseEntitlementsReportGetProjectReportParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEnterpriseEntitlementsReportGetProjectReportParams() *EnterpriseEntitlementsReportGetProjectReportParams {
	return &EnterpriseEntitlementsReportGetProjectReportParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEnterpriseEntitlementsReportGetProjectReportParamsWithTimeout creates a new EnterpriseEntitlementsReportGetProjectReportParams object
// with the ability to set a timeout on a request.
func NewEnterpriseEntitlementsReportGetProjectReportParamsWithTimeout(timeout time.Duration) *EnterpriseEntitlementsReportGetProjectReportParams {
	return &EnterpriseEntitlementsReportGetProjectReportParams{
		timeout: timeout,
	}
}

// NewEnterpriseEntitlementsReportGetProjectReportParamsWithContext creates a new EnterpriseEntitlementsReportGetProjectReportParams object
// with the ability to set a context for a request.
func NewEnterpriseEntitlementsReportGetProjectReportParamsWithContext(ctx context.Context) *EnterpriseEntitlementsReportGetProjectReportParams {
	return &EnterpriseEntitlementsReportGetProjectReportParams{
		Context: ctx,
	}
}

// NewEnterpriseEntitlementsReportGetProjectReportParamsWithHTTPClient creates a new EnterpriseEntitlementsReportGetProjectReportParams object
// with the ability to set a custom HTTPClient for a request.
func NewEnterpriseEntitlementsReportGetProjectReportParamsWithHTTPClient(client *http.Client) *EnterpriseEntitlementsReportGetProjectReportParams {
	return &EnterpriseEntitlementsReportGetProjectReportParams{
		HTTPClient: client,
	}
}

/*
EnterpriseEntitlementsReportGetProjectReportParams contains all the parameters to send to the API endpoint

	for the enterprise entitlements report get project report operation.

	Typically these are written to a http.Request.
*/
type EnterpriseEntitlementsReportGetProjectReportParams struct {

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

// WithDefaults hydrates default values in the enterprise entitlements report get project report params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnterpriseEntitlementsReportGetProjectReportParams) WithDefaults() *EnterpriseEntitlementsReportGetProjectReportParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the enterprise entitlements report get project report params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnterpriseEntitlementsReportGetProjectReportParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the enterprise entitlements report get project report params
func (o *EnterpriseEntitlementsReportGetProjectReportParams) WithTimeout(timeout time.Duration) *EnterpriseEntitlementsReportGetProjectReportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the enterprise entitlements report get project report params
func (o *EnterpriseEntitlementsReportGetProjectReportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the enterprise entitlements report get project report params
func (o *EnterpriseEntitlementsReportGetProjectReportParams) WithContext(ctx context.Context) *EnterpriseEntitlementsReportGetProjectReportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the enterprise entitlements report get project report params
func (o *EnterpriseEntitlementsReportGetProjectReportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the enterprise entitlements report get project report params
func (o *EnterpriseEntitlementsReportGetProjectReportParams) WithHTTPClient(client *http.Client) *EnterpriseEntitlementsReportGetProjectReportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the enterprise entitlements report get project report params
func (o *EnterpriseEntitlementsReportGetProjectReportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the enterprise entitlements report get project report params
func (o *EnterpriseEntitlementsReportGetProjectReportParams) WithXRequestID(xRequestID *string) *EnterpriseEntitlementsReportGetProjectReportParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the enterprise entitlements report get project report params
func (o *EnterpriseEntitlementsReportGetProjectReportParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithTenantID adds the tenantID to the enterprise entitlements report get project report params
func (o *EnterpriseEntitlementsReportGetProjectReportParams) WithTenantID(tenantID *string) *EnterpriseEntitlementsReportGetProjectReportParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the enterprise entitlements report get project report params
func (o *EnterpriseEntitlementsReportGetProjectReportParams) SetTenantID(tenantID *string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *EnterpriseEntitlementsReportGetProjectReportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
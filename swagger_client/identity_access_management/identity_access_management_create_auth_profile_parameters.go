// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package identity_access_management

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

	"github.com/zededa/zedcloud-api/swagger_models"
)

// NewIdentityAccessManagementCreateAuthProfileParams creates a new IdentityAccessManagementCreateAuthProfileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIdentityAccessManagementCreateAuthProfileParams() *IdentityAccessManagementCreateAuthProfileParams {
	return &IdentityAccessManagementCreateAuthProfileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIdentityAccessManagementCreateAuthProfileParamsWithTimeout creates a new IdentityAccessManagementCreateAuthProfileParams object
// with the ability to set a timeout on a request.
func NewIdentityAccessManagementCreateAuthProfileParamsWithTimeout(timeout time.Duration) *IdentityAccessManagementCreateAuthProfileParams {
	return &IdentityAccessManagementCreateAuthProfileParams{
		timeout: timeout,
	}
}

// NewIdentityAccessManagementCreateAuthProfileParamsWithContext creates a new IdentityAccessManagementCreateAuthProfileParams object
// with the ability to set a context for a request.
func NewIdentityAccessManagementCreateAuthProfileParamsWithContext(ctx context.Context) *IdentityAccessManagementCreateAuthProfileParams {
	return &IdentityAccessManagementCreateAuthProfileParams{
		Context: ctx,
	}
}

// NewIdentityAccessManagementCreateAuthProfileParamsWithHTTPClient creates a new IdentityAccessManagementCreateAuthProfileParams object
// with the ability to set a custom HTTPClient for a request.
func NewIdentityAccessManagementCreateAuthProfileParamsWithHTTPClient(client *http.Client) *IdentityAccessManagementCreateAuthProfileParams {
	return &IdentityAccessManagementCreateAuthProfileParams{
		HTTPClient: client,
	}
}

/*
IdentityAccessManagementCreateAuthProfileParams contains all the parameters to send to the API endpoint

	for the identity access management create auth profile operation.

	Typically these are written to a http.Request.
*/
type IdentityAccessManagementCreateAuthProfileParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body *swagger_models.AuthorizationProfile

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the identity access management create auth profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IdentityAccessManagementCreateAuthProfileParams) WithDefaults() *IdentityAccessManagementCreateAuthProfileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the identity access management create auth profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IdentityAccessManagementCreateAuthProfileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the identity access management create auth profile params
func (o *IdentityAccessManagementCreateAuthProfileParams) WithTimeout(timeout time.Duration) *IdentityAccessManagementCreateAuthProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the identity access management create auth profile params
func (o *IdentityAccessManagementCreateAuthProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the identity access management create auth profile params
func (o *IdentityAccessManagementCreateAuthProfileParams) WithContext(ctx context.Context) *IdentityAccessManagementCreateAuthProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the identity access management create auth profile params
func (o *IdentityAccessManagementCreateAuthProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the identity access management create auth profile params
func (o *IdentityAccessManagementCreateAuthProfileParams) WithHTTPClient(client *http.Client) *IdentityAccessManagementCreateAuthProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the identity access management create auth profile params
func (o *IdentityAccessManagementCreateAuthProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the identity access management create auth profile params
func (o *IdentityAccessManagementCreateAuthProfileParams) WithXRequestID(xRequestID *string) *IdentityAccessManagementCreateAuthProfileParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the identity access management create auth profile params
func (o *IdentityAccessManagementCreateAuthProfileParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the identity access management create auth profile params
func (o *IdentityAccessManagementCreateAuthProfileParams) WithBody(body *swagger_models.AuthorizationProfile) *IdentityAccessManagementCreateAuthProfileParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the identity access management create auth profile params
func (o *IdentityAccessManagementCreateAuthProfileParams) SetBody(body *swagger_models.AuthorizationProfile) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *IdentityAccessManagementCreateAuthProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

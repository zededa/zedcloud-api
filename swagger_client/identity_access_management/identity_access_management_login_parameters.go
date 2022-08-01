// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

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

// NewIdentityAccessManagementLoginParams creates a new IdentityAccessManagementLoginParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIdentityAccessManagementLoginParams() *IdentityAccessManagementLoginParams {
	return &IdentityAccessManagementLoginParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIdentityAccessManagementLoginParamsWithTimeout creates a new IdentityAccessManagementLoginParams object
// with the ability to set a timeout on a request.
func NewIdentityAccessManagementLoginParamsWithTimeout(timeout time.Duration) *IdentityAccessManagementLoginParams {
	return &IdentityAccessManagementLoginParams{
		timeout: timeout,
	}
}

// NewIdentityAccessManagementLoginParamsWithContext creates a new IdentityAccessManagementLoginParams object
// with the ability to set a context for a request.
func NewIdentityAccessManagementLoginParamsWithContext(ctx context.Context) *IdentityAccessManagementLoginParams {
	return &IdentityAccessManagementLoginParams{
		Context: ctx,
	}
}

// NewIdentityAccessManagementLoginParamsWithHTTPClient creates a new IdentityAccessManagementLoginParams object
// with the ability to set a custom HTTPClient for a request.
func NewIdentityAccessManagementLoginParamsWithHTTPClient(client *http.Client) *IdentityAccessManagementLoginParams {
	return &IdentityAccessManagementLoginParams{
		HTTPClient: client,
	}
}

/* IdentityAccessManagementLoginParams contains all the parameters to send to the API endpoint
   for the identity access management login operation.

   Typically these are written to a http.Request.
*/
type IdentityAccessManagementLoginParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body *swagger_models.AAAFrontendLoginWithPasswordRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the identity access management login params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IdentityAccessManagementLoginParams) WithDefaults() *IdentityAccessManagementLoginParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the identity access management login params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IdentityAccessManagementLoginParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the identity access management login params
func (o *IdentityAccessManagementLoginParams) WithTimeout(timeout time.Duration) *IdentityAccessManagementLoginParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the identity access management login params
func (o *IdentityAccessManagementLoginParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the identity access management login params
func (o *IdentityAccessManagementLoginParams) WithContext(ctx context.Context) *IdentityAccessManagementLoginParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the identity access management login params
func (o *IdentityAccessManagementLoginParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the identity access management login params
func (o *IdentityAccessManagementLoginParams) WithHTTPClient(client *http.Client) *IdentityAccessManagementLoginParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the identity access management login params
func (o *IdentityAccessManagementLoginParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the identity access management login params
func (o *IdentityAccessManagementLoginParams) WithXRequestID(xRequestID *string) *IdentityAccessManagementLoginParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the identity access management login params
func (o *IdentityAccessManagementLoginParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the identity access management login params
func (o *IdentityAccessManagementLoginParams) WithBody(body *swagger_models.AAAFrontendLoginWithPasswordRequest) *IdentityAccessManagementLoginParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the identity access management login params
func (o *IdentityAccessManagementLoginParams) SetBody(body *swagger_models.AAAFrontendLoginWithPasswordRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *IdentityAccessManagementLoginParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

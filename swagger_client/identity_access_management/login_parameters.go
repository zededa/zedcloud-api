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

// NewLoginParams creates a new LoginParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLoginParams() *LoginParams {
	return &LoginParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLoginParamsWithTimeout creates a new LoginParams object
// with the ability to set a timeout on a request.
func NewLoginParamsWithTimeout(timeout time.Duration) *LoginParams {
	return &LoginParams{
		timeout: timeout,
	}
}

// NewLoginParamsWithContext creates a new LoginParams object
// with the ability to set a context for a request.
func NewLoginParamsWithContext(ctx context.Context) *LoginParams {
	return &LoginParams{
		Context: ctx,
	}
}

// NewLoginParamsWithHTTPClient creates a new LoginParams object
// with the ability to set a custom HTTPClient for a request.
func NewLoginParamsWithHTTPClient(client *http.Client) *LoginParams {
	return &LoginParams{
		HTTPClient: client,
	}
}

/* LoginParams contains all the parameters to send to the API endpoint
   for the login operation.

   Typically these are written to a http.Request.
*/
type LoginParams struct {

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

// WithDefaults hydrates default values in the login params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LoginParams) WithDefaults() *LoginParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the login params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LoginParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the login params
func (o *LoginParams) WithTimeout(timeout time.Duration) *LoginParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the login params
func (o *LoginParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the login params
func (o *LoginParams) WithContext(ctx context.Context) *LoginParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the login params
func (o *LoginParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the login params
func (o *LoginParams) WithHTTPClient(client *http.Client) *LoginParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the login params
func (o *LoginParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the login params
func (o *LoginParams) WithXRequestID(xRequestID *string) *LoginParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the login params
func (o *LoginParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the login params
func (o *LoginParams) WithBody(body *swagger_models.AAAFrontendLoginWithPasswordRequest) *LoginParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the login params
func (o *LoginParams) SetBody(body *swagger_models.AAAFrontendLoginWithPasswordRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *LoginParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

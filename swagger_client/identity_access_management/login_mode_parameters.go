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

// NewLoginModeParams creates a new LoginModeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLoginModeParams() *LoginModeParams {
	return &LoginModeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLoginModeParamsWithTimeout creates a new LoginModeParams object
// with the ability to set a timeout on a request.
func NewLoginModeParamsWithTimeout(timeout time.Duration) *LoginModeParams {
	return &LoginModeParams{
		timeout: timeout,
	}
}

// NewLoginModeParamsWithContext creates a new LoginModeParams object
// with the ability to set a context for a request.
func NewLoginModeParamsWithContext(ctx context.Context) *LoginModeParams {
	return &LoginModeParams{
		Context: ctx,
	}
}

// NewLoginModeParamsWithHTTPClient creates a new LoginModeParams object
// with the ability to set a custom HTTPClient for a request.
func NewLoginModeParamsWithHTTPClient(client *http.Client) *LoginModeParams {
	return &LoginModeParams{
		HTTPClient: client,
	}
}

/* LoginModeParams contains all the parameters to send to the API endpoint
   for the login mode operation.

   Typically these are written to a http.Request.
*/
type LoginModeParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body *swagger_models.AAAFrontendLoginModeRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the login mode params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LoginModeParams) WithDefaults() *LoginModeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the login mode params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LoginModeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the login mode params
func (o *LoginModeParams) WithTimeout(timeout time.Duration) *LoginModeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the login mode params
func (o *LoginModeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the login mode params
func (o *LoginModeParams) WithContext(ctx context.Context) *LoginModeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the login mode params
func (o *LoginModeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the login mode params
func (o *LoginModeParams) WithHTTPClient(client *http.Client) *LoginModeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the login mode params
func (o *LoginModeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the login mode params
func (o *LoginModeParams) WithXRequestID(xRequestID *string) *LoginModeParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the login mode params
func (o *LoginModeParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the login mode params
func (o *LoginModeParams) WithBody(body *swagger_models.AAAFrontendLoginModeRequest) *LoginModeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the login mode params
func (o *LoginModeParams) SetBody(body *swagger_models.AAAFrontendLoginModeRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *LoginModeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

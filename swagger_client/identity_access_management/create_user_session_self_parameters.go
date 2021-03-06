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
)

// NewCreateUserSessionSelfParams creates a new CreateUserSessionSelfParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateUserSessionSelfParams() *CreateUserSessionSelfParams {
	return &CreateUserSessionSelfParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateUserSessionSelfParamsWithTimeout creates a new CreateUserSessionSelfParams object
// with the ability to set a timeout on a request.
func NewCreateUserSessionSelfParamsWithTimeout(timeout time.Duration) *CreateUserSessionSelfParams {
	return &CreateUserSessionSelfParams{
		timeout: timeout,
	}
}

// NewCreateUserSessionSelfParamsWithContext creates a new CreateUserSessionSelfParams object
// with the ability to set a context for a request.
func NewCreateUserSessionSelfParamsWithContext(ctx context.Context) *CreateUserSessionSelfParams {
	return &CreateUserSessionSelfParams{
		Context: ctx,
	}
}

// NewCreateUserSessionSelfParamsWithHTTPClient creates a new CreateUserSessionSelfParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateUserSessionSelfParamsWithHTTPClient(client *http.Client) *CreateUserSessionSelfParams {
	return &CreateUserSessionSelfParams{
		HTTPClient: client,
	}
}

/* CreateUserSessionSelfParams contains all the parameters to send to the API endpoint
   for the create user session self operation.

   Typically these are written to a http.Request.
*/
type CreateUserSessionSelfParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create user session self params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateUserSessionSelfParams) WithDefaults() *CreateUserSessionSelfParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create user session self params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateUserSessionSelfParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create user session self params
func (o *CreateUserSessionSelfParams) WithTimeout(timeout time.Duration) *CreateUserSessionSelfParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create user session self params
func (o *CreateUserSessionSelfParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create user session self params
func (o *CreateUserSessionSelfParams) WithContext(ctx context.Context) *CreateUserSessionSelfParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create user session self params
func (o *CreateUserSessionSelfParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create user session self params
func (o *CreateUserSessionSelfParams) WithHTTPClient(client *http.Client) *CreateUserSessionSelfParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create user session self params
func (o *CreateUserSessionSelfParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the create user session self params
func (o *CreateUserSessionSelfParams) WithXRequestID(xRequestID *string) *CreateUserSessionSelfParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the create user session self params
func (o *CreateUserSessionSelfParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateUserSessionSelfParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

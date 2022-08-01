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

// NewIdentityAccessManagementUpdateEnterprise2Params creates a new IdentityAccessManagementUpdateEnterprise2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIdentityAccessManagementUpdateEnterprise2Params() *IdentityAccessManagementUpdateEnterprise2Params {
	return &IdentityAccessManagementUpdateEnterprise2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewIdentityAccessManagementUpdateEnterprise2ParamsWithTimeout creates a new IdentityAccessManagementUpdateEnterprise2Params object
// with the ability to set a timeout on a request.
func NewIdentityAccessManagementUpdateEnterprise2ParamsWithTimeout(timeout time.Duration) *IdentityAccessManagementUpdateEnterprise2Params {
	return &IdentityAccessManagementUpdateEnterprise2Params{
		timeout: timeout,
	}
}

// NewIdentityAccessManagementUpdateEnterprise2ParamsWithContext creates a new IdentityAccessManagementUpdateEnterprise2Params object
// with the ability to set a context for a request.
func NewIdentityAccessManagementUpdateEnterprise2ParamsWithContext(ctx context.Context) *IdentityAccessManagementUpdateEnterprise2Params {
	return &IdentityAccessManagementUpdateEnterprise2Params{
		Context: ctx,
	}
}

// NewIdentityAccessManagementUpdateEnterprise2ParamsWithHTTPClient creates a new IdentityAccessManagementUpdateEnterprise2Params object
// with the ability to set a custom HTTPClient for a request.
func NewIdentityAccessManagementUpdateEnterprise2ParamsWithHTTPClient(client *http.Client) *IdentityAccessManagementUpdateEnterprise2Params {
	return &IdentityAccessManagementUpdateEnterprise2Params{
		HTTPClient: client,
	}
}

/* IdentityAccessManagementUpdateEnterprise2Params contains all the parameters to send to the API endpoint
   for the identity access management update enterprise2 operation.

   Typically these are written to a http.Request.
*/
type IdentityAccessManagementUpdateEnterprise2Params struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body IdentityAccessManagementUpdateEnterprise2Body

	/* ID.

	   Unique system defined enterprise ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the identity access management update enterprise2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IdentityAccessManagementUpdateEnterprise2Params) WithDefaults() *IdentityAccessManagementUpdateEnterprise2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the identity access management update enterprise2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IdentityAccessManagementUpdateEnterprise2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the identity access management update enterprise2 params
func (o *IdentityAccessManagementUpdateEnterprise2Params) WithTimeout(timeout time.Duration) *IdentityAccessManagementUpdateEnterprise2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the identity access management update enterprise2 params
func (o *IdentityAccessManagementUpdateEnterprise2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the identity access management update enterprise2 params
func (o *IdentityAccessManagementUpdateEnterprise2Params) WithContext(ctx context.Context) *IdentityAccessManagementUpdateEnterprise2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the identity access management update enterprise2 params
func (o *IdentityAccessManagementUpdateEnterprise2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the identity access management update enterprise2 params
func (o *IdentityAccessManagementUpdateEnterprise2Params) WithHTTPClient(client *http.Client) *IdentityAccessManagementUpdateEnterprise2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the identity access management update enterprise2 params
func (o *IdentityAccessManagementUpdateEnterprise2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the identity access management update enterprise2 params
func (o *IdentityAccessManagementUpdateEnterprise2Params) WithXRequestID(xRequestID *string) *IdentityAccessManagementUpdateEnterprise2Params {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the identity access management update enterprise2 params
func (o *IdentityAccessManagementUpdateEnterprise2Params) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the identity access management update enterprise2 params
func (o *IdentityAccessManagementUpdateEnterprise2Params) WithBody(body IdentityAccessManagementUpdateEnterprise2Body) *IdentityAccessManagementUpdateEnterprise2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the identity access management update enterprise2 params
func (o *IdentityAccessManagementUpdateEnterprise2Params) SetBody(body IdentityAccessManagementUpdateEnterprise2Body) {
	o.Body = body
}

// WithID adds the id to the identity access management update enterprise2 params
func (o *IdentityAccessManagementUpdateEnterprise2Params) WithID(id string) *IdentityAccessManagementUpdateEnterprise2Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the identity access management update enterprise2 params
func (o *IdentityAccessManagementUpdateEnterprise2Params) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IdentityAccessManagementUpdateEnterprise2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
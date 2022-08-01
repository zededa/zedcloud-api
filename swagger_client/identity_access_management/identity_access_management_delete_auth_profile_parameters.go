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

// NewIdentityAccessManagementDeleteAuthProfileParams creates a new IdentityAccessManagementDeleteAuthProfileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIdentityAccessManagementDeleteAuthProfileParams() *IdentityAccessManagementDeleteAuthProfileParams {
	return &IdentityAccessManagementDeleteAuthProfileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIdentityAccessManagementDeleteAuthProfileParamsWithTimeout creates a new IdentityAccessManagementDeleteAuthProfileParams object
// with the ability to set a timeout on a request.
func NewIdentityAccessManagementDeleteAuthProfileParamsWithTimeout(timeout time.Duration) *IdentityAccessManagementDeleteAuthProfileParams {
	return &IdentityAccessManagementDeleteAuthProfileParams{
		timeout: timeout,
	}
}

// NewIdentityAccessManagementDeleteAuthProfileParamsWithContext creates a new IdentityAccessManagementDeleteAuthProfileParams object
// with the ability to set a context for a request.
func NewIdentityAccessManagementDeleteAuthProfileParamsWithContext(ctx context.Context) *IdentityAccessManagementDeleteAuthProfileParams {
	return &IdentityAccessManagementDeleteAuthProfileParams{
		Context: ctx,
	}
}

// NewIdentityAccessManagementDeleteAuthProfileParamsWithHTTPClient creates a new IdentityAccessManagementDeleteAuthProfileParams object
// with the ability to set a custom HTTPClient for a request.
func NewIdentityAccessManagementDeleteAuthProfileParamsWithHTTPClient(client *http.Client) *IdentityAccessManagementDeleteAuthProfileParams {
	return &IdentityAccessManagementDeleteAuthProfileParams{
		HTTPClient: client,
	}
}

/* IdentityAccessManagementDeleteAuthProfileParams contains all the parameters to send to the API endpoint
   for the identity access management delete auth profile operation.

   Typically these are written to a http.Request.
*/
type IdentityAccessManagementDeleteAuthProfileParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the identity access management delete auth profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IdentityAccessManagementDeleteAuthProfileParams) WithDefaults() *IdentityAccessManagementDeleteAuthProfileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the identity access management delete auth profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IdentityAccessManagementDeleteAuthProfileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the identity access management delete auth profile params
func (o *IdentityAccessManagementDeleteAuthProfileParams) WithTimeout(timeout time.Duration) *IdentityAccessManagementDeleteAuthProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the identity access management delete auth profile params
func (o *IdentityAccessManagementDeleteAuthProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the identity access management delete auth profile params
func (o *IdentityAccessManagementDeleteAuthProfileParams) WithContext(ctx context.Context) *IdentityAccessManagementDeleteAuthProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the identity access management delete auth profile params
func (o *IdentityAccessManagementDeleteAuthProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the identity access management delete auth profile params
func (o *IdentityAccessManagementDeleteAuthProfileParams) WithHTTPClient(client *http.Client) *IdentityAccessManagementDeleteAuthProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the identity access management delete auth profile params
func (o *IdentityAccessManagementDeleteAuthProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the identity access management delete auth profile params
func (o *IdentityAccessManagementDeleteAuthProfileParams) WithXRequestID(xRequestID *string) *IdentityAccessManagementDeleteAuthProfileParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the identity access management delete auth profile params
func (o *IdentityAccessManagementDeleteAuthProfileParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the identity access management delete auth profile params
func (o *IdentityAccessManagementDeleteAuthProfileParams) WithID(id string) *IdentityAccessManagementDeleteAuthProfileParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the identity access management delete auth profile params
func (o *IdentityAccessManagementDeleteAuthProfileParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IdentityAccessManagementDeleteAuthProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

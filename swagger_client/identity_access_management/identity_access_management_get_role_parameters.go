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

// NewIdentityAccessManagementGetRoleParams creates a new IdentityAccessManagementGetRoleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIdentityAccessManagementGetRoleParams() *IdentityAccessManagementGetRoleParams {
	return &IdentityAccessManagementGetRoleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIdentityAccessManagementGetRoleParamsWithTimeout creates a new IdentityAccessManagementGetRoleParams object
// with the ability to set a timeout on a request.
func NewIdentityAccessManagementGetRoleParamsWithTimeout(timeout time.Duration) *IdentityAccessManagementGetRoleParams {
	return &IdentityAccessManagementGetRoleParams{
		timeout: timeout,
	}
}

// NewIdentityAccessManagementGetRoleParamsWithContext creates a new IdentityAccessManagementGetRoleParams object
// with the ability to set a context for a request.
func NewIdentityAccessManagementGetRoleParamsWithContext(ctx context.Context) *IdentityAccessManagementGetRoleParams {
	return &IdentityAccessManagementGetRoleParams{
		Context: ctx,
	}
}

// NewIdentityAccessManagementGetRoleParamsWithHTTPClient creates a new IdentityAccessManagementGetRoleParams object
// with the ability to set a custom HTTPClient for a request.
func NewIdentityAccessManagementGetRoleParamsWithHTTPClient(client *http.Client) *IdentityAccessManagementGetRoleParams {
	return &IdentityAccessManagementGetRoleParams{
		HTTPClient: client,
	}
}

/* IdentityAccessManagementGetRoleParams contains all the parameters to send to the API endpoint
   for the identity access management get role operation.

   Typically these are written to a http.Request.
*/
type IdentityAccessManagementGetRoleParams struct {

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

// WithDefaults hydrates default values in the identity access management get role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IdentityAccessManagementGetRoleParams) WithDefaults() *IdentityAccessManagementGetRoleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the identity access management get role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IdentityAccessManagementGetRoleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the identity access management get role params
func (o *IdentityAccessManagementGetRoleParams) WithTimeout(timeout time.Duration) *IdentityAccessManagementGetRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the identity access management get role params
func (o *IdentityAccessManagementGetRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the identity access management get role params
func (o *IdentityAccessManagementGetRoleParams) WithContext(ctx context.Context) *IdentityAccessManagementGetRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the identity access management get role params
func (o *IdentityAccessManagementGetRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the identity access management get role params
func (o *IdentityAccessManagementGetRoleParams) WithHTTPClient(client *http.Client) *IdentityAccessManagementGetRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the identity access management get role params
func (o *IdentityAccessManagementGetRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the identity access management get role params
func (o *IdentityAccessManagementGetRoleParams) WithXRequestID(xRequestID *string) *IdentityAccessManagementGetRoleParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the identity access management get role params
func (o *IdentityAccessManagementGetRoleParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the identity access management get role params
func (o *IdentityAccessManagementGetRoleParams) WithID(id string) *IdentityAccessManagementGetRoleParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the identity access management get role params
func (o *IdentityAccessManagementGetRoleParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IdentityAccessManagementGetRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

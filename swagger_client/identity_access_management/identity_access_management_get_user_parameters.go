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
)

// NewIdentityAccessManagementGetUserParams creates a new IdentityAccessManagementGetUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIdentityAccessManagementGetUserParams() *IdentityAccessManagementGetUserParams {
	return &IdentityAccessManagementGetUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIdentityAccessManagementGetUserParamsWithTimeout creates a new IdentityAccessManagementGetUserParams object
// with the ability to set a timeout on a request.
func NewIdentityAccessManagementGetUserParamsWithTimeout(timeout time.Duration) *IdentityAccessManagementGetUserParams {
	return &IdentityAccessManagementGetUserParams{
		timeout: timeout,
	}
}

// NewIdentityAccessManagementGetUserParamsWithContext creates a new IdentityAccessManagementGetUserParams object
// with the ability to set a context for a request.
func NewIdentityAccessManagementGetUserParamsWithContext(ctx context.Context) *IdentityAccessManagementGetUserParams {
	return &IdentityAccessManagementGetUserParams{
		Context: ctx,
	}
}

// NewIdentityAccessManagementGetUserParamsWithHTTPClient creates a new IdentityAccessManagementGetUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewIdentityAccessManagementGetUserParamsWithHTTPClient(client *http.Client) *IdentityAccessManagementGetUserParams {
	return &IdentityAccessManagementGetUserParams{
		HTTPClient: client,
	}
}

/*
IdentityAccessManagementGetUserParams contains all the parameters to send to the API endpoint

	for the identity access management get user operation.

	Typically these are written to a http.Request.
*/
type IdentityAccessManagementGetUserParams struct {

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

// WithDefaults hydrates default values in the identity access management get user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IdentityAccessManagementGetUserParams) WithDefaults() *IdentityAccessManagementGetUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the identity access management get user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IdentityAccessManagementGetUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the identity access management get user params
func (o *IdentityAccessManagementGetUserParams) WithTimeout(timeout time.Duration) *IdentityAccessManagementGetUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the identity access management get user params
func (o *IdentityAccessManagementGetUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the identity access management get user params
func (o *IdentityAccessManagementGetUserParams) WithContext(ctx context.Context) *IdentityAccessManagementGetUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the identity access management get user params
func (o *IdentityAccessManagementGetUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the identity access management get user params
func (o *IdentityAccessManagementGetUserParams) WithHTTPClient(client *http.Client) *IdentityAccessManagementGetUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the identity access management get user params
func (o *IdentityAccessManagementGetUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the identity access management get user params
func (o *IdentityAccessManagementGetUserParams) WithXRequestID(xRequestID *string) *IdentityAccessManagementGetUserParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the identity access management get user params
func (o *IdentityAccessManagementGetUserParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the identity access management get user params
func (o *IdentityAccessManagementGetUserParams) WithID(id string) *IdentityAccessManagementGetUserParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the identity access management get user params
func (o *IdentityAccessManagementGetUserParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IdentityAccessManagementGetUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

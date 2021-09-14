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

// NewGetRoleByNameParams creates a new GetRoleByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRoleByNameParams() *GetRoleByNameParams {
	return &GetRoleByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRoleByNameParamsWithTimeout creates a new GetRoleByNameParams object
// with the ability to set a timeout on a request.
func NewGetRoleByNameParamsWithTimeout(timeout time.Duration) *GetRoleByNameParams {
	return &GetRoleByNameParams{
		timeout: timeout,
	}
}

// NewGetRoleByNameParamsWithContext creates a new GetRoleByNameParams object
// with the ability to set a context for a request.
func NewGetRoleByNameParamsWithContext(ctx context.Context) *GetRoleByNameParams {
	return &GetRoleByNameParams{
		Context: ctx,
	}
}

// NewGetRoleByNameParamsWithHTTPClient creates a new GetRoleByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRoleByNameParamsWithHTTPClient(client *http.Client) *GetRoleByNameParams {
	return &GetRoleByNameParams{
		HTTPClient: client,
	}
}

/* GetRoleByNameParams contains all the parameters to send to the API endpoint
   for the get role by name operation.

   Typically these are written to a http.Request.
*/
type GetRoleByNameParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Name.
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get role by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRoleByNameParams) WithDefaults() *GetRoleByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get role by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRoleByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get role by name params
func (o *GetRoleByNameParams) WithTimeout(timeout time.Duration) *GetRoleByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get role by name params
func (o *GetRoleByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get role by name params
func (o *GetRoleByNameParams) WithContext(ctx context.Context) *GetRoleByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get role by name params
func (o *GetRoleByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get role by name params
func (o *GetRoleByNameParams) WithHTTPClient(client *http.Client) *GetRoleByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get role by name params
func (o *GetRoleByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get role by name params
func (o *GetRoleByNameParams) WithXRequestID(xRequestID *string) *GetRoleByNameParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get role by name params
func (o *GetRoleByNameParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithName adds the name to the get role by name params
func (o *GetRoleByNameParams) WithName(name string) *GetRoleByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get role by name params
func (o *GetRoleByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetRoleByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

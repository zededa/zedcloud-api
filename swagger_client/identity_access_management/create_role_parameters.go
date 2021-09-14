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

// NewCreateRoleParams creates a new CreateRoleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateRoleParams() *CreateRoleParams {
	return &CreateRoleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRoleParamsWithTimeout creates a new CreateRoleParams object
// with the ability to set a timeout on a request.
func NewCreateRoleParamsWithTimeout(timeout time.Duration) *CreateRoleParams {
	return &CreateRoleParams{
		timeout: timeout,
	}
}

// NewCreateRoleParamsWithContext creates a new CreateRoleParams object
// with the ability to set a context for a request.
func NewCreateRoleParamsWithContext(ctx context.Context) *CreateRoleParams {
	return &CreateRoleParams{
		Context: ctx,
	}
}

// NewCreateRoleParamsWithHTTPClient creates a new CreateRoleParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateRoleParamsWithHTTPClient(client *http.Client) *CreateRoleParams {
	return &CreateRoleParams{
		HTTPClient: client,
	}
}

/* CreateRoleParams contains all the parameters to send to the API endpoint
   for the create role operation.

   Typically these are written to a http.Request.
*/
type CreateRoleParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body *swagger_models.Role

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRoleParams) WithDefaults() *CreateRoleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRoleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create role params
func (o *CreateRoleParams) WithTimeout(timeout time.Duration) *CreateRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create role params
func (o *CreateRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create role params
func (o *CreateRoleParams) WithContext(ctx context.Context) *CreateRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create role params
func (o *CreateRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create role params
func (o *CreateRoleParams) WithHTTPClient(client *http.Client) *CreateRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create role params
func (o *CreateRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the create role params
func (o *CreateRoleParams) WithXRequestID(xRequestID *string) *CreateRoleParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the create role params
func (o *CreateRoleParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the create role params
func (o *CreateRoleParams) WithBody(body *swagger_models.Role) *CreateRoleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create role params
func (o *CreateRoleParams) SetBody(body *swagger_models.Role) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

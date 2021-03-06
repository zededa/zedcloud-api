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

// NewUpdateAuthProfileParams creates a new UpdateAuthProfileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateAuthProfileParams() *UpdateAuthProfileParams {
	return &UpdateAuthProfileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAuthProfileParamsWithTimeout creates a new UpdateAuthProfileParams object
// with the ability to set a timeout on a request.
func NewUpdateAuthProfileParamsWithTimeout(timeout time.Duration) *UpdateAuthProfileParams {
	return &UpdateAuthProfileParams{
		timeout: timeout,
	}
}

// NewUpdateAuthProfileParamsWithContext creates a new UpdateAuthProfileParams object
// with the ability to set a context for a request.
func NewUpdateAuthProfileParamsWithContext(ctx context.Context) *UpdateAuthProfileParams {
	return &UpdateAuthProfileParams{
		Context: ctx,
	}
}

// NewUpdateAuthProfileParamsWithHTTPClient creates a new UpdateAuthProfileParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateAuthProfileParamsWithHTTPClient(client *http.Client) *UpdateAuthProfileParams {
	return &UpdateAuthProfileParams{
		HTTPClient: client,
	}
}

/* UpdateAuthProfileParams contains all the parameters to send to the API endpoint
   for the update auth profile operation.

   Typically these are written to a http.Request.
*/
type UpdateAuthProfileParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body *swagger_models.AuthorizationProfile

	/* ID.

	   Unique system defined profile ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update auth profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAuthProfileParams) WithDefaults() *UpdateAuthProfileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update auth profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAuthProfileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update auth profile params
func (o *UpdateAuthProfileParams) WithTimeout(timeout time.Duration) *UpdateAuthProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update auth profile params
func (o *UpdateAuthProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update auth profile params
func (o *UpdateAuthProfileParams) WithContext(ctx context.Context) *UpdateAuthProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update auth profile params
func (o *UpdateAuthProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update auth profile params
func (o *UpdateAuthProfileParams) WithHTTPClient(client *http.Client) *UpdateAuthProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update auth profile params
func (o *UpdateAuthProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the update auth profile params
func (o *UpdateAuthProfileParams) WithXRequestID(xRequestID *string) *UpdateAuthProfileParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the update auth profile params
func (o *UpdateAuthProfileParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the update auth profile params
func (o *UpdateAuthProfileParams) WithBody(body *swagger_models.AuthorizationProfile) *UpdateAuthProfileParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update auth profile params
func (o *UpdateAuthProfileParams) SetBody(body *swagger_models.AuthorizationProfile) {
	o.Body = body
}

// WithID adds the id to the update auth profile params
func (o *UpdateAuthProfileParams) WithID(id string) *UpdateAuthProfileParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update auth profile params
func (o *UpdateAuthProfileParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAuthProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

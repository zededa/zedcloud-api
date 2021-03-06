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

// NewUpdateEnterprise2Params creates a new UpdateEnterprise2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateEnterprise2Params() *UpdateEnterprise2Params {
	return &UpdateEnterprise2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateEnterprise2ParamsWithTimeout creates a new UpdateEnterprise2Params object
// with the ability to set a timeout on a request.
func NewUpdateEnterprise2ParamsWithTimeout(timeout time.Duration) *UpdateEnterprise2Params {
	return &UpdateEnterprise2Params{
		timeout: timeout,
	}
}

// NewUpdateEnterprise2ParamsWithContext creates a new UpdateEnterprise2Params object
// with the ability to set a context for a request.
func NewUpdateEnterprise2ParamsWithContext(ctx context.Context) *UpdateEnterprise2Params {
	return &UpdateEnterprise2Params{
		Context: ctx,
	}
}

// NewUpdateEnterprise2ParamsWithHTTPClient creates a new UpdateEnterprise2Params object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateEnterprise2ParamsWithHTTPClient(client *http.Client) *UpdateEnterprise2Params {
	return &UpdateEnterprise2Params{
		HTTPClient: client,
	}
}

/* UpdateEnterprise2Params contains all the parameters to send to the API endpoint
   for the update enterprise2 operation.

   Typically these are written to a http.Request.
*/
type UpdateEnterprise2Params struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body *swagger_models.Enterprise

	/* ID.

	   Unique system defined enterprise ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update enterprise2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateEnterprise2Params) WithDefaults() *UpdateEnterprise2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update enterprise2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateEnterprise2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update enterprise2 params
func (o *UpdateEnterprise2Params) WithTimeout(timeout time.Duration) *UpdateEnterprise2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update enterprise2 params
func (o *UpdateEnterprise2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update enterprise2 params
func (o *UpdateEnterprise2Params) WithContext(ctx context.Context) *UpdateEnterprise2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update enterprise2 params
func (o *UpdateEnterprise2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update enterprise2 params
func (o *UpdateEnterprise2Params) WithHTTPClient(client *http.Client) *UpdateEnterprise2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update enterprise2 params
func (o *UpdateEnterprise2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the update enterprise2 params
func (o *UpdateEnterprise2Params) WithXRequestID(xRequestID *string) *UpdateEnterprise2Params {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the update enterprise2 params
func (o *UpdateEnterprise2Params) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the update enterprise2 params
func (o *UpdateEnterprise2Params) WithBody(body *swagger_models.Enterprise) *UpdateEnterprise2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update enterprise2 params
func (o *UpdateEnterprise2Params) SetBody(body *swagger_models.Enterprise) {
	o.Body = body
}

// WithID adds the id to the update enterprise2 params
func (o *UpdateEnterprise2Params) WithID(id string) *UpdateEnterprise2Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the update enterprise2 params
func (o *UpdateEnterprise2Params) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateEnterprise2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

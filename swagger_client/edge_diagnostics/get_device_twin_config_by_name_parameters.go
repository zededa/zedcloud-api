// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package edge_diagnostics

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

// NewGetDeviceTwinConfigByNameParams creates a new GetDeviceTwinConfigByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeviceTwinConfigByNameParams() *GetDeviceTwinConfigByNameParams {
	return &GetDeviceTwinConfigByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceTwinConfigByNameParamsWithTimeout creates a new GetDeviceTwinConfigByNameParams object
// with the ability to set a timeout on a request.
func NewGetDeviceTwinConfigByNameParamsWithTimeout(timeout time.Duration) *GetDeviceTwinConfigByNameParams {
	return &GetDeviceTwinConfigByNameParams{
		timeout: timeout,
	}
}

// NewGetDeviceTwinConfigByNameParamsWithContext creates a new GetDeviceTwinConfigByNameParams object
// with the ability to set a context for a request.
func NewGetDeviceTwinConfigByNameParamsWithContext(ctx context.Context) *GetDeviceTwinConfigByNameParams {
	return &GetDeviceTwinConfigByNameParams{
		Context: ctx,
	}
}

// NewGetDeviceTwinConfigByNameParamsWithHTTPClient creates a new GetDeviceTwinConfigByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeviceTwinConfigByNameParamsWithHTTPClient(client *http.Client) *GetDeviceTwinConfigByNameParams {
	return &GetDeviceTwinConfigByNameParams{
		HTTPClient: client,
	}
}

/* GetDeviceTwinConfigByNameParams contains all the parameters to send to the API endpoint
   for the get device twin config by name operation.

   Typically these are written to a http.Request.
*/
type GetDeviceTwinConfigByNameParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* Name.

	   user defined device name for a device
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get device twin config by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceTwinConfigByNameParams) WithDefaults() *GetDeviceTwinConfigByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get device twin config by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceTwinConfigByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get device twin config by name params
func (o *GetDeviceTwinConfigByNameParams) WithTimeout(timeout time.Duration) *GetDeviceTwinConfigByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device twin config by name params
func (o *GetDeviceTwinConfigByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device twin config by name params
func (o *GetDeviceTwinConfigByNameParams) WithContext(ctx context.Context) *GetDeviceTwinConfigByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device twin config by name params
func (o *GetDeviceTwinConfigByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device twin config by name params
func (o *GetDeviceTwinConfigByNameParams) WithHTTPClient(client *http.Client) *GetDeviceTwinConfigByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device twin config by name params
func (o *GetDeviceTwinConfigByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get device twin config by name params
func (o *GetDeviceTwinConfigByNameParams) WithXRequestID(xRequestID *string) *GetDeviceTwinConfigByNameParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get device twin config by name params
func (o *GetDeviceTwinConfigByNameParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithName adds the name to the get device twin config by name params
func (o *GetDeviceTwinConfigByNameParams) WithName(name string) *GetDeviceTwinConfigByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get device twin config by name params
func (o *GetDeviceTwinConfigByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceTwinConfigByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package hardware_model

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

// NewGetGlobalHardwareModelParams creates a new GetGlobalHardwareModelParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetGlobalHardwareModelParams() *GetGlobalHardwareModelParams {
	return &GetGlobalHardwareModelParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetGlobalHardwareModelParamsWithTimeout creates a new GetGlobalHardwareModelParams object
// with the ability to set a timeout on a request.
func NewGetGlobalHardwareModelParamsWithTimeout(timeout time.Duration) *GetGlobalHardwareModelParams {
	return &GetGlobalHardwareModelParams{
		timeout: timeout,
	}
}

// NewGetGlobalHardwareModelParamsWithContext creates a new GetGlobalHardwareModelParams object
// with the ability to set a context for a request.
func NewGetGlobalHardwareModelParamsWithContext(ctx context.Context) *GetGlobalHardwareModelParams {
	return &GetGlobalHardwareModelParams{
		Context: ctx,
	}
}

// NewGetGlobalHardwareModelParamsWithHTTPClient creates a new GetGlobalHardwareModelParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetGlobalHardwareModelParamsWithHTTPClient(client *http.Client) *GetGlobalHardwareModelParams {
	return &GetGlobalHardwareModelParams{
		HTTPClient: client,
	}
}

/* GetGlobalHardwareModelParams contains all the parameters to send to the API endpoint
   for the get global hardware model operation.

   Typically these are written to a http.Request.
*/
type GetGlobalHardwareModelParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* ID.

	   System defined universally unique Id of the  model
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get global hardware model params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGlobalHardwareModelParams) WithDefaults() *GetGlobalHardwareModelParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get global hardware model params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGlobalHardwareModelParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get global hardware model params
func (o *GetGlobalHardwareModelParams) WithTimeout(timeout time.Duration) *GetGlobalHardwareModelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get global hardware model params
func (o *GetGlobalHardwareModelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get global hardware model params
func (o *GetGlobalHardwareModelParams) WithContext(ctx context.Context) *GetGlobalHardwareModelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get global hardware model params
func (o *GetGlobalHardwareModelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get global hardware model params
func (o *GetGlobalHardwareModelParams) WithHTTPClient(client *http.Client) *GetGlobalHardwareModelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get global hardware model params
func (o *GetGlobalHardwareModelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get global hardware model params
func (o *GetGlobalHardwareModelParams) WithXRequestID(xRequestID *string) *GetGlobalHardwareModelParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get global hardware model params
func (o *GetGlobalHardwareModelParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the get global hardware model params
func (o *GetGlobalHardwareModelParams) WithID(id string) *GetGlobalHardwareModelParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get global hardware model params
func (o *GetGlobalHardwareModelParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetGlobalHardwareModelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

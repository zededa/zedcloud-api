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

// NewHardwareModelUpdateHardwareBrandParams creates a new HardwareModelUpdateHardwareBrandParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHardwareModelUpdateHardwareBrandParams() *HardwareModelUpdateHardwareBrandParams {
	return &HardwareModelUpdateHardwareBrandParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHardwareModelUpdateHardwareBrandParamsWithTimeout creates a new HardwareModelUpdateHardwareBrandParams object
// with the ability to set a timeout on a request.
func NewHardwareModelUpdateHardwareBrandParamsWithTimeout(timeout time.Duration) *HardwareModelUpdateHardwareBrandParams {
	return &HardwareModelUpdateHardwareBrandParams{
		timeout: timeout,
	}
}

// NewHardwareModelUpdateHardwareBrandParamsWithContext creates a new HardwareModelUpdateHardwareBrandParams object
// with the ability to set a context for a request.
func NewHardwareModelUpdateHardwareBrandParamsWithContext(ctx context.Context) *HardwareModelUpdateHardwareBrandParams {
	return &HardwareModelUpdateHardwareBrandParams{
		Context: ctx,
	}
}

// NewHardwareModelUpdateHardwareBrandParamsWithHTTPClient creates a new HardwareModelUpdateHardwareBrandParams object
// with the ability to set a custom HTTPClient for a request.
func NewHardwareModelUpdateHardwareBrandParamsWithHTTPClient(client *http.Client) *HardwareModelUpdateHardwareBrandParams {
	return &HardwareModelUpdateHardwareBrandParams{
		HTTPClient: client,
	}
}

/* HardwareModelUpdateHardwareBrandParams contains all the parameters to send to the API endpoint
   for the hardware model update hardware brand operation.

   Typically these are written to a http.Request.
*/
type HardwareModelUpdateHardwareBrandParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body HardwareModelUpdateHardwareBrandBody

	/* ID.

	   System defined universally unique Id of the brand.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the hardware model update hardware brand params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HardwareModelUpdateHardwareBrandParams) WithDefaults() *HardwareModelUpdateHardwareBrandParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the hardware model update hardware brand params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HardwareModelUpdateHardwareBrandParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the hardware model update hardware brand params
func (o *HardwareModelUpdateHardwareBrandParams) WithTimeout(timeout time.Duration) *HardwareModelUpdateHardwareBrandParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the hardware model update hardware brand params
func (o *HardwareModelUpdateHardwareBrandParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the hardware model update hardware brand params
func (o *HardwareModelUpdateHardwareBrandParams) WithContext(ctx context.Context) *HardwareModelUpdateHardwareBrandParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the hardware model update hardware brand params
func (o *HardwareModelUpdateHardwareBrandParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the hardware model update hardware brand params
func (o *HardwareModelUpdateHardwareBrandParams) WithHTTPClient(client *http.Client) *HardwareModelUpdateHardwareBrandParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the hardware model update hardware brand params
func (o *HardwareModelUpdateHardwareBrandParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the hardware model update hardware brand params
func (o *HardwareModelUpdateHardwareBrandParams) WithXRequestID(xRequestID *string) *HardwareModelUpdateHardwareBrandParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the hardware model update hardware brand params
func (o *HardwareModelUpdateHardwareBrandParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the hardware model update hardware brand params
func (o *HardwareModelUpdateHardwareBrandParams) WithBody(body HardwareModelUpdateHardwareBrandBody) *HardwareModelUpdateHardwareBrandParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the hardware model update hardware brand params
func (o *HardwareModelUpdateHardwareBrandParams) SetBody(body HardwareModelUpdateHardwareBrandBody) {
	o.Body = body
}

// WithID adds the id to the hardware model update hardware brand params
func (o *HardwareModelUpdateHardwareBrandParams) WithID(id string) *HardwareModelUpdateHardwareBrandParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the hardware model update hardware brand params
func (o *HardwareModelUpdateHardwareBrandParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *HardwareModelUpdateHardwareBrandParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

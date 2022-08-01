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

	"github.com/zededa/zedcloud-api/swagger_models"
)

// NewHardwareModelCreateHardwareBrandParams creates a new HardwareModelCreateHardwareBrandParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHardwareModelCreateHardwareBrandParams() *HardwareModelCreateHardwareBrandParams {
	return &HardwareModelCreateHardwareBrandParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHardwareModelCreateHardwareBrandParamsWithTimeout creates a new HardwareModelCreateHardwareBrandParams object
// with the ability to set a timeout on a request.
func NewHardwareModelCreateHardwareBrandParamsWithTimeout(timeout time.Duration) *HardwareModelCreateHardwareBrandParams {
	return &HardwareModelCreateHardwareBrandParams{
		timeout: timeout,
	}
}

// NewHardwareModelCreateHardwareBrandParamsWithContext creates a new HardwareModelCreateHardwareBrandParams object
// with the ability to set a context for a request.
func NewHardwareModelCreateHardwareBrandParamsWithContext(ctx context.Context) *HardwareModelCreateHardwareBrandParams {
	return &HardwareModelCreateHardwareBrandParams{
		Context: ctx,
	}
}

// NewHardwareModelCreateHardwareBrandParamsWithHTTPClient creates a new HardwareModelCreateHardwareBrandParams object
// with the ability to set a custom HTTPClient for a request.
func NewHardwareModelCreateHardwareBrandParamsWithHTTPClient(client *http.Client) *HardwareModelCreateHardwareBrandParams {
	return &HardwareModelCreateHardwareBrandParams{
		HTTPClient: client,
	}
}

/* HardwareModelCreateHardwareBrandParams contains all the parameters to send to the API endpoint
   for the hardware model create hardware brand operation.

   Typically these are written to a http.Request.
*/
type HardwareModelCreateHardwareBrandParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body *swagger_models.SysBrand

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the hardware model create hardware brand params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HardwareModelCreateHardwareBrandParams) WithDefaults() *HardwareModelCreateHardwareBrandParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the hardware model create hardware brand params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HardwareModelCreateHardwareBrandParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the hardware model create hardware brand params
func (o *HardwareModelCreateHardwareBrandParams) WithTimeout(timeout time.Duration) *HardwareModelCreateHardwareBrandParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the hardware model create hardware brand params
func (o *HardwareModelCreateHardwareBrandParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the hardware model create hardware brand params
func (o *HardwareModelCreateHardwareBrandParams) WithContext(ctx context.Context) *HardwareModelCreateHardwareBrandParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the hardware model create hardware brand params
func (o *HardwareModelCreateHardwareBrandParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the hardware model create hardware brand params
func (o *HardwareModelCreateHardwareBrandParams) WithHTTPClient(client *http.Client) *HardwareModelCreateHardwareBrandParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the hardware model create hardware brand params
func (o *HardwareModelCreateHardwareBrandParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the hardware model create hardware brand params
func (o *HardwareModelCreateHardwareBrandParams) WithXRequestID(xRequestID *string) *HardwareModelCreateHardwareBrandParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the hardware model create hardware brand params
func (o *HardwareModelCreateHardwareBrandParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the hardware model create hardware brand params
func (o *HardwareModelCreateHardwareBrandParams) WithBody(body *swagger_models.SysBrand) *HardwareModelCreateHardwareBrandParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the hardware model create hardware brand params
func (o *HardwareModelCreateHardwareBrandParams) SetBody(body *swagger_models.SysBrand) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *HardwareModelCreateHardwareBrandParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
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

// NewHardwareModelGetGlobalHardwareModelByNameParams creates a new HardwareModelGetGlobalHardwareModelByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHardwareModelGetGlobalHardwareModelByNameParams() *HardwareModelGetGlobalHardwareModelByNameParams {
	return &HardwareModelGetGlobalHardwareModelByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHardwareModelGetGlobalHardwareModelByNameParamsWithTimeout creates a new HardwareModelGetGlobalHardwareModelByNameParams object
// with the ability to set a timeout on a request.
func NewHardwareModelGetGlobalHardwareModelByNameParamsWithTimeout(timeout time.Duration) *HardwareModelGetGlobalHardwareModelByNameParams {
	return &HardwareModelGetGlobalHardwareModelByNameParams{
		timeout: timeout,
	}
}

// NewHardwareModelGetGlobalHardwareModelByNameParamsWithContext creates a new HardwareModelGetGlobalHardwareModelByNameParams object
// with the ability to set a context for a request.
func NewHardwareModelGetGlobalHardwareModelByNameParamsWithContext(ctx context.Context) *HardwareModelGetGlobalHardwareModelByNameParams {
	return &HardwareModelGetGlobalHardwareModelByNameParams{
		Context: ctx,
	}
}

// NewHardwareModelGetGlobalHardwareModelByNameParamsWithHTTPClient creates a new HardwareModelGetGlobalHardwareModelByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewHardwareModelGetGlobalHardwareModelByNameParamsWithHTTPClient(client *http.Client) *HardwareModelGetGlobalHardwareModelByNameParams {
	return &HardwareModelGetGlobalHardwareModelByNameParams{
		HTTPClient: client,
	}
}

/*
HardwareModelGetGlobalHardwareModelByNameParams contains all the parameters to send to the API endpoint

	for the hardware model get global hardware model by name operation.

	Typically these are written to a http.Request.
*/
type HardwareModelGetGlobalHardwareModelByNameParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* Name.

	   User defined name of the model, unique across the enterprise. Once model is created, name can’t be changed.
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the hardware model get global hardware model by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HardwareModelGetGlobalHardwareModelByNameParams) WithDefaults() *HardwareModelGetGlobalHardwareModelByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the hardware model get global hardware model by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HardwareModelGetGlobalHardwareModelByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the hardware model get global hardware model by name params
func (o *HardwareModelGetGlobalHardwareModelByNameParams) WithTimeout(timeout time.Duration) *HardwareModelGetGlobalHardwareModelByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the hardware model get global hardware model by name params
func (o *HardwareModelGetGlobalHardwareModelByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the hardware model get global hardware model by name params
func (o *HardwareModelGetGlobalHardwareModelByNameParams) WithContext(ctx context.Context) *HardwareModelGetGlobalHardwareModelByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the hardware model get global hardware model by name params
func (o *HardwareModelGetGlobalHardwareModelByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the hardware model get global hardware model by name params
func (o *HardwareModelGetGlobalHardwareModelByNameParams) WithHTTPClient(client *http.Client) *HardwareModelGetGlobalHardwareModelByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the hardware model get global hardware model by name params
func (o *HardwareModelGetGlobalHardwareModelByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the hardware model get global hardware model by name params
func (o *HardwareModelGetGlobalHardwareModelByNameParams) WithXRequestID(xRequestID *string) *HardwareModelGetGlobalHardwareModelByNameParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the hardware model get global hardware model by name params
func (o *HardwareModelGetGlobalHardwareModelByNameParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithName adds the name to the hardware model get global hardware model by name params
func (o *HardwareModelGetGlobalHardwareModelByNameParams) WithName(name string) *HardwareModelGetGlobalHardwareModelByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the hardware model get global hardware model by name params
func (o *HardwareModelGetGlobalHardwareModelByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *HardwareModelGetGlobalHardwareModelByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

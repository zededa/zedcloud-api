// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package volume_instance_configuration

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

// NewVolumeInstanceConfigurationGetVolumeInstanceByNameParams creates a new VolumeInstanceConfigurationGetVolumeInstanceByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVolumeInstanceConfigurationGetVolumeInstanceByNameParams() *VolumeInstanceConfigurationGetVolumeInstanceByNameParams {
	return &VolumeInstanceConfigurationGetVolumeInstanceByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVolumeInstanceConfigurationGetVolumeInstanceByNameParamsWithTimeout creates a new VolumeInstanceConfigurationGetVolumeInstanceByNameParams object
// with the ability to set a timeout on a request.
func NewVolumeInstanceConfigurationGetVolumeInstanceByNameParamsWithTimeout(timeout time.Duration) *VolumeInstanceConfigurationGetVolumeInstanceByNameParams {
	return &VolumeInstanceConfigurationGetVolumeInstanceByNameParams{
		timeout: timeout,
	}
}

// NewVolumeInstanceConfigurationGetVolumeInstanceByNameParamsWithContext creates a new VolumeInstanceConfigurationGetVolumeInstanceByNameParams object
// with the ability to set a context for a request.
func NewVolumeInstanceConfigurationGetVolumeInstanceByNameParamsWithContext(ctx context.Context) *VolumeInstanceConfigurationGetVolumeInstanceByNameParams {
	return &VolumeInstanceConfigurationGetVolumeInstanceByNameParams{
		Context: ctx,
	}
}

// NewVolumeInstanceConfigurationGetVolumeInstanceByNameParamsWithHTTPClient creates a new VolumeInstanceConfigurationGetVolumeInstanceByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewVolumeInstanceConfigurationGetVolumeInstanceByNameParamsWithHTTPClient(client *http.Client) *VolumeInstanceConfigurationGetVolumeInstanceByNameParams {
	return &VolumeInstanceConfigurationGetVolumeInstanceByNameParams{
		HTTPClient: client,
	}
}

/* VolumeInstanceConfigurationGetVolumeInstanceByNameParams contains all the parameters to send to the API endpoint
   for the volume instance configuration get volume instance by name operation.

   Typically these are written to a http.Request.
*/
type VolumeInstanceConfigurationGetVolumeInstanceByNameParams struct {

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

// WithDefaults hydrates default values in the volume instance configuration get volume instance by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameParams) WithDefaults() *VolumeInstanceConfigurationGetVolumeInstanceByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the volume instance configuration get volume instance by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the volume instance configuration get volume instance by name params
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameParams) WithTimeout(timeout time.Duration) *VolumeInstanceConfigurationGetVolumeInstanceByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the volume instance configuration get volume instance by name params
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the volume instance configuration get volume instance by name params
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameParams) WithContext(ctx context.Context) *VolumeInstanceConfigurationGetVolumeInstanceByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the volume instance configuration get volume instance by name params
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the volume instance configuration get volume instance by name params
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameParams) WithHTTPClient(client *http.Client) *VolumeInstanceConfigurationGetVolumeInstanceByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the volume instance configuration get volume instance by name params
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the volume instance configuration get volume instance by name params
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameParams) WithXRequestID(xRequestID *string) *VolumeInstanceConfigurationGetVolumeInstanceByNameParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the volume instance configuration get volume instance by name params
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithName adds the name to the volume instance configuration get volume instance by name params
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameParams) WithName(name string) *VolumeInstanceConfigurationGetVolumeInstanceByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the volume instance configuration get volume instance by name params
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
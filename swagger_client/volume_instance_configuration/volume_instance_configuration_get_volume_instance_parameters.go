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

// NewVolumeInstanceConfigurationGetVolumeInstanceParams creates a new VolumeInstanceConfigurationGetVolumeInstanceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVolumeInstanceConfigurationGetVolumeInstanceParams() *VolumeInstanceConfigurationGetVolumeInstanceParams {
	return &VolumeInstanceConfigurationGetVolumeInstanceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVolumeInstanceConfigurationGetVolumeInstanceParamsWithTimeout creates a new VolumeInstanceConfigurationGetVolumeInstanceParams object
// with the ability to set a timeout on a request.
func NewVolumeInstanceConfigurationGetVolumeInstanceParamsWithTimeout(timeout time.Duration) *VolumeInstanceConfigurationGetVolumeInstanceParams {
	return &VolumeInstanceConfigurationGetVolumeInstanceParams{
		timeout: timeout,
	}
}

// NewVolumeInstanceConfigurationGetVolumeInstanceParamsWithContext creates a new VolumeInstanceConfigurationGetVolumeInstanceParams object
// with the ability to set a context for a request.
func NewVolumeInstanceConfigurationGetVolumeInstanceParamsWithContext(ctx context.Context) *VolumeInstanceConfigurationGetVolumeInstanceParams {
	return &VolumeInstanceConfigurationGetVolumeInstanceParams{
		Context: ctx,
	}
}

// NewVolumeInstanceConfigurationGetVolumeInstanceParamsWithHTTPClient creates a new VolumeInstanceConfigurationGetVolumeInstanceParams object
// with the ability to set a custom HTTPClient for a request.
func NewVolumeInstanceConfigurationGetVolumeInstanceParamsWithHTTPClient(client *http.Client) *VolumeInstanceConfigurationGetVolumeInstanceParams {
	return &VolumeInstanceConfigurationGetVolumeInstanceParams{
		HTTPClient: client,
	}
}

/* VolumeInstanceConfigurationGetVolumeInstanceParams contains all the parameters to send to the API endpoint
   for the volume instance configuration get volume instance operation.

   Typically these are written to a http.Request.
*/
type VolumeInstanceConfigurationGetVolumeInstanceParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the volume instance configuration get volume instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VolumeInstanceConfigurationGetVolumeInstanceParams) WithDefaults() *VolumeInstanceConfigurationGetVolumeInstanceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the volume instance configuration get volume instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VolumeInstanceConfigurationGetVolumeInstanceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the volume instance configuration get volume instance params
func (o *VolumeInstanceConfigurationGetVolumeInstanceParams) WithTimeout(timeout time.Duration) *VolumeInstanceConfigurationGetVolumeInstanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the volume instance configuration get volume instance params
func (o *VolumeInstanceConfigurationGetVolumeInstanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the volume instance configuration get volume instance params
func (o *VolumeInstanceConfigurationGetVolumeInstanceParams) WithContext(ctx context.Context) *VolumeInstanceConfigurationGetVolumeInstanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the volume instance configuration get volume instance params
func (o *VolumeInstanceConfigurationGetVolumeInstanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the volume instance configuration get volume instance params
func (o *VolumeInstanceConfigurationGetVolumeInstanceParams) WithHTTPClient(client *http.Client) *VolumeInstanceConfigurationGetVolumeInstanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the volume instance configuration get volume instance params
func (o *VolumeInstanceConfigurationGetVolumeInstanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the volume instance configuration get volume instance params
func (o *VolumeInstanceConfigurationGetVolumeInstanceParams) WithXRequestID(xRequestID *string) *VolumeInstanceConfigurationGetVolumeInstanceParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the volume instance configuration get volume instance params
func (o *VolumeInstanceConfigurationGetVolumeInstanceParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the volume instance configuration get volume instance params
func (o *VolumeInstanceConfigurationGetVolumeInstanceParams) WithID(id string) *VolumeInstanceConfigurationGetVolumeInstanceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the volume instance configuration get volume instance params
func (o *VolumeInstanceConfigurationGetVolumeInstanceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *VolumeInstanceConfigurationGetVolumeInstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

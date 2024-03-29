// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package volume_instance_status

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

// NewVolumeInstanceStatusGetVolumeInstanceStatusParams creates a new VolumeInstanceStatusGetVolumeInstanceStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVolumeInstanceStatusGetVolumeInstanceStatusParams() *VolumeInstanceStatusGetVolumeInstanceStatusParams {
	return &VolumeInstanceStatusGetVolumeInstanceStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVolumeInstanceStatusGetVolumeInstanceStatusParamsWithTimeout creates a new VolumeInstanceStatusGetVolumeInstanceStatusParams object
// with the ability to set a timeout on a request.
func NewVolumeInstanceStatusGetVolumeInstanceStatusParamsWithTimeout(timeout time.Duration) *VolumeInstanceStatusGetVolumeInstanceStatusParams {
	return &VolumeInstanceStatusGetVolumeInstanceStatusParams{
		timeout: timeout,
	}
}

// NewVolumeInstanceStatusGetVolumeInstanceStatusParamsWithContext creates a new VolumeInstanceStatusGetVolumeInstanceStatusParams object
// with the ability to set a context for a request.
func NewVolumeInstanceStatusGetVolumeInstanceStatusParamsWithContext(ctx context.Context) *VolumeInstanceStatusGetVolumeInstanceStatusParams {
	return &VolumeInstanceStatusGetVolumeInstanceStatusParams{
		Context: ctx,
	}
}

// NewVolumeInstanceStatusGetVolumeInstanceStatusParamsWithHTTPClient creates a new VolumeInstanceStatusGetVolumeInstanceStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewVolumeInstanceStatusGetVolumeInstanceStatusParamsWithHTTPClient(client *http.Client) *VolumeInstanceStatusGetVolumeInstanceStatusParams {
	return &VolumeInstanceStatusGetVolumeInstanceStatusParams{
		HTTPClient: client,
	}
}

/*
VolumeInstanceStatusGetVolumeInstanceStatusParams contains all the parameters to send to the API endpoint

	for the volume instance status get volume instance status operation.

	Typically these are written to a http.Request.
*/
type VolumeInstanceStatusGetVolumeInstanceStatusParams struct {

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

// WithDefaults hydrates default values in the volume instance status get volume instance status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VolumeInstanceStatusGetVolumeInstanceStatusParams) WithDefaults() *VolumeInstanceStatusGetVolumeInstanceStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the volume instance status get volume instance status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VolumeInstanceStatusGetVolumeInstanceStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the volume instance status get volume instance status params
func (o *VolumeInstanceStatusGetVolumeInstanceStatusParams) WithTimeout(timeout time.Duration) *VolumeInstanceStatusGetVolumeInstanceStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the volume instance status get volume instance status params
func (o *VolumeInstanceStatusGetVolumeInstanceStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the volume instance status get volume instance status params
func (o *VolumeInstanceStatusGetVolumeInstanceStatusParams) WithContext(ctx context.Context) *VolumeInstanceStatusGetVolumeInstanceStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the volume instance status get volume instance status params
func (o *VolumeInstanceStatusGetVolumeInstanceStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the volume instance status get volume instance status params
func (o *VolumeInstanceStatusGetVolumeInstanceStatusParams) WithHTTPClient(client *http.Client) *VolumeInstanceStatusGetVolumeInstanceStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the volume instance status get volume instance status params
func (o *VolumeInstanceStatusGetVolumeInstanceStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the volume instance status get volume instance status params
func (o *VolumeInstanceStatusGetVolumeInstanceStatusParams) WithXRequestID(xRequestID *string) *VolumeInstanceStatusGetVolumeInstanceStatusParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the volume instance status get volume instance status params
func (o *VolumeInstanceStatusGetVolumeInstanceStatusParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the volume instance status get volume instance status params
func (o *VolumeInstanceStatusGetVolumeInstanceStatusParams) WithID(id string) *VolumeInstanceStatusGetVolumeInstanceStatusParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the volume instance status get volume instance status params
func (o *VolumeInstanceStatusGetVolumeInstanceStatusParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *VolumeInstanceStatusGetVolumeInstanceStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

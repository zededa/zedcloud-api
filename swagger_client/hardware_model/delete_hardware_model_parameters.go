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

// NewDeleteHardwareModelParams creates a new DeleteHardwareModelParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteHardwareModelParams() *DeleteHardwareModelParams {
	return &DeleteHardwareModelParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteHardwareModelParamsWithTimeout creates a new DeleteHardwareModelParams object
// with the ability to set a timeout on a request.
func NewDeleteHardwareModelParamsWithTimeout(timeout time.Duration) *DeleteHardwareModelParams {
	return &DeleteHardwareModelParams{
		timeout: timeout,
	}
}

// NewDeleteHardwareModelParamsWithContext creates a new DeleteHardwareModelParams object
// with the ability to set a context for a request.
func NewDeleteHardwareModelParamsWithContext(ctx context.Context) *DeleteHardwareModelParams {
	return &DeleteHardwareModelParams{
		Context: ctx,
	}
}

// NewDeleteHardwareModelParamsWithHTTPClient creates a new DeleteHardwareModelParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteHardwareModelParamsWithHTTPClient(client *http.Client) *DeleteHardwareModelParams {
	return &DeleteHardwareModelParams{
		HTTPClient: client,
	}
}

/* DeleteHardwareModelParams contains all the parameters to send to the API endpoint
   for the delete hardware model operation.

   Typically these are written to a http.Request.
*/
type DeleteHardwareModelParams struct {

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

// WithDefaults hydrates default values in the delete hardware model params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteHardwareModelParams) WithDefaults() *DeleteHardwareModelParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete hardware model params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteHardwareModelParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete hardware model params
func (o *DeleteHardwareModelParams) WithTimeout(timeout time.Duration) *DeleteHardwareModelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete hardware model params
func (o *DeleteHardwareModelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete hardware model params
func (o *DeleteHardwareModelParams) WithContext(ctx context.Context) *DeleteHardwareModelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete hardware model params
func (o *DeleteHardwareModelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete hardware model params
func (o *DeleteHardwareModelParams) WithHTTPClient(client *http.Client) *DeleteHardwareModelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete hardware model params
func (o *DeleteHardwareModelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the delete hardware model params
func (o *DeleteHardwareModelParams) WithXRequestID(xRequestID *string) *DeleteHardwareModelParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the delete hardware model params
func (o *DeleteHardwareModelParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the delete hardware model params
func (o *DeleteHardwareModelParams) WithID(id string) *DeleteHardwareModelParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete hardware model params
func (o *DeleteHardwareModelParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteHardwareModelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

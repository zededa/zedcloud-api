// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package image_configuration

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

// NewUpdateImageParams creates a new UpdateImageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateImageParams() *UpdateImageParams {
	return &UpdateImageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateImageParamsWithTimeout creates a new UpdateImageParams object
// with the ability to set a timeout on a request.
func NewUpdateImageParamsWithTimeout(timeout time.Duration) *UpdateImageParams {
	return &UpdateImageParams{
		timeout: timeout,
	}
}

// NewUpdateImageParamsWithContext creates a new UpdateImageParams object
// with the ability to set a context for a request.
func NewUpdateImageParamsWithContext(ctx context.Context) *UpdateImageParams {
	return &UpdateImageParams{
		Context: ctx,
	}
}

// NewUpdateImageParamsWithHTTPClient creates a new UpdateImageParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateImageParamsWithHTTPClient(client *http.Client) *UpdateImageParams {
	return &UpdateImageParams{
		HTTPClient: client,
	}
}

/* UpdateImageParams contains all the parameters to send to the API endpoint
   for the update image operation.

   Typically these are written to a http.Request.
*/
type UpdateImageParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body *swagger_models.ImageConfig

	/* ID.

	   System defined universally unique Id of the image.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateImageParams) WithDefaults() *UpdateImageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateImageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update image params
func (o *UpdateImageParams) WithTimeout(timeout time.Duration) *UpdateImageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update image params
func (o *UpdateImageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update image params
func (o *UpdateImageParams) WithContext(ctx context.Context) *UpdateImageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update image params
func (o *UpdateImageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update image params
func (o *UpdateImageParams) WithHTTPClient(client *http.Client) *UpdateImageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update image params
func (o *UpdateImageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the update image params
func (o *UpdateImageParams) WithXRequestID(xRequestID *string) *UpdateImageParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the update image params
func (o *UpdateImageParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the update image params
func (o *UpdateImageParams) WithBody(body *swagger_models.ImageConfig) *UpdateImageParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update image params
func (o *UpdateImageParams) SetBody(body *swagger_models.ImageConfig) {
	o.Body = body
}

// WithID adds the id to the update image params
func (o *UpdateImageParams) WithID(id string) *UpdateImageParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update image params
func (o *UpdateImageParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateImageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

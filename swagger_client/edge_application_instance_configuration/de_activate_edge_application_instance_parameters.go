// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package edge_application_instance_configuration

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

// NewDeActivateEdgeApplicationInstanceParams creates a new DeActivateEdgeApplicationInstanceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeActivateEdgeApplicationInstanceParams() *DeActivateEdgeApplicationInstanceParams {
	return &DeActivateEdgeApplicationInstanceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeActivateEdgeApplicationInstanceParamsWithTimeout creates a new DeActivateEdgeApplicationInstanceParams object
// with the ability to set a timeout on a request.
func NewDeActivateEdgeApplicationInstanceParamsWithTimeout(timeout time.Duration) *DeActivateEdgeApplicationInstanceParams {
	return &DeActivateEdgeApplicationInstanceParams{
		timeout: timeout,
	}
}

// NewDeActivateEdgeApplicationInstanceParamsWithContext creates a new DeActivateEdgeApplicationInstanceParams object
// with the ability to set a context for a request.
func NewDeActivateEdgeApplicationInstanceParamsWithContext(ctx context.Context) *DeActivateEdgeApplicationInstanceParams {
	return &DeActivateEdgeApplicationInstanceParams{
		Context: ctx,
	}
}

// NewDeActivateEdgeApplicationInstanceParamsWithHTTPClient creates a new DeActivateEdgeApplicationInstanceParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeActivateEdgeApplicationInstanceParamsWithHTTPClient(client *http.Client) *DeActivateEdgeApplicationInstanceParams {
	return &DeActivateEdgeApplicationInstanceParams{
		HTTPClient: client,
	}
}

/* DeActivateEdgeApplicationInstanceParams contains all the parameters to send to the API endpoint
   for the de activate edge application instance operation.

   Typically these are written to a http.Request.
*/
type DeActivateEdgeApplicationInstanceParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* ID.

	   System defined universally unique Id of the app instance
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the de activate edge application instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeActivateEdgeApplicationInstanceParams) WithDefaults() *DeActivateEdgeApplicationInstanceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the de activate edge application instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeActivateEdgeApplicationInstanceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the de activate edge application instance params
func (o *DeActivateEdgeApplicationInstanceParams) WithTimeout(timeout time.Duration) *DeActivateEdgeApplicationInstanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the de activate edge application instance params
func (o *DeActivateEdgeApplicationInstanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the de activate edge application instance params
func (o *DeActivateEdgeApplicationInstanceParams) WithContext(ctx context.Context) *DeActivateEdgeApplicationInstanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the de activate edge application instance params
func (o *DeActivateEdgeApplicationInstanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the de activate edge application instance params
func (o *DeActivateEdgeApplicationInstanceParams) WithHTTPClient(client *http.Client) *DeActivateEdgeApplicationInstanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the de activate edge application instance params
func (o *DeActivateEdgeApplicationInstanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the de activate edge application instance params
func (o *DeActivateEdgeApplicationInstanceParams) WithXRequestID(xRequestID *string) *DeActivateEdgeApplicationInstanceParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the de activate edge application instance params
func (o *DeActivateEdgeApplicationInstanceParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the de activate edge application instance params
func (o *DeActivateEdgeApplicationInstanceParams) WithID(id string) *DeActivateEdgeApplicationInstanceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the de activate edge application instance params
func (o *DeActivateEdgeApplicationInstanceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeActivateEdgeApplicationInstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

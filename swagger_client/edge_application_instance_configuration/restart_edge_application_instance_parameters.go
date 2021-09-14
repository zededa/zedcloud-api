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

// NewRestartEdgeApplicationInstanceParams creates a new RestartEdgeApplicationInstanceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRestartEdgeApplicationInstanceParams() *RestartEdgeApplicationInstanceParams {
	return &RestartEdgeApplicationInstanceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRestartEdgeApplicationInstanceParamsWithTimeout creates a new RestartEdgeApplicationInstanceParams object
// with the ability to set a timeout on a request.
func NewRestartEdgeApplicationInstanceParamsWithTimeout(timeout time.Duration) *RestartEdgeApplicationInstanceParams {
	return &RestartEdgeApplicationInstanceParams{
		timeout: timeout,
	}
}

// NewRestartEdgeApplicationInstanceParamsWithContext creates a new RestartEdgeApplicationInstanceParams object
// with the ability to set a context for a request.
func NewRestartEdgeApplicationInstanceParamsWithContext(ctx context.Context) *RestartEdgeApplicationInstanceParams {
	return &RestartEdgeApplicationInstanceParams{
		Context: ctx,
	}
}

// NewRestartEdgeApplicationInstanceParamsWithHTTPClient creates a new RestartEdgeApplicationInstanceParams object
// with the ability to set a custom HTTPClient for a request.
func NewRestartEdgeApplicationInstanceParamsWithHTTPClient(client *http.Client) *RestartEdgeApplicationInstanceParams {
	return &RestartEdgeApplicationInstanceParams{
		HTTPClient: client,
	}
}

/* RestartEdgeApplicationInstanceParams contains all the parameters to send to the API endpoint
   for the restart edge application instance operation.

   Typically these are written to a http.Request.
*/
type RestartEdgeApplicationInstanceParams struct {

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

// WithDefaults hydrates default values in the restart edge application instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestartEdgeApplicationInstanceParams) WithDefaults() *RestartEdgeApplicationInstanceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the restart edge application instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestartEdgeApplicationInstanceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the restart edge application instance params
func (o *RestartEdgeApplicationInstanceParams) WithTimeout(timeout time.Duration) *RestartEdgeApplicationInstanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the restart edge application instance params
func (o *RestartEdgeApplicationInstanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the restart edge application instance params
func (o *RestartEdgeApplicationInstanceParams) WithContext(ctx context.Context) *RestartEdgeApplicationInstanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the restart edge application instance params
func (o *RestartEdgeApplicationInstanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the restart edge application instance params
func (o *RestartEdgeApplicationInstanceParams) WithHTTPClient(client *http.Client) *RestartEdgeApplicationInstanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the restart edge application instance params
func (o *RestartEdgeApplicationInstanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the restart edge application instance params
func (o *RestartEdgeApplicationInstanceParams) WithXRequestID(xRequestID *string) *RestartEdgeApplicationInstanceParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the restart edge application instance params
func (o *RestartEdgeApplicationInstanceParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the restart edge application instance params
func (o *RestartEdgeApplicationInstanceParams) WithID(id string) *RestartEdgeApplicationInstanceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the restart edge application instance params
func (o *RestartEdgeApplicationInstanceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RestartEdgeApplicationInstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

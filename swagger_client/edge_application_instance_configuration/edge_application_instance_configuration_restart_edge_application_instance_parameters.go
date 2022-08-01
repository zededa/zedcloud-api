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

// NewEdgeApplicationInstanceConfigurationRestartEdgeApplicationInstanceParams creates a new EdgeApplicationInstanceConfigurationRestartEdgeApplicationInstanceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEdgeApplicationInstanceConfigurationRestartEdgeApplicationInstanceParams() *EdgeApplicationInstanceConfigurationRestartEdgeApplicationInstanceParams {
	return &EdgeApplicationInstanceConfigurationRestartEdgeApplicationInstanceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeApplicationInstanceConfigurationRestartEdgeApplicationInstanceParamsWithTimeout creates a new EdgeApplicationInstanceConfigurationRestartEdgeApplicationInstanceParams object
// with the ability to set a timeout on a request.
func NewEdgeApplicationInstanceConfigurationRestartEdgeApplicationInstanceParamsWithTimeout(timeout time.Duration) *EdgeApplicationInstanceConfigurationRestartEdgeApplicationInstanceParams {
	return &EdgeApplicationInstanceConfigurationRestartEdgeApplicationInstanceParams{
		timeout: timeout,
	}
}

// NewEdgeApplicationInstanceConfigurationRestartEdgeApplicationInstanceParamsWithContext creates a new EdgeApplicationInstanceConfigurationRestartEdgeApplicationInstanceParams object
// with the ability to set a context for a request.
func NewEdgeApplicationInstanceConfigurationRestartEdgeApplicationInstanceParamsWithContext(ctx context.Context) *EdgeApplicationInstanceConfigurationRestartEdgeApplicationInstanceParams {
	return &EdgeApplicationInstanceConfigurationRestartEdgeApplicationInstanceParams{
		Context: ctx,
	}
}

// NewEdgeApplicationInstanceConfigurationRestartEdgeApplicationInstanceParamsWithHTTPClient creates a new EdgeApplicationInstanceConfigurationRestartEdgeApplicationInstanceParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeApplicationInstanceConfigurationRestartEdgeApplicationInstanceParamsWithHTTPClient(client *http.Client) *EdgeApplicationInstanceConfigurationRestartEdgeApplicationInstanceParams {
	return &EdgeApplicationInstanceConfigurationRestartEdgeApplicationInstanceParams{
		HTTPClient: client,
	}
}

/* EdgeApplicationInstanceConfigurationRestartEdgeApplicationInstanceParams contains all the parameters to send to the API endpoint
   for the edge application instance configuration restart edge application instance operation.

   Typically these are written to a http.Request.
*/
type EdgeApplicationInstanceConfigurationRestartEdgeApplicationInstanceParams struct {

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

// WithDefaults hydrates default values in the edge application instance configuration restart edge application instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeApplicationInstanceConfigurationRestartEdgeApplicationInstanceParams) WithDefaults() *EdgeApplicationInstanceConfigurationRestartEdgeApplicationInstanceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge application instance configuration restart edge application instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeApplicationInstanceConfigurationRestartEdgeApplicationInstanceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge application instance configuration restart edge application instance params
func (o *EdgeApplicationInstanceConfigurationRestartEdgeApplicationInstanceParams) WithTimeout(timeout time.Duration) *EdgeApplicationInstanceConfigurationRestartEdgeApplicationInstanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge application instance configuration restart edge application instance params
func (o *EdgeApplicationInstanceConfigurationRestartEdgeApplicationInstanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge application instance configuration restart edge application instance params
func (o *EdgeApplicationInstanceConfigurationRestartEdgeApplicationInstanceParams) WithContext(ctx context.Context) *EdgeApplicationInstanceConfigurationRestartEdgeApplicationInstanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge application instance configuration restart edge application instance params
func (o *EdgeApplicationInstanceConfigurationRestartEdgeApplicationInstanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge application instance configuration restart edge application instance params
func (o *EdgeApplicationInstanceConfigurationRestartEdgeApplicationInstanceParams) WithHTTPClient(client *http.Client) *EdgeApplicationInstanceConfigurationRestartEdgeApplicationInstanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge application instance configuration restart edge application instance params
func (o *EdgeApplicationInstanceConfigurationRestartEdgeApplicationInstanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the edge application instance configuration restart edge application instance params
func (o *EdgeApplicationInstanceConfigurationRestartEdgeApplicationInstanceParams) WithXRequestID(xRequestID *string) *EdgeApplicationInstanceConfigurationRestartEdgeApplicationInstanceParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the edge application instance configuration restart edge application instance params
func (o *EdgeApplicationInstanceConfigurationRestartEdgeApplicationInstanceParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the edge application instance configuration restart edge application instance params
func (o *EdgeApplicationInstanceConfigurationRestartEdgeApplicationInstanceParams) WithID(id string) *EdgeApplicationInstanceConfigurationRestartEdgeApplicationInstanceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the edge application instance configuration restart edge application instance params
func (o *EdgeApplicationInstanceConfigurationRestartEdgeApplicationInstanceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeApplicationInstanceConfigurationRestartEdgeApplicationInstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

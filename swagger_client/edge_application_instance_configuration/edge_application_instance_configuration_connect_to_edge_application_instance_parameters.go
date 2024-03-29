// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
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

// NewEdgeApplicationInstanceConfigurationConnectToEdgeApplicationInstanceParams creates a new EdgeApplicationInstanceConfigurationConnectToEdgeApplicationInstanceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEdgeApplicationInstanceConfigurationConnectToEdgeApplicationInstanceParams() *EdgeApplicationInstanceConfigurationConnectToEdgeApplicationInstanceParams {
	return &EdgeApplicationInstanceConfigurationConnectToEdgeApplicationInstanceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeApplicationInstanceConfigurationConnectToEdgeApplicationInstanceParamsWithTimeout creates a new EdgeApplicationInstanceConfigurationConnectToEdgeApplicationInstanceParams object
// with the ability to set a timeout on a request.
func NewEdgeApplicationInstanceConfigurationConnectToEdgeApplicationInstanceParamsWithTimeout(timeout time.Duration) *EdgeApplicationInstanceConfigurationConnectToEdgeApplicationInstanceParams {
	return &EdgeApplicationInstanceConfigurationConnectToEdgeApplicationInstanceParams{
		timeout: timeout,
	}
}

// NewEdgeApplicationInstanceConfigurationConnectToEdgeApplicationInstanceParamsWithContext creates a new EdgeApplicationInstanceConfigurationConnectToEdgeApplicationInstanceParams object
// with the ability to set a context for a request.
func NewEdgeApplicationInstanceConfigurationConnectToEdgeApplicationInstanceParamsWithContext(ctx context.Context) *EdgeApplicationInstanceConfigurationConnectToEdgeApplicationInstanceParams {
	return &EdgeApplicationInstanceConfigurationConnectToEdgeApplicationInstanceParams{
		Context: ctx,
	}
}

// NewEdgeApplicationInstanceConfigurationConnectToEdgeApplicationInstanceParamsWithHTTPClient creates a new EdgeApplicationInstanceConfigurationConnectToEdgeApplicationInstanceParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeApplicationInstanceConfigurationConnectToEdgeApplicationInstanceParamsWithHTTPClient(client *http.Client) *EdgeApplicationInstanceConfigurationConnectToEdgeApplicationInstanceParams {
	return &EdgeApplicationInstanceConfigurationConnectToEdgeApplicationInstanceParams{
		HTTPClient: client,
	}
}

/*
EdgeApplicationInstanceConfigurationConnectToEdgeApplicationInstanceParams contains all the parameters to send to the API endpoint

	for the edge application instance configuration connect to edge application instance operation.

	Typically these are written to a http.Request.
*/
type EdgeApplicationInstanceConfigurationConnectToEdgeApplicationInstanceParams struct {

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

// WithDefaults hydrates default values in the edge application instance configuration connect to edge application instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeApplicationInstanceConfigurationConnectToEdgeApplicationInstanceParams) WithDefaults() *EdgeApplicationInstanceConfigurationConnectToEdgeApplicationInstanceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge application instance configuration connect to edge application instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeApplicationInstanceConfigurationConnectToEdgeApplicationInstanceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge application instance configuration connect to edge application instance params
func (o *EdgeApplicationInstanceConfigurationConnectToEdgeApplicationInstanceParams) WithTimeout(timeout time.Duration) *EdgeApplicationInstanceConfigurationConnectToEdgeApplicationInstanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge application instance configuration connect to edge application instance params
func (o *EdgeApplicationInstanceConfigurationConnectToEdgeApplicationInstanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge application instance configuration connect to edge application instance params
func (o *EdgeApplicationInstanceConfigurationConnectToEdgeApplicationInstanceParams) WithContext(ctx context.Context) *EdgeApplicationInstanceConfigurationConnectToEdgeApplicationInstanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge application instance configuration connect to edge application instance params
func (o *EdgeApplicationInstanceConfigurationConnectToEdgeApplicationInstanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge application instance configuration connect to edge application instance params
func (o *EdgeApplicationInstanceConfigurationConnectToEdgeApplicationInstanceParams) WithHTTPClient(client *http.Client) *EdgeApplicationInstanceConfigurationConnectToEdgeApplicationInstanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge application instance configuration connect to edge application instance params
func (o *EdgeApplicationInstanceConfigurationConnectToEdgeApplicationInstanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the edge application instance configuration connect to edge application instance params
func (o *EdgeApplicationInstanceConfigurationConnectToEdgeApplicationInstanceParams) WithXRequestID(xRequestID *string) *EdgeApplicationInstanceConfigurationConnectToEdgeApplicationInstanceParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the edge application instance configuration connect to edge application instance params
func (o *EdgeApplicationInstanceConfigurationConnectToEdgeApplicationInstanceParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the edge application instance configuration connect to edge application instance params
func (o *EdgeApplicationInstanceConfigurationConnectToEdgeApplicationInstanceParams) WithID(id string) *EdgeApplicationInstanceConfigurationConnectToEdgeApplicationInstanceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the edge application instance configuration connect to edge application instance params
func (o *EdgeApplicationInstanceConfigurationConnectToEdgeApplicationInstanceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeApplicationInstanceConfigurationConnectToEdgeApplicationInstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

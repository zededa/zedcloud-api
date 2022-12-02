// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package edge_application_instance_status

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

// NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusByNameParams creates a new EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusByNameParams() *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusByNameParams {
	return &EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusByNameParamsWithTimeout creates a new EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusByNameParams object
// with the ability to set a timeout on a request.
func NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusByNameParamsWithTimeout(timeout time.Duration) *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusByNameParams {
	return &EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusByNameParams{
		timeout: timeout,
	}
}

// NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusByNameParamsWithContext creates a new EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusByNameParams object
// with the ability to set a context for a request.
func NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusByNameParamsWithContext(ctx context.Context) *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusByNameParams {
	return &EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusByNameParams{
		Context: ctx,
	}
}

// NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusByNameParamsWithHTTPClient creates a new EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusByNameParamsWithHTTPClient(client *http.Client) *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusByNameParams {
	return &EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusByNameParams{
		HTTPClient: client,
	}
}

/*
EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusByNameParams contains all the parameters to send to the API endpoint

	for the edge application instance status get edge application instance status by name operation.

	Typically these are written to a http.Request.
*/
type EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusByNameParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* Name.

	   User defined name of the app instance, unique across the enterprise. Once app instance is created, name can’t be changed
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edge application instance status get edge application instance status by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusByNameParams) WithDefaults() *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge application instance status get edge application instance status by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge application instance status get edge application instance status by name params
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusByNameParams) WithTimeout(timeout time.Duration) *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge application instance status get edge application instance status by name params
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge application instance status get edge application instance status by name params
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusByNameParams) WithContext(ctx context.Context) *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge application instance status get edge application instance status by name params
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge application instance status get edge application instance status by name params
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusByNameParams) WithHTTPClient(client *http.Client) *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge application instance status get edge application instance status by name params
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the edge application instance status get edge application instance status by name params
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusByNameParams) WithXRequestID(xRequestID *string) *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusByNameParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the edge application instance status get edge application instance status by name params
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusByNameParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithName adds the name to the edge application instance status get edge application instance status by name params
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusByNameParams) WithName(name string) *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the edge application instance status get edge application instance status by name params
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

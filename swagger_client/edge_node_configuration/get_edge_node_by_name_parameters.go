// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package edge_node_configuration

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

// NewGetEdgeNodeByNameParams creates a new GetEdgeNodeByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetEdgeNodeByNameParams() *GetEdgeNodeByNameParams {
	return &GetEdgeNodeByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetEdgeNodeByNameParamsWithTimeout creates a new GetEdgeNodeByNameParams object
// with the ability to set a timeout on a request.
func NewGetEdgeNodeByNameParamsWithTimeout(timeout time.Duration) *GetEdgeNodeByNameParams {
	return &GetEdgeNodeByNameParams{
		timeout: timeout,
	}
}

// NewGetEdgeNodeByNameParamsWithContext creates a new GetEdgeNodeByNameParams object
// with the ability to set a context for a request.
func NewGetEdgeNodeByNameParamsWithContext(ctx context.Context) *GetEdgeNodeByNameParams {
	return &GetEdgeNodeByNameParams{
		Context: ctx,
	}
}

// NewGetEdgeNodeByNameParamsWithHTTPClient creates a new GetEdgeNodeByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetEdgeNodeByNameParamsWithHTTPClient(client *http.Client) *GetEdgeNodeByNameParams {
	return &GetEdgeNodeByNameParams{
		HTTPClient: client,
	}
}

/* GetEdgeNodeByNameParams contains all the parameters to send to the API endpoint
   for the get edge node by name operation.

   Typically these are written to a http.Request.
*/
type GetEdgeNodeByNameParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* Name.

	   user defined device name for a device
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get edge node by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEdgeNodeByNameParams) WithDefaults() *GetEdgeNodeByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get edge node by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEdgeNodeByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get edge node by name params
func (o *GetEdgeNodeByNameParams) WithTimeout(timeout time.Duration) *GetEdgeNodeByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get edge node by name params
func (o *GetEdgeNodeByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get edge node by name params
func (o *GetEdgeNodeByNameParams) WithContext(ctx context.Context) *GetEdgeNodeByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get edge node by name params
func (o *GetEdgeNodeByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get edge node by name params
func (o *GetEdgeNodeByNameParams) WithHTTPClient(client *http.Client) *GetEdgeNodeByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get edge node by name params
func (o *GetEdgeNodeByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get edge node by name params
func (o *GetEdgeNodeByNameParams) WithXRequestID(xRequestID *string) *GetEdgeNodeByNameParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get edge node by name params
func (o *GetEdgeNodeByNameParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithName adds the name to the get edge node by name params
func (o *GetEdgeNodeByNameParams) WithName(name string) *GetEdgeNodeByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get edge node by name params
func (o *GetEdgeNodeByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetEdgeNodeByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

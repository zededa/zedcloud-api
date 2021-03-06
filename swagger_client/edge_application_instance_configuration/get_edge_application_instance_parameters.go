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

// NewGetEdgeApplicationInstanceParams creates a new GetEdgeApplicationInstanceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetEdgeApplicationInstanceParams() *GetEdgeApplicationInstanceParams {
	return &GetEdgeApplicationInstanceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetEdgeApplicationInstanceParamsWithTimeout creates a new GetEdgeApplicationInstanceParams object
// with the ability to set a timeout on a request.
func NewGetEdgeApplicationInstanceParamsWithTimeout(timeout time.Duration) *GetEdgeApplicationInstanceParams {
	return &GetEdgeApplicationInstanceParams{
		timeout: timeout,
	}
}

// NewGetEdgeApplicationInstanceParamsWithContext creates a new GetEdgeApplicationInstanceParams object
// with the ability to set a context for a request.
func NewGetEdgeApplicationInstanceParamsWithContext(ctx context.Context) *GetEdgeApplicationInstanceParams {
	return &GetEdgeApplicationInstanceParams{
		Context: ctx,
	}
}

// NewGetEdgeApplicationInstanceParamsWithHTTPClient creates a new GetEdgeApplicationInstanceParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetEdgeApplicationInstanceParamsWithHTTPClient(client *http.Client) *GetEdgeApplicationInstanceParams {
	return &GetEdgeApplicationInstanceParams{
		HTTPClient: client,
	}
}

/* GetEdgeApplicationInstanceParams contains all the parameters to send to the API endpoint
   for the get edge application instance operation.

   Typically these are written to a http.Request.
*/
type GetEdgeApplicationInstanceParams struct {

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

// WithDefaults hydrates default values in the get edge application instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEdgeApplicationInstanceParams) WithDefaults() *GetEdgeApplicationInstanceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get edge application instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEdgeApplicationInstanceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get edge application instance params
func (o *GetEdgeApplicationInstanceParams) WithTimeout(timeout time.Duration) *GetEdgeApplicationInstanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get edge application instance params
func (o *GetEdgeApplicationInstanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get edge application instance params
func (o *GetEdgeApplicationInstanceParams) WithContext(ctx context.Context) *GetEdgeApplicationInstanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get edge application instance params
func (o *GetEdgeApplicationInstanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get edge application instance params
func (o *GetEdgeApplicationInstanceParams) WithHTTPClient(client *http.Client) *GetEdgeApplicationInstanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get edge application instance params
func (o *GetEdgeApplicationInstanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get edge application instance params
func (o *GetEdgeApplicationInstanceParams) WithXRequestID(xRequestID *string) *GetEdgeApplicationInstanceParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get edge application instance params
func (o *GetEdgeApplicationInstanceParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the get edge application instance params
func (o *GetEdgeApplicationInstanceParams) WithID(id string) *GetEdgeApplicationInstanceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get edge application instance params
func (o *GetEdgeApplicationInstanceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetEdgeApplicationInstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

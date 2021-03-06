// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package edge_network_configuration

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

// NewGetEdgeNetworkParams creates a new GetEdgeNetworkParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetEdgeNetworkParams() *GetEdgeNetworkParams {
	return &GetEdgeNetworkParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetEdgeNetworkParamsWithTimeout creates a new GetEdgeNetworkParams object
// with the ability to set a timeout on a request.
func NewGetEdgeNetworkParamsWithTimeout(timeout time.Duration) *GetEdgeNetworkParams {
	return &GetEdgeNetworkParams{
		timeout: timeout,
	}
}

// NewGetEdgeNetworkParamsWithContext creates a new GetEdgeNetworkParams object
// with the ability to set a context for a request.
func NewGetEdgeNetworkParamsWithContext(ctx context.Context) *GetEdgeNetworkParams {
	return &GetEdgeNetworkParams{
		Context: ctx,
	}
}

// NewGetEdgeNetworkParamsWithHTTPClient creates a new GetEdgeNetworkParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetEdgeNetworkParamsWithHTTPClient(client *http.Client) *GetEdgeNetworkParams {
	return &GetEdgeNetworkParams{
		HTTPClient: client,
	}
}

/* GetEdgeNetworkParams contains all the parameters to send to the API endpoint
   for the get edge network operation.

   Typically these are written to a http.Request.
*/
type GetEdgeNetworkParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* ID.

	   System defined universally unique Id of the network
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get edge network params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEdgeNetworkParams) WithDefaults() *GetEdgeNetworkParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get edge network params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEdgeNetworkParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get edge network params
func (o *GetEdgeNetworkParams) WithTimeout(timeout time.Duration) *GetEdgeNetworkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get edge network params
func (o *GetEdgeNetworkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get edge network params
func (o *GetEdgeNetworkParams) WithContext(ctx context.Context) *GetEdgeNetworkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get edge network params
func (o *GetEdgeNetworkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get edge network params
func (o *GetEdgeNetworkParams) WithHTTPClient(client *http.Client) *GetEdgeNetworkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get edge network params
func (o *GetEdgeNetworkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get edge network params
func (o *GetEdgeNetworkParams) WithXRequestID(xRequestID *string) *GetEdgeNetworkParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get edge network params
func (o *GetEdgeNetworkParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the get edge network params
func (o *GetEdgeNetworkParams) WithID(id string) *GetEdgeNetworkParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get edge network params
func (o *GetEdgeNetworkParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetEdgeNetworkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

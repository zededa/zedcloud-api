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

// NewDeleteEdgeNetworkParams creates a new DeleteEdgeNetworkParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteEdgeNetworkParams() *DeleteEdgeNetworkParams {
	return &DeleteEdgeNetworkParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteEdgeNetworkParamsWithTimeout creates a new DeleteEdgeNetworkParams object
// with the ability to set a timeout on a request.
func NewDeleteEdgeNetworkParamsWithTimeout(timeout time.Duration) *DeleteEdgeNetworkParams {
	return &DeleteEdgeNetworkParams{
		timeout: timeout,
	}
}

// NewDeleteEdgeNetworkParamsWithContext creates a new DeleteEdgeNetworkParams object
// with the ability to set a context for a request.
func NewDeleteEdgeNetworkParamsWithContext(ctx context.Context) *DeleteEdgeNetworkParams {
	return &DeleteEdgeNetworkParams{
		Context: ctx,
	}
}

// NewDeleteEdgeNetworkParamsWithHTTPClient creates a new DeleteEdgeNetworkParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteEdgeNetworkParamsWithHTTPClient(client *http.Client) *DeleteEdgeNetworkParams {
	return &DeleteEdgeNetworkParams{
		HTTPClient: client,
	}
}

/* DeleteEdgeNetworkParams contains all the parameters to send to the API endpoint
   for the delete edge network operation.

   Typically these are written to a http.Request.
*/
type DeleteEdgeNetworkParams struct {

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

// WithDefaults hydrates default values in the delete edge network params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteEdgeNetworkParams) WithDefaults() *DeleteEdgeNetworkParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete edge network params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteEdgeNetworkParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete edge network params
func (o *DeleteEdgeNetworkParams) WithTimeout(timeout time.Duration) *DeleteEdgeNetworkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete edge network params
func (o *DeleteEdgeNetworkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete edge network params
func (o *DeleteEdgeNetworkParams) WithContext(ctx context.Context) *DeleteEdgeNetworkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete edge network params
func (o *DeleteEdgeNetworkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete edge network params
func (o *DeleteEdgeNetworkParams) WithHTTPClient(client *http.Client) *DeleteEdgeNetworkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete edge network params
func (o *DeleteEdgeNetworkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the delete edge network params
func (o *DeleteEdgeNetworkParams) WithXRequestID(xRequestID *string) *DeleteEdgeNetworkParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the delete edge network params
func (o *DeleteEdgeNetworkParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the delete edge network params
func (o *DeleteEdgeNetworkParams) WithID(id string) *DeleteEdgeNetworkParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete edge network params
func (o *DeleteEdgeNetworkParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteEdgeNetworkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

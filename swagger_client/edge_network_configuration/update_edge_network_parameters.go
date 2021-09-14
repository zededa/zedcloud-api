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

	"github.com/zededa/zedcloud-api/swagger_models"
)

// NewUpdateEdgeNetworkParams creates a new UpdateEdgeNetworkParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateEdgeNetworkParams() *UpdateEdgeNetworkParams {
	return &UpdateEdgeNetworkParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateEdgeNetworkParamsWithTimeout creates a new UpdateEdgeNetworkParams object
// with the ability to set a timeout on a request.
func NewUpdateEdgeNetworkParamsWithTimeout(timeout time.Duration) *UpdateEdgeNetworkParams {
	return &UpdateEdgeNetworkParams{
		timeout: timeout,
	}
}

// NewUpdateEdgeNetworkParamsWithContext creates a new UpdateEdgeNetworkParams object
// with the ability to set a context for a request.
func NewUpdateEdgeNetworkParamsWithContext(ctx context.Context) *UpdateEdgeNetworkParams {
	return &UpdateEdgeNetworkParams{
		Context: ctx,
	}
}

// NewUpdateEdgeNetworkParamsWithHTTPClient creates a new UpdateEdgeNetworkParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateEdgeNetworkParamsWithHTTPClient(client *http.Client) *UpdateEdgeNetworkParams {
	return &UpdateEdgeNetworkParams{
		HTTPClient: client,
	}
}

/* UpdateEdgeNetworkParams contains all the parameters to send to the API endpoint
   for the update edge network operation.

   Typically these are written to a http.Request.
*/
type UpdateEdgeNetworkParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body *swagger_models.NetConfig

	/* ID.

	   System defined universally unique Id of the network
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update edge network params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateEdgeNetworkParams) WithDefaults() *UpdateEdgeNetworkParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update edge network params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateEdgeNetworkParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update edge network params
func (o *UpdateEdgeNetworkParams) WithTimeout(timeout time.Duration) *UpdateEdgeNetworkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update edge network params
func (o *UpdateEdgeNetworkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update edge network params
func (o *UpdateEdgeNetworkParams) WithContext(ctx context.Context) *UpdateEdgeNetworkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update edge network params
func (o *UpdateEdgeNetworkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update edge network params
func (o *UpdateEdgeNetworkParams) WithHTTPClient(client *http.Client) *UpdateEdgeNetworkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update edge network params
func (o *UpdateEdgeNetworkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the update edge network params
func (o *UpdateEdgeNetworkParams) WithXRequestID(xRequestID *string) *UpdateEdgeNetworkParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the update edge network params
func (o *UpdateEdgeNetworkParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the update edge network params
func (o *UpdateEdgeNetworkParams) WithBody(body *swagger_models.NetConfig) *UpdateEdgeNetworkParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update edge network params
func (o *UpdateEdgeNetworkParams) SetBody(body *swagger_models.NetConfig) {
	o.Body = body
}

// WithID adds the id to the update edge network params
func (o *UpdateEdgeNetworkParams) WithID(id string) *UpdateEdgeNetworkParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update edge network params
func (o *UpdateEdgeNetworkParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateEdgeNetworkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

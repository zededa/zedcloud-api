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

// NewCreateEdgeNetworkParams creates a new CreateEdgeNetworkParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateEdgeNetworkParams() *CreateEdgeNetworkParams {
	return &CreateEdgeNetworkParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateEdgeNetworkParamsWithTimeout creates a new CreateEdgeNetworkParams object
// with the ability to set a timeout on a request.
func NewCreateEdgeNetworkParamsWithTimeout(timeout time.Duration) *CreateEdgeNetworkParams {
	return &CreateEdgeNetworkParams{
		timeout: timeout,
	}
}

// NewCreateEdgeNetworkParamsWithContext creates a new CreateEdgeNetworkParams object
// with the ability to set a context for a request.
func NewCreateEdgeNetworkParamsWithContext(ctx context.Context) *CreateEdgeNetworkParams {
	return &CreateEdgeNetworkParams{
		Context: ctx,
	}
}

// NewCreateEdgeNetworkParamsWithHTTPClient creates a new CreateEdgeNetworkParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateEdgeNetworkParamsWithHTTPClient(client *http.Client) *CreateEdgeNetworkParams {
	return &CreateEdgeNetworkParams{
		HTTPClient: client,
	}
}

/* CreateEdgeNetworkParams contains all the parameters to send to the API endpoint
   for the create edge network operation.

   Typically these are written to a http.Request.
*/
type CreateEdgeNetworkParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body *swagger_models.NetConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create edge network params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateEdgeNetworkParams) WithDefaults() *CreateEdgeNetworkParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create edge network params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateEdgeNetworkParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create edge network params
func (o *CreateEdgeNetworkParams) WithTimeout(timeout time.Duration) *CreateEdgeNetworkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create edge network params
func (o *CreateEdgeNetworkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create edge network params
func (o *CreateEdgeNetworkParams) WithContext(ctx context.Context) *CreateEdgeNetworkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create edge network params
func (o *CreateEdgeNetworkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create edge network params
func (o *CreateEdgeNetworkParams) WithHTTPClient(client *http.Client) *CreateEdgeNetworkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create edge network params
func (o *CreateEdgeNetworkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the create edge network params
func (o *CreateEdgeNetworkParams) WithXRequestID(xRequestID *string) *CreateEdgeNetworkParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the create edge network params
func (o *CreateEdgeNetworkParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the create edge network params
func (o *CreateEdgeNetworkParams) WithBody(body *swagger_models.NetConfig) *CreateEdgeNetworkParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create edge network params
func (o *CreateEdgeNetworkParams) SetBody(body *swagger_models.NetConfig) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateEdgeNetworkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package edge_network_instance_configuration

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

// NewCreateEdgeNetworkInstanceParams creates a new CreateEdgeNetworkInstanceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateEdgeNetworkInstanceParams() *CreateEdgeNetworkInstanceParams {
	return &CreateEdgeNetworkInstanceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateEdgeNetworkInstanceParamsWithTimeout creates a new CreateEdgeNetworkInstanceParams object
// with the ability to set a timeout on a request.
func NewCreateEdgeNetworkInstanceParamsWithTimeout(timeout time.Duration) *CreateEdgeNetworkInstanceParams {
	return &CreateEdgeNetworkInstanceParams{
		timeout: timeout,
	}
}

// NewCreateEdgeNetworkInstanceParamsWithContext creates a new CreateEdgeNetworkInstanceParams object
// with the ability to set a context for a request.
func NewCreateEdgeNetworkInstanceParamsWithContext(ctx context.Context) *CreateEdgeNetworkInstanceParams {
	return &CreateEdgeNetworkInstanceParams{
		Context: ctx,
	}
}

// NewCreateEdgeNetworkInstanceParamsWithHTTPClient creates a new CreateEdgeNetworkInstanceParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateEdgeNetworkInstanceParamsWithHTTPClient(client *http.Client) *CreateEdgeNetworkInstanceParams {
	return &CreateEdgeNetworkInstanceParams{
		HTTPClient: client,
	}
}

/* CreateEdgeNetworkInstanceParams contains all the parameters to send to the API endpoint
   for the create edge network instance operation.

   Typically these are written to a http.Request.
*/
type CreateEdgeNetworkInstanceParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body *swagger_models.NetInstConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create edge network instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateEdgeNetworkInstanceParams) WithDefaults() *CreateEdgeNetworkInstanceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create edge network instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateEdgeNetworkInstanceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create edge network instance params
func (o *CreateEdgeNetworkInstanceParams) WithTimeout(timeout time.Duration) *CreateEdgeNetworkInstanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create edge network instance params
func (o *CreateEdgeNetworkInstanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create edge network instance params
func (o *CreateEdgeNetworkInstanceParams) WithContext(ctx context.Context) *CreateEdgeNetworkInstanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create edge network instance params
func (o *CreateEdgeNetworkInstanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create edge network instance params
func (o *CreateEdgeNetworkInstanceParams) WithHTTPClient(client *http.Client) *CreateEdgeNetworkInstanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create edge network instance params
func (o *CreateEdgeNetworkInstanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the create edge network instance params
func (o *CreateEdgeNetworkInstanceParams) WithXRequestID(xRequestID *string) *CreateEdgeNetworkInstanceParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the create edge network instance params
func (o *CreateEdgeNetworkInstanceParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the create edge network instance params
func (o *CreateEdgeNetworkInstanceParams) WithBody(body *swagger_models.NetInstConfig) *CreateEdgeNetworkInstanceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create edge network instance params
func (o *CreateEdgeNetworkInstanceParams) SetBody(body *swagger_models.NetInstConfig) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateEdgeNetworkInstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

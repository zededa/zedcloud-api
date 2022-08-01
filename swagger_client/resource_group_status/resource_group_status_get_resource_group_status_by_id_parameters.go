// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package resource_group_status

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

// NewResourceGroupStatusGetResourceGroupStatusByIDParams creates a new ResourceGroupStatusGetResourceGroupStatusByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewResourceGroupStatusGetResourceGroupStatusByIDParams() *ResourceGroupStatusGetResourceGroupStatusByIDParams {
	return &ResourceGroupStatusGetResourceGroupStatusByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewResourceGroupStatusGetResourceGroupStatusByIDParamsWithTimeout creates a new ResourceGroupStatusGetResourceGroupStatusByIDParams object
// with the ability to set a timeout on a request.
func NewResourceGroupStatusGetResourceGroupStatusByIDParamsWithTimeout(timeout time.Duration) *ResourceGroupStatusGetResourceGroupStatusByIDParams {
	return &ResourceGroupStatusGetResourceGroupStatusByIDParams{
		timeout: timeout,
	}
}

// NewResourceGroupStatusGetResourceGroupStatusByIDParamsWithContext creates a new ResourceGroupStatusGetResourceGroupStatusByIDParams object
// with the ability to set a context for a request.
func NewResourceGroupStatusGetResourceGroupStatusByIDParamsWithContext(ctx context.Context) *ResourceGroupStatusGetResourceGroupStatusByIDParams {
	return &ResourceGroupStatusGetResourceGroupStatusByIDParams{
		Context: ctx,
	}
}

// NewResourceGroupStatusGetResourceGroupStatusByIDParamsWithHTTPClient creates a new ResourceGroupStatusGetResourceGroupStatusByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewResourceGroupStatusGetResourceGroupStatusByIDParamsWithHTTPClient(client *http.Client) *ResourceGroupStatusGetResourceGroupStatusByIDParams {
	return &ResourceGroupStatusGetResourceGroupStatusByIDParams{
		HTTPClient: client,
	}
}

/* ResourceGroupStatusGetResourceGroupStatusByIDParams contains all the parameters to send to the API endpoint
   for the resource group status get resource group status by Id operation.

   Typically these are written to a http.Request.
*/
type ResourceGroupStatusGetResourceGroupStatusByIDParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* ID.

	   System defined universally unique Id of the resource group
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the resource group status get resource group status by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResourceGroupStatusGetResourceGroupStatusByIDParams) WithDefaults() *ResourceGroupStatusGetResourceGroupStatusByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the resource group status get resource group status by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResourceGroupStatusGetResourceGroupStatusByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the resource group status get resource group status by Id params
func (o *ResourceGroupStatusGetResourceGroupStatusByIDParams) WithTimeout(timeout time.Duration) *ResourceGroupStatusGetResourceGroupStatusByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the resource group status get resource group status by Id params
func (o *ResourceGroupStatusGetResourceGroupStatusByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the resource group status get resource group status by Id params
func (o *ResourceGroupStatusGetResourceGroupStatusByIDParams) WithContext(ctx context.Context) *ResourceGroupStatusGetResourceGroupStatusByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the resource group status get resource group status by Id params
func (o *ResourceGroupStatusGetResourceGroupStatusByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the resource group status get resource group status by Id params
func (o *ResourceGroupStatusGetResourceGroupStatusByIDParams) WithHTTPClient(client *http.Client) *ResourceGroupStatusGetResourceGroupStatusByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the resource group status get resource group status by Id params
func (o *ResourceGroupStatusGetResourceGroupStatusByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the resource group status get resource group status by Id params
func (o *ResourceGroupStatusGetResourceGroupStatusByIDParams) WithXRequestID(xRequestID *string) *ResourceGroupStatusGetResourceGroupStatusByIDParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the resource group status get resource group status by Id params
func (o *ResourceGroupStatusGetResourceGroupStatusByIDParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the resource group status get resource group status by Id params
func (o *ResourceGroupStatusGetResourceGroupStatusByIDParams) WithID(id string) *ResourceGroupStatusGetResourceGroupStatusByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the resource group status get resource group status by Id params
func (o *ResourceGroupStatusGetResourceGroupStatusByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ResourceGroupStatusGetResourceGroupStatusByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

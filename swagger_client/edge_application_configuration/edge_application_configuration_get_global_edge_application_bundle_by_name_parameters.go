// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package edge_application_configuration

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

// NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameParams creates a new EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameParams() *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameParams {
	return &EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameParamsWithTimeout creates a new EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameParams object
// with the ability to set a timeout on a request.
func NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameParamsWithTimeout(timeout time.Duration) *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameParams {
	return &EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameParams{
		timeout: timeout,
	}
}

// NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameParamsWithContext creates a new EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameParams object
// with the ability to set a context for a request.
func NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameParamsWithContext(ctx context.Context) *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameParams {
	return &EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameParams{
		Context: ctx,
	}
}

// NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameParamsWithHTTPClient creates a new EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameParamsWithHTTPClient(client *http.Client) *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameParams {
	return &EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameParams{
		HTTPClient: client,
	}
}

/*
EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameParams contains all the parameters to send to the API endpoint

	for the edge application configuration get global edge application bundle by name operation.

	Typically these are written to a http.Request.
*/
type EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* Name.

	   User defined name of the edge application, unique across the enterprise. Once object is created, name can’t be changed
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edge application configuration get global edge application bundle by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameParams) WithDefaults() *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge application configuration get global edge application bundle by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge application configuration get global edge application bundle by name params
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameParams) WithTimeout(timeout time.Duration) *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge application configuration get global edge application bundle by name params
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge application configuration get global edge application bundle by name params
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameParams) WithContext(ctx context.Context) *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge application configuration get global edge application bundle by name params
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge application configuration get global edge application bundle by name params
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameParams) WithHTTPClient(client *http.Client) *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge application configuration get global edge application bundle by name params
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the edge application configuration get global edge application bundle by name params
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameParams) WithXRequestID(xRequestID *string) *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the edge application configuration get global edge application bundle by name params
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithName adds the name to the edge application configuration get global edge application bundle by name params
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameParams) WithName(name string) *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the edge application configuration get global edge application bundle by name params
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

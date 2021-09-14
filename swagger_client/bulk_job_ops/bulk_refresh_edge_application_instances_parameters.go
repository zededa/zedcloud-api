// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package bulk_job_ops

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

// NewBulkRefreshEdgeApplicationInstancesParams creates a new BulkRefreshEdgeApplicationInstancesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBulkRefreshEdgeApplicationInstancesParams() *BulkRefreshEdgeApplicationInstancesParams {
	return &BulkRefreshEdgeApplicationInstancesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBulkRefreshEdgeApplicationInstancesParamsWithTimeout creates a new BulkRefreshEdgeApplicationInstancesParams object
// with the ability to set a timeout on a request.
func NewBulkRefreshEdgeApplicationInstancesParamsWithTimeout(timeout time.Duration) *BulkRefreshEdgeApplicationInstancesParams {
	return &BulkRefreshEdgeApplicationInstancesParams{
		timeout: timeout,
	}
}

// NewBulkRefreshEdgeApplicationInstancesParamsWithContext creates a new BulkRefreshEdgeApplicationInstancesParams object
// with the ability to set a context for a request.
func NewBulkRefreshEdgeApplicationInstancesParamsWithContext(ctx context.Context) *BulkRefreshEdgeApplicationInstancesParams {
	return &BulkRefreshEdgeApplicationInstancesParams{
		Context: ctx,
	}
}

// NewBulkRefreshEdgeApplicationInstancesParamsWithHTTPClient creates a new BulkRefreshEdgeApplicationInstancesParams object
// with the ability to set a custom HTTPClient for a request.
func NewBulkRefreshEdgeApplicationInstancesParamsWithHTTPClient(client *http.Client) *BulkRefreshEdgeApplicationInstancesParams {
	return &BulkRefreshEdgeApplicationInstancesParams{
		HTTPClient: client,
	}
}

/* BulkRefreshEdgeApplicationInstancesParams contains all the parameters to send to the API endpoint
   for the bulk refresh edge application instances operation.

   Typically these are written to a http.Request.
*/
type BulkRefreshEdgeApplicationInstancesParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body *swagger_models.BulkConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the bulk refresh edge application instances params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BulkRefreshEdgeApplicationInstancesParams) WithDefaults() *BulkRefreshEdgeApplicationInstancesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the bulk refresh edge application instances params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BulkRefreshEdgeApplicationInstancesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the bulk refresh edge application instances params
func (o *BulkRefreshEdgeApplicationInstancesParams) WithTimeout(timeout time.Duration) *BulkRefreshEdgeApplicationInstancesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bulk refresh edge application instances params
func (o *BulkRefreshEdgeApplicationInstancesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bulk refresh edge application instances params
func (o *BulkRefreshEdgeApplicationInstancesParams) WithContext(ctx context.Context) *BulkRefreshEdgeApplicationInstancesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bulk refresh edge application instances params
func (o *BulkRefreshEdgeApplicationInstancesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bulk refresh edge application instances params
func (o *BulkRefreshEdgeApplicationInstancesParams) WithHTTPClient(client *http.Client) *BulkRefreshEdgeApplicationInstancesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bulk refresh edge application instances params
func (o *BulkRefreshEdgeApplicationInstancesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the bulk refresh edge application instances params
func (o *BulkRefreshEdgeApplicationInstancesParams) WithXRequestID(xRequestID *string) *BulkRefreshEdgeApplicationInstancesParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the bulk refresh edge application instances params
func (o *BulkRefreshEdgeApplicationInstancesParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the bulk refresh edge application instances params
func (o *BulkRefreshEdgeApplicationInstancesParams) WithBody(body *swagger_models.BulkConfig) *BulkRefreshEdgeApplicationInstancesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the bulk refresh edge application instances params
func (o *BulkRefreshEdgeApplicationInstancesParams) SetBody(body *swagger_models.BulkConfig) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *BulkRefreshEdgeApplicationInstancesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

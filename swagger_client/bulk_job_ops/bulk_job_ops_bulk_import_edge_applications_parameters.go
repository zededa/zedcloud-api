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

// NewBulkJobOpsBulkImportEdgeApplicationsParams creates a new BulkJobOpsBulkImportEdgeApplicationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBulkJobOpsBulkImportEdgeApplicationsParams() *BulkJobOpsBulkImportEdgeApplicationsParams {
	return &BulkJobOpsBulkImportEdgeApplicationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBulkJobOpsBulkImportEdgeApplicationsParamsWithTimeout creates a new BulkJobOpsBulkImportEdgeApplicationsParams object
// with the ability to set a timeout on a request.
func NewBulkJobOpsBulkImportEdgeApplicationsParamsWithTimeout(timeout time.Duration) *BulkJobOpsBulkImportEdgeApplicationsParams {
	return &BulkJobOpsBulkImportEdgeApplicationsParams{
		timeout: timeout,
	}
}

// NewBulkJobOpsBulkImportEdgeApplicationsParamsWithContext creates a new BulkJobOpsBulkImportEdgeApplicationsParams object
// with the ability to set a context for a request.
func NewBulkJobOpsBulkImportEdgeApplicationsParamsWithContext(ctx context.Context) *BulkJobOpsBulkImportEdgeApplicationsParams {
	return &BulkJobOpsBulkImportEdgeApplicationsParams{
		Context: ctx,
	}
}

// NewBulkJobOpsBulkImportEdgeApplicationsParamsWithHTTPClient creates a new BulkJobOpsBulkImportEdgeApplicationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewBulkJobOpsBulkImportEdgeApplicationsParamsWithHTTPClient(client *http.Client) *BulkJobOpsBulkImportEdgeApplicationsParams {
	return &BulkJobOpsBulkImportEdgeApplicationsParams{
		HTTPClient: client,
	}
}

/* BulkJobOpsBulkImportEdgeApplicationsParams contains all the parameters to send to the API endpoint
   for the bulk job ops bulk import edge applications operation.

   Typically these are written to a http.Request.
*/
type BulkJobOpsBulkImportEdgeApplicationsParams struct {

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

// WithDefaults hydrates default values in the bulk job ops bulk import edge applications params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BulkJobOpsBulkImportEdgeApplicationsParams) WithDefaults() *BulkJobOpsBulkImportEdgeApplicationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the bulk job ops bulk import edge applications params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BulkJobOpsBulkImportEdgeApplicationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the bulk job ops bulk import edge applications params
func (o *BulkJobOpsBulkImportEdgeApplicationsParams) WithTimeout(timeout time.Duration) *BulkJobOpsBulkImportEdgeApplicationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bulk job ops bulk import edge applications params
func (o *BulkJobOpsBulkImportEdgeApplicationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bulk job ops bulk import edge applications params
func (o *BulkJobOpsBulkImportEdgeApplicationsParams) WithContext(ctx context.Context) *BulkJobOpsBulkImportEdgeApplicationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bulk job ops bulk import edge applications params
func (o *BulkJobOpsBulkImportEdgeApplicationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bulk job ops bulk import edge applications params
func (o *BulkJobOpsBulkImportEdgeApplicationsParams) WithHTTPClient(client *http.Client) *BulkJobOpsBulkImportEdgeApplicationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bulk job ops bulk import edge applications params
func (o *BulkJobOpsBulkImportEdgeApplicationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the bulk job ops bulk import edge applications params
func (o *BulkJobOpsBulkImportEdgeApplicationsParams) WithXRequestID(xRequestID *string) *BulkJobOpsBulkImportEdgeApplicationsParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the bulk job ops bulk import edge applications params
func (o *BulkJobOpsBulkImportEdgeApplicationsParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the bulk job ops bulk import edge applications params
func (o *BulkJobOpsBulkImportEdgeApplicationsParams) WithBody(body *swagger_models.BulkConfig) *BulkJobOpsBulkImportEdgeApplicationsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the bulk job ops bulk import edge applications params
func (o *BulkJobOpsBulkImportEdgeApplicationsParams) SetBody(body *swagger_models.BulkConfig) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *BulkJobOpsBulkImportEdgeApplicationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
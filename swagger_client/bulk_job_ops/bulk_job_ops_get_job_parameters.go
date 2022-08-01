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
)

// NewBulkJobOpsGetJobParams creates a new BulkJobOpsGetJobParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBulkJobOpsGetJobParams() *BulkJobOpsGetJobParams {
	return &BulkJobOpsGetJobParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBulkJobOpsGetJobParamsWithTimeout creates a new BulkJobOpsGetJobParams object
// with the ability to set a timeout on a request.
func NewBulkJobOpsGetJobParamsWithTimeout(timeout time.Duration) *BulkJobOpsGetJobParams {
	return &BulkJobOpsGetJobParams{
		timeout: timeout,
	}
}

// NewBulkJobOpsGetJobParamsWithContext creates a new BulkJobOpsGetJobParams object
// with the ability to set a context for a request.
func NewBulkJobOpsGetJobParamsWithContext(ctx context.Context) *BulkJobOpsGetJobParams {
	return &BulkJobOpsGetJobParams{
		Context: ctx,
	}
}

// NewBulkJobOpsGetJobParamsWithHTTPClient creates a new BulkJobOpsGetJobParams object
// with the ability to set a custom HTTPClient for a request.
func NewBulkJobOpsGetJobParamsWithHTTPClient(client *http.Client) *BulkJobOpsGetJobParams {
	return &BulkJobOpsGetJobParams{
		HTTPClient: client,
	}
}

/* BulkJobOpsGetJobParams contains all the parameters to send to the API endpoint
   for the bulk job ops get job operation.

   Typically these are written to a http.Request.
*/
type BulkJobOpsGetJobParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* ID.

	   System defined universally unique Id of the job request
	*/
	ID string

	/* ObjectType.

	   object type for which job has been created
	*/
	ObjectType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the bulk job ops get job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BulkJobOpsGetJobParams) WithDefaults() *BulkJobOpsGetJobParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the bulk job ops get job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BulkJobOpsGetJobParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the bulk job ops get job params
func (o *BulkJobOpsGetJobParams) WithTimeout(timeout time.Duration) *BulkJobOpsGetJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bulk job ops get job params
func (o *BulkJobOpsGetJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bulk job ops get job params
func (o *BulkJobOpsGetJobParams) WithContext(ctx context.Context) *BulkJobOpsGetJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bulk job ops get job params
func (o *BulkJobOpsGetJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bulk job ops get job params
func (o *BulkJobOpsGetJobParams) WithHTTPClient(client *http.Client) *BulkJobOpsGetJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bulk job ops get job params
func (o *BulkJobOpsGetJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the bulk job ops get job params
func (o *BulkJobOpsGetJobParams) WithXRequestID(xRequestID *string) *BulkJobOpsGetJobParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the bulk job ops get job params
func (o *BulkJobOpsGetJobParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the bulk job ops get job params
func (o *BulkJobOpsGetJobParams) WithID(id string) *BulkJobOpsGetJobParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the bulk job ops get job params
func (o *BulkJobOpsGetJobParams) SetID(id string) {
	o.ID = id
}

// WithObjectType adds the objectType to the bulk job ops get job params
func (o *BulkJobOpsGetJobParams) WithObjectType(objectType string) *BulkJobOpsGetJobParams {
	o.SetObjectType(objectType)
	return o
}

// SetObjectType adds the objectType to the bulk job ops get job params
func (o *BulkJobOpsGetJobParams) SetObjectType(objectType string) {
	o.ObjectType = objectType
}

// WriteToRequest writes these params to a swagger request
func (o *BulkJobOpsGetJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param objectType
	if err := r.SetPathParam("objectType", o.ObjectType); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

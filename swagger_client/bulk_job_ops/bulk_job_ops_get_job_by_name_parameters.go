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

// NewBulkJobOpsGetJobByNameParams creates a new BulkJobOpsGetJobByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBulkJobOpsGetJobByNameParams() *BulkJobOpsGetJobByNameParams {
	return &BulkJobOpsGetJobByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBulkJobOpsGetJobByNameParamsWithTimeout creates a new BulkJobOpsGetJobByNameParams object
// with the ability to set a timeout on a request.
func NewBulkJobOpsGetJobByNameParamsWithTimeout(timeout time.Duration) *BulkJobOpsGetJobByNameParams {
	return &BulkJobOpsGetJobByNameParams{
		timeout: timeout,
	}
}

// NewBulkJobOpsGetJobByNameParamsWithContext creates a new BulkJobOpsGetJobByNameParams object
// with the ability to set a context for a request.
func NewBulkJobOpsGetJobByNameParamsWithContext(ctx context.Context) *BulkJobOpsGetJobByNameParams {
	return &BulkJobOpsGetJobByNameParams{
		Context: ctx,
	}
}

// NewBulkJobOpsGetJobByNameParamsWithHTTPClient creates a new BulkJobOpsGetJobByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewBulkJobOpsGetJobByNameParamsWithHTTPClient(client *http.Client) *BulkJobOpsGetJobByNameParams {
	return &BulkJobOpsGetJobByNameParams{
		HTTPClient: client,
	}
}

/* BulkJobOpsGetJobByNameParams contains all the parameters to send to the API endpoint
   for the bulk job ops get job by name operation.

   Typically these are written to a http.Request.
*/
type BulkJobOpsGetJobByNameParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* Name.

	   User defined name of the job request, unique across the enterprise. Once object is created, name can’t be changed
	*/
	Name string

	/* ObjectType.

	   object type for which job has been created
	*/
	ObjectType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the bulk job ops get job by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BulkJobOpsGetJobByNameParams) WithDefaults() *BulkJobOpsGetJobByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the bulk job ops get job by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BulkJobOpsGetJobByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the bulk job ops get job by name params
func (o *BulkJobOpsGetJobByNameParams) WithTimeout(timeout time.Duration) *BulkJobOpsGetJobByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bulk job ops get job by name params
func (o *BulkJobOpsGetJobByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bulk job ops get job by name params
func (o *BulkJobOpsGetJobByNameParams) WithContext(ctx context.Context) *BulkJobOpsGetJobByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bulk job ops get job by name params
func (o *BulkJobOpsGetJobByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bulk job ops get job by name params
func (o *BulkJobOpsGetJobByNameParams) WithHTTPClient(client *http.Client) *BulkJobOpsGetJobByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bulk job ops get job by name params
func (o *BulkJobOpsGetJobByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the bulk job ops get job by name params
func (o *BulkJobOpsGetJobByNameParams) WithXRequestID(xRequestID *string) *BulkJobOpsGetJobByNameParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the bulk job ops get job by name params
func (o *BulkJobOpsGetJobByNameParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithName adds the name to the bulk job ops get job by name params
func (o *BulkJobOpsGetJobByNameParams) WithName(name string) *BulkJobOpsGetJobByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the bulk job ops get job by name params
func (o *BulkJobOpsGetJobByNameParams) SetName(name string) {
	o.Name = name
}

// WithObjectType adds the objectType to the bulk job ops get job by name params
func (o *BulkJobOpsGetJobByNameParams) WithObjectType(objectType string) *BulkJobOpsGetJobByNameParams {
	o.SetObjectType(objectType)
	return o
}

// SetObjectType adds the objectType to the bulk job ops get job by name params
func (o *BulkJobOpsGetJobByNameParams) SetObjectType(objectType string) {
	o.ObjectType = objectType
}

// WriteToRequest writes these params to a swagger request
func (o *BulkJobOpsGetJobByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param objectType
	if err := r.SetPathParam("objectType", o.ObjectType); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

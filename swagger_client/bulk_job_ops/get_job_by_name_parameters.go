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

// NewGetJobByNameParams creates a new GetJobByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetJobByNameParams() *GetJobByNameParams {
	return &GetJobByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetJobByNameParamsWithTimeout creates a new GetJobByNameParams object
// with the ability to set a timeout on a request.
func NewGetJobByNameParamsWithTimeout(timeout time.Duration) *GetJobByNameParams {
	return &GetJobByNameParams{
		timeout: timeout,
	}
}

// NewGetJobByNameParamsWithContext creates a new GetJobByNameParams object
// with the ability to set a context for a request.
func NewGetJobByNameParamsWithContext(ctx context.Context) *GetJobByNameParams {
	return &GetJobByNameParams{
		Context: ctx,
	}
}

// NewGetJobByNameParamsWithHTTPClient creates a new GetJobByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetJobByNameParamsWithHTTPClient(client *http.Client) *GetJobByNameParams {
	return &GetJobByNameParams{
		HTTPClient: client,
	}
}

/* GetJobByNameParams contains all the parameters to send to the API endpoint
   for the get job by name operation.

   Typically these are written to a http.Request.
*/
type GetJobByNameParams struct {

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

// WithDefaults hydrates default values in the get job by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetJobByNameParams) WithDefaults() *GetJobByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get job by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetJobByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get job by name params
func (o *GetJobByNameParams) WithTimeout(timeout time.Duration) *GetJobByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get job by name params
func (o *GetJobByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get job by name params
func (o *GetJobByNameParams) WithContext(ctx context.Context) *GetJobByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get job by name params
func (o *GetJobByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get job by name params
func (o *GetJobByNameParams) WithHTTPClient(client *http.Client) *GetJobByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get job by name params
func (o *GetJobByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get job by name params
func (o *GetJobByNameParams) WithXRequestID(xRequestID *string) *GetJobByNameParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get job by name params
func (o *GetJobByNameParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithName adds the name to the get job by name params
func (o *GetJobByNameParams) WithName(name string) *GetJobByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get job by name params
func (o *GetJobByNameParams) SetName(name string) {
	o.Name = name
}

// WithObjectType adds the objectType to the get job by name params
func (o *GetJobByNameParams) WithObjectType(objectType string) *GetJobByNameParams {
	o.SetObjectType(objectType)
	return o
}

// SetObjectType adds the objectType to the get job by name params
func (o *GetJobByNameParams) SetObjectType(objectType string) {
	o.ObjectType = objectType
}

// WriteToRequest writes these params to a swagger request
func (o *GetJobByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

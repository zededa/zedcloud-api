// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package datastore_configuration

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

// NewDeleteDatastoreParams creates a new DeleteDatastoreParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteDatastoreParams() *DeleteDatastoreParams {
	return &DeleteDatastoreParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDatastoreParamsWithTimeout creates a new DeleteDatastoreParams object
// with the ability to set a timeout on a request.
func NewDeleteDatastoreParamsWithTimeout(timeout time.Duration) *DeleteDatastoreParams {
	return &DeleteDatastoreParams{
		timeout: timeout,
	}
}

// NewDeleteDatastoreParamsWithContext creates a new DeleteDatastoreParams object
// with the ability to set a context for a request.
func NewDeleteDatastoreParamsWithContext(ctx context.Context) *DeleteDatastoreParams {
	return &DeleteDatastoreParams{
		Context: ctx,
	}
}

// NewDeleteDatastoreParamsWithHTTPClient creates a new DeleteDatastoreParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteDatastoreParamsWithHTTPClient(client *http.Client) *DeleteDatastoreParams {
	return &DeleteDatastoreParams{
		HTTPClient: client,
	}
}

/* DeleteDatastoreParams contains all the parameters to send to the API endpoint
   for the delete datastore operation.

   Typically these are written to a http.Request.
*/
type DeleteDatastoreParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* ID.

	   System defined universally unique Id of the datastore
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete datastore params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDatastoreParams) WithDefaults() *DeleteDatastoreParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete datastore params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDatastoreParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete datastore params
func (o *DeleteDatastoreParams) WithTimeout(timeout time.Duration) *DeleteDatastoreParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete datastore params
func (o *DeleteDatastoreParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete datastore params
func (o *DeleteDatastoreParams) WithContext(ctx context.Context) *DeleteDatastoreParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete datastore params
func (o *DeleteDatastoreParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete datastore params
func (o *DeleteDatastoreParams) WithHTTPClient(client *http.Client) *DeleteDatastoreParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete datastore params
func (o *DeleteDatastoreParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the delete datastore params
func (o *DeleteDatastoreParams) WithXRequestID(xRequestID *string) *DeleteDatastoreParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the delete datastore params
func (o *DeleteDatastoreParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the delete datastore params
func (o *DeleteDatastoreParams) WithID(id string) *DeleteDatastoreParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete datastore params
func (o *DeleteDatastoreParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDatastoreParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

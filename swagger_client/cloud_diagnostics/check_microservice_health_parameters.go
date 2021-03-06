// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package cloud_diagnostics

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

// NewCheckMicroserviceHealthParams creates a new CheckMicroserviceHealthParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCheckMicroserviceHealthParams() *CheckMicroserviceHealthParams {
	return &CheckMicroserviceHealthParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCheckMicroserviceHealthParamsWithTimeout creates a new CheckMicroserviceHealthParams object
// with the ability to set a timeout on a request.
func NewCheckMicroserviceHealthParamsWithTimeout(timeout time.Duration) *CheckMicroserviceHealthParams {
	return &CheckMicroserviceHealthParams{
		timeout: timeout,
	}
}

// NewCheckMicroserviceHealthParamsWithContext creates a new CheckMicroserviceHealthParams object
// with the ability to set a context for a request.
func NewCheckMicroserviceHealthParamsWithContext(ctx context.Context) *CheckMicroserviceHealthParams {
	return &CheckMicroserviceHealthParams{
		Context: ctx,
	}
}

// NewCheckMicroserviceHealthParamsWithHTTPClient creates a new CheckMicroserviceHealthParams object
// with the ability to set a custom HTTPClient for a request.
func NewCheckMicroserviceHealthParamsWithHTTPClient(client *http.Client) *CheckMicroserviceHealthParams {
	return &CheckMicroserviceHealthParams{
		HTTPClient: client,
	}
}

/* CheckMicroserviceHealthParams contains all the parameters to send to the API endpoint
   for the check microservice health operation.

   Typically these are written to a http.Request.
*/
type CheckMicroserviceHealthParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body *swagger_models.HelloName

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the check microservice health params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CheckMicroserviceHealthParams) WithDefaults() *CheckMicroserviceHealthParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the check microservice health params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CheckMicroserviceHealthParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the check microservice health params
func (o *CheckMicroserviceHealthParams) WithTimeout(timeout time.Duration) *CheckMicroserviceHealthParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the check microservice health params
func (o *CheckMicroserviceHealthParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the check microservice health params
func (o *CheckMicroserviceHealthParams) WithContext(ctx context.Context) *CheckMicroserviceHealthParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the check microservice health params
func (o *CheckMicroserviceHealthParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the check microservice health params
func (o *CheckMicroserviceHealthParams) WithHTTPClient(client *http.Client) *CheckMicroserviceHealthParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the check microservice health params
func (o *CheckMicroserviceHealthParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the check microservice health params
func (o *CheckMicroserviceHealthParams) WithXRequestID(xRequestID *string) *CheckMicroserviceHealthParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the check microservice health params
func (o *CheckMicroserviceHealthParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the check microservice health params
func (o *CheckMicroserviceHealthParams) WithBody(body *swagger_models.HelloName) *CheckMicroserviceHealthParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the check microservice health params
func (o *CheckMicroserviceHealthParams) SetBody(body *swagger_models.HelloName) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CheckMicroserviceHealthParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

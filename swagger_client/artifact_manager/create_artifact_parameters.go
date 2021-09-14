// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package artifact_manager

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

// NewCreateArtifactParams creates a new CreateArtifactParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateArtifactParams() *CreateArtifactParams {
	return &CreateArtifactParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateArtifactParamsWithTimeout creates a new CreateArtifactParams object
// with the ability to set a timeout on a request.
func NewCreateArtifactParamsWithTimeout(timeout time.Duration) *CreateArtifactParams {
	return &CreateArtifactParams{
		timeout: timeout,
	}
}

// NewCreateArtifactParamsWithContext creates a new CreateArtifactParams object
// with the ability to set a context for a request.
func NewCreateArtifactParamsWithContext(ctx context.Context) *CreateArtifactParams {
	return &CreateArtifactParams{
		Context: ctx,
	}
}

// NewCreateArtifactParamsWithHTTPClient creates a new CreateArtifactParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateArtifactParamsWithHTTPClient(client *http.Client) *CreateArtifactParams {
	return &CreateArtifactParams{
		HTTPClient: client,
	}
}

/* CreateArtifactParams contains all the parameters to send to the API endpoint
   for the create artifact operation.

   Typically these are written to a http.Request.
*/
type CreateArtifactParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body *swagger_models.Artifact

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create artifact params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateArtifactParams) WithDefaults() *CreateArtifactParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create artifact params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateArtifactParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create artifact params
func (o *CreateArtifactParams) WithTimeout(timeout time.Duration) *CreateArtifactParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create artifact params
func (o *CreateArtifactParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create artifact params
func (o *CreateArtifactParams) WithContext(ctx context.Context) *CreateArtifactParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create artifact params
func (o *CreateArtifactParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create artifact params
func (o *CreateArtifactParams) WithHTTPClient(client *http.Client) *CreateArtifactParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create artifact params
func (o *CreateArtifactParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the create artifact params
func (o *CreateArtifactParams) WithXRequestID(xRequestID *string) *CreateArtifactParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the create artifact params
func (o *CreateArtifactParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the create artifact params
func (o *CreateArtifactParams) WithBody(body *swagger_models.Artifact) *CreateArtifactParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create artifact params
func (o *CreateArtifactParams) SetBody(body *swagger_models.Artifact) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateArtifactParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

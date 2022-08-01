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
)

// NewArtifactManagerUploadArtifactParams creates a new ArtifactManagerUploadArtifactParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewArtifactManagerUploadArtifactParams() *ArtifactManagerUploadArtifactParams {
	return &ArtifactManagerUploadArtifactParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewArtifactManagerUploadArtifactParamsWithTimeout creates a new ArtifactManagerUploadArtifactParams object
// with the ability to set a timeout on a request.
func NewArtifactManagerUploadArtifactParamsWithTimeout(timeout time.Duration) *ArtifactManagerUploadArtifactParams {
	return &ArtifactManagerUploadArtifactParams{
		timeout: timeout,
	}
}

// NewArtifactManagerUploadArtifactParamsWithContext creates a new ArtifactManagerUploadArtifactParams object
// with the ability to set a context for a request.
func NewArtifactManagerUploadArtifactParamsWithContext(ctx context.Context) *ArtifactManagerUploadArtifactParams {
	return &ArtifactManagerUploadArtifactParams{
		Context: ctx,
	}
}

// NewArtifactManagerUploadArtifactParamsWithHTTPClient creates a new ArtifactManagerUploadArtifactParams object
// with the ability to set a custom HTTPClient for a request.
func NewArtifactManagerUploadArtifactParamsWithHTTPClient(client *http.Client) *ArtifactManagerUploadArtifactParams {
	return &ArtifactManagerUploadArtifactParams{
		HTTPClient: client,
	}
}

/* ArtifactManagerUploadArtifactParams contains all the parameters to send to the API endpoint
   for the artifact manager upload artifact operation.

   Typically these are written to a http.Request.
*/
type ArtifactManagerUploadArtifactParams struct {

	/* ContentRange.

	   Content range in this request. Example : bytes 1024-2047/8192
	*/
	ContentRange string

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* Body.

	   artifact chunk data

	   Format: byte
	*/
	Body strfmt.Base64

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the artifact manager upload artifact params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ArtifactManagerUploadArtifactParams) WithDefaults() *ArtifactManagerUploadArtifactParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the artifact manager upload artifact params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ArtifactManagerUploadArtifactParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the artifact manager upload artifact params
func (o *ArtifactManagerUploadArtifactParams) WithTimeout(timeout time.Duration) *ArtifactManagerUploadArtifactParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the artifact manager upload artifact params
func (o *ArtifactManagerUploadArtifactParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the artifact manager upload artifact params
func (o *ArtifactManagerUploadArtifactParams) WithContext(ctx context.Context) *ArtifactManagerUploadArtifactParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the artifact manager upload artifact params
func (o *ArtifactManagerUploadArtifactParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the artifact manager upload artifact params
func (o *ArtifactManagerUploadArtifactParams) WithHTTPClient(client *http.Client) *ArtifactManagerUploadArtifactParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the artifact manager upload artifact params
func (o *ArtifactManagerUploadArtifactParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentRange adds the contentRange to the artifact manager upload artifact params
func (o *ArtifactManagerUploadArtifactParams) WithContentRange(contentRange string) *ArtifactManagerUploadArtifactParams {
	o.SetContentRange(contentRange)
	return o
}

// SetContentRange adds the contentRange to the artifact manager upload artifact params
func (o *ArtifactManagerUploadArtifactParams) SetContentRange(contentRange string) {
	o.ContentRange = contentRange
}

// WithXRequestID adds the xRequestID to the artifact manager upload artifact params
func (o *ArtifactManagerUploadArtifactParams) WithXRequestID(xRequestID *string) *ArtifactManagerUploadArtifactParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the artifact manager upload artifact params
func (o *ArtifactManagerUploadArtifactParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the artifact manager upload artifact params
func (o *ArtifactManagerUploadArtifactParams) WithBody(body strfmt.Base64) *ArtifactManagerUploadArtifactParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the artifact manager upload artifact params
func (o *ArtifactManagerUploadArtifactParams) SetBody(body strfmt.Base64) {
	o.Body = body
}

// WithID adds the id to the artifact manager upload artifact params
func (o *ArtifactManagerUploadArtifactParams) WithID(id string) *ArtifactManagerUploadArtifactParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the artifact manager upload artifact params
func (o *ArtifactManagerUploadArtifactParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ArtifactManagerUploadArtifactParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Content-Range
	if err := r.SetHeaderParam("Content-Range", o.ContentRange); err != nil {
		return err
	}

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}
	}
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
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

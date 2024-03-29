// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package identity_access_management

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

// NewIdentityAccessManagementUpdateDocPolicyLatestParams creates a new IdentityAccessManagementUpdateDocPolicyLatestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIdentityAccessManagementUpdateDocPolicyLatestParams() *IdentityAccessManagementUpdateDocPolicyLatestParams {
	return &IdentityAccessManagementUpdateDocPolicyLatestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIdentityAccessManagementUpdateDocPolicyLatestParamsWithTimeout creates a new IdentityAccessManagementUpdateDocPolicyLatestParams object
// with the ability to set a timeout on a request.
func NewIdentityAccessManagementUpdateDocPolicyLatestParamsWithTimeout(timeout time.Duration) *IdentityAccessManagementUpdateDocPolicyLatestParams {
	return &IdentityAccessManagementUpdateDocPolicyLatestParams{
		timeout: timeout,
	}
}

// NewIdentityAccessManagementUpdateDocPolicyLatestParamsWithContext creates a new IdentityAccessManagementUpdateDocPolicyLatestParams object
// with the ability to set a context for a request.
func NewIdentityAccessManagementUpdateDocPolicyLatestParamsWithContext(ctx context.Context) *IdentityAccessManagementUpdateDocPolicyLatestParams {
	return &IdentityAccessManagementUpdateDocPolicyLatestParams{
		Context: ctx,
	}
}

// NewIdentityAccessManagementUpdateDocPolicyLatestParamsWithHTTPClient creates a new IdentityAccessManagementUpdateDocPolicyLatestParams object
// with the ability to set a custom HTTPClient for a request.
func NewIdentityAccessManagementUpdateDocPolicyLatestParamsWithHTTPClient(client *http.Client) *IdentityAccessManagementUpdateDocPolicyLatestParams {
	return &IdentityAccessManagementUpdateDocPolicyLatestParams{
		HTTPClient: client,
	}
}

/*
IdentityAccessManagementUpdateDocPolicyLatestParams contains all the parameters to send to the API endpoint

	for the identity access management update doc policy latest operation.

	Typically these are written to a http.Request.
*/
type IdentityAccessManagementUpdateDocPolicyLatestParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body *swagger_models.DocPolicy

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the identity access management update doc policy latest params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IdentityAccessManagementUpdateDocPolicyLatestParams) WithDefaults() *IdentityAccessManagementUpdateDocPolicyLatestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the identity access management update doc policy latest params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IdentityAccessManagementUpdateDocPolicyLatestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the identity access management update doc policy latest params
func (o *IdentityAccessManagementUpdateDocPolicyLatestParams) WithTimeout(timeout time.Duration) *IdentityAccessManagementUpdateDocPolicyLatestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the identity access management update doc policy latest params
func (o *IdentityAccessManagementUpdateDocPolicyLatestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the identity access management update doc policy latest params
func (o *IdentityAccessManagementUpdateDocPolicyLatestParams) WithContext(ctx context.Context) *IdentityAccessManagementUpdateDocPolicyLatestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the identity access management update doc policy latest params
func (o *IdentityAccessManagementUpdateDocPolicyLatestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the identity access management update doc policy latest params
func (o *IdentityAccessManagementUpdateDocPolicyLatestParams) WithHTTPClient(client *http.Client) *IdentityAccessManagementUpdateDocPolicyLatestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the identity access management update doc policy latest params
func (o *IdentityAccessManagementUpdateDocPolicyLatestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the identity access management update doc policy latest params
func (o *IdentityAccessManagementUpdateDocPolicyLatestParams) WithXRequestID(xRequestID *string) *IdentityAccessManagementUpdateDocPolicyLatestParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the identity access management update doc policy latest params
func (o *IdentityAccessManagementUpdateDocPolicyLatestParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the identity access management update doc policy latest params
func (o *IdentityAccessManagementUpdateDocPolicyLatestParams) WithBody(body *swagger_models.DocPolicy) *IdentityAccessManagementUpdateDocPolicyLatestParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the identity access management update doc policy latest params
func (o *IdentityAccessManagementUpdateDocPolicyLatestParams) SetBody(body *swagger_models.DocPolicy) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *IdentityAccessManagementUpdateDocPolicyLatestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

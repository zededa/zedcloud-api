// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package resource_group

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

// NewResourceGroupGetResourceGroupEventsParams creates a new ResourceGroupGetResourceGroupEventsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewResourceGroupGetResourceGroupEventsParams() *ResourceGroupGetResourceGroupEventsParams {
	return &ResourceGroupGetResourceGroupEventsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewResourceGroupGetResourceGroupEventsParamsWithTimeout creates a new ResourceGroupGetResourceGroupEventsParams object
// with the ability to set a timeout on a request.
func NewResourceGroupGetResourceGroupEventsParamsWithTimeout(timeout time.Duration) *ResourceGroupGetResourceGroupEventsParams {
	return &ResourceGroupGetResourceGroupEventsParams{
		timeout: timeout,
	}
}

// NewResourceGroupGetResourceGroupEventsParamsWithContext creates a new ResourceGroupGetResourceGroupEventsParams object
// with the ability to set a context for a request.
func NewResourceGroupGetResourceGroupEventsParamsWithContext(ctx context.Context) *ResourceGroupGetResourceGroupEventsParams {
	return &ResourceGroupGetResourceGroupEventsParams{
		Context: ctx,
	}
}

// NewResourceGroupGetResourceGroupEventsParamsWithHTTPClient creates a new ResourceGroupGetResourceGroupEventsParams object
// with the ability to set a custom HTTPClient for a request.
func NewResourceGroupGetResourceGroupEventsParamsWithHTTPClient(client *http.Client) *ResourceGroupGetResourceGroupEventsParams {
	return &ResourceGroupGetResourceGroupEventsParams{
		HTTPClient: client,
	}
}

/*
ResourceGroupGetResourceGroupEventsParams contains all the parameters to send to the API endpoint

	for the resource group get resource group events operation.

	Typically these are written to a http.Request.
*/
type ResourceGroupGetResourceGroupEventsParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* EnterpriseID.

	   system generated unique id for an enterprise (deprecated)
	*/
	EnterpriseID *string

	/* Objid.

	   Object id
	*/
	Objid string

	/* Objname.

	   Object name
	*/
	Objname *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the resource group get resource group events params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResourceGroupGetResourceGroupEventsParams) WithDefaults() *ResourceGroupGetResourceGroupEventsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the resource group get resource group events params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResourceGroupGetResourceGroupEventsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the resource group get resource group events params
func (o *ResourceGroupGetResourceGroupEventsParams) WithTimeout(timeout time.Duration) *ResourceGroupGetResourceGroupEventsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the resource group get resource group events params
func (o *ResourceGroupGetResourceGroupEventsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the resource group get resource group events params
func (o *ResourceGroupGetResourceGroupEventsParams) WithContext(ctx context.Context) *ResourceGroupGetResourceGroupEventsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the resource group get resource group events params
func (o *ResourceGroupGetResourceGroupEventsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the resource group get resource group events params
func (o *ResourceGroupGetResourceGroupEventsParams) WithHTTPClient(client *http.Client) *ResourceGroupGetResourceGroupEventsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the resource group get resource group events params
func (o *ResourceGroupGetResourceGroupEventsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the resource group get resource group events params
func (o *ResourceGroupGetResourceGroupEventsParams) WithXRequestID(xRequestID *string) *ResourceGroupGetResourceGroupEventsParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the resource group get resource group events params
func (o *ResourceGroupGetResourceGroupEventsParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithEnterpriseID adds the enterpriseID to the resource group get resource group events params
func (o *ResourceGroupGetResourceGroupEventsParams) WithEnterpriseID(enterpriseID *string) *ResourceGroupGetResourceGroupEventsParams {
	o.SetEnterpriseID(enterpriseID)
	return o
}

// SetEnterpriseID adds the enterpriseId to the resource group get resource group events params
func (o *ResourceGroupGetResourceGroupEventsParams) SetEnterpriseID(enterpriseID *string) {
	o.EnterpriseID = enterpriseID
}

// WithObjid adds the objid to the resource group get resource group events params
func (o *ResourceGroupGetResourceGroupEventsParams) WithObjid(objid string) *ResourceGroupGetResourceGroupEventsParams {
	o.SetObjid(objid)
	return o
}

// SetObjid adds the objid to the resource group get resource group events params
func (o *ResourceGroupGetResourceGroupEventsParams) SetObjid(objid string) {
	o.Objid = objid
}

// WithObjname adds the objname to the resource group get resource group events params
func (o *ResourceGroupGetResourceGroupEventsParams) WithObjname(objname *string) *ResourceGroupGetResourceGroupEventsParams {
	o.SetObjname(objname)
	return o
}

// SetObjname adds the objname to the resource group get resource group events params
func (o *ResourceGroupGetResourceGroupEventsParams) SetObjname(objname *string) {
	o.Objname = objname
}

// WriteToRequest writes these params to a swagger request
func (o *ResourceGroupGetResourceGroupEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.EnterpriseID != nil {

		// query param enterpriseId
		var qrEnterpriseID string

		if o.EnterpriseID != nil {
			qrEnterpriseID = *o.EnterpriseID
		}
		qEnterpriseID := qrEnterpriseID
		if qEnterpriseID != "" {

			if err := r.SetQueryParam("enterpriseId", qEnterpriseID); err != nil {
				return err
			}
		}
	}

	// path param objid
	if err := r.SetPathParam("objid", o.Objid); err != nil {
		return err
	}

	if o.Objname != nil {

		// query param objname
		var qrObjname string

		if o.Objname != nil {
			qrObjname = *o.Objname
		}
		qObjname := qrObjname
		if qObjname != "" {

			if err := r.SetQueryParam("objname", qObjname); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

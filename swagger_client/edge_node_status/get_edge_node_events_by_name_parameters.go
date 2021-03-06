// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package edge_node_status

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

// NewGetEdgeNodeEventsByNameParams creates a new GetEdgeNodeEventsByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetEdgeNodeEventsByNameParams() *GetEdgeNodeEventsByNameParams {
	return &GetEdgeNodeEventsByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetEdgeNodeEventsByNameParamsWithTimeout creates a new GetEdgeNodeEventsByNameParams object
// with the ability to set a timeout on a request.
func NewGetEdgeNodeEventsByNameParamsWithTimeout(timeout time.Duration) *GetEdgeNodeEventsByNameParams {
	return &GetEdgeNodeEventsByNameParams{
		timeout: timeout,
	}
}

// NewGetEdgeNodeEventsByNameParamsWithContext creates a new GetEdgeNodeEventsByNameParams object
// with the ability to set a context for a request.
func NewGetEdgeNodeEventsByNameParamsWithContext(ctx context.Context) *GetEdgeNodeEventsByNameParams {
	return &GetEdgeNodeEventsByNameParams{
		Context: ctx,
	}
}

// NewGetEdgeNodeEventsByNameParamsWithHTTPClient creates a new GetEdgeNodeEventsByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetEdgeNodeEventsByNameParamsWithHTTPClient(client *http.Client) *GetEdgeNodeEventsByNameParams {
	return &GetEdgeNodeEventsByNameParams{
		HTTPClient: client,
	}
}

/* GetEdgeNodeEventsByNameParams contains all the parameters to send to the API endpoint
   for the get edge node events by name operation.

   Typically these are written to a http.Request.
*/
type GetEdgeNodeEventsByNameParams struct {

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
	Objid *string

	/* Objname.

	   Object name
	*/
	Objname string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get edge node events by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEdgeNodeEventsByNameParams) WithDefaults() *GetEdgeNodeEventsByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get edge node events by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEdgeNodeEventsByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get edge node events by name params
func (o *GetEdgeNodeEventsByNameParams) WithTimeout(timeout time.Duration) *GetEdgeNodeEventsByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get edge node events by name params
func (o *GetEdgeNodeEventsByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get edge node events by name params
func (o *GetEdgeNodeEventsByNameParams) WithContext(ctx context.Context) *GetEdgeNodeEventsByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get edge node events by name params
func (o *GetEdgeNodeEventsByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get edge node events by name params
func (o *GetEdgeNodeEventsByNameParams) WithHTTPClient(client *http.Client) *GetEdgeNodeEventsByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get edge node events by name params
func (o *GetEdgeNodeEventsByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get edge node events by name params
func (o *GetEdgeNodeEventsByNameParams) WithXRequestID(xRequestID *string) *GetEdgeNodeEventsByNameParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get edge node events by name params
func (o *GetEdgeNodeEventsByNameParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithEnterpriseID adds the enterpriseID to the get edge node events by name params
func (o *GetEdgeNodeEventsByNameParams) WithEnterpriseID(enterpriseID *string) *GetEdgeNodeEventsByNameParams {
	o.SetEnterpriseID(enterpriseID)
	return o
}

// SetEnterpriseID adds the enterpriseId to the get edge node events by name params
func (o *GetEdgeNodeEventsByNameParams) SetEnterpriseID(enterpriseID *string) {
	o.EnterpriseID = enterpriseID
}

// WithObjid adds the objid to the get edge node events by name params
func (o *GetEdgeNodeEventsByNameParams) WithObjid(objid *string) *GetEdgeNodeEventsByNameParams {
	o.SetObjid(objid)
	return o
}

// SetObjid adds the objid to the get edge node events by name params
func (o *GetEdgeNodeEventsByNameParams) SetObjid(objid *string) {
	o.Objid = objid
}

// WithObjname adds the objname to the get edge node events by name params
func (o *GetEdgeNodeEventsByNameParams) WithObjname(objname string) *GetEdgeNodeEventsByNameParams {
	o.SetObjname(objname)
	return o
}

// SetObjname adds the objname to the get edge node events by name params
func (o *GetEdgeNodeEventsByNameParams) SetObjname(objname string) {
	o.Objname = objname
}

// WriteToRequest writes these params to a swagger request
func (o *GetEdgeNodeEventsByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Objid != nil {

		// query param objid
		var qrObjid string

		if o.Objid != nil {
			qrObjid = *o.Objid
		}
		qObjid := qrObjid
		if qObjid != "" {

			if err := r.SetQueryParam("objid", qObjid); err != nil {
				return err
			}
		}
	}

	// path param objname
	if err := r.SetPathParam("objname", o.Objname); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

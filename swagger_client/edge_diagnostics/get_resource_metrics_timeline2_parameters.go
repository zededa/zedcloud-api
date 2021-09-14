// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package edge_diagnostics

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

// NewGetResourceMetricsTimeline2Params creates a new GetResourceMetricsTimeline2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetResourceMetricsTimeline2Params() *GetResourceMetricsTimeline2Params {
	return &GetResourceMetricsTimeline2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetResourceMetricsTimeline2ParamsWithTimeout creates a new GetResourceMetricsTimeline2Params object
// with the ability to set a timeout on a request.
func NewGetResourceMetricsTimeline2ParamsWithTimeout(timeout time.Duration) *GetResourceMetricsTimeline2Params {
	return &GetResourceMetricsTimeline2Params{
		timeout: timeout,
	}
}

// NewGetResourceMetricsTimeline2ParamsWithContext creates a new GetResourceMetricsTimeline2Params object
// with the ability to set a context for a request.
func NewGetResourceMetricsTimeline2ParamsWithContext(ctx context.Context) *GetResourceMetricsTimeline2Params {
	return &GetResourceMetricsTimeline2Params{
		Context: ctx,
	}
}

// NewGetResourceMetricsTimeline2ParamsWithHTTPClient creates a new GetResourceMetricsTimeline2Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetResourceMetricsTimeline2ParamsWithHTTPClient(client *http.Client) *GetResourceMetricsTimeline2Params {
	return &GetResourceMetricsTimeline2Params{
		HTTPClient: client,
	}
}

/* GetResourceMetricsTimeline2Params contains all the parameters to send to the API endpoint
   for the get resource metrics timeline2 operation.

   Typically these are written to a http.Request.
*/
type GetResourceMetricsTimeline2Params struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// EndTime.
	//
	// Format: date-time
	EndTime *strfmt.DateTime

	// EnterpriseID.
	EnterpriseID *string

	// Interval.
	//
	// Format: date-time
	Interval *strfmt.DateTime

	// MType.
	MType string

	// Objid.
	Objid *string

	// Objname.
	Objname *string

	// Objtype.
	//
	// Default: "OBJECT_TYPE_UNSPECIFIED"
	Objtype *string

	// StartTime.
	//
	// Format: date-time
	StartTime *strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get resource metrics timeline2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetResourceMetricsTimeline2Params) WithDefaults() *GetResourceMetricsTimeline2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get resource metrics timeline2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetResourceMetricsTimeline2Params) SetDefaults() {
	var (
		objtypeDefault = string("OBJECT_TYPE_UNSPECIFIED")
	)

	val := GetResourceMetricsTimeline2Params{
		Objtype: &objtypeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get resource metrics timeline2 params
func (o *GetResourceMetricsTimeline2Params) WithTimeout(timeout time.Duration) *GetResourceMetricsTimeline2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get resource metrics timeline2 params
func (o *GetResourceMetricsTimeline2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get resource metrics timeline2 params
func (o *GetResourceMetricsTimeline2Params) WithContext(ctx context.Context) *GetResourceMetricsTimeline2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get resource metrics timeline2 params
func (o *GetResourceMetricsTimeline2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get resource metrics timeline2 params
func (o *GetResourceMetricsTimeline2Params) WithHTTPClient(client *http.Client) *GetResourceMetricsTimeline2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get resource metrics timeline2 params
func (o *GetResourceMetricsTimeline2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get resource metrics timeline2 params
func (o *GetResourceMetricsTimeline2Params) WithXRequestID(xRequestID *string) *GetResourceMetricsTimeline2Params {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get resource metrics timeline2 params
func (o *GetResourceMetricsTimeline2Params) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithEndTime adds the endTime to the get resource metrics timeline2 params
func (o *GetResourceMetricsTimeline2Params) WithEndTime(endTime *strfmt.DateTime) *GetResourceMetricsTimeline2Params {
	o.SetEndTime(endTime)
	return o
}

// SetEndTime adds the endTime to the get resource metrics timeline2 params
func (o *GetResourceMetricsTimeline2Params) SetEndTime(endTime *strfmt.DateTime) {
	o.EndTime = endTime
}

// WithEnterpriseID adds the enterpriseID to the get resource metrics timeline2 params
func (o *GetResourceMetricsTimeline2Params) WithEnterpriseID(enterpriseID *string) *GetResourceMetricsTimeline2Params {
	o.SetEnterpriseID(enterpriseID)
	return o
}

// SetEnterpriseID adds the enterpriseId to the get resource metrics timeline2 params
func (o *GetResourceMetricsTimeline2Params) SetEnterpriseID(enterpriseID *string) {
	o.EnterpriseID = enterpriseID
}

// WithInterval adds the interval to the get resource metrics timeline2 params
func (o *GetResourceMetricsTimeline2Params) WithInterval(interval *strfmt.DateTime) *GetResourceMetricsTimeline2Params {
	o.SetInterval(interval)
	return o
}

// SetInterval adds the interval to the get resource metrics timeline2 params
func (o *GetResourceMetricsTimeline2Params) SetInterval(interval *strfmt.DateTime) {
	o.Interval = interval
}

// WithMType adds the mType to the get resource metrics timeline2 params
func (o *GetResourceMetricsTimeline2Params) WithMType(mType string) *GetResourceMetricsTimeline2Params {
	o.SetMType(mType)
	return o
}

// SetMType adds the mType to the get resource metrics timeline2 params
func (o *GetResourceMetricsTimeline2Params) SetMType(mType string) {
	o.MType = mType
}

// WithObjid adds the objid to the get resource metrics timeline2 params
func (o *GetResourceMetricsTimeline2Params) WithObjid(objid *string) *GetResourceMetricsTimeline2Params {
	o.SetObjid(objid)
	return o
}

// SetObjid adds the objid to the get resource metrics timeline2 params
func (o *GetResourceMetricsTimeline2Params) SetObjid(objid *string) {
	o.Objid = objid
}

// WithObjname adds the objname to the get resource metrics timeline2 params
func (o *GetResourceMetricsTimeline2Params) WithObjname(objname *string) *GetResourceMetricsTimeline2Params {
	o.SetObjname(objname)
	return o
}

// SetObjname adds the objname to the get resource metrics timeline2 params
func (o *GetResourceMetricsTimeline2Params) SetObjname(objname *string) {
	o.Objname = objname
}

// WithObjtype adds the objtype to the get resource metrics timeline2 params
func (o *GetResourceMetricsTimeline2Params) WithObjtype(objtype *string) *GetResourceMetricsTimeline2Params {
	o.SetObjtype(objtype)
	return o
}

// SetObjtype adds the objtype to the get resource metrics timeline2 params
func (o *GetResourceMetricsTimeline2Params) SetObjtype(objtype *string) {
	o.Objtype = objtype
}

// WithStartTime adds the startTime to the get resource metrics timeline2 params
func (o *GetResourceMetricsTimeline2Params) WithStartTime(startTime *strfmt.DateTime) *GetResourceMetricsTimeline2Params {
	o.SetStartTime(startTime)
	return o
}

// SetStartTime adds the startTime to the get resource metrics timeline2 params
func (o *GetResourceMetricsTimeline2Params) SetStartTime(startTime *strfmt.DateTime) {
	o.StartTime = startTime
}

// WriteToRequest writes these params to a swagger request
func (o *GetResourceMetricsTimeline2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.EndTime != nil {

		// query param endTime
		var qrEndTime strfmt.DateTime

		if o.EndTime != nil {
			qrEndTime = *o.EndTime
		}
		qEndTime := qrEndTime.String()
		if qEndTime != "" {

			if err := r.SetQueryParam("endTime", qEndTime); err != nil {
				return err
			}
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

	if o.Interval != nil {

		// query param interval
		var qrInterval strfmt.DateTime

		if o.Interval != nil {
			qrInterval = *o.Interval
		}
		qInterval := qrInterval.String()
		if qInterval != "" {

			if err := r.SetQueryParam("interval", qInterval); err != nil {
				return err
			}
		}
	}

	// path param mType
	if err := r.SetPathParam("mType", o.MType); err != nil {
		return err
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

	if o.Objtype != nil {

		// query param objtype
		var qrObjtype string

		if o.Objtype != nil {
			qrObjtype = *o.Objtype
		}
		qObjtype := qrObjtype
		if qObjtype != "" {

			if err := r.SetQueryParam("objtype", qObjtype); err != nil {
				return err
			}
		}
	}

	if o.StartTime != nil {

		// query param startTime
		var qrStartTime strfmt.DateTime

		if o.StartTime != nil {
			qrStartTime = *o.StartTime
		}
		qStartTime := qrStartTime.String()
		if qStartTime != "" {

			if err := r.SetQueryParam("startTime", qStartTime); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

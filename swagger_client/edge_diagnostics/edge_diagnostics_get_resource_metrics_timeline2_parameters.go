// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
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

// NewEdgeDiagnosticsGetResourceMetricsTimeline2Params creates a new EdgeDiagnosticsGetResourceMetricsTimeline2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEdgeDiagnosticsGetResourceMetricsTimeline2Params() *EdgeDiagnosticsGetResourceMetricsTimeline2Params {
	return &EdgeDiagnosticsGetResourceMetricsTimeline2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeDiagnosticsGetResourceMetricsTimeline2ParamsWithTimeout creates a new EdgeDiagnosticsGetResourceMetricsTimeline2Params object
// with the ability to set a timeout on a request.
func NewEdgeDiagnosticsGetResourceMetricsTimeline2ParamsWithTimeout(timeout time.Duration) *EdgeDiagnosticsGetResourceMetricsTimeline2Params {
	return &EdgeDiagnosticsGetResourceMetricsTimeline2Params{
		timeout: timeout,
	}
}

// NewEdgeDiagnosticsGetResourceMetricsTimeline2ParamsWithContext creates a new EdgeDiagnosticsGetResourceMetricsTimeline2Params object
// with the ability to set a context for a request.
func NewEdgeDiagnosticsGetResourceMetricsTimeline2ParamsWithContext(ctx context.Context) *EdgeDiagnosticsGetResourceMetricsTimeline2Params {
	return &EdgeDiagnosticsGetResourceMetricsTimeline2Params{
		Context: ctx,
	}
}

// NewEdgeDiagnosticsGetResourceMetricsTimeline2ParamsWithHTTPClient creates a new EdgeDiagnosticsGetResourceMetricsTimeline2Params object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeDiagnosticsGetResourceMetricsTimeline2ParamsWithHTTPClient(client *http.Client) *EdgeDiagnosticsGetResourceMetricsTimeline2Params {
	return &EdgeDiagnosticsGetResourceMetricsTimeline2Params{
		HTTPClient: client,
	}
}

/*
EdgeDiagnosticsGetResourceMetricsTimeline2Params contains all the parameters to send to the API endpoint

	for the edge diagnostics get resource metrics timeline2 operation.

	Typically these are written to a http.Request.
*/
type EdgeDiagnosticsGetResourceMetricsTimeline2Params struct {

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

// WithDefaults hydrates default values in the edge diagnostics get resource metrics timeline2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeDiagnosticsGetResourceMetricsTimeline2Params) WithDefaults() *EdgeDiagnosticsGetResourceMetricsTimeline2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge diagnostics get resource metrics timeline2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeDiagnosticsGetResourceMetricsTimeline2Params) SetDefaults() {
	var (
		objtypeDefault = string("OBJECT_TYPE_UNSPECIFIED")
	)

	val := EdgeDiagnosticsGetResourceMetricsTimeline2Params{
		Objtype: &objtypeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the edge diagnostics get resource metrics timeline2 params
func (o *EdgeDiagnosticsGetResourceMetricsTimeline2Params) WithTimeout(timeout time.Duration) *EdgeDiagnosticsGetResourceMetricsTimeline2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge diagnostics get resource metrics timeline2 params
func (o *EdgeDiagnosticsGetResourceMetricsTimeline2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge diagnostics get resource metrics timeline2 params
func (o *EdgeDiagnosticsGetResourceMetricsTimeline2Params) WithContext(ctx context.Context) *EdgeDiagnosticsGetResourceMetricsTimeline2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge diagnostics get resource metrics timeline2 params
func (o *EdgeDiagnosticsGetResourceMetricsTimeline2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge diagnostics get resource metrics timeline2 params
func (o *EdgeDiagnosticsGetResourceMetricsTimeline2Params) WithHTTPClient(client *http.Client) *EdgeDiagnosticsGetResourceMetricsTimeline2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge diagnostics get resource metrics timeline2 params
func (o *EdgeDiagnosticsGetResourceMetricsTimeline2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the edge diagnostics get resource metrics timeline2 params
func (o *EdgeDiagnosticsGetResourceMetricsTimeline2Params) WithXRequestID(xRequestID *string) *EdgeDiagnosticsGetResourceMetricsTimeline2Params {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the edge diagnostics get resource metrics timeline2 params
func (o *EdgeDiagnosticsGetResourceMetricsTimeline2Params) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithEndTime adds the endTime to the edge diagnostics get resource metrics timeline2 params
func (o *EdgeDiagnosticsGetResourceMetricsTimeline2Params) WithEndTime(endTime *strfmt.DateTime) *EdgeDiagnosticsGetResourceMetricsTimeline2Params {
	o.SetEndTime(endTime)
	return o
}

// SetEndTime adds the endTime to the edge diagnostics get resource metrics timeline2 params
func (o *EdgeDiagnosticsGetResourceMetricsTimeline2Params) SetEndTime(endTime *strfmt.DateTime) {
	o.EndTime = endTime
}

// WithEnterpriseID adds the enterpriseID to the edge diagnostics get resource metrics timeline2 params
func (o *EdgeDiagnosticsGetResourceMetricsTimeline2Params) WithEnterpriseID(enterpriseID *string) *EdgeDiagnosticsGetResourceMetricsTimeline2Params {
	o.SetEnterpriseID(enterpriseID)
	return o
}

// SetEnterpriseID adds the enterpriseId to the edge diagnostics get resource metrics timeline2 params
func (o *EdgeDiagnosticsGetResourceMetricsTimeline2Params) SetEnterpriseID(enterpriseID *string) {
	o.EnterpriseID = enterpriseID
}

// WithInterval adds the interval to the edge diagnostics get resource metrics timeline2 params
func (o *EdgeDiagnosticsGetResourceMetricsTimeline2Params) WithInterval(interval *strfmt.DateTime) *EdgeDiagnosticsGetResourceMetricsTimeline2Params {
	o.SetInterval(interval)
	return o
}

// SetInterval adds the interval to the edge diagnostics get resource metrics timeline2 params
func (o *EdgeDiagnosticsGetResourceMetricsTimeline2Params) SetInterval(interval *strfmt.DateTime) {
	o.Interval = interval
}

// WithMType adds the mType to the edge diagnostics get resource metrics timeline2 params
func (o *EdgeDiagnosticsGetResourceMetricsTimeline2Params) WithMType(mType string) *EdgeDiagnosticsGetResourceMetricsTimeline2Params {
	o.SetMType(mType)
	return o
}

// SetMType adds the mType to the edge diagnostics get resource metrics timeline2 params
func (o *EdgeDiagnosticsGetResourceMetricsTimeline2Params) SetMType(mType string) {
	o.MType = mType
}

// WithObjid adds the objid to the edge diagnostics get resource metrics timeline2 params
func (o *EdgeDiagnosticsGetResourceMetricsTimeline2Params) WithObjid(objid *string) *EdgeDiagnosticsGetResourceMetricsTimeline2Params {
	o.SetObjid(objid)
	return o
}

// SetObjid adds the objid to the edge diagnostics get resource metrics timeline2 params
func (o *EdgeDiagnosticsGetResourceMetricsTimeline2Params) SetObjid(objid *string) {
	o.Objid = objid
}

// WithObjname adds the objname to the edge diagnostics get resource metrics timeline2 params
func (o *EdgeDiagnosticsGetResourceMetricsTimeline2Params) WithObjname(objname *string) *EdgeDiagnosticsGetResourceMetricsTimeline2Params {
	o.SetObjname(objname)
	return o
}

// SetObjname adds the objname to the edge diagnostics get resource metrics timeline2 params
func (o *EdgeDiagnosticsGetResourceMetricsTimeline2Params) SetObjname(objname *string) {
	o.Objname = objname
}

// WithObjtype adds the objtype to the edge diagnostics get resource metrics timeline2 params
func (o *EdgeDiagnosticsGetResourceMetricsTimeline2Params) WithObjtype(objtype *string) *EdgeDiagnosticsGetResourceMetricsTimeline2Params {
	o.SetObjtype(objtype)
	return o
}

// SetObjtype adds the objtype to the edge diagnostics get resource metrics timeline2 params
func (o *EdgeDiagnosticsGetResourceMetricsTimeline2Params) SetObjtype(objtype *string) {
	o.Objtype = objtype
}

// WithStartTime adds the startTime to the edge diagnostics get resource metrics timeline2 params
func (o *EdgeDiagnosticsGetResourceMetricsTimeline2Params) WithStartTime(startTime *strfmt.DateTime) *EdgeDiagnosticsGetResourceMetricsTimeline2Params {
	o.SetStartTime(startTime)
	return o
}

// SetStartTime adds the startTime to the edge diagnostics get resource metrics timeline2 params
func (o *EdgeDiagnosticsGetResourceMetricsTimeline2Params) SetStartTime(startTime *strfmt.DateTime) {
	o.StartTime = startTime
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeDiagnosticsGetResourceMetricsTimeline2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

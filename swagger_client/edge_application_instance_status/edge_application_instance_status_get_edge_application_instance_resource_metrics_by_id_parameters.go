// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package edge_application_instance_status

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

// NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams creates a new EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams() *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams {
	return &EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParamsWithTimeout creates a new EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams object
// with the ability to set a timeout on a request.
func NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParamsWithTimeout(timeout time.Duration) *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams {
	return &EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams{
		timeout: timeout,
	}
}

// NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParamsWithContext creates a new EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams object
// with the ability to set a context for a request.
func NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParamsWithContext(ctx context.Context) *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams {
	return &EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams{
		Context: ctx,
	}
}

// NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParamsWithHTTPClient creates a new EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParamsWithHTTPClient(client *http.Client) *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams {
	return &EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams{
		HTTPClient: client,
	}
}

/* EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams contains all the parameters to send to the API endpoint
   for the edge application instance status get edge application instance resource metrics by Id operation.

   Typically these are written to a http.Request.
*/
type EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams struct {

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
	Objid string

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

// WithDefaults hydrates default values in the edge application instance status get edge application instance resource metrics by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams) WithDefaults() *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge application instance status get edge application instance resource metrics by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams) SetDefaults() {
	var (
		objtypeDefault = string("OBJECT_TYPE_UNSPECIFIED")
	)

	val := EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams{
		Objtype: &objtypeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the edge application instance status get edge application instance resource metrics by Id params
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams) WithTimeout(timeout time.Duration) *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge application instance status get edge application instance resource metrics by Id params
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge application instance status get edge application instance resource metrics by Id params
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams) WithContext(ctx context.Context) *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge application instance status get edge application instance resource metrics by Id params
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge application instance status get edge application instance resource metrics by Id params
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams) WithHTTPClient(client *http.Client) *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge application instance status get edge application instance resource metrics by Id params
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the edge application instance status get edge application instance resource metrics by Id params
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams) WithXRequestID(xRequestID *string) *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the edge application instance status get edge application instance resource metrics by Id params
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithEndTime adds the endTime to the edge application instance status get edge application instance resource metrics by Id params
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams) WithEndTime(endTime *strfmt.DateTime) *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams {
	o.SetEndTime(endTime)
	return o
}

// SetEndTime adds the endTime to the edge application instance status get edge application instance resource metrics by Id params
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams) SetEndTime(endTime *strfmt.DateTime) {
	o.EndTime = endTime
}

// WithEnterpriseID adds the enterpriseID to the edge application instance status get edge application instance resource metrics by Id params
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams) WithEnterpriseID(enterpriseID *string) *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams {
	o.SetEnterpriseID(enterpriseID)
	return o
}

// SetEnterpriseID adds the enterpriseId to the edge application instance status get edge application instance resource metrics by Id params
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams) SetEnterpriseID(enterpriseID *string) {
	o.EnterpriseID = enterpriseID
}

// WithInterval adds the interval to the edge application instance status get edge application instance resource metrics by Id params
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams) WithInterval(interval *strfmt.DateTime) *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams {
	o.SetInterval(interval)
	return o
}

// SetInterval adds the interval to the edge application instance status get edge application instance resource metrics by Id params
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams) SetInterval(interval *strfmt.DateTime) {
	o.Interval = interval
}

// WithMType adds the mType to the edge application instance status get edge application instance resource metrics by Id params
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams) WithMType(mType string) *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams {
	o.SetMType(mType)
	return o
}

// SetMType adds the mType to the edge application instance status get edge application instance resource metrics by Id params
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams) SetMType(mType string) {
	o.MType = mType
}

// WithObjid adds the objid to the edge application instance status get edge application instance resource metrics by Id params
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams) WithObjid(objid string) *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams {
	o.SetObjid(objid)
	return o
}

// SetObjid adds the objid to the edge application instance status get edge application instance resource metrics by Id params
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams) SetObjid(objid string) {
	o.Objid = objid
}

// WithObjname adds the objname to the edge application instance status get edge application instance resource metrics by Id params
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams) WithObjname(objname *string) *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams {
	o.SetObjname(objname)
	return o
}

// SetObjname adds the objname to the edge application instance status get edge application instance resource metrics by Id params
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams) SetObjname(objname *string) {
	o.Objname = objname
}

// WithObjtype adds the objtype to the edge application instance status get edge application instance resource metrics by Id params
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams) WithObjtype(objtype *string) *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams {
	o.SetObjtype(objtype)
	return o
}

// SetObjtype adds the objtype to the edge application instance status get edge application instance resource metrics by Id params
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams) SetObjtype(objtype *string) {
	o.Objtype = objtype
}

// WithStartTime adds the startTime to the edge application instance status get edge application instance resource metrics by Id params
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams) WithStartTime(startTime *strfmt.DateTime) *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams {
	o.SetStartTime(startTime)
	return o
}

// SetStartTime adds the startTime to the edge application instance status get edge application instance resource metrics by Id params
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams) SetStartTime(startTime *strfmt.DateTime) {
	o.StartTime = startTime
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

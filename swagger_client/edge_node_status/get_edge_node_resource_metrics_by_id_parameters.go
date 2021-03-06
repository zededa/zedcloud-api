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

// NewGetEdgeNodeResourceMetricsByIDParams creates a new GetEdgeNodeResourceMetricsByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetEdgeNodeResourceMetricsByIDParams() *GetEdgeNodeResourceMetricsByIDParams {
	return &GetEdgeNodeResourceMetricsByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetEdgeNodeResourceMetricsByIDParamsWithTimeout creates a new GetEdgeNodeResourceMetricsByIDParams object
// with the ability to set a timeout on a request.
func NewGetEdgeNodeResourceMetricsByIDParamsWithTimeout(timeout time.Duration) *GetEdgeNodeResourceMetricsByIDParams {
	return &GetEdgeNodeResourceMetricsByIDParams{
		timeout: timeout,
	}
}

// NewGetEdgeNodeResourceMetricsByIDParamsWithContext creates a new GetEdgeNodeResourceMetricsByIDParams object
// with the ability to set a context for a request.
func NewGetEdgeNodeResourceMetricsByIDParamsWithContext(ctx context.Context) *GetEdgeNodeResourceMetricsByIDParams {
	return &GetEdgeNodeResourceMetricsByIDParams{
		Context: ctx,
	}
}

// NewGetEdgeNodeResourceMetricsByIDParamsWithHTTPClient creates a new GetEdgeNodeResourceMetricsByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetEdgeNodeResourceMetricsByIDParamsWithHTTPClient(client *http.Client) *GetEdgeNodeResourceMetricsByIDParams {
	return &GetEdgeNodeResourceMetricsByIDParams{
		HTTPClient: client,
	}
}

/* GetEdgeNodeResourceMetricsByIDParams contains all the parameters to send to the API endpoint
   for the get edge node resource metrics by Id operation.

   Typically these are written to a http.Request.
*/
type GetEdgeNodeResourceMetricsByIDParams struct {

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

// WithDefaults hydrates default values in the get edge node resource metrics by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEdgeNodeResourceMetricsByIDParams) WithDefaults() *GetEdgeNodeResourceMetricsByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get edge node resource metrics by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEdgeNodeResourceMetricsByIDParams) SetDefaults() {
	var (
		objtypeDefault = string("OBJECT_TYPE_UNSPECIFIED")
	)

	val := GetEdgeNodeResourceMetricsByIDParams{
		Objtype: &objtypeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get edge node resource metrics by Id params
func (o *GetEdgeNodeResourceMetricsByIDParams) WithTimeout(timeout time.Duration) *GetEdgeNodeResourceMetricsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get edge node resource metrics by Id params
func (o *GetEdgeNodeResourceMetricsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get edge node resource metrics by Id params
func (o *GetEdgeNodeResourceMetricsByIDParams) WithContext(ctx context.Context) *GetEdgeNodeResourceMetricsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get edge node resource metrics by Id params
func (o *GetEdgeNodeResourceMetricsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get edge node resource metrics by Id params
func (o *GetEdgeNodeResourceMetricsByIDParams) WithHTTPClient(client *http.Client) *GetEdgeNodeResourceMetricsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get edge node resource metrics by Id params
func (o *GetEdgeNodeResourceMetricsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get edge node resource metrics by Id params
func (o *GetEdgeNodeResourceMetricsByIDParams) WithXRequestID(xRequestID *string) *GetEdgeNodeResourceMetricsByIDParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get edge node resource metrics by Id params
func (o *GetEdgeNodeResourceMetricsByIDParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithEndTime adds the endTime to the get edge node resource metrics by Id params
func (o *GetEdgeNodeResourceMetricsByIDParams) WithEndTime(endTime *strfmt.DateTime) *GetEdgeNodeResourceMetricsByIDParams {
	o.SetEndTime(endTime)
	return o
}

// SetEndTime adds the endTime to the get edge node resource metrics by Id params
func (o *GetEdgeNodeResourceMetricsByIDParams) SetEndTime(endTime *strfmt.DateTime) {
	o.EndTime = endTime
}

// WithEnterpriseID adds the enterpriseID to the get edge node resource metrics by Id params
func (o *GetEdgeNodeResourceMetricsByIDParams) WithEnterpriseID(enterpriseID *string) *GetEdgeNodeResourceMetricsByIDParams {
	o.SetEnterpriseID(enterpriseID)
	return o
}

// SetEnterpriseID adds the enterpriseId to the get edge node resource metrics by Id params
func (o *GetEdgeNodeResourceMetricsByIDParams) SetEnterpriseID(enterpriseID *string) {
	o.EnterpriseID = enterpriseID
}

// WithInterval adds the interval to the get edge node resource metrics by Id params
func (o *GetEdgeNodeResourceMetricsByIDParams) WithInterval(interval *strfmt.DateTime) *GetEdgeNodeResourceMetricsByIDParams {
	o.SetInterval(interval)
	return o
}

// SetInterval adds the interval to the get edge node resource metrics by Id params
func (o *GetEdgeNodeResourceMetricsByIDParams) SetInterval(interval *strfmt.DateTime) {
	o.Interval = interval
}

// WithMType adds the mType to the get edge node resource metrics by Id params
func (o *GetEdgeNodeResourceMetricsByIDParams) WithMType(mType string) *GetEdgeNodeResourceMetricsByIDParams {
	o.SetMType(mType)
	return o
}

// SetMType adds the mType to the get edge node resource metrics by Id params
func (o *GetEdgeNodeResourceMetricsByIDParams) SetMType(mType string) {
	o.MType = mType
}

// WithObjid adds the objid to the get edge node resource metrics by Id params
func (o *GetEdgeNodeResourceMetricsByIDParams) WithObjid(objid string) *GetEdgeNodeResourceMetricsByIDParams {
	o.SetObjid(objid)
	return o
}

// SetObjid adds the objid to the get edge node resource metrics by Id params
func (o *GetEdgeNodeResourceMetricsByIDParams) SetObjid(objid string) {
	o.Objid = objid
}

// WithObjname adds the objname to the get edge node resource metrics by Id params
func (o *GetEdgeNodeResourceMetricsByIDParams) WithObjname(objname *string) *GetEdgeNodeResourceMetricsByIDParams {
	o.SetObjname(objname)
	return o
}

// SetObjname adds the objname to the get edge node resource metrics by Id params
func (o *GetEdgeNodeResourceMetricsByIDParams) SetObjname(objname *string) {
	o.Objname = objname
}

// WithObjtype adds the objtype to the get edge node resource metrics by Id params
func (o *GetEdgeNodeResourceMetricsByIDParams) WithObjtype(objtype *string) *GetEdgeNodeResourceMetricsByIDParams {
	o.SetObjtype(objtype)
	return o
}

// SetObjtype adds the objtype to the get edge node resource metrics by Id params
func (o *GetEdgeNodeResourceMetricsByIDParams) SetObjtype(objtype *string) {
	o.Objtype = objtype
}

// WithStartTime adds the startTime to the get edge node resource metrics by Id params
func (o *GetEdgeNodeResourceMetricsByIDParams) WithStartTime(startTime *strfmt.DateTime) *GetEdgeNodeResourceMetricsByIDParams {
	o.SetStartTime(startTime)
	return o
}

// SetStartTime adds the startTime to the get edge node resource metrics by Id params
func (o *GetEdgeNodeResourceMetricsByIDParams) SetStartTime(startTime *strfmt.DateTime) {
	o.StartTime = startTime
}

// WriteToRequest writes these params to a swagger request
func (o *GetEdgeNodeResourceMetricsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

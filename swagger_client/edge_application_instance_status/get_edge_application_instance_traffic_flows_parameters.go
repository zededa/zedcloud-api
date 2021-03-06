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
	"github.com/go-openapi/swag"
)

// NewGetEdgeApplicationInstanceTrafficFlowsParams creates a new GetEdgeApplicationInstanceTrafficFlowsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetEdgeApplicationInstanceTrafficFlowsParams() *GetEdgeApplicationInstanceTrafficFlowsParams {
	return &GetEdgeApplicationInstanceTrafficFlowsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetEdgeApplicationInstanceTrafficFlowsParamsWithTimeout creates a new GetEdgeApplicationInstanceTrafficFlowsParams object
// with the ability to set a timeout on a request.
func NewGetEdgeApplicationInstanceTrafficFlowsParamsWithTimeout(timeout time.Duration) *GetEdgeApplicationInstanceTrafficFlowsParams {
	return &GetEdgeApplicationInstanceTrafficFlowsParams{
		timeout: timeout,
	}
}

// NewGetEdgeApplicationInstanceTrafficFlowsParamsWithContext creates a new GetEdgeApplicationInstanceTrafficFlowsParams object
// with the ability to set a context for a request.
func NewGetEdgeApplicationInstanceTrafficFlowsParamsWithContext(ctx context.Context) *GetEdgeApplicationInstanceTrafficFlowsParams {
	return &GetEdgeApplicationInstanceTrafficFlowsParams{
		Context: ctx,
	}
}

// NewGetEdgeApplicationInstanceTrafficFlowsParamsWithHTTPClient creates a new GetEdgeApplicationInstanceTrafficFlowsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetEdgeApplicationInstanceTrafficFlowsParamsWithHTTPClient(client *http.Client) *GetEdgeApplicationInstanceTrafficFlowsParams {
	return &GetEdgeApplicationInstanceTrafficFlowsParams{
		HTTPClient: client,
	}
}

/* GetEdgeApplicationInstanceTrafficFlowsParams contains all the parameters to send to the API endpoint
   for the get edge application instance traffic flows operation.

   Typically these are written to a http.Request.
*/
type GetEdgeApplicationInstanceTrafficFlowsParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Aclid.
	//
	// Format: int64
	Aclid *int64

	// Action.
	//
	// Default: "FLOW_LOG_ACTION_UNSPECIFIED"
	Action *string

	// Direction.
	//
	// Default: "FLOW_LOG_DIRECTION_UNSPECIFIED"
	Direction *string

	// EndTime.
	//
	// Format: date-time
	EndTime *strfmt.DateTime

	// EnterpriseID.
	EnterpriseID *string

	// Fqdn.
	Fqdn *string

	// ID.
	ID string

	// Levels.
	Levels []string

	// MetricType.
	//
	// Default: "FLOW_LOG_METRIC_UNSPECIFIED"
	MetricType *string

	// Name.
	Name *string

	// Proto.
	//
	// Format: int64
	Proto *int64

	// Remoteip.
	Remoteip *string

	// Remoteport.
	//
	// Format: int64
	Remoteport *int64

	// StartTime.
	//
	// Format: date-time
	StartTime *strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get edge application instance traffic flows params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEdgeApplicationInstanceTrafficFlowsParams) WithDefaults() *GetEdgeApplicationInstanceTrafficFlowsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get edge application instance traffic flows params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEdgeApplicationInstanceTrafficFlowsParams) SetDefaults() {
	var (
		actionDefault = string("FLOW_LOG_ACTION_UNSPECIFIED")

		directionDefault = string("FLOW_LOG_DIRECTION_UNSPECIFIED")

		metricTypeDefault = string("FLOW_LOG_METRIC_UNSPECIFIED")
	)

	val := GetEdgeApplicationInstanceTrafficFlowsParams{
		Action:     &actionDefault,
		Direction:  &directionDefault,
		MetricType: &metricTypeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get edge application instance traffic flows params
func (o *GetEdgeApplicationInstanceTrafficFlowsParams) WithTimeout(timeout time.Duration) *GetEdgeApplicationInstanceTrafficFlowsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get edge application instance traffic flows params
func (o *GetEdgeApplicationInstanceTrafficFlowsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get edge application instance traffic flows params
func (o *GetEdgeApplicationInstanceTrafficFlowsParams) WithContext(ctx context.Context) *GetEdgeApplicationInstanceTrafficFlowsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get edge application instance traffic flows params
func (o *GetEdgeApplicationInstanceTrafficFlowsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get edge application instance traffic flows params
func (o *GetEdgeApplicationInstanceTrafficFlowsParams) WithHTTPClient(client *http.Client) *GetEdgeApplicationInstanceTrafficFlowsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get edge application instance traffic flows params
func (o *GetEdgeApplicationInstanceTrafficFlowsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get edge application instance traffic flows params
func (o *GetEdgeApplicationInstanceTrafficFlowsParams) WithXRequestID(xRequestID *string) *GetEdgeApplicationInstanceTrafficFlowsParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get edge application instance traffic flows params
func (o *GetEdgeApplicationInstanceTrafficFlowsParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithAclid adds the aclid to the get edge application instance traffic flows params
func (o *GetEdgeApplicationInstanceTrafficFlowsParams) WithAclid(aclid *int64) *GetEdgeApplicationInstanceTrafficFlowsParams {
	o.SetAclid(aclid)
	return o
}

// SetAclid adds the aclid to the get edge application instance traffic flows params
func (o *GetEdgeApplicationInstanceTrafficFlowsParams) SetAclid(aclid *int64) {
	o.Aclid = aclid
}

// WithAction adds the action to the get edge application instance traffic flows params
func (o *GetEdgeApplicationInstanceTrafficFlowsParams) WithAction(action *string) *GetEdgeApplicationInstanceTrafficFlowsParams {
	o.SetAction(action)
	return o
}

// SetAction adds the action to the get edge application instance traffic flows params
func (o *GetEdgeApplicationInstanceTrafficFlowsParams) SetAction(action *string) {
	o.Action = action
}

// WithDirection adds the direction to the get edge application instance traffic flows params
func (o *GetEdgeApplicationInstanceTrafficFlowsParams) WithDirection(direction *string) *GetEdgeApplicationInstanceTrafficFlowsParams {
	o.SetDirection(direction)
	return o
}

// SetDirection adds the direction to the get edge application instance traffic flows params
func (o *GetEdgeApplicationInstanceTrafficFlowsParams) SetDirection(direction *string) {
	o.Direction = direction
}

// WithEndTime adds the endTime to the get edge application instance traffic flows params
func (o *GetEdgeApplicationInstanceTrafficFlowsParams) WithEndTime(endTime *strfmt.DateTime) *GetEdgeApplicationInstanceTrafficFlowsParams {
	o.SetEndTime(endTime)
	return o
}

// SetEndTime adds the endTime to the get edge application instance traffic flows params
func (o *GetEdgeApplicationInstanceTrafficFlowsParams) SetEndTime(endTime *strfmt.DateTime) {
	o.EndTime = endTime
}

// WithEnterpriseID adds the enterpriseID to the get edge application instance traffic flows params
func (o *GetEdgeApplicationInstanceTrafficFlowsParams) WithEnterpriseID(enterpriseID *string) *GetEdgeApplicationInstanceTrafficFlowsParams {
	o.SetEnterpriseID(enterpriseID)
	return o
}

// SetEnterpriseID adds the enterpriseId to the get edge application instance traffic flows params
func (o *GetEdgeApplicationInstanceTrafficFlowsParams) SetEnterpriseID(enterpriseID *string) {
	o.EnterpriseID = enterpriseID
}

// WithFqdn adds the fqdn to the get edge application instance traffic flows params
func (o *GetEdgeApplicationInstanceTrafficFlowsParams) WithFqdn(fqdn *string) *GetEdgeApplicationInstanceTrafficFlowsParams {
	o.SetFqdn(fqdn)
	return o
}

// SetFqdn adds the fqdn to the get edge application instance traffic flows params
func (o *GetEdgeApplicationInstanceTrafficFlowsParams) SetFqdn(fqdn *string) {
	o.Fqdn = fqdn
}

// WithID adds the id to the get edge application instance traffic flows params
func (o *GetEdgeApplicationInstanceTrafficFlowsParams) WithID(id string) *GetEdgeApplicationInstanceTrafficFlowsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get edge application instance traffic flows params
func (o *GetEdgeApplicationInstanceTrafficFlowsParams) SetID(id string) {
	o.ID = id
}

// WithLevels adds the levels to the get edge application instance traffic flows params
func (o *GetEdgeApplicationInstanceTrafficFlowsParams) WithLevels(levels []string) *GetEdgeApplicationInstanceTrafficFlowsParams {
	o.SetLevels(levels)
	return o
}

// SetLevels adds the levels to the get edge application instance traffic flows params
func (o *GetEdgeApplicationInstanceTrafficFlowsParams) SetLevels(levels []string) {
	o.Levels = levels
}

// WithMetricType adds the metricType to the get edge application instance traffic flows params
func (o *GetEdgeApplicationInstanceTrafficFlowsParams) WithMetricType(metricType *string) *GetEdgeApplicationInstanceTrafficFlowsParams {
	o.SetMetricType(metricType)
	return o
}

// SetMetricType adds the metricType to the get edge application instance traffic flows params
func (o *GetEdgeApplicationInstanceTrafficFlowsParams) SetMetricType(metricType *string) {
	o.MetricType = metricType
}

// WithName adds the name to the get edge application instance traffic flows params
func (o *GetEdgeApplicationInstanceTrafficFlowsParams) WithName(name *string) *GetEdgeApplicationInstanceTrafficFlowsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get edge application instance traffic flows params
func (o *GetEdgeApplicationInstanceTrafficFlowsParams) SetName(name *string) {
	o.Name = name
}

// WithProto adds the proto to the get edge application instance traffic flows params
func (o *GetEdgeApplicationInstanceTrafficFlowsParams) WithProto(proto *int64) *GetEdgeApplicationInstanceTrafficFlowsParams {
	o.SetProto(proto)
	return o
}

// SetProto adds the proto to the get edge application instance traffic flows params
func (o *GetEdgeApplicationInstanceTrafficFlowsParams) SetProto(proto *int64) {
	o.Proto = proto
}

// WithRemoteip adds the remoteip to the get edge application instance traffic flows params
func (o *GetEdgeApplicationInstanceTrafficFlowsParams) WithRemoteip(remoteip *string) *GetEdgeApplicationInstanceTrafficFlowsParams {
	o.SetRemoteip(remoteip)
	return o
}

// SetRemoteip adds the remoteip to the get edge application instance traffic flows params
func (o *GetEdgeApplicationInstanceTrafficFlowsParams) SetRemoteip(remoteip *string) {
	o.Remoteip = remoteip
}

// WithRemoteport adds the remoteport to the get edge application instance traffic flows params
func (o *GetEdgeApplicationInstanceTrafficFlowsParams) WithRemoteport(remoteport *int64) *GetEdgeApplicationInstanceTrafficFlowsParams {
	o.SetRemoteport(remoteport)
	return o
}

// SetRemoteport adds the remoteport to the get edge application instance traffic flows params
func (o *GetEdgeApplicationInstanceTrafficFlowsParams) SetRemoteport(remoteport *int64) {
	o.Remoteport = remoteport
}

// WithStartTime adds the startTime to the get edge application instance traffic flows params
func (o *GetEdgeApplicationInstanceTrafficFlowsParams) WithStartTime(startTime *strfmt.DateTime) *GetEdgeApplicationInstanceTrafficFlowsParams {
	o.SetStartTime(startTime)
	return o
}

// SetStartTime adds the startTime to the get edge application instance traffic flows params
func (o *GetEdgeApplicationInstanceTrafficFlowsParams) SetStartTime(startTime *strfmt.DateTime) {
	o.StartTime = startTime
}

// WriteToRequest writes these params to a swagger request
func (o *GetEdgeApplicationInstanceTrafficFlowsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Aclid != nil {

		// query param aclid
		var qrAclid int64

		if o.Aclid != nil {
			qrAclid = *o.Aclid
		}
		qAclid := swag.FormatInt64(qrAclid)
		if qAclid != "" {

			if err := r.SetQueryParam("aclid", qAclid); err != nil {
				return err
			}
		}
	}

	if o.Action != nil {

		// query param action
		var qrAction string

		if o.Action != nil {
			qrAction = *o.Action
		}
		qAction := qrAction
		if qAction != "" {

			if err := r.SetQueryParam("action", qAction); err != nil {
				return err
			}
		}
	}

	if o.Direction != nil {

		// query param direction
		var qrDirection string

		if o.Direction != nil {
			qrDirection = *o.Direction
		}
		qDirection := qrDirection
		if qDirection != "" {

			if err := r.SetQueryParam("direction", qDirection); err != nil {
				return err
			}
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

	if o.Fqdn != nil {

		// query param fqdn
		var qrFqdn string

		if o.Fqdn != nil {
			qrFqdn = *o.Fqdn
		}
		qFqdn := qrFqdn
		if qFqdn != "" {

			if err := r.SetQueryParam("fqdn", qFqdn); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Levels != nil {

		// binding items for levels
		joinedLevels := o.bindParamLevels(reg)

		// query array param levels
		if err := r.SetQueryParam("levels", joinedLevels...); err != nil {
			return err
		}
	}

	if o.MetricType != nil {

		// query param metricType
		var qrMetricType string

		if o.MetricType != nil {
			qrMetricType = *o.MetricType
		}
		qMetricType := qrMetricType
		if qMetricType != "" {

			if err := r.SetQueryParam("metricType", qMetricType); err != nil {
				return err
			}
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.Proto != nil {

		// query param proto
		var qrProto int64

		if o.Proto != nil {
			qrProto = *o.Proto
		}
		qProto := swag.FormatInt64(qrProto)
		if qProto != "" {

			if err := r.SetQueryParam("proto", qProto); err != nil {
				return err
			}
		}
	}

	if o.Remoteip != nil {

		// query param remoteip
		var qrRemoteip string

		if o.Remoteip != nil {
			qrRemoteip = *o.Remoteip
		}
		qRemoteip := qrRemoteip
		if qRemoteip != "" {

			if err := r.SetQueryParam("remoteip", qRemoteip); err != nil {
				return err
			}
		}
	}

	if o.Remoteport != nil {

		// query param remoteport
		var qrRemoteport int64

		if o.Remoteport != nil {
			qrRemoteport = *o.Remoteport
		}
		qRemoteport := swag.FormatInt64(qrRemoteport)
		if qRemoteport != "" {

			if err := r.SetQueryParam("remoteport", qRemoteport); err != nil {
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

// bindParamGetEdgeApplicationInstanceTrafficFlows binds the parameter levels
func (o *GetEdgeApplicationInstanceTrafficFlowsParams) bindParamLevels(formats strfmt.Registry) []string {
	levelsIR := o.Levels

	var levelsIC []string
	for _, levelsIIR := range levelsIR { // explode []string

		levelsIIV := levelsIIR // string as string
		levelsIC = append(levelsIC, levelsIIV)
	}

	// items.CollectionFormat: "multi"
	levelsIS := swag.JoinByFormat(levelsIC, "multi")

	return levelsIS
}

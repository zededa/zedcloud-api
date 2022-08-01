// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package edge_application_instance_configuration

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

// NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams creates a new EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams() *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams {
	return &EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParamsWithTimeout creates a new EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams object
// with the ability to set a timeout on a request.
func NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParamsWithTimeout(timeout time.Duration) *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams {
	return &EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams{
		timeout: timeout,
	}
}

// NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParamsWithContext creates a new EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams object
// with the ability to set a context for a request.
func NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParamsWithContext(ctx context.Context) *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams {
	return &EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams{
		Context: ctx,
	}
}

// NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParamsWithHTTPClient creates a new EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParamsWithHTTPClient(client *http.Client) *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams {
	return &EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams{
		HTTPClient: client,
	}
}

/* EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams contains all the parameters to send to the API endpoint
   for the edge application instance configuration query edge application instances operation.

   Typically these are written to a http.Request.
*/
type EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* Fields.

	   Fields to select: id, name, title, description, projectId, tags, clusterID, appId, deviceId, manifestInfo, appType, appPolicyId, userDefinedVersion, remoteConsole
	*/
	Fields []string

	/* FilterAppName.

	   User defined name of the app instance, unique across the enterprise. Once app instance is created, name can’t be changed
	*/
	FilterAppName *string

	/* FilterAppType.

	   type of bundle

	   Default: "APP_TYPE_UNSPECIFIED"
	*/
	FilterAppType *string

	/* FilterDeploymentType.

	   type of deployment for the app, eg: azure, k3s, standalone

	   Default: "DEPLOYMENT_TYPE_UNSPECIFIED"
	*/
	FilterDeploymentType *string

	/* FilterDeviceName.

	   User defined name of the device, unique across the enterprise. Once device is created, name can’t be changed
	*/
	FilterDeviceName *string

	/* FilterDeviceNamePattern.

	   device name pattern
	*/
	FilterDeviceNamePattern *string

	/* FilterNamePattern.

	   name pattern
	*/
	FilterNamePattern *string

	/* FilterProjectName.

	   User defined name of the project, unique across the enterprise. Once project is created, name can’t be changed
	*/
	FilterProjectName *string

	/* FilterProjectNamePattern.

	   project name pattern
	*/
	FilterProjectNamePattern *string

	/* NextOrderBy.

	   OrderBy helps in sorting the list response
	*/
	NextOrderBy []string

	/* NextPageNum.

	   Page Number

	   Format: int64
	*/
	NextPageNum *int64

	/* NextPageSize.

	   Defines the page size

	   Format: int64
	*/
	NextPageSize *int64

	/* NextPageToken.

	   Page Token
	*/
	NextPageToken *string

	/* NextTotalPages.

	   Total number of pages to be fetched.

	   Format: int64
	*/
	NextTotalPages *int64

	/* Summary.

	   summary flag
	*/
	Summary *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edge application instance configuration query edge application instances params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams) WithDefaults() *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge application instance configuration query edge application instances params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams) SetDefaults() {
	var (
		filterAppTypeDefault = string("APP_TYPE_UNSPECIFIED")

		filterDeploymentTypeDefault = string("DEPLOYMENT_TYPE_UNSPECIFIED")
	)

	val := EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams{
		FilterAppType:        &filterAppTypeDefault,
		FilterDeploymentType: &filterDeploymentTypeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the edge application instance configuration query edge application instances params
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams) WithTimeout(timeout time.Duration) *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge application instance configuration query edge application instances params
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge application instance configuration query edge application instances params
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams) WithContext(ctx context.Context) *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge application instance configuration query edge application instances params
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge application instance configuration query edge application instances params
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams) WithHTTPClient(client *http.Client) *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge application instance configuration query edge application instances params
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the edge application instance configuration query edge application instances params
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams) WithXRequestID(xRequestID *string) *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the edge application instance configuration query edge application instances params
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithFields adds the fields to the edge application instance configuration query edge application instances params
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams) WithFields(fields []string) *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the edge application instance configuration query edge application instances params
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithFilterAppName adds the filterAppName to the edge application instance configuration query edge application instances params
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams) WithFilterAppName(filterAppName *string) *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams {
	o.SetFilterAppName(filterAppName)
	return o
}

// SetFilterAppName adds the filterAppName to the edge application instance configuration query edge application instances params
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams) SetFilterAppName(filterAppName *string) {
	o.FilterAppName = filterAppName
}

// WithFilterAppType adds the filterAppType to the edge application instance configuration query edge application instances params
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams) WithFilterAppType(filterAppType *string) *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams {
	o.SetFilterAppType(filterAppType)
	return o
}

// SetFilterAppType adds the filterAppType to the edge application instance configuration query edge application instances params
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams) SetFilterAppType(filterAppType *string) {
	o.FilterAppType = filterAppType
}

// WithFilterDeploymentType adds the filterDeploymentType to the edge application instance configuration query edge application instances params
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams) WithFilterDeploymentType(filterDeploymentType *string) *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams {
	o.SetFilterDeploymentType(filterDeploymentType)
	return o
}

// SetFilterDeploymentType adds the filterDeploymentType to the edge application instance configuration query edge application instances params
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams) SetFilterDeploymentType(filterDeploymentType *string) {
	o.FilterDeploymentType = filterDeploymentType
}

// WithFilterDeviceName adds the filterDeviceName to the edge application instance configuration query edge application instances params
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams) WithFilterDeviceName(filterDeviceName *string) *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams {
	o.SetFilterDeviceName(filterDeviceName)
	return o
}

// SetFilterDeviceName adds the filterDeviceName to the edge application instance configuration query edge application instances params
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams) SetFilterDeviceName(filterDeviceName *string) {
	o.FilterDeviceName = filterDeviceName
}

// WithFilterDeviceNamePattern adds the filterDeviceNamePattern to the edge application instance configuration query edge application instances params
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams) WithFilterDeviceNamePattern(filterDeviceNamePattern *string) *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams {
	o.SetFilterDeviceNamePattern(filterDeviceNamePattern)
	return o
}

// SetFilterDeviceNamePattern adds the filterDeviceNamePattern to the edge application instance configuration query edge application instances params
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams) SetFilterDeviceNamePattern(filterDeviceNamePattern *string) {
	o.FilterDeviceNamePattern = filterDeviceNamePattern
}

// WithFilterNamePattern adds the filterNamePattern to the edge application instance configuration query edge application instances params
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams) WithFilterNamePattern(filterNamePattern *string) *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams {
	o.SetFilterNamePattern(filterNamePattern)
	return o
}

// SetFilterNamePattern adds the filterNamePattern to the edge application instance configuration query edge application instances params
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams) SetFilterNamePattern(filterNamePattern *string) {
	o.FilterNamePattern = filterNamePattern
}

// WithFilterProjectName adds the filterProjectName to the edge application instance configuration query edge application instances params
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams) WithFilterProjectName(filterProjectName *string) *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams {
	o.SetFilterProjectName(filterProjectName)
	return o
}

// SetFilterProjectName adds the filterProjectName to the edge application instance configuration query edge application instances params
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams) SetFilterProjectName(filterProjectName *string) {
	o.FilterProjectName = filterProjectName
}

// WithFilterProjectNamePattern adds the filterProjectNamePattern to the edge application instance configuration query edge application instances params
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams) WithFilterProjectNamePattern(filterProjectNamePattern *string) *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams {
	o.SetFilterProjectNamePattern(filterProjectNamePattern)
	return o
}

// SetFilterProjectNamePattern adds the filterProjectNamePattern to the edge application instance configuration query edge application instances params
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams) SetFilterProjectNamePattern(filterProjectNamePattern *string) {
	o.FilterProjectNamePattern = filterProjectNamePattern
}

// WithNextOrderBy adds the nextOrderBy to the edge application instance configuration query edge application instances params
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams) WithNextOrderBy(nextOrderBy []string) *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams {
	o.SetNextOrderBy(nextOrderBy)
	return o
}

// SetNextOrderBy adds the nextOrderBy to the edge application instance configuration query edge application instances params
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams) SetNextOrderBy(nextOrderBy []string) {
	o.NextOrderBy = nextOrderBy
}

// WithNextPageNum adds the nextPageNum to the edge application instance configuration query edge application instances params
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams) WithNextPageNum(nextPageNum *int64) *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams {
	o.SetNextPageNum(nextPageNum)
	return o
}

// SetNextPageNum adds the nextPageNum to the edge application instance configuration query edge application instances params
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams) SetNextPageNum(nextPageNum *int64) {
	o.NextPageNum = nextPageNum
}

// WithNextPageSize adds the nextPageSize to the edge application instance configuration query edge application instances params
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams) WithNextPageSize(nextPageSize *int64) *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams {
	o.SetNextPageSize(nextPageSize)
	return o
}

// SetNextPageSize adds the nextPageSize to the edge application instance configuration query edge application instances params
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams) SetNextPageSize(nextPageSize *int64) {
	o.NextPageSize = nextPageSize
}

// WithNextPageToken adds the nextPageToken to the edge application instance configuration query edge application instances params
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams) WithNextPageToken(nextPageToken *string) *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams {
	o.SetNextPageToken(nextPageToken)
	return o
}

// SetNextPageToken adds the nextPageToken to the edge application instance configuration query edge application instances params
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams) SetNextPageToken(nextPageToken *string) {
	o.NextPageToken = nextPageToken
}

// WithNextTotalPages adds the nextTotalPages to the edge application instance configuration query edge application instances params
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams) WithNextTotalPages(nextTotalPages *int64) *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams {
	o.SetNextTotalPages(nextTotalPages)
	return o
}

// SetNextTotalPages adds the nextTotalPages to the edge application instance configuration query edge application instances params
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams) SetNextTotalPages(nextTotalPages *int64) {
	o.NextTotalPages = nextTotalPages
}

// WithSummary adds the summary to the edge application instance configuration query edge application instances params
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams) WithSummary(summary *bool) *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams {
	o.SetSummary(summary)
	return o
}

// SetSummary adds the summary to the edge application instance configuration query edge application instances params
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams) SetSummary(summary *bool) {
	o.Summary = summary
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.FilterAppName != nil {

		// query param filter.appName
		var qrFilterAppName string

		if o.FilterAppName != nil {
			qrFilterAppName = *o.FilterAppName
		}
		qFilterAppName := qrFilterAppName
		if qFilterAppName != "" {

			if err := r.SetQueryParam("filter.appName", qFilterAppName); err != nil {
				return err
			}
		}
	}

	if o.FilterAppType != nil {

		// query param filter.appType
		var qrFilterAppType string

		if o.FilterAppType != nil {
			qrFilterAppType = *o.FilterAppType
		}
		qFilterAppType := qrFilterAppType
		if qFilterAppType != "" {

			if err := r.SetQueryParam("filter.appType", qFilterAppType); err != nil {
				return err
			}
		}
	}

	if o.FilterDeploymentType != nil {

		// query param filter.deploymentType
		var qrFilterDeploymentType string

		if o.FilterDeploymentType != nil {
			qrFilterDeploymentType = *o.FilterDeploymentType
		}
		qFilterDeploymentType := qrFilterDeploymentType
		if qFilterDeploymentType != "" {

			if err := r.SetQueryParam("filter.deploymentType", qFilterDeploymentType); err != nil {
				return err
			}
		}
	}

	if o.FilterDeviceName != nil {

		// query param filter.deviceName
		var qrFilterDeviceName string

		if o.FilterDeviceName != nil {
			qrFilterDeviceName = *o.FilterDeviceName
		}
		qFilterDeviceName := qrFilterDeviceName
		if qFilterDeviceName != "" {

			if err := r.SetQueryParam("filter.deviceName", qFilterDeviceName); err != nil {
				return err
			}
		}
	}

	if o.FilterDeviceNamePattern != nil {

		// query param filter.deviceNamePattern
		var qrFilterDeviceNamePattern string

		if o.FilterDeviceNamePattern != nil {
			qrFilterDeviceNamePattern = *o.FilterDeviceNamePattern
		}
		qFilterDeviceNamePattern := qrFilterDeviceNamePattern
		if qFilterDeviceNamePattern != "" {

			if err := r.SetQueryParam("filter.deviceNamePattern", qFilterDeviceNamePattern); err != nil {
				return err
			}
		}
	}

	if o.FilterNamePattern != nil {

		// query param filter.namePattern
		var qrFilterNamePattern string

		if o.FilterNamePattern != nil {
			qrFilterNamePattern = *o.FilterNamePattern
		}
		qFilterNamePattern := qrFilterNamePattern
		if qFilterNamePattern != "" {

			if err := r.SetQueryParam("filter.namePattern", qFilterNamePattern); err != nil {
				return err
			}
		}
	}

	if o.FilterProjectName != nil {

		// query param filter.projectName
		var qrFilterProjectName string

		if o.FilterProjectName != nil {
			qrFilterProjectName = *o.FilterProjectName
		}
		qFilterProjectName := qrFilterProjectName
		if qFilterProjectName != "" {

			if err := r.SetQueryParam("filter.projectName", qFilterProjectName); err != nil {
				return err
			}
		}
	}

	if o.FilterProjectNamePattern != nil {

		// query param filter.projectNamePattern
		var qrFilterProjectNamePattern string

		if o.FilterProjectNamePattern != nil {
			qrFilterProjectNamePattern = *o.FilterProjectNamePattern
		}
		qFilterProjectNamePattern := qrFilterProjectNamePattern
		if qFilterProjectNamePattern != "" {

			if err := r.SetQueryParam("filter.projectNamePattern", qFilterProjectNamePattern); err != nil {
				return err
			}
		}
	}

	if o.NextOrderBy != nil {

		// binding items for next.orderBy
		joinedNextOrderBy := o.bindParamNextOrderBy(reg)

		// query array param next.orderBy
		if err := r.SetQueryParam("next.orderBy", joinedNextOrderBy...); err != nil {
			return err
		}
	}

	if o.NextPageNum != nil {

		// query param next.pageNum
		var qrNextPageNum int64

		if o.NextPageNum != nil {
			qrNextPageNum = *o.NextPageNum
		}
		qNextPageNum := swag.FormatInt64(qrNextPageNum)
		if qNextPageNum != "" {

			if err := r.SetQueryParam("next.pageNum", qNextPageNum); err != nil {
				return err
			}
		}
	}

	if o.NextPageSize != nil {

		// query param next.pageSize
		var qrNextPageSize int64

		if o.NextPageSize != nil {
			qrNextPageSize = *o.NextPageSize
		}
		qNextPageSize := swag.FormatInt64(qrNextPageSize)
		if qNextPageSize != "" {

			if err := r.SetQueryParam("next.pageSize", qNextPageSize); err != nil {
				return err
			}
		}
	}

	if o.NextPageToken != nil {

		// query param next.pageToken
		var qrNextPageToken string

		if o.NextPageToken != nil {
			qrNextPageToken = *o.NextPageToken
		}
		qNextPageToken := qrNextPageToken
		if qNextPageToken != "" {

			if err := r.SetQueryParam("next.pageToken", qNextPageToken); err != nil {
				return err
			}
		}
	}

	if o.NextTotalPages != nil {

		// query param next.totalPages
		var qrNextTotalPages int64

		if o.NextTotalPages != nil {
			qrNextTotalPages = *o.NextTotalPages
		}
		qNextTotalPages := swag.FormatInt64(qrNextTotalPages)
		if qNextTotalPages != "" {

			if err := r.SetQueryParam("next.totalPages", qNextTotalPages); err != nil {
				return err
			}
		}
	}

	if o.Summary != nil {

		// query param summary
		var qrSummary bool

		if o.Summary != nil {
			qrSummary = *o.Summary
		}
		qSummary := swag.FormatBool(qrSummary)
		if qSummary != "" {

			if err := r.SetQueryParam("summary", qSummary); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstances binds the parameter fields
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.Fields

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "multi"
	fieldsIS := swag.JoinByFormat(fieldsIC, "multi")

	return fieldsIS
}

// bindParamEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstances binds the parameter next.orderBy
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesParams) bindParamNextOrderBy(formats strfmt.Registry) []string {
	nextOrderByIR := o.NextOrderBy

	var nextOrderByIC []string
	for _, nextOrderByIIR := range nextOrderByIR { // explode []string

		nextOrderByIIV := nextOrderByIIR // string as string
		nextOrderByIC = append(nextOrderByIC, nextOrderByIIV)
	}

	// items.CollectionFormat: "multi"
	nextOrderByIS := swag.JoinByFormat(nextOrderByIC, "multi")

	return nextOrderByIS
}

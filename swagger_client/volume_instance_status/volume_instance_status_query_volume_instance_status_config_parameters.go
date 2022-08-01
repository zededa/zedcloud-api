// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package volume_instance_status

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

// NewVolumeInstanceStatusQueryVolumeInstanceStatusConfigParams creates a new VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVolumeInstanceStatusQueryVolumeInstanceStatusConfigParams() *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams {
	return &VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVolumeInstanceStatusQueryVolumeInstanceStatusConfigParamsWithTimeout creates a new VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams object
// with the ability to set a timeout on a request.
func NewVolumeInstanceStatusQueryVolumeInstanceStatusConfigParamsWithTimeout(timeout time.Duration) *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams {
	return &VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams{
		timeout: timeout,
	}
}

// NewVolumeInstanceStatusQueryVolumeInstanceStatusConfigParamsWithContext creates a new VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams object
// with the ability to set a context for a request.
func NewVolumeInstanceStatusQueryVolumeInstanceStatusConfigParamsWithContext(ctx context.Context) *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams {
	return &VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams{
		Context: ctx,
	}
}

// NewVolumeInstanceStatusQueryVolumeInstanceStatusConfigParamsWithHTTPClient creates a new VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewVolumeInstanceStatusQueryVolumeInstanceStatusConfigParamsWithHTTPClient(client *http.Client) *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams {
	return &VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams{
		HTTPClient: client,
	}
}

/* VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams contains all the parameters to send to the API endpoint
   for the volume instance status query volume instance status config operation.

   Typically these are written to a http.Request.
*/
type VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Fields.
	Fields []string

	// FilterAppInstName.
	FilterAppInstName *string

	// FilterDeviceName.
	FilterDeviceName *string

	// FilterDeviceNamePattern.
	FilterDeviceNamePattern *string

	// FilterImageName.
	FilterImageName *string

	// FilterNamePattern.
	FilterNamePattern *string

	// FilterProjectName.
	FilterProjectName *string

	// FilterProjectNamePattern.
	FilterProjectNamePattern *string

	/* FilterRunState.

	     - RUN_STATE_ONLINE: Entity Online
	- RUN_STATE_HALTED: Entity Halted
	- RUN_STATE_INIT: Entity Initializing
	- RUN_STATE_REBOOTING: Entity Rebooting
	- RUN_STATE_OFFLINE: Entity Offline
	- RUN_STATE_UNKNOWN: Entity state Unknown
	- RUN_STATE_UNPROVISIONED: Entity Unprovisioned
	- RUN_STATE_PROVISIONED: Entity Provisioned
	- RUN_STATE_SUSPECT: Entity Suspect
	- RUN_STATE_DOWNLOADING: Edge-node downloading entity artifacts
	- RUN_STATE_RESTARTING: Entity Restarting
	- RUN_STATE_PURGING: Entity Purging
	- RUN_STATE_HALTING: Entity Halting
	- RUN_STATE_ERROR: Entity encountered an error
	- RUN_STATE_VERIFYING: Verification of downloaded Artifacts in Progress.
	- RUN_STATE_LOADING: Loading of Artifacts into local datastore in Progress.
	- RUN_STATE_CREATING_VOLUME: Volume creation from artifacts in Progress
	- RUN_STATE_BOOTING: Entity booting up
	- RUN_STATE_MAINTENANCE_MODE: Entity maintenance mode
	- RUN_STATE_START_DELAYED: Application start delayed as per configuration
	- RUN_STATE_BASEOS_UPDATING: Device BaseOs Update in Progress
	- RUN_STATE_PREPARING_POWEROFF: device preparing power off - shutting down all app instances
	- RUN_STATE_POWERING_OFF: device powering off from local profile server
	- RUN_STATE_PREPARED_POWEROFF: device prepared power off - all app instances are shut down

	    Default: "RUN_STATE_UNSPECIFIED"
	*/
	FilterRunState *string

	// FilterType.
	//
	// Default: "VOLUME_INSTANCE_TYPE_UNSPECIFIED"
	FilterType *string

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

	// Summary.
	Summary *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the volume instance status query volume instance status config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams) WithDefaults() *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the volume instance status query volume instance status config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams) SetDefaults() {
	var (
		filterRunStateDefault = string("RUN_STATE_UNSPECIFIED")

		filterTypeDefault = string("VOLUME_INSTANCE_TYPE_UNSPECIFIED")
	)

	val := VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams{
		FilterRunState: &filterRunStateDefault,
		FilterType:     &filterTypeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the volume instance status query volume instance status config params
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams) WithTimeout(timeout time.Duration) *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the volume instance status query volume instance status config params
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the volume instance status query volume instance status config params
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams) WithContext(ctx context.Context) *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the volume instance status query volume instance status config params
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the volume instance status query volume instance status config params
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams) WithHTTPClient(client *http.Client) *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the volume instance status query volume instance status config params
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the volume instance status query volume instance status config params
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams) WithXRequestID(xRequestID *string) *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the volume instance status query volume instance status config params
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithFields adds the fields to the volume instance status query volume instance status config params
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams) WithFields(fields []string) *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the volume instance status query volume instance status config params
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithFilterAppInstName adds the filterAppInstName to the volume instance status query volume instance status config params
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams) WithFilterAppInstName(filterAppInstName *string) *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams {
	o.SetFilterAppInstName(filterAppInstName)
	return o
}

// SetFilterAppInstName adds the filterAppInstName to the volume instance status query volume instance status config params
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams) SetFilterAppInstName(filterAppInstName *string) {
	o.FilterAppInstName = filterAppInstName
}

// WithFilterDeviceName adds the filterDeviceName to the volume instance status query volume instance status config params
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams) WithFilterDeviceName(filterDeviceName *string) *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams {
	o.SetFilterDeviceName(filterDeviceName)
	return o
}

// SetFilterDeviceName adds the filterDeviceName to the volume instance status query volume instance status config params
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams) SetFilterDeviceName(filterDeviceName *string) {
	o.FilterDeviceName = filterDeviceName
}

// WithFilterDeviceNamePattern adds the filterDeviceNamePattern to the volume instance status query volume instance status config params
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams) WithFilterDeviceNamePattern(filterDeviceNamePattern *string) *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams {
	o.SetFilterDeviceNamePattern(filterDeviceNamePattern)
	return o
}

// SetFilterDeviceNamePattern adds the filterDeviceNamePattern to the volume instance status query volume instance status config params
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams) SetFilterDeviceNamePattern(filterDeviceNamePattern *string) {
	o.FilterDeviceNamePattern = filterDeviceNamePattern
}

// WithFilterImageName adds the filterImageName to the volume instance status query volume instance status config params
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams) WithFilterImageName(filterImageName *string) *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams {
	o.SetFilterImageName(filterImageName)
	return o
}

// SetFilterImageName adds the filterImageName to the volume instance status query volume instance status config params
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams) SetFilterImageName(filterImageName *string) {
	o.FilterImageName = filterImageName
}

// WithFilterNamePattern adds the filterNamePattern to the volume instance status query volume instance status config params
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams) WithFilterNamePattern(filterNamePattern *string) *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams {
	o.SetFilterNamePattern(filterNamePattern)
	return o
}

// SetFilterNamePattern adds the filterNamePattern to the volume instance status query volume instance status config params
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams) SetFilterNamePattern(filterNamePattern *string) {
	o.FilterNamePattern = filterNamePattern
}

// WithFilterProjectName adds the filterProjectName to the volume instance status query volume instance status config params
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams) WithFilterProjectName(filterProjectName *string) *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams {
	o.SetFilterProjectName(filterProjectName)
	return o
}

// SetFilterProjectName adds the filterProjectName to the volume instance status query volume instance status config params
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams) SetFilterProjectName(filterProjectName *string) {
	o.FilterProjectName = filterProjectName
}

// WithFilterProjectNamePattern adds the filterProjectNamePattern to the volume instance status query volume instance status config params
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams) WithFilterProjectNamePattern(filterProjectNamePattern *string) *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams {
	o.SetFilterProjectNamePattern(filterProjectNamePattern)
	return o
}

// SetFilterProjectNamePattern adds the filterProjectNamePattern to the volume instance status query volume instance status config params
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams) SetFilterProjectNamePattern(filterProjectNamePattern *string) {
	o.FilterProjectNamePattern = filterProjectNamePattern
}

// WithFilterRunState adds the filterRunState to the volume instance status query volume instance status config params
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams) WithFilterRunState(filterRunState *string) *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams {
	o.SetFilterRunState(filterRunState)
	return o
}

// SetFilterRunState adds the filterRunState to the volume instance status query volume instance status config params
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams) SetFilterRunState(filterRunState *string) {
	o.FilterRunState = filterRunState
}

// WithFilterType adds the filterType to the volume instance status query volume instance status config params
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams) WithFilterType(filterType *string) *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams {
	o.SetFilterType(filterType)
	return o
}

// SetFilterType adds the filterType to the volume instance status query volume instance status config params
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams) SetFilterType(filterType *string) {
	o.FilterType = filterType
}

// WithNextOrderBy adds the nextOrderBy to the volume instance status query volume instance status config params
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams) WithNextOrderBy(nextOrderBy []string) *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams {
	o.SetNextOrderBy(nextOrderBy)
	return o
}

// SetNextOrderBy adds the nextOrderBy to the volume instance status query volume instance status config params
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams) SetNextOrderBy(nextOrderBy []string) {
	o.NextOrderBy = nextOrderBy
}

// WithNextPageNum adds the nextPageNum to the volume instance status query volume instance status config params
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams) WithNextPageNum(nextPageNum *int64) *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams {
	o.SetNextPageNum(nextPageNum)
	return o
}

// SetNextPageNum adds the nextPageNum to the volume instance status query volume instance status config params
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams) SetNextPageNum(nextPageNum *int64) {
	o.NextPageNum = nextPageNum
}

// WithNextPageSize adds the nextPageSize to the volume instance status query volume instance status config params
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams) WithNextPageSize(nextPageSize *int64) *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams {
	o.SetNextPageSize(nextPageSize)
	return o
}

// SetNextPageSize adds the nextPageSize to the volume instance status query volume instance status config params
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams) SetNextPageSize(nextPageSize *int64) {
	o.NextPageSize = nextPageSize
}

// WithNextPageToken adds the nextPageToken to the volume instance status query volume instance status config params
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams) WithNextPageToken(nextPageToken *string) *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams {
	o.SetNextPageToken(nextPageToken)
	return o
}

// SetNextPageToken adds the nextPageToken to the volume instance status query volume instance status config params
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams) SetNextPageToken(nextPageToken *string) {
	o.NextPageToken = nextPageToken
}

// WithNextTotalPages adds the nextTotalPages to the volume instance status query volume instance status config params
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams) WithNextTotalPages(nextTotalPages *int64) *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams {
	o.SetNextTotalPages(nextTotalPages)
	return o
}

// SetNextTotalPages adds the nextTotalPages to the volume instance status query volume instance status config params
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams) SetNextTotalPages(nextTotalPages *int64) {
	o.NextTotalPages = nextTotalPages
}

// WithSummary adds the summary to the volume instance status query volume instance status config params
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams) WithSummary(summary *bool) *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams {
	o.SetSummary(summary)
	return o
}

// SetSummary adds the summary to the volume instance status query volume instance status config params
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams) SetSummary(summary *bool) {
	o.Summary = summary
}

// WriteToRequest writes these params to a swagger request
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.FilterAppInstName != nil {

		// query param filter.appInstName
		var qrFilterAppInstName string

		if o.FilterAppInstName != nil {
			qrFilterAppInstName = *o.FilterAppInstName
		}
		qFilterAppInstName := qrFilterAppInstName
		if qFilterAppInstName != "" {

			if err := r.SetQueryParam("filter.appInstName", qFilterAppInstName); err != nil {
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

	if o.FilterImageName != nil {

		// query param filter.imageName
		var qrFilterImageName string

		if o.FilterImageName != nil {
			qrFilterImageName = *o.FilterImageName
		}
		qFilterImageName := qrFilterImageName
		if qFilterImageName != "" {

			if err := r.SetQueryParam("filter.imageName", qFilterImageName); err != nil {
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

	if o.FilterRunState != nil {

		// query param filter.runState
		var qrFilterRunState string

		if o.FilterRunState != nil {
			qrFilterRunState = *o.FilterRunState
		}
		qFilterRunState := qrFilterRunState
		if qFilterRunState != "" {

			if err := r.SetQueryParam("filter.runState", qFilterRunState); err != nil {
				return err
			}
		}
	}

	if o.FilterType != nil {

		// query param filter.type
		var qrFilterType string

		if o.FilterType != nil {
			qrFilterType = *o.FilterType
		}
		qFilterType := qrFilterType
		if qFilterType != "" {

			if err := r.SetQueryParam("filter.type", qFilterType); err != nil {
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

// bindParamVolumeInstanceStatusQueryVolumeInstanceStatusConfig binds the parameter fields
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamVolumeInstanceStatusQueryVolumeInstanceStatusConfig binds the parameter next.orderBy
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusConfigParams) bindParamNextOrderBy(formats strfmt.Registry) []string {
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
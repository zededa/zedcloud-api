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
	"github.com/go-openapi/swag"
)

// NewQueryEdgeNodeStatusParams creates a new QueryEdgeNodeStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQueryEdgeNodeStatusParams() *QueryEdgeNodeStatusParams {
	return &QueryEdgeNodeStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQueryEdgeNodeStatusParamsWithTimeout creates a new QueryEdgeNodeStatusParams object
// with the ability to set a timeout on a request.
func NewQueryEdgeNodeStatusParamsWithTimeout(timeout time.Duration) *QueryEdgeNodeStatusParams {
	return &QueryEdgeNodeStatusParams{
		timeout: timeout,
	}
}

// NewQueryEdgeNodeStatusParamsWithContext creates a new QueryEdgeNodeStatusParams object
// with the ability to set a context for a request.
func NewQueryEdgeNodeStatusParamsWithContext(ctx context.Context) *QueryEdgeNodeStatusParams {
	return &QueryEdgeNodeStatusParams{
		Context: ctx,
	}
}

// NewQueryEdgeNodeStatusParamsWithHTTPClient creates a new QueryEdgeNodeStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewQueryEdgeNodeStatusParamsWithHTTPClient(client *http.Client) *QueryEdgeNodeStatusParams {
	return &QueryEdgeNodeStatusParams{
		HTTPClient: client,
	}
}

/* QueryEdgeNodeStatusParams contains all the parameters to send to the API endpoint
   for the query edge node status operation.

   Typically these are written to a http.Request.
*/
type QueryEdgeNodeStatusParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// FilterLoad.
	//
	// Default: "DEVICE_LOAD_UNSPECIFIED"
	FilterLoad *string

	// FilterNamePattern.
	FilterNamePattern *string

	// FilterProjectName.
	FilterProjectName *string

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

	    Default: "RUN_STATE_UNSPECIFIED"
	*/
	FilterRunState *string

	/* NextOrderBy.

	   OrderBy helps in sorting the list response
	*/
	NextOrderBy *string

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
	//
	// Format: boolean
	Summary *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the query edge node status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryEdgeNodeStatusParams) WithDefaults() *QueryEdgeNodeStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the query edge node status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryEdgeNodeStatusParams) SetDefaults() {
	var (
		filterLoadDefault = string("DEVICE_LOAD_UNSPECIFIED")

		filterRunStateDefault = string("RUN_STATE_UNSPECIFIED")
	)

	val := QueryEdgeNodeStatusParams{
		FilterLoad:     &filterLoadDefault,
		FilterRunState: &filterRunStateDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the query edge node status params
func (o *QueryEdgeNodeStatusParams) WithTimeout(timeout time.Duration) *QueryEdgeNodeStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query edge node status params
func (o *QueryEdgeNodeStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query edge node status params
func (o *QueryEdgeNodeStatusParams) WithContext(ctx context.Context) *QueryEdgeNodeStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query edge node status params
func (o *QueryEdgeNodeStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query edge node status params
func (o *QueryEdgeNodeStatusParams) WithHTTPClient(client *http.Client) *QueryEdgeNodeStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query edge node status params
func (o *QueryEdgeNodeStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the query edge node status params
func (o *QueryEdgeNodeStatusParams) WithXRequestID(xRequestID *string) *QueryEdgeNodeStatusParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the query edge node status params
func (o *QueryEdgeNodeStatusParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithFilterLoad adds the filterLoad to the query edge node status params
func (o *QueryEdgeNodeStatusParams) WithFilterLoad(filterLoad *string) *QueryEdgeNodeStatusParams {
	o.SetFilterLoad(filterLoad)
	return o
}

// SetFilterLoad adds the filterLoad to the query edge node status params
func (o *QueryEdgeNodeStatusParams) SetFilterLoad(filterLoad *string) {
	o.FilterLoad = filterLoad
}

// WithFilterNamePattern adds the filterNamePattern to the query edge node status params
func (o *QueryEdgeNodeStatusParams) WithFilterNamePattern(filterNamePattern *string) *QueryEdgeNodeStatusParams {
	o.SetFilterNamePattern(filterNamePattern)
	return o
}

// SetFilterNamePattern adds the filterNamePattern to the query edge node status params
func (o *QueryEdgeNodeStatusParams) SetFilterNamePattern(filterNamePattern *string) {
	o.FilterNamePattern = filterNamePattern
}

// WithFilterProjectName adds the filterProjectName to the query edge node status params
func (o *QueryEdgeNodeStatusParams) WithFilterProjectName(filterProjectName *string) *QueryEdgeNodeStatusParams {
	o.SetFilterProjectName(filterProjectName)
	return o
}

// SetFilterProjectName adds the filterProjectName to the query edge node status params
func (o *QueryEdgeNodeStatusParams) SetFilterProjectName(filterProjectName *string) {
	o.FilterProjectName = filterProjectName
}

// WithFilterRunState adds the filterRunState to the query edge node status params
func (o *QueryEdgeNodeStatusParams) WithFilterRunState(filterRunState *string) *QueryEdgeNodeStatusParams {
	o.SetFilterRunState(filterRunState)
	return o
}

// SetFilterRunState adds the filterRunState to the query edge node status params
func (o *QueryEdgeNodeStatusParams) SetFilterRunState(filterRunState *string) {
	o.FilterRunState = filterRunState
}

// WithNextOrderBy adds the nextOrderBy to the query edge node status params
func (o *QueryEdgeNodeStatusParams) WithNextOrderBy(nextOrderBy *string) *QueryEdgeNodeStatusParams {
	o.SetNextOrderBy(nextOrderBy)
	return o
}

// SetNextOrderBy adds the nextOrderBy to the query edge node status params
func (o *QueryEdgeNodeStatusParams) SetNextOrderBy(nextOrderBy *string) {
	o.NextOrderBy = nextOrderBy
}

// WithNextPageNum adds the nextPageNum to the query edge node status params
func (o *QueryEdgeNodeStatusParams) WithNextPageNum(nextPageNum *int64) *QueryEdgeNodeStatusParams {
	o.SetNextPageNum(nextPageNum)
	return o
}

// SetNextPageNum adds the nextPageNum to the query edge node status params
func (o *QueryEdgeNodeStatusParams) SetNextPageNum(nextPageNum *int64) {
	o.NextPageNum = nextPageNum
}

// WithNextPageSize adds the nextPageSize to the query edge node status params
func (o *QueryEdgeNodeStatusParams) WithNextPageSize(nextPageSize *int64) *QueryEdgeNodeStatusParams {
	o.SetNextPageSize(nextPageSize)
	return o
}

// SetNextPageSize adds the nextPageSize to the query edge node status params
func (o *QueryEdgeNodeStatusParams) SetNextPageSize(nextPageSize *int64) {
	o.NextPageSize = nextPageSize
}

// WithNextPageToken adds the nextPageToken to the query edge node status params
func (o *QueryEdgeNodeStatusParams) WithNextPageToken(nextPageToken *string) *QueryEdgeNodeStatusParams {
	o.SetNextPageToken(nextPageToken)
	return o
}

// SetNextPageToken adds the nextPageToken to the query edge node status params
func (o *QueryEdgeNodeStatusParams) SetNextPageToken(nextPageToken *string) {
	o.NextPageToken = nextPageToken
}

// WithNextTotalPages adds the nextTotalPages to the query edge node status params
func (o *QueryEdgeNodeStatusParams) WithNextTotalPages(nextTotalPages *int64) *QueryEdgeNodeStatusParams {
	o.SetNextTotalPages(nextTotalPages)
	return o
}

// SetNextTotalPages adds the nextTotalPages to the query edge node status params
func (o *QueryEdgeNodeStatusParams) SetNextTotalPages(nextTotalPages *int64) {
	o.NextTotalPages = nextTotalPages
}

// WithSummary adds the summary to the query edge node status params
func (o *QueryEdgeNodeStatusParams) WithSummary(summary *bool) *QueryEdgeNodeStatusParams {
	o.SetSummary(summary)
	return o
}

// SetSummary adds the summary to the query edge node status params
func (o *QueryEdgeNodeStatusParams) SetSummary(summary *bool) {
	o.Summary = summary
}

// WriteToRequest writes these params to a swagger request
func (o *QueryEdgeNodeStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.FilterLoad != nil {

		// query param filter.load
		var qrFilterLoad string

		if o.FilterLoad != nil {
			qrFilterLoad = *o.FilterLoad
		}
		qFilterLoad := qrFilterLoad
		if qFilterLoad != "" {

			if err := r.SetQueryParam("filter.load", qFilterLoad); err != nil {
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

	if o.NextOrderBy != nil {

		// query param next.orderBy
		var qrNextOrderBy string

		if o.NextOrderBy != nil {
			qrNextOrderBy = *o.NextOrderBy
		}
		qNextOrderBy := qrNextOrderBy
		if qNextOrderBy != "" {

			if err := r.SetQueryParam("next.orderBy", qNextOrderBy); err != nil {
				return err
			}
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

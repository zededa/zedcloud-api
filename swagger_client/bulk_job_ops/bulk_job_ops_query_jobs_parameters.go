// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package bulk_job_ops

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

// NewBulkJobOpsQueryJobsParams creates a new BulkJobOpsQueryJobsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBulkJobOpsQueryJobsParams() *BulkJobOpsQueryJobsParams {
	return &BulkJobOpsQueryJobsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBulkJobOpsQueryJobsParamsWithTimeout creates a new BulkJobOpsQueryJobsParams object
// with the ability to set a timeout on a request.
func NewBulkJobOpsQueryJobsParamsWithTimeout(timeout time.Duration) *BulkJobOpsQueryJobsParams {
	return &BulkJobOpsQueryJobsParams{
		timeout: timeout,
	}
}

// NewBulkJobOpsQueryJobsParamsWithContext creates a new BulkJobOpsQueryJobsParams object
// with the ability to set a context for a request.
func NewBulkJobOpsQueryJobsParamsWithContext(ctx context.Context) *BulkJobOpsQueryJobsParams {
	return &BulkJobOpsQueryJobsParams{
		Context: ctx,
	}
}

// NewBulkJobOpsQueryJobsParamsWithHTTPClient creates a new BulkJobOpsQueryJobsParams object
// with the ability to set a custom HTTPClient for a request.
func NewBulkJobOpsQueryJobsParamsWithHTTPClient(client *http.Client) *BulkJobOpsQueryJobsParams {
	return &BulkJobOpsQueryJobsParams{
		HTTPClient: client,
	}
}

/* BulkJobOpsQueryJobsParams contains all the parameters to send to the API endpoint
   for the bulk job ops query jobs operation.

   Typically these are written to a http.Request.
*/
type BulkJobOpsQueryJobsParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* FilterJobStatus.

	   status of the job.

	   Default: "JOB_STATUS_UNSPECIFIED"
	*/
	FilterJobStatus *string

	/* FilterNamePattern.

	   name pattern of the job.
	*/
	FilterNamePattern *string

	/* FilterProjectName.

	   project name.
	*/
	FilterProjectName *string

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

	   Only summary count is required.
	*/
	Summary *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the bulk job ops query jobs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BulkJobOpsQueryJobsParams) WithDefaults() *BulkJobOpsQueryJobsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the bulk job ops query jobs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BulkJobOpsQueryJobsParams) SetDefaults() {
	var (
		filterJobStatusDefault = string("JOB_STATUS_UNSPECIFIED")
	)

	val := BulkJobOpsQueryJobsParams{
		FilterJobStatus: &filterJobStatusDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the bulk job ops query jobs params
func (o *BulkJobOpsQueryJobsParams) WithTimeout(timeout time.Duration) *BulkJobOpsQueryJobsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bulk job ops query jobs params
func (o *BulkJobOpsQueryJobsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bulk job ops query jobs params
func (o *BulkJobOpsQueryJobsParams) WithContext(ctx context.Context) *BulkJobOpsQueryJobsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bulk job ops query jobs params
func (o *BulkJobOpsQueryJobsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bulk job ops query jobs params
func (o *BulkJobOpsQueryJobsParams) WithHTTPClient(client *http.Client) *BulkJobOpsQueryJobsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bulk job ops query jobs params
func (o *BulkJobOpsQueryJobsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the bulk job ops query jobs params
func (o *BulkJobOpsQueryJobsParams) WithXRequestID(xRequestID *string) *BulkJobOpsQueryJobsParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the bulk job ops query jobs params
func (o *BulkJobOpsQueryJobsParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithFilterJobStatus adds the filterJobStatus to the bulk job ops query jobs params
func (o *BulkJobOpsQueryJobsParams) WithFilterJobStatus(filterJobStatus *string) *BulkJobOpsQueryJobsParams {
	o.SetFilterJobStatus(filterJobStatus)
	return o
}

// SetFilterJobStatus adds the filterJobStatus to the bulk job ops query jobs params
func (o *BulkJobOpsQueryJobsParams) SetFilterJobStatus(filterJobStatus *string) {
	o.FilterJobStatus = filterJobStatus
}

// WithFilterNamePattern adds the filterNamePattern to the bulk job ops query jobs params
func (o *BulkJobOpsQueryJobsParams) WithFilterNamePattern(filterNamePattern *string) *BulkJobOpsQueryJobsParams {
	o.SetFilterNamePattern(filterNamePattern)
	return o
}

// SetFilterNamePattern adds the filterNamePattern to the bulk job ops query jobs params
func (o *BulkJobOpsQueryJobsParams) SetFilterNamePattern(filterNamePattern *string) {
	o.FilterNamePattern = filterNamePattern
}

// WithFilterProjectName adds the filterProjectName to the bulk job ops query jobs params
func (o *BulkJobOpsQueryJobsParams) WithFilterProjectName(filterProjectName *string) *BulkJobOpsQueryJobsParams {
	o.SetFilterProjectName(filterProjectName)
	return o
}

// SetFilterProjectName adds the filterProjectName to the bulk job ops query jobs params
func (o *BulkJobOpsQueryJobsParams) SetFilterProjectName(filterProjectName *string) {
	o.FilterProjectName = filterProjectName
}

// WithNextOrderBy adds the nextOrderBy to the bulk job ops query jobs params
func (o *BulkJobOpsQueryJobsParams) WithNextOrderBy(nextOrderBy []string) *BulkJobOpsQueryJobsParams {
	o.SetNextOrderBy(nextOrderBy)
	return o
}

// SetNextOrderBy adds the nextOrderBy to the bulk job ops query jobs params
func (o *BulkJobOpsQueryJobsParams) SetNextOrderBy(nextOrderBy []string) {
	o.NextOrderBy = nextOrderBy
}

// WithNextPageNum adds the nextPageNum to the bulk job ops query jobs params
func (o *BulkJobOpsQueryJobsParams) WithNextPageNum(nextPageNum *int64) *BulkJobOpsQueryJobsParams {
	o.SetNextPageNum(nextPageNum)
	return o
}

// SetNextPageNum adds the nextPageNum to the bulk job ops query jobs params
func (o *BulkJobOpsQueryJobsParams) SetNextPageNum(nextPageNum *int64) {
	o.NextPageNum = nextPageNum
}

// WithNextPageSize adds the nextPageSize to the bulk job ops query jobs params
func (o *BulkJobOpsQueryJobsParams) WithNextPageSize(nextPageSize *int64) *BulkJobOpsQueryJobsParams {
	o.SetNextPageSize(nextPageSize)
	return o
}

// SetNextPageSize adds the nextPageSize to the bulk job ops query jobs params
func (o *BulkJobOpsQueryJobsParams) SetNextPageSize(nextPageSize *int64) {
	o.NextPageSize = nextPageSize
}

// WithNextPageToken adds the nextPageToken to the bulk job ops query jobs params
func (o *BulkJobOpsQueryJobsParams) WithNextPageToken(nextPageToken *string) *BulkJobOpsQueryJobsParams {
	o.SetNextPageToken(nextPageToken)
	return o
}

// SetNextPageToken adds the nextPageToken to the bulk job ops query jobs params
func (o *BulkJobOpsQueryJobsParams) SetNextPageToken(nextPageToken *string) {
	o.NextPageToken = nextPageToken
}

// WithNextTotalPages adds the nextTotalPages to the bulk job ops query jobs params
func (o *BulkJobOpsQueryJobsParams) WithNextTotalPages(nextTotalPages *int64) *BulkJobOpsQueryJobsParams {
	o.SetNextTotalPages(nextTotalPages)
	return o
}

// SetNextTotalPages adds the nextTotalPages to the bulk job ops query jobs params
func (o *BulkJobOpsQueryJobsParams) SetNextTotalPages(nextTotalPages *int64) {
	o.NextTotalPages = nextTotalPages
}

// WithSummary adds the summary to the bulk job ops query jobs params
func (o *BulkJobOpsQueryJobsParams) WithSummary(summary *bool) *BulkJobOpsQueryJobsParams {
	o.SetSummary(summary)
	return o
}

// SetSummary adds the summary to the bulk job ops query jobs params
func (o *BulkJobOpsQueryJobsParams) SetSummary(summary *bool) {
	o.Summary = summary
}

// WriteToRequest writes these params to a swagger request
func (o *BulkJobOpsQueryJobsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.FilterJobStatus != nil {

		// query param filter.jobStatus
		var qrFilterJobStatus string

		if o.FilterJobStatus != nil {
			qrFilterJobStatus = *o.FilterJobStatus
		}
		qFilterJobStatus := qrFilterJobStatus
		if qFilterJobStatus != "" {

			if err := r.SetQueryParam("filter.jobStatus", qFilterJobStatus); err != nil {
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

// bindParamBulkJobOpsQueryJobs binds the parameter next.orderBy
func (o *BulkJobOpsQueryJobsParams) bindParamNextOrderBy(formats strfmt.Registry) []string {
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

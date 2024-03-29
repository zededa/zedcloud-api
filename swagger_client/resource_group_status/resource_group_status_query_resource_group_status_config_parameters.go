// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package resource_group_status

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

// NewResourceGroupStatusQueryResourceGroupStatusConfigParams creates a new ResourceGroupStatusQueryResourceGroupStatusConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewResourceGroupStatusQueryResourceGroupStatusConfigParams() *ResourceGroupStatusQueryResourceGroupStatusConfigParams {
	return &ResourceGroupStatusQueryResourceGroupStatusConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewResourceGroupStatusQueryResourceGroupStatusConfigParamsWithTimeout creates a new ResourceGroupStatusQueryResourceGroupStatusConfigParams object
// with the ability to set a timeout on a request.
func NewResourceGroupStatusQueryResourceGroupStatusConfigParamsWithTimeout(timeout time.Duration) *ResourceGroupStatusQueryResourceGroupStatusConfigParams {
	return &ResourceGroupStatusQueryResourceGroupStatusConfigParams{
		timeout: timeout,
	}
}

// NewResourceGroupStatusQueryResourceGroupStatusConfigParamsWithContext creates a new ResourceGroupStatusQueryResourceGroupStatusConfigParams object
// with the ability to set a context for a request.
func NewResourceGroupStatusQueryResourceGroupStatusConfigParamsWithContext(ctx context.Context) *ResourceGroupStatusQueryResourceGroupStatusConfigParams {
	return &ResourceGroupStatusQueryResourceGroupStatusConfigParams{
		Context: ctx,
	}
}

// NewResourceGroupStatusQueryResourceGroupStatusConfigParamsWithHTTPClient creates a new ResourceGroupStatusQueryResourceGroupStatusConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewResourceGroupStatusQueryResourceGroupStatusConfigParamsWithHTTPClient(client *http.Client) *ResourceGroupStatusQueryResourceGroupStatusConfigParams {
	return &ResourceGroupStatusQueryResourceGroupStatusConfigParams{
		HTTPClient: client,
	}
}

/*
ResourceGroupStatusQueryResourceGroupStatusConfigParams contains all the parameters to send to the API endpoint

	for the resource group status query resource group status config operation.

	Typically these are written to a http.Request.
*/
type ResourceGroupStatusQueryResourceGroupStatusConfigParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* Fields.

	   Fields to select: id, name, status
	*/
	Fields []string

	/* FilterNamePattern.

	   Resource group name pattern to be matched.
	*/
	FilterNamePattern *string

	/* FilterStatus.

	    Resource group status to be matched.

	- TAG_STATUS_UNSPECIFIED: Tag Status : UNSPECIFIED
	- TAG_STATUS_ARCHIVE: Tag Status : ARCHIVE
	- TAG_STATUS_ACTIVE: Tag Status : ACTIVE
	- TAG_STATUS_INACTIVE: Tag Status : INACTIVE

	    Default: "TAG_STATUS_UNSPECIFIED"
	*/
	FilterStatus *string

	/* FilterType.

	    Resource group type to ne matched.

	- TAG_TYPE_UNSPECIFIED: Unspecified
	- TAG_TYPE_GENERIC: Generic resource group
	- TAG_TYPE_PROJECT: Project resource group
	- TAG_TYPE_AZURE: Project resource group
	- TAG_TYPE_DEPLOYMENT: Project with deployment

	    Default: "TAG_TYPE_UNSPECIFIED"
	*/
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

	/* Summary.

	   Only summary of the records required
	*/
	Summary *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the resource group status query resource group status config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResourceGroupStatusQueryResourceGroupStatusConfigParams) WithDefaults() *ResourceGroupStatusQueryResourceGroupStatusConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the resource group status query resource group status config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResourceGroupStatusQueryResourceGroupStatusConfigParams) SetDefaults() {
	var (
		filterStatusDefault = string("TAG_STATUS_UNSPECIFIED")

		filterTypeDefault = string("TAG_TYPE_UNSPECIFIED")
	)

	val := ResourceGroupStatusQueryResourceGroupStatusConfigParams{
		FilterStatus: &filterStatusDefault,
		FilterType:   &filterTypeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the resource group status query resource group status config params
func (o *ResourceGroupStatusQueryResourceGroupStatusConfigParams) WithTimeout(timeout time.Duration) *ResourceGroupStatusQueryResourceGroupStatusConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the resource group status query resource group status config params
func (o *ResourceGroupStatusQueryResourceGroupStatusConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the resource group status query resource group status config params
func (o *ResourceGroupStatusQueryResourceGroupStatusConfigParams) WithContext(ctx context.Context) *ResourceGroupStatusQueryResourceGroupStatusConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the resource group status query resource group status config params
func (o *ResourceGroupStatusQueryResourceGroupStatusConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the resource group status query resource group status config params
func (o *ResourceGroupStatusQueryResourceGroupStatusConfigParams) WithHTTPClient(client *http.Client) *ResourceGroupStatusQueryResourceGroupStatusConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the resource group status query resource group status config params
func (o *ResourceGroupStatusQueryResourceGroupStatusConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the resource group status query resource group status config params
func (o *ResourceGroupStatusQueryResourceGroupStatusConfigParams) WithXRequestID(xRequestID *string) *ResourceGroupStatusQueryResourceGroupStatusConfigParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the resource group status query resource group status config params
func (o *ResourceGroupStatusQueryResourceGroupStatusConfigParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithFields adds the fields to the resource group status query resource group status config params
func (o *ResourceGroupStatusQueryResourceGroupStatusConfigParams) WithFields(fields []string) *ResourceGroupStatusQueryResourceGroupStatusConfigParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the resource group status query resource group status config params
func (o *ResourceGroupStatusQueryResourceGroupStatusConfigParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithFilterNamePattern adds the filterNamePattern to the resource group status query resource group status config params
func (o *ResourceGroupStatusQueryResourceGroupStatusConfigParams) WithFilterNamePattern(filterNamePattern *string) *ResourceGroupStatusQueryResourceGroupStatusConfigParams {
	o.SetFilterNamePattern(filterNamePattern)
	return o
}

// SetFilterNamePattern adds the filterNamePattern to the resource group status query resource group status config params
func (o *ResourceGroupStatusQueryResourceGroupStatusConfigParams) SetFilterNamePattern(filterNamePattern *string) {
	o.FilterNamePattern = filterNamePattern
}

// WithFilterStatus adds the filterStatus to the resource group status query resource group status config params
func (o *ResourceGroupStatusQueryResourceGroupStatusConfigParams) WithFilterStatus(filterStatus *string) *ResourceGroupStatusQueryResourceGroupStatusConfigParams {
	o.SetFilterStatus(filterStatus)
	return o
}

// SetFilterStatus adds the filterStatus to the resource group status query resource group status config params
func (o *ResourceGroupStatusQueryResourceGroupStatusConfigParams) SetFilterStatus(filterStatus *string) {
	o.FilterStatus = filterStatus
}

// WithFilterType adds the filterType to the resource group status query resource group status config params
func (o *ResourceGroupStatusQueryResourceGroupStatusConfigParams) WithFilterType(filterType *string) *ResourceGroupStatusQueryResourceGroupStatusConfigParams {
	o.SetFilterType(filterType)
	return o
}

// SetFilterType adds the filterType to the resource group status query resource group status config params
func (o *ResourceGroupStatusQueryResourceGroupStatusConfigParams) SetFilterType(filterType *string) {
	o.FilterType = filterType
}

// WithNextOrderBy adds the nextOrderBy to the resource group status query resource group status config params
func (o *ResourceGroupStatusQueryResourceGroupStatusConfigParams) WithNextOrderBy(nextOrderBy []string) *ResourceGroupStatusQueryResourceGroupStatusConfigParams {
	o.SetNextOrderBy(nextOrderBy)
	return o
}

// SetNextOrderBy adds the nextOrderBy to the resource group status query resource group status config params
func (o *ResourceGroupStatusQueryResourceGroupStatusConfigParams) SetNextOrderBy(nextOrderBy []string) {
	o.NextOrderBy = nextOrderBy
}

// WithNextPageNum adds the nextPageNum to the resource group status query resource group status config params
func (o *ResourceGroupStatusQueryResourceGroupStatusConfigParams) WithNextPageNum(nextPageNum *int64) *ResourceGroupStatusQueryResourceGroupStatusConfigParams {
	o.SetNextPageNum(nextPageNum)
	return o
}

// SetNextPageNum adds the nextPageNum to the resource group status query resource group status config params
func (o *ResourceGroupStatusQueryResourceGroupStatusConfigParams) SetNextPageNum(nextPageNum *int64) {
	o.NextPageNum = nextPageNum
}

// WithNextPageSize adds the nextPageSize to the resource group status query resource group status config params
func (o *ResourceGroupStatusQueryResourceGroupStatusConfigParams) WithNextPageSize(nextPageSize *int64) *ResourceGroupStatusQueryResourceGroupStatusConfigParams {
	o.SetNextPageSize(nextPageSize)
	return o
}

// SetNextPageSize adds the nextPageSize to the resource group status query resource group status config params
func (o *ResourceGroupStatusQueryResourceGroupStatusConfigParams) SetNextPageSize(nextPageSize *int64) {
	o.NextPageSize = nextPageSize
}

// WithNextPageToken adds the nextPageToken to the resource group status query resource group status config params
func (o *ResourceGroupStatusQueryResourceGroupStatusConfigParams) WithNextPageToken(nextPageToken *string) *ResourceGroupStatusQueryResourceGroupStatusConfigParams {
	o.SetNextPageToken(nextPageToken)
	return o
}

// SetNextPageToken adds the nextPageToken to the resource group status query resource group status config params
func (o *ResourceGroupStatusQueryResourceGroupStatusConfigParams) SetNextPageToken(nextPageToken *string) {
	o.NextPageToken = nextPageToken
}

// WithNextTotalPages adds the nextTotalPages to the resource group status query resource group status config params
func (o *ResourceGroupStatusQueryResourceGroupStatusConfigParams) WithNextTotalPages(nextTotalPages *int64) *ResourceGroupStatusQueryResourceGroupStatusConfigParams {
	o.SetNextTotalPages(nextTotalPages)
	return o
}

// SetNextTotalPages adds the nextTotalPages to the resource group status query resource group status config params
func (o *ResourceGroupStatusQueryResourceGroupStatusConfigParams) SetNextTotalPages(nextTotalPages *int64) {
	o.NextTotalPages = nextTotalPages
}

// WithSummary adds the summary to the resource group status query resource group status config params
func (o *ResourceGroupStatusQueryResourceGroupStatusConfigParams) WithSummary(summary *bool) *ResourceGroupStatusQueryResourceGroupStatusConfigParams {
	o.SetSummary(summary)
	return o
}

// SetSummary adds the summary to the resource group status query resource group status config params
func (o *ResourceGroupStatusQueryResourceGroupStatusConfigParams) SetSummary(summary *bool) {
	o.Summary = summary
}

// WriteToRequest writes these params to a swagger request
func (o *ResourceGroupStatusQueryResourceGroupStatusConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.FilterStatus != nil {

		// query param filter.status
		var qrFilterStatus string

		if o.FilterStatus != nil {
			qrFilterStatus = *o.FilterStatus
		}
		qFilterStatus := qrFilterStatus
		if qFilterStatus != "" {

			if err := r.SetQueryParam("filter.status", qFilterStatus); err != nil {
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

// bindParamResourceGroupStatusQueryResourceGroupStatusConfig binds the parameter fields
func (o *ResourceGroupStatusQueryResourceGroupStatusConfigParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamResourceGroupStatusQueryResourceGroupStatusConfig binds the parameter next.orderBy
func (o *ResourceGroupStatusQueryResourceGroupStatusConfigParams) bindParamNextOrderBy(formats strfmt.Registry) []string {
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

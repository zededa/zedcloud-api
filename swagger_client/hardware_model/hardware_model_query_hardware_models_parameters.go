// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package hardware_model

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

// NewHardwareModelQueryHardwareModelsParams creates a new HardwareModelQueryHardwareModelsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHardwareModelQueryHardwareModelsParams() *HardwareModelQueryHardwareModelsParams {
	return &HardwareModelQueryHardwareModelsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHardwareModelQueryHardwareModelsParamsWithTimeout creates a new HardwareModelQueryHardwareModelsParams object
// with the ability to set a timeout on a request.
func NewHardwareModelQueryHardwareModelsParamsWithTimeout(timeout time.Duration) *HardwareModelQueryHardwareModelsParams {
	return &HardwareModelQueryHardwareModelsParams{
		timeout: timeout,
	}
}

// NewHardwareModelQueryHardwareModelsParamsWithContext creates a new HardwareModelQueryHardwareModelsParams object
// with the ability to set a context for a request.
func NewHardwareModelQueryHardwareModelsParamsWithContext(ctx context.Context) *HardwareModelQueryHardwareModelsParams {
	return &HardwareModelQueryHardwareModelsParams{
		Context: ctx,
	}
}

// NewHardwareModelQueryHardwareModelsParamsWithHTTPClient creates a new HardwareModelQueryHardwareModelsParams object
// with the ability to set a custom HTTPClient for a request.
func NewHardwareModelQueryHardwareModelsParamsWithHTTPClient(client *http.Client) *HardwareModelQueryHardwareModelsParams {
	return &HardwareModelQueryHardwareModelsParams{
		HTTPClient: client,
	}
}

/*
HardwareModelQueryHardwareModelsParams contains all the parameters to send to the API endpoint

	for the hardware model query hardware models operation.

	Typically these are written to a http.Request.
*/
type HardwareModelQueryHardwareModelsParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* FilterBrandID.

	   System defined universally unique Id of the brand.
	*/
	FilterBrandID *string

	/* FilterBrandIds.

	   System defined universally unique Ids of the brand.
	*/
	FilterBrandIds []string

	/* FilterNamePattern.

	   Model name pattern to be matched.
	*/
	FilterNamePattern *string

	/* FilterOriginType.

	     origin of object

	 - ORIGIN_UNSPECIFIED: default options, which says no Operation/Invalid Operation
	 - ORIGIN_IMPORTED: Object imported from global enterprise.
	 - ORIGIN_LOCAL: Objectl created locally.
	 - ORIGIN_GLOBAL: Object created in global store,
	to use this type user should have root previlage.

	     Default: "ORIGIN_UNSPECIFIED"
	*/
	FilterOriginType *string

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

// WithDefaults hydrates default values in the hardware model query hardware models params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HardwareModelQueryHardwareModelsParams) WithDefaults() *HardwareModelQueryHardwareModelsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the hardware model query hardware models params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HardwareModelQueryHardwareModelsParams) SetDefaults() {
	var (
		filterOriginTypeDefault = string("ORIGIN_UNSPECIFIED")
	)

	val := HardwareModelQueryHardwareModelsParams{
		FilterOriginType: &filterOriginTypeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the hardware model query hardware models params
func (o *HardwareModelQueryHardwareModelsParams) WithTimeout(timeout time.Duration) *HardwareModelQueryHardwareModelsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the hardware model query hardware models params
func (o *HardwareModelQueryHardwareModelsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the hardware model query hardware models params
func (o *HardwareModelQueryHardwareModelsParams) WithContext(ctx context.Context) *HardwareModelQueryHardwareModelsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the hardware model query hardware models params
func (o *HardwareModelQueryHardwareModelsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the hardware model query hardware models params
func (o *HardwareModelQueryHardwareModelsParams) WithHTTPClient(client *http.Client) *HardwareModelQueryHardwareModelsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the hardware model query hardware models params
func (o *HardwareModelQueryHardwareModelsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the hardware model query hardware models params
func (o *HardwareModelQueryHardwareModelsParams) WithXRequestID(xRequestID *string) *HardwareModelQueryHardwareModelsParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the hardware model query hardware models params
func (o *HardwareModelQueryHardwareModelsParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithFilterBrandID adds the filterBrandID to the hardware model query hardware models params
func (o *HardwareModelQueryHardwareModelsParams) WithFilterBrandID(filterBrandID *string) *HardwareModelQueryHardwareModelsParams {
	o.SetFilterBrandID(filterBrandID)
	return o
}

// SetFilterBrandID adds the filterBrandId to the hardware model query hardware models params
func (o *HardwareModelQueryHardwareModelsParams) SetFilterBrandID(filterBrandID *string) {
	o.FilterBrandID = filterBrandID
}

// WithFilterBrandIds adds the filterBrandIds to the hardware model query hardware models params
func (o *HardwareModelQueryHardwareModelsParams) WithFilterBrandIds(filterBrandIds []string) *HardwareModelQueryHardwareModelsParams {
	o.SetFilterBrandIds(filterBrandIds)
	return o
}

// SetFilterBrandIds adds the filterBrandIds to the hardware model query hardware models params
func (o *HardwareModelQueryHardwareModelsParams) SetFilterBrandIds(filterBrandIds []string) {
	o.FilterBrandIds = filterBrandIds
}

// WithFilterNamePattern adds the filterNamePattern to the hardware model query hardware models params
func (o *HardwareModelQueryHardwareModelsParams) WithFilterNamePattern(filterNamePattern *string) *HardwareModelQueryHardwareModelsParams {
	o.SetFilterNamePattern(filterNamePattern)
	return o
}

// SetFilterNamePattern adds the filterNamePattern to the hardware model query hardware models params
func (o *HardwareModelQueryHardwareModelsParams) SetFilterNamePattern(filterNamePattern *string) {
	o.FilterNamePattern = filterNamePattern
}

// WithFilterOriginType adds the filterOriginType to the hardware model query hardware models params
func (o *HardwareModelQueryHardwareModelsParams) WithFilterOriginType(filterOriginType *string) *HardwareModelQueryHardwareModelsParams {
	o.SetFilterOriginType(filterOriginType)
	return o
}

// SetFilterOriginType adds the filterOriginType to the hardware model query hardware models params
func (o *HardwareModelQueryHardwareModelsParams) SetFilterOriginType(filterOriginType *string) {
	o.FilterOriginType = filterOriginType
}

// WithNextOrderBy adds the nextOrderBy to the hardware model query hardware models params
func (o *HardwareModelQueryHardwareModelsParams) WithNextOrderBy(nextOrderBy []string) *HardwareModelQueryHardwareModelsParams {
	o.SetNextOrderBy(nextOrderBy)
	return o
}

// SetNextOrderBy adds the nextOrderBy to the hardware model query hardware models params
func (o *HardwareModelQueryHardwareModelsParams) SetNextOrderBy(nextOrderBy []string) {
	o.NextOrderBy = nextOrderBy
}

// WithNextPageNum adds the nextPageNum to the hardware model query hardware models params
func (o *HardwareModelQueryHardwareModelsParams) WithNextPageNum(nextPageNum *int64) *HardwareModelQueryHardwareModelsParams {
	o.SetNextPageNum(nextPageNum)
	return o
}

// SetNextPageNum adds the nextPageNum to the hardware model query hardware models params
func (o *HardwareModelQueryHardwareModelsParams) SetNextPageNum(nextPageNum *int64) {
	o.NextPageNum = nextPageNum
}

// WithNextPageSize adds the nextPageSize to the hardware model query hardware models params
func (o *HardwareModelQueryHardwareModelsParams) WithNextPageSize(nextPageSize *int64) *HardwareModelQueryHardwareModelsParams {
	o.SetNextPageSize(nextPageSize)
	return o
}

// SetNextPageSize adds the nextPageSize to the hardware model query hardware models params
func (o *HardwareModelQueryHardwareModelsParams) SetNextPageSize(nextPageSize *int64) {
	o.NextPageSize = nextPageSize
}

// WithNextPageToken adds the nextPageToken to the hardware model query hardware models params
func (o *HardwareModelQueryHardwareModelsParams) WithNextPageToken(nextPageToken *string) *HardwareModelQueryHardwareModelsParams {
	o.SetNextPageToken(nextPageToken)
	return o
}

// SetNextPageToken adds the nextPageToken to the hardware model query hardware models params
func (o *HardwareModelQueryHardwareModelsParams) SetNextPageToken(nextPageToken *string) {
	o.NextPageToken = nextPageToken
}

// WithNextTotalPages adds the nextTotalPages to the hardware model query hardware models params
func (o *HardwareModelQueryHardwareModelsParams) WithNextTotalPages(nextTotalPages *int64) *HardwareModelQueryHardwareModelsParams {
	o.SetNextTotalPages(nextTotalPages)
	return o
}

// SetNextTotalPages adds the nextTotalPages to the hardware model query hardware models params
func (o *HardwareModelQueryHardwareModelsParams) SetNextTotalPages(nextTotalPages *int64) {
	o.NextTotalPages = nextTotalPages
}

// WithSummary adds the summary to the hardware model query hardware models params
func (o *HardwareModelQueryHardwareModelsParams) WithSummary(summary *bool) *HardwareModelQueryHardwareModelsParams {
	o.SetSummary(summary)
	return o
}

// SetSummary adds the summary to the hardware model query hardware models params
func (o *HardwareModelQueryHardwareModelsParams) SetSummary(summary *bool) {
	o.Summary = summary
}

// WriteToRequest writes these params to a swagger request
func (o *HardwareModelQueryHardwareModelsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.FilterBrandID != nil {

		// query param filter.brandId
		var qrFilterBrandID string

		if o.FilterBrandID != nil {
			qrFilterBrandID = *o.FilterBrandID
		}
		qFilterBrandID := qrFilterBrandID
		if qFilterBrandID != "" {

			if err := r.SetQueryParam("filter.brandId", qFilterBrandID); err != nil {
				return err
			}
		}
	}

	if o.FilterBrandIds != nil {

		// binding items for filter.brandIds
		joinedFilterBrandIds := o.bindParamFilterBrandIds(reg)

		// query array param filter.brandIds
		if err := r.SetQueryParam("filter.brandIds", joinedFilterBrandIds...); err != nil {
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

	if o.FilterOriginType != nil {

		// query param filter.originType
		var qrFilterOriginType string

		if o.FilterOriginType != nil {
			qrFilterOriginType = *o.FilterOriginType
		}
		qFilterOriginType := qrFilterOriginType
		if qFilterOriginType != "" {

			if err := r.SetQueryParam("filter.originType", qFilterOriginType); err != nil {
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

// bindParamHardwareModelQueryHardwareModels binds the parameter filter.brandIds
func (o *HardwareModelQueryHardwareModelsParams) bindParamFilterBrandIds(formats strfmt.Registry) []string {
	filterBrandIdsIR := o.FilterBrandIds

	var filterBrandIdsIC []string
	for _, filterBrandIdsIIR := range filterBrandIdsIR { // explode []string

		filterBrandIdsIIV := filterBrandIdsIIR // string as string
		filterBrandIdsIC = append(filterBrandIdsIC, filterBrandIdsIIV)
	}

	// items.CollectionFormat: "multi"
	filterBrandIdsIS := swag.JoinByFormat(filterBrandIdsIC, "multi")

	return filterBrandIdsIS
}

// bindParamHardwareModelQueryHardwareModels binds the parameter next.orderBy
func (o *HardwareModelQueryHardwareModelsParams) bindParamNextOrderBy(formats strfmt.Registry) []string {
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

// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package enterprise_entitlements_report

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

// NewEnterpriseEntitlementsReportGetAllowedEnterprisesParams creates a new EnterpriseEntitlementsReportGetAllowedEnterprisesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEnterpriseEntitlementsReportGetAllowedEnterprisesParams() *EnterpriseEntitlementsReportGetAllowedEnterprisesParams {
	return &EnterpriseEntitlementsReportGetAllowedEnterprisesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEnterpriseEntitlementsReportGetAllowedEnterprisesParamsWithTimeout creates a new EnterpriseEntitlementsReportGetAllowedEnterprisesParams object
// with the ability to set a timeout on a request.
func NewEnterpriseEntitlementsReportGetAllowedEnterprisesParamsWithTimeout(timeout time.Duration) *EnterpriseEntitlementsReportGetAllowedEnterprisesParams {
	return &EnterpriseEntitlementsReportGetAllowedEnterprisesParams{
		timeout: timeout,
	}
}

// NewEnterpriseEntitlementsReportGetAllowedEnterprisesParamsWithContext creates a new EnterpriseEntitlementsReportGetAllowedEnterprisesParams object
// with the ability to set a context for a request.
func NewEnterpriseEntitlementsReportGetAllowedEnterprisesParamsWithContext(ctx context.Context) *EnterpriseEntitlementsReportGetAllowedEnterprisesParams {
	return &EnterpriseEntitlementsReportGetAllowedEnterprisesParams{
		Context: ctx,
	}
}

// NewEnterpriseEntitlementsReportGetAllowedEnterprisesParamsWithHTTPClient creates a new EnterpriseEntitlementsReportGetAllowedEnterprisesParams object
// with the ability to set a custom HTTPClient for a request.
func NewEnterpriseEntitlementsReportGetAllowedEnterprisesParamsWithHTTPClient(client *http.Client) *EnterpriseEntitlementsReportGetAllowedEnterprisesParams {
	return &EnterpriseEntitlementsReportGetAllowedEnterprisesParams{
		HTTPClient: client,
	}
}

/*
EnterpriseEntitlementsReportGetAllowedEnterprisesParams contains all the parameters to send to the API endpoint

	for the enterprise entitlements report get allowed enterprises operation.

	Typically these are written to a http.Request.
*/
type EnterpriseEntitlementsReportGetAllowedEnterprisesParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// FilterAll.
	FilterAll *bool

	// FilterEntpstate.
	FilterEntpstate *string

	// FilterHubspotid.
	FilterHubspotid *string

	// FilterNamePattern.
	FilterNamePattern *string

	// FilterProjectid.
	FilterProjectid []string

	// FilterRoleName.
	FilterRoleName *string

	// FilterSfdcid.
	FilterSfdcid *string

	// FilterUserstate.
	FilterUserstate *string

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

// WithDefaults hydrates default values in the enterprise entitlements report get allowed enterprises params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesParams) WithDefaults() *EnterpriseEntitlementsReportGetAllowedEnterprisesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the enterprise entitlements report get allowed enterprises params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the enterprise entitlements report get allowed enterprises params
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesParams) WithTimeout(timeout time.Duration) *EnterpriseEntitlementsReportGetAllowedEnterprisesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the enterprise entitlements report get allowed enterprises params
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the enterprise entitlements report get allowed enterprises params
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesParams) WithContext(ctx context.Context) *EnterpriseEntitlementsReportGetAllowedEnterprisesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the enterprise entitlements report get allowed enterprises params
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the enterprise entitlements report get allowed enterprises params
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesParams) WithHTTPClient(client *http.Client) *EnterpriseEntitlementsReportGetAllowedEnterprisesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the enterprise entitlements report get allowed enterprises params
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the enterprise entitlements report get allowed enterprises params
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesParams) WithXRequestID(xRequestID *string) *EnterpriseEntitlementsReportGetAllowedEnterprisesParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the enterprise entitlements report get allowed enterprises params
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithFilterAll adds the filterAll to the enterprise entitlements report get allowed enterprises params
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesParams) WithFilterAll(filterAll *bool) *EnterpriseEntitlementsReportGetAllowedEnterprisesParams {
	o.SetFilterAll(filterAll)
	return o
}

// SetFilterAll adds the filterAll to the enterprise entitlements report get allowed enterprises params
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesParams) SetFilterAll(filterAll *bool) {
	o.FilterAll = filterAll
}

// WithFilterEntpstate adds the filterEntpstate to the enterprise entitlements report get allowed enterprises params
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesParams) WithFilterEntpstate(filterEntpstate *string) *EnterpriseEntitlementsReportGetAllowedEnterprisesParams {
	o.SetFilterEntpstate(filterEntpstate)
	return o
}

// SetFilterEntpstate adds the filterEntpstate to the enterprise entitlements report get allowed enterprises params
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesParams) SetFilterEntpstate(filterEntpstate *string) {
	o.FilterEntpstate = filterEntpstate
}

// WithFilterHubspotid adds the filterHubspotid to the enterprise entitlements report get allowed enterprises params
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesParams) WithFilterHubspotid(filterHubspotid *string) *EnterpriseEntitlementsReportGetAllowedEnterprisesParams {
	o.SetFilterHubspotid(filterHubspotid)
	return o
}

// SetFilterHubspotid adds the filterHubspotid to the enterprise entitlements report get allowed enterprises params
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesParams) SetFilterHubspotid(filterHubspotid *string) {
	o.FilterHubspotid = filterHubspotid
}

// WithFilterNamePattern adds the filterNamePattern to the enterprise entitlements report get allowed enterprises params
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesParams) WithFilterNamePattern(filterNamePattern *string) *EnterpriseEntitlementsReportGetAllowedEnterprisesParams {
	o.SetFilterNamePattern(filterNamePattern)
	return o
}

// SetFilterNamePattern adds the filterNamePattern to the enterprise entitlements report get allowed enterprises params
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesParams) SetFilterNamePattern(filterNamePattern *string) {
	o.FilterNamePattern = filterNamePattern
}

// WithFilterProjectid adds the filterProjectid to the enterprise entitlements report get allowed enterprises params
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesParams) WithFilterProjectid(filterProjectid []string) *EnterpriseEntitlementsReportGetAllowedEnterprisesParams {
	o.SetFilterProjectid(filterProjectid)
	return o
}

// SetFilterProjectid adds the filterProjectid to the enterprise entitlements report get allowed enterprises params
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesParams) SetFilterProjectid(filterProjectid []string) {
	o.FilterProjectid = filterProjectid
}

// WithFilterRoleName adds the filterRoleName to the enterprise entitlements report get allowed enterprises params
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesParams) WithFilterRoleName(filterRoleName *string) *EnterpriseEntitlementsReportGetAllowedEnterprisesParams {
	o.SetFilterRoleName(filterRoleName)
	return o
}

// SetFilterRoleName adds the filterRoleName to the enterprise entitlements report get allowed enterprises params
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesParams) SetFilterRoleName(filterRoleName *string) {
	o.FilterRoleName = filterRoleName
}

// WithFilterSfdcid adds the filterSfdcid to the enterprise entitlements report get allowed enterprises params
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesParams) WithFilterSfdcid(filterSfdcid *string) *EnterpriseEntitlementsReportGetAllowedEnterprisesParams {
	o.SetFilterSfdcid(filterSfdcid)
	return o
}

// SetFilterSfdcid adds the filterSfdcid to the enterprise entitlements report get allowed enterprises params
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesParams) SetFilterSfdcid(filterSfdcid *string) {
	o.FilterSfdcid = filterSfdcid
}

// WithFilterUserstate adds the filterUserstate to the enterprise entitlements report get allowed enterprises params
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesParams) WithFilterUserstate(filterUserstate *string) *EnterpriseEntitlementsReportGetAllowedEnterprisesParams {
	o.SetFilterUserstate(filterUserstate)
	return o
}

// SetFilterUserstate adds the filterUserstate to the enterprise entitlements report get allowed enterprises params
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesParams) SetFilterUserstate(filterUserstate *string) {
	o.FilterUserstate = filterUserstate
}

// WithNextOrderBy adds the nextOrderBy to the enterprise entitlements report get allowed enterprises params
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesParams) WithNextOrderBy(nextOrderBy []string) *EnterpriseEntitlementsReportGetAllowedEnterprisesParams {
	o.SetNextOrderBy(nextOrderBy)
	return o
}

// SetNextOrderBy adds the nextOrderBy to the enterprise entitlements report get allowed enterprises params
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesParams) SetNextOrderBy(nextOrderBy []string) {
	o.NextOrderBy = nextOrderBy
}

// WithNextPageNum adds the nextPageNum to the enterprise entitlements report get allowed enterprises params
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesParams) WithNextPageNum(nextPageNum *int64) *EnterpriseEntitlementsReportGetAllowedEnterprisesParams {
	o.SetNextPageNum(nextPageNum)
	return o
}

// SetNextPageNum adds the nextPageNum to the enterprise entitlements report get allowed enterprises params
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesParams) SetNextPageNum(nextPageNum *int64) {
	o.NextPageNum = nextPageNum
}

// WithNextPageSize adds the nextPageSize to the enterprise entitlements report get allowed enterprises params
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesParams) WithNextPageSize(nextPageSize *int64) *EnterpriseEntitlementsReportGetAllowedEnterprisesParams {
	o.SetNextPageSize(nextPageSize)
	return o
}

// SetNextPageSize adds the nextPageSize to the enterprise entitlements report get allowed enterprises params
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesParams) SetNextPageSize(nextPageSize *int64) {
	o.NextPageSize = nextPageSize
}

// WithNextPageToken adds the nextPageToken to the enterprise entitlements report get allowed enterprises params
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesParams) WithNextPageToken(nextPageToken *string) *EnterpriseEntitlementsReportGetAllowedEnterprisesParams {
	o.SetNextPageToken(nextPageToken)
	return o
}

// SetNextPageToken adds the nextPageToken to the enterprise entitlements report get allowed enterprises params
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesParams) SetNextPageToken(nextPageToken *string) {
	o.NextPageToken = nextPageToken
}

// WithNextTotalPages adds the nextTotalPages to the enterprise entitlements report get allowed enterprises params
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesParams) WithNextTotalPages(nextTotalPages *int64) *EnterpriseEntitlementsReportGetAllowedEnterprisesParams {
	o.SetNextTotalPages(nextTotalPages)
	return o
}

// SetNextTotalPages adds the nextTotalPages to the enterprise entitlements report get allowed enterprises params
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesParams) SetNextTotalPages(nextTotalPages *int64) {
	o.NextTotalPages = nextTotalPages
}

// WithSummary adds the summary to the enterprise entitlements report get allowed enterprises params
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesParams) WithSummary(summary *bool) *EnterpriseEntitlementsReportGetAllowedEnterprisesParams {
	o.SetSummary(summary)
	return o
}

// SetSummary adds the summary to the enterprise entitlements report get allowed enterprises params
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesParams) SetSummary(summary *bool) {
	o.Summary = summary
}

// WriteToRequest writes these params to a swagger request
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.FilterAll != nil {

		// query param filter.all
		var qrFilterAll bool

		if o.FilterAll != nil {
			qrFilterAll = *o.FilterAll
		}
		qFilterAll := swag.FormatBool(qrFilterAll)
		if qFilterAll != "" {

			if err := r.SetQueryParam("filter.all", qFilterAll); err != nil {
				return err
			}
		}
	}

	if o.FilterEntpstate != nil {

		// query param filter.entpstate
		var qrFilterEntpstate string

		if o.FilterEntpstate != nil {
			qrFilterEntpstate = *o.FilterEntpstate
		}
		qFilterEntpstate := qrFilterEntpstate
		if qFilterEntpstate != "" {

			if err := r.SetQueryParam("filter.entpstate", qFilterEntpstate); err != nil {
				return err
			}
		}
	}

	if o.FilterHubspotid != nil {

		// query param filter.hubspotid
		var qrFilterHubspotid string

		if o.FilterHubspotid != nil {
			qrFilterHubspotid = *o.FilterHubspotid
		}
		qFilterHubspotid := qrFilterHubspotid
		if qFilterHubspotid != "" {

			if err := r.SetQueryParam("filter.hubspotid", qFilterHubspotid); err != nil {
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

	if o.FilterProjectid != nil {

		// binding items for filter.projectid
		joinedFilterProjectid := o.bindParamFilterProjectid(reg)

		// query array param filter.projectid
		if err := r.SetQueryParam("filter.projectid", joinedFilterProjectid...); err != nil {
			return err
		}
	}

	if o.FilterRoleName != nil {

		// query param filter.roleName
		var qrFilterRoleName string

		if o.FilterRoleName != nil {
			qrFilterRoleName = *o.FilterRoleName
		}
		qFilterRoleName := qrFilterRoleName
		if qFilterRoleName != "" {

			if err := r.SetQueryParam("filter.roleName", qFilterRoleName); err != nil {
				return err
			}
		}
	}

	if o.FilterSfdcid != nil {

		// query param filter.sfdcid
		var qrFilterSfdcid string

		if o.FilterSfdcid != nil {
			qrFilterSfdcid = *o.FilterSfdcid
		}
		qFilterSfdcid := qrFilterSfdcid
		if qFilterSfdcid != "" {

			if err := r.SetQueryParam("filter.sfdcid", qFilterSfdcid); err != nil {
				return err
			}
		}
	}

	if o.FilterUserstate != nil {

		// query param filter.userstate
		var qrFilterUserstate string

		if o.FilterUserstate != nil {
			qrFilterUserstate = *o.FilterUserstate
		}
		qFilterUserstate := qrFilterUserstate
		if qFilterUserstate != "" {

			if err := r.SetQueryParam("filter.userstate", qFilterUserstate); err != nil {
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

// bindParamEnterpriseEntitlementsReportGetAllowedEnterprises binds the parameter filter.projectid
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesParams) bindParamFilterProjectid(formats strfmt.Registry) []string {
	filterProjectidIR := o.FilterProjectid

	var filterProjectidIC []string
	for _, filterProjectidIIR := range filterProjectidIR { // explode []string

		filterProjectidIIV := filterProjectidIIR // string as string
		filterProjectidIC = append(filterProjectidIC, filterProjectidIIV)
	}

	// items.CollectionFormat: "multi"
	filterProjectidIS := swag.JoinByFormat(filterProjectidIC, "multi")

	return filterProjectidIS
}

// bindParamEnterpriseEntitlementsReportGetAllowedEnterprises binds the parameter next.orderBy
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesParams) bindParamNextOrderBy(formats strfmt.Registry) []string {
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

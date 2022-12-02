// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package identity_access_management

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

// NewIdentityAccessManagementQueryAuthProfilesParams creates a new IdentityAccessManagementQueryAuthProfilesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIdentityAccessManagementQueryAuthProfilesParams() *IdentityAccessManagementQueryAuthProfilesParams {
	return &IdentityAccessManagementQueryAuthProfilesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIdentityAccessManagementQueryAuthProfilesParamsWithTimeout creates a new IdentityAccessManagementQueryAuthProfilesParams object
// with the ability to set a timeout on a request.
func NewIdentityAccessManagementQueryAuthProfilesParamsWithTimeout(timeout time.Duration) *IdentityAccessManagementQueryAuthProfilesParams {
	return &IdentityAccessManagementQueryAuthProfilesParams{
		timeout: timeout,
	}
}

// NewIdentityAccessManagementQueryAuthProfilesParamsWithContext creates a new IdentityAccessManagementQueryAuthProfilesParams object
// with the ability to set a context for a request.
func NewIdentityAccessManagementQueryAuthProfilesParamsWithContext(ctx context.Context) *IdentityAccessManagementQueryAuthProfilesParams {
	return &IdentityAccessManagementQueryAuthProfilesParams{
		Context: ctx,
	}
}

// NewIdentityAccessManagementQueryAuthProfilesParamsWithHTTPClient creates a new IdentityAccessManagementQueryAuthProfilesParams object
// with the ability to set a custom HTTPClient for a request.
func NewIdentityAccessManagementQueryAuthProfilesParamsWithHTTPClient(client *http.Client) *IdentityAccessManagementQueryAuthProfilesParams {
	return &IdentityAccessManagementQueryAuthProfilesParams{
		HTTPClient: client,
	}
}

/*
IdentityAccessManagementQueryAuthProfilesParams contains all the parameters to send to the API endpoint

	for the identity access management query auth profiles operation.

	Typically these are written to a http.Request.
*/
type IdentityAccessManagementQueryAuthProfilesParams struct {

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

// WithDefaults hydrates default values in the identity access management query auth profiles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IdentityAccessManagementQueryAuthProfilesParams) WithDefaults() *IdentityAccessManagementQueryAuthProfilesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the identity access management query auth profiles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IdentityAccessManagementQueryAuthProfilesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the identity access management query auth profiles params
func (o *IdentityAccessManagementQueryAuthProfilesParams) WithTimeout(timeout time.Duration) *IdentityAccessManagementQueryAuthProfilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the identity access management query auth profiles params
func (o *IdentityAccessManagementQueryAuthProfilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the identity access management query auth profiles params
func (o *IdentityAccessManagementQueryAuthProfilesParams) WithContext(ctx context.Context) *IdentityAccessManagementQueryAuthProfilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the identity access management query auth profiles params
func (o *IdentityAccessManagementQueryAuthProfilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the identity access management query auth profiles params
func (o *IdentityAccessManagementQueryAuthProfilesParams) WithHTTPClient(client *http.Client) *IdentityAccessManagementQueryAuthProfilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the identity access management query auth profiles params
func (o *IdentityAccessManagementQueryAuthProfilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the identity access management query auth profiles params
func (o *IdentityAccessManagementQueryAuthProfilesParams) WithXRequestID(xRequestID *string) *IdentityAccessManagementQueryAuthProfilesParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the identity access management query auth profiles params
func (o *IdentityAccessManagementQueryAuthProfilesParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithFilterAll adds the filterAll to the identity access management query auth profiles params
func (o *IdentityAccessManagementQueryAuthProfilesParams) WithFilterAll(filterAll *bool) *IdentityAccessManagementQueryAuthProfilesParams {
	o.SetFilterAll(filterAll)
	return o
}

// SetFilterAll adds the filterAll to the identity access management query auth profiles params
func (o *IdentityAccessManagementQueryAuthProfilesParams) SetFilterAll(filterAll *bool) {
	o.FilterAll = filterAll
}

// WithFilterEntpstate adds the filterEntpstate to the identity access management query auth profiles params
func (o *IdentityAccessManagementQueryAuthProfilesParams) WithFilterEntpstate(filterEntpstate *string) *IdentityAccessManagementQueryAuthProfilesParams {
	o.SetFilterEntpstate(filterEntpstate)
	return o
}

// SetFilterEntpstate adds the filterEntpstate to the identity access management query auth profiles params
func (o *IdentityAccessManagementQueryAuthProfilesParams) SetFilterEntpstate(filterEntpstate *string) {
	o.FilterEntpstate = filterEntpstate
}

// WithFilterHubspotid adds the filterHubspotid to the identity access management query auth profiles params
func (o *IdentityAccessManagementQueryAuthProfilesParams) WithFilterHubspotid(filterHubspotid *string) *IdentityAccessManagementQueryAuthProfilesParams {
	o.SetFilterHubspotid(filterHubspotid)
	return o
}

// SetFilterHubspotid adds the filterHubspotid to the identity access management query auth profiles params
func (o *IdentityAccessManagementQueryAuthProfilesParams) SetFilterHubspotid(filterHubspotid *string) {
	o.FilterHubspotid = filterHubspotid
}

// WithFilterNamePattern adds the filterNamePattern to the identity access management query auth profiles params
func (o *IdentityAccessManagementQueryAuthProfilesParams) WithFilterNamePattern(filterNamePattern *string) *IdentityAccessManagementQueryAuthProfilesParams {
	o.SetFilterNamePattern(filterNamePattern)
	return o
}

// SetFilterNamePattern adds the filterNamePattern to the identity access management query auth profiles params
func (o *IdentityAccessManagementQueryAuthProfilesParams) SetFilterNamePattern(filterNamePattern *string) {
	o.FilterNamePattern = filterNamePattern
}

// WithFilterProjectid adds the filterProjectid to the identity access management query auth profiles params
func (o *IdentityAccessManagementQueryAuthProfilesParams) WithFilterProjectid(filterProjectid []string) *IdentityAccessManagementQueryAuthProfilesParams {
	o.SetFilterProjectid(filterProjectid)
	return o
}

// SetFilterProjectid adds the filterProjectid to the identity access management query auth profiles params
func (o *IdentityAccessManagementQueryAuthProfilesParams) SetFilterProjectid(filterProjectid []string) {
	o.FilterProjectid = filterProjectid
}

// WithFilterRoleName adds the filterRoleName to the identity access management query auth profiles params
func (o *IdentityAccessManagementQueryAuthProfilesParams) WithFilterRoleName(filterRoleName *string) *IdentityAccessManagementQueryAuthProfilesParams {
	o.SetFilterRoleName(filterRoleName)
	return o
}

// SetFilterRoleName adds the filterRoleName to the identity access management query auth profiles params
func (o *IdentityAccessManagementQueryAuthProfilesParams) SetFilterRoleName(filterRoleName *string) {
	o.FilterRoleName = filterRoleName
}

// WithFilterSfdcid adds the filterSfdcid to the identity access management query auth profiles params
func (o *IdentityAccessManagementQueryAuthProfilesParams) WithFilterSfdcid(filterSfdcid *string) *IdentityAccessManagementQueryAuthProfilesParams {
	o.SetFilterSfdcid(filterSfdcid)
	return o
}

// SetFilterSfdcid adds the filterSfdcid to the identity access management query auth profiles params
func (o *IdentityAccessManagementQueryAuthProfilesParams) SetFilterSfdcid(filterSfdcid *string) {
	o.FilterSfdcid = filterSfdcid
}

// WithFilterUserstate adds the filterUserstate to the identity access management query auth profiles params
func (o *IdentityAccessManagementQueryAuthProfilesParams) WithFilterUserstate(filterUserstate *string) *IdentityAccessManagementQueryAuthProfilesParams {
	o.SetFilterUserstate(filterUserstate)
	return o
}

// SetFilterUserstate adds the filterUserstate to the identity access management query auth profiles params
func (o *IdentityAccessManagementQueryAuthProfilesParams) SetFilterUserstate(filterUserstate *string) {
	o.FilterUserstate = filterUserstate
}

// WithNextOrderBy adds the nextOrderBy to the identity access management query auth profiles params
func (o *IdentityAccessManagementQueryAuthProfilesParams) WithNextOrderBy(nextOrderBy []string) *IdentityAccessManagementQueryAuthProfilesParams {
	o.SetNextOrderBy(nextOrderBy)
	return o
}

// SetNextOrderBy adds the nextOrderBy to the identity access management query auth profiles params
func (o *IdentityAccessManagementQueryAuthProfilesParams) SetNextOrderBy(nextOrderBy []string) {
	o.NextOrderBy = nextOrderBy
}

// WithNextPageNum adds the nextPageNum to the identity access management query auth profiles params
func (o *IdentityAccessManagementQueryAuthProfilesParams) WithNextPageNum(nextPageNum *int64) *IdentityAccessManagementQueryAuthProfilesParams {
	o.SetNextPageNum(nextPageNum)
	return o
}

// SetNextPageNum adds the nextPageNum to the identity access management query auth profiles params
func (o *IdentityAccessManagementQueryAuthProfilesParams) SetNextPageNum(nextPageNum *int64) {
	o.NextPageNum = nextPageNum
}

// WithNextPageSize adds the nextPageSize to the identity access management query auth profiles params
func (o *IdentityAccessManagementQueryAuthProfilesParams) WithNextPageSize(nextPageSize *int64) *IdentityAccessManagementQueryAuthProfilesParams {
	o.SetNextPageSize(nextPageSize)
	return o
}

// SetNextPageSize adds the nextPageSize to the identity access management query auth profiles params
func (o *IdentityAccessManagementQueryAuthProfilesParams) SetNextPageSize(nextPageSize *int64) {
	o.NextPageSize = nextPageSize
}

// WithNextPageToken adds the nextPageToken to the identity access management query auth profiles params
func (o *IdentityAccessManagementQueryAuthProfilesParams) WithNextPageToken(nextPageToken *string) *IdentityAccessManagementQueryAuthProfilesParams {
	o.SetNextPageToken(nextPageToken)
	return o
}

// SetNextPageToken adds the nextPageToken to the identity access management query auth profiles params
func (o *IdentityAccessManagementQueryAuthProfilesParams) SetNextPageToken(nextPageToken *string) {
	o.NextPageToken = nextPageToken
}

// WithNextTotalPages adds the nextTotalPages to the identity access management query auth profiles params
func (o *IdentityAccessManagementQueryAuthProfilesParams) WithNextTotalPages(nextTotalPages *int64) *IdentityAccessManagementQueryAuthProfilesParams {
	o.SetNextTotalPages(nextTotalPages)
	return o
}

// SetNextTotalPages adds the nextTotalPages to the identity access management query auth profiles params
func (o *IdentityAccessManagementQueryAuthProfilesParams) SetNextTotalPages(nextTotalPages *int64) {
	o.NextTotalPages = nextTotalPages
}

// WithSummary adds the summary to the identity access management query auth profiles params
func (o *IdentityAccessManagementQueryAuthProfilesParams) WithSummary(summary *bool) *IdentityAccessManagementQueryAuthProfilesParams {
	o.SetSummary(summary)
	return o
}

// SetSummary adds the summary to the identity access management query auth profiles params
func (o *IdentityAccessManagementQueryAuthProfilesParams) SetSummary(summary *bool) {
	o.Summary = summary
}

// WriteToRequest writes these params to a swagger request
func (o *IdentityAccessManagementQueryAuthProfilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// bindParamIdentityAccessManagementQueryAuthProfiles binds the parameter filter.projectid
func (o *IdentityAccessManagementQueryAuthProfilesParams) bindParamFilterProjectid(formats strfmt.Registry) []string {
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

// bindParamIdentityAccessManagementQueryAuthProfiles binds the parameter next.orderBy
func (o *IdentityAccessManagementQueryAuthProfilesParams) bindParamNextOrderBy(formats strfmt.Registry) []string {
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

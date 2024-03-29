// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package edge_network_configuration

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

// NewEdgeNetworkConfigurationQueryEdgeNetworksParams creates a new EdgeNetworkConfigurationQueryEdgeNetworksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEdgeNetworkConfigurationQueryEdgeNetworksParams() *EdgeNetworkConfigurationQueryEdgeNetworksParams {
	return &EdgeNetworkConfigurationQueryEdgeNetworksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeNetworkConfigurationQueryEdgeNetworksParamsWithTimeout creates a new EdgeNetworkConfigurationQueryEdgeNetworksParams object
// with the ability to set a timeout on a request.
func NewEdgeNetworkConfigurationQueryEdgeNetworksParamsWithTimeout(timeout time.Duration) *EdgeNetworkConfigurationQueryEdgeNetworksParams {
	return &EdgeNetworkConfigurationQueryEdgeNetworksParams{
		timeout: timeout,
	}
}

// NewEdgeNetworkConfigurationQueryEdgeNetworksParamsWithContext creates a new EdgeNetworkConfigurationQueryEdgeNetworksParams object
// with the ability to set a context for a request.
func NewEdgeNetworkConfigurationQueryEdgeNetworksParamsWithContext(ctx context.Context) *EdgeNetworkConfigurationQueryEdgeNetworksParams {
	return &EdgeNetworkConfigurationQueryEdgeNetworksParams{
		Context: ctx,
	}
}

// NewEdgeNetworkConfigurationQueryEdgeNetworksParamsWithHTTPClient creates a new EdgeNetworkConfigurationQueryEdgeNetworksParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeNetworkConfigurationQueryEdgeNetworksParamsWithHTTPClient(client *http.Client) *EdgeNetworkConfigurationQueryEdgeNetworksParams {
	return &EdgeNetworkConfigurationQueryEdgeNetworksParams{
		HTTPClient: client,
	}
}

/*
EdgeNetworkConfigurationQueryEdgeNetworksParams contains all the parameters to send to the API endpoint

	for the edge network configuration query edge networks operation.

	Typically these are written to a http.Request.
*/
type EdgeNetworkConfigurationQueryEdgeNetworksParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// FilterDist.
	//
	// Default: "NETWORK_WIRELESS_TYPE_UNSPECIFIED"
	FilterDist *string

	// FilterKind.
	//
	// Default: "NETWORK_KIND_UNSPECIFIED"
	FilterKind *string

	// FilterNamePattern.
	FilterNamePattern *string

	// FilterProjectName.
	FilterProjectName *string

	// FilterProjectNamePattern.
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

	// Summary.
	Summary *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edge network configuration query edge networks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNetworkConfigurationQueryEdgeNetworksParams) WithDefaults() *EdgeNetworkConfigurationQueryEdgeNetworksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge network configuration query edge networks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNetworkConfigurationQueryEdgeNetworksParams) SetDefaults() {
	var (
		filterDistDefault = string("NETWORK_WIRELESS_TYPE_UNSPECIFIED")

		filterKindDefault = string("NETWORK_KIND_UNSPECIFIED")
	)

	val := EdgeNetworkConfigurationQueryEdgeNetworksParams{
		FilterDist: &filterDistDefault,
		FilterKind: &filterKindDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the edge network configuration query edge networks params
func (o *EdgeNetworkConfigurationQueryEdgeNetworksParams) WithTimeout(timeout time.Duration) *EdgeNetworkConfigurationQueryEdgeNetworksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge network configuration query edge networks params
func (o *EdgeNetworkConfigurationQueryEdgeNetworksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge network configuration query edge networks params
func (o *EdgeNetworkConfigurationQueryEdgeNetworksParams) WithContext(ctx context.Context) *EdgeNetworkConfigurationQueryEdgeNetworksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge network configuration query edge networks params
func (o *EdgeNetworkConfigurationQueryEdgeNetworksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge network configuration query edge networks params
func (o *EdgeNetworkConfigurationQueryEdgeNetworksParams) WithHTTPClient(client *http.Client) *EdgeNetworkConfigurationQueryEdgeNetworksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge network configuration query edge networks params
func (o *EdgeNetworkConfigurationQueryEdgeNetworksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the edge network configuration query edge networks params
func (o *EdgeNetworkConfigurationQueryEdgeNetworksParams) WithXRequestID(xRequestID *string) *EdgeNetworkConfigurationQueryEdgeNetworksParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the edge network configuration query edge networks params
func (o *EdgeNetworkConfigurationQueryEdgeNetworksParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithFilterDist adds the filterDist to the edge network configuration query edge networks params
func (o *EdgeNetworkConfigurationQueryEdgeNetworksParams) WithFilterDist(filterDist *string) *EdgeNetworkConfigurationQueryEdgeNetworksParams {
	o.SetFilterDist(filterDist)
	return o
}

// SetFilterDist adds the filterDist to the edge network configuration query edge networks params
func (o *EdgeNetworkConfigurationQueryEdgeNetworksParams) SetFilterDist(filterDist *string) {
	o.FilterDist = filterDist
}

// WithFilterKind adds the filterKind to the edge network configuration query edge networks params
func (o *EdgeNetworkConfigurationQueryEdgeNetworksParams) WithFilterKind(filterKind *string) *EdgeNetworkConfigurationQueryEdgeNetworksParams {
	o.SetFilterKind(filterKind)
	return o
}

// SetFilterKind adds the filterKind to the edge network configuration query edge networks params
func (o *EdgeNetworkConfigurationQueryEdgeNetworksParams) SetFilterKind(filterKind *string) {
	o.FilterKind = filterKind
}

// WithFilterNamePattern adds the filterNamePattern to the edge network configuration query edge networks params
func (o *EdgeNetworkConfigurationQueryEdgeNetworksParams) WithFilterNamePattern(filterNamePattern *string) *EdgeNetworkConfigurationQueryEdgeNetworksParams {
	o.SetFilterNamePattern(filterNamePattern)
	return o
}

// SetFilterNamePattern adds the filterNamePattern to the edge network configuration query edge networks params
func (o *EdgeNetworkConfigurationQueryEdgeNetworksParams) SetFilterNamePattern(filterNamePattern *string) {
	o.FilterNamePattern = filterNamePattern
}

// WithFilterProjectName adds the filterProjectName to the edge network configuration query edge networks params
func (o *EdgeNetworkConfigurationQueryEdgeNetworksParams) WithFilterProjectName(filterProjectName *string) *EdgeNetworkConfigurationQueryEdgeNetworksParams {
	o.SetFilterProjectName(filterProjectName)
	return o
}

// SetFilterProjectName adds the filterProjectName to the edge network configuration query edge networks params
func (o *EdgeNetworkConfigurationQueryEdgeNetworksParams) SetFilterProjectName(filterProjectName *string) {
	o.FilterProjectName = filterProjectName
}

// WithFilterProjectNamePattern adds the filterProjectNamePattern to the edge network configuration query edge networks params
func (o *EdgeNetworkConfigurationQueryEdgeNetworksParams) WithFilterProjectNamePattern(filterProjectNamePattern *string) *EdgeNetworkConfigurationQueryEdgeNetworksParams {
	o.SetFilterProjectNamePattern(filterProjectNamePattern)
	return o
}

// SetFilterProjectNamePattern adds the filterProjectNamePattern to the edge network configuration query edge networks params
func (o *EdgeNetworkConfigurationQueryEdgeNetworksParams) SetFilterProjectNamePattern(filterProjectNamePattern *string) {
	o.FilterProjectNamePattern = filterProjectNamePattern
}

// WithNextOrderBy adds the nextOrderBy to the edge network configuration query edge networks params
func (o *EdgeNetworkConfigurationQueryEdgeNetworksParams) WithNextOrderBy(nextOrderBy []string) *EdgeNetworkConfigurationQueryEdgeNetworksParams {
	o.SetNextOrderBy(nextOrderBy)
	return o
}

// SetNextOrderBy adds the nextOrderBy to the edge network configuration query edge networks params
func (o *EdgeNetworkConfigurationQueryEdgeNetworksParams) SetNextOrderBy(nextOrderBy []string) {
	o.NextOrderBy = nextOrderBy
}

// WithNextPageNum adds the nextPageNum to the edge network configuration query edge networks params
func (o *EdgeNetworkConfigurationQueryEdgeNetworksParams) WithNextPageNum(nextPageNum *int64) *EdgeNetworkConfigurationQueryEdgeNetworksParams {
	o.SetNextPageNum(nextPageNum)
	return o
}

// SetNextPageNum adds the nextPageNum to the edge network configuration query edge networks params
func (o *EdgeNetworkConfigurationQueryEdgeNetworksParams) SetNextPageNum(nextPageNum *int64) {
	o.NextPageNum = nextPageNum
}

// WithNextPageSize adds the nextPageSize to the edge network configuration query edge networks params
func (o *EdgeNetworkConfigurationQueryEdgeNetworksParams) WithNextPageSize(nextPageSize *int64) *EdgeNetworkConfigurationQueryEdgeNetworksParams {
	o.SetNextPageSize(nextPageSize)
	return o
}

// SetNextPageSize adds the nextPageSize to the edge network configuration query edge networks params
func (o *EdgeNetworkConfigurationQueryEdgeNetworksParams) SetNextPageSize(nextPageSize *int64) {
	o.NextPageSize = nextPageSize
}

// WithNextPageToken adds the nextPageToken to the edge network configuration query edge networks params
func (o *EdgeNetworkConfigurationQueryEdgeNetworksParams) WithNextPageToken(nextPageToken *string) *EdgeNetworkConfigurationQueryEdgeNetworksParams {
	o.SetNextPageToken(nextPageToken)
	return o
}

// SetNextPageToken adds the nextPageToken to the edge network configuration query edge networks params
func (o *EdgeNetworkConfigurationQueryEdgeNetworksParams) SetNextPageToken(nextPageToken *string) {
	o.NextPageToken = nextPageToken
}

// WithNextTotalPages adds the nextTotalPages to the edge network configuration query edge networks params
func (o *EdgeNetworkConfigurationQueryEdgeNetworksParams) WithNextTotalPages(nextTotalPages *int64) *EdgeNetworkConfigurationQueryEdgeNetworksParams {
	o.SetNextTotalPages(nextTotalPages)
	return o
}

// SetNextTotalPages adds the nextTotalPages to the edge network configuration query edge networks params
func (o *EdgeNetworkConfigurationQueryEdgeNetworksParams) SetNextTotalPages(nextTotalPages *int64) {
	o.NextTotalPages = nextTotalPages
}

// WithSummary adds the summary to the edge network configuration query edge networks params
func (o *EdgeNetworkConfigurationQueryEdgeNetworksParams) WithSummary(summary *bool) *EdgeNetworkConfigurationQueryEdgeNetworksParams {
	o.SetSummary(summary)
	return o
}

// SetSummary adds the summary to the edge network configuration query edge networks params
func (o *EdgeNetworkConfigurationQueryEdgeNetworksParams) SetSummary(summary *bool) {
	o.Summary = summary
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeNetworkConfigurationQueryEdgeNetworksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.FilterDist != nil {

		// query param filter.dist
		var qrFilterDist string

		if o.FilterDist != nil {
			qrFilterDist = *o.FilterDist
		}
		qFilterDist := qrFilterDist
		if qFilterDist != "" {

			if err := r.SetQueryParam("filter.dist", qFilterDist); err != nil {
				return err
			}
		}
	}

	if o.FilterKind != nil {

		// query param filter.kind
		var qrFilterKind string

		if o.FilterKind != nil {
			qrFilterKind = *o.FilterKind
		}
		qFilterKind := qrFilterKind
		if qFilterKind != "" {

			if err := r.SetQueryParam("filter.kind", qFilterKind); err != nil {
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

// bindParamEdgeNetworkConfigurationQueryEdgeNetworks binds the parameter next.orderBy
func (o *EdgeNetworkConfigurationQueryEdgeNetworksParams) bindParamNextOrderBy(formats strfmt.Registry) []string {
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

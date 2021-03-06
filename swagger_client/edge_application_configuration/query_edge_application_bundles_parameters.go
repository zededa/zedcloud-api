// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package edge_application_configuration

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

// NewQueryEdgeApplicationBundlesParams creates a new QueryEdgeApplicationBundlesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQueryEdgeApplicationBundlesParams() *QueryEdgeApplicationBundlesParams {
	return &QueryEdgeApplicationBundlesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQueryEdgeApplicationBundlesParamsWithTimeout creates a new QueryEdgeApplicationBundlesParams object
// with the ability to set a timeout on a request.
func NewQueryEdgeApplicationBundlesParamsWithTimeout(timeout time.Duration) *QueryEdgeApplicationBundlesParams {
	return &QueryEdgeApplicationBundlesParams{
		timeout: timeout,
	}
}

// NewQueryEdgeApplicationBundlesParamsWithContext creates a new QueryEdgeApplicationBundlesParams object
// with the ability to set a context for a request.
func NewQueryEdgeApplicationBundlesParamsWithContext(ctx context.Context) *QueryEdgeApplicationBundlesParams {
	return &QueryEdgeApplicationBundlesParams{
		Context: ctx,
	}
}

// NewQueryEdgeApplicationBundlesParamsWithHTTPClient creates a new QueryEdgeApplicationBundlesParams object
// with the ability to set a custom HTTPClient for a request.
func NewQueryEdgeApplicationBundlesParamsWithHTTPClient(client *http.Client) *QueryEdgeApplicationBundlesParams {
	return &QueryEdgeApplicationBundlesParams{
		HTTPClient: client,
	}
}

/* QueryEdgeApplicationBundlesParams contains all the parameters to send to the API endpoint
   for the query edge application bundles operation.

   Typically these are written to a http.Request.
*/
type QueryEdgeApplicationBundlesParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* FilterAppCategory.

	    category type of the bundle

	- APP_CATEGORY_UNSPECIFIED: Invalid Edge Application Category
	- APP_CATEGORY_OPERATING_SYSTEM: Edge operating systems
	- APP_CATEGORY_INDUSTRIAL: Connectity solution for industrial protocols
	- APP_CATEGORY_EDGE_APPLICATION: Edge application framework featuring composite functions for data ingestion, buffering, analytics and export
	- APP_CATEGORY_NETWORKING: Edge networking services (e.g. SD-WAN, NFV)
	- APP_CATEGORY_SECURITY: Edge security services (e.g. protocol inspection, firewall)
	- APP_CATEGORY_DATA_ANALYTICS: Edge analytics (e.g. AI/ML) and data management (e.g. database, connectors) services
	- APP_CATEGORY_CLOUD_APPLICATION: Edge application runtimes, cloud connectivity and general application enablement
	- APP_CATEGORY_DEVOPS: Tools for Edge Application CI/CD and performance optimization
	- APP_CATEGORY_OTHERS: Miscellaneous functionality

	    Default: "APP_CATEGORY_UNSPECIFIED"
	*/
	FilterAppCategory *string

	/* FilterAppType.

	   app type, eg: vm, container, module

	   Default: "APP_TYPE_UNSPECIFIED"
	*/
	FilterAppType *string

	/* FilterCategory.

	   category type of the bundle
	*/
	FilterCategory *string

	/* FilterDeploymentType.

	   type of deployment for the app, eg: azure, k3s, standalone

	   Default: "DEPLOYMENT_TYPE_UNSPECIFIED"
	*/
	FilterDeploymentType *string

	/* FilterNamePattern.

	   query param : "name-pattern" . Search * namePattern * in name field to filter records
	*/
	FilterNamePattern *string

	/* FilterOriginType.

	     origin of bundle

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

	/* Summary.

	   Only summary of the records required

	   Format: boolean
	*/
	Summary *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the query edge application bundles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryEdgeApplicationBundlesParams) WithDefaults() *QueryEdgeApplicationBundlesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the query edge application bundles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryEdgeApplicationBundlesParams) SetDefaults() {
	var (
		filterAppCategoryDefault = string("APP_CATEGORY_UNSPECIFIED")

		filterAppTypeDefault = string("APP_TYPE_UNSPECIFIED")

		filterDeploymentTypeDefault = string("DEPLOYMENT_TYPE_UNSPECIFIED")

		filterOriginTypeDefault = string("ORIGIN_UNSPECIFIED")
	)

	val := QueryEdgeApplicationBundlesParams{
		FilterAppCategory:    &filterAppCategoryDefault,
		FilterAppType:        &filterAppTypeDefault,
		FilterDeploymentType: &filterDeploymentTypeDefault,
		FilterOriginType:     &filterOriginTypeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the query edge application bundles params
func (o *QueryEdgeApplicationBundlesParams) WithTimeout(timeout time.Duration) *QueryEdgeApplicationBundlesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query edge application bundles params
func (o *QueryEdgeApplicationBundlesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query edge application bundles params
func (o *QueryEdgeApplicationBundlesParams) WithContext(ctx context.Context) *QueryEdgeApplicationBundlesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query edge application bundles params
func (o *QueryEdgeApplicationBundlesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query edge application bundles params
func (o *QueryEdgeApplicationBundlesParams) WithHTTPClient(client *http.Client) *QueryEdgeApplicationBundlesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query edge application bundles params
func (o *QueryEdgeApplicationBundlesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the query edge application bundles params
func (o *QueryEdgeApplicationBundlesParams) WithXRequestID(xRequestID *string) *QueryEdgeApplicationBundlesParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the query edge application bundles params
func (o *QueryEdgeApplicationBundlesParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithFilterAppCategory adds the filterAppCategory to the query edge application bundles params
func (o *QueryEdgeApplicationBundlesParams) WithFilterAppCategory(filterAppCategory *string) *QueryEdgeApplicationBundlesParams {
	o.SetFilterAppCategory(filterAppCategory)
	return o
}

// SetFilterAppCategory adds the filterAppCategory to the query edge application bundles params
func (o *QueryEdgeApplicationBundlesParams) SetFilterAppCategory(filterAppCategory *string) {
	o.FilterAppCategory = filterAppCategory
}

// WithFilterAppType adds the filterAppType to the query edge application bundles params
func (o *QueryEdgeApplicationBundlesParams) WithFilterAppType(filterAppType *string) *QueryEdgeApplicationBundlesParams {
	o.SetFilterAppType(filterAppType)
	return o
}

// SetFilterAppType adds the filterAppType to the query edge application bundles params
func (o *QueryEdgeApplicationBundlesParams) SetFilterAppType(filterAppType *string) {
	o.FilterAppType = filterAppType
}

// WithFilterCategory adds the filterCategory to the query edge application bundles params
func (o *QueryEdgeApplicationBundlesParams) WithFilterCategory(filterCategory *string) *QueryEdgeApplicationBundlesParams {
	o.SetFilterCategory(filterCategory)
	return o
}

// SetFilterCategory adds the filterCategory to the query edge application bundles params
func (o *QueryEdgeApplicationBundlesParams) SetFilterCategory(filterCategory *string) {
	o.FilterCategory = filterCategory
}

// WithFilterDeploymentType adds the filterDeploymentType to the query edge application bundles params
func (o *QueryEdgeApplicationBundlesParams) WithFilterDeploymentType(filterDeploymentType *string) *QueryEdgeApplicationBundlesParams {
	o.SetFilterDeploymentType(filterDeploymentType)
	return o
}

// SetFilterDeploymentType adds the filterDeploymentType to the query edge application bundles params
func (o *QueryEdgeApplicationBundlesParams) SetFilterDeploymentType(filterDeploymentType *string) {
	o.FilterDeploymentType = filterDeploymentType
}

// WithFilterNamePattern adds the filterNamePattern to the query edge application bundles params
func (o *QueryEdgeApplicationBundlesParams) WithFilterNamePattern(filterNamePattern *string) *QueryEdgeApplicationBundlesParams {
	o.SetFilterNamePattern(filterNamePattern)
	return o
}

// SetFilterNamePattern adds the filterNamePattern to the query edge application bundles params
func (o *QueryEdgeApplicationBundlesParams) SetFilterNamePattern(filterNamePattern *string) {
	o.FilterNamePattern = filterNamePattern
}

// WithFilterOriginType adds the filterOriginType to the query edge application bundles params
func (o *QueryEdgeApplicationBundlesParams) WithFilterOriginType(filterOriginType *string) *QueryEdgeApplicationBundlesParams {
	o.SetFilterOriginType(filterOriginType)
	return o
}

// SetFilterOriginType adds the filterOriginType to the query edge application bundles params
func (o *QueryEdgeApplicationBundlesParams) SetFilterOriginType(filterOriginType *string) {
	o.FilterOriginType = filterOriginType
}

// WithNextOrderBy adds the nextOrderBy to the query edge application bundles params
func (o *QueryEdgeApplicationBundlesParams) WithNextOrderBy(nextOrderBy *string) *QueryEdgeApplicationBundlesParams {
	o.SetNextOrderBy(nextOrderBy)
	return o
}

// SetNextOrderBy adds the nextOrderBy to the query edge application bundles params
func (o *QueryEdgeApplicationBundlesParams) SetNextOrderBy(nextOrderBy *string) {
	o.NextOrderBy = nextOrderBy
}

// WithNextPageNum adds the nextPageNum to the query edge application bundles params
func (o *QueryEdgeApplicationBundlesParams) WithNextPageNum(nextPageNum *int64) *QueryEdgeApplicationBundlesParams {
	o.SetNextPageNum(nextPageNum)
	return o
}

// SetNextPageNum adds the nextPageNum to the query edge application bundles params
func (o *QueryEdgeApplicationBundlesParams) SetNextPageNum(nextPageNum *int64) {
	o.NextPageNum = nextPageNum
}

// WithNextPageSize adds the nextPageSize to the query edge application bundles params
func (o *QueryEdgeApplicationBundlesParams) WithNextPageSize(nextPageSize *int64) *QueryEdgeApplicationBundlesParams {
	o.SetNextPageSize(nextPageSize)
	return o
}

// SetNextPageSize adds the nextPageSize to the query edge application bundles params
func (o *QueryEdgeApplicationBundlesParams) SetNextPageSize(nextPageSize *int64) {
	o.NextPageSize = nextPageSize
}

// WithNextPageToken adds the nextPageToken to the query edge application bundles params
func (o *QueryEdgeApplicationBundlesParams) WithNextPageToken(nextPageToken *string) *QueryEdgeApplicationBundlesParams {
	o.SetNextPageToken(nextPageToken)
	return o
}

// SetNextPageToken adds the nextPageToken to the query edge application bundles params
func (o *QueryEdgeApplicationBundlesParams) SetNextPageToken(nextPageToken *string) {
	o.NextPageToken = nextPageToken
}

// WithNextTotalPages adds the nextTotalPages to the query edge application bundles params
func (o *QueryEdgeApplicationBundlesParams) WithNextTotalPages(nextTotalPages *int64) *QueryEdgeApplicationBundlesParams {
	o.SetNextTotalPages(nextTotalPages)
	return o
}

// SetNextTotalPages adds the nextTotalPages to the query edge application bundles params
func (o *QueryEdgeApplicationBundlesParams) SetNextTotalPages(nextTotalPages *int64) {
	o.NextTotalPages = nextTotalPages
}

// WithSummary adds the summary to the query edge application bundles params
func (o *QueryEdgeApplicationBundlesParams) WithSummary(summary *bool) *QueryEdgeApplicationBundlesParams {
	o.SetSummary(summary)
	return o
}

// SetSummary adds the summary to the query edge application bundles params
func (o *QueryEdgeApplicationBundlesParams) SetSummary(summary *bool) {
	o.Summary = summary
}

// WriteToRequest writes these params to a swagger request
func (o *QueryEdgeApplicationBundlesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.FilterAppCategory != nil {

		// query param filter.appCategory
		var qrFilterAppCategory string

		if o.FilterAppCategory != nil {
			qrFilterAppCategory = *o.FilterAppCategory
		}
		qFilterAppCategory := qrFilterAppCategory
		if qFilterAppCategory != "" {

			if err := r.SetQueryParam("filter.appCategory", qFilterAppCategory); err != nil {
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

	if o.FilterCategory != nil {

		// query param filter.category
		var qrFilterCategory string

		if o.FilterCategory != nil {
			qrFilterCategory = *o.FilterCategory
		}
		qFilterCategory := qrFilterCategory
		if qFilterCategory != "" {

			if err := r.SetQueryParam("filter.category", qFilterCategory); err != nil {
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

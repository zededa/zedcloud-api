// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package artifact_manager

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

// NewQueryArtifactsParams creates a new QueryArtifactsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQueryArtifactsParams() *QueryArtifactsParams {
	return &QueryArtifactsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQueryArtifactsParamsWithTimeout creates a new QueryArtifactsParams object
// with the ability to set a timeout on a request.
func NewQueryArtifactsParamsWithTimeout(timeout time.Duration) *QueryArtifactsParams {
	return &QueryArtifactsParams{
		timeout: timeout,
	}
}

// NewQueryArtifactsParamsWithContext creates a new QueryArtifactsParams object
// with the ability to set a context for a request.
func NewQueryArtifactsParamsWithContext(ctx context.Context) *QueryArtifactsParams {
	return &QueryArtifactsParams{
		Context: ctx,
	}
}

// NewQueryArtifactsParamsWithHTTPClient creates a new QueryArtifactsParams object
// with the ability to set a custom HTTPClient for a request.
func NewQueryArtifactsParamsWithHTTPClient(client *http.Client) *QueryArtifactsParams {
	return &QueryArtifactsParams{
		HTTPClient: client,
	}
}

/* QueryArtifactsParams contains all the parameters to send to the API endpoint
   for the query artifacts operation.

   Typically these are written to a http.Request.
*/
type QueryArtifactsParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

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

// WithDefaults hydrates default values in the query artifacts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryArtifactsParams) WithDefaults() *QueryArtifactsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the query artifacts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryArtifactsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the query artifacts params
func (o *QueryArtifactsParams) WithTimeout(timeout time.Duration) *QueryArtifactsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query artifacts params
func (o *QueryArtifactsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query artifacts params
func (o *QueryArtifactsParams) WithContext(ctx context.Context) *QueryArtifactsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query artifacts params
func (o *QueryArtifactsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query artifacts params
func (o *QueryArtifactsParams) WithHTTPClient(client *http.Client) *QueryArtifactsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query artifacts params
func (o *QueryArtifactsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the query artifacts params
func (o *QueryArtifactsParams) WithXRequestID(xRequestID *string) *QueryArtifactsParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the query artifacts params
func (o *QueryArtifactsParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithNextOrderBy adds the nextOrderBy to the query artifacts params
func (o *QueryArtifactsParams) WithNextOrderBy(nextOrderBy *string) *QueryArtifactsParams {
	o.SetNextOrderBy(nextOrderBy)
	return o
}

// SetNextOrderBy adds the nextOrderBy to the query artifacts params
func (o *QueryArtifactsParams) SetNextOrderBy(nextOrderBy *string) {
	o.NextOrderBy = nextOrderBy
}

// WithNextPageNum adds the nextPageNum to the query artifacts params
func (o *QueryArtifactsParams) WithNextPageNum(nextPageNum *int64) *QueryArtifactsParams {
	o.SetNextPageNum(nextPageNum)
	return o
}

// SetNextPageNum adds the nextPageNum to the query artifacts params
func (o *QueryArtifactsParams) SetNextPageNum(nextPageNum *int64) {
	o.NextPageNum = nextPageNum
}

// WithNextPageSize adds the nextPageSize to the query artifacts params
func (o *QueryArtifactsParams) WithNextPageSize(nextPageSize *int64) *QueryArtifactsParams {
	o.SetNextPageSize(nextPageSize)
	return o
}

// SetNextPageSize adds the nextPageSize to the query artifacts params
func (o *QueryArtifactsParams) SetNextPageSize(nextPageSize *int64) {
	o.NextPageSize = nextPageSize
}

// WithNextPageToken adds the nextPageToken to the query artifacts params
func (o *QueryArtifactsParams) WithNextPageToken(nextPageToken *string) *QueryArtifactsParams {
	o.SetNextPageToken(nextPageToken)
	return o
}

// SetNextPageToken adds the nextPageToken to the query artifacts params
func (o *QueryArtifactsParams) SetNextPageToken(nextPageToken *string) {
	o.NextPageToken = nextPageToken
}

// WithNextTotalPages adds the nextTotalPages to the query artifacts params
func (o *QueryArtifactsParams) WithNextTotalPages(nextTotalPages *int64) *QueryArtifactsParams {
	o.SetNextTotalPages(nextTotalPages)
	return o
}

// SetNextTotalPages adds the nextTotalPages to the query artifacts params
func (o *QueryArtifactsParams) SetNextTotalPages(nextTotalPages *int64) {
	o.NextTotalPages = nextTotalPages
}

// WithSummary adds the summary to the query artifacts params
func (o *QueryArtifactsParams) WithSummary(summary *bool) *QueryArtifactsParams {
	o.SetSummary(summary)
	return o
}

// SetSummary adds the summary to the query artifacts params
func (o *QueryArtifactsParams) SetSummary(summary *bool) {
	o.Summary = summary
}

// WriteToRequest writes these params to a swagger request
func (o *QueryArtifactsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

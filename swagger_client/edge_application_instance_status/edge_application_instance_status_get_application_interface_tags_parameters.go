// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package edge_application_instance_status

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

// NewEdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams creates a new EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams() *EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams {
	return &EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeApplicationInstanceStatusGetApplicationInterfaceTagsParamsWithTimeout creates a new EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams object
// with the ability to set a timeout on a request.
func NewEdgeApplicationInstanceStatusGetApplicationInterfaceTagsParamsWithTimeout(timeout time.Duration) *EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams {
	return &EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams{
		timeout: timeout,
	}
}

// NewEdgeApplicationInstanceStatusGetApplicationInterfaceTagsParamsWithContext creates a new EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams object
// with the ability to set a context for a request.
func NewEdgeApplicationInstanceStatusGetApplicationInterfaceTagsParamsWithContext(ctx context.Context) *EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams {
	return &EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams{
		Context: ctx,
	}
}

// NewEdgeApplicationInstanceStatusGetApplicationInterfaceTagsParamsWithHTTPClient creates a new EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeApplicationInstanceStatusGetApplicationInterfaceTagsParamsWithHTTPClient(client *http.Client) *EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams {
	return &EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams{
		HTTPClient: client,
	}
}

/*
EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams contains all the parameters to send to the API endpoint

	for the edge application instance status get application interface tags operation.

	Typically these are written to a http.Request.
*/
type EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* FilterObjID.

	   Object Id which tags are associated.
	*/
	FilterObjID *string

	/* FilterObjName.

	   Object name which tags are associated.
	*/
	FilterObjName *string

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edge application instance status get application interface tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams) WithDefaults() *EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge application instance status get application interface tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge application instance status get application interface tags params
func (o *EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams) WithTimeout(timeout time.Duration) *EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge application instance status get application interface tags params
func (o *EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge application instance status get application interface tags params
func (o *EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams) WithContext(ctx context.Context) *EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge application instance status get application interface tags params
func (o *EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge application instance status get application interface tags params
func (o *EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams) WithHTTPClient(client *http.Client) *EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge application instance status get application interface tags params
func (o *EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the edge application instance status get application interface tags params
func (o *EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams) WithXRequestID(xRequestID *string) *EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the edge application instance status get application interface tags params
func (o *EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithFilterObjID adds the filterObjID to the edge application instance status get application interface tags params
func (o *EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams) WithFilterObjID(filterObjID *string) *EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams {
	o.SetFilterObjID(filterObjID)
	return o
}

// SetFilterObjID adds the filterObjId to the edge application instance status get application interface tags params
func (o *EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams) SetFilterObjID(filterObjID *string) {
	o.FilterObjID = filterObjID
}

// WithFilterObjName adds the filterObjName to the edge application instance status get application interface tags params
func (o *EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams) WithFilterObjName(filterObjName *string) *EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams {
	o.SetFilterObjName(filterObjName)
	return o
}

// SetFilterObjName adds the filterObjName to the edge application instance status get application interface tags params
func (o *EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams) SetFilterObjName(filterObjName *string) {
	o.FilterObjName = filterObjName
}

// WithNextOrderBy adds the nextOrderBy to the edge application instance status get application interface tags params
func (o *EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams) WithNextOrderBy(nextOrderBy []string) *EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams {
	o.SetNextOrderBy(nextOrderBy)
	return o
}

// SetNextOrderBy adds the nextOrderBy to the edge application instance status get application interface tags params
func (o *EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams) SetNextOrderBy(nextOrderBy []string) {
	o.NextOrderBy = nextOrderBy
}

// WithNextPageNum adds the nextPageNum to the edge application instance status get application interface tags params
func (o *EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams) WithNextPageNum(nextPageNum *int64) *EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams {
	o.SetNextPageNum(nextPageNum)
	return o
}

// SetNextPageNum adds the nextPageNum to the edge application instance status get application interface tags params
func (o *EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams) SetNextPageNum(nextPageNum *int64) {
	o.NextPageNum = nextPageNum
}

// WithNextPageSize adds the nextPageSize to the edge application instance status get application interface tags params
func (o *EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams) WithNextPageSize(nextPageSize *int64) *EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams {
	o.SetNextPageSize(nextPageSize)
	return o
}

// SetNextPageSize adds the nextPageSize to the edge application instance status get application interface tags params
func (o *EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams) SetNextPageSize(nextPageSize *int64) {
	o.NextPageSize = nextPageSize
}

// WithNextPageToken adds the nextPageToken to the edge application instance status get application interface tags params
func (o *EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams) WithNextPageToken(nextPageToken *string) *EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams {
	o.SetNextPageToken(nextPageToken)
	return o
}

// SetNextPageToken adds the nextPageToken to the edge application instance status get application interface tags params
func (o *EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams) SetNextPageToken(nextPageToken *string) {
	o.NextPageToken = nextPageToken
}

// WithNextTotalPages adds the nextTotalPages to the edge application instance status get application interface tags params
func (o *EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams) WithNextTotalPages(nextTotalPages *int64) *EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams {
	o.SetNextTotalPages(nextTotalPages)
	return o
}

// SetNextTotalPages adds the nextTotalPages to the edge application instance status get application interface tags params
func (o *EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams) SetNextTotalPages(nextTotalPages *int64) {
	o.NextTotalPages = nextTotalPages
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.FilterObjID != nil {

		// query param filter.objId
		var qrFilterObjID string

		if o.FilterObjID != nil {
			qrFilterObjID = *o.FilterObjID
		}
		qFilterObjID := qrFilterObjID
		if qFilterObjID != "" {

			if err := r.SetQueryParam("filter.objId", qFilterObjID); err != nil {
				return err
			}
		}
	}

	if o.FilterObjName != nil {

		// query param filter.objName
		var qrFilterObjName string

		if o.FilterObjName != nil {
			qrFilterObjName = *o.FilterObjName
		}
		qFilterObjName := qrFilterObjName
		if qFilterObjName != "" {

			if err := r.SetQueryParam("filter.objName", qFilterObjName); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamEdgeApplicationInstanceStatusGetApplicationInterfaceTags binds the parameter next.orderBy
func (o *EdgeApplicationInstanceStatusGetApplicationInterfaceTagsParams) bindParamNextOrderBy(formats strfmt.Registry) []string {
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

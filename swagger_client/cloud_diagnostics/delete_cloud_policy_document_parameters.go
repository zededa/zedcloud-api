// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package cloud_diagnostics

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
)

// NewDeleteCloudPolicyDocumentParams creates a new DeleteCloudPolicyDocumentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteCloudPolicyDocumentParams() *DeleteCloudPolicyDocumentParams {
	return &DeleteCloudPolicyDocumentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCloudPolicyDocumentParamsWithTimeout creates a new DeleteCloudPolicyDocumentParams object
// with the ability to set a timeout on a request.
func NewDeleteCloudPolicyDocumentParamsWithTimeout(timeout time.Duration) *DeleteCloudPolicyDocumentParams {
	return &DeleteCloudPolicyDocumentParams{
		timeout: timeout,
	}
}

// NewDeleteCloudPolicyDocumentParamsWithContext creates a new DeleteCloudPolicyDocumentParams object
// with the ability to set a context for a request.
func NewDeleteCloudPolicyDocumentParamsWithContext(ctx context.Context) *DeleteCloudPolicyDocumentParams {
	return &DeleteCloudPolicyDocumentParams{
		Context: ctx,
	}
}

// NewDeleteCloudPolicyDocumentParamsWithHTTPClient creates a new DeleteCloudPolicyDocumentParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteCloudPolicyDocumentParamsWithHTTPClient(client *http.Client) *DeleteCloudPolicyDocumentParams {
	return &DeleteCloudPolicyDocumentParams{
		HTTPClient: client,
	}
}

/* DeleteCloudPolicyDocumentParams contains all the parameters to send to the API endpoint
   for the delete cloud policy document operation.

   Typically these are written to a http.Request.
*/
type DeleteCloudPolicyDocumentParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// FileURL.
	FileURL string

	// Policy.
	Policy *string

	/* RevisionCreatedBy.

	   User data: Created By
	*/
	RevisionCreatedBy *string

	/* RevisionCurr.

	   Current Database version of the record
	*/
	RevisionCurr *string

	/* RevisionPrev.

	   Previous
	*/
	RevisionPrev *string

	/* RevisionUpdatedBy.

	   User data: Updated By
	*/
	RevisionUpdatedBy *string

	// Version.
	Version *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete cloud policy document params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCloudPolicyDocumentParams) WithDefaults() *DeleteCloudPolicyDocumentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete cloud policy document params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCloudPolicyDocumentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete cloud policy document params
func (o *DeleteCloudPolicyDocumentParams) WithTimeout(timeout time.Duration) *DeleteCloudPolicyDocumentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete cloud policy document params
func (o *DeleteCloudPolicyDocumentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete cloud policy document params
func (o *DeleteCloudPolicyDocumentParams) WithContext(ctx context.Context) *DeleteCloudPolicyDocumentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete cloud policy document params
func (o *DeleteCloudPolicyDocumentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete cloud policy document params
func (o *DeleteCloudPolicyDocumentParams) WithHTTPClient(client *http.Client) *DeleteCloudPolicyDocumentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete cloud policy document params
func (o *DeleteCloudPolicyDocumentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the delete cloud policy document params
func (o *DeleteCloudPolicyDocumentParams) WithXRequestID(xRequestID *string) *DeleteCloudPolicyDocumentParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the delete cloud policy document params
func (o *DeleteCloudPolicyDocumentParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithFileURL adds the fileURL to the delete cloud policy document params
func (o *DeleteCloudPolicyDocumentParams) WithFileURL(fileURL string) *DeleteCloudPolicyDocumentParams {
	o.SetFileURL(fileURL)
	return o
}

// SetFileURL adds the fileUrl to the delete cloud policy document params
func (o *DeleteCloudPolicyDocumentParams) SetFileURL(fileURL string) {
	o.FileURL = fileURL
}

// WithPolicy adds the policy to the delete cloud policy document params
func (o *DeleteCloudPolicyDocumentParams) WithPolicy(policy *string) *DeleteCloudPolicyDocumentParams {
	o.SetPolicy(policy)
	return o
}

// SetPolicy adds the policy to the delete cloud policy document params
func (o *DeleteCloudPolicyDocumentParams) SetPolicy(policy *string) {
	o.Policy = policy
}

// WithRevisionCreatedBy adds the revisionCreatedBy to the delete cloud policy document params
func (o *DeleteCloudPolicyDocumentParams) WithRevisionCreatedBy(revisionCreatedBy *string) *DeleteCloudPolicyDocumentParams {
	o.SetRevisionCreatedBy(revisionCreatedBy)
	return o
}

// SetRevisionCreatedBy adds the revisionCreatedBy to the delete cloud policy document params
func (o *DeleteCloudPolicyDocumentParams) SetRevisionCreatedBy(revisionCreatedBy *string) {
	o.RevisionCreatedBy = revisionCreatedBy
}

// WithRevisionCurr adds the revisionCurr to the delete cloud policy document params
func (o *DeleteCloudPolicyDocumentParams) WithRevisionCurr(revisionCurr *string) *DeleteCloudPolicyDocumentParams {
	o.SetRevisionCurr(revisionCurr)
	return o
}

// SetRevisionCurr adds the revisionCurr to the delete cloud policy document params
func (o *DeleteCloudPolicyDocumentParams) SetRevisionCurr(revisionCurr *string) {
	o.RevisionCurr = revisionCurr
}

// WithRevisionPrev adds the revisionPrev to the delete cloud policy document params
func (o *DeleteCloudPolicyDocumentParams) WithRevisionPrev(revisionPrev *string) *DeleteCloudPolicyDocumentParams {
	o.SetRevisionPrev(revisionPrev)
	return o
}

// SetRevisionPrev adds the revisionPrev to the delete cloud policy document params
func (o *DeleteCloudPolicyDocumentParams) SetRevisionPrev(revisionPrev *string) {
	o.RevisionPrev = revisionPrev
}

// WithRevisionUpdatedBy adds the revisionUpdatedBy to the delete cloud policy document params
func (o *DeleteCloudPolicyDocumentParams) WithRevisionUpdatedBy(revisionUpdatedBy *string) *DeleteCloudPolicyDocumentParams {
	o.SetRevisionUpdatedBy(revisionUpdatedBy)
	return o
}

// SetRevisionUpdatedBy adds the revisionUpdatedBy to the delete cloud policy document params
func (o *DeleteCloudPolicyDocumentParams) SetRevisionUpdatedBy(revisionUpdatedBy *string) {
	o.RevisionUpdatedBy = revisionUpdatedBy
}

// WithVersion adds the version to the delete cloud policy document params
func (o *DeleteCloudPolicyDocumentParams) WithVersion(version *string) *DeleteCloudPolicyDocumentParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete cloud policy document params
func (o *DeleteCloudPolicyDocumentParams) SetVersion(version *string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCloudPolicyDocumentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param fileURL
	if err := r.SetPathParam("fileURL", o.FileURL); err != nil {
		return err
	}

	if o.Policy != nil {

		// query param policy
		var qrPolicy string

		if o.Policy != nil {
			qrPolicy = *o.Policy
		}
		qPolicy := qrPolicy
		if qPolicy != "" {

			if err := r.SetQueryParam("policy", qPolicy); err != nil {
				return err
			}
		}
	}

	if o.RevisionCreatedBy != nil {

		// query param revision.createdBy
		var qrRevisionCreatedBy string

		if o.RevisionCreatedBy != nil {
			qrRevisionCreatedBy = *o.RevisionCreatedBy
		}
		qRevisionCreatedBy := qrRevisionCreatedBy
		if qRevisionCreatedBy != "" {

			if err := r.SetQueryParam("revision.createdBy", qRevisionCreatedBy); err != nil {
				return err
			}
		}
	}

	if o.RevisionCurr != nil {

		// query param revision.curr
		var qrRevisionCurr string

		if o.RevisionCurr != nil {
			qrRevisionCurr = *o.RevisionCurr
		}
		qRevisionCurr := qrRevisionCurr
		if qRevisionCurr != "" {

			if err := r.SetQueryParam("revision.curr", qRevisionCurr); err != nil {
				return err
			}
		}
	}

	if o.RevisionPrev != nil {

		// query param revision.prev
		var qrRevisionPrev string

		if o.RevisionPrev != nil {
			qrRevisionPrev = *o.RevisionPrev
		}
		qRevisionPrev := qrRevisionPrev
		if qRevisionPrev != "" {

			if err := r.SetQueryParam("revision.prev", qRevisionPrev); err != nil {
				return err
			}
		}
	}

	if o.RevisionUpdatedBy != nil {

		// query param revision.updatedBy
		var qrRevisionUpdatedBy string

		if o.RevisionUpdatedBy != nil {
			qrRevisionUpdatedBy = *o.RevisionUpdatedBy
		}
		qRevisionUpdatedBy := qrRevisionUpdatedBy
		if qRevisionUpdatedBy != "" {

			if err := r.SetQueryParam("revision.updatedBy", qRevisionUpdatedBy); err != nil {
				return err
			}
		}
	}

	if o.Version != nil {

		// query param version
		var qrVersion string

		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := qrVersion
		if qVersion != "" {

			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

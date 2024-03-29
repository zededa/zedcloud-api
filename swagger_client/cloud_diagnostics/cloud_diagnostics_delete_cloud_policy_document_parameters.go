// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
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

// NewCloudDiagnosticsDeleteCloudPolicyDocumentParams creates a new CloudDiagnosticsDeleteCloudPolicyDocumentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCloudDiagnosticsDeleteCloudPolicyDocumentParams() *CloudDiagnosticsDeleteCloudPolicyDocumentParams {
	return &CloudDiagnosticsDeleteCloudPolicyDocumentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCloudDiagnosticsDeleteCloudPolicyDocumentParamsWithTimeout creates a new CloudDiagnosticsDeleteCloudPolicyDocumentParams object
// with the ability to set a timeout on a request.
func NewCloudDiagnosticsDeleteCloudPolicyDocumentParamsWithTimeout(timeout time.Duration) *CloudDiagnosticsDeleteCloudPolicyDocumentParams {
	return &CloudDiagnosticsDeleteCloudPolicyDocumentParams{
		timeout: timeout,
	}
}

// NewCloudDiagnosticsDeleteCloudPolicyDocumentParamsWithContext creates a new CloudDiagnosticsDeleteCloudPolicyDocumentParams object
// with the ability to set a context for a request.
func NewCloudDiagnosticsDeleteCloudPolicyDocumentParamsWithContext(ctx context.Context) *CloudDiagnosticsDeleteCloudPolicyDocumentParams {
	return &CloudDiagnosticsDeleteCloudPolicyDocumentParams{
		Context: ctx,
	}
}

// NewCloudDiagnosticsDeleteCloudPolicyDocumentParamsWithHTTPClient creates a new CloudDiagnosticsDeleteCloudPolicyDocumentParams object
// with the ability to set a custom HTTPClient for a request.
func NewCloudDiagnosticsDeleteCloudPolicyDocumentParamsWithHTTPClient(client *http.Client) *CloudDiagnosticsDeleteCloudPolicyDocumentParams {
	return &CloudDiagnosticsDeleteCloudPolicyDocumentParams{
		HTTPClient: client,
	}
}

/*
CloudDiagnosticsDeleteCloudPolicyDocumentParams contains all the parameters to send to the API endpoint

	for the cloud diagnostics delete cloud policy document operation.

	Typically these are written to a http.Request.
*/
type CloudDiagnosticsDeleteCloudPolicyDocumentParams struct {

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

// WithDefaults hydrates default values in the cloud diagnostics delete cloud policy document params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CloudDiagnosticsDeleteCloudPolicyDocumentParams) WithDefaults() *CloudDiagnosticsDeleteCloudPolicyDocumentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cloud diagnostics delete cloud policy document params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CloudDiagnosticsDeleteCloudPolicyDocumentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cloud diagnostics delete cloud policy document params
func (o *CloudDiagnosticsDeleteCloudPolicyDocumentParams) WithTimeout(timeout time.Duration) *CloudDiagnosticsDeleteCloudPolicyDocumentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cloud diagnostics delete cloud policy document params
func (o *CloudDiagnosticsDeleteCloudPolicyDocumentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cloud diagnostics delete cloud policy document params
func (o *CloudDiagnosticsDeleteCloudPolicyDocumentParams) WithContext(ctx context.Context) *CloudDiagnosticsDeleteCloudPolicyDocumentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cloud diagnostics delete cloud policy document params
func (o *CloudDiagnosticsDeleteCloudPolicyDocumentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cloud diagnostics delete cloud policy document params
func (o *CloudDiagnosticsDeleteCloudPolicyDocumentParams) WithHTTPClient(client *http.Client) *CloudDiagnosticsDeleteCloudPolicyDocumentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cloud diagnostics delete cloud policy document params
func (o *CloudDiagnosticsDeleteCloudPolicyDocumentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the cloud diagnostics delete cloud policy document params
func (o *CloudDiagnosticsDeleteCloudPolicyDocumentParams) WithXRequestID(xRequestID *string) *CloudDiagnosticsDeleteCloudPolicyDocumentParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the cloud diagnostics delete cloud policy document params
func (o *CloudDiagnosticsDeleteCloudPolicyDocumentParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithFileURL adds the fileURL to the cloud diagnostics delete cloud policy document params
func (o *CloudDiagnosticsDeleteCloudPolicyDocumentParams) WithFileURL(fileURL string) *CloudDiagnosticsDeleteCloudPolicyDocumentParams {
	o.SetFileURL(fileURL)
	return o
}

// SetFileURL adds the fileUrl to the cloud diagnostics delete cloud policy document params
func (o *CloudDiagnosticsDeleteCloudPolicyDocumentParams) SetFileURL(fileURL string) {
	o.FileURL = fileURL
}

// WithPolicy adds the policy to the cloud diagnostics delete cloud policy document params
func (o *CloudDiagnosticsDeleteCloudPolicyDocumentParams) WithPolicy(policy *string) *CloudDiagnosticsDeleteCloudPolicyDocumentParams {
	o.SetPolicy(policy)
	return o
}

// SetPolicy adds the policy to the cloud diagnostics delete cloud policy document params
func (o *CloudDiagnosticsDeleteCloudPolicyDocumentParams) SetPolicy(policy *string) {
	o.Policy = policy
}

// WithRevisionCreatedBy adds the revisionCreatedBy to the cloud diagnostics delete cloud policy document params
func (o *CloudDiagnosticsDeleteCloudPolicyDocumentParams) WithRevisionCreatedBy(revisionCreatedBy *string) *CloudDiagnosticsDeleteCloudPolicyDocumentParams {
	o.SetRevisionCreatedBy(revisionCreatedBy)
	return o
}

// SetRevisionCreatedBy adds the revisionCreatedBy to the cloud diagnostics delete cloud policy document params
func (o *CloudDiagnosticsDeleteCloudPolicyDocumentParams) SetRevisionCreatedBy(revisionCreatedBy *string) {
	o.RevisionCreatedBy = revisionCreatedBy
}

// WithRevisionCurr adds the revisionCurr to the cloud diagnostics delete cloud policy document params
func (o *CloudDiagnosticsDeleteCloudPolicyDocumentParams) WithRevisionCurr(revisionCurr *string) *CloudDiagnosticsDeleteCloudPolicyDocumentParams {
	o.SetRevisionCurr(revisionCurr)
	return o
}

// SetRevisionCurr adds the revisionCurr to the cloud diagnostics delete cloud policy document params
func (o *CloudDiagnosticsDeleteCloudPolicyDocumentParams) SetRevisionCurr(revisionCurr *string) {
	o.RevisionCurr = revisionCurr
}

// WithRevisionPrev adds the revisionPrev to the cloud diagnostics delete cloud policy document params
func (o *CloudDiagnosticsDeleteCloudPolicyDocumentParams) WithRevisionPrev(revisionPrev *string) *CloudDiagnosticsDeleteCloudPolicyDocumentParams {
	o.SetRevisionPrev(revisionPrev)
	return o
}

// SetRevisionPrev adds the revisionPrev to the cloud diagnostics delete cloud policy document params
func (o *CloudDiagnosticsDeleteCloudPolicyDocumentParams) SetRevisionPrev(revisionPrev *string) {
	o.RevisionPrev = revisionPrev
}

// WithRevisionUpdatedBy adds the revisionUpdatedBy to the cloud diagnostics delete cloud policy document params
func (o *CloudDiagnosticsDeleteCloudPolicyDocumentParams) WithRevisionUpdatedBy(revisionUpdatedBy *string) *CloudDiagnosticsDeleteCloudPolicyDocumentParams {
	o.SetRevisionUpdatedBy(revisionUpdatedBy)
	return o
}

// SetRevisionUpdatedBy adds the revisionUpdatedBy to the cloud diagnostics delete cloud policy document params
func (o *CloudDiagnosticsDeleteCloudPolicyDocumentParams) SetRevisionUpdatedBy(revisionUpdatedBy *string) {
	o.RevisionUpdatedBy = revisionUpdatedBy
}

// WithVersion adds the version to the cloud diagnostics delete cloud policy document params
func (o *CloudDiagnosticsDeleteCloudPolicyDocumentParams) WithVersion(version *string) *CloudDiagnosticsDeleteCloudPolicyDocumentParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the cloud diagnostics delete cloud policy document params
func (o *CloudDiagnosticsDeleteCloudPolicyDocumentParams) SetVersion(version *string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *CloudDiagnosticsDeleteCloudPolicyDocumentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package swagger_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// JobStatus job status.
//
// swagger:model JobStatus
type JobStatus string

func NewJobStatus(value JobStatus) *JobStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated JobStatus.
func (m JobStatus) Pointer() *JobStatus {
	return &m
}

const (

	// JobStatusJOBSTATUSUNSPECIFIED captures enum value "JOB_STATUS_UNSPECIFIED"
	JobStatusJOBSTATUSUNSPECIFIED JobStatus = "JOB_STATUS_UNSPECIFIED"

	// JobStatusJOBSTATUSINITIALIZED captures enum value "JOB_STATUS_INITIALIZED"
	JobStatusJOBSTATUSINITIALIZED JobStatus = "JOB_STATUS_INITIALIZED"

	// JobStatusJOBSTATUSREADY captures enum value "JOB_STATUS_READY"
	JobStatusJOBSTATUSREADY JobStatus = "JOB_STATUS_READY"

	// JobStatusJOBSTATUSINPROGRESS captures enum value "JOB_STATUS_INPROGRESS"
	JobStatusJOBSTATUSINPROGRESS JobStatus = "JOB_STATUS_INPROGRESS"

	// JobStatusJOBSTATUSCOMPLETED captures enum value "JOB_STATUS_COMPLETED"
	JobStatusJOBSTATUSCOMPLETED JobStatus = "JOB_STATUS_COMPLETED"

	// JobStatusJOBSTATUSFAILED captures enum value "JOB_STATUS_FAILED"
	JobStatusJOBSTATUSFAILED JobStatus = "JOB_STATUS_FAILED"

	// JobStatusJOBSTATUSABORTED captures enum value "JOB_STATUS_ABORTED"
	JobStatusJOBSTATUSABORTED JobStatus = "JOB_STATUS_ABORTED"
)

// for schema
var jobStatusEnum []interface{}

func init() {
	var res []JobStatus
	if err := json.Unmarshal([]byte(`["JOB_STATUS_UNSPECIFIED","JOB_STATUS_INITIALIZED","JOB_STATUS_READY","JOB_STATUS_INPROGRESS","JOB_STATUS_COMPLETED","JOB_STATUS_FAILED","JOB_STATUS_ABORTED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		jobStatusEnum = append(jobStatusEnum, v)
	}
}

func (m JobStatus) validateJobStatusEnum(path, location string, value JobStatus) error {
	if err := validate.EnumCase(path, location, value, jobStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this job status
func (m JobStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateJobStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this job status based on context it is used
func (m JobStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

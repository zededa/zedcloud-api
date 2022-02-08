// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

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

// DatastoreStatus Datastore status
//
// swagger:model DatastoreStatus
type DatastoreStatus string

func NewDatastoreStatus(value DatastoreStatus) *DatastoreStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated DatastoreStatus.
func (m DatastoreStatus) Pointer() *DatastoreStatus {
	return &m
}

const (

	// DatastoreStatusDATASTORESTATUSUNSPECIFIED captures enum value "DATASTORE_STATUS_UNSPECIFIED"
	DatastoreStatusDATASTORESTATUSUNSPECIFIED DatastoreStatus = "DATASTORE_STATUS_UNSPECIFIED"

	// DatastoreStatusDATASTORESTATUSCREATED captures enum value "DATASTORE_STATUS_CREATED"
	DatastoreStatusDATASTORESTATUSCREATED DatastoreStatus = "DATASTORE_STATUS_CREATED"

	// DatastoreStatusDATASTORESTATUSVERIFYING captures enum value "DATASTORE_STATUS_VERIFYING"
	DatastoreStatusDATASTORESTATUSVERIFYING DatastoreStatus = "DATASTORE_STATUS_VERIFYING"

	// DatastoreStatusDATASTORESTATUSACTIVE captures enum value "DATASTORE_STATUS_ACTIVE"
	DatastoreStatusDATASTORESTATUSACTIVE DatastoreStatus = "DATASTORE_STATUS_ACTIVE"

	// DatastoreStatusDATASTORESTATUSINACTIVE captures enum value "DATASTORE_STATUS_INACTIVE"
	DatastoreStatusDATASTORESTATUSINACTIVE DatastoreStatus = "DATASTORE_STATUS_INACTIVE"

	// DatastoreStatusDATASTORESTATUSFAILED captures enum value "DATASTORE_STATUS_FAILED"
	DatastoreStatusDATASTORESTATUSFAILED DatastoreStatus = "DATASTORE_STATUS_FAILED"
)

// for schema
var datastoreStatusEnum []interface{}

func init() {
	var res []DatastoreStatus
	if err := json.Unmarshal([]byte(`["DATASTORE_STATUS_UNSPECIFIED","DATASTORE_STATUS_CREATED","DATASTORE_STATUS_VERIFYING","DATASTORE_STATUS_ACTIVE","DATASTORE_STATUS_INACTIVE","DATASTORE_STATUS_FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		datastoreStatusEnum = append(datastoreStatusEnum, v)
	}
}

func (m DatastoreStatus) validateDatastoreStatusEnum(path, location string, value DatastoreStatus) error {
	if err := validate.EnumCase(path, location, value, datastoreStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this datastore status
func (m DatastoreStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDatastoreStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this datastore status based on context it is used
func (m DatastoreStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

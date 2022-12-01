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

// PhysicalStorageStatus physical storage status
//
// swagger:model PhysicalStorageStatus
type PhysicalStorageStatus string

func NewPhysicalStorageStatus(value PhysicalStorageStatus) *PhysicalStorageStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated PhysicalStorageStatus.
func (m PhysicalStorageStatus) Pointer() *PhysicalStorageStatus {
	return &m
}

const (

	// PhysicalStorageStatusPHYSICALSTORAGESTATUSUNSPECIFIED captures enum value "PHYSICAL_STORAGE_STATUS_UNSPECIFIED"
	PhysicalStorageStatusPHYSICALSTORAGESTATUSUNSPECIFIED PhysicalStorageStatus = "PHYSICAL_STORAGE_STATUS_UNSPECIFIED"

	// PhysicalStorageStatusPHYSICALSTORAGESTATUSONLINE captures enum value "PHYSICAL_STORAGE_STATUS_ONLINE"
	PhysicalStorageStatusPHYSICALSTORAGESTATUSONLINE PhysicalStorageStatus = "PHYSICAL_STORAGE_STATUS_ONLINE"

	// PhysicalStorageStatusPHYSICALSTORAGESTATUSDEGRADED captures enum value "PHYSICAL_STORAGE_STATUS_DEGRADED"
	PhysicalStorageStatusPHYSICALSTORAGESTATUSDEGRADED PhysicalStorageStatus = "PHYSICAL_STORAGE_STATUS_DEGRADED"

	// PhysicalStorageStatusPHYSICALSTORAGESTATUSFAULTED captures enum value "PHYSICAL_STORAGE_STATUS_FAULTED"
	PhysicalStorageStatusPHYSICALSTORAGESTATUSFAULTED PhysicalStorageStatus = "PHYSICAL_STORAGE_STATUS_FAULTED"

	// PhysicalStorageStatusPHYSICALSTORAGESTATUSOFFLINE captures enum value "PHYSICAL_STORAGE_STATUS_OFFLINE"
	PhysicalStorageStatusPHYSICALSTORAGESTATUSOFFLINE PhysicalStorageStatus = "PHYSICAL_STORAGE_STATUS_OFFLINE"

	// PhysicalStorageStatusPHYSICALSTORAGESTATUSUNAVAIL captures enum value "PHYSICAL_STORAGE_STATUS_UNAVAIL"
	PhysicalStorageStatusPHYSICALSTORAGESTATUSUNAVAIL PhysicalStorageStatus = "PHYSICAL_STORAGE_STATUS_UNAVAIL"

	// PhysicalStorageStatusPHYSICALSTORAGESTATUSREMOVED captures enum value "PHYSICAL_STORAGE_STATUS_REMOVED"
	PhysicalStorageStatusPHYSICALSTORAGESTATUSREMOVED PhysicalStorageStatus = "PHYSICAL_STORAGE_STATUS_REMOVED"

	// PhysicalStorageStatusPHYSICALSTORAGESTATUSSUSPENDED captures enum value "PHYSICAL_STORAGE_STATUS_SUSPENDED"
	PhysicalStorageStatusPHYSICALSTORAGESTATUSSUSPENDED PhysicalStorageStatus = "PHYSICAL_STORAGE_STATUS_SUSPENDED"
)

// for schema
var physicalStorageStatusEnum []interface{}

func init() {
	var res []PhysicalStorageStatus
	if err := json.Unmarshal([]byte(`["PHYSICAL_STORAGE_STATUS_UNSPECIFIED","PHYSICAL_STORAGE_STATUS_ONLINE","PHYSICAL_STORAGE_STATUS_DEGRADED","PHYSICAL_STORAGE_STATUS_FAULTED","PHYSICAL_STORAGE_STATUS_OFFLINE","PHYSICAL_STORAGE_STATUS_UNAVAIL","PHYSICAL_STORAGE_STATUS_REMOVED","PHYSICAL_STORAGE_STATUS_SUSPENDED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		physicalStorageStatusEnum = append(physicalStorageStatusEnum, v)
	}
}

func (m PhysicalStorageStatus) validatePhysicalStorageStatusEnum(path, location string, value PhysicalStorageStatus) error {
	if err := validate.EnumCase(path, location, value, physicalStorageStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this physical storage status
func (m PhysicalStorageStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePhysicalStorageStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this physical storage status based on context it is used
func (m PhysicalStorageStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

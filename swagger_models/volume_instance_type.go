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

// VolumeInstanceType volume instance type
//
// swagger:model VolumeInstanceType
type VolumeInstanceType string

func NewVolumeInstanceType(value VolumeInstanceType) *VolumeInstanceType {
	v := value
	return &v
}

const (

	// VolumeInstanceTypeVOLUMEINSTANCETYPEUNSPECIFIED captures enum value "VOLUME_INSTANCE_TYPE_UNSPECIFIED"
	VolumeInstanceTypeVOLUMEINSTANCETYPEUNSPECIFIED VolumeInstanceType = "VOLUME_INSTANCE_TYPE_UNSPECIFIED"

	// VolumeInstanceTypeVOLUMEINSTANCETYPEEMPTYDIR captures enum value "VOLUME_INSTANCE_TYPE_EMPTYDIR"
	VolumeInstanceTypeVOLUMEINSTANCETYPEEMPTYDIR VolumeInstanceType = "VOLUME_INSTANCE_TYPE_EMPTYDIR"

	// VolumeInstanceTypeVOLUMEINSTANCETYPEBLOCKSTORAGE captures enum value "VOLUME_INSTANCE_TYPE_BLOCKSTORAGE"
	VolumeInstanceTypeVOLUMEINSTANCETYPEBLOCKSTORAGE VolumeInstanceType = "VOLUME_INSTANCE_TYPE_BLOCKSTORAGE"

	// VolumeInstanceTypeVOLUMEINSTANCETYPEHOSTFS captures enum value "VOLUME_INSTANCE_TYPE_HOSTFS"
	VolumeInstanceTypeVOLUMEINSTANCETYPEHOSTFS VolumeInstanceType = "VOLUME_INSTANCE_TYPE_HOSTFS"

	// VolumeInstanceTypeVOLUMEINSTANCETYPETMPFS captures enum value "VOLUME_INSTANCE_TYPE_TMPFS"
	VolumeInstanceTypeVOLUMEINSTANCETYPETMPFS VolumeInstanceType = "VOLUME_INSTANCE_TYPE_TMPFS"

	// VolumeInstanceTypeVOLUMEINSTANCETYPESECRET captures enum value "VOLUME_INSTANCE_TYPE_SECRET"
	VolumeInstanceTypeVOLUMEINSTANCETYPESECRET VolumeInstanceType = "VOLUME_INSTANCE_TYPE_SECRET"

	// VolumeInstanceTypeVOLUMEINSTANCETYPENFS captures enum value "VOLUME_INSTANCE_TYPE_NFS"
	VolumeInstanceTypeVOLUMEINSTANCETYPENFS VolumeInstanceType = "VOLUME_INSTANCE_TYPE_NFS"

	// VolumeInstanceTypeVOLUMEINSTANCETYPEAWSBLOCKSTORAGE captures enum value "VOLUME_INSTANCE_TYPE_AWS_BLOCK_STORAGE"
	VolumeInstanceTypeVOLUMEINSTANCETYPEAWSBLOCKSTORAGE VolumeInstanceType = "VOLUME_INSTANCE_TYPE_AWS_BLOCK_STORAGE"

	// VolumeInstanceTypeVOLUMEINSTANCETYPECONTENTTREE captures enum value "VOLUME_INSTANCE_TYPE_CONTENT_TREE"
	VolumeInstanceTypeVOLUMEINSTANCETYPECONTENTTREE VolumeInstanceType = "VOLUME_INSTANCE_TYPE_CONTENT_TREE"
)

// for schema
var volumeInstanceTypeEnum []interface{}

func init() {
	var res []VolumeInstanceType
	if err := json.Unmarshal([]byte(`["VOLUME_INSTANCE_TYPE_UNSPECIFIED","VOLUME_INSTANCE_TYPE_EMPTYDIR","VOLUME_INSTANCE_TYPE_BLOCKSTORAGE","VOLUME_INSTANCE_TYPE_HOSTFS","VOLUME_INSTANCE_TYPE_TMPFS","VOLUME_INSTANCE_TYPE_SECRET","VOLUME_INSTANCE_TYPE_NFS","VOLUME_INSTANCE_TYPE_AWS_BLOCK_STORAGE","VOLUME_INSTANCE_TYPE_CONTENT_TREE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		volumeInstanceTypeEnum = append(volumeInstanceTypeEnum, v)
	}
}

func (m VolumeInstanceType) validateVolumeInstanceTypeEnum(path, location string, value VolumeInstanceType) error {
	if err := validate.EnumCase(path, location, value, volumeInstanceTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this volume instance type
func (m VolumeInstanceType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateVolumeInstanceTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this volume instance type based on context it is used
func (m VolumeInstanceType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

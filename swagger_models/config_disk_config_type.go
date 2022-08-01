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

// ConfigDiskConfigType DiskConfigType is the desired configuration of disks
//
// swagger:model configDiskConfigType
type ConfigDiskConfigType string

func NewConfigDiskConfigType(value ConfigDiskConfigType) *ConfigDiskConfigType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ConfigDiskConfigType.
func (m ConfigDiskConfigType) Pointer() *ConfigDiskConfigType {
	return &m
}

const (

	// ConfigDiskConfigTypeDISKCONFIGTYPEUNSPECIFIED captures enum value "DISK_CONFIG_TYPE_UNSPECIFIED"
	ConfigDiskConfigTypeDISKCONFIGTYPEUNSPECIFIED ConfigDiskConfigType = "DISK_CONFIG_TYPE_UNSPECIFIED"

	// ConfigDiskConfigTypeDISKCONFIGTYPEEVEOS captures enum value "DISK_CONFIG_TYPE_EVEOS"
	ConfigDiskConfigTypeDISKCONFIGTYPEEVEOS ConfigDiskConfigType = "DISK_CONFIG_TYPE_EVEOS"

	// ConfigDiskConfigTypeDISKCONFIGTYPEPERSIST captures enum value "DISK_CONFIG_TYPE_PERSIST"
	ConfigDiskConfigTypeDISKCONFIGTYPEPERSIST ConfigDiskConfigType = "DISK_CONFIG_TYPE_PERSIST"

	// ConfigDiskConfigTypeDISKCONFIGTYPEZFSONLINE captures enum value "DISK_CONFIG_TYPE_ZFS_ONLINE"
	ConfigDiskConfigTypeDISKCONFIGTYPEZFSONLINE ConfigDiskConfigType = "DISK_CONFIG_TYPE_ZFS_ONLINE"

	// ConfigDiskConfigTypeDISKCONFIGTYPEZFSOFFLINE captures enum value "DISK_CONFIG_TYPE_ZFS_OFFLINE"
	ConfigDiskConfigTypeDISKCONFIGTYPEZFSOFFLINE ConfigDiskConfigType = "DISK_CONFIG_TYPE_ZFS_OFFLINE"

	// ConfigDiskConfigTypeDISKCONFIGTYPEAPPDIRECT captures enum value "DISK_CONFIG_TYPE_APPDIRECT"
	ConfigDiskConfigTypeDISKCONFIGTYPEAPPDIRECT ConfigDiskConfigType = "DISK_CONFIG_TYPE_APPDIRECT"

	// ConfigDiskConfigTypeDISKCONFIGTYPEUNUSED captures enum value "DISK_CONFIG_TYPE_UNUSED"
	ConfigDiskConfigTypeDISKCONFIGTYPEUNUSED ConfigDiskConfigType = "DISK_CONFIG_TYPE_UNUSED"
)

// for schema
var configDiskConfigTypeEnum []interface{}

func init() {
	var res []ConfigDiskConfigType
	if err := json.Unmarshal([]byte(`["DISK_CONFIG_TYPE_UNSPECIFIED","DISK_CONFIG_TYPE_EVEOS","DISK_CONFIG_TYPE_PERSIST","DISK_CONFIG_TYPE_ZFS_ONLINE","DISK_CONFIG_TYPE_ZFS_OFFLINE","DISK_CONFIG_TYPE_APPDIRECT","DISK_CONFIG_TYPE_UNUSED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		configDiskConfigTypeEnum = append(configDiskConfigTypeEnum, v)
	}
}

func (m ConfigDiskConfigType) validateConfigDiskConfigTypeEnum(path, location string, value ConfigDiskConfigType) error {
	if err := validate.EnumCase(path, location, value, configDiskConfigTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this config disk config type
func (m ConfigDiskConfigType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateConfigDiskConfigTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this config disk config type based on context it is used
func (m ConfigDiskConfigType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

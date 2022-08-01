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

// DeploymentType deployment type
//
// swagger:model DeploymentType
type DeploymentType string

func NewDeploymentType(value DeploymentType) *DeploymentType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated DeploymentType.
func (m DeploymentType) Pointer() *DeploymentType {
	return &m
}

const (

	// DeploymentTypeDEPLOYMENTTYPEUNSPECIFIED captures enum value "DEPLOYMENT_TYPE_UNSPECIFIED"
	DeploymentTypeDEPLOYMENTTYPEUNSPECIFIED DeploymentType = "DEPLOYMENT_TYPE_UNSPECIFIED"

	// DeploymentTypeDEPLOYMENTTYPESTANDALONE captures enum value "DEPLOYMENT_TYPE_STAND_ALONE"
	DeploymentTypeDEPLOYMENTTYPESTANDALONE DeploymentType = "DEPLOYMENT_TYPE_STAND_ALONE"

	// DeploymentTypeDEPLOYMENTTYPEAZURE captures enum value "DEPLOYMENT_TYPE_AZURE"
	DeploymentTypeDEPLOYMENTTYPEAZURE DeploymentType = "DEPLOYMENT_TYPE_AZURE"

	// DeploymentTypeDEPLOYMENTTYPEK3S captures enum value "DEPLOYMENT_TYPE_K3S"
	DeploymentTypeDEPLOYMENTTYPEK3S DeploymentType = "DEPLOYMENT_TYPE_K3S"

	// DeploymentTypeDEPLOYMENTTYPEAWS captures enum value "DEPLOYMENT_TYPE_AWS"
	DeploymentTypeDEPLOYMENTTYPEAWS DeploymentType = "DEPLOYMENT_TYPE_AWS"

	// DeploymentTypeDEPLOYMENTTYPEK3SAZURE captures enum value "DEPLOYMENT_TYPE_K3S_AZURE"
	DeploymentTypeDEPLOYMENTTYPEK3SAZURE DeploymentType = "DEPLOYMENT_TYPE_K3S_AZURE"

	// DeploymentTypeDEPLOYMENTTYPEK3SAWS captures enum value "DEPLOYMENT_TYPE_K3S_AWS"
	DeploymentTypeDEPLOYMENTTYPEK3SAWS DeploymentType = "DEPLOYMENT_TYPE_K3S_AWS"

	// DeploymentTypeDEPLOYMENTTYPEVMWAREVCE captures enum value "DEPLOYMENT_TYPE_VMWARE_VCE"
	DeploymentTypeDEPLOYMENTTYPEVMWAREVCE DeploymentType = "DEPLOYMENT_TYPE_VMWARE_VCE"
)

// for schema
var deploymentTypeEnum []interface{}

func init() {
	var res []DeploymentType
	if err := json.Unmarshal([]byte(`["DEPLOYMENT_TYPE_UNSPECIFIED","DEPLOYMENT_TYPE_STAND_ALONE","DEPLOYMENT_TYPE_AZURE","DEPLOYMENT_TYPE_K3S","DEPLOYMENT_TYPE_AWS","DEPLOYMENT_TYPE_K3S_AZURE","DEPLOYMENT_TYPE_K3S_AWS","DEPLOYMENT_TYPE_VMWARE_VCE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deploymentTypeEnum = append(deploymentTypeEnum, v)
	}
}

func (m DeploymentType) validateDeploymentTypeEnum(path, location string, value DeploymentType) error {
	if err := validate.EnumCase(path, location, value, deploymentTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this deployment type
func (m DeploymentType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDeploymentTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this deployment type based on context it is used
func (m DeploymentType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

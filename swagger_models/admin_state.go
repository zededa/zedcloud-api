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

// AdminState Configured state of an object
//
// - ADMIN_STATE_CREATED: Entity Created in the controller
//   - ADMIN_STATE_DELETED: Entity Deleted in the controller
//   - ADMIN_STATE_ACTIVE: Entity Activated in the controller
//   - ADMIN_STATE_INACTIVE: Entity Deactivated in the controller
//   - ADMIN_STATE_REGISTERED: Specific to Edge-node - Edge-node Registered with the controller
//   - ADMIN_STATE_ARCHIVED: Entity Archived in the controller
//
// swagger:model AdminState
type AdminState string

func NewAdminState(value AdminState) *AdminState {
	return &value
}

// Pointer returns a pointer to a freshly-allocated AdminState.
func (m AdminState) Pointer() *AdminState {
	return &m
}

const (

	// AdminStateADMINSTATEUNSPECIFIED captures enum value "ADMIN_STATE_UNSPECIFIED"
	AdminStateADMINSTATEUNSPECIFIED AdminState = "ADMIN_STATE_UNSPECIFIED"

	// AdminStateADMINSTATECREATED captures enum value "ADMIN_STATE_CREATED"
	AdminStateADMINSTATECREATED AdminState = "ADMIN_STATE_CREATED"

	// AdminStateADMINSTATEDELETED captures enum value "ADMIN_STATE_DELETED"
	AdminStateADMINSTATEDELETED AdminState = "ADMIN_STATE_DELETED"

	// AdminStateADMINSTATEACTIVE captures enum value "ADMIN_STATE_ACTIVE"
	AdminStateADMINSTATEACTIVE AdminState = "ADMIN_STATE_ACTIVE"

	// AdminStateADMINSTATEINACTIVE captures enum value "ADMIN_STATE_INACTIVE"
	AdminStateADMINSTATEINACTIVE AdminState = "ADMIN_STATE_INACTIVE"

	// AdminStateADMINSTATEREGISTERED captures enum value "ADMIN_STATE_REGISTERED"
	AdminStateADMINSTATEREGISTERED AdminState = "ADMIN_STATE_REGISTERED"

	// AdminStateADMINSTATEARCHIVED captures enum value "ADMIN_STATE_ARCHIVED"
	AdminStateADMINSTATEARCHIVED AdminState = "ADMIN_STATE_ARCHIVED"
)

// for schema
var adminStateEnum []interface{}

func init() {
	var res []AdminState
	if err := json.Unmarshal([]byte(`["ADMIN_STATE_UNSPECIFIED","ADMIN_STATE_CREATED","ADMIN_STATE_DELETED","ADMIN_STATE_ACTIVE","ADMIN_STATE_INACTIVE","ADMIN_STATE_REGISTERED","ADMIN_STATE_ARCHIVED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		adminStateEnum = append(adminStateEnum, v)
	}
}

func (m AdminState) validateAdminStateEnum(path, location string, value AdminState) error {
	if err := validate.EnumCase(path, location, value, adminStateEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this admin state
func (m AdminState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAdminStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this admin state based on context it is used
func (m AdminState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

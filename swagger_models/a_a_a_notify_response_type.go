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

// AAANotifyResponseType a a a notify response type
//
// swagger:model AAANotifyResponseType
type AAANotifyResponseType string

func NewAAANotifyResponseType(value AAANotifyResponseType) *AAANotifyResponseType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated AAANotifyResponseType.
func (m AAANotifyResponseType) Pointer() *AAANotifyResponseType {
	return &m
}

const (

	// AAANotifyResponseTypeINVALID captures enum value "__INVALID__"
	AAANotifyResponseTypeINVALID AAANotifyResponseType = "__INVALID__"

	// AAANotifyResponseTypeAAANotifyTypeLogin captures enum value "AAANotifyTypeLogin"
	AAANotifyResponseTypeAAANotifyTypeLogin AAANotifyResponseType = "AAANotifyTypeLogin"

	// AAANotifyResponseTypeAAANotifyTypeLogout captures enum value "AAANotifyTypeLogout"
	AAANotifyResponseTypeAAANotifyTypeLogout AAANotifyResponseType = "AAANotifyTypeLogout"
)

// for schema
var aAANotifyResponseTypeEnum []interface{}

func init() {
	var res []AAANotifyResponseType
	if err := json.Unmarshal([]byte(`["__INVALID__","AAANotifyTypeLogin","AAANotifyTypeLogout"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		aAANotifyResponseTypeEnum = append(aAANotifyResponseTypeEnum, v)
	}
}

func (m AAANotifyResponseType) validateAAANotifyResponseTypeEnum(path, location string, value AAANotifyResponseType) error {
	if err := validate.EnumCase(path, location, value, aAANotifyResponseTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this a a a notify response type
func (m AAANotifyResponseType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAAANotifyResponseTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this a a a notify response type based on context it is used
func (m AAANotifyResponseType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

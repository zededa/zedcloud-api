// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package swagger_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// JWTInfo j w t info
//
// swagger:model JWTInfo
type JWTInfo struct {

	// allow sec
	AllowSec int64 `json:"allowSec,omitempty"`

	// disp Url
	DispURL string `json:"dispUrl,omitempty"`

	// encrypt
	Encrypt bool `json:"encrypt,omitempty"`

	// expire sec
	ExpireSec string `json:"expireSec,omitempty"`

	// num inst
	NumInst int64 `json:"numInst,omitempty"`
}

// Validate validates this j w t info
func (m *JWTInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this j w t info based on context it is used
func (m *JWTInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *JWTInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JWTInfo) UnmarshalBinary(b []byte) error {
	var res JWTInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

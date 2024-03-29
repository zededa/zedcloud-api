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

// TwinDetail twin detail
//
// swagger:model TwinDetail
type TwinDetail struct {

	// authentication type
	AuthenticationType string `json:"authenticationType,omitempty"`

	// cloud to device message count
	CloudToDeviceMessageCount int64 `json:"cloudToDeviceMessageCount,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// desired
	// Format: byte
	Desired strfmt.Base64 `json:"desired,omitempty"`

	// last desired status
	// Format: byte
	LastDesiredStatus strfmt.Base64 `json:"lastDesiredStatus,omitempty"`

	// module count
	ModuleCount int64 `json:"moduleCount,omitempty"`

	// reported
	// Format: byte
	Reported strfmt.Base64 `json:"reported,omitempty"`

	// status code
	StatusCode int32 `json:"statusCode,omitempty"`

	// tags
	// Format: byte
	Tags strfmt.Base64 `json:"tags,omitempty"`
}

// Validate validates this twin detail
func (m *TwinDetail) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this twin detail based on context it is used
func (m *TwinDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TwinDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TwinDetail) UnmarshalBinary(b []byte) error {
	var res TwinDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

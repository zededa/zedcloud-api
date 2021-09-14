// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package swagger_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConfigSignatureInfo XXX this will be deprecated when all deployed instances of EVE
// no longer expect it. 5.6.X depend on it. 5.7.1 does not.
//
// swagger:model configSignatureInfo
type ConfigSignatureInfo struct {

	// intercertsurl
	Intercertsurl string `json:"intercertsurl,omitempty"`

	// signature
	// Format: byte
	Signature strfmt.Base64 `json:"signature,omitempty"`

	// signercerturl
	Signercerturl string `json:"signercerturl,omitempty"`
}

// Validate validates this config signature info
func (m *ConfigSignatureInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this config signature info based on context it is used
func (m *ConfigSignatureInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConfigSignatureInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigSignatureInfo) UnmarshalBinary(b []byte) error {
	var res ConfigSignatureInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

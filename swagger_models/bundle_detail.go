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

// BundleDetail bundle detail
//
// swagger:model BundleDetail
type BundleDetail struct {

	// name
	Name string `json:"name,omitempty"`

	// parent bundle Id
	ParentBundleID string `json:"parentBundleId,omitempty"`
}

// Validate validates this bundle detail
func (m *BundleDetail) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this bundle detail based on context it is used
func (m *BundleDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BundleDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BundleDetail) UnmarshalBinary(b []byte) error {
	var res BundleDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

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

// Cursor Cursor is used as filter in list operation.
//
// Cursor helps in filtering the various list response like edge-app bundle list, model list, bundle list etc.
//
// swagger:model Cursor
type Cursor struct {

	// OrderBy helps in sorting the list response
	OrderBy []string `json:"orderBy"`

	// Page Number
	PageNum int64 `json:"pageNum,omitempty"`

	// Defines the page size
	PageSize int64 `json:"pageSize,omitempty"`

	// Page Token
	PageToken string `json:"pageToken,omitempty"`

	// Total number of pages to be fetched.
	TotalPages int64 `json:"totalPages,omitempty"`
}

// Validate validates this cursor
func (m *Cursor) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cursor based on context it is used
func (m *Cursor) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Cursor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Cursor) UnmarshalBinary(b []byte) error {
	var res Cursor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

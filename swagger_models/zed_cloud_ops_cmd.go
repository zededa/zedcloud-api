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

// ZedCloudOpsCmd Zed cloud Operation command
//
// # Zed cloud Operation command
//
// swagger:model ZedCloudOpsCmd
type ZedCloudOpsCmd struct {

	// counter
	Counter int64 `json:"counter,omitempty"`

	// Timestamp: Operational time
	OpsTime interface{} `json:"opsTime,omitempty"`
}

// Validate validates this zed cloud ops cmd
func (m *ZedCloudOpsCmd) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this zed cloud ops cmd based on context it is used
func (m *ZedCloudOpsCmd) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ZedCloudOpsCmd) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ZedCloudOpsCmd) UnmarshalBinary(b []byte) error {
	var res ZedCloudOpsCmd
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

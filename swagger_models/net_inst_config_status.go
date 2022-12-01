// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package swagger_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NetInstConfigStatus Network instance config status
//
// NetInstConfigStatus stores the summary of the network instance config status.
//
// swagger:model NetInstConfigStatus
type NetInstConfigStatus struct {

	// User defined name of the clusterInstance, unique across the enterprise.
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9][a-zA-Z0-9_.-]+
	ClusterID string `json:"clusterID,omitempty"`

	// device default
	DeviceDefault bool `json:"deviceDefault,omitempty"`

	// device Id
	DeviceID string `json:"deviceId,omitempty"`

	// device name
	DeviceName string `json:"deviceName,omitempty"`

	// System defined universally unique Id of the network instance
	// Read Only: true
	// Pattern: [0-9A-Za-z-]+
	ID string `json:"id,omitempty"`

	// kind
	Kind *NetworkInstanceKind `json:"kind,omitempty"`

	// User defined name of the network instance, unique across the enterprise. Once object is created, name can’t be changed
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9][a-zA-Z0-9_.-]+
	Name string `json:"name,omitempty"`

	// network policy Id
	NetworkPolicyID string `json:"networkPolicyId,omitempty"`

	// project Id
	ProjectID string `json:"projectId,omitempty"`

	// project name
	ProjectName string `json:"projectName,omitempty"`

	// run state
	RunState *RunState `json:"runState,omitempty"`

	// Tags are name/value pairs that enable you to categorize resources. Tag names are case insensitive with max_length 512 and min_length 3. Tag values are case sensitive with max_length 256 and min_length 3.
	Tags map[string]string `json:"tags,omitempty"`

	// type
	Type *NetworkInstanceDhcpType `json:"type,omitempty"`

	// up time stamp
	// Format: date-time
	UpTimeStamp strfmt.DateTime `json:"upTimeStamp,omitempty"`

	// uplink intf
	UplinkIntf string `json:"uplinkIntf,omitempty"`
}

// Validate validates this net inst config status
func (m *NetInstConfigStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpTimeStamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetInstConfigStatus) validateClusterID(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterID) { // not required
		return nil
	}

	if err := validate.MinLength("clusterID", "body", m.ClusterID, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("clusterID", "body", m.ClusterID, 256); err != nil {
		return err
	}

	if err := validate.Pattern("clusterID", "body", m.ClusterID, `[a-zA-Z0-9][a-zA-Z0-9_.-]+`); err != nil {
		return err
	}

	return nil
}

func (m *NetInstConfigStatus) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.Pattern("id", "body", m.ID, `[0-9A-Za-z-]+`); err != nil {
		return err
	}

	return nil
}

func (m *NetInstConfigStatus) validateKind(formats strfmt.Registry) error {
	if swag.IsZero(m.Kind) { // not required
		return nil
	}

	if m.Kind != nil {
		if err := m.Kind.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kind")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kind")
			}
			return err
		}
	}

	return nil
}

func (m *NetInstConfigStatus) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", m.Name, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", m.Name, 256); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", m.Name, `[a-zA-Z0-9][a-zA-Z0-9_.-]+`); err != nil {
		return err
	}

	return nil
}

func (m *NetInstConfigStatus) validateRunState(formats strfmt.Registry) error {
	if swag.IsZero(m.RunState) { // not required
		return nil
	}

	if m.RunState != nil {
		if err := m.RunState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("runState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("runState")
			}
			return err
		}
	}

	return nil
}

func (m *NetInstConfigStatus) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *NetInstConfigStatus) validateUpTimeStamp(formats strfmt.Registry) error {
	if swag.IsZero(m.UpTimeStamp) { // not required
		return nil
	}

	if err := validate.FormatOf("upTimeStamp", "body", "date-time", m.UpTimeStamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this net inst config status based on the context it is used
func (m *NetInstConfigStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKind(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRunState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetInstConfigStatus) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *NetInstConfigStatus) contextValidateKind(ctx context.Context, formats strfmt.Registry) error {

	if m.Kind != nil {
		if err := m.Kind.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kind")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kind")
			}
			return err
		}
	}

	return nil
}

func (m *NetInstConfigStatus) contextValidateRunState(ctx context.Context, formats strfmt.Registry) error {

	if m.RunState != nil {
		if err := m.RunState.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("runState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("runState")
			}
			return err
		}
	}

	return nil
}

func (m *NetInstConfigStatus) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetInstConfigStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetInstConfigStatus) UnmarshalBinary(b []byte) error {
	var res NetInstConfigStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

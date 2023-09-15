// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package swagger_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AppInstStatusSummaryMsg App instance status summary
//
// # AppInstStatusSummaryMsg stores the summary of the app instance status
//
// swagger:model AppInstStatusSummaryMsg
type AppInstStatusSummaryMsg struct {

	// cpu summary
	CPU *CPUSummary `json:"Cpu,omitempty"`

	// memory summary
	Memory *MemorySummary `json:"Memory,omitempty"`

	// storage summary
	Storage *StorageSummary `json:"Storage,omitempty"`

	// SPU details
	AdminState *AdminState `json:"adminState,omitempty"`

	// User defined name of the edge app, unique across the enterprise. Once edge app is created, name can’t be changed
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9][a-zA-Z0-9_.-]+
	AppID string `json:"appId,omitempty"`

	// User defined name of the edge app, unique across the enterprise. Once edge app is created, name can’t be changed
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9][a-zA-Z0-9_.-]+
	AppName string `json:"appName,omitempty"`

	// type of app
	AppType *AppType `json:"appType,omitempty"`

	// System defined universally unique clusterInstance ID, unique across the enterprise.
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9][a-zA-Z0-9_.-]+
	ClusterID string `json:"clusterID,omitempty"`

	// type of deployment for the app, eg: azure, k3s, standalone
	DeploymentType *DeploymentType `json:"deploymentType,omitempty"`

	// User defined name of the device, unique across the enterprise. Once device is created, name can’t be changed
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9][a-zA-Z0-9_.-]+
	DeviceID string `json:"deviceId,omitempty"`

	// User defined name of the device, unique across the enterprise. Once device is created, name can’t be changed
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9][a-zA-Z0-9_.-]+
	DeviceName string `json:"deviceName,omitempty"`

	// System defined universally unique Id of the app instance
	// Read Only: true
	// Pattern: [0-9A-Za-z-]+
	ID string `json:"id,omitempty"`

	// memory summary
	MemorySummary *AppInstMemorySummary `json:"memorySummary,omitempty"`

	// User defined name of the app instance, unique across the enterprise. Once app instance is created, name can’t be changed
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9][a-zA-Z0-9_.-]+
	Name string `json:"name,omitempty"`

	// User defined name of the project, unique across the enterprise. Once project is created, name can’t be changed
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9][a-zA-Z0-9_.-]+
	ProjectID string `json:"projectId,omitempty"`

	// User defined name of the project, unique across the enterprise. Once project is created, name can’t be changed
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9][a-zA-Z0-9_.-]+
	ProjectName string `json:"projectName,omitempty"`

	// operation status
	RunState *RunState `json:"runState,omitempty"`

	// Software details
	SwInfo []*SWInfo `json:"swInfo"`

	// sotware state
	SwState *SWState `json:"swState,omitempty"`

	// Tags are name/value pairs that enable you to categorize resources. Tag names are case insensitive with max_length 512 and min_length 3. Tag values are case sensitive with max_length 256 and min_length 3.
	Tags map[string]string `json:"tags,omitempty"`

	// User defined title of the app instance. Title can be changed at any time
	// Max Length: 256
	// Min Length: 3
	// Pattern: ^[a-zA-Z0-9]+[a-zA-Z0-9!-~ ]+$
	Title string `json:"title,omitempty"`
}

// Validate validates this app inst status summary msg
func (m *AppInstStatusSummaryMsg) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCPU(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAdminState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAppID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAppName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAppType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeploymentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemorySummary(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSwInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSwState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppInstStatusSummaryMsg) validateCPU(formats strfmt.Registry) error {
	if swag.IsZero(m.CPU) { // not required
		return nil
	}

	if m.CPU != nil {
		if err := m.CPU.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Cpu")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Cpu")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstStatusSummaryMsg) validateMemory(formats strfmt.Registry) error {
	if swag.IsZero(m.Memory) { // not required
		return nil
	}

	if m.Memory != nil {
		if err := m.Memory.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Memory")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Memory")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstStatusSummaryMsg) validateStorage(formats strfmt.Registry) error {
	if swag.IsZero(m.Storage) { // not required
		return nil
	}

	if m.Storage != nil {
		if err := m.Storage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Storage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Storage")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstStatusSummaryMsg) validateAdminState(formats strfmt.Registry) error {
	if swag.IsZero(m.AdminState) { // not required
		return nil
	}

	if m.AdminState != nil {
		if err := m.AdminState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("adminState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("adminState")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstStatusSummaryMsg) validateAppID(formats strfmt.Registry) error {
	if swag.IsZero(m.AppID) { // not required
		return nil
	}

	if err := validate.MinLength("appId", "body", m.AppID, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("appId", "body", m.AppID, 256); err != nil {
		return err
	}

	if err := validate.Pattern("appId", "body", m.AppID, `[a-zA-Z0-9][a-zA-Z0-9_.-]+`); err != nil {
		return err
	}

	return nil
}

func (m *AppInstStatusSummaryMsg) validateAppName(formats strfmt.Registry) error {
	if swag.IsZero(m.AppName) { // not required
		return nil
	}

	if err := validate.MinLength("appName", "body", m.AppName, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("appName", "body", m.AppName, 256); err != nil {
		return err
	}

	if err := validate.Pattern("appName", "body", m.AppName, `[a-zA-Z0-9][a-zA-Z0-9_.-]+`); err != nil {
		return err
	}

	return nil
}

func (m *AppInstStatusSummaryMsg) validateAppType(formats strfmt.Registry) error {
	if swag.IsZero(m.AppType) { // not required
		return nil
	}

	if m.AppType != nil {
		if err := m.AppType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("appType")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstStatusSummaryMsg) validateClusterID(formats strfmt.Registry) error {
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

func (m *AppInstStatusSummaryMsg) validateDeploymentType(formats strfmt.Registry) error {
	if swag.IsZero(m.DeploymentType) { // not required
		return nil
	}

	if m.DeploymentType != nil {
		if err := m.DeploymentType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deploymentType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deploymentType")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstStatusSummaryMsg) validateDeviceID(formats strfmt.Registry) error {
	if swag.IsZero(m.DeviceID) { // not required
		return nil
	}

	if err := validate.MinLength("deviceId", "body", m.DeviceID, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("deviceId", "body", m.DeviceID, 256); err != nil {
		return err
	}

	if err := validate.Pattern("deviceId", "body", m.DeviceID, `[a-zA-Z0-9][a-zA-Z0-9_.-]+`); err != nil {
		return err
	}

	return nil
}

func (m *AppInstStatusSummaryMsg) validateDeviceName(formats strfmt.Registry) error {
	if swag.IsZero(m.DeviceName) { // not required
		return nil
	}

	if err := validate.MinLength("deviceName", "body", m.DeviceName, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("deviceName", "body", m.DeviceName, 256); err != nil {
		return err
	}

	if err := validate.Pattern("deviceName", "body", m.DeviceName, `[a-zA-Z0-9][a-zA-Z0-9_.-]+`); err != nil {
		return err
	}

	return nil
}

func (m *AppInstStatusSummaryMsg) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.Pattern("id", "body", m.ID, `[0-9A-Za-z-]+`); err != nil {
		return err
	}

	return nil
}

func (m *AppInstStatusSummaryMsg) validateMemorySummary(formats strfmt.Registry) error {
	if swag.IsZero(m.MemorySummary) { // not required
		return nil
	}

	if m.MemorySummary != nil {
		if err := m.MemorySummary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("memorySummary")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("memorySummary")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstStatusSummaryMsg) validateName(formats strfmt.Registry) error {
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

func (m *AppInstStatusSummaryMsg) validateProjectID(formats strfmt.Registry) error {
	if swag.IsZero(m.ProjectID) { // not required
		return nil
	}

	if err := validate.MinLength("projectId", "body", m.ProjectID, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("projectId", "body", m.ProjectID, 256); err != nil {
		return err
	}

	if err := validate.Pattern("projectId", "body", m.ProjectID, `[a-zA-Z0-9][a-zA-Z0-9_.-]+`); err != nil {
		return err
	}

	return nil
}

func (m *AppInstStatusSummaryMsg) validateProjectName(formats strfmt.Registry) error {
	if swag.IsZero(m.ProjectName) { // not required
		return nil
	}

	if err := validate.MinLength("projectName", "body", m.ProjectName, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("projectName", "body", m.ProjectName, 256); err != nil {
		return err
	}

	if err := validate.Pattern("projectName", "body", m.ProjectName, `[a-zA-Z0-9][a-zA-Z0-9_.-]+`); err != nil {
		return err
	}

	return nil
}

func (m *AppInstStatusSummaryMsg) validateRunState(formats strfmt.Registry) error {
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

func (m *AppInstStatusSummaryMsg) validateSwInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.SwInfo) { // not required
		return nil
	}

	for i := 0; i < len(m.SwInfo); i++ {
		if swag.IsZero(m.SwInfo[i]) { // not required
			continue
		}

		if m.SwInfo[i] != nil {
			if err := m.SwInfo[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("swInfo" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("swInfo" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AppInstStatusSummaryMsg) validateSwState(formats strfmt.Registry) error {
	if swag.IsZero(m.SwState) { // not required
		return nil
	}

	if m.SwState != nil {
		if err := m.SwState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("swState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("swState")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstStatusSummaryMsg) validateTitle(formats strfmt.Registry) error {
	if swag.IsZero(m.Title) { // not required
		return nil
	}

	if err := validate.MinLength("title", "body", m.Title, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("title", "body", m.Title, 256); err != nil {
		return err
	}

	if err := validate.Pattern("title", "body", m.Title, `^[a-zA-Z0-9]+[a-zA-Z0-9!-~ ]+$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this app inst status summary msg based on the context it is used
func (m *AppInstStatusSummaryMsg) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCPU(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMemory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStorage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAdminState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAppType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeploymentType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMemorySummary(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRunState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSwInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSwState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppInstStatusSummaryMsg) contextValidateCPU(ctx context.Context, formats strfmt.Registry) error {

	if m.CPU != nil {

		if swag.IsZero(m.CPU) { // not required
			return nil
		}

		if err := m.CPU.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Cpu")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Cpu")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstStatusSummaryMsg) contextValidateMemory(ctx context.Context, formats strfmt.Registry) error {

	if m.Memory != nil {

		if swag.IsZero(m.Memory) { // not required
			return nil
		}

		if err := m.Memory.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Memory")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Memory")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstStatusSummaryMsg) contextValidateStorage(ctx context.Context, formats strfmt.Registry) error {

	if m.Storage != nil {

		if swag.IsZero(m.Storage) { // not required
			return nil
		}

		if err := m.Storage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Storage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Storage")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstStatusSummaryMsg) contextValidateAdminState(ctx context.Context, formats strfmt.Registry) error {

	if m.AdminState != nil {

		if swag.IsZero(m.AdminState) { // not required
			return nil
		}

		if err := m.AdminState.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("adminState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("adminState")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstStatusSummaryMsg) contextValidateAppType(ctx context.Context, formats strfmt.Registry) error {

	if m.AppType != nil {

		if swag.IsZero(m.AppType) { // not required
			return nil
		}

		if err := m.AppType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("appType")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstStatusSummaryMsg) contextValidateDeploymentType(ctx context.Context, formats strfmt.Registry) error {

	if m.DeploymentType != nil {

		if swag.IsZero(m.DeploymentType) { // not required
			return nil
		}

		if err := m.DeploymentType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deploymentType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deploymentType")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstStatusSummaryMsg) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *AppInstStatusSummaryMsg) contextValidateMemorySummary(ctx context.Context, formats strfmt.Registry) error {

	if m.MemorySummary != nil {

		if swag.IsZero(m.MemorySummary) { // not required
			return nil
		}

		if err := m.MemorySummary.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("memorySummary")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("memorySummary")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstStatusSummaryMsg) contextValidateRunState(ctx context.Context, formats strfmt.Registry) error {

	if m.RunState != nil {

		if swag.IsZero(m.RunState) { // not required
			return nil
		}

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

func (m *AppInstStatusSummaryMsg) contextValidateSwInfo(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SwInfo); i++ {

		if m.SwInfo[i] != nil {

			if swag.IsZero(m.SwInfo[i]) { // not required
				return nil
			}

			if err := m.SwInfo[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("swInfo" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("swInfo" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AppInstStatusSummaryMsg) contextValidateSwState(ctx context.Context, formats strfmt.Registry) error {

	if m.SwState != nil {

		if swag.IsZero(m.SwState) { // not required
			return nil
		}

		if err := m.SwState.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("swState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("swState")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppInstStatusSummaryMsg) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppInstStatusSummaryMsg) UnmarshalBinary(b []byte) error {
	var res AppInstStatusSummaryMsg
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

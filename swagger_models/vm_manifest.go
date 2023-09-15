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

// VMManifest VM manifest
//
// swagger:model VMManifest
type VMManifest struct {

	// Edge Application type
	//
	// UI map: N/A - not exposed to users
	AcKind *string `json:"acKind,omitempty"`

	// Manifest version
	//
	// UI map: N/A - not exposed to users
	AcVersion *string `json:"acVersion,omitempty"`

	// bundle type, eg: vm, container, module
	AppType *AppType `json:"appType,omitempty"`

	// Annotations   types.Annotations    `json:"Annotations,omitempty"`
	// Labels        types.Labels         `json:"Labels,omitempty"`
	//
	// Template for Custom Configuration. Used for Cloud-Init
	Configuration *UserDataTemplate `json:"configuration,omitempty"`

	// Container specific details
	//
	// Create options direct the creation of the Docker container
	ContainerDetail *ContainerDetail `json:"containerDetail,omitempty"`

	// Enable CpuPinning
	CPUPinningEnabled bool `json:"cpuPinningEnabled,omitempty"`

	// type of deployment for the app, eg: azure, k3s, standalone
	DeploymentType *DeploymentType `json:"deploymentType,omitempty"`

	// Description of the application
	Desc *Details `json:"desc,omitempty"`

	// Description of the Edge application
	//
	// UI map: AppDetailsPage:IdentityPane:DescriptionField, AppMarketplacePage:AppCard:DescriptionField
	// Pattern: [0-9A-Za-z-]+
	Description string `json:"description,omitempty"`

	// Display name or title of app manifest
	//
	// UI map: AppEditPage:IdentityPane:Title_Field, AppDetailsPage:IdentityPane:Title_Field
	DisplayName string `json:"displayName,omitempty"`

	// Enable VNC for the app
	//
	// UI map: AppEditPage:IdentityPane:VNC_Field, AppDetailsPage:IdentityPane:VNC_Field
	Enablevnc bool `json:"enablevnc,omitempty"`

	// Images used by the app
	//
	// UI map: AppEditPage:DrivesPane, AppDetailsPage:DrivesPane
	Images []*VMManifestImage `json:"images"`

	// I/O adapter settings
	//
	// UI map: AppEditPage:EnvironmentsPane, AppDetailsPage:EnvironmentsPane
	Interfaces []*Interface `json:"interfaces"`

	// Module specific details
	//
	// Azure module specific details like module twin, environment variable, routes
	Module *ModuleDetail `json:"module,omitempty"`

	// Unique id of app manifest, should match object name
	//
	// UI map: AppEditPage:IdentityPane:Name_Field, AppDetailsPage:IdentityPane:Name_Field
	Name string `json:"name,omitempty"`

	// Owner of the application
	Owner *Author `json:"owner,omitempty"`

	// permissions
	Permissions []Permission `json:"permissions"`

	// Hardware resource requirement (CPU, Memory, Storage) for the app
	//
	// UI map: AppEditPage:ResourcesPane, AppDetailsPage:ResourcesPane
	Resources []*Resource `json:"resources"`

	// VM mode for VM-based app
	//
	// UI map: AppEditPage:IdentityPane:VM_Mode_Field, AppDetailsPage:IdentityPane:VM_Mode_Field
	Vmmode *string `json:"vmmode,omitempty"`
}

// Validate validates this VM manifest
func (m *VMManifest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContainerDetail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeploymentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDesc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterfaces(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMManifest) validateAppType(formats strfmt.Registry) error {
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

func (m *VMManifest) validateConfiguration(formats strfmt.Registry) error {
	if swag.IsZero(m.Configuration) { // not required
		return nil
	}

	if m.Configuration != nil {
		if err := m.Configuration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("configuration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("configuration")
			}
			return err
		}
	}

	return nil
}

func (m *VMManifest) validateContainerDetail(formats strfmt.Registry) error {
	if swag.IsZero(m.ContainerDetail) { // not required
		return nil
	}

	if m.ContainerDetail != nil {
		if err := m.ContainerDetail.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("containerDetail")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("containerDetail")
			}
			return err
		}
	}

	return nil
}

func (m *VMManifest) validateDeploymentType(formats strfmt.Registry) error {
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

func (m *VMManifest) validateDesc(formats strfmt.Registry) error {
	if swag.IsZero(m.Desc) { // not required
		return nil
	}

	if m.Desc != nil {
		if err := m.Desc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("desc")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("desc")
			}
			return err
		}
	}

	return nil
}

func (m *VMManifest) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.Pattern("description", "body", m.Description, `[0-9A-Za-z-]+`); err != nil {
		return err
	}

	return nil
}

func (m *VMManifest) validateImages(formats strfmt.Registry) error {
	if swag.IsZero(m.Images) { // not required
		return nil
	}

	for i := 0; i < len(m.Images); i++ {
		if swag.IsZero(m.Images[i]) { // not required
			continue
		}

		if m.Images[i] != nil {
			if err := m.Images[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("images" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("images" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VMManifest) validateInterfaces(formats strfmt.Registry) error {
	if swag.IsZero(m.Interfaces) { // not required
		return nil
	}

	for i := 0; i < len(m.Interfaces); i++ {
		if swag.IsZero(m.Interfaces[i]) { // not required
			continue
		}

		if m.Interfaces[i] != nil {
			if err := m.Interfaces[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("interfaces" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("interfaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VMManifest) validateModule(formats strfmt.Registry) error {
	if swag.IsZero(m.Module) { // not required
		return nil
	}

	if m.Module != nil {
		if err := m.Module.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("module")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("module")
			}
			return err
		}
	}

	return nil
}

func (m *VMManifest) validateOwner(formats strfmt.Registry) error {
	if swag.IsZero(m.Owner) { // not required
		return nil
	}

	if m.Owner != nil {
		if err := m.Owner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

func (m *VMManifest) validateResources(formats strfmt.Registry) error {
	if swag.IsZero(m.Resources) { // not required
		return nil
	}

	for i := 0; i < len(m.Resources); i++ {
		if swag.IsZero(m.Resources[i]) { // not required
			continue
		}

		if m.Resources[i] != nil {
			if err := m.Resources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("resources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this VM manifest based on the context it is used
func (m *VMManifest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAppType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConfiguration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateContainerDetail(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeploymentType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDesc(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateImages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInterfaces(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOwner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMManifest) contextValidateAppType(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMManifest) contextValidateConfiguration(ctx context.Context, formats strfmt.Registry) error {

	if m.Configuration != nil {

		if swag.IsZero(m.Configuration) { // not required
			return nil
		}

		if err := m.Configuration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("configuration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("configuration")
			}
			return err
		}
	}

	return nil
}

func (m *VMManifest) contextValidateContainerDetail(ctx context.Context, formats strfmt.Registry) error {

	if m.ContainerDetail != nil {

		if swag.IsZero(m.ContainerDetail) { // not required
			return nil
		}

		if err := m.ContainerDetail.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("containerDetail")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("containerDetail")
			}
			return err
		}
	}

	return nil
}

func (m *VMManifest) contextValidateDeploymentType(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMManifest) contextValidateDesc(ctx context.Context, formats strfmt.Registry) error {

	if m.Desc != nil {

		if swag.IsZero(m.Desc) { // not required
			return nil
		}

		if err := m.Desc.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("desc")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("desc")
			}
			return err
		}
	}

	return nil
}

func (m *VMManifest) contextValidateImages(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Images); i++ {

		if m.Images[i] != nil {

			if swag.IsZero(m.Images[i]) { // not required
				return nil
			}

			if err := m.Images[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("images" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("images" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VMManifest) contextValidateInterfaces(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Interfaces); i++ {

		if m.Interfaces[i] != nil {

			if swag.IsZero(m.Interfaces[i]) { // not required
				return nil
			}

			if err := m.Interfaces[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("interfaces" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("interfaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VMManifest) contextValidateModule(ctx context.Context, formats strfmt.Registry) error {

	if m.Module != nil {

		if swag.IsZero(m.Module) { // not required
			return nil
		}

		if err := m.Module.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("module")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("module")
			}
			return err
		}
	}

	return nil
}

func (m *VMManifest) contextValidateOwner(ctx context.Context, formats strfmt.Registry) error {

	if m.Owner != nil {

		if swag.IsZero(m.Owner) { // not required
			return nil
		}

		if err := m.Owner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

func (m *VMManifest) contextValidateResources(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Resources); i++ {

		if m.Resources[i] != nil {

			if swag.IsZero(m.Resources[i]) { // not required
				return nil
			}

			if err := m.Resources[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("resources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VMManifest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMManifest) UnmarshalBinary(b []byte) error {
	var res VMManifest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

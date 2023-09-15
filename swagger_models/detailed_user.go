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

// DetailedUser User detail
//
// User meta data
// Example: {"HubspotId":"","LastLoginTime":"2020-07-17T06:02:36Z","LastLogoutTime":"1970-01-01T00:00:01Z","SfdcId":"","customUserInput":{},"email":"us.root@zededa.com","emailState":"ID_STATE_VERIFIED","firstName":"ROOT","fullName":"ZEDEDA root user","id":"AAF1ABCtMaCnVHZN_b9cm2yUEkgp","locale":"EN","notifyPref":"email","phone":"","phoneState":"ID_STATE_UNSPECIFIED","revision":{"createdAt":"2020-07-16T18:19:56Z","createdBy":"SYSTEM_ROOT","curr":"3","prev":"","updatedAt":"1970-01-01T00:00:01Z","updatedBy":"SYSTEM_ROOT"},"roleName":"SysRoot","state":"USER_STATE_ACTIVE","timeZone":"","type":"AUTH_TYPE_UNSPECIFIED","username":"test@zededa.com"}
//
// swagger:model DetailedUser
type DetailedUser struct {

	// hubspot Id
	HubspotID string `json:"HubspotId,omitempty"`

	// Operational Status to be returned to CLI/UI
	//
	// Last login time of the user
	// Format: date-time
	LastLoginTime strfmt.DateTime `json:"LastLoginTime,omitempty"`

	// Last logout time of the user
	// Format: date-time
	LastLogoutTime strfmt.DateTime `json:"LastLogoutTime,omitempty"`

	// sfdc Id
	SfdcID string `json:"SfdcId,omitempty"`

	// Permitted list of enterprises with their associated roles
	AllowedEnterprises []*AllowedEnterprise `json:"allowedEnterprises"`

	// Custom user parameters
	CustomUserInput map[string]string `json:"customUserInput,omitempty"`

	// Email of the user
	// Required: true
	Email *string `json:"email"`

	// Email state
	// Read Only: true
	EmailState *IDState `json:"emailState,omitempty"`

	// Origin enterprise of the user
	// Read Only: true
	EnterpriseID string `json:"enterpriseId,omitempty"`

	// First name of the user
	FirstName string `json:"firstName,omitempty"`

	// Full name of the user
	FullName string `json:"fullName,omitempty"`

	// Unique system defined user ID
	// Read Only: true
	// Pattern: [0-9A-Za-z_=-]{28}
	ID string `json:"id,omitempty"`

	// Locale of the user
	Locale string `json:"locale,omitempty"`

	// Notification preference of the user
	NotifyPref string `json:"notifyPref,omitempty"`

	// Phone number of the user
	Phone string `json:"phone,omitempty"`

	// Phone state
	// Read Only: true
	PhoneState *IDState `json:"phoneState,omitempty"`

	// system defined info
	// Read Only: true
	Revision *ObjectRevision `json:"revision,omitempty"`

	// Role associated with the user
	// Required: true
	// Pattern: [0-9A-Za-z_=-]{28}
	RoleID *string `json:"roleId"`

	// User state
	// Read Only: true
	State *UserState `json:"state,omitempty"`

	// Preferred time zone of the user
	TimeZone string `json:"timeZone,omitempty"`

	// Type of the user
	Type *AuthType `json:"type,omitempty"`

	// User defined name
	// Required: true
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9][a-zA-Z0-9_.-]+
	Username *string `json:"username"`
}

// Validate validates this detailed user
func (m *DetailedUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastLoginTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastLogoutTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAllowedEnterprises(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmailState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhoneState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRevision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoleID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DetailedUser) validateLastLoginTime(formats strfmt.Registry) error {
	if swag.IsZero(m.LastLoginTime) { // not required
		return nil
	}

	if err := validate.FormatOf("LastLoginTime", "body", "date-time", m.LastLoginTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DetailedUser) validateLastLogoutTime(formats strfmt.Registry) error {
	if swag.IsZero(m.LastLogoutTime) { // not required
		return nil
	}

	if err := validate.FormatOf("LastLogoutTime", "body", "date-time", m.LastLogoutTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DetailedUser) validateAllowedEnterprises(formats strfmt.Registry) error {
	if swag.IsZero(m.AllowedEnterprises) { // not required
		return nil
	}

	for i := 0; i < len(m.AllowedEnterprises); i++ {
		if swag.IsZero(m.AllowedEnterprises[i]) { // not required
			continue
		}

		if m.AllowedEnterprises[i] != nil {
			if err := m.AllowedEnterprises[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("allowedEnterprises" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("allowedEnterprises" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DetailedUser) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	return nil
}

func (m *DetailedUser) validateEmailState(formats strfmt.Registry) error {
	if swag.IsZero(m.EmailState) { // not required
		return nil
	}

	if m.EmailState != nil {
		if err := m.EmailState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("emailState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("emailState")
			}
			return err
		}
	}

	return nil
}

func (m *DetailedUser) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.Pattern("id", "body", m.ID, `[0-9A-Za-z_=-]{28}`); err != nil {
		return err
	}

	return nil
}

func (m *DetailedUser) validatePhoneState(formats strfmt.Registry) error {
	if swag.IsZero(m.PhoneState) { // not required
		return nil
	}

	if m.PhoneState != nil {
		if err := m.PhoneState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("phoneState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("phoneState")
			}
			return err
		}
	}

	return nil
}

func (m *DetailedUser) validateRevision(formats strfmt.Registry) error {
	if swag.IsZero(m.Revision) { // not required
		return nil
	}

	if m.Revision != nil {
		if err := m.Revision.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("revision")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("revision")
			}
			return err
		}
	}

	return nil
}

func (m *DetailedUser) validateRoleID(formats strfmt.Registry) error {

	if err := validate.Required("roleId", "body", m.RoleID); err != nil {
		return err
	}

	if err := validate.Pattern("roleId", "body", *m.RoleID, `[0-9A-Za-z_=-]{28}`); err != nil {
		return err
	}

	return nil
}

func (m *DetailedUser) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	if m.State != nil {
		if err := m.State.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

func (m *DetailedUser) validateType(formats strfmt.Registry) error {
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

func (m *DetailedUser) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	if err := validate.MinLength("username", "body", *m.Username, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("username", "body", *m.Username, 256); err != nil {
		return err
	}

	if err := validate.Pattern("username", "body", *m.Username, `[a-zA-Z0-9][a-zA-Z0-9_.-]+`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this detailed user based on the context it is used
func (m *DetailedUser) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAllowedEnterprises(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEmailState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEnterpriseID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePhoneState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRevision(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
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

func (m *DetailedUser) contextValidateAllowedEnterprises(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AllowedEnterprises); i++ {

		if m.AllowedEnterprises[i] != nil {

			if swag.IsZero(m.AllowedEnterprises[i]) { // not required
				return nil
			}

			if err := m.AllowedEnterprises[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("allowedEnterprises" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("allowedEnterprises" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DetailedUser) contextValidateEmailState(ctx context.Context, formats strfmt.Registry) error {

	if m.EmailState != nil {

		if swag.IsZero(m.EmailState) { // not required
			return nil
		}

		if err := m.EmailState.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("emailState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("emailState")
			}
			return err
		}
	}

	return nil
}

func (m *DetailedUser) contextValidateEnterpriseID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "enterpriseId", "body", string(m.EnterpriseID)); err != nil {
		return err
	}

	return nil
}

func (m *DetailedUser) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *DetailedUser) contextValidatePhoneState(ctx context.Context, formats strfmt.Registry) error {

	if m.PhoneState != nil {

		if swag.IsZero(m.PhoneState) { // not required
			return nil
		}

		if err := m.PhoneState.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("phoneState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("phoneState")
			}
			return err
		}
	}

	return nil
}

func (m *DetailedUser) contextValidateRevision(ctx context.Context, formats strfmt.Registry) error {

	if m.Revision != nil {

		if swag.IsZero(m.Revision) { // not required
			return nil
		}

		if err := m.Revision.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("revision")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("revision")
			}
			return err
		}
	}

	return nil
}

func (m *DetailedUser) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if m.State != nil {

		if swag.IsZero(m.State) { // not required
			return nil
		}

		if err := m.State.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

func (m *DetailedUser) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {

		if swag.IsZero(m.Type) { // not required
			return nil
		}

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
func (m *DetailedUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DetailedUser) UnmarshalBinary(b []byte) error {
	var res DetailedUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

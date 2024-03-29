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

// DatastoreInfo Datastore
//
// Datastore idetail
// Example: {"description":"My test datastore hosted on HTTP server for Edge computing images.","dsError":"Datastore validated successfully...","dsFQDN":"http://my-datastore.my-company.com","dsPath":"download/AMD64","dsStatus":"DATASTORE_STATUS_ACTIVE","dsType":"DATASTORE_TYPE_HTTP","id":"d1125b0f-633d-459c-99c6-f47e7467cebc","name":"my-test-datastore","originType":"ORIGIN_LOCAL","revision":{"createdAt":{"seconds":1592068270},"createdBy":"admin@my-company.com","curr":"1","updatedAt":{"seconds":1592068271},"updatedBy":"admin@my-company.com"},"secret":{"apiKey":"my-datastore-api-key","apiPasswd":"my-datastore-api-password"},"title":"My test datatsore for Edge computing images"}
//
// swagger:model DatastoreInfo
type DatastoreInfo struct {

	// api key
	APIKey string `json:"apiKey,omitempty"`

	// Certificate chain of the certificate
	CertificateChain *CertificateChain `json:"certificateChain,omitempty"`

	// Internal - Encryption Key context
	// Read Only: true
	CryptoKey string `json:"cryptoKey,omitempty"`

	// Detailed description of the datastore.
	// Max Length: 256
	Description string `json:"description,omitempty"`

	// Datastore validation detailed error/status message
	// Read Only: true
	DsErr string `json:"dsErr,omitempty"`

	// Datastore Fully Qualified Domain Name
	// Required: true
	DsFQDN *string `json:"dsFQDN"`

	// Datastore relative path w.r.t. Datastore root
	// Required: true
	DsPath *string `json:"dsPath"`

	// Datastore status
	DsStatus *DatastoreStatus `json:"dsStatus,omitempty"`

	// Datastore type
	// Required: true
	DsType *DatastoreType `json:"dsType"`

	// Internal - Encrypted sensitive data
	// Read Only: true
	EncryptedSecrets map[string]string `json:"encryptedSecrets,omitempty"`

	// enterprise Id
	EnterpriseID string `json:"enterpriseId,omitempty"`

	// System defined universally unique Id of the datastore.
	// Read Only: true
	// Pattern: [a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}
	ID string `json:"id,omitempty"`

	// User defined name of the datastore, unique across the enterprise. Once datastore is created, name can’t be changed.
	// Required: true
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9][a-zA-Z0-9_.-]+
	Name *string `json:"name"`

	// knob for sending creds in clear text
	NeedClearText bool `json:"needClearText,omitempty"`

	// Origin type of datastore.
	// Read Only: true
	OriginType *Origin `json:"originType,omitempty"`

	// project access list of the datastore
	ProjectAccessList []string `json:"projectAccessList"`

	// Datastore region - valid for AWS S3 and Azure BlobStorage
	Region string `json:"region,omitempty"`

	// system defined info
	// Read Only: true
	Revision *ObjectRevision `json:"revision,omitempty"`

	// Plain-text sensitive data
	Secret *DatastoreInfoSecrets `json:"secret,omitempty"`

	// User defined title of the datastore. Title can be changed at any time.
	// Required: true
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9]+[a-zA-Z0-9!-~ ]+
	Title *string `json:"title"`
}

// Validate validates this datastore info
func (m *DatastoreInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCertificateChain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDsFQDN(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDsPath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDsStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDsType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRevision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecret(formats); err != nil {
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

func (m *DatastoreInfo) validateCertificateChain(formats strfmt.Registry) error {
	if swag.IsZero(m.CertificateChain) { // not required
		return nil
	}

	if m.CertificateChain != nil {
		if err := m.CertificateChain.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("certificateChain")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("certificateChain")
			}
			return err
		}
	}

	return nil
}

func (m *DatastoreInfo) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 256); err != nil {
		return err
	}

	return nil
}

func (m *DatastoreInfo) validateDsFQDN(formats strfmt.Registry) error {

	if err := validate.Required("dsFQDN", "body", m.DsFQDN); err != nil {
		return err
	}

	return nil
}

func (m *DatastoreInfo) validateDsPath(formats strfmt.Registry) error {

	if err := validate.Required("dsPath", "body", m.DsPath); err != nil {
		return err
	}

	return nil
}

func (m *DatastoreInfo) validateDsStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.DsStatus) { // not required
		return nil
	}

	if m.DsStatus != nil {
		if err := m.DsStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dsStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dsStatus")
			}
			return err
		}
	}

	return nil
}

func (m *DatastoreInfo) validateDsType(formats strfmt.Registry) error {

	if err := validate.Required("dsType", "body", m.DsType); err != nil {
		return err
	}

	if err := validate.Required("dsType", "body", m.DsType); err != nil {
		return err
	}

	if m.DsType != nil {
		if err := m.DsType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dsType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dsType")
			}
			return err
		}
	}

	return nil
}

func (m *DatastoreInfo) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.Pattern("id", "body", m.ID, `[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}`); err != nil {
		return err
	}

	return nil
}

func (m *DatastoreInfo) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 256); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", *m.Name, `[a-zA-Z0-9][a-zA-Z0-9_.-]+`); err != nil {
		return err
	}

	return nil
}

func (m *DatastoreInfo) validateOriginType(formats strfmt.Registry) error {
	if swag.IsZero(m.OriginType) { // not required
		return nil
	}

	if m.OriginType != nil {
		if err := m.OriginType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("originType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("originType")
			}
			return err
		}
	}

	return nil
}

func (m *DatastoreInfo) validateRevision(formats strfmt.Registry) error {
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

func (m *DatastoreInfo) validateSecret(formats strfmt.Registry) error {
	if swag.IsZero(m.Secret) { // not required
		return nil
	}

	if m.Secret != nil {
		if err := m.Secret.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secret")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("secret")
			}
			return err
		}
	}

	return nil
}

func (m *DatastoreInfo) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	if err := validate.MinLength("title", "body", *m.Title, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("title", "body", *m.Title, 256); err != nil {
		return err
	}

	if err := validate.Pattern("title", "body", *m.Title, `[a-zA-Z0-9]+[a-zA-Z0-9!-~ ]+`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this datastore info based on the context it is used
func (m *DatastoreInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCertificateChain(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCryptoKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDsErr(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDsStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDsType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEncryptedSecrets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOriginType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRevision(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecret(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatastoreInfo) contextValidateCertificateChain(ctx context.Context, formats strfmt.Registry) error {

	if m.CertificateChain != nil {

		if swag.IsZero(m.CertificateChain) { // not required
			return nil
		}

		if err := m.CertificateChain.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("certificateChain")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("certificateChain")
			}
			return err
		}
	}

	return nil
}

func (m *DatastoreInfo) contextValidateCryptoKey(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "cryptoKey", "body", string(m.CryptoKey)); err != nil {
		return err
	}

	return nil
}

func (m *DatastoreInfo) contextValidateDsErr(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dsErr", "body", string(m.DsErr)); err != nil {
		return err
	}

	return nil
}

func (m *DatastoreInfo) contextValidateDsStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.DsStatus != nil {

		if swag.IsZero(m.DsStatus) { // not required
			return nil
		}

		if err := m.DsStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dsStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dsStatus")
			}
			return err
		}
	}

	return nil
}

func (m *DatastoreInfo) contextValidateDsType(ctx context.Context, formats strfmt.Registry) error {

	if m.DsType != nil {

		if err := m.DsType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dsType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dsType")
			}
			return err
		}
	}

	return nil
}

func (m *DatastoreInfo) contextValidateEncryptedSecrets(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *DatastoreInfo) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *DatastoreInfo) contextValidateOriginType(ctx context.Context, formats strfmt.Registry) error {

	if m.OriginType != nil {

		if swag.IsZero(m.OriginType) { // not required
			return nil
		}

		if err := m.OriginType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("originType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("originType")
			}
			return err
		}
	}

	return nil
}

func (m *DatastoreInfo) contextValidateRevision(ctx context.Context, formats strfmt.Registry) error {

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

func (m *DatastoreInfo) contextValidateSecret(ctx context.Context, formats strfmt.Registry) error {

	if m.Secret != nil {

		if swag.IsZero(m.Secret) { // not required
			return nil
		}

		if err := m.Secret.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secret")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("secret")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DatastoreInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatastoreInfo) UnmarshalBinary(b []byte) error {
	var res DatastoreInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

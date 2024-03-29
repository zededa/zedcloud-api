// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package swagger_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// DatastoreType Datastore type
//
// - DATASTORE_TYPE_HTTP: Datastore hosted on HTTP server
//   - DATASTORE_TYPE_HTTPS: Datastore hosted on HTTPS server
//   - DATASTORE_TYPE_AWSS3: Datastore hosted on AWS S3
//   - DATASTORE_TYPE_SFTP: Datastore hosted on SFTP server
//   - DATASTORE_TYPE_CONTAINERREGISTRY: Datastore hosted on Container Registry (e.g. Dockerhub, Azure Container Registry etc.)
//   - DATASTORE_TYPE_AZUREBLOB: Datastore hosted on Azure BlobStorage
//   - DATASTORE_TYPE_ZEDEDAS3: ZEDEDA Owned and Operated datastore hosted on AWS S3
//   - DATASTORE_TYPE_ZEDEDABLOB: ZEDEDA Owned and Operated datastore hosted on Azure BlobStorage
//   - DATASTORE_TYPE_FILE_STORAGE: Datastore hosted on File Storage (e.g. HTTP, HTTPS, AWS S3, SFTP, Azure Blob, etc.)
//
// swagger:model DatastoreType
type DatastoreType string

func NewDatastoreType(value DatastoreType) *DatastoreType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated DatastoreType.
func (m DatastoreType) Pointer() *DatastoreType {
	return &m
}

const (

	// DatastoreTypeDATASTORETYPEUNSPECIFIED captures enum value "DATASTORE_TYPE_UNSPECIFIED"
	DatastoreTypeDATASTORETYPEUNSPECIFIED DatastoreType = "DATASTORE_TYPE_UNSPECIFIED"

	// DatastoreTypeDATASTORETYPEHTTP captures enum value "DATASTORE_TYPE_HTTP"
	DatastoreTypeDATASTORETYPEHTTP DatastoreType = "DATASTORE_TYPE_HTTP"

	// DatastoreTypeDATASTORETYPEHTTPS captures enum value "DATASTORE_TYPE_HTTPS"
	DatastoreTypeDATASTORETYPEHTTPS DatastoreType = "DATASTORE_TYPE_HTTPS"

	// DatastoreTypeDATASTORETYPEAWSS3 captures enum value "DATASTORE_TYPE_AWSS3"
	DatastoreTypeDATASTORETYPEAWSS3 DatastoreType = "DATASTORE_TYPE_AWSS3"

	// DatastoreTypeDATASTORETYPESFTP captures enum value "DATASTORE_TYPE_SFTP"
	DatastoreTypeDATASTORETYPESFTP DatastoreType = "DATASTORE_TYPE_SFTP"

	// DatastoreTypeDATASTORETYPECONTAINERREGISTRY captures enum value "DATASTORE_TYPE_CONTAINERREGISTRY"
	DatastoreTypeDATASTORETYPECONTAINERREGISTRY DatastoreType = "DATASTORE_TYPE_CONTAINERREGISTRY"

	// DatastoreTypeDATASTORETYPEAZUREBLOB captures enum value "DATASTORE_TYPE_AZUREBLOB"
	DatastoreTypeDATASTORETYPEAZUREBLOB DatastoreType = "DATASTORE_TYPE_AZUREBLOB"

	// DatastoreTypeDATASTORETYPEZEDEDAS3 captures enum value "DATASTORE_TYPE_ZEDEDAS3"
	DatastoreTypeDATASTORETYPEZEDEDAS3 DatastoreType = "DATASTORE_TYPE_ZEDEDAS3"

	// DatastoreTypeDATASTORETYPEZEDEDABLOB captures enum value "DATASTORE_TYPE_ZEDEDABLOB"
	DatastoreTypeDATASTORETYPEZEDEDABLOB DatastoreType = "DATASTORE_TYPE_ZEDEDABLOB"

	// DatastoreTypeDATASTORETYPEFILESTORAGE captures enum value "DATASTORE_TYPE_FILE_STORAGE"
	DatastoreTypeDATASTORETYPEFILESTORAGE DatastoreType = "DATASTORE_TYPE_FILE_STORAGE"
)

// for schema
var datastoreTypeEnum []interface{}

func init() {
	var res []DatastoreType
	if err := json.Unmarshal([]byte(`["DATASTORE_TYPE_UNSPECIFIED","DATASTORE_TYPE_HTTP","DATASTORE_TYPE_HTTPS","DATASTORE_TYPE_AWSS3","DATASTORE_TYPE_SFTP","DATASTORE_TYPE_CONTAINERREGISTRY","DATASTORE_TYPE_AZUREBLOB","DATASTORE_TYPE_ZEDEDAS3","DATASTORE_TYPE_ZEDEDABLOB","DATASTORE_TYPE_FILE_STORAGE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		datastoreTypeEnum = append(datastoreTypeEnum, v)
	}
}

func (m DatastoreType) validateDatastoreTypeEnum(path, location string, value DatastoreType) error {
	if err := validate.EnumCase(path, location, value, datastoreTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this datastore type
func (m DatastoreType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDatastoreTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this datastore type based on context it is used
func (m DatastoreType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

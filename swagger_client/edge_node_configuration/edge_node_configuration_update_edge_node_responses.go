// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package edge_node_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// EdgeNodeConfigurationUpdateEdgeNodeReader is a Reader for the EdgeNodeConfigurationUpdateEdgeNode structure.
type EdgeNodeConfigurationUpdateEdgeNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeNodeConfigurationUpdateEdgeNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeNodeConfigurationUpdateEdgeNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeNodeConfigurationUpdateEdgeNodeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeNodeConfigurationUpdateEdgeNodeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeNodeConfigurationUpdateEdgeNodeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewEdgeNodeConfigurationUpdateEdgeNodeConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeNodeConfigurationUpdateEdgeNodeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeNodeConfigurationUpdateEdgeNodeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeNodeConfigurationUpdateEdgeNodeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeNodeConfigurationUpdateEdgeNodeOK creates a EdgeNodeConfigurationUpdateEdgeNodeOK with default headers values
func NewEdgeNodeConfigurationUpdateEdgeNodeOK() *EdgeNodeConfigurationUpdateEdgeNodeOK {
	return &EdgeNodeConfigurationUpdateEdgeNodeOK{}
}

/* EdgeNodeConfigurationUpdateEdgeNodeOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeNodeConfigurationUpdateEdgeNodeOK struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeOK) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}][%d] edgeNodeConfigurationUpdateEdgeNodeOK  %+v", 200, o.Payload)
}
func (o *EdgeNodeConfigurationUpdateEdgeNodeOK) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationUpdateEdgeNodeUnauthorized creates a EdgeNodeConfigurationUpdateEdgeNodeUnauthorized with default headers values
func NewEdgeNodeConfigurationUpdateEdgeNodeUnauthorized() *EdgeNodeConfigurationUpdateEdgeNodeUnauthorized {
	return &EdgeNodeConfigurationUpdateEdgeNodeUnauthorized{}
}

/* EdgeNodeConfigurationUpdateEdgeNodeUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeNodeConfigurationUpdateEdgeNodeUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}][%d] edgeNodeConfigurationUpdateEdgeNodeUnauthorized  %+v", 401, o.Payload)
}
func (o *EdgeNodeConfigurationUpdateEdgeNodeUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationUpdateEdgeNodeForbidden creates a EdgeNodeConfigurationUpdateEdgeNodeForbidden with default headers values
func NewEdgeNodeConfigurationUpdateEdgeNodeForbidden() *EdgeNodeConfigurationUpdateEdgeNodeForbidden {
	return &EdgeNodeConfigurationUpdateEdgeNodeForbidden{}
}

/* EdgeNodeConfigurationUpdateEdgeNodeForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeNodeConfigurationUpdateEdgeNodeForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}][%d] edgeNodeConfigurationUpdateEdgeNodeForbidden  %+v", 403, o.Payload)
}
func (o *EdgeNodeConfigurationUpdateEdgeNodeForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationUpdateEdgeNodeNotFound creates a EdgeNodeConfigurationUpdateEdgeNodeNotFound with default headers values
func NewEdgeNodeConfigurationUpdateEdgeNodeNotFound() *EdgeNodeConfigurationUpdateEdgeNodeNotFound {
	return &EdgeNodeConfigurationUpdateEdgeNodeNotFound{}
}

/* EdgeNodeConfigurationUpdateEdgeNodeNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeNodeConfigurationUpdateEdgeNodeNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}][%d] edgeNodeConfigurationUpdateEdgeNodeNotFound  %+v", 404, o.Payload)
}
func (o *EdgeNodeConfigurationUpdateEdgeNodeNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationUpdateEdgeNodeConflict creates a EdgeNodeConfigurationUpdateEdgeNodeConflict with default headers values
func NewEdgeNodeConfigurationUpdateEdgeNodeConflict() *EdgeNodeConfigurationUpdateEdgeNodeConflict {
	return &EdgeNodeConfigurationUpdateEdgeNodeConflict{}
}

/* EdgeNodeConfigurationUpdateEdgeNodeConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this operation will conflict with an already existing edge node record.
*/
type EdgeNodeConfigurationUpdateEdgeNodeConflict struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}][%d] edgeNodeConfigurationUpdateEdgeNodeConflict  %+v", 409, o.Payload)
}
func (o *EdgeNodeConfigurationUpdateEdgeNodeConflict) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationUpdateEdgeNodeInternalServerError creates a EdgeNodeConfigurationUpdateEdgeNodeInternalServerError with default headers values
func NewEdgeNodeConfigurationUpdateEdgeNodeInternalServerError() *EdgeNodeConfigurationUpdateEdgeNodeInternalServerError {
	return &EdgeNodeConfigurationUpdateEdgeNodeInternalServerError{}
}

/* EdgeNodeConfigurationUpdateEdgeNodeInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeNodeConfigurationUpdateEdgeNodeInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}][%d] edgeNodeConfigurationUpdateEdgeNodeInternalServerError  %+v", 500, o.Payload)
}
func (o *EdgeNodeConfigurationUpdateEdgeNodeInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationUpdateEdgeNodeGatewayTimeout creates a EdgeNodeConfigurationUpdateEdgeNodeGatewayTimeout with default headers values
func NewEdgeNodeConfigurationUpdateEdgeNodeGatewayTimeout() *EdgeNodeConfigurationUpdateEdgeNodeGatewayTimeout {
	return &EdgeNodeConfigurationUpdateEdgeNodeGatewayTimeout{}
}

/* EdgeNodeConfigurationUpdateEdgeNodeGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeNodeConfigurationUpdateEdgeNodeGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}][%d] edgeNodeConfigurationUpdateEdgeNodeGatewayTimeout  %+v", 504, o.Payload)
}
func (o *EdgeNodeConfigurationUpdateEdgeNodeGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationUpdateEdgeNodeDefault creates a EdgeNodeConfigurationUpdateEdgeNodeDefault with default headers values
func NewEdgeNodeConfigurationUpdateEdgeNodeDefault(code int) *EdgeNodeConfigurationUpdateEdgeNodeDefault {
	return &EdgeNodeConfigurationUpdateEdgeNodeDefault{
		_statusCode: code,
	}
}

/* EdgeNodeConfigurationUpdateEdgeNodeDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeNodeConfigurationUpdateEdgeNodeDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the edge node configuration update edge node default response
func (o *EdgeNodeConfigurationUpdateEdgeNodeDefault) Code() int {
	return o._statusCode
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}][%d] EdgeNodeConfiguration_UpdateEdgeNode default  %+v", o._statusCode, o.Payload)
}
func (o *EdgeNodeConfigurationUpdateEdgeNodeDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*EdgeNodeConfigurationUpdateEdgeNodeBody Device Configuration payload detail
//
// Device Configuration request paylod holds the device properties
swagger:model EdgeNodeConfigurationUpdateEdgeNodeBody
*/
type EdgeNodeConfigurationUpdateEdgeNodeBody struct {

	// administrative state of device
	AdminState *swagger_models.AdminState `json:"adminState,omitempty"`

	// Device asset ID
	AssetID string `json:"assetId,omitempty"`

	// base images
	BaseImage []*swagger_models.BaseOSImage `json:"baseImage"`

	// device baseos retry counter
	BaseOsRetryCounter int64 `json:"baseOsRetryCounter,omitempty"`

	// device baseos retry time
	BaseOsRetryTime string `json:"baseOsRetryTime,omitempty"`

	// Client IP
	ClientIP string `json:"clientIp,omitempty"`

	// System defined universally unique clusterInstance ID, unique across the enterprise.
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9][a-zA-Z0-9_.-]+
	ClusterID string `json:"clusterID,omitempty"`

	// ED configurations
	ConfigItem []*swagger_models.EDConfigItem `json:"configItem"`

	// CPU (configured values)
	CPU int64 `json:"cpu,omitempty"`

	// debug knob details for the device
	DebugKnob *swagger_models.DebugKnobDetail `json:"debugKnob,omitempty"`

	// default network instance details
	DefaultNetInst *swagger_models.NetInstConfig `json:"defaultNetInst,omitempty"`

	// deprecated field
	Deprecated string `json:"deprecated,omitempty"`

	// user specified description
	Description string `json:"description,omitempty"`

	// User specified geo location
	DevLocation *swagger_models.GeoLocation `json:"devLocation,omitempty"`

	// device Lisp
	Dlisp *swagger_models.DeviceLisp `json:"dlisp,omitempty"`

	// Device identity
	// Format: byte
	Identity strfmt.Base64 `json:"identity,omitempty"`

	// System Interface list
	Interfaces []*swagger_models.SysInterface `json:"interfaces"`

	// Device location: deprecated
	Location string `json:"location,omitempty"`

	// Device memory in MBs
	Memory int64 `json:"memory,omitempty"`

	// device model
	// Required: true
	ModelID *string `json:"modelId"`

	// user specified device name
	// Required: true
	Name *string `json:"name"`

	// Object key
	Obkey string `json:"obkey,omitempty"`

	// Device level certificates used while onboarding
	Onboarding *swagger_models.DeviceCerts `json:"onboarding,omitempty"`

	// prepare poweroff counter
	PreparePowerOffCounter int64 `json:"preparePowerOffCounter,omitempty"`

	// prepare poweroff time
	PreparePowerOffTime string `json:"preparePowerOffTime,omitempty"`

	// project name
	// Required: true
	ProjectID *string `json:"projectId"`

	// devicereset counter
	ResetCounter int64 `json:"resetCounter,omitempty"`

	// device reset time
	ResetTime string `json:"resetTime,omitempty"`

	// Object revision details
	Revision *swagger_models.ObjectRevision `json:"revision,omitempty"`

	// Device serial number
	Serialno string `json:"serialno,omitempty"`

	// Site captured pictures
	SitePictures []string `json:"sitePictures"`

	// Device storage in GBs
	Storage int64 `json:"storage,omitempty"`

	// Tags are name/value pairs that enable you to categorize resources. Tag names are case insensitive with max_length 512 and min_length 3. Tag values are case sensitive with max_length 256 and min_length 3.
	Tags map[string]string `json:"tags,omitempty"`

	// Threads
	Thread int64 `json:"thread,omitempty"`

	// user specified title
	// Required: true
	Title *string `json:"title"`

	// Single use token
	Token string `json:"token,omitempty"`

	// device model arch type
	Utype *swagger_models.ModelArchType `json:"utype,omitempty"`
}

// Validate validates this edge node configuration update edge node body
func (o *EdgeNodeConfigurationUpdateEdgeNodeBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAdminState(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateBaseImage(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateConfigItem(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDebugKnob(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDefaultNetInst(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDevLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDlisp(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateInterfaces(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateModelID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOnboarding(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateProjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRevision(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUtype(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBody) validateAdminState(formats strfmt.Registry) error {
	if swag.IsZero(o.AdminState) { // not required
		return nil
	}

	if o.AdminState != nil {
		if err := o.AdminState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "adminState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "adminState")
			}
			return err
		}
	}

	return nil
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBody) validateBaseImage(formats strfmt.Registry) error {
	if swag.IsZero(o.BaseImage) { // not required
		return nil
	}

	for i := 0; i < len(o.BaseImage); i++ {
		if swag.IsZero(o.BaseImage[i]) { // not required
			continue
		}

		if o.BaseImage[i] != nil {
			if err := o.BaseImage[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "baseImage" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "baseImage" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBody) validateClusterID(formats strfmt.Registry) error {
	if swag.IsZero(o.ClusterID) { // not required
		return nil
	}

	if err := validate.MinLength("body"+"."+"clusterID", "body", o.ClusterID, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("body"+"."+"clusterID", "body", o.ClusterID, 256); err != nil {
		return err
	}

	if err := validate.Pattern("body"+"."+"clusterID", "body", o.ClusterID, `[a-zA-Z0-9][a-zA-Z0-9_.-]+`); err != nil {
		return err
	}

	return nil
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBody) validateConfigItem(formats strfmt.Registry) error {
	if swag.IsZero(o.ConfigItem) { // not required
		return nil
	}

	for i := 0; i < len(o.ConfigItem); i++ {
		if swag.IsZero(o.ConfigItem[i]) { // not required
			continue
		}

		if o.ConfigItem[i] != nil {
			if err := o.ConfigItem[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "configItem" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "configItem" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBody) validateDebugKnob(formats strfmt.Registry) error {
	if swag.IsZero(o.DebugKnob) { // not required
		return nil
	}

	if o.DebugKnob != nil {
		if err := o.DebugKnob.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "debugKnob")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "debugKnob")
			}
			return err
		}
	}

	return nil
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBody) validateDefaultNetInst(formats strfmt.Registry) error {
	if swag.IsZero(o.DefaultNetInst) { // not required
		return nil
	}

	if o.DefaultNetInst != nil {
		if err := o.DefaultNetInst.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "defaultNetInst")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "defaultNetInst")
			}
			return err
		}
	}

	return nil
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBody) validateDevLocation(formats strfmt.Registry) error {
	if swag.IsZero(o.DevLocation) { // not required
		return nil
	}

	if o.DevLocation != nil {
		if err := o.DevLocation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "devLocation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "devLocation")
			}
			return err
		}
	}

	return nil
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBody) validateDlisp(formats strfmt.Registry) error {
	if swag.IsZero(o.Dlisp) { // not required
		return nil
	}

	if o.Dlisp != nil {
		if err := o.Dlisp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "dlisp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "dlisp")
			}
			return err
		}
	}

	return nil
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBody) validateInterfaces(formats strfmt.Registry) error {
	if swag.IsZero(o.Interfaces) { // not required
		return nil
	}

	for i := 0; i < len(o.Interfaces); i++ {
		if swag.IsZero(o.Interfaces[i]) { // not required
			continue
		}

		if o.Interfaces[i] != nil {
			if err := o.Interfaces[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "interfaces" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "interfaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBody) validateModelID(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"modelId", "body", o.ModelID); err != nil {
		return err
	}

	return nil
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBody) validateOnboarding(formats strfmt.Registry) error {
	if swag.IsZero(o.Onboarding) { // not required
		return nil
	}

	if o.Onboarding != nil {
		if err := o.Onboarding.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "onboarding")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "onboarding")
			}
			return err
		}
	}

	return nil
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBody) validateProjectID(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"projectId", "body", o.ProjectID); err != nil {
		return err
	}

	return nil
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBody) validateRevision(formats strfmt.Registry) error {
	if swag.IsZero(o.Revision) { // not required
		return nil
	}

	if o.Revision != nil {
		if err := o.Revision.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "revision")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "revision")
			}
			return err
		}
	}

	return nil
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBody) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"title", "body", o.Title); err != nil {
		return err
	}

	return nil
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBody) validateUtype(formats strfmt.Registry) error {
	if swag.IsZero(o.Utype) { // not required
		return nil
	}

	if o.Utype != nil {
		if err := o.Utype.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "utype")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "utype")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this edge node configuration update edge node body based on the context it is used
func (o *EdgeNodeConfigurationUpdateEdgeNodeBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAdminState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateBaseImage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateConfigItem(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateDebugKnob(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateDefaultNetInst(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateDevLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateDlisp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateInterfaces(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateOnboarding(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateRevision(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateUtype(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBody) contextValidateAdminState(ctx context.Context, formats strfmt.Registry) error {

	if o.AdminState != nil {
		if err := o.AdminState.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "adminState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "adminState")
			}
			return err
		}
	}

	return nil
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBody) contextValidateBaseImage(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.BaseImage); i++ {

		if o.BaseImage[i] != nil {
			if err := o.BaseImage[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "baseImage" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "baseImage" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBody) contextValidateConfigItem(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.ConfigItem); i++ {

		if o.ConfigItem[i] != nil {
			if err := o.ConfigItem[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "configItem" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "configItem" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBody) contextValidateDebugKnob(ctx context.Context, formats strfmt.Registry) error {

	if o.DebugKnob != nil {
		if err := o.DebugKnob.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "debugKnob")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "debugKnob")
			}
			return err
		}
	}

	return nil
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBody) contextValidateDefaultNetInst(ctx context.Context, formats strfmt.Registry) error {

	if o.DefaultNetInst != nil {
		if err := o.DefaultNetInst.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "defaultNetInst")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "defaultNetInst")
			}
			return err
		}
	}

	return nil
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBody) contextValidateDevLocation(ctx context.Context, formats strfmt.Registry) error {

	if o.DevLocation != nil {
		if err := o.DevLocation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "devLocation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "devLocation")
			}
			return err
		}
	}

	return nil
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBody) contextValidateDlisp(ctx context.Context, formats strfmt.Registry) error {

	if o.Dlisp != nil {
		if err := o.Dlisp.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "dlisp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "dlisp")
			}
			return err
		}
	}

	return nil
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBody) contextValidateInterfaces(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Interfaces); i++ {

		if o.Interfaces[i] != nil {
			if err := o.Interfaces[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "interfaces" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "interfaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBody) contextValidateOnboarding(ctx context.Context, formats strfmt.Registry) error {

	if o.Onboarding != nil {
		if err := o.Onboarding.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "onboarding")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "onboarding")
			}
			return err
		}
	}

	return nil
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBody) contextValidateRevision(ctx context.Context, formats strfmt.Registry) error {

	if o.Revision != nil {
		if err := o.Revision.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "revision")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "revision")
			}
			return err
		}
	}

	return nil
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBody) contextValidateUtype(ctx context.Context, formats strfmt.Registry) error {

	if o.Utype != nil {
		if err := o.Utype.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "utype")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "utype")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *EdgeNodeConfigurationUpdateEdgeNodeBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *EdgeNodeConfigurationUpdateEdgeNodeBody) UnmarshalBinary(b []byte) error {
	var res EdgeNodeConfigurationUpdateEdgeNodeBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
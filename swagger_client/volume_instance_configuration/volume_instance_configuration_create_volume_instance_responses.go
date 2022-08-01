// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package volume_instance_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// VolumeInstanceConfigurationCreateVolumeInstanceReader is a Reader for the VolumeInstanceConfigurationCreateVolumeInstance structure.
type VolumeInstanceConfigurationCreateVolumeInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VolumeInstanceConfigurationCreateVolumeInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVolumeInstanceConfigurationCreateVolumeInstanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewVolumeInstanceConfigurationCreateVolumeInstanceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewVolumeInstanceConfigurationCreateVolumeInstanceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewVolumeInstanceConfigurationCreateVolumeInstanceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewVolumeInstanceConfigurationCreateVolumeInstanceConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewVolumeInstanceConfigurationCreateVolumeInstanceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewVolumeInstanceConfigurationCreateVolumeInstanceGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewVolumeInstanceConfigurationCreateVolumeInstanceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVolumeInstanceConfigurationCreateVolumeInstanceOK creates a VolumeInstanceConfigurationCreateVolumeInstanceOK with default headers values
func NewVolumeInstanceConfigurationCreateVolumeInstanceOK() *VolumeInstanceConfigurationCreateVolumeInstanceOK {
	return &VolumeInstanceConfigurationCreateVolumeInstanceOK{}
}

/* VolumeInstanceConfigurationCreateVolumeInstanceOK describes a response with status code 200, with default header values.

A successful response.
*/
type VolumeInstanceConfigurationCreateVolumeInstanceOK struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceOK) Error() string {
	return fmt.Sprintf("[POST /v1/volumes/instances][%d] volumeInstanceConfigurationCreateVolumeInstanceOK  %+v", 200, o.Payload)
}
func (o *VolumeInstanceConfigurationCreateVolumeInstanceOK) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceConfigurationCreateVolumeInstanceBadRequest creates a VolumeInstanceConfigurationCreateVolumeInstanceBadRequest with default headers values
func NewVolumeInstanceConfigurationCreateVolumeInstanceBadRequest() *VolumeInstanceConfigurationCreateVolumeInstanceBadRequest {
	return &VolumeInstanceConfigurationCreateVolumeInstanceBadRequest{}
}

/* VolumeInstanceConfigurationCreateVolumeInstanceBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type VolumeInstanceConfigurationCreateVolumeInstanceBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/volumes/instances][%d] volumeInstanceConfigurationCreateVolumeInstanceBadRequest  %+v", 400, o.Payload)
}
func (o *VolumeInstanceConfigurationCreateVolumeInstanceBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceConfigurationCreateVolumeInstanceUnauthorized creates a VolumeInstanceConfigurationCreateVolumeInstanceUnauthorized with default headers values
func NewVolumeInstanceConfigurationCreateVolumeInstanceUnauthorized() *VolumeInstanceConfigurationCreateVolumeInstanceUnauthorized {
	return &VolumeInstanceConfigurationCreateVolumeInstanceUnauthorized{}
}

/* VolumeInstanceConfigurationCreateVolumeInstanceUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type VolumeInstanceConfigurationCreateVolumeInstanceUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/volumes/instances][%d] volumeInstanceConfigurationCreateVolumeInstanceUnauthorized  %+v", 401, o.Payload)
}
func (o *VolumeInstanceConfigurationCreateVolumeInstanceUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceConfigurationCreateVolumeInstanceForbidden creates a VolumeInstanceConfigurationCreateVolumeInstanceForbidden with default headers values
func NewVolumeInstanceConfigurationCreateVolumeInstanceForbidden() *VolumeInstanceConfigurationCreateVolumeInstanceForbidden {
	return &VolumeInstanceConfigurationCreateVolumeInstanceForbidden{}
}

/* VolumeInstanceConfigurationCreateVolumeInstanceForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type VolumeInstanceConfigurationCreateVolumeInstanceForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/volumes/instances][%d] volumeInstanceConfigurationCreateVolumeInstanceForbidden  %+v", 403, o.Payload)
}
func (o *VolumeInstanceConfigurationCreateVolumeInstanceForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceConfigurationCreateVolumeInstanceConflict creates a VolumeInstanceConfigurationCreateVolumeInstanceConflict with default headers values
func NewVolumeInstanceConfigurationCreateVolumeInstanceConflict() *VolumeInstanceConfigurationCreateVolumeInstanceConflict {
	return &VolumeInstanceConfigurationCreateVolumeInstanceConflict{}
}

/* VolumeInstanceConfigurationCreateVolumeInstanceConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this edge volume instance record will conflict with an already existing edge volume instance record.
*/
type VolumeInstanceConfigurationCreateVolumeInstanceConflict struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceConflict) Error() string {
	return fmt.Sprintf("[POST /v1/volumes/instances][%d] volumeInstanceConfigurationCreateVolumeInstanceConflict  %+v", 409, o.Payload)
}
func (o *VolumeInstanceConfigurationCreateVolumeInstanceConflict) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceConfigurationCreateVolumeInstanceInternalServerError creates a VolumeInstanceConfigurationCreateVolumeInstanceInternalServerError with default headers values
func NewVolumeInstanceConfigurationCreateVolumeInstanceInternalServerError() *VolumeInstanceConfigurationCreateVolumeInstanceInternalServerError {
	return &VolumeInstanceConfigurationCreateVolumeInstanceInternalServerError{}
}

/* VolumeInstanceConfigurationCreateVolumeInstanceInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type VolumeInstanceConfigurationCreateVolumeInstanceInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/volumes/instances][%d] volumeInstanceConfigurationCreateVolumeInstanceInternalServerError  %+v", 500, o.Payload)
}
func (o *VolumeInstanceConfigurationCreateVolumeInstanceInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceConfigurationCreateVolumeInstanceGatewayTimeout creates a VolumeInstanceConfigurationCreateVolumeInstanceGatewayTimeout with default headers values
func NewVolumeInstanceConfigurationCreateVolumeInstanceGatewayTimeout() *VolumeInstanceConfigurationCreateVolumeInstanceGatewayTimeout {
	return &VolumeInstanceConfigurationCreateVolumeInstanceGatewayTimeout{}
}

/* VolumeInstanceConfigurationCreateVolumeInstanceGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type VolumeInstanceConfigurationCreateVolumeInstanceGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v1/volumes/instances][%d] volumeInstanceConfigurationCreateVolumeInstanceGatewayTimeout  %+v", 504, o.Payload)
}
func (o *VolumeInstanceConfigurationCreateVolumeInstanceGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceConfigurationCreateVolumeInstanceDefault creates a VolumeInstanceConfigurationCreateVolumeInstanceDefault with default headers values
func NewVolumeInstanceConfigurationCreateVolumeInstanceDefault(code int) *VolumeInstanceConfigurationCreateVolumeInstanceDefault {
	return &VolumeInstanceConfigurationCreateVolumeInstanceDefault{
		_statusCode: code,
	}
}

/* VolumeInstanceConfigurationCreateVolumeInstanceDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type VolumeInstanceConfigurationCreateVolumeInstanceDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the volume instance configuration create volume instance default response
func (o *VolumeInstanceConfigurationCreateVolumeInstanceDefault) Code() int {
	return o._statusCode
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceDefault) Error() string {
	return fmt.Sprintf("[POST /v1/volumes/instances][%d] VolumeInstanceConfiguration_CreateVolumeInstance default  %+v", o._statusCode, o.Payload)
}
func (o *VolumeInstanceConfigurationCreateVolumeInstanceDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

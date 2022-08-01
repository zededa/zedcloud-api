// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package volume_instance_status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// VolumeInstanceStatusGetVolumeInstanceEventsByNameReader is a Reader for the VolumeInstanceStatusGetVolumeInstanceEventsByName structure.
type VolumeInstanceStatusGetVolumeInstanceEventsByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VolumeInstanceStatusGetVolumeInstanceEventsByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVolumeInstanceStatusGetVolumeInstanceEventsByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewVolumeInstanceStatusGetVolumeInstanceEventsByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewVolumeInstanceStatusGetVolumeInstanceEventsByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewVolumeInstanceStatusGetVolumeInstanceEventsByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewVolumeInstanceStatusGetVolumeInstanceEventsByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewVolumeInstanceStatusGetVolumeInstanceEventsByNameGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewVolumeInstanceStatusGetVolumeInstanceEventsByNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVolumeInstanceStatusGetVolumeInstanceEventsByNameOK creates a VolumeInstanceStatusGetVolumeInstanceEventsByNameOK with default headers values
func NewVolumeInstanceStatusGetVolumeInstanceEventsByNameOK() *VolumeInstanceStatusGetVolumeInstanceEventsByNameOK {
	return &VolumeInstanceStatusGetVolumeInstanceEventsByNameOK{}
}

/* VolumeInstanceStatusGetVolumeInstanceEventsByNameOK describes a response with status code 200, with default header values.

A successful response.
*/
type VolumeInstanceStatusGetVolumeInstanceEventsByNameOK struct {
	Payload *swagger_models.EventQueryResponse
}

func (o *VolumeInstanceStatusGetVolumeInstanceEventsByNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/name/{objname}/events][%d] volumeInstanceStatusGetVolumeInstanceEventsByNameOK  %+v", 200, o.Payload)
}
func (o *VolumeInstanceStatusGetVolumeInstanceEventsByNameOK) GetPayload() *swagger_models.EventQueryResponse {
	return o.Payload
}

func (o *VolumeInstanceStatusGetVolumeInstanceEventsByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.EventQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceStatusGetVolumeInstanceEventsByNameUnauthorized creates a VolumeInstanceStatusGetVolumeInstanceEventsByNameUnauthorized with default headers values
func NewVolumeInstanceStatusGetVolumeInstanceEventsByNameUnauthorized() *VolumeInstanceStatusGetVolumeInstanceEventsByNameUnauthorized {
	return &VolumeInstanceStatusGetVolumeInstanceEventsByNameUnauthorized{}
}

/* VolumeInstanceStatusGetVolumeInstanceEventsByNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type VolumeInstanceStatusGetVolumeInstanceEventsByNameUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *VolumeInstanceStatusGetVolumeInstanceEventsByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/name/{objname}/events][%d] volumeInstanceStatusGetVolumeInstanceEventsByNameUnauthorized  %+v", 401, o.Payload)
}
func (o *VolumeInstanceStatusGetVolumeInstanceEventsByNameUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceStatusGetVolumeInstanceEventsByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceStatusGetVolumeInstanceEventsByNameForbidden creates a VolumeInstanceStatusGetVolumeInstanceEventsByNameForbidden with default headers values
func NewVolumeInstanceStatusGetVolumeInstanceEventsByNameForbidden() *VolumeInstanceStatusGetVolumeInstanceEventsByNameForbidden {
	return &VolumeInstanceStatusGetVolumeInstanceEventsByNameForbidden{}
}

/* VolumeInstanceStatusGetVolumeInstanceEventsByNameForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type VolumeInstanceStatusGetVolumeInstanceEventsByNameForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *VolumeInstanceStatusGetVolumeInstanceEventsByNameForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/name/{objname}/events][%d] volumeInstanceStatusGetVolumeInstanceEventsByNameForbidden  %+v", 403, o.Payload)
}
func (o *VolumeInstanceStatusGetVolumeInstanceEventsByNameForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceStatusGetVolumeInstanceEventsByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceStatusGetVolumeInstanceEventsByNameNotFound creates a VolumeInstanceStatusGetVolumeInstanceEventsByNameNotFound with default headers values
func NewVolumeInstanceStatusGetVolumeInstanceEventsByNameNotFound() *VolumeInstanceStatusGetVolumeInstanceEventsByNameNotFound {
	return &VolumeInstanceStatusGetVolumeInstanceEventsByNameNotFound{}
}

/* VolumeInstanceStatusGetVolumeInstanceEventsByNameNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type VolumeInstanceStatusGetVolumeInstanceEventsByNameNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *VolumeInstanceStatusGetVolumeInstanceEventsByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/name/{objname}/events][%d] volumeInstanceStatusGetVolumeInstanceEventsByNameNotFound  %+v", 404, o.Payload)
}
func (o *VolumeInstanceStatusGetVolumeInstanceEventsByNameNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceStatusGetVolumeInstanceEventsByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceStatusGetVolumeInstanceEventsByNameInternalServerError creates a VolumeInstanceStatusGetVolumeInstanceEventsByNameInternalServerError with default headers values
func NewVolumeInstanceStatusGetVolumeInstanceEventsByNameInternalServerError() *VolumeInstanceStatusGetVolumeInstanceEventsByNameInternalServerError {
	return &VolumeInstanceStatusGetVolumeInstanceEventsByNameInternalServerError{}
}

/* VolumeInstanceStatusGetVolumeInstanceEventsByNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type VolumeInstanceStatusGetVolumeInstanceEventsByNameInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *VolumeInstanceStatusGetVolumeInstanceEventsByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/name/{objname}/events][%d] volumeInstanceStatusGetVolumeInstanceEventsByNameInternalServerError  %+v", 500, o.Payload)
}
func (o *VolumeInstanceStatusGetVolumeInstanceEventsByNameInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceStatusGetVolumeInstanceEventsByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceStatusGetVolumeInstanceEventsByNameGatewayTimeout creates a VolumeInstanceStatusGetVolumeInstanceEventsByNameGatewayTimeout with default headers values
func NewVolumeInstanceStatusGetVolumeInstanceEventsByNameGatewayTimeout() *VolumeInstanceStatusGetVolumeInstanceEventsByNameGatewayTimeout {
	return &VolumeInstanceStatusGetVolumeInstanceEventsByNameGatewayTimeout{}
}

/* VolumeInstanceStatusGetVolumeInstanceEventsByNameGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type VolumeInstanceStatusGetVolumeInstanceEventsByNameGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *VolumeInstanceStatusGetVolumeInstanceEventsByNameGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/name/{objname}/events][%d] volumeInstanceStatusGetVolumeInstanceEventsByNameGatewayTimeout  %+v", 504, o.Payload)
}
func (o *VolumeInstanceStatusGetVolumeInstanceEventsByNameGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceStatusGetVolumeInstanceEventsByNameGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceStatusGetVolumeInstanceEventsByNameDefault creates a VolumeInstanceStatusGetVolumeInstanceEventsByNameDefault with default headers values
func NewVolumeInstanceStatusGetVolumeInstanceEventsByNameDefault(code int) *VolumeInstanceStatusGetVolumeInstanceEventsByNameDefault {
	return &VolumeInstanceStatusGetVolumeInstanceEventsByNameDefault{
		_statusCode: code,
	}
}

/* VolumeInstanceStatusGetVolumeInstanceEventsByNameDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type VolumeInstanceStatusGetVolumeInstanceEventsByNameDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the volume instance status get volume instance events by name default response
func (o *VolumeInstanceStatusGetVolumeInstanceEventsByNameDefault) Code() int {
	return o._statusCode
}

func (o *VolumeInstanceStatusGetVolumeInstanceEventsByNameDefault) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/name/{objname}/events][%d] VolumeInstanceStatus_GetVolumeInstanceEventsByName default  %+v", o._statusCode, o.Payload)
}
func (o *VolumeInstanceStatusGetVolumeInstanceEventsByNameDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *VolumeInstanceStatusGetVolumeInstanceEventsByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

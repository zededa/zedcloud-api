// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package hardware_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// HardwareModelGetDeviceStatusConfigReader is a Reader for the HardwareModelGetDeviceStatusConfig structure.
type HardwareModelGetDeviceStatusConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HardwareModelGetDeviceStatusConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHardwareModelGetDeviceStatusConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewHardwareModelGetDeviceStatusConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewHardwareModelGetDeviceStatusConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewHardwareModelGetDeviceStatusConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewHardwareModelGetDeviceStatusConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewHardwareModelGetDeviceStatusConfigGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewHardwareModelGetDeviceStatusConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHardwareModelGetDeviceStatusConfigOK creates a HardwareModelGetDeviceStatusConfigOK with default headers values
func NewHardwareModelGetDeviceStatusConfigOK() *HardwareModelGetDeviceStatusConfigOK {
	return &HardwareModelGetDeviceStatusConfigOK{}
}

/* HardwareModelGetDeviceStatusConfigOK describes a response with status code 200, with default header values.

A successful response.
*/
type HardwareModelGetDeviceStatusConfigOK struct {
	Payload *swagger_models.DeviceStatusConfigList
}

func (o *HardwareModelGetDeviceStatusConfigOK) Error() string {
	return fmt.Sprintf("[GET /v1/devices/status-config][%d] hardwareModelGetDeviceStatusConfigOK  %+v", 200, o.Payload)
}
func (o *HardwareModelGetDeviceStatusConfigOK) GetPayload() *swagger_models.DeviceStatusConfigList {
	return o.Payload
}

func (o *HardwareModelGetDeviceStatusConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.DeviceStatusConfigList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetDeviceStatusConfigBadRequest creates a HardwareModelGetDeviceStatusConfigBadRequest with default headers values
func NewHardwareModelGetDeviceStatusConfigBadRequest() *HardwareModelGetDeviceStatusConfigBadRequest {
	return &HardwareModelGetDeviceStatusConfigBadRequest{}
}

/* HardwareModelGetDeviceStatusConfigBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of invalid value of filter parameters.
*/
type HardwareModelGetDeviceStatusConfigBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *HardwareModelGetDeviceStatusConfigBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/devices/status-config][%d] hardwareModelGetDeviceStatusConfigBadRequest  %+v", 400, o.Payload)
}
func (o *HardwareModelGetDeviceStatusConfigBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetDeviceStatusConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetDeviceStatusConfigUnauthorized creates a HardwareModelGetDeviceStatusConfigUnauthorized with default headers values
func NewHardwareModelGetDeviceStatusConfigUnauthorized() *HardwareModelGetDeviceStatusConfigUnauthorized {
	return &HardwareModelGetDeviceStatusConfigUnauthorized{}
}

/* HardwareModelGetDeviceStatusConfigUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type HardwareModelGetDeviceStatusConfigUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *HardwareModelGetDeviceStatusConfigUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/devices/status-config][%d] hardwareModelGetDeviceStatusConfigUnauthorized  %+v", 401, o.Payload)
}
func (o *HardwareModelGetDeviceStatusConfigUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetDeviceStatusConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetDeviceStatusConfigForbidden creates a HardwareModelGetDeviceStatusConfigForbidden with default headers values
func NewHardwareModelGetDeviceStatusConfigForbidden() *HardwareModelGetDeviceStatusConfigForbidden {
	return &HardwareModelGetDeviceStatusConfigForbidden{}
}

/* HardwareModelGetDeviceStatusConfigForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type HardwareModelGetDeviceStatusConfigForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *HardwareModelGetDeviceStatusConfigForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/devices/status-config][%d] hardwareModelGetDeviceStatusConfigForbidden  %+v", 403, o.Payload)
}
func (o *HardwareModelGetDeviceStatusConfigForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetDeviceStatusConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetDeviceStatusConfigInternalServerError creates a HardwareModelGetDeviceStatusConfigInternalServerError with default headers values
func NewHardwareModelGetDeviceStatusConfigInternalServerError() *HardwareModelGetDeviceStatusConfigInternalServerError {
	return &HardwareModelGetDeviceStatusConfigInternalServerError{}
}

/* HardwareModelGetDeviceStatusConfigInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type HardwareModelGetDeviceStatusConfigInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *HardwareModelGetDeviceStatusConfigInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/devices/status-config][%d] hardwareModelGetDeviceStatusConfigInternalServerError  %+v", 500, o.Payload)
}
func (o *HardwareModelGetDeviceStatusConfigInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetDeviceStatusConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetDeviceStatusConfigGatewayTimeout creates a HardwareModelGetDeviceStatusConfigGatewayTimeout with default headers values
func NewHardwareModelGetDeviceStatusConfigGatewayTimeout() *HardwareModelGetDeviceStatusConfigGatewayTimeout {
	return &HardwareModelGetDeviceStatusConfigGatewayTimeout{}
}

/* HardwareModelGetDeviceStatusConfigGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type HardwareModelGetDeviceStatusConfigGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *HardwareModelGetDeviceStatusConfigGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/devices/status-config][%d] hardwareModelGetDeviceStatusConfigGatewayTimeout  %+v", 504, o.Payload)
}
func (o *HardwareModelGetDeviceStatusConfigGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetDeviceStatusConfigGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetDeviceStatusConfigDefault creates a HardwareModelGetDeviceStatusConfigDefault with default headers values
func NewHardwareModelGetDeviceStatusConfigDefault(code int) *HardwareModelGetDeviceStatusConfigDefault {
	return &HardwareModelGetDeviceStatusConfigDefault{
		_statusCode: code,
	}
}

/* HardwareModelGetDeviceStatusConfigDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type HardwareModelGetDeviceStatusConfigDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the hardware model get device status config default response
func (o *HardwareModelGetDeviceStatusConfigDefault) Code() int {
	return o._statusCode
}

func (o *HardwareModelGetDeviceStatusConfigDefault) Error() string {
	return fmt.Sprintf("[GET /v1/devices/status-config][%d] HardwareModel_GetDeviceStatusConfig default  %+v", o._statusCode, o.Payload)
}
func (o *HardwareModelGetDeviceStatusConfigDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *HardwareModelGetDeviceStatusConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

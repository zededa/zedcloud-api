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

// HardwareModelGetDeviceTagsReader is a Reader for the HardwareModelGetDeviceTags structure.
type HardwareModelGetDeviceTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HardwareModelGetDeviceTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHardwareModelGetDeviceTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewHardwareModelGetDeviceTagsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewHardwareModelGetDeviceTagsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewHardwareModelGetDeviceTagsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewHardwareModelGetDeviceTagsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewHardwareModelGetDeviceTagsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewHardwareModelGetDeviceTagsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHardwareModelGetDeviceTagsOK creates a HardwareModelGetDeviceTagsOK with default headers values
func NewHardwareModelGetDeviceTagsOK() *HardwareModelGetDeviceTagsOK {
	return &HardwareModelGetDeviceTagsOK{}
}

/* HardwareModelGetDeviceTagsOK describes a response with status code 200, with default header values.

A successful response.
*/
type HardwareModelGetDeviceTagsOK struct {
	Payload *swagger_models.ObjectTagsList
}

func (o *HardwareModelGetDeviceTagsOK) Error() string {
	return fmt.Sprintf("[GET /v1/devices/tags][%d] hardwareModelGetDeviceTagsOK  %+v", 200, o.Payload)
}
func (o *HardwareModelGetDeviceTagsOK) GetPayload() *swagger_models.ObjectTagsList {
	return o.Payload
}

func (o *HardwareModelGetDeviceTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ObjectTagsList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetDeviceTagsBadRequest creates a HardwareModelGetDeviceTagsBadRequest with default headers values
func NewHardwareModelGetDeviceTagsBadRequest() *HardwareModelGetDeviceTagsBadRequest {
	return &HardwareModelGetDeviceTagsBadRequest{}
}

/* HardwareModelGetDeviceTagsBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of invalid value of filter parameters.
*/
type HardwareModelGetDeviceTagsBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *HardwareModelGetDeviceTagsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/devices/tags][%d] hardwareModelGetDeviceTagsBadRequest  %+v", 400, o.Payload)
}
func (o *HardwareModelGetDeviceTagsBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetDeviceTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetDeviceTagsUnauthorized creates a HardwareModelGetDeviceTagsUnauthorized with default headers values
func NewHardwareModelGetDeviceTagsUnauthorized() *HardwareModelGetDeviceTagsUnauthorized {
	return &HardwareModelGetDeviceTagsUnauthorized{}
}

/* HardwareModelGetDeviceTagsUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type HardwareModelGetDeviceTagsUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *HardwareModelGetDeviceTagsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/devices/tags][%d] hardwareModelGetDeviceTagsUnauthorized  %+v", 401, o.Payload)
}
func (o *HardwareModelGetDeviceTagsUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetDeviceTagsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetDeviceTagsForbidden creates a HardwareModelGetDeviceTagsForbidden with default headers values
func NewHardwareModelGetDeviceTagsForbidden() *HardwareModelGetDeviceTagsForbidden {
	return &HardwareModelGetDeviceTagsForbidden{}
}

/* HardwareModelGetDeviceTagsForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type HardwareModelGetDeviceTagsForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *HardwareModelGetDeviceTagsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/devices/tags][%d] hardwareModelGetDeviceTagsForbidden  %+v", 403, o.Payload)
}
func (o *HardwareModelGetDeviceTagsForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetDeviceTagsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetDeviceTagsInternalServerError creates a HardwareModelGetDeviceTagsInternalServerError with default headers values
func NewHardwareModelGetDeviceTagsInternalServerError() *HardwareModelGetDeviceTagsInternalServerError {
	return &HardwareModelGetDeviceTagsInternalServerError{}
}

/* HardwareModelGetDeviceTagsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type HardwareModelGetDeviceTagsInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *HardwareModelGetDeviceTagsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/devices/tags][%d] hardwareModelGetDeviceTagsInternalServerError  %+v", 500, o.Payload)
}
func (o *HardwareModelGetDeviceTagsInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetDeviceTagsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetDeviceTagsGatewayTimeout creates a HardwareModelGetDeviceTagsGatewayTimeout with default headers values
func NewHardwareModelGetDeviceTagsGatewayTimeout() *HardwareModelGetDeviceTagsGatewayTimeout {
	return &HardwareModelGetDeviceTagsGatewayTimeout{}
}

/* HardwareModelGetDeviceTagsGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type HardwareModelGetDeviceTagsGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *HardwareModelGetDeviceTagsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/devices/tags][%d] hardwareModelGetDeviceTagsGatewayTimeout  %+v", 504, o.Payload)
}
func (o *HardwareModelGetDeviceTagsGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetDeviceTagsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetDeviceTagsDefault creates a HardwareModelGetDeviceTagsDefault with default headers values
func NewHardwareModelGetDeviceTagsDefault(code int) *HardwareModelGetDeviceTagsDefault {
	return &HardwareModelGetDeviceTagsDefault{
		_statusCode: code,
	}
}

/* HardwareModelGetDeviceTagsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type HardwareModelGetDeviceTagsDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the hardware model get device tags default response
func (o *HardwareModelGetDeviceTagsDefault) Code() int {
	return o._statusCode
}

func (o *HardwareModelGetDeviceTagsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/devices/tags][%d] HardwareModel_GetDeviceTags default  %+v", o._statusCode, o.Payload)
}
func (o *HardwareModelGetDeviceTagsDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *HardwareModelGetDeviceTagsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
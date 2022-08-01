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

// HardwareModelGetHardwareModelReader is a Reader for the HardwareModelGetHardwareModel structure.
type HardwareModelGetHardwareModelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HardwareModelGetHardwareModelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHardwareModelGetHardwareModelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewHardwareModelGetHardwareModelUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewHardwareModelGetHardwareModelForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewHardwareModelGetHardwareModelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewHardwareModelGetHardwareModelInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewHardwareModelGetHardwareModelGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewHardwareModelGetHardwareModelDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHardwareModelGetHardwareModelOK creates a HardwareModelGetHardwareModelOK with default headers values
func NewHardwareModelGetHardwareModelOK() *HardwareModelGetHardwareModelOK {
	return &HardwareModelGetHardwareModelOK{}
}

/* HardwareModelGetHardwareModelOK describes a response with status code 200, with default header values.

A successful response.
*/
type HardwareModelGetHardwareModelOK struct {
	Payload *swagger_models.SysModel
}

func (o *HardwareModelGetHardwareModelOK) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/id/{id}][%d] hardwareModelGetHardwareModelOK  %+v", 200, o.Payload)
}
func (o *HardwareModelGetHardwareModelOK) GetPayload() *swagger_models.SysModel {
	return o.Payload
}

func (o *HardwareModelGetHardwareModelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.SysModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetHardwareModelUnauthorized creates a HardwareModelGetHardwareModelUnauthorized with default headers values
func NewHardwareModelGetHardwareModelUnauthorized() *HardwareModelGetHardwareModelUnauthorized {
	return &HardwareModelGetHardwareModelUnauthorized{}
}

/* HardwareModelGetHardwareModelUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type HardwareModelGetHardwareModelUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *HardwareModelGetHardwareModelUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/id/{id}][%d] hardwareModelGetHardwareModelUnauthorized  %+v", 401, o.Payload)
}
func (o *HardwareModelGetHardwareModelUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetHardwareModelUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetHardwareModelForbidden creates a HardwareModelGetHardwareModelForbidden with default headers values
func NewHardwareModelGetHardwareModelForbidden() *HardwareModelGetHardwareModelForbidden {
	return &HardwareModelGetHardwareModelForbidden{}
}

/* HardwareModelGetHardwareModelForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type HardwareModelGetHardwareModelForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *HardwareModelGetHardwareModelForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/id/{id}][%d] hardwareModelGetHardwareModelForbidden  %+v", 403, o.Payload)
}
func (o *HardwareModelGetHardwareModelForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetHardwareModelForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetHardwareModelNotFound creates a HardwareModelGetHardwareModelNotFound with default headers values
func NewHardwareModelGetHardwareModelNotFound() *HardwareModelGetHardwareModelNotFound {
	return &HardwareModelGetHardwareModelNotFound{}
}

/* HardwareModelGetHardwareModelNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type HardwareModelGetHardwareModelNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *HardwareModelGetHardwareModelNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/id/{id}][%d] hardwareModelGetHardwareModelNotFound  %+v", 404, o.Payload)
}
func (o *HardwareModelGetHardwareModelNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetHardwareModelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetHardwareModelInternalServerError creates a HardwareModelGetHardwareModelInternalServerError with default headers values
func NewHardwareModelGetHardwareModelInternalServerError() *HardwareModelGetHardwareModelInternalServerError {
	return &HardwareModelGetHardwareModelInternalServerError{}
}

/* HardwareModelGetHardwareModelInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type HardwareModelGetHardwareModelInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *HardwareModelGetHardwareModelInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/id/{id}][%d] hardwareModelGetHardwareModelInternalServerError  %+v", 500, o.Payload)
}
func (o *HardwareModelGetHardwareModelInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetHardwareModelInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetHardwareModelGatewayTimeout creates a HardwareModelGetHardwareModelGatewayTimeout with default headers values
func NewHardwareModelGetHardwareModelGatewayTimeout() *HardwareModelGetHardwareModelGatewayTimeout {
	return &HardwareModelGetHardwareModelGatewayTimeout{}
}

/* HardwareModelGetHardwareModelGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type HardwareModelGetHardwareModelGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *HardwareModelGetHardwareModelGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/id/{id}][%d] hardwareModelGetHardwareModelGatewayTimeout  %+v", 504, o.Payload)
}
func (o *HardwareModelGetHardwareModelGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetHardwareModelGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetHardwareModelDefault creates a HardwareModelGetHardwareModelDefault with default headers values
func NewHardwareModelGetHardwareModelDefault(code int) *HardwareModelGetHardwareModelDefault {
	return &HardwareModelGetHardwareModelDefault{
		_statusCode: code,
	}
}

/* HardwareModelGetHardwareModelDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type HardwareModelGetHardwareModelDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the hardware model get hardware model default response
func (o *HardwareModelGetHardwareModelDefault) Code() int {
	return o._statusCode
}

func (o *HardwareModelGetHardwareModelDefault) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/id/{id}][%d] HardwareModel_GetHardwareModel default  %+v", o._statusCode, o.Payload)
}
func (o *HardwareModelGetHardwareModelDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *HardwareModelGetHardwareModelDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
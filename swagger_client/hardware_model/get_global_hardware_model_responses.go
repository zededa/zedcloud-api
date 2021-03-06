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

// GetGlobalHardwareModelReader is a Reader for the GetGlobalHardwareModel structure.
type GetGlobalHardwareModelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGlobalHardwareModelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGlobalHardwareModelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetGlobalHardwareModelUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGlobalHardwareModelForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGlobalHardwareModelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGlobalHardwareModelInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetGlobalHardwareModelGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGlobalHardwareModelOK creates a GetGlobalHardwareModelOK with default headers values
func NewGetGlobalHardwareModelOK() *GetGlobalHardwareModelOK {
	return &GetGlobalHardwareModelOK{}
}

/* GetGlobalHardwareModelOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetGlobalHardwareModelOK struct {
	Payload *swagger_models.SysModel
}

func (o *GetGlobalHardwareModelOK) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/global/id/{id}][%d] getGlobalHardwareModelOK  %+v", 200, o.Payload)
}
func (o *GetGlobalHardwareModelOK) GetPayload() *swagger_models.SysModel {
	return o.Payload
}

func (o *GetGlobalHardwareModelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.SysModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGlobalHardwareModelUnauthorized creates a GetGlobalHardwareModelUnauthorized with default headers values
func NewGetGlobalHardwareModelUnauthorized() *GetGlobalHardwareModelUnauthorized {
	return &GetGlobalHardwareModelUnauthorized{}
}

/* GetGlobalHardwareModelUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type GetGlobalHardwareModelUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetGlobalHardwareModelUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/global/id/{id}][%d] getGlobalHardwareModelUnauthorized  %+v", 401, o.Payload)
}
func (o *GetGlobalHardwareModelUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetGlobalHardwareModelUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGlobalHardwareModelForbidden creates a GetGlobalHardwareModelForbidden with default headers values
func NewGetGlobalHardwareModelForbidden() *GetGlobalHardwareModelForbidden {
	return &GetGlobalHardwareModelForbidden{}
}

/* GetGlobalHardwareModelForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type GetGlobalHardwareModelForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetGlobalHardwareModelForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/global/id/{id}][%d] getGlobalHardwareModelForbidden  %+v", 403, o.Payload)
}
func (o *GetGlobalHardwareModelForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetGlobalHardwareModelForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGlobalHardwareModelNotFound creates a GetGlobalHardwareModelNotFound with default headers values
func NewGetGlobalHardwareModelNotFound() *GetGlobalHardwareModelNotFound {
	return &GetGlobalHardwareModelNotFound{}
}

/* GetGlobalHardwareModelNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type GetGlobalHardwareModelNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetGlobalHardwareModelNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/global/id/{id}][%d] getGlobalHardwareModelNotFound  %+v", 404, o.Payload)
}
func (o *GetGlobalHardwareModelNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetGlobalHardwareModelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGlobalHardwareModelInternalServerError creates a GetGlobalHardwareModelInternalServerError with default headers values
func NewGetGlobalHardwareModelInternalServerError() *GetGlobalHardwareModelInternalServerError {
	return &GetGlobalHardwareModelInternalServerError{}
}

/* GetGlobalHardwareModelInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type GetGlobalHardwareModelInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetGlobalHardwareModelInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/global/id/{id}][%d] getGlobalHardwareModelInternalServerError  %+v", 500, o.Payload)
}
func (o *GetGlobalHardwareModelInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetGlobalHardwareModelInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGlobalHardwareModelGatewayTimeout creates a GetGlobalHardwareModelGatewayTimeout with default headers values
func NewGetGlobalHardwareModelGatewayTimeout() *GetGlobalHardwareModelGatewayTimeout {
	return &GetGlobalHardwareModelGatewayTimeout{}
}

/* GetGlobalHardwareModelGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type GetGlobalHardwareModelGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetGlobalHardwareModelGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/global/id/{id}][%d] getGlobalHardwareModelGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetGlobalHardwareModelGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetGlobalHardwareModelGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

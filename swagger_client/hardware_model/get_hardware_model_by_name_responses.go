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

// GetHardwareModelByNameReader is a Reader for the GetHardwareModelByName structure.
type GetHardwareModelByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHardwareModelByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHardwareModelByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetHardwareModelByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetHardwareModelByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetHardwareModelByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetHardwareModelByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetHardwareModelByNameGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetHardwareModelByNameOK creates a GetHardwareModelByNameOK with default headers values
func NewGetHardwareModelByNameOK() *GetHardwareModelByNameOK {
	return &GetHardwareModelByNameOK{}
}

/* GetHardwareModelByNameOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetHardwareModelByNameOK struct {
	Payload *swagger_models.SysModel
}

func (o *GetHardwareModelByNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/name/{name}][%d] getHardwareModelByNameOK  %+v", 200, o.Payload)
}
func (o *GetHardwareModelByNameOK) GetPayload() *swagger_models.SysModel {
	return o.Payload
}

func (o *GetHardwareModelByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.SysModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHardwareModelByNameUnauthorized creates a GetHardwareModelByNameUnauthorized with default headers values
func NewGetHardwareModelByNameUnauthorized() *GetHardwareModelByNameUnauthorized {
	return &GetHardwareModelByNameUnauthorized{}
}

/* GetHardwareModelByNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type GetHardwareModelByNameUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetHardwareModelByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/name/{name}][%d] getHardwareModelByNameUnauthorized  %+v", 401, o.Payload)
}
func (o *GetHardwareModelByNameUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetHardwareModelByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHardwareModelByNameForbidden creates a GetHardwareModelByNameForbidden with default headers values
func NewGetHardwareModelByNameForbidden() *GetHardwareModelByNameForbidden {
	return &GetHardwareModelByNameForbidden{}
}

/* GetHardwareModelByNameForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type GetHardwareModelByNameForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetHardwareModelByNameForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/name/{name}][%d] getHardwareModelByNameForbidden  %+v", 403, o.Payload)
}
func (o *GetHardwareModelByNameForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetHardwareModelByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHardwareModelByNameNotFound creates a GetHardwareModelByNameNotFound with default headers values
func NewGetHardwareModelByNameNotFound() *GetHardwareModelByNameNotFound {
	return &GetHardwareModelByNameNotFound{}
}

/* GetHardwareModelByNameNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type GetHardwareModelByNameNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetHardwareModelByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/name/{name}][%d] getHardwareModelByNameNotFound  %+v", 404, o.Payload)
}
func (o *GetHardwareModelByNameNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetHardwareModelByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHardwareModelByNameInternalServerError creates a GetHardwareModelByNameInternalServerError with default headers values
func NewGetHardwareModelByNameInternalServerError() *GetHardwareModelByNameInternalServerError {
	return &GetHardwareModelByNameInternalServerError{}
}

/* GetHardwareModelByNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type GetHardwareModelByNameInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetHardwareModelByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/name/{name}][%d] getHardwareModelByNameInternalServerError  %+v", 500, o.Payload)
}
func (o *GetHardwareModelByNameInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetHardwareModelByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHardwareModelByNameGatewayTimeout creates a GetHardwareModelByNameGatewayTimeout with default headers values
func NewGetHardwareModelByNameGatewayTimeout() *GetHardwareModelByNameGatewayTimeout {
	return &GetHardwareModelByNameGatewayTimeout{}
}

/* GetHardwareModelByNameGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type GetHardwareModelByNameGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetHardwareModelByNameGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/name/{name}][%d] getHardwareModelByNameGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetHardwareModelByNameGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetHardwareModelByNameGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package edge_diagnostics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// GetDeviceTwinNextConfigByNameReader is a Reader for the GetDeviceTwinNextConfigByName structure.
type GetDeviceTwinNextConfigByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceTwinNextConfigByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceTwinNextConfigByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDeviceTwinNextConfigByNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDeviceTwinNextConfigByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDeviceTwinNextConfigByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDeviceTwinNextConfigByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDeviceTwinNextConfigByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetDeviceTwinNextConfigByNameGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDeviceTwinNextConfigByNameOK creates a GetDeviceTwinNextConfigByNameOK with default headers values
func NewGetDeviceTwinNextConfigByNameOK() *GetDeviceTwinNextConfigByNameOK {
	return &GetDeviceTwinNextConfigByNameOK{}
}

/* GetDeviceTwinNextConfigByNameOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetDeviceTwinNextConfigByNameOK struct {
	Payload *swagger_models.ConfigServiceResp
}

func (o *GetDeviceTwinNextConfigByNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/config/next][%d] getDeviceTwinNextConfigByNameOK  %+v", 200, o.Payload)
}
func (o *GetDeviceTwinNextConfigByNameOK) GetPayload() *swagger_models.ConfigServiceResp {
	return o.Payload
}

func (o *GetDeviceTwinNextConfigByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ConfigServiceResp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceTwinNextConfigByNameBadRequest creates a GetDeviceTwinNextConfigByNameBadRequest with default headers values
func NewGetDeviceTwinNextConfigByNameBadRequest() *GetDeviceTwinNextConfigByNameBadRequest {
	return &GetDeviceTwinNextConfigByNameBadRequest{}
}

/* GetDeviceTwinNextConfigByNameBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type GetDeviceTwinNextConfigByNameBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetDeviceTwinNextConfigByNameBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/config/next][%d] getDeviceTwinNextConfigByNameBadRequest  %+v", 400, o.Payload)
}
func (o *GetDeviceTwinNextConfigByNameBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetDeviceTwinNextConfigByNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceTwinNextConfigByNameUnauthorized creates a GetDeviceTwinNextConfigByNameUnauthorized with default headers values
func NewGetDeviceTwinNextConfigByNameUnauthorized() *GetDeviceTwinNextConfigByNameUnauthorized {
	return &GetDeviceTwinNextConfigByNameUnauthorized{}
}

/* GetDeviceTwinNextConfigByNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type GetDeviceTwinNextConfigByNameUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetDeviceTwinNextConfigByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/config/next][%d] getDeviceTwinNextConfigByNameUnauthorized  %+v", 401, o.Payload)
}
func (o *GetDeviceTwinNextConfigByNameUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetDeviceTwinNextConfigByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceTwinNextConfigByNameForbidden creates a GetDeviceTwinNextConfigByNameForbidden with default headers values
func NewGetDeviceTwinNextConfigByNameForbidden() *GetDeviceTwinNextConfigByNameForbidden {
	return &GetDeviceTwinNextConfigByNameForbidden{}
}

/* GetDeviceTwinNextConfigByNameForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type GetDeviceTwinNextConfigByNameForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetDeviceTwinNextConfigByNameForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/config/next][%d] getDeviceTwinNextConfigByNameForbidden  %+v", 403, o.Payload)
}
func (o *GetDeviceTwinNextConfigByNameForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetDeviceTwinNextConfigByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceTwinNextConfigByNameNotFound creates a GetDeviceTwinNextConfigByNameNotFound with default headers values
func NewGetDeviceTwinNextConfigByNameNotFound() *GetDeviceTwinNextConfigByNameNotFound {
	return &GetDeviceTwinNextConfigByNameNotFound{}
}

/* GetDeviceTwinNextConfigByNameNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type GetDeviceTwinNextConfigByNameNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetDeviceTwinNextConfigByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/config/next][%d] getDeviceTwinNextConfigByNameNotFound  %+v", 404, o.Payload)
}
func (o *GetDeviceTwinNextConfigByNameNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetDeviceTwinNextConfigByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceTwinNextConfigByNameInternalServerError creates a GetDeviceTwinNextConfigByNameInternalServerError with default headers values
func NewGetDeviceTwinNextConfigByNameInternalServerError() *GetDeviceTwinNextConfigByNameInternalServerError {
	return &GetDeviceTwinNextConfigByNameInternalServerError{}
}

/* GetDeviceTwinNextConfigByNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type GetDeviceTwinNextConfigByNameInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetDeviceTwinNextConfigByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/config/next][%d] getDeviceTwinNextConfigByNameInternalServerError  %+v", 500, o.Payload)
}
func (o *GetDeviceTwinNextConfigByNameInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetDeviceTwinNextConfigByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceTwinNextConfigByNameGatewayTimeout creates a GetDeviceTwinNextConfigByNameGatewayTimeout with default headers values
func NewGetDeviceTwinNextConfigByNameGatewayTimeout() *GetDeviceTwinNextConfigByNameGatewayTimeout {
	return &GetDeviceTwinNextConfigByNameGatewayTimeout{}
}

/* GetDeviceTwinNextConfigByNameGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type GetDeviceTwinNextConfigByNameGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetDeviceTwinNextConfigByNameGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/config/next][%d] getDeviceTwinNextConfigByNameGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetDeviceTwinNextConfigByNameGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetDeviceTwinNextConfigByNameGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

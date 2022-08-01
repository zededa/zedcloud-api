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

// EdgeDiagnosticsGetDeviceTwinConfigByNameReader is a Reader for the EdgeDiagnosticsGetDeviceTwinConfigByName structure.
type EdgeDiagnosticsGetDeviceTwinConfigByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeDiagnosticsGetDeviceTwinConfigByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeDiagnosticsGetDeviceTwinConfigByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEdgeDiagnosticsGetDeviceTwinConfigByNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewEdgeDiagnosticsGetDeviceTwinConfigByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeDiagnosticsGetDeviceTwinConfigByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeDiagnosticsGetDeviceTwinConfigByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeDiagnosticsGetDeviceTwinConfigByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeDiagnosticsGetDeviceTwinConfigByNameGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeDiagnosticsGetDeviceTwinConfigByNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeDiagnosticsGetDeviceTwinConfigByNameOK creates a EdgeDiagnosticsGetDeviceTwinConfigByNameOK with default headers values
func NewEdgeDiagnosticsGetDeviceTwinConfigByNameOK() *EdgeDiagnosticsGetDeviceTwinConfigByNameOK {
	return &EdgeDiagnosticsGetDeviceTwinConfigByNameOK{}
}

/* EdgeDiagnosticsGetDeviceTwinConfigByNameOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeDiagnosticsGetDeviceTwinConfigByNameOK struct {
	Payload *swagger_models.ConfigServiceResp
}

func (o *EdgeDiagnosticsGetDeviceTwinConfigByNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/config][%d] edgeDiagnosticsGetDeviceTwinConfigByNameOK  %+v", 200, o.Payload)
}
func (o *EdgeDiagnosticsGetDeviceTwinConfigByNameOK) GetPayload() *swagger_models.ConfigServiceResp {
	return o.Payload
}

func (o *EdgeDiagnosticsGetDeviceTwinConfigByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ConfigServiceResp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeDiagnosticsGetDeviceTwinConfigByNameBadRequest creates a EdgeDiagnosticsGetDeviceTwinConfigByNameBadRequest with default headers values
func NewEdgeDiagnosticsGetDeviceTwinConfigByNameBadRequest() *EdgeDiagnosticsGetDeviceTwinConfigByNameBadRequest {
	return &EdgeDiagnosticsGetDeviceTwinConfigByNameBadRequest{}
}

/* EdgeDiagnosticsGetDeviceTwinConfigByNameBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type EdgeDiagnosticsGetDeviceTwinConfigByNameBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeDiagnosticsGetDeviceTwinConfigByNameBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/config][%d] edgeDiagnosticsGetDeviceTwinConfigByNameBadRequest  %+v", 400, o.Payload)
}
func (o *EdgeDiagnosticsGetDeviceTwinConfigByNameBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeDiagnosticsGetDeviceTwinConfigByNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeDiagnosticsGetDeviceTwinConfigByNameUnauthorized creates a EdgeDiagnosticsGetDeviceTwinConfigByNameUnauthorized with default headers values
func NewEdgeDiagnosticsGetDeviceTwinConfigByNameUnauthorized() *EdgeDiagnosticsGetDeviceTwinConfigByNameUnauthorized {
	return &EdgeDiagnosticsGetDeviceTwinConfigByNameUnauthorized{}
}

/* EdgeDiagnosticsGetDeviceTwinConfigByNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeDiagnosticsGetDeviceTwinConfigByNameUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeDiagnosticsGetDeviceTwinConfigByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/config][%d] edgeDiagnosticsGetDeviceTwinConfigByNameUnauthorized  %+v", 401, o.Payload)
}
func (o *EdgeDiagnosticsGetDeviceTwinConfigByNameUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeDiagnosticsGetDeviceTwinConfigByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeDiagnosticsGetDeviceTwinConfigByNameForbidden creates a EdgeDiagnosticsGetDeviceTwinConfigByNameForbidden with default headers values
func NewEdgeDiagnosticsGetDeviceTwinConfigByNameForbidden() *EdgeDiagnosticsGetDeviceTwinConfigByNameForbidden {
	return &EdgeDiagnosticsGetDeviceTwinConfigByNameForbidden{}
}

/* EdgeDiagnosticsGetDeviceTwinConfigByNameForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeDiagnosticsGetDeviceTwinConfigByNameForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeDiagnosticsGetDeviceTwinConfigByNameForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/config][%d] edgeDiagnosticsGetDeviceTwinConfigByNameForbidden  %+v", 403, o.Payload)
}
func (o *EdgeDiagnosticsGetDeviceTwinConfigByNameForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeDiagnosticsGetDeviceTwinConfigByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeDiagnosticsGetDeviceTwinConfigByNameNotFound creates a EdgeDiagnosticsGetDeviceTwinConfigByNameNotFound with default headers values
func NewEdgeDiagnosticsGetDeviceTwinConfigByNameNotFound() *EdgeDiagnosticsGetDeviceTwinConfigByNameNotFound {
	return &EdgeDiagnosticsGetDeviceTwinConfigByNameNotFound{}
}

/* EdgeDiagnosticsGetDeviceTwinConfigByNameNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeDiagnosticsGetDeviceTwinConfigByNameNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeDiagnosticsGetDeviceTwinConfigByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/config][%d] edgeDiagnosticsGetDeviceTwinConfigByNameNotFound  %+v", 404, o.Payload)
}
func (o *EdgeDiagnosticsGetDeviceTwinConfigByNameNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeDiagnosticsGetDeviceTwinConfigByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeDiagnosticsGetDeviceTwinConfigByNameInternalServerError creates a EdgeDiagnosticsGetDeviceTwinConfigByNameInternalServerError with default headers values
func NewEdgeDiagnosticsGetDeviceTwinConfigByNameInternalServerError() *EdgeDiagnosticsGetDeviceTwinConfigByNameInternalServerError {
	return &EdgeDiagnosticsGetDeviceTwinConfigByNameInternalServerError{}
}

/* EdgeDiagnosticsGetDeviceTwinConfigByNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeDiagnosticsGetDeviceTwinConfigByNameInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeDiagnosticsGetDeviceTwinConfigByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/config][%d] edgeDiagnosticsGetDeviceTwinConfigByNameInternalServerError  %+v", 500, o.Payload)
}
func (o *EdgeDiagnosticsGetDeviceTwinConfigByNameInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeDiagnosticsGetDeviceTwinConfigByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeDiagnosticsGetDeviceTwinConfigByNameGatewayTimeout creates a EdgeDiagnosticsGetDeviceTwinConfigByNameGatewayTimeout with default headers values
func NewEdgeDiagnosticsGetDeviceTwinConfigByNameGatewayTimeout() *EdgeDiagnosticsGetDeviceTwinConfigByNameGatewayTimeout {
	return &EdgeDiagnosticsGetDeviceTwinConfigByNameGatewayTimeout{}
}

/* EdgeDiagnosticsGetDeviceTwinConfigByNameGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeDiagnosticsGetDeviceTwinConfigByNameGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeDiagnosticsGetDeviceTwinConfigByNameGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/config][%d] edgeDiagnosticsGetDeviceTwinConfigByNameGatewayTimeout  %+v", 504, o.Payload)
}
func (o *EdgeDiagnosticsGetDeviceTwinConfigByNameGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeDiagnosticsGetDeviceTwinConfigByNameGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeDiagnosticsGetDeviceTwinConfigByNameDefault creates a EdgeDiagnosticsGetDeviceTwinConfigByNameDefault with default headers values
func NewEdgeDiagnosticsGetDeviceTwinConfigByNameDefault(code int) *EdgeDiagnosticsGetDeviceTwinConfigByNameDefault {
	return &EdgeDiagnosticsGetDeviceTwinConfigByNameDefault{
		_statusCode: code,
	}
}

/* EdgeDiagnosticsGetDeviceTwinConfigByNameDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeDiagnosticsGetDeviceTwinConfigByNameDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the edge diagnostics get device twin config by name default response
func (o *EdgeDiagnosticsGetDeviceTwinConfigByNameDefault) Code() int {
	return o._statusCode
}

func (o *EdgeDiagnosticsGetDeviceTwinConfigByNameDefault) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/config][%d] EdgeDiagnostics_GetDeviceTwinConfigByName default  %+v", o._statusCode, o.Payload)
}
func (o *EdgeDiagnosticsGetDeviceTwinConfigByNameDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeDiagnosticsGetDeviceTwinConfigByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

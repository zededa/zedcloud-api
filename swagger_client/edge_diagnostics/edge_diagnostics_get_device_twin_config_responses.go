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

// EdgeDiagnosticsGetDeviceTwinConfigReader is a Reader for the EdgeDiagnosticsGetDeviceTwinConfig structure.
type EdgeDiagnosticsGetDeviceTwinConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeDiagnosticsGetDeviceTwinConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeDiagnosticsGetDeviceTwinConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEdgeDiagnosticsGetDeviceTwinConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewEdgeDiagnosticsGetDeviceTwinConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeDiagnosticsGetDeviceTwinConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeDiagnosticsGetDeviceTwinConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeDiagnosticsGetDeviceTwinConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeDiagnosticsGetDeviceTwinConfigGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeDiagnosticsGetDeviceTwinConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeDiagnosticsGetDeviceTwinConfigOK creates a EdgeDiagnosticsGetDeviceTwinConfigOK with default headers values
func NewEdgeDiagnosticsGetDeviceTwinConfigOK() *EdgeDiagnosticsGetDeviceTwinConfigOK {
	return &EdgeDiagnosticsGetDeviceTwinConfigOK{}
}

/* EdgeDiagnosticsGetDeviceTwinConfigOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeDiagnosticsGetDeviceTwinConfigOK struct {
	Payload *swagger_models.ConfigServiceResp
}

func (o *EdgeDiagnosticsGetDeviceTwinConfigOK) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/config][%d] edgeDiagnosticsGetDeviceTwinConfigOK  %+v", 200, o.Payload)
}
func (o *EdgeDiagnosticsGetDeviceTwinConfigOK) GetPayload() *swagger_models.ConfigServiceResp {
	return o.Payload
}

func (o *EdgeDiagnosticsGetDeviceTwinConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ConfigServiceResp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeDiagnosticsGetDeviceTwinConfigBadRequest creates a EdgeDiagnosticsGetDeviceTwinConfigBadRequest with default headers values
func NewEdgeDiagnosticsGetDeviceTwinConfigBadRequest() *EdgeDiagnosticsGetDeviceTwinConfigBadRequest {
	return &EdgeDiagnosticsGetDeviceTwinConfigBadRequest{}
}

/* EdgeDiagnosticsGetDeviceTwinConfigBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type EdgeDiagnosticsGetDeviceTwinConfigBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeDiagnosticsGetDeviceTwinConfigBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/config][%d] edgeDiagnosticsGetDeviceTwinConfigBadRequest  %+v", 400, o.Payload)
}
func (o *EdgeDiagnosticsGetDeviceTwinConfigBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeDiagnosticsGetDeviceTwinConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeDiagnosticsGetDeviceTwinConfigUnauthorized creates a EdgeDiagnosticsGetDeviceTwinConfigUnauthorized with default headers values
func NewEdgeDiagnosticsGetDeviceTwinConfigUnauthorized() *EdgeDiagnosticsGetDeviceTwinConfigUnauthorized {
	return &EdgeDiagnosticsGetDeviceTwinConfigUnauthorized{}
}

/* EdgeDiagnosticsGetDeviceTwinConfigUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeDiagnosticsGetDeviceTwinConfigUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeDiagnosticsGetDeviceTwinConfigUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/config][%d] edgeDiagnosticsGetDeviceTwinConfigUnauthorized  %+v", 401, o.Payload)
}
func (o *EdgeDiagnosticsGetDeviceTwinConfigUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeDiagnosticsGetDeviceTwinConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeDiagnosticsGetDeviceTwinConfigForbidden creates a EdgeDiagnosticsGetDeviceTwinConfigForbidden with default headers values
func NewEdgeDiagnosticsGetDeviceTwinConfigForbidden() *EdgeDiagnosticsGetDeviceTwinConfigForbidden {
	return &EdgeDiagnosticsGetDeviceTwinConfigForbidden{}
}

/* EdgeDiagnosticsGetDeviceTwinConfigForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeDiagnosticsGetDeviceTwinConfigForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeDiagnosticsGetDeviceTwinConfigForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/config][%d] edgeDiagnosticsGetDeviceTwinConfigForbidden  %+v", 403, o.Payload)
}
func (o *EdgeDiagnosticsGetDeviceTwinConfigForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeDiagnosticsGetDeviceTwinConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeDiagnosticsGetDeviceTwinConfigNotFound creates a EdgeDiagnosticsGetDeviceTwinConfigNotFound with default headers values
func NewEdgeDiagnosticsGetDeviceTwinConfigNotFound() *EdgeDiagnosticsGetDeviceTwinConfigNotFound {
	return &EdgeDiagnosticsGetDeviceTwinConfigNotFound{}
}

/* EdgeDiagnosticsGetDeviceTwinConfigNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeDiagnosticsGetDeviceTwinConfigNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeDiagnosticsGetDeviceTwinConfigNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/config][%d] edgeDiagnosticsGetDeviceTwinConfigNotFound  %+v", 404, o.Payload)
}
func (o *EdgeDiagnosticsGetDeviceTwinConfigNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeDiagnosticsGetDeviceTwinConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeDiagnosticsGetDeviceTwinConfigInternalServerError creates a EdgeDiagnosticsGetDeviceTwinConfigInternalServerError with default headers values
func NewEdgeDiagnosticsGetDeviceTwinConfigInternalServerError() *EdgeDiagnosticsGetDeviceTwinConfigInternalServerError {
	return &EdgeDiagnosticsGetDeviceTwinConfigInternalServerError{}
}

/* EdgeDiagnosticsGetDeviceTwinConfigInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeDiagnosticsGetDeviceTwinConfigInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeDiagnosticsGetDeviceTwinConfigInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/config][%d] edgeDiagnosticsGetDeviceTwinConfigInternalServerError  %+v", 500, o.Payload)
}
func (o *EdgeDiagnosticsGetDeviceTwinConfigInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeDiagnosticsGetDeviceTwinConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeDiagnosticsGetDeviceTwinConfigGatewayTimeout creates a EdgeDiagnosticsGetDeviceTwinConfigGatewayTimeout with default headers values
func NewEdgeDiagnosticsGetDeviceTwinConfigGatewayTimeout() *EdgeDiagnosticsGetDeviceTwinConfigGatewayTimeout {
	return &EdgeDiagnosticsGetDeviceTwinConfigGatewayTimeout{}
}

/* EdgeDiagnosticsGetDeviceTwinConfigGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeDiagnosticsGetDeviceTwinConfigGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeDiagnosticsGetDeviceTwinConfigGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/config][%d] edgeDiagnosticsGetDeviceTwinConfigGatewayTimeout  %+v", 504, o.Payload)
}
func (o *EdgeDiagnosticsGetDeviceTwinConfigGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeDiagnosticsGetDeviceTwinConfigGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeDiagnosticsGetDeviceTwinConfigDefault creates a EdgeDiagnosticsGetDeviceTwinConfigDefault with default headers values
func NewEdgeDiagnosticsGetDeviceTwinConfigDefault(code int) *EdgeDiagnosticsGetDeviceTwinConfigDefault {
	return &EdgeDiagnosticsGetDeviceTwinConfigDefault{
		_statusCode: code,
	}
}

/* EdgeDiagnosticsGetDeviceTwinConfigDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeDiagnosticsGetDeviceTwinConfigDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the edge diagnostics get device twin config default response
func (o *EdgeDiagnosticsGetDeviceTwinConfigDefault) Code() int {
	return o._statusCode
}

func (o *EdgeDiagnosticsGetDeviceTwinConfigDefault) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/config][%d] EdgeDiagnostics_GetDeviceTwinConfig default  %+v", o._statusCode, o.Payload)
}
func (o *EdgeDiagnosticsGetDeviceTwinConfigDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeDiagnosticsGetDeviceTwinConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

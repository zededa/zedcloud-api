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

// EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameReader is a Reader for the EdgeDiagnosticsGetDeviceTwinOfflineConfigByName structure.
type EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeDiagnosticsGetDeviceTwinOfflineConfigByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEdgeDiagnosticsGetDeviceTwinOfflineConfigByNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewEdgeDiagnosticsGetDeviceTwinOfflineConfigByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeDiagnosticsGetDeviceTwinOfflineConfigByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeDiagnosticsGetDeviceTwinOfflineConfigByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeDiagnosticsGetDeviceTwinOfflineConfigByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeDiagnosticsGetDeviceTwinOfflineConfigByNameGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeDiagnosticsGetDeviceTwinOfflineConfigByNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeDiagnosticsGetDeviceTwinOfflineConfigByNameOK creates a EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameOK with default headers values
func NewEdgeDiagnosticsGetDeviceTwinOfflineConfigByNameOK() *EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameOK {
	return &EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameOK{}
}

/* EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameOK struct {
	Payload *swagger_models.ConfigServiceResp
}

func (o *EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/config/offline][%d] edgeDiagnosticsGetDeviceTwinOfflineConfigByNameOK  %+v", 200, o.Payload)
}
func (o *EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameOK) GetPayload() *swagger_models.ConfigServiceResp {
	return o.Payload
}

func (o *EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ConfigServiceResp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeDiagnosticsGetDeviceTwinOfflineConfigByNameBadRequest creates a EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameBadRequest with default headers values
func NewEdgeDiagnosticsGetDeviceTwinOfflineConfigByNameBadRequest() *EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameBadRequest {
	return &EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameBadRequest{}
}

/* EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/config/offline][%d] edgeDiagnosticsGetDeviceTwinOfflineConfigByNameBadRequest  %+v", 400, o.Payload)
}
func (o *EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeDiagnosticsGetDeviceTwinOfflineConfigByNameUnauthorized creates a EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameUnauthorized with default headers values
func NewEdgeDiagnosticsGetDeviceTwinOfflineConfigByNameUnauthorized() *EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameUnauthorized {
	return &EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameUnauthorized{}
}

/* EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/config/offline][%d] edgeDiagnosticsGetDeviceTwinOfflineConfigByNameUnauthorized  %+v", 401, o.Payload)
}
func (o *EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeDiagnosticsGetDeviceTwinOfflineConfigByNameForbidden creates a EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameForbidden with default headers values
func NewEdgeDiagnosticsGetDeviceTwinOfflineConfigByNameForbidden() *EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameForbidden {
	return &EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameForbidden{}
}

/* EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/config/offline][%d] edgeDiagnosticsGetDeviceTwinOfflineConfigByNameForbidden  %+v", 403, o.Payload)
}
func (o *EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeDiagnosticsGetDeviceTwinOfflineConfigByNameNotFound creates a EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameNotFound with default headers values
func NewEdgeDiagnosticsGetDeviceTwinOfflineConfigByNameNotFound() *EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameNotFound {
	return &EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameNotFound{}
}

/* EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/config/offline][%d] edgeDiagnosticsGetDeviceTwinOfflineConfigByNameNotFound  %+v", 404, o.Payload)
}
func (o *EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeDiagnosticsGetDeviceTwinOfflineConfigByNameInternalServerError creates a EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameInternalServerError with default headers values
func NewEdgeDiagnosticsGetDeviceTwinOfflineConfigByNameInternalServerError() *EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameInternalServerError {
	return &EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameInternalServerError{}
}

/* EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/config/offline][%d] edgeDiagnosticsGetDeviceTwinOfflineConfigByNameInternalServerError  %+v", 500, o.Payload)
}
func (o *EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeDiagnosticsGetDeviceTwinOfflineConfigByNameGatewayTimeout creates a EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameGatewayTimeout with default headers values
func NewEdgeDiagnosticsGetDeviceTwinOfflineConfigByNameGatewayTimeout() *EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameGatewayTimeout {
	return &EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameGatewayTimeout{}
}

/* EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/config/offline][%d] edgeDiagnosticsGetDeviceTwinOfflineConfigByNameGatewayTimeout  %+v", 504, o.Payload)
}
func (o *EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeDiagnosticsGetDeviceTwinOfflineConfigByNameDefault creates a EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameDefault with default headers values
func NewEdgeDiagnosticsGetDeviceTwinOfflineConfigByNameDefault(code int) *EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameDefault {
	return &EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameDefault{
		_statusCode: code,
	}
}

/* EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the edge diagnostics get device twin offline config by name default response
func (o *EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameDefault) Code() int {
	return o._statusCode
}

func (o *EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameDefault) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/config/offline][%d] EdgeDiagnostics_GetDeviceTwinOfflineConfigByName default  %+v", o._statusCode, o.Payload)
}
func (o *EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeDiagnosticsGetDeviceTwinOfflineConfigByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
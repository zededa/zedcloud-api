// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
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

// EdgeDiagnosticsGetDeviceTwinNextConfigByNameReader is a Reader for the EdgeDiagnosticsGetDeviceTwinNextConfigByName structure.
type EdgeDiagnosticsGetDeviceTwinNextConfigByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeDiagnosticsGetDeviceTwinNextConfigByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEdgeDiagnosticsGetDeviceTwinNextConfigByNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewEdgeDiagnosticsGetDeviceTwinNextConfigByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeDiagnosticsGetDeviceTwinNextConfigByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeDiagnosticsGetDeviceTwinNextConfigByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeDiagnosticsGetDeviceTwinNextConfigByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeDiagnosticsGetDeviceTwinNextConfigByNameGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeDiagnosticsGetDeviceTwinNextConfigByNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeDiagnosticsGetDeviceTwinNextConfigByNameOK creates a EdgeDiagnosticsGetDeviceTwinNextConfigByNameOK with default headers values
func NewEdgeDiagnosticsGetDeviceTwinNextConfigByNameOK() *EdgeDiagnosticsGetDeviceTwinNextConfigByNameOK {
	return &EdgeDiagnosticsGetDeviceTwinNextConfigByNameOK{}
}

/*
EdgeDiagnosticsGetDeviceTwinNextConfigByNameOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeDiagnosticsGetDeviceTwinNextConfigByNameOK struct {
	Payload *swagger_models.ConfigServiceResp
}

// IsSuccess returns true when this edge diagnostics get device twin next config by name o k response has a 2xx status code
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge diagnostics get device twin next config by name o k response has a 3xx status code
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge diagnostics get device twin next config by name o k response has a 4xx status code
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge diagnostics get device twin next config by name o k response has a 5xx status code
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge diagnostics get device twin next config by name o k response a status code equal to that given
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameOK) IsCode(code int) bool {
	return code == 200
}

func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/config/next][%d] edgeDiagnosticsGetDeviceTwinNextConfigByNameOK  %+v", 200, o.Payload)
}

func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameOK) String() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/config/next][%d] edgeDiagnosticsGetDeviceTwinNextConfigByNameOK  %+v", 200, o.Payload)
}

func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameOK) GetPayload() *swagger_models.ConfigServiceResp {
	return o.Payload
}

func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ConfigServiceResp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeDiagnosticsGetDeviceTwinNextConfigByNameBadRequest creates a EdgeDiagnosticsGetDeviceTwinNextConfigByNameBadRequest with default headers values
func NewEdgeDiagnosticsGetDeviceTwinNextConfigByNameBadRequest() *EdgeDiagnosticsGetDeviceTwinNextConfigByNameBadRequest {
	return &EdgeDiagnosticsGetDeviceTwinNextConfigByNameBadRequest{}
}

/*
EdgeDiagnosticsGetDeviceTwinNextConfigByNameBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type EdgeDiagnosticsGetDeviceTwinNextConfigByNameBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge diagnostics get device twin next config by name bad request response has a 2xx status code
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge diagnostics get device twin next config by name bad request response has a 3xx status code
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge diagnostics get device twin next config by name bad request response has a 4xx status code
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge diagnostics get device twin next config by name bad request response has a 5xx status code
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this edge diagnostics get device twin next config by name bad request response a status code equal to that given
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/config/next][%d] edgeDiagnosticsGetDeviceTwinNextConfigByNameBadRequest  %+v", 400, o.Payload)
}

func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/config/next][%d] edgeDiagnosticsGetDeviceTwinNextConfigByNameBadRequest  %+v", 400, o.Payload)
}

func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeDiagnosticsGetDeviceTwinNextConfigByNameUnauthorized creates a EdgeDiagnosticsGetDeviceTwinNextConfigByNameUnauthorized with default headers values
func NewEdgeDiagnosticsGetDeviceTwinNextConfigByNameUnauthorized() *EdgeDiagnosticsGetDeviceTwinNextConfigByNameUnauthorized {
	return &EdgeDiagnosticsGetDeviceTwinNextConfigByNameUnauthorized{}
}

/*
EdgeDiagnosticsGetDeviceTwinNextConfigByNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeDiagnosticsGetDeviceTwinNextConfigByNameUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge diagnostics get device twin next config by name unauthorized response has a 2xx status code
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge diagnostics get device twin next config by name unauthorized response has a 3xx status code
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge diagnostics get device twin next config by name unauthorized response has a 4xx status code
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge diagnostics get device twin next config by name unauthorized response has a 5xx status code
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge diagnostics get device twin next config by name unauthorized response a status code equal to that given
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/config/next][%d] edgeDiagnosticsGetDeviceTwinNextConfigByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/config/next][%d] edgeDiagnosticsGetDeviceTwinNextConfigByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeDiagnosticsGetDeviceTwinNextConfigByNameForbidden creates a EdgeDiagnosticsGetDeviceTwinNextConfigByNameForbidden with default headers values
func NewEdgeDiagnosticsGetDeviceTwinNextConfigByNameForbidden() *EdgeDiagnosticsGetDeviceTwinNextConfigByNameForbidden {
	return &EdgeDiagnosticsGetDeviceTwinNextConfigByNameForbidden{}
}

/*
EdgeDiagnosticsGetDeviceTwinNextConfigByNameForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeDiagnosticsGetDeviceTwinNextConfigByNameForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge diagnostics get device twin next config by name forbidden response has a 2xx status code
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge diagnostics get device twin next config by name forbidden response has a 3xx status code
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge diagnostics get device twin next config by name forbidden response has a 4xx status code
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge diagnostics get device twin next config by name forbidden response has a 5xx status code
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge diagnostics get device twin next config by name forbidden response a status code equal to that given
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/config/next][%d] edgeDiagnosticsGetDeviceTwinNextConfigByNameForbidden  %+v", 403, o.Payload)
}

func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameForbidden) String() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/config/next][%d] edgeDiagnosticsGetDeviceTwinNextConfigByNameForbidden  %+v", 403, o.Payload)
}

func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeDiagnosticsGetDeviceTwinNextConfigByNameNotFound creates a EdgeDiagnosticsGetDeviceTwinNextConfigByNameNotFound with default headers values
func NewEdgeDiagnosticsGetDeviceTwinNextConfigByNameNotFound() *EdgeDiagnosticsGetDeviceTwinNextConfigByNameNotFound {
	return &EdgeDiagnosticsGetDeviceTwinNextConfigByNameNotFound{}
}

/*
EdgeDiagnosticsGetDeviceTwinNextConfigByNameNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeDiagnosticsGetDeviceTwinNextConfigByNameNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge diagnostics get device twin next config by name not found response has a 2xx status code
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge diagnostics get device twin next config by name not found response has a 3xx status code
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge diagnostics get device twin next config by name not found response has a 4xx status code
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge diagnostics get device twin next config by name not found response has a 5xx status code
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge diagnostics get device twin next config by name not found response a status code equal to that given
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/config/next][%d] edgeDiagnosticsGetDeviceTwinNextConfigByNameNotFound  %+v", 404, o.Payload)
}

func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameNotFound) String() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/config/next][%d] edgeDiagnosticsGetDeviceTwinNextConfigByNameNotFound  %+v", 404, o.Payload)
}

func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeDiagnosticsGetDeviceTwinNextConfigByNameInternalServerError creates a EdgeDiagnosticsGetDeviceTwinNextConfigByNameInternalServerError with default headers values
func NewEdgeDiagnosticsGetDeviceTwinNextConfigByNameInternalServerError() *EdgeDiagnosticsGetDeviceTwinNextConfigByNameInternalServerError {
	return &EdgeDiagnosticsGetDeviceTwinNextConfigByNameInternalServerError{}
}

/*
EdgeDiagnosticsGetDeviceTwinNextConfigByNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeDiagnosticsGetDeviceTwinNextConfigByNameInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge diagnostics get device twin next config by name internal server error response has a 2xx status code
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge diagnostics get device twin next config by name internal server error response has a 3xx status code
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge diagnostics get device twin next config by name internal server error response has a 4xx status code
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge diagnostics get device twin next config by name internal server error response has a 5xx status code
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge diagnostics get device twin next config by name internal server error response a status code equal to that given
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/config/next][%d] edgeDiagnosticsGetDeviceTwinNextConfigByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/config/next][%d] edgeDiagnosticsGetDeviceTwinNextConfigByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeDiagnosticsGetDeviceTwinNextConfigByNameGatewayTimeout creates a EdgeDiagnosticsGetDeviceTwinNextConfigByNameGatewayTimeout with default headers values
func NewEdgeDiagnosticsGetDeviceTwinNextConfigByNameGatewayTimeout() *EdgeDiagnosticsGetDeviceTwinNextConfigByNameGatewayTimeout {
	return &EdgeDiagnosticsGetDeviceTwinNextConfigByNameGatewayTimeout{}
}

/*
EdgeDiagnosticsGetDeviceTwinNextConfigByNameGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeDiagnosticsGetDeviceTwinNextConfigByNameGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge diagnostics get device twin next config by name gateway timeout response has a 2xx status code
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge diagnostics get device twin next config by name gateway timeout response has a 3xx status code
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge diagnostics get device twin next config by name gateway timeout response has a 4xx status code
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge diagnostics get device twin next config by name gateway timeout response has a 5xx status code
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge diagnostics get device twin next config by name gateway timeout response a status code equal to that given
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/config/next][%d] edgeDiagnosticsGetDeviceTwinNextConfigByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/config/next][%d] edgeDiagnosticsGetDeviceTwinNextConfigByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeDiagnosticsGetDeviceTwinNextConfigByNameDefault creates a EdgeDiagnosticsGetDeviceTwinNextConfigByNameDefault with default headers values
func NewEdgeDiagnosticsGetDeviceTwinNextConfigByNameDefault(code int) *EdgeDiagnosticsGetDeviceTwinNextConfigByNameDefault {
	return &EdgeDiagnosticsGetDeviceTwinNextConfigByNameDefault{
		_statusCode: code,
	}
}

/*
EdgeDiagnosticsGetDeviceTwinNextConfigByNameDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeDiagnosticsGetDeviceTwinNextConfigByNameDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the edge diagnostics get device twin next config by name default response
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this edge diagnostics get device twin next config by name default response has a 2xx status code
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge diagnostics get device twin next config by name default response has a 3xx status code
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge diagnostics get device twin next config by name default response has a 4xx status code
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge diagnostics get device twin next config by name default response has a 5xx status code
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge diagnostics get device twin next config by name default response a status code equal to that given
func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameDefault) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/config/next][%d] EdgeDiagnostics_GetDeviceTwinNextConfigByName default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameDefault) String() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/config/next][%d] EdgeDiagnostics_GetDeviceTwinNextConfigByName default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeDiagnosticsGetDeviceTwinNextConfigByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

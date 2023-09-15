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

// EdgeDiagnosticsRegenDeviceConfigReader is a Reader for the EdgeDiagnosticsRegenDeviceConfig structure.
type EdgeDiagnosticsRegenDeviceConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeDiagnosticsRegenDeviceConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeDiagnosticsRegenDeviceConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEdgeDiagnosticsRegenDeviceConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewEdgeDiagnosticsRegenDeviceConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeDiagnosticsRegenDeviceConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeDiagnosticsRegenDeviceConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeDiagnosticsRegenDeviceConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeDiagnosticsRegenDeviceConfigGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeDiagnosticsRegenDeviceConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeDiagnosticsRegenDeviceConfigOK creates a EdgeDiagnosticsRegenDeviceConfigOK with default headers values
func NewEdgeDiagnosticsRegenDeviceConfigOK() *EdgeDiagnosticsRegenDeviceConfigOK {
	return &EdgeDiagnosticsRegenDeviceConfigOK{}
}

/*
EdgeDiagnosticsRegenDeviceConfigOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeDiagnosticsRegenDeviceConfigOK struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge diagnostics regen device config o k response has a 2xx status code
func (o *EdgeDiagnosticsRegenDeviceConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge diagnostics regen device config o k response has a 3xx status code
func (o *EdgeDiagnosticsRegenDeviceConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge diagnostics regen device config o k response has a 4xx status code
func (o *EdgeDiagnosticsRegenDeviceConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge diagnostics regen device config o k response has a 5xx status code
func (o *EdgeDiagnosticsRegenDeviceConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge diagnostics regen device config o k response a status code equal to that given
func (o *EdgeDiagnosticsRegenDeviceConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the edge diagnostics regen device config o k response
func (o *EdgeDiagnosticsRegenDeviceConfigOK) Code() int {
	return 200
}

func (o *EdgeDiagnosticsRegenDeviceConfigOK) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/config][%d] edgeDiagnosticsRegenDeviceConfigOK  %+v", 200, o.Payload)
}

func (o *EdgeDiagnosticsRegenDeviceConfigOK) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/config][%d] edgeDiagnosticsRegenDeviceConfigOK  %+v", 200, o.Payload)
}

func (o *EdgeDiagnosticsRegenDeviceConfigOK) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeDiagnosticsRegenDeviceConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeDiagnosticsRegenDeviceConfigBadRequest creates a EdgeDiagnosticsRegenDeviceConfigBadRequest with default headers values
func NewEdgeDiagnosticsRegenDeviceConfigBadRequest() *EdgeDiagnosticsRegenDeviceConfigBadRequest {
	return &EdgeDiagnosticsRegenDeviceConfigBadRequest{}
}

/*
EdgeDiagnosticsRegenDeviceConfigBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type EdgeDiagnosticsRegenDeviceConfigBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge diagnostics regen device config bad request response has a 2xx status code
func (o *EdgeDiagnosticsRegenDeviceConfigBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge diagnostics regen device config bad request response has a 3xx status code
func (o *EdgeDiagnosticsRegenDeviceConfigBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge diagnostics regen device config bad request response has a 4xx status code
func (o *EdgeDiagnosticsRegenDeviceConfigBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge diagnostics regen device config bad request response has a 5xx status code
func (o *EdgeDiagnosticsRegenDeviceConfigBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this edge diagnostics regen device config bad request response a status code equal to that given
func (o *EdgeDiagnosticsRegenDeviceConfigBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the edge diagnostics regen device config bad request response
func (o *EdgeDiagnosticsRegenDeviceConfigBadRequest) Code() int {
	return 400
}

func (o *EdgeDiagnosticsRegenDeviceConfigBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/config][%d] edgeDiagnosticsRegenDeviceConfigBadRequest  %+v", 400, o.Payload)
}

func (o *EdgeDiagnosticsRegenDeviceConfigBadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/config][%d] edgeDiagnosticsRegenDeviceConfigBadRequest  %+v", 400, o.Payload)
}

func (o *EdgeDiagnosticsRegenDeviceConfigBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeDiagnosticsRegenDeviceConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeDiagnosticsRegenDeviceConfigUnauthorized creates a EdgeDiagnosticsRegenDeviceConfigUnauthorized with default headers values
func NewEdgeDiagnosticsRegenDeviceConfigUnauthorized() *EdgeDiagnosticsRegenDeviceConfigUnauthorized {
	return &EdgeDiagnosticsRegenDeviceConfigUnauthorized{}
}

/*
EdgeDiagnosticsRegenDeviceConfigUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeDiagnosticsRegenDeviceConfigUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge diagnostics regen device config unauthorized response has a 2xx status code
func (o *EdgeDiagnosticsRegenDeviceConfigUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge diagnostics regen device config unauthorized response has a 3xx status code
func (o *EdgeDiagnosticsRegenDeviceConfigUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge diagnostics regen device config unauthorized response has a 4xx status code
func (o *EdgeDiagnosticsRegenDeviceConfigUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge diagnostics regen device config unauthorized response has a 5xx status code
func (o *EdgeDiagnosticsRegenDeviceConfigUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge diagnostics regen device config unauthorized response a status code equal to that given
func (o *EdgeDiagnosticsRegenDeviceConfigUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the edge diagnostics regen device config unauthorized response
func (o *EdgeDiagnosticsRegenDeviceConfigUnauthorized) Code() int {
	return 401
}

func (o *EdgeDiagnosticsRegenDeviceConfigUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/config][%d] edgeDiagnosticsRegenDeviceConfigUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeDiagnosticsRegenDeviceConfigUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/config][%d] edgeDiagnosticsRegenDeviceConfigUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeDiagnosticsRegenDeviceConfigUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeDiagnosticsRegenDeviceConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeDiagnosticsRegenDeviceConfigForbidden creates a EdgeDiagnosticsRegenDeviceConfigForbidden with default headers values
func NewEdgeDiagnosticsRegenDeviceConfigForbidden() *EdgeDiagnosticsRegenDeviceConfigForbidden {
	return &EdgeDiagnosticsRegenDeviceConfigForbidden{}
}

/*
EdgeDiagnosticsRegenDeviceConfigForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeDiagnosticsRegenDeviceConfigForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge diagnostics regen device config forbidden response has a 2xx status code
func (o *EdgeDiagnosticsRegenDeviceConfigForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge diagnostics regen device config forbidden response has a 3xx status code
func (o *EdgeDiagnosticsRegenDeviceConfigForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge diagnostics regen device config forbidden response has a 4xx status code
func (o *EdgeDiagnosticsRegenDeviceConfigForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge diagnostics regen device config forbidden response has a 5xx status code
func (o *EdgeDiagnosticsRegenDeviceConfigForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge diagnostics regen device config forbidden response a status code equal to that given
func (o *EdgeDiagnosticsRegenDeviceConfigForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the edge diagnostics regen device config forbidden response
func (o *EdgeDiagnosticsRegenDeviceConfigForbidden) Code() int {
	return 403
}

func (o *EdgeDiagnosticsRegenDeviceConfigForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/config][%d] edgeDiagnosticsRegenDeviceConfigForbidden  %+v", 403, o.Payload)
}

func (o *EdgeDiagnosticsRegenDeviceConfigForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/config][%d] edgeDiagnosticsRegenDeviceConfigForbidden  %+v", 403, o.Payload)
}

func (o *EdgeDiagnosticsRegenDeviceConfigForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeDiagnosticsRegenDeviceConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeDiagnosticsRegenDeviceConfigNotFound creates a EdgeDiagnosticsRegenDeviceConfigNotFound with default headers values
func NewEdgeDiagnosticsRegenDeviceConfigNotFound() *EdgeDiagnosticsRegenDeviceConfigNotFound {
	return &EdgeDiagnosticsRegenDeviceConfigNotFound{}
}

/*
EdgeDiagnosticsRegenDeviceConfigNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeDiagnosticsRegenDeviceConfigNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge diagnostics regen device config not found response has a 2xx status code
func (o *EdgeDiagnosticsRegenDeviceConfigNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge diagnostics regen device config not found response has a 3xx status code
func (o *EdgeDiagnosticsRegenDeviceConfigNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge diagnostics regen device config not found response has a 4xx status code
func (o *EdgeDiagnosticsRegenDeviceConfigNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge diagnostics regen device config not found response has a 5xx status code
func (o *EdgeDiagnosticsRegenDeviceConfigNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge diagnostics regen device config not found response a status code equal to that given
func (o *EdgeDiagnosticsRegenDeviceConfigNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the edge diagnostics regen device config not found response
func (o *EdgeDiagnosticsRegenDeviceConfigNotFound) Code() int {
	return 404
}

func (o *EdgeDiagnosticsRegenDeviceConfigNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/config][%d] edgeDiagnosticsRegenDeviceConfigNotFound  %+v", 404, o.Payload)
}

func (o *EdgeDiagnosticsRegenDeviceConfigNotFound) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/config][%d] edgeDiagnosticsRegenDeviceConfigNotFound  %+v", 404, o.Payload)
}

func (o *EdgeDiagnosticsRegenDeviceConfigNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeDiagnosticsRegenDeviceConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeDiagnosticsRegenDeviceConfigInternalServerError creates a EdgeDiagnosticsRegenDeviceConfigInternalServerError with default headers values
func NewEdgeDiagnosticsRegenDeviceConfigInternalServerError() *EdgeDiagnosticsRegenDeviceConfigInternalServerError {
	return &EdgeDiagnosticsRegenDeviceConfigInternalServerError{}
}

/*
EdgeDiagnosticsRegenDeviceConfigInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeDiagnosticsRegenDeviceConfigInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge diagnostics regen device config internal server error response has a 2xx status code
func (o *EdgeDiagnosticsRegenDeviceConfigInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge diagnostics regen device config internal server error response has a 3xx status code
func (o *EdgeDiagnosticsRegenDeviceConfigInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge diagnostics regen device config internal server error response has a 4xx status code
func (o *EdgeDiagnosticsRegenDeviceConfigInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge diagnostics regen device config internal server error response has a 5xx status code
func (o *EdgeDiagnosticsRegenDeviceConfigInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge diagnostics regen device config internal server error response a status code equal to that given
func (o *EdgeDiagnosticsRegenDeviceConfigInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the edge diagnostics regen device config internal server error response
func (o *EdgeDiagnosticsRegenDeviceConfigInternalServerError) Code() int {
	return 500
}

func (o *EdgeDiagnosticsRegenDeviceConfigInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/config][%d] edgeDiagnosticsRegenDeviceConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeDiagnosticsRegenDeviceConfigInternalServerError) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/config][%d] edgeDiagnosticsRegenDeviceConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeDiagnosticsRegenDeviceConfigInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeDiagnosticsRegenDeviceConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeDiagnosticsRegenDeviceConfigGatewayTimeout creates a EdgeDiagnosticsRegenDeviceConfigGatewayTimeout with default headers values
func NewEdgeDiagnosticsRegenDeviceConfigGatewayTimeout() *EdgeDiagnosticsRegenDeviceConfigGatewayTimeout {
	return &EdgeDiagnosticsRegenDeviceConfigGatewayTimeout{}
}

/*
EdgeDiagnosticsRegenDeviceConfigGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeDiagnosticsRegenDeviceConfigGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge diagnostics regen device config gateway timeout response has a 2xx status code
func (o *EdgeDiagnosticsRegenDeviceConfigGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge diagnostics regen device config gateway timeout response has a 3xx status code
func (o *EdgeDiagnosticsRegenDeviceConfigGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge diagnostics regen device config gateway timeout response has a 4xx status code
func (o *EdgeDiagnosticsRegenDeviceConfigGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge diagnostics regen device config gateway timeout response has a 5xx status code
func (o *EdgeDiagnosticsRegenDeviceConfigGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge diagnostics regen device config gateway timeout response a status code equal to that given
func (o *EdgeDiagnosticsRegenDeviceConfigGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the edge diagnostics regen device config gateway timeout response
func (o *EdgeDiagnosticsRegenDeviceConfigGatewayTimeout) Code() int {
	return 504
}

func (o *EdgeDiagnosticsRegenDeviceConfigGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/config][%d] edgeDiagnosticsRegenDeviceConfigGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeDiagnosticsRegenDeviceConfigGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/config][%d] edgeDiagnosticsRegenDeviceConfigGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeDiagnosticsRegenDeviceConfigGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeDiagnosticsRegenDeviceConfigGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeDiagnosticsRegenDeviceConfigDefault creates a EdgeDiagnosticsRegenDeviceConfigDefault with default headers values
func NewEdgeDiagnosticsRegenDeviceConfigDefault(code int) *EdgeDiagnosticsRegenDeviceConfigDefault {
	return &EdgeDiagnosticsRegenDeviceConfigDefault{
		_statusCode: code,
	}
}

/*
EdgeDiagnosticsRegenDeviceConfigDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeDiagnosticsRegenDeviceConfigDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// IsSuccess returns true when this edge diagnostics regen device config default response has a 2xx status code
func (o *EdgeDiagnosticsRegenDeviceConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge diagnostics regen device config default response has a 3xx status code
func (o *EdgeDiagnosticsRegenDeviceConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge diagnostics regen device config default response has a 4xx status code
func (o *EdgeDiagnosticsRegenDeviceConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge diagnostics regen device config default response has a 5xx status code
func (o *EdgeDiagnosticsRegenDeviceConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge diagnostics regen device config default response a status code equal to that given
func (o *EdgeDiagnosticsRegenDeviceConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the edge diagnostics regen device config default response
func (o *EdgeDiagnosticsRegenDeviceConfigDefault) Code() int {
	return o._statusCode
}

func (o *EdgeDiagnosticsRegenDeviceConfigDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/config][%d] EdgeDiagnostics_RegenDeviceConfig default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeDiagnosticsRegenDeviceConfigDefault) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/config][%d] EdgeDiagnostics_RegenDeviceConfig default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeDiagnosticsRegenDeviceConfigDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeDiagnosticsRegenDeviceConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

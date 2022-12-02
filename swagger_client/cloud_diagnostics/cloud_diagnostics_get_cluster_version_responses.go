// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package cloud_diagnostics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// CloudDiagnosticsGetClusterVersionReader is a Reader for the CloudDiagnosticsGetClusterVersion structure.
type CloudDiagnosticsGetClusterVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloudDiagnosticsGetClusterVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCloudDiagnosticsGetClusterVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewCloudDiagnosticsGetClusterVersionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewCloudDiagnosticsGetClusterVersionGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCloudDiagnosticsGetClusterVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCloudDiagnosticsGetClusterVersionOK creates a CloudDiagnosticsGetClusterVersionOK with default headers values
func NewCloudDiagnosticsGetClusterVersionOK() *CloudDiagnosticsGetClusterVersionOK {
	return &CloudDiagnosticsGetClusterVersionOK{}
}

/*
CloudDiagnosticsGetClusterVersionOK describes a response with status code 200, with default header values.

A successful response.
*/
type CloudDiagnosticsGetClusterVersionOK struct {
	Payload *swagger_models.CloudVersionResp
}

// IsSuccess returns true when this cloud diagnostics get cluster version o k response has a 2xx status code
func (o *CloudDiagnosticsGetClusterVersionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cloud diagnostics get cluster version o k response has a 3xx status code
func (o *CloudDiagnosticsGetClusterVersionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud diagnostics get cluster version o k response has a 4xx status code
func (o *CloudDiagnosticsGetClusterVersionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cloud diagnostics get cluster version o k response has a 5xx status code
func (o *CloudDiagnosticsGetClusterVersionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud diagnostics get cluster version o k response a status code equal to that given
func (o *CloudDiagnosticsGetClusterVersionOK) IsCode(code int) bool {
	return code == 200
}

func (o *CloudDiagnosticsGetClusterVersionOK) Error() string {
	return fmt.Sprintf("[GET /v1/cloud/version][%d] cloudDiagnosticsGetClusterVersionOK  %+v", 200, o.Payload)
}

func (o *CloudDiagnosticsGetClusterVersionOK) String() string {
	return fmt.Sprintf("[GET /v1/cloud/version][%d] cloudDiagnosticsGetClusterVersionOK  %+v", 200, o.Payload)
}

func (o *CloudDiagnosticsGetClusterVersionOK) GetPayload() *swagger_models.CloudVersionResp {
	return o.Payload
}

func (o *CloudDiagnosticsGetClusterVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.CloudVersionResp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudDiagnosticsGetClusterVersionInternalServerError creates a CloudDiagnosticsGetClusterVersionInternalServerError with default headers values
func NewCloudDiagnosticsGetClusterVersionInternalServerError() *CloudDiagnosticsGetClusterVersionInternalServerError {
	return &CloudDiagnosticsGetClusterVersionInternalServerError{}
}

/*
CloudDiagnosticsGetClusterVersionInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type CloudDiagnosticsGetClusterVersionInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this cloud diagnostics get cluster version internal server error response has a 2xx status code
func (o *CloudDiagnosticsGetClusterVersionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud diagnostics get cluster version internal server error response has a 3xx status code
func (o *CloudDiagnosticsGetClusterVersionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud diagnostics get cluster version internal server error response has a 4xx status code
func (o *CloudDiagnosticsGetClusterVersionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this cloud diagnostics get cluster version internal server error response has a 5xx status code
func (o *CloudDiagnosticsGetClusterVersionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this cloud diagnostics get cluster version internal server error response a status code equal to that given
func (o *CloudDiagnosticsGetClusterVersionInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CloudDiagnosticsGetClusterVersionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/cloud/version][%d] cloudDiagnosticsGetClusterVersionInternalServerError  %+v", 500, o.Payload)
}

func (o *CloudDiagnosticsGetClusterVersionInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/cloud/version][%d] cloudDiagnosticsGetClusterVersionInternalServerError  %+v", 500, o.Payload)
}

func (o *CloudDiagnosticsGetClusterVersionInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CloudDiagnosticsGetClusterVersionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudDiagnosticsGetClusterVersionGatewayTimeout creates a CloudDiagnosticsGetClusterVersionGatewayTimeout with default headers values
func NewCloudDiagnosticsGetClusterVersionGatewayTimeout() *CloudDiagnosticsGetClusterVersionGatewayTimeout {
	return &CloudDiagnosticsGetClusterVersionGatewayTimeout{}
}

/*
CloudDiagnosticsGetClusterVersionGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type CloudDiagnosticsGetClusterVersionGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this cloud diagnostics get cluster version gateway timeout response has a 2xx status code
func (o *CloudDiagnosticsGetClusterVersionGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud diagnostics get cluster version gateway timeout response has a 3xx status code
func (o *CloudDiagnosticsGetClusterVersionGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud diagnostics get cluster version gateway timeout response has a 4xx status code
func (o *CloudDiagnosticsGetClusterVersionGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this cloud diagnostics get cluster version gateway timeout response has a 5xx status code
func (o *CloudDiagnosticsGetClusterVersionGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this cloud diagnostics get cluster version gateway timeout response a status code equal to that given
func (o *CloudDiagnosticsGetClusterVersionGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *CloudDiagnosticsGetClusterVersionGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/cloud/version][%d] cloudDiagnosticsGetClusterVersionGatewayTimeout  %+v", 504, o.Payload)
}

func (o *CloudDiagnosticsGetClusterVersionGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/cloud/version][%d] cloudDiagnosticsGetClusterVersionGatewayTimeout  %+v", 504, o.Payload)
}

func (o *CloudDiagnosticsGetClusterVersionGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CloudDiagnosticsGetClusterVersionGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudDiagnosticsGetClusterVersionDefault creates a CloudDiagnosticsGetClusterVersionDefault with default headers values
func NewCloudDiagnosticsGetClusterVersionDefault(code int) *CloudDiagnosticsGetClusterVersionDefault {
	return &CloudDiagnosticsGetClusterVersionDefault{
		_statusCode: code,
	}
}

/*
CloudDiagnosticsGetClusterVersionDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CloudDiagnosticsGetClusterVersionDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the cloud diagnostics get cluster version default response
func (o *CloudDiagnosticsGetClusterVersionDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this cloud diagnostics get cluster version default response has a 2xx status code
func (o *CloudDiagnosticsGetClusterVersionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cloud diagnostics get cluster version default response has a 3xx status code
func (o *CloudDiagnosticsGetClusterVersionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cloud diagnostics get cluster version default response has a 4xx status code
func (o *CloudDiagnosticsGetClusterVersionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cloud diagnostics get cluster version default response has a 5xx status code
func (o *CloudDiagnosticsGetClusterVersionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cloud diagnostics get cluster version default response a status code equal to that given
func (o *CloudDiagnosticsGetClusterVersionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CloudDiagnosticsGetClusterVersionDefault) Error() string {
	return fmt.Sprintf("[GET /v1/cloud/version][%d] CloudDiagnostics_GetClusterVersion default  %+v", o._statusCode, o.Payload)
}

func (o *CloudDiagnosticsGetClusterVersionDefault) String() string {
	return fmt.Sprintf("[GET /v1/cloud/version][%d] CloudDiagnostics_GetClusterVersion default  %+v", o._statusCode, o.Payload)
}

func (o *CloudDiagnosticsGetClusterVersionDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *CloudDiagnosticsGetClusterVersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

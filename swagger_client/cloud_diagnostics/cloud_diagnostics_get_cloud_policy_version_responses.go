// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

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

// CloudDiagnosticsGetCloudPolicyVersionReader is a Reader for the CloudDiagnosticsGetCloudPolicyVersion structure.
type CloudDiagnosticsGetCloudPolicyVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloudDiagnosticsGetCloudPolicyVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCloudDiagnosticsGetCloudPolicyVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewCloudDiagnosticsGetCloudPolicyVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCloudDiagnosticsGetCloudPolicyVersionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewCloudDiagnosticsGetCloudPolicyVersionGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCloudDiagnosticsGetCloudPolicyVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCloudDiagnosticsGetCloudPolicyVersionOK creates a CloudDiagnosticsGetCloudPolicyVersionOK with default headers values
func NewCloudDiagnosticsGetCloudPolicyVersionOK() *CloudDiagnosticsGetCloudPolicyVersionOK {
	return &CloudDiagnosticsGetCloudPolicyVersionOK{}
}

/* CloudDiagnosticsGetCloudPolicyVersionOK describes a response with status code 200, with default header values.

A successful response.
*/
type CloudDiagnosticsGetCloudPolicyVersionOK struct {
	Payload *swagger_models.PolicyDocVersionResp
}

func (o *CloudDiagnosticsGetCloudPolicyVersionOK) Error() string {
	return fmt.Sprintf("[GET /v1/cloud/policies/name/{policy}/{version}][%d] cloudDiagnosticsGetCloudPolicyVersionOK  %+v", 200, o.Payload)
}
func (o *CloudDiagnosticsGetCloudPolicyVersionOK) GetPayload() *swagger_models.PolicyDocVersionResp {
	return o.Payload
}

func (o *CloudDiagnosticsGetCloudPolicyVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.PolicyDocVersionResp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudDiagnosticsGetCloudPolicyVersionNotFound creates a CloudDiagnosticsGetCloudPolicyVersionNotFound with default headers values
func NewCloudDiagnosticsGetCloudPolicyVersionNotFound() *CloudDiagnosticsGetCloudPolicyVersionNotFound {
	return &CloudDiagnosticsGetCloudPolicyVersionNotFound{}
}

/* CloudDiagnosticsGetCloudPolicyVersionNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type CloudDiagnosticsGetCloudPolicyVersionNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CloudDiagnosticsGetCloudPolicyVersionNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/cloud/policies/name/{policy}/{version}][%d] cloudDiagnosticsGetCloudPolicyVersionNotFound  %+v", 404, o.Payload)
}
func (o *CloudDiagnosticsGetCloudPolicyVersionNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CloudDiagnosticsGetCloudPolicyVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudDiagnosticsGetCloudPolicyVersionInternalServerError creates a CloudDiagnosticsGetCloudPolicyVersionInternalServerError with default headers values
func NewCloudDiagnosticsGetCloudPolicyVersionInternalServerError() *CloudDiagnosticsGetCloudPolicyVersionInternalServerError {
	return &CloudDiagnosticsGetCloudPolicyVersionInternalServerError{}
}

/* CloudDiagnosticsGetCloudPolicyVersionInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type CloudDiagnosticsGetCloudPolicyVersionInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CloudDiagnosticsGetCloudPolicyVersionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/cloud/policies/name/{policy}/{version}][%d] cloudDiagnosticsGetCloudPolicyVersionInternalServerError  %+v", 500, o.Payload)
}
func (o *CloudDiagnosticsGetCloudPolicyVersionInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CloudDiagnosticsGetCloudPolicyVersionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudDiagnosticsGetCloudPolicyVersionGatewayTimeout creates a CloudDiagnosticsGetCloudPolicyVersionGatewayTimeout with default headers values
func NewCloudDiagnosticsGetCloudPolicyVersionGatewayTimeout() *CloudDiagnosticsGetCloudPolicyVersionGatewayTimeout {
	return &CloudDiagnosticsGetCloudPolicyVersionGatewayTimeout{}
}

/* CloudDiagnosticsGetCloudPolicyVersionGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type CloudDiagnosticsGetCloudPolicyVersionGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CloudDiagnosticsGetCloudPolicyVersionGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/cloud/policies/name/{policy}/{version}][%d] cloudDiagnosticsGetCloudPolicyVersionGatewayTimeout  %+v", 504, o.Payload)
}
func (o *CloudDiagnosticsGetCloudPolicyVersionGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CloudDiagnosticsGetCloudPolicyVersionGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudDiagnosticsGetCloudPolicyVersionDefault creates a CloudDiagnosticsGetCloudPolicyVersionDefault with default headers values
func NewCloudDiagnosticsGetCloudPolicyVersionDefault(code int) *CloudDiagnosticsGetCloudPolicyVersionDefault {
	return &CloudDiagnosticsGetCloudPolicyVersionDefault{
		_statusCode: code,
	}
}

/* CloudDiagnosticsGetCloudPolicyVersionDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CloudDiagnosticsGetCloudPolicyVersionDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the cloud diagnostics get cloud policy version default response
func (o *CloudDiagnosticsGetCloudPolicyVersionDefault) Code() int {
	return o._statusCode
}

func (o *CloudDiagnosticsGetCloudPolicyVersionDefault) Error() string {
	return fmt.Sprintf("[GET /v1/cloud/policies/name/{policy}/{version}][%d] CloudDiagnostics_GetCloudPolicyVersion default  %+v", o._statusCode, o.Payload)
}
func (o *CloudDiagnosticsGetCloudPolicyVersionDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *CloudDiagnosticsGetCloudPolicyVersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
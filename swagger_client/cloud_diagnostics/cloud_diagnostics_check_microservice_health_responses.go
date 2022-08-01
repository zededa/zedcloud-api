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

// CloudDiagnosticsCheckMicroserviceHealthReader is a Reader for the CloudDiagnosticsCheckMicroserviceHealth structure.
type CloudDiagnosticsCheckMicroserviceHealthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloudDiagnosticsCheckMicroserviceHealthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCloudDiagnosticsCheckMicroserviceHealthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCloudDiagnosticsCheckMicroserviceHealthBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCloudDiagnosticsCheckMicroserviceHealthInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewCloudDiagnosticsCheckMicroserviceHealthGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCloudDiagnosticsCheckMicroserviceHealthDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCloudDiagnosticsCheckMicroserviceHealthOK creates a CloudDiagnosticsCheckMicroserviceHealthOK with default headers values
func NewCloudDiagnosticsCheckMicroserviceHealthOK() *CloudDiagnosticsCheckMicroserviceHealthOK {
	return &CloudDiagnosticsCheckMicroserviceHealthOK{}
}

/* CloudDiagnosticsCheckMicroserviceHealthOK describes a response with status code 200, with default header values.

A successful response.
*/
type CloudDiagnosticsCheckMicroserviceHealthOK struct {
	Payload *swagger_models.HelloResp
}

func (o *CloudDiagnosticsCheckMicroserviceHealthOK) Error() string {
	return fmt.Sprintf("[POST /v1/hello][%d] cloudDiagnosticsCheckMicroserviceHealthOK  %+v", 200, o.Payload)
}
func (o *CloudDiagnosticsCheckMicroserviceHealthOK) GetPayload() *swagger_models.HelloResp {
	return o.Payload
}

func (o *CloudDiagnosticsCheckMicroserviceHealthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.HelloResp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudDiagnosticsCheckMicroserviceHealthBadRequest creates a CloudDiagnosticsCheckMicroserviceHealthBadRequest with default headers values
func NewCloudDiagnosticsCheckMicroserviceHealthBadRequest() *CloudDiagnosticsCheckMicroserviceHealthBadRequest {
	return &CloudDiagnosticsCheckMicroserviceHealthBadRequest{}
}

/* CloudDiagnosticsCheckMicroserviceHealthBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type CloudDiagnosticsCheckMicroserviceHealthBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CloudDiagnosticsCheckMicroserviceHealthBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/hello][%d] cloudDiagnosticsCheckMicroserviceHealthBadRequest  %+v", 400, o.Payload)
}
func (o *CloudDiagnosticsCheckMicroserviceHealthBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CloudDiagnosticsCheckMicroserviceHealthBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudDiagnosticsCheckMicroserviceHealthInternalServerError creates a CloudDiagnosticsCheckMicroserviceHealthInternalServerError with default headers values
func NewCloudDiagnosticsCheckMicroserviceHealthInternalServerError() *CloudDiagnosticsCheckMicroserviceHealthInternalServerError {
	return &CloudDiagnosticsCheckMicroserviceHealthInternalServerError{}
}

/* CloudDiagnosticsCheckMicroserviceHealthInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type CloudDiagnosticsCheckMicroserviceHealthInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CloudDiagnosticsCheckMicroserviceHealthInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/hello][%d] cloudDiagnosticsCheckMicroserviceHealthInternalServerError  %+v", 500, o.Payload)
}
func (o *CloudDiagnosticsCheckMicroserviceHealthInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CloudDiagnosticsCheckMicroserviceHealthInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudDiagnosticsCheckMicroserviceHealthGatewayTimeout creates a CloudDiagnosticsCheckMicroserviceHealthGatewayTimeout with default headers values
func NewCloudDiagnosticsCheckMicroserviceHealthGatewayTimeout() *CloudDiagnosticsCheckMicroserviceHealthGatewayTimeout {
	return &CloudDiagnosticsCheckMicroserviceHealthGatewayTimeout{}
}

/* CloudDiagnosticsCheckMicroserviceHealthGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type CloudDiagnosticsCheckMicroserviceHealthGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CloudDiagnosticsCheckMicroserviceHealthGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v1/hello][%d] cloudDiagnosticsCheckMicroserviceHealthGatewayTimeout  %+v", 504, o.Payload)
}
func (o *CloudDiagnosticsCheckMicroserviceHealthGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CloudDiagnosticsCheckMicroserviceHealthGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudDiagnosticsCheckMicroserviceHealthDefault creates a CloudDiagnosticsCheckMicroserviceHealthDefault with default headers values
func NewCloudDiagnosticsCheckMicroserviceHealthDefault(code int) *CloudDiagnosticsCheckMicroserviceHealthDefault {
	return &CloudDiagnosticsCheckMicroserviceHealthDefault{
		_statusCode: code,
	}
}

/* CloudDiagnosticsCheckMicroserviceHealthDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CloudDiagnosticsCheckMicroserviceHealthDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the cloud diagnostics check microservice health default response
func (o *CloudDiagnosticsCheckMicroserviceHealthDefault) Code() int {
	return o._statusCode
}

func (o *CloudDiagnosticsCheckMicroserviceHealthDefault) Error() string {
	return fmt.Sprintf("[POST /v1/hello][%d] CloudDiagnostics_checkMicroserviceHealth default  %+v", o._statusCode, o.Payload)
}
func (o *CloudDiagnosticsCheckMicroserviceHealthDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *CloudDiagnosticsCheckMicroserviceHealthDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

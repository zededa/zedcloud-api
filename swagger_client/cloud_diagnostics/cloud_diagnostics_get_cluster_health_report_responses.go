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

// CloudDiagnosticsGetClusterHealthReportReader is a Reader for the CloudDiagnosticsGetClusterHealthReport structure.
type CloudDiagnosticsGetClusterHealthReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloudDiagnosticsGetClusterHealthReportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCloudDiagnosticsGetClusterHealthReportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewCloudDiagnosticsGetClusterHealthReportInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewCloudDiagnosticsGetClusterHealthReportGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCloudDiagnosticsGetClusterHealthReportDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCloudDiagnosticsGetClusterHealthReportOK creates a CloudDiagnosticsGetClusterHealthReportOK with default headers values
func NewCloudDiagnosticsGetClusterHealthReportOK() *CloudDiagnosticsGetClusterHealthReportOK {
	return &CloudDiagnosticsGetClusterHealthReportOK{}
}

/*
CloudDiagnosticsGetClusterHealthReportOK describes a response with status code 200, with default header values.

A successful response.
*/
type CloudDiagnosticsGetClusterHealthReportOK struct {
	Payload *swagger_models.HealthServiceResp
}

// IsSuccess returns true when this cloud diagnostics get cluster health report o k response has a 2xx status code
func (o *CloudDiagnosticsGetClusterHealthReportOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cloud diagnostics get cluster health report o k response has a 3xx status code
func (o *CloudDiagnosticsGetClusterHealthReportOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud diagnostics get cluster health report o k response has a 4xx status code
func (o *CloudDiagnosticsGetClusterHealthReportOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cloud diagnostics get cluster health report o k response has a 5xx status code
func (o *CloudDiagnosticsGetClusterHealthReportOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud diagnostics get cluster health report o k response a status code equal to that given
func (o *CloudDiagnosticsGetClusterHealthReportOK) IsCode(code int) bool {
	return code == 200
}

func (o *CloudDiagnosticsGetClusterHealthReportOK) Error() string {
	return fmt.Sprintf("[GET /v1/cloud/healthreport][%d] cloudDiagnosticsGetClusterHealthReportOK  %+v", 200, o.Payload)
}

func (o *CloudDiagnosticsGetClusterHealthReportOK) String() string {
	return fmt.Sprintf("[GET /v1/cloud/healthreport][%d] cloudDiagnosticsGetClusterHealthReportOK  %+v", 200, o.Payload)
}

func (o *CloudDiagnosticsGetClusterHealthReportOK) GetPayload() *swagger_models.HealthServiceResp {
	return o.Payload
}

func (o *CloudDiagnosticsGetClusterHealthReportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.HealthServiceResp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudDiagnosticsGetClusterHealthReportInternalServerError creates a CloudDiagnosticsGetClusterHealthReportInternalServerError with default headers values
func NewCloudDiagnosticsGetClusterHealthReportInternalServerError() *CloudDiagnosticsGetClusterHealthReportInternalServerError {
	return &CloudDiagnosticsGetClusterHealthReportInternalServerError{}
}

/*
CloudDiagnosticsGetClusterHealthReportInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type CloudDiagnosticsGetClusterHealthReportInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this cloud diagnostics get cluster health report internal server error response has a 2xx status code
func (o *CloudDiagnosticsGetClusterHealthReportInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud diagnostics get cluster health report internal server error response has a 3xx status code
func (o *CloudDiagnosticsGetClusterHealthReportInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud diagnostics get cluster health report internal server error response has a 4xx status code
func (o *CloudDiagnosticsGetClusterHealthReportInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this cloud diagnostics get cluster health report internal server error response has a 5xx status code
func (o *CloudDiagnosticsGetClusterHealthReportInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this cloud diagnostics get cluster health report internal server error response a status code equal to that given
func (o *CloudDiagnosticsGetClusterHealthReportInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CloudDiagnosticsGetClusterHealthReportInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/cloud/healthreport][%d] cloudDiagnosticsGetClusterHealthReportInternalServerError  %+v", 500, o.Payload)
}

func (o *CloudDiagnosticsGetClusterHealthReportInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/cloud/healthreport][%d] cloudDiagnosticsGetClusterHealthReportInternalServerError  %+v", 500, o.Payload)
}

func (o *CloudDiagnosticsGetClusterHealthReportInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CloudDiagnosticsGetClusterHealthReportInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudDiagnosticsGetClusterHealthReportGatewayTimeout creates a CloudDiagnosticsGetClusterHealthReportGatewayTimeout with default headers values
func NewCloudDiagnosticsGetClusterHealthReportGatewayTimeout() *CloudDiagnosticsGetClusterHealthReportGatewayTimeout {
	return &CloudDiagnosticsGetClusterHealthReportGatewayTimeout{}
}

/*
CloudDiagnosticsGetClusterHealthReportGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type CloudDiagnosticsGetClusterHealthReportGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this cloud diagnostics get cluster health report gateway timeout response has a 2xx status code
func (o *CloudDiagnosticsGetClusterHealthReportGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud diagnostics get cluster health report gateway timeout response has a 3xx status code
func (o *CloudDiagnosticsGetClusterHealthReportGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud diagnostics get cluster health report gateway timeout response has a 4xx status code
func (o *CloudDiagnosticsGetClusterHealthReportGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this cloud diagnostics get cluster health report gateway timeout response has a 5xx status code
func (o *CloudDiagnosticsGetClusterHealthReportGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this cloud diagnostics get cluster health report gateway timeout response a status code equal to that given
func (o *CloudDiagnosticsGetClusterHealthReportGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *CloudDiagnosticsGetClusterHealthReportGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/cloud/healthreport][%d] cloudDiagnosticsGetClusterHealthReportGatewayTimeout  %+v", 504, o.Payload)
}

func (o *CloudDiagnosticsGetClusterHealthReportGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/cloud/healthreport][%d] cloudDiagnosticsGetClusterHealthReportGatewayTimeout  %+v", 504, o.Payload)
}

func (o *CloudDiagnosticsGetClusterHealthReportGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CloudDiagnosticsGetClusterHealthReportGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudDiagnosticsGetClusterHealthReportDefault creates a CloudDiagnosticsGetClusterHealthReportDefault with default headers values
func NewCloudDiagnosticsGetClusterHealthReportDefault(code int) *CloudDiagnosticsGetClusterHealthReportDefault {
	return &CloudDiagnosticsGetClusterHealthReportDefault{
		_statusCode: code,
	}
}

/*
CloudDiagnosticsGetClusterHealthReportDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CloudDiagnosticsGetClusterHealthReportDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the cloud diagnostics get cluster health report default response
func (o *CloudDiagnosticsGetClusterHealthReportDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this cloud diagnostics get cluster health report default response has a 2xx status code
func (o *CloudDiagnosticsGetClusterHealthReportDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cloud diagnostics get cluster health report default response has a 3xx status code
func (o *CloudDiagnosticsGetClusterHealthReportDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cloud diagnostics get cluster health report default response has a 4xx status code
func (o *CloudDiagnosticsGetClusterHealthReportDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cloud diagnostics get cluster health report default response has a 5xx status code
func (o *CloudDiagnosticsGetClusterHealthReportDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cloud diagnostics get cluster health report default response a status code equal to that given
func (o *CloudDiagnosticsGetClusterHealthReportDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CloudDiagnosticsGetClusterHealthReportDefault) Error() string {
	return fmt.Sprintf("[GET /v1/cloud/healthreport][%d] CloudDiagnostics_GetClusterHealthReport default  %+v", o._statusCode, o.Payload)
}

func (o *CloudDiagnosticsGetClusterHealthReportDefault) String() string {
	return fmt.Sprintf("[GET /v1/cloud/healthreport][%d] CloudDiagnostics_GetClusterHealthReport default  %+v", o._statusCode, o.Payload)
}

func (o *CloudDiagnosticsGetClusterHealthReportDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *CloudDiagnosticsGetClusterHealthReportDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// CloudDiagnosticsGetCloudPolicyDocumentReader is a Reader for the CloudDiagnosticsGetCloudPolicyDocument structure.
type CloudDiagnosticsGetCloudPolicyDocumentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloudDiagnosticsGetCloudPolicyDocumentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCloudDiagnosticsGetCloudPolicyDocumentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewCloudDiagnosticsGetCloudPolicyDocumentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCloudDiagnosticsGetCloudPolicyDocumentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewCloudDiagnosticsGetCloudPolicyDocumentGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCloudDiagnosticsGetCloudPolicyDocumentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCloudDiagnosticsGetCloudPolicyDocumentOK creates a CloudDiagnosticsGetCloudPolicyDocumentOK with default headers values
func NewCloudDiagnosticsGetCloudPolicyDocumentOK() *CloudDiagnosticsGetCloudPolicyDocumentOK {
	return &CloudDiagnosticsGetCloudPolicyDocumentOK{}
}

/*
CloudDiagnosticsGetCloudPolicyDocumentOK describes a response with status code 200, with default header values.

A successful response.
*/
type CloudDiagnosticsGetCloudPolicyDocumentOK struct {
	Payload *swagger_models.PolicyDocVersionResp
}

// IsSuccess returns true when this cloud diagnostics get cloud policy document o k response has a 2xx status code
func (o *CloudDiagnosticsGetCloudPolicyDocumentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cloud diagnostics get cloud policy document o k response has a 3xx status code
func (o *CloudDiagnosticsGetCloudPolicyDocumentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud diagnostics get cloud policy document o k response has a 4xx status code
func (o *CloudDiagnosticsGetCloudPolicyDocumentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cloud diagnostics get cloud policy document o k response has a 5xx status code
func (o *CloudDiagnosticsGetCloudPolicyDocumentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud diagnostics get cloud policy document o k response a status code equal to that given
func (o *CloudDiagnosticsGetCloudPolicyDocumentOK) IsCode(code int) bool {
	return code == 200
}

func (o *CloudDiagnosticsGetCloudPolicyDocumentOK) Error() string {
	return fmt.Sprintf("[GET /v1/cloud/policies/id/{fileURL}][%d] cloudDiagnosticsGetCloudPolicyDocumentOK  %+v", 200, o.Payload)
}

func (o *CloudDiagnosticsGetCloudPolicyDocumentOK) String() string {
	return fmt.Sprintf("[GET /v1/cloud/policies/id/{fileURL}][%d] cloudDiagnosticsGetCloudPolicyDocumentOK  %+v", 200, o.Payload)
}

func (o *CloudDiagnosticsGetCloudPolicyDocumentOK) GetPayload() *swagger_models.PolicyDocVersionResp {
	return o.Payload
}

func (o *CloudDiagnosticsGetCloudPolicyDocumentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.PolicyDocVersionResp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudDiagnosticsGetCloudPolicyDocumentNotFound creates a CloudDiagnosticsGetCloudPolicyDocumentNotFound with default headers values
func NewCloudDiagnosticsGetCloudPolicyDocumentNotFound() *CloudDiagnosticsGetCloudPolicyDocumentNotFound {
	return &CloudDiagnosticsGetCloudPolicyDocumentNotFound{}
}

/*
CloudDiagnosticsGetCloudPolicyDocumentNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type CloudDiagnosticsGetCloudPolicyDocumentNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this cloud diagnostics get cloud policy document not found response has a 2xx status code
func (o *CloudDiagnosticsGetCloudPolicyDocumentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud diagnostics get cloud policy document not found response has a 3xx status code
func (o *CloudDiagnosticsGetCloudPolicyDocumentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud diagnostics get cloud policy document not found response has a 4xx status code
func (o *CloudDiagnosticsGetCloudPolicyDocumentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this cloud diagnostics get cloud policy document not found response has a 5xx status code
func (o *CloudDiagnosticsGetCloudPolicyDocumentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud diagnostics get cloud policy document not found response a status code equal to that given
func (o *CloudDiagnosticsGetCloudPolicyDocumentNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CloudDiagnosticsGetCloudPolicyDocumentNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/cloud/policies/id/{fileURL}][%d] cloudDiagnosticsGetCloudPolicyDocumentNotFound  %+v", 404, o.Payload)
}

func (o *CloudDiagnosticsGetCloudPolicyDocumentNotFound) String() string {
	return fmt.Sprintf("[GET /v1/cloud/policies/id/{fileURL}][%d] cloudDiagnosticsGetCloudPolicyDocumentNotFound  %+v", 404, o.Payload)
}

func (o *CloudDiagnosticsGetCloudPolicyDocumentNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CloudDiagnosticsGetCloudPolicyDocumentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudDiagnosticsGetCloudPolicyDocumentInternalServerError creates a CloudDiagnosticsGetCloudPolicyDocumentInternalServerError with default headers values
func NewCloudDiagnosticsGetCloudPolicyDocumentInternalServerError() *CloudDiagnosticsGetCloudPolicyDocumentInternalServerError {
	return &CloudDiagnosticsGetCloudPolicyDocumentInternalServerError{}
}

/*
CloudDiagnosticsGetCloudPolicyDocumentInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type CloudDiagnosticsGetCloudPolicyDocumentInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this cloud diagnostics get cloud policy document internal server error response has a 2xx status code
func (o *CloudDiagnosticsGetCloudPolicyDocumentInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud diagnostics get cloud policy document internal server error response has a 3xx status code
func (o *CloudDiagnosticsGetCloudPolicyDocumentInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud diagnostics get cloud policy document internal server error response has a 4xx status code
func (o *CloudDiagnosticsGetCloudPolicyDocumentInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this cloud diagnostics get cloud policy document internal server error response has a 5xx status code
func (o *CloudDiagnosticsGetCloudPolicyDocumentInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this cloud diagnostics get cloud policy document internal server error response a status code equal to that given
func (o *CloudDiagnosticsGetCloudPolicyDocumentInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CloudDiagnosticsGetCloudPolicyDocumentInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/cloud/policies/id/{fileURL}][%d] cloudDiagnosticsGetCloudPolicyDocumentInternalServerError  %+v", 500, o.Payload)
}

func (o *CloudDiagnosticsGetCloudPolicyDocumentInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/cloud/policies/id/{fileURL}][%d] cloudDiagnosticsGetCloudPolicyDocumentInternalServerError  %+v", 500, o.Payload)
}

func (o *CloudDiagnosticsGetCloudPolicyDocumentInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CloudDiagnosticsGetCloudPolicyDocumentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudDiagnosticsGetCloudPolicyDocumentGatewayTimeout creates a CloudDiagnosticsGetCloudPolicyDocumentGatewayTimeout with default headers values
func NewCloudDiagnosticsGetCloudPolicyDocumentGatewayTimeout() *CloudDiagnosticsGetCloudPolicyDocumentGatewayTimeout {
	return &CloudDiagnosticsGetCloudPolicyDocumentGatewayTimeout{}
}

/*
CloudDiagnosticsGetCloudPolicyDocumentGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type CloudDiagnosticsGetCloudPolicyDocumentGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this cloud diagnostics get cloud policy document gateway timeout response has a 2xx status code
func (o *CloudDiagnosticsGetCloudPolicyDocumentGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud diagnostics get cloud policy document gateway timeout response has a 3xx status code
func (o *CloudDiagnosticsGetCloudPolicyDocumentGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud diagnostics get cloud policy document gateway timeout response has a 4xx status code
func (o *CloudDiagnosticsGetCloudPolicyDocumentGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this cloud diagnostics get cloud policy document gateway timeout response has a 5xx status code
func (o *CloudDiagnosticsGetCloudPolicyDocumentGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this cloud diagnostics get cloud policy document gateway timeout response a status code equal to that given
func (o *CloudDiagnosticsGetCloudPolicyDocumentGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *CloudDiagnosticsGetCloudPolicyDocumentGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/cloud/policies/id/{fileURL}][%d] cloudDiagnosticsGetCloudPolicyDocumentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *CloudDiagnosticsGetCloudPolicyDocumentGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/cloud/policies/id/{fileURL}][%d] cloudDiagnosticsGetCloudPolicyDocumentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *CloudDiagnosticsGetCloudPolicyDocumentGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CloudDiagnosticsGetCloudPolicyDocumentGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudDiagnosticsGetCloudPolicyDocumentDefault creates a CloudDiagnosticsGetCloudPolicyDocumentDefault with default headers values
func NewCloudDiagnosticsGetCloudPolicyDocumentDefault(code int) *CloudDiagnosticsGetCloudPolicyDocumentDefault {
	return &CloudDiagnosticsGetCloudPolicyDocumentDefault{
		_statusCode: code,
	}
}

/*
CloudDiagnosticsGetCloudPolicyDocumentDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CloudDiagnosticsGetCloudPolicyDocumentDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the cloud diagnostics get cloud policy document default response
func (o *CloudDiagnosticsGetCloudPolicyDocumentDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this cloud diagnostics get cloud policy document default response has a 2xx status code
func (o *CloudDiagnosticsGetCloudPolicyDocumentDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cloud diagnostics get cloud policy document default response has a 3xx status code
func (o *CloudDiagnosticsGetCloudPolicyDocumentDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cloud diagnostics get cloud policy document default response has a 4xx status code
func (o *CloudDiagnosticsGetCloudPolicyDocumentDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cloud diagnostics get cloud policy document default response has a 5xx status code
func (o *CloudDiagnosticsGetCloudPolicyDocumentDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cloud diagnostics get cloud policy document default response a status code equal to that given
func (o *CloudDiagnosticsGetCloudPolicyDocumentDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CloudDiagnosticsGetCloudPolicyDocumentDefault) Error() string {
	return fmt.Sprintf("[GET /v1/cloud/policies/id/{fileURL}][%d] CloudDiagnostics_GetCloudPolicyDocument default  %+v", o._statusCode, o.Payload)
}

func (o *CloudDiagnosticsGetCloudPolicyDocumentDefault) String() string {
	return fmt.Sprintf("[GET /v1/cloud/policies/id/{fileURL}][%d] CloudDiagnostics_GetCloudPolicyDocument default  %+v", o._statusCode, o.Payload)
}

func (o *CloudDiagnosticsGetCloudPolicyDocumentDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *CloudDiagnosticsGetCloudPolicyDocumentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

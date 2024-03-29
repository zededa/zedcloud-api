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

// CloudDiagnosticsCheckClusterHealth2Reader is a Reader for the CloudDiagnosticsCheckClusterHealth2 structure.
type CloudDiagnosticsCheckClusterHealth2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloudDiagnosticsCheckClusterHealth2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCloudDiagnosticsCheckClusterHealth2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewCloudDiagnosticsCheckClusterHealth2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewCloudDiagnosticsCheckClusterHealth2GatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCloudDiagnosticsCheckClusterHealth2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCloudDiagnosticsCheckClusterHealth2OK creates a CloudDiagnosticsCheckClusterHealth2OK with default headers values
func NewCloudDiagnosticsCheckClusterHealth2OK() *CloudDiagnosticsCheckClusterHealth2OK {
	return &CloudDiagnosticsCheckClusterHealth2OK{}
}

/*
CloudDiagnosticsCheckClusterHealth2OK describes a response with status code 200, with default header values.

A successful response.
*/
type CloudDiagnosticsCheckClusterHealth2OK struct {
	Payload *swagger_models.PingMsgSendResp
}

// IsSuccess returns true when this cloud diagnostics check cluster health2 o k response has a 2xx status code
func (o *CloudDiagnosticsCheckClusterHealth2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cloud diagnostics check cluster health2 o k response has a 3xx status code
func (o *CloudDiagnosticsCheckClusterHealth2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud diagnostics check cluster health2 o k response has a 4xx status code
func (o *CloudDiagnosticsCheckClusterHealth2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cloud diagnostics check cluster health2 o k response has a 5xx status code
func (o *CloudDiagnosticsCheckClusterHealth2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud diagnostics check cluster health2 o k response a status code equal to that given
func (o *CloudDiagnosticsCheckClusterHealth2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cloud diagnostics check cluster health2 o k response
func (o *CloudDiagnosticsCheckClusterHealth2OK) Code() int {
	return 200
}

func (o *CloudDiagnosticsCheckClusterHealth2OK) Error() string {
	return fmt.Sprintf("[GET /v1/cloud/ping/id/{pingId}][%d] cloudDiagnosticsCheckClusterHealth2OK  %+v", 200, o.Payload)
}

func (o *CloudDiagnosticsCheckClusterHealth2OK) String() string {
	return fmt.Sprintf("[GET /v1/cloud/ping/id/{pingId}][%d] cloudDiagnosticsCheckClusterHealth2OK  %+v", 200, o.Payload)
}

func (o *CloudDiagnosticsCheckClusterHealth2OK) GetPayload() *swagger_models.PingMsgSendResp {
	return o.Payload
}

func (o *CloudDiagnosticsCheckClusterHealth2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.PingMsgSendResp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudDiagnosticsCheckClusterHealth2InternalServerError creates a CloudDiagnosticsCheckClusterHealth2InternalServerError with default headers values
func NewCloudDiagnosticsCheckClusterHealth2InternalServerError() *CloudDiagnosticsCheckClusterHealth2InternalServerError {
	return &CloudDiagnosticsCheckClusterHealth2InternalServerError{}
}

/*
CloudDiagnosticsCheckClusterHealth2InternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type CloudDiagnosticsCheckClusterHealth2InternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this cloud diagnostics check cluster health2 internal server error response has a 2xx status code
func (o *CloudDiagnosticsCheckClusterHealth2InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud diagnostics check cluster health2 internal server error response has a 3xx status code
func (o *CloudDiagnosticsCheckClusterHealth2InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud diagnostics check cluster health2 internal server error response has a 4xx status code
func (o *CloudDiagnosticsCheckClusterHealth2InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this cloud diagnostics check cluster health2 internal server error response has a 5xx status code
func (o *CloudDiagnosticsCheckClusterHealth2InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this cloud diagnostics check cluster health2 internal server error response a status code equal to that given
func (o *CloudDiagnosticsCheckClusterHealth2InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the cloud diagnostics check cluster health2 internal server error response
func (o *CloudDiagnosticsCheckClusterHealth2InternalServerError) Code() int {
	return 500
}

func (o *CloudDiagnosticsCheckClusterHealth2InternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/cloud/ping/id/{pingId}][%d] cloudDiagnosticsCheckClusterHealth2InternalServerError  %+v", 500, o.Payload)
}

func (o *CloudDiagnosticsCheckClusterHealth2InternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/cloud/ping/id/{pingId}][%d] cloudDiagnosticsCheckClusterHealth2InternalServerError  %+v", 500, o.Payload)
}

func (o *CloudDiagnosticsCheckClusterHealth2InternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CloudDiagnosticsCheckClusterHealth2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudDiagnosticsCheckClusterHealth2GatewayTimeout creates a CloudDiagnosticsCheckClusterHealth2GatewayTimeout with default headers values
func NewCloudDiagnosticsCheckClusterHealth2GatewayTimeout() *CloudDiagnosticsCheckClusterHealth2GatewayTimeout {
	return &CloudDiagnosticsCheckClusterHealth2GatewayTimeout{}
}

/*
CloudDiagnosticsCheckClusterHealth2GatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type CloudDiagnosticsCheckClusterHealth2GatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this cloud diagnostics check cluster health2 gateway timeout response has a 2xx status code
func (o *CloudDiagnosticsCheckClusterHealth2GatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud diagnostics check cluster health2 gateway timeout response has a 3xx status code
func (o *CloudDiagnosticsCheckClusterHealth2GatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud diagnostics check cluster health2 gateway timeout response has a 4xx status code
func (o *CloudDiagnosticsCheckClusterHealth2GatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this cloud diagnostics check cluster health2 gateway timeout response has a 5xx status code
func (o *CloudDiagnosticsCheckClusterHealth2GatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this cloud diagnostics check cluster health2 gateway timeout response a status code equal to that given
func (o *CloudDiagnosticsCheckClusterHealth2GatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the cloud diagnostics check cluster health2 gateway timeout response
func (o *CloudDiagnosticsCheckClusterHealth2GatewayTimeout) Code() int {
	return 504
}

func (o *CloudDiagnosticsCheckClusterHealth2GatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/cloud/ping/id/{pingId}][%d] cloudDiagnosticsCheckClusterHealth2GatewayTimeout  %+v", 504, o.Payload)
}

func (o *CloudDiagnosticsCheckClusterHealth2GatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/cloud/ping/id/{pingId}][%d] cloudDiagnosticsCheckClusterHealth2GatewayTimeout  %+v", 504, o.Payload)
}

func (o *CloudDiagnosticsCheckClusterHealth2GatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CloudDiagnosticsCheckClusterHealth2GatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudDiagnosticsCheckClusterHealth2Default creates a CloudDiagnosticsCheckClusterHealth2Default with default headers values
func NewCloudDiagnosticsCheckClusterHealth2Default(code int) *CloudDiagnosticsCheckClusterHealth2Default {
	return &CloudDiagnosticsCheckClusterHealth2Default{
		_statusCode: code,
	}
}

/*
CloudDiagnosticsCheckClusterHealth2Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CloudDiagnosticsCheckClusterHealth2Default struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// IsSuccess returns true when this cloud diagnostics check cluster health2 default response has a 2xx status code
func (o *CloudDiagnosticsCheckClusterHealth2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cloud diagnostics check cluster health2 default response has a 3xx status code
func (o *CloudDiagnosticsCheckClusterHealth2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cloud diagnostics check cluster health2 default response has a 4xx status code
func (o *CloudDiagnosticsCheckClusterHealth2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cloud diagnostics check cluster health2 default response has a 5xx status code
func (o *CloudDiagnosticsCheckClusterHealth2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cloud diagnostics check cluster health2 default response a status code equal to that given
func (o *CloudDiagnosticsCheckClusterHealth2Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cloud diagnostics check cluster health2 default response
func (o *CloudDiagnosticsCheckClusterHealth2Default) Code() int {
	return o._statusCode
}

func (o *CloudDiagnosticsCheckClusterHealth2Default) Error() string {
	return fmt.Sprintf("[GET /v1/cloud/ping/id/{pingId}][%d] CloudDiagnostics_checkClusterHealth2 default  %+v", o._statusCode, o.Payload)
}

func (o *CloudDiagnosticsCheckClusterHealth2Default) String() string {
	return fmt.Sprintf("[GET /v1/cloud/ping/id/{pingId}][%d] CloudDiagnostics_checkClusterHealth2 default  %+v", o._statusCode, o.Payload)
}

func (o *CloudDiagnosticsCheckClusterHealth2Default) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *CloudDiagnosticsCheckClusterHealth2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

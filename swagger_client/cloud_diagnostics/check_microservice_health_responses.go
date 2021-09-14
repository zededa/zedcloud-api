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

// CheckMicroserviceHealthReader is a Reader for the CheckMicroserviceHealth structure.
type CheckMicroserviceHealthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckMicroserviceHealthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckMicroserviceHealthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCheckMicroserviceHealthBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCheckMicroserviceHealthInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewCheckMicroserviceHealthGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCheckMicroserviceHealthOK creates a CheckMicroserviceHealthOK with default headers values
func NewCheckMicroserviceHealthOK() *CheckMicroserviceHealthOK {
	return &CheckMicroserviceHealthOK{}
}

/* CheckMicroserviceHealthOK describes a response with status code 200, with default header values.

A successful response.
*/
type CheckMicroserviceHealthOK struct {
	Payload *swagger_models.HelloResp
}

func (o *CheckMicroserviceHealthOK) Error() string {
	return fmt.Sprintf("[POST /v1/hello][%d] checkMicroserviceHealthOK  %+v", 200, o.Payload)
}
func (o *CheckMicroserviceHealthOK) GetPayload() *swagger_models.HelloResp {
	return o.Payload
}

func (o *CheckMicroserviceHealthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.HelloResp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckMicroserviceHealthBadRequest creates a CheckMicroserviceHealthBadRequest with default headers values
func NewCheckMicroserviceHealthBadRequest() *CheckMicroserviceHealthBadRequest {
	return &CheckMicroserviceHealthBadRequest{}
}

/* CheckMicroserviceHealthBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type CheckMicroserviceHealthBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CheckMicroserviceHealthBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/hello][%d] checkMicroserviceHealthBadRequest  %+v", 400, o.Payload)
}
func (o *CheckMicroserviceHealthBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CheckMicroserviceHealthBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckMicroserviceHealthInternalServerError creates a CheckMicroserviceHealthInternalServerError with default headers values
func NewCheckMicroserviceHealthInternalServerError() *CheckMicroserviceHealthInternalServerError {
	return &CheckMicroserviceHealthInternalServerError{}
}

/* CheckMicroserviceHealthInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type CheckMicroserviceHealthInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CheckMicroserviceHealthInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/hello][%d] checkMicroserviceHealthInternalServerError  %+v", 500, o.Payload)
}
func (o *CheckMicroserviceHealthInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CheckMicroserviceHealthInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckMicroserviceHealthGatewayTimeout creates a CheckMicroserviceHealthGatewayTimeout with default headers values
func NewCheckMicroserviceHealthGatewayTimeout() *CheckMicroserviceHealthGatewayTimeout {
	return &CheckMicroserviceHealthGatewayTimeout{}
}

/* CheckMicroserviceHealthGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type CheckMicroserviceHealthGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CheckMicroserviceHealthGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v1/hello][%d] checkMicroserviceHealthGatewayTimeout  %+v", 504, o.Payload)
}
func (o *CheckMicroserviceHealthGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CheckMicroserviceHealthGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

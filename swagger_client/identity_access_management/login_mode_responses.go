// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package identity_access_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// LoginModeReader is a Reader for the LoginMode structure.
type LoginModeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LoginModeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLoginModeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewLoginModeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewLoginModeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewLoginModeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewLoginModeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLoginModeOK creates a LoginModeOK with default headers values
func NewLoginModeOK() *LoginModeOK {
	return &LoginModeOK{}
}

/* LoginModeOK describes a response with status code 200, with default header values.

A successful response.
*/
type LoginModeOK struct {
	Payload *swagger_models.AAAFrontendLoginModeResponse
}

func (o *LoginModeOK) Error() string {
	return fmt.Sprintf("[POST /v1/login/mode][%d] loginModeOK  %+v", 200, o.Payload)
}
func (o *LoginModeOK) GetPayload() *swagger_models.AAAFrontendLoginModeResponse {
	return o.Payload
}

func (o *LoginModeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.AAAFrontendLoginModeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoginModeUnauthorized creates a LoginModeUnauthorized with default headers values
func NewLoginModeUnauthorized() *LoginModeUnauthorized {
	return &LoginModeUnauthorized{}
}

/* LoginModeUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type LoginModeUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *LoginModeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/login/mode][%d] loginModeUnauthorized  %+v", 401, o.Payload)
}
func (o *LoginModeUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *LoginModeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoginModeForbidden creates a LoginModeForbidden with default headers values
func NewLoginModeForbidden() *LoginModeForbidden {
	return &LoginModeForbidden{}
}

/* LoginModeForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type LoginModeForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *LoginModeForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/login/mode][%d] loginModeForbidden  %+v", 403, o.Payload)
}
func (o *LoginModeForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *LoginModeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoginModeInternalServerError creates a LoginModeInternalServerError with default headers values
func NewLoginModeInternalServerError() *LoginModeInternalServerError {
	return &LoginModeInternalServerError{}
}

/* LoginModeInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type LoginModeInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *LoginModeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/login/mode][%d] loginModeInternalServerError  %+v", 500, o.Payload)
}
func (o *LoginModeInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *LoginModeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoginModeGatewayTimeout creates a LoginModeGatewayTimeout with default headers values
func NewLoginModeGatewayTimeout() *LoginModeGatewayTimeout {
	return &LoginModeGatewayTimeout{}
}

/* LoginModeGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type LoginModeGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *LoginModeGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v1/login/mode][%d] loginModeGatewayTimeout  %+v", 504, o.Payload)
}
func (o *LoginModeGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *LoginModeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

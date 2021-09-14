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

// LogoutReader is a Reader for the Logout structure.
type LogoutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LogoutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLogoutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewLogoutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewLogoutForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewLogoutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewLogoutGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLogoutOK creates a LogoutOK with default headers values
func NewLogoutOK() *LogoutOK {
	return &LogoutOK{}
}

/* LogoutOK describes a response with status code 200, with default header values.

A successful response.
*/
type LogoutOK struct {
	Payload *swagger_models.AAAFrontendLogoutResponse
}

func (o *LogoutOK) Error() string {
	return fmt.Sprintf("[POST /v1/logout][%d] logoutOK  %+v", 200, o.Payload)
}
func (o *LogoutOK) GetPayload() *swagger_models.AAAFrontendLogoutResponse {
	return o.Payload
}

func (o *LogoutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.AAAFrontendLogoutResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLogoutUnauthorized creates a LogoutUnauthorized with default headers values
func NewLogoutUnauthorized() *LogoutUnauthorized {
	return &LogoutUnauthorized{}
}

/* LogoutUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type LogoutUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *LogoutUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/logout][%d] logoutUnauthorized  %+v", 401, o.Payload)
}
func (o *LogoutUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *LogoutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLogoutForbidden creates a LogoutForbidden with default headers values
func NewLogoutForbidden() *LogoutForbidden {
	return &LogoutForbidden{}
}

/* LogoutForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type LogoutForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *LogoutForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/logout][%d] logoutForbidden  %+v", 403, o.Payload)
}
func (o *LogoutForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *LogoutForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLogoutInternalServerError creates a LogoutInternalServerError with default headers values
func NewLogoutInternalServerError() *LogoutInternalServerError {
	return &LogoutInternalServerError{}
}

/* LogoutInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type LogoutInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *LogoutInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/logout][%d] logoutInternalServerError  %+v", 500, o.Payload)
}
func (o *LogoutInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *LogoutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLogoutGatewayTimeout creates a LogoutGatewayTimeout with default headers values
func NewLogoutGatewayTimeout() *LogoutGatewayTimeout {
	return &LogoutGatewayTimeout{}
}

/* LogoutGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type LogoutGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *LogoutGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v1/logout][%d] logoutGatewayTimeout  %+v", 504, o.Payload)
}
func (o *LogoutGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *LogoutGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package edge_application_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// EdgeApplicationConfigurationGetEdgeApplicationBundleReader is a Reader for the EdgeApplicationConfigurationGetEdgeApplicationBundle structure.
type EdgeApplicationConfigurationGetEdgeApplicationBundleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeApplicationConfigurationGetEdgeApplicationBundleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeApplicationConfigurationGetEdgeApplicationBundleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeApplicationConfigurationGetEdgeApplicationBundleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeApplicationConfigurationGetEdgeApplicationBundleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeApplicationConfigurationGetEdgeApplicationBundleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeApplicationConfigurationGetEdgeApplicationBundleGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeApplicationConfigurationGetEdgeApplicationBundleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeApplicationConfigurationGetEdgeApplicationBundleOK creates a EdgeApplicationConfigurationGetEdgeApplicationBundleOK with default headers values
func NewEdgeApplicationConfigurationGetEdgeApplicationBundleOK() *EdgeApplicationConfigurationGetEdgeApplicationBundleOK {
	return &EdgeApplicationConfigurationGetEdgeApplicationBundleOK{}
}

/* EdgeApplicationConfigurationGetEdgeApplicationBundleOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeApplicationConfigurationGetEdgeApplicationBundleOK struct {
	Payload *swagger_models.App
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleOK) Error() string {
	return fmt.Sprintf("[GET /v1/apps/id/{id}][%d] edgeApplicationConfigurationGetEdgeApplicationBundleOK  %+v", 200, o.Payload)
}
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleOK) GetPayload() *swagger_models.App {
	return o.Payload
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.App)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationGetEdgeApplicationBundleUnauthorized creates a EdgeApplicationConfigurationGetEdgeApplicationBundleUnauthorized with default headers values
func NewEdgeApplicationConfigurationGetEdgeApplicationBundleUnauthorized() *EdgeApplicationConfigurationGetEdgeApplicationBundleUnauthorized {
	return &EdgeApplicationConfigurationGetEdgeApplicationBundleUnauthorized{}
}

/* EdgeApplicationConfigurationGetEdgeApplicationBundleUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeApplicationConfigurationGetEdgeApplicationBundleUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/apps/id/{id}][%d] edgeApplicationConfigurationGetEdgeApplicationBundleUnauthorized  %+v", 401, o.Payload)
}
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationGetEdgeApplicationBundleForbidden creates a EdgeApplicationConfigurationGetEdgeApplicationBundleForbidden with default headers values
func NewEdgeApplicationConfigurationGetEdgeApplicationBundleForbidden() *EdgeApplicationConfigurationGetEdgeApplicationBundleForbidden {
	return &EdgeApplicationConfigurationGetEdgeApplicationBundleForbidden{}
}

/* EdgeApplicationConfigurationGetEdgeApplicationBundleForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have application level access permission for the operation or does not have access scope to the project.
*/
type EdgeApplicationConfigurationGetEdgeApplicationBundleForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/apps/id/{id}][%d] edgeApplicationConfigurationGetEdgeApplicationBundleForbidden  %+v", 403, o.Payload)
}
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationGetEdgeApplicationBundleNotFound creates a EdgeApplicationConfigurationGetEdgeApplicationBundleNotFound with default headers values
func NewEdgeApplicationConfigurationGetEdgeApplicationBundleNotFound() *EdgeApplicationConfigurationGetEdgeApplicationBundleNotFound {
	return &EdgeApplicationConfigurationGetEdgeApplicationBundleNotFound{}
}

/* EdgeApplicationConfigurationGetEdgeApplicationBundleNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeApplicationConfigurationGetEdgeApplicationBundleNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/apps/id/{id}][%d] edgeApplicationConfigurationGetEdgeApplicationBundleNotFound  %+v", 404, o.Payload)
}
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationGetEdgeApplicationBundleInternalServerError creates a EdgeApplicationConfigurationGetEdgeApplicationBundleInternalServerError with default headers values
func NewEdgeApplicationConfigurationGetEdgeApplicationBundleInternalServerError() *EdgeApplicationConfigurationGetEdgeApplicationBundleInternalServerError {
	return &EdgeApplicationConfigurationGetEdgeApplicationBundleInternalServerError{}
}

/* EdgeApplicationConfigurationGetEdgeApplicationBundleInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeApplicationConfigurationGetEdgeApplicationBundleInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/apps/id/{id}][%d] edgeApplicationConfigurationGetEdgeApplicationBundleInternalServerError  %+v", 500, o.Payload)
}
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationGetEdgeApplicationBundleGatewayTimeout creates a EdgeApplicationConfigurationGetEdgeApplicationBundleGatewayTimeout with default headers values
func NewEdgeApplicationConfigurationGetEdgeApplicationBundleGatewayTimeout() *EdgeApplicationConfigurationGetEdgeApplicationBundleGatewayTimeout {
	return &EdgeApplicationConfigurationGetEdgeApplicationBundleGatewayTimeout{}
}

/* EdgeApplicationConfigurationGetEdgeApplicationBundleGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeApplicationConfigurationGetEdgeApplicationBundleGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/apps/id/{id}][%d] edgeApplicationConfigurationGetEdgeApplicationBundleGatewayTimeout  %+v", 504, o.Payload)
}
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationGetEdgeApplicationBundleDefault creates a EdgeApplicationConfigurationGetEdgeApplicationBundleDefault with default headers values
func NewEdgeApplicationConfigurationGetEdgeApplicationBundleDefault(code int) *EdgeApplicationConfigurationGetEdgeApplicationBundleDefault {
	return &EdgeApplicationConfigurationGetEdgeApplicationBundleDefault{
		_statusCode: code,
	}
}

/* EdgeApplicationConfigurationGetEdgeApplicationBundleDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeApplicationConfigurationGetEdgeApplicationBundleDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the edge application configuration get edge application bundle default response
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleDefault) Code() int {
	return o._statusCode
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleDefault) Error() string {
	return fmt.Sprintf("[GET /v1/apps/id/{id}][%d] EdgeApplicationConfiguration_GetEdgeApplicationBundle default  %+v", o._statusCode, o.Payload)
}
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
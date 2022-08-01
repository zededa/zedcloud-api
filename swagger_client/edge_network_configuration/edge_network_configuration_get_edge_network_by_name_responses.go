// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package edge_network_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// EdgeNetworkConfigurationGetEdgeNetworkByNameReader is a Reader for the EdgeNetworkConfigurationGetEdgeNetworkByName structure.
type EdgeNetworkConfigurationGetEdgeNetworkByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeNetworkConfigurationGetEdgeNetworkByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeNetworkConfigurationGetEdgeNetworkByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeNetworkConfigurationGetEdgeNetworkByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeNetworkConfigurationGetEdgeNetworkByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeNetworkConfigurationGetEdgeNetworkByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeNetworkConfigurationGetEdgeNetworkByNameGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeNetworkConfigurationGetEdgeNetworkByNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeNetworkConfigurationGetEdgeNetworkByNameOK creates a EdgeNetworkConfigurationGetEdgeNetworkByNameOK with default headers values
func NewEdgeNetworkConfigurationGetEdgeNetworkByNameOK() *EdgeNetworkConfigurationGetEdgeNetworkByNameOK {
	return &EdgeNetworkConfigurationGetEdgeNetworkByNameOK{}
}

/* EdgeNetworkConfigurationGetEdgeNetworkByNameOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeNetworkConfigurationGetEdgeNetworkByNameOK struct {
	Payload *swagger_models.NetConfig
}

func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/networks/name/{name}][%d] edgeNetworkConfigurationGetEdgeNetworkByNameOK  %+v", 200, o.Payload)
}
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameOK) GetPayload() *swagger_models.NetConfig {
	return o.Payload
}

func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.NetConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkConfigurationGetEdgeNetworkByNameUnauthorized creates a EdgeNetworkConfigurationGetEdgeNetworkByNameUnauthorized with default headers values
func NewEdgeNetworkConfigurationGetEdgeNetworkByNameUnauthorized() *EdgeNetworkConfigurationGetEdgeNetworkByNameUnauthorized {
	return &EdgeNetworkConfigurationGetEdgeNetworkByNameUnauthorized{}
}

/* EdgeNetworkConfigurationGetEdgeNetworkByNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeNetworkConfigurationGetEdgeNetworkByNameUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/networks/name/{name}][%d] edgeNetworkConfigurationGetEdgeNetworkByNameUnauthorized  %+v", 401, o.Payload)
}
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkConfigurationGetEdgeNetworkByNameForbidden creates a EdgeNetworkConfigurationGetEdgeNetworkByNameForbidden with default headers values
func NewEdgeNetworkConfigurationGetEdgeNetworkByNameForbidden() *EdgeNetworkConfigurationGetEdgeNetworkByNameForbidden {
	return &EdgeNetworkConfigurationGetEdgeNetworkByNameForbidden{}
}

/* EdgeNetworkConfigurationGetEdgeNetworkByNameForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeNetworkConfigurationGetEdgeNetworkByNameForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/networks/name/{name}][%d] edgeNetworkConfigurationGetEdgeNetworkByNameForbidden  %+v", 403, o.Payload)
}
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkConfigurationGetEdgeNetworkByNameNotFound creates a EdgeNetworkConfigurationGetEdgeNetworkByNameNotFound with default headers values
func NewEdgeNetworkConfigurationGetEdgeNetworkByNameNotFound() *EdgeNetworkConfigurationGetEdgeNetworkByNameNotFound {
	return &EdgeNetworkConfigurationGetEdgeNetworkByNameNotFound{}
}

/* EdgeNetworkConfigurationGetEdgeNetworkByNameNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeNetworkConfigurationGetEdgeNetworkByNameNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/networks/name/{name}][%d] edgeNetworkConfigurationGetEdgeNetworkByNameNotFound  %+v", 404, o.Payload)
}
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkConfigurationGetEdgeNetworkByNameInternalServerError creates a EdgeNetworkConfigurationGetEdgeNetworkByNameInternalServerError with default headers values
func NewEdgeNetworkConfigurationGetEdgeNetworkByNameInternalServerError() *EdgeNetworkConfigurationGetEdgeNetworkByNameInternalServerError {
	return &EdgeNetworkConfigurationGetEdgeNetworkByNameInternalServerError{}
}

/* EdgeNetworkConfigurationGetEdgeNetworkByNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeNetworkConfigurationGetEdgeNetworkByNameInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/networks/name/{name}][%d] edgeNetworkConfigurationGetEdgeNetworkByNameInternalServerError  %+v", 500, o.Payload)
}
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkConfigurationGetEdgeNetworkByNameGatewayTimeout creates a EdgeNetworkConfigurationGetEdgeNetworkByNameGatewayTimeout with default headers values
func NewEdgeNetworkConfigurationGetEdgeNetworkByNameGatewayTimeout() *EdgeNetworkConfigurationGetEdgeNetworkByNameGatewayTimeout {
	return &EdgeNetworkConfigurationGetEdgeNetworkByNameGatewayTimeout{}
}

/* EdgeNetworkConfigurationGetEdgeNetworkByNameGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeNetworkConfigurationGetEdgeNetworkByNameGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/networks/name/{name}][%d] edgeNetworkConfigurationGetEdgeNetworkByNameGatewayTimeout  %+v", 504, o.Payload)
}
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkConfigurationGetEdgeNetworkByNameDefault creates a EdgeNetworkConfigurationGetEdgeNetworkByNameDefault with default headers values
func NewEdgeNetworkConfigurationGetEdgeNetworkByNameDefault(code int) *EdgeNetworkConfigurationGetEdgeNetworkByNameDefault {
	return &EdgeNetworkConfigurationGetEdgeNetworkByNameDefault{
		_statusCode: code,
	}
}

/* EdgeNetworkConfigurationGetEdgeNetworkByNameDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeNetworkConfigurationGetEdgeNetworkByNameDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the edge network configuration get edge network by name default response
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameDefault) Code() int {
	return o._statusCode
}

func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameDefault) Error() string {
	return fmt.Sprintf("[GET /v1/networks/name/{name}][%d] EdgeNetworkConfiguration_GetEdgeNetworkByName default  %+v", o._statusCode, o.Payload)
}
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

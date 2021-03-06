// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package edge_node_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// GetEdgeNodeBySerialReader is a Reader for the GetEdgeNodeBySerial structure.
type GetEdgeNodeBySerialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEdgeNodeBySerialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEdgeNodeBySerialOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetEdgeNodeBySerialUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetEdgeNodeBySerialForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetEdgeNodeBySerialNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetEdgeNodeBySerialInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetEdgeNodeBySerialGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetEdgeNodeBySerialOK creates a GetEdgeNodeBySerialOK with default headers values
func NewGetEdgeNodeBySerialOK() *GetEdgeNodeBySerialOK {
	return &GetEdgeNodeBySerialOK{}
}

/* GetEdgeNodeBySerialOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetEdgeNodeBySerialOK struct {
	Payload *swagger_models.DeviceConfig
}

func (o *GetEdgeNodeBySerialOK) Error() string {
	return fmt.Sprintf("[GET /v1/devices/serial/{serialno}][%d] getEdgeNodeBySerialOK  %+v", 200, o.Payload)
}
func (o *GetEdgeNodeBySerialOK) GetPayload() *swagger_models.DeviceConfig {
	return o.Payload
}

func (o *GetEdgeNodeBySerialOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.DeviceConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeNodeBySerialUnauthorized creates a GetEdgeNodeBySerialUnauthorized with default headers values
func NewGetEdgeNodeBySerialUnauthorized() *GetEdgeNodeBySerialUnauthorized {
	return &GetEdgeNodeBySerialUnauthorized{}
}

/* GetEdgeNodeBySerialUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type GetEdgeNodeBySerialUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEdgeNodeBySerialUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/devices/serial/{serialno}][%d] getEdgeNodeBySerialUnauthorized  %+v", 401, o.Payload)
}
func (o *GetEdgeNodeBySerialUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEdgeNodeBySerialUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeNodeBySerialForbidden creates a GetEdgeNodeBySerialForbidden with default headers values
func NewGetEdgeNodeBySerialForbidden() *GetEdgeNodeBySerialForbidden {
	return &GetEdgeNodeBySerialForbidden{}
}

/* GetEdgeNodeBySerialForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type GetEdgeNodeBySerialForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEdgeNodeBySerialForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/devices/serial/{serialno}][%d] getEdgeNodeBySerialForbidden  %+v", 403, o.Payload)
}
func (o *GetEdgeNodeBySerialForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEdgeNodeBySerialForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeNodeBySerialNotFound creates a GetEdgeNodeBySerialNotFound with default headers values
func NewGetEdgeNodeBySerialNotFound() *GetEdgeNodeBySerialNotFound {
	return &GetEdgeNodeBySerialNotFound{}
}

/* GetEdgeNodeBySerialNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type GetEdgeNodeBySerialNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEdgeNodeBySerialNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/devices/serial/{serialno}][%d] getEdgeNodeBySerialNotFound  %+v", 404, o.Payload)
}
func (o *GetEdgeNodeBySerialNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEdgeNodeBySerialNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeNodeBySerialInternalServerError creates a GetEdgeNodeBySerialInternalServerError with default headers values
func NewGetEdgeNodeBySerialInternalServerError() *GetEdgeNodeBySerialInternalServerError {
	return &GetEdgeNodeBySerialInternalServerError{}
}

/* GetEdgeNodeBySerialInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type GetEdgeNodeBySerialInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEdgeNodeBySerialInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/devices/serial/{serialno}][%d] getEdgeNodeBySerialInternalServerError  %+v", 500, o.Payload)
}
func (o *GetEdgeNodeBySerialInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEdgeNodeBySerialInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeNodeBySerialGatewayTimeout creates a GetEdgeNodeBySerialGatewayTimeout with default headers values
func NewGetEdgeNodeBySerialGatewayTimeout() *GetEdgeNodeBySerialGatewayTimeout {
	return &GetEdgeNodeBySerialGatewayTimeout{}
}

/* GetEdgeNodeBySerialGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type GetEdgeNodeBySerialGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEdgeNodeBySerialGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/devices/serial/{serialno}][%d] getEdgeNodeBySerialGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetEdgeNodeBySerialGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEdgeNodeBySerialGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

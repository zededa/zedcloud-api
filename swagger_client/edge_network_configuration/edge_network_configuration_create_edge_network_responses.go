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

// EdgeNetworkConfigurationCreateEdgeNetworkReader is a Reader for the EdgeNetworkConfigurationCreateEdgeNetwork structure.
type EdgeNetworkConfigurationCreateEdgeNetworkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeNetworkConfigurationCreateEdgeNetworkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeNetworkConfigurationCreateEdgeNetworkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEdgeNetworkConfigurationCreateEdgeNetworkBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewEdgeNetworkConfigurationCreateEdgeNetworkUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeNetworkConfigurationCreateEdgeNetworkForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewEdgeNetworkConfigurationCreateEdgeNetworkConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeNetworkConfigurationCreateEdgeNetworkInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeNetworkConfigurationCreateEdgeNetworkGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeNetworkConfigurationCreateEdgeNetworkDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeNetworkConfigurationCreateEdgeNetworkOK creates a EdgeNetworkConfigurationCreateEdgeNetworkOK with default headers values
func NewEdgeNetworkConfigurationCreateEdgeNetworkOK() *EdgeNetworkConfigurationCreateEdgeNetworkOK {
	return &EdgeNetworkConfigurationCreateEdgeNetworkOK{}
}

/* EdgeNetworkConfigurationCreateEdgeNetworkOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeNetworkConfigurationCreateEdgeNetworkOK struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkOK) Error() string {
	return fmt.Sprintf("[POST /v1/networks][%d] edgeNetworkConfigurationCreateEdgeNetworkOK  %+v", 200, o.Payload)
}
func (o *EdgeNetworkConfigurationCreateEdgeNetworkOK) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkConfigurationCreateEdgeNetworkBadRequest creates a EdgeNetworkConfigurationCreateEdgeNetworkBadRequest with default headers values
func NewEdgeNetworkConfigurationCreateEdgeNetworkBadRequest() *EdgeNetworkConfigurationCreateEdgeNetworkBadRequest {
	return &EdgeNetworkConfigurationCreateEdgeNetworkBadRequest{}
}

/* EdgeNetworkConfigurationCreateEdgeNetworkBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type EdgeNetworkConfigurationCreateEdgeNetworkBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/networks][%d] edgeNetworkConfigurationCreateEdgeNetworkBadRequest  %+v", 400, o.Payload)
}
func (o *EdgeNetworkConfigurationCreateEdgeNetworkBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkConfigurationCreateEdgeNetworkUnauthorized creates a EdgeNetworkConfigurationCreateEdgeNetworkUnauthorized with default headers values
func NewEdgeNetworkConfigurationCreateEdgeNetworkUnauthorized() *EdgeNetworkConfigurationCreateEdgeNetworkUnauthorized {
	return &EdgeNetworkConfigurationCreateEdgeNetworkUnauthorized{}
}

/* EdgeNetworkConfigurationCreateEdgeNetworkUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeNetworkConfigurationCreateEdgeNetworkUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/networks][%d] edgeNetworkConfigurationCreateEdgeNetworkUnauthorized  %+v", 401, o.Payload)
}
func (o *EdgeNetworkConfigurationCreateEdgeNetworkUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkConfigurationCreateEdgeNetworkForbidden creates a EdgeNetworkConfigurationCreateEdgeNetworkForbidden with default headers values
func NewEdgeNetworkConfigurationCreateEdgeNetworkForbidden() *EdgeNetworkConfigurationCreateEdgeNetworkForbidden {
	return &EdgeNetworkConfigurationCreateEdgeNetworkForbidden{}
}

/* EdgeNetworkConfigurationCreateEdgeNetworkForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeNetworkConfigurationCreateEdgeNetworkForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/networks][%d] edgeNetworkConfigurationCreateEdgeNetworkForbidden  %+v", 403, o.Payload)
}
func (o *EdgeNetworkConfigurationCreateEdgeNetworkForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkConfigurationCreateEdgeNetworkConflict creates a EdgeNetworkConfigurationCreateEdgeNetworkConflict with default headers values
func NewEdgeNetworkConfigurationCreateEdgeNetworkConflict() *EdgeNetworkConfigurationCreateEdgeNetworkConflict {
	return &EdgeNetworkConfigurationCreateEdgeNetworkConflict{}
}

/* EdgeNetworkConfigurationCreateEdgeNetworkConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this edge network record will conflict with an already existing edge network record.
*/
type EdgeNetworkConfigurationCreateEdgeNetworkConflict struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkConflict) Error() string {
	return fmt.Sprintf("[POST /v1/networks][%d] edgeNetworkConfigurationCreateEdgeNetworkConflict  %+v", 409, o.Payload)
}
func (o *EdgeNetworkConfigurationCreateEdgeNetworkConflict) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkConfigurationCreateEdgeNetworkInternalServerError creates a EdgeNetworkConfigurationCreateEdgeNetworkInternalServerError with default headers values
func NewEdgeNetworkConfigurationCreateEdgeNetworkInternalServerError() *EdgeNetworkConfigurationCreateEdgeNetworkInternalServerError {
	return &EdgeNetworkConfigurationCreateEdgeNetworkInternalServerError{}
}

/* EdgeNetworkConfigurationCreateEdgeNetworkInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeNetworkConfigurationCreateEdgeNetworkInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/networks][%d] edgeNetworkConfigurationCreateEdgeNetworkInternalServerError  %+v", 500, o.Payload)
}
func (o *EdgeNetworkConfigurationCreateEdgeNetworkInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkConfigurationCreateEdgeNetworkGatewayTimeout creates a EdgeNetworkConfigurationCreateEdgeNetworkGatewayTimeout with default headers values
func NewEdgeNetworkConfigurationCreateEdgeNetworkGatewayTimeout() *EdgeNetworkConfigurationCreateEdgeNetworkGatewayTimeout {
	return &EdgeNetworkConfigurationCreateEdgeNetworkGatewayTimeout{}
}

/* EdgeNetworkConfigurationCreateEdgeNetworkGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeNetworkConfigurationCreateEdgeNetworkGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v1/networks][%d] edgeNetworkConfigurationCreateEdgeNetworkGatewayTimeout  %+v", 504, o.Payload)
}
func (o *EdgeNetworkConfigurationCreateEdgeNetworkGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkConfigurationCreateEdgeNetworkDefault creates a EdgeNetworkConfigurationCreateEdgeNetworkDefault with default headers values
func NewEdgeNetworkConfigurationCreateEdgeNetworkDefault(code int) *EdgeNetworkConfigurationCreateEdgeNetworkDefault {
	return &EdgeNetworkConfigurationCreateEdgeNetworkDefault{
		_statusCode: code,
	}
}

/* EdgeNetworkConfigurationCreateEdgeNetworkDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeNetworkConfigurationCreateEdgeNetworkDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the edge network configuration create edge network default response
func (o *EdgeNetworkConfigurationCreateEdgeNetworkDefault) Code() int {
	return o._statusCode
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkDefault) Error() string {
	return fmt.Sprintf("[POST /v1/networks][%d] EdgeNetworkConfiguration_CreateEdgeNetwork default  %+v", o._statusCode, o.Payload)
}
func (o *EdgeNetworkConfigurationCreateEdgeNetworkDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeNetworkConfigurationCreateEdgeNetworkDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
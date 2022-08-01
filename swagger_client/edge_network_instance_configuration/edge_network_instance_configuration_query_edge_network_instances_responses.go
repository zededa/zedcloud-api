// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package edge_network_instance_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesReader is a Reader for the EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstances structure.
type EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewEdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesOK creates a EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesOK with default headers values
func NewEdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesOK() *EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesOK {
	return &EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesOK{}
}

/* EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesOK struct {
	Payload *swagger_models.NetInstList
}

func (o *EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesOK) Error() string {
	return fmt.Sprintf("[GET /v1/netinsts][%d] edgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesOK  %+v", 200, o.Payload)
}
func (o *EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesOK) GetPayload() *swagger_models.NetInstList {
	return o.Payload
}

func (o *EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.NetInstList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesBadRequest creates a EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesBadRequest with default headers values
func NewEdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesBadRequest() *EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesBadRequest {
	return &EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesBadRequest{}
}

/* EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of invalid value of filter parameters.
*/
type EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/netinsts][%d] edgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesBadRequest  %+v", 400, o.Payload)
}
func (o *EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesUnauthorized creates a EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesUnauthorized with default headers values
func NewEdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesUnauthorized() *EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesUnauthorized {
	return &EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesUnauthorized{}
}

/* EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/netinsts][%d] edgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesUnauthorized  %+v", 401, o.Payload)
}
func (o *EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesForbidden creates a EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesForbidden with default headers values
func NewEdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesForbidden() *EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesForbidden {
	return &EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesForbidden{}
}

/* EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/netinsts][%d] edgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesForbidden  %+v", 403, o.Payload)
}
func (o *EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesInternalServerError creates a EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesInternalServerError with default headers values
func NewEdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesInternalServerError() *EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesInternalServerError {
	return &EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesInternalServerError{}
}

/* EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/netinsts][%d] edgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesInternalServerError  %+v", 500, o.Payload)
}
func (o *EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesGatewayTimeout creates a EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesGatewayTimeout with default headers values
func NewEdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesGatewayTimeout() *EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesGatewayTimeout {
	return &EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesGatewayTimeout{}
}

/* EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/netinsts][%d] edgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesGatewayTimeout  %+v", 504, o.Payload)
}
func (o *EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesDefault creates a EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesDefault with default headers values
func NewEdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesDefault(code int) *EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesDefault {
	return &EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesDefault{
		_statusCode: code,
	}
}

/* EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the edge network instance configuration query edge network instances default response
func (o *EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesDefault) Code() int {
	return o._statusCode
}

func (o *EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesDefault) Error() string {
	return fmt.Sprintf("[GET /v1/netinsts][%d] EdgeNetworkInstanceConfiguration_QueryEdgeNetworkInstances default  %+v", o._statusCode, o.Payload)
}
func (o *EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeNetworkInstanceConfigurationQueryEdgeNetworkInstancesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
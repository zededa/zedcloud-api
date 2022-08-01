// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package edge_application_instance_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesReader is a Reader for the EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstances structure.
type EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesOK creates a EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesOK with default headers values
func NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesOK() *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesOK {
	return &EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesOK{}
}

/* EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesOK struct {
	Payload *swagger_models.AppInstances
}

func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesOK) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances][%d] edgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesOK  %+v", 200, o.Payload)
}
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesOK) GetPayload() *swagger_models.AppInstances {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.AppInstances)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesBadRequest creates a EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesBadRequest with default headers values
func NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesBadRequest() *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesBadRequest {
	return &EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesBadRequest{}
}

/* EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of invalid value of filter parameters.
*/
type EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances][%d] edgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesBadRequest  %+v", 400, o.Payload)
}
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesUnauthorized creates a EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesUnauthorized with default headers values
func NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesUnauthorized() *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesUnauthorized {
	return &EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesUnauthorized{}
}

/* EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances][%d] edgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesUnauthorized  %+v", 401, o.Payload)
}
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesForbidden creates a EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesForbidden with default headers values
func NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesForbidden() *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesForbidden {
	return &EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesForbidden{}
}

/* EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have application level access permission for the operation or does not have access scope to the project.
*/
type EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances][%d] edgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesForbidden  %+v", 403, o.Payload)
}
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesInternalServerError creates a EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesInternalServerError with default headers values
func NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesInternalServerError() *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesInternalServerError {
	return &EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesInternalServerError{}
}

/* EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances][%d] edgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesInternalServerError  %+v", 500, o.Payload)
}
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesGatewayTimeout creates a EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesGatewayTimeout with default headers values
func NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesGatewayTimeout() *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesGatewayTimeout {
	return &EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesGatewayTimeout{}
}

/* EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances][%d] edgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesGatewayTimeout  %+v", 504, o.Payload)
}
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesDefault creates a EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesDefault with default headers values
func NewEdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesDefault(code int) *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesDefault {
	return &EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesDefault{
		_statusCode: code,
	}
}

/* EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the edge application instance configuration query edge application instances default response
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesDefault) Code() int {
	return o._statusCode
}

func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesDefault) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances][%d] EdgeApplicationInstanceConfiguration_QueryEdgeApplicationInstances default  %+v", o._statusCode, o.Payload)
}
func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationQueryEdgeApplicationInstancesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

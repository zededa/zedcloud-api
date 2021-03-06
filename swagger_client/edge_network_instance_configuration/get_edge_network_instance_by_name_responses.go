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

// GetEdgeNetworkInstanceByNameReader is a Reader for the GetEdgeNetworkInstanceByName structure.
type GetEdgeNetworkInstanceByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEdgeNetworkInstanceByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEdgeNetworkInstanceByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetEdgeNetworkInstanceByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetEdgeNetworkInstanceByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetEdgeNetworkInstanceByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetEdgeNetworkInstanceByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetEdgeNetworkInstanceByNameGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetEdgeNetworkInstanceByNameOK creates a GetEdgeNetworkInstanceByNameOK with default headers values
func NewGetEdgeNetworkInstanceByNameOK() *GetEdgeNetworkInstanceByNameOK {
	return &GetEdgeNetworkInstanceByNameOK{}
}

/* GetEdgeNetworkInstanceByNameOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetEdgeNetworkInstanceByNameOK struct {
	Payload *swagger_models.NetInstConfig
}

func (o *GetEdgeNetworkInstanceByNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/netinsts/name/{name}][%d] getEdgeNetworkInstanceByNameOK  %+v", 200, o.Payload)
}
func (o *GetEdgeNetworkInstanceByNameOK) GetPayload() *swagger_models.NetInstConfig {
	return o.Payload
}

func (o *GetEdgeNetworkInstanceByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.NetInstConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeNetworkInstanceByNameUnauthorized creates a GetEdgeNetworkInstanceByNameUnauthorized with default headers values
func NewGetEdgeNetworkInstanceByNameUnauthorized() *GetEdgeNetworkInstanceByNameUnauthorized {
	return &GetEdgeNetworkInstanceByNameUnauthorized{}
}

/* GetEdgeNetworkInstanceByNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type GetEdgeNetworkInstanceByNameUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEdgeNetworkInstanceByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/netinsts/name/{name}][%d] getEdgeNetworkInstanceByNameUnauthorized  %+v", 401, o.Payload)
}
func (o *GetEdgeNetworkInstanceByNameUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEdgeNetworkInstanceByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeNetworkInstanceByNameForbidden creates a GetEdgeNetworkInstanceByNameForbidden with default headers values
func NewGetEdgeNetworkInstanceByNameForbidden() *GetEdgeNetworkInstanceByNameForbidden {
	return &GetEdgeNetworkInstanceByNameForbidden{}
}

/* GetEdgeNetworkInstanceByNameForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type GetEdgeNetworkInstanceByNameForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEdgeNetworkInstanceByNameForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/netinsts/name/{name}][%d] getEdgeNetworkInstanceByNameForbidden  %+v", 403, o.Payload)
}
func (o *GetEdgeNetworkInstanceByNameForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEdgeNetworkInstanceByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeNetworkInstanceByNameNotFound creates a GetEdgeNetworkInstanceByNameNotFound with default headers values
func NewGetEdgeNetworkInstanceByNameNotFound() *GetEdgeNetworkInstanceByNameNotFound {
	return &GetEdgeNetworkInstanceByNameNotFound{}
}

/* GetEdgeNetworkInstanceByNameNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type GetEdgeNetworkInstanceByNameNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEdgeNetworkInstanceByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/netinsts/name/{name}][%d] getEdgeNetworkInstanceByNameNotFound  %+v", 404, o.Payload)
}
func (o *GetEdgeNetworkInstanceByNameNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEdgeNetworkInstanceByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeNetworkInstanceByNameInternalServerError creates a GetEdgeNetworkInstanceByNameInternalServerError with default headers values
func NewGetEdgeNetworkInstanceByNameInternalServerError() *GetEdgeNetworkInstanceByNameInternalServerError {
	return &GetEdgeNetworkInstanceByNameInternalServerError{}
}

/* GetEdgeNetworkInstanceByNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type GetEdgeNetworkInstanceByNameInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEdgeNetworkInstanceByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/netinsts/name/{name}][%d] getEdgeNetworkInstanceByNameInternalServerError  %+v", 500, o.Payload)
}
func (o *GetEdgeNetworkInstanceByNameInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEdgeNetworkInstanceByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeNetworkInstanceByNameGatewayTimeout creates a GetEdgeNetworkInstanceByNameGatewayTimeout with default headers values
func NewGetEdgeNetworkInstanceByNameGatewayTimeout() *GetEdgeNetworkInstanceByNameGatewayTimeout {
	return &GetEdgeNetworkInstanceByNameGatewayTimeout{}
}

/* GetEdgeNetworkInstanceByNameGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type GetEdgeNetworkInstanceByNameGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEdgeNetworkInstanceByNameGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/netinsts/name/{name}][%d] getEdgeNetworkInstanceByNameGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetEdgeNetworkInstanceByNameGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEdgeNetworkInstanceByNameGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// GetEdgeNetworkByNameReader is a Reader for the GetEdgeNetworkByName structure.
type GetEdgeNetworkByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEdgeNetworkByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEdgeNetworkByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetEdgeNetworkByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetEdgeNetworkByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetEdgeNetworkByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetEdgeNetworkByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetEdgeNetworkByNameGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetEdgeNetworkByNameOK creates a GetEdgeNetworkByNameOK with default headers values
func NewGetEdgeNetworkByNameOK() *GetEdgeNetworkByNameOK {
	return &GetEdgeNetworkByNameOK{}
}

/* GetEdgeNetworkByNameOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetEdgeNetworkByNameOK struct {
	Payload *swagger_models.NetConfig
}

func (o *GetEdgeNetworkByNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/networks/name/{name}][%d] getEdgeNetworkByNameOK  %+v", 200, o.Payload)
}
func (o *GetEdgeNetworkByNameOK) GetPayload() *swagger_models.NetConfig {
	return o.Payload
}

func (o *GetEdgeNetworkByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.NetConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeNetworkByNameUnauthorized creates a GetEdgeNetworkByNameUnauthorized with default headers values
func NewGetEdgeNetworkByNameUnauthorized() *GetEdgeNetworkByNameUnauthorized {
	return &GetEdgeNetworkByNameUnauthorized{}
}

/* GetEdgeNetworkByNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type GetEdgeNetworkByNameUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEdgeNetworkByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/networks/name/{name}][%d] getEdgeNetworkByNameUnauthorized  %+v", 401, o.Payload)
}
func (o *GetEdgeNetworkByNameUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEdgeNetworkByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeNetworkByNameForbidden creates a GetEdgeNetworkByNameForbidden with default headers values
func NewGetEdgeNetworkByNameForbidden() *GetEdgeNetworkByNameForbidden {
	return &GetEdgeNetworkByNameForbidden{}
}

/* GetEdgeNetworkByNameForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type GetEdgeNetworkByNameForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEdgeNetworkByNameForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/networks/name/{name}][%d] getEdgeNetworkByNameForbidden  %+v", 403, o.Payload)
}
func (o *GetEdgeNetworkByNameForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEdgeNetworkByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeNetworkByNameNotFound creates a GetEdgeNetworkByNameNotFound with default headers values
func NewGetEdgeNetworkByNameNotFound() *GetEdgeNetworkByNameNotFound {
	return &GetEdgeNetworkByNameNotFound{}
}

/* GetEdgeNetworkByNameNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type GetEdgeNetworkByNameNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEdgeNetworkByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/networks/name/{name}][%d] getEdgeNetworkByNameNotFound  %+v", 404, o.Payload)
}
func (o *GetEdgeNetworkByNameNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEdgeNetworkByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeNetworkByNameInternalServerError creates a GetEdgeNetworkByNameInternalServerError with default headers values
func NewGetEdgeNetworkByNameInternalServerError() *GetEdgeNetworkByNameInternalServerError {
	return &GetEdgeNetworkByNameInternalServerError{}
}

/* GetEdgeNetworkByNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type GetEdgeNetworkByNameInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEdgeNetworkByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/networks/name/{name}][%d] getEdgeNetworkByNameInternalServerError  %+v", 500, o.Payload)
}
func (o *GetEdgeNetworkByNameInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEdgeNetworkByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeNetworkByNameGatewayTimeout creates a GetEdgeNetworkByNameGatewayTimeout with default headers values
func NewGetEdgeNetworkByNameGatewayTimeout() *GetEdgeNetworkByNameGatewayTimeout {
	return &GetEdgeNetworkByNameGatewayTimeout{}
}

/* GetEdgeNetworkByNameGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type GetEdgeNetworkByNameGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEdgeNetworkByNameGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/networks/name/{name}][%d] getEdgeNetworkByNameGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetEdgeNetworkByNameGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEdgeNetworkByNameGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

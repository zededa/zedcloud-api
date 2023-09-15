// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
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

// EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameReader is a Reader for the EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByName structure.
type EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameOK creates a EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameOK with default headers values
func NewEdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameOK() *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameOK {
	return &EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameOK{}
}

/*
EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameOK struct {
	Payload *swagger_models.NetInstConfig
}

// IsSuccess returns true when this edge network instance configuration get edge network instance by name o k response has a 2xx status code
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge network instance configuration get edge network instance by name o k response has a 3xx status code
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network instance configuration get edge network instance by name o k response has a 4xx status code
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge network instance configuration get edge network instance by name o k response has a 5xx status code
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge network instance configuration get edge network instance by name o k response a status code equal to that given
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the edge network instance configuration get edge network instance by name o k response
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameOK) Code() int {
	return 200
}

func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/netinsts/name/{name}][%d] edgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameOK  %+v", 200, o.Payload)
}

func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameOK) String() string {
	return fmt.Sprintf("[GET /v1/netinsts/name/{name}][%d] edgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameOK  %+v", 200, o.Payload)
}

func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameOK) GetPayload() *swagger_models.NetInstConfig {
	return o.Payload
}

func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.NetInstConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameUnauthorized creates a EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameUnauthorized with default headers values
func NewEdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameUnauthorized() *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameUnauthorized {
	return &EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameUnauthorized{}
}

/*
EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge network instance configuration get edge network instance by name unauthorized response has a 2xx status code
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network instance configuration get edge network instance by name unauthorized response has a 3xx status code
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network instance configuration get edge network instance by name unauthorized response has a 4xx status code
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge network instance configuration get edge network instance by name unauthorized response has a 5xx status code
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge network instance configuration get edge network instance by name unauthorized response a status code equal to that given
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the edge network instance configuration get edge network instance by name unauthorized response
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameUnauthorized) Code() int {
	return 401
}

func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/netinsts/name/{name}][%d] edgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/netinsts/name/{name}][%d] edgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameForbidden creates a EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameForbidden with default headers values
func NewEdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameForbidden() *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameForbidden {
	return &EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameForbidden{}
}

/*
EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge network instance configuration get edge network instance by name forbidden response has a 2xx status code
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network instance configuration get edge network instance by name forbidden response has a 3xx status code
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network instance configuration get edge network instance by name forbidden response has a 4xx status code
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge network instance configuration get edge network instance by name forbidden response has a 5xx status code
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge network instance configuration get edge network instance by name forbidden response a status code equal to that given
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the edge network instance configuration get edge network instance by name forbidden response
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameForbidden) Code() int {
	return 403
}

func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/netinsts/name/{name}][%d] edgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameForbidden) String() string {
	return fmt.Sprintf("[GET /v1/netinsts/name/{name}][%d] edgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameNotFound creates a EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameNotFound with default headers values
func NewEdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameNotFound() *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameNotFound {
	return &EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameNotFound{}
}

/*
EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge network instance configuration get edge network instance by name not found response has a 2xx status code
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network instance configuration get edge network instance by name not found response has a 3xx status code
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network instance configuration get edge network instance by name not found response has a 4xx status code
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge network instance configuration get edge network instance by name not found response has a 5xx status code
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge network instance configuration get edge network instance by name not found response a status code equal to that given
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the edge network instance configuration get edge network instance by name not found response
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameNotFound) Code() int {
	return 404
}

func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/netinsts/name/{name}][%d] edgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameNotFound) String() string {
	return fmt.Sprintf("[GET /v1/netinsts/name/{name}][%d] edgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameInternalServerError creates a EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameInternalServerError with default headers values
func NewEdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameInternalServerError() *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameInternalServerError {
	return &EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameInternalServerError{}
}

/*
EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge network instance configuration get edge network instance by name internal server error response has a 2xx status code
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network instance configuration get edge network instance by name internal server error response has a 3xx status code
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network instance configuration get edge network instance by name internal server error response has a 4xx status code
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge network instance configuration get edge network instance by name internal server error response has a 5xx status code
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge network instance configuration get edge network instance by name internal server error response a status code equal to that given
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the edge network instance configuration get edge network instance by name internal server error response
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameInternalServerError) Code() int {
	return 500
}

func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/netinsts/name/{name}][%d] edgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/netinsts/name/{name}][%d] edgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameGatewayTimeout creates a EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameGatewayTimeout with default headers values
func NewEdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameGatewayTimeout() *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameGatewayTimeout {
	return &EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameGatewayTimeout{}
}

/*
EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge network instance configuration get edge network instance by name gateway timeout response has a 2xx status code
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network instance configuration get edge network instance by name gateway timeout response has a 3xx status code
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network instance configuration get edge network instance by name gateway timeout response has a 4xx status code
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge network instance configuration get edge network instance by name gateway timeout response has a 5xx status code
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge network instance configuration get edge network instance by name gateway timeout response a status code equal to that given
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the edge network instance configuration get edge network instance by name gateway timeout response
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameGatewayTimeout) Code() int {
	return 504
}

func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/netinsts/name/{name}][%d] edgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/netinsts/name/{name}][%d] edgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameDefault creates a EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameDefault with default headers values
func NewEdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameDefault(code int) *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameDefault {
	return &EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameDefault{
		_statusCode: code,
	}
}

/*
EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// IsSuccess returns true when this edge network instance configuration get edge network instance by name default response has a 2xx status code
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge network instance configuration get edge network instance by name default response has a 3xx status code
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge network instance configuration get edge network instance by name default response has a 4xx status code
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge network instance configuration get edge network instance by name default response has a 5xx status code
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge network instance configuration get edge network instance by name default response a status code equal to that given
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the edge network instance configuration get edge network instance by name default response
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameDefault) Code() int {
	return o._statusCode
}

func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameDefault) Error() string {
	return fmt.Sprintf("[GET /v1/netinsts/name/{name}][%d] EdgeNetworkInstanceConfiguration_GetEdgeNetworkInstanceByName default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameDefault) String() string {
	return fmt.Sprintf("[GET /v1/netinsts/name/{name}][%d] EdgeNetworkInstanceConfiguration_GetEdgeNetworkInstanceByName default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

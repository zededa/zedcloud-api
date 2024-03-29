// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package edge_node_status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// EdgeNodeStatusGetEdgeNodeInfoByNameReader is a Reader for the EdgeNodeStatusGetEdgeNodeInfoByName structure.
type EdgeNodeStatusGetEdgeNodeInfoByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeNodeStatusGetEdgeNodeInfoByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeNodeStatusGetEdgeNodeInfoByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeNodeStatusGetEdgeNodeInfoByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeNodeStatusGetEdgeNodeInfoByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeNodeStatusGetEdgeNodeInfoByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeNodeStatusGetEdgeNodeInfoByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeNodeStatusGetEdgeNodeInfoByNameGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeNodeStatusGetEdgeNodeInfoByNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeNodeStatusGetEdgeNodeInfoByNameOK creates a EdgeNodeStatusGetEdgeNodeInfoByNameOK with default headers values
func NewEdgeNodeStatusGetEdgeNodeInfoByNameOK() *EdgeNodeStatusGetEdgeNodeInfoByNameOK {
	return &EdgeNodeStatusGetEdgeNodeInfoByNameOK{}
}

/*
EdgeNodeStatusGetEdgeNodeInfoByNameOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeNodeStatusGetEdgeNodeInfoByNameOK struct {
	Payload *swagger_models.DeviceInfoMsg
}

// IsSuccess returns true when this edge node status get edge node info by name o k response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge node status get edge node info by name o k response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node info by name o k response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node status get edge node info by name o k response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node status get edge node info by name o k response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeInfoByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the edge node status get edge node info by name o k response
func (o *EdgeNodeStatusGetEdgeNodeInfoByNameOK) Code() int {
	return 200
}

func (o *EdgeNodeStatusGetEdgeNodeInfoByNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status/info][%d] edgeNodeStatusGetEdgeNodeInfoByNameOK  %+v", 200, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeInfoByNameOK) String() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status/info][%d] edgeNodeStatusGetEdgeNodeInfoByNameOK  %+v", 200, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeInfoByNameOK) GetPayload() *swagger_models.DeviceInfoMsg {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeInfoByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.DeviceInfoMsg)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeInfoByNameUnauthorized creates a EdgeNodeStatusGetEdgeNodeInfoByNameUnauthorized with default headers values
func NewEdgeNodeStatusGetEdgeNodeInfoByNameUnauthorized() *EdgeNodeStatusGetEdgeNodeInfoByNameUnauthorized {
	return &EdgeNodeStatusGetEdgeNodeInfoByNameUnauthorized{}
}

/*
EdgeNodeStatusGetEdgeNodeInfoByNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeNodeStatusGetEdgeNodeInfoByNameUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge node status get edge node info by name unauthorized response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoByNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node status get edge node info by name unauthorized response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoByNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node info by name unauthorized response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoByNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node status get edge node info by name unauthorized response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoByNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node status get edge node info by name unauthorized response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeInfoByNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the edge node status get edge node info by name unauthorized response
func (o *EdgeNodeStatusGetEdgeNodeInfoByNameUnauthorized) Code() int {
	return 401
}

func (o *EdgeNodeStatusGetEdgeNodeInfoByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status/info][%d] edgeNodeStatusGetEdgeNodeInfoByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeInfoByNameUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status/info][%d] edgeNodeStatusGetEdgeNodeInfoByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeInfoByNameUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeInfoByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeInfoByNameForbidden creates a EdgeNodeStatusGetEdgeNodeInfoByNameForbidden with default headers values
func NewEdgeNodeStatusGetEdgeNodeInfoByNameForbidden() *EdgeNodeStatusGetEdgeNodeInfoByNameForbidden {
	return &EdgeNodeStatusGetEdgeNodeInfoByNameForbidden{}
}

/*
EdgeNodeStatusGetEdgeNodeInfoByNameForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeNodeStatusGetEdgeNodeInfoByNameForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge node status get edge node info by name forbidden response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoByNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node status get edge node info by name forbidden response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoByNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node info by name forbidden response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoByNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node status get edge node info by name forbidden response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoByNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node status get edge node info by name forbidden response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeInfoByNameForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the edge node status get edge node info by name forbidden response
func (o *EdgeNodeStatusGetEdgeNodeInfoByNameForbidden) Code() int {
	return 403
}

func (o *EdgeNodeStatusGetEdgeNodeInfoByNameForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status/info][%d] edgeNodeStatusGetEdgeNodeInfoByNameForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeInfoByNameForbidden) String() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status/info][%d] edgeNodeStatusGetEdgeNodeInfoByNameForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeInfoByNameForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeInfoByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeInfoByNameNotFound creates a EdgeNodeStatusGetEdgeNodeInfoByNameNotFound with default headers values
func NewEdgeNodeStatusGetEdgeNodeInfoByNameNotFound() *EdgeNodeStatusGetEdgeNodeInfoByNameNotFound {
	return &EdgeNodeStatusGetEdgeNodeInfoByNameNotFound{}
}

/*
EdgeNodeStatusGetEdgeNodeInfoByNameNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeNodeStatusGetEdgeNodeInfoByNameNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge node status get edge node info by name not found response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoByNameNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node status get edge node info by name not found response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoByNameNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node info by name not found response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoByNameNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node status get edge node info by name not found response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoByNameNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node status get edge node info by name not found response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeInfoByNameNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the edge node status get edge node info by name not found response
func (o *EdgeNodeStatusGetEdgeNodeInfoByNameNotFound) Code() int {
	return 404
}

func (o *EdgeNodeStatusGetEdgeNodeInfoByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status/info][%d] edgeNodeStatusGetEdgeNodeInfoByNameNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeInfoByNameNotFound) String() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status/info][%d] edgeNodeStatusGetEdgeNodeInfoByNameNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeInfoByNameNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeInfoByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeInfoByNameInternalServerError creates a EdgeNodeStatusGetEdgeNodeInfoByNameInternalServerError with default headers values
func NewEdgeNodeStatusGetEdgeNodeInfoByNameInternalServerError() *EdgeNodeStatusGetEdgeNodeInfoByNameInternalServerError {
	return &EdgeNodeStatusGetEdgeNodeInfoByNameInternalServerError{}
}

/*
EdgeNodeStatusGetEdgeNodeInfoByNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeNodeStatusGetEdgeNodeInfoByNameInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge node status get edge node info by name internal server error response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoByNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node status get edge node info by name internal server error response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoByNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node info by name internal server error response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoByNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node status get edge node info by name internal server error response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoByNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node status get edge node info by name internal server error response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeInfoByNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the edge node status get edge node info by name internal server error response
func (o *EdgeNodeStatusGetEdgeNodeInfoByNameInternalServerError) Code() int {
	return 500
}

func (o *EdgeNodeStatusGetEdgeNodeInfoByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status/info][%d] edgeNodeStatusGetEdgeNodeInfoByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeInfoByNameInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status/info][%d] edgeNodeStatusGetEdgeNodeInfoByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeInfoByNameInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeInfoByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeInfoByNameGatewayTimeout creates a EdgeNodeStatusGetEdgeNodeInfoByNameGatewayTimeout with default headers values
func NewEdgeNodeStatusGetEdgeNodeInfoByNameGatewayTimeout() *EdgeNodeStatusGetEdgeNodeInfoByNameGatewayTimeout {
	return &EdgeNodeStatusGetEdgeNodeInfoByNameGatewayTimeout{}
}

/*
EdgeNodeStatusGetEdgeNodeInfoByNameGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeNodeStatusGetEdgeNodeInfoByNameGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge node status get edge node info by name gateway timeout response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoByNameGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node status get edge node info by name gateway timeout response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoByNameGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node info by name gateway timeout response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoByNameGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node status get edge node info by name gateway timeout response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoByNameGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node status get edge node info by name gateway timeout response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeInfoByNameGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the edge node status get edge node info by name gateway timeout response
func (o *EdgeNodeStatusGetEdgeNodeInfoByNameGatewayTimeout) Code() int {
	return 504
}

func (o *EdgeNodeStatusGetEdgeNodeInfoByNameGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status/info][%d] edgeNodeStatusGetEdgeNodeInfoByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeInfoByNameGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status/info][%d] edgeNodeStatusGetEdgeNodeInfoByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeInfoByNameGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeInfoByNameGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeInfoByNameDefault creates a EdgeNodeStatusGetEdgeNodeInfoByNameDefault with default headers values
func NewEdgeNodeStatusGetEdgeNodeInfoByNameDefault(code int) *EdgeNodeStatusGetEdgeNodeInfoByNameDefault {
	return &EdgeNodeStatusGetEdgeNodeInfoByNameDefault{
		_statusCode: code,
	}
}

/*
EdgeNodeStatusGetEdgeNodeInfoByNameDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeNodeStatusGetEdgeNodeInfoByNameDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// IsSuccess returns true when this edge node status get edge node info by name default response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoByNameDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge node status get edge node info by name default response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoByNameDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge node status get edge node info by name default response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoByNameDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge node status get edge node info by name default response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoByNameDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge node status get edge node info by name default response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeInfoByNameDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the edge node status get edge node info by name default response
func (o *EdgeNodeStatusGetEdgeNodeInfoByNameDefault) Code() int {
	return o._statusCode
}

func (o *EdgeNodeStatusGetEdgeNodeInfoByNameDefault) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status/info][%d] EdgeNodeStatus_GetEdgeNodeInfoByName default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeInfoByNameDefault) String() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status/info][%d] EdgeNodeStatus_GetEdgeNodeInfoByName default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeInfoByNameDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeInfoByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

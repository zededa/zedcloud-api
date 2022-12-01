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

// EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameReader is a Reader for the EdgeNodeStatusGetEdgeNodeEdgeviewStatusByName structure.
type EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameOK creates a EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameOK with default headers values
func NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameOK() *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameOK {
	return &EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameOK{}
}

/*
EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameOK struct {
	Payload *swagger_models.DeviceRawMetrics
}

// IsSuccess returns true when this edge node status get edge node edgeview status by name o k response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge node status get edge node edgeview status by name o k response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node edgeview status by name o k response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node status get edge node edgeview status by name o k response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node status get edge node edgeview status by name o k response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameOK) IsCode(code int) bool {
	return code == 200
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status/edgeview][%d] edgeNodeStatusGetEdgeNodeEdgeviewStatusByNameOK  %+v", 200, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameOK) String() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status/edgeview][%d] edgeNodeStatusGetEdgeNodeEdgeviewStatusByNameOK  %+v", 200, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameOK) GetPayload() *swagger_models.DeviceRawMetrics {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.DeviceRawMetrics)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameUnauthorized creates a EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameUnauthorized with default headers values
func NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameUnauthorized() *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameUnauthorized {
	return &EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameUnauthorized{}
}

/*
EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge node status get edge node edgeview status by name unauthorized response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node status get edge node edgeview status by name unauthorized response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node edgeview status by name unauthorized response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node status get edge node edgeview status by name unauthorized response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node status get edge node edgeview status by name unauthorized response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status/edgeview][%d] edgeNodeStatusGetEdgeNodeEdgeviewStatusByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status/edgeview][%d] edgeNodeStatusGetEdgeNodeEdgeviewStatusByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameForbidden creates a EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameForbidden with default headers values
func NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameForbidden() *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameForbidden {
	return &EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameForbidden{}
}

/*
EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge node status get edge node edgeview status by name forbidden response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node status get edge node edgeview status by name forbidden response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node edgeview status by name forbidden response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node status get edge node edgeview status by name forbidden response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node status get edge node edgeview status by name forbidden response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status/edgeview][%d] edgeNodeStatusGetEdgeNodeEdgeviewStatusByNameForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameForbidden) String() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status/edgeview][%d] edgeNodeStatusGetEdgeNodeEdgeviewStatusByNameForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameNotFound creates a EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameNotFound with default headers values
func NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameNotFound() *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameNotFound {
	return &EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameNotFound{}
}

/*
EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge node status get edge node edgeview status by name not found response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node status get edge node edgeview status by name not found response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node edgeview status by name not found response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node status get edge node edgeview status by name not found response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node status get edge node edgeview status by name not found response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status/edgeview][%d] edgeNodeStatusGetEdgeNodeEdgeviewStatusByNameNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameNotFound) String() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status/edgeview][%d] edgeNodeStatusGetEdgeNodeEdgeviewStatusByNameNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameInternalServerError creates a EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameInternalServerError with default headers values
func NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameInternalServerError() *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameInternalServerError {
	return &EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameInternalServerError{}
}

/*
EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge node status get edge node edgeview status by name internal server error response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node status get edge node edgeview status by name internal server error response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node edgeview status by name internal server error response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node status get edge node edgeview status by name internal server error response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node status get edge node edgeview status by name internal server error response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status/edgeview][%d] edgeNodeStatusGetEdgeNodeEdgeviewStatusByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status/edgeview][%d] edgeNodeStatusGetEdgeNodeEdgeviewStatusByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameGatewayTimeout creates a EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameGatewayTimeout with default headers values
func NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameGatewayTimeout() *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameGatewayTimeout {
	return &EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameGatewayTimeout{}
}

/*
EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge node status get edge node edgeview status by name gateway timeout response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node status get edge node edgeview status by name gateway timeout response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node edgeview status by name gateway timeout response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node status get edge node edgeview status by name gateway timeout response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node status get edge node edgeview status by name gateway timeout response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status/edgeview][%d] edgeNodeStatusGetEdgeNodeEdgeviewStatusByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status/edgeview][%d] edgeNodeStatusGetEdgeNodeEdgeviewStatusByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameDefault creates a EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameDefault with default headers values
func NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameDefault(code int) *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameDefault {
	return &EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameDefault{
		_statusCode: code,
	}
}

/*
EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the edge node status get edge node edgeview status by name default response
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this edge node status get edge node edgeview status by name default response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge node status get edge node edgeview status by name default response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge node status get edge node edgeview status by name default response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge node status get edge node edgeview status by name default response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge node status get edge node edgeview status by name default response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameDefault) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status/edgeview][%d] EdgeNodeStatus_GetEdgeNodeEdgeviewStatusByName default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameDefault) String() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status/edgeview][%d] EdgeNodeStatus_GetEdgeNodeEdgeviewStatusByName default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

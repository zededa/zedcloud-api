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

// EdgeNodeStatusGetEdgeNodeInfoReader is a Reader for the EdgeNodeStatusGetEdgeNodeInfo structure.
type EdgeNodeStatusGetEdgeNodeInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeNodeStatusGetEdgeNodeInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeNodeStatusGetEdgeNodeInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeNodeStatusGetEdgeNodeInfoUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeNodeStatusGetEdgeNodeInfoForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeNodeStatusGetEdgeNodeInfoNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeNodeStatusGetEdgeNodeInfoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeNodeStatusGetEdgeNodeInfoGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeNodeStatusGetEdgeNodeInfoDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeNodeStatusGetEdgeNodeInfoOK creates a EdgeNodeStatusGetEdgeNodeInfoOK with default headers values
func NewEdgeNodeStatusGetEdgeNodeInfoOK() *EdgeNodeStatusGetEdgeNodeInfoOK {
	return &EdgeNodeStatusGetEdgeNodeInfoOK{}
}

/*
EdgeNodeStatusGetEdgeNodeInfoOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeNodeStatusGetEdgeNodeInfoOK struct {
	Payload *swagger_models.DeviceInfoMsg
}

// IsSuccess returns true when this edge node status get edge node info o k response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge node status get edge node info o k response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node info o k response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node status get edge node info o k response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node status get edge node info o k response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeInfoOK) IsCode(code int) bool {
	return code == 200
}

func (o *EdgeNodeStatusGetEdgeNodeInfoOK) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/status/info][%d] edgeNodeStatusGetEdgeNodeInfoOK  %+v", 200, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeInfoOK) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/status/info][%d] edgeNodeStatusGetEdgeNodeInfoOK  %+v", 200, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeInfoOK) GetPayload() *swagger_models.DeviceInfoMsg {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.DeviceInfoMsg)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeInfoUnauthorized creates a EdgeNodeStatusGetEdgeNodeInfoUnauthorized with default headers values
func NewEdgeNodeStatusGetEdgeNodeInfoUnauthorized() *EdgeNodeStatusGetEdgeNodeInfoUnauthorized {
	return &EdgeNodeStatusGetEdgeNodeInfoUnauthorized{}
}

/*
EdgeNodeStatusGetEdgeNodeInfoUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeNodeStatusGetEdgeNodeInfoUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge node status get edge node info unauthorized response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node status get edge node info unauthorized response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node info unauthorized response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node status get edge node info unauthorized response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node status get edge node info unauthorized response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeInfoUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *EdgeNodeStatusGetEdgeNodeInfoUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/status/info][%d] edgeNodeStatusGetEdgeNodeInfoUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeInfoUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/status/info][%d] edgeNodeStatusGetEdgeNodeInfoUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeInfoUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeInfoUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeInfoForbidden creates a EdgeNodeStatusGetEdgeNodeInfoForbidden with default headers values
func NewEdgeNodeStatusGetEdgeNodeInfoForbidden() *EdgeNodeStatusGetEdgeNodeInfoForbidden {
	return &EdgeNodeStatusGetEdgeNodeInfoForbidden{}
}

/*
EdgeNodeStatusGetEdgeNodeInfoForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeNodeStatusGetEdgeNodeInfoForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge node status get edge node info forbidden response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node status get edge node info forbidden response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node info forbidden response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node status get edge node info forbidden response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node status get edge node info forbidden response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeInfoForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *EdgeNodeStatusGetEdgeNodeInfoForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/status/info][%d] edgeNodeStatusGetEdgeNodeInfoForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeInfoForbidden) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/status/info][%d] edgeNodeStatusGetEdgeNodeInfoForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeInfoForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeInfoForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeInfoNotFound creates a EdgeNodeStatusGetEdgeNodeInfoNotFound with default headers values
func NewEdgeNodeStatusGetEdgeNodeInfoNotFound() *EdgeNodeStatusGetEdgeNodeInfoNotFound {
	return &EdgeNodeStatusGetEdgeNodeInfoNotFound{}
}

/*
EdgeNodeStatusGetEdgeNodeInfoNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeNodeStatusGetEdgeNodeInfoNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge node status get edge node info not found response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node status get edge node info not found response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node info not found response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node status get edge node info not found response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node status get edge node info not found response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeInfoNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *EdgeNodeStatusGetEdgeNodeInfoNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/status/info][%d] edgeNodeStatusGetEdgeNodeInfoNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeInfoNotFound) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/status/info][%d] edgeNodeStatusGetEdgeNodeInfoNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeInfoNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeInfoNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeInfoInternalServerError creates a EdgeNodeStatusGetEdgeNodeInfoInternalServerError with default headers values
func NewEdgeNodeStatusGetEdgeNodeInfoInternalServerError() *EdgeNodeStatusGetEdgeNodeInfoInternalServerError {
	return &EdgeNodeStatusGetEdgeNodeInfoInternalServerError{}
}

/*
EdgeNodeStatusGetEdgeNodeInfoInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeNodeStatusGetEdgeNodeInfoInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge node status get edge node info internal server error response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node status get edge node info internal server error response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node info internal server error response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node status get edge node info internal server error response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node status get edge node info internal server error response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeInfoInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EdgeNodeStatusGetEdgeNodeInfoInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/status/info][%d] edgeNodeStatusGetEdgeNodeInfoInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeInfoInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/status/info][%d] edgeNodeStatusGetEdgeNodeInfoInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeInfoInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeInfoInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeInfoGatewayTimeout creates a EdgeNodeStatusGetEdgeNodeInfoGatewayTimeout with default headers values
func NewEdgeNodeStatusGetEdgeNodeInfoGatewayTimeout() *EdgeNodeStatusGetEdgeNodeInfoGatewayTimeout {
	return &EdgeNodeStatusGetEdgeNodeInfoGatewayTimeout{}
}

/*
EdgeNodeStatusGetEdgeNodeInfoGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeNodeStatusGetEdgeNodeInfoGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge node status get edge node info gateway timeout response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node status get edge node info gateway timeout response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node info gateway timeout response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node status get edge node info gateway timeout response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node status get edge node info gateway timeout response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeInfoGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *EdgeNodeStatusGetEdgeNodeInfoGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/status/info][%d] edgeNodeStatusGetEdgeNodeInfoGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeInfoGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/status/info][%d] edgeNodeStatusGetEdgeNodeInfoGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeInfoGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeInfoGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeInfoDefault creates a EdgeNodeStatusGetEdgeNodeInfoDefault with default headers values
func NewEdgeNodeStatusGetEdgeNodeInfoDefault(code int) *EdgeNodeStatusGetEdgeNodeInfoDefault {
	return &EdgeNodeStatusGetEdgeNodeInfoDefault{
		_statusCode: code,
	}
}

/*
EdgeNodeStatusGetEdgeNodeInfoDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeNodeStatusGetEdgeNodeInfoDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the edge node status get edge node info default response
func (o *EdgeNodeStatusGetEdgeNodeInfoDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this edge node status get edge node info default response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge node status get edge node info default response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge node status get edge node info default response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge node status get edge node info default response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeInfoDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge node status get edge node info default response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeInfoDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *EdgeNodeStatusGetEdgeNodeInfoDefault) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/status/info][%d] EdgeNodeStatus_GetEdgeNodeInfo default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeInfoDefault) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/status/info][%d] EdgeNodeStatus_GetEdgeNodeInfo default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeInfoDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeInfoDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package edge_network_instance_status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigReader is a Reader for the EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfig structure.
type EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewEdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigOK creates a EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigOK with default headers values
func NewEdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigOK() *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigOK {
	return &EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigOK{}
}

/*
EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigOK struct {
	Payload *swagger_models.NetInstConfigStatusList
}

// IsSuccess returns true when this edge network instance status query edge network instance status config o k response has a 2xx status code
func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge network instance status query edge network instance status config o k response has a 3xx status code
func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network instance status query edge network instance status config o k response has a 4xx status code
func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge network instance status query edge network instance status config o k response has a 5xx status code
func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge network instance status query edge network instance status config o k response a status code equal to that given
func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigOK) IsCode(code int) bool {
	return code == 200
}

func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigOK) Error() string {
	return fmt.Sprintf("[GET /v1/netinsts/status-config][%d] edgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigOK  %+v", 200, o.Payload)
}

func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigOK) String() string {
	return fmt.Sprintf("[GET /v1/netinsts/status-config][%d] edgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigOK  %+v", 200, o.Payload)
}

func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigOK) GetPayload() *swagger_models.NetInstConfigStatusList {
	return o.Payload
}

func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.NetInstConfigStatusList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigBadRequest creates a EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigBadRequest with default headers values
func NewEdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigBadRequest() *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigBadRequest {
	return &EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigBadRequest{}
}

/*
EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of invalid value of filter parameters.
*/
type EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge network instance status query edge network instance status config bad request response has a 2xx status code
func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network instance status query edge network instance status config bad request response has a 3xx status code
func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network instance status query edge network instance status config bad request response has a 4xx status code
func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge network instance status query edge network instance status config bad request response has a 5xx status code
func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this edge network instance status query edge network instance status config bad request response a status code equal to that given
func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/netinsts/status-config][%d] edgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigBadRequest  %+v", 400, o.Payload)
}

func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/netinsts/status-config][%d] edgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigBadRequest  %+v", 400, o.Payload)
}

func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigUnauthorized creates a EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigUnauthorized with default headers values
func NewEdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigUnauthorized() *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigUnauthorized {
	return &EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigUnauthorized{}
}

/*
EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge network instance status query edge network instance status config unauthorized response has a 2xx status code
func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network instance status query edge network instance status config unauthorized response has a 3xx status code
func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network instance status query edge network instance status config unauthorized response has a 4xx status code
func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge network instance status query edge network instance status config unauthorized response has a 5xx status code
func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge network instance status query edge network instance status config unauthorized response a status code equal to that given
func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/netinsts/status-config][%d] edgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/netinsts/status-config][%d] edgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigForbidden creates a EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigForbidden with default headers values
func NewEdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigForbidden() *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigForbidden {
	return &EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigForbidden{}
}

/*
EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge network instance status query edge network instance status config forbidden response has a 2xx status code
func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network instance status query edge network instance status config forbidden response has a 3xx status code
func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network instance status query edge network instance status config forbidden response has a 4xx status code
func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge network instance status query edge network instance status config forbidden response has a 5xx status code
func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge network instance status query edge network instance status config forbidden response a status code equal to that given
func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/netinsts/status-config][%d] edgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigForbidden) String() string {
	return fmt.Sprintf("[GET /v1/netinsts/status-config][%d] edgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigInternalServerError creates a EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigInternalServerError with default headers values
func NewEdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigInternalServerError() *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigInternalServerError {
	return &EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigInternalServerError{}
}

/*
EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge network instance status query edge network instance status config internal server error response has a 2xx status code
func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network instance status query edge network instance status config internal server error response has a 3xx status code
func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network instance status query edge network instance status config internal server error response has a 4xx status code
func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge network instance status query edge network instance status config internal server error response has a 5xx status code
func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge network instance status query edge network instance status config internal server error response a status code equal to that given
func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/netinsts/status-config][%d] edgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/netinsts/status-config][%d] edgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigGatewayTimeout creates a EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigGatewayTimeout with default headers values
func NewEdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigGatewayTimeout() *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigGatewayTimeout {
	return &EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigGatewayTimeout{}
}

/*
EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge network instance status query edge network instance status config gateway timeout response has a 2xx status code
func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network instance status query edge network instance status config gateway timeout response has a 3xx status code
func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network instance status query edge network instance status config gateway timeout response has a 4xx status code
func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge network instance status query edge network instance status config gateway timeout response has a 5xx status code
func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge network instance status query edge network instance status config gateway timeout response a status code equal to that given
func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/netinsts/status-config][%d] edgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/netinsts/status-config][%d] edgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigDefault creates a EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigDefault with default headers values
func NewEdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigDefault(code int) *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigDefault {
	return &EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigDefault{
		_statusCode: code,
	}
}

/*
EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the edge network instance status query edge network instance status config default response
func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this edge network instance status query edge network instance status config default response has a 2xx status code
func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge network instance status query edge network instance status config default response has a 3xx status code
func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge network instance status query edge network instance status config default response has a 4xx status code
func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge network instance status query edge network instance status config default response has a 5xx status code
func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge network instance status query edge network instance status config default response a status code equal to that given
func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigDefault) Error() string {
	return fmt.Sprintf("[GET /v1/netinsts/status-config][%d] EdgeNetworkInstanceStatus_QueryEdgeNetworkInstanceStatusConfig default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigDefault) String() string {
	return fmt.Sprintf("[GET /v1/netinsts/status-config][%d] EdgeNetworkInstanceStatus_QueryEdgeNetworkInstanceStatusConfig default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeNetworkInstanceStatusQueryEdgeNetworkInstanceStatusConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

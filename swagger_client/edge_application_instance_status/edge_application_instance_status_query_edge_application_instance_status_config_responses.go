// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package edge_application_instance_status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigReader is a Reader for the EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfig structure.
type EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewEdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigOK creates a EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigOK with default headers values
func NewEdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigOK() *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigOK {
	return &EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigOK{}
}

/*
EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigOK struct {
	Payload *swagger_models.AppInstConfigStatusList
}

// IsSuccess returns true when this edge application instance status query edge application instance status config o k response has a 2xx status code
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge application instance status query edge application instance status config o k response has a 3xx status code
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance status query edge application instance status config o k response has a 4xx status code
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application instance status query edge application instance status config o k response has a 5xx status code
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application instance status query edge application instance status config o k response a status code equal to that given
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the edge application instance status query edge application instance status config o k response
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigOK) Code() int {
	return 200
}

func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigOK) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/status-config][%d] edgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigOK  %+v", 200, o.Payload)
}

func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigOK) String() string {
	return fmt.Sprintf("[GET /v1/apps/instances/status-config][%d] edgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigOK  %+v", 200, o.Payload)
}

func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigOK) GetPayload() *swagger_models.AppInstConfigStatusList {
	return o.Payload
}

func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.AppInstConfigStatusList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigBadRequest creates a EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigBadRequest with default headers values
func NewEdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigBadRequest() *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigBadRequest {
	return &EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigBadRequest{}
}

/*
EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of invalid value of filter parameters.
*/
type EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance status query edge application instance status config bad request response has a 2xx status code
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance status query edge application instance status config bad request response has a 3xx status code
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance status query edge application instance status config bad request response has a 4xx status code
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application instance status query edge application instance status config bad request response has a 5xx status code
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application instance status query edge application instance status config bad request response a status code equal to that given
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the edge application instance status query edge application instance status config bad request response
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigBadRequest) Code() int {
	return 400
}

func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/status-config][%d] edgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigBadRequest  %+v", 400, o.Payload)
}

func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/apps/instances/status-config][%d] edgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigBadRequest  %+v", 400, o.Payload)
}

func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigUnauthorized creates a EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigUnauthorized with default headers values
func NewEdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigUnauthorized() *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigUnauthorized {
	return &EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigUnauthorized{}
}

/*
EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance status query edge application instance status config unauthorized response has a 2xx status code
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance status query edge application instance status config unauthorized response has a 3xx status code
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance status query edge application instance status config unauthorized response has a 4xx status code
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application instance status query edge application instance status config unauthorized response has a 5xx status code
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application instance status query edge application instance status config unauthorized response a status code equal to that given
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the edge application instance status query edge application instance status config unauthorized response
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigUnauthorized) Code() int {
	return 401
}

func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/status-config][%d] edgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/apps/instances/status-config][%d] edgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigForbidden creates a EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigForbidden with default headers values
func NewEdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigForbidden() *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigForbidden {
	return &EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigForbidden{}
}

/*
EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have application level access permission for the operation or does not have access scope to the project.
*/
type EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance status query edge application instance status config forbidden response has a 2xx status code
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance status query edge application instance status config forbidden response has a 3xx status code
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance status query edge application instance status config forbidden response has a 4xx status code
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application instance status query edge application instance status config forbidden response has a 5xx status code
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application instance status query edge application instance status config forbidden response a status code equal to that given
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the edge application instance status query edge application instance status config forbidden response
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigForbidden) Code() int {
	return 403
}

func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/status-config][%d] edgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigForbidden  %+v", 403, o.Payload)
}

func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigForbidden) String() string {
	return fmt.Sprintf("[GET /v1/apps/instances/status-config][%d] edgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigForbidden  %+v", 403, o.Payload)
}

func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigInternalServerError creates a EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigInternalServerError with default headers values
func NewEdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigInternalServerError() *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigInternalServerError {
	return &EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigInternalServerError{}
}

/*
EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance status query edge application instance status config internal server error response has a 2xx status code
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance status query edge application instance status config internal server error response has a 3xx status code
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance status query edge application instance status config internal server error response has a 4xx status code
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application instance status query edge application instance status config internal server error response has a 5xx status code
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge application instance status query edge application instance status config internal server error response a status code equal to that given
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the edge application instance status query edge application instance status config internal server error response
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigInternalServerError) Code() int {
	return 500
}

func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/status-config][%d] edgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/apps/instances/status-config][%d] edgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigGatewayTimeout creates a EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigGatewayTimeout with default headers values
func NewEdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigGatewayTimeout() *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigGatewayTimeout {
	return &EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigGatewayTimeout{}
}

/*
EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance status query edge application instance status config gateway timeout response has a 2xx status code
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance status query edge application instance status config gateway timeout response has a 3xx status code
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance status query edge application instance status config gateway timeout response has a 4xx status code
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application instance status query edge application instance status config gateway timeout response has a 5xx status code
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge application instance status query edge application instance status config gateway timeout response a status code equal to that given
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the edge application instance status query edge application instance status config gateway timeout response
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigGatewayTimeout) Code() int {
	return 504
}

func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/status-config][%d] edgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/apps/instances/status-config][%d] edgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigDefault creates a EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigDefault with default headers values
func NewEdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigDefault(code int) *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigDefault {
	return &EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigDefault{
		_statusCode: code,
	}
}

/*
EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// IsSuccess returns true when this edge application instance status query edge application instance status config default response has a 2xx status code
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge application instance status query edge application instance status config default response has a 3xx status code
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge application instance status query edge application instance status config default response has a 4xx status code
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge application instance status query edge application instance status config default response has a 5xx status code
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge application instance status query edge application instance status config default response a status code equal to that given
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the edge application instance status query edge application instance status config default response
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigDefault) Code() int {
	return o._statusCode
}

func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigDefault) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/status-config][%d] EdgeApplicationInstanceStatus_QueryEdgeApplicationInstanceStatusConfig default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigDefault) String() string {
	return fmt.Sprintf("[GET /v1/apps/instances/status-config][%d] EdgeApplicationInstanceStatus_QueryEdgeApplicationInstanceStatusConfig default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

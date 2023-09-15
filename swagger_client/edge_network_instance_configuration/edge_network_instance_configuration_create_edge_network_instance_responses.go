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

// EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceReader is a Reader for the EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstance structure.
type EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewEdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewEdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceOK creates a EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceOK with default headers values
func NewEdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceOK() *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceOK {
	return &EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceOK{}
}

/*
EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceOK struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge network instance configuration create edge network instance o k response has a 2xx status code
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge network instance configuration create edge network instance o k response has a 3xx status code
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network instance configuration create edge network instance o k response has a 4xx status code
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge network instance configuration create edge network instance o k response has a 5xx status code
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge network instance configuration create edge network instance o k response a status code equal to that given
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the edge network instance configuration create edge network instance o k response
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceOK) Code() int {
	return 200
}

func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceOK) Error() string {
	return fmt.Sprintf("[POST /v1/netinsts][%d] edgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceOK  %+v", 200, o.Payload)
}

func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceOK) String() string {
	return fmt.Sprintf("[POST /v1/netinsts][%d] edgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceOK  %+v", 200, o.Payload)
}

func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceOK) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceBadRequest creates a EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceBadRequest with default headers values
func NewEdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceBadRequest() *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceBadRequest {
	return &EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceBadRequest{}
}

/*
EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge network instance configuration create edge network instance bad request response has a 2xx status code
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network instance configuration create edge network instance bad request response has a 3xx status code
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network instance configuration create edge network instance bad request response has a 4xx status code
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge network instance configuration create edge network instance bad request response has a 5xx status code
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this edge network instance configuration create edge network instance bad request response a status code equal to that given
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the edge network instance configuration create edge network instance bad request response
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceBadRequest) Code() int {
	return 400
}

func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/netinsts][%d] edgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceBadRequest  %+v", 400, o.Payload)
}

func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/netinsts][%d] edgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceBadRequest  %+v", 400, o.Payload)
}

func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceUnauthorized creates a EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceUnauthorized with default headers values
func NewEdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceUnauthorized() *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceUnauthorized {
	return &EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceUnauthorized{}
}

/*
EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge network instance configuration create edge network instance unauthorized response has a 2xx status code
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network instance configuration create edge network instance unauthorized response has a 3xx status code
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network instance configuration create edge network instance unauthorized response has a 4xx status code
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge network instance configuration create edge network instance unauthorized response has a 5xx status code
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge network instance configuration create edge network instance unauthorized response a status code equal to that given
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the edge network instance configuration create edge network instance unauthorized response
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceUnauthorized) Code() int {
	return 401
}

func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/netinsts][%d] edgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/netinsts][%d] edgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceForbidden creates a EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceForbidden with default headers values
func NewEdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceForbidden() *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceForbidden {
	return &EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceForbidden{}
}

/*
EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge network instance configuration create edge network instance forbidden response has a 2xx status code
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network instance configuration create edge network instance forbidden response has a 3xx status code
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network instance configuration create edge network instance forbidden response has a 4xx status code
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge network instance configuration create edge network instance forbidden response has a 5xx status code
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge network instance configuration create edge network instance forbidden response a status code equal to that given
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the edge network instance configuration create edge network instance forbidden response
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceForbidden) Code() int {
	return 403
}

func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/netinsts][%d] edgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceForbidden) String() string {
	return fmt.Sprintf("[POST /v1/netinsts][%d] edgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceConflict creates a EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceConflict with default headers values
func NewEdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceConflict() *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceConflict {
	return &EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceConflict{}
}

/*
EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this edge network record will conflict with an already existing edge network record.
*/
type EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceConflict struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge network instance configuration create edge network instance conflict response has a 2xx status code
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network instance configuration create edge network instance conflict response has a 3xx status code
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network instance configuration create edge network instance conflict response has a 4xx status code
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge network instance configuration create edge network instance conflict response has a 5xx status code
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this edge network instance configuration create edge network instance conflict response a status code equal to that given
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the edge network instance configuration create edge network instance conflict response
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceConflict) Code() int {
	return 409
}

func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceConflict) Error() string {
	return fmt.Sprintf("[POST /v1/netinsts][%d] edgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceConflict  %+v", 409, o.Payload)
}

func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceConflict) String() string {
	return fmt.Sprintf("[POST /v1/netinsts][%d] edgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceConflict  %+v", 409, o.Payload)
}

func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceConflict) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceInternalServerError creates a EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceInternalServerError with default headers values
func NewEdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceInternalServerError() *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceInternalServerError {
	return &EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceInternalServerError{}
}

/*
EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge network instance configuration create edge network instance internal server error response has a 2xx status code
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network instance configuration create edge network instance internal server error response has a 3xx status code
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network instance configuration create edge network instance internal server error response has a 4xx status code
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge network instance configuration create edge network instance internal server error response has a 5xx status code
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge network instance configuration create edge network instance internal server error response a status code equal to that given
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the edge network instance configuration create edge network instance internal server error response
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceInternalServerError) Code() int {
	return 500
}

func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/netinsts][%d] edgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/netinsts][%d] edgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceGatewayTimeout creates a EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceGatewayTimeout with default headers values
func NewEdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceGatewayTimeout() *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceGatewayTimeout {
	return &EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceGatewayTimeout{}
}

/*
EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge network instance configuration create edge network instance gateway timeout response has a 2xx status code
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network instance configuration create edge network instance gateway timeout response has a 3xx status code
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network instance configuration create edge network instance gateway timeout response has a 4xx status code
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge network instance configuration create edge network instance gateway timeout response has a 5xx status code
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge network instance configuration create edge network instance gateway timeout response a status code equal to that given
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the edge network instance configuration create edge network instance gateway timeout response
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceGatewayTimeout) Code() int {
	return 504
}

func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v1/netinsts][%d] edgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /v1/netinsts][%d] edgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceDefault creates a EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceDefault with default headers values
func NewEdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceDefault(code int) *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceDefault {
	return &EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceDefault{
		_statusCode: code,
	}
}

/*
EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// IsSuccess returns true when this edge network instance configuration create edge network instance default response has a 2xx status code
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge network instance configuration create edge network instance default response has a 3xx status code
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge network instance configuration create edge network instance default response has a 4xx status code
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge network instance configuration create edge network instance default response has a 5xx status code
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge network instance configuration create edge network instance default response a status code equal to that given
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the edge network instance configuration create edge network instance default response
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceDefault) Code() int {
	return o._statusCode
}

func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceDefault) Error() string {
	return fmt.Sprintf("[POST /v1/netinsts][%d] EdgeNetworkInstanceConfiguration_CreateEdgeNetworkInstance default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceDefault) String() string {
	return fmt.Sprintf("[POST /v1/netinsts][%d] EdgeNetworkInstanceConfiguration_CreateEdgeNetworkInstance default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

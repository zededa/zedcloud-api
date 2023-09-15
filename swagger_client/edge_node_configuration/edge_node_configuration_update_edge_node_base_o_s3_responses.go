// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package edge_node_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Reader is a Reader for the EdgeNodeConfigurationUpdateEdgeNodeBaseOS3 structure.
type EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS3Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS3NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS3Conflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS3InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS3GatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS3Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS3OK creates a EdgeNodeConfigurationUpdateEdgeNodeBaseOS3OK with default headers values
func NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS3OK() *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3OK {
	return &EdgeNodeConfigurationUpdateEdgeNodeBaseOS3OK{}
}

/*
EdgeNodeConfigurationUpdateEdgeNodeBaseOS3OK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeNodeConfigurationUpdateEdgeNodeBaseOS3OK struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration update edge node base o s3 o k response has a 2xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge node configuration update edge node base o s3 o k response has a 3xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration update edge node base o s3 o k response has a 4xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration update edge node base o s3 o k response has a 5xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3OK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration update edge node base o s3 o k response a status code equal to that given
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the edge node configuration update edge node base o s3 o k response
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3OK) Code() int {
	return 200
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3OK) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/unpublish][%d] edgeNodeConfigurationUpdateEdgeNodeBaseOS3OK  %+v", 200, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3OK) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/unpublish][%d] edgeNodeConfigurationUpdateEdgeNodeBaseOS3OK  %+v", 200, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3OK) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS3Unauthorized creates a EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Unauthorized with default headers values
func NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS3Unauthorized() *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Unauthorized {
	return &EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Unauthorized{}
}

/*
EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Unauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Unauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration update edge node base o s3 unauthorized response has a 2xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration update edge node base o s3 unauthorized response has a 3xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration update edge node base o s3 unauthorized response has a 4xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration update edge node base o s3 unauthorized response has a 5xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration update edge node base o s3 unauthorized response a status code equal to that given
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the edge node configuration update edge node base o s3 unauthorized response
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Unauthorized) Code() int {
	return 401
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/unpublish][%d] edgeNodeConfigurationUpdateEdgeNodeBaseOS3Unauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Unauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/unpublish][%d] edgeNodeConfigurationUpdateEdgeNodeBaseOS3Unauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Unauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS3Forbidden creates a EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Forbidden with default headers values
func NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS3Forbidden() *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Forbidden {
	return &EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Forbidden{}
}

/*
EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Forbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Forbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration update edge node base o s3 forbidden response has a 2xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration update edge node base o s3 forbidden response has a 3xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration update edge node base o s3 forbidden response has a 4xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration update edge node base o s3 forbidden response has a 5xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration update edge node base o s3 forbidden response a status code equal to that given
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the edge node configuration update edge node base o s3 forbidden response
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Forbidden) Code() int {
	return 403
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Forbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/unpublish][%d] edgeNodeConfigurationUpdateEdgeNodeBaseOS3Forbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Forbidden) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/unpublish][%d] edgeNodeConfigurationUpdateEdgeNodeBaseOS3Forbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Forbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS3NotFound creates a EdgeNodeConfigurationUpdateEdgeNodeBaseOS3NotFound with default headers values
func NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS3NotFound() *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3NotFound {
	return &EdgeNodeConfigurationUpdateEdgeNodeBaseOS3NotFound{}
}

/*
EdgeNodeConfigurationUpdateEdgeNodeBaseOS3NotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeNodeConfigurationUpdateEdgeNodeBaseOS3NotFound struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration update edge node base o s3 not found response has a 2xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration update edge node base o s3 not found response has a 3xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration update edge node base o s3 not found response has a 4xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration update edge node base o s3 not found response has a 5xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration update edge node base o s3 not found response a status code equal to that given
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3NotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the edge node configuration update edge node base o s3 not found response
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3NotFound) Code() int {
	return 404
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3NotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/unpublish][%d] edgeNodeConfigurationUpdateEdgeNodeBaseOS3NotFound  %+v", 404, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3NotFound) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/unpublish][%d] edgeNodeConfigurationUpdateEdgeNodeBaseOS3NotFound  %+v", 404, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3NotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS3Conflict creates a EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Conflict with default headers values
func NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS3Conflict() *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Conflict {
	return &EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Conflict{}
}

/*
EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Conflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this operation will conflict with an already existing edge node record.
*/
type EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Conflict struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration update edge node base o s3 conflict response has a 2xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Conflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration update edge node base o s3 conflict response has a 3xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Conflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration update edge node base o s3 conflict response has a 4xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Conflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration update edge node base o s3 conflict response has a 5xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Conflict) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration update edge node base o s3 conflict response a status code equal to that given
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Conflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the edge node configuration update edge node base o s3 conflict response
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Conflict) Code() int {
	return 409
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Conflict) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/unpublish][%d] edgeNodeConfigurationUpdateEdgeNodeBaseOS3Conflict  %+v", 409, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Conflict) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/unpublish][%d] edgeNodeConfigurationUpdateEdgeNodeBaseOS3Conflict  %+v", 409, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Conflict) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Conflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS3InternalServerError creates a EdgeNodeConfigurationUpdateEdgeNodeBaseOS3InternalServerError with default headers values
func NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS3InternalServerError() *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3InternalServerError {
	return &EdgeNodeConfigurationUpdateEdgeNodeBaseOS3InternalServerError{}
}

/*
EdgeNodeConfigurationUpdateEdgeNodeBaseOS3InternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeNodeConfigurationUpdateEdgeNodeBaseOS3InternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration update edge node base o s3 internal server error response has a 2xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration update edge node base o s3 internal server error response has a 3xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration update edge node base o s3 internal server error response has a 4xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration update edge node base o s3 internal server error response has a 5xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node configuration update edge node base o s3 internal server error response a status code equal to that given
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the edge node configuration update edge node base o s3 internal server error response
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3InternalServerError) Code() int {
	return 500
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3InternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/unpublish][%d] edgeNodeConfigurationUpdateEdgeNodeBaseOS3InternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3InternalServerError) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/unpublish][%d] edgeNodeConfigurationUpdateEdgeNodeBaseOS3InternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3InternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS3GatewayTimeout creates a EdgeNodeConfigurationUpdateEdgeNodeBaseOS3GatewayTimeout with default headers values
func NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS3GatewayTimeout() *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3GatewayTimeout {
	return &EdgeNodeConfigurationUpdateEdgeNodeBaseOS3GatewayTimeout{}
}

/*
EdgeNodeConfigurationUpdateEdgeNodeBaseOS3GatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeNodeConfigurationUpdateEdgeNodeBaseOS3GatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration update edge node base o s3 gateway timeout response has a 2xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3GatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration update edge node base o s3 gateway timeout response has a 3xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3GatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration update edge node base o s3 gateway timeout response has a 4xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3GatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration update edge node base o s3 gateway timeout response has a 5xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3GatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node configuration update edge node base o s3 gateway timeout response a status code equal to that given
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3GatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the edge node configuration update edge node base o s3 gateway timeout response
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3GatewayTimeout) Code() int {
	return 504
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3GatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/unpublish][%d] edgeNodeConfigurationUpdateEdgeNodeBaseOS3GatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3GatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/unpublish][%d] edgeNodeConfigurationUpdateEdgeNodeBaseOS3GatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3GatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3GatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS3Default creates a EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Default with default headers values
func NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS3Default(code int) *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Default {
	return &EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Default{
		_statusCode: code,
	}
}

/*
EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Default struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// IsSuccess returns true when this edge node configuration update edge node base o s3 default response has a 2xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge node configuration update edge node base o s3 default response has a 3xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge node configuration update edge node base o s3 default response has a 4xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge node configuration update edge node base o s3 default response has a 5xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge node configuration update edge node base o s3 default response a status code equal to that given
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the edge node configuration update edge node base o s3 default response
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Default) Code() int {
	return o._statusCode
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Default) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/unpublish][%d] EdgeNodeConfiguration_UpdateEdgeNodeBaseOS3 default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Default) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/unpublish][%d] EdgeNodeConfiguration_UpdateEdgeNodeBaseOS3 default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Default) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS3Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

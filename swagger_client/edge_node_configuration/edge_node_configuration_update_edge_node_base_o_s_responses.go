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

// EdgeNodeConfigurationUpdateEdgeNodeBaseOSReader is a Reader for the EdgeNodeConfigurationUpdateEdgeNodeBaseOS structure.
type EdgeNodeConfigurationUpdateEdgeNodeBaseOSReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeNodeConfigurationUpdateEdgeNodeBaseOSOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeNodeConfigurationUpdateEdgeNodeBaseOSUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeNodeConfigurationUpdateEdgeNodeBaseOSForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeNodeConfigurationUpdateEdgeNodeBaseOSNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewEdgeNodeConfigurationUpdateEdgeNodeBaseOSConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeNodeConfigurationUpdateEdgeNodeBaseOSInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeNodeConfigurationUpdateEdgeNodeBaseOSGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeNodeConfigurationUpdateEdgeNodeBaseOSDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeNodeConfigurationUpdateEdgeNodeBaseOSOK creates a EdgeNodeConfigurationUpdateEdgeNodeBaseOSOK with default headers values
func NewEdgeNodeConfigurationUpdateEdgeNodeBaseOSOK() *EdgeNodeConfigurationUpdateEdgeNodeBaseOSOK {
	return &EdgeNodeConfigurationUpdateEdgeNodeBaseOSOK{}
}

/*
EdgeNodeConfigurationUpdateEdgeNodeBaseOSOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeNodeConfigurationUpdateEdgeNodeBaseOSOK struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration update edge node base o s o k response has a 2xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge node configuration update edge node base o s o k response has a 3xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration update edge node base o s o k response has a 4xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration update edge node base o s o k response has a 5xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration update edge node base o s o k response a status code equal to that given
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the edge node configuration update edge node base o s o k response
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSOK) Code() int {
	return 200
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSOK) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/apply][%d] edgeNodeConfigurationUpdateEdgeNodeBaseOSOK  %+v", 200, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSOK) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/apply][%d] edgeNodeConfigurationUpdateEdgeNodeBaseOSOK  %+v", 200, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSOK) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationUpdateEdgeNodeBaseOSUnauthorized creates a EdgeNodeConfigurationUpdateEdgeNodeBaseOSUnauthorized with default headers values
func NewEdgeNodeConfigurationUpdateEdgeNodeBaseOSUnauthorized() *EdgeNodeConfigurationUpdateEdgeNodeBaseOSUnauthorized {
	return &EdgeNodeConfigurationUpdateEdgeNodeBaseOSUnauthorized{}
}

/*
EdgeNodeConfigurationUpdateEdgeNodeBaseOSUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeNodeConfigurationUpdateEdgeNodeBaseOSUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration update edge node base o s unauthorized response has a 2xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration update edge node base o s unauthorized response has a 3xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration update edge node base o s unauthorized response has a 4xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration update edge node base o s unauthorized response has a 5xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration update edge node base o s unauthorized response a status code equal to that given
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the edge node configuration update edge node base o s unauthorized response
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSUnauthorized) Code() int {
	return 401
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/apply][%d] edgeNodeConfigurationUpdateEdgeNodeBaseOSUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/apply][%d] edgeNodeConfigurationUpdateEdgeNodeBaseOSUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationUpdateEdgeNodeBaseOSForbidden creates a EdgeNodeConfigurationUpdateEdgeNodeBaseOSForbidden with default headers values
func NewEdgeNodeConfigurationUpdateEdgeNodeBaseOSForbidden() *EdgeNodeConfigurationUpdateEdgeNodeBaseOSForbidden {
	return &EdgeNodeConfigurationUpdateEdgeNodeBaseOSForbidden{}
}

/*
EdgeNodeConfigurationUpdateEdgeNodeBaseOSForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeNodeConfigurationUpdateEdgeNodeBaseOSForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration update edge node base o s forbidden response has a 2xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration update edge node base o s forbidden response has a 3xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration update edge node base o s forbidden response has a 4xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration update edge node base o s forbidden response has a 5xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration update edge node base o s forbidden response a status code equal to that given
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the edge node configuration update edge node base o s forbidden response
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSForbidden) Code() int {
	return 403
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/apply][%d] edgeNodeConfigurationUpdateEdgeNodeBaseOSForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/apply][%d] edgeNodeConfigurationUpdateEdgeNodeBaseOSForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationUpdateEdgeNodeBaseOSNotFound creates a EdgeNodeConfigurationUpdateEdgeNodeBaseOSNotFound with default headers values
func NewEdgeNodeConfigurationUpdateEdgeNodeBaseOSNotFound() *EdgeNodeConfigurationUpdateEdgeNodeBaseOSNotFound {
	return &EdgeNodeConfigurationUpdateEdgeNodeBaseOSNotFound{}
}

/*
EdgeNodeConfigurationUpdateEdgeNodeBaseOSNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeNodeConfigurationUpdateEdgeNodeBaseOSNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration update edge node base o s not found response has a 2xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration update edge node base o s not found response has a 3xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration update edge node base o s not found response has a 4xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration update edge node base o s not found response has a 5xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration update edge node base o s not found response a status code equal to that given
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the edge node configuration update edge node base o s not found response
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSNotFound) Code() int {
	return 404
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/apply][%d] edgeNodeConfigurationUpdateEdgeNodeBaseOSNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSNotFound) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/apply][%d] edgeNodeConfigurationUpdateEdgeNodeBaseOSNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationUpdateEdgeNodeBaseOSConflict creates a EdgeNodeConfigurationUpdateEdgeNodeBaseOSConflict with default headers values
func NewEdgeNodeConfigurationUpdateEdgeNodeBaseOSConflict() *EdgeNodeConfigurationUpdateEdgeNodeBaseOSConflict {
	return &EdgeNodeConfigurationUpdateEdgeNodeBaseOSConflict{}
}

/*
EdgeNodeConfigurationUpdateEdgeNodeBaseOSConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this operation will conflict with an already existing edge node record.
*/
type EdgeNodeConfigurationUpdateEdgeNodeBaseOSConflict struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration update edge node base o s conflict response has a 2xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration update edge node base o s conflict response has a 3xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration update edge node base o s conflict response has a 4xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration update edge node base o s conflict response has a 5xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration update edge node base o s conflict response a status code equal to that given
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the edge node configuration update edge node base o s conflict response
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSConflict) Code() int {
	return 409
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/apply][%d] edgeNodeConfigurationUpdateEdgeNodeBaseOSConflict  %+v", 409, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSConflict) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/apply][%d] edgeNodeConfigurationUpdateEdgeNodeBaseOSConflict  %+v", 409, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSConflict) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationUpdateEdgeNodeBaseOSInternalServerError creates a EdgeNodeConfigurationUpdateEdgeNodeBaseOSInternalServerError with default headers values
func NewEdgeNodeConfigurationUpdateEdgeNodeBaseOSInternalServerError() *EdgeNodeConfigurationUpdateEdgeNodeBaseOSInternalServerError {
	return &EdgeNodeConfigurationUpdateEdgeNodeBaseOSInternalServerError{}
}

/*
EdgeNodeConfigurationUpdateEdgeNodeBaseOSInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeNodeConfigurationUpdateEdgeNodeBaseOSInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration update edge node base o s internal server error response has a 2xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration update edge node base o s internal server error response has a 3xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration update edge node base o s internal server error response has a 4xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration update edge node base o s internal server error response has a 5xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node configuration update edge node base o s internal server error response a status code equal to that given
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the edge node configuration update edge node base o s internal server error response
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSInternalServerError) Code() int {
	return 500
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/apply][%d] edgeNodeConfigurationUpdateEdgeNodeBaseOSInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSInternalServerError) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/apply][%d] edgeNodeConfigurationUpdateEdgeNodeBaseOSInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationUpdateEdgeNodeBaseOSGatewayTimeout creates a EdgeNodeConfigurationUpdateEdgeNodeBaseOSGatewayTimeout with default headers values
func NewEdgeNodeConfigurationUpdateEdgeNodeBaseOSGatewayTimeout() *EdgeNodeConfigurationUpdateEdgeNodeBaseOSGatewayTimeout {
	return &EdgeNodeConfigurationUpdateEdgeNodeBaseOSGatewayTimeout{}
}

/*
EdgeNodeConfigurationUpdateEdgeNodeBaseOSGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeNodeConfigurationUpdateEdgeNodeBaseOSGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration update edge node base o s gateway timeout response has a 2xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration update edge node base o s gateway timeout response has a 3xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration update edge node base o s gateway timeout response has a 4xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration update edge node base o s gateway timeout response has a 5xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node configuration update edge node base o s gateway timeout response a status code equal to that given
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the edge node configuration update edge node base o s gateway timeout response
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSGatewayTimeout) Code() int {
	return 504
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/apply][%d] edgeNodeConfigurationUpdateEdgeNodeBaseOSGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/apply][%d] edgeNodeConfigurationUpdateEdgeNodeBaseOSGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationUpdateEdgeNodeBaseOSDefault creates a EdgeNodeConfigurationUpdateEdgeNodeBaseOSDefault with default headers values
func NewEdgeNodeConfigurationUpdateEdgeNodeBaseOSDefault(code int) *EdgeNodeConfigurationUpdateEdgeNodeBaseOSDefault {
	return &EdgeNodeConfigurationUpdateEdgeNodeBaseOSDefault{
		_statusCode: code,
	}
}

/*
EdgeNodeConfigurationUpdateEdgeNodeBaseOSDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeNodeConfigurationUpdateEdgeNodeBaseOSDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// IsSuccess returns true when this edge node configuration update edge node base o s default response has a 2xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge node configuration update edge node base o s default response has a 3xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge node configuration update edge node base o s default response has a 4xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge node configuration update edge node base o s default response has a 5xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge node configuration update edge node base o s default response a status code equal to that given
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the edge node configuration update edge node base o s default response
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSDefault) Code() int {
	return o._statusCode
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/apply][%d] EdgeNodeConfiguration_UpdateEdgeNodeBaseOS default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSDefault) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/apply][%d] EdgeNodeConfiguration_UpdateEdgeNodeBaseOS default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package identity_access_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// IdentityAccessManagementQueryUsersReader is a Reader for the IdentityAccessManagementQueryUsers structure.
type IdentityAccessManagementQueryUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IdentityAccessManagementQueryUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIdentityAccessManagementQueryUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewIdentityAccessManagementQueryUsersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewIdentityAccessManagementQueryUsersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewIdentityAccessManagementQueryUsersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewIdentityAccessManagementQueryUsersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewIdentityAccessManagementQueryUsersGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIdentityAccessManagementQueryUsersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIdentityAccessManagementQueryUsersOK creates a IdentityAccessManagementQueryUsersOK with default headers values
func NewIdentityAccessManagementQueryUsersOK() *IdentityAccessManagementQueryUsersOK {
	return &IdentityAccessManagementQueryUsersOK{}
}

/*
IdentityAccessManagementQueryUsersOK describes a response with status code 200, with default header values.

A successful response.
*/
type IdentityAccessManagementQueryUsersOK struct {
	Payload *swagger_models.CrudResponse
}

// IsSuccess returns true when this identity access management query users o k response has a 2xx status code
func (o *IdentityAccessManagementQueryUsersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this identity access management query users o k response has a 3xx status code
func (o *IdentityAccessManagementQueryUsersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management query users o k response has a 4xx status code
func (o *IdentityAccessManagementQueryUsersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management query users o k response has a 5xx status code
func (o *IdentityAccessManagementQueryUsersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management query users o k response a status code equal to that given
func (o *IdentityAccessManagementQueryUsersOK) IsCode(code int) bool {
	return code == 200
}

func (o *IdentityAccessManagementQueryUsersOK) Error() string {
	return fmt.Sprintf("[GET /v1/users][%d] identityAccessManagementQueryUsersOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementQueryUsersOK) String() string {
	return fmt.Sprintf("[GET /v1/users][%d] identityAccessManagementQueryUsersOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementQueryUsersOK) GetPayload() *swagger_models.CrudResponse {
	return o.Payload
}

func (o *IdentityAccessManagementQueryUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.CrudResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementQueryUsersBadRequest creates a IdentityAccessManagementQueryUsersBadRequest with default headers values
func NewIdentityAccessManagementQueryUsersBadRequest() *IdentityAccessManagementQueryUsersBadRequest {
	return &IdentityAccessManagementQueryUsersBadRequest{}
}

/*
IdentityAccessManagementQueryUsersBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of invalid value of filter parameters.
*/
type IdentityAccessManagementQueryUsersBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management query users bad request response has a 2xx status code
func (o *IdentityAccessManagementQueryUsersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management query users bad request response has a 3xx status code
func (o *IdentityAccessManagementQueryUsersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management query users bad request response has a 4xx status code
func (o *IdentityAccessManagementQueryUsersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management query users bad request response has a 5xx status code
func (o *IdentityAccessManagementQueryUsersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management query users bad request response a status code equal to that given
func (o *IdentityAccessManagementQueryUsersBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *IdentityAccessManagementQueryUsersBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/users][%d] identityAccessManagementQueryUsersBadRequest  %+v", 400, o.Payload)
}

func (o *IdentityAccessManagementQueryUsersBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/users][%d] identityAccessManagementQueryUsersBadRequest  %+v", 400, o.Payload)
}

func (o *IdentityAccessManagementQueryUsersBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementQueryUsersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementQueryUsersUnauthorized creates a IdentityAccessManagementQueryUsersUnauthorized with default headers values
func NewIdentityAccessManagementQueryUsersUnauthorized() *IdentityAccessManagementQueryUsersUnauthorized {
	return &IdentityAccessManagementQueryUsersUnauthorized{}
}

/*
IdentityAccessManagementQueryUsersUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type IdentityAccessManagementQueryUsersUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management query users unauthorized response has a 2xx status code
func (o *IdentityAccessManagementQueryUsersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management query users unauthorized response has a 3xx status code
func (o *IdentityAccessManagementQueryUsersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management query users unauthorized response has a 4xx status code
func (o *IdentityAccessManagementQueryUsersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management query users unauthorized response has a 5xx status code
func (o *IdentityAccessManagementQueryUsersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management query users unauthorized response a status code equal to that given
func (o *IdentityAccessManagementQueryUsersUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *IdentityAccessManagementQueryUsersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/users][%d] identityAccessManagementQueryUsersUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementQueryUsersUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/users][%d] identityAccessManagementQueryUsersUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementQueryUsersUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementQueryUsersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementQueryUsersForbidden creates a IdentityAccessManagementQueryUsersForbidden with default headers values
func NewIdentityAccessManagementQueryUsersForbidden() *IdentityAccessManagementQueryUsersForbidden {
	return &IdentityAccessManagementQueryUsersForbidden{}
}

/*
IdentityAccessManagementQueryUsersForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type IdentityAccessManagementQueryUsersForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management query users forbidden response has a 2xx status code
func (o *IdentityAccessManagementQueryUsersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management query users forbidden response has a 3xx status code
func (o *IdentityAccessManagementQueryUsersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management query users forbidden response has a 4xx status code
func (o *IdentityAccessManagementQueryUsersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management query users forbidden response has a 5xx status code
func (o *IdentityAccessManagementQueryUsersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management query users forbidden response a status code equal to that given
func (o *IdentityAccessManagementQueryUsersForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *IdentityAccessManagementQueryUsersForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/users][%d] identityAccessManagementQueryUsersForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementQueryUsersForbidden) String() string {
	return fmt.Sprintf("[GET /v1/users][%d] identityAccessManagementQueryUsersForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementQueryUsersForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementQueryUsersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementQueryUsersInternalServerError creates a IdentityAccessManagementQueryUsersInternalServerError with default headers values
func NewIdentityAccessManagementQueryUsersInternalServerError() *IdentityAccessManagementQueryUsersInternalServerError {
	return &IdentityAccessManagementQueryUsersInternalServerError{}
}

/*
IdentityAccessManagementQueryUsersInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type IdentityAccessManagementQueryUsersInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management query users internal server error response has a 2xx status code
func (o *IdentityAccessManagementQueryUsersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management query users internal server error response has a 3xx status code
func (o *IdentityAccessManagementQueryUsersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management query users internal server error response has a 4xx status code
func (o *IdentityAccessManagementQueryUsersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management query users internal server error response has a 5xx status code
func (o *IdentityAccessManagementQueryUsersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management query users internal server error response a status code equal to that given
func (o *IdentityAccessManagementQueryUsersInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *IdentityAccessManagementQueryUsersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/users][%d] identityAccessManagementQueryUsersInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementQueryUsersInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/users][%d] identityAccessManagementQueryUsersInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementQueryUsersInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementQueryUsersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementQueryUsersGatewayTimeout creates a IdentityAccessManagementQueryUsersGatewayTimeout with default headers values
func NewIdentityAccessManagementQueryUsersGatewayTimeout() *IdentityAccessManagementQueryUsersGatewayTimeout {
	return &IdentityAccessManagementQueryUsersGatewayTimeout{}
}

/*
IdentityAccessManagementQueryUsersGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type IdentityAccessManagementQueryUsersGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management query users gateway timeout response has a 2xx status code
func (o *IdentityAccessManagementQueryUsersGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management query users gateway timeout response has a 3xx status code
func (o *IdentityAccessManagementQueryUsersGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management query users gateway timeout response has a 4xx status code
func (o *IdentityAccessManagementQueryUsersGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management query users gateway timeout response has a 5xx status code
func (o *IdentityAccessManagementQueryUsersGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management query users gateway timeout response a status code equal to that given
func (o *IdentityAccessManagementQueryUsersGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *IdentityAccessManagementQueryUsersGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/users][%d] identityAccessManagementQueryUsersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementQueryUsersGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/users][%d] identityAccessManagementQueryUsersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementQueryUsersGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementQueryUsersGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementQueryUsersDefault creates a IdentityAccessManagementQueryUsersDefault with default headers values
func NewIdentityAccessManagementQueryUsersDefault(code int) *IdentityAccessManagementQueryUsersDefault {
	return &IdentityAccessManagementQueryUsersDefault{
		_statusCode: code,
	}
}

/*
IdentityAccessManagementQueryUsersDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type IdentityAccessManagementQueryUsersDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the identity access management query users default response
func (o *IdentityAccessManagementQueryUsersDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this identity access management query users default response has a 2xx status code
func (o *IdentityAccessManagementQueryUsersDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this identity access management query users default response has a 3xx status code
func (o *IdentityAccessManagementQueryUsersDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this identity access management query users default response has a 4xx status code
func (o *IdentityAccessManagementQueryUsersDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this identity access management query users default response has a 5xx status code
func (o *IdentityAccessManagementQueryUsersDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this identity access management query users default response a status code equal to that given
func (o *IdentityAccessManagementQueryUsersDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IdentityAccessManagementQueryUsersDefault) Error() string {
	return fmt.Sprintf("[GET /v1/users][%d] IdentityAccessManagement_QueryUsers default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementQueryUsersDefault) String() string {
	return fmt.Sprintf("[GET /v1/users][%d] IdentityAccessManagement_QueryUsers default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementQueryUsersDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *IdentityAccessManagementQueryUsersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

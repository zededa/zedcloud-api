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

// IdentityAccessManagementQueryRolesReader is a Reader for the IdentityAccessManagementQueryRoles structure.
type IdentityAccessManagementQueryRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IdentityAccessManagementQueryRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIdentityAccessManagementQueryRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewIdentityAccessManagementQueryRolesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewIdentityAccessManagementQueryRolesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewIdentityAccessManagementQueryRolesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewIdentityAccessManagementQueryRolesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewIdentityAccessManagementQueryRolesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIdentityAccessManagementQueryRolesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIdentityAccessManagementQueryRolesOK creates a IdentityAccessManagementQueryRolesOK with default headers values
func NewIdentityAccessManagementQueryRolesOK() *IdentityAccessManagementQueryRolesOK {
	return &IdentityAccessManagementQueryRolesOK{}
}

/*
IdentityAccessManagementQueryRolesOK describes a response with status code 200, with default header values.

A successful response.
*/
type IdentityAccessManagementQueryRolesOK struct {
	Payload *swagger_models.CrudResponse
}

// IsSuccess returns true when this identity access management query roles o k response has a 2xx status code
func (o *IdentityAccessManagementQueryRolesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this identity access management query roles o k response has a 3xx status code
func (o *IdentityAccessManagementQueryRolesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management query roles o k response has a 4xx status code
func (o *IdentityAccessManagementQueryRolesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management query roles o k response has a 5xx status code
func (o *IdentityAccessManagementQueryRolesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management query roles o k response a status code equal to that given
func (o *IdentityAccessManagementQueryRolesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the identity access management query roles o k response
func (o *IdentityAccessManagementQueryRolesOK) Code() int {
	return 200
}

func (o *IdentityAccessManagementQueryRolesOK) Error() string {
	return fmt.Sprintf("[GET /v1/roles][%d] identityAccessManagementQueryRolesOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementQueryRolesOK) String() string {
	return fmt.Sprintf("[GET /v1/roles][%d] identityAccessManagementQueryRolesOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementQueryRolesOK) GetPayload() *swagger_models.CrudResponse {
	return o.Payload
}

func (o *IdentityAccessManagementQueryRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.CrudResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementQueryRolesBadRequest creates a IdentityAccessManagementQueryRolesBadRequest with default headers values
func NewIdentityAccessManagementQueryRolesBadRequest() *IdentityAccessManagementQueryRolesBadRequest {
	return &IdentityAccessManagementQueryRolesBadRequest{}
}

/*
IdentityAccessManagementQueryRolesBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of invalid value of filter parameters.
*/
type IdentityAccessManagementQueryRolesBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management query roles bad request response has a 2xx status code
func (o *IdentityAccessManagementQueryRolesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management query roles bad request response has a 3xx status code
func (o *IdentityAccessManagementQueryRolesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management query roles bad request response has a 4xx status code
func (o *IdentityAccessManagementQueryRolesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management query roles bad request response has a 5xx status code
func (o *IdentityAccessManagementQueryRolesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management query roles bad request response a status code equal to that given
func (o *IdentityAccessManagementQueryRolesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the identity access management query roles bad request response
func (o *IdentityAccessManagementQueryRolesBadRequest) Code() int {
	return 400
}

func (o *IdentityAccessManagementQueryRolesBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/roles][%d] identityAccessManagementQueryRolesBadRequest  %+v", 400, o.Payload)
}

func (o *IdentityAccessManagementQueryRolesBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/roles][%d] identityAccessManagementQueryRolesBadRequest  %+v", 400, o.Payload)
}

func (o *IdentityAccessManagementQueryRolesBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementQueryRolesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementQueryRolesUnauthorized creates a IdentityAccessManagementQueryRolesUnauthorized with default headers values
func NewIdentityAccessManagementQueryRolesUnauthorized() *IdentityAccessManagementQueryRolesUnauthorized {
	return &IdentityAccessManagementQueryRolesUnauthorized{}
}

/*
IdentityAccessManagementQueryRolesUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type IdentityAccessManagementQueryRolesUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management query roles unauthorized response has a 2xx status code
func (o *IdentityAccessManagementQueryRolesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management query roles unauthorized response has a 3xx status code
func (o *IdentityAccessManagementQueryRolesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management query roles unauthorized response has a 4xx status code
func (o *IdentityAccessManagementQueryRolesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management query roles unauthorized response has a 5xx status code
func (o *IdentityAccessManagementQueryRolesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management query roles unauthorized response a status code equal to that given
func (o *IdentityAccessManagementQueryRolesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the identity access management query roles unauthorized response
func (o *IdentityAccessManagementQueryRolesUnauthorized) Code() int {
	return 401
}

func (o *IdentityAccessManagementQueryRolesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/roles][%d] identityAccessManagementQueryRolesUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementQueryRolesUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/roles][%d] identityAccessManagementQueryRolesUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementQueryRolesUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementQueryRolesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementQueryRolesForbidden creates a IdentityAccessManagementQueryRolesForbidden with default headers values
func NewIdentityAccessManagementQueryRolesForbidden() *IdentityAccessManagementQueryRolesForbidden {
	return &IdentityAccessManagementQueryRolesForbidden{}
}

/*
IdentityAccessManagementQueryRolesForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type IdentityAccessManagementQueryRolesForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management query roles forbidden response has a 2xx status code
func (o *IdentityAccessManagementQueryRolesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management query roles forbidden response has a 3xx status code
func (o *IdentityAccessManagementQueryRolesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management query roles forbidden response has a 4xx status code
func (o *IdentityAccessManagementQueryRolesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management query roles forbidden response has a 5xx status code
func (o *IdentityAccessManagementQueryRolesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management query roles forbidden response a status code equal to that given
func (o *IdentityAccessManagementQueryRolesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the identity access management query roles forbidden response
func (o *IdentityAccessManagementQueryRolesForbidden) Code() int {
	return 403
}

func (o *IdentityAccessManagementQueryRolesForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/roles][%d] identityAccessManagementQueryRolesForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementQueryRolesForbidden) String() string {
	return fmt.Sprintf("[GET /v1/roles][%d] identityAccessManagementQueryRolesForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementQueryRolesForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementQueryRolesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementQueryRolesInternalServerError creates a IdentityAccessManagementQueryRolesInternalServerError with default headers values
func NewIdentityAccessManagementQueryRolesInternalServerError() *IdentityAccessManagementQueryRolesInternalServerError {
	return &IdentityAccessManagementQueryRolesInternalServerError{}
}

/*
IdentityAccessManagementQueryRolesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type IdentityAccessManagementQueryRolesInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management query roles internal server error response has a 2xx status code
func (o *IdentityAccessManagementQueryRolesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management query roles internal server error response has a 3xx status code
func (o *IdentityAccessManagementQueryRolesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management query roles internal server error response has a 4xx status code
func (o *IdentityAccessManagementQueryRolesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management query roles internal server error response has a 5xx status code
func (o *IdentityAccessManagementQueryRolesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management query roles internal server error response a status code equal to that given
func (o *IdentityAccessManagementQueryRolesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the identity access management query roles internal server error response
func (o *IdentityAccessManagementQueryRolesInternalServerError) Code() int {
	return 500
}

func (o *IdentityAccessManagementQueryRolesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/roles][%d] identityAccessManagementQueryRolesInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementQueryRolesInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/roles][%d] identityAccessManagementQueryRolesInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementQueryRolesInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementQueryRolesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementQueryRolesGatewayTimeout creates a IdentityAccessManagementQueryRolesGatewayTimeout with default headers values
func NewIdentityAccessManagementQueryRolesGatewayTimeout() *IdentityAccessManagementQueryRolesGatewayTimeout {
	return &IdentityAccessManagementQueryRolesGatewayTimeout{}
}

/*
IdentityAccessManagementQueryRolesGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type IdentityAccessManagementQueryRolesGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management query roles gateway timeout response has a 2xx status code
func (o *IdentityAccessManagementQueryRolesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management query roles gateway timeout response has a 3xx status code
func (o *IdentityAccessManagementQueryRolesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management query roles gateway timeout response has a 4xx status code
func (o *IdentityAccessManagementQueryRolesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management query roles gateway timeout response has a 5xx status code
func (o *IdentityAccessManagementQueryRolesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management query roles gateway timeout response a status code equal to that given
func (o *IdentityAccessManagementQueryRolesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the identity access management query roles gateway timeout response
func (o *IdentityAccessManagementQueryRolesGatewayTimeout) Code() int {
	return 504
}

func (o *IdentityAccessManagementQueryRolesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/roles][%d] identityAccessManagementQueryRolesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementQueryRolesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/roles][%d] identityAccessManagementQueryRolesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementQueryRolesGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementQueryRolesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementQueryRolesDefault creates a IdentityAccessManagementQueryRolesDefault with default headers values
func NewIdentityAccessManagementQueryRolesDefault(code int) *IdentityAccessManagementQueryRolesDefault {
	return &IdentityAccessManagementQueryRolesDefault{
		_statusCode: code,
	}
}

/*
IdentityAccessManagementQueryRolesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type IdentityAccessManagementQueryRolesDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// IsSuccess returns true when this identity access management query roles default response has a 2xx status code
func (o *IdentityAccessManagementQueryRolesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this identity access management query roles default response has a 3xx status code
func (o *IdentityAccessManagementQueryRolesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this identity access management query roles default response has a 4xx status code
func (o *IdentityAccessManagementQueryRolesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this identity access management query roles default response has a 5xx status code
func (o *IdentityAccessManagementQueryRolesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this identity access management query roles default response a status code equal to that given
func (o *IdentityAccessManagementQueryRolesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the identity access management query roles default response
func (o *IdentityAccessManagementQueryRolesDefault) Code() int {
	return o._statusCode
}

func (o *IdentityAccessManagementQueryRolesDefault) Error() string {
	return fmt.Sprintf("[GET /v1/roles][%d] IdentityAccessManagement_QueryRoles default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementQueryRolesDefault) String() string {
	return fmt.Sprintf("[GET /v1/roles][%d] IdentityAccessManagement_QueryRoles default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementQueryRolesDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *IdentityAccessManagementQueryRolesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

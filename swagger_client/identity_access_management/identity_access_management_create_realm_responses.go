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

// IdentityAccessManagementCreateRealmReader is a Reader for the IdentityAccessManagementCreateRealm structure.
type IdentityAccessManagementCreateRealmReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IdentityAccessManagementCreateRealmReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIdentityAccessManagementCreateRealmOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewIdentityAccessManagementCreateRealmBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewIdentityAccessManagementCreateRealmUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewIdentityAccessManagementCreateRealmForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewIdentityAccessManagementCreateRealmConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewIdentityAccessManagementCreateRealmInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewIdentityAccessManagementCreateRealmGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIdentityAccessManagementCreateRealmDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIdentityAccessManagementCreateRealmOK creates a IdentityAccessManagementCreateRealmOK with default headers values
func NewIdentityAccessManagementCreateRealmOK() *IdentityAccessManagementCreateRealmOK {
	return &IdentityAccessManagementCreateRealmOK{}
}

/*
IdentityAccessManagementCreateRealmOK describes a response with status code 200, with default header values.

A successful response.
*/
type IdentityAccessManagementCreateRealmOK struct {
	Payload *swagger_models.CrudResponse
}

// IsSuccess returns true when this identity access management create realm o k response has a 2xx status code
func (o *IdentityAccessManagementCreateRealmOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this identity access management create realm o k response has a 3xx status code
func (o *IdentityAccessManagementCreateRealmOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management create realm o k response has a 4xx status code
func (o *IdentityAccessManagementCreateRealmOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management create realm o k response has a 5xx status code
func (o *IdentityAccessManagementCreateRealmOK) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management create realm o k response a status code equal to that given
func (o *IdentityAccessManagementCreateRealmOK) IsCode(code int) bool {
	return code == 200
}

func (o *IdentityAccessManagementCreateRealmOK) Error() string {
	return fmt.Sprintf("[POST /v1/realms][%d] identityAccessManagementCreateRealmOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementCreateRealmOK) String() string {
	return fmt.Sprintf("[POST /v1/realms][%d] identityAccessManagementCreateRealmOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementCreateRealmOK) GetPayload() *swagger_models.CrudResponse {
	return o.Payload
}

func (o *IdentityAccessManagementCreateRealmOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.CrudResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementCreateRealmBadRequest creates a IdentityAccessManagementCreateRealmBadRequest with default headers values
func NewIdentityAccessManagementCreateRealmBadRequest() *IdentityAccessManagementCreateRealmBadRequest {
	return &IdentityAccessManagementCreateRealmBadRequest{}
}

/*
IdentityAccessManagementCreateRealmBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type IdentityAccessManagementCreateRealmBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management create realm bad request response has a 2xx status code
func (o *IdentityAccessManagementCreateRealmBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management create realm bad request response has a 3xx status code
func (o *IdentityAccessManagementCreateRealmBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management create realm bad request response has a 4xx status code
func (o *IdentityAccessManagementCreateRealmBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management create realm bad request response has a 5xx status code
func (o *IdentityAccessManagementCreateRealmBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management create realm bad request response a status code equal to that given
func (o *IdentityAccessManagementCreateRealmBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *IdentityAccessManagementCreateRealmBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/realms][%d] identityAccessManagementCreateRealmBadRequest  %+v", 400, o.Payload)
}

func (o *IdentityAccessManagementCreateRealmBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/realms][%d] identityAccessManagementCreateRealmBadRequest  %+v", 400, o.Payload)
}

func (o *IdentityAccessManagementCreateRealmBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementCreateRealmBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementCreateRealmUnauthorized creates a IdentityAccessManagementCreateRealmUnauthorized with default headers values
func NewIdentityAccessManagementCreateRealmUnauthorized() *IdentityAccessManagementCreateRealmUnauthorized {
	return &IdentityAccessManagementCreateRealmUnauthorized{}
}

/*
IdentityAccessManagementCreateRealmUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type IdentityAccessManagementCreateRealmUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management create realm unauthorized response has a 2xx status code
func (o *IdentityAccessManagementCreateRealmUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management create realm unauthorized response has a 3xx status code
func (o *IdentityAccessManagementCreateRealmUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management create realm unauthorized response has a 4xx status code
func (o *IdentityAccessManagementCreateRealmUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management create realm unauthorized response has a 5xx status code
func (o *IdentityAccessManagementCreateRealmUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management create realm unauthorized response a status code equal to that given
func (o *IdentityAccessManagementCreateRealmUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *IdentityAccessManagementCreateRealmUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/realms][%d] identityAccessManagementCreateRealmUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementCreateRealmUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/realms][%d] identityAccessManagementCreateRealmUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementCreateRealmUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementCreateRealmUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementCreateRealmForbidden creates a IdentityAccessManagementCreateRealmForbidden with default headers values
func NewIdentityAccessManagementCreateRealmForbidden() *IdentityAccessManagementCreateRealmForbidden {
	return &IdentityAccessManagementCreateRealmForbidden{}
}

/*
IdentityAccessManagementCreateRealmForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type IdentityAccessManagementCreateRealmForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management create realm forbidden response has a 2xx status code
func (o *IdentityAccessManagementCreateRealmForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management create realm forbidden response has a 3xx status code
func (o *IdentityAccessManagementCreateRealmForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management create realm forbidden response has a 4xx status code
func (o *IdentityAccessManagementCreateRealmForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management create realm forbidden response has a 5xx status code
func (o *IdentityAccessManagementCreateRealmForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management create realm forbidden response a status code equal to that given
func (o *IdentityAccessManagementCreateRealmForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *IdentityAccessManagementCreateRealmForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/realms][%d] identityAccessManagementCreateRealmForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementCreateRealmForbidden) String() string {
	return fmt.Sprintf("[POST /v1/realms][%d] identityAccessManagementCreateRealmForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementCreateRealmForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementCreateRealmForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementCreateRealmConflict creates a IdentityAccessManagementCreateRealmConflict with default headers values
func NewIdentityAccessManagementCreateRealmConflict() *IdentityAccessManagementCreateRealmConflict {
	return &IdentityAccessManagementCreateRealmConflict{}
}

/*
IdentityAccessManagementCreateRealmConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this realm will conflict with an already existing realm.
*/
type IdentityAccessManagementCreateRealmConflict struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management create realm conflict response has a 2xx status code
func (o *IdentityAccessManagementCreateRealmConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management create realm conflict response has a 3xx status code
func (o *IdentityAccessManagementCreateRealmConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management create realm conflict response has a 4xx status code
func (o *IdentityAccessManagementCreateRealmConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management create realm conflict response has a 5xx status code
func (o *IdentityAccessManagementCreateRealmConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management create realm conflict response a status code equal to that given
func (o *IdentityAccessManagementCreateRealmConflict) IsCode(code int) bool {
	return code == 409
}

func (o *IdentityAccessManagementCreateRealmConflict) Error() string {
	return fmt.Sprintf("[POST /v1/realms][%d] identityAccessManagementCreateRealmConflict  %+v", 409, o.Payload)
}

func (o *IdentityAccessManagementCreateRealmConflict) String() string {
	return fmt.Sprintf("[POST /v1/realms][%d] identityAccessManagementCreateRealmConflict  %+v", 409, o.Payload)
}

func (o *IdentityAccessManagementCreateRealmConflict) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementCreateRealmConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementCreateRealmInternalServerError creates a IdentityAccessManagementCreateRealmInternalServerError with default headers values
func NewIdentityAccessManagementCreateRealmInternalServerError() *IdentityAccessManagementCreateRealmInternalServerError {
	return &IdentityAccessManagementCreateRealmInternalServerError{}
}

/*
IdentityAccessManagementCreateRealmInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type IdentityAccessManagementCreateRealmInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management create realm internal server error response has a 2xx status code
func (o *IdentityAccessManagementCreateRealmInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management create realm internal server error response has a 3xx status code
func (o *IdentityAccessManagementCreateRealmInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management create realm internal server error response has a 4xx status code
func (o *IdentityAccessManagementCreateRealmInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management create realm internal server error response has a 5xx status code
func (o *IdentityAccessManagementCreateRealmInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management create realm internal server error response a status code equal to that given
func (o *IdentityAccessManagementCreateRealmInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *IdentityAccessManagementCreateRealmInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/realms][%d] identityAccessManagementCreateRealmInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementCreateRealmInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/realms][%d] identityAccessManagementCreateRealmInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementCreateRealmInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementCreateRealmInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementCreateRealmGatewayTimeout creates a IdentityAccessManagementCreateRealmGatewayTimeout with default headers values
func NewIdentityAccessManagementCreateRealmGatewayTimeout() *IdentityAccessManagementCreateRealmGatewayTimeout {
	return &IdentityAccessManagementCreateRealmGatewayTimeout{}
}

/*
IdentityAccessManagementCreateRealmGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type IdentityAccessManagementCreateRealmGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management create realm gateway timeout response has a 2xx status code
func (o *IdentityAccessManagementCreateRealmGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management create realm gateway timeout response has a 3xx status code
func (o *IdentityAccessManagementCreateRealmGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management create realm gateway timeout response has a 4xx status code
func (o *IdentityAccessManagementCreateRealmGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management create realm gateway timeout response has a 5xx status code
func (o *IdentityAccessManagementCreateRealmGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management create realm gateway timeout response a status code equal to that given
func (o *IdentityAccessManagementCreateRealmGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *IdentityAccessManagementCreateRealmGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v1/realms][%d] identityAccessManagementCreateRealmGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementCreateRealmGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /v1/realms][%d] identityAccessManagementCreateRealmGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementCreateRealmGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementCreateRealmGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementCreateRealmDefault creates a IdentityAccessManagementCreateRealmDefault with default headers values
func NewIdentityAccessManagementCreateRealmDefault(code int) *IdentityAccessManagementCreateRealmDefault {
	return &IdentityAccessManagementCreateRealmDefault{
		_statusCode: code,
	}
}

/*
IdentityAccessManagementCreateRealmDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type IdentityAccessManagementCreateRealmDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the identity access management create realm default response
func (o *IdentityAccessManagementCreateRealmDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this identity access management create realm default response has a 2xx status code
func (o *IdentityAccessManagementCreateRealmDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this identity access management create realm default response has a 3xx status code
func (o *IdentityAccessManagementCreateRealmDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this identity access management create realm default response has a 4xx status code
func (o *IdentityAccessManagementCreateRealmDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this identity access management create realm default response has a 5xx status code
func (o *IdentityAccessManagementCreateRealmDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this identity access management create realm default response a status code equal to that given
func (o *IdentityAccessManagementCreateRealmDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IdentityAccessManagementCreateRealmDefault) Error() string {
	return fmt.Sprintf("[POST /v1/realms][%d] IdentityAccessManagement_CreateRealm default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementCreateRealmDefault) String() string {
	return fmt.Sprintf("[POST /v1/realms][%d] IdentityAccessManagement_CreateRealm default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementCreateRealmDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *IdentityAccessManagementCreateRealmDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

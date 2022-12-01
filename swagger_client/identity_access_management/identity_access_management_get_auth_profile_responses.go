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

// IdentityAccessManagementGetAuthProfileReader is a Reader for the IdentityAccessManagementGetAuthProfile structure.
type IdentityAccessManagementGetAuthProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IdentityAccessManagementGetAuthProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIdentityAccessManagementGetAuthProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewIdentityAccessManagementGetAuthProfileUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewIdentityAccessManagementGetAuthProfileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewIdentityAccessManagementGetAuthProfileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewIdentityAccessManagementGetAuthProfileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewIdentityAccessManagementGetAuthProfileGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIdentityAccessManagementGetAuthProfileDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIdentityAccessManagementGetAuthProfileOK creates a IdentityAccessManagementGetAuthProfileOK with default headers values
func NewIdentityAccessManagementGetAuthProfileOK() *IdentityAccessManagementGetAuthProfileOK {
	return &IdentityAccessManagementGetAuthProfileOK{}
}

/*
IdentityAccessManagementGetAuthProfileOK describes a response with status code 200, with default header values.

A successful response.
*/
type IdentityAccessManagementGetAuthProfileOK struct {
	Payload *swagger_models.CrudResponseRead
}

// IsSuccess returns true when this identity access management get auth profile o k response has a 2xx status code
func (o *IdentityAccessManagementGetAuthProfileOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this identity access management get auth profile o k response has a 3xx status code
func (o *IdentityAccessManagementGetAuthProfileOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get auth profile o k response has a 4xx status code
func (o *IdentityAccessManagementGetAuthProfileOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management get auth profile o k response has a 5xx status code
func (o *IdentityAccessManagementGetAuthProfileOK) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management get auth profile o k response a status code equal to that given
func (o *IdentityAccessManagementGetAuthProfileOK) IsCode(code int) bool {
	return code == 200
}

func (o *IdentityAccessManagementGetAuthProfileOK) Error() string {
	return fmt.Sprintf("[GET /v1/authorization/profiles/id/{id}][%d] identityAccessManagementGetAuthProfileOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementGetAuthProfileOK) String() string {
	return fmt.Sprintf("[GET /v1/authorization/profiles/id/{id}][%d] identityAccessManagementGetAuthProfileOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementGetAuthProfileOK) GetPayload() *swagger_models.CrudResponseRead {
	return o.Payload
}

func (o *IdentityAccessManagementGetAuthProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.CrudResponseRead)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetAuthProfileUnauthorized creates a IdentityAccessManagementGetAuthProfileUnauthorized with default headers values
func NewIdentityAccessManagementGetAuthProfileUnauthorized() *IdentityAccessManagementGetAuthProfileUnauthorized {
	return &IdentityAccessManagementGetAuthProfileUnauthorized{}
}

/*
IdentityAccessManagementGetAuthProfileUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type IdentityAccessManagementGetAuthProfileUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management get auth profile unauthorized response has a 2xx status code
func (o *IdentityAccessManagementGetAuthProfileUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management get auth profile unauthorized response has a 3xx status code
func (o *IdentityAccessManagementGetAuthProfileUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get auth profile unauthorized response has a 4xx status code
func (o *IdentityAccessManagementGetAuthProfileUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management get auth profile unauthorized response has a 5xx status code
func (o *IdentityAccessManagementGetAuthProfileUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management get auth profile unauthorized response a status code equal to that given
func (o *IdentityAccessManagementGetAuthProfileUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *IdentityAccessManagementGetAuthProfileUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/authorization/profiles/id/{id}][%d] identityAccessManagementGetAuthProfileUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementGetAuthProfileUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/authorization/profiles/id/{id}][%d] identityAccessManagementGetAuthProfileUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementGetAuthProfileUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementGetAuthProfileUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetAuthProfileForbidden creates a IdentityAccessManagementGetAuthProfileForbidden with default headers values
func NewIdentityAccessManagementGetAuthProfileForbidden() *IdentityAccessManagementGetAuthProfileForbidden {
	return &IdentityAccessManagementGetAuthProfileForbidden{}
}

/*
IdentityAccessManagementGetAuthProfileForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type IdentityAccessManagementGetAuthProfileForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management get auth profile forbidden response has a 2xx status code
func (o *IdentityAccessManagementGetAuthProfileForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management get auth profile forbidden response has a 3xx status code
func (o *IdentityAccessManagementGetAuthProfileForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get auth profile forbidden response has a 4xx status code
func (o *IdentityAccessManagementGetAuthProfileForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management get auth profile forbidden response has a 5xx status code
func (o *IdentityAccessManagementGetAuthProfileForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management get auth profile forbidden response a status code equal to that given
func (o *IdentityAccessManagementGetAuthProfileForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *IdentityAccessManagementGetAuthProfileForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/authorization/profiles/id/{id}][%d] identityAccessManagementGetAuthProfileForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementGetAuthProfileForbidden) String() string {
	return fmt.Sprintf("[GET /v1/authorization/profiles/id/{id}][%d] identityAccessManagementGetAuthProfileForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementGetAuthProfileForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementGetAuthProfileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetAuthProfileNotFound creates a IdentityAccessManagementGetAuthProfileNotFound with default headers values
func NewIdentityAccessManagementGetAuthProfileNotFound() *IdentityAccessManagementGetAuthProfileNotFound {
	return &IdentityAccessManagementGetAuthProfileNotFound{}
}

/*
IdentityAccessManagementGetAuthProfileNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type IdentityAccessManagementGetAuthProfileNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management get auth profile not found response has a 2xx status code
func (o *IdentityAccessManagementGetAuthProfileNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management get auth profile not found response has a 3xx status code
func (o *IdentityAccessManagementGetAuthProfileNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get auth profile not found response has a 4xx status code
func (o *IdentityAccessManagementGetAuthProfileNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management get auth profile not found response has a 5xx status code
func (o *IdentityAccessManagementGetAuthProfileNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management get auth profile not found response a status code equal to that given
func (o *IdentityAccessManagementGetAuthProfileNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *IdentityAccessManagementGetAuthProfileNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/authorization/profiles/id/{id}][%d] identityAccessManagementGetAuthProfileNotFound  %+v", 404, o.Payload)
}

func (o *IdentityAccessManagementGetAuthProfileNotFound) String() string {
	return fmt.Sprintf("[GET /v1/authorization/profiles/id/{id}][%d] identityAccessManagementGetAuthProfileNotFound  %+v", 404, o.Payload)
}

func (o *IdentityAccessManagementGetAuthProfileNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementGetAuthProfileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetAuthProfileInternalServerError creates a IdentityAccessManagementGetAuthProfileInternalServerError with default headers values
func NewIdentityAccessManagementGetAuthProfileInternalServerError() *IdentityAccessManagementGetAuthProfileInternalServerError {
	return &IdentityAccessManagementGetAuthProfileInternalServerError{}
}

/*
IdentityAccessManagementGetAuthProfileInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type IdentityAccessManagementGetAuthProfileInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management get auth profile internal server error response has a 2xx status code
func (o *IdentityAccessManagementGetAuthProfileInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management get auth profile internal server error response has a 3xx status code
func (o *IdentityAccessManagementGetAuthProfileInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get auth profile internal server error response has a 4xx status code
func (o *IdentityAccessManagementGetAuthProfileInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management get auth profile internal server error response has a 5xx status code
func (o *IdentityAccessManagementGetAuthProfileInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management get auth profile internal server error response a status code equal to that given
func (o *IdentityAccessManagementGetAuthProfileInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *IdentityAccessManagementGetAuthProfileInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/authorization/profiles/id/{id}][%d] identityAccessManagementGetAuthProfileInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementGetAuthProfileInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/authorization/profiles/id/{id}][%d] identityAccessManagementGetAuthProfileInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementGetAuthProfileInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementGetAuthProfileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetAuthProfileGatewayTimeout creates a IdentityAccessManagementGetAuthProfileGatewayTimeout with default headers values
func NewIdentityAccessManagementGetAuthProfileGatewayTimeout() *IdentityAccessManagementGetAuthProfileGatewayTimeout {
	return &IdentityAccessManagementGetAuthProfileGatewayTimeout{}
}

/*
IdentityAccessManagementGetAuthProfileGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type IdentityAccessManagementGetAuthProfileGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management get auth profile gateway timeout response has a 2xx status code
func (o *IdentityAccessManagementGetAuthProfileGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management get auth profile gateway timeout response has a 3xx status code
func (o *IdentityAccessManagementGetAuthProfileGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get auth profile gateway timeout response has a 4xx status code
func (o *IdentityAccessManagementGetAuthProfileGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management get auth profile gateway timeout response has a 5xx status code
func (o *IdentityAccessManagementGetAuthProfileGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management get auth profile gateway timeout response a status code equal to that given
func (o *IdentityAccessManagementGetAuthProfileGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *IdentityAccessManagementGetAuthProfileGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/authorization/profiles/id/{id}][%d] identityAccessManagementGetAuthProfileGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementGetAuthProfileGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/authorization/profiles/id/{id}][%d] identityAccessManagementGetAuthProfileGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementGetAuthProfileGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementGetAuthProfileGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetAuthProfileDefault creates a IdentityAccessManagementGetAuthProfileDefault with default headers values
func NewIdentityAccessManagementGetAuthProfileDefault(code int) *IdentityAccessManagementGetAuthProfileDefault {
	return &IdentityAccessManagementGetAuthProfileDefault{
		_statusCode: code,
	}
}

/*
IdentityAccessManagementGetAuthProfileDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type IdentityAccessManagementGetAuthProfileDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the identity access management get auth profile default response
func (o *IdentityAccessManagementGetAuthProfileDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this identity access management get auth profile default response has a 2xx status code
func (o *IdentityAccessManagementGetAuthProfileDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this identity access management get auth profile default response has a 3xx status code
func (o *IdentityAccessManagementGetAuthProfileDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this identity access management get auth profile default response has a 4xx status code
func (o *IdentityAccessManagementGetAuthProfileDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this identity access management get auth profile default response has a 5xx status code
func (o *IdentityAccessManagementGetAuthProfileDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this identity access management get auth profile default response a status code equal to that given
func (o *IdentityAccessManagementGetAuthProfileDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IdentityAccessManagementGetAuthProfileDefault) Error() string {
	return fmt.Sprintf("[GET /v1/authorization/profiles/id/{id}][%d] IdentityAccessManagement_GetAuthProfile default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementGetAuthProfileDefault) String() string {
	return fmt.Sprintf("[GET /v1/authorization/profiles/id/{id}][%d] IdentityAccessManagement_GetAuthProfile default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementGetAuthProfileDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *IdentityAccessManagementGetAuthProfileDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

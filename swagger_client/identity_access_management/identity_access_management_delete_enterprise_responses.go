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

// IdentityAccessManagementDeleteEnterpriseReader is a Reader for the IdentityAccessManagementDeleteEnterprise structure.
type IdentityAccessManagementDeleteEnterpriseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IdentityAccessManagementDeleteEnterpriseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIdentityAccessManagementDeleteEnterpriseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewIdentityAccessManagementDeleteEnterpriseUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewIdentityAccessManagementDeleteEnterpriseForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewIdentityAccessManagementDeleteEnterpriseNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewIdentityAccessManagementDeleteEnterpriseConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewIdentityAccessManagementDeleteEnterpriseInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewIdentityAccessManagementDeleteEnterpriseGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIdentityAccessManagementDeleteEnterpriseDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIdentityAccessManagementDeleteEnterpriseOK creates a IdentityAccessManagementDeleteEnterpriseOK with default headers values
func NewIdentityAccessManagementDeleteEnterpriseOK() *IdentityAccessManagementDeleteEnterpriseOK {
	return &IdentityAccessManagementDeleteEnterpriseOK{}
}

/*
IdentityAccessManagementDeleteEnterpriseOK describes a response with status code 200, with default header values.

A successful response.
*/
type IdentityAccessManagementDeleteEnterpriseOK struct {
	Payload *swagger_models.CrudResponse
}

// IsSuccess returns true when this identity access management delete enterprise o k response has a 2xx status code
func (o *IdentityAccessManagementDeleteEnterpriseOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this identity access management delete enterprise o k response has a 3xx status code
func (o *IdentityAccessManagementDeleteEnterpriseOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management delete enterprise o k response has a 4xx status code
func (o *IdentityAccessManagementDeleteEnterpriseOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management delete enterprise o k response has a 5xx status code
func (o *IdentityAccessManagementDeleteEnterpriseOK) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management delete enterprise o k response a status code equal to that given
func (o *IdentityAccessManagementDeleteEnterpriseOK) IsCode(code int) bool {
	return code == 200
}

func (o *IdentityAccessManagementDeleteEnterpriseOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/enterprises/id/{id}][%d] identityAccessManagementDeleteEnterpriseOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementDeleteEnterpriseOK) String() string {
	return fmt.Sprintf("[DELETE /v1/enterprises/id/{id}][%d] identityAccessManagementDeleteEnterpriseOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementDeleteEnterpriseOK) GetPayload() *swagger_models.CrudResponse {
	return o.Payload
}

func (o *IdentityAccessManagementDeleteEnterpriseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.CrudResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementDeleteEnterpriseUnauthorized creates a IdentityAccessManagementDeleteEnterpriseUnauthorized with default headers values
func NewIdentityAccessManagementDeleteEnterpriseUnauthorized() *IdentityAccessManagementDeleteEnterpriseUnauthorized {
	return &IdentityAccessManagementDeleteEnterpriseUnauthorized{}
}

/*
IdentityAccessManagementDeleteEnterpriseUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type IdentityAccessManagementDeleteEnterpriseUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management delete enterprise unauthorized response has a 2xx status code
func (o *IdentityAccessManagementDeleteEnterpriseUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management delete enterprise unauthorized response has a 3xx status code
func (o *IdentityAccessManagementDeleteEnterpriseUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management delete enterprise unauthorized response has a 4xx status code
func (o *IdentityAccessManagementDeleteEnterpriseUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management delete enterprise unauthorized response has a 5xx status code
func (o *IdentityAccessManagementDeleteEnterpriseUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management delete enterprise unauthorized response a status code equal to that given
func (o *IdentityAccessManagementDeleteEnterpriseUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *IdentityAccessManagementDeleteEnterpriseUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v1/enterprises/id/{id}][%d] identityAccessManagementDeleteEnterpriseUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementDeleteEnterpriseUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /v1/enterprises/id/{id}][%d] identityAccessManagementDeleteEnterpriseUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementDeleteEnterpriseUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementDeleteEnterpriseUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementDeleteEnterpriseForbidden creates a IdentityAccessManagementDeleteEnterpriseForbidden with default headers values
func NewIdentityAccessManagementDeleteEnterpriseForbidden() *IdentityAccessManagementDeleteEnterpriseForbidden {
	return &IdentityAccessManagementDeleteEnterpriseForbidden{}
}

/*
IdentityAccessManagementDeleteEnterpriseForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type IdentityAccessManagementDeleteEnterpriseForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management delete enterprise forbidden response has a 2xx status code
func (o *IdentityAccessManagementDeleteEnterpriseForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management delete enterprise forbidden response has a 3xx status code
func (o *IdentityAccessManagementDeleteEnterpriseForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management delete enterprise forbidden response has a 4xx status code
func (o *IdentityAccessManagementDeleteEnterpriseForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management delete enterprise forbidden response has a 5xx status code
func (o *IdentityAccessManagementDeleteEnterpriseForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management delete enterprise forbidden response a status code equal to that given
func (o *IdentityAccessManagementDeleteEnterpriseForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *IdentityAccessManagementDeleteEnterpriseForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/enterprises/id/{id}][%d] identityAccessManagementDeleteEnterpriseForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementDeleteEnterpriseForbidden) String() string {
	return fmt.Sprintf("[DELETE /v1/enterprises/id/{id}][%d] identityAccessManagementDeleteEnterpriseForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementDeleteEnterpriseForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementDeleteEnterpriseForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementDeleteEnterpriseNotFound creates a IdentityAccessManagementDeleteEnterpriseNotFound with default headers values
func NewIdentityAccessManagementDeleteEnterpriseNotFound() *IdentityAccessManagementDeleteEnterpriseNotFound {
	return &IdentityAccessManagementDeleteEnterpriseNotFound{}
}

/*
IdentityAccessManagementDeleteEnterpriseNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type IdentityAccessManagementDeleteEnterpriseNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management delete enterprise not found response has a 2xx status code
func (o *IdentityAccessManagementDeleteEnterpriseNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management delete enterprise not found response has a 3xx status code
func (o *IdentityAccessManagementDeleteEnterpriseNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management delete enterprise not found response has a 4xx status code
func (o *IdentityAccessManagementDeleteEnterpriseNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management delete enterprise not found response has a 5xx status code
func (o *IdentityAccessManagementDeleteEnterpriseNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management delete enterprise not found response a status code equal to that given
func (o *IdentityAccessManagementDeleteEnterpriseNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *IdentityAccessManagementDeleteEnterpriseNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/enterprises/id/{id}][%d] identityAccessManagementDeleteEnterpriseNotFound  %+v", 404, o.Payload)
}

func (o *IdentityAccessManagementDeleteEnterpriseNotFound) String() string {
	return fmt.Sprintf("[DELETE /v1/enterprises/id/{id}][%d] identityAccessManagementDeleteEnterpriseNotFound  %+v", 404, o.Payload)
}

func (o *IdentityAccessManagementDeleteEnterpriseNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementDeleteEnterpriseNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementDeleteEnterpriseConflict creates a IdentityAccessManagementDeleteEnterpriseConflict with default headers values
func NewIdentityAccessManagementDeleteEnterpriseConflict() *IdentityAccessManagementDeleteEnterpriseConflict {
	return &IdentityAccessManagementDeleteEnterpriseConflict{}
}

/*
IdentityAccessManagementDeleteEnterpriseConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because there are IAM users of this IAM role
*/
type IdentityAccessManagementDeleteEnterpriseConflict struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management delete enterprise conflict response has a 2xx status code
func (o *IdentityAccessManagementDeleteEnterpriseConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management delete enterprise conflict response has a 3xx status code
func (o *IdentityAccessManagementDeleteEnterpriseConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management delete enterprise conflict response has a 4xx status code
func (o *IdentityAccessManagementDeleteEnterpriseConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management delete enterprise conflict response has a 5xx status code
func (o *IdentityAccessManagementDeleteEnterpriseConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management delete enterprise conflict response a status code equal to that given
func (o *IdentityAccessManagementDeleteEnterpriseConflict) IsCode(code int) bool {
	return code == 409
}

func (o *IdentityAccessManagementDeleteEnterpriseConflict) Error() string {
	return fmt.Sprintf("[DELETE /v1/enterprises/id/{id}][%d] identityAccessManagementDeleteEnterpriseConflict  %+v", 409, o.Payload)
}

func (o *IdentityAccessManagementDeleteEnterpriseConflict) String() string {
	return fmt.Sprintf("[DELETE /v1/enterprises/id/{id}][%d] identityAccessManagementDeleteEnterpriseConflict  %+v", 409, o.Payload)
}

func (o *IdentityAccessManagementDeleteEnterpriseConflict) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementDeleteEnterpriseConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementDeleteEnterpriseInternalServerError creates a IdentityAccessManagementDeleteEnterpriseInternalServerError with default headers values
func NewIdentityAccessManagementDeleteEnterpriseInternalServerError() *IdentityAccessManagementDeleteEnterpriseInternalServerError {
	return &IdentityAccessManagementDeleteEnterpriseInternalServerError{}
}

/*
IdentityAccessManagementDeleteEnterpriseInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type IdentityAccessManagementDeleteEnterpriseInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management delete enterprise internal server error response has a 2xx status code
func (o *IdentityAccessManagementDeleteEnterpriseInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management delete enterprise internal server error response has a 3xx status code
func (o *IdentityAccessManagementDeleteEnterpriseInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management delete enterprise internal server error response has a 4xx status code
func (o *IdentityAccessManagementDeleteEnterpriseInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management delete enterprise internal server error response has a 5xx status code
func (o *IdentityAccessManagementDeleteEnterpriseInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management delete enterprise internal server error response a status code equal to that given
func (o *IdentityAccessManagementDeleteEnterpriseInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *IdentityAccessManagementDeleteEnterpriseInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/enterprises/id/{id}][%d] identityAccessManagementDeleteEnterpriseInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementDeleteEnterpriseInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /v1/enterprises/id/{id}][%d] identityAccessManagementDeleteEnterpriseInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementDeleteEnterpriseInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementDeleteEnterpriseInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementDeleteEnterpriseGatewayTimeout creates a IdentityAccessManagementDeleteEnterpriseGatewayTimeout with default headers values
func NewIdentityAccessManagementDeleteEnterpriseGatewayTimeout() *IdentityAccessManagementDeleteEnterpriseGatewayTimeout {
	return &IdentityAccessManagementDeleteEnterpriseGatewayTimeout{}
}

/*
IdentityAccessManagementDeleteEnterpriseGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type IdentityAccessManagementDeleteEnterpriseGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management delete enterprise gateway timeout response has a 2xx status code
func (o *IdentityAccessManagementDeleteEnterpriseGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management delete enterprise gateway timeout response has a 3xx status code
func (o *IdentityAccessManagementDeleteEnterpriseGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management delete enterprise gateway timeout response has a 4xx status code
func (o *IdentityAccessManagementDeleteEnterpriseGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management delete enterprise gateway timeout response has a 5xx status code
func (o *IdentityAccessManagementDeleteEnterpriseGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management delete enterprise gateway timeout response a status code equal to that given
func (o *IdentityAccessManagementDeleteEnterpriseGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *IdentityAccessManagementDeleteEnterpriseGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /v1/enterprises/id/{id}][%d] identityAccessManagementDeleteEnterpriseGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementDeleteEnterpriseGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /v1/enterprises/id/{id}][%d] identityAccessManagementDeleteEnterpriseGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementDeleteEnterpriseGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementDeleteEnterpriseGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementDeleteEnterpriseDefault creates a IdentityAccessManagementDeleteEnterpriseDefault with default headers values
func NewIdentityAccessManagementDeleteEnterpriseDefault(code int) *IdentityAccessManagementDeleteEnterpriseDefault {
	return &IdentityAccessManagementDeleteEnterpriseDefault{
		_statusCode: code,
	}
}

/*
IdentityAccessManagementDeleteEnterpriseDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type IdentityAccessManagementDeleteEnterpriseDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the identity access management delete enterprise default response
func (o *IdentityAccessManagementDeleteEnterpriseDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this identity access management delete enterprise default response has a 2xx status code
func (o *IdentityAccessManagementDeleteEnterpriseDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this identity access management delete enterprise default response has a 3xx status code
func (o *IdentityAccessManagementDeleteEnterpriseDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this identity access management delete enterprise default response has a 4xx status code
func (o *IdentityAccessManagementDeleteEnterpriseDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this identity access management delete enterprise default response has a 5xx status code
func (o *IdentityAccessManagementDeleteEnterpriseDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this identity access management delete enterprise default response a status code equal to that given
func (o *IdentityAccessManagementDeleteEnterpriseDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IdentityAccessManagementDeleteEnterpriseDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/enterprises/id/{id}][%d] IdentityAccessManagement_DeleteEnterprise default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementDeleteEnterpriseDefault) String() string {
	return fmt.Sprintf("[DELETE /v1/enterprises/id/{id}][%d] IdentityAccessManagement_DeleteEnterprise default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementDeleteEnterpriseDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *IdentityAccessManagementDeleteEnterpriseDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

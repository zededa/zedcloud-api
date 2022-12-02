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

// IdentityAccessManagementQueryEnterprisesReader is a Reader for the IdentityAccessManagementQueryEnterprises structure.
type IdentityAccessManagementQueryEnterprisesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IdentityAccessManagementQueryEnterprisesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIdentityAccessManagementQueryEnterprisesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewIdentityAccessManagementQueryEnterprisesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewIdentityAccessManagementQueryEnterprisesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewIdentityAccessManagementQueryEnterprisesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewIdentityAccessManagementQueryEnterprisesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewIdentityAccessManagementQueryEnterprisesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIdentityAccessManagementQueryEnterprisesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIdentityAccessManagementQueryEnterprisesOK creates a IdentityAccessManagementQueryEnterprisesOK with default headers values
func NewIdentityAccessManagementQueryEnterprisesOK() *IdentityAccessManagementQueryEnterprisesOK {
	return &IdentityAccessManagementQueryEnterprisesOK{}
}

/*
IdentityAccessManagementQueryEnterprisesOK describes a response with status code 200, with default header values.

A successful response.
*/
type IdentityAccessManagementQueryEnterprisesOK struct {
	Payload *swagger_models.CrudResponse
}

// IsSuccess returns true when this identity access management query enterprises o k response has a 2xx status code
func (o *IdentityAccessManagementQueryEnterprisesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this identity access management query enterprises o k response has a 3xx status code
func (o *IdentityAccessManagementQueryEnterprisesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management query enterprises o k response has a 4xx status code
func (o *IdentityAccessManagementQueryEnterprisesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management query enterprises o k response has a 5xx status code
func (o *IdentityAccessManagementQueryEnterprisesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management query enterprises o k response a status code equal to that given
func (o *IdentityAccessManagementQueryEnterprisesOK) IsCode(code int) bool {
	return code == 200
}

func (o *IdentityAccessManagementQueryEnterprisesOK) Error() string {
	return fmt.Sprintf("[GET /v1/enterprises][%d] identityAccessManagementQueryEnterprisesOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementQueryEnterprisesOK) String() string {
	return fmt.Sprintf("[GET /v1/enterprises][%d] identityAccessManagementQueryEnterprisesOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementQueryEnterprisesOK) GetPayload() *swagger_models.CrudResponse {
	return o.Payload
}

func (o *IdentityAccessManagementQueryEnterprisesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.CrudResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementQueryEnterprisesBadRequest creates a IdentityAccessManagementQueryEnterprisesBadRequest with default headers values
func NewIdentityAccessManagementQueryEnterprisesBadRequest() *IdentityAccessManagementQueryEnterprisesBadRequest {
	return &IdentityAccessManagementQueryEnterprisesBadRequest{}
}

/*
IdentityAccessManagementQueryEnterprisesBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of invalid value of filter parameters.
*/
type IdentityAccessManagementQueryEnterprisesBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management query enterprises bad request response has a 2xx status code
func (o *IdentityAccessManagementQueryEnterprisesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management query enterprises bad request response has a 3xx status code
func (o *IdentityAccessManagementQueryEnterprisesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management query enterprises bad request response has a 4xx status code
func (o *IdentityAccessManagementQueryEnterprisesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management query enterprises bad request response has a 5xx status code
func (o *IdentityAccessManagementQueryEnterprisesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management query enterprises bad request response a status code equal to that given
func (o *IdentityAccessManagementQueryEnterprisesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *IdentityAccessManagementQueryEnterprisesBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/enterprises][%d] identityAccessManagementQueryEnterprisesBadRequest  %+v", 400, o.Payload)
}

func (o *IdentityAccessManagementQueryEnterprisesBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/enterprises][%d] identityAccessManagementQueryEnterprisesBadRequest  %+v", 400, o.Payload)
}

func (o *IdentityAccessManagementQueryEnterprisesBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementQueryEnterprisesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementQueryEnterprisesUnauthorized creates a IdentityAccessManagementQueryEnterprisesUnauthorized with default headers values
func NewIdentityAccessManagementQueryEnterprisesUnauthorized() *IdentityAccessManagementQueryEnterprisesUnauthorized {
	return &IdentityAccessManagementQueryEnterprisesUnauthorized{}
}

/*
IdentityAccessManagementQueryEnterprisesUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type IdentityAccessManagementQueryEnterprisesUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management query enterprises unauthorized response has a 2xx status code
func (o *IdentityAccessManagementQueryEnterprisesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management query enterprises unauthorized response has a 3xx status code
func (o *IdentityAccessManagementQueryEnterprisesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management query enterprises unauthorized response has a 4xx status code
func (o *IdentityAccessManagementQueryEnterprisesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management query enterprises unauthorized response has a 5xx status code
func (o *IdentityAccessManagementQueryEnterprisesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management query enterprises unauthorized response a status code equal to that given
func (o *IdentityAccessManagementQueryEnterprisesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *IdentityAccessManagementQueryEnterprisesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/enterprises][%d] identityAccessManagementQueryEnterprisesUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementQueryEnterprisesUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/enterprises][%d] identityAccessManagementQueryEnterprisesUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementQueryEnterprisesUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementQueryEnterprisesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementQueryEnterprisesForbidden creates a IdentityAccessManagementQueryEnterprisesForbidden with default headers values
func NewIdentityAccessManagementQueryEnterprisesForbidden() *IdentityAccessManagementQueryEnterprisesForbidden {
	return &IdentityAccessManagementQueryEnterprisesForbidden{}
}

/*
IdentityAccessManagementQueryEnterprisesForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type IdentityAccessManagementQueryEnterprisesForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management query enterprises forbidden response has a 2xx status code
func (o *IdentityAccessManagementQueryEnterprisesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management query enterprises forbidden response has a 3xx status code
func (o *IdentityAccessManagementQueryEnterprisesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management query enterprises forbidden response has a 4xx status code
func (o *IdentityAccessManagementQueryEnterprisesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management query enterprises forbidden response has a 5xx status code
func (o *IdentityAccessManagementQueryEnterprisesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management query enterprises forbidden response a status code equal to that given
func (o *IdentityAccessManagementQueryEnterprisesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *IdentityAccessManagementQueryEnterprisesForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/enterprises][%d] identityAccessManagementQueryEnterprisesForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementQueryEnterprisesForbidden) String() string {
	return fmt.Sprintf("[GET /v1/enterprises][%d] identityAccessManagementQueryEnterprisesForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementQueryEnterprisesForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementQueryEnterprisesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementQueryEnterprisesInternalServerError creates a IdentityAccessManagementQueryEnterprisesInternalServerError with default headers values
func NewIdentityAccessManagementQueryEnterprisesInternalServerError() *IdentityAccessManagementQueryEnterprisesInternalServerError {
	return &IdentityAccessManagementQueryEnterprisesInternalServerError{}
}

/*
IdentityAccessManagementQueryEnterprisesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type IdentityAccessManagementQueryEnterprisesInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management query enterprises internal server error response has a 2xx status code
func (o *IdentityAccessManagementQueryEnterprisesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management query enterprises internal server error response has a 3xx status code
func (o *IdentityAccessManagementQueryEnterprisesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management query enterprises internal server error response has a 4xx status code
func (o *IdentityAccessManagementQueryEnterprisesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management query enterprises internal server error response has a 5xx status code
func (o *IdentityAccessManagementQueryEnterprisesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management query enterprises internal server error response a status code equal to that given
func (o *IdentityAccessManagementQueryEnterprisesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *IdentityAccessManagementQueryEnterprisesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/enterprises][%d] identityAccessManagementQueryEnterprisesInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementQueryEnterprisesInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/enterprises][%d] identityAccessManagementQueryEnterprisesInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementQueryEnterprisesInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementQueryEnterprisesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementQueryEnterprisesGatewayTimeout creates a IdentityAccessManagementQueryEnterprisesGatewayTimeout with default headers values
func NewIdentityAccessManagementQueryEnterprisesGatewayTimeout() *IdentityAccessManagementQueryEnterprisesGatewayTimeout {
	return &IdentityAccessManagementQueryEnterprisesGatewayTimeout{}
}

/*
IdentityAccessManagementQueryEnterprisesGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type IdentityAccessManagementQueryEnterprisesGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management query enterprises gateway timeout response has a 2xx status code
func (o *IdentityAccessManagementQueryEnterprisesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management query enterprises gateway timeout response has a 3xx status code
func (o *IdentityAccessManagementQueryEnterprisesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management query enterprises gateway timeout response has a 4xx status code
func (o *IdentityAccessManagementQueryEnterprisesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management query enterprises gateway timeout response has a 5xx status code
func (o *IdentityAccessManagementQueryEnterprisesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management query enterprises gateway timeout response a status code equal to that given
func (o *IdentityAccessManagementQueryEnterprisesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *IdentityAccessManagementQueryEnterprisesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/enterprises][%d] identityAccessManagementQueryEnterprisesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementQueryEnterprisesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/enterprises][%d] identityAccessManagementQueryEnterprisesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementQueryEnterprisesGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementQueryEnterprisesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementQueryEnterprisesDefault creates a IdentityAccessManagementQueryEnterprisesDefault with default headers values
func NewIdentityAccessManagementQueryEnterprisesDefault(code int) *IdentityAccessManagementQueryEnterprisesDefault {
	return &IdentityAccessManagementQueryEnterprisesDefault{
		_statusCode: code,
	}
}

/*
IdentityAccessManagementQueryEnterprisesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type IdentityAccessManagementQueryEnterprisesDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the identity access management query enterprises default response
func (o *IdentityAccessManagementQueryEnterprisesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this identity access management query enterprises default response has a 2xx status code
func (o *IdentityAccessManagementQueryEnterprisesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this identity access management query enterprises default response has a 3xx status code
func (o *IdentityAccessManagementQueryEnterprisesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this identity access management query enterprises default response has a 4xx status code
func (o *IdentityAccessManagementQueryEnterprisesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this identity access management query enterprises default response has a 5xx status code
func (o *IdentityAccessManagementQueryEnterprisesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this identity access management query enterprises default response a status code equal to that given
func (o *IdentityAccessManagementQueryEnterprisesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IdentityAccessManagementQueryEnterprisesDefault) Error() string {
	return fmt.Sprintf("[GET /v1/enterprises][%d] IdentityAccessManagement_QueryEnterprises default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementQueryEnterprisesDefault) String() string {
	return fmt.Sprintf("[GET /v1/enterprises][%d] IdentityAccessManagement_QueryEnterprises default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementQueryEnterprisesDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *IdentityAccessManagementQueryEnterprisesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

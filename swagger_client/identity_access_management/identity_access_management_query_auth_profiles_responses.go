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

// IdentityAccessManagementQueryAuthProfilesReader is a Reader for the IdentityAccessManagementQueryAuthProfiles structure.
type IdentityAccessManagementQueryAuthProfilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IdentityAccessManagementQueryAuthProfilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIdentityAccessManagementQueryAuthProfilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewIdentityAccessManagementQueryAuthProfilesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewIdentityAccessManagementQueryAuthProfilesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewIdentityAccessManagementQueryAuthProfilesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewIdentityAccessManagementQueryAuthProfilesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewIdentityAccessManagementQueryAuthProfilesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIdentityAccessManagementQueryAuthProfilesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIdentityAccessManagementQueryAuthProfilesOK creates a IdentityAccessManagementQueryAuthProfilesOK with default headers values
func NewIdentityAccessManagementQueryAuthProfilesOK() *IdentityAccessManagementQueryAuthProfilesOK {
	return &IdentityAccessManagementQueryAuthProfilesOK{}
}

/*
IdentityAccessManagementQueryAuthProfilesOK describes a response with status code 200, with default header values.

A successful response.
*/
type IdentityAccessManagementQueryAuthProfilesOK struct {
	Payload *swagger_models.CrudResponse
}

// IsSuccess returns true when this identity access management query auth profiles o k response has a 2xx status code
func (o *IdentityAccessManagementQueryAuthProfilesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this identity access management query auth profiles o k response has a 3xx status code
func (o *IdentityAccessManagementQueryAuthProfilesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management query auth profiles o k response has a 4xx status code
func (o *IdentityAccessManagementQueryAuthProfilesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management query auth profiles o k response has a 5xx status code
func (o *IdentityAccessManagementQueryAuthProfilesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management query auth profiles o k response a status code equal to that given
func (o *IdentityAccessManagementQueryAuthProfilesOK) IsCode(code int) bool {
	return code == 200
}

func (o *IdentityAccessManagementQueryAuthProfilesOK) Error() string {
	return fmt.Sprintf("[GET /v1/authorization/profiles][%d] identityAccessManagementQueryAuthProfilesOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementQueryAuthProfilesOK) String() string {
	return fmt.Sprintf("[GET /v1/authorization/profiles][%d] identityAccessManagementQueryAuthProfilesOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementQueryAuthProfilesOK) GetPayload() *swagger_models.CrudResponse {
	return o.Payload
}

func (o *IdentityAccessManagementQueryAuthProfilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.CrudResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementQueryAuthProfilesBadRequest creates a IdentityAccessManagementQueryAuthProfilesBadRequest with default headers values
func NewIdentityAccessManagementQueryAuthProfilesBadRequest() *IdentityAccessManagementQueryAuthProfilesBadRequest {
	return &IdentityAccessManagementQueryAuthProfilesBadRequest{}
}

/*
IdentityAccessManagementQueryAuthProfilesBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of invalid value of filter parameters.
*/
type IdentityAccessManagementQueryAuthProfilesBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management query auth profiles bad request response has a 2xx status code
func (o *IdentityAccessManagementQueryAuthProfilesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management query auth profiles bad request response has a 3xx status code
func (o *IdentityAccessManagementQueryAuthProfilesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management query auth profiles bad request response has a 4xx status code
func (o *IdentityAccessManagementQueryAuthProfilesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management query auth profiles bad request response has a 5xx status code
func (o *IdentityAccessManagementQueryAuthProfilesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management query auth profiles bad request response a status code equal to that given
func (o *IdentityAccessManagementQueryAuthProfilesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *IdentityAccessManagementQueryAuthProfilesBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/authorization/profiles][%d] identityAccessManagementQueryAuthProfilesBadRequest  %+v", 400, o.Payload)
}

func (o *IdentityAccessManagementQueryAuthProfilesBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/authorization/profiles][%d] identityAccessManagementQueryAuthProfilesBadRequest  %+v", 400, o.Payload)
}

func (o *IdentityAccessManagementQueryAuthProfilesBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementQueryAuthProfilesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementQueryAuthProfilesUnauthorized creates a IdentityAccessManagementQueryAuthProfilesUnauthorized with default headers values
func NewIdentityAccessManagementQueryAuthProfilesUnauthorized() *IdentityAccessManagementQueryAuthProfilesUnauthorized {
	return &IdentityAccessManagementQueryAuthProfilesUnauthorized{}
}

/*
IdentityAccessManagementQueryAuthProfilesUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type IdentityAccessManagementQueryAuthProfilesUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management query auth profiles unauthorized response has a 2xx status code
func (o *IdentityAccessManagementQueryAuthProfilesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management query auth profiles unauthorized response has a 3xx status code
func (o *IdentityAccessManagementQueryAuthProfilesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management query auth profiles unauthorized response has a 4xx status code
func (o *IdentityAccessManagementQueryAuthProfilesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management query auth profiles unauthorized response has a 5xx status code
func (o *IdentityAccessManagementQueryAuthProfilesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management query auth profiles unauthorized response a status code equal to that given
func (o *IdentityAccessManagementQueryAuthProfilesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *IdentityAccessManagementQueryAuthProfilesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/authorization/profiles][%d] identityAccessManagementQueryAuthProfilesUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementQueryAuthProfilesUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/authorization/profiles][%d] identityAccessManagementQueryAuthProfilesUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementQueryAuthProfilesUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementQueryAuthProfilesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementQueryAuthProfilesForbidden creates a IdentityAccessManagementQueryAuthProfilesForbidden with default headers values
func NewIdentityAccessManagementQueryAuthProfilesForbidden() *IdentityAccessManagementQueryAuthProfilesForbidden {
	return &IdentityAccessManagementQueryAuthProfilesForbidden{}
}

/*
IdentityAccessManagementQueryAuthProfilesForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type IdentityAccessManagementQueryAuthProfilesForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management query auth profiles forbidden response has a 2xx status code
func (o *IdentityAccessManagementQueryAuthProfilesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management query auth profiles forbidden response has a 3xx status code
func (o *IdentityAccessManagementQueryAuthProfilesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management query auth profiles forbidden response has a 4xx status code
func (o *IdentityAccessManagementQueryAuthProfilesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management query auth profiles forbidden response has a 5xx status code
func (o *IdentityAccessManagementQueryAuthProfilesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management query auth profiles forbidden response a status code equal to that given
func (o *IdentityAccessManagementQueryAuthProfilesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *IdentityAccessManagementQueryAuthProfilesForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/authorization/profiles][%d] identityAccessManagementQueryAuthProfilesForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementQueryAuthProfilesForbidden) String() string {
	return fmt.Sprintf("[GET /v1/authorization/profiles][%d] identityAccessManagementQueryAuthProfilesForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementQueryAuthProfilesForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementQueryAuthProfilesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementQueryAuthProfilesInternalServerError creates a IdentityAccessManagementQueryAuthProfilesInternalServerError with default headers values
func NewIdentityAccessManagementQueryAuthProfilesInternalServerError() *IdentityAccessManagementQueryAuthProfilesInternalServerError {
	return &IdentityAccessManagementQueryAuthProfilesInternalServerError{}
}

/*
IdentityAccessManagementQueryAuthProfilesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type IdentityAccessManagementQueryAuthProfilesInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management query auth profiles internal server error response has a 2xx status code
func (o *IdentityAccessManagementQueryAuthProfilesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management query auth profiles internal server error response has a 3xx status code
func (o *IdentityAccessManagementQueryAuthProfilesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management query auth profiles internal server error response has a 4xx status code
func (o *IdentityAccessManagementQueryAuthProfilesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management query auth profiles internal server error response has a 5xx status code
func (o *IdentityAccessManagementQueryAuthProfilesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management query auth profiles internal server error response a status code equal to that given
func (o *IdentityAccessManagementQueryAuthProfilesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *IdentityAccessManagementQueryAuthProfilesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/authorization/profiles][%d] identityAccessManagementQueryAuthProfilesInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementQueryAuthProfilesInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/authorization/profiles][%d] identityAccessManagementQueryAuthProfilesInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementQueryAuthProfilesInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementQueryAuthProfilesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementQueryAuthProfilesGatewayTimeout creates a IdentityAccessManagementQueryAuthProfilesGatewayTimeout with default headers values
func NewIdentityAccessManagementQueryAuthProfilesGatewayTimeout() *IdentityAccessManagementQueryAuthProfilesGatewayTimeout {
	return &IdentityAccessManagementQueryAuthProfilesGatewayTimeout{}
}

/*
IdentityAccessManagementQueryAuthProfilesGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type IdentityAccessManagementQueryAuthProfilesGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management query auth profiles gateway timeout response has a 2xx status code
func (o *IdentityAccessManagementQueryAuthProfilesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management query auth profiles gateway timeout response has a 3xx status code
func (o *IdentityAccessManagementQueryAuthProfilesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management query auth profiles gateway timeout response has a 4xx status code
func (o *IdentityAccessManagementQueryAuthProfilesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management query auth profiles gateway timeout response has a 5xx status code
func (o *IdentityAccessManagementQueryAuthProfilesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management query auth profiles gateway timeout response a status code equal to that given
func (o *IdentityAccessManagementQueryAuthProfilesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *IdentityAccessManagementQueryAuthProfilesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/authorization/profiles][%d] identityAccessManagementQueryAuthProfilesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementQueryAuthProfilesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/authorization/profiles][%d] identityAccessManagementQueryAuthProfilesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementQueryAuthProfilesGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementQueryAuthProfilesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementQueryAuthProfilesDefault creates a IdentityAccessManagementQueryAuthProfilesDefault with default headers values
func NewIdentityAccessManagementQueryAuthProfilesDefault(code int) *IdentityAccessManagementQueryAuthProfilesDefault {
	return &IdentityAccessManagementQueryAuthProfilesDefault{
		_statusCode: code,
	}
}

/*
IdentityAccessManagementQueryAuthProfilesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type IdentityAccessManagementQueryAuthProfilesDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the identity access management query auth profiles default response
func (o *IdentityAccessManagementQueryAuthProfilesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this identity access management query auth profiles default response has a 2xx status code
func (o *IdentityAccessManagementQueryAuthProfilesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this identity access management query auth profiles default response has a 3xx status code
func (o *IdentityAccessManagementQueryAuthProfilesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this identity access management query auth profiles default response has a 4xx status code
func (o *IdentityAccessManagementQueryAuthProfilesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this identity access management query auth profiles default response has a 5xx status code
func (o *IdentityAccessManagementQueryAuthProfilesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this identity access management query auth profiles default response a status code equal to that given
func (o *IdentityAccessManagementQueryAuthProfilesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IdentityAccessManagementQueryAuthProfilesDefault) Error() string {
	return fmt.Sprintf("[GET /v1/authorization/profiles][%d] IdentityAccessManagement_QueryAuthProfiles default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementQueryAuthProfilesDefault) String() string {
	return fmt.Sprintf("[GET /v1/authorization/profiles][%d] IdentityAccessManagement_QueryAuthProfiles default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementQueryAuthProfilesDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *IdentityAccessManagementQueryAuthProfilesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

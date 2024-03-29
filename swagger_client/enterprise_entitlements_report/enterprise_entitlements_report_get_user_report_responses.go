// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package enterprise_entitlements_report

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// EnterpriseEntitlementsReportGetUserReportReader is a Reader for the EnterpriseEntitlementsReportGetUserReport structure.
type EnterpriseEntitlementsReportGetUserReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnterpriseEntitlementsReportGetUserReportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEnterpriseEntitlementsReportGetUserReportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEnterpriseEntitlementsReportGetUserReportUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEnterpriseEntitlementsReportGetUserReportForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEnterpriseEntitlementsReportGetUserReportNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEnterpriseEntitlementsReportGetUserReportInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEnterpriseEntitlementsReportGetUserReportGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEnterpriseEntitlementsReportGetUserReportDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEnterpriseEntitlementsReportGetUserReportOK creates a EnterpriseEntitlementsReportGetUserReportOK with default headers values
func NewEnterpriseEntitlementsReportGetUserReportOK() *EnterpriseEntitlementsReportGetUserReportOK {
	return &EnterpriseEntitlementsReportGetUserReportOK{}
}

/*
EnterpriseEntitlementsReportGetUserReportOK describes a response with status code 200, with default header values.

A successful response.
*/
type EnterpriseEntitlementsReportGetUserReportOK struct {
	Payload *swagger_models.UserReport
}

// IsSuccess returns true when this enterprise entitlements report get user report o k response has a 2xx status code
func (o *EnterpriseEntitlementsReportGetUserReportOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this enterprise entitlements report get user report o k response has a 3xx status code
func (o *EnterpriseEntitlementsReportGetUserReportOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enterprise entitlements report get user report o k response has a 4xx status code
func (o *EnterpriseEntitlementsReportGetUserReportOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this enterprise entitlements report get user report o k response has a 5xx status code
func (o *EnterpriseEntitlementsReportGetUserReportOK) IsServerError() bool {
	return false
}

// IsCode returns true when this enterprise entitlements report get user report o k response a status code equal to that given
func (o *EnterpriseEntitlementsReportGetUserReportOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the enterprise entitlements report get user report o k response
func (o *EnterpriseEntitlementsReportGetUserReportOK) Code() int {
	return 200
}

func (o *EnterpriseEntitlementsReportGetUserReportOK) Error() string {
	return fmt.Sprintf("[GET /v1/reports/user][%d] enterpriseEntitlementsReportGetUserReportOK  %+v", 200, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetUserReportOK) String() string {
	return fmt.Sprintf("[GET /v1/reports/user][%d] enterpriseEntitlementsReportGetUserReportOK  %+v", 200, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetUserReportOK) GetPayload() *swagger_models.UserReport {
	return o.Payload
}

func (o *EnterpriseEntitlementsReportGetUserReportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.UserReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnterpriseEntitlementsReportGetUserReportUnauthorized creates a EnterpriseEntitlementsReportGetUserReportUnauthorized with default headers values
func NewEnterpriseEntitlementsReportGetUserReportUnauthorized() *EnterpriseEntitlementsReportGetUserReportUnauthorized {
	return &EnterpriseEntitlementsReportGetUserReportUnauthorized{}
}

/*
EnterpriseEntitlementsReportGetUserReportUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EnterpriseEntitlementsReportGetUserReportUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this enterprise entitlements report get user report unauthorized response has a 2xx status code
func (o *EnterpriseEntitlementsReportGetUserReportUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enterprise entitlements report get user report unauthorized response has a 3xx status code
func (o *EnterpriseEntitlementsReportGetUserReportUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enterprise entitlements report get user report unauthorized response has a 4xx status code
func (o *EnterpriseEntitlementsReportGetUserReportUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this enterprise entitlements report get user report unauthorized response has a 5xx status code
func (o *EnterpriseEntitlementsReportGetUserReportUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this enterprise entitlements report get user report unauthorized response a status code equal to that given
func (o *EnterpriseEntitlementsReportGetUserReportUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the enterprise entitlements report get user report unauthorized response
func (o *EnterpriseEntitlementsReportGetUserReportUnauthorized) Code() int {
	return 401
}

func (o *EnterpriseEntitlementsReportGetUserReportUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/reports/user][%d] enterpriseEntitlementsReportGetUserReportUnauthorized  %+v", 401, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetUserReportUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/reports/user][%d] enterpriseEntitlementsReportGetUserReportUnauthorized  %+v", 401, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetUserReportUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EnterpriseEntitlementsReportGetUserReportUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnterpriseEntitlementsReportGetUserReportForbidden creates a EnterpriseEntitlementsReportGetUserReportForbidden with default headers values
func NewEnterpriseEntitlementsReportGetUserReportForbidden() *EnterpriseEntitlementsReportGetUserReportForbidden {
	return &EnterpriseEntitlementsReportGetUserReportForbidden{}
}

/*
EnterpriseEntitlementsReportGetUserReportForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EnterpriseEntitlementsReportGetUserReportForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this enterprise entitlements report get user report forbidden response has a 2xx status code
func (o *EnterpriseEntitlementsReportGetUserReportForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enterprise entitlements report get user report forbidden response has a 3xx status code
func (o *EnterpriseEntitlementsReportGetUserReportForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enterprise entitlements report get user report forbidden response has a 4xx status code
func (o *EnterpriseEntitlementsReportGetUserReportForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this enterprise entitlements report get user report forbidden response has a 5xx status code
func (o *EnterpriseEntitlementsReportGetUserReportForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this enterprise entitlements report get user report forbidden response a status code equal to that given
func (o *EnterpriseEntitlementsReportGetUserReportForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the enterprise entitlements report get user report forbidden response
func (o *EnterpriseEntitlementsReportGetUserReportForbidden) Code() int {
	return 403
}

func (o *EnterpriseEntitlementsReportGetUserReportForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/reports/user][%d] enterpriseEntitlementsReportGetUserReportForbidden  %+v", 403, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetUserReportForbidden) String() string {
	return fmt.Sprintf("[GET /v1/reports/user][%d] enterpriseEntitlementsReportGetUserReportForbidden  %+v", 403, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetUserReportForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EnterpriseEntitlementsReportGetUserReportForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnterpriseEntitlementsReportGetUserReportNotFound creates a EnterpriseEntitlementsReportGetUserReportNotFound with default headers values
func NewEnterpriseEntitlementsReportGetUserReportNotFound() *EnterpriseEntitlementsReportGetUserReportNotFound {
	return &EnterpriseEntitlementsReportGetUserReportNotFound{}
}

/*
EnterpriseEntitlementsReportGetUserReportNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EnterpriseEntitlementsReportGetUserReportNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this enterprise entitlements report get user report not found response has a 2xx status code
func (o *EnterpriseEntitlementsReportGetUserReportNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enterprise entitlements report get user report not found response has a 3xx status code
func (o *EnterpriseEntitlementsReportGetUserReportNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enterprise entitlements report get user report not found response has a 4xx status code
func (o *EnterpriseEntitlementsReportGetUserReportNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this enterprise entitlements report get user report not found response has a 5xx status code
func (o *EnterpriseEntitlementsReportGetUserReportNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this enterprise entitlements report get user report not found response a status code equal to that given
func (o *EnterpriseEntitlementsReportGetUserReportNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the enterprise entitlements report get user report not found response
func (o *EnterpriseEntitlementsReportGetUserReportNotFound) Code() int {
	return 404
}

func (o *EnterpriseEntitlementsReportGetUserReportNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/reports/user][%d] enterpriseEntitlementsReportGetUserReportNotFound  %+v", 404, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetUserReportNotFound) String() string {
	return fmt.Sprintf("[GET /v1/reports/user][%d] enterpriseEntitlementsReportGetUserReportNotFound  %+v", 404, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetUserReportNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EnterpriseEntitlementsReportGetUserReportNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnterpriseEntitlementsReportGetUserReportInternalServerError creates a EnterpriseEntitlementsReportGetUserReportInternalServerError with default headers values
func NewEnterpriseEntitlementsReportGetUserReportInternalServerError() *EnterpriseEntitlementsReportGetUserReportInternalServerError {
	return &EnterpriseEntitlementsReportGetUserReportInternalServerError{}
}

/*
EnterpriseEntitlementsReportGetUserReportInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EnterpriseEntitlementsReportGetUserReportInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this enterprise entitlements report get user report internal server error response has a 2xx status code
func (o *EnterpriseEntitlementsReportGetUserReportInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enterprise entitlements report get user report internal server error response has a 3xx status code
func (o *EnterpriseEntitlementsReportGetUserReportInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enterprise entitlements report get user report internal server error response has a 4xx status code
func (o *EnterpriseEntitlementsReportGetUserReportInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this enterprise entitlements report get user report internal server error response has a 5xx status code
func (o *EnterpriseEntitlementsReportGetUserReportInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this enterprise entitlements report get user report internal server error response a status code equal to that given
func (o *EnterpriseEntitlementsReportGetUserReportInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the enterprise entitlements report get user report internal server error response
func (o *EnterpriseEntitlementsReportGetUserReportInternalServerError) Code() int {
	return 500
}

func (o *EnterpriseEntitlementsReportGetUserReportInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/reports/user][%d] enterpriseEntitlementsReportGetUserReportInternalServerError  %+v", 500, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetUserReportInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/reports/user][%d] enterpriseEntitlementsReportGetUserReportInternalServerError  %+v", 500, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetUserReportInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EnterpriseEntitlementsReportGetUserReportInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnterpriseEntitlementsReportGetUserReportGatewayTimeout creates a EnterpriseEntitlementsReportGetUserReportGatewayTimeout with default headers values
func NewEnterpriseEntitlementsReportGetUserReportGatewayTimeout() *EnterpriseEntitlementsReportGetUserReportGatewayTimeout {
	return &EnterpriseEntitlementsReportGetUserReportGatewayTimeout{}
}

/*
EnterpriseEntitlementsReportGetUserReportGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EnterpriseEntitlementsReportGetUserReportGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this enterprise entitlements report get user report gateway timeout response has a 2xx status code
func (o *EnterpriseEntitlementsReportGetUserReportGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enterprise entitlements report get user report gateway timeout response has a 3xx status code
func (o *EnterpriseEntitlementsReportGetUserReportGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enterprise entitlements report get user report gateway timeout response has a 4xx status code
func (o *EnterpriseEntitlementsReportGetUserReportGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this enterprise entitlements report get user report gateway timeout response has a 5xx status code
func (o *EnterpriseEntitlementsReportGetUserReportGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this enterprise entitlements report get user report gateway timeout response a status code equal to that given
func (o *EnterpriseEntitlementsReportGetUserReportGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the enterprise entitlements report get user report gateway timeout response
func (o *EnterpriseEntitlementsReportGetUserReportGatewayTimeout) Code() int {
	return 504
}

func (o *EnterpriseEntitlementsReportGetUserReportGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/reports/user][%d] enterpriseEntitlementsReportGetUserReportGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetUserReportGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/reports/user][%d] enterpriseEntitlementsReportGetUserReportGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetUserReportGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EnterpriseEntitlementsReportGetUserReportGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnterpriseEntitlementsReportGetUserReportDefault creates a EnterpriseEntitlementsReportGetUserReportDefault with default headers values
func NewEnterpriseEntitlementsReportGetUserReportDefault(code int) *EnterpriseEntitlementsReportGetUserReportDefault {
	return &EnterpriseEntitlementsReportGetUserReportDefault{
		_statusCode: code,
	}
}

/*
EnterpriseEntitlementsReportGetUserReportDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EnterpriseEntitlementsReportGetUserReportDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// IsSuccess returns true when this enterprise entitlements report get user report default response has a 2xx status code
func (o *EnterpriseEntitlementsReportGetUserReportDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this enterprise entitlements report get user report default response has a 3xx status code
func (o *EnterpriseEntitlementsReportGetUserReportDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this enterprise entitlements report get user report default response has a 4xx status code
func (o *EnterpriseEntitlementsReportGetUserReportDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this enterprise entitlements report get user report default response has a 5xx status code
func (o *EnterpriseEntitlementsReportGetUserReportDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this enterprise entitlements report get user report default response a status code equal to that given
func (o *EnterpriseEntitlementsReportGetUserReportDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the enterprise entitlements report get user report default response
func (o *EnterpriseEntitlementsReportGetUserReportDefault) Code() int {
	return o._statusCode
}

func (o *EnterpriseEntitlementsReportGetUserReportDefault) Error() string {
	return fmt.Sprintf("[GET /v1/reports/user][%d] EnterpriseEntitlementsReport_GetUserReport default  %+v", o._statusCode, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetUserReportDefault) String() string {
	return fmt.Sprintf("[GET /v1/reports/user][%d] EnterpriseEntitlementsReport_GetUserReport default  %+v", o._statusCode, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetUserReportDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *EnterpriseEntitlementsReportGetUserReportDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

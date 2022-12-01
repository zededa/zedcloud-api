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

// EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsReader is a Reader for the EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlements structure.
type EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsOK creates a EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsOK with default headers values
func NewEnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsOK() *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsOK {
	return &EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsOK{}
}

/*
EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsOK describes a response with status code 200, with default header values.

A successful response.
*/
type EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsOK struct {
	Payload *swagger_models.CrudResponse
}

// IsSuccess returns true when this enterprise entitlements report get allowed enterprises for entitlements o k response has a 2xx status code
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this enterprise entitlements report get allowed enterprises for entitlements o k response has a 3xx status code
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enterprise entitlements report get allowed enterprises for entitlements o k response has a 4xx status code
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this enterprise entitlements report get allowed enterprises for entitlements o k response has a 5xx status code
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this enterprise entitlements report get allowed enterprises for entitlements o k response a status code equal to that given
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsOK) IsCode(code int) bool {
	return code == 200
}

func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsOK) Error() string {
	return fmt.Sprintf("[GET /v1/entitlements/allowedenterprises][%d] enterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsOK  %+v", 200, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsOK) String() string {
	return fmt.Sprintf("[GET /v1/entitlements/allowedenterprises][%d] enterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsOK  %+v", 200, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsOK) GetPayload() *swagger_models.CrudResponse {
	return o.Payload
}

func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.CrudResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsUnauthorized creates a EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsUnauthorized with default headers values
func NewEnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsUnauthorized() *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsUnauthorized {
	return &EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsUnauthorized{}
}

/*
EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this enterprise entitlements report get allowed enterprises for entitlements unauthorized response has a 2xx status code
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enterprise entitlements report get allowed enterprises for entitlements unauthorized response has a 3xx status code
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enterprise entitlements report get allowed enterprises for entitlements unauthorized response has a 4xx status code
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this enterprise entitlements report get allowed enterprises for entitlements unauthorized response has a 5xx status code
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this enterprise entitlements report get allowed enterprises for entitlements unauthorized response a status code equal to that given
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/entitlements/allowedenterprises][%d] enterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsUnauthorized  %+v", 401, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/entitlements/allowedenterprises][%d] enterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsUnauthorized  %+v", 401, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsForbidden creates a EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsForbidden with default headers values
func NewEnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsForbidden() *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsForbidden {
	return &EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsForbidden{}
}

/*
EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this enterprise entitlements report get allowed enterprises for entitlements forbidden response has a 2xx status code
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enterprise entitlements report get allowed enterprises for entitlements forbidden response has a 3xx status code
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enterprise entitlements report get allowed enterprises for entitlements forbidden response has a 4xx status code
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this enterprise entitlements report get allowed enterprises for entitlements forbidden response has a 5xx status code
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this enterprise entitlements report get allowed enterprises for entitlements forbidden response a status code equal to that given
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/entitlements/allowedenterprises][%d] enterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsForbidden  %+v", 403, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsForbidden) String() string {
	return fmt.Sprintf("[GET /v1/entitlements/allowedenterprises][%d] enterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsForbidden  %+v", 403, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsNotFound creates a EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsNotFound with default headers values
func NewEnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsNotFound() *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsNotFound {
	return &EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsNotFound{}
}

/*
EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this enterprise entitlements report get allowed enterprises for entitlements not found response has a 2xx status code
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enterprise entitlements report get allowed enterprises for entitlements not found response has a 3xx status code
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enterprise entitlements report get allowed enterprises for entitlements not found response has a 4xx status code
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this enterprise entitlements report get allowed enterprises for entitlements not found response has a 5xx status code
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this enterprise entitlements report get allowed enterprises for entitlements not found response a status code equal to that given
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/entitlements/allowedenterprises][%d] enterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsNotFound  %+v", 404, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsNotFound) String() string {
	return fmt.Sprintf("[GET /v1/entitlements/allowedenterprises][%d] enterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsNotFound  %+v", 404, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsInternalServerError creates a EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsInternalServerError with default headers values
func NewEnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsInternalServerError() *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsInternalServerError {
	return &EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsInternalServerError{}
}

/*
EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this enterprise entitlements report get allowed enterprises for entitlements internal server error response has a 2xx status code
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enterprise entitlements report get allowed enterprises for entitlements internal server error response has a 3xx status code
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enterprise entitlements report get allowed enterprises for entitlements internal server error response has a 4xx status code
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this enterprise entitlements report get allowed enterprises for entitlements internal server error response has a 5xx status code
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this enterprise entitlements report get allowed enterprises for entitlements internal server error response a status code equal to that given
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/entitlements/allowedenterprises][%d] enterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsInternalServerError  %+v", 500, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/entitlements/allowedenterprises][%d] enterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsInternalServerError  %+v", 500, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsGatewayTimeout creates a EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsGatewayTimeout with default headers values
func NewEnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsGatewayTimeout() *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsGatewayTimeout {
	return &EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsGatewayTimeout{}
}

/*
EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this enterprise entitlements report get allowed enterprises for entitlements gateway timeout response has a 2xx status code
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enterprise entitlements report get allowed enterprises for entitlements gateway timeout response has a 3xx status code
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enterprise entitlements report get allowed enterprises for entitlements gateway timeout response has a 4xx status code
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this enterprise entitlements report get allowed enterprises for entitlements gateway timeout response has a 5xx status code
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this enterprise entitlements report get allowed enterprises for entitlements gateway timeout response a status code equal to that given
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/entitlements/allowedenterprises][%d] enterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/entitlements/allowedenterprises][%d] enterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsDefault creates a EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsDefault with default headers values
func NewEnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsDefault(code int) *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsDefault {
	return &EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsDefault{
		_statusCode: code,
	}
}

/*
EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the enterprise entitlements report get allowed enterprises for entitlements default response
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this enterprise entitlements report get allowed enterprises for entitlements default response has a 2xx status code
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this enterprise entitlements report get allowed enterprises for entitlements default response has a 3xx status code
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this enterprise entitlements report get allowed enterprises for entitlements default response has a 4xx status code
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this enterprise entitlements report get allowed enterprises for entitlements default response has a 5xx status code
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this enterprise entitlements report get allowed enterprises for entitlements default response a status code equal to that given
func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/entitlements/allowedenterprises][%d] EnterpriseEntitlementsReport_GetAllowedEnterprisesForEntitlements default  %+v", o._statusCode, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsDefault) String() string {
	return fmt.Sprintf("[GET /v1/entitlements/allowedenterprises][%d] EnterpriseEntitlementsReport_GetAllowedEnterprisesForEntitlements default  %+v", o._statusCode, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *EnterpriseEntitlementsReportGetAllowedEnterprisesForEntitlementsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

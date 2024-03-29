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

// EnterpriseEntitlementsReportGetProjectReportReader is a Reader for the EnterpriseEntitlementsReportGetProjectReport structure.
type EnterpriseEntitlementsReportGetProjectReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnterpriseEntitlementsReportGetProjectReportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEnterpriseEntitlementsReportGetProjectReportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEnterpriseEntitlementsReportGetProjectReportUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEnterpriseEntitlementsReportGetProjectReportForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEnterpriseEntitlementsReportGetProjectReportNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEnterpriseEntitlementsReportGetProjectReportInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEnterpriseEntitlementsReportGetProjectReportGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEnterpriseEntitlementsReportGetProjectReportDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEnterpriseEntitlementsReportGetProjectReportOK creates a EnterpriseEntitlementsReportGetProjectReportOK with default headers values
func NewEnterpriseEntitlementsReportGetProjectReportOK() *EnterpriseEntitlementsReportGetProjectReportOK {
	return &EnterpriseEntitlementsReportGetProjectReportOK{}
}

/*
EnterpriseEntitlementsReportGetProjectReportOK describes a response with status code 200, with default header values.

A successful response.
*/
type EnterpriseEntitlementsReportGetProjectReportOK struct {
	Payload *swagger_models.ProjectReport
}

// IsSuccess returns true when this enterprise entitlements report get project report o k response has a 2xx status code
func (o *EnterpriseEntitlementsReportGetProjectReportOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this enterprise entitlements report get project report o k response has a 3xx status code
func (o *EnterpriseEntitlementsReportGetProjectReportOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enterprise entitlements report get project report o k response has a 4xx status code
func (o *EnterpriseEntitlementsReportGetProjectReportOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this enterprise entitlements report get project report o k response has a 5xx status code
func (o *EnterpriseEntitlementsReportGetProjectReportOK) IsServerError() bool {
	return false
}

// IsCode returns true when this enterprise entitlements report get project report o k response a status code equal to that given
func (o *EnterpriseEntitlementsReportGetProjectReportOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the enterprise entitlements report get project report o k response
func (o *EnterpriseEntitlementsReportGetProjectReportOK) Code() int {
	return 200
}

func (o *EnterpriseEntitlementsReportGetProjectReportOK) Error() string {
	return fmt.Sprintf("[GET /v1/reports/project][%d] enterpriseEntitlementsReportGetProjectReportOK  %+v", 200, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetProjectReportOK) String() string {
	return fmt.Sprintf("[GET /v1/reports/project][%d] enterpriseEntitlementsReportGetProjectReportOK  %+v", 200, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetProjectReportOK) GetPayload() *swagger_models.ProjectReport {
	return o.Payload
}

func (o *EnterpriseEntitlementsReportGetProjectReportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ProjectReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnterpriseEntitlementsReportGetProjectReportUnauthorized creates a EnterpriseEntitlementsReportGetProjectReportUnauthorized with default headers values
func NewEnterpriseEntitlementsReportGetProjectReportUnauthorized() *EnterpriseEntitlementsReportGetProjectReportUnauthorized {
	return &EnterpriseEntitlementsReportGetProjectReportUnauthorized{}
}

/*
EnterpriseEntitlementsReportGetProjectReportUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EnterpriseEntitlementsReportGetProjectReportUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this enterprise entitlements report get project report unauthorized response has a 2xx status code
func (o *EnterpriseEntitlementsReportGetProjectReportUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enterprise entitlements report get project report unauthorized response has a 3xx status code
func (o *EnterpriseEntitlementsReportGetProjectReportUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enterprise entitlements report get project report unauthorized response has a 4xx status code
func (o *EnterpriseEntitlementsReportGetProjectReportUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this enterprise entitlements report get project report unauthorized response has a 5xx status code
func (o *EnterpriseEntitlementsReportGetProjectReportUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this enterprise entitlements report get project report unauthorized response a status code equal to that given
func (o *EnterpriseEntitlementsReportGetProjectReportUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the enterprise entitlements report get project report unauthorized response
func (o *EnterpriseEntitlementsReportGetProjectReportUnauthorized) Code() int {
	return 401
}

func (o *EnterpriseEntitlementsReportGetProjectReportUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/reports/project][%d] enterpriseEntitlementsReportGetProjectReportUnauthorized  %+v", 401, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetProjectReportUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/reports/project][%d] enterpriseEntitlementsReportGetProjectReportUnauthorized  %+v", 401, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetProjectReportUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EnterpriseEntitlementsReportGetProjectReportUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnterpriseEntitlementsReportGetProjectReportForbidden creates a EnterpriseEntitlementsReportGetProjectReportForbidden with default headers values
func NewEnterpriseEntitlementsReportGetProjectReportForbidden() *EnterpriseEntitlementsReportGetProjectReportForbidden {
	return &EnterpriseEntitlementsReportGetProjectReportForbidden{}
}

/*
EnterpriseEntitlementsReportGetProjectReportForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EnterpriseEntitlementsReportGetProjectReportForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this enterprise entitlements report get project report forbidden response has a 2xx status code
func (o *EnterpriseEntitlementsReportGetProjectReportForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enterprise entitlements report get project report forbidden response has a 3xx status code
func (o *EnterpriseEntitlementsReportGetProjectReportForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enterprise entitlements report get project report forbidden response has a 4xx status code
func (o *EnterpriseEntitlementsReportGetProjectReportForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this enterprise entitlements report get project report forbidden response has a 5xx status code
func (o *EnterpriseEntitlementsReportGetProjectReportForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this enterprise entitlements report get project report forbidden response a status code equal to that given
func (o *EnterpriseEntitlementsReportGetProjectReportForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the enterprise entitlements report get project report forbidden response
func (o *EnterpriseEntitlementsReportGetProjectReportForbidden) Code() int {
	return 403
}

func (o *EnterpriseEntitlementsReportGetProjectReportForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/reports/project][%d] enterpriseEntitlementsReportGetProjectReportForbidden  %+v", 403, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetProjectReportForbidden) String() string {
	return fmt.Sprintf("[GET /v1/reports/project][%d] enterpriseEntitlementsReportGetProjectReportForbidden  %+v", 403, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetProjectReportForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EnterpriseEntitlementsReportGetProjectReportForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnterpriseEntitlementsReportGetProjectReportNotFound creates a EnterpriseEntitlementsReportGetProjectReportNotFound with default headers values
func NewEnterpriseEntitlementsReportGetProjectReportNotFound() *EnterpriseEntitlementsReportGetProjectReportNotFound {
	return &EnterpriseEntitlementsReportGetProjectReportNotFound{}
}

/*
EnterpriseEntitlementsReportGetProjectReportNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EnterpriseEntitlementsReportGetProjectReportNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this enterprise entitlements report get project report not found response has a 2xx status code
func (o *EnterpriseEntitlementsReportGetProjectReportNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enterprise entitlements report get project report not found response has a 3xx status code
func (o *EnterpriseEntitlementsReportGetProjectReportNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enterprise entitlements report get project report not found response has a 4xx status code
func (o *EnterpriseEntitlementsReportGetProjectReportNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this enterprise entitlements report get project report not found response has a 5xx status code
func (o *EnterpriseEntitlementsReportGetProjectReportNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this enterprise entitlements report get project report not found response a status code equal to that given
func (o *EnterpriseEntitlementsReportGetProjectReportNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the enterprise entitlements report get project report not found response
func (o *EnterpriseEntitlementsReportGetProjectReportNotFound) Code() int {
	return 404
}

func (o *EnterpriseEntitlementsReportGetProjectReportNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/reports/project][%d] enterpriseEntitlementsReportGetProjectReportNotFound  %+v", 404, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetProjectReportNotFound) String() string {
	return fmt.Sprintf("[GET /v1/reports/project][%d] enterpriseEntitlementsReportGetProjectReportNotFound  %+v", 404, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetProjectReportNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EnterpriseEntitlementsReportGetProjectReportNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnterpriseEntitlementsReportGetProjectReportInternalServerError creates a EnterpriseEntitlementsReportGetProjectReportInternalServerError with default headers values
func NewEnterpriseEntitlementsReportGetProjectReportInternalServerError() *EnterpriseEntitlementsReportGetProjectReportInternalServerError {
	return &EnterpriseEntitlementsReportGetProjectReportInternalServerError{}
}

/*
EnterpriseEntitlementsReportGetProjectReportInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EnterpriseEntitlementsReportGetProjectReportInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this enterprise entitlements report get project report internal server error response has a 2xx status code
func (o *EnterpriseEntitlementsReportGetProjectReportInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enterprise entitlements report get project report internal server error response has a 3xx status code
func (o *EnterpriseEntitlementsReportGetProjectReportInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enterprise entitlements report get project report internal server error response has a 4xx status code
func (o *EnterpriseEntitlementsReportGetProjectReportInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this enterprise entitlements report get project report internal server error response has a 5xx status code
func (o *EnterpriseEntitlementsReportGetProjectReportInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this enterprise entitlements report get project report internal server error response a status code equal to that given
func (o *EnterpriseEntitlementsReportGetProjectReportInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the enterprise entitlements report get project report internal server error response
func (o *EnterpriseEntitlementsReportGetProjectReportInternalServerError) Code() int {
	return 500
}

func (o *EnterpriseEntitlementsReportGetProjectReportInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/reports/project][%d] enterpriseEntitlementsReportGetProjectReportInternalServerError  %+v", 500, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetProjectReportInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/reports/project][%d] enterpriseEntitlementsReportGetProjectReportInternalServerError  %+v", 500, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetProjectReportInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EnterpriseEntitlementsReportGetProjectReportInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnterpriseEntitlementsReportGetProjectReportGatewayTimeout creates a EnterpriseEntitlementsReportGetProjectReportGatewayTimeout with default headers values
func NewEnterpriseEntitlementsReportGetProjectReportGatewayTimeout() *EnterpriseEntitlementsReportGetProjectReportGatewayTimeout {
	return &EnterpriseEntitlementsReportGetProjectReportGatewayTimeout{}
}

/*
EnterpriseEntitlementsReportGetProjectReportGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EnterpriseEntitlementsReportGetProjectReportGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this enterprise entitlements report get project report gateway timeout response has a 2xx status code
func (o *EnterpriseEntitlementsReportGetProjectReportGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enterprise entitlements report get project report gateway timeout response has a 3xx status code
func (o *EnterpriseEntitlementsReportGetProjectReportGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enterprise entitlements report get project report gateway timeout response has a 4xx status code
func (o *EnterpriseEntitlementsReportGetProjectReportGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this enterprise entitlements report get project report gateway timeout response has a 5xx status code
func (o *EnterpriseEntitlementsReportGetProjectReportGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this enterprise entitlements report get project report gateway timeout response a status code equal to that given
func (o *EnterpriseEntitlementsReportGetProjectReportGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the enterprise entitlements report get project report gateway timeout response
func (o *EnterpriseEntitlementsReportGetProjectReportGatewayTimeout) Code() int {
	return 504
}

func (o *EnterpriseEntitlementsReportGetProjectReportGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/reports/project][%d] enterpriseEntitlementsReportGetProjectReportGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetProjectReportGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/reports/project][%d] enterpriseEntitlementsReportGetProjectReportGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetProjectReportGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EnterpriseEntitlementsReportGetProjectReportGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnterpriseEntitlementsReportGetProjectReportDefault creates a EnterpriseEntitlementsReportGetProjectReportDefault with default headers values
func NewEnterpriseEntitlementsReportGetProjectReportDefault(code int) *EnterpriseEntitlementsReportGetProjectReportDefault {
	return &EnterpriseEntitlementsReportGetProjectReportDefault{
		_statusCode: code,
	}
}

/*
EnterpriseEntitlementsReportGetProjectReportDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EnterpriseEntitlementsReportGetProjectReportDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// IsSuccess returns true when this enterprise entitlements report get project report default response has a 2xx status code
func (o *EnterpriseEntitlementsReportGetProjectReportDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this enterprise entitlements report get project report default response has a 3xx status code
func (o *EnterpriseEntitlementsReportGetProjectReportDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this enterprise entitlements report get project report default response has a 4xx status code
func (o *EnterpriseEntitlementsReportGetProjectReportDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this enterprise entitlements report get project report default response has a 5xx status code
func (o *EnterpriseEntitlementsReportGetProjectReportDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this enterprise entitlements report get project report default response a status code equal to that given
func (o *EnterpriseEntitlementsReportGetProjectReportDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the enterprise entitlements report get project report default response
func (o *EnterpriseEntitlementsReportGetProjectReportDefault) Code() int {
	return o._statusCode
}

func (o *EnterpriseEntitlementsReportGetProjectReportDefault) Error() string {
	return fmt.Sprintf("[GET /v1/reports/project][%d] EnterpriseEntitlementsReport_GetProjectReport default  %+v", o._statusCode, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetProjectReportDefault) String() string {
	return fmt.Sprintf("[GET /v1/reports/project][%d] EnterpriseEntitlementsReport_GetProjectReport default  %+v", o._statusCode, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetProjectReportDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *EnterpriseEntitlementsReportGetProjectReportDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

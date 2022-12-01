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

// EnterpriseEntitlementsReportGetAppInstReportReader is a Reader for the EnterpriseEntitlementsReportGetAppInstReport structure.
type EnterpriseEntitlementsReportGetAppInstReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnterpriseEntitlementsReportGetAppInstReportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEnterpriseEntitlementsReportGetAppInstReportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEnterpriseEntitlementsReportGetAppInstReportUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEnterpriseEntitlementsReportGetAppInstReportForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEnterpriseEntitlementsReportGetAppInstReportNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEnterpriseEntitlementsReportGetAppInstReportInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEnterpriseEntitlementsReportGetAppInstReportGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEnterpriseEntitlementsReportGetAppInstReportDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEnterpriseEntitlementsReportGetAppInstReportOK creates a EnterpriseEntitlementsReportGetAppInstReportOK with default headers values
func NewEnterpriseEntitlementsReportGetAppInstReportOK() *EnterpriseEntitlementsReportGetAppInstReportOK {
	return &EnterpriseEntitlementsReportGetAppInstReportOK{}
}

/*
EnterpriseEntitlementsReportGetAppInstReportOK describes a response with status code 200, with default header values.

A successful response.
*/
type EnterpriseEntitlementsReportGetAppInstReportOK struct {
	Payload *swagger_models.AppInstReport
}

// IsSuccess returns true when this enterprise entitlements report get app inst report o k response has a 2xx status code
func (o *EnterpriseEntitlementsReportGetAppInstReportOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this enterprise entitlements report get app inst report o k response has a 3xx status code
func (o *EnterpriseEntitlementsReportGetAppInstReportOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enterprise entitlements report get app inst report o k response has a 4xx status code
func (o *EnterpriseEntitlementsReportGetAppInstReportOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this enterprise entitlements report get app inst report o k response has a 5xx status code
func (o *EnterpriseEntitlementsReportGetAppInstReportOK) IsServerError() bool {
	return false
}

// IsCode returns true when this enterprise entitlements report get app inst report o k response a status code equal to that given
func (o *EnterpriseEntitlementsReportGetAppInstReportOK) IsCode(code int) bool {
	return code == 200
}

func (o *EnterpriseEntitlementsReportGetAppInstReportOK) Error() string {
	return fmt.Sprintf("[GET /v1/reports/apps/instance][%d] enterpriseEntitlementsReportGetAppInstReportOK  %+v", 200, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetAppInstReportOK) String() string {
	return fmt.Sprintf("[GET /v1/reports/apps/instance][%d] enterpriseEntitlementsReportGetAppInstReportOK  %+v", 200, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetAppInstReportOK) GetPayload() *swagger_models.AppInstReport {
	return o.Payload
}

func (o *EnterpriseEntitlementsReportGetAppInstReportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.AppInstReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnterpriseEntitlementsReportGetAppInstReportUnauthorized creates a EnterpriseEntitlementsReportGetAppInstReportUnauthorized with default headers values
func NewEnterpriseEntitlementsReportGetAppInstReportUnauthorized() *EnterpriseEntitlementsReportGetAppInstReportUnauthorized {
	return &EnterpriseEntitlementsReportGetAppInstReportUnauthorized{}
}

/*
EnterpriseEntitlementsReportGetAppInstReportUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EnterpriseEntitlementsReportGetAppInstReportUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this enterprise entitlements report get app inst report unauthorized response has a 2xx status code
func (o *EnterpriseEntitlementsReportGetAppInstReportUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enterprise entitlements report get app inst report unauthorized response has a 3xx status code
func (o *EnterpriseEntitlementsReportGetAppInstReportUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enterprise entitlements report get app inst report unauthorized response has a 4xx status code
func (o *EnterpriseEntitlementsReportGetAppInstReportUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this enterprise entitlements report get app inst report unauthorized response has a 5xx status code
func (o *EnterpriseEntitlementsReportGetAppInstReportUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this enterprise entitlements report get app inst report unauthorized response a status code equal to that given
func (o *EnterpriseEntitlementsReportGetAppInstReportUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *EnterpriseEntitlementsReportGetAppInstReportUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/reports/apps/instance][%d] enterpriseEntitlementsReportGetAppInstReportUnauthorized  %+v", 401, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetAppInstReportUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/reports/apps/instance][%d] enterpriseEntitlementsReportGetAppInstReportUnauthorized  %+v", 401, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetAppInstReportUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EnterpriseEntitlementsReportGetAppInstReportUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnterpriseEntitlementsReportGetAppInstReportForbidden creates a EnterpriseEntitlementsReportGetAppInstReportForbidden with default headers values
func NewEnterpriseEntitlementsReportGetAppInstReportForbidden() *EnterpriseEntitlementsReportGetAppInstReportForbidden {
	return &EnterpriseEntitlementsReportGetAppInstReportForbidden{}
}

/*
EnterpriseEntitlementsReportGetAppInstReportForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EnterpriseEntitlementsReportGetAppInstReportForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this enterprise entitlements report get app inst report forbidden response has a 2xx status code
func (o *EnterpriseEntitlementsReportGetAppInstReportForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enterprise entitlements report get app inst report forbidden response has a 3xx status code
func (o *EnterpriseEntitlementsReportGetAppInstReportForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enterprise entitlements report get app inst report forbidden response has a 4xx status code
func (o *EnterpriseEntitlementsReportGetAppInstReportForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this enterprise entitlements report get app inst report forbidden response has a 5xx status code
func (o *EnterpriseEntitlementsReportGetAppInstReportForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this enterprise entitlements report get app inst report forbidden response a status code equal to that given
func (o *EnterpriseEntitlementsReportGetAppInstReportForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *EnterpriseEntitlementsReportGetAppInstReportForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/reports/apps/instance][%d] enterpriseEntitlementsReportGetAppInstReportForbidden  %+v", 403, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetAppInstReportForbidden) String() string {
	return fmt.Sprintf("[GET /v1/reports/apps/instance][%d] enterpriseEntitlementsReportGetAppInstReportForbidden  %+v", 403, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetAppInstReportForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EnterpriseEntitlementsReportGetAppInstReportForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnterpriseEntitlementsReportGetAppInstReportNotFound creates a EnterpriseEntitlementsReportGetAppInstReportNotFound with default headers values
func NewEnterpriseEntitlementsReportGetAppInstReportNotFound() *EnterpriseEntitlementsReportGetAppInstReportNotFound {
	return &EnterpriseEntitlementsReportGetAppInstReportNotFound{}
}

/*
EnterpriseEntitlementsReportGetAppInstReportNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EnterpriseEntitlementsReportGetAppInstReportNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this enterprise entitlements report get app inst report not found response has a 2xx status code
func (o *EnterpriseEntitlementsReportGetAppInstReportNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enterprise entitlements report get app inst report not found response has a 3xx status code
func (o *EnterpriseEntitlementsReportGetAppInstReportNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enterprise entitlements report get app inst report not found response has a 4xx status code
func (o *EnterpriseEntitlementsReportGetAppInstReportNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this enterprise entitlements report get app inst report not found response has a 5xx status code
func (o *EnterpriseEntitlementsReportGetAppInstReportNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this enterprise entitlements report get app inst report not found response a status code equal to that given
func (o *EnterpriseEntitlementsReportGetAppInstReportNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *EnterpriseEntitlementsReportGetAppInstReportNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/reports/apps/instance][%d] enterpriseEntitlementsReportGetAppInstReportNotFound  %+v", 404, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetAppInstReportNotFound) String() string {
	return fmt.Sprintf("[GET /v1/reports/apps/instance][%d] enterpriseEntitlementsReportGetAppInstReportNotFound  %+v", 404, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetAppInstReportNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EnterpriseEntitlementsReportGetAppInstReportNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnterpriseEntitlementsReportGetAppInstReportInternalServerError creates a EnterpriseEntitlementsReportGetAppInstReportInternalServerError with default headers values
func NewEnterpriseEntitlementsReportGetAppInstReportInternalServerError() *EnterpriseEntitlementsReportGetAppInstReportInternalServerError {
	return &EnterpriseEntitlementsReportGetAppInstReportInternalServerError{}
}

/*
EnterpriseEntitlementsReportGetAppInstReportInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EnterpriseEntitlementsReportGetAppInstReportInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this enterprise entitlements report get app inst report internal server error response has a 2xx status code
func (o *EnterpriseEntitlementsReportGetAppInstReportInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enterprise entitlements report get app inst report internal server error response has a 3xx status code
func (o *EnterpriseEntitlementsReportGetAppInstReportInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enterprise entitlements report get app inst report internal server error response has a 4xx status code
func (o *EnterpriseEntitlementsReportGetAppInstReportInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this enterprise entitlements report get app inst report internal server error response has a 5xx status code
func (o *EnterpriseEntitlementsReportGetAppInstReportInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this enterprise entitlements report get app inst report internal server error response a status code equal to that given
func (o *EnterpriseEntitlementsReportGetAppInstReportInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EnterpriseEntitlementsReportGetAppInstReportInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/reports/apps/instance][%d] enterpriseEntitlementsReportGetAppInstReportInternalServerError  %+v", 500, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetAppInstReportInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/reports/apps/instance][%d] enterpriseEntitlementsReportGetAppInstReportInternalServerError  %+v", 500, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetAppInstReportInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EnterpriseEntitlementsReportGetAppInstReportInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnterpriseEntitlementsReportGetAppInstReportGatewayTimeout creates a EnterpriseEntitlementsReportGetAppInstReportGatewayTimeout with default headers values
func NewEnterpriseEntitlementsReportGetAppInstReportGatewayTimeout() *EnterpriseEntitlementsReportGetAppInstReportGatewayTimeout {
	return &EnterpriseEntitlementsReportGetAppInstReportGatewayTimeout{}
}

/*
EnterpriseEntitlementsReportGetAppInstReportGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EnterpriseEntitlementsReportGetAppInstReportGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this enterprise entitlements report get app inst report gateway timeout response has a 2xx status code
func (o *EnterpriseEntitlementsReportGetAppInstReportGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enterprise entitlements report get app inst report gateway timeout response has a 3xx status code
func (o *EnterpriseEntitlementsReportGetAppInstReportGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enterprise entitlements report get app inst report gateway timeout response has a 4xx status code
func (o *EnterpriseEntitlementsReportGetAppInstReportGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this enterprise entitlements report get app inst report gateway timeout response has a 5xx status code
func (o *EnterpriseEntitlementsReportGetAppInstReportGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this enterprise entitlements report get app inst report gateway timeout response a status code equal to that given
func (o *EnterpriseEntitlementsReportGetAppInstReportGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *EnterpriseEntitlementsReportGetAppInstReportGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/reports/apps/instance][%d] enterpriseEntitlementsReportGetAppInstReportGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetAppInstReportGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/reports/apps/instance][%d] enterpriseEntitlementsReportGetAppInstReportGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetAppInstReportGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EnterpriseEntitlementsReportGetAppInstReportGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnterpriseEntitlementsReportGetAppInstReportDefault creates a EnterpriseEntitlementsReportGetAppInstReportDefault with default headers values
func NewEnterpriseEntitlementsReportGetAppInstReportDefault(code int) *EnterpriseEntitlementsReportGetAppInstReportDefault {
	return &EnterpriseEntitlementsReportGetAppInstReportDefault{
		_statusCode: code,
	}
}

/*
EnterpriseEntitlementsReportGetAppInstReportDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EnterpriseEntitlementsReportGetAppInstReportDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the enterprise entitlements report get app inst report default response
func (o *EnterpriseEntitlementsReportGetAppInstReportDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this enterprise entitlements report get app inst report default response has a 2xx status code
func (o *EnterpriseEntitlementsReportGetAppInstReportDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this enterprise entitlements report get app inst report default response has a 3xx status code
func (o *EnterpriseEntitlementsReportGetAppInstReportDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this enterprise entitlements report get app inst report default response has a 4xx status code
func (o *EnterpriseEntitlementsReportGetAppInstReportDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this enterprise entitlements report get app inst report default response has a 5xx status code
func (o *EnterpriseEntitlementsReportGetAppInstReportDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this enterprise entitlements report get app inst report default response a status code equal to that given
func (o *EnterpriseEntitlementsReportGetAppInstReportDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *EnterpriseEntitlementsReportGetAppInstReportDefault) Error() string {
	return fmt.Sprintf("[GET /v1/reports/apps/instance][%d] EnterpriseEntitlementsReport_GetAppInstReport default  %+v", o._statusCode, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetAppInstReportDefault) String() string {
	return fmt.Sprintf("[GET /v1/reports/apps/instance][%d] EnterpriseEntitlementsReport_GetAppInstReport default  %+v", o._statusCode, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetAppInstReportDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *EnterpriseEntitlementsReportGetAppInstReportDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
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

// EnterpriseEntitlementsReportGetDeviceReportReader is a Reader for the EnterpriseEntitlementsReportGetDeviceReport structure.
type EnterpriseEntitlementsReportGetDeviceReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnterpriseEntitlementsReportGetDeviceReportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEnterpriseEntitlementsReportGetDeviceReportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEnterpriseEntitlementsReportGetDeviceReportUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEnterpriseEntitlementsReportGetDeviceReportForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEnterpriseEntitlementsReportGetDeviceReportNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEnterpriseEntitlementsReportGetDeviceReportInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEnterpriseEntitlementsReportGetDeviceReportGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEnterpriseEntitlementsReportGetDeviceReportDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEnterpriseEntitlementsReportGetDeviceReportOK creates a EnterpriseEntitlementsReportGetDeviceReportOK with default headers values
func NewEnterpriseEntitlementsReportGetDeviceReportOK() *EnterpriseEntitlementsReportGetDeviceReportOK {
	return &EnterpriseEntitlementsReportGetDeviceReportOK{}
}

/*
EnterpriseEntitlementsReportGetDeviceReportOK describes a response with status code 200, with default header values.

A successful response.
*/
type EnterpriseEntitlementsReportGetDeviceReportOK struct {
	Payload *swagger_models.DeviceReport
}

// IsSuccess returns true when this enterprise entitlements report get device report o k response has a 2xx status code
func (o *EnterpriseEntitlementsReportGetDeviceReportOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this enterprise entitlements report get device report o k response has a 3xx status code
func (o *EnterpriseEntitlementsReportGetDeviceReportOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enterprise entitlements report get device report o k response has a 4xx status code
func (o *EnterpriseEntitlementsReportGetDeviceReportOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this enterprise entitlements report get device report o k response has a 5xx status code
func (o *EnterpriseEntitlementsReportGetDeviceReportOK) IsServerError() bool {
	return false
}

// IsCode returns true when this enterprise entitlements report get device report o k response a status code equal to that given
func (o *EnterpriseEntitlementsReportGetDeviceReportOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the enterprise entitlements report get device report o k response
func (o *EnterpriseEntitlementsReportGetDeviceReportOK) Code() int {
	return 200
}

func (o *EnterpriseEntitlementsReportGetDeviceReportOK) Error() string {
	return fmt.Sprintf("[GET /v1/reports/device][%d] enterpriseEntitlementsReportGetDeviceReportOK  %+v", 200, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetDeviceReportOK) String() string {
	return fmt.Sprintf("[GET /v1/reports/device][%d] enterpriseEntitlementsReportGetDeviceReportOK  %+v", 200, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetDeviceReportOK) GetPayload() *swagger_models.DeviceReport {
	return o.Payload
}

func (o *EnterpriseEntitlementsReportGetDeviceReportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.DeviceReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnterpriseEntitlementsReportGetDeviceReportUnauthorized creates a EnterpriseEntitlementsReportGetDeviceReportUnauthorized with default headers values
func NewEnterpriseEntitlementsReportGetDeviceReportUnauthorized() *EnterpriseEntitlementsReportGetDeviceReportUnauthorized {
	return &EnterpriseEntitlementsReportGetDeviceReportUnauthorized{}
}

/*
EnterpriseEntitlementsReportGetDeviceReportUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EnterpriseEntitlementsReportGetDeviceReportUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this enterprise entitlements report get device report unauthorized response has a 2xx status code
func (o *EnterpriseEntitlementsReportGetDeviceReportUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enterprise entitlements report get device report unauthorized response has a 3xx status code
func (o *EnterpriseEntitlementsReportGetDeviceReportUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enterprise entitlements report get device report unauthorized response has a 4xx status code
func (o *EnterpriseEntitlementsReportGetDeviceReportUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this enterprise entitlements report get device report unauthorized response has a 5xx status code
func (o *EnterpriseEntitlementsReportGetDeviceReportUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this enterprise entitlements report get device report unauthorized response a status code equal to that given
func (o *EnterpriseEntitlementsReportGetDeviceReportUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the enterprise entitlements report get device report unauthorized response
func (o *EnterpriseEntitlementsReportGetDeviceReportUnauthorized) Code() int {
	return 401
}

func (o *EnterpriseEntitlementsReportGetDeviceReportUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/reports/device][%d] enterpriseEntitlementsReportGetDeviceReportUnauthorized  %+v", 401, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetDeviceReportUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/reports/device][%d] enterpriseEntitlementsReportGetDeviceReportUnauthorized  %+v", 401, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetDeviceReportUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EnterpriseEntitlementsReportGetDeviceReportUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnterpriseEntitlementsReportGetDeviceReportForbidden creates a EnterpriseEntitlementsReportGetDeviceReportForbidden with default headers values
func NewEnterpriseEntitlementsReportGetDeviceReportForbidden() *EnterpriseEntitlementsReportGetDeviceReportForbidden {
	return &EnterpriseEntitlementsReportGetDeviceReportForbidden{}
}

/*
EnterpriseEntitlementsReportGetDeviceReportForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EnterpriseEntitlementsReportGetDeviceReportForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this enterprise entitlements report get device report forbidden response has a 2xx status code
func (o *EnterpriseEntitlementsReportGetDeviceReportForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enterprise entitlements report get device report forbidden response has a 3xx status code
func (o *EnterpriseEntitlementsReportGetDeviceReportForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enterprise entitlements report get device report forbidden response has a 4xx status code
func (o *EnterpriseEntitlementsReportGetDeviceReportForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this enterprise entitlements report get device report forbidden response has a 5xx status code
func (o *EnterpriseEntitlementsReportGetDeviceReportForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this enterprise entitlements report get device report forbidden response a status code equal to that given
func (o *EnterpriseEntitlementsReportGetDeviceReportForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the enterprise entitlements report get device report forbidden response
func (o *EnterpriseEntitlementsReportGetDeviceReportForbidden) Code() int {
	return 403
}

func (o *EnterpriseEntitlementsReportGetDeviceReportForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/reports/device][%d] enterpriseEntitlementsReportGetDeviceReportForbidden  %+v", 403, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetDeviceReportForbidden) String() string {
	return fmt.Sprintf("[GET /v1/reports/device][%d] enterpriseEntitlementsReportGetDeviceReportForbidden  %+v", 403, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetDeviceReportForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EnterpriseEntitlementsReportGetDeviceReportForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnterpriseEntitlementsReportGetDeviceReportNotFound creates a EnterpriseEntitlementsReportGetDeviceReportNotFound with default headers values
func NewEnterpriseEntitlementsReportGetDeviceReportNotFound() *EnterpriseEntitlementsReportGetDeviceReportNotFound {
	return &EnterpriseEntitlementsReportGetDeviceReportNotFound{}
}

/*
EnterpriseEntitlementsReportGetDeviceReportNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EnterpriseEntitlementsReportGetDeviceReportNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this enterprise entitlements report get device report not found response has a 2xx status code
func (o *EnterpriseEntitlementsReportGetDeviceReportNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enterprise entitlements report get device report not found response has a 3xx status code
func (o *EnterpriseEntitlementsReportGetDeviceReportNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enterprise entitlements report get device report not found response has a 4xx status code
func (o *EnterpriseEntitlementsReportGetDeviceReportNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this enterprise entitlements report get device report not found response has a 5xx status code
func (o *EnterpriseEntitlementsReportGetDeviceReportNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this enterprise entitlements report get device report not found response a status code equal to that given
func (o *EnterpriseEntitlementsReportGetDeviceReportNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the enterprise entitlements report get device report not found response
func (o *EnterpriseEntitlementsReportGetDeviceReportNotFound) Code() int {
	return 404
}

func (o *EnterpriseEntitlementsReportGetDeviceReportNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/reports/device][%d] enterpriseEntitlementsReportGetDeviceReportNotFound  %+v", 404, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetDeviceReportNotFound) String() string {
	return fmt.Sprintf("[GET /v1/reports/device][%d] enterpriseEntitlementsReportGetDeviceReportNotFound  %+v", 404, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetDeviceReportNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EnterpriseEntitlementsReportGetDeviceReportNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnterpriseEntitlementsReportGetDeviceReportInternalServerError creates a EnterpriseEntitlementsReportGetDeviceReportInternalServerError with default headers values
func NewEnterpriseEntitlementsReportGetDeviceReportInternalServerError() *EnterpriseEntitlementsReportGetDeviceReportInternalServerError {
	return &EnterpriseEntitlementsReportGetDeviceReportInternalServerError{}
}

/*
EnterpriseEntitlementsReportGetDeviceReportInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EnterpriseEntitlementsReportGetDeviceReportInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this enterprise entitlements report get device report internal server error response has a 2xx status code
func (o *EnterpriseEntitlementsReportGetDeviceReportInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enterprise entitlements report get device report internal server error response has a 3xx status code
func (o *EnterpriseEntitlementsReportGetDeviceReportInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enterprise entitlements report get device report internal server error response has a 4xx status code
func (o *EnterpriseEntitlementsReportGetDeviceReportInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this enterprise entitlements report get device report internal server error response has a 5xx status code
func (o *EnterpriseEntitlementsReportGetDeviceReportInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this enterprise entitlements report get device report internal server error response a status code equal to that given
func (o *EnterpriseEntitlementsReportGetDeviceReportInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the enterprise entitlements report get device report internal server error response
func (o *EnterpriseEntitlementsReportGetDeviceReportInternalServerError) Code() int {
	return 500
}

func (o *EnterpriseEntitlementsReportGetDeviceReportInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/reports/device][%d] enterpriseEntitlementsReportGetDeviceReportInternalServerError  %+v", 500, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetDeviceReportInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/reports/device][%d] enterpriseEntitlementsReportGetDeviceReportInternalServerError  %+v", 500, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetDeviceReportInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EnterpriseEntitlementsReportGetDeviceReportInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnterpriseEntitlementsReportGetDeviceReportGatewayTimeout creates a EnterpriseEntitlementsReportGetDeviceReportGatewayTimeout with default headers values
func NewEnterpriseEntitlementsReportGetDeviceReportGatewayTimeout() *EnterpriseEntitlementsReportGetDeviceReportGatewayTimeout {
	return &EnterpriseEntitlementsReportGetDeviceReportGatewayTimeout{}
}

/*
EnterpriseEntitlementsReportGetDeviceReportGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EnterpriseEntitlementsReportGetDeviceReportGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this enterprise entitlements report get device report gateway timeout response has a 2xx status code
func (o *EnterpriseEntitlementsReportGetDeviceReportGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enterprise entitlements report get device report gateway timeout response has a 3xx status code
func (o *EnterpriseEntitlementsReportGetDeviceReportGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enterprise entitlements report get device report gateway timeout response has a 4xx status code
func (o *EnterpriseEntitlementsReportGetDeviceReportGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this enterprise entitlements report get device report gateway timeout response has a 5xx status code
func (o *EnterpriseEntitlementsReportGetDeviceReportGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this enterprise entitlements report get device report gateway timeout response a status code equal to that given
func (o *EnterpriseEntitlementsReportGetDeviceReportGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the enterprise entitlements report get device report gateway timeout response
func (o *EnterpriseEntitlementsReportGetDeviceReportGatewayTimeout) Code() int {
	return 504
}

func (o *EnterpriseEntitlementsReportGetDeviceReportGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/reports/device][%d] enterpriseEntitlementsReportGetDeviceReportGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetDeviceReportGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/reports/device][%d] enterpriseEntitlementsReportGetDeviceReportGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetDeviceReportGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EnterpriseEntitlementsReportGetDeviceReportGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnterpriseEntitlementsReportGetDeviceReportDefault creates a EnterpriseEntitlementsReportGetDeviceReportDefault with default headers values
func NewEnterpriseEntitlementsReportGetDeviceReportDefault(code int) *EnterpriseEntitlementsReportGetDeviceReportDefault {
	return &EnterpriseEntitlementsReportGetDeviceReportDefault{
		_statusCode: code,
	}
}

/*
EnterpriseEntitlementsReportGetDeviceReportDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EnterpriseEntitlementsReportGetDeviceReportDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// IsSuccess returns true when this enterprise entitlements report get device report default response has a 2xx status code
func (o *EnterpriseEntitlementsReportGetDeviceReportDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this enterprise entitlements report get device report default response has a 3xx status code
func (o *EnterpriseEntitlementsReportGetDeviceReportDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this enterprise entitlements report get device report default response has a 4xx status code
func (o *EnterpriseEntitlementsReportGetDeviceReportDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this enterprise entitlements report get device report default response has a 5xx status code
func (o *EnterpriseEntitlementsReportGetDeviceReportDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this enterprise entitlements report get device report default response a status code equal to that given
func (o *EnterpriseEntitlementsReportGetDeviceReportDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the enterprise entitlements report get device report default response
func (o *EnterpriseEntitlementsReportGetDeviceReportDefault) Code() int {
	return o._statusCode
}

func (o *EnterpriseEntitlementsReportGetDeviceReportDefault) Error() string {
	return fmt.Sprintf("[GET /v1/reports/device][%d] EnterpriseEntitlementsReport_GetDeviceReport default  %+v", o._statusCode, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetDeviceReportDefault) String() string {
	return fmt.Sprintf("[GET /v1/reports/device][%d] EnterpriseEntitlementsReport_GetDeviceReport default  %+v", o._statusCode, o.Payload)
}

func (o *EnterpriseEntitlementsReportGetDeviceReportDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *EnterpriseEntitlementsReportGetDeviceReportDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

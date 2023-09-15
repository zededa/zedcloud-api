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

// EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataReader is a Reader for the EnterpriseEntitlementsReportCreateEnterpriseEntitlementData structure.
type EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEnterpriseEntitlementsReportCreateEnterpriseEntitlementDataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEnterpriseEntitlementsReportCreateEnterpriseEntitlementDataBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewEnterpriseEntitlementsReportCreateEnterpriseEntitlementDataUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEnterpriseEntitlementsReportCreateEnterpriseEntitlementDataForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewEnterpriseEntitlementsReportCreateEnterpriseEntitlementDataConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEnterpriseEntitlementsReportCreateEnterpriseEntitlementDataInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEnterpriseEntitlementsReportCreateEnterpriseEntitlementDataGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEnterpriseEntitlementsReportCreateEnterpriseEntitlementDataDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEnterpriseEntitlementsReportCreateEnterpriseEntitlementDataOK creates a EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataOK with default headers values
func NewEnterpriseEntitlementsReportCreateEnterpriseEntitlementDataOK() *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataOK {
	return &EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataOK{}
}

/*
EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataOK describes a response with status code 200, with default header values.

A successful response.
*/
type EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataOK struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this enterprise entitlements report create enterprise entitlement data o k response has a 2xx status code
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this enterprise entitlements report create enterprise entitlement data o k response has a 3xx status code
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enterprise entitlements report create enterprise entitlement data o k response has a 4xx status code
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this enterprise entitlements report create enterprise entitlement data o k response has a 5xx status code
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataOK) IsServerError() bool {
	return false
}

// IsCode returns true when this enterprise entitlements report create enterprise entitlement data o k response a status code equal to that given
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the enterprise entitlements report create enterprise entitlement data o k response
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataOK) Code() int {
	return 200
}

func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataOK) Error() string {
	return fmt.Sprintf("[PUT /v1/entitlements][%d] enterpriseEntitlementsReportCreateEnterpriseEntitlementDataOK  %+v", 200, o.Payload)
}

func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataOK) String() string {
	return fmt.Sprintf("[PUT /v1/entitlements][%d] enterpriseEntitlementsReportCreateEnterpriseEntitlementDataOK  %+v", 200, o.Payload)
}

func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataOK) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnterpriseEntitlementsReportCreateEnterpriseEntitlementDataBadRequest creates a EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataBadRequest with default headers values
func NewEnterpriseEntitlementsReportCreateEnterpriseEntitlementDataBadRequest() *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataBadRequest {
	return &EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataBadRequest{}
}

/*
EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this enterprise entitlements report create enterprise entitlement data bad request response has a 2xx status code
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enterprise entitlements report create enterprise entitlement data bad request response has a 3xx status code
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enterprise entitlements report create enterprise entitlement data bad request response has a 4xx status code
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this enterprise entitlements report create enterprise entitlement data bad request response has a 5xx status code
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this enterprise entitlements report create enterprise entitlement data bad request response a status code equal to that given
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the enterprise entitlements report create enterprise entitlement data bad request response
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataBadRequest) Code() int {
	return 400
}

func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/entitlements][%d] enterpriseEntitlementsReportCreateEnterpriseEntitlementDataBadRequest  %+v", 400, o.Payload)
}

func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataBadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/entitlements][%d] enterpriseEntitlementsReportCreateEnterpriseEntitlementDataBadRequest  %+v", 400, o.Payload)
}

func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnterpriseEntitlementsReportCreateEnterpriseEntitlementDataUnauthorized creates a EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataUnauthorized with default headers values
func NewEnterpriseEntitlementsReportCreateEnterpriseEntitlementDataUnauthorized() *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataUnauthorized {
	return &EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataUnauthorized{}
}

/*
EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this enterprise entitlements report create enterprise entitlement data unauthorized response has a 2xx status code
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enterprise entitlements report create enterprise entitlement data unauthorized response has a 3xx status code
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enterprise entitlements report create enterprise entitlement data unauthorized response has a 4xx status code
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this enterprise entitlements report create enterprise entitlement data unauthorized response has a 5xx status code
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this enterprise entitlements report create enterprise entitlement data unauthorized response a status code equal to that given
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the enterprise entitlements report create enterprise entitlement data unauthorized response
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataUnauthorized) Code() int {
	return 401
}

func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/entitlements][%d] enterpriseEntitlementsReportCreateEnterpriseEntitlementDataUnauthorized  %+v", 401, o.Payload)
}

func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/entitlements][%d] enterpriseEntitlementsReportCreateEnterpriseEntitlementDataUnauthorized  %+v", 401, o.Payload)
}

func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnterpriseEntitlementsReportCreateEnterpriseEntitlementDataForbidden creates a EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataForbidden with default headers values
func NewEnterpriseEntitlementsReportCreateEnterpriseEntitlementDataForbidden() *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataForbidden {
	return &EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataForbidden{}
}

/*
EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this enterprise entitlements report create enterprise entitlement data forbidden response has a 2xx status code
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enterprise entitlements report create enterprise entitlement data forbidden response has a 3xx status code
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enterprise entitlements report create enterprise entitlement data forbidden response has a 4xx status code
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this enterprise entitlements report create enterprise entitlement data forbidden response has a 5xx status code
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this enterprise entitlements report create enterprise entitlement data forbidden response a status code equal to that given
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the enterprise entitlements report create enterprise entitlement data forbidden response
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataForbidden) Code() int {
	return 403
}

func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/entitlements][%d] enterpriseEntitlementsReportCreateEnterpriseEntitlementDataForbidden  %+v", 403, o.Payload)
}

func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/entitlements][%d] enterpriseEntitlementsReportCreateEnterpriseEntitlementDataForbidden  %+v", 403, o.Payload)
}

func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnterpriseEntitlementsReportCreateEnterpriseEntitlementDataConflict creates a EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataConflict with default headers values
func NewEnterpriseEntitlementsReportCreateEnterpriseEntitlementDataConflict() *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataConflict {
	return &EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataConflict{}
}

/*
EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this entitlement record will conflict with an already entitlement record.
*/
type EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataConflict struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this enterprise entitlements report create enterprise entitlement data conflict response has a 2xx status code
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enterprise entitlements report create enterprise entitlement data conflict response has a 3xx status code
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enterprise entitlements report create enterprise entitlement data conflict response has a 4xx status code
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this enterprise entitlements report create enterprise entitlement data conflict response has a 5xx status code
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this enterprise entitlements report create enterprise entitlement data conflict response a status code equal to that given
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the enterprise entitlements report create enterprise entitlement data conflict response
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataConflict) Code() int {
	return 409
}

func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/entitlements][%d] enterpriseEntitlementsReportCreateEnterpriseEntitlementDataConflict  %+v", 409, o.Payload)
}

func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataConflict) String() string {
	return fmt.Sprintf("[PUT /v1/entitlements][%d] enterpriseEntitlementsReportCreateEnterpriseEntitlementDataConflict  %+v", 409, o.Payload)
}

func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataConflict) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnterpriseEntitlementsReportCreateEnterpriseEntitlementDataInternalServerError creates a EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataInternalServerError with default headers values
func NewEnterpriseEntitlementsReportCreateEnterpriseEntitlementDataInternalServerError() *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataInternalServerError {
	return &EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataInternalServerError{}
}

/*
EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this enterprise entitlements report create enterprise entitlement data internal server error response has a 2xx status code
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enterprise entitlements report create enterprise entitlement data internal server error response has a 3xx status code
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enterprise entitlements report create enterprise entitlement data internal server error response has a 4xx status code
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this enterprise entitlements report create enterprise entitlement data internal server error response has a 5xx status code
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this enterprise entitlements report create enterprise entitlement data internal server error response a status code equal to that given
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the enterprise entitlements report create enterprise entitlement data internal server error response
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataInternalServerError) Code() int {
	return 500
}

func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/entitlements][%d] enterpriseEntitlementsReportCreateEnterpriseEntitlementDataInternalServerError  %+v", 500, o.Payload)
}

func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataInternalServerError) String() string {
	return fmt.Sprintf("[PUT /v1/entitlements][%d] enterpriseEntitlementsReportCreateEnterpriseEntitlementDataInternalServerError  %+v", 500, o.Payload)
}

func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnterpriseEntitlementsReportCreateEnterpriseEntitlementDataGatewayTimeout creates a EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataGatewayTimeout with default headers values
func NewEnterpriseEntitlementsReportCreateEnterpriseEntitlementDataGatewayTimeout() *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataGatewayTimeout {
	return &EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataGatewayTimeout{}
}

/*
EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this enterprise entitlements report create enterprise entitlement data gateway timeout response has a 2xx status code
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enterprise entitlements report create enterprise entitlement data gateway timeout response has a 3xx status code
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enterprise entitlements report create enterprise entitlement data gateway timeout response has a 4xx status code
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this enterprise entitlements report create enterprise entitlement data gateway timeout response has a 5xx status code
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this enterprise entitlements report create enterprise entitlement data gateway timeout response a status code equal to that given
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the enterprise entitlements report create enterprise entitlement data gateway timeout response
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataGatewayTimeout) Code() int {
	return 504
}

func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/entitlements][%d] enterpriseEntitlementsReportCreateEnterpriseEntitlementDataGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /v1/entitlements][%d] enterpriseEntitlementsReportCreateEnterpriseEntitlementDataGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnterpriseEntitlementsReportCreateEnterpriseEntitlementDataDefault creates a EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataDefault with default headers values
func NewEnterpriseEntitlementsReportCreateEnterpriseEntitlementDataDefault(code int) *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataDefault {
	return &EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataDefault{
		_statusCode: code,
	}
}

/*
EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// IsSuccess returns true when this enterprise entitlements report create enterprise entitlement data default response has a 2xx status code
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this enterprise entitlements report create enterprise entitlement data default response has a 3xx status code
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this enterprise entitlements report create enterprise entitlement data default response has a 4xx status code
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this enterprise entitlements report create enterprise entitlement data default response has a 5xx status code
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this enterprise entitlements report create enterprise entitlement data default response a status code equal to that given
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the enterprise entitlements report create enterprise entitlement data default response
func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataDefault) Code() int {
	return o._statusCode
}

func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/entitlements][%d] EnterpriseEntitlementsReport_CreateEnterpriseEntitlementData default  %+v", o._statusCode, o.Payload)
}

func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataDefault) String() string {
	return fmt.Sprintf("[PUT /v1/entitlements][%d] EnterpriseEntitlementsReport_CreateEnterpriseEntitlementData default  %+v", o._statusCode, o.Payload)
}

func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *EnterpriseEntitlementsReportCreateEnterpriseEntitlementDataDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

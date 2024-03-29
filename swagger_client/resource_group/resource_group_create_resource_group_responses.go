// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package resource_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// ResourceGroupCreateResourceGroupReader is a Reader for the ResourceGroupCreateResourceGroup structure.
type ResourceGroupCreateResourceGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResourceGroupCreateResourceGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResourceGroupCreateResourceGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewResourceGroupCreateResourceGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewResourceGroupCreateResourceGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewResourceGroupCreateResourceGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewResourceGroupCreateResourceGroupConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewResourceGroupCreateResourceGroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewResourceGroupCreateResourceGroupGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewResourceGroupCreateResourceGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewResourceGroupCreateResourceGroupOK creates a ResourceGroupCreateResourceGroupOK with default headers values
func NewResourceGroupCreateResourceGroupOK() *ResourceGroupCreateResourceGroupOK {
	return &ResourceGroupCreateResourceGroupOK{}
}

/*
ResourceGroupCreateResourceGroupOK describes a response with status code 200, with default header values.

A successful response.
*/
type ResourceGroupCreateResourceGroupOK struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this resource group create resource group o k response has a 2xx status code
func (o *ResourceGroupCreateResourceGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this resource group create resource group o k response has a 3xx status code
func (o *ResourceGroupCreateResourceGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource group create resource group o k response has a 4xx status code
func (o *ResourceGroupCreateResourceGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this resource group create resource group o k response has a 5xx status code
func (o *ResourceGroupCreateResourceGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this resource group create resource group o k response a status code equal to that given
func (o *ResourceGroupCreateResourceGroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the resource group create resource group o k response
func (o *ResourceGroupCreateResourceGroupOK) Code() int {
	return 200
}

func (o *ResourceGroupCreateResourceGroupOK) Error() string {
	return fmt.Sprintf("[POST /v1/projects][%d] resourceGroupCreateResourceGroupOK  %+v", 200, o.Payload)
}

func (o *ResourceGroupCreateResourceGroupOK) String() string {
	return fmt.Sprintf("[POST /v1/projects][%d] resourceGroupCreateResourceGroupOK  %+v", 200, o.Payload)
}

func (o *ResourceGroupCreateResourceGroupOK) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ResourceGroupCreateResourceGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResourceGroupCreateResourceGroupBadRequest creates a ResourceGroupCreateResourceGroupBadRequest with default headers values
func NewResourceGroupCreateResourceGroupBadRequest() *ResourceGroupCreateResourceGroupBadRequest {
	return &ResourceGroupCreateResourceGroupBadRequest{}
}

/*
ResourceGroupCreateResourceGroupBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type ResourceGroupCreateResourceGroupBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this resource group create resource group bad request response has a 2xx status code
func (o *ResourceGroupCreateResourceGroupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resource group create resource group bad request response has a 3xx status code
func (o *ResourceGroupCreateResourceGroupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource group create resource group bad request response has a 4xx status code
func (o *ResourceGroupCreateResourceGroupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this resource group create resource group bad request response has a 5xx status code
func (o *ResourceGroupCreateResourceGroupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this resource group create resource group bad request response a status code equal to that given
func (o *ResourceGroupCreateResourceGroupBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the resource group create resource group bad request response
func (o *ResourceGroupCreateResourceGroupBadRequest) Code() int {
	return 400
}

func (o *ResourceGroupCreateResourceGroupBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/projects][%d] resourceGroupCreateResourceGroupBadRequest  %+v", 400, o.Payload)
}

func (o *ResourceGroupCreateResourceGroupBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/projects][%d] resourceGroupCreateResourceGroupBadRequest  %+v", 400, o.Payload)
}

func (o *ResourceGroupCreateResourceGroupBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ResourceGroupCreateResourceGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResourceGroupCreateResourceGroupUnauthorized creates a ResourceGroupCreateResourceGroupUnauthorized with default headers values
func NewResourceGroupCreateResourceGroupUnauthorized() *ResourceGroupCreateResourceGroupUnauthorized {
	return &ResourceGroupCreateResourceGroupUnauthorized{}
}

/*
ResourceGroupCreateResourceGroupUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type ResourceGroupCreateResourceGroupUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this resource group create resource group unauthorized response has a 2xx status code
func (o *ResourceGroupCreateResourceGroupUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resource group create resource group unauthorized response has a 3xx status code
func (o *ResourceGroupCreateResourceGroupUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource group create resource group unauthorized response has a 4xx status code
func (o *ResourceGroupCreateResourceGroupUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this resource group create resource group unauthorized response has a 5xx status code
func (o *ResourceGroupCreateResourceGroupUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this resource group create resource group unauthorized response a status code equal to that given
func (o *ResourceGroupCreateResourceGroupUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the resource group create resource group unauthorized response
func (o *ResourceGroupCreateResourceGroupUnauthorized) Code() int {
	return 401
}

func (o *ResourceGroupCreateResourceGroupUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/projects][%d] resourceGroupCreateResourceGroupUnauthorized  %+v", 401, o.Payload)
}

func (o *ResourceGroupCreateResourceGroupUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/projects][%d] resourceGroupCreateResourceGroupUnauthorized  %+v", 401, o.Payload)
}

func (o *ResourceGroupCreateResourceGroupUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ResourceGroupCreateResourceGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResourceGroupCreateResourceGroupForbidden creates a ResourceGroupCreateResourceGroupForbidden with default headers values
func NewResourceGroupCreateResourceGroupForbidden() *ResourceGroupCreateResourceGroupForbidden {
	return &ResourceGroupCreateResourceGroupForbidden{}
}

/*
ResourceGroupCreateResourceGroupForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type ResourceGroupCreateResourceGroupForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this resource group create resource group forbidden response has a 2xx status code
func (o *ResourceGroupCreateResourceGroupForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resource group create resource group forbidden response has a 3xx status code
func (o *ResourceGroupCreateResourceGroupForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource group create resource group forbidden response has a 4xx status code
func (o *ResourceGroupCreateResourceGroupForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this resource group create resource group forbidden response has a 5xx status code
func (o *ResourceGroupCreateResourceGroupForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this resource group create resource group forbidden response a status code equal to that given
func (o *ResourceGroupCreateResourceGroupForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the resource group create resource group forbidden response
func (o *ResourceGroupCreateResourceGroupForbidden) Code() int {
	return 403
}

func (o *ResourceGroupCreateResourceGroupForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/projects][%d] resourceGroupCreateResourceGroupForbidden  %+v", 403, o.Payload)
}

func (o *ResourceGroupCreateResourceGroupForbidden) String() string {
	return fmt.Sprintf("[POST /v1/projects][%d] resourceGroupCreateResourceGroupForbidden  %+v", 403, o.Payload)
}

func (o *ResourceGroupCreateResourceGroupForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ResourceGroupCreateResourceGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResourceGroupCreateResourceGroupConflict creates a ResourceGroupCreateResourceGroupConflict with default headers values
func NewResourceGroupCreateResourceGroupConflict() *ResourceGroupCreateResourceGroupConflict {
	return &ResourceGroupCreateResourceGroupConflict{}
}

/*
ResourceGroupCreateResourceGroupConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this resource group record will conflict with an already existing resource group record.
*/
type ResourceGroupCreateResourceGroupConflict struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this resource group create resource group conflict response has a 2xx status code
func (o *ResourceGroupCreateResourceGroupConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resource group create resource group conflict response has a 3xx status code
func (o *ResourceGroupCreateResourceGroupConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource group create resource group conflict response has a 4xx status code
func (o *ResourceGroupCreateResourceGroupConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this resource group create resource group conflict response has a 5xx status code
func (o *ResourceGroupCreateResourceGroupConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this resource group create resource group conflict response a status code equal to that given
func (o *ResourceGroupCreateResourceGroupConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the resource group create resource group conflict response
func (o *ResourceGroupCreateResourceGroupConflict) Code() int {
	return 409
}

func (o *ResourceGroupCreateResourceGroupConflict) Error() string {
	return fmt.Sprintf("[POST /v1/projects][%d] resourceGroupCreateResourceGroupConflict  %+v", 409, o.Payload)
}

func (o *ResourceGroupCreateResourceGroupConflict) String() string {
	return fmt.Sprintf("[POST /v1/projects][%d] resourceGroupCreateResourceGroupConflict  %+v", 409, o.Payload)
}

func (o *ResourceGroupCreateResourceGroupConflict) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ResourceGroupCreateResourceGroupConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResourceGroupCreateResourceGroupInternalServerError creates a ResourceGroupCreateResourceGroupInternalServerError with default headers values
func NewResourceGroupCreateResourceGroupInternalServerError() *ResourceGroupCreateResourceGroupInternalServerError {
	return &ResourceGroupCreateResourceGroupInternalServerError{}
}

/*
ResourceGroupCreateResourceGroupInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type ResourceGroupCreateResourceGroupInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this resource group create resource group internal server error response has a 2xx status code
func (o *ResourceGroupCreateResourceGroupInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resource group create resource group internal server error response has a 3xx status code
func (o *ResourceGroupCreateResourceGroupInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource group create resource group internal server error response has a 4xx status code
func (o *ResourceGroupCreateResourceGroupInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this resource group create resource group internal server error response has a 5xx status code
func (o *ResourceGroupCreateResourceGroupInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this resource group create resource group internal server error response a status code equal to that given
func (o *ResourceGroupCreateResourceGroupInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the resource group create resource group internal server error response
func (o *ResourceGroupCreateResourceGroupInternalServerError) Code() int {
	return 500
}

func (o *ResourceGroupCreateResourceGroupInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/projects][%d] resourceGroupCreateResourceGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *ResourceGroupCreateResourceGroupInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/projects][%d] resourceGroupCreateResourceGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *ResourceGroupCreateResourceGroupInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ResourceGroupCreateResourceGroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResourceGroupCreateResourceGroupGatewayTimeout creates a ResourceGroupCreateResourceGroupGatewayTimeout with default headers values
func NewResourceGroupCreateResourceGroupGatewayTimeout() *ResourceGroupCreateResourceGroupGatewayTimeout {
	return &ResourceGroupCreateResourceGroupGatewayTimeout{}
}

/*
ResourceGroupCreateResourceGroupGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type ResourceGroupCreateResourceGroupGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this resource group create resource group gateway timeout response has a 2xx status code
func (o *ResourceGroupCreateResourceGroupGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resource group create resource group gateway timeout response has a 3xx status code
func (o *ResourceGroupCreateResourceGroupGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource group create resource group gateway timeout response has a 4xx status code
func (o *ResourceGroupCreateResourceGroupGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this resource group create resource group gateway timeout response has a 5xx status code
func (o *ResourceGroupCreateResourceGroupGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this resource group create resource group gateway timeout response a status code equal to that given
func (o *ResourceGroupCreateResourceGroupGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the resource group create resource group gateway timeout response
func (o *ResourceGroupCreateResourceGroupGatewayTimeout) Code() int {
	return 504
}

func (o *ResourceGroupCreateResourceGroupGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v1/projects][%d] resourceGroupCreateResourceGroupGatewayTimeout  %+v", 504, o.Payload)
}

func (o *ResourceGroupCreateResourceGroupGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /v1/projects][%d] resourceGroupCreateResourceGroupGatewayTimeout  %+v", 504, o.Payload)
}

func (o *ResourceGroupCreateResourceGroupGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ResourceGroupCreateResourceGroupGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResourceGroupCreateResourceGroupDefault creates a ResourceGroupCreateResourceGroupDefault with default headers values
func NewResourceGroupCreateResourceGroupDefault(code int) *ResourceGroupCreateResourceGroupDefault {
	return &ResourceGroupCreateResourceGroupDefault{
		_statusCode: code,
	}
}

/*
ResourceGroupCreateResourceGroupDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ResourceGroupCreateResourceGroupDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// IsSuccess returns true when this resource group create resource group default response has a 2xx status code
func (o *ResourceGroupCreateResourceGroupDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this resource group create resource group default response has a 3xx status code
func (o *ResourceGroupCreateResourceGroupDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this resource group create resource group default response has a 4xx status code
func (o *ResourceGroupCreateResourceGroupDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this resource group create resource group default response has a 5xx status code
func (o *ResourceGroupCreateResourceGroupDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this resource group create resource group default response a status code equal to that given
func (o *ResourceGroupCreateResourceGroupDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the resource group create resource group default response
func (o *ResourceGroupCreateResourceGroupDefault) Code() int {
	return o._statusCode
}

func (o *ResourceGroupCreateResourceGroupDefault) Error() string {
	return fmt.Sprintf("[POST /v1/projects][%d] ResourceGroup_CreateResourceGroup default  %+v", o._statusCode, o.Payload)
}

func (o *ResourceGroupCreateResourceGroupDefault) String() string {
	return fmt.Sprintf("[POST /v1/projects][%d] ResourceGroup_CreateResourceGroup default  %+v", o._statusCode, o.Payload)
}

func (o *ResourceGroupCreateResourceGroupDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *ResourceGroupCreateResourceGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

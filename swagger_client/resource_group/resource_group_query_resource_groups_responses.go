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

// ResourceGroupQueryResourceGroupsReader is a Reader for the ResourceGroupQueryResourceGroups structure.
type ResourceGroupQueryResourceGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResourceGroupQueryResourceGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResourceGroupQueryResourceGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewResourceGroupQueryResourceGroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewResourceGroupQueryResourceGroupsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewResourceGroupQueryResourceGroupsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewResourceGroupQueryResourceGroupsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewResourceGroupQueryResourceGroupsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewResourceGroupQueryResourceGroupsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewResourceGroupQueryResourceGroupsOK creates a ResourceGroupQueryResourceGroupsOK with default headers values
func NewResourceGroupQueryResourceGroupsOK() *ResourceGroupQueryResourceGroupsOK {
	return &ResourceGroupQueryResourceGroupsOK{}
}

/*
ResourceGroupQueryResourceGroupsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ResourceGroupQueryResourceGroupsOK struct {
	Payload *swagger_models.Tags
}

// IsSuccess returns true when this resource group query resource groups o k response has a 2xx status code
func (o *ResourceGroupQueryResourceGroupsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this resource group query resource groups o k response has a 3xx status code
func (o *ResourceGroupQueryResourceGroupsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource group query resource groups o k response has a 4xx status code
func (o *ResourceGroupQueryResourceGroupsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this resource group query resource groups o k response has a 5xx status code
func (o *ResourceGroupQueryResourceGroupsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this resource group query resource groups o k response a status code equal to that given
func (o *ResourceGroupQueryResourceGroupsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ResourceGroupQueryResourceGroupsOK) Error() string {
	return fmt.Sprintf("[GET /v1/projects][%d] resourceGroupQueryResourceGroupsOK  %+v", 200, o.Payload)
}

func (o *ResourceGroupQueryResourceGroupsOK) String() string {
	return fmt.Sprintf("[GET /v1/projects][%d] resourceGroupQueryResourceGroupsOK  %+v", 200, o.Payload)
}

func (o *ResourceGroupQueryResourceGroupsOK) GetPayload() *swagger_models.Tags {
	return o.Payload
}

func (o *ResourceGroupQueryResourceGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.Tags)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResourceGroupQueryResourceGroupsBadRequest creates a ResourceGroupQueryResourceGroupsBadRequest with default headers values
func NewResourceGroupQueryResourceGroupsBadRequest() *ResourceGroupQueryResourceGroupsBadRequest {
	return &ResourceGroupQueryResourceGroupsBadRequest{}
}

/*
ResourceGroupQueryResourceGroupsBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of invalid value of filter parameters.
*/
type ResourceGroupQueryResourceGroupsBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this resource group query resource groups bad request response has a 2xx status code
func (o *ResourceGroupQueryResourceGroupsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resource group query resource groups bad request response has a 3xx status code
func (o *ResourceGroupQueryResourceGroupsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource group query resource groups bad request response has a 4xx status code
func (o *ResourceGroupQueryResourceGroupsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this resource group query resource groups bad request response has a 5xx status code
func (o *ResourceGroupQueryResourceGroupsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this resource group query resource groups bad request response a status code equal to that given
func (o *ResourceGroupQueryResourceGroupsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ResourceGroupQueryResourceGroupsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/projects][%d] resourceGroupQueryResourceGroupsBadRequest  %+v", 400, o.Payload)
}

func (o *ResourceGroupQueryResourceGroupsBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/projects][%d] resourceGroupQueryResourceGroupsBadRequest  %+v", 400, o.Payload)
}

func (o *ResourceGroupQueryResourceGroupsBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ResourceGroupQueryResourceGroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResourceGroupQueryResourceGroupsUnauthorized creates a ResourceGroupQueryResourceGroupsUnauthorized with default headers values
func NewResourceGroupQueryResourceGroupsUnauthorized() *ResourceGroupQueryResourceGroupsUnauthorized {
	return &ResourceGroupQueryResourceGroupsUnauthorized{}
}

/*
ResourceGroupQueryResourceGroupsUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type ResourceGroupQueryResourceGroupsUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this resource group query resource groups unauthorized response has a 2xx status code
func (o *ResourceGroupQueryResourceGroupsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resource group query resource groups unauthorized response has a 3xx status code
func (o *ResourceGroupQueryResourceGroupsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource group query resource groups unauthorized response has a 4xx status code
func (o *ResourceGroupQueryResourceGroupsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this resource group query resource groups unauthorized response has a 5xx status code
func (o *ResourceGroupQueryResourceGroupsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this resource group query resource groups unauthorized response a status code equal to that given
func (o *ResourceGroupQueryResourceGroupsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ResourceGroupQueryResourceGroupsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/projects][%d] resourceGroupQueryResourceGroupsUnauthorized  %+v", 401, o.Payload)
}

func (o *ResourceGroupQueryResourceGroupsUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/projects][%d] resourceGroupQueryResourceGroupsUnauthorized  %+v", 401, o.Payload)
}

func (o *ResourceGroupQueryResourceGroupsUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ResourceGroupQueryResourceGroupsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResourceGroupQueryResourceGroupsForbidden creates a ResourceGroupQueryResourceGroupsForbidden with default headers values
func NewResourceGroupQueryResourceGroupsForbidden() *ResourceGroupQueryResourceGroupsForbidden {
	return &ResourceGroupQueryResourceGroupsForbidden{}
}

/*
ResourceGroupQueryResourceGroupsForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type ResourceGroupQueryResourceGroupsForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this resource group query resource groups forbidden response has a 2xx status code
func (o *ResourceGroupQueryResourceGroupsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resource group query resource groups forbidden response has a 3xx status code
func (o *ResourceGroupQueryResourceGroupsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource group query resource groups forbidden response has a 4xx status code
func (o *ResourceGroupQueryResourceGroupsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this resource group query resource groups forbidden response has a 5xx status code
func (o *ResourceGroupQueryResourceGroupsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this resource group query resource groups forbidden response a status code equal to that given
func (o *ResourceGroupQueryResourceGroupsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ResourceGroupQueryResourceGroupsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/projects][%d] resourceGroupQueryResourceGroupsForbidden  %+v", 403, o.Payload)
}

func (o *ResourceGroupQueryResourceGroupsForbidden) String() string {
	return fmt.Sprintf("[GET /v1/projects][%d] resourceGroupQueryResourceGroupsForbidden  %+v", 403, o.Payload)
}

func (o *ResourceGroupQueryResourceGroupsForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ResourceGroupQueryResourceGroupsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResourceGroupQueryResourceGroupsInternalServerError creates a ResourceGroupQueryResourceGroupsInternalServerError with default headers values
func NewResourceGroupQueryResourceGroupsInternalServerError() *ResourceGroupQueryResourceGroupsInternalServerError {
	return &ResourceGroupQueryResourceGroupsInternalServerError{}
}

/*
ResourceGroupQueryResourceGroupsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type ResourceGroupQueryResourceGroupsInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this resource group query resource groups internal server error response has a 2xx status code
func (o *ResourceGroupQueryResourceGroupsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resource group query resource groups internal server error response has a 3xx status code
func (o *ResourceGroupQueryResourceGroupsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource group query resource groups internal server error response has a 4xx status code
func (o *ResourceGroupQueryResourceGroupsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this resource group query resource groups internal server error response has a 5xx status code
func (o *ResourceGroupQueryResourceGroupsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this resource group query resource groups internal server error response a status code equal to that given
func (o *ResourceGroupQueryResourceGroupsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ResourceGroupQueryResourceGroupsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/projects][%d] resourceGroupQueryResourceGroupsInternalServerError  %+v", 500, o.Payload)
}

func (o *ResourceGroupQueryResourceGroupsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/projects][%d] resourceGroupQueryResourceGroupsInternalServerError  %+v", 500, o.Payload)
}

func (o *ResourceGroupQueryResourceGroupsInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ResourceGroupQueryResourceGroupsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResourceGroupQueryResourceGroupsGatewayTimeout creates a ResourceGroupQueryResourceGroupsGatewayTimeout with default headers values
func NewResourceGroupQueryResourceGroupsGatewayTimeout() *ResourceGroupQueryResourceGroupsGatewayTimeout {
	return &ResourceGroupQueryResourceGroupsGatewayTimeout{}
}

/*
ResourceGroupQueryResourceGroupsGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type ResourceGroupQueryResourceGroupsGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this resource group query resource groups gateway timeout response has a 2xx status code
func (o *ResourceGroupQueryResourceGroupsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resource group query resource groups gateway timeout response has a 3xx status code
func (o *ResourceGroupQueryResourceGroupsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource group query resource groups gateway timeout response has a 4xx status code
func (o *ResourceGroupQueryResourceGroupsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this resource group query resource groups gateway timeout response has a 5xx status code
func (o *ResourceGroupQueryResourceGroupsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this resource group query resource groups gateway timeout response a status code equal to that given
func (o *ResourceGroupQueryResourceGroupsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *ResourceGroupQueryResourceGroupsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/projects][%d] resourceGroupQueryResourceGroupsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *ResourceGroupQueryResourceGroupsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/projects][%d] resourceGroupQueryResourceGroupsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *ResourceGroupQueryResourceGroupsGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ResourceGroupQueryResourceGroupsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResourceGroupQueryResourceGroupsDefault creates a ResourceGroupQueryResourceGroupsDefault with default headers values
func NewResourceGroupQueryResourceGroupsDefault(code int) *ResourceGroupQueryResourceGroupsDefault {
	return &ResourceGroupQueryResourceGroupsDefault{
		_statusCode: code,
	}
}

/*
ResourceGroupQueryResourceGroupsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ResourceGroupQueryResourceGroupsDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the resource group query resource groups default response
func (o *ResourceGroupQueryResourceGroupsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this resource group query resource groups default response has a 2xx status code
func (o *ResourceGroupQueryResourceGroupsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this resource group query resource groups default response has a 3xx status code
func (o *ResourceGroupQueryResourceGroupsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this resource group query resource groups default response has a 4xx status code
func (o *ResourceGroupQueryResourceGroupsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this resource group query resource groups default response has a 5xx status code
func (o *ResourceGroupQueryResourceGroupsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this resource group query resource groups default response a status code equal to that given
func (o *ResourceGroupQueryResourceGroupsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ResourceGroupQueryResourceGroupsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/projects][%d] ResourceGroup_QueryResourceGroups default  %+v", o._statusCode, o.Payload)
}

func (o *ResourceGroupQueryResourceGroupsDefault) String() string {
	return fmt.Sprintf("[GET /v1/projects][%d] ResourceGroup_QueryResourceGroups default  %+v", o._statusCode, o.Payload)
}

func (o *ResourceGroupQueryResourceGroupsDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *ResourceGroupQueryResourceGroupsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

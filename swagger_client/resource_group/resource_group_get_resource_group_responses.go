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

// ResourceGroupGetResourceGroupReader is a Reader for the ResourceGroupGetResourceGroup structure.
type ResourceGroupGetResourceGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResourceGroupGetResourceGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResourceGroupGetResourceGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewResourceGroupGetResourceGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewResourceGroupGetResourceGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewResourceGroupGetResourceGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewResourceGroupGetResourceGroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewResourceGroupGetResourceGroupGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewResourceGroupGetResourceGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewResourceGroupGetResourceGroupOK creates a ResourceGroupGetResourceGroupOK with default headers values
func NewResourceGroupGetResourceGroupOK() *ResourceGroupGetResourceGroupOK {
	return &ResourceGroupGetResourceGroupOK{}
}

/*
ResourceGroupGetResourceGroupOK describes a response with status code 200, with default header values.

A successful response.
*/
type ResourceGroupGetResourceGroupOK struct {
	Payload *swagger_models.Tag
}

// IsSuccess returns true when this resource group get resource group o k response has a 2xx status code
func (o *ResourceGroupGetResourceGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this resource group get resource group o k response has a 3xx status code
func (o *ResourceGroupGetResourceGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource group get resource group o k response has a 4xx status code
func (o *ResourceGroupGetResourceGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this resource group get resource group o k response has a 5xx status code
func (o *ResourceGroupGetResourceGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this resource group get resource group o k response a status code equal to that given
func (o *ResourceGroupGetResourceGroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the resource group get resource group o k response
func (o *ResourceGroupGetResourceGroupOK) Code() int {
	return 200
}

func (o *ResourceGroupGetResourceGroupOK) Error() string {
	return fmt.Sprintf("[GET /v1/projects/id/{id}][%d] resourceGroupGetResourceGroupOK  %+v", 200, o.Payload)
}

func (o *ResourceGroupGetResourceGroupOK) String() string {
	return fmt.Sprintf("[GET /v1/projects/id/{id}][%d] resourceGroupGetResourceGroupOK  %+v", 200, o.Payload)
}

func (o *ResourceGroupGetResourceGroupOK) GetPayload() *swagger_models.Tag {
	return o.Payload
}

func (o *ResourceGroupGetResourceGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.Tag)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResourceGroupGetResourceGroupUnauthorized creates a ResourceGroupGetResourceGroupUnauthorized with default headers values
func NewResourceGroupGetResourceGroupUnauthorized() *ResourceGroupGetResourceGroupUnauthorized {
	return &ResourceGroupGetResourceGroupUnauthorized{}
}

/*
ResourceGroupGetResourceGroupUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type ResourceGroupGetResourceGroupUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this resource group get resource group unauthorized response has a 2xx status code
func (o *ResourceGroupGetResourceGroupUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resource group get resource group unauthorized response has a 3xx status code
func (o *ResourceGroupGetResourceGroupUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource group get resource group unauthorized response has a 4xx status code
func (o *ResourceGroupGetResourceGroupUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this resource group get resource group unauthorized response has a 5xx status code
func (o *ResourceGroupGetResourceGroupUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this resource group get resource group unauthorized response a status code equal to that given
func (o *ResourceGroupGetResourceGroupUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the resource group get resource group unauthorized response
func (o *ResourceGroupGetResourceGroupUnauthorized) Code() int {
	return 401
}

func (o *ResourceGroupGetResourceGroupUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/projects/id/{id}][%d] resourceGroupGetResourceGroupUnauthorized  %+v", 401, o.Payload)
}

func (o *ResourceGroupGetResourceGroupUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/projects/id/{id}][%d] resourceGroupGetResourceGroupUnauthorized  %+v", 401, o.Payload)
}

func (o *ResourceGroupGetResourceGroupUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ResourceGroupGetResourceGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResourceGroupGetResourceGroupForbidden creates a ResourceGroupGetResourceGroupForbidden with default headers values
func NewResourceGroupGetResourceGroupForbidden() *ResourceGroupGetResourceGroupForbidden {
	return &ResourceGroupGetResourceGroupForbidden{}
}

/*
ResourceGroupGetResourceGroupForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type ResourceGroupGetResourceGroupForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this resource group get resource group forbidden response has a 2xx status code
func (o *ResourceGroupGetResourceGroupForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resource group get resource group forbidden response has a 3xx status code
func (o *ResourceGroupGetResourceGroupForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource group get resource group forbidden response has a 4xx status code
func (o *ResourceGroupGetResourceGroupForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this resource group get resource group forbidden response has a 5xx status code
func (o *ResourceGroupGetResourceGroupForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this resource group get resource group forbidden response a status code equal to that given
func (o *ResourceGroupGetResourceGroupForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the resource group get resource group forbidden response
func (o *ResourceGroupGetResourceGroupForbidden) Code() int {
	return 403
}

func (o *ResourceGroupGetResourceGroupForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/projects/id/{id}][%d] resourceGroupGetResourceGroupForbidden  %+v", 403, o.Payload)
}

func (o *ResourceGroupGetResourceGroupForbidden) String() string {
	return fmt.Sprintf("[GET /v1/projects/id/{id}][%d] resourceGroupGetResourceGroupForbidden  %+v", 403, o.Payload)
}

func (o *ResourceGroupGetResourceGroupForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ResourceGroupGetResourceGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResourceGroupGetResourceGroupNotFound creates a ResourceGroupGetResourceGroupNotFound with default headers values
func NewResourceGroupGetResourceGroupNotFound() *ResourceGroupGetResourceGroupNotFound {
	return &ResourceGroupGetResourceGroupNotFound{}
}

/*
ResourceGroupGetResourceGroupNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type ResourceGroupGetResourceGroupNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this resource group get resource group not found response has a 2xx status code
func (o *ResourceGroupGetResourceGroupNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resource group get resource group not found response has a 3xx status code
func (o *ResourceGroupGetResourceGroupNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource group get resource group not found response has a 4xx status code
func (o *ResourceGroupGetResourceGroupNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this resource group get resource group not found response has a 5xx status code
func (o *ResourceGroupGetResourceGroupNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this resource group get resource group not found response a status code equal to that given
func (o *ResourceGroupGetResourceGroupNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the resource group get resource group not found response
func (o *ResourceGroupGetResourceGroupNotFound) Code() int {
	return 404
}

func (o *ResourceGroupGetResourceGroupNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/projects/id/{id}][%d] resourceGroupGetResourceGroupNotFound  %+v", 404, o.Payload)
}

func (o *ResourceGroupGetResourceGroupNotFound) String() string {
	return fmt.Sprintf("[GET /v1/projects/id/{id}][%d] resourceGroupGetResourceGroupNotFound  %+v", 404, o.Payload)
}

func (o *ResourceGroupGetResourceGroupNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ResourceGroupGetResourceGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResourceGroupGetResourceGroupInternalServerError creates a ResourceGroupGetResourceGroupInternalServerError with default headers values
func NewResourceGroupGetResourceGroupInternalServerError() *ResourceGroupGetResourceGroupInternalServerError {
	return &ResourceGroupGetResourceGroupInternalServerError{}
}

/*
ResourceGroupGetResourceGroupInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type ResourceGroupGetResourceGroupInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this resource group get resource group internal server error response has a 2xx status code
func (o *ResourceGroupGetResourceGroupInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resource group get resource group internal server error response has a 3xx status code
func (o *ResourceGroupGetResourceGroupInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource group get resource group internal server error response has a 4xx status code
func (o *ResourceGroupGetResourceGroupInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this resource group get resource group internal server error response has a 5xx status code
func (o *ResourceGroupGetResourceGroupInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this resource group get resource group internal server error response a status code equal to that given
func (o *ResourceGroupGetResourceGroupInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the resource group get resource group internal server error response
func (o *ResourceGroupGetResourceGroupInternalServerError) Code() int {
	return 500
}

func (o *ResourceGroupGetResourceGroupInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/projects/id/{id}][%d] resourceGroupGetResourceGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *ResourceGroupGetResourceGroupInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/projects/id/{id}][%d] resourceGroupGetResourceGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *ResourceGroupGetResourceGroupInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ResourceGroupGetResourceGroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResourceGroupGetResourceGroupGatewayTimeout creates a ResourceGroupGetResourceGroupGatewayTimeout with default headers values
func NewResourceGroupGetResourceGroupGatewayTimeout() *ResourceGroupGetResourceGroupGatewayTimeout {
	return &ResourceGroupGetResourceGroupGatewayTimeout{}
}

/*
ResourceGroupGetResourceGroupGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type ResourceGroupGetResourceGroupGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this resource group get resource group gateway timeout response has a 2xx status code
func (o *ResourceGroupGetResourceGroupGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resource group get resource group gateway timeout response has a 3xx status code
func (o *ResourceGroupGetResourceGroupGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource group get resource group gateway timeout response has a 4xx status code
func (o *ResourceGroupGetResourceGroupGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this resource group get resource group gateway timeout response has a 5xx status code
func (o *ResourceGroupGetResourceGroupGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this resource group get resource group gateway timeout response a status code equal to that given
func (o *ResourceGroupGetResourceGroupGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the resource group get resource group gateway timeout response
func (o *ResourceGroupGetResourceGroupGatewayTimeout) Code() int {
	return 504
}

func (o *ResourceGroupGetResourceGroupGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/projects/id/{id}][%d] resourceGroupGetResourceGroupGatewayTimeout  %+v", 504, o.Payload)
}

func (o *ResourceGroupGetResourceGroupGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/projects/id/{id}][%d] resourceGroupGetResourceGroupGatewayTimeout  %+v", 504, o.Payload)
}

func (o *ResourceGroupGetResourceGroupGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ResourceGroupGetResourceGroupGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResourceGroupGetResourceGroupDefault creates a ResourceGroupGetResourceGroupDefault with default headers values
func NewResourceGroupGetResourceGroupDefault(code int) *ResourceGroupGetResourceGroupDefault {
	return &ResourceGroupGetResourceGroupDefault{
		_statusCode: code,
	}
}

/*
ResourceGroupGetResourceGroupDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ResourceGroupGetResourceGroupDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// IsSuccess returns true when this resource group get resource group default response has a 2xx status code
func (o *ResourceGroupGetResourceGroupDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this resource group get resource group default response has a 3xx status code
func (o *ResourceGroupGetResourceGroupDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this resource group get resource group default response has a 4xx status code
func (o *ResourceGroupGetResourceGroupDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this resource group get resource group default response has a 5xx status code
func (o *ResourceGroupGetResourceGroupDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this resource group get resource group default response a status code equal to that given
func (o *ResourceGroupGetResourceGroupDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the resource group get resource group default response
func (o *ResourceGroupGetResourceGroupDefault) Code() int {
	return o._statusCode
}

func (o *ResourceGroupGetResourceGroupDefault) Error() string {
	return fmt.Sprintf("[GET /v1/projects/id/{id}][%d] ResourceGroup_GetResourceGroup default  %+v", o._statusCode, o.Payload)
}

func (o *ResourceGroupGetResourceGroupDefault) String() string {
	return fmt.Sprintf("[GET /v1/projects/id/{id}][%d] ResourceGroup_GetResourceGroup default  %+v", o._statusCode, o.Payload)
}

func (o *ResourceGroupGetResourceGroupDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *ResourceGroupGetResourceGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

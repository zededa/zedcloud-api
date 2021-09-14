// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package resource_group_status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// GetResourceGroupStatusByNameReader is a Reader for the GetResourceGroupStatusByName structure.
type GetResourceGroupStatusByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResourceGroupStatusByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetResourceGroupStatusByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetResourceGroupStatusByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetResourceGroupStatusByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetResourceGroupStatusByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetResourceGroupStatusByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetResourceGroupStatusByNameGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetResourceGroupStatusByNameOK creates a GetResourceGroupStatusByNameOK with default headers values
func NewGetResourceGroupStatusByNameOK() *GetResourceGroupStatusByNameOK {
	return &GetResourceGroupStatusByNameOK{}
}

/* GetResourceGroupStatusByNameOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetResourceGroupStatusByNameOK struct {
	Payload *swagger_models.TagStatusMsg
}

func (o *GetResourceGroupStatusByNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/projects/name/{name}/status][%d] getResourceGroupStatusByNameOK  %+v", 200, o.Payload)
}
func (o *GetResourceGroupStatusByNameOK) GetPayload() *swagger_models.TagStatusMsg {
	return o.Payload
}

func (o *GetResourceGroupStatusByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.TagStatusMsg)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourceGroupStatusByNameUnauthorized creates a GetResourceGroupStatusByNameUnauthorized with default headers values
func NewGetResourceGroupStatusByNameUnauthorized() *GetResourceGroupStatusByNameUnauthorized {
	return &GetResourceGroupStatusByNameUnauthorized{}
}

/* GetResourceGroupStatusByNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type GetResourceGroupStatusByNameUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetResourceGroupStatusByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/projects/name/{name}/status][%d] getResourceGroupStatusByNameUnauthorized  %+v", 401, o.Payload)
}
func (o *GetResourceGroupStatusByNameUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetResourceGroupStatusByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourceGroupStatusByNameForbidden creates a GetResourceGroupStatusByNameForbidden with default headers values
func NewGetResourceGroupStatusByNameForbidden() *GetResourceGroupStatusByNameForbidden {
	return &GetResourceGroupStatusByNameForbidden{}
}

/* GetResourceGroupStatusByNameForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type GetResourceGroupStatusByNameForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetResourceGroupStatusByNameForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/projects/name/{name}/status][%d] getResourceGroupStatusByNameForbidden  %+v", 403, o.Payload)
}
func (o *GetResourceGroupStatusByNameForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetResourceGroupStatusByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourceGroupStatusByNameNotFound creates a GetResourceGroupStatusByNameNotFound with default headers values
func NewGetResourceGroupStatusByNameNotFound() *GetResourceGroupStatusByNameNotFound {
	return &GetResourceGroupStatusByNameNotFound{}
}

/* GetResourceGroupStatusByNameNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type GetResourceGroupStatusByNameNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetResourceGroupStatusByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/projects/name/{name}/status][%d] getResourceGroupStatusByNameNotFound  %+v", 404, o.Payload)
}
func (o *GetResourceGroupStatusByNameNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetResourceGroupStatusByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourceGroupStatusByNameInternalServerError creates a GetResourceGroupStatusByNameInternalServerError with default headers values
func NewGetResourceGroupStatusByNameInternalServerError() *GetResourceGroupStatusByNameInternalServerError {
	return &GetResourceGroupStatusByNameInternalServerError{}
}

/* GetResourceGroupStatusByNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type GetResourceGroupStatusByNameInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetResourceGroupStatusByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/projects/name/{name}/status][%d] getResourceGroupStatusByNameInternalServerError  %+v", 500, o.Payload)
}
func (o *GetResourceGroupStatusByNameInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetResourceGroupStatusByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourceGroupStatusByNameGatewayTimeout creates a GetResourceGroupStatusByNameGatewayTimeout with default headers values
func NewGetResourceGroupStatusByNameGatewayTimeout() *GetResourceGroupStatusByNameGatewayTimeout {
	return &GetResourceGroupStatusByNameGatewayTimeout{}
}

/* GetResourceGroupStatusByNameGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type GetResourceGroupStatusByNameGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetResourceGroupStatusByNameGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/projects/name/{name}/status][%d] getResourceGroupStatusByNameGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetResourceGroupStatusByNameGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetResourceGroupStatusByNameGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

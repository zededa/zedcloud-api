// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package hardware_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// GetGlobalHardwareBrandByNameReader is a Reader for the GetGlobalHardwareBrandByName structure.
type GetGlobalHardwareBrandByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGlobalHardwareBrandByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGlobalHardwareBrandByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetGlobalHardwareBrandByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGlobalHardwareBrandByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGlobalHardwareBrandByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGlobalHardwareBrandByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetGlobalHardwareBrandByNameGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGlobalHardwareBrandByNameOK creates a GetGlobalHardwareBrandByNameOK with default headers values
func NewGetGlobalHardwareBrandByNameOK() *GetGlobalHardwareBrandByNameOK {
	return &GetGlobalHardwareBrandByNameOK{}
}

/* GetGlobalHardwareBrandByNameOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetGlobalHardwareBrandByNameOK struct {
	Payload *swagger_models.SysBrand
}

func (o *GetGlobalHardwareBrandByNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/brands/global/name/{name}][%d] getGlobalHardwareBrandByNameOK  %+v", 200, o.Payload)
}
func (o *GetGlobalHardwareBrandByNameOK) GetPayload() *swagger_models.SysBrand {
	return o.Payload
}

func (o *GetGlobalHardwareBrandByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.SysBrand)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGlobalHardwareBrandByNameUnauthorized creates a GetGlobalHardwareBrandByNameUnauthorized with default headers values
func NewGetGlobalHardwareBrandByNameUnauthorized() *GetGlobalHardwareBrandByNameUnauthorized {
	return &GetGlobalHardwareBrandByNameUnauthorized{}
}

/* GetGlobalHardwareBrandByNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type GetGlobalHardwareBrandByNameUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetGlobalHardwareBrandByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/brands/global/name/{name}][%d] getGlobalHardwareBrandByNameUnauthorized  %+v", 401, o.Payload)
}
func (o *GetGlobalHardwareBrandByNameUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetGlobalHardwareBrandByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGlobalHardwareBrandByNameForbidden creates a GetGlobalHardwareBrandByNameForbidden with default headers values
func NewGetGlobalHardwareBrandByNameForbidden() *GetGlobalHardwareBrandByNameForbidden {
	return &GetGlobalHardwareBrandByNameForbidden{}
}

/* GetGlobalHardwareBrandByNameForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type GetGlobalHardwareBrandByNameForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetGlobalHardwareBrandByNameForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/brands/global/name/{name}][%d] getGlobalHardwareBrandByNameForbidden  %+v", 403, o.Payload)
}
func (o *GetGlobalHardwareBrandByNameForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetGlobalHardwareBrandByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGlobalHardwareBrandByNameNotFound creates a GetGlobalHardwareBrandByNameNotFound with default headers values
func NewGetGlobalHardwareBrandByNameNotFound() *GetGlobalHardwareBrandByNameNotFound {
	return &GetGlobalHardwareBrandByNameNotFound{}
}

/* GetGlobalHardwareBrandByNameNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type GetGlobalHardwareBrandByNameNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetGlobalHardwareBrandByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/brands/global/name/{name}][%d] getGlobalHardwareBrandByNameNotFound  %+v", 404, o.Payload)
}
func (o *GetGlobalHardwareBrandByNameNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetGlobalHardwareBrandByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGlobalHardwareBrandByNameInternalServerError creates a GetGlobalHardwareBrandByNameInternalServerError with default headers values
func NewGetGlobalHardwareBrandByNameInternalServerError() *GetGlobalHardwareBrandByNameInternalServerError {
	return &GetGlobalHardwareBrandByNameInternalServerError{}
}

/* GetGlobalHardwareBrandByNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type GetGlobalHardwareBrandByNameInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetGlobalHardwareBrandByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/brands/global/name/{name}][%d] getGlobalHardwareBrandByNameInternalServerError  %+v", 500, o.Payload)
}
func (o *GetGlobalHardwareBrandByNameInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetGlobalHardwareBrandByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGlobalHardwareBrandByNameGatewayTimeout creates a GetGlobalHardwareBrandByNameGatewayTimeout with default headers values
func NewGetGlobalHardwareBrandByNameGatewayTimeout() *GetGlobalHardwareBrandByNameGatewayTimeout {
	return &GetGlobalHardwareBrandByNameGatewayTimeout{}
}

/* GetGlobalHardwareBrandByNameGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type GetGlobalHardwareBrandByNameGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetGlobalHardwareBrandByNameGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/brands/global/name/{name}][%d] getGlobalHardwareBrandByNameGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetGlobalHardwareBrandByNameGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetGlobalHardwareBrandByNameGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

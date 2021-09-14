// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package edge_application_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// GetEdgeApplicationBundleReader is a Reader for the GetEdgeApplicationBundle structure.
type GetEdgeApplicationBundleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEdgeApplicationBundleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEdgeApplicationBundleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetEdgeApplicationBundleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetEdgeApplicationBundleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetEdgeApplicationBundleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetEdgeApplicationBundleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetEdgeApplicationBundleGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetEdgeApplicationBundleOK creates a GetEdgeApplicationBundleOK with default headers values
func NewGetEdgeApplicationBundleOK() *GetEdgeApplicationBundleOK {
	return &GetEdgeApplicationBundleOK{}
}

/* GetEdgeApplicationBundleOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetEdgeApplicationBundleOK struct {
	Payload *swagger_models.App
}

func (o *GetEdgeApplicationBundleOK) Error() string {
	return fmt.Sprintf("[GET /v1/apps/id/{id}][%d] getEdgeApplicationBundleOK  %+v", 200, o.Payload)
}
func (o *GetEdgeApplicationBundleOK) GetPayload() *swagger_models.App {
	return o.Payload
}

func (o *GetEdgeApplicationBundleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.App)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeApplicationBundleUnauthorized creates a GetEdgeApplicationBundleUnauthorized with default headers values
func NewGetEdgeApplicationBundleUnauthorized() *GetEdgeApplicationBundleUnauthorized {
	return &GetEdgeApplicationBundleUnauthorized{}
}

/* GetEdgeApplicationBundleUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type GetEdgeApplicationBundleUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEdgeApplicationBundleUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/apps/id/{id}][%d] getEdgeApplicationBundleUnauthorized  %+v", 401, o.Payload)
}
func (o *GetEdgeApplicationBundleUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEdgeApplicationBundleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeApplicationBundleForbidden creates a GetEdgeApplicationBundleForbidden with default headers values
func NewGetEdgeApplicationBundleForbidden() *GetEdgeApplicationBundleForbidden {
	return &GetEdgeApplicationBundleForbidden{}
}

/* GetEdgeApplicationBundleForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have application level access permission for the operation or does not have access scope to the project.
*/
type GetEdgeApplicationBundleForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEdgeApplicationBundleForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/apps/id/{id}][%d] getEdgeApplicationBundleForbidden  %+v", 403, o.Payload)
}
func (o *GetEdgeApplicationBundleForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEdgeApplicationBundleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeApplicationBundleNotFound creates a GetEdgeApplicationBundleNotFound with default headers values
func NewGetEdgeApplicationBundleNotFound() *GetEdgeApplicationBundleNotFound {
	return &GetEdgeApplicationBundleNotFound{}
}

/* GetEdgeApplicationBundleNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type GetEdgeApplicationBundleNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEdgeApplicationBundleNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/apps/id/{id}][%d] getEdgeApplicationBundleNotFound  %+v", 404, o.Payload)
}
func (o *GetEdgeApplicationBundleNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEdgeApplicationBundleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeApplicationBundleInternalServerError creates a GetEdgeApplicationBundleInternalServerError with default headers values
func NewGetEdgeApplicationBundleInternalServerError() *GetEdgeApplicationBundleInternalServerError {
	return &GetEdgeApplicationBundleInternalServerError{}
}

/* GetEdgeApplicationBundleInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type GetEdgeApplicationBundleInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEdgeApplicationBundleInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/apps/id/{id}][%d] getEdgeApplicationBundleInternalServerError  %+v", 500, o.Payload)
}
func (o *GetEdgeApplicationBundleInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEdgeApplicationBundleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeApplicationBundleGatewayTimeout creates a GetEdgeApplicationBundleGatewayTimeout with default headers values
func NewGetEdgeApplicationBundleGatewayTimeout() *GetEdgeApplicationBundleGatewayTimeout {
	return &GetEdgeApplicationBundleGatewayTimeout{}
}

/* GetEdgeApplicationBundleGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type GetEdgeApplicationBundleGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEdgeApplicationBundleGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/apps/id/{id}][%d] getEdgeApplicationBundleGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetEdgeApplicationBundleGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEdgeApplicationBundleGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

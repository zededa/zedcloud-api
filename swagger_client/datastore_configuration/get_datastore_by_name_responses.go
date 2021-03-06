// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package datastore_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// GetDatastoreByNameReader is a Reader for the GetDatastoreByName structure.
type GetDatastoreByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDatastoreByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDatastoreByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetDatastoreByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDatastoreByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDatastoreByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDatastoreByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetDatastoreByNameGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDatastoreByNameOK creates a GetDatastoreByNameOK with default headers values
func NewGetDatastoreByNameOK() *GetDatastoreByNameOK {
	return &GetDatastoreByNameOK{}
}

/* GetDatastoreByNameOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetDatastoreByNameOK struct {
	Payload *swagger_models.DatastoreInfo
}

func (o *GetDatastoreByNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/datastores/name/{name}][%d] getDatastoreByNameOK  %+v", 200, o.Payload)
}
func (o *GetDatastoreByNameOK) GetPayload() *swagger_models.DatastoreInfo {
	return o.Payload
}

func (o *GetDatastoreByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.DatastoreInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDatastoreByNameUnauthorized creates a GetDatastoreByNameUnauthorized with default headers values
func NewGetDatastoreByNameUnauthorized() *GetDatastoreByNameUnauthorized {
	return &GetDatastoreByNameUnauthorized{}
}

/* GetDatastoreByNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type GetDatastoreByNameUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetDatastoreByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/datastores/name/{name}][%d] getDatastoreByNameUnauthorized  %+v", 401, o.Payload)
}
func (o *GetDatastoreByNameUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetDatastoreByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDatastoreByNameForbidden creates a GetDatastoreByNameForbidden with default headers values
func NewGetDatastoreByNameForbidden() *GetDatastoreByNameForbidden {
	return &GetDatastoreByNameForbidden{}
}

/* GetDatastoreByNameForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type GetDatastoreByNameForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetDatastoreByNameForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/datastores/name/{name}][%d] getDatastoreByNameForbidden  %+v", 403, o.Payload)
}
func (o *GetDatastoreByNameForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetDatastoreByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDatastoreByNameNotFound creates a GetDatastoreByNameNotFound with default headers values
func NewGetDatastoreByNameNotFound() *GetDatastoreByNameNotFound {
	return &GetDatastoreByNameNotFound{}
}

/* GetDatastoreByNameNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type GetDatastoreByNameNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetDatastoreByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/datastores/name/{name}][%d] getDatastoreByNameNotFound  %+v", 404, o.Payload)
}
func (o *GetDatastoreByNameNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetDatastoreByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDatastoreByNameInternalServerError creates a GetDatastoreByNameInternalServerError with default headers values
func NewGetDatastoreByNameInternalServerError() *GetDatastoreByNameInternalServerError {
	return &GetDatastoreByNameInternalServerError{}
}

/* GetDatastoreByNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type GetDatastoreByNameInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetDatastoreByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/datastores/name/{name}][%d] getDatastoreByNameInternalServerError  %+v", 500, o.Payload)
}
func (o *GetDatastoreByNameInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetDatastoreByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDatastoreByNameGatewayTimeout creates a GetDatastoreByNameGatewayTimeout with default headers values
func NewGetDatastoreByNameGatewayTimeout() *GetDatastoreByNameGatewayTimeout {
	return &GetDatastoreByNameGatewayTimeout{}
}

/* GetDatastoreByNameGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type GetDatastoreByNameGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetDatastoreByNameGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/datastores/name/{name}][%d] getDatastoreByNameGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetDatastoreByNameGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetDatastoreByNameGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

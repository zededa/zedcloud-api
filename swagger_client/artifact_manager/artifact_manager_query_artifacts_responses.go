// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package artifact_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// ArtifactManagerQueryArtifactsReader is a Reader for the ArtifactManagerQueryArtifacts structure.
type ArtifactManagerQueryArtifactsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ArtifactManagerQueryArtifactsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewArtifactManagerQueryArtifactsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewArtifactManagerQueryArtifactsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewArtifactManagerQueryArtifactsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewArtifactManagerQueryArtifactsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewArtifactManagerQueryArtifactsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewArtifactManagerQueryArtifactsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewArtifactManagerQueryArtifactsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewArtifactManagerQueryArtifactsOK creates a ArtifactManagerQueryArtifactsOK with default headers values
func NewArtifactManagerQueryArtifactsOK() *ArtifactManagerQueryArtifactsOK {
	return &ArtifactManagerQueryArtifactsOK{}
}

/* ArtifactManagerQueryArtifactsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ArtifactManagerQueryArtifactsOK struct {
	Payload *swagger_models.ArtifactList
}

func (o *ArtifactManagerQueryArtifactsOK) Error() string {
	return fmt.Sprintf("[GET /v1/artifacts][%d] artifactManagerQueryArtifactsOK  %+v", 200, o.Payload)
}
func (o *ArtifactManagerQueryArtifactsOK) GetPayload() *swagger_models.ArtifactList {
	return o.Payload
}

func (o *ArtifactManagerQueryArtifactsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ArtifactList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArtifactManagerQueryArtifactsBadRequest creates a ArtifactManagerQueryArtifactsBadRequest with default headers values
func NewArtifactManagerQueryArtifactsBadRequest() *ArtifactManagerQueryArtifactsBadRequest {
	return &ArtifactManagerQueryArtifactsBadRequest{}
}

/* ArtifactManagerQueryArtifactsBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of invalid value of filter parameters.
*/
type ArtifactManagerQueryArtifactsBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ArtifactManagerQueryArtifactsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/artifacts][%d] artifactManagerQueryArtifactsBadRequest  %+v", 400, o.Payload)
}
func (o *ArtifactManagerQueryArtifactsBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ArtifactManagerQueryArtifactsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArtifactManagerQueryArtifactsUnauthorized creates a ArtifactManagerQueryArtifactsUnauthorized with default headers values
func NewArtifactManagerQueryArtifactsUnauthorized() *ArtifactManagerQueryArtifactsUnauthorized {
	return &ArtifactManagerQueryArtifactsUnauthorized{}
}

/* ArtifactManagerQueryArtifactsUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type ArtifactManagerQueryArtifactsUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ArtifactManagerQueryArtifactsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/artifacts][%d] artifactManagerQueryArtifactsUnauthorized  %+v", 401, o.Payload)
}
func (o *ArtifactManagerQueryArtifactsUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ArtifactManagerQueryArtifactsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArtifactManagerQueryArtifactsForbidden creates a ArtifactManagerQueryArtifactsForbidden with default headers values
func NewArtifactManagerQueryArtifactsForbidden() *ArtifactManagerQueryArtifactsForbidden {
	return &ArtifactManagerQueryArtifactsForbidden{}
}

/* ArtifactManagerQueryArtifactsForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type ArtifactManagerQueryArtifactsForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ArtifactManagerQueryArtifactsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/artifacts][%d] artifactManagerQueryArtifactsForbidden  %+v", 403, o.Payload)
}
func (o *ArtifactManagerQueryArtifactsForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ArtifactManagerQueryArtifactsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArtifactManagerQueryArtifactsInternalServerError creates a ArtifactManagerQueryArtifactsInternalServerError with default headers values
func NewArtifactManagerQueryArtifactsInternalServerError() *ArtifactManagerQueryArtifactsInternalServerError {
	return &ArtifactManagerQueryArtifactsInternalServerError{}
}

/* ArtifactManagerQueryArtifactsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type ArtifactManagerQueryArtifactsInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ArtifactManagerQueryArtifactsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/artifacts][%d] artifactManagerQueryArtifactsInternalServerError  %+v", 500, o.Payload)
}
func (o *ArtifactManagerQueryArtifactsInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ArtifactManagerQueryArtifactsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArtifactManagerQueryArtifactsGatewayTimeout creates a ArtifactManagerQueryArtifactsGatewayTimeout with default headers values
func NewArtifactManagerQueryArtifactsGatewayTimeout() *ArtifactManagerQueryArtifactsGatewayTimeout {
	return &ArtifactManagerQueryArtifactsGatewayTimeout{}
}

/* ArtifactManagerQueryArtifactsGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type ArtifactManagerQueryArtifactsGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ArtifactManagerQueryArtifactsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/artifacts][%d] artifactManagerQueryArtifactsGatewayTimeout  %+v", 504, o.Payload)
}
func (o *ArtifactManagerQueryArtifactsGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ArtifactManagerQueryArtifactsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArtifactManagerQueryArtifactsDefault creates a ArtifactManagerQueryArtifactsDefault with default headers values
func NewArtifactManagerQueryArtifactsDefault(code int) *ArtifactManagerQueryArtifactsDefault {
	return &ArtifactManagerQueryArtifactsDefault{
		_statusCode: code,
	}
}

/* ArtifactManagerQueryArtifactsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ArtifactManagerQueryArtifactsDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the artifact manager query artifacts default response
func (o *ArtifactManagerQueryArtifactsDefault) Code() int {
	return o._statusCode
}

func (o *ArtifactManagerQueryArtifactsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/artifacts][%d] ArtifactManager_QueryArtifacts default  %+v", o._statusCode, o.Payload)
}
func (o *ArtifactManagerQueryArtifactsDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *ArtifactManagerQueryArtifactsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

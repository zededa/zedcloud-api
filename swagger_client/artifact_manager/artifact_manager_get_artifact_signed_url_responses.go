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

// ArtifactManagerGetArtifactSignedURLReader is a Reader for the ArtifactManagerGetArtifactSignedURL structure.
type ArtifactManagerGetArtifactSignedURLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ArtifactManagerGetArtifactSignedURLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewArtifactManagerGetArtifactSignedURLOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 307:
		result := NewArtifactManagerGetArtifactSignedURLTemporaryRedirect()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewArtifactManagerGetArtifactSignedURLBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewArtifactManagerGetArtifactSignedURLUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewArtifactManagerGetArtifactSignedURLForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewArtifactManagerGetArtifactSignedURLInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewArtifactManagerGetArtifactSignedURLGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewArtifactManagerGetArtifactSignedURLDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewArtifactManagerGetArtifactSignedURLOK creates a ArtifactManagerGetArtifactSignedURLOK with default headers values
func NewArtifactManagerGetArtifactSignedURLOK() *ArtifactManagerGetArtifactSignedURLOK {
	return &ArtifactManagerGetArtifactSignedURLOK{}
}

/* ArtifactManagerGetArtifactSignedURLOK describes a response with status code 200, with default header values.

A successful response.
*/
type ArtifactManagerGetArtifactSignedURLOK struct {
	Payload *swagger_models.Artifact
}

func (o *ArtifactManagerGetArtifactSignedURLOK) Error() string {
	return fmt.Sprintf("[GET /v1/artifacts/id/{id}/url][%d] artifactManagerGetArtifactSignedUrlOK  %+v", 200, o.Payload)
}
func (o *ArtifactManagerGetArtifactSignedURLOK) GetPayload() *swagger_models.Artifact {
	return o.Payload
}

func (o *ArtifactManagerGetArtifactSignedURLOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.Artifact)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArtifactManagerGetArtifactSignedURLTemporaryRedirect creates a ArtifactManagerGetArtifactSignedURLTemporaryRedirect with default headers values
func NewArtifactManagerGetArtifactSignedURLTemporaryRedirect() *ArtifactManagerGetArtifactSignedURLTemporaryRedirect {
	return &ArtifactManagerGetArtifactSignedURLTemporaryRedirect{}
}

/* ArtifactManagerGetArtifactSignedURLTemporaryRedirect describes a response with status code 307, with default header values.

Temporary Redirect. Returned when the requested artifactId is not available at the requested time
*/
type ArtifactManagerGetArtifactSignedURLTemporaryRedirect struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ArtifactManagerGetArtifactSignedURLTemporaryRedirect) Error() string {
	return fmt.Sprintf("[GET /v1/artifacts/id/{id}/url][%d] artifactManagerGetArtifactSignedUrlTemporaryRedirect  %+v", 307, o.Payload)
}
func (o *ArtifactManagerGetArtifactSignedURLTemporaryRedirect) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ArtifactManagerGetArtifactSignedURLTemporaryRedirect) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArtifactManagerGetArtifactSignedURLBadRequest creates a ArtifactManagerGetArtifactSignedURLBadRequest with default headers values
func NewArtifactManagerGetArtifactSignedURLBadRequest() *ArtifactManagerGetArtifactSignedURLBadRequest {
	return &ArtifactManagerGetArtifactSignedURLBadRequest{}
}

/* ArtifactManagerGetArtifactSignedURLBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of invalid value of parameters.
*/
type ArtifactManagerGetArtifactSignedURLBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ArtifactManagerGetArtifactSignedURLBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/artifacts/id/{id}/url][%d] artifactManagerGetArtifactSignedUrlBadRequest  %+v", 400, o.Payload)
}
func (o *ArtifactManagerGetArtifactSignedURLBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ArtifactManagerGetArtifactSignedURLBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArtifactManagerGetArtifactSignedURLUnauthorized creates a ArtifactManagerGetArtifactSignedURLUnauthorized with default headers values
func NewArtifactManagerGetArtifactSignedURLUnauthorized() *ArtifactManagerGetArtifactSignedURLUnauthorized {
	return &ArtifactManagerGetArtifactSignedURLUnauthorized{}
}

/* ArtifactManagerGetArtifactSignedURLUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type ArtifactManagerGetArtifactSignedURLUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ArtifactManagerGetArtifactSignedURLUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/artifacts/id/{id}/url][%d] artifactManagerGetArtifactSignedUrlUnauthorized  %+v", 401, o.Payload)
}
func (o *ArtifactManagerGetArtifactSignedURLUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ArtifactManagerGetArtifactSignedURLUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArtifactManagerGetArtifactSignedURLForbidden creates a ArtifactManagerGetArtifactSignedURLForbidden with default headers values
func NewArtifactManagerGetArtifactSignedURLForbidden() *ArtifactManagerGetArtifactSignedURLForbidden {
	return &ArtifactManagerGetArtifactSignedURLForbidden{}
}

/* ArtifactManagerGetArtifactSignedURLForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type ArtifactManagerGetArtifactSignedURLForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ArtifactManagerGetArtifactSignedURLForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/artifacts/id/{id}/url][%d] artifactManagerGetArtifactSignedUrlForbidden  %+v", 403, o.Payload)
}
func (o *ArtifactManagerGetArtifactSignedURLForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ArtifactManagerGetArtifactSignedURLForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArtifactManagerGetArtifactSignedURLInternalServerError creates a ArtifactManagerGetArtifactSignedURLInternalServerError with default headers values
func NewArtifactManagerGetArtifactSignedURLInternalServerError() *ArtifactManagerGetArtifactSignedURLInternalServerError {
	return &ArtifactManagerGetArtifactSignedURLInternalServerError{}
}

/* ArtifactManagerGetArtifactSignedURLInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type ArtifactManagerGetArtifactSignedURLInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ArtifactManagerGetArtifactSignedURLInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/artifacts/id/{id}/url][%d] artifactManagerGetArtifactSignedUrlInternalServerError  %+v", 500, o.Payload)
}
func (o *ArtifactManagerGetArtifactSignedURLInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ArtifactManagerGetArtifactSignedURLInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArtifactManagerGetArtifactSignedURLGatewayTimeout creates a ArtifactManagerGetArtifactSignedURLGatewayTimeout with default headers values
func NewArtifactManagerGetArtifactSignedURLGatewayTimeout() *ArtifactManagerGetArtifactSignedURLGatewayTimeout {
	return &ArtifactManagerGetArtifactSignedURLGatewayTimeout{}
}

/* ArtifactManagerGetArtifactSignedURLGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type ArtifactManagerGetArtifactSignedURLGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ArtifactManagerGetArtifactSignedURLGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/artifacts/id/{id}/url][%d] artifactManagerGetArtifactSignedUrlGatewayTimeout  %+v", 504, o.Payload)
}
func (o *ArtifactManagerGetArtifactSignedURLGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ArtifactManagerGetArtifactSignedURLGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArtifactManagerGetArtifactSignedURLDefault creates a ArtifactManagerGetArtifactSignedURLDefault with default headers values
func NewArtifactManagerGetArtifactSignedURLDefault(code int) *ArtifactManagerGetArtifactSignedURLDefault {
	return &ArtifactManagerGetArtifactSignedURLDefault{
		_statusCode: code,
	}
}

/* ArtifactManagerGetArtifactSignedURLDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ArtifactManagerGetArtifactSignedURLDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the artifact manager get artifact signed Url default response
func (o *ArtifactManagerGetArtifactSignedURLDefault) Code() int {
	return o._statusCode
}

func (o *ArtifactManagerGetArtifactSignedURLDefault) Error() string {
	return fmt.Sprintf("[GET /v1/artifacts/id/{id}/url][%d] ArtifactManager_GetArtifactSignedUrl default  %+v", o._statusCode, o.Payload)
}
func (o *ArtifactManagerGetArtifactSignedURLDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *ArtifactManagerGetArtifactSignedURLDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
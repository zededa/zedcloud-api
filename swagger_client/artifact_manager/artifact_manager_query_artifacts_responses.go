// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
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

/*
ArtifactManagerQueryArtifactsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ArtifactManagerQueryArtifactsOK struct {
	Payload *swagger_models.ArtifactList
}

// IsSuccess returns true when this artifact manager query artifacts o k response has a 2xx status code
func (o *ArtifactManagerQueryArtifactsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this artifact manager query artifacts o k response has a 3xx status code
func (o *ArtifactManagerQueryArtifactsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this artifact manager query artifacts o k response has a 4xx status code
func (o *ArtifactManagerQueryArtifactsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this artifact manager query artifacts o k response has a 5xx status code
func (o *ArtifactManagerQueryArtifactsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this artifact manager query artifacts o k response a status code equal to that given
func (o *ArtifactManagerQueryArtifactsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the artifact manager query artifacts o k response
func (o *ArtifactManagerQueryArtifactsOK) Code() int {
	return 200
}

func (o *ArtifactManagerQueryArtifactsOK) Error() string {
	return fmt.Sprintf("[GET /v1/artifacts][%d] artifactManagerQueryArtifactsOK  %+v", 200, o.Payload)
}

func (o *ArtifactManagerQueryArtifactsOK) String() string {
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

/*
ArtifactManagerQueryArtifactsBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of invalid value of filter parameters.
*/
type ArtifactManagerQueryArtifactsBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this artifact manager query artifacts bad request response has a 2xx status code
func (o *ArtifactManagerQueryArtifactsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this artifact manager query artifacts bad request response has a 3xx status code
func (o *ArtifactManagerQueryArtifactsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this artifact manager query artifacts bad request response has a 4xx status code
func (o *ArtifactManagerQueryArtifactsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this artifact manager query artifacts bad request response has a 5xx status code
func (o *ArtifactManagerQueryArtifactsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this artifact manager query artifacts bad request response a status code equal to that given
func (o *ArtifactManagerQueryArtifactsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the artifact manager query artifacts bad request response
func (o *ArtifactManagerQueryArtifactsBadRequest) Code() int {
	return 400
}

func (o *ArtifactManagerQueryArtifactsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/artifacts][%d] artifactManagerQueryArtifactsBadRequest  %+v", 400, o.Payload)
}

func (o *ArtifactManagerQueryArtifactsBadRequest) String() string {
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

/*
ArtifactManagerQueryArtifactsUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type ArtifactManagerQueryArtifactsUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this artifact manager query artifacts unauthorized response has a 2xx status code
func (o *ArtifactManagerQueryArtifactsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this artifact manager query artifacts unauthorized response has a 3xx status code
func (o *ArtifactManagerQueryArtifactsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this artifact manager query artifacts unauthorized response has a 4xx status code
func (o *ArtifactManagerQueryArtifactsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this artifact manager query artifacts unauthorized response has a 5xx status code
func (o *ArtifactManagerQueryArtifactsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this artifact manager query artifacts unauthorized response a status code equal to that given
func (o *ArtifactManagerQueryArtifactsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the artifact manager query artifacts unauthorized response
func (o *ArtifactManagerQueryArtifactsUnauthorized) Code() int {
	return 401
}

func (o *ArtifactManagerQueryArtifactsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/artifacts][%d] artifactManagerQueryArtifactsUnauthorized  %+v", 401, o.Payload)
}

func (o *ArtifactManagerQueryArtifactsUnauthorized) String() string {
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

/*
ArtifactManagerQueryArtifactsForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type ArtifactManagerQueryArtifactsForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this artifact manager query artifacts forbidden response has a 2xx status code
func (o *ArtifactManagerQueryArtifactsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this artifact manager query artifacts forbidden response has a 3xx status code
func (o *ArtifactManagerQueryArtifactsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this artifact manager query artifacts forbidden response has a 4xx status code
func (o *ArtifactManagerQueryArtifactsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this artifact manager query artifacts forbidden response has a 5xx status code
func (o *ArtifactManagerQueryArtifactsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this artifact manager query artifacts forbidden response a status code equal to that given
func (o *ArtifactManagerQueryArtifactsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the artifact manager query artifacts forbidden response
func (o *ArtifactManagerQueryArtifactsForbidden) Code() int {
	return 403
}

func (o *ArtifactManagerQueryArtifactsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/artifacts][%d] artifactManagerQueryArtifactsForbidden  %+v", 403, o.Payload)
}

func (o *ArtifactManagerQueryArtifactsForbidden) String() string {
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

/*
ArtifactManagerQueryArtifactsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type ArtifactManagerQueryArtifactsInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this artifact manager query artifacts internal server error response has a 2xx status code
func (o *ArtifactManagerQueryArtifactsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this artifact manager query artifacts internal server error response has a 3xx status code
func (o *ArtifactManagerQueryArtifactsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this artifact manager query artifacts internal server error response has a 4xx status code
func (o *ArtifactManagerQueryArtifactsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this artifact manager query artifacts internal server error response has a 5xx status code
func (o *ArtifactManagerQueryArtifactsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this artifact manager query artifacts internal server error response a status code equal to that given
func (o *ArtifactManagerQueryArtifactsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the artifact manager query artifacts internal server error response
func (o *ArtifactManagerQueryArtifactsInternalServerError) Code() int {
	return 500
}

func (o *ArtifactManagerQueryArtifactsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/artifacts][%d] artifactManagerQueryArtifactsInternalServerError  %+v", 500, o.Payload)
}

func (o *ArtifactManagerQueryArtifactsInternalServerError) String() string {
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

/*
ArtifactManagerQueryArtifactsGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type ArtifactManagerQueryArtifactsGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this artifact manager query artifacts gateway timeout response has a 2xx status code
func (o *ArtifactManagerQueryArtifactsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this artifact manager query artifacts gateway timeout response has a 3xx status code
func (o *ArtifactManagerQueryArtifactsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this artifact manager query artifacts gateway timeout response has a 4xx status code
func (o *ArtifactManagerQueryArtifactsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this artifact manager query artifacts gateway timeout response has a 5xx status code
func (o *ArtifactManagerQueryArtifactsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this artifact manager query artifacts gateway timeout response a status code equal to that given
func (o *ArtifactManagerQueryArtifactsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the artifact manager query artifacts gateway timeout response
func (o *ArtifactManagerQueryArtifactsGatewayTimeout) Code() int {
	return 504
}

func (o *ArtifactManagerQueryArtifactsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/artifacts][%d] artifactManagerQueryArtifactsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *ArtifactManagerQueryArtifactsGatewayTimeout) String() string {
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

/*
ArtifactManagerQueryArtifactsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ArtifactManagerQueryArtifactsDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// IsSuccess returns true when this artifact manager query artifacts default response has a 2xx status code
func (o *ArtifactManagerQueryArtifactsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this artifact manager query artifacts default response has a 3xx status code
func (o *ArtifactManagerQueryArtifactsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this artifact manager query artifacts default response has a 4xx status code
func (o *ArtifactManagerQueryArtifactsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this artifact manager query artifacts default response has a 5xx status code
func (o *ArtifactManagerQueryArtifactsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this artifact manager query artifacts default response a status code equal to that given
func (o *ArtifactManagerQueryArtifactsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the artifact manager query artifacts default response
func (o *ArtifactManagerQueryArtifactsDefault) Code() int {
	return o._statusCode
}

func (o *ArtifactManagerQueryArtifactsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/artifacts][%d] ArtifactManager_QueryArtifacts default  %+v", o._statusCode, o.Payload)
}

func (o *ArtifactManagerQueryArtifactsDefault) String() string {
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

// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package artifact_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// ArtifactManagerGetArtifactStreamReader is a Reader for the ArtifactManagerGetArtifactStream structure.
type ArtifactManagerGetArtifactStreamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ArtifactManagerGetArtifactStreamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewArtifactManagerGetArtifactStreamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 206:
		result := NewArtifactManagerGetArtifactStreamPartialContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 307:
		result := NewArtifactManagerGetArtifactStreamTemporaryRedirect()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewArtifactManagerGetArtifactStreamBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewArtifactManagerGetArtifactStreamUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewArtifactManagerGetArtifactStreamForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewArtifactManagerGetArtifactStreamInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewArtifactManagerGetArtifactStreamGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewArtifactManagerGetArtifactStreamDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewArtifactManagerGetArtifactStreamOK creates a ArtifactManagerGetArtifactStreamOK with default headers values
func NewArtifactManagerGetArtifactStreamOK() *ArtifactManagerGetArtifactStreamOK {
	return &ArtifactManagerGetArtifactStreamOK{}
}

/*
ArtifactManagerGetArtifactStreamOK describes a response with status code 200, with default header values.

artifact chunk data(streaming responses)
*/
type ArtifactManagerGetArtifactStreamOK struct {
	Payload *ArtifactManagerGetArtifactStreamOKBody
}

// IsSuccess returns true when this artifact manager get artifact stream o k response has a 2xx status code
func (o *ArtifactManagerGetArtifactStreamOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this artifact manager get artifact stream o k response has a 3xx status code
func (o *ArtifactManagerGetArtifactStreamOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this artifact manager get artifact stream o k response has a 4xx status code
func (o *ArtifactManagerGetArtifactStreamOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this artifact manager get artifact stream o k response has a 5xx status code
func (o *ArtifactManagerGetArtifactStreamOK) IsServerError() bool {
	return false
}

// IsCode returns true when this artifact manager get artifact stream o k response a status code equal to that given
func (o *ArtifactManagerGetArtifactStreamOK) IsCode(code int) bool {
	return code == 200
}

func (o *ArtifactManagerGetArtifactStreamOK) Error() string {
	return fmt.Sprintf("[GET /v1/artifacts/id/{id}][%d] artifactManagerGetArtifactStreamOK  %+v", 200, o.Payload)
}

func (o *ArtifactManagerGetArtifactStreamOK) String() string {
	return fmt.Sprintf("[GET /v1/artifacts/id/{id}][%d] artifactManagerGetArtifactStreamOK  %+v", 200, o.Payload)
}

func (o *ArtifactManagerGetArtifactStreamOK) GetPayload() *ArtifactManagerGetArtifactStreamOKBody {
	return o.Payload
}

func (o *ArtifactManagerGetArtifactStreamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ArtifactManagerGetArtifactStreamOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArtifactManagerGetArtifactStreamPartialContent creates a ArtifactManagerGetArtifactStreamPartialContent with default headers values
func NewArtifactManagerGetArtifactStreamPartialContent() *ArtifactManagerGetArtifactStreamPartialContent {
	return &ArtifactManagerGetArtifactStreamPartialContent{}
}

/*
ArtifactManagerGetArtifactStreamPartialContent describes a response with status code 206, with default header values.

Partial Content. The API gateway has fulfilled the partial GET request for the resource.
*/
type ArtifactManagerGetArtifactStreamPartialContent struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this artifact manager get artifact stream partial content response has a 2xx status code
func (o *ArtifactManagerGetArtifactStreamPartialContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this artifact manager get artifact stream partial content response has a 3xx status code
func (o *ArtifactManagerGetArtifactStreamPartialContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this artifact manager get artifact stream partial content response has a 4xx status code
func (o *ArtifactManagerGetArtifactStreamPartialContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this artifact manager get artifact stream partial content response has a 5xx status code
func (o *ArtifactManagerGetArtifactStreamPartialContent) IsServerError() bool {
	return false
}

// IsCode returns true when this artifact manager get artifact stream partial content response a status code equal to that given
func (o *ArtifactManagerGetArtifactStreamPartialContent) IsCode(code int) bool {
	return code == 206
}

func (o *ArtifactManagerGetArtifactStreamPartialContent) Error() string {
	return fmt.Sprintf("[GET /v1/artifacts/id/{id}][%d] artifactManagerGetArtifactStreamPartialContent  %+v", 206, o.Payload)
}

func (o *ArtifactManagerGetArtifactStreamPartialContent) String() string {
	return fmt.Sprintf("[GET /v1/artifacts/id/{id}][%d] artifactManagerGetArtifactStreamPartialContent  %+v", 206, o.Payload)
}

func (o *ArtifactManagerGetArtifactStreamPartialContent) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ArtifactManagerGetArtifactStreamPartialContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArtifactManagerGetArtifactStreamTemporaryRedirect creates a ArtifactManagerGetArtifactStreamTemporaryRedirect with default headers values
func NewArtifactManagerGetArtifactStreamTemporaryRedirect() *ArtifactManagerGetArtifactStreamTemporaryRedirect {
	return &ArtifactManagerGetArtifactStreamTemporaryRedirect{}
}

/*
ArtifactManagerGetArtifactStreamTemporaryRedirect describes a response with status code 307, with default header values.

Temporary Redirect. Returned when the requested artifactId is not available at the requested time
*/
type ArtifactManagerGetArtifactStreamTemporaryRedirect struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this artifact manager get artifact stream temporary redirect response has a 2xx status code
func (o *ArtifactManagerGetArtifactStreamTemporaryRedirect) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this artifact manager get artifact stream temporary redirect response has a 3xx status code
func (o *ArtifactManagerGetArtifactStreamTemporaryRedirect) IsRedirect() bool {
	return true
}

// IsClientError returns true when this artifact manager get artifact stream temporary redirect response has a 4xx status code
func (o *ArtifactManagerGetArtifactStreamTemporaryRedirect) IsClientError() bool {
	return false
}

// IsServerError returns true when this artifact manager get artifact stream temporary redirect response has a 5xx status code
func (o *ArtifactManagerGetArtifactStreamTemporaryRedirect) IsServerError() bool {
	return false
}

// IsCode returns true when this artifact manager get artifact stream temporary redirect response a status code equal to that given
func (o *ArtifactManagerGetArtifactStreamTemporaryRedirect) IsCode(code int) bool {
	return code == 307
}

func (o *ArtifactManagerGetArtifactStreamTemporaryRedirect) Error() string {
	return fmt.Sprintf("[GET /v1/artifacts/id/{id}][%d] artifactManagerGetArtifactStreamTemporaryRedirect  %+v", 307, o.Payload)
}

func (o *ArtifactManagerGetArtifactStreamTemporaryRedirect) String() string {
	return fmt.Sprintf("[GET /v1/artifacts/id/{id}][%d] artifactManagerGetArtifactStreamTemporaryRedirect  %+v", 307, o.Payload)
}

func (o *ArtifactManagerGetArtifactStreamTemporaryRedirect) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ArtifactManagerGetArtifactStreamTemporaryRedirect) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArtifactManagerGetArtifactStreamBadRequest creates a ArtifactManagerGetArtifactStreamBadRequest with default headers values
func NewArtifactManagerGetArtifactStreamBadRequest() *ArtifactManagerGetArtifactStreamBadRequest {
	return &ArtifactManagerGetArtifactStreamBadRequest{}
}

/*
ArtifactManagerGetArtifactStreamBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of invalid value of parameters.
*/
type ArtifactManagerGetArtifactStreamBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this artifact manager get artifact stream bad request response has a 2xx status code
func (o *ArtifactManagerGetArtifactStreamBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this artifact manager get artifact stream bad request response has a 3xx status code
func (o *ArtifactManagerGetArtifactStreamBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this artifact manager get artifact stream bad request response has a 4xx status code
func (o *ArtifactManagerGetArtifactStreamBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this artifact manager get artifact stream bad request response has a 5xx status code
func (o *ArtifactManagerGetArtifactStreamBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this artifact manager get artifact stream bad request response a status code equal to that given
func (o *ArtifactManagerGetArtifactStreamBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ArtifactManagerGetArtifactStreamBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/artifacts/id/{id}][%d] artifactManagerGetArtifactStreamBadRequest  %+v", 400, o.Payload)
}

func (o *ArtifactManagerGetArtifactStreamBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/artifacts/id/{id}][%d] artifactManagerGetArtifactStreamBadRequest  %+v", 400, o.Payload)
}

func (o *ArtifactManagerGetArtifactStreamBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ArtifactManagerGetArtifactStreamBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArtifactManagerGetArtifactStreamUnauthorized creates a ArtifactManagerGetArtifactStreamUnauthorized with default headers values
func NewArtifactManagerGetArtifactStreamUnauthorized() *ArtifactManagerGetArtifactStreamUnauthorized {
	return &ArtifactManagerGetArtifactStreamUnauthorized{}
}

/*
ArtifactManagerGetArtifactStreamUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type ArtifactManagerGetArtifactStreamUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this artifact manager get artifact stream unauthorized response has a 2xx status code
func (o *ArtifactManagerGetArtifactStreamUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this artifact manager get artifact stream unauthorized response has a 3xx status code
func (o *ArtifactManagerGetArtifactStreamUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this artifact manager get artifact stream unauthorized response has a 4xx status code
func (o *ArtifactManagerGetArtifactStreamUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this artifact manager get artifact stream unauthorized response has a 5xx status code
func (o *ArtifactManagerGetArtifactStreamUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this artifact manager get artifact stream unauthorized response a status code equal to that given
func (o *ArtifactManagerGetArtifactStreamUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ArtifactManagerGetArtifactStreamUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/artifacts/id/{id}][%d] artifactManagerGetArtifactStreamUnauthorized  %+v", 401, o.Payload)
}

func (o *ArtifactManagerGetArtifactStreamUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/artifacts/id/{id}][%d] artifactManagerGetArtifactStreamUnauthorized  %+v", 401, o.Payload)
}

func (o *ArtifactManagerGetArtifactStreamUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ArtifactManagerGetArtifactStreamUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArtifactManagerGetArtifactStreamForbidden creates a ArtifactManagerGetArtifactStreamForbidden with default headers values
func NewArtifactManagerGetArtifactStreamForbidden() *ArtifactManagerGetArtifactStreamForbidden {
	return &ArtifactManagerGetArtifactStreamForbidden{}
}

/*
ArtifactManagerGetArtifactStreamForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type ArtifactManagerGetArtifactStreamForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this artifact manager get artifact stream forbidden response has a 2xx status code
func (o *ArtifactManagerGetArtifactStreamForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this artifact manager get artifact stream forbidden response has a 3xx status code
func (o *ArtifactManagerGetArtifactStreamForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this artifact manager get artifact stream forbidden response has a 4xx status code
func (o *ArtifactManagerGetArtifactStreamForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this artifact manager get artifact stream forbidden response has a 5xx status code
func (o *ArtifactManagerGetArtifactStreamForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this artifact manager get artifact stream forbidden response a status code equal to that given
func (o *ArtifactManagerGetArtifactStreamForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ArtifactManagerGetArtifactStreamForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/artifacts/id/{id}][%d] artifactManagerGetArtifactStreamForbidden  %+v", 403, o.Payload)
}

func (o *ArtifactManagerGetArtifactStreamForbidden) String() string {
	return fmt.Sprintf("[GET /v1/artifacts/id/{id}][%d] artifactManagerGetArtifactStreamForbidden  %+v", 403, o.Payload)
}

func (o *ArtifactManagerGetArtifactStreamForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ArtifactManagerGetArtifactStreamForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArtifactManagerGetArtifactStreamInternalServerError creates a ArtifactManagerGetArtifactStreamInternalServerError with default headers values
func NewArtifactManagerGetArtifactStreamInternalServerError() *ArtifactManagerGetArtifactStreamInternalServerError {
	return &ArtifactManagerGetArtifactStreamInternalServerError{}
}

/*
ArtifactManagerGetArtifactStreamInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type ArtifactManagerGetArtifactStreamInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this artifact manager get artifact stream internal server error response has a 2xx status code
func (o *ArtifactManagerGetArtifactStreamInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this artifact manager get artifact stream internal server error response has a 3xx status code
func (o *ArtifactManagerGetArtifactStreamInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this artifact manager get artifact stream internal server error response has a 4xx status code
func (o *ArtifactManagerGetArtifactStreamInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this artifact manager get artifact stream internal server error response has a 5xx status code
func (o *ArtifactManagerGetArtifactStreamInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this artifact manager get artifact stream internal server error response a status code equal to that given
func (o *ArtifactManagerGetArtifactStreamInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ArtifactManagerGetArtifactStreamInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/artifacts/id/{id}][%d] artifactManagerGetArtifactStreamInternalServerError  %+v", 500, o.Payload)
}

func (o *ArtifactManagerGetArtifactStreamInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/artifacts/id/{id}][%d] artifactManagerGetArtifactStreamInternalServerError  %+v", 500, o.Payload)
}

func (o *ArtifactManagerGetArtifactStreamInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ArtifactManagerGetArtifactStreamInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArtifactManagerGetArtifactStreamGatewayTimeout creates a ArtifactManagerGetArtifactStreamGatewayTimeout with default headers values
func NewArtifactManagerGetArtifactStreamGatewayTimeout() *ArtifactManagerGetArtifactStreamGatewayTimeout {
	return &ArtifactManagerGetArtifactStreamGatewayTimeout{}
}

/*
ArtifactManagerGetArtifactStreamGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type ArtifactManagerGetArtifactStreamGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this artifact manager get artifact stream gateway timeout response has a 2xx status code
func (o *ArtifactManagerGetArtifactStreamGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this artifact manager get artifact stream gateway timeout response has a 3xx status code
func (o *ArtifactManagerGetArtifactStreamGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this artifact manager get artifact stream gateway timeout response has a 4xx status code
func (o *ArtifactManagerGetArtifactStreamGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this artifact manager get artifact stream gateway timeout response has a 5xx status code
func (o *ArtifactManagerGetArtifactStreamGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this artifact manager get artifact stream gateway timeout response a status code equal to that given
func (o *ArtifactManagerGetArtifactStreamGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *ArtifactManagerGetArtifactStreamGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/artifacts/id/{id}][%d] artifactManagerGetArtifactStreamGatewayTimeout  %+v", 504, o.Payload)
}

func (o *ArtifactManagerGetArtifactStreamGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/artifacts/id/{id}][%d] artifactManagerGetArtifactStreamGatewayTimeout  %+v", 504, o.Payload)
}

func (o *ArtifactManagerGetArtifactStreamGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ArtifactManagerGetArtifactStreamGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArtifactManagerGetArtifactStreamDefault creates a ArtifactManagerGetArtifactStreamDefault with default headers values
func NewArtifactManagerGetArtifactStreamDefault(code int) *ArtifactManagerGetArtifactStreamDefault {
	return &ArtifactManagerGetArtifactStreamDefault{
		_statusCode: code,
	}
}

/*
ArtifactManagerGetArtifactStreamDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ArtifactManagerGetArtifactStreamDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the artifact manager get artifact stream default response
func (o *ArtifactManagerGetArtifactStreamDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this artifact manager get artifact stream default response has a 2xx status code
func (o *ArtifactManagerGetArtifactStreamDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this artifact manager get artifact stream default response has a 3xx status code
func (o *ArtifactManagerGetArtifactStreamDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this artifact manager get artifact stream default response has a 4xx status code
func (o *ArtifactManagerGetArtifactStreamDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this artifact manager get artifact stream default response has a 5xx status code
func (o *ArtifactManagerGetArtifactStreamDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this artifact manager get artifact stream default response a status code equal to that given
func (o *ArtifactManagerGetArtifactStreamDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ArtifactManagerGetArtifactStreamDefault) Error() string {
	return fmt.Sprintf("[GET /v1/artifacts/id/{id}][%d] ArtifactManager_GetArtifactStream default  %+v", o._statusCode, o.Payload)
}

func (o *ArtifactManagerGetArtifactStreamDefault) String() string {
	return fmt.Sprintf("[GET /v1/artifacts/id/{id}][%d] ArtifactManager_GetArtifactStream default  %+v", o._statusCode, o.Payload)
}

func (o *ArtifactManagerGetArtifactStreamDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *ArtifactManagerGetArtifactStreamDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ArtifactManagerGetArtifactStreamOKBody Stream result of ArtifactStream
//
// artifact chunk data
swagger:model ArtifactManagerGetArtifactStreamOKBody
*/
type ArtifactManagerGetArtifactStreamOKBody struct {

	// error
	Error *swagger_models.GooglerpcStatus `json:"error,omitempty"`

	// result
	Result interface{} `json:"result,omitempty"`
}

// Validate validates this artifact manager get artifact stream o k body
func (o *ArtifactManagerGetArtifactStreamOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ArtifactManagerGetArtifactStreamOKBody) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("artifactManagerGetArtifactStreamOK" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("artifactManagerGetArtifactStreamOK" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this artifact manager get artifact stream o k body based on the context it is used
func (o *ArtifactManagerGetArtifactStreamOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ArtifactManagerGetArtifactStreamOKBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {
		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("artifactManagerGetArtifactStreamOK" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("artifactManagerGetArtifactStreamOK" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ArtifactManagerGetArtifactStreamOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ArtifactManagerGetArtifactStreamOKBody) UnmarshalBinary(b []byte) error {
	var res ArtifactManagerGetArtifactStreamOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

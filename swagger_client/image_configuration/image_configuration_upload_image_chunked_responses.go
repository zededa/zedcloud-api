// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package image_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// ImageConfigurationUploadImageChunkedReader is a Reader for the ImageConfigurationUploadImageChunked structure.
type ImageConfigurationUploadImageChunkedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImageConfigurationUploadImageChunkedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImageConfigurationUploadImageChunkedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewImageConfigurationUploadImageChunkedAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewImageConfigurationUploadImageChunkedBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewImageConfigurationUploadImageChunkedUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewImageConfigurationUploadImageChunkedForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewImageConfigurationUploadImageChunkedNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewImageConfigurationUploadImageChunkedConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImageConfigurationUploadImageChunkedInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewImageConfigurationUploadImageChunkedGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewImageConfigurationUploadImageChunkedDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewImageConfigurationUploadImageChunkedOK creates a ImageConfigurationUploadImageChunkedOK with default headers values
func NewImageConfigurationUploadImageChunkedOK() *ImageConfigurationUploadImageChunkedOK {
	return &ImageConfigurationUploadImageChunkedOK{}
}

/*
ImageConfigurationUploadImageChunkedOK describes a response with status code 200, with default header values.

A successful response.
*/
type ImageConfigurationUploadImageChunkedOK struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this image configuration upload image chunked o k response has a 2xx status code
func (o *ImageConfigurationUploadImageChunkedOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this image configuration upload image chunked o k response has a 3xx status code
func (o *ImageConfigurationUploadImageChunkedOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration upload image chunked o k response has a 4xx status code
func (o *ImageConfigurationUploadImageChunkedOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this image configuration upload image chunked o k response has a 5xx status code
func (o *ImageConfigurationUploadImageChunkedOK) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration upload image chunked o k response a status code equal to that given
func (o *ImageConfigurationUploadImageChunkedOK) IsCode(code int) bool {
	return code == 200
}

func (o *ImageConfigurationUploadImageChunkedOK) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/upload/chunked][%d] imageConfigurationUploadImageChunkedOK  %+v", 200, o.Payload)
}

func (o *ImageConfigurationUploadImageChunkedOK) String() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/upload/chunked][%d] imageConfigurationUploadImageChunkedOK  %+v", 200, o.Payload)
}

func (o *ImageConfigurationUploadImageChunkedOK) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationUploadImageChunkedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationUploadImageChunkedAccepted creates a ImageConfigurationUploadImageChunkedAccepted with default headers values
func NewImageConfigurationUploadImageChunkedAccepted() *ImageConfigurationUploadImageChunkedAccepted {
	return &ImageConfigurationUploadImageChunkedAccepted{}
}

/*
ImageConfigurationUploadImageChunkedAccepted describes a response with status code 202, with default header values.

Accepted. The API gateway accepted the request for uploading but the upload process has not been completed. Please check ImageStatus and ImageError fields to track the status of upload process and any error messages.
*/
type ImageConfigurationUploadImageChunkedAccepted struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this image configuration upload image chunked accepted response has a 2xx status code
func (o *ImageConfigurationUploadImageChunkedAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this image configuration upload image chunked accepted response has a 3xx status code
func (o *ImageConfigurationUploadImageChunkedAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration upload image chunked accepted response has a 4xx status code
func (o *ImageConfigurationUploadImageChunkedAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this image configuration upload image chunked accepted response has a 5xx status code
func (o *ImageConfigurationUploadImageChunkedAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration upload image chunked accepted response a status code equal to that given
func (o *ImageConfigurationUploadImageChunkedAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *ImageConfigurationUploadImageChunkedAccepted) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/upload/chunked][%d] imageConfigurationUploadImageChunkedAccepted  %+v", 202, o.Payload)
}

func (o *ImageConfigurationUploadImageChunkedAccepted) String() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/upload/chunked][%d] imageConfigurationUploadImageChunkedAccepted  %+v", 202, o.Payload)
}

func (o *ImageConfigurationUploadImageChunkedAccepted) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationUploadImageChunkedAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationUploadImageChunkedBadRequest creates a ImageConfigurationUploadImageChunkedBadRequest with default headers values
func NewImageConfigurationUploadImageChunkedBadRequest() *ImageConfigurationUploadImageChunkedBadRequest {
	return &ImageConfigurationUploadImageChunkedBadRequest{}
}

/*
ImageConfigurationUploadImageChunkedBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type ImageConfigurationUploadImageChunkedBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this image configuration upload image chunked bad request response has a 2xx status code
func (o *ImageConfigurationUploadImageChunkedBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration upload image chunked bad request response has a 3xx status code
func (o *ImageConfigurationUploadImageChunkedBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration upload image chunked bad request response has a 4xx status code
func (o *ImageConfigurationUploadImageChunkedBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration upload image chunked bad request response has a 5xx status code
func (o *ImageConfigurationUploadImageChunkedBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration upload image chunked bad request response a status code equal to that given
func (o *ImageConfigurationUploadImageChunkedBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ImageConfigurationUploadImageChunkedBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/upload/chunked][%d] imageConfigurationUploadImageChunkedBadRequest  %+v", 400, o.Payload)
}

func (o *ImageConfigurationUploadImageChunkedBadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/upload/chunked][%d] imageConfigurationUploadImageChunkedBadRequest  %+v", 400, o.Payload)
}

func (o *ImageConfigurationUploadImageChunkedBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationUploadImageChunkedBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationUploadImageChunkedUnauthorized creates a ImageConfigurationUploadImageChunkedUnauthorized with default headers values
func NewImageConfigurationUploadImageChunkedUnauthorized() *ImageConfigurationUploadImageChunkedUnauthorized {
	return &ImageConfigurationUploadImageChunkedUnauthorized{}
}

/*
ImageConfigurationUploadImageChunkedUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type ImageConfigurationUploadImageChunkedUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this image configuration upload image chunked unauthorized response has a 2xx status code
func (o *ImageConfigurationUploadImageChunkedUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration upload image chunked unauthorized response has a 3xx status code
func (o *ImageConfigurationUploadImageChunkedUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration upload image chunked unauthorized response has a 4xx status code
func (o *ImageConfigurationUploadImageChunkedUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration upload image chunked unauthorized response has a 5xx status code
func (o *ImageConfigurationUploadImageChunkedUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration upload image chunked unauthorized response a status code equal to that given
func (o *ImageConfigurationUploadImageChunkedUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ImageConfigurationUploadImageChunkedUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/upload/chunked][%d] imageConfigurationUploadImageChunkedUnauthorized  %+v", 401, o.Payload)
}

func (o *ImageConfigurationUploadImageChunkedUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/upload/chunked][%d] imageConfigurationUploadImageChunkedUnauthorized  %+v", 401, o.Payload)
}

func (o *ImageConfigurationUploadImageChunkedUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationUploadImageChunkedUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationUploadImageChunkedForbidden creates a ImageConfigurationUploadImageChunkedForbidden with default headers values
func NewImageConfigurationUploadImageChunkedForbidden() *ImageConfigurationUploadImageChunkedForbidden {
	return &ImageConfigurationUploadImageChunkedForbidden{}
}

/*
ImageConfigurationUploadImageChunkedForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type ImageConfigurationUploadImageChunkedForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this image configuration upload image chunked forbidden response has a 2xx status code
func (o *ImageConfigurationUploadImageChunkedForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration upload image chunked forbidden response has a 3xx status code
func (o *ImageConfigurationUploadImageChunkedForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration upload image chunked forbidden response has a 4xx status code
func (o *ImageConfigurationUploadImageChunkedForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration upload image chunked forbidden response has a 5xx status code
func (o *ImageConfigurationUploadImageChunkedForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration upload image chunked forbidden response a status code equal to that given
func (o *ImageConfigurationUploadImageChunkedForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ImageConfigurationUploadImageChunkedForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/upload/chunked][%d] imageConfigurationUploadImageChunkedForbidden  %+v", 403, o.Payload)
}

func (o *ImageConfigurationUploadImageChunkedForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/upload/chunked][%d] imageConfigurationUploadImageChunkedForbidden  %+v", 403, o.Payload)
}

func (o *ImageConfigurationUploadImageChunkedForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationUploadImageChunkedForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationUploadImageChunkedNotFound creates a ImageConfigurationUploadImageChunkedNotFound with default headers values
func NewImageConfigurationUploadImageChunkedNotFound() *ImageConfigurationUploadImageChunkedNotFound {
	return &ImageConfigurationUploadImageChunkedNotFound{}
}

/*
ImageConfigurationUploadImageChunkedNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type ImageConfigurationUploadImageChunkedNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this image configuration upload image chunked not found response has a 2xx status code
func (o *ImageConfigurationUploadImageChunkedNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration upload image chunked not found response has a 3xx status code
func (o *ImageConfigurationUploadImageChunkedNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration upload image chunked not found response has a 4xx status code
func (o *ImageConfigurationUploadImageChunkedNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration upload image chunked not found response has a 5xx status code
func (o *ImageConfigurationUploadImageChunkedNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration upload image chunked not found response a status code equal to that given
func (o *ImageConfigurationUploadImageChunkedNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ImageConfigurationUploadImageChunkedNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/upload/chunked][%d] imageConfigurationUploadImageChunkedNotFound  %+v", 404, o.Payload)
}

func (o *ImageConfigurationUploadImageChunkedNotFound) String() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/upload/chunked][%d] imageConfigurationUploadImageChunkedNotFound  %+v", 404, o.Payload)
}

func (o *ImageConfigurationUploadImageChunkedNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationUploadImageChunkedNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationUploadImageChunkedConflict creates a ImageConfigurationUploadImageChunkedConflict with default headers values
func NewImageConfigurationUploadImageChunkedConflict() *ImageConfigurationUploadImageChunkedConflict {
	return &ImageConfigurationUploadImageChunkedConflict{}
}

/*
ImageConfigurationUploadImageChunkedConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because
-- another request for uplink / upload is already in progress
-- image has been already uploaded, can't be modified again
*/
type ImageConfigurationUploadImageChunkedConflict struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this image configuration upload image chunked conflict response has a 2xx status code
func (o *ImageConfigurationUploadImageChunkedConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration upload image chunked conflict response has a 3xx status code
func (o *ImageConfigurationUploadImageChunkedConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration upload image chunked conflict response has a 4xx status code
func (o *ImageConfigurationUploadImageChunkedConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration upload image chunked conflict response has a 5xx status code
func (o *ImageConfigurationUploadImageChunkedConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration upload image chunked conflict response a status code equal to that given
func (o *ImageConfigurationUploadImageChunkedConflict) IsCode(code int) bool {
	return code == 409
}

func (o *ImageConfigurationUploadImageChunkedConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/upload/chunked][%d] imageConfigurationUploadImageChunkedConflict  %+v", 409, o.Payload)
}

func (o *ImageConfigurationUploadImageChunkedConflict) String() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/upload/chunked][%d] imageConfigurationUploadImageChunkedConflict  %+v", 409, o.Payload)
}

func (o *ImageConfigurationUploadImageChunkedConflict) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationUploadImageChunkedConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationUploadImageChunkedInternalServerError creates a ImageConfigurationUploadImageChunkedInternalServerError with default headers values
func NewImageConfigurationUploadImageChunkedInternalServerError() *ImageConfigurationUploadImageChunkedInternalServerError {
	return &ImageConfigurationUploadImageChunkedInternalServerError{}
}

/*
ImageConfigurationUploadImageChunkedInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type ImageConfigurationUploadImageChunkedInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this image configuration upload image chunked internal server error response has a 2xx status code
func (o *ImageConfigurationUploadImageChunkedInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration upload image chunked internal server error response has a 3xx status code
func (o *ImageConfigurationUploadImageChunkedInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration upload image chunked internal server error response has a 4xx status code
func (o *ImageConfigurationUploadImageChunkedInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this image configuration upload image chunked internal server error response has a 5xx status code
func (o *ImageConfigurationUploadImageChunkedInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this image configuration upload image chunked internal server error response a status code equal to that given
func (o *ImageConfigurationUploadImageChunkedInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ImageConfigurationUploadImageChunkedInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/upload/chunked][%d] imageConfigurationUploadImageChunkedInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageConfigurationUploadImageChunkedInternalServerError) String() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/upload/chunked][%d] imageConfigurationUploadImageChunkedInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageConfigurationUploadImageChunkedInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationUploadImageChunkedInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationUploadImageChunkedGatewayTimeout creates a ImageConfigurationUploadImageChunkedGatewayTimeout with default headers values
func NewImageConfigurationUploadImageChunkedGatewayTimeout() *ImageConfigurationUploadImageChunkedGatewayTimeout {
	return &ImageConfigurationUploadImageChunkedGatewayTimeout{}
}

/*
ImageConfigurationUploadImageChunkedGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type ImageConfigurationUploadImageChunkedGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this image configuration upload image chunked gateway timeout response has a 2xx status code
func (o *ImageConfigurationUploadImageChunkedGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration upload image chunked gateway timeout response has a 3xx status code
func (o *ImageConfigurationUploadImageChunkedGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration upload image chunked gateway timeout response has a 4xx status code
func (o *ImageConfigurationUploadImageChunkedGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this image configuration upload image chunked gateway timeout response has a 5xx status code
func (o *ImageConfigurationUploadImageChunkedGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this image configuration upload image chunked gateway timeout response a status code equal to that given
func (o *ImageConfigurationUploadImageChunkedGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *ImageConfigurationUploadImageChunkedGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/upload/chunked][%d] imageConfigurationUploadImageChunkedGatewayTimeout  %+v", 504, o.Payload)
}

func (o *ImageConfigurationUploadImageChunkedGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/upload/chunked][%d] imageConfigurationUploadImageChunkedGatewayTimeout  %+v", 504, o.Payload)
}

func (o *ImageConfigurationUploadImageChunkedGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationUploadImageChunkedGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationUploadImageChunkedDefault creates a ImageConfigurationUploadImageChunkedDefault with default headers values
func NewImageConfigurationUploadImageChunkedDefault(code int) *ImageConfigurationUploadImageChunkedDefault {
	return &ImageConfigurationUploadImageChunkedDefault{
		_statusCode: code,
	}
}

/*
ImageConfigurationUploadImageChunkedDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ImageConfigurationUploadImageChunkedDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the image configuration upload image chunked default response
func (o *ImageConfigurationUploadImageChunkedDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this image configuration upload image chunked default response has a 2xx status code
func (o *ImageConfigurationUploadImageChunkedDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this image configuration upload image chunked default response has a 3xx status code
func (o *ImageConfigurationUploadImageChunkedDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this image configuration upload image chunked default response has a 4xx status code
func (o *ImageConfigurationUploadImageChunkedDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this image configuration upload image chunked default response has a 5xx status code
func (o *ImageConfigurationUploadImageChunkedDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this image configuration upload image chunked default response a status code equal to that given
func (o *ImageConfigurationUploadImageChunkedDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ImageConfigurationUploadImageChunkedDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/upload/chunked][%d] ImageConfiguration_UploadImageChunked default  %+v", o._statusCode, o.Payload)
}

func (o *ImageConfigurationUploadImageChunkedDefault) String() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/upload/chunked][%d] ImageConfiguration_UploadImageChunked default  %+v", o._statusCode, o.Payload)
}

func (o *ImageConfigurationUploadImageChunkedDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *ImageConfigurationUploadImageChunkedDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

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

// DeleteImageReader is a Reader for the DeleteImage structure.
type DeleteImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteImageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteImageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteImageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteImageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteImageConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteImageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteImageGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteImageOK creates a DeleteImageOK with default headers values
func NewDeleteImageOK() *DeleteImageOK {
	return &DeleteImageOK{}
}

/* DeleteImageOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteImageOK struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *DeleteImageOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/images/id/{id}][%d] deleteImageOK  %+v", 200, o.Payload)
}
func (o *DeleteImageOK) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *DeleteImageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteImageUnauthorized creates a DeleteImageUnauthorized with default headers values
func NewDeleteImageUnauthorized() *DeleteImageUnauthorized {
	return &DeleteImageUnauthorized{}
}

/* DeleteImageUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type DeleteImageUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *DeleteImageUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/images/id/{id}][%d] deleteImageUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteImageUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *DeleteImageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteImageForbidden creates a DeleteImageForbidden with default headers values
func NewDeleteImageForbidden() *DeleteImageForbidden {
	return &DeleteImageForbidden{}
}

/* DeleteImageForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type DeleteImageForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *DeleteImageForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/images/id/{id}][%d] deleteImageForbidden  %+v", 403, o.Payload)
}
func (o *DeleteImageForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *DeleteImageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteImageNotFound creates a DeleteImageNotFound with default headers values
func NewDeleteImageNotFound() *DeleteImageNotFound {
	return &DeleteImageNotFound{}
}

/* DeleteImageNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type DeleteImageNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *DeleteImageNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/images/id/{id}][%d] deleteImageNotFound  %+v", 404, o.Payload)
}
func (o *DeleteImageNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *DeleteImageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteImageConflict creates a DeleteImageConflict with default headers values
func NewDeleteImageConflict() *DeleteImageConflict {
	return &DeleteImageConflict{}
}

/* DeleteImageConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because there are edge application bundles using this edge application image
*/
type DeleteImageConflict struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *DeleteImageConflict) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/images/id/{id}][%d] deleteImageConflict  %+v", 409, o.Payload)
}
func (o *DeleteImageConflict) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *DeleteImageConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteImageInternalServerError creates a DeleteImageInternalServerError with default headers values
func NewDeleteImageInternalServerError() *DeleteImageInternalServerError {
	return &DeleteImageInternalServerError{}
}

/* DeleteImageInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type DeleteImageInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *DeleteImageInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/images/id/{id}][%d] deleteImageInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteImageInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *DeleteImageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteImageGatewayTimeout creates a DeleteImageGatewayTimeout with default headers values
func NewDeleteImageGatewayTimeout() *DeleteImageGatewayTimeout {
	return &DeleteImageGatewayTimeout{}
}

/* DeleteImageGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type DeleteImageGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *DeleteImageGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/images/id/{id}][%d] deleteImageGatewayTimeout  %+v", 504, o.Payload)
}
func (o *DeleteImageGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *DeleteImageGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

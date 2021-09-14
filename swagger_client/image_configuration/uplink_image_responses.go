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

// UplinkImageReader is a Reader for the UplinkImage structure.
type UplinkImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UplinkImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUplinkImageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewUplinkImageAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUplinkImageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUplinkImageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUplinkImageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUplinkImageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUplinkImageConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUplinkImageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewUplinkImageGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUplinkImageOK creates a UplinkImageOK with default headers values
func NewUplinkImageOK() *UplinkImageOK {
	return &UplinkImageOK{}
}

/* UplinkImageOK describes a response with status code 200, with default header values.

A successful response.
*/
type UplinkImageOK struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UplinkImageOK) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/uplink][%d] uplinkImageOK  %+v", 200, o.Payload)
}
func (o *UplinkImageOK) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UplinkImageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUplinkImageAccepted creates a UplinkImageAccepted with default headers values
func NewUplinkImageAccepted() *UplinkImageAccepted {
	return &UplinkImageAccepted{}
}

/* UplinkImageAccepted describes a response with status code 202, with default header values.

Accepted. The API gateway accepted the request for uplinking but the uplinking process has not been completed. Please check ImageStatus and ImageError fields to track the status of uplinking process and any error messages.
*/
type UplinkImageAccepted struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UplinkImageAccepted) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/uplink][%d] uplinkImageAccepted  %+v", 202, o.Payload)
}
func (o *UplinkImageAccepted) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UplinkImageAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUplinkImageBadRequest creates a UplinkImageBadRequest with default headers values
func NewUplinkImageBadRequest() *UplinkImageBadRequest {
	return &UplinkImageBadRequest{}
}

/* UplinkImageBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type UplinkImageBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UplinkImageBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/uplink][%d] uplinkImageBadRequest  %+v", 400, o.Payload)
}
func (o *UplinkImageBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UplinkImageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUplinkImageUnauthorized creates a UplinkImageUnauthorized with default headers values
func NewUplinkImageUnauthorized() *UplinkImageUnauthorized {
	return &UplinkImageUnauthorized{}
}

/* UplinkImageUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type UplinkImageUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UplinkImageUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/uplink][%d] uplinkImageUnauthorized  %+v", 401, o.Payload)
}
func (o *UplinkImageUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UplinkImageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUplinkImageForbidden creates a UplinkImageForbidden with default headers values
func NewUplinkImageForbidden() *UplinkImageForbidden {
	return &UplinkImageForbidden{}
}

/* UplinkImageForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type UplinkImageForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UplinkImageForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/uplink][%d] uplinkImageForbidden  %+v", 403, o.Payload)
}
func (o *UplinkImageForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UplinkImageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUplinkImageNotFound creates a UplinkImageNotFound with default headers values
func NewUplinkImageNotFound() *UplinkImageNotFound {
	return &UplinkImageNotFound{}
}

/* UplinkImageNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type UplinkImageNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UplinkImageNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/uplink][%d] uplinkImageNotFound  %+v", 404, o.Payload)
}
func (o *UplinkImageNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UplinkImageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUplinkImageConflict creates a UplinkImageConflict with default headers values
func NewUplinkImageConflict() *UplinkImageConflict {
	return &UplinkImageConflict{}
}

/* UplinkImageConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because another request for uplink / upload is already in progress
*/
type UplinkImageConflict struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UplinkImageConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/uplink][%d] uplinkImageConflict  %+v", 409, o.Payload)
}
func (o *UplinkImageConflict) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UplinkImageConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUplinkImageInternalServerError creates a UplinkImageInternalServerError with default headers values
func NewUplinkImageInternalServerError() *UplinkImageInternalServerError {
	return &UplinkImageInternalServerError{}
}

/* UplinkImageInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type UplinkImageInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UplinkImageInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/uplink][%d] uplinkImageInternalServerError  %+v", 500, o.Payload)
}
func (o *UplinkImageInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UplinkImageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUplinkImageGatewayTimeout creates a UplinkImageGatewayTimeout with default headers values
func NewUplinkImageGatewayTimeout() *UplinkImageGatewayTimeout {
	return &UplinkImageGatewayTimeout{}
}

/* UplinkImageGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type UplinkImageGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UplinkImageGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/uplink][%d] uplinkImageGatewayTimeout  %+v", 504, o.Payload)
}
func (o *UplinkImageGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UplinkImageGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

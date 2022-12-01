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

// ImageConfigurationGetImageReader is a Reader for the ImageConfigurationGetImage structure.
type ImageConfigurationGetImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImageConfigurationGetImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImageConfigurationGetImageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewImageConfigurationGetImageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewImageConfigurationGetImageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewImageConfigurationGetImageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImageConfigurationGetImageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewImageConfigurationGetImageGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewImageConfigurationGetImageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewImageConfigurationGetImageOK creates a ImageConfigurationGetImageOK with default headers values
func NewImageConfigurationGetImageOK() *ImageConfigurationGetImageOK {
	return &ImageConfigurationGetImageOK{}
}

/*
ImageConfigurationGetImageOK describes a response with status code 200, with default header values.

A successful response.
*/
type ImageConfigurationGetImageOK struct {
	Payload *swagger_models.ImageConfig
}

// IsSuccess returns true when this image configuration get image o k response has a 2xx status code
func (o *ImageConfigurationGetImageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this image configuration get image o k response has a 3xx status code
func (o *ImageConfigurationGetImageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration get image o k response has a 4xx status code
func (o *ImageConfigurationGetImageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this image configuration get image o k response has a 5xx status code
func (o *ImageConfigurationGetImageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration get image o k response a status code equal to that given
func (o *ImageConfigurationGetImageOK) IsCode(code int) bool {
	return code == 200
}

func (o *ImageConfigurationGetImageOK) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/id/{id}][%d] imageConfigurationGetImageOK  %+v", 200, o.Payload)
}

func (o *ImageConfigurationGetImageOK) String() string {
	return fmt.Sprintf("[GET /v1/apps/images/id/{id}][%d] imageConfigurationGetImageOK  %+v", 200, o.Payload)
}

func (o *ImageConfigurationGetImageOK) GetPayload() *swagger_models.ImageConfig {
	return o.Payload
}

func (o *ImageConfigurationGetImageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ImageConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationGetImageUnauthorized creates a ImageConfigurationGetImageUnauthorized with default headers values
func NewImageConfigurationGetImageUnauthorized() *ImageConfigurationGetImageUnauthorized {
	return &ImageConfigurationGetImageUnauthorized{}
}

/*
ImageConfigurationGetImageUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type ImageConfigurationGetImageUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this image configuration get image unauthorized response has a 2xx status code
func (o *ImageConfigurationGetImageUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration get image unauthorized response has a 3xx status code
func (o *ImageConfigurationGetImageUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration get image unauthorized response has a 4xx status code
func (o *ImageConfigurationGetImageUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration get image unauthorized response has a 5xx status code
func (o *ImageConfigurationGetImageUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration get image unauthorized response a status code equal to that given
func (o *ImageConfigurationGetImageUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ImageConfigurationGetImageUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/id/{id}][%d] imageConfigurationGetImageUnauthorized  %+v", 401, o.Payload)
}

func (o *ImageConfigurationGetImageUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/apps/images/id/{id}][%d] imageConfigurationGetImageUnauthorized  %+v", 401, o.Payload)
}

func (o *ImageConfigurationGetImageUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationGetImageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationGetImageForbidden creates a ImageConfigurationGetImageForbidden with default headers values
func NewImageConfigurationGetImageForbidden() *ImageConfigurationGetImageForbidden {
	return &ImageConfigurationGetImageForbidden{}
}

/*
ImageConfigurationGetImageForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type ImageConfigurationGetImageForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this image configuration get image forbidden response has a 2xx status code
func (o *ImageConfigurationGetImageForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration get image forbidden response has a 3xx status code
func (o *ImageConfigurationGetImageForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration get image forbidden response has a 4xx status code
func (o *ImageConfigurationGetImageForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration get image forbidden response has a 5xx status code
func (o *ImageConfigurationGetImageForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration get image forbidden response a status code equal to that given
func (o *ImageConfigurationGetImageForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ImageConfigurationGetImageForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/id/{id}][%d] imageConfigurationGetImageForbidden  %+v", 403, o.Payload)
}

func (o *ImageConfigurationGetImageForbidden) String() string {
	return fmt.Sprintf("[GET /v1/apps/images/id/{id}][%d] imageConfigurationGetImageForbidden  %+v", 403, o.Payload)
}

func (o *ImageConfigurationGetImageForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationGetImageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationGetImageNotFound creates a ImageConfigurationGetImageNotFound with default headers values
func NewImageConfigurationGetImageNotFound() *ImageConfigurationGetImageNotFound {
	return &ImageConfigurationGetImageNotFound{}
}

/*
ImageConfigurationGetImageNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type ImageConfigurationGetImageNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this image configuration get image not found response has a 2xx status code
func (o *ImageConfigurationGetImageNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration get image not found response has a 3xx status code
func (o *ImageConfigurationGetImageNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration get image not found response has a 4xx status code
func (o *ImageConfigurationGetImageNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration get image not found response has a 5xx status code
func (o *ImageConfigurationGetImageNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration get image not found response a status code equal to that given
func (o *ImageConfigurationGetImageNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ImageConfigurationGetImageNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/id/{id}][%d] imageConfigurationGetImageNotFound  %+v", 404, o.Payload)
}

func (o *ImageConfigurationGetImageNotFound) String() string {
	return fmt.Sprintf("[GET /v1/apps/images/id/{id}][%d] imageConfigurationGetImageNotFound  %+v", 404, o.Payload)
}

func (o *ImageConfigurationGetImageNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationGetImageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationGetImageInternalServerError creates a ImageConfigurationGetImageInternalServerError with default headers values
func NewImageConfigurationGetImageInternalServerError() *ImageConfigurationGetImageInternalServerError {
	return &ImageConfigurationGetImageInternalServerError{}
}

/*
ImageConfigurationGetImageInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type ImageConfigurationGetImageInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this image configuration get image internal server error response has a 2xx status code
func (o *ImageConfigurationGetImageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration get image internal server error response has a 3xx status code
func (o *ImageConfigurationGetImageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration get image internal server error response has a 4xx status code
func (o *ImageConfigurationGetImageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this image configuration get image internal server error response has a 5xx status code
func (o *ImageConfigurationGetImageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this image configuration get image internal server error response a status code equal to that given
func (o *ImageConfigurationGetImageInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ImageConfigurationGetImageInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/id/{id}][%d] imageConfigurationGetImageInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageConfigurationGetImageInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/apps/images/id/{id}][%d] imageConfigurationGetImageInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageConfigurationGetImageInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationGetImageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationGetImageGatewayTimeout creates a ImageConfigurationGetImageGatewayTimeout with default headers values
func NewImageConfigurationGetImageGatewayTimeout() *ImageConfigurationGetImageGatewayTimeout {
	return &ImageConfigurationGetImageGatewayTimeout{}
}

/*
ImageConfigurationGetImageGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type ImageConfigurationGetImageGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this image configuration get image gateway timeout response has a 2xx status code
func (o *ImageConfigurationGetImageGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration get image gateway timeout response has a 3xx status code
func (o *ImageConfigurationGetImageGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration get image gateway timeout response has a 4xx status code
func (o *ImageConfigurationGetImageGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this image configuration get image gateway timeout response has a 5xx status code
func (o *ImageConfigurationGetImageGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this image configuration get image gateway timeout response a status code equal to that given
func (o *ImageConfigurationGetImageGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *ImageConfigurationGetImageGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/id/{id}][%d] imageConfigurationGetImageGatewayTimeout  %+v", 504, o.Payload)
}

func (o *ImageConfigurationGetImageGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/apps/images/id/{id}][%d] imageConfigurationGetImageGatewayTimeout  %+v", 504, o.Payload)
}

func (o *ImageConfigurationGetImageGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationGetImageGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationGetImageDefault creates a ImageConfigurationGetImageDefault with default headers values
func NewImageConfigurationGetImageDefault(code int) *ImageConfigurationGetImageDefault {
	return &ImageConfigurationGetImageDefault{
		_statusCode: code,
	}
}

/*
ImageConfigurationGetImageDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ImageConfigurationGetImageDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the image configuration get image default response
func (o *ImageConfigurationGetImageDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this image configuration get image default response has a 2xx status code
func (o *ImageConfigurationGetImageDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this image configuration get image default response has a 3xx status code
func (o *ImageConfigurationGetImageDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this image configuration get image default response has a 4xx status code
func (o *ImageConfigurationGetImageDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this image configuration get image default response has a 5xx status code
func (o *ImageConfigurationGetImageDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this image configuration get image default response a status code equal to that given
func (o *ImageConfigurationGetImageDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ImageConfigurationGetImageDefault) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/id/{id}][%d] ImageConfiguration_GetImage default  %+v", o._statusCode, o.Payload)
}

func (o *ImageConfigurationGetImageDefault) String() string {
	return fmt.Sprintf("[GET /v1/apps/images/id/{id}][%d] ImageConfiguration_GetImage default  %+v", o._statusCode, o.Payload)
}

func (o *ImageConfigurationGetImageDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *ImageConfigurationGetImageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

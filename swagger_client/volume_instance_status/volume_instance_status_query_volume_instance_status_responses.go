// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package volume_instance_status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// VolumeInstanceStatusQueryVolumeInstanceStatusReader is a Reader for the VolumeInstanceStatusQueryVolumeInstanceStatus structure.
type VolumeInstanceStatusQueryVolumeInstanceStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVolumeInstanceStatusQueryVolumeInstanceStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewVolumeInstanceStatusQueryVolumeInstanceStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewVolumeInstanceStatusQueryVolumeInstanceStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewVolumeInstanceStatusQueryVolumeInstanceStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewVolumeInstanceStatusQueryVolumeInstanceStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewVolumeInstanceStatusQueryVolumeInstanceStatusGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewVolumeInstanceStatusQueryVolumeInstanceStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVolumeInstanceStatusQueryVolumeInstanceStatusOK creates a VolumeInstanceStatusQueryVolumeInstanceStatusOK with default headers values
func NewVolumeInstanceStatusQueryVolumeInstanceStatusOK() *VolumeInstanceStatusQueryVolumeInstanceStatusOK {
	return &VolumeInstanceStatusQueryVolumeInstanceStatusOK{}
}

/*
VolumeInstanceStatusQueryVolumeInstanceStatusOK describes a response with status code 200, with default header values.

A successful response.
*/
type VolumeInstanceStatusQueryVolumeInstanceStatusOK struct {
	Payload *swagger_models.VolInstStatusListMsg
}

// IsSuccess returns true when this volume instance status query volume instance status o k response has a 2xx status code
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this volume instance status query volume instance status o k response has a 3xx status code
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume instance status query volume instance status o k response has a 4xx status code
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this volume instance status query volume instance status o k response has a 5xx status code
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this volume instance status query volume instance status o k response a status code equal to that given
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusOK) IsCode(code int) bool {
	return code == 200
}

func (o *VolumeInstanceStatusQueryVolumeInstanceStatusOK) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/status][%d] volumeInstanceStatusQueryVolumeInstanceStatusOK  %+v", 200, o.Payload)
}

func (o *VolumeInstanceStatusQueryVolumeInstanceStatusOK) String() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/status][%d] volumeInstanceStatusQueryVolumeInstanceStatusOK  %+v", 200, o.Payload)
}

func (o *VolumeInstanceStatusQueryVolumeInstanceStatusOK) GetPayload() *swagger_models.VolInstStatusListMsg {
	return o.Payload
}

func (o *VolumeInstanceStatusQueryVolumeInstanceStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.VolInstStatusListMsg)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceStatusQueryVolumeInstanceStatusBadRequest creates a VolumeInstanceStatusQueryVolumeInstanceStatusBadRequest with default headers values
func NewVolumeInstanceStatusQueryVolumeInstanceStatusBadRequest() *VolumeInstanceStatusQueryVolumeInstanceStatusBadRequest {
	return &VolumeInstanceStatusQueryVolumeInstanceStatusBadRequest{}
}

/*
VolumeInstanceStatusQueryVolumeInstanceStatusBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of invalid value of filter parameters.
*/
type VolumeInstanceStatusQueryVolumeInstanceStatusBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this volume instance status query volume instance status bad request response has a 2xx status code
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this volume instance status query volume instance status bad request response has a 3xx status code
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume instance status query volume instance status bad request response has a 4xx status code
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this volume instance status query volume instance status bad request response has a 5xx status code
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this volume instance status query volume instance status bad request response a status code equal to that given
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *VolumeInstanceStatusQueryVolumeInstanceStatusBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/status][%d] volumeInstanceStatusQueryVolumeInstanceStatusBadRequest  %+v", 400, o.Payload)
}

func (o *VolumeInstanceStatusQueryVolumeInstanceStatusBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/status][%d] volumeInstanceStatusQueryVolumeInstanceStatusBadRequest  %+v", 400, o.Payload)
}

func (o *VolumeInstanceStatusQueryVolumeInstanceStatusBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceStatusQueryVolumeInstanceStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceStatusQueryVolumeInstanceStatusUnauthorized creates a VolumeInstanceStatusQueryVolumeInstanceStatusUnauthorized with default headers values
func NewVolumeInstanceStatusQueryVolumeInstanceStatusUnauthorized() *VolumeInstanceStatusQueryVolumeInstanceStatusUnauthorized {
	return &VolumeInstanceStatusQueryVolumeInstanceStatusUnauthorized{}
}

/*
VolumeInstanceStatusQueryVolumeInstanceStatusUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type VolumeInstanceStatusQueryVolumeInstanceStatusUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this volume instance status query volume instance status unauthorized response has a 2xx status code
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this volume instance status query volume instance status unauthorized response has a 3xx status code
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume instance status query volume instance status unauthorized response has a 4xx status code
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this volume instance status query volume instance status unauthorized response has a 5xx status code
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this volume instance status query volume instance status unauthorized response a status code equal to that given
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *VolumeInstanceStatusQueryVolumeInstanceStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/status][%d] volumeInstanceStatusQueryVolumeInstanceStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *VolumeInstanceStatusQueryVolumeInstanceStatusUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/status][%d] volumeInstanceStatusQueryVolumeInstanceStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *VolumeInstanceStatusQueryVolumeInstanceStatusUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceStatusQueryVolumeInstanceStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceStatusQueryVolumeInstanceStatusForbidden creates a VolumeInstanceStatusQueryVolumeInstanceStatusForbidden with default headers values
func NewVolumeInstanceStatusQueryVolumeInstanceStatusForbidden() *VolumeInstanceStatusQueryVolumeInstanceStatusForbidden {
	return &VolumeInstanceStatusQueryVolumeInstanceStatusForbidden{}
}

/*
VolumeInstanceStatusQueryVolumeInstanceStatusForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type VolumeInstanceStatusQueryVolumeInstanceStatusForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this volume instance status query volume instance status forbidden response has a 2xx status code
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this volume instance status query volume instance status forbidden response has a 3xx status code
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume instance status query volume instance status forbidden response has a 4xx status code
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this volume instance status query volume instance status forbidden response has a 5xx status code
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this volume instance status query volume instance status forbidden response a status code equal to that given
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *VolumeInstanceStatusQueryVolumeInstanceStatusForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/status][%d] volumeInstanceStatusQueryVolumeInstanceStatusForbidden  %+v", 403, o.Payload)
}

func (o *VolumeInstanceStatusQueryVolumeInstanceStatusForbidden) String() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/status][%d] volumeInstanceStatusQueryVolumeInstanceStatusForbidden  %+v", 403, o.Payload)
}

func (o *VolumeInstanceStatusQueryVolumeInstanceStatusForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceStatusQueryVolumeInstanceStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceStatusQueryVolumeInstanceStatusInternalServerError creates a VolumeInstanceStatusQueryVolumeInstanceStatusInternalServerError with default headers values
func NewVolumeInstanceStatusQueryVolumeInstanceStatusInternalServerError() *VolumeInstanceStatusQueryVolumeInstanceStatusInternalServerError {
	return &VolumeInstanceStatusQueryVolumeInstanceStatusInternalServerError{}
}

/*
VolumeInstanceStatusQueryVolumeInstanceStatusInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type VolumeInstanceStatusQueryVolumeInstanceStatusInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this volume instance status query volume instance status internal server error response has a 2xx status code
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this volume instance status query volume instance status internal server error response has a 3xx status code
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume instance status query volume instance status internal server error response has a 4xx status code
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this volume instance status query volume instance status internal server error response has a 5xx status code
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this volume instance status query volume instance status internal server error response a status code equal to that given
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *VolumeInstanceStatusQueryVolumeInstanceStatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/status][%d] volumeInstanceStatusQueryVolumeInstanceStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *VolumeInstanceStatusQueryVolumeInstanceStatusInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/status][%d] volumeInstanceStatusQueryVolumeInstanceStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *VolumeInstanceStatusQueryVolumeInstanceStatusInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceStatusQueryVolumeInstanceStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceStatusQueryVolumeInstanceStatusGatewayTimeout creates a VolumeInstanceStatusQueryVolumeInstanceStatusGatewayTimeout with default headers values
func NewVolumeInstanceStatusQueryVolumeInstanceStatusGatewayTimeout() *VolumeInstanceStatusQueryVolumeInstanceStatusGatewayTimeout {
	return &VolumeInstanceStatusQueryVolumeInstanceStatusGatewayTimeout{}
}

/*
VolumeInstanceStatusQueryVolumeInstanceStatusGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type VolumeInstanceStatusQueryVolumeInstanceStatusGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this volume instance status query volume instance status gateway timeout response has a 2xx status code
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this volume instance status query volume instance status gateway timeout response has a 3xx status code
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume instance status query volume instance status gateway timeout response has a 4xx status code
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this volume instance status query volume instance status gateway timeout response has a 5xx status code
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this volume instance status query volume instance status gateway timeout response a status code equal to that given
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *VolumeInstanceStatusQueryVolumeInstanceStatusGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/status][%d] volumeInstanceStatusQueryVolumeInstanceStatusGatewayTimeout  %+v", 504, o.Payload)
}

func (o *VolumeInstanceStatusQueryVolumeInstanceStatusGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/status][%d] volumeInstanceStatusQueryVolumeInstanceStatusGatewayTimeout  %+v", 504, o.Payload)
}

func (o *VolumeInstanceStatusQueryVolumeInstanceStatusGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceStatusQueryVolumeInstanceStatusGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceStatusQueryVolumeInstanceStatusDefault creates a VolumeInstanceStatusQueryVolumeInstanceStatusDefault with default headers values
func NewVolumeInstanceStatusQueryVolumeInstanceStatusDefault(code int) *VolumeInstanceStatusQueryVolumeInstanceStatusDefault {
	return &VolumeInstanceStatusQueryVolumeInstanceStatusDefault{
		_statusCode: code,
	}
}

/*
VolumeInstanceStatusQueryVolumeInstanceStatusDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type VolumeInstanceStatusQueryVolumeInstanceStatusDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the volume instance status query volume instance status default response
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this volume instance status query volume instance status default response has a 2xx status code
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this volume instance status query volume instance status default response has a 3xx status code
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this volume instance status query volume instance status default response has a 4xx status code
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this volume instance status query volume instance status default response has a 5xx status code
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this volume instance status query volume instance status default response a status code equal to that given
func (o *VolumeInstanceStatusQueryVolumeInstanceStatusDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *VolumeInstanceStatusQueryVolumeInstanceStatusDefault) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/status][%d] VolumeInstanceStatus_QueryVolumeInstanceStatus default  %+v", o._statusCode, o.Payload)
}

func (o *VolumeInstanceStatusQueryVolumeInstanceStatusDefault) String() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/status][%d] VolumeInstanceStatus_QueryVolumeInstanceStatus default  %+v", o._statusCode, o.Payload)
}

func (o *VolumeInstanceStatusQueryVolumeInstanceStatusDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *VolumeInstanceStatusQueryVolumeInstanceStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

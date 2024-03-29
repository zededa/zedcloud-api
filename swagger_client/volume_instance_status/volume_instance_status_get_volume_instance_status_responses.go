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

// VolumeInstanceStatusGetVolumeInstanceStatusReader is a Reader for the VolumeInstanceStatusGetVolumeInstanceStatus structure.
type VolumeInstanceStatusGetVolumeInstanceStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VolumeInstanceStatusGetVolumeInstanceStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVolumeInstanceStatusGetVolumeInstanceStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewVolumeInstanceStatusGetVolumeInstanceStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewVolumeInstanceStatusGetVolumeInstanceStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewVolumeInstanceStatusGetVolumeInstanceStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewVolumeInstanceStatusGetVolumeInstanceStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewVolumeInstanceStatusGetVolumeInstanceStatusGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewVolumeInstanceStatusGetVolumeInstanceStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVolumeInstanceStatusGetVolumeInstanceStatusOK creates a VolumeInstanceStatusGetVolumeInstanceStatusOK with default headers values
func NewVolumeInstanceStatusGetVolumeInstanceStatusOK() *VolumeInstanceStatusGetVolumeInstanceStatusOK {
	return &VolumeInstanceStatusGetVolumeInstanceStatusOK{}
}

/*
VolumeInstanceStatusGetVolumeInstanceStatusOK describes a response with status code 200, with default header values.

A successful response.
*/
type VolumeInstanceStatusGetVolumeInstanceStatusOK struct {
	Payload *swagger_models.VolInstStatusMsg
}

// IsSuccess returns true when this volume instance status get volume instance status o k response has a 2xx status code
func (o *VolumeInstanceStatusGetVolumeInstanceStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this volume instance status get volume instance status o k response has a 3xx status code
func (o *VolumeInstanceStatusGetVolumeInstanceStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume instance status get volume instance status o k response has a 4xx status code
func (o *VolumeInstanceStatusGetVolumeInstanceStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this volume instance status get volume instance status o k response has a 5xx status code
func (o *VolumeInstanceStatusGetVolumeInstanceStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this volume instance status get volume instance status o k response a status code equal to that given
func (o *VolumeInstanceStatusGetVolumeInstanceStatusOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the volume instance status get volume instance status o k response
func (o *VolumeInstanceStatusGetVolumeInstanceStatusOK) Code() int {
	return 200
}

func (o *VolumeInstanceStatusGetVolumeInstanceStatusOK) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/id/{id}/status][%d] volumeInstanceStatusGetVolumeInstanceStatusOK  %+v", 200, o.Payload)
}

func (o *VolumeInstanceStatusGetVolumeInstanceStatusOK) String() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/id/{id}/status][%d] volumeInstanceStatusGetVolumeInstanceStatusOK  %+v", 200, o.Payload)
}

func (o *VolumeInstanceStatusGetVolumeInstanceStatusOK) GetPayload() *swagger_models.VolInstStatusMsg {
	return o.Payload
}

func (o *VolumeInstanceStatusGetVolumeInstanceStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.VolInstStatusMsg)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceStatusGetVolumeInstanceStatusUnauthorized creates a VolumeInstanceStatusGetVolumeInstanceStatusUnauthorized with default headers values
func NewVolumeInstanceStatusGetVolumeInstanceStatusUnauthorized() *VolumeInstanceStatusGetVolumeInstanceStatusUnauthorized {
	return &VolumeInstanceStatusGetVolumeInstanceStatusUnauthorized{}
}

/*
VolumeInstanceStatusGetVolumeInstanceStatusUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type VolumeInstanceStatusGetVolumeInstanceStatusUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this volume instance status get volume instance status unauthorized response has a 2xx status code
func (o *VolumeInstanceStatusGetVolumeInstanceStatusUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this volume instance status get volume instance status unauthorized response has a 3xx status code
func (o *VolumeInstanceStatusGetVolumeInstanceStatusUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume instance status get volume instance status unauthorized response has a 4xx status code
func (o *VolumeInstanceStatusGetVolumeInstanceStatusUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this volume instance status get volume instance status unauthorized response has a 5xx status code
func (o *VolumeInstanceStatusGetVolumeInstanceStatusUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this volume instance status get volume instance status unauthorized response a status code equal to that given
func (o *VolumeInstanceStatusGetVolumeInstanceStatusUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the volume instance status get volume instance status unauthorized response
func (o *VolumeInstanceStatusGetVolumeInstanceStatusUnauthorized) Code() int {
	return 401
}

func (o *VolumeInstanceStatusGetVolumeInstanceStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/id/{id}/status][%d] volumeInstanceStatusGetVolumeInstanceStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *VolumeInstanceStatusGetVolumeInstanceStatusUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/id/{id}/status][%d] volumeInstanceStatusGetVolumeInstanceStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *VolumeInstanceStatusGetVolumeInstanceStatusUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceStatusGetVolumeInstanceStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceStatusGetVolumeInstanceStatusForbidden creates a VolumeInstanceStatusGetVolumeInstanceStatusForbidden with default headers values
func NewVolumeInstanceStatusGetVolumeInstanceStatusForbidden() *VolumeInstanceStatusGetVolumeInstanceStatusForbidden {
	return &VolumeInstanceStatusGetVolumeInstanceStatusForbidden{}
}

/*
VolumeInstanceStatusGetVolumeInstanceStatusForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type VolumeInstanceStatusGetVolumeInstanceStatusForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this volume instance status get volume instance status forbidden response has a 2xx status code
func (o *VolumeInstanceStatusGetVolumeInstanceStatusForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this volume instance status get volume instance status forbidden response has a 3xx status code
func (o *VolumeInstanceStatusGetVolumeInstanceStatusForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume instance status get volume instance status forbidden response has a 4xx status code
func (o *VolumeInstanceStatusGetVolumeInstanceStatusForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this volume instance status get volume instance status forbidden response has a 5xx status code
func (o *VolumeInstanceStatusGetVolumeInstanceStatusForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this volume instance status get volume instance status forbidden response a status code equal to that given
func (o *VolumeInstanceStatusGetVolumeInstanceStatusForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the volume instance status get volume instance status forbidden response
func (o *VolumeInstanceStatusGetVolumeInstanceStatusForbidden) Code() int {
	return 403
}

func (o *VolumeInstanceStatusGetVolumeInstanceStatusForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/id/{id}/status][%d] volumeInstanceStatusGetVolumeInstanceStatusForbidden  %+v", 403, o.Payload)
}

func (o *VolumeInstanceStatusGetVolumeInstanceStatusForbidden) String() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/id/{id}/status][%d] volumeInstanceStatusGetVolumeInstanceStatusForbidden  %+v", 403, o.Payload)
}

func (o *VolumeInstanceStatusGetVolumeInstanceStatusForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceStatusGetVolumeInstanceStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceStatusGetVolumeInstanceStatusNotFound creates a VolumeInstanceStatusGetVolumeInstanceStatusNotFound with default headers values
func NewVolumeInstanceStatusGetVolumeInstanceStatusNotFound() *VolumeInstanceStatusGetVolumeInstanceStatusNotFound {
	return &VolumeInstanceStatusGetVolumeInstanceStatusNotFound{}
}

/*
VolumeInstanceStatusGetVolumeInstanceStatusNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type VolumeInstanceStatusGetVolumeInstanceStatusNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this volume instance status get volume instance status not found response has a 2xx status code
func (o *VolumeInstanceStatusGetVolumeInstanceStatusNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this volume instance status get volume instance status not found response has a 3xx status code
func (o *VolumeInstanceStatusGetVolumeInstanceStatusNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume instance status get volume instance status not found response has a 4xx status code
func (o *VolumeInstanceStatusGetVolumeInstanceStatusNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this volume instance status get volume instance status not found response has a 5xx status code
func (o *VolumeInstanceStatusGetVolumeInstanceStatusNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this volume instance status get volume instance status not found response a status code equal to that given
func (o *VolumeInstanceStatusGetVolumeInstanceStatusNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the volume instance status get volume instance status not found response
func (o *VolumeInstanceStatusGetVolumeInstanceStatusNotFound) Code() int {
	return 404
}

func (o *VolumeInstanceStatusGetVolumeInstanceStatusNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/id/{id}/status][%d] volumeInstanceStatusGetVolumeInstanceStatusNotFound  %+v", 404, o.Payload)
}

func (o *VolumeInstanceStatusGetVolumeInstanceStatusNotFound) String() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/id/{id}/status][%d] volumeInstanceStatusGetVolumeInstanceStatusNotFound  %+v", 404, o.Payload)
}

func (o *VolumeInstanceStatusGetVolumeInstanceStatusNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceStatusGetVolumeInstanceStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceStatusGetVolumeInstanceStatusInternalServerError creates a VolumeInstanceStatusGetVolumeInstanceStatusInternalServerError with default headers values
func NewVolumeInstanceStatusGetVolumeInstanceStatusInternalServerError() *VolumeInstanceStatusGetVolumeInstanceStatusInternalServerError {
	return &VolumeInstanceStatusGetVolumeInstanceStatusInternalServerError{}
}

/*
VolumeInstanceStatusGetVolumeInstanceStatusInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type VolumeInstanceStatusGetVolumeInstanceStatusInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this volume instance status get volume instance status internal server error response has a 2xx status code
func (o *VolumeInstanceStatusGetVolumeInstanceStatusInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this volume instance status get volume instance status internal server error response has a 3xx status code
func (o *VolumeInstanceStatusGetVolumeInstanceStatusInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume instance status get volume instance status internal server error response has a 4xx status code
func (o *VolumeInstanceStatusGetVolumeInstanceStatusInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this volume instance status get volume instance status internal server error response has a 5xx status code
func (o *VolumeInstanceStatusGetVolumeInstanceStatusInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this volume instance status get volume instance status internal server error response a status code equal to that given
func (o *VolumeInstanceStatusGetVolumeInstanceStatusInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the volume instance status get volume instance status internal server error response
func (o *VolumeInstanceStatusGetVolumeInstanceStatusInternalServerError) Code() int {
	return 500
}

func (o *VolumeInstanceStatusGetVolumeInstanceStatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/id/{id}/status][%d] volumeInstanceStatusGetVolumeInstanceStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *VolumeInstanceStatusGetVolumeInstanceStatusInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/id/{id}/status][%d] volumeInstanceStatusGetVolumeInstanceStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *VolumeInstanceStatusGetVolumeInstanceStatusInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceStatusGetVolumeInstanceStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceStatusGetVolumeInstanceStatusGatewayTimeout creates a VolumeInstanceStatusGetVolumeInstanceStatusGatewayTimeout with default headers values
func NewVolumeInstanceStatusGetVolumeInstanceStatusGatewayTimeout() *VolumeInstanceStatusGetVolumeInstanceStatusGatewayTimeout {
	return &VolumeInstanceStatusGetVolumeInstanceStatusGatewayTimeout{}
}

/*
VolumeInstanceStatusGetVolumeInstanceStatusGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type VolumeInstanceStatusGetVolumeInstanceStatusGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this volume instance status get volume instance status gateway timeout response has a 2xx status code
func (o *VolumeInstanceStatusGetVolumeInstanceStatusGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this volume instance status get volume instance status gateway timeout response has a 3xx status code
func (o *VolumeInstanceStatusGetVolumeInstanceStatusGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume instance status get volume instance status gateway timeout response has a 4xx status code
func (o *VolumeInstanceStatusGetVolumeInstanceStatusGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this volume instance status get volume instance status gateway timeout response has a 5xx status code
func (o *VolumeInstanceStatusGetVolumeInstanceStatusGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this volume instance status get volume instance status gateway timeout response a status code equal to that given
func (o *VolumeInstanceStatusGetVolumeInstanceStatusGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the volume instance status get volume instance status gateway timeout response
func (o *VolumeInstanceStatusGetVolumeInstanceStatusGatewayTimeout) Code() int {
	return 504
}

func (o *VolumeInstanceStatusGetVolumeInstanceStatusGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/id/{id}/status][%d] volumeInstanceStatusGetVolumeInstanceStatusGatewayTimeout  %+v", 504, o.Payload)
}

func (o *VolumeInstanceStatusGetVolumeInstanceStatusGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/id/{id}/status][%d] volumeInstanceStatusGetVolumeInstanceStatusGatewayTimeout  %+v", 504, o.Payload)
}

func (o *VolumeInstanceStatusGetVolumeInstanceStatusGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceStatusGetVolumeInstanceStatusGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceStatusGetVolumeInstanceStatusDefault creates a VolumeInstanceStatusGetVolumeInstanceStatusDefault with default headers values
func NewVolumeInstanceStatusGetVolumeInstanceStatusDefault(code int) *VolumeInstanceStatusGetVolumeInstanceStatusDefault {
	return &VolumeInstanceStatusGetVolumeInstanceStatusDefault{
		_statusCode: code,
	}
}

/*
VolumeInstanceStatusGetVolumeInstanceStatusDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type VolumeInstanceStatusGetVolumeInstanceStatusDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// IsSuccess returns true when this volume instance status get volume instance status default response has a 2xx status code
func (o *VolumeInstanceStatusGetVolumeInstanceStatusDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this volume instance status get volume instance status default response has a 3xx status code
func (o *VolumeInstanceStatusGetVolumeInstanceStatusDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this volume instance status get volume instance status default response has a 4xx status code
func (o *VolumeInstanceStatusGetVolumeInstanceStatusDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this volume instance status get volume instance status default response has a 5xx status code
func (o *VolumeInstanceStatusGetVolumeInstanceStatusDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this volume instance status get volume instance status default response a status code equal to that given
func (o *VolumeInstanceStatusGetVolumeInstanceStatusDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the volume instance status get volume instance status default response
func (o *VolumeInstanceStatusGetVolumeInstanceStatusDefault) Code() int {
	return o._statusCode
}

func (o *VolumeInstanceStatusGetVolumeInstanceStatusDefault) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/id/{id}/status][%d] VolumeInstanceStatus_GetVolumeInstanceStatus default  %+v", o._statusCode, o.Payload)
}

func (o *VolumeInstanceStatusGetVolumeInstanceStatusDefault) String() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/id/{id}/status][%d] VolumeInstanceStatus_GetVolumeInstanceStatus default  %+v", o._statusCode, o.Payload)
}

func (o *VolumeInstanceStatusGetVolumeInstanceStatusDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *VolumeInstanceStatusGetVolumeInstanceStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

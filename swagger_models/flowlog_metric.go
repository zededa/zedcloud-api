// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package swagger_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// FlowlogMetric flowlog metric
//
// swagger:model FlowlogMetric
type FlowlogMetric string

func NewFlowlogMetric(value FlowlogMetric) *FlowlogMetric {
	v := value
	return &v
}

const (

	// FlowlogMetricFLOWLOGMETRICUNSPECIFIED captures enum value "FLOW_LOG_METRIC_UNSPECIFIED"
	FlowlogMetricFLOWLOGMETRICUNSPECIFIED FlowlogMetric = "FLOW_LOG_METRIC_UNSPECIFIED"

	// FlowlogMetricFLOWLOGMETRICBYTES captures enum value "FLOW_LOG_METRIC_BYTES"
	FlowlogMetricFLOWLOGMETRICBYTES FlowlogMetric = "FLOW_LOG_METRIC_BYTES"

	// FlowlogMetricFLOWLOGMETRICPACKETS captures enum value "FLOW_LOG_METRIC_PACKETS"
	FlowlogMetricFLOWLOGMETRICPACKETS FlowlogMetric = "FLOW_LOG_METRIC_PACKETS"
)

// for schema
var flowlogMetricEnum []interface{}

func init() {
	var res []FlowlogMetric
	if err := json.Unmarshal([]byte(`["FLOW_LOG_METRIC_UNSPECIFIED","FLOW_LOG_METRIC_BYTES","FLOW_LOG_METRIC_PACKETS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		flowlogMetricEnum = append(flowlogMetricEnum, v)
	}
}

func (m FlowlogMetric) validateFlowlogMetricEnum(path, location string, value FlowlogMetric) error {
	if err := validate.EnumCase(path, location, value, flowlogMetricEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this flowlog metric
func (m FlowlogMetric) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateFlowlogMetricEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this flowlog metric based on context it is used
func (m FlowlogMetric) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
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

// RunState Run-time state of an object, reported by EVE
//
// - RUN_STATE_ONLINE: Entity Online
//   - RUN_STATE_HALTED: Entity Halted
//   - RUN_STATE_INIT: Entity Initializing
//   - RUN_STATE_REBOOTING: Entity Rebooting
//   - RUN_STATE_OFFLINE: Entity Offline
//   - RUN_STATE_UNKNOWN: Entity state Unknown
//   - RUN_STATE_UNPROVISIONED: Entity Unprovisioned
//   - RUN_STATE_PROVISIONED: Entity Provisioned
//   - RUN_STATE_SUSPECT: Entity Suspect
//   - RUN_STATE_DOWNLOADING: Edge-node downloading entity artifacts
//   - RUN_STATE_RESTARTING: Entity Restarting
//   - RUN_STATE_PURGING: Entity Purging
//   - RUN_STATE_HALTING: Entity Halting
//   - RUN_STATE_ERROR: Entity encountered an error
//   - RUN_STATE_VERIFYING: Verification of downloaded Artifacts in Progress.
//   - RUN_STATE_LOADING: Loading of Artifacts into local datastore in Progress.
//   - RUN_STATE_CREATING_VOLUME: Volume creation from artifacts in Progress
//   - RUN_STATE_BOOTING: Entity booting up
//   - RUN_STATE_MAINTENANCE_MODE: Entity maintenance mode
//   - RUN_STATE_START_DELAYED: Application start delayed as per configuration
//   - RUN_STATE_BASEOS_UPDATING: Device BaseOs Update in Progress
//   - RUN_STATE_PREPARING_POWEROFF: device preparing power off - shutting down all app instances
//   - RUN_STATE_POWERING_OFF: device powering off from local profile server
//   - RUN_STATE_PREPARED_POWEROFF: device prepared power off - all app instances are shut down
//
// swagger:model RunState
type RunState string

func NewRunState(value RunState) *RunState {
	return &value
}

// Pointer returns a pointer to a freshly-allocated RunState.
func (m RunState) Pointer() *RunState {
	return &m
}

const (

	// RunStateRUNSTATEUNSPECIFIED captures enum value "RUN_STATE_UNSPECIFIED"
	RunStateRUNSTATEUNSPECIFIED RunState = "RUN_STATE_UNSPECIFIED"

	// RunStateRUNSTATEONLINE captures enum value "RUN_STATE_ONLINE"
	RunStateRUNSTATEONLINE RunState = "RUN_STATE_ONLINE"

	// RunStateRUNSTATEHALTED captures enum value "RUN_STATE_HALTED"
	RunStateRUNSTATEHALTED RunState = "RUN_STATE_HALTED"

	// RunStateRUNSTATEINIT captures enum value "RUN_STATE_INIT"
	RunStateRUNSTATEINIT RunState = "RUN_STATE_INIT"

	// RunStateRUNSTATEREBOOTING captures enum value "RUN_STATE_REBOOTING"
	RunStateRUNSTATEREBOOTING RunState = "RUN_STATE_REBOOTING"

	// RunStateRUNSTATEOFFLINE captures enum value "RUN_STATE_OFFLINE"
	RunStateRUNSTATEOFFLINE RunState = "RUN_STATE_OFFLINE"

	// RunStateRUNSTATEUNKNOWN captures enum value "RUN_STATE_UNKNOWN"
	RunStateRUNSTATEUNKNOWN RunState = "RUN_STATE_UNKNOWN"

	// RunStateRUNSTATEUNPROVISIONED captures enum value "RUN_STATE_UNPROVISIONED"
	RunStateRUNSTATEUNPROVISIONED RunState = "RUN_STATE_UNPROVISIONED"

	// RunStateRUNSTATEPROVISIONED captures enum value "RUN_STATE_PROVISIONED"
	RunStateRUNSTATEPROVISIONED RunState = "RUN_STATE_PROVISIONED"

	// RunStateRUNSTATESUSPECT captures enum value "RUN_STATE_SUSPECT"
	RunStateRUNSTATESUSPECT RunState = "RUN_STATE_SUSPECT"

	// RunStateRUNSTATEDOWNLOADING captures enum value "RUN_STATE_DOWNLOADING"
	RunStateRUNSTATEDOWNLOADING RunState = "RUN_STATE_DOWNLOADING"

	// RunStateRUNSTATERESTARTING captures enum value "RUN_STATE_RESTARTING"
	RunStateRUNSTATERESTARTING RunState = "RUN_STATE_RESTARTING"

	// RunStateRUNSTATEPURGING captures enum value "RUN_STATE_PURGING"
	RunStateRUNSTATEPURGING RunState = "RUN_STATE_PURGING"

	// RunStateRUNSTATEHALTING captures enum value "RUN_STATE_HALTING"
	RunStateRUNSTATEHALTING RunState = "RUN_STATE_HALTING"

	// RunStateRUNSTATEERROR captures enum value "RUN_STATE_ERROR"
	RunStateRUNSTATEERROR RunState = "RUN_STATE_ERROR"

	// RunStateRUNSTATEVERIFYING captures enum value "RUN_STATE_VERIFYING"
	RunStateRUNSTATEVERIFYING RunState = "RUN_STATE_VERIFYING"

	// RunStateRUNSTATELOADING captures enum value "RUN_STATE_LOADING"
	RunStateRUNSTATELOADING RunState = "RUN_STATE_LOADING"

	// RunStateRUNSTATECREATINGVOLUME captures enum value "RUN_STATE_CREATING_VOLUME"
	RunStateRUNSTATECREATINGVOLUME RunState = "RUN_STATE_CREATING_VOLUME"

	// RunStateRUNSTATEBOOTING captures enum value "RUN_STATE_BOOTING"
	RunStateRUNSTATEBOOTING RunState = "RUN_STATE_BOOTING"

	// RunStateRUNSTATEMAINTENANCEMODE captures enum value "RUN_STATE_MAINTENANCE_MODE"
	RunStateRUNSTATEMAINTENANCEMODE RunState = "RUN_STATE_MAINTENANCE_MODE"

	// RunStateRUNSTATESTARTDELAYED captures enum value "RUN_STATE_START_DELAYED"
	RunStateRUNSTATESTARTDELAYED RunState = "RUN_STATE_START_DELAYED"

	// RunStateRUNSTATEBASEOSUPDATING captures enum value "RUN_STATE_BASEOS_UPDATING"
	RunStateRUNSTATEBASEOSUPDATING RunState = "RUN_STATE_BASEOS_UPDATING"

	// RunStateRUNSTATEPREPARINGPOWEROFF captures enum value "RUN_STATE_PREPARING_POWEROFF"
	RunStateRUNSTATEPREPARINGPOWEROFF RunState = "RUN_STATE_PREPARING_POWEROFF"

	// RunStateRUNSTATEPOWERINGOFF captures enum value "RUN_STATE_POWERING_OFF"
	RunStateRUNSTATEPOWERINGOFF RunState = "RUN_STATE_POWERING_OFF"

	// RunStateRUNSTATEPREPAREDPOWEROFF captures enum value "RUN_STATE_PREPARED_POWEROFF"
	RunStateRUNSTATEPREPAREDPOWEROFF RunState = "RUN_STATE_PREPARED_POWEROFF"
)

// for schema
var runStateEnum []interface{}

func init() {
	var res []RunState
	if err := json.Unmarshal([]byte(`["RUN_STATE_UNSPECIFIED","RUN_STATE_ONLINE","RUN_STATE_HALTED","RUN_STATE_INIT","RUN_STATE_REBOOTING","RUN_STATE_OFFLINE","RUN_STATE_UNKNOWN","RUN_STATE_UNPROVISIONED","RUN_STATE_PROVISIONED","RUN_STATE_SUSPECT","RUN_STATE_DOWNLOADING","RUN_STATE_RESTARTING","RUN_STATE_PURGING","RUN_STATE_HALTING","RUN_STATE_ERROR","RUN_STATE_VERIFYING","RUN_STATE_LOADING","RUN_STATE_CREATING_VOLUME","RUN_STATE_BOOTING","RUN_STATE_MAINTENANCE_MODE","RUN_STATE_START_DELAYED","RUN_STATE_BASEOS_UPDATING","RUN_STATE_PREPARING_POWEROFF","RUN_STATE_POWERING_OFF","RUN_STATE_PREPARED_POWEROFF"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		runStateEnum = append(runStateEnum, v)
	}
}

func (m RunState) validateRunStateEnum(path, location string, value RunState) error {
	if err := validate.EnumCase(path, location, value, runStateEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this run state
func (m RunState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateRunStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this run state based on context it is used
func (m RunState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

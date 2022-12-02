// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package swagger_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DeviceInfoMsg device info msg
//
// swagger:model DeviceInfoMsg
type DeviceInfoMsg struct {

	// Cpu
	CPU *CPUSummary `json:"Cpu,omitempty"`

	// memory - OBSOLETE. Use memorySummary instead.
	Memory *MemorySummary `json:"Memory,omitempty"`

	// storage
	Storage *StorageSummary `json:"Storage,omitempty"`

	// admin state
	AdminState *AdminState `json:"adminState,omitempty"`

	// attest state
	AttestState *AttestState `json:"attestState,omitempty"`

	// blob list
	BlobList []*BlobStatus `json:"blobList"`

	// boot time
	// Format: date-time
	BootTime strfmt.DateTime `json:"bootTime,omitempty"`

	// Information about hardware capabilities
	//
	// Edge node virtualization capabilities.
	Capabilities *Capabilities `json:"capabilities,omitempty"`

	// System defined universally unique clusterInstance ID, unique across the enterprise.
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9][a-zA-Z0-9_.-]+
	ClusterID string `json:"clusterID,omitempty"`

	// data sec info
	DataSecInfo []*DevDataSecAtRest `json:"dataSecInfo"`

	// debug knob expiry time
	DebugKnob bool `json:"debugKnob,omitempty"`

	// debug knob expiry time
	// Format: date-time
	DebugKnobExpiryTime strfmt.DateTime `json:"debugKnobExpiryTime,omitempty"`

	// dev error
	DevError []*DeviceError `json:"devError"`

	// device reboot reason
	DeviceRebootReason *DeviceBootReason `json:"deviceRebootReason,omitempty"`

	// dinfo
	Dinfo *DeviceInfo `json:"dinfo,omitempty"`

	// dns
	DNS *DNSInfo `json:"dns,omitempty"`

	// host name
	HostName string `json:"hostName,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// io status list
	IoStatusList []*IoBundleStatus `json:"ioStatusList"`

	// last reboot reason
	LastRebootReason string `json:"lastRebootReason,omitempty"`

	// last reboot time
	// Format: date-time
	LastRebootTime strfmt.DateTime `json:"lastRebootTime,omitempty"`

	// last update
	// Format: date-time
	LastUpdate strfmt.DateTime `json:"lastUpdate,omitempty"`

	// Device memory Info
	MemorySummary *DeviceMemorySummary `json:"memorySummary,omitempty"`

	// minfo
	Minfo *ZManufacturerInfo `json:"minfo,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// net counter list
	NetCounterList []*NetworkCounters `json:"netCounterList"`

	// net status list
	NetStatusList []*NetworkStatus `json:"netStatusList"`

	// physical storage
	PhysicalStorage []*PhysicalStorage `json:"physicalStorage"`

	// project Id
	ProjectID string `json:"projectId,omitempty"`

	// raw status
	RawStatus string `json:"rawStatus,omitempty"`

	// run state
	RunState *RunState `json:"runState,omitempty"`

	// storage list
	StorageList []*StorageStatus `json:"storageList"`

	// sw info
	SwInfo []*DeviceSWInfo `json:"swInfo"`

	// deprecated = 6;
	//
	// Tags are name/value pairs that enable you to categorize resources. Tag names are case insensitive with max_length 512 and min_length 3. Tag values are case sensitive with max_length 256 and min_length 3.
	Tags map[string]string `json:"tags,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// up time
	// Format: date-time
	UpTime strfmt.DateTime `json:"upTime,omitempty"`

	// zc counters
	ZcCounters []*ZedcloudCounters `json:"zcCounters"`

	// Last received counters for zpool metrics.
	ZpoolMetrics *StorageDeviceMetrics `json:"zpoolMetrics,omitempty"`
}

// Validate validates this device info msg
func (m *DeviceInfoMsg) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCPU(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAdminState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAttestState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlobList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBootTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCapabilities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataSecInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDebugKnobExpiryTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceRebootReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDinfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDNS(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIoStatusList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastRebootTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemorySummary(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetCounterList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetStatusList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhysicalStorage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSwInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZcCounters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZpoolMetrics(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceInfoMsg) validateCPU(formats strfmt.Registry) error {
	if swag.IsZero(m.CPU) { // not required
		return nil
	}

	if m.CPU != nil {
		if err := m.CPU.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Cpu")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Cpu")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInfoMsg) validateMemory(formats strfmt.Registry) error {
	if swag.IsZero(m.Memory) { // not required
		return nil
	}

	if m.Memory != nil {
		if err := m.Memory.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Memory")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Memory")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInfoMsg) validateStorage(formats strfmt.Registry) error {
	if swag.IsZero(m.Storage) { // not required
		return nil
	}

	if m.Storage != nil {
		if err := m.Storage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Storage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Storage")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInfoMsg) validateAdminState(formats strfmt.Registry) error {
	if swag.IsZero(m.AdminState) { // not required
		return nil
	}

	if m.AdminState != nil {
		if err := m.AdminState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("adminState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("adminState")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInfoMsg) validateAttestState(formats strfmt.Registry) error {
	if swag.IsZero(m.AttestState) { // not required
		return nil
	}

	if m.AttestState != nil {
		if err := m.AttestState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attestState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attestState")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInfoMsg) validateBlobList(formats strfmt.Registry) error {
	if swag.IsZero(m.BlobList) { // not required
		return nil
	}

	for i := 0; i < len(m.BlobList); i++ {
		if swag.IsZero(m.BlobList[i]) { // not required
			continue
		}

		if m.BlobList[i] != nil {
			if err := m.BlobList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("blobList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("blobList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeviceInfoMsg) validateBootTime(formats strfmt.Registry) error {
	if swag.IsZero(m.BootTime) { // not required
		return nil
	}

	if err := validate.FormatOf("bootTime", "body", "date-time", m.BootTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeviceInfoMsg) validateCapabilities(formats strfmt.Registry) error {
	if swag.IsZero(m.Capabilities) { // not required
		return nil
	}

	if m.Capabilities != nil {
		if err := m.Capabilities.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("capabilities")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("capabilities")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInfoMsg) validateClusterID(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterID) { // not required
		return nil
	}

	if err := validate.MinLength("clusterID", "body", m.ClusterID, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("clusterID", "body", m.ClusterID, 256); err != nil {
		return err
	}

	if err := validate.Pattern("clusterID", "body", m.ClusterID, `[a-zA-Z0-9][a-zA-Z0-9_.-]+`); err != nil {
		return err
	}

	return nil
}

func (m *DeviceInfoMsg) validateDataSecInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.DataSecInfo) { // not required
		return nil
	}

	for i := 0; i < len(m.DataSecInfo); i++ {
		if swag.IsZero(m.DataSecInfo[i]) { // not required
			continue
		}

		if m.DataSecInfo[i] != nil {
			if err := m.DataSecInfo[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dataSecInfo" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dataSecInfo" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeviceInfoMsg) validateDebugKnobExpiryTime(formats strfmt.Registry) error {
	if swag.IsZero(m.DebugKnobExpiryTime) { // not required
		return nil
	}

	if err := validate.FormatOf("debugKnobExpiryTime", "body", "date-time", m.DebugKnobExpiryTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeviceInfoMsg) validateDevError(formats strfmt.Registry) error {
	if swag.IsZero(m.DevError) { // not required
		return nil
	}

	for i := 0; i < len(m.DevError); i++ {
		if swag.IsZero(m.DevError[i]) { // not required
			continue
		}

		if m.DevError[i] != nil {
			if err := m.DevError[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("devError" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("devError" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeviceInfoMsg) validateDeviceRebootReason(formats strfmt.Registry) error {
	if swag.IsZero(m.DeviceRebootReason) { // not required
		return nil
	}

	if m.DeviceRebootReason != nil {
		if err := m.DeviceRebootReason.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deviceRebootReason")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deviceRebootReason")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInfoMsg) validateDinfo(formats strfmt.Registry) error {
	if swag.IsZero(m.Dinfo) { // not required
		return nil
	}

	if m.Dinfo != nil {
		if err := m.Dinfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dinfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dinfo")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInfoMsg) validateDNS(formats strfmt.Registry) error {
	if swag.IsZero(m.DNS) { // not required
		return nil
	}

	if m.DNS != nil {
		if err := m.DNS.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dns")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dns")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInfoMsg) validateIoStatusList(formats strfmt.Registry) error {
	if swag.IsZero(m.IoStatusList) { // not required
		return nil
	}

	for i := 0; i < len(m.IoStatusList); i++ {
		if swag.IsZero(m.IoStatusList[i]) { // not required
			continue
		}

		if m.IoStatusList[i] != nil {
			if err := m.IoStatusList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ioStatusList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ioStatusList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeviceInfoMsg) validateLastRebootTime(formats strfmt.Registry) error {
	if swag.IsZero(m.LastRebootTime) { // not required
		return nil
	}

	if err := validate.FormatOf("lastRebootTime", "body", "date-time", m.LastRebootTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeviceInfoMsg) validateLastUpdate(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdate) { // not required
		return nil
	}

	if err := validate.FormatOf("lastUpdate", "body", "date-time", m.LastUpdate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeviceInfoMsg) validateMemorySummary(formats strfmt.Registry) error {
	if swag.IsZero(m.MemorySummary) { // not required
		return nil
	}

	if m.MemorySummary != nil {
		if err := m.MemorySummary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("memorySummary")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("memorySummary")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInfoMsg) validateMinfo(formats strfmt.Registry) error {
	if swag.IsZero(m.Minfo) { // not required
		return nil
	}

	if m.Minfo != nil {
		if err := m.Minfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("minfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("minfo")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInfoMsg) validateNetCounterList(formats strfmt.Registry) error {
	if swag.IsZero(m.NetCounterList) { // not required
		return nil
	}

	for i := 0; i < len(m.NetCounterList); i++ {
		if swag.IsZero(m.NetCounterList[i]) { // not required
			continue
		}

		if m.NetCounterList[i] != nil {
			if err := m.NetCounterList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("netCounterList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("netCounterList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeviceInfoMsg) validateNetStatusList(formats strfmt.Registry) error {
	if swag.IsZero(m.NetStatusList) { // not required
		return nil
	}

	for i := 0; i < len(m.NetStatusList); i++ {
		if swag.IsZero(m.NetStatusList[i]) { // not required
			continue
		}

		if m.NetStatusList[i] != nil {
			if err := m.NetStatusList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("netStatusList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("netStatusList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeviceInfoMsg) validatePhysicalStorage(formats strfmt.Registry) error {
	if swag.IsZero(m.PhysicalStorage) { // not required
		return nil
	}

	for i := 0; i < len(m.PhysicalStorage); i++ {
		if swag.IsZero(m.PhysicalStorage[i]) { // not required
			continue
		}

		if m.PhysicalStorage[i] != nil {
			if err := m.PhysicalStorage[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("physicalStorage" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("physicalStorage" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeviceInfoMsg) validateRunState(formats strfmt.Registry) error {
	if swag.IsZero(m.RunState) { // not required
		return nil
	}

	if m.RunState != nil {
		if err := m.RunState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("runState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("runState")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInfoMsg) validateStorageList(formats strfmt.Registry) error {
	if swag.IsZero(m.StorageList) { // not required
		return nil
	}

	for i := 0; i < len(m.StorageList); i++ {
		if swag.IsZero(m.StorageList[i]) { // not required
			continue
		}

		if m.StorageList[i] != nil {
			if err := m.StorageList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("storageList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("storageList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeviceInfoMsg) validateSwInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.SwInfo) { // not required
		return nil
	}

	for i := 0; i < len(m.SwInfo); i++ {
		if swag.IsZero(m.SwInfo[i]) { // not required
			continue
		}

		if m.SwInfo[i] != nil {
			if err := m.SwInfo[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("swInfo" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("swInfo" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeviceInfoMsg) validateUpTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpTime) { // not required
		return nil
	}

	if err := validate.FormatOf("upTime", "body", "date-time", m.UpTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeviceInfoMsg) validateZcCounters(formats strfmt.Registry) error {
	if swag.IsZero(m.ZcCounters) { // not required
		return nil
	}

	for i := 0; i < len(m.ZcCounters); i++ {
		if swag.IsZero(m.ZcCounters[i]) { // not required
			continue
		}

		if m.ZcCounters[i] != nil {
			if err := m.ZcCounters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("zcCounters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("zcCounters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeviceInfoMsg) validateZpoolMetrics(formats strfmt.Registry) error {
	if swag.IsZero(m.ZpoolMetrics) { // not required
		return nil
	}

	if m.ZpoolMetrics != nil {
		if err := m.ZpoolMetrics.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("zpoolMetrics")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("zpoolMetrics")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this device info msg based on the context it is used
func (m *DeviceInfoMsg) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCPU(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMemory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStorage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAdminState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAttestState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBlobList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCapabilities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDataSecInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDevError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeviceRebootReason(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDinfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDNS(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIoStatusList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMemorySummary(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMinfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetCounterList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetStatusList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePhysicalStorage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRunState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStorageList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSwInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateZcCounters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateZpoolMetrics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceInfoMsg) contextValidateCPU(ctx context.Context, formats strfmt.Registry) error {

	if m.CPU != nil {
		if err := m.CPU.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Cpu")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Cpu")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInfoMsg) contextValidateMemory(ctx context.Context, formats strfmt.Registry) error {

	if m.Memory != nil {
		if err := m.Memory.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Memory")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Memory")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInfoMsg) contextValidateStorage(ctx context.Context, formats strfmt.Registry) error {

	if m.Storage != nil {
		if err := m.Storage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Storage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Storage")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInfoMsg) contextValidateAdminState(ctx context.Context, formats strfmt.Registry) error {

	if m.AdminState != nil {
		if err := m.AdminState.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("adminState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("adminState")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInfoMsg) contextValidateAttestState(ctx context.Context, formats strfmt.Registry) error {

	if m.AttestState != nil {
		if err := m.AttestState.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attestState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attestState")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInfoMsg) contextValidateBlobList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BlobList); i++ {

		if m.BlobList[i] != nil {
			if err := m.BlobList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("blobList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("blobList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeviceInfoMsg) contextValidateCapabilities(ctx context.Context, formats strfmt.Registry) error {

	if m.Capabilities != nil {
		if err := m.Capabilities.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("capabilities")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("capabilities")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInfoMsg) contextValidateDataSecInfo(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DataSecInfo); i++ {

		if m.DataSecInfo[i] != nil {
			if err := m.DataSecInfo[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dataSecInfo" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dataSecInfo" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeviceInfoMsg) contextValidateDevError(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DevError); i++ {

		if m.DevError[i] != nil {
			if err := m.DevError[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("devError" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("devError" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeviceInfoMsg) contextValidateDeviceRebootReason(ctx context.Context, formats strfmt.Registry) error {

	if m.DeviceRebootReason != nil {
		if err := m.DeviceRebootReason.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deviceRebootReason")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deviceRebootReason")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInfoMsg) contextValidateDinfo(ctx context.Context, formats strfmt.Registry) error {

	if m.Dinfo != nil {
		if err := m.Dinfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dinfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dinfo")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInfoMsg) contextValidateDNS(ctx context.Context, formats strfmt.Registry) error {

	if m.DNS != nil {
		if err := m.DNS.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dns")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dns")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInfoMsg) contextValidateIoStatusList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IoStatusList); i++ {

		if m.IoStatusList[i] != nil {
			if err := m.IoStatusList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ioStatusList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ioStatusList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeviceInfoMsg) contextValidateMemorySummary(ctx context.Context, formats strfmt.Registry) error {

	if m.MemorySummary != nil {
		if err := m.MemorySummary.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("memorySummary")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("memorySummary")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInfoMsg) contextValidateMinfo(ctx context.Context, formats strfmt.Registry) error {

	if m.Minfo != nil {
		if err := m.Minfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("minfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("minfo")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInfoMsg) contextValidateNetCounterList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NetCounterList); i++ {

		if m.NetCounterList[i] != nil {
			if err := m.NetCounterList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("netCounterList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("netCounterList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeviceInfoMsg) contextValidateNetStatusList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NetStatusList); i++ {

		if m.NetStatusList[i] != nil {
			if err := m.NetStatusList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("netStatusList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("netStatusList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeviceInfoMsg) contextValidatePhysicalStorage(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PhysicalStorage); i++ {

		if m.PhysicalStorage[i] != nil {
			if err := m.PhysicalStorage[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("physicalStorage" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("physicalStorage" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeviceInfoMsg) contextValidateRunState(ctx context.Context, formats strfmt.Registry) error {

	if m.RunState != nil {
		if err := m.RunState.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("runState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("runState")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInfoMsg) contextValidateStorageList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StorageList); i++ {

		if m.StorageList[i] != nil {
			if err := m.StorageList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("storageList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("storageList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeviceInfoMsg) contextValidateSwInfo(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SwInfo); i++ {

		if m.SwInfo[i] != nil {
			if err := m.SwInfo[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("swInfo" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("swInfo" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeviceInfoMsg) contextValidateZcCounters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ZcCounters); i++ {

		if m.ZcCounters[i] != nil {
			if err := m.ZcCounters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("zcCounters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("zcCounters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeviceInfoMsg) contextValidateZpoolMetrics(ctx context.Context, formats strfmt.Registry) error {

	if m.ZpoolMetrics != nil {
		if err := m.ZpoolMetrics.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("zpoolMetrics")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("zpoolMetrics")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceInfoMsg) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceInfoMsg) UnmarshalBinary(b []byte) error {
	var res DeviceInfoMsg
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

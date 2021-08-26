/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"time"
)

// UsageSNMPHour The number of SNMP devices for each hour for a given organization.
type UsageSNMPHour struct {
	// The hour for the usage.
	Hour *time.Time `json:"hour,omitempty"`
	// Contains the number of SNMP devices.
	SnmpDevices *int64 `json:"snmp_devices,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:-`
}

// NewUsageSNMPHour instantiates a new UsageSNMPHour object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageSNMPHour() *UsageSNMPHour {
	this := UsageSNMPHour{}
	return &this
}

// NewUsageSNMPHourWithDefaults instantiates a new UsageSNMPHour object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageSNMPHourWithDefaults() *UsageSNMPHour {
	this := UsageSNMPHour{}
	return &this
}

// GetHour returns the Hour field value if set, zero value otherwise.
func (o *UsageSNMPHour) GetHour() time.Time {
	if o == nil || o.Hour == nil {
		var ret time.Time
		return ret
	}
	return *o.Hour
}

// GetHourOk returns a tuple with the Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSNMPHour) GetHourOk() (*time.Time, bool) {
	if o == nil || o.Hour == nil {
		return nil, false
	}
	return o.Hour, true
}

// HasHour returns a boolean if a field has been set.
func (o *UsageSNMPHour) HasHour() bool {
	if o != nil && o.Hour != nil {
		return true
	}

	return false
}

// SetHour gets a reference to the given time.Time and assigns it to the Hour field.
func (o *UsageSNMPHour) SetHour(v time.Time) {
	o.Hour = &v
}

// GetSnmpDevices returns the SnmpDevices field value if set, zero value otherwise.
func (o *UsageSNMPHour) GetSnmpDevices() int64 {
	if o == nil || o.SnmpDevices == nil {
		var ret int64
		return ret
	}
	return *o.SnmpDevices
}

// GetSnmpDevicesOk returns a tuple with the SnmpDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSNMPHour) GetSnmpDevicesOk() (*int64, bool) {
	if o == nil || o.SnmpDevices == nil {
		return nil, false
	}
	return o.SnmpDevices, true
}

// HasSnmpDevices returns a boolean if a field has been set.
func (o *UsageSNMPHour) HasSnmpDevices() bool {
	if o != nil && o.SnmpDevices != nil {
		return true
	}

	return false
}

// SetSnmpDevices gets a reference to the given int64 and assigns it to the SnmpDevices field.
func (o *UsageSNMPHour) SetSnmpDevices(v int64) {
	o.SnmpDevices = &v
}

func (o UsageSNMPHour) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Hour != nil {
		toSerialize["hour"] = o.Hour
	}
	if o.SnmpDevices != nil {
		toSerialize["snmp_devices"] = o.SnmpDevices
	}
	return json.Marshal(toSerialize)
}

func (o *UsageSNMPHour) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Hour        *time.Time `json:"hour,omitempty"`
		SnmpDevices *int64     `json:"snmp_devices,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.Hour = all.Hour
	o.SnmpDevices = all.SnmpDevices
	return nil
}

type NullableUsageSNMPHour struct {
	value *UsageSNMPHour
	isSet bool
}

func (v NullableUsageSNMPHour) Get() *UsageSNMPHour {
	return v.value
}

func (v *NullableUsageSNMPHour) Set(val *UsageSNMPHour) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageSNMPHour) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageSNMPHour) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageSNMPHour(val *UsageSNMPHour) *NullableUsageSNMPHour {
	return &NullableUsageSNMPHour{value: val, isSet: true}
}

func (v NullableUsageSNMPHour) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageSNMPHour) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// SecurityMonitoringSignalsListResponseLinks Links attributes.
type SecurityMonitoringSignalsListResponseLinks struct {
	// Link for the next set of results. Note that the request can also be made using the POST endpoint.
	Next *string `json:"next,omitempty"`
}

// NewSecurityMonitoringSignalsListResponseLinks instantiates a new SecurityMonitoringSignalsListResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityMonitoringSignalsListResponseLinks() *SecurityMonitoringSignalsListResponseLinks {
	this := SecurityMonitoringSignalsListResponseLinks{}
	return &this
}

// NewSecurityMonitoringSignalsListResponseLinksWithDefaults instantiates a new SecurityMonitoringSignalsListResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityMonitoringSignalsListResponseLinksWithDefaults() *SecurityMonitoringSignalsListResponseLinks {
	this := SecurityMonitoringSignalsListResponseLinks{}
	return &this
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *SecurityMonitoringSignalsListResponseLinks) GetNext() string {
	if o == nil || o.Next == nil {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalsListResponseLinks) GetNextOk() (*string, bool) {
	if o == nil || o.Next == nil {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *SecurityMonitoringSignalsListResponseLinks) HasNext() bool {
	if o != nil && o.Next != nil {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *SecurityMonitoringSignalsListResponseLinks) SetNext(v string) {
	o.Next = &v
}

func (o SecurityMonitoringSignalsListResponseLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Next != nil {
		toSerialize["next"] = o.Next
	}
	return json.Marshal(toSerialize)
}

type NullableSecurityMonitoringSignalsListResponseLinks struct {
	value *SecurityMonitoringSignalsListResponseLinks
	isSet bool
}

func (v NullableSecurityMonitoringSignalsListResponseLinks) Get() *SecurityMonitoringSignalsListResponseLinks {
	return v.value
}

func (v *NullableSecurityMonitoringSignalsListResponseLinks) Set(val *SecurityMonitoringSignalsListResponseLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityMonitoringSignalsListResponseLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityMonitoringSignalsListResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityMonitoringSignalsListResponseLinks(val *SecurityMonitoringSignalsListResponseLinks) *NullableSecurityMonitoringSignalsListResponseLinks {
	return &NullableSecurityMonitoringSignalsListResponseLinks{value: val, isSet: true}
}

func (v NullableSecurityMonitoringSignalsListResponseLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityMonitoringSignalsListResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
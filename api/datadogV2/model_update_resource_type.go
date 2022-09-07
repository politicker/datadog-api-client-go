// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// UpdateResourceType The JSON:API type for this request.
type UpdateResourceType string

// List of UpdateResourceType.
const (
	UPDATERESOURCETYPE_CONFLUENT_CLOUD_RESOURCES UpdateResourceType = "confluent-cloud-resources"
)

var allowedUpdateResourceTypeEnumValues = []UpdateResourceType{
	UPDATERESOURCETYPE_CONFLUENT_CLOUD_RESOURCES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *UpdateResourceType) GetAllowedValues() []UpdateResourceType {
	return allowedUpdateResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *UpdateResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UpdateResourceType(value)
	return nil
}

// NewUpdateResourceTypeFromValue returns a pointer to a valid UpdateResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewUpdateResourceTypeFromValue(v string) (*UpdateResourceType, error) {
	ev := UpdateResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for UpdateResourceType: valid values are %v", v, allowedUpdateResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v UpdateResourceType) IsValid() bool {
	for _, existing := range allowedUpdateResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UpdateResourceType value.
func (v UpdateResourceType) Ptr() *UpdateResourceType {
	return &v
}

// NullableUpdateResourceType handles when a null is used for UpdateResourceType.
type NullableUpdateResourceType struct {
	value *UpdateResourceType
	isSet bool
}

// Get returns the associated value.
func (v NullableUpdateResourceType) Get() *UpdateResourceType {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableUpdateResourceType) Set(val *UpdateResourceType) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableUpdateResourceType) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableUpdateResourceType) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableUpdateResourceType initializes the struct as if Set has been called.
func NewNullableUpdateResourceType(val *UpdateResourceType) *NullableUpdateResourceType {
	return &NullableUpdateResourceType{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableUpdateResourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableUpdateResourceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

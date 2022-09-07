// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// UpdateAccountType The JSON:API type for this request.
type UpdateAccountType string

// List of UpdateAccountType.
const (
	UPDATEACCOUNTTYPE_CONFLUENT_CLOUD_ACCOUNTS UpdateAccountType = "confluent-cloud-accounts"
)

var allowedUpdateAccountTypeEnumValues = []UpdateAccountType{
	UPDATEACCOUNTTYPE_CONFLUENT_CLOUD_ACCOUNTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *UpdateAccountType) GetAllowedValues() []UpdateAccountType {
	return allowedUpdateAccountTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *UpdateAccountType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UpdateAccountType(value)
	return nil
}

// NewUpdateAccountTypeFromValue returns a pointer to a valid UpdateAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewUpdateAccountTypeFromValue(v string) (*UpdateAccountType, error) {
	ev := UpdateAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for UpdateAccountType: valid values are %v", v, allowedUpdateAccountTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v UpdateAccountType) IsValid() bool {
	for _, existing := range allowedUpdateAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UpdateAccountType value.
func (v UpdateAccountType) Ptr() *UpdateAccountType {
	return &v
}

// NullableUpdateAccountType handles when a null is used for UpdateAccountType.
type NullableUpdateAccountType struct {
	value *UpdateAccountType
	isSet bool
}

// Get returns the associated value.
func (v NullableUpdateAccountType) Get() *UpdateAccountType {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableUpdateAccountType) Set(val *UpdateAccountType) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableUpdateAccountType) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableUpdateAccountType) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableUpdateAccountType initializes the struct as if Set has been called.
func NewNullableUpdateAccountType(val *UpdateAccountType) *NullableUpdateAccountType {
	return &NullableUpdateAccountType{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableUpdateAccountType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableUpdateAccountType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

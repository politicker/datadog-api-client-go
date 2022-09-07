// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// AddAccountType The JSON:API type. Will always be `confluent-cloud-accounts`.
type AddAccountType string

// List of AddAccountType.
const (
	ADDACCOUNTTYPE_CONFLUENT_CLOUD_ACCOUNTS AddAccountType = "confluent-cloud-accounts"
)

var allowedAddAccountTypeEnumValues = []AddAccountType{
	ADDACCOUNTTYPE_CONFLUENT_CLOUD_ACCOUNTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AddAccountType) GetAllowedValues() []AddAccountType {
	return allowedAddAccountTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AddAccountType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AddAccountType(value)
	return nil
}

// NewAddAccountTypeFromValue returns a pointer to a valid AddAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAddAccountTypeFromValue(v string) (*AddAccountType, error) {
	ev := AddAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AddAccountType: valid values are %v", v, allowedAddAccountTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AddAccountType) IsValid() bool {
	for _, existing := range allowedAddAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AddAccountType value.
func (v AddAccountType) Ptr() *AddAccountType {
	return &v
}

// NullableAddAccountType handles when a null is used for AddAccountType.
type NullableAddAccountType struct {
	value *AddAccountType
	isSet bool
}

// Get returns the associated value.
func (v NullableAddAccountType) Get() *AddAccountType {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableAddAccountType) Set(val *AddAccountType) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableAddAccountType) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableAddAccountType) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableAddAccountType initializes the struct as if Set has been called.
func NewNullableAddAccountType(val *AddAccountType) *NullableAddAccountType {
	return &NullableAddAccountType{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableAddAccountType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableAddAccountType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

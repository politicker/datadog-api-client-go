// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// AccountConfigType The JSON:API type for this API. Should always be `confluent-cloud-accounts`.
type AccountConfigType string

// List of AccountConfigType.
const (
	ACCOUNTCONFIGTYPE_CONFLUENT_CLOUD_ACCOUNTS AccountConfigType = "confluent-cloud-accounts"
)

var allowedAccountConfigTypeEnumValues = []AccountConfigType{
	ACCOUNTCONFIGTYPE_CONFLUENT_CLOUD_ACCOUNTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AccountConfigType) GetAllowedValues() []AccountConfigType {
	return allowedAccountConfigTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AccountConfigType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AccountConfigType(value)
	return nil
}

// NewAccountConfigTypeFromValue returns a pointer to a valid AccountConfigType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAccountConfigTypeFromValue(v string) (*AccountConfigType, error) {
	ev := AccountConfigType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AccountConfigType: valid values are %v", v, allowedAccountConfigTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AccountConfigType) IsValid() bool {
	for _, existing := range allowedAccountConfigTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccountConfigType value.
func (v AccountConfigType) Ptr() *AccountConfigType {
	return &v
}

// NullableAccountConfigType handles when a null is used for AccountConfigType.
type NullableAccountConfigType struct {
	value *AccountConfigType
	isSet bool
}

// Get returns the associated value.
func (v NullableAccountConfigType) Get() *AccountConfigType {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableAccountConfigType) Set(val *AccountConfigType) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableAccountConfigType) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableAccountConfigType) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableAccountConfigType initializes the struct as if Set has been called.
func NewNullableAccountConfigType(val *AccountConfigType) *NullableAccountConfigType {
	return &NullableAccountConfigType{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableAccountConfigType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableAccountConfigType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

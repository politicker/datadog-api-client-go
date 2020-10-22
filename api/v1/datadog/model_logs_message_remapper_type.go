/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// LogsMessageRemapperType Type of logs message remapper.
type LogsMessageRemapperType string

// List of LogsMessageRemapperType
const (
	LOGSMESSAGEREMAPPERTYPE_MESSAGE_REMAPPER LogsMessageRemapperType = "message-remapper"
)

func (v *LogsMessageRemapperType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LogsMessageRemapperType(value)
	for _, existing := range []LogsMessageRemapperType{"message-remapper"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LogsMessageRemapperType", value)
}

// Ptr returns reference to LogsMessageRemapperType value
func (v LogsMessageRemapperType) Ptr() *LogsMessageRemapperType {
	return &v
}

type NullableLogsMessageRemapperType struct {
	value *LogsMessageRemapperType
	isSet bool
}

func (v NullableLogsMessageRemapperType) Get() *LogsMessageRemapperType {
	return v.value
}

func (v *NullableLogsMessageRemapperType) Set(val *LogsMessageRemapperType) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsMessageRemapperType) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsMessageRemapperType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsMessageRemapperType(val *LogsMessageRemapperType) *NullableLogsMessageRemapperType {
	return &NullableLogsMessageRemapperType{value: val, isSet: true}
}

func (v NullableLogsMessageRemapperType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsMessageRemapperType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
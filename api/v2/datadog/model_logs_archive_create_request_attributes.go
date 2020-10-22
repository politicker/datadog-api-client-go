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

// LogsArchiveCreateRequestAttributes The attributes associated with the archive.
type LogsArchiveCreateRequestAttributes struct {
	Destination LogsArchiveCreateRequestDestination `json:"destination"`
	// To store the tags in the archive, set the value \"true\". If it is set to \"false\", the tags will be deleted when the logs are sent to the archive.
	IncludeTags *bool `json:"include_tags,omitempty"`
	// The archive name.
	Name string `json:"name"`
	// The archive query/filter. Logs matching this query are included in the archive.
	Query string `json:"query"`
	// An array of tags to add to rehydrated logs from an archive.
	RehydrationTags *[]string `json:"rehydration_tags,omitempty"`
}

// NewLogsArchiveCreateRequestAttributes instantiates a new LogsArchiveCreateRequestAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogsArchiveCreateRequestAttributes(destination LogsArchiveCreateRequestDestination, name string, query string) *LogsArchiveCreateRequestAttributes {
	this := LogsArchiveCreateRequestAttributes{}
	this.Destination = destination
	var includeTags bool = false
	this.IncludeTags = &includeTags
	this.Name = name
	this.Query = query
	return &this
}

// NewLogsArchiveCreateRequestAttributesWithDefaults instantiates a new LogsArchiveCreateRequestAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogsArchiveCreateRequestAttributesWithDefaults() *LogsArchiveCreateRequestAttributes {
	this := LogsArchiveCreateRequestAttributes{}
	var includeTags bool = false
	this.IncludeTags = &includeTags
	return &this
}

// GetDestination returns the Destination field value
func (o *LogsArchiveCreateRequestAttributes) GetDestination() LogsArchiveCreateRequestDestination {
	if o == nil {
		var ret LogsArchiveCreateRequestDestination
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *LogsArchiveCreateRequestAttributes) GetDestinationOk() (*LogsArchiveCreateRequestDestination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *LogsArchiveCreateRequestAttributes) SetDestination(v LogsArchiveCreateRequestDestination) {
	o.Destination = v
}

// GetIncludeTags returns the IncludeTags field value if set, zero value otherwise.
func (o *LogsArchiveCreateRequestAttributes) GetIncludeTags() bool {
	if o == nil || o.IncludeTags == nil {
		var ret bool
		return ret
	}
	return *o.IncludeTags
}

// GetIncludeTagsOk returns a tuple with the IncludeTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsArchiveCreateRequestAttributes) GetIncludeTagsOk() (*bool, bool) {
	if o == nil || o.IncludeTags == nil {
		return nil, false
	}
	return o.IncludeTags, true
}

// HasIncludeTags returns a boolean if a field has been set.
func (o *LogsArchiveCreateRequestAttributes) HasIncludeTags() bool {
	if o != nil && o.IncludeTags != nil {
		return true
	}

	return false
}

// SetIncludeTags gets a reference to the given bool and assigns it to the IncludeTags field.
func (o *LogsArchiveCreateRequestAttributes) SetIncludeTags(v bool) {
	o.IncludeTags = &v
}

// GetName returns the Name field value
func (o *LogsArchiveCreateRequestAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LogsArchiveCreateRequestAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LogsArchiveCreateRequestAttributes) SetName(v string) {
	o.Name = v
}

// GetQuery returns the Query field value
func (o *LogsArchiveCreateRequestAttributes) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *LogsArchiveCreateRequestAttributes) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *LogsArchiveCreateRequestAttributes) SetQuery(v string) {
	o.Query = v
}

// GetRehydrationTags returns the RehydrationTags field value if set, zero value otherwise.
func (o *LogsArchiveCreateRequestAttributes) GetRehydrationTags() []string {
	if o == nil || o.RehydrationTags == nil {
		var ret []string
		return ret
	}
	return *o.RehydrationTags
}

// GetRehydrationTagsOk returns a tuple with the RehydrationTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsArchiveCreateRequestAttributes) GetRehydrationTagsOk() (*[]string, bool) {
	if o == nil || o.RehydrationTags == nil {
		return nil, false
	}
	return o.RehydrationTags, true
}

// HasRehydrationTags returns a boolean if a field has been set.
func (o *LogsArchiveCreateRequestAttributes) HasRehydrationTags() bool {
	if o != nil && o.RehydrationTags != nil {
		return true
	}

	return false
}

// SetRehydrationTags gets a reference to the given []string and assigns it to the RehydrationTags field.
func (o *LogsArchiveCreateRequestAttributes) SetRehydrationTags(v []string) {
	o.RehydrationTags = &v
}

func (o LogsArchiveCreateRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["destination"] = o.Destination
	}
	if o.IncludeTags != nil {
		toSerialize["include_tags"] = o.IncludeTags
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["query"] = o.Query
	}
	if o.RehydrationTags != nil {
		toSerialize["rehydration_tags"] = o.RehydrationTags
	}
	return json.Marshal(toSerialize)
}

type NullableLogsArchiveCreateRequestAttributes struct {
	value *LogsArchiveCreateRequestAttributes
	isSet bool
}

func (v NullableLogsArchiveCreateRequestAttributes) Get() *LogsArchiveCreateRequestAttributes {
	return v.value
}

func (v *NullableLogsArchiveCreateRequestAttributes) Set(val *LogsArchiveCreateRequestAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsArchiveCreateRequestAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsArchiveCreateRequestAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsArchiveCreateRequestAttributes(val *LogsArchiveCreateRequestAttributes) *NullableLogsArchiveCreateRequestAttributes {
	return &NullableLogsArchiveCreateRequestAttributes{value: val, isSet: true}
}

func (v NullableLogsArchiveCreateRequestAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsArchiveCreateRequestAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
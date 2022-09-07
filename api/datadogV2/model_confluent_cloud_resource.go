// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// ConfluentCloudResource Model representation of a Confluent Cloud resource.
type ConfluentCloudResource struct {
	// The ID associated with the Confluent resource.
	Id string `json:"id"`
	// The resource type of the Resource. Can be `kafka`, `connector`, `ksql`, or `schema_registry`.
	ResourceType string `json:"resource_type"`
	// A comma-delimited string representation of tags. Can be a single key, or key-value pairs separated by a colon.
	Tags []string `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewConfluentCloudResource instantiates a new ConfluentCloudResource object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewConfluentCloudResource(id string, resourceType string) *ConfluentCloudResource {
	this := ConfluentCloudResource{}
	this.Id = id
	this.ResourceType = resourceType
	return &this
}

// NewConfluentCloudResourceWithDefaults instantiates a new ConfluentCloudResource object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewConfluentCloudResourceWithDefaults() *ConfluentCloudResource {
	this := ConfluentCloudResource{}
	return &this
}

// GetId returns the Id field value.
func (o *ConfluentCloudResource) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ConfluentCloudResource) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ConfluentCloudResource) SetId(v string) {
	o.Id = v
}

// GetResourceType returns the ResourceType field value.
func (o *ConfluentCloudResource) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *ConfluentCloudResource) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value.
func (o *ConfluentCloudResource) SetResourceType(v string) {
	o.ResourceType = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ConfluentCloudResource) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfluentCloudResource) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ConfluentCloudResource) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ConfluentCloudResource) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ConfluentCloudResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["resource_type"] = o.ResourceType
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ConfluentCloudResource) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		Id           *string `json:"id"`
		ResourceType *string `json:"resource_type"`
	}{}
	all := struct {
		Id           string   `json:"id"`
		ResourceType string   `json:"resource_type"`
		Tags         []string `json:"tags,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Id == nil {
		return fmt.Errorf("Required field id missing")
	}
	if required.ResourceType == nil {
		return fmt.Errorf("Required field resource_type missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.Id = all.Id
	o.ResourceType = all.ResourceType
	o.Tags = all.Tags
	return nil
}

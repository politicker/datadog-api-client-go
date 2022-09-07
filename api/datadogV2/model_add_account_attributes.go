// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// AddAccountAttributes Attributes associated with the account creation request.
type AddAccountAttributes struct {
	// The API key associated with your Confluent account.
	ApiKey string `json:"api_key"`
	// The API secret associated with your Confluent account.
	ApiSecret string `json:"api_secret"`
	// A list of Confluent resources associated with the Confluent account.
	Resources []ConfluentCloudResource `json:"resources,omitempty"`
	// A comma-delimited string representation of tags. Can be a single key, or key-value pairs separated by a colon.
	Tags []string `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewAddAccountAttributes instantiates a new AddAccountAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAddAccountAttributes(apiKey string, apiSecret string) *AddAccountAttributes {
	this := AddAccountAttributes{}
	this.ApiKey = apiKey
	this.ApiSecret = apiSecret
	return &this
}

// NewAddAccountAttributesWithDefaults instantiates a new AddAccountAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAddAccountAttributesWithDefaults() *AddAccountAttributes {
	this := AddAccountAttributes{}
	return &this
}

// GetApiKey returns the ApiKey field value.
func (o *AddAccountAttributes) GetApiKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *AddAccountAttributes) GetApiKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value.
func (o *AddAccountAttributes) SetApiKey(v string) {
	o.ApiKey = v
}

// GetApiSecret returns the ApiSecret field value.
func (o *AddAccountAttributes) GetApiSecret() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ApiSecret
}

// GetApiSecretOk returns a tuple with the ApiSecret field value
// and a boolean to check if the value has been set.
func (o *AddAccountAttributes) GetApiSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiSecret, true
}

// SetApiSecret sets field value.
func (o *AddAccountAttributes) SetApiSecret(v string) {
	o.ApiSecret = v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *AddAccountAttributes) GetResources() []ConfluentCloudResource {
	if o == nil || o.Resources == nil {
		var ret []ConfluentCloudResource
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAccountAttributes) GetResourcesOk() (*[]ConfluentCloudResource, bool) {
	if o == nil || o.Resources == nil {
		return nil, false
	}
	return &o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *AddAccountAttributes) HasResources() bool {
	if o != nil && o.Resources != nil {
		return true
	}

	return false
}

// SetResources gets a reference to the given []ConfluentCloudResource and assigns it to the Resources field.
func (o *AddAccountAttributes) SetResources(v []ConfluentCloudResource) {
	o.Resources = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *AddAccountAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAccountAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AddAccountAttributes) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *AddAccountAttributes) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AddAccountAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["api_key"] = o.ApiKey
	toSerialize["api_secret"] = o.ApiSecret
	if o.Resources != nil {
		toSerialize["resources"] = o.Resources
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AddAccountAttributes) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		ApiKey    *string `json:"api_key"`
		ApiSecret *string `json:"api_secret"`
	}{}
	all := struct {
		ApiKey    string                   `json:"api_key"`
		ApiSecret string                   `json:"api_secret"`
		Resources []ConfluentCloudResource `json:"resources,omitempty"`
		Tags      []string                 `json:"tags,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.ApiKey == nil {
		return fmt.Errorf("Required field api_key missing")
	}
	if required.ApiSecret == nil {
		return fmt.Errorf("Required field api_secret missing")
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
	o.ApiKey = all.ApiKey
	o.ApiSecret = all.ApiSecret
	o.Resources = all.Resources
	o.Tags = all.Tags
	return nil
}

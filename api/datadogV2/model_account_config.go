// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// AccountConfig An API key and API secret pair that represents a Confluent account.
type AccountConfig struct {
	// The attributes of a Confluent account.
	Attributes AccountConfigAttributes `json:"attributes"`
	// A randomly generated ID associated with a Confluent account.
	Id string `json:"id"`
	// The JSON:API type for this API. Should always be `confluent-cloud-accounts`.
	Type AccountConfigType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewAccountConfig instantiates a new AccountConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAccountConfig(attributes AccountConfigAttributes, id string, typeVar AccountConfigType) *AccountConfig {
	this := AccountConfig{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewAccountConfigWithDefaults instantiates a new AccountConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAccountConfigWithDefaults() *AccountConfig {
	this := AccountConfig{}
	var typeVar AccountConfigType = ACCOUNTCONFIGTYPE_CONFLUENT_CLOUD_ACCOUNTS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *AccountConfig) GetAttributes() AccountConfigAttributes {
	if o == nil {
		var ret AccountConfigAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AccountConfig) GetAttributesOk() (*AccountConfigAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *AccountConfig) SetAttributes(v AccountConfigAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *AccountConfig) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AccountConfig) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *AccountConfig) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *AccountConfig) GetType() AccountConfigType {
	if o == nil {
		var ret AccountConfigType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AccountConfig) GetTypeOk() (*AccountConfigType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AccountConfig) SetType(v AccountConfigType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AccountConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AccountConfig) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		Attributes *AccountConfigAttributes `json:"attributes"`
		Id         *string                  `json:"id"`
		Type       *AccountConfigType       `json:"type"`
	}{}
	all := struct {
		Attributes AccountConfigAttributes `json:"attributes"`
		Id         string                  `json:"id"`
		Type       AccountConfigType       `json:"type"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Attributes == nil {
		return fmt.Errorf("Required field attributes missing")
	}
	if required.Id == nil {
		return fmt.Errorf("Required field id missing")
	}
	if required.Type == nil {
		return fmt.Errorf("Required field type missing")
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
	if v := all.Type; !v.IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Attributes = all.Attributes
	o.Id = all.Id
	o.Type = all.Type
	return nil
}

/*
Copyright 2022 Keyfactor
Licensed under the Apache License, Version 2.0 (the "License"); you may
not use this file except in compliance with the License.  You may obtain a
copy of the License at http://www.apache.org/licenses/LICENSE-2.0.  Unless
required by applicable law or agreed to in writing, software distributed
under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES
OR CONDITIONS OF ANY KIND, either express or implied. See the License for
thespecific language governing permissions and limitations under the
License.
Keyfactor-v1

This reference serves to document REST-based methods to manage and integrate with Keyfactor. In addition, an embedded interface allows for the execution of calls against the current Keyfactor API instance.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package keyfactor

import (
	"encoding/json"
)

// checks if the KeyfactorApiModelsConfigurationTenantConfigurationTenantRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorApiModelsConfigurationTenantConfigurationTenantRequest{}

// KeyfactorApiModelsConfigurationTenantConfigurationTenantRequest struct for KeyfactorApiModelsConfigurationTenantConfigurationTenantRequest
type KeyfactorApiModelsConfigurationTenantConfigurationTenantRequest struct {
	ConfigurationTenant *string `json:"ConfigurationTenant,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KeyfactorApiModelsConfigurationTenantConfigurationTenantRequest KeyfactorApiModelsConfigurationTenantConfigurationTenantRequest

// NewKeyfactorApiModelsConfigurationTenantConfigurationTenantRequest instantiates a new KeyfactorApiModelsConfigurationTenantConfigurationTenantRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorApiModelsConfigurationTenantConfigurationTenantRequest() *KeyfactorApiModelsConfigurationTenantConfigurationTenantRequest {
	this := KeyfactorApiModelsConfigurationTenantConfigurationTenantRequest{}
	return &this
}

// NewKeyfactorApiModelsConfigurationTenantConfigurationTenantRequestWithDefaults instantiates a new KeyfactorApiModelsConfigurationTenantConfigurationTenantRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorApiModelsConfigurationTenantConfigurationTenantRequestWithDefaults() *KeyfactorApiModelsConfigurationTenantConfigurationTenantRequest {
	this := KeyfactorApiModelsConfigurationTenantConfigurationTenantRequest{}
	return &this
}

// GetConfigurationTenant returns the ConfigurationTenant field value if set, zero value otherwise.
func (o *KeyfactorApiModelsConfigurationTenantConfigurationTenantRequest) GetConfigurationTenant() string {
	if o == nil || isNil(o.ConfigurationTenant) {
		var ret string
		return ret
	}
	return *o.ConfigurationTenant
}

// GetConfigurationTenantOk returns a tuple with the ConfigurationTenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsConfigurationTenantConfigurationTenantRequest) GetConfigurationTenantOk() (*string, bool) {
	if o == nil || isNil(o.ConfigurationTenant) {
		return nil, false
	}
	return o.ConfigurationTenant, true
}

// HasConfigurationTenant returns a boolean if a field has been set.
func (o *KeyfactorApiModelsConfigurationTenantConfigurationTenantRequest) HasConfigurationTenant() bool {
	if o != nil && !isNil(o.ConfigurationTenant) {
		return true
	}

	return false
}

// SetConfigurationTenant gets a reference to the given string and assigns it to the ConfigurationTenant field.
func (o *KeyfactorApiModelsConfigurationTenantConfigurationTenantRequest) SetConfigurationTenant(v string) {
	o.ConfigurationTenant = &v
}

func (o KeyfactorApiModelsConfigurationTenantConfigurationTenantRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorApiModelsConfigurationTenantConfigurationTenantRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ConfigurationTenant) {
		toSerialize["ConfigurationTenant"] = o.ConfigurationTenant
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KeyfactorApiModelsConfigurationTenantConfigurationTenantRequest) UnmarshalJSON(bytes []byte) (err error) {
	varKeyfactorApiModelsConfigurationTenantConfigurationTenantRequest := _KeyfactorApiModelsConfigurationTenantConfigurationTenantRequest{}

	if err = json.Unmarshal(bytes, &varKeyfactorApiModelsConfigurationTenantConfigurationTenantRequest); err == nil {
		*o = KeyfactorApiModelsConfigurationTenantConfigurationTenantRequest(varKeyfactorApiModelsConfigurationTenantConfigurationTenantRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ConfigurationTenant")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKeyfactorApiModelsConfigurationTenantConfigurationTenantRequest struct {
	value *KeyfactorApiModelsConfigurationTenantConfigurationTenantRequest
	isSet bool
}

func (v NullableKeyfactorApiModelsConfigurationTenantConfigurationTenantRequest) Get() *KeyfactorApiModelsConfigurationTenantConfigurationTenantRequest {
	return v.value
}

func (v *NullableKeyfactorApiModelsConfigurationTenantConfigurationTenantRequest) Set(val *KeyfactorApiModelsConfigurationTenantConfigurationTenantRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorApiModelsConfigurationTenantConfigurationTenantRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorApiModelsConfigurationTenantConfigurationTenantRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorApiModelsConfigurationTenantConfigurationTenantRequest(val *KeyfactorApiModelsConfigurationTenantConfigurationTenantRequest) *NullableKeyfactorApiModelsConfigurationTenantConfigurationTenantRequest {
	return &NullableKeyfactorApiModelsConfigurationTenantConfigurationTenantRequest{value: val, isSet: true}
}

func (v NullableKeyfactorApiModelsConfigurationTenantConfigurationTenantRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorApiModelsConfigurationTenantConfigurationTenantRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


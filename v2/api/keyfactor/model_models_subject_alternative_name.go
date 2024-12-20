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

// checks if the ModelsSubjectAlternativeName type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsSubjectAlternativeName{}

// ModelsSubjectAlternativeName struct for ModelsSubjectAlternativeName
type ModelsSubjectAlternativeName struct {
	Id                   *int32  `json:"Id,omitempty"`
	Value                *string `json:"Value,omitempty"`
	Type                 *int32  `json:"Type,omitempty"`
	ValueHash            *string `json:"ValueHash,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelsSubjectAlternativeName ModelsSubjectAlternativeName

// NewModelsSubjectAlternativeName instantiates a new ModelsSubjectAlternativeName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsSubjectAlternativeName() *ModelsSubjectAlternativeName {
	this := ModelsSubjectAlternativeName{}
	return &this
}

// NewModelsSubjectAlternativeNameWithDefaults instantiates a new ModelsSubjectAlternativeName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsSubjectAlternativeNameWithDefaults() *ModelsSubjectAlternativeName {
	this := ModelsSubjectAlternativeName{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelsSubjectAlternativeName) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSubjectAlternativeName) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelsSubjectAlternativeName) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ModelsSubjectAlternativeName) SetId(v int32) {
	o.Id = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ModelsSubjectAlternativeName) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSubjectAlternativeName) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ModelsSubjectAlternativeName) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ModelsSubjectAlternativeName) SetValue(v string) {
	o.Value = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ModelsSubjectAlternativeName) GetType() int32 {
	if o == nil || isNil(o.Type) {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSubjectAlternativeName) GetTypeOk() (*int32, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ModelsSubjectAlternativeName) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *ModelsSubjectAlternativeName) SetType(v int32) {
	o.Type = &v
}

// GetValueHash returns the ValueHash field value if set, zero value otherwise.
func (o *ModelsSubjectAlternativeName) GetValueHash() string {
	if o == nil || isNil(o.ValueHash) {
		var ret string
		return ret
	}
	return *o.ValueHash
}

// GetValueHashOk returns a tuple with the ValueHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSubjectAlternativeName) GetValueHashOk() (*string, bool) {
	if o == nil || isNil(o.ValueHash) {
		return nil, false
	}
	return o.ValueHash, true
}

// HasValueHash returns a boolean if a field has been set.
func (o *ModelsSubjectAlternativeName) HasValueHash() bool {
	if o != nil && !isNil(o.ValueHash) {
		return true
	}

	return false
}

// SetValueHash gets a reference to the given string and assigns it to the ValueHash field.
func (o *ModelsSubjectAlternativeName) SetValueHash(v string) {
	o.ValueHash = &v
}

func (o ModelsSubjectAlternativeName) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsSubjectAlternativeName) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !isNil(o.Value) {
		toSerialize["Value"] = o.Value
	}
	if !isNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if !isNil(o.ValueHash) {
		toSerialize["ValueHash"] = o.ValueHash
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModelsSubjectAlternativeName) UnmarshalJSON(bytes []byte) (err error) {
	varModelsSubjectAlternativeName := _ModelsSubjectAlternativeName{}

	if err = json.Unmarshal(bytes, &varModelsSubjectAlternativeName); err == nil {
		*o = ModelsSubjectAlternativeName(varModelsSubjectAlternativeName)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Id")
		delete(additionalProperties, "Value")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "ValueHash")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelsSubjectAlternativeName struct {
	value *ModelsSubjectAlternativeName
	isSet bool
}

func (v NullableModelsSubjectAlternativeName) Get() *ModelsSubjectAlternativeName {
	return v.value
}

func (v *NullableModelsSubjectAlternativeName) Set(val *ModelsSubjectAlternativeName) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsSubjectAlternativeName) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsSubjectAlternativeName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsSubjectAlternativeName(val *ModelsSubjectAlternativeName) *NullableModelsSubjectAlternativeName {
	return &NullableModelsSubjectAlternativeName{value: val, isSet: true}
}

func (v NullableModelsSubjectAlternativeName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsSubjectAlternativeName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

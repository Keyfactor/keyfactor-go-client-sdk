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

package command

import (
	"encoding/json"
)

// checks if the KeyfactorApiPAMProviderTypeParameterResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorApiPAMProviderTypeParameterResponse{}

// KeyfactorApiPAMProviderTypeParameterResponse struct for KeyfactorApiPAMProviderTypeParameterResponse
type KeyfactorApiPAMProviderTypeParameterResponse struct {
	Id                   *int32  `json:"Id,omitempty"`
	Name                 *string `json:"Name,omitempty"`
	DisplayName          *string `json:"DisplayName,omitempty"`
	DataType             *int32  `json:"DataType,omitempty"`
	InstanceLevel        *bool   `json:"InstanceLevel,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KeyfactorApiPAMProviderTypeParameterResponse KeyfactorApiPAMProviderTypeParameterResponse

// NewKeyfactorApiPAMProviderTypeParameterResponse instantiates a new KeyfactorApiPAMProviderTypeParameterResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorApiPAMProviderTypeParameterResponse() *KeyfactorApiPAMProviderTypeParameterResponse {
	this := KeyfactorApiPAMProviderTypeParameterResponse{}
	return &this
}

// NewKeyfactorApiPAMProviderTypeParameterResponseWithDefaults instantiates a new KeyfactorApiPAMProviderTypeParameterResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorApiPAMProviderTypeParameterResponseWithDefaults() *KeyfactorApiPAMProviderTypeParameterResponse {
	this := KeyfactorApiPAMProviderTypeParameterResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KeyfactorApiPAMProviderTypeParameterResponse) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiPAMProviderTypeParameterResponse) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KeyfactorApiPAMProviderTypeParameterResponse) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *KeyfactorApiPAMProviderTypeParameterResponse) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KeyfactorApiPAMProviderTypeParameterResponse) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiPAMProviderTypeParameterResponse) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KeyfactorApiPAMProviderTypeParameterResponse) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KeyfactorApiPAMProviderTypeParameterResponse) SetName(v string) {
	o.Name = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *KeyfactorApiPAMProviderTypeParameterResponse) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiPAMProviderTypeParameterResponse) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *KeyfactorApiPAMProviderTypeParameterResponse) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *KeyfactorApiPAMProviderTypeParameterResponse) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDataType returns the DataType field value if set, zero value otherwise.
func (o *KeyfactorApiPAMProviderTypeParameterResponse) GetDataType() int32 {
	if o == nil || isNil(o.DataType) {
		var ret int32
		return ret
	}
	return *o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiPAMProviderTypeParameterResponse) GetDataTypeOk() (*int32, bool) {
	if o == nil || isNil(o.DataType) {
		return nil, false
	}
	return o.DataType, true
}

// HasDataType returns a boolean if a field has been set.
func (o *KeyfactorApiPAMProviderTypeParameterResponse) HasDataType() bool {
	if o != nil && !isNil(o.DataType) {
		return true
	}

	return false
}

// SetDataType gets a reference to the given int32 and assigns it to the DataType field.
func (o *KeyfactorApiPAMProviderTypeParameterResponse) SetDataType(v int32) {
	o.DataType = &v
}

// GetInstanceLevel returns the InstanceLevel field value if set, zero value otherwise.
func (o *KeyfactorApiPAMProviderTypeParameterResponse) GetInstanceLevel() bool {
	if o == nil || isNil(o.InstanceLevel) {
		var ret bool
		return ret
	}
	return *o.InstanceLevel
}

// GetInstanceLevelOk returns a tuple with the InstanceLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiPAMProviderTypeParameterResponse) GetInstanceLevelOk() (*bool, bool) {
	if o == nil || isNil(o.InstanceLevel) {
		return nil, false
	}
	return o.InstanceLevel, true
}

// HasInstanceLevel returns a boolean if a field has been set.
func (o *KeyfactorApiPAMProviderTypeParameterResponse) HasInstanceLevel() bool {
	if o != nil && !isNil(o.InstanceLevel) {
		return true
	}

	return false
}

// SetInstanceLevel gets a reference to the given bool and assigns it to the InstanceLevel field.
func (o *KeyfactorApiPAMProviderTypeParameterResponse) SetInstanceLevel(v bool) {
	o.InstanceLevel = &v
}

func (o KeyfactorApiPAMProviderTypeParameterResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorApiPAMProviderTypeParameterResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !isNil(o.DisplayName) {
		toSerialize["DisplayName"] = o.DisplayName
	}
	if !isNil(o.DataType) {
		toSerialize["DataType"] = o.DataType
	}
	if !isNil(o.InstanceLevel) {
		toSerialize["InstanceLevel"] = o.InstanceLevel
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KeyfactorApiPAMProviderTypeParameterResponse) UnmarshalJSON(bytes []byte) (err error) {
	varKeyfactorApiPAMProviderTypeParameterResponse := _KeyfactorApiPAMProviderTypeParameterResponse{}

	if err = json.Unmarshal(bytes, &varKeyfactorApiPAMProviderTypeParameterResponse); err == nil {
		*o = KeyfactorApiPAMProviderTypeParameterResponse(varKeyfactorApiPAMProviderTypeParameterResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Id")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "DisplayName")
		delete(additionalProperties, "DataType")
		delete(additionalProperties, "InstanceLevel")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKeyfactorApiPAMProviderTypeParameterResponse struct {
	value *KeyfactorApiPAMProviderTypeParameterResponse
	isSet bool
}

func (v NullableKeyfactorApiPAMProviderTypeParameterResponse) Get() *KeyfactorApiPAMProviderTypeParameterResponse {
	return v.value
}

func (v *NullableKeyfactorApiPAMProviderTypeParameterResponse) Set(val *KeyfactorApiPAMProviderTypeParameterResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorApiPAMProviderTypeParameterResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorApiPAMProviderTypeParameterResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorApiPAMProviderTypeParameterResponse(val *KeyfactorApiPAMProviderTypeParameterResponse) *NullableKeyfactorApiPAMProviderTypeParameterResponse {
	return &NullableKeyfactorApiPAMProviderTypeParameterResponse{value: val, isSet: true}
}

func (v NullableKeyfactorApiPAMProviderTypeParameterResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorApiPAMProviderTypeParameterResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
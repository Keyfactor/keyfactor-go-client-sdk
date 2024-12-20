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

// checks if the ModelsMetadataUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsMetadataUpdateRequest{}

// ModelsMetadataUpdateRequest struct for ModelsMetadataUpdateRequest
type ModelsMetadataUpdateRequest struct {
	Id                   *int32            `json:"Id,omitempty"`
	Metadata             map[string]string `json:"Metadata"`
	AdditionalProperties map[string]interface{}
}

type _ModelsMetadataUpdateRequest ModelsMetadataUpdateRequest

// NewModelsMetadataUpdateRequest instantiates a new ModelsMetadataUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsMetadataUpdateRequest(metadata map[string]string) *ModelsMetadataUpdateRequest {
	this := ModelsMetadataUpdateRequest{}
	this.Metadata = metadata
	return &this
}

// NewModelsMetadataUpdateRequestWithDefaults instantiates a new ModelsMetadataUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsMetadataUpdateRequestWithDefaults() *ModelsMetadataUpdateRequest {
	this := ModelsMetadataUpdateRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelsMetadataUpdateRequest) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsMetadataUpdateRequest) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelsMetadataUpdateRequest) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ModelsMetadataUpdateRequest) SetId(v int32) {
	o.Id = &v
}

// GetMetadata returns the Metadata field value
func (o *ModelsMetadataUpdateRequest) GetMetadata() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *ModelsMetadataUpdateRequest) GetMetadataOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *ModelsMetadataUpdateRequest) SetMetadata(v map[string]string) {
	o.Metadata = v
}

func (o ModelsMetadataUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsMetadataUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	toSerialize["Metadata"] = o.Metadata

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModelsMetadataUpdateRequest) UnmarshalJSON(bytes []byte) (err error) {
	varModelsMetadataUpdateRequest := _ModelsMetadataUpdateRequest{}

	if err = json.Unmarshal(bytes, &varModelsMetadataUpdateRequest); err == nil {
		*o = ModelsMetadataUpdateRequest(varModelsMetadataUpdateRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Id")
		delete(additionalProperties, "Metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelsMetadataUpdateRequest struct {
	value *ModelsMetadataUpdateRequest
	isSet bool
}

func (v NullableModelsMetadataUpdateRequest) Get() *ModelsMetadataUpdateRequest {
	return v.value
}

func (v *NullableModelsMetadataUpdateRequest) Set(val *ModelsMetadataUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsMetadataUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsMetadataUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsMetadataUpdateRequest(val *ModelsMetadataUpdateRequest) *NullableModelsMetadataUpdateRequest {
	return &NullableModelsMetadataUpdateRequest{value: val, isSet: true}
}

func (v NullableModelsMetadataUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsMetadataUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

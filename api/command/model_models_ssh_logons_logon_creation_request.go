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

// checks if the ModelsSSHLogonsLogonCreationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsSSHLogonsLogonCreationRequest{}

// ModelsSSHLogonsLogonCreationRequest struct for ModelsSSHLogonsLogonCreationRequest
type ModelsSSHLogonsLogonCreationRequest struct {
	Username             string  `json:"Username"`
	ServerId             int32   `json:"ServerId"`
	UserIds              []int32 `json:"UserIds,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelsSSHLogonsLogonCreationRequest ModelsSSHLogonsLogonCreationRequest

// NewModelsSSHLogonsLogonCreationRequest instantiates a new ModelsSSHLogonsLogonCreationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsSSHLogonsLogonCreationRequest(username string, serverId int32) *ModelsSSHLogonsLogonCreationRequest {
	this := ModelsSSHLogonsLogonCreationRequest{}
	this.Username = username
	this.ServerId = serverId
	return &this
}

// NewModelsSSHLogonsLogonCreationRequestWithDefaults instantiates a new ModelsSSHLogonsLogonCreationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsSSHLogonsLogonCreationRequestWithDefaults() *ModelsSSHLogonsLogonCreationRequest {
	this := ModelsSSHLogonsLogonCreationRequest{}
	return &this
}

// GetUsername returns the Username field value
func (o *ModelsSSHLogonsLogonCreationRequest) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *ModelsSSHLogonsLogonCreationRequest) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *ModelsSSHLogonsLogonCreationRequest) SetUsername(v string) {
	o.Username = v
}

// GetServerId returns the ServerId field value
func (o *ModelsSSHLogonsLogonCreationRequest) GetServerId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ServerId
}

// GetServerIdOk returns a tuple with the ServerId field value
// and a boolean to check if the value has been set.
func (o *ModelsSSHLogonsLogonCreationRequest) GetServerIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerId, true
}

// SetServerId sets field value
func (o *ModelsSSHLogonsLogonCreationRequest) SetServerId(v int32) {
	o.ServerId = v
}

// GetUserIds returns the UserIds field value if set, zero value otherwise.
func (o *ModelsSSHLogonsLogonCreationRequest) GetUserIds() []int32 {
	if o == nil || isNil(o.UserIds) {
		var ret []int32
		return ret
	}
	return o.UserIds
}

// GetUserIdsOk returns a tuple with the UserIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSSHLogonsLogonCreationRequest) GetUserIdsOk() ([]int32, bool) {
	if o == nil || isNil(o.UserIds) {
		return nil, false
	}
	return o.UserIds, true
}

// HasUserIds returns a boolean if a field has been set.
func (o *ModelsSSHLogonsLogonCreationRequest) HasUserIds() bool {
	if o != nil && !isNil(o.UserIds) {
		return true
	}

	return false
}

// SetUserIds gets a reference to the given []int32 and assigns it to the UserIds field.
func (o *ModelsSSHLogonsLogonCreationRequest) SetUserIds(v []int32) {
	o.UserIds = v
}

func (o ModelsSSHLogonsLogonCreationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsSSHLogonsLogonCreationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Username"] = o.Username
	toSerialize["ServerId"] = o.ServerId
	if !isNil(o.UserIds) {
		toSerialize["UserIds"] = o.UserIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModelsSSHLogonsLogonCreationRequest) UnmarshalJSON(bytes []byte) (err error) {
	varModelsSSHLogonsLogonCreationRequest := _ModelsSSHLogonsLogonCreationRequest{}

	if err = json.Unmarshal(bytes, &varModelsSSHLogonsLogonCreationRequest); err == nil {
		*o = ModelsSSHLogonsLogonCreationRequest(varModelsSSHLogonsLogonCreationRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Username")
		delete(additionalProperties, "ServerId")
		delete(additionalProperties, "UserIds")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelsSSHLogonsLogonCreationRequest struct {
	value *ModelsSSHLogonsLogonCreationRequest
	isSet bool
}

func (v NullableModelsSSHLogonsLogonCreationRequest) Get() *ModelsSSHLogonsLogonCreationRequest {
	return v.value
}

func (v *NullableModelsSSHLogonsLogonCreationRequest) Set(val *ModelsSSHLogonsLogonCreationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsSSHLogonsLogonCreationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsSSHLogonsLogonCreationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsSSHLogonsLogonCreationRequest(val *ModelsSSHLogonsLogonCreationRequest) *NullableModelsSSHLogonsLogonCreationRequest {
	return &NullableModelsSSHLogonsLogonCreationRequest{value: val, isSet: true}
}

func (v NullableModelsSSHLogonsLogonCreationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsSSHLogonsLogonCreationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
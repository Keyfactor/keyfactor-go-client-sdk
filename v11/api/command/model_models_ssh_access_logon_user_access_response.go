/*

Copyright 2023 Keyfactor
Licensed under the Apache License, Version 2.0 (the License); you may
not use this file except in compliance with the License.  You may obtain a
copy of the License at http://www.apache.org/licenses/LICENSE-2.0.  Unless
required by applicable law or agreed to in writing, software distributed
under the License is distributed on an AS IS BASIS, WITHOUT WARRANTIES
OR CONDITIONS OF ANY KIND, either express or implied. See the License for
the specific language governing permissions and limitations under the
License.

Keyfactor API Reference and Utility

<p>This page provides a utility through which the Keyfactor API endpoints can be called and results returned.                                                           It is intended to be used primarily for validation, testing and workflow development.                                                           It also serves secondarily as documentation for the API.</p>                                                          <p>If you would like to view documentation containing details on the Keyfactor API and endpoints,                                                           please refer to the Web API section of the Keyfactor Command documentation.</p>

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package command

import (
	"encoding/json"
)

// checks if the ModelsSSHAccessLogonUserAccessResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsSSHAccessLogonUserAccessResponse{}

// ModelsSSHAccessLogonUserAccessResponse struct for ModelsSSHAccessLogonUserAccessResponse
type ModelsSSHAccessLogonUserAccessResponse struct {
	LogonId   NullableInt32                   `json:"logonId,omitempty"`
	LogonName NullableString                  `json:"logonName,omitempty"`
	Users     []ModelsSSHUsersSshUserResponse `json:"users,omitempty"`
}

// NewModelsSSHAccessLogonUserAccessResponse instantiates a new ModelsSSHAccessLogonUserAccessResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsSSHAccessLogonUserAccessResponse() *ModelsSSHAccessLogonUserAccessResponse {
	this := ModelsSSHAccessLogonUserAccessResponse{}
	return &this
}

// NewModelsSSHAccessLogonUserAccessResponseWithDefaults instantiates a new ModelsSSHAccessLogonUserAccessResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsSSHAccessLogonUserAccessResponseWithDefaults() *ModelsSSHAccessLogonUserAccessResponse {
	this := ModelsSSHAccessLogonUserAccessResponse{}
	return &this
}

// GetLogonId returns the LogonId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelsSSHAccessLogonUserAccessResponse) GetLogonId() int32 {
	if o == nil || isNil(o.LogonId.Get()) {
		var ret int32
		return ret
	}
	return *o.LogonId.Get()
}

// GetLogonIdOk returns a tuple with the LogonId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelsSSHAccessLogonUserAccessResponse) GetLogonIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogonId.Get(), o.LogonId.IsSet()
}

// HasLogonId returns a boolean if a field has been set.
func (o *ModelsSSHAccessLogonUserAccessResponse) HasLogonId() bool {
	if o != nil && o.LogonId.IsSet() {
		return true
	}

	return false
}

// SetLogonId gets a reference to the given NullableInt32 and assigns it to the LogonId field.
func (o *ModelsSSHAccessLogonUserAccessResponse) SetLogonId(v int32) {
	o.LogonId.Set(&v)
}

// SetLogonIdNil sets the value for LogonId to be an explicit nil
func (o *ModelsSSHAccessLogonUserAccessResponse) SetLogonIdNil() {
	o.LogonId.Set(nil)
}

// UnsetLogonId ensures that no value is present for LogonId, not even an explicit nil
func (o *ModelsSSHAccessLogonUserAccessResponse) UnsetLogonId() {
	o.LogonId.Unset()
}

// GetLogonName returns the LogonName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelsSSHAccessLogonUserAccessResponse) GetLogonName() string {
	if o == nil || isNil(o.LogonName.Get()) {
		var ret string
		return ret
	}
	return *o.LogonName.Get()
}

// GetLogonNameOk returns a tuple with the LogonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelsSSHAccessLogonUserAccessResponse) GetLogonNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogonName.Get(), o.LogonName.IsSet()
}

// HasLogonName returns a boolean if a field has been set.
func (o *ModelsSSHAccessLogonUserAccessResponse) HasLogonName() bool {
	if o != nil && o.LogonName.IsSet() {
		return true
	}

	return false
}

// SetLogonName gets a reference to the given NullableString and assigns it to the LogonName field.
func (o *ModelsSSHAccessLogonUserAccessResponse) SetLogonName(v string) {
	o.LogonName.Set(&v)
}

// SetLogonNameNil sets the value for LogonName to be an explicit nil
func (o *ModelsSSHAccessLogonUserAccessResponse) SetLogonNameNil() {
	o.LogonName.Set(nil)
}

// UnsetLogonName ensures that no value is present for LogonName, not even an explicit nil
func (o *ModelsSSHAccessLogonUserAccessResponse) UnsetLogonName() {
	o.LogonName.Unset()
}

// GetUsers returns the Users field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelsSSHAccessLogonUserAccessResponse) GetUsers() []ModelsSSHUsersSshUserResponse {
	if o == nil {
		var ret []ModelsSSHUsersSshUserResponse
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelsSSHAccessLogonUserAccessResponse) GetUsersOk() ([]ModelsSSHUsersSshUserResponse, bool) {
	if o == nil || isNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *ModelsSSHAccessLogonUserAccessResponse) HasUsers() bool {
	if o != nil && isNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []ModelsSSHUsersSshUserResponse and assigns it to the Users field.
func (o *ModelsSSHAccessLogonUserAccessResponse) SetUsers(v []ModelsSSHUsersSshUserResponse) {
	o.Users = v
}

func (o ModelsSSHAccessLogonUserAccessResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsSSHAccessLogonUserAccessResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.LogonId.IsSet() {
		toSerialize["logonId"] = o.LogonId.Get()
	}
	if o.LogonName.IsSet() {
		toSerialize["logonName"] = o.LogonName.Get()
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	return toSerialize, nil
}

type NullableModelsSSHAccessLogonUserAccessResponse struct {
	value *ModelsSSHAccessLogonUserAccessResponse
	isSet bool
}

func (v NullableModelsSSHAccessLogonUserAccessResponse) Get() *ModelsSSHAccessLogonUserAccessResponse {
	return v.value
}

func (v *NullableModelsSSHAccessLogonUserAccessResponse) Set(val *ModelsSSHAccessLogonUserAccessResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsSSHAccessLogonUserAccessResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsSSHAccessLogonUserAccessResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsSSHAccessLogonUserAccessResponse(val *ModelsSSHAccessLogonUserAccessResponse) *NullableModelsSSHAccessLogonUserAccessResponse {
	return &NullableModelsSSHAccessLogonUserAccessResponse{value: val, isSet: true}
}

func (v NullableModelsSSHAccessLogonUserAccessResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsSSHAccessLogonUserAccessResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
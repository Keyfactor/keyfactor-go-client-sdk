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

// checks if the ModelsSSHAccessServerAccessResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsSSHAccessServerAccessResponse{}

// ModelsSSHAccessServerAccessResponse struct for ModelsSSHAccessServerAccessResponse
type ModelsSSHAccessServerAccessResponse struct {
	ServerId   *int32                                   `json:"serverId,omitempty"`
	LogonUsers []ModelsSSHAccessLogonUserAccessResponse `json:"logonUsers,omitempty"`
}

// NewModelsSSHAccessServerAccessResponse instantiates a new ModelsSSHAccessServerAccessResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsSSHAccessServerAccessResponse() *ModelsSSHAccessServerAccessResponse {
	this := ModelsSSHAccessServerAccessResponse{}
	return &this
}

// NewModelsSSHAccessServerAccessResponseWithDefaults instantiates a new ModelsSSHAccessServerAccessResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsSSHAccessServerAccessResponseWithDefaults() *ModelsSSHAccessServerAccessResponse {
	this := ModelsSSHAccessServerAccessResponse{}
	return &this
}

// GetServerId returns the ServerId field value if set, zero value otherwise.
func (o *ModelsSSHAccessServerAccessResponse) GetServerId() int32 {
	if o == nil || isNil(o.ServerId) {
		var ret int32
		return ret
	}
	return *o.ServerId
}

// GetServerIdOk returns a tuple with the ServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSSHAccessServerAccessResponse) GetServerIdOk() (*int32, bool) {
	if o == nil || isNil(o.ServerId) {
		return nil, false
	}
	return o.ServerId, true
}

// HasServerId returns a boolean if a field has been set.
func (o *ModelsSSHAccessServerAccessResponse) HasServerId() bool {
	if o != nil && !isNil(o.ServerId) {
		return true
	}

	return false
}

// SetServerId gets a reference to the given int32 and assigns it to the ServerId field.
func (o *ModelsSSHAccessServerAccessResponse) SetServerId(v int32) {
	o.ServerId = &v
}

// GetLogonUsers returns the LogonUsers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelsSSHAccessServerAccessResponse) GetLogonUsers() []ModelsSSHAccessLogonUserAccessResponse {
	if o == nil {
		var ret []ModelsSSHAccessLogonUserAccessResponse
		return ret
	}
	return o.LogonUsers
}

// GetLogonUsersOk returns a tuple with the LogonUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelsSSHAccessServerAccessResponse) GetLogonUsersOk() ([]ModelsSSHAccessLogonUserAccessResponse, bool) {
	if o == nil || isNil(o.LogonUsers) {
		return nil, false
	}
	return o.LogonUsers, true
}

// HasLogonUsers returns a boolean if a field has been set.
func (o *ModelsSSHAccessServerAccessResponse) HasLogonUsers() bool {
	if o != nil && isNil(o.LogonUsers) {
		return true
	}

	return false
}

// SetLogonUsers gets a reference to the given []ModelsSSHAccessLogonUserAccessResponse and assigns it to the LogonUsers field.
func (o *ModelsSSHAccessServerAccessResponse) SetLogonUsers(v []ModelsSSHAccessLogonUserAccessResponse) {
	o.LogonUsers = v
}

func (o ModelsSSHAccessServerAccessResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsSSHAccessServerAccessResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ServerId) {
		toSerialize["serverId"] = o.ServerId
	}
	if o.LogonUsers != nil {
		toSerialize["logonUsers"] = o.LogonUsers
	}
	return toSerialize, nil
}

type NullableModelsSSHAccessServerAccessResponse struct {
	value *ModelsSSHAccessServerAccessResponse
	isSet bool
}

func (v NullableModelsSSHAccessServerAccessResponse) Get() *ModelsSSHAccessServerAccessResponse {
	return v.value
}

func (v *NullableModelsSSHAccessServerAccessResponse) Set(val *ModelsSSHAccessServerAccessResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsSSHAccessServerAccessResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsSSHAccessServerAccessResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsSSHAccessServerAccessResponse(val *ModelsSSHAccessServerAccessResponse) *NullableModelsSSHAccessServerAccessResponse {
	return &NullableModelsSSHAccessServerAccessResponse{value: val, isSet: true}
}

func (v NullableModelsSSHAccessServerAccessResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsSSHAccessServerAccessResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
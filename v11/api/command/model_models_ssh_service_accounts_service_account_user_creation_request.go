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

// checks if the ModelsSSHServiceAccountsServiceAccountUserCreationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsSSHServiceAccountsServiceAccountUserCreationRequest{}

// ModelsSSHServiceAccountsServiceAccountUserCreationRequest struct for ModelsSSHServiceAccountsServiceAccountUserCreationRequest
type ModelsSSHServiceAccountsServiceAccountUserCreationRequest struct {
	Username string  `json:"username"`
	LogonIds []int32 `json:"logonIds,omitempty"`
}

// NewModelsSSHServiceAccountsServiceAccountUserCreationRequest instantiates a new ModelsSSHServiceAccountsServiceAccountUserCreationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsSSHServiceAccountsServiceAccountUserCreationRequest(username string) *ModelsSSHServiceAccountsServiceAccountUserCreationRequest {
	this := ModelsSSHServiceAccountsServiceAccountUserCreationRequest{}
	this.Username = username
	return &this
}

// NewModelsSSHServiceAccountsServiceAccountUserCreationRequestWithDefaults instantiates a new ModelsSSHServiceAccountsServiceAccountUserCreationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsSSHServiceAccountsServiceAccountUserCreationRequestWithDefaults() *ModelsSSHServiceAccountsServiceAccountUserCreationRequest {
	this := ModelsSSHServiceAccountsServiceAccountUserCreationRequest{}
	return &this
}

// GetUsername returns the Username field value
func (o *ModelsSSHServiceAccountsServiceAccountUserCreationRequest) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *ModelsSSHServiceAccountsServiceAccountUserCreationRequest) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *ModelsSSHServiceAccountsServiceAccountUserCreationRequest) SetUsername(v string) {
	o.Username = v
}

// GetLogonIds returns the LogonIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelsSSHServiceAccountsServiceAccountUserCreationRequest) GetLogonIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.LogonIds
}

// GetLogonIdsOk returns a tuple with the LogonIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelsSSHServiceAccountsServiceAccountUserCreationRequest) GetLogonIdsOk() ([]int32, bool) {
	if o == nil || isNil(o.LogonIds) {
		return nil, false
	}
	return o.LogonIds, true
}

// HasLogonIds returns a boolean if a field has been set.
func (o *ModelsSSHServiceAccountsServiceAccountUserCreationRequest) HasLogonIds() bool {
	if o != nil && isNil(o.LogonIds) {
		return true
	}

	return false
}

// SetLogonIds gets a reference to the given []int32 and assigns it to the LogonIds field.
func (o *ModelsSSHServiceAccountsServiceAccountUserCreationRequest) SetLogonIds(v []int32) {
	o.LogonIds = v
}

func (o ModelsSSHServiceAccountsServiceAccountUserCreationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsSSHServiceAccountsServiceAccountUserCreationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["username"] = o.Username
	if o.LogonIds != nil {
		toSerialize["logonIds"] = o.LogonIds
	}
	return toSerialize, nil
}

type NullableModelsSSHServiceAccountsServiceAccountUserCreationRequest struct {
	value *ModelsSSHServiceAccountsServiceAccountUserCreationRequest
	isSet bool
}

func (v NullableModelsSSHServiceAccountsServiceAccountUserCreationRequest) Get() *ModelsSSHServiceAccountsServiceAccountUserCreationRequest {
	return v.value
}

func (v *NullableModelsSSHServiceAccountsServiceAccountUserCreationRequest) Set(val *ModelsSSHServiceAccountsServiceAccountUserCreationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsSSHServiceAccountsServiceAccountUserCreationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsSSHServiceAccountsServiceAccountUserCreationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsSSHServiceAccountsServiceAccountUserCreationRequest(val *ModelsSSHServiceAccountsServiceAccountUserCreationRequest) *NullableModelsSSHServiceAccountsServiceAccountUserCreationRequest {
	return &NullableModelsSSHServiceAccountsServiceAccountUserCreationRequest{value: val, isSet: true}
}

func (v NullableModelsSSHServiceAccountsServiceAccountUserCreationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsSSHServiceAccountsServiceAccountUserCreationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
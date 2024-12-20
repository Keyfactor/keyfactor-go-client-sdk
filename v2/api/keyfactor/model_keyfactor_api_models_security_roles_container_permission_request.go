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

// checks if the KeyfactorApiModelsSecurityRolesContainerPermissionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorApiModelsSecurityRolesContainerPermissionRequest{}

// KeyfactorApiModelsSecurityRolesContainerPermissionRequest struct for KeyfactorApiModelsSecurityRolesContainerPermissionRequest
type KeyfactorApiModelsSecurityRolesContainerPermissionRequest struct {
	ContainerId          *int32  `json:"ContainerId,omitempty"`
	Permission           *string `json:"Permission,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KeyfactorApiModelsSecurityRolesContainerPermissionRequest KeyfactorApiModelsSecurityRolesContainerPermissionRequest

// NewKeyfactorApiModelsSecurityRolesContainerPermissionRequest instantiates a new KeyfactorApiModelsSecurityRolesContainerPermissionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorApiModelsSecurityRolesContainerPermissionRequest() *KeyfactorApiModelsSecurityRolesContainerPermissionRequest {
	this := KeyfactorApiModelsSecurityRolesContainerPermissionRequest{}
	return &this
}

// NewKeyfactorApiModelsSecurityRolesContainerPermissionRequestWithDefaults instantiates a new KeyfactorApiModelsSecurityRolesContainerPermissionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorApiModelsSecurityRolesContainerPermissionRequestWithDefaults() *KeyfactorApiModelsSecurityRolesContainerPermissionRequest {
	this := KeyfactorApiModelsSecurityRolesContainerPermissionRequest{}
	return &this
}

// GetContainerId returns the ContainerId field value if set, zero value otherwise.
func (o *KeyfactorApiModelsSecurityRolesContainerPermissionRequest) GetContainerId() int32 {
	if o == nil || isNil(o.ContainerId) {
		var ret int32
		return ret
	}
	return *o.ContainerId
}

// GetContainerIdOk returns a tuple with the ContainerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSecurityRolesContainerPermissionRequest) GetContainerIdOk() (*int32, bool) {
	if o == nil || isNil(o.ContainerId) {
		return nil, false
	}
	return o.ContainerId, true
}

// HasContainerId returns a boolean if a field has been set.
func (o *KeyfactorApiModelsSecurityRolesContainerPermissionRequest) HasContainerId() bool {
	if o != nil && !isNil(o.ContainerId) {
		return true
	}

	return false
}

// SetContainerId gets a reference to the given int32 and assigns it to the ContainerId field.
func (o *KeyfactorApiModelsSecurityRolesContainerPermissionRequest) SetContainerId(v int32) {
	o.ContainerId = &v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *KeyfactorApiModelsSecurityRolesContainerPermissionRequest) GetPermission() string {
	if o == nil || isNil(o.Permission) {
		var ret string
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSecurityRolesContainerPermissionRequest) GetPermissionOk() (*string, bool) {
	if o == nil || isNil(o.Permission) {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *KeyfactorApiModelsSecurityRolesContainerPermissionRequest) HasPermission() bool {
	if o != nil && !isNil(o.Permission) {
		return true
	}

	return false
}

// SetPermission gets a reference to the given string and assigns it to the Permission field.
func (o *KeyfactorApiModelsSecurityRolesContainerPermissionRequest) SetPermission(v string) {
	o.Permission = &v
}

func (o KeyfactorApiModelsSecurityRolesContainerPermissionRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorApiModelsSecurityRolesContainerPermissionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ContainerId) {
		toSerialize["ContainerId"] = o.ContainerId
	}
	if !isNil(o.Permission) {
		toSerialize["Permission"] = o.Permission
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KeyfactorApiModelsSecurityRolesContainerPermissionRequest) UnmarshalJSON(bytes []byte) (err error) {
	varKeyfactorApiModelsSecurityRolesContainerPermissionRequest := _KeyfactorApiModelsSecurityRolesContainerPermissionRequest{}

	if err = json.Unmarshal(bytes, &varKeyfactorApiModelsSecurityRolesContainerPermissionRequest); err == nil {
		*o = KeyfactorApiModelsSecurityRolesContainerPermissionRequest(varKeyfactorApiModelsSecurityRolesContainerPermissionRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ContainerId")
		delete(additionalProperties, "Permission")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKeyfactorApiModelsSecurityRolesContainerPermissionRequest struct {
	value *KeyfactorApiModelsSecurityRolesContainerPermissionRequest
	isSet bool
}

func (v NullableKeyfactorApiModelsSecurityRolesContainerPermissionRequest) Get() *KeyfactorApiModelsSecurityRolesContainerPermissionRequest {
	return v.value
}

func (v *NullableKeyfactorApiModelsSecurityRolesContainerPermissionRequest) Set(val *KeyfactorApiModelsSecurityRolesContainerPermissionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorApiModelsSecurityRolesContainerPermissionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorApiModelsSecurityRolesContainerPermissionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorApiModelsSecurityRolesContainerPermissionRequest(val *KeyfactorApiModelsSecurityRolesContainerPermissionRequest) *NullableKeyfactorApiModelsSecurityRolesContainerPermissionRequest {
	return &NullableKeyfactorApiModelsSecurityRolesContainerPermissionRequest{value: val, isSet: true}
}

func (v NullableKeyfactorApiModelsSecurityRolesContainerPermissionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorApiModelsSecurityRolesContainerPermissionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

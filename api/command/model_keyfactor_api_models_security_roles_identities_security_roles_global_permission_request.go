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

// checks if the KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest{}

// KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest struct for KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest
type KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest struct {
	Area                 *string `json:"Area,omitempty"`
	Permission           *string `json:"Permission,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest

// NewKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest instantiates a new KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest() *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest {
	this := KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest{}
	return &this
}

// NewKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequestWithDefaults instantiates a new KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequestWithDefaults() *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest {
	this := KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest{}
	return &this
}

// GetArea returns the Area field value if set, zero value otherwise.
func (o *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest) GetArea() string {
	if o == nil || isNil(o.Area) {
		var ret string
		return ret
	}
	return *o.Area
}

// GetAreaOk returns a tuple with the Area field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest) GetAreaOk() (*string, bool) {
	if o == nil || isNil(o.Area) {
		return nil, false
	}
	return o.Area, true
}

// HasArea returns a boolean if a field has been set.
func (o *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest) HasArea() bool {
	if o != nil && !isNil(o.Area) {
		return true
	}

	return false
}

// SetArea gets a reference to the given string and assigns it to the Area field.
func (o *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest) SetArea(v string) {
	o.Area = &v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest) GetPermission() string {
	if o == nil || isNil(o.Permission) {
		var ret string
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest) GetPermissionOk() (*string, bool) {
	if o == nil || isNil(o.Permission) {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest) HasPermission() bool {
	if o != nil && !isNil(o.Permission) {
		return true
	}

	return false
}

// SetPermission gets a reference to the given string and assigns it to the Permission field.
func (o *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest) SetPermission(v string) {
	o.Permission = &v
}

func (o KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Area) {
		toSerialize["Area"] = o.Area
	}
	if !isNil(o.Permission) {
		toSerialize["Permission"] = o.Permission
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest) UnmarshalJSON(bytes []byte) (err error) {
	varKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest := _KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest{}

	if err = json.Unmarshal(bytes, &varKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest); err == nil {
		*o = KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest(varKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Area")
		delete(additionalProperties, "Permission")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest struct {
	value *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest
	isSet bool
}

func (v NullableKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest) Get() *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest {
	return v.value
}

func (v *NullableKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest) Set(val *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest(val *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest) *NullableKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest {
	return &NullableKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest{value: val, isSet: true}
}

func (v NullableKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
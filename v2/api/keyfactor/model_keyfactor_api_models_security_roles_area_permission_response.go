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

// checks if the KeyfactorApiModelsSecurityRolesAreaPermissionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorApiModelsSecurityRolesAreaPermissionResponse{}

// KeyfactorApiModelsSecurityRolesAreaPermissionResponse struct for KeyfactorApiModelsSecurityRolesAreaPermissionResponse
type KeyfactorApiModelsSecurityRolesAreaPermissionResponse struct {
	Type                 *string `json:"Type,omitempty"`
	Area                 *string `json:"Area,omitempty"`
	Permission           *string `json:"Permission,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KeyfactorApiModelsSecurityRolesAreaPermissionResponse KeyfactorApiModelsSecurityRolesAreaPermissionResponse

// NewKeyfactorApiModelsSecurityRolesAreaPermissionResponse instantiates a new KeyfactorApiModelsSecurityRolesAreaPermissionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorApiModelsSecurityRolesAreaPermissionResponse() *KeyfactorApiModelsSecurityRolesAreaPermissionResponse {
	this := KeyfactorApiModelsSecurityRolesAreaPermissionResponse{}
	return &this
}

// NewKeyfactorApiModelsSecurityRolesAreaPermissionResponseWithDefaults instantiates a new KeyfactorApiModelsSecurityRolesAreaPermissionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorApiModelsSecurityRolesAreaPermissionResponseWithDefaults() *KeyfactorApiModelsSecurityRolesAreaPermissionResponse {
	this := KeyfactorApiModelsSecurityRolesAreaPermissionResponse{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *KeyfactorApiModelsSecurityRolesAreaPermissionResponse) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSecurityRolesAreaPermissionResponse) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *KeyfactorApiModelsSecurityRolesAreaPermissionResponse) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *KeyfactorApiModelsSecurityRolesAreaPermissionResponse) SetType(v string) {
	o.Type = &v
}

// GetArea returns the Area field value if set, zero value otherwise.
func (o *KeyfactorApiModelsSecurityRolesAreaPermissionResponse) GetArea() string {
	if o == nil || isNil(o.Area) {
		var ret string
		return ret
	}
	return *o.Area
}

// GetAreaOk returns a tuple with the Area field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSecurityRolesAreaPermissionResponse) GetAreaOk() (*string, bool) {
	if o == nil || isNil(o.Area) {
		return nil, false
	}
	return o.Area, true
}

// HasArea returns a boolean if a field has been set.
func (o *KeyfactorApiModelsSecurityRolesAreaPermissionResponse) HasArea() bool {
	if o != nil && !isNil(o.Area) {
		return true
	}

	return false
}

// SetArea gets a reference to the given string and assigns it to the Area field.
func (o *KeyfactorApiModelsSecurityRolesAreaPermissionResponse) SetArea(v string) {
	o.Area = &v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *KeyfactorApiModelsSecurityRolesAreaPermissionResponse) GetPermission() string {
	if o == nil || isNil(o.Permission) {
		var ret string
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSecurityRolesAreaPermissionResponse) GetPermissionOk() (*string, bool) {
	if o == nil || isNil(o.Permission) {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *KeyfactorApiModelsSecurityRolesAreaPermissionResponse) HasPermission() bool {
	if o != nil && !isNil(o.Permission) {
		return true
	}

	return false
}

// SetPermission gets a reference to the given string and assigns it to the Permission field.
func (o *KeyfactorApiModelsSecurityRolesAreaPermissionResponse) SetPermission(v string) {
	o.Permission = &v
}

func (o KeyfactorApiModelsSecurityRolesAreaPermissionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorApiModelsSecurityRolesAreaPermissionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
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

func (o *KeyfactorApiModelsSecurityRolesAreaPermissionResponse) UnmarshalJSON(bytes []byte) (err error) {
	varKeyfactorApiModelsSecurityRolesAreaPermissionResponse := _KeyfactorApiModelsSecurityRolesAreaPermissionResponse{}

	if err = json.Unmarshal(bytes, &varKeyfactorApiModelsSecurityRolesAreaPermissionResponse); err == nil {
		*o = KeyfactorApiModelsSecurityRolesAreaPermissionResponse(varKeyfactorApiModelsSecurityRolesAreaPermissionResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Type")
		delete(additionalProperties, "Area")
		delete(additionalProperties, "Permission")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKeyfactorApiModelsSecurityRolesAreaPermissionResponse struct {
	value *KeyfactorApiModelsSecurityRolesAreaPermissionResponse
	isSet bool
}

func (v NullableKeyfactorApiModelsSecurityRolesAreaPermissionResponse) Get() *KeyfactorApiModelsSecurityRolesAreaPermissionResponse {
	return v.value
}

func (v *NullableKeyfactorApiModelsSecurityRolesAreaPermissionResponse) Set(val *KeyfactorApiModelsSecurityRolesAreaPermissionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorApiModelsSecurityRolesAreaPermissionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorApiModelsSecurityRolesAreaPermissionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorApiModelsSecurityRolesAreaPermissionResponse(val *KeyfactorApiModelsSecurityRolesAreaPermissionResponse) *NullableKeyfactorApiModelsSecurityRolesAreaPermissionResponse {
	return &NullableKeyfactorApiModelsSecurityRolesAreaPermissionResponse{value: val, isSet: true}
}

func (v NullableKeyfactorApiModelsSecurityRolesAreaPermissionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorApiModelsSecurityRolesAreaPermissionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

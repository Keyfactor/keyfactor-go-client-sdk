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

// checks if the KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest{}

// KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest struct for KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest
type KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest struct {
	CollectionId         *int32  `json:"CollectionId,omitempty"`
	Permission           *string `json:"Permission,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest

// NewKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest instantiates a new KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest() *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest {
	this := KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest{}
	return &this
}

// NewKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequestWithDefaults instantiates a new KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequestWithDefaults() *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest {
	this := KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest{}
	return &this
}

// GetCollectionId returns the CollectionId field value if set, zero value otherwise.
func (o *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest) GetCollectionId() int32 {
	if o == nil || isNil(o.CollectionId) {
		var ret int32
		return ret
	}
	return *o.CollectionId
}

// GetCollectionIdOk returns a tuple with the CollectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest) GetCollectionIdOk() (*int32, bool) {
	if o == nil || isNil(o.CollectionId) {
		return nil, false
	}
	return o.CollectionId, true
}

// HasCollectionId returns a boolean if a field has been set.
func (o *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest) HasCollectionId() bool {
	if o != nil && !isNil(o.CollectionId) {
		return true
	}

	return false
}

// SetCollectionId gets a reference to the given int32 and assigns it to the CollectionId field.
func (o *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest) SetCollectionId(v int32) {
	o.CollectionId = &v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest) GetPermission() string {
	if o == nil || isNil(o.Permission) {
		var ret string
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest) GetPermissionOk() (*string, bool) {
	if o == nil || isNil(o.Permission) {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest) HasPermission() bool {
	if o != nil && !isNil(o.Permission) {
		return true
	}

	return false
}

// SetPermission gets a reference to the given string and assigns it to the Permission field.
func (o *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest) SetPermission(v string) {
	o.Permission = &v
}

func (o KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CollectionId) {
		toSerialize["CollectionId"] = o.CollectionId
	}
	if !isNil(o.Permission) {
		toSerialize["Permission"] = o.Permission
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest) UnmarshalJSON(bytes []byte) (err error) {
	varKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest := _KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest{}

	if err = json.Unmarshal(bytes, &varKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest); err == nil {
		*o = KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest(varKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "CollectionId")
		delete(additionalProperties, "Permission")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest struct {
	value *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest
	isSet bool
}

func (v NullableKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest) Get() *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest {
	return v.value
}

func (v *NullableKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest) Set(val *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest(val *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest) *NullableKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest {
	return &NullableKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest{value: val, isSet: true}
}

func (v NullableKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

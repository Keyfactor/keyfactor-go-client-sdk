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

// checks if the KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse{}

// KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse struct for KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse
type KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse struct {
	CollectionId         *int32  `json:"CollectionId,omitempty"`
	Name                 *string `json:"Name,omitempty"`
	Permission           *string `json:"Permission,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse

// NewKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse instantiates a new KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse() *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse {
	this := KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse{}
	return &this
}

// NewKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponseWithDefaults instantiates a new KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponseWithDefaults() *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse {
	this := KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse{}
	return &this
}

// GetCollectionId returns the CollectionId field value if set, zero value otherwise.
func (o *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse) GetCollectionId() int32 {
	if o == nil || isNil(o.CollectionId) {
		var ret int32
		return ret
	}
	return *o.CollectionId
}

// GetCollectionIdOk returns a tuple with the CollectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse) GetCollectionIdOk() (*int32, bool) {
	if o == nil || isNil(o.CollectionId) {
		return nil, false
	}
	return o.CollectionId, true
}

// HasCollectionId returns a boolean if a field has been set.
func (o *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse) HasCollectionId() bool {
	if o != nil && !isNil(o.CollectionId) {
		return true
	}

	return false
}

// SetCollectionId gets a reference to the given int32 and assigns it to the CollectionId field.
func (o *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse) SetCollectionId(v int32) {
	o.CollectionId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse) SetName(v string) {
	o.Name = &v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse) GetPermission() string {
	if o == nil || isNil(o.Permission) {
		var ret string
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse) GetPermissionOk() (*string, bool) {
	if o == nil || isNil(o.Permission) {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse) HasPermission() bool {
	if o != nil && !isNil(o.Permission) {
		return true
	}

	return false
}

// SetPermission gets a reference to the given string and assigns it to the Permission field.
func (o *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse) SetPermission(v string) {
	o.Permission = &v
}

func (o KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CollectionId) {
		toSerialize["CollectionId"] = o.CollectionId
	}
	if !isNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !isNil(o.Permission) {
		toSerialize["Permission"] = o.Permission
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse) UnmarshalJSON(bytes []byte) (err error) {
	varKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse := _KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse{}

	if err = json.Unmarshal(bytes, &varKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse); err == nil {
		*o = KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse(varKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "CollectionId")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Permission")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse struct {
	value *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse
	isSet bool
}

func (v NullableKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse) Get() *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse {
	return v.value
}

func (v *NullableKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse) Set(val *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse(val *KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse) *NullableKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse {
	return &NullableKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse{value: val, isSet: true}
}

func (v NullableKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
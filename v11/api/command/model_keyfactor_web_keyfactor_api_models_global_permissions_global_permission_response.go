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

// checks if the KeyfactorWebKeyfactorApiModelsGlobalPermissionsGlobalPermissionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorWebKeyfactorApiModelsGlobalPermissionsGlobalPermissionResponse{}

// KeyfactorWebKeyfactorApiModelsGlobalPermissionsGlobalPermissionResponse struct for KeyfactorWebKeyfactorApiModelsGlobalPermissionsGlobalPermissionResponse
type KeyfactorWebKeyfactorApiModelsGlobalPermissionsGlobalPermissionResponse struct {
	Area       NullableString `json:"area,omitempty"`
	Permission NullableString `json:"permission,omitempty"`
}

// NewKeyfactorWebKeyfactorApiModelsGlobalPermissionsGlobalPermissionResponse instantiates a new KeyfactorWebKeyfactorApiModelsGlobalPermissionsGlobalPermissionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorWebKeyfactorApiModelsGlobalPermissionsGlobalPermissionResponse() *KeyfactorWebKeyfactorApiModelsGlobalPermissionsGlobalPermissionResponse {
	this := KeyfactorWebKeyfactorApiModelsGlobalPermissionsGlobalPermissionResponse{}
	return &this
}

// NewKeyfactorWebKeyfactorApiModelsGlobalPermissionsGlobalPermissionResponseWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsGlobalPermissionsGlobalPermissionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorWebKeyfactorApiModelsGlobalPermissionsGlobalPermissionResponseWithDefaults() *KeyfactorWebKeyfactorApiModelsGlobalPermissionsGlobalPermissionResponse {
	this := KeyfactorWebKeyfactorApiModelsGlobalPermissionsGlobalPermissionResponse{}
	return &this
}

// GetArea returns the Area field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorWebKeyfactorApiModelsGlobalPermissionsGlobalPermissionResponse) GetArea() string {
	if o == nil || isNil(o.Area.Get()) {
		var ret string
		return ret
	}
	return *o.Area.Get()
}

// GetAreaOk returns a tuple with the Area field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorWebKeyfactorApiModelsGlobalPermissionsGlobalPermissionResponse) GetAreaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Area.Get(), o.Area.IsSet()
}

// HasArea returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsGlobalPermissionsGlobalPermissionResponse) HasArea() bool {
	if o != nil && o.Area.IsSet() {
		return true
	}

	return false
}

// SetArea gets a reference to the given NullableString and assigns it to the Area field.
func (o *KeyfactorWebKeyfactorApiModelsGlobalPermissionsGlobalPermissionResponse) SetArea(v string) {
	o.Area.Set(&v)
}

// SetAreaNil sets the value for Area to be an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsGlobalPermissionsGlobalPermissionResponse) SetAreaNil() {
	o.Area.Set(nil)
}

// UnsetArea ensures that no value is present for Area, not even an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsGlobalPermissionsGlobalPermissionResponse) UnsetArea() {
	o.Area.Unset()
}

// GetPermission returns the Permission field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorWebKeyfactorApiModelsGlobalPermissionsGlobalPermissionResponse) GetPermission() string {
	if o == nil || isNil(o.Permission.Get()) {
		var ret string
		return ret
	}
	return *o.Permission.Get()
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorWebKeyfactorApiModelsGlobalPermissionsGlobalPermissionResponse) GetPermissionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Permission.Get(), o.Permission.IsSet()
}

// HasPermission returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsGlobalPermissionsGlobalPermissionResponse) HasPermission() bool {
	if o != nil && o.Permission.IsSet() {
		return true
	}

	return false
}

// SetPermission gets a reference to the given NullableString and assigns it to the Permission field.
func (o *KeyfactorWebKeyfactorApiModelsGlobalPermissionsGlobalPermissionResponse) SetPermission(v string) {
	o.Permission.Set(&v)
}

// SetPermissionNil sets the value for Permission to be an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsGlobalPermissionsGlobalPermissionResponse) SetPermissionNil() {
	o.Permission.Set(nil)
}

// UnsetPermission ensures that no value is present for Permission, not even an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsGlobalPermissionsGlobalPermissionResponse) UnsetPermission() {
	o.Permission.Unset()
}

func (o KeyfactorWebKeyfactorApiModelsGlobalPermissionsGlobalPermissionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorWebKeyfactorApiModelsGlobalPermissionsGlobalPermissionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Area.IsSet() {
		toSerialize["area"] = o.Area.Get()
	}
	if o.Permission.IsSet() {
		toSerialize["permission"] = o.Permission.Get()
	}
	return toSerialize, nil
}

type NullableKeyfactorWebKeyfactorApiModelsGlobalPermissionsGlobalPermissionResponse struct {
	value *KeyfactorWebKeyfactorApiModelsGlobalPermissionsGlobalPermissionResponse
	isSet bool
}

func (v NullableKeyfactorWebKeyfactorApiModelsGlobalPermissionsGlobalPermissionResponse) Get() *KeyfactorWebKeyfactorApiModelsGlobalPermissionsGlobalPermissionResponse {
	return v.value
}

func (v *NullableKeyfactorWebKeyfactorApiModelsGlobalPermissionsGlobalPermissionResponse) Set(val *KeyfactorWebKeyfactorApiModelsGlobalPermissionsGlobalPermissionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorWebKeyfactorApiModelsGlobalPermissionsGlobalPermissionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorWebKeyfactorApiModelsGlobalPermissionsGlobalPermissionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorWebKeyfactorApiModelsGlobalPermissionsGlobalPermissionResponse(val *KeyfactorWebKeyfactorApiModelsGlobalPermissionsGlobalPermissionResponse) *NullableKeyfactorWebKeyfactorApiModelsGlobalPermissionsGlobalPermissionResponse {
	return &NullableKeyfactorWebKeyfactorApiModelsGlobalPermissionsGlobalPermissionResponse{value: val, isSet: true}
}

func (v NullableKeyfactorWebKeyfactorApiModelsGlobalPermissionsGlobalPermissionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorWebKeyfactorApiModelsGlobalPermissionsGlobalPermissionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
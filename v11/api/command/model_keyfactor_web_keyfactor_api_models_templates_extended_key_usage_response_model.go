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

// checks if the KeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel{}

// KeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel struct for KeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel
type KeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel struct {
	Id          *int32         `json:"id,omitempty"`
	Oid         NullableString `json:"oid,omitempty"`
	DisplayName NullableString `json:"displayName,omitempty"`
}

// NewKeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel instantiates a new KeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel() *KeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel {
	this := KeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel{}
	return &this
}

// NewKeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModelWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModelWithDefaults() *KeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel {
	this := KeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *KeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel) SetId(v int32) {
	o.Id = &v
}

// GetOid returns the Oid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel) GetOid() string {
	if o == nil || isNil(o.Oid.Get()) {
		var ret string
		return ret
	}
	return *o.Oid.Get()
}

// GetOidOk returns a tuple with the Oid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel) GetOidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Oid.Get(), o.Oid.IsSet()
}

// HasOid returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel) HasOid() bool {
	if o != nil && o.Oid.IsSet() {
		return true
	}

	return false
}

// SetOid gets a reference to the given NullableString and assigns it to the Oid field.
func (o *KeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel) SetOid(v string) {
	o.Oid.Set(&v)
}

// SetOidNil sets the value for Oid to be an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel) SetOidNil() {
	o.Oid.Set(nil)
}

// UnsetOid ensures that no value is present for Oid, not even an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel) UnsetOid() {
	o.Oid.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName.Get()) {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *KeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}

// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel) UnsetDisplayName() {
	o.DisplayName.Unset()
}

func (o KeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Oid.IsSet() {
		toSerialize["oid"] = o.Oid.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	return toSerialize, nil
}

type NullableKeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel struct {
	value *KeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel
	isSet bool
}

func (v NullableKeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel) Get() *KeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel {
	return v.value
}

func (v *NullableKeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel) Set(val *KeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel(val *KeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel) *NullableKeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel {
	return &NullableKeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel{value: val, isSet: true}
}

func (v NullableKeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorWebKeyfactorApiModelsTemplatesExtendedKeyUsageResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
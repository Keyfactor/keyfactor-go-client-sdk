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

// checks if the CSSCMSDataModelModelsProviderType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CSSCMSDataModelModelsProviderType{}

// CSSCMSDataModelModelsProviderType struct for CSSCMSDataModelModelsProviderType
type CSSCMSDataModelModelsProviderType struct {
	Id                   *string                                  `json:"Id,omitempty"`
	Name                 *string                                  `json:"Name,omitempty"`
	ProviderTypeParams   []CSSCMSDataModelModelsProviderTypeParam `json:"ProviderTypeParams,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CSSCMSDataModelModelsProviderType CSSCMSDataModelModelsProviderType

// NewCSSCMSDataModelModelsProviderType instantiates a new CSSCMSDataModelModelsProviderType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCSSCMSDataModelModelsProviderType() *CSSCMSDataModelModelsProviderType {
	this := CSSCMSDataModelModelsProviderType{}
	return &this
}

// NewCSSCMSDataModelModelsProviderTypeWithDefaults instantiates a new CSSCMSDataModelModelsProviderType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCSSCMSDataModelModelsProviderTypeWithDefaults() *CSSCMSDataModelModelsProviderType {
	this := CSSCMSDataModelModelsProviderType{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CSSCMSDataModelModelsProviderType) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSSCMSDataModelModelsProviderType) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CSSCMSDataModelModelsProviderType) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CSSCMSDataModelModelsProviderType) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CSSCMSDataModelModelsProviderType) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSSCMSDataModelModelsProviderType) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CSSCMSDataModelModelsProviderType) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CSSCMSDataModelModelsProviderType) SetName(v string) {
	o.Name = &v
}

// GetProviderTypeParams returns the ProviderTypeParams field value if set, zero value otherwise.
func (o *CSSCMSDataModelModelsProviderType) GetProviderTypeParams() []CSSCMSDataModelModelsProviderTypeParam {
	if o == nil || isNil(o.ProviderTypeParams) {
		var ret []CSSCMSDataModelModelsProviderTypeParam
		return ret
	}
	return o.ProviderTypeParams
}

// GetProviderTypeParamsOk returns a tuple with the ProviderTypeParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSSCMSDataModelModelsProviderType) GetProviderTypeParamsOk() ([]CSSCMSDataModelModelsProviderTypeParam, bool) {
	if o == nil || isNil(o.ProviderTypeParams) {
		return nil, false
	}
	return o.ProviderTypeParams, true
}

// HasProviderTypeParams returns a boolean if a field has been set.
func (o *CSSCMSDataModelModelsProviderType) HasProviderTypeParams() bool {
	if o != nil && !isNil(o.ProviderTypeParams) {
		return true
	}

	return false
}

// SetProviderTypeParams gets a reference to the given []CSSCMSDataModelModelsProviderTypeParam and assigns it to the ProviderTypeParams field.
func (o *CSSCMSDataModelModelsProviderType) SetProviderTypeParams(v []CSSCMSDataModelModelsProviderTypeParam) {
	o.ProviderTypeParams = v
}

func (o CSSCMSDataModelModelsProviderType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CSSCMSDataModelModelsProviderType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !isNil(o.ProviderTypeParams) {
		toSerialize["ProviderTypeParams"] = o.ProviderTypeParams
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CSSCMSDataModelModelsProviderType) UnmarshalJSON(bytes []byte) (err error) {
	varCSSCMSDataModelModelsProviderType := _CSSCMSDataModelModelsProviderType{}

	if err = json.Unmarshal(bytes, &varCSSCMSDataModelModelsProviderType); err == nil {
		*o = CSSCMSDataModelModelsProviderType(varCSSCMSDataModelModelsProviderType)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Id")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "ProviderTypeParams")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCSSCMSDataModelModelsProviderType struct {
	value *CSSCMSDataModelModelsProviderType
	isSet bool
}

func (v NullableCSSCMSDataModelModelsProviderType) Get() *CSSCMSDataModelModelsProviderType {
	return v.value
}

func (v *NullableCSSCMSDataModelModelsProviderType) Set(val *CSSCMSDataModelModelsProviderType) {
	v.value = val
	v.isSet = true
}

func (v NullableCSSCMSDataModelModelsProviderType) IsSet() bool {
	return v.isSet
}

func (v *NullableCSSCMSDataModelModelsProviderType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCSSCMSDataModelModelsProviderType(val *CSSCMSDataModelModelsProviderType) *NullableCSSCMSDataModelModelsProviderType {
	return &NullableCSSCMSDataModelModelsProviderType{value: val, isSet: true}
}

func (v NullableCSSCMSDataModelModelsProviderType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCSSCMSDataModelModelsProviderType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
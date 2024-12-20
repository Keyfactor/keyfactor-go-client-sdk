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

// checks if the KeyfactorApiModelsWorkflowsInstanceDefinitionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorApiModelsWorkflowsInstanceDefinitionResponse{}

// KeyfactorApiModelsWorkflowsInstanceDefinitionResponse struct for KeyfactorApiModelsWorkflowsInstanceDefinitionResponse
type KeyfactorApiModelsWorkflowsInstanceDefinitionResponse struct {
	Id                   *string `json:"Id,omitempty"`
	DisplayName          *string `json:"DisplayName,omitempty"`
	Version              *int32  `json:"Version,omitempty"`
	WorkflowType         *string `json:"WorkflowType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KeyfactorApiModelsWorkflowsInstanceDefinitionResponse KeyfactorApiModelsWorkflowsInstanceDefinitionResponse

// NewKeyfactorApiModelsWorkflowsInstanceDefinitionResponse instantiates a new KeyfactorApiModelsWorkflowsInstanceDefinitionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorApiModelsWorkflowsInstanceDefinitionResponse() *KeyfactorApiModelsWorkflowsInstanceDefinitionResponse {
	this := KeyfactorApiModelsWorkflowsInstanceDefinitionResponse{}
	return &this
}

// NewKeyfactorApiModelsWorkflowsInstanceDefinitionResponseWithDefaults instantiates a new KeyfactorApiModelsWorkflowsInstanceDefinitionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorApiModelsWorkflowsInstanceDefinitionResponseWithDefaults() *KeyfactorApiModelsWorkflowsInstanceDefinitionResponse {
	this := KeyfactorApiModelsWorkflowsInstanceDefinitionResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KeyfactorApiModelsWorkflowsInstanceDefinitionResponse) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsWorkflowsInstanceDefinitionResponse) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KeyfactorApiModelsWorkflowsInstanceDefinitionResponse) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *KeyfactorApiModelsWorkflowsInstanceDefinitionResponse) SetId(v string) {
	o.Id = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *KeyfactorApiModelsWorkflowsInstanceDefinitionResponse) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsWorkflowsInstanceDefinitionResponse) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *KeyfactorApiModelsWorkflowsInstanceDefinitionResponse) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *KeyfactorApiModelsWorkflowsInstanceDefinitionResponse) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *KeyfactorApiModelsWorkflowsInstanceDefinitionResponse) GetVersion() int32 {
	if o == nil || isNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsWorkflowsInstanceDefinitionResponse) GetVersionOk() (*int32, bool) {
	if o == nil || isNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *KeyfactorApiModelsWorkflowsInstanceDefinitionResponse) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *KeyfactorApiModelsWorkflowsInstanceDefinitionResponse) SetVersion(v int32) {
	o.Version = &v
}

// GetWorkflowType returns the WorkflowType field value if set, zero value otherwise.
func (o *KeyfactorApiModelsWorkflowsInstanceDefinitionResponse) GetWorkflowType() string {
	if o == nil || isNil(o.WorkflowType) {
		var ret string
		return ret
	}
	return *o.WorkflowType
}

// GetWorkflowTypeOk returns a tuple with the WorkflowType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsWorkflowsInstanceDefinitionResponse) GetWorkflowTypeOk() (*string, bool) {
	if o == nil || isNil(o.WorkflowType) {
		return nil, false
	}
	return o.WorkflowType, true
}

// HasWorkflowType returns a boolean if a field has been set.
func (o *KeyfactorApiModelsWorkflowsInstanceDefinitionResponse) HasWorkflowType() bool {
	if o != nil && !isNil(o.WorkflowType) {
		return true
	}

	return false
}

// SetWorkflowType gets a reference to the given string and assigns it to the WorkflowType field.
func (o *KeyfactorApiModelsWorkflowsInstanceDefinitionResponse) SetWorkflowType(v string) {
	o.WorkflowType = &v
}

func (o KeyfactorApiModelsWorkflowsInstanceDefinitionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorApiModelsWorkflowsInstanceDefinitionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !isNil(o.DisplayName) {
		toSerialize["DisplayName"] = o.DisplayName
	}
	if !isNil(o.Version) {
		toSerialize["Version"] = o.Version
	}
	if !isNil(o.WorkflowType) {
		toSerialize["WorkflowType"] = o.WorkflowType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KeyfactorApiModelsWorkflowsInstanceDefinitionResponse) UnmarshalJSON(bytes []byte) (err error) {
	varKeyfactorApiModelsWorkflowsInstanceDefinitionResponse := _KeyfactorApiModelsWorkflowsInstanceDefinitionResponse{}

	if err = json.Unmarshal(bytes, &varKeyfactorApiModelsWorkflowsInstanceDefinitionResponse); err == nil {
		*o = KeyfactorApiModelsWorkflowsInstanceDefinitionResponse(varKeyfactorApiModelsWorkflowsInstanceDefinitionResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Id")
		delete(additionalProperties, "DisplayName")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "WorkflowType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKeyfactorApiModelsWorkflowsInstanceDefinitionResponse struct {
	value *KeyfactorApiModelsWorkflowsInstanceDefinitionResponse
	isSet bool
}

func (v NullableKeyfactorApiModelsWorkflowsInstanceDefinitionResponse) Get() *KeyfactorApiModelsWorkflowsInstanceDefinitionResponse {
	return v.value
}

func (v *NullableKeyfactorApiModelsWorkflowsInstanceDefinitionResponse) Set(val *KeyfactorApiModelsWorkflowsInstanceDefinitionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorApiModelsWorkflowsInstanceDefinitionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorApiModelsWorkflowsInstanceDefinitionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorApiModelsWorkflowsInstanceDefinitionResponse(val *KeyfactorApiModelsWorkflowsInstanceDefinitionResponse) *NullableKeyfactorApiModelsWorkflowsInstanceDefinitionResponse {
	return &NullableKeyfactorApiModelsWorkflowsInstanceDefinitionResponse{value: val, isSet: true}
}

func (v NullableKeyfactorApiModelsWorkflowsInstanceDefinitionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorApiModelsWorkflowsInstanceDefinitionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

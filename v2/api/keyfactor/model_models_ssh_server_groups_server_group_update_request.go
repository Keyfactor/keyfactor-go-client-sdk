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

// checks if the ModelsSSHServerGroupsServerGroupUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsSSHServerGroupsServerGroupUpdateRequest{}

// ModelsSSHServerGroupsServerGroupUpdateRequest struct for ModelsSSHServerGroupsServerGroupUpdateRequest
type ModelsSSHServerGroupsServerGroupUpdateRequest struct {
	Id                   string                                      `json:"Id"`
	OwnerName            string                                      `json:"OwnerName"`
	GroupName            string                                      `json:"GroupName"`
	SyncSchedule         *KeyfactorCommonSchedulingKeyfactorSchedule `json:"SyncSchedule,omitempty"`
	UnderManagement      bool                                        `json:"UnderManagement"`
	AdditionalProperties map[string]interface{}
}

type _ModelsSSHServerGroupsServerGroupUpdateRequest ModelsSSHServerGroupsServerGroupUpdateRequest

// NewModelsSSHServerGroupsServerGroupUpdateRequest instantiates a new ModelsSSHServerGroupsServerGroupUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsSSHServerGroupsServerGroupUpdateRequest(id string, ownerName string, groupName string, underManagement bool) *ModelsSSHServerGroupsServerGroupUpdateRequest {
	this := ModelsSSHServerGroupsServerGroupUpdateRequest{}
	this.Id = id
	this.OwnerName = ownerName
	this.GroupName = groupName
	this.UnderManagement = underManagement
	return &this
}

// NewModelsSSHServerGroupsServerGroupUpdateRequestWithDefaults instantiates a new ModelsSSHServerGroupsServerGroupUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsSSHServerGroupsServerGroupUpdateRequestWithDefaults() *ModelsSSHServerGroupsServerGroupUpdateRequest {
	this := ModelsSSHServerGroupsServerGroupUpdateRequest{}
	return &this
}

// GetId returns the Id field value
func (o *ModelsSSHServerGroupsServerGroupUpdateRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ModelsSSHServerGroupsServerGroupUpdateRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ModelsSSHServerGroupsServerGroupUpdateRequest) SetId(v string) {
	o.Id = v
}

// GetOwnerName returns the OwnerName field value
func (o *ModelsSSHServerGroupsServerGroupUpdateRequest) GetOwnerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerName
}

// GetOwnerNameOk returns a tuple with the OwnerName field value
// and a boolean to check if the value has been set.
func (o *ModelsSSHServerGroupsServerGroupUpdateRequest) GetOwnerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerName, true
}

// SetOwnerName sets field value
func (o *ModelsSSHServerGroupsServerGroupUpdateRequest) SetOwnerName(v string) {
	o.OwnerName = v
}

// GetGroupName returns the GroupName field value
func (o *ModelsSSHServerGroupsServerGroupUpdateRequest) GetGroupName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value
// and a boolean to check if the value has been set.
func (o *ModelsSSHServerGroupsServerGroupUpdateRequest) GetGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupName, true
}

// SetGroupName sets field value
func (o *ModelsSSHServerGroupsServerGroupUpdateRequest) SetGroupName(v string) {
	o.GroupName = v
}

// GetSyncSchedule returns the SyncSchedule field value if set, zero value otherwise.
func (o *ModelsSSHServerGroupsServerGroupUpdateRequest) GetSyncSchedule() KeyfactorCommonSchedulingKeyfactorSchedule {
	if o == nil || isNil(o.SyncSchedule) {
		var ret KeyfactorCommonSchedulingKeyfactorSchedule
		return ret
	}
	return *o.SyncSchedule
}

// GetSyncScheduleOk returns a tuple with the SyncSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSSHServerGroupsServerGroupUpdateRequest) GetSyncScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool) {
	if o == nil || isNil(o.SyncSchedule) {
		return nil, false
	}
	return o.SyncSchedule, true
}

// HasSyncSchedule returns a boolean if a field has been set.
func (o *ModelsSSHServerGroupsServerGroupUpdateRequest) HasSyncSchedule() bool {
	if o != nil && !isNil(o.SyncSchedule) {
		return true
	}

	return false
}

// SetSyncSchedule gets a reference to the given KeyfactorCommonSchedulingKeyfactorSchedule and assigns it to the SyncSchedule field.
func (o *ModelsSSHServerGroupsServerGroupUpdateRequest) SetSyncSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule) {
	o.SyncSchedule = &v
}

// GetUnderManagement returns the UnderManagement field value
func (o *ModelsSSHServerGroupsServerGroupUpdateRequest) GetUnderManagement() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UnderManagement
}

// GetUnderManagementOk returns a tuple with the UnderManagement field value
// and a boolean to check if the value has been set.
func (o *ModelsSSHServerGroupsServerGroupUpdateRequest) GetUnderManagementOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnderManagement, true
}

// SetUnderManagement sets field value
func (o *ModelsSSHServerGroupsServerGroupUpdateRequest) SetUnderManagement(v bool) {
	o.UnderManagement = v
}

func (o ModelsSSHServerGroupsServerGroupUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsSSHServerGroupsServerGroupUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Id"] = o.Id
	toSerialize["OwnerName"] = o.OwnerName
	toSerialize["GroupName"] = o.GroupName
	if !isNil(o.SyncSchedule) {
		toSerialize["SyncSchedule"] = o.SyncSchedule
	}
	toSerialize["UnderManagement"] = o.UnderManagement

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModelsSSHServerGroupsServerGroupUpdateRequest) UnmarshalJSON(bytes []byte) (err error) {
	varModelsSSHServerGroupsServerGroupUpdateRequest := _ModelsSSHServerGroupsServerGroupUpdateRequest{}

	if err = json.Unmarshal(bytes, &varModelsSSHServerGroupsServerGroupUpdateRequest); err == nil {
		*o = ModelsSSHServerGroupsServerGroupUpdateRequest(varModelsSSHServerGroupsServerGroupUpdateRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Id")
		delete(additionalProperties, "OwnerName")
		delete(additionalProperties, "GroupName")
		delete(additionalProperties, "SyncSchedule")
		delete(additionalProperties, "UnderManagement")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelsSSHServerGroupsServerGroupUpdateRequest struct {
	value *ModelsSSHServerGroupsServerGroupUpdateRequest
	isSet bool
}

func (v NullableModelsSSHServerGroupsServerGroupUpdateRequest) Get() *ModelsSSHServerGroupsServerGroupUpdateRequest {
	return v.value
}

func (v *NullableModelsSSHServerGroupsServerGroupUpdateRequest) Set(val *ModelsSSHServerGroupsServerGroupUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsSSHServerGroupsServerGroupUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsSSHServerGroupsServerGroupUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsSSHServerGroupsServerGroupUpdateRequest(val *ModelsSSHServerGroupsServerGroupUpdateRequest) *NullableModelsSSHServerGroupsServerGroupUpdateRequest {
	return &NullableModelsSSHServerGroupsServerGroupUpdateRequest{value: val, isSet: true}
}

func (v NullableModelsSSHServerGroupsServerGroupUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsSSHServerGroupsServerGroupUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

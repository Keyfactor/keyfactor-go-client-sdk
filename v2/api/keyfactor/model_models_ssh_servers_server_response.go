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

// checks if the ModelsSSHServersServerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsSSHServersServerResponse{}

// ModelsSSHServersServerResponse struct for ModelsSSHServersServerResponse
type ModelsSSHServersServerResponse struct {
	Id                   *int32                                      `json:"Id,omitempty"`
	AgentId              *string                                     `json:"AgentId,omitempty"`
	Hostname             *string                                     `json:"Hostname,omitempty"`
	ServerGroupId        *string                                     `json:"ServerGroupId,omitempty"`
	SyncSchedule         *KeyfactorCommonSchedulingKeyfactorSchedule `json:"SyncSchedule,omitempty"`
	UnderManagement      *bool                                       `json:"UnderManagement,omitempty"`
	Owner                *ModelsSSHUsersSshUserResponse              `json:"Owner,omitempty"`
	GroupName            *string                                     `json:"GroupName,omitempty"`
	Orchestrator         *string                                     `json:"Orchestrator,omitempty"`
	Port                 *int32                                      `json:"Port,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelsSSHServersServerResponse ModelsSSHServersServerResponse

// NewModelsSSHServersServerResponse instantiates a new ModelsSSHServersServerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsSSHServersServerResponse() *ModelsSSHServersServerResponse {
	this := ModelsSSHServersServerResponse{}
	return &this
}

// NewModelsSSHServersServerResponseWithDefaults instantiates a new ModelsSSHServersServerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsSSHServersServerResponseWithDefaults() *ModelsSSHServersServerResponse {
	this := ModelsSSHServersServerResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelsSSHServersServerResponse) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSSHServersServerResponse) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelsSSHServersServerResponse) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ModelsSSHServersServerResponse) SetId(v int32) {
	o.Id = &v
}

// GetAgentId returns the AgentId field value if set, zero value otherwise.
func (o *ModelsSSHServersServerResponse) GetAgentId() string {
	if o == nil || isNil(o.AgentId) {
		var ret string
		return ret
	}
	return *o.AgentId
}

// GetAgentIdOk returns a tuple with the AgentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSSHServersServerResponse) GetAgentIdOk() (*string, bool) {
	if o == nil || isNil(o.AgentId) {
		return nil, false
	}
	return o.AgentId, true
}

// HasAgentId returns a boolean if a field has been set.
func (o *ModelsSSHServersServerResponse) HasAgentId() bool {
	if o != nil && !isNil(o.AgentId) {
		return true
	}

	return false
}

// SetAgentId gets a reference to the given string and assigns it to the AgentId field.
func (o *ModelsSSHServersServerResponse) SetAgentId(v string) {
	o.AgentId = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *ModelsSSHServersServerResponse) GetHostname() string {
	if o == nil || isNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSSHServersServerResponse) GetHostnameOk() (*string, bool) {
	if o == nil || isNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *ModelsSSHServersServerResponse) HasHostname() bool {
	if o != nil && !isNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *ModelsSSHServersServerResponse) SetHostname(v string) {
	o.Hostname = &v
}

// GetServerGroupId returns the ServerGroupId field value if set, zero value otherwise.
func (o *ModelsSSHServersServerResponse) GetServerGroupId() string {
	if o == nil || isNil(o.ServerGroupId) {
		var ret string
		return ret
	}
	return *o.ServerGroupId
}

// GetServerGroupIdOk returns a tuple with the ServerGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSSHServersServerResponse) GetServerGroupIdOk() (*string, bool) {
	if o == nil || isNil(o.ServerGroupId) {
		return nil, false
	}
	return o.ServerGroupId, true
}

// HasServerGroupId returns a boolean if a field has been set.
func (o *ModelsSSHServersServerResponse) HasServerGroupId() bool {
	if o != nil && !isNil(o.ServerGroupId) {
		return true
	}

	return false
}

// SetServerGroupId gets a reference to the given string and assigns it to the ServerGroupId field.
func (o *ModelsSSHServersServerResponse) SetServerGroupId(v string) {
	o.ServerGroupId = &v
}

// GetSyncSchedule returns the SyncSchedule field value if set, zero value otherwise.
func (o *ModelsSSHServersServerResponse) GetSyncSchedule() KeyfactorCommonSchedulingKeyfactorSchedule {
	if o == nil || isNil(o.SyncSchedule) {
		var ret KeyfactorCommonSchedulingKeyfactorSchedule
		return ret
	}
	return *o.SyncSchedule
}

// GetSyncScheduleOk returns a tuple with the SyncSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSSHServersServerResponse) GetSyncScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool) {
	if o == nil || isNil(o.SyncSchedule) {
		return nil, false
	}
	return o.SyncSchedule, true
}

// HasSyncSchedule returns a boolean if a field has been set.
func (o *ModelsSSHServersServerResponse) HasSyncSchedule() bool {
	if o != nil && !isNil(o.SyncSchedule) {
		return true
	}

	return false
}

// SetSyncSchedule gets a reference to the given KeyfactorCommonSchedulingKeyfactorSchedule and assigns it to the SyncSchedule field.
func (o *ModelsSSHServersServerResponse) SetSyncSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule) {
	o.SyncSchedule = &v
}

// GetUnderManagement returns the UnderManagement field value if set, zero value otherwise.
func (o *ModelsSSHServersServerResponse) GetUnderManagement() bool {
	if o == nil || isNil(o.UnderManagement) {
		var ret bool
		return ret
	}
	return *o.UnderManagement
}

// GetUnderManagementOk returns a tuple with the UnderManagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSSHServersServerResponse) GetUnderManagementOk() (*bool, bool) {
	if o == nil || isNil(o.UnderManagement) {
		return nil, false
	}
	return o.UnderManagement, true
}

// HasUnderManagement returns a boolean if a field has been set.
func (o *ModelsSSHServersServerResponse) HasUnderManagement() bool {
	if o != nil && !isNil(o.UnderManagement) {
		return true
	}

	return false
}

// SetUnderManagement gets a reference to the given bool and assigns it to the UnderManagement field.
func (o *ModelsSSHServersServerResponse) SetUnderManagement(v bool) {
	o.UnderManagement = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *ModelsSSHServersServerResponse) GetOwner() ModelsSSHUsersSshUserResponse {
	if o == nil || isNil(o.Owner) {
		var ret ModelsSSHUsersSshUserResponse
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSSHServersServerResponse) GetOwnerOk() (*ModelsSSHUsersSshUserResponse, bool) {
	if o == nil || isNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *ModelsSSHServersServerResponse) HasOwner() bool {
	if o != nil && !isNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given ModelsSSHUsersSshUserResponse and assigns it to the Owner field.
func (o *ModelsSSHServersServerResponse) SetOwner(v ModelsSSHUsersSshUserResponse) {
	o.Owner = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *ModelsSSHServersServerResponse) GetGroupName() string {
	if o == nil || isNil(o.GroupName) {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSSHServersServerResponse) GetGroupNameOk() (*string, bool) {
	if o == nil || isNil(o.GroupName) {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *ModelsSSHServersServerResponse) HasGroupName() bool {
	if o != nil && !isNil(o.GroupName) {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *ModelsSSHServersServerResponse) SetGroupName(v string) {
	o.GroupName = &v
}

// GetOrchestrator returns the Orchestrator field value if set, zero value otherwise.
func (o *ModelsSSHServersServerResponse) GetOrchestrator() string {
	if o == nil || isNil(o.Orchestrator) {
		var ret string
		return ret
	}
	return *o.Orchestrator
}

// GetOrchestratorOk returns a tuple with the Orchestrator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSSHServersServerResponse) GetOrchestratorOk() (*string, bool) {
	if o == nil || isNil(o.Orchestrator) {
		return nil, false
	}
	return o.Orchestrator, true
}

// HasOrchestrator returns a boolean if a field has been set.
func (o *ModelsSSHServersServerResponse) HasOrchestrator() bool {
	if o != nil && !isNil(o.Orchestrator) {
		return true
	}

	return false
}

// SetOrchestrator gets a reference to the given string and assigns it to the Orchestrator field.
func (o *ModelsSSHServersServerResponse) SetOrchestrator(v string) {
	o.Orchestrator = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *ModelsSSHServersServerResponse) GetPort() int32 {
	if o == nil || isNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSSHServersServerResponse) GetPortOk() (*int32, bool) {
	if o == nil || isNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *ModelsSSHServersServerResponse) HasPort() bool {
	if o != nil && !isNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *ModelsSSHServersServerResponse) SetPort(v int32) {
	o.Port = &v
}

func (o ModelsSSHServersServerResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsSSHServersServerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !isNil(o.AgentId) {
		toSerialize["AgentId"] = o.AgentId
	}
	if !isNil(o.Hostname) {
		toSerialize["Hostname"] = o.Hostname
	}
	if !isNil(o.ServerGroupId) {
		toSerialize["ServerGroupId"] = o.ServerGroupId
	}
	if !isNil(o.SyncSchedule) {
		toSerialize["SyncSchedule"] = o.SyncSchedule
	}
	if !isNil(o.UnderManagement) {
		toSerialize["UnderManagement"] = o.UnderManagement
	}
	if !isNil(o.Owner) {
		toSerialize["Owner"] = o.Owner
	}
	if !isNil(o.GroupName) {
		toSerialize["GroupName"] = o.GroupName
	}
	if !isNil(o.Orchestrator) {
		toSerialize["Orchestrator"] = o.Orchestrator
	}
	if !isNil(o.Port) {
		toSerialize["Port"] = o.Port
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModelsSSHServersServerResponse) UnmarshalJSON(bytes []byte) (err error) {
	varModelsSSHServersServerResponse := _ModelsSSHServersServerResponse{}

	if err = json.Unmarshal(bytes, &varModelsSSHServersServerResponse); err == nil {
		*o = ModelsSSHServersServerResponse(varModelsSSHServersServerResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Id")
		delete(additionalProperties, "AgentId")
		delete(additionalProperties, "Hostname")
		delete(additionalProperties, "ServerGroupId")
		delete(additionalProperties, "SyncSchedule")
		delete(additionalProperties, "UnderManagement")
		delete(additionalProperties, "Owner")
		delete(additionalProperties, "GroupName")
		delete(additionalProperties, "Orchestrator")
		delete(additionalProperties, "Port")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelsSSHServersServerResponse struct {
	value *ModelsSSHServersServerResponse
	isSet bool
}

func (v NullableModelsSSHServersServerResponse) Get() *ModelsSSHServersServerResponse {
	return v.value
}

func (v *NullableModelsSSHServersServerResponse) Set(val *ModelsSSHServersServerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsSSHServersServerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsSSHServersServerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsSSHServersServerResponse(val *ModelsSSHServersServerResponse) *NullableModelsSSHServersServerResponse {
	return &NullableModelsSSHServersServerResponse{value: val, isSet: true}
}

func (v NullableModelsSSHServersServerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsSSHServersServerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

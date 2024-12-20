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

// checks if the ModelsSSHServersServerCreationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsSSHServersServerCreationRequest{}

// ModelsSSHServersServerCreationRequest struct for ModelsSSHServersServerCreationRequest
type ModelsSSHServersServerCreationRequest struct {
	AgentId              string `json:"AgentId"`
	Hostname             string `json:"Hostname"`
	ServerGroupId        string `json:"ServerGroupId"`
	UnderManagement      *bool  `json:"UnderManagement,omitempty"`
	Port                 *int32 `json:"Port,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelsSSHServersServerCreationRequest ModelsSSHServersServerCreationRequest

// NewModelsSSHServersServerCreationRequest instantiates a new ModelsSSHServersServerCreationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsSSHServersServerCreationRequest(agentId string, hostname string, serverGroupId string) *ModelsSSHServersServerCreationRequest {
	this := ModelsSSHServersServerCreationRequest{}
	this.AgentId = agentId
	this.Hostname = hostname
	this.ServerGroupId = serverGroupId
	return &this
}

// NewModelsSSHServersServerCreationRequestWithDefaults instantiates a new ModelsSSHServersServerCreationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsSSHServersServerCreationRequestWithDefaults() *ModelsSSHServersServerCreationRequest {
	this := ModelsSSHServersServerCreationRequest{}
	return &this
}

// GetAgentId returns the AgentId field value
func (o *ModelsSSHServersServerCreationRequest) GetAgentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AgentId
}

// GetAgentIdOk returns a tuple with the AgentId field value
// and a boolean to check if the value has been set.
func (o *ModelsSSHServersServerCreationRequest) GetAgentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentId, true
}

// SetAgentId sets field value
func (o *ModelsSSHServersServerCreationRequest) SetAgentId(v string) {
	o.AgentId = v
}

// GetHostname returns the Hostname field value
func (o *ModelsSSHServersServerCreationRequest) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *ModelsSSHServersServerCreationRequest) GetHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *ModelsSSHServersServerCreationRequest) SetHostname(v string) {
	o.Hostname = v
}

// GetServerGroupId returns the ServerGroupId field value
func (o *ModelsSSHServersServerCreationRequest) GetServerGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerGroupId
}

// GetServerGroupIdOk returns a tuple with the ServerGroupId field value
// and a boolean to check if the value has been set.
func (o *ModelsSSHServersServerCreationRequest) GetServerGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerGroupId, true
}

// SetServerGroupId sets field value
func (o *ModelsSSHServersServerCreationRequest) SetServerGroupId(v string) {
	o.ServerGroupId = v
}

// GetUnderManagement returns the UnderManagement field value if set, zero value otherwise.
func (o *ModelsSSHServersServerCreationRequest) GetUnderManagement() bool {
	if o == nil || isNil(o.UnderManagement) {
		var ret bool
		return ret
	}
	return *o.UnderManagement
}

// GetUnderManagementOk returns a tuple with the UnderManagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSSHServersServerCreationRequest) GetUnderManagementOk() (*bool, bool) {
	if o == nil || isNil(o.UnderManagement) {
		return nil, false
	}
	return o.UnderManagement, true
}

// HasUnderManagement returns a boolean if a field has been set.
func (o *ModelsSSHServersServerCreationRequest) HasUnderManagement() bool {
	if o != nil && !isNil(o.UnderManagement) {
		return true
	}

	return false
}

// SetUnderManagement gets a reference to the given bool and assigns it to the UnderManagement field.
func (o *ModelsSSHServersServerCreationRequest) SetUnderManagement(v bool) {
	o.UnderManagement = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *ModelsSSHServersServerCreationRequest) GetPort() int32 {
	if o == nil || isNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSSHServersServerCreationRequest) GetPortOk() (*int32, bool) {
	if o == nil || isNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *ModelsSSHServersServerCreationRequest) HasPort() bool {
	if o != nil && !isNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *ModelsSSHServersServerCreationRequest) SetPort(v int32) {
	o.Port = &v
}

func (o ModelsSSHServersServerCreationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsSSHServersServerCreationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["AgentId"] = o.AgentId
	toSerialize["Hostname"] = o.Hostname
	toSerialize["ServerGroupId"] = o.ServerGroupId
	if !isNil(o.UnderManagement) {
		toSerialize["UnderManagement"] = o.UnderManagement
	}
	if !isNil(o.Port) {
		toSerialize["Port"] = o.Port
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModelsSSHServersServerCreationRequest) UnmarshalJSON(bytes []byte) (err error) {
	varModelsSSHServersServerCreationRequest := _ModelsSSHServersServerCreationRequest{}

	if err = json.Unmarshal(bytes, &varModelsSSHServersServerCreationRequest); err == nil {
		*o = ModelsSSHServersServerCreationRequest(varModelsSSHServersServerCreationRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "AgentId")
		delete(additionalProperties, "Hostname")
		delete(additionalProperties, "ServerGroupId")
		delete(additionalProperties, "UnderManagement")
		delete(additionalProperties, "Port")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelsSSHServersServerCreationRequest struct {
	value *ModelsSSHServersServerCreationRequest
	isSet bool
}

func (v NullableModelsSSHServersServerCreationRequest) Get() *ModelsSSHServersServerCreationRequest {
	return v.value
}

func (v *NullableModelsSSHServersServerCreationRequest) Set(val *ModelsSSHServersServerCreationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsSSHServersServerCreationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsSSHServersServerCreationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsSSHServersServerCreationRequest(val *ModelsSSHServersServerCreationRequest) *NullableModelsSSHServersServerCreationRequest {
	return &NullableModelsSSHServersServerCreationRequest{value: val, isSet: true}
}

func (v NullableModelsSSHServersServerCreationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsSSHServersServerCreationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

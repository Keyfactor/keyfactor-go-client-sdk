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

Keyfactor-v1

This reference serves to document REST-based methods to manage and integrate with Keyfactor. In addition, an embedded interface allows for the execution of calls against the current Keyfactor API instance.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package command

import (
	"encoding/json"
)

// checks if the ModelsCertificateStoresCertificateStoreCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsCertificateStoresCertificateStoreCreateRequest{}

// ModelsCertificateStoresCertificateStoreCreateRequest struct for ModelsCertificateStoresCertificateStoreCreateRequest
type ModelsCertificateStoresCertificateStoreCreateRequest struct {
	ContainerId          *int32                                      `json:"ContainerId,omitempty"`
	ClientMachine        *string                                     `json:"ClientMachine,omitempty"`
	Storepath            *string                                     `json:"Storepath,omitempty"`
	CertStoreType        *int32                                      `json:"CertStoreType,omitempty"`
	CreateIfMissing      *bool                                       `json:"CreateIfMissing,omitempty"`
	Properties           *string                                     `json:"Properties,omitempty"`
	AgentId              *string                                     `json:"AgentId,omitempty"`
	AgentAssigned        *bool                                       `json:"AgentAssigned,omitempty"`
	InventorySchedule    *KeyfactorCommonSchedulingKeyfactorSchedule `json:"InventorySchedule,omitempty"`
	Password             *ModelsKeyfactorAPISecret                   `json:"Password,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelsCertificateStoresCertificateStoreCreateRequest ModelsCertificateStoresCertificateStoreCreateRequest

// NewModelsCertificateStoresCertificateStoreCreateRequest instantiates a new ModelsCertificateStoresCertificateStoreCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsCertificateStoresCertificateStoreCreateRequest() *ModelsCertificateStoresCertificateStoreCreateRequest {
	this := ModelsCertificateStoresCertificateStoreCreateRequest{}
	return &this
}

// NewModelsCertificateStoresCertificateStoreCreateRequestWithDefaults instantiates a new ModelsCertificateStoresCertificateStoreCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsCertificateStoresCertificateStoreCreateRequestWithDefaults() *ModelsCertificateStoresCertificateStoreCreateRequest {
	this := ModelsCertificateStoresCertificateStoreCreateRequest{}
	return &this
}

// GetContainerId returns the ContainerId field value if set, zero value otherwise.
func (o *ModelsCertificateStoresCertificateStoreCreateRequest) GetContainerId() int32 {
	if o == nil || isNil(o.ContainerId) {
		var ret int32
		return ret
	}
	return *o.ContainerId
}

// GetContainerIdOk returns a tuple with the ContainerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertificateStoresCertificateStoreCreateRequest) GetContainerIdOk() (*int32, bool) {
	if o == nil || isNil(o.ContainerId) {
		return nil, false
	}
	return o.ContainerId, true
}

// HasContainerId returns a boolean if a field has been set.
func (o *ModelsCertificateStoresCertificateStoreCreateRequest) HasContainerId() bool {
	if o != nil && !isNil(o.ContainerId) {
		return true
	}

	return false
}

// SetContainerId gets a reference to the given int32 and assigns it to the ContainerId field.
func (o *ModelsCertificateStoresCertificateStoreCreateRequest) SetContainerId(v int32) {
	o.ContainerId = &v
}

// GetClientMachine returns the ClientMachine field value if set, zero value otherwise.
func (o *ModelsCertificateStoresCertificateStoreCreateRequest) GetClientMachine() string {
	if o == nil || isNil(o.ClientMachine) {
		var ret string
		return ret
	}
	return *o.ClientMachine
}

// GetClientMachineOk returns a tuple with the ClientMachine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertificateStoresCertificateStoreCreateRequest) GetClientMachineOk() (*string, bool) {
	if o == nil || isNil(o.ClientMachine) {
		return nil, false
	}
	return o.ClientMachine, true
}

// HasClientMachine returns a boolean if a field has been set.
func (o *ModelsCertificateStoresCertificateStoreCreateRequest) HasClientMachine() bool {
	if o != nil && !isNil(o.ClientMachine) {
		return true
	}

	return false
}

// SetClientMachine gets a reference to the given string and assigns it to the ClientMachine field.
func (o *ModelsCertificateStoresCertificateStoreCreateRequest) SetClientMachine(v string) {
	o.ClientMachine = &v
}

// GetStorepath returns the Storepath field value if set, zero value otherwise.
func (o *ModelsCertificateStoresCertificateStoreCreateRequest) GetStorepath() string {
	if o == nil || isNil(o.Storepath) {
		var ret string
		return ret
	}
	return *o.Storepath
}

// GetStorepathOk returns a tuple with the Storepath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertificateStoresCertificateStoreCreateRequest) GetStorepathOk() (*string, bool) {
	if o == nil || isNil(o.Storepath) {
		return nil, false
	}
	return o.Storepath, true
}

// HasStorepath returns a boolean if a field has been set.
func (o *ModelsCertificateStoresCertificateStoreCreateRequest) HasStorepath() bool {
	if o != nil && !isNil(o.Storepath) {
		return true
	}

	return false
}

// SetStorepath gets a reference to the given string and assigns it to the Storepath field.
func (o *ModelsCertificateStoresCertificateStoreCreateRequest) SetStorepath(v string) {
	o.Storepath = &v
}

// GetCertStoreType returns the CertStoreType field value if set, zero value otherwise.
func (o *ModelsCertificateStoresCertificateStoreCreateRequest) GetCertStoreType() int32 {
	if o == nil || isNil(o.CertStoreType) {
		var ret int32
		return ret
	}
	return *o.CertStoreType
}

// GetCertStoreTypeOk returns a tuple with the CertStoreType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertificateStoresCertificateStoreCreateRequest) GetCertStoreTypeOk() (*int32, bool) {
	if o == nil || isNil(o.CertStoreType) {
		return nil, false
	}
	return o.CertStoreType, true
}

// HasCertStoreType returns a boolean if a field has been set.
func (o *ModelsCertificateStoresCertificateStoreCreateRequest) HasCertStoreType() bool {
	if o != nil && !isNil(o.CertStoreType) {
		return true
	}

	return false
}

// SetCertStoreType gets a reference to the given int32 and assigns it to the CertStoreType field.
func (o *ModelsCertificateStoresCertificateStoreCreateRequest) SetCertStoreType(v int32) {
	o.CertStoreType = &v
}

// GetCreateIfMissing returns the CreateIfMissing field value if set, zero value otherwise.
func (o *ModelsCertificateStoresCertificateStoreCreateRequest) GetCreateIfMissing() bool {
	if o == nil || isNil(o.CreateIfMissing) {
		var ret bool
		return ret
	}
	return *o.CreateIfMissing
}

// GetCreateIfMissingOk returns a tuple with the CreateIfMissing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertificateStoresCertificateStoreCreateRequest) GetCreateIfMissingOk() (*bool, bool) {
	if o == nil || isNil(o.CreateIfMissing) {
		return nil, false
	}
	return o.CreateIfMissing, true
}

// HasCreateIfMissing returns a boolean if a field has been set.
func (o *ModelsCertificateStoresCertificateStoreCreateRequest) HasCreateIfMissing() bool {
	if o != nil && !isNil(o.CreateIfMissing) {
		return true
	}

	return false
}

// SetCreateIfMissing gets a reference to the given bool and assigns it to the CreateIfMissing field.
func (o *ModelsCertificateStoresCertificateStoreCreateRequest) SetCreateIfMissing(v bool) {
	o.CreateIfMissing = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *ModelsCertificateStoresCertificateStoreCreateRequest) GetProperties() string {
	if o == nil || isNil(o.Properties) {
		var ret string
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertificateStoresCertificateStoreCreateRequest) GetPropertiesOk() (*string, bool) {
	if o == nil || isNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *ModelsCertificateStoresCertificateStoreCreateRequest) HasProperties() bool {
	if o != nil && !isNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given string and assigns it to the Properties field.
func (o *ModelsCertificateStoresCertificateStoreCreateRequest) SetProperties(v string) {
	o.Properties = &v
}

// GetAgentId returns the AgentId field value if set, zero value otherwise.
func (o *ModelsCertificateStoresCertificateStoreCreateRequest) GetAgentId() string {
	if o == nil || isNil(o.AgentId) {
		var ret string
		return ret
	}
	return *o.AgentId
}

// GetAgentIdOk returns a tuple with the AgentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertificateStoresCertificateStoreCreateRequest) GetAgentIdOk() (*string, bool) {
	if o == nil || isNil(o.AgentId) {
		return nil, false
	}
	return o.AgentId, true
}

// HasAgentId returns a boolean if a field has been set.
func (o *ModelsCertificateStoresCertificateStoreCreateRequest) HasAgentId() bool {
	if o != nil && !isNil(o.AgentId) {
		return true
	}

	return false
}

// SetAgentId gets a reference to the given string and assigns it to the AgentId field.
func (o *ModelsCertificateStoresCertificateStoreCreateRequest) SetAgentId(v string) {
	o.AgentId = &v
}

// GetAgentAssigned returns the AgentAssigned field value if set, zero value otherwise.
func (o *ModelsCertificateStoresCertificateStoreCreateRequest) GetAgentAssigned() bool {
	if o == nil || isNil(o.AgentAssigned) {
		var ret bool
		return ret
	}
	return *o.AgentAssigned
}

// GetAgentAssignedOk returns a tuple with the AgentAssigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertificateStoresCertificateStoreCreateRequest) GetAgentAssignedOk() (*bool, bool) {
	if o == nil || isNil(o.AgentAssigned) {
		return nil, false
	}
	return o.AgentAssigned, true
}

// HasAgentAssigned returns a boolean if a field has been set.
func (o *ModelsCertificateStoresCertificateStoreCreateRequest) HasAgentAssigned() bool {
	if o != nil && !isNil(o.AgentAssigned) {
		return true
	}

	return false
}

// SetAgentAssigned gets a reference to the given bool and assigns it to the AgentAssigned field.
func (o *ModelsCertificateStoresCertificateStoreCreateRequest) SetAgentAssigned(v bool) {
	o.AgentAssigned = &v
}

// GetInventorySchedule returns the InventorySchedule field value if set, zero value otherwise.
func (o *ModelsCertificateStoresCertificateStoreCreateRequest) GetInventorySchedule() KeyfactorCommonSchedulingKeyfactorSchedule {
	if o == nil || isNil(o.InventorySchedule) {
		var ret KeyfactorCommonSchedulingKeyfactorSchedule
		return ret
	}
	return *o.InventorySchedule
}

// GetInventoryScheduleOk returns a tuple with the InventorySchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertificateStoresCertificateStoreCreateRequest) GetInventoryScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool) {
	if o == nil || isNil(o.InventorySchedule) {
		return nil, false
	}
	return o.InventorySchedule, true
}

// HasInventorySchedule returns a boolean if a field has been set.
func (o *ModelsCertificateStoresCertificateStoreCreateRequest) HasInventorySchedule() bool {
	if o != nil && !isNil(o.InventorySchedule) {
		return true
	}

	return false
}

// SetInventorySchedule gets a reference to the given KeyfactorCommonSchedulingKeyfactorSchedule and assigns it to the InventorySchedule field.
func (o *ModelsCertificateStoresCertificateStoreCreateRequest) SetInventorySchedule(v KeyfactorCommonSchedulingKeyfactorSchedule) {
	o.InventorySchedule = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ModelsCertificateStoresCertificateStoreCreateRequest) GetPassword() ModelsKeyfactorAPISecret {
	if o == nil || isNil(o.Password) {
		var ret ModelsKeyfactorAPISecret
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertificateStoresCertificateStoreCreateRequest) GetPasswordOk() (*ModelsKeyfactorAPISecret, bool) {
	if o == nil || isNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ModelsCertificateStoresCertificateStoreCreateRequest) HasPassword() bool {
	if o != nil && !isNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given ModelsKeyfactorAPISecret and assigns it to the Password field.
func (o *ModelsCertificateStoresCertificateStoreCreateRequest) SetPassword(v ModelsKeyfactorAPISecret) {
	o.Password = &v
}

func (o ModelsCertificateStoresCertificateStoreCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsCertificateStoresCertificateStoreCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ContainerId) {
		toSerialize["ContainerId"] = o.ContainerId
	}
	if !isNil(o.ClientMachine) {
		toSerialize["ClientMachine"] = o.ClientMachine
	}
	if !isNil(o.Storepath) {
		toSerialize["Storepath"] = o.Storepath
	}
	if !isNil(o.CertStoreType) {
		toSerialize["CertStoreType"] = o.CertStoreType
	}
	if !isNil(o.CreateIfMissing) {
		toSerialize["CreateIfMissing"] = o.CreateIfMissing
	}
	if !isNil(o.Properties) {
		toSerialize["Properties"] = o.Properties
	}
	if !isNil(o.AgentId) {
		toSerialize["AgentId"] = o.AgentId
	}
	if !isNil(o.AgentAssigned) {
		toSerialize["AgentAssigned"] = o.AgentAssigned
	}
	if !isNil(o.InventorySchedule) {
		toSerialize["InventorySchedule"] = o.InventorySchedule
	}
	if !isNil(o.Password) {
		toSerialize["Password"] = o.Password
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModelsCertificateStoresCertificateStoreCreateRequest) UnmarshalJSON(bytes []byte) (err error) {
	varModelsCertificateStoresCertificateStoreCreateRequest := _ModelsCertificateStoresCertificateStoreCreateRequest{}

	if err = json.Unmarshal(bytes, &varModelsCertificateStoresCertificateStoreCreateRequest); err == nil {
		*o = ModelsCertificateStoresCertificateStoreCreateRequest(varModelsCertificateStoresCertificateStoreCreateRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ContainerId")
		delete(additionalProperties, "ClientMachine")
		delete(additionalProperties, "Storepath")
		delete(additionalProperties, "CertStoreType")
		delete(additionalProperties, "CreateIfMissing")
		delete(additionalProperties, "Properties")
		delete(additionalProperties, "AgentId")
		delete(additionalProperties, "AgentAssigned")
		delete(additionalProperties, "InventorySchedule")
		delete(additionalProperties, "Password")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelsCertificateStoresCertificateStoreCreateRequest struct {
	value *ModelsCertificateStoresCertificateStoreCreateRequest
	isSet bool
}

func (v NullableModelsCertificateStoresCertificateStoreCreateRequest) Get() *ModelsCertificateStoresCertificateStoreCreateRequest {
	return v.value
}

func (v *NullableModelsCertificateStoresCertificateStoreCreateRequest) Set(val *ModelsCertificateStoresCertificateStoreCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsCertificateStoresCertificateStoreCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsCertificateStoresCertificateStoreCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsCertificateStoresCertificateStoreCreateRequest(val *ModelsCertificateStoresCertificateStoreCreateRequest) *NullableModelsCertificateStoresCertificateStoreCreateRequest {
	return &NullableModelsCertificateStoresCertificateStoreCreateRequest{value: val, isSet: true}
}

func (v NullableModelsCertificateStoresCertificateStoreCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsCertificateStoresCertificateStoreCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
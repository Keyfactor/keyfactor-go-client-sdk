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
	"time"
)

// checks if the KeyfactorApiModelsOrchestratorsAgentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorApiModelsOrchestratorsAgentResponse{}

// KeyfactorApiModelsOrchestratorsAgentResponse struct for KeyfactorApiModelsOrchestratorsAgentResponse
type KeyfactorApiModelsOrchestratorsAgentResponse struct {
	AgentId *string `json:"AgentId,omitempty"`
	ClientMachine *string `json:"ClientMachine,omitempty"`
	Username *string `json:"Username,omitempty"`
	AgentPlatform *int32 `json:"AgentPlatform,omitempty"`
	Version *string `json:"Version,omitempty"`
	Status *int32 `json:"Status,omitempty"`
	LastSeen *time.Time `json:"LastSeen,omitempty"`
	Capabilities []string `json:"Capabilities,omitempty"`
	Blueprint *string `json:"Blueprint,omitempty"`
	Thumbprint *string `json:"Thumbprint,omitempty"`
	LegacyThumbprint *string `json:"LegacyThumbprint,omitempty"`
	AuthCertificateReenrollment *string `json:"AuthCertificateReenrollment,omitempty"`
	LastThumbprintUsed *string `json:"LastThumbprintUsed,omitempty"`
	LastErrorCode *int64 `json:"LastErrorCode,omitempty"`
	LastErrorMessage *string `json:"LastErrorMessage,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KeyfactorApiModelsOrchestratorsAgentResponse KeyfactorApiModelsOrchestratorsAgentResponse

// NewKeyfactorApiModelsOrchestratorsAgentResponse instantiates a new KeyfactorApiModelsOrchestratorsAgentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorApiModelsOrchestratorsAgentResponse() *KeyfactorApiModelsOrchestratorsAgentResponse {
	this := KeyfactorApiModelsOrchestratorsAgentResponse{}
	return &this
}

// NewKeyfactorApiModelsOrchestratorsAgentResponseWithDefaults instantiates a new KeyfactorApiModelsOrchestratorsAgentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorApiModelsOrchestratorsAgentResponseWithDefaults() *KeyfactorApiModelsOrchestratorsAgentResponse {
	this := KeyfactorApiModelsOrchestratorsAgentResponse{}
	return &this
}

// GetAgentId returns the AgentId field value if set, zero value otherwise.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetAgentId() string {
	if o == nil || isNil(o.AgentId) {
		var ret string
		return ret
	}
	return *o.AgentId
}

// GetAgentIdOk returns a tuple with the AgentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetAgentIdOk() (*string, bool) {
	if o == nil || isNil(o.AgentId) {
		return nil, false
	}
	return o.AgentId, true
}

// HasAgentId returns a boolean if a field has been set.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) HasAgentId() bool {
	if o != nil && !isNil(o.AgentId) {
		return true
	}

	return false
}

// SetAgentId gets a reference to the given string and assigns it to the AgentId field.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) SetAgentId(v string) {
	o.AgentId = &v
}

// GetClientMachine returns the ClientMachine field value if set, zero value otherwise.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetClientMachine() string {
	if o == nil || isNil(o.ClientMachine) {
		var ret string
		return ret
	}
	return *o.ClientMachine
}

// GetClientMachineOk returns a tuple with the ClientMachine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetClientMachineOk() (*string, bool) {
	if o == nil || isNil(o.ClientMachine) {
		return nil, false
	}
	return o.ClientMachine, true
}

// HasClientMachine returns a boolean if a field has been set.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) HasClientMachine() bool {
	if o != nil && !isNil(o.ClientMachine) {
		return true
	}

	return false
}

// SetClientMachine gets a reference to the given string and assigns it to the ClientMachine field.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) SetClientMachine(v string) {
	o.ClientMachine = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetUsername() string {
	if o == nil || isNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetUsernameOk() (*string, bool) {
	if o == nil || isNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) HasUsername() bool {
	if o != nil && !isNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) SetUsername(v string) {
	o.Username = &v
}

// GetAgentPlatform returns the AgentPlatform field value if set, zero value otherwise.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetAgentPlatform() int32 {
	if o == nil || isNil(o.AgentPlatform) {
		var ret int32
		return ret
	}
	return *o.AgentPlatform
}

// GetAgentPlatformOk returns a tuple with the AgentPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetAgentPlatformOk() (*int32, bool) {
	if o == nil || isNil(o.AgentPlatform) {
		return nil, false
	}
	return o.AgentPlatform, true
}

// HasAgentPlatform returns a boolean if a field has been set.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) HasAgentPlatform() bool {
	if o != nil && !isNil(o.AgentPlatform) {
		return true
	}

	return false
}

// SetAgentPlatform gets a reference to the given int32 and assigns it to the AgentPlatform field.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) SetAgentPlatform(v int32) {
	o.AgentPlatform = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetVersion() string {
	if o == nil || isNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetVersionOk() (*string, bool) {
	if o == nil || isNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) SetVersion(v string) {
	o.Version = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetStatus() int32 {
	if o == nil || isNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetStatusOk() (*int32, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) SetStatus(v int32) {
	o.Status = &v
}

// GetLastSeen returns the LastSeen field value if set, zero value otherwise.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetLastSeen() time.Time {
	if o == nil || isNil(o.LastSeen) {
		var ret time.Time
		return ret
	}
	return *o.LastSeen
}

// GetLastSeenOk returns a tuple with the LastSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetLastSeenOk() (*time.Time, bool) {
	if o == nil || isNil(o.LastSeen) {
		return nil, false
	}
	return o.LastSeen, true
}

// HasLastSeen returns a boolean if a field has been set.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) HasLastSeen() bool {
	if o != nil && !isNil(o.LastSeen) {
		return true
	}

	return false
}

// SetLastSeen gets a reference to the given time.Time and assigns it to the LastSeen field.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) SetLastSeen(v time.Time) {
	o.LastSeen = &v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetCapabilities() []string {
	if o == nil || isNil(o.Capabilities) {
		var ret []string
		return ret
	}
	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetCapabilitiesOk() ([]string, bool) {
	if o == nil || isNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) HasCapabilities() bool {
	if o != nil && !isNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given []string and assigns it to the Capabilities field.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) SetCapabilities(v []string) {
	o.Capabilities = v
}

// GetBlueprint returns the Blueprint field value if set, zero value otherwise.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetBlueprint() string {
	if o == nil || isNil(o.Blueprint) {
		var ret string
		return ret
	}
	return *o.Blueprint
}

// GetBlueprintOk returns a tuple with the Blueprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetBlueprintOk() (*string, bool) {
	if o == nil || isNil(o.Blueprint) {
		return nil, false
	}
	return o.Blueprint, true
}

// HasBlueprint returns a boolean if a field has been set.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) HasBlueprint() bool {
	if o != nil && !isNil(o.Blueprint) {
		return true
	}

	return false
}

// SetBlueprint gets a reference to the given string and assigns it to the Blueprint field.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) SetBlueprint(v string) {
	o.Blueprint = &v
}

// GetThumbprint returns the Thumbprint field value if set, zero value otherwise.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetThumbprint() string {
	if o == nil || isNil(o.Thumbprint) {
		var ret string
		return ret
	}
	return *o.Thumbprint
}

// GetThumbprintOk returns a tuple with the Thumbprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetThumbprintOk() (*string, bool) {
	if o == nil || isNil(o.Thumbprint) {
		return nil, false
	}
	return o.Thumbprint, true
}

// HasThumbprint returns a boolean if a field has been set.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) HasThumbprint() bool {
	if o != nil && !isNil(o.Thumbprint) {
		return true
	}

	return false
}

// SetThumbprint gets a reference to the given string and assigns it to the Thumbprint field.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) SetThumbprint(v string) {
	o.Thumbprint = &v
}

// GetLegacyThumbprint returns the LegacyThumbprint field value if set, zero value otherwise.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetLegacyThumbprint() string {
	if o == nil || isNil(o.LegacyThumbprint) {
		var ret string
		return ret
	}
	return *o.LegacyThumbprint
}

// GetLegacyThumbprintOk returns a tuple with the LegacyThumbprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetLegacyThumbprintOk() (*string, bool) {
	if o == nil || isNil(o.LegacyThumbprint) {
		return nil, false
	}
	return o.LegacyThumbprint, true
}

// HasLegacyThumbprint returns a boolean if a field has been set.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) HasLegacyThumbprint() bool {
	if o != nil && !isNil(o.LegacyThumbprint) {
		return true
	}

	return false
}

// SetLegacyThumbprint gets a reference to the given string and assigns it to the LegacyThumbprint field.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) SetLegacyThumbprint(v string) {
	o.LegacyThumbprint = &v
}

// GetAuthCertificateReenrollment returns the AuthCertificateReenrollment field value if set, zero value otherwise.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetAuthCertificateReenrollment() string {
	if o == nil || isNil(o.AuthCertificateReenrollment) {
		var ret string
		return ret
	}
	return *o.AuthCertificateReenrollment
}

// GetAuthCertificateReenrollmentOk returns a tuple with the AuthCertificateReenrollment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetAuthCertificateReenrollmentOk() (*string, bool) {
	if o == nil || isNil(o.AuthCertificateReenrollment) {
		return nil, false
	}
	return o.AuthCertificateReenrollment, true
}

// HasAuthCertificateReenrollment returns a boolean if a field has been set.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) HasAuthCertificateReenrollment() bool {
	if o != nil && !isNil(o.AuthCertificateReenrollment) {
		return true
	}

	return false
}

// SetAuthCertificateReenrollment gets a reference to the given string and assigns it to the AuthCertificateReenrollment field.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) SetAuthCertificateReenrollment(v string) {
	o.AuthCertificateReenrollment = &v
}

// GetLastThumbprintUsed returns the LastThumbprintUsed field value if set, zero value otherwise.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetLastThumbprintUsed() string {
	if o == nil || isNil(o.LastThumbprintUsed) {
		var ret string
		return ret
	}
	return *o.LastThumbprintUsed
}

// GetLastThumbprintUsedOk returns a tuple with the LastThumbprintUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetLastThumbprintUsedOk() (*string, bool) {
	if o == nil || isNil(o.LastThumbprintUsed) {
		return nil, false
	}
	return o.LastThumbprintUsed, true
}

// HasLastThumbprintUsed returns a boolean if a field has been set.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) HasLastThumbprintUsed() bool {
	if o != nil && !isNil(o.LastThumbprintUsed) {
		return true
	}

	return false
}

// SetLastThumbprintUsed gets a reference to the given string and assigns it to the LastThumbprintUsed field.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) SetLastThumbprintUsed(v string) {
	o.LastThumbprintUsed = &v
}

// GetLastErrorCode returns the LastErrorCode field value if set, zero value otherwise.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetLastErrorCode() int64 {
	if o == nil || isNil(o.LastErrorCode) {
		var ret int64
		return ret
	}
	return *o.LastErrorCode
}

// GetLastErrorCodeOk returns a tuple with the LastErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetLastErrorCodeOk() (*int64, bool) {
	if o == nil || isNil(o.LastErrorCode) {
		return nil, false
	}
	return o.LastErrorCode, true
}

// HasLastErrorCode returns a boolean if a field has been set.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) HasLastErrorCode() bool {
	if o != nil && !isNil(o.LastErrorCode) {
		return true
	}

	return false
}

// SetLastErrorCode gets a reference to the given int64 and assigns it to the LastErrorCode field.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) SetLastErrorCode(v int64) {
	o.LastErrorCode = &v
}

// GetLastErrorMessage returns the LastErrorMessage field value if set, zero value otherwise.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetLastErrorMessage() string {
	if o == nil || isNil(o.LastErrorMessage) {
		var ret string
		return ret
	}
	return *o.LastErrorMessage
}

// GetLastErrorMessageOk returns a tuple with the LastErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) GetLastErrorMessageOk() (*string, bool) {
	if o == nil || isNil(o.LastErrorMessage) {
		return nil, false
	}
	return o.LastErrorMessage, true
}

// HasLastErrorMessage returns a boolean if a field has been set.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) HasLastErrorMessage() bool {
	if o != nil && !isNil(o.LastErrorMessage) {
		return true
	}

	return false
}

// SetLastErrorMessage gets a reference to the given string and assigns it to the LastErrorMessage field.
func (o *KeyfactorApiModelsOrchestratorsAgentResponse) SetLastErrorMessage(v string) {
	o.LastErrorMessage = &v
}

func (o KeyfactorApiModelsOrchestratorsAgentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorApiModelsOrchestratorsAgentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AgentId) {
		toSerialize["AgentId"] = o.AgentId
	}
	if !isNil(o.ClientMachine) {
		toSerialize["ClientMachine"] = o.ClientMachine
	}
	if !isNil(o.Username) {
		toSerialize["Username"] = o.Username
	}
	if !isNil(o.AgentPlatform) {
		toSerialize["AgentPlatform"] = o.AgentPlatform
	}
	if !isNil(o.Version) {
		toSerialize["Version"] = o.Version
	}
	if !isNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if !isNil(o.LastSeen) {
		toSerialize["LastSeen"] = o.LastSeen
	}
	if !isNil(o.Capabilities) {
		toSerialize["Capabilities"] = o.Capabilities
	}
	if !isNil(o.Blueprint) {
		toSerialize["Blueprint"] = o.Blueprint
	}
	if !isNil(o.Thumbprint) {
		toSerialize["Thumbprint"] = o.Thumbprint
	}
	if !isNil(o.LegacyThumbprint) {
		toSerialize["LegacyThumbprint"] = o.LegacyThumbprint
	}
	if !isNil(o.AuthCertificateReenrollment) {
		toSerialize["AuthCertificateReenrollment"] = o.AuthCertificateReenrollment
	}
	if !isNil(o.LastThumbprintUsed) {
		toSerialize["LastThumbprintUsed"] = o.LastThumbprintUsed
	}
	if !isNil(o.LastErrorCode) {
		toSerialize["LastErrorCode"] = o.LastErrorCode
	}
	if !isNil(o.LastErrorMessage) {
		toSerialize["LastErrorMessage"] = o.LastErrorMessage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KeyfactorApiModelsOrchestratorsAgentResponse) UnmarshalJSON(bytes []byte) (err error) {
	varKeyfactorApiModelsOrchestratorsAgentResponse := _KeyfactorApiModelsOrchestratorsAgentResponse{}

	if err = json.Unmarshal(bytes, &varKeyfactorApiModelsOrchestratorsAgentResponse); err == nil {
		*o = KeyfactorApiModelsOrchestratorsAgentResponse(varKeyfactorApiModelsOrchestratorsAgentResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "AgentId")
		delete(additionalProperties, "ClientMachine")
		delete(additionalProperties, "Username")
		delete(additionalProperties, "AgentPlatform")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "LastSeen")
		delete(additionalProperties, "Capabilities")
		delete(additionalProperties, "Blueprint")
		delete(additionalProperties, "Thumbprint")
		delete(additionalProperties, "LegacyThumbprint")
		delete(additionalProperties, "AuthCertificateReenrollment")
		delete(additionalProperties, "LastThumbprintUsed")
		delete(additionalProperties, "LastErrorCode")
		delete(additionalProperties, "LastErrorMessage")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKeyfactorApiModelsOrchestratorsAgentResponse struct {
	value *KeyfactorApiModelsOrchestratorsAgentResponse
	isSet bool
}

func (v NullableKeyfactorApiModelsOrchestratorsAgentResponse) Get() *KeyfactorApiModelsOrchestratorsAgentResponse {
	return v.value
}

func (v *NullableKeyfactorApiModelsOrchestratorsAgentResponse) Set(val *KeyfactorApiModelsOrchestratorsAgentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorApiModelsOrchestratorsAgentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorApiModelsOrchestratorsAgentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorApiModelsOrchestratorsAgentResponse(val *KeyfactorApiModelsOrchestratorsAgentResponse) *NullableKeyfactorApiModelsOrchestratorsAgentResponse {
	return &NullableKeyfactorApiModelsOrchestratorsAgentResponse{value: val, isSet: true}
}

func (v NullableKeyfactorApiModelsOrchestratorsAgentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorApiModelsOrchestratorsAgentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


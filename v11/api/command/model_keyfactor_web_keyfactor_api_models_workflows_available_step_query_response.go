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

// checks if the KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse{}

// KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse struct for KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse
type KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse struct {
	// The display name of the step.
	DisplayName NullableString `json:"displayName,omitempty"`
	// The extension name of the step.
	ExtensionName NullableString `json:"extensionName,omitempty"`
	// The workflow types which this step can be a part of.
	SupportedWorkflowTypes            []string                                                                      `json:"supportedWorkflowTypes,omitempty"`
	ConfigurationParametersDefinition map[string]KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse `json:"configurationParametersDefinition,omitempty"`
	SignalsDefinition                 map[string]KeyfactorWebKeyfactorApiModelsWorkflowsSignalDefinitionResponse    `json:"signalsDefinition,omitempty"`
}

// NewKeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse instantiates a new KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse() *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse {
	this := KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse{}
	return &this
}

// NewKeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponseWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponseWithDefaults() *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse {
	this := KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName.Get()) {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}

// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetExtensionName returns the ExtensionName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse) GetExtensionName() string {
	if o == nil || isNil(o.ExtensionName.Get()) {
		var ret string
		return ret
	}
	return *o.ExtensionName.Get()
}

// GetExtensionNameOk returns a tuple with the ExtensionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse) GetExtensionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExtensionName.Get(), o.ExtensionName.IsSet()
}

// HasExtensionName returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse) HasExtensionName() bool {
	if o != nil && o.ExtensionName.IsSet() {
		return true
	}

	return false
}

// SetExtensionName gets a reference to the given NullableString and assigns it to the ExtensionName field.
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse) SetExtensionName(v string) {
	o.ExtensionName.Set(&v)
}

// SetExtensionNameNil sets the value for ExtensionName to be an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse) SetExtensionNameNil() {
	o.ExtensionName.Set(nil)
}

// UnsetExtensionName ensures that no value is present for ExtensionName, not even an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse) UnsetExtensionName() {
	o.ExtensionName.Unset()
}

// GetSupportedWorkflowTypes returns the SupportedWorkflowTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse) GetSupportedWorkflowTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SupportedWorkflowTypes
}

// GetSupportedWorkflowTypesOk returns a tuple with the SupportedWorkflowTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse) GetSupportedWorkflowTypesOk() ([]string, bool) {
	if o == nil || isNil(o.SupportedWorkflowTypes) {
		return nil, false
	}
	return o.SupportedWorkflowTypes, true
}

// HasSupportedWorkflowTypes returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse) HasSupportedWorkflowTypes() bool {
	if o != nil && isNil(o.SupportedWorkflowTypes) {
		return true
	}

	return false
}

// SetSupportedWorkflowTypes gets a reference to the given []string and assigns it to the SupportedWorkflowTypes field.
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse) SetSupportedWorkflowTypes(v []string) {
	o.SupportedWorkflowTypes = v
}

// GetConfigurationParametersDefinition returns the ConfigurationParametersDefinition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse) GetConfigurationParametersDefinition() map[string]KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse {
	if o == nil {
		var ret map[string]KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse
		return ret
	}
	return o.ConfigurationParametersDefinition
}

// GetConfigurationParametersDefinitionOk returns a tuple with the ConfigurationParametersDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse) GetConfigurationParametersDefinitionOk() (*map[string]KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse, bool) {
	if o == nil || isNil(o.ConfigurationParametersDefinition) {
		return nil, false
	}
	return &o.ConfigurationParametersDefinition, true
}

// HasConfigurationParametersDefinition returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse) HasConfigurationParametersDefinition() bool {
	if o != nil && isNil(o.ConfigurationParametersDefinition) {
		return true
	}

	return false
}

// SetConfigurationParametersDefinition gets a reference to the given map[string]KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse and assigns it to the ConfigurationParametersDefinition field.
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse) SetConfigurationParametersDefinition(v map[string]KeyfactorWebKeyfactorApiModelsWorkflowsParameterDefinitionResponse) {
	o.ConfigurationParametersDefinition = v
}

// GetSignalsDefinition returns the SignalsDefinition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse) GetSignalsDefinition() map[string]KeyfactorWebKeyfactorApiModelsWorkflowsSignalDefinitionResponse {
	if o == nil {
		var ret map[string]KeyfactorWebKeyfactorApiModelsWorkflowsSignalDefinitionResponse
		return ret
	}
	return o.SignalsDefinition
}

// GetSignalsDefinitionOk returns a tuple with the SignalsDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse) GetSignalsDefinitionOk() (*map[string]KeyfactorWebKeyfactorApiModelsWorkflowsSignalDefinitionResponse, bool) {
	if o == nil || isNil(o.SignalsDefinition) {
		return nil, false
	}
	return &o.SignalsDefinition, true
}

// HasSignalsDefinition returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse) HasSignalsDefinition() bool {
	if o != nil && isNil(o.SignalsDefinition) {
		return true
	}

	return false
}

// SetSignalsDefinition gets a reference to the given map[string]KeyfactorWebKeyfactorApiModelsWorkflowsSignalDefinitionResponse and assigns it to the SignalsDefinition field.
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse) SetSignalsDefinition(v map[string]KeyfactorWebKeyfactorApiModelsWorkflowsSignalDefinitionResponse) {
	o.SignalsDefinition = v
}

func (o KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.ExtensionName.IsSet() {
		toSerialize["extensionName"] = o.ExtensionName.Get()
	}
	if o.SupportedWorkflowTypes != nil {
		toSerialize["supportedWorkflowTypes"] = o.SupportedWorkflowTypes
	}
	if o.ConfigurationParametersDefinition != nil {
		toSerialize["configurationParametersDefinition"] = o.ConfigurationParametersDefinition
	}
	if o.SignalsDefinition != nil {
		toSerialize["signalsDefinition"] = o.SignalsDefinition
	}
	return toSerialize, nil
}

type NullableKeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse struct {
	value *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse
	isSet bool
}

func (v NullableKeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse) Get() *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse {
	return v.value
}

func (v *NullableKeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse) Set(val *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse(val *KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse) *NullableKeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse {
	return &NullableKeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse{value: val, isSet: true}
}

func (v NullableKeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
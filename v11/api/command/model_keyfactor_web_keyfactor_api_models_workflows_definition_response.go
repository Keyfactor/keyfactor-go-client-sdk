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

// checks if the KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse{}

// KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse struct for KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse
type KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse struct {
	Id               *string                                                         `json:"id,omitempty"`
	DisplayName      NullableString                                                  `json:"displayName,omitempty"`
	Description      NullableString                                                  `json:"description,omitempty"`
	Key              NullableString                                                  `json:"key,omitempty"`
	KeyDisplayName   NullableString                                                  `json:"keyDisplayName,omitempty"`
	IsPublished      *bool                                                           `json:"isPublished,omitempty"`
	WorkflowType     NullableString                                                  `json:"workflowType,omitempty"`
	Steps            []KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse `json:"steps,omitempty"`
	DraftVersion     *int32                                                          `json:"draftVersion,omitempty"`
	PublishedVersion NullableInt32                                                   `json:"publishedVersion,omitempty"`
}

// NewKeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse instantiates a new KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse() *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse {
	this := KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse{}
	return &this
}

// NewKeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponseWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponseWithDefaults() *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse {
	this := KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) SetId(v string) {
	o.Id = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName.Get()) {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}

// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) GetDescription() string {
	if o == nil || isNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) UnsetDescription() {
	o.Description.Unset()
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) GetKey() string {
	if o == nil || isNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) SetKey(v string) {
	o.Key.Set(&v)
}

// SetKeyNil sets the value for Key to be an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) UnsetKey() {
	o.Key.Unset()
}

// GetKeyDisplayName returns the KeyDisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) GetKeyDisplayName() string {
	if o == nil || isNil(o.KeyDisplayName.Get()) {
		var ret string
		return ret
	}
	return *o.KeyDisplayName.Get()
}

// GetKeyDisplayNameOk returns a tuple with the KeyDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) GetKeyDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.KeyDisplayName.Get(), o.KeyDisplayName.IsSet()
}

// HasKeyDisplayName returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) HasKeyDisplayName() bool {
	if o != nil && o.KeyDisplayName.IsSet() {
		return true
	}

	return false
}

// SetKeyDisplayName gets a reference to the given NullableString and assigns it to the KeyDisplayName field.
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) SetKeyDisplayName(v string) {
	o.KeyDisplayName.Set(&v)
}

// SetKeyDisplayNameNil sets the value for KeyDisplayName to be an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) SetKeyDisplayNameNil() {
	o.KeyDisplayName.Set(nil)
}

// UnsetKeyDisplayName ensures that no value is present for KeyDisplayName, not even an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) UnsetKeyDisplayName() {
	o.KeyDisplayName.Unset()
}

// GetIsPublished returns the IsPublished field value if set, zero value otherwise.
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) GetIsPublished() bool {
	if o == nil || isNil(o.IsPublished) {
		var ret bool
		return ret
	}
	return *o.IsPublished
}

// GetIsPublishedOk returns a tuple with the IsPublished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) GetIsPublishedOk() (*bool, bool) {
	if o == nil || isNil(o.IsPublished) {
		return nil, false
	}
	return o.IsPublished, true
}

// HasIsPublished returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) HasIsPublished() bool {
	if o != nil && !isNil(o.IsPublished) {
		return true
	}

	return false
}

// SetIsPublished gets a reference to the given bool and assigns it to the IsPublished field.
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) SetIsPublished(v bool) {
	o.IsPublished = &v
}

// GetWorkflowType returns the WorkflowType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) GetWorkflowType() string {
	if o == nil || isNil(o.WorkflowType.Get()) {
		var ret string
		return ret
	}
	return *o.WorkflowType.Get()
}

// GetWorkflowTypeOk returns a tuple with the WorkflowType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) GetWorkflowTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkflowType.Get(), o.WorkflowType.IsSet()
}

// HasWorkflowType returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) HasWorkflowType() bool {
	if o != nil && o.WorkflowType.IsSet() {
		return true
	}

	return false
}

// SetWorkflowType gets a reference to the given NullableString and assigns it to the WorkflowType field.
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) SetWorkflowType(v string) {
	o.WorkflowType.Set(&v)
}

// SetWorkflowTypeNil sets the value for WorkflowType to be an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) SetWorkflowTypeNil() {
	o.WorkflowType.Set(nil)
}

// UnsetWorkflowType ensures that no value is present for WorkflowType, not even an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) UnsetWorkflowType() {
	o.WorkflowType.Unset()
}

// GetSteps returns the Steps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) GetSteps() []KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse {
	if o == nil {
		var ret []KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) GetStepsOk() ([]KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse, bool) {
	if o == nil || isNil(o.Steps) {
		return nil, false
	}
	return o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) HasSteps() bool {
	if o != nil && isNil(o.Steps) {
		return true
	}

	return false
}

// SetSteps gets a reference to the given []KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse and assigns it to the Steps field.
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) SetSteps(v []KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepResponse) {
	o.Steps = v
}

// GetDraftVersion returns the DraftVersion field value if set, zero value otherwise.
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) GetDraftVersion() int32 {
	if o == nil || isNil(o.DraftVersion) {
		var ret int32
		return ret
	}
	return *o.DraftVersion
}

// GetDraftVersionOk returns a tuple with the DraftVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) GetDraftVersionOk() (*int32, bool) {
	if o == nil || isNil(o.DraftVersion) {
		return nil, false
	}
	return o.DraftVersion, true
}

// HasDraftVersion returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) HasDraftVersion() bool {
	if o != nil && !isNil(o.DraftVersion) {
		return true
	}

	return false
}

// SetDraftVersion gets a reference to the given int32 and assigns it to the DraftVersion field.
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) SetDraftVersion(v int32) {
	o.DraftVersion = &v
}

// GetPublishedVersion returns the PublishedVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) GetPublishedVersion() int32 {
	if o == nil || isNil(o.PublishedVersion.Get()) {
		var ret int32
		return ret
	}
	return *o.PublishedVersion.Get()
}

// GetPublishedVersionOk returns a tuple with the PublishedVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) GetPublishedVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PublishedVersion.Get(), o.PublishedVersion.IsSet()
}

// HasPublishedVersion returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) HasPublishedVersion() bool {
	if o != nil && o.PublishedVersion.IsSet() {
		return true
	}

	return false
}

// SetPublishedVersion gets a reference to the given NullableInt32 and assigns it to the PublishedVersion field.
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) SetPublishedVersion(v int32) {
	o.PublishedVersion.Set(&v)
}

// SetPublishedVersionNil sets the value for PublishedVersion to be an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) SetPublishedVersionNil() {
	o.PublishedVersion.Set(nil)
}

// UnsetPublishedVersion ensures that no value is present for PublishedVersion, not even an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) UnsetPublishedVersion() {
	o.PublishedVersion.Unset()
}

func (o KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if o.KeyDisplayName.IsSet() {
		toSerialize["keyDisplayName"] = o.KeyDisplayName.Get()
	}
	if !isNil(o.IsPublished) {
		toSerialize["isPublished"] = o.IsPublished
	}
	if o.WorkflowType.IsSet() {
		toSerialize["workflowType"] = o.WorkflowType.Get()
	}
	if o.Steps != nil {
		toSerialize["steps"] = o.Steps
	}
	if !isNil(o.DraftVersion) {
		toSerialize["draftVersion"] = o.DraftVersion
	}
	if o.PublishedVersion.IsSet() {
		toSerialize["publishedVersion"] = o.PublishedVersion.Get()
	}
	return toSerialize, nil
}

type NullableKeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse struct {
	value *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse
	isSet bool
}

func (v NullableKeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) Get() *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse {
	return v.value
}

func (v *NullableKeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) Set(val *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse(val *KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) *NullableKeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse {
	return &NullableKeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse{value: val, isSet: true}
}

func (v NullableKeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
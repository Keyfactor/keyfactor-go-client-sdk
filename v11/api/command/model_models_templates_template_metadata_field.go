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

// checks if the ModelsTemplatesTemplateMetadataField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsTemplatesTemplateMetadataField{}

// ModelsTemplatesTemplateMetadataField struct for ModelsTemplatesTemplateMetadataField
type ModelsTemplatesTemplateMetadataField struct {
	Id           *int32         `json:"id,omitempty"`
	DefaultValue NullableString `json:"defaultValue,omitempty"`
	MetadataId   *int32         `json:"metadataId,omitempty"`
	Name         NullableString `json:"name,omitempty"`
	Hint         NullableString `json:"hint,omitempty"`
	Validation   NullableString `json:"validation,omitempty"`
	Enrollment   *int32         `json:"enrollment,omitempty"`
	Message      NullableString `json:"message,omitempty"`
}

// NewModelsTemplatesTemplateMetadataField instantiates a new ModelsTemplatesTemplateMetadataField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsTemplatesTemplateMetadataField() *ModelsTemplatesTemplateMetadataField {
	this := ModelsTemplatesTemplateMetadataField{}
	return &this
}

// NewModelsTemplatesTemplateMetadataFieldWithDefaults instantiates a new ModelsTemplatesTemplateMetadataField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsTemplatesTemplateMetadataFieldWithDefaults() *ModelsTemplatesTemplateMetadataField {
	this := ModelsTemplatesTemplateMetadataField{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelsTemplatesTemplateMetadataField) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsTemplatesTemplateMetadataField) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelsTemplatesTemplateMetadataField) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ModelsTemplatesTemplateMetadataField) SetId(v int32) {
	o.Id = &v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelsTemplatesTemplateMetadataField) GetDefaultValue() string {
	if o == nil || isNil(o.DefaultValue.Get()) {
		var ret string
		return ret
	}
	return *o.DefaultValue.Get()
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelsTemplatesTemplateMetadataField) GetDefaultValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultValue.Get(), o.DefaultValue.IsSet()
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *ModelsTemplatesTemplateMetadataField) HasDefaultValue() bool {
	if o != nil && o.DefaultValue.IsSet() {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given NullableString and assigns it to the DefaultValue field.
func (o *ModelsTemplatesTemplateMetadataField) SetDefaultValue(v string) {
	o.DefaultValue.Set(&v)
}

// SetDefaultValueNil sets the value for DefaultValue to be an explicit nil
func (o *ModelsTemplatesTemplateMetadataField) SetDefaultValueNil() {
	o.DefaultValue.Set(nil)
}

// UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
func (o *ModelsTemplatesTemplateMetadataField) UnsetDefaultValue() {
	o.DefaultValue.Unset()
}

// GetMetadataId returns the MetadataId field value if set, zero value otherwise.
func (o *ModelsTemplatesTemplateMetadataField) GetMetadataId() int32 {
	if o == nil || isNil(o.MetadataId) {
		var ret int32
		return ret
	}
	return *o.MetadataId
}

// GetMetadataIdOk returns a tuple with the MetadataId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsTemplatesTemplateMetadataField) GetMetadataIdOk() (*int32, bool) {
	if o == nil || isNil(o.MetadataId) {
		return nil, false
	}
	return o.MetadataId, true
}

// HasMetadataId returns a boolean if a field has been set.
func (o *ModelsTemplatesTemplateMetadataField) HasMetadataId() bool {
	if o != nil && !isNil(o.MetadataId) {
		return true
	}

	return false
}

// SetMetadataId gets a reference to the given int32 and assigns it to the MetadataId field.
func (o *ModelsTemplatesTemplateMetadataField) SetMetadataId(v int32) {
	o.MetadataId = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelsTemplatesTemplateMetadataField) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelsTemplatesTemplateMetadataField) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ModelsTemplatesTemplateMetadataField) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ModelsTemplatesTemplateMetadataField) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *ModelsTemplatesTemplateMetadataField) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ModelsTemplatesTemplateMetadataField) UnsetName() {
	o.Name.Unset()
}

// GetHint returns the Hint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelsTemplatesTemplateMetadataField) GetHint() string {
	if o == nil || isNil(o.Hint.Get()) {
		var ret string
		return ret
	}
	return *o.Hint.Get()
}

// GetHintOk returns a tuple with the Hint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelsTemplatesTemplateMetadataField) GetHintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hint.Get(), o.Hint.IsSet()
}

// HasHint returns a boolean if a field has been set.
func (o *ModelsTemplatesTemplateMetadataField) HasHint() bool {
	if o != nil && o.Hint.IsSet() {
		return true
	}

	return false
}

// SetHint gets a reference to the given NullableString and assigns it to the Hint field.
func (o *ModelsTemplatesTemplateMetadataField) SetHint(v string) {
	o.Hint.Set(&v)
}

// SetHintNil sets the value for Hint to be an explicit nil
func (o *ModelsTemplatesTemplateMetadataField) SetHintNil() {
	o.Hint.Set(nil)
}

// UnsetHint ensures that no value is present for Hint, not even an explicit nil
func (o *ModelsTemplatesTemplateMetadataField) UnsetHint() {
	o.Hint.Unset()
}

// GetValidation returns the Validation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelsTemplatesTemplateMetadataField) GetValidation() string {
	if o == nil || isNil(o.Validation.Get()) {
		var ret string
		return ret
	}
	return *o.Validation.Get()
}

// GetValidationOk returns a tuple with the Validation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelsTemplatesTemplateMetadataField) GetValidationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Validation.Get(), o.Validation.IsSet()
}

// HasValidation returns a boolean if a field has been set.
func (o *ModelsTemplatesTemplateMetadataField) HasValidation() bool {
	if o != nil && o.Validation.IsSet() {
		return true
	}

	return false
}

// SetValidation gets a reference to the given NullableString and assigns it to the Validation field.
func (o *ModelsTemplatesTemplateMetadataField) SetValidation(v string) {
	o.Validation.Set(&v)
}

// SetValidationNil sets the value for Validation to be an explicit nil
func (o *ModelsTemplatesTemplateMetadataField) SetValidationNil() {
	o.Validation.Set(nil)
}

// UnsetValidation ensures that no value is present for Validation, not even an explicit nil
func (o *ModelsTemplatesTemplateMetadataField) UnsetValidation() {
	o.Validation.Unset()
}

// GetEnrollment returns the Enrollment field value if set, zero value otherwise.
func (o *ModelsTemplatesTemplateMetadataField) GetEnrollment() int32 {
	if o == nil || isNil(o.Enrollment) {
		var ret int32
		return ret
	}
	return *o.Enrollment
}

// GetEnrollmentOk returns a tuple with the Enrollment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsTemplatesTemplateMetadataField) GetEnrollmentOk() (*int32, bool) {
	if o == nil || isNil(o.Enrollment) {
		return nil, false
	}
	return o.Enrollment, true
}

// HasEnrollment returns a boolean if a field has been set.
func (o *ModelsTemplatesTemplateMetadataField) HasEnrollment() bool {
	if o != nil && !isNil(o.Enrollment) {
		return true
	}

	return false
}

// SetEnrollment gets a reference to the given int32 and assigns it to the Enrollment field.
func (o *ModelsTemplatesTemplateMetadataField) SetEnrollment(v int32) {
	o.Enrollment = &v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelsTemplatesTemplateMetadataField) GetMessage() string {
	if o == nil || isNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelsTemplatesTemplateMetadataField) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *ModelsTemplatesTemplateMetadataField) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *ModelsTemplatesTemplateMetadataField) SetMessage(v string) {
	o.Message.Set(&v)
}

// SetMessageNil sets the value for Message to be an explicit nil
func (o *ModelsTemplatesTemplateMetadataField) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *ModelsTemplatesTemplateMetadataField) UnsetMessage() {
	o.Message.Unset()
}

func (o ModelsTemplatesTemplateMetadataField) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsTemplatesTemplateMetadataField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.DefaultValue.IsSet() {
		toSerialize["defaultValue"] = o.DefaultValue.Get()
	}
	if !isNil(o.MetadataId) {
		toSerialize["metadataId"] = o.MetadataId
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Hint.IsSet() {
		toSerialize["hint"] = o.Hint.Get()
	}
	if o.Validation.IsSet() {
		toSerialize["validation"] = o.Validation.Get()
	}
	if !isNil(o.Enrollment) {
		toSerialize["enrollment"] = o.Enrollment
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	return toSerialize, nil
}

type NullableModelsTemplatesTemplateMetadataField struct {
	value *ModelsTemplatesTemplateMetadataField
	isSet bool
}

func (v NullableModelsTemplatesTemplateMetadataField) Get() *ModelsTemplatesTemplateMetadataField {
	return v.value
}

func (v *NullableModelsTemplatesTemplateMetadataField) Set(val *ModelsTemplatesTemplateMetadataField) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsTemplatesTemplateMetadataField) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsTemplatesTemplateMetadataField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsTemplatesTemplateMetadataField(val *ModelsTemplatesTemplateMetadataField) *NullableModelsTemplatesTemplateMetadataField {
	return &NullableModelsTemplatesTemplateMetadataField{value: val, isSet: true}
}

func (v NullableModelsTemplatesTemplateMetadataField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsTemplatesTemplateMetadataField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
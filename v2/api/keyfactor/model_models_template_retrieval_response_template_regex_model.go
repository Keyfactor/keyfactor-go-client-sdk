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

// checks if the ModelsTemplateRetrievalResponseTemplateRegexModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsTemplateRetrievalResponseTemplateRegexModel{}

// ModelsTemplateRetrievalResponseTemplateRegexModel struct for ModelsTemplateRetrievalResponseTemplateRegexModel
type ModelsTemplateRetrievalResponseTemplateRegexModel struct {
	TemplateId           *int32  `json:"TemplateId,omitempty"`
	SubjectPart          *string `json:"SubjectPart,omitempty"`
	Regex                *string `json:"Regex,omitempty"`
	Error                *string `json:"Error,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelsTemplateRetrievalResponseTemplateRegexModel ModelsTemplateRetrievalResponseTemplateRegexModel

// NewModelsTemplateRetrievalResponseTemplateRegexModel instantiates a new ModelsTemplateRetrievalResponseTemplateRegexModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsTemplateRetrievalResponseTemplateRegexModel() *ModelsTemplateRetrievalResponseTemplateRegexModel {
	this := ModelsTemplateRetrievalResponseTemplateRegexModel{}
	return &this
}

// NewModelsTemplateRetrievalResponseTemplateRegexModelWithDefaults instantiates a new ModelsTemplateRetrievalResponseTemplateRegexModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsTemplateRetrievalResponseTemplateRegexModelWithDefaults() *ModelsTemplateRetrievalResponseTemplateRegexModel {
	this := ModelsTemplateRetrievalResponseTemplateRegexModel{}
	return &this
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *ModelsTemplateRetrievalResponseTemplateRegexModel) GetTemplateId() int32 {
	if o == nil || isNil(o.TemplateId) {
		var ret int32
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsTemplateRetrievalResponseTemplateRegexModel) GetTemplateIdOk() (*int32, bool) {
	if o == nil || isNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *ModelsTemplateRetrievalResponseTemplateRegexModel) HasTemplateId() bool {
	if o != nil && !isNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given int32 and assigns it to the TemplateId field.
func (o *ModelsTemplateRetrievalResponseTemplateRegexModel) SetTemplateId(v int32) {
	o.TemplateId = &v
}

// GetSubjectPart returns the SubjectPart field value if set, zero value otherwise.
func (o *ModelsTemplateRetrievalResponseTemplateRegexModel) GetSubjectPart() string {
	if o == nil || isNil(o.SubjectPart) {
		var ret string
		return ret
	}
	return *o.SubjectPart
}

// GetSubjectPartOk returns a tuple with the SubjectPart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsTemplateRetrievalResponseTemplateRegexModel) GetSubjectPartOk() (*string, bool) {
	if o == nil || isNil(o.SubjectPart) {
		return nil, false
	}
	return o.SubjectPart, true
}

// HasSubjectPart returns a boolean if a field has been set.
func (o *ModelsTemplateRetrievalResponseTemplateRegexModel) HasSubjectPart() bool {
	if o != nil && !isNil(o.SubjectPart) {
		return true
	}

	return false
}

// SetSubjectPart gets a reference to the given string and assigns it to the SubjectPart field.
func (o *ModelsTemplateRetrievalResponseTemplateRegexModel) SetSubjectPart(v string) {
	o.SubjectPart = &v
}

// GetRegex returns the Regex field value if set, zero value otherwise.
func (o *ModelsTemplateRetrievalResponseTemplateRegexModel) GetRegex() string {
	if o == nil || isNil(o.Regex) {
		var ret string
		return ret
	}
	return *o.Regex
}

// GetRegexOk returns a tuple with the Regex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsTemplateRetrievalResponseTemplateRegexModel) GetRegexOk() (*string, bool) {
	if o == nil || isNil(o.Regex) {
		return nil, false
	}
	return o.Regex, true
}

// HasRegex returns a boolean if a field has been set.
func (o *ModelsTemplateRetrievalResponseTemplateRegexModel) HasRegex() bool {
	if o != nil && !isNil(o.Regex) {
		return true
	}

	return false
}

// SetRegex gets a reference to the given string and assigns it to the Regex field.
func (o *ModelsTemplateRetrievalResponseTemplateRegexModel) SetRegex(v string) {
	o.Regex = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ModelsTemplateRetrievalResponseTemplateRegexModel) GetError() string {
	if o == nil || isNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsTemplateRetrievalResponseTemplateRegexModel) GetErrorOk() (*string, bool) {
	if o == nil || isNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ModelsTemplateRetrievalResponseTemplateRegexModel) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *ModelsTemplateRetrievalResponseTemplateRegexModel) SetError(v string) {
	o.Error = &v
}

func (o ModelsTemplateRetrievalResponseTemplateRegexModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsTemplateRetrievalResponseTemplateRegexModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TemplateId) {
		toSerialize["TemplateId"] = o.TemplateId
	}
	if !isNil(o.SubjectPart) {
		toSerialize["SubjectPart"] = o.SubjectPart
	}
	if !isNil(o.Regex) {
		toSerialize["Regex"] = o.Regex
	}
	if !isNil(o.Error) {
		toSerialize["Error"] = o.Error
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModelsTemplateRetrievalResponseTemplateRegexModel) UnmarshalJSON(bytes []byte) (err error) {
	varModelsTemplateRetrievalResponseTemplateRegexModel := _ModelsTemplateRetrievalResponseTemplateRegexModel{}

	if err = json.Unmarshal(bytes, &varModelsTemplateRetrievalResponseTemplateRegexModel); err == nil {
		*o = ModelsTemplateRetrievalResponseTemplateRegexModel(varModelsTemplateRetrievalResponseTemplateRegexModel)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "TemplateId")
		delete(additionalProperties, "SubjectPart")
		delete(additionalProperties, "Regex")
		delete(additionalProperties, "Error")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelsTemplateRetrievalResponseTemplateRegexModel struct {
	value *ModelsTemplateRetrievalResponseTemplateRegexModel
	isSet bool
}

func (v NullableModelsTemplateRetrievalResponseTemplateRegexModel) Get() *ModelsTemplateRetrievalResponseTemplateRegexModel {
	return v.value
}

func (v *NullableModelsTemplateRetrievalResponseTemplateRegexModel) Set(val *ModelsTemplateRetrievalResponseTemplateRegexModel) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsTemplateRetrievalResponseTemplateRegexModel) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsTemplateRetrievalResponseTemplateRegexModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsTemplateRetrievalResponseTemplateRegexModel(val *ModelsTemplateRetrievalResponseTemplateRegexModel) *NullableModelsTemplateRetrievalResponseTemplateRegexModel {
	return &NullableModelsTemplateRetrievalResponseTemplateRegexModel{value: val, isSet: true}
}

func (v NullableModelsTemplateRetrievalResponseTemplateRegexModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsTemplateRetrievalResponseTemplateRegexModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

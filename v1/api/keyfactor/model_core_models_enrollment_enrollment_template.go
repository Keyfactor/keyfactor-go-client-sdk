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

// checks if the CoreModelsEnrollmentEnrollmentTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreModelsEnrollmentEnrollmentTemplate{}

// CoreModelsEnrollmentEnrollmentTemplate struct for CoreModelsEnrollmentEnrollmentTemplate
type CoreModelsEnrollmentEnrollmentTemplate struct {
	Id                   *int32                             `json:"Id,omitempty"`
	Name                 *string                            `json:"Name,omitempty"`
	DisplayName          *string                            `json:"DisplayName,omitempty"`
	Forest               *string                            `json:"Forest,omitempty"`
	KeySize              *string                            `json:"KeySize,omitempty"`
	KeyType              *string                            `json:"KeyType,omitempty"`
	RequiresApproval     *bool                              `json:"RequiresApproval,omitempty"`
	RFCEnforcement       *bool                              `json:"RFCEnforcement,omitempty"`
	CAs                  []CoreModelsEnrollmentEnrollmentCA `json:"CAs,omitempty"`
	EnrollmentFields     []ModelsTemplateEnrollmentField    `json:"EnrollmentFields,omitempty"`
	MetadataFields       []ModelsTemplateMetadataField      `json:"MetadataFields,omitempty"`
	Regexes              []ModelsTemplateRegex              `json:"Regexes,omitempty"`
	ExtendedKeyUsages    []ModelsExtendedKeyUsage           `json:"ExtendedKeyUsages,omitempty"`
	Curve                NullableString                     `json:"Curve,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CoreModelsEnrollmentEnrollmentTemplate CoreModelsEnrollmentEnrollmentTemplate

// NewCoreModelsEnrollmentEnrollmentTemplate instantiates a new CoreModelsEnrollmentEnrollmentTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreModelsEnrollmentEnrollmentTemplate() *CoreModelsEnrollmentEnrollmentTemplate {
	this := CoreModelsEnrollmentEnrollmentTemplate{}
	return &this
}

// NewCoreModelsEnrollmentEnrollmentTemplateWithDefaults instantiates a new CoreModelsEnrollmentEnrollmentTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreModelsEnrollmentEnrollmentTemplateWithDefaults() *CoreModelsEnrollmentEnrollmentTemplate {
	this := CoreModelsEnrollmentEnrollmentTemplate{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CoreModelsEnrollmentEnrollmentTemplate) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreModelsEnrollmentEnrollmentTemplate) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CoreModelsEnrollmentEnrollmentTemplate) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CoreModelsEnrollmentEnrollmentTemplate) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CoreModelsEnrollmentEnrollmentTemplate) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreModelsEnrollmentEnrollmentTemplate) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CoreModelsEnrollmentEnrollmentTemplate) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CoreModelsEnrollmentEnrollmentTemplate) SetName(v string) {
	o.Name = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *CoreModelsEnrollmentEnrollmentTemplate) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreModelsEnrollmentEnrollmentTemplate) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CoreModelsEnrollmentEnrollmentTemplate) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *CoreModelsEnrollmentEnrollmentTemplate) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetForest returns the Forest field value if set, zero value otherwise.
func (o *CoreModelsEnrollmentEnrollmentTemplate) GetForest() string {
	if o == nil || isNil(o.Forest) {
		var ret string
		return ret
	}
	return *o.Forest
}

// GetForestOk returns a tuple with the Forest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreModelsEnrollmentEnrollmentTemplate) GetForestOk() (*string, bool) {
	if o == nil || isNil(o.Forest) {
		return nil, false
	}
	return o.Forest, true
}

// HasForest returns a boolean if a field has been set.
func (o *CoreModelsEnrollmentEnrollmentTemplate) HasForest() bool {
	if o != nil && !isNil(o.Forest) {
		return true
	}

	return false
}

// SetForest gets a reference to the given string and assigns it to the Forest field.
func (o *CoreModelsEnrollmentEnrollmentTemplate) SetForest(v string) {
	o.Forest = &v
}

// GetKeySize returns the KeySize field value if set, zero value otherwise.
func (o *CoreModelsEnrollmentEnrollmentTemplate) GetKeySize() string {
	if o == nil || isNil(o.KeySize) {
		var ret string
		return ret
	}
	return *o.KeySize
}

// GetKeySizeOk returns a tuple with the KeySize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreModelsEnrollmentEnrollmentTemplate) GetKeySizeOk() (*string, bool) {
	if o == nil || isNil(o.KeySize) {
		return nil, false
	}
	return o.KeySize, true
}

// HasKeySize returns a boolean if a field has been set.
func (o *CoreModelsEnrollmentEnrollmentTemplate) HasKeySize() bool {
	if o != nil && !isNil(o.KeySize) {
		return true
	}

	return false
}

// SetKeySize gets a reference to the given string and assigns it to the KeySize field.
func (o *CoreModelsEnrollmentEnrollmentTemplate) SetKeySize(v string) {
	o.KeySize = &v
}

// GetKeyType returns the KeyType field value if set, zero value otherwise.
func (o *CoreModelsEnrollmentEnrollmentTemplate) GetKeyType() string {
	if o == nil || isNil(o.KeyType) {
		var ret string
		return ret
	}
	return *o.KeyType
}

// GetKeyTypeOk returns a tuple with the KeyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreModelsEnrollmentEnrollmentTemplate) GetKeyTypeOk() (*string, bool) {
	if o == nil || isNil(o.KeyType) {
		return nil, false
	}
	return o.KeyType, true
}

// HasKeyType returns a boolean if a field has been set.
func (o *CoreModelsEnrollmentEnrollmentTemplate) HasKeyType() bool {
	if o != nil && !isNil(o.KeyType) {
		return true
	}

	return false
}

// SetKeyType gets a reference to the given string and assigns it to the KeyType field.
func (o *CoreModelsEnrollmentEnrollmentTemplate) SetKeyType(v string) {
	o.KeyType = &v
}

// GetRequiresApproval returns the RequiresApproval field value if set, zero value otherwise.
func (o *CoreModelsEnrollmentEnrollmentTemplate) GetRequiresApproval() bool {
	if o == nil || isNil(o.RequiresApproval) {
		var ret bool
		return ret
	}
	return *o.RequiresApproval
}

// GetRequiresApprovalOk returns a tuple with the RequiresApproval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreModelsEnrollmentEnrollmentTemplate) GetRequiresApprovalOk() (*bool, bool) {
	if o == nil || isNil(o.RequiresApproval) {
		return nil, false
	}
	return o.RequiresApproval, true
}

// HasRequiresApproval returns a boolean if a field has been set.
func (o *CoreModelsEnrollmentEnrollmentTemplate) HasRequiresApproval() bool {
	if o != nil && !isNil(o.RequiresApproval) {
		return true
	}

	return false
}

// SetRequiresApproval gets a reference to the given bool and assigns it to the RequiresApproval field.
func (o *CoreModelsEnrollmentEnrollmentTemplate) SetRequiresApproval(v bool) {
	o.RequiresApproval = &v
}

// GetRFCEnforcement returns the RFCEnforcement field value if set, zero value otherwise.
func (o *CoreModelsEnrollmentEnrollmentTemplate) GetRFCEnforcement() bool {
	if o == nil || isNil(o.RFCEnforcement) {
		var ret bool
		return ret
	}
	return *o.RFCEnforcement
}

// GetRFCEnforcementOk returns a tuple with the RFCEnforcement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreModelsEnrollmentEnrollmentTemplate) GetRFCEnforcementOk() (*bool, bool) {
	if o == nil || isNil(o.RFCEnforcement) {
		return nil, false
	}
	return o.RFCEnforcement, true
}

// HasRFCEnforcement returns a boolean if a field has been set.
func (o *CoreModelsEnrollmentEnrollmentTemplate) HasRFCEnforcement() bool {
	if o != nil && !isNil(o.RFCEnforcement) {
		return true
	}

	return false
}

// SetRFCEnforcement gets a reference to the given bool and assigns it to the RFCEnforcement field.
func (o *CoreModelsEnrollmentEnrollmentTemplate) SetRFCEnforcement(v bool) {
	o.RFCEnforcement = &v
}

// GetCAs returns the CAs field value if set, zero value otherwise.
func (o *CoreModelsEnrollmentEnrollmentTemplate) GetCAs() []CoreModelsEnrollmentEnrollmentCA {
	if o == nil || isNil(o.CAs) {
		var ret []CoreModelsEnrollmentEnrollmentCA
		return ret
	}
	return o.CAs
}

// GetCAsOk returns a tuple with the CAs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreModelsEnrollmentEnrollmentTemplate) GetCAsOk() ([]CoreModelsEnrollmentEnrollmentCA, bool) {
	if o == nil || isNil(o.CAs) {
		return nil, false
	}
	return o.CAs, true
}

// HasCAs returns a boolean if a field has been set.
func (o *CoreModelsEnrollmentEnrollmentTemplate) HasCAs() bool {
	if o != nil && !isNil(o.CAs) {
		return true
	}

	return false
}

// SetCAs gets a reference to the given []CoreModelsEnrollmentEnrollmentCA and assigns it to the CAs field.
func (o *CoreModelsEnrollmentEnrollmentTemplate) SetCAs(v []CoreModelsEnrollmentEnrollmentCA) {
	o.CAs = v
}

// GetEnrollmentFields returns the EnrollmentFields field value if set, zero value otherwise.
func (o *CoreModelsEnrollmentEnrollmentTemplate) GetEnrollmentFields() []ModelsTemplateEnrollmentField {
	if o == nil || isNil(o.EnrollmentFields) {
		var ret []ModelsTemplateEnrollmentField
		return ret
	}
	return o.EnrollmentFields
}

// GetEnrollmentFieldsOk returns a tuple with the EnrollmentFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreModelsEnrollmentEnrollmentTemplate) GetEnrollmentFieldsOk() ([]ModelsTemplateEnrollmentField, bool) {
	if o == nil || isNil(o.EnrollmentFields) {
		return nil, false
	}
	return o.EnrollmentFields, true
}

// HasEnrollmentFields returns a boolean if a field has been set.
func (o *CoreModelsEnrollmentEnrollmentTemplate) HasEnrollmentFields() bool {
	if o != nil && !isNil(o.EnrollmentFields) {
		return true
	}

	return false
}

// SetEnrollmentFields gets a reference to the given []ModelsTemplateEnrollmentField and assigns it to the EnrollmentFields field.
func (o *CoreModelsEnrollmentEnrollmentTemplate) SetEnrollmentFields(v []ModelsTemplateEnrollmentField) {
	o.EnrollmentFields = v
}

// GetMetadataFields returns the MetadataFields field value if set, zero value otherwise.
func (o *CoreModelsEnrollmentEnrollmentTemplate) GetMetadataFields() []ModelsTemplateMetadataField {
	if o == nil || isNil(o.MetadataFields) {
		var ret []ModelsTemplateMetadataField
		return ret
	}
	return o.MetadataFields
}

// GetMetadataFieldsOk returns a tuple with the MetadataFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreModelsEnrollmentEnrollmentTemplate) GetMetadataFieldsOk() ([]ModelsTemplateMetadataField, bool) {
	if o == nil || isNil(o.MetadataFields) {
		return nil, false
	}
	return o.MetadataFields, true
}

// HasMetadataFields returns a boolean if a field has been set.
func (o *CoreModelsEnrollmentEnrollmentTemplate) HasMetadataFields() bool {
	if o != nil && !isNil(o.MetadataFields) {
		return true
	}

	return false
}

// SetMetadataFields gets a reference to the given []ModelsTemplateMetadataField and assigns it to the MetadataFields field.
func (o *CoreModelsEnrollmentEnrollmentTemplate) SetMetadataFields(v []ModelsTemplateMetadataField) {
	o.MetadataFields = v
}

// GetRegexes returns the Regexes field value if set, zero value otherwise.
func (o *CoreModelsEnrollmentEnrollmentTemplate) GetRegexes() []ModelsTemplateRegex {
	if o == nil || isNil(o.Regexes) {
		var ret []ModelsTemplateRegex
		return ret
	}
	return o.Regexes
}

// GetRegexesOk returns a tuple with the Regexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreModelsEnrollmentEnrollmentTemplate) GetRegexesOk() ([]ModelsTemplateRegex, bool) {
	if o == nil || isNil(o.Regexes) {
		return nil, false
	}
	return o.Regexes, true
}

// HasRegexes returns a boolean if a field has been set.
func (o *CoreModelsEnrollmentEnrollmentTemplate) HasRegexes() bool {
	if o != nil && !isNil(o.Regexes) {
		return true
	}

	return false
}

// SetRegexes gets a reference to the given []ModelsTemplateRegex and assigns it to the Regexes field.
func (o *CoreModelsEnrollmentEnrollmentTemplate) SetRegexes(v []ModelsTemplateRegex) {
	o.Regexes = v
}

// GetExtendedKeyUsages returns the ExtendedKeyUsages field value if set, zero value otherwise.
func (o *CoreModelsEnrollmentEnrollmentTemplate) GetExtendedKeyUsages() []ModelsExtendedKeyUsage {
	if o == nil || isNil(o.ExtendedKeyUsages) {
		var ret []ModelsExtendedKeyUsage
		return ret
	}
	return o.ExtendedKeyUsages
}

// GetExtendedKeyUsagesOk returns a tuple with the ExtendedKeyUsages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreModelsEnrollmentEnrollmentTemplate) GetExtendedKeyUsagesOk() ([]ModelsExtendedKeyUsage, bool) {
	if o == nil || isNil(o.ExtendedKeyUsages) {
		return nil, false
	}
	return o.ExtendedKeyUsages, true
}

// HasExtendedKeyUsages returns a boolean if a field has been set.
func (o *CoreModelsEnrollmentEnrollmentTemplate) HasExtendedKeyUsages() bool {
	if o != nil && !isNil(o.ExtendedKeyUsages) {
		return true
	}

	return false
}

// SetExtendedKeyUsages gets a reference to the given []ModelsExtendedKeyUsage and assigns it to the ExtendedKeyUsages field.
func (o *CoreModelsEnrollmentEnrollmentTemplate) SetExtendedKeyUsages(v []ModelsExtendedKeyUsage) {
	o.ExtendedKeyUsages = v
}

// GetCurve returns the Curve field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CoreModelsEnrollmentEnrollmentTemplate) GetCurve() string {
	if o == nil || isNil(o.Curve.Get()) {
		var ret string
		return ret
	}
	return *o.Curve.Get()
}

// GetCurveOk returns a tuple with the Curve field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CoreModelsEnrollmentEnrollmentTemplate) GetCurveOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Curve.Get(), o.Curve.IsSet()
}

// HasCurve returns a boolean if a field has been set.
func (o *CoreModelsEnrollmentEnrollmentTemplate) HasCurve() bool {
	if o != nil && o.Curve.IsSet() {
		return true
	}

	return false
}

// SetCurve gets a reference to the given NullableString and assigns it to the Curve field.
func (o *CoreModelsEnrollmentEnrollmentTemplate) SetCurve(v string) {
	o.Curve.Set(&v)
}

// SetCurveNil sets the value for Curve to be an explicit nil
func (o *CoreModelsEnrollmentEnrollmentTemplate) SetCurveNil() {
	o.Curve.Set(nil)
}

// UnsetCurve ensures that no value is present for Curve, not even an explicit nil
func (o *CoreModelsEnrollmentEnrollmentTemplate) UnsetCurve() {
	o.Curve.Unset()
}

func (o CoreModelsEnrollmentEnrollmentTemplate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreModelsEnrollmentEnrollmentTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !isNil(o.DisplayName) {
		toSerialize["DisplayName"] = o.DisplayName
	}
	if !isNil(o.Forest) {
		toSerialize["Forest"] = o.Forest
	}
	if !isNil(o.KeySize) {
		toSerialize["KeySize"] = o.KeySize
	}
	if !isNil(o.KeyType) {
		toSerialize["KeyType"] = o.KeyType
	}
	if !isNil(o.RequiresApproval) {
		toSerialize["RequiresApproval"] = o.RequiresApproval
	}
	if !isNil(o.RFCEnforcement) {
		toSerialize["RFCEnforcement"] = o.RFCEnforcement
	}
	if !isNil(o.CAs) {
		toSerialize["CAs"] = o.CAs
	}
	if !isNil(o.EnrollmentFields) {
		toSerialize["EnrollmentFields"] = o.EnrollmentFields
	}
	if !isNil(o.MetadataFields) {
		toSerialize["MetadataFields"] = o.MetadataFields
	}
	if !isNil(o.Regexes) {
		toSerialize["Regexes"] = o.Regexes
	}
	if !isNil(o.ExtendedKeyUsages) {
		toSerialize["ExtendedKeyUsages"] = o.ExtendedKeyUsages
	}
	if o.Curve.IsSet() {
		toSerialize["Curve"] = o.Curve.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CoreModelsEnrollmentEnrollmentTemplate) UnmarshalJSON(bytes []byte) (err error) {
	varCoreModelsEnrollmentEnrollmentTemplate := _CoreModelsEnrollmentEnrollmentTemplate{}

	if err = json.Unmarshal(bytes, &varCoreModelsEnrollmentEnrollmentTemplate); err == nil {
		*o = CoreModelsEnrollmentEnrollmentTemplate(varCoreModelsEnrollmentEnrollmentTemplate)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Id")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "DisplayName")
		delete(additionalProperties, "Forest")
		delete(additionalProperties, "KeySize")
		delete(additionalProperties, "KeyType")
		delete(additionalProperties, "RequiresApproval")
		delete(additionalProperties, "RFCEnforcement")
		delete(additionalProperties, "CAs")
		delete(additionalProperties, "EnrollmentFields")
		delete(additionalProperties, "MetadataFields")
		delete(additionalProperties, "Regexes")
		delete(additionalProperties, "ExtendedKeyUsages")
		delete(additionalProperties, "Curve")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCoreModelsEnrollmentEnrollmentTemplate struct {
	value *CoreModelsEnrollmentEnrollmentTemplate
	isSet bool
}

func (v NullableCoreModelsEnrollmentEnrollmentTemplate) Get() *CoreModelsEnrollmentEnrollmentTemplate {
	return v.value
}

func (v *NullableCoreModelsEnrollmentEnrollmentTemplate) Set(val *CoreModelsEnrollmentEnrollmentTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreModelsEnrollmentEnrollmentTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreModelsEnrollmentEnrollmentTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreModelsEnrollmentEnrollmentTemplate(val *CoreModelsEnrollmentEnrollmentTemplate) *NullableCoreModelsEnrollmentEnrollmentTemplate {
	return &NullableCoreModelsEnrollmentEnrollmentTemplate{value: val, isSet: true}
}

func (v NullableCoreModelsEnrollmentEnrollmentTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreModelsEnrollmentEnrollmentTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
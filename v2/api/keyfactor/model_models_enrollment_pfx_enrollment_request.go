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

// checks if the ModelsEnrollmentPFXEnrollmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsEnrollmentPFXEnrollmentRequest{}

// ModelsEnrollmentPFXEnrollmentRequest struct for ModelsEnrollmentPFXEnrollmentRequest
type ModelsEnrollmentPFXEnrollmentRequest struct {
	CustomFriendlyName          *string                           `json:"CustomFriendlyName,omitempty"`
	Password                    *string                           `json:"Password,omitempty"`
	PopulateMissingValuesFromAD *bool                             `json:"PopulateMissingValuesFromAD,omitempty"`
	Subject                     *string                           `json:"Subject,omitempty"`
	IncludeChain                *bool                             `json:"IncludeChain,omitempty"`
	RenewalCertificateId        *int32                            `json:"RenewalCertificateId,omitempty"`
	CertificateAuthority        *string                           `json:"CertificateAuthority,omitempty"`
	Metadata                    map[string]interface{}            `json:"Metadata,omitempty"`
	AdditionalEnrollmentFields  map[string]map[string]interface{} `json:"AdditionalEnrollmentFields,omitempty"`
	Timestamp                   *time.Time                        `json:"Timestamp,omitempty"`
	Template                    *string                           `json:"Template,omitempty"`
	SANs                        *map[string][]string              `json:"SANs,omitempty"`
	AdditionalProperties        map[string]interface{}
}

type _ModelsEnrollmentPFXEnrollmentRequest ModelsEnrollmentPFXEnrollmentRequest

// NewModelsEnrollmentPFXEnrollmentRequest instantiates a new ModelsEnrollmentPFXEnrollmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsEnrollmentPFXEnrollmentRequest() *ModelsEnrollmentPFXEnrollmentRequest {
	this := ModelsEnrollmentPFXEnrollmentRequest{}
	return &this
}

// NewModelsEnrollmentPFXEnrollmentRequestWithDefaults instantiates a new ModelsEnrollmentPFXEnrollmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsEnrollmentPFXEnrollmentRequestWithDefaults() *ModelsEnrollmentPFXEnrollmentRequest {
	this := ModelsEnrollmentPFXEnrollmentRequest{}
	return &this
}

// GetCustomFriendlyName returns the CustomFriendlyName field value if set, zero value otherwise.
func (o *ModelsEnrollmentPFXEnrollmentRequest) GetCustomFriendlyName() string {
	if o == nil || isNil(o.CustomFriendlyName) {
		var ret string
		return ret
	}
	return *o.CustomFriendlyName
}

// GetCustomFriendlyNameOk returns a tuple with the CustomFriendlyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsEnrollmentPFXEnrollmentRequest) GetCustomFriendlyNameOk() (*string, bool) {
	if o == nil || isNil(o.CustomFriendlyName) {
		return nil, false
	}
	return o.CustomFriendlyName, true
}

// HasCustomFriendlyName returns a boolean if a field has been set.
func (o *ModelsEnrollmentPFXEnrollmentRequest) HasCustomFriendlyName() bool {
	if o != nil && !isNil(o.CustomFriendlyName) {
		return true
	}

	return false
}

// SetCustomFriendlyName gets a reference to the given string and assigns it to the CustomFriendlyName field.
func (o *ModelsEnrollmentPFXEnrollmentRequest) SetCustomFriendlyName(v string) {
	o.CustomFriendlyName = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ModelsEnrollmentPFXEnrollmentRequest) GetPassword() string {
	if o == nil || isNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsEnrollmentPFXEnrollmentRequest) GetPasswordOk() (*string, bool) {
	if o == nil || isNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ModelsEnrollmentPFXEnrollmentRequest) HasPassword() bool {
	if o != nil && !isNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ModelsEnrollmentPFXEnrollmentRequest) SetPassword(v string) {
	o.Password = &v
}

// GetPopulateMissingValuesFromAD returns the PopulateMissingValuesFromAD field value if set, zero value otherwise.
func (o *ModelsEnrollmentPFXEnrollmentRequest) GetPopulateMissingValuesFromAD() bool {
	if o == nil || isNil(o.PopulateMissingValuesFromAD) {
		var ret bool
		return ret
	}
	return *o.PopulateMissingValuesFromAD
}

// GetPopulateMissingValuesFromADOk returns a tuple with the PopulateMissingValuesFromAD field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsEnrollmentPFXEnrollmentRequest) GetPopulateMissingValuesFromADOk() (*bool, bool) {
	if o == nil || isNil(o.PopulateMissingValuesFromAD) {
		return nil, false
	}
	return o.PopulateMissingValuesFromAD, true
}

// HasPopulateMissingValuesFromAD returns a boolean if a field has been set.
func (o *ModelsEnrollmentPFXEnrollmentRequest) HasPopulateMissingValuesFromAD() bool {
	if o != nil && !isNil(o.PopulateMissingValuesFromAD) {
		return true
	}

	return false
}

// SetPopulateMissingValuesFromAD gets a reference to the given bool and assigns it to the PopulateMissingValuesFromAD field.
func (o *ModelsEnrollmentPFXEnrollmentRequest) SetPopulateMissingValuesFromAD(v bool) {
	o.PopulateMissingValuesFromAD = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *ModelsEnrollmentPFXEnrollmentRequest) GetSubject() string {
	if o == nil || isNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsEnrollmentPFXEnrollmentRequest) GetSubjectOk() (*string, bool) {
	if o == nil || isNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *ModelsEnrollmentPFXEnrollmentRequest) HasSubject() bool {
	if o != nil && !isNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *ModelsEnrollmentPFXEnrollmentRequest) SetSubject(v string) {
	o.Subject = &v
}

// GetIncludeChain returns the IncludeChain field value if set, zero value otherwise.
func (o *ModelsEnrollmentPFXEnrollmentRequest) GetIncludeChain() bool {
	if o == nil || isNil(o.IncludeChain) {
		var ret bool
		return ret
	}
	return *o.IncludeChain
}

// GetIncludeChainOk returns a tuple with the IncludeChain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsEnrollmentPFXEnrollmentRequest) GetIncludeChainOk() (*bool, bool) {
	if o == nil || isNil(o.IncludeChain) {
		return nil, false
	}
	return o.IncludeChain, true
}

// HasIncludeChain returns a boolean if a field has been set.
func (o *ModelsEnrollmentPFXEnrollmentRequest) HasIncludeChain() bool {
	if o != nil && !isNil(o.IncludeChain) {
		return true
	}

	return false
}

// SetIncludeChain gets a reference to the given bool and assigns it to the IncludeChain field.
func (o *ModelsEnrollmentPFXEnrollmentRequest) SetIncludeChain(v bool) {
	o.IncludeChain = &v
}

// GetRenewalCertificateId returns the RenewalCertificateId field value if set, zero value otherwise.
func (o *ModelsEnrollmentPFXEnrollmentRequest) GetRenewalCertificateId() int32 {
	if o == nil || isNil(o.RenewalCertificateId) {
		var ret int32
		return ret
	}
	return *o.RenewalCertificateId
}

// GetRenewalCertificateIdOk returns a tuple with the RenewalCertificateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsEnrollmentPFXEnrollmentRequest) GetRenewalCertificateIdOk() (*int32, bool) {
	if o == nil || isNil(o.RenewalCertificateId) {
		return nil, false
	}
	return o.RenewalCertificateId, true
}

// HasRenewalCertificateId returns a boolean if a field has been set.
func (o *ModelsEnrollmentPFXEnrollmentRequest) HasRenewalCertificateId() bool {
	if o != nil && !isNil(o.RenewalCertificateId) {
		return true
	}

	return false
}

// SetRenewalCertificateId gets a reference to the given int32 and assigns it to the RenewalCertificateId field.
func (o *ModelsEnrollmentPFXEnrollmentRequest) SetRenewalCertificateId(v int32) {
	o.RenewalCertificateId = &v
}

// GetCertificateAuthority returns the CertificateAuthority field value if set, zero value otherwise.
func (o *ModelsEnrollmentPFXEnrollmentRequest) GetCertificateAuthority() string {
	if o == nil || isNil(o.CertificateAuthority) {
		var ret string
		return ret
	}
	return *o.CertificateAuthority
}

// GetCertificateAuthorityOk returns a tuple with the CertificateAuthority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsEnrollmentPFXEnrollmentRequest) GetCertificateAuthorityOk() (*string, bool) {
	if o == nil || isNil(o.CertificateAuthority) {
		return nil, false
	}
	return o.CertificateAuthority, true
}

// HasCertificateAuthority returns a boolean if a field has been set.
func (o *ModelsEnrollmentPFXEnrollmentRequest) HasCertificateAuthority() bool {
	if o != nil && !isNil(o.CertificateAuthority) {
		return true
	}

	return false
}

// SetCertificateAuthority gets a reference to the given string and assigns it to the CertificateAuthority field.
func (o *ModelsEnrollmentPFXEnrollmentRequest) SetCertificateAuthority(v string) {
	o.CertificateAuthority = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ModelsEnrollmentPFXEnrollmentRequest) GetMetadata() map[string]interface{} {
	if o == nil || isNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsEnrollmentPFXEnrollmentRequest) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ModelsEnrollmentPFXEnrollmentRequest) HasMetadata() bool {
	if o != nil && !isNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *ModelsEnrollmentPFXEnrollmentRequest) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetAdditionalEnrollmentFields returns the AdditionalEnrollmentFields field value if set, zero value otherwise.
func (o *ModelsEnrollmentPFXEnrollmentRequest) GetAdditionalEnrollmentFields() map[string]map[string]interface{} {
	if o == nil || isNil(o.AdditionalEnrollmentFields) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.AdditionalEnrollmentFields
}

// GetAdditionalEnrollmentFieldsOk returns a tuple with the AdditionalEnrollmentFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsEnrollmentPFXEnrollmentRequest) GetAdditionalEnrollmentFieldsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || isNil(o.AdditionalEnrollmentFields) {
		return map[string]map[string]interface{}{}, false
	}
	return o.AdditionalEnrollmentFields, true
}

// HasAdditionalEnrollmentFields returns a boolean if a field has been set.
func (o *ModelsEnrollmentPFXEnrollmentRequest) HasAdditionalEnrollmentFields() bool {
	if o != nil && !isNil(o.AdditionalEnrollmentFields) {
		return true
	}

	return false
}

// SetAdditionalEnrollmentFields gets a reference to the given map[string]map[string]interface{} and assigns it to the AdditionalEnrollmentFields field.
func (o *ModelsEnrollmentPFXEnrollmentRequest) SetAdditionalEnrollmentFields(v map[string]map[string]interface{}) {
	o.AdditionalEnrollmentFields = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ModelsEnrollmentPFXEnrollmentRequest) GetTimestamp() time.Time {
	if o == nil || isNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsEnrollmentPFXEnrollmentRequest) GetTimestampOk() (*time.Time, bool) {
	if o == nil || isNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ModelsEnrollmentPFXEnrollmentRequest) HasTimestamp() bool {
	if o != nil && !isNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *ModelsEnrollmentPFXEnrollmentRequest) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *ModelsEnrollmentPFXEnrollmentRequest) GetTemplate() string {
	if o == nil || isNil(o.Template) {
		var ret string
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsEnrollmentPFXEnrollmentRequest) GetTemplateOk() (*string, bool) {
	if o == nil || isNil(o.Template) {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *ModelsEnrollmentPFXEnrollmentRequest) HasTemplate() bool {
	if o != nil && !isNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given string and assigns it to the Template field.
func (o *ModelsEnrollmentPFXEnrollmentRequest) SetTemplate(v string) {
	o.Template = &v
}

// GetSANs returns the SANs field value if set, zero value otherwise.
func (o *ModelsEnrollmentPFXEnrollmentRequest) GetSANs() map[string][]string {
	if o == nil || isNil(o.SANs) {
		var ret map[string][]string
		return ret
	}
	return *o.SANs
}

// GetSANsOk returns a tuple with the SANs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsEnrollmentPFXEnrollmentRequest) GetSANsOk() (*map[string][]string, bool) {
	if o == nil || isNil(o.SANs) {
		return nil, false
	}
	return o.SANs, true
}

// HasSANs returns a boolean if a field has been set.
func (o *ModelsEnrollmentPFXEnrollmentRequest) HasSANs() bool {
	if o != nil && !isNil(o.SANs) {
		return true
	}

	return false
}

// SetSANs gets a reference to the given map[string][]string and assigns it to the SANs field.
func (o *ModelsEnrollmentPFXEnrollmentRequest) SetSANs(v map[string][]string) {
	o.SANs = &v
}

func (o ModelsEnrollmentPFXEnrollmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsEnrollmentPFXEnrollmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CustomFriendlyName) {
		toSerialize["CustomFriendlyName"] = o.CustomFriendlyName
	}
	if !isNil(o.Password) {
		toSerialize["Password"] = o.Password
	}
	if !isNil(o.PopulateMissingValuesFromAD) {
		toSerialize["PopulateMissingValuesFromAD"] = o.PopulateMissingValuesFromAD
	}
	if !isNil(o.Subject) {
		toSerialize["Subject"] = o.Subject
	}
	if !isNil(o.IncludeChain) {
		toSerialize["IncludeChain"] = o.IncludeChain
	}
	if !isNil(o.RenewalCertificateId) {
		toSerialize["RenewalCertificateId"] = o.RenewalCertificateId
	}
	if !isNil(o.CertificateAuthority) {
		toSerialize["CertificateAuthority"] = o.CertificateAuthority
	}
	if !isNil(o.Metadata) {
		toSerialize["Metadata"] = o.Metadata
	}
	if !isNil(o.AdditionalEnrollmentFields) {
		toSerialize["AdditionalEnrollmentFields"] = o.AdditionalEnrollmentFields
	}
	if !isNil(o.Timestamp) {
		toSerialize["Timestamp"] = o.Timestamp
	}
	if !isNil(o.Template) {
		toSerialize["Template"] = o.Template
	}
	if !isNil(o.SANs) {
		toSerialize["SANs"] = o.SANs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModelsEnrollmentPFXEnrollmentRequest) UnmarshalJSON(bytes []byte) (err error) {
	varModelsEnrollmentPFXEnrollmentRequest := _ModelsEnrollmentPFXEnrollmentRequest{}

	if err = json.Unmarshal(bytes, &varModelsEnrollmentPFXEnrollmentRequest); err == nil {
		*o = ModelsEnrollmentPFXEnrollmentRequest(varModelsEnrollmentPFXEnrollmentRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "CustomFriendlyName")
		delete(additionalProperties, "Password")
		delete(additionalProperties, "PopulateMissingValuesFromAD")
		delete(additionalProperties, "Subject")
		delete(additionalProperties, "IncludeChain")
		delete(additionalProperties, "RenewalCertificateId")
		delete(additionalProperties, "CertificateAuthority")
		delete(additionalProperties, "Metadata")
		delete(additionalProperties, "AdditionalEnrollmentFields")
		delete(additionalProperties, "Timestamp")
		delete(additionalProperties, "Template")
		delete(additionalProperties, "SANs")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelsEnrollmentPFXEnrollmentRequest struct {
	value *ModelsEnrollmentPFXEnrollmentRequest
	isSet bool
}

func (v NullableModelsEnrollmentPFXEnrollmentRequest) Get() *ModelsEnrollmentPFXEnrollmentRequest {
	return v.value
}

func (v *NullableModelsEnrollmentPFXEnrollmentRequest) Set(val *ModelsEnrollmentPFXEnrollmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsEnrollmentPFXEnrollmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsEnrollmentPFXEnrollmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsEnrollmentPFXEnrollmentRequest(val *ModelsEnrollmentPFXEnrollmentRequest) *NullableModelsEnrollmentPFXEnrollmentRequest {
	return &NullableModelsEnrollmentPFXEnrollmentRequest{value: val, isSet: true}
}

func (v NullableModelsEnrollmentPFXEnrollmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsEnrollmentPFXEnrollmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

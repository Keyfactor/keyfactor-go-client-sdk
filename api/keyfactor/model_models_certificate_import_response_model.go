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

// checks if the ModelsCertificateImportResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsCertificateImportResponseModel{}

// ModelsCertificateImportResponseModel struct for ModelsCertificateImportResponseModel
type ModelsCertificateImportResponseModel struct {
	ImportStatus *int32 `json:"ImportStatus,omitempty"`
	JobStatus *int32 `json:"JobStatus,omitempty"`
	InvalidKeystores []ModelsInvalidKeystore `json:"InvalidKeystores,omitempty"`
	Thumbprint *string `json:"Thumbprint,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelsCertificateImportResponseModel ModelsCertificateImportResponseModel

// NewModelsCertificateImportResponseModel instantiates a new ModelsCertificateImportResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsCertificateImportResponseModel() *ModelsCertificateImportResponseModel {
	this := ModelsCertificateImportResponseModel{}
	return &this
}

// NewModelsCertificateImportResponseModelWithDefaults instantiates a new ModelsCertificateImportResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsCertificateImportResponseModelWithDefaults() *ModelsCertificateImportResponseModel {
	this := ModelsCertificateImportResponseModel{}
	return &this
}

// GetImportStatus returns the ImportStatus field value if set, zero value otherwise.
func (o *ModelsCertificateImportResponseModel) GetImportStatus() int32 {
	if o == nil || isNil(o.ImportStatus) {
		var ret int32
		return ret
	}
	return *o.ImportStatus
}

// GetImportStatusOk returns a tuple with the ImportStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertificateImportResponseModel) GetImportStatusOk() (*int32, bool) {
	if o == nil || isNil(o.ImportStatus) {
		return nil, false
	}
	return o.ImportStatus, true
}

// HasImportStatus returns a boolean if a field has been set.
func (o *ModelsCertificateImportResponseModel) HasImportStatus() bool {
	if o != nil && !isNil(o.ImportStatus) {
		return true
	}

	return false
}

// SetImportStatus gets a reference to the given int32 and assigns it to the ImportStatus field.
func (o *ModelsCertificateImportResponseModel) SetImportStatus(v int32) {
	o.ImportStatus = &v
}

// GetJobStatus returns the JobStatus field value if set, zero value otherwise.
func (o *ModelsCertificateImportResponseModel) GetJobStatus() int32 {
	if o == nil || isNil(o.JobStatus) {
		var ret int32
		return ret
	}
	return *o.JobStatus
}

// GetJobStatusOk returns a tuple with the JobStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertificateImportResponseModel) GetJobStatusOk() (*int32, bool) {
	if o == nil || isNil(o.JobStatus) {
		return nil, false
	}
	return o.JobStatus, true
}

// HasJobStatus returns a boolean if a field has been set.
func (o *ModelsCertificateImportResponseModel) HasJobStatus() bool {
	if o != nil && !isNil(o.JobStatus) {
		return true
	}

	return false
}

// SetJobStatus gets a reference to the given int32 and assigns it to the JobStatus field.
func (o *ModelsCertificateImportResponseModel) SetJobStatus(v int32) {
	o.JobStatus = &v
}

// GetInvalidKeystores returns the InvalidKeystores field value if set, zero value otherwise.
func (o *ModelsCertificateImportResponseModel) GetInvalidKeystores() []ModelsInvalidKeystore {
	if o == nil || isNil(o.InvalidKeystores) {
		var ret []ModelsInvalidKeystore
		return ret
	}
	return o.InvalidKeystores
}

// GetInvalidKeystoresOk returns a tuple with the InvalidKeystores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertificateImportResponseModel) GetInvalidKeystoresOk() ([]ModelsInvalidKeystore, bool) {
	if o == nil || isNil(o.InvalidKeystores) {
		return nil, false
	}
	return o.InvalidKeystores, true
}

// HasInvalidKeystores returns a boolean if a field has been set.
func (o *ModelsCertificateImportResponseModel) HasInvalidKeystores() bool {
	if o != nil && !isNil(o.InvalidKeystores) {
		return true
	}

	return false
}

// SetInvalidKeystores gets a reference to the given []ModelsInvalidKeystore and assigns it to the InvalidKeystores field.
func (o *ModelsCertificateImportResponseModel) SetInvalidKeystores(v []ModelsInvalidKeystore) {
	o.InvalidKeystores = v
}

// GetThumbprint returns the Thumbprint field value if set, zero value otherwise.
func (o *ModelsCertificateImportResponseModel) GetThumbprint() string {
	if o == nil || isNil(o.Thumbprint) {
		var ret string
		return ret
	}
	return *o.Thumbprint
}

// GetThumbprintOk returns a tuple with the Thumbprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertificateImportResponseModel) GetThumbprintOk() (*string, bool) {
	if o == nil || isNil(o.Thumbprint) {
		return nil, false
	}
	return o.Thumbprint, true
}

// HasThumbprint returns a boolean if a field has been set.
func (o *ModelsCertificateImportResponseModel) HasThumbprint() bool {
	if o != nil && !isNil(o.Thumbprint) {
		return true
	}

	return false
}

// SetThumbprint gets a reference to the given string and assigns it to the Thumbprint field.
func (o *ModelsCertificateImportResponseModel) SetThumbprint(v string) {
	o.Thumbprint = &v
}

func (o ModelsCertificateImportResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsCertificateImportResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ImportStatus) {
		toSerialize["ImportStatus"] = o.ImportStatus
	}
	if !isNil(o.JobStatus) {
		toSerialize["JobStatus"] = o.JobStatus
	}
	if !isNil(o.InvalidKeystores) {
		toSerialize["InvalidKeystores"] = o.InvalidKeystores
	}
	if !isNil(o.Thumbprint) {
		toSerialize["Thumbprint"] = o.Thumbprint
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModelsCertificateImportResponseModel) UnmarshalJSON(bytes []byte) (err error) {
	varModelsCertificateImportResponseModel := _ModelsCertificateImportResponseModel{}

	if err = json.Unmarshal(bytes, &varModelsCertificateImportResponseModel); err == nil {
		*o = ModelsCertificateImportResponseModel(varModelsCertificateImportResponseModel)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ImportStatus")
		delete(additionalProperties, "JobStatus")
		delete(additionalProperties, "InvalidKeystores")
		delete(additionalProperties, "Thumbprint")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelsCertificateImportResponseModel struct {
	value *ModelsCertificateImportResponseModel
	isSet bool
}

func (v NullableModelsCertificateImportResponseModel) Get() *ModelsCertificateImportResponseModel {
	return v.value
}

func (v *NullableModelsCertificateImportResponseModel) Set(val *ModelsCertificateImportResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsCertificateImportResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsCertificateImportResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsCertificateImportResponseModel(val *ModelsCertificateImportResponseModel) *NullableModelsCertificateImportResponseModel {
	return &NullableModelsCertificateImportResponseModel{value: val, isSet: true}
}

func (v NullableModelsCertificateImportResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsCertificateImportResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


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

// checks if the ModelsCertificateStoreEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsCertificateStoreEntry{}

// ModelsCertificateStoreEntry struct for ModelsCertificateStoreEntry
type ModelsCertificateStoreEntry struct {
	CertificateStoreId string                            `json:"CertificateStoreId"`
	Alias              *string                           `json:"Alias,omitempty"`
	JobFields          map[string]map[string]interface{} `json:"JobFields,omitempty"`
	Overwrite          *bool                             `json:"Overwrite,omitempty"`
	EntryPassword      *ModelsKeyfactorAPISecret         `json:"EntryPassword,omitempty"`
	// The PFX password.
	PfxPassword          *string `json:"PfxPassword,omitempty"`
	IncludePrivateKey    *bool   `json:"IncludePrivateKey,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelsCertificateStoreEntry ModelsCertificateStoreEntry

// NewModelsCertificateStoreEntry instantiates a new ModelsCertificateStoreEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsCertificateStoreEntry(certificateStoreId string) *ModelsCertificateStoreEntry {
	this := ModelsCertificateStoreEntry{}
	this.CertificateStoreId = certificateStoreId
	return &this
}

// NewModelsCertificateStoreEntryWithDefaults instantiates a new ModelsCertificateStoreEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsCertificateStoreEntryWithDefaults() *ModelsCertificateStoreEntry {
	this := ModelsCertificateStoreEntry{}
	return &this
}

// GetCertificateStoreId returns the CertificateStoreId field value
func (o *ModelsCertificateStoreEntry) GetCertificateStoreId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CertificateStoreId
}

// GetCertificateStoreIdOk returns a tuple with the CertificateStoreId field value
// and a boolean to check if the value has been set.
func (o *ModelsCertificateStoreEntry) GetCertificateStoreIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CertificateStoreId, true
}

// SetCertificateStoreId sets field value
func (o *ModelsCertificateStoreEntry) SetCertificateStoreId(v string) {
	o.CertificateStoreId = v
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *ModelsCertificateStoreEntry) GetAlias() string {
	if o == nil || isNil(o.Alias) {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertificateStoreEntry) GetAliasOk() (*string, bool) {
	if o == nil || isNil(o.Alias) {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *ModelsCertificateStoreEntry) HasAlias() bool {
	if o != nil && !isNil(o.Alias) {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *ModelsCertificateStoreEntry) SetAlias(v string) {
	o.Alias = &v
}

// GetJobFields returns the JobFields field value if set, zero value otherwise.
func (o *ModelsCertificateStoreEntry) GetJobFields() map[string]map[string]interface{} {
	if o == nil || isNil(o.JobFields) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.JobFields
}

// GetJobFieldsOk returns a tuple with the JobFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertificateStoreEntry) GetJobFieldsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || isNil(o.JobFields) {
		return map[string]map[string]interface{}{}, false
	}
	return o.JobFields, true
}

// HasJobFields returns a boolean if a field has been set.
func (o *ModelsCertificateStoreEntry) HasJobFields() bool {
	if o != nil && !isNil(o.JobFields) {
		return true
	}

	return false
}

// SetJobFields gets a reference to the given map[string]map[string]interface{} and assigns it to the JobFields field.
func (o *ModelsCertificateStoreEntry) SetJobFields(v map[string]map[string]interface{}) {
	o.JobFields = v
}

// GetOverwrite returns the Overwrite field value if set, zero value otherwise.
func (o *ModelsCertificateStoreEntry) GetOverwrite() bool {
	if o == nil || isNil(o.Overwrite) {
		var ret bool
		return ret
	}
	return *o.Overwrite
}

// GetOverwriteOk returns a tuple with the Overwrite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertificateStoreEntry) GetOverwriteOk() (*bool, bool) {
	if o == nil || isNil(o.Overwrite) {
		return nil, false
	}
	return o.Overwrite, true
}

// HasOverwrite returns a boolean if a field has been set.
func (o *ModelsCertificateStoreEntry) HasOverwrite() bool {
	if o != nil && !isNil(o.Overwrite) {
		return true
	}

	return false
}

// SetOverwrite gets a reference to the given bool and assigns it to the Overwrite field.
func (o *ModelsCertificateStoreEntry) SetOverwrite(v bool) {
	o.Overwrite = &v
}

// GetEntryPassword returns the EntryPassword field value if set, zero value otherwise.
func (o *ModelsCertificateStoreEntry) GetEntryPassword() ModelsKeyfactorAPISecret {
	if o == nil || isNil(o.EntryPassword) {
		var ret ModelsKeyfactorAPISecret
		return ret
	}
	return *o.EntryPassword
}

// GetEntryPasswordOk returns a tuple with the EntryPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertificateStoreEntry) GetEntryPasswordOk() (*ModelsKeyfactorAPISecret, bool) {
	if o == nil || isNil(o.EntryPassword) {
		return nil, false
	}
	return o.EntryPassword, true
}

// HasEntryPassword returns a boolean if a field has been set.
func (o *ModelsCertificateStoreEntry) HasEntryPassword() bool {
	if o != nil && !isNil(o.EntryPassword) {
		return true
	}

	return false
}

// SetEntryPassword gets a reference to the given ModelsKeyfactorAPISecret and assigns it to the EntryPassword field.
func (o *ModelsCertificateStoreEntry) SetEntryPassword(v ModelsKeyfactorAPISecret) {
	o.EntryPassword = &v
}

// GetPfxPassword returns the PfxPassword field value if set, zero value otherwise.
func (o *ModelsCertificateStoreEntry) GetPfxPassword() string {
	if o == nil || isNil(o.PfxPassword) {
		var ret string
		return ret
	}
	return *o.PfxPassword
}

// GetPfxPasswordOk returns a tuple with the PfxPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertificateStoreEntry) GetPfxPasswordOk() (*string, bool) {
	if o == nil || isNil(o.PfxPassword) {
		return nil, false
	}
	return o.PfxPassword, true
}

// HasPfxPassword returns a boolean if a field has been set.
func (o *ModelsCertificateStoreEntry) HasPfxPassword() bool {
	if o != nil && !isNil(o.PfxPassword) {
		return true
	}

	return false
}

// SetPfxPassword gets a reference to the given string and assigns it to the PfxPassword field.
func (o *ModelsCertificateStoreEntry) SetPfxPassword(v string) {
	o.PfxPassword = &v
}

// GetIncludePrivateKey returns the IncludePrivateKey field value if set, zero value otherwise.
func (o *ModelsCertificateStoreEntry) GetIncludePrivateKey() bool {
	if o == nil || isNil(o.IncludePrivateKey) {
		var ret bool
		return ret
	}
	return *o.IncludePrivateKey
}

// GetIncludePrivateKeyOk returns a tuple with the IncludePrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertificateStoreEntry) GetIncludePrivateKeyOk() (*bool, bool) {
	if o == nil || isNil(o.IncludePrivateKey) {
		return nil, false
	}
	return o.IncludePrivateKey, true
}

// HasIncludePrivateKey returns a boolean if a field has been set.
func (o *ModelsCertificateStoreEntry) HasIncludePrivateKey() bool {
	if o != nil && !isNil(o.IncludePrivateKey) {
		return true
	}

	return false
}

// SetIncludePrivateKey gets a reference to the given bool and assigns it to the IncludePrivateKey field.
func (o *ModelsCertificateStoreEntry) SetIncludePrivateKey(v bool) {
	o.IncludePrivateKey = &v
}

func (o ModelsCertificateStoreEntry) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsCertificateStoreEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["CertificateStoreId"] = o.CertificateStoreId
	if !isNil(o.Alias) {
		toSerialize["Alias"] = o.Alias
	}
	if !isNil(o.JobFields) {
		toSerialize["JobFields"] = o.JobFields
	}
	if !isNil(o.Overwrite) {
		toSerialize["Overwrite"] = o.Overwrite
	}
	if !isNil(o.EntryPassword) {
		toSerialize["EntryPassword"] = o.EntryPassword
	}
	if !isNil(o.PfxPassword) {
		toSerialize["PfxPassword"] = o.PfxPassword
	}
	if !isNil(o.IncludePrivateKey) {
		toSerialize["IncludePrivateKey"] = o.IncludePrivateKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModelsCertificateStoreEntry) UnmarshalJSON(bytes []byte) (err error) {
	varModelsCertificateStoreEntry := _ModelsCertificateStoreEntry{}

	if err = json.Unmarshal(bytes, &varModelsCertificateStoreEntry); err == nil {
		*o = ModelsCertificateStoreEntry(varModelsCertificateStoreEntry)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "CertificateStoreId")
		delete(additionalProperties, "Alias")
		delete(additionalProperties, "JobFields")
		delete(additionalProperties, "Overwrite")
		delete(additionalProperties, "EntryPassword")
		delete(additionalProperties, "PfxPassword")
		delete(additionalProperties, "IncludePrivateKey")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelsCertificateStoreEntry struct {
	value *ModelsCertificateStoreEntry
	isSet bool
}

func (v NullableModelsCertificateStoreEntry) Get() *ModelsCertificateStoreEntry {
	return v.value
}

func (v *NullableModelsCertificateStoreEntry) Set(val *ModelsCertificateStoreEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsCertificateStoreEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsCertificateStoreEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsCertificateStoreEntry(val *ModelsCertificateStoreEntry) *NullableModelsCertificateStoreEntry {
	return &NullableModelsCertificateStoreEntry{value: val, isSet: true}
}

func (v NullableModelsCertificateStoreEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsCertificateStoreEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
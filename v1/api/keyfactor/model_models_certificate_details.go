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

// checks if the ModelsCertificateDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsCertificateDetails{}

// ModelsCertificateDetails struct for ModelsCertificateDetails
type ModelsCertificateDetails struct {
	IssuedDN             NullableString     `json:"IssuedDN,omitempty"`
	IssuerDN             NullableString     `json:"IssuerDN,omitempty"`
	Thumbprint           *string            `json:"Thumbprint,omitempty"`
	NotAfter             *time.Time         `json:"NotAfter,omitempty"`
	NotBefore            *time.Time         `json:"NotBefore,omitempty"`
	Metadata             *map[string]string `json:"Metadata,omitempty"`
	IsEndEntity          *bool              `json:"IsEndEntity,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelsCertificateDetails ModelsCertificateDetails

// NewModelsCertificateDetails instantiates a new ModelsCertificateDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsCertificateDetails() *ModelsCertificateDetails {
	this := ModelsCertificateDetails{}
	return &this
}

// NewModelsCertificateDetailsWithDefaults instantiates a new ModelsCertificateDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsCertificateDetailsWithDefaults() *ModelsCertificateDetails {
	this := ModelsCertificateDetails{}
	return &this
}

// GetIssuedDN returns the IssuedDN field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelsCertificateDetails) GetIssuedDN() string {
	if o == nil || isNil(o.IssuedDN.Get()) {
		var ret string
		return ret
	}
	return *o.IssuedDN.Get()
}

// GetIssuedDNOk returns a tuple with the IssuedDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelsCertificateDetails) GetIssuedDNOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IssuedDN.Get(), o.IssuedDN.IsSet()
}

// HasIssuedDN returns a boolean if a field has been set.
func (o *ModelsCertificateDetails) HasIssuedDN() bool {
	if o != nil && o.IssuedDN.IsSet() {
		return true
	}

	return false
}

// SetIssuedDN gets a reference to the given NullableString and assigns it to the IssuedDN field.
func (o *ModelsCertificateDetails) SetIssuedDN(v string) {
	o.IssuedDN.Set(&v)
}

// SetIssuedDNNil sets the value for IssuedDN to be an explicit nil
func (o *ModelsCertificateDetails) SetIssuedDNNil() {
	o.IssuedDN.Set(nil)
}

// UnsetIssuedDN ensures that no value is present for IssuedDN, not even an explicit nil
func (o *ModelsCertificateDetails) UnsetIssuedDN() {
	o.IssuedDN.Unset()
}

// GetIssuerDN returns the IssuerDN field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelsCertificateDetails) GetIssuerDN() string {
	if o == nil || isNil(o.IssuerDN.Get()) {
		var ret string
		return ret
	}
	return *o.IssuerDN.Get()
}

// GetIssuerDNOk returns a tuple with the IssuerDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelsCertificateDetails) GetIssuerDNOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IssuerDN.Get(), o.IssuerDN.IsSet()
}

// HasIssuerDN returns a boolean if a field has been set.
func (o *ModelsCertificateDetails) HasIssuerDN() bool {
	if o != nil && o.IssuerDN.IsSet() {
		return true
	}

	return false
}

// SetIssuerDN gets a reference to the given NullableString and assigns it to the IssuerDN field.
func (o *ModelsCertificateDetails) SetIssuerDN(v string) {
	o.IssuerDN.Set(&v)
}

// SetIssuerDNNil sets the value for IssuerDN to be an explicit nil
func (o *ModelsCertificateDetails) SetIssuerDNNil() {
	o.IssuerDN.Set(nil)
}

// UnsetIssuerDN ensures that no value is present for IssuerDN, not even an explicit nil
func (o *ModelsCertificateDetails) UnsetIssuerDN() {
	o.IssuerDN.Unset()
}

// GetThumbprint returns the Thumbprint field value if set, zero value otherwise.
func (o *ModelsCertificateDetails) GetThumbprint() string {
	if o == nil || isNil(o.Thumbprint) {
		var ret string
		return ret
	}
	return *o.Thumbprint
}

// GetThumbprintOk returns a tuple with the Thumbprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertificateDetails) GetThumbprintOk() (*string, bool) {
	if o == nil || isNil(o.Thumbprint) {
		return nil, false
	}
	return o.Thumbprint, true
}

// HasThumbprint returns a boolean if a field has been set.
func (o *ModelsCertificateDetails) HasThumbprint() bool {
	if o != nil && !isNil(o.Thumbprint) {
		return true
	}

	return false
}

// SetThumbprint gets a reference to the given string and assigns it to the Thumbprint field.
func (o *ModelsCertificateDetails) SetThumbprint(v string) {
	o.Thumbprint = &v
}

// GetNotAfter returns the NotAfter field value if set, zero value otherwise.
func (o *ModelsCertificateDetails) GetNotAfter() time.Time {
	if o == nil || isNil(o.NotAfter) {
		var ret time.Time
		return ret
	}
	return *o.NotAfter
}

// GetNotAfterOk returns a tuple with the NotAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertificateDetails) GetNotAfterOk() (*time.Time, bool) {
	if o == nil || isNil(o.NotAfter) {
		return nil, false
	}
	return o.NotAfter, true
}

// HasNotAfter returns a boolean if a field has been set.
func (o *ModelsCertificateDetails) HasNotAfter() bool {
	if o != nil && !isNil(o.NotAfter) {
		return true
	}

	return false
}

// SetNotAfter gets a reference to the given time.Time and assigns it to the NotAfter field.
func (o *ModelsCertificateDetails) SetNotAfter(v time.Time) {
	o.NotAfter = &v
}

// GetNotBefore returns the NotBefore field value if set, zero value otherwise.
func (o *ModelsCertificateDetails) GetNotBefore() time.Time {
	if o == nil || isNil(o.NotBefore) {
		var ret time.Time
		return ret
	}
	return *o.NotBefore
}

// GetNotBeforeOk returns a tuple with the NotBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertificateDetails) GetNotBeforeOk() (*time.Time, bool) {
	if o == nil || isNil(o.NotBefore) {
		return nil, false
	}
	return o.NotBefore, true
}

// HasNotBefore returns a boolean if a field has been set.
func (o *ModelsCertificateDetails) HasNotBefore() bool {
	if o != nil && !isNil(o.NotBefore) {
		return true
	}

	return false
}

// SetNotBefore gets a reference to the given time.Time and assigns it to the NotBefore field.
func (o *ModelsCertificateDetails) SetNotBefore(v time.Time) {
	o.NotBefore = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ModelsCertificateDetails) GetMetadata() map[string]string {
	if o == nil || isNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertificateDetails) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || isNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ModelsCertificateDetails) HasMetadata() bool {
	if o != nil && !isNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *ModelsCertificateDetails) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetIsEndEntity returns the IsEndEntity field value if set, zero value otherwise.
func (o *ModelsCertificateDetails) GetIsEndEntity() bool {
	if o == nil || isNil(o.IsEndEntity) {
		var ret bool
		return ret
	}
	return *o.IsEndEntity
}

// GetIsEndEntityOk returns a tuple with the IsEndEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertificateDetails) GetIsEndEntityOk() (*bool, bool) {
	if o == nil || isNil(o.IsEndEntity) {
		return nil, false
	}
	return o.IsEndEntity, true
}

// HasIsEndEntity returns a boolean if a field has been set.
func (o *ModelsCertificateDetails) HasIsEndEntity() bool {
	if o != nil && !isNil(o.IsEndEntity) {
		return true
	}

	return false
}

// SetIsEndEntity gets a reference to the given bool and assigns it to the IsEndEntity field.
func (o *ModelsCertificateDetails) SetIsEndEntity(v bool) {
	o.IsEndEntity = &v
}

func (o ModelsCertificateDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsCertificateDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.IssuedDN.IsSet() {
		toSerialize["IssuedDN"] = o.IssuedDN.Get()
	}
	if o.IssuerDN.IsSet() {
		toSerialize["IssuerDN"] = o.IssuerDN.Get()
	}
	if !isNil(o.Thumbprint) {
		toSerialize["Thumbprint"] = o.Thumbprint
	}
	if !isNil(o.NotAfter) {
		toSerialize["NotAfter"] = o.NotAfter
	}
	if !isNil(o.NotBefore) {
		toSerialize["NotBefore"] = o.NotBefore
	}
	if !isNil(o.Metadata) {
		toSerialize["Metadata"] = o.Metadata
	}
	if !isNil(o.IsEndEntity) {
		toSerialize["IsEndEntity"] = o.IsEndEntity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModelsCertificateDetails) UnmarshalJSON(bytes []byte) (err error) {
	varModelsCertificateDetails := _ModelsCertificateDetails{}

	if err = json.Unmarshal(bytes, &varModelsCertificateDetails); err == nil {
		*o = ModelsCertificateDetails(varModelsCertificateDetails)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "IssuedDN")
		delete(additionalProperties, "IssuerDN")
		delete(additionalProperties, "Thumbprint")
		delete(additionalProperties, "NotAfter")
		delete(additionalProperties, "NotBefore")
		delete(additionalProperties, "Metadata")
		delete(additionalProperties, "IsEndEntity")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelsCertificateDetails struct {
	value *ModelsCertificateDetails
	isSet bool
}

func (v NullableModelsCertificateDetails) Get() *ModelsCertificateDetails {
	return v.value
}

func (v *NullableModelsCertificateDetails) Set(val *ModelsCertificateDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsCertificateDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsCertificateDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsCertificateDetails(val *ModelsCertificateDetails) *NullableModelsCertificateDetails {
	return &NullableModelsCertificateDetails{value: val, isSet: true}
}

func (v NullableModelsCertificateDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsCertificateDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
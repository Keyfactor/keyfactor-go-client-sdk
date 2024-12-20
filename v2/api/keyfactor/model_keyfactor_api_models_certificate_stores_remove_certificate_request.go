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

// checks if the KeyfactorApiModelsCertificateStoresRemoveCertificateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorApiModelsCertificateStoresRemoveCertificateRequest{}

// KeyfactorApiModelsCertificateStoresRemoveCertificateRequest struct for KeyfactorApiModelsCertificateStoresRemoveCertificateRequest
type KeyfactorApiModelsCertificateStoresRemoveCertificateRequest struct {
	CertificateStores    []ModelsCertificateLocationSpecifier       `json:"CertificateStores"`
	Schedule             KeyfactorCommonSchedulingKeyfactorSchedule `json:"Schedule"`
	CollectionId         *int32                                     `json:"CollectionId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KeyfactorApiModelsCertificateStoresRemoveCertificateRequest KeyfactorApiModelsCertificateStoresRemoveCertificateRequest

// NewKeyfactorApiModelsCertificateStoresRemoveCertificateRequest instantiates a new KeyfactorApiModelsCertificateStoresRemoveCertificateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorApiModelsCertificateStoresRemoveCertificateRequest(certificateStores []ModelsCertificateLocationSpecifier, schedule KeyfactorCommonSchedulingKeyfactorSchedule) *KeyfactorApiModelsCertificateStoresRemoveCertificateRequest {
	this := KeyfactorApiModelsCertificateStoresRemoveCertificateRequest{}
	this.CertificateStores = certificateStores
	this.Schedule = schedule
	return &this
}

// NewKeyfactorApiModelsCertificateStoresRemoveCertificateRequestWithDefaults instantiates a new KeyfactorApiModelsCertificateStoresRemoveCertificateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorApiModelsCertificateStoresRemoveCertificateRequestWithDefaults() *KeyfactorApiModelsCertificateStoresRemoveCertificateRequest {
	this := KeyfactorApiModelsCertificateStoresRemoveCertificateRequest{}
	return &this
}

// GetCertificateStores returns the CertificateStores field value
func (o *KeyfactorApiModelsCertificateStoresRemoveCertificateRequest) GetCertificateStores() []ModelsCertificateLocationSpecifier {
	if o == nil {
		var ret []ModelsCertificateLocationSpecifier
		return ret
	}

	return o.CertificateStores
}

// GetCertificateStoresOk returns a tuple with the CertificateStores field value
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsCertificateStoresRemoveCertificateRequest) GetCertificateStoresOk() ([]ModelsCertificateLocationSpecifier, bool) {
	if o == nil {
		return nil, false
	}
	return o.CertificateStores, true
}

// SetCertificateStores sets field value
func (o *KeyfactorApiModelsCertificateStoresRemoveCertificateRequest) SetCertificateStores(v []ModelsCertificateLocationSpecifier) {
	o.CertificateStores = v
}

// GetSchedule returns the Schedule field value
func (o *KeyfactorApiModelsCertificateStoresRemoveCertificateRequest) GetSchedule() KeyfactorCommonSchedulingKeyfactorSchedule {
	if o == nil {
		var ret KeyfactorCommonSchedulingKeyfactorSchedule
		return ret
	}

	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsCertificateStoresRemoveCertificateRequest) GetScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schedule, true
}

// SetSchedule sets field value
func (o *KeyfactorApiModelsCertificateStoresRemoveCertificateRequest) SetSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule) {
	o.Schedule = v
}

// GetCollectionId returns the CollectionId field value if set, zero value otherwise.
func (o *KeyfactorApiModelsCertificateStoresRemoveCertificateRequest) GetCollectionId() int32 {
	if o == nil || isNil(o.CollectionId) {
		var ret int32
		return ret
	}
	return *o.CollectionId
}

// GetCollectionIdOk returns a tuple with the CollectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsCertificateStoresRemoveCertificateRequest) GetCollectionIdOk() (*int32, bool) {
	if o == nil || isNil(o.CollectionId) {
		return nil, false
	}
	return o.CollectionId, true
}

// HasCollectionId returns a boolean if a field has been set.
func (o *KeyfactorApiModelsCertificateStoresRemoveCertificateRequest) HasCollectionId() bool {
	if o != nil && !isNil(o.CollectionId) {
		return true
	}

	return false
}

// SetCollectionId gets a reference to the given int32 and assigns it to the CollectionId field.
func (o *KeyfactorApiModelsCertificateStoresRemoveCertificateRequest) SetCollectionId(v int32) {
	o.CollectionId = &v
}

func (o KeyfactorApiModelsCertificateStoresRemoveCertificateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorApiModelsCertificateStoresRemoveCertificateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["CertificateStores"] = o.CertificateStores
	toSerialize["Schedule"] = o.Schedule
	if !isNil(o.CollectionId) {
		toSerialize["CollectionId"] = o.CollectionId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KeyfactorApiModelsCertificateStoresRemoveCertificateRequest) UnmarshalJSON(bytes []byte) (err error) {
	varKeyfactorApiModelsCertificateStoresRemoveCertificateRequest := _KeyfactorApiModelsCertificateStoresRemoveCertificateRequest{}

	if err = json.Unmarshal(bytes, &varKeyfactorApiModelsCertificateStoresRemoveCertificateRequest); err == nil {
		*o = KeyfactorApiModelsCertificateStoresRemoveCertificateRequest(varKeyfactorApiModelsCertificateStoresRemoveCertificateRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "CertificateStores")
		delete(additionalProperties, "Schedule")
		delete(additionalProperties, "CollectionId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKeyfactorApiModelsCertificateStoresRemoveCertificateRequest struct {
	value *KeyfactorApiModelsCertificateStoresRemoveCertificateRequest
	isSet bool
}

func (v NullableKeyfactorApiModelsCertificateStoresRemoveCertificateRequest) Get() *KeyfactorApiModelsCertificateStoresRemoveCertificateRequest {
	return v.value
}

func (v *NullableKeyfactorApiModelsCertificateStoresRemoveCertificateRequest) Set(val *KeyfactorApiModelsCertificateStoresRemoveCertificateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorApiModelsCertificateStoresRemoveCertificateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorApiModelsCertificateStoresRemoveCertificateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorApiModelsCertificateStoresRemoveCertificateRequest(val *KeyfactorApiModelsCertificateStoresRemoveCertificateRequest) *NullableKeyfactorApiModelsCertificateStoresRemoveCertificateRequest {
	return &NullableKeyfactorApiModelsCertificateStoresRemoveCertificateRequest{value: val, isSet: true}
}

func (v NullableKeyfactorApiModelsCertificateStoresRemoveCertificateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorApiModelsCertificateStoresRemoveCertificateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

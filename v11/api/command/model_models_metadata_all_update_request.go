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

// checks if the ModelsMetadataAllUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsMetadataAllUpdateRequest{}

// ModelsMetadataAllUpdateRequest struct for ModelsMetadataAllUpdateRequest
type ModelsMetadataAllUpdateRequest struct {
	Query          NullableString                      `json:"query,omitempty"`
	CertificateIds []int32                             `json:"certificateIds,omitempty"`
	Metadata       []ModelsMetadataSingleUpdateRequest `json:"metadata"`
}

// NewModelsMetadataAllUpdateRequest instantiates a new ModelsMetadataAllUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsMetadataAllUpdateRequest(metadata []ModelsMetadataSingleUpdateRequest) *ModelsMetadataAllUpdateRequest {
	this := ModelsMetadataAllUpdateRequest{}
	this.Metadata = metadata
	return &this
}

// NewModelsMetadataAllUpdateRequestWithDefaults instantiates a new ModelsMetadataAllUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsMetadataAllUpdateRequestWithDefaults() *ModelsMetadataAllUpdateRequest {
	this := ModelsMetadataAllUpdateRequest{}
	return &this
}

// GetQuery returns the Query field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelsMetadataAllUpdateRequest) GetQuery() string {
	if o == nil || isNil(o.Query.Get()) {
		var ret string
		return ret
	}
	return *o.Query.Get()
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelsMetadataAllUpdateRequest) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Query.Get(), o.Query.IsSet()
}

// HasQuery returns a boolean if a field has been set.
func (o *ModelsMetadataAllUpdateRequest) HasQuery() bool {
	if o != nil && o.Query.IsSet() {
		return true
	}

	return false
}

// SetQuery gets a reference to the given NullableString and assigns it to the Query field.
func (o *ModelsMetadataAllUpdateRequest) SetQuery(v string) {
	o.Query.Set(&v)
}

// SetQueryNil sets the value for Query to be an explicit nil
func (o *ModelsMetadataAllUpdateRequest) SetQueryNil() {
	o.Query.Set(nil)
}

// UnsetQuery ensures that no value is present for Query, not even an explicit nil
func (o *ModelsMetadataAllUpdateRequest) UnsetQuery() {
	o.Query.Unset()
}

// GetCertificateIds returns the CertificateIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelsMetadataAllUpdateRequest) GetCertificateIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.CertificateIds
}

// GetCertificateIdsOk returns a tuple with the CertificateIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelsMetadataAllUpdateRequest) GetCertificateIdsOk() ([]int32, bool) {
	if o == nil || isNil(o.CertificateIds) {
		return nil, false
	}
	return o.CertificateIds, true
}

// HasCertificateIds returns a boolean if a field has been set.
func (o *ModelsMetadataAllUpdateRequest) HasCertificateIds() bool {
	if o != nil && isNil(o.CertificateIds) {
		return true
	}

	return false
}

// SetCertificateIds gets a reference to the given []int32 and assigns it to the CertificateIds field.
func (o *ModelsMetadataAllUpdateRequest) SetCertificateIds(v []int32) {
	o.CertificateIds = v
}

// GetMetadata returns the Metadata field value
func (o *ModelsMetadataAllUpdateRequest) GetMetadata() []ModelsMetadataSingleUpdateRequest {
	if o == nil {
		var ret []ModelsMetadataSingleUpdateRequest
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *ModelsMetadataAllUpdateRequest) GetMetadataOk() ([]ModelsMetadataSingleUpdateRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metadata, true
}

// SetMetadata sets field value
func (o *ModelsMetadataAllUpdateRequest) SetMetadata(v []ModelsMetadataSingleUpdateRequest) {
	o.Metadata = v
}

func (o ModelsMetadataAllUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsMetadataAllUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Query.IsSet() {
		toSerialize["query"] = o.Query.Get()
	}
	if o.CertificateIds != nil {
		toSerialize["certificateIds"] = o.CertificateIds
	}
	toSerialize["metadata"] = o.Metadata
	return toSerialize, nil
}

type NullableModelsMetadataAllUpdateRequest struct {
	value *ModelsMetadataAllUpdateRequest
	isSet bool
}

func (v NullableModelsMetadataAllUpdateRequest) Get() *ModelsMetadataAllUpdateRequest {
	return v.value
}

func (v *NullableModelsMetadataAllUpdateRequest) Set(val *ModelsMetadataAllUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsMetadataAllUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsMetadataAllUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsMetadataAllUpdateRequest(val *ModelsMetadataAllUpdateRequest) *NullableModelsMetadataAllUpdateRequest {
	return &NullableModelsMetadataAllUpdateRequest{value: val, isSet: true}
}

func (v NullableModelsMetadataAllUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsMetadataAllUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
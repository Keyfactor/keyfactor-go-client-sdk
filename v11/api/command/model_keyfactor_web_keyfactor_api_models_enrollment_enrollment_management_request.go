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
	"time"
)

// checks if the KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest{}

// KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest struct for KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest
type KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest struct {
	// The stores to add the certificate to. Values in this collection will take precedence over values in CSS.CMS.Data.Model.Models.Enrollment.SpecificEnrollmentManagementRequest.StoreTypes.
	Stores        []KeyfactorWebKeyfactorApiModelsEnrollmentManagementStoreRequest     `json:"stores,omitempty"`
	StoreIds      []string                                                             `json:"storeIds,omitempty"`
	StoreTypes    []KeyfactorWebKeyfactorApiModelsEnrollmentManagementStoreTypeRequest `json:"storeTypes,omitempty"`
	CertificateId *int32                                                               `json:"certificateId,omitempty"`
	RequestId     *int32                                                               `json:"requestId,omitempty"`
	Password      string                                                               `json:"password"`
	JobTime       *time.Time                                                           `json:"jobTime,omitempty"`
}

// NewKeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest instantiates a new KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest(password string) *KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest {
	this := KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest{}
	this.Password = password
	return &this
}

// NewKeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequestWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequestWithDefaults() *KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest {
	this := KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest{}
	return &this
}

// GetStores returns the Stores field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest) GetStores() []KeyfactorWebKeyfactorApiModelsEnrollmentManagementStoreRequest {
	if o == nil {
		var ret []KeyfactorWebKeyfactorApiModelsEnrollmentManagementStoreRequest
		return ret
	}
	return o.Stores
}

// GetStoresOk returns a tuple with the Stores field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest) GetStoresOk() ([]KeyfactorWebKeyfactorApiModelsEnrollmentManagementStoreRequest, bool) {
	if o == nil || isNil(o.Stores) {
		return nil, false
	}
	return o.Stores, true
}

// HasStores returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest) HasStores() bool {
	if o != nil && isNil(o.Stores) {
		return true
	}

	return false
}

// SetStores gets a reference to the given []KeyfactorWebKeyfactorApiModelsEnrollmentManagementStoreRequest and assigns it to the Stores field.
func (o *KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest) SetStores(v []KeyfactorWebKeyfactorApiModelsEnrollmentManagementStoreRequest) {
	o.Stores = v
}

// GetStoreIds returns the StoreIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest) GetStoreIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.StoreIds
}

// GetStoreIdsOk returns a tuple with the StoreIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest) GetStoreIdsOk() ([]string, bool) {
	if o == nil || isNil(o.StoreIds) {
		return nil, false
	}
	return o.StoreIds, true
}

// HasStoreIds returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest) HasStoreIds() bool {
	if o != nil && isNil(o.StoreIds) {
		return true
	}

	return false
}

// SetStoreIds gets a reference to the given []string and assigns it to the StoreIds field.
func (o *KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest) SetStoreIds(v []string) {
	o.StoreIds = v
}

// GetStoreTypes returns the StoreTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest) GetStoreTypes() []KeyfactorWebKeyfactorApiModelsEnrollmentManagementStoreTypeRequest {
	if o == nil {
		var ret []KeyfactorWebKeyfactorApiModelsEnrollmentManagementStoreTypeRequest
		return ret
	}
	return o.StoreTypes
}

// GetStoreTypesOk returns a tuple with the StoreTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest) GetStoreTypesOk() ([]KeyfactorWebKeyfactorApiModelsEnrollmentManagementStoreTypeRequest, bool) {
	if o == nil || isNil(o.StoreTypes) {
		return nil, false
	}
	return o.StoreTypes, true
}

// HasStoreTypes returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest) HasStoreTypes() bool {
	if o != nil && isNil(o.StoreTypes) {
		return true
	}

	return false
}

// SetStoreTypes gets a reference to the given []KeyfactorWebKeyfactorApiModelsEnrollmentManagementStoreTypeRequest and assigns it to the StoreTypes field.
func (o *KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest) SetStoreTypes(v []KeyfactorWebKeyfactorApiModelsEnrollmentManagementStoreTypeRequest) {
	o.StoreTypes = v
}

// GetCertificateId returns the CertificateId field value if set, zero value otherwise.
func (o *KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest) GetCertificateId() int32 {
	if o == nil || isNil(o.CertificateId) {
		var ret int32
		return ret
	}
	return *o.CertificateId
}

// GetCertificateIdOk returns a tuple with the CertificateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest) GetCertificateIdOk() (*int32, bool) {
	if o == nil || isNil(o.CertificateId) {
		return nil, false
	}
	return o.CertificateId, true
}

// HasCertificateId returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest) HasCertificateId() bool {
	if o != nil && !isNil(o.CertificateId) {
		return true
	}

	return false
}

// SetCertificateId gets a reference to the given int32 and assigns it to the CertificateId field.
func (o *KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest) SetCertificateId(v int32) {
	o.CertificateId = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest) GetRequestId() int32 {
	if o == nil || isNil(o.RequestId) {
		var ret int32
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest) GetRequestIdOk() (*int32, bool) {
	if o == nil || isNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest) HasRequestId() bool {
	if o != nil && !isNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given int32 and assigns it to the RequestId field.
func (o *KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest) SetRequestId(v int32) {
	o.RequestId = &v
}

// GetPassword returns the Password field value
func (o *KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest) SetPassword(v string) {
	o.Password = v
}

// GetJobTime returns the JobTime field value if set, zero value otherwise.
func (o *KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest) GetJobTime() time.Time {
	if o == nil || isNil(o.JobTime) {
		var ret time.Time
		return ret
	}
	return *o.JobTime
}

// GetJobTimeOk returns a tuple with the JobTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest) GetJobTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.JobTime) {
		return nil, false
	}
	return o.JobTime, true
}

// HasJobTime returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest) HasJobTime() bool {
	if o != nil && !isNil(o.JobTime) {
		return true
	}

	return false
}

// SetJobTime gets a reference to the given time.Time and assigns it to the JobTime field.
func (o *KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest) SetJobTime(v time.Time) {
	o.JobTime = &v
}

func (o KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Stores != nil {
		toSerialize["stores"] = o.Stores
	}
	if o.StoreIds != nil {
		toSerialize["storeIds"] = o.StoreIds
	}
	if o.StoreTypes != nil {
		toSerialize["storeTypes"] = o.StoreTypes
	}
	if !isNil(o.CertificateId) {
		toSerialize["certificateId"] = o.CertificateId
	}
	if !isNil(o.RequestId) {
		toSerialize["requestId"] = o.RequestId
	}
	toSerialize["password"] = o.Password
	if !isNil(o.JobTime) {
		toSerialize["jobTime"] = o.JobTime
	}
	return toSerialize, nil
}

type NullableKeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest struct {
	value *KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest
	isSet bool
}

func (v NullableKeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest) Get() *KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest {
	return v.value
}

func (v *NullableKeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest) Set(val *KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest(val *KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest) *NullableKeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest {
	return &NullableKeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest{value: val, isSet: true}
}

func (v NullableKeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
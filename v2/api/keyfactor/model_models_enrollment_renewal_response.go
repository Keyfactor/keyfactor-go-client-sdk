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

// checks if the ModelsEnrollmentRenewalResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsEnrollmentRenewalResponse{}

// ModelsEnrollmentRenewalResponse struct for ModelsEnrollmentRenewalResponse
type ModelsEnrollmentRenewalResponse struct {
	KeyfactorId          *int32         `json:"KeyfactorId,omitempty"`
	KeyfactorRequestId   *int32         `json:"KeyfactorRequestId,omitempty"`
	Thumbprint           *string        `json:"Thumbprint,omitempty"`
	SerialNumber         *string        `json:"SerialNumber,omitempty"`
	IssuerDN             NullableString `json:"IssuerDN,omitempty"`
	RequestDisposition   *string        `json:"RequestDisposition,omitempty"`
	DispositionMessage   *string        `json:"DispositionMessage,omitempty"`
	Password             *string        `json:"Password,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelsEnrollmentRenewalResponse ModelsEnrollmentRenewalResponse

// NewModelsEnrollmentRenewalResponse instantiates a new ModelsEnrollmentRenewalResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsEnrollmentRenewalResponse() *ModelsEnrollmentRenewalResponse {
	this := ModelsEnrollmentRenewalResponse{}
	return &this
}

// NewModelsEnrollmentRenewalResponseWithDefaults instantiates a new ModelsEnrollmentRenewalResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsEnrollmentRenewalResponseWithDefaults() *ModelsEnrollmentRenewalResponse {
	this := ModelsEnrollmentRenewalResponse{}
	return &this
}

// GetKeyfactorId returns the KeyfactorId field value if set, zero value otherwise.
func (o *ModelsEnrollmentRenewalResponse) GetKeyfactorId() int32 {
	if o == nil || isNil(o.KeyfactorId) {
		var ret int32
		return ret
	}
	return *o.KeyfactorId
}

// GetKeyfactorIdOk returns a tuple with the KeyfactorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsEnrollmentRenewalResponse) GetKeyfactorIdOk() (*int32, bool) {
	if o == nil || isNil(o.KeyfactorId) {
		return nil, false
	}
	return o.KeyfactorId, true
}

// HasKeyfactorId returns a boolean if a field has been set.
func (o *ModelsEnrollmentRenewalResponse) HasKeyfactorId() bool {
	if o != nil && !isNil(o.KeyfactorId) {
		return true
	}

	return false
}

// SetKeyfactorId gets a reference to the given int32 and assigns it to the KeyfactorId field.
func (o *ModelsEnrollmentRenewalResponse) SetKeyfactorId(v int32) {
	o.KeyfactorId = &v
}

// GetKeyfactorRequestId returns the KeyfactorRequestId field value if set, zero value otherwise.
func (o *ModelsEnrollmentRenewalResponse) GetKeyfactorRequestId() int32 {
	if o == nil || isNil(o.KeyfactorRequestId) {
		var ret int32
		return ret
	}
	return *o.KeyfactorRequestId
}

// GetKeyfactorRequestIdOk returns a tuple with the KeyfactorRequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsEnrollmentRenewalResponse) GetKeyfactorRequestIdOk() (*int32, bool) {
	if o == nil || isNil(o.KeyfactorRequestId) {
		return nil, false
	}
	return o.KeyfactorRequestId, true
}

// HasKeyfactorRequestId returns a boolean if a field has been set.
func (o *ModelsEnrollmentRenewalResponse) HasKeyfactorRequestId() bool {
	if o != nil && !isNil(o.KeyfactorRequestId) {
		return true
	}

	return false
}

// SetKeyfactorRequestId gets a reference to the given int32 and assigns it to the KeyfactorRequestId field.
func (o *ModelsEnrollmentRenewalResponse) SetKeyfactorRequestId(v int32) {
	o.KeyfactorRequestId = &v
}

// GetThumbprint returns the Thumbprint field value if set, zero value otherwise.
func (o *ModelsEnrollmentRenewalResponse) GetThumbprint() string {
	if o == nil || isNil(o.Thumbprint) {
		var ret string
		return ret
	}
	return *o.Thumbprint
}

// GetThumbprintOk returns a tuple with the Thumbprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsEnrollmentRenewalResponse) GetThumbprintOk() (*string, bool) {
	if o == nil || isNil(o.Thumbprint) {
		return nil, false
	}
	return o.Thumbprint, true
}

// HasThumbprint returns a boolean if a field has been set.
func (o *ModelsEnrollmentRenewalResponse) HasThumbprint() bool {
	if o != nil && !isNil(o.Thumbprint) {
		return true
	}

	return false
}

// SetThumbprint gets a reference to the given string and assigns it to the Thumbprint field.
func (o *ModelsEnrollmentRenewalResponse) SetThumbprint(v string) {
	o.Thumbprint = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *ModelsEnrollmentRenewalResponse) GetSerialNumber() string {
	if o == nil || isNil(o.SerialNumber) {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsEnrollmentRenewalResponse) GetSerialNumberOk() (*string, bool) {
	if o == nil || isNil(o.SerialNumber) {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *ModelsEnrollmentRenewalResponse) HasSerialNumber() bool {
	if o != nil && !isNil(o.SerialNumber) {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *ModelsEnrollmentRenewalResponse) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetIssuerDN returns the IssuerDN field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelsEnrollmentRenewalResponse) GetIssuerDN() string {
	if o == nil || isNil(o.IssuerDN.Get()) {
		var ret string
		return ret
	}
	return *o.IssuerDN.Get()
}

// GetIssuerDNOk returns a tuple with the IssuerDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelsEnrollmentRenewalResponse) GetIssuerDNOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IssuerDN.Get(), o.IssuerDN.IsSet()
}

// HasIssuerDN returns a boolean if a field has been set.
func (o *ModelsEnrollmentRenewalResponse) HasIssuerDN() bool {
	if o != nil && o.IssuerDN.IsSet() {
		return true
	}

	return false
}

// SetIssuerDN gets a reference to the given NullableString and assigns it to the IssuerDN field.
func (o *ModelsEnrollmentRenewalResponse) SetIssuerDN(v string) {
	o.IssuerDN.Set(&v)
}

// SetIssuerDNNil sets the value for IssuerDN to be an explicit nil
func (o *ModelsEnrollmentRenewalResponse) SetIssuerDNNil() {
	o.IssuerDN.Set(nil)
}

// UnsetIssuerDN ensures that no value is present for IssuerDN, not even an explicit nil
func (o *ModelsEnrollmentRenewalResponse) UnsetIssuerDN() {
	o.IssuerDN.Unset()
}

// GetRequestDisposition returns the RequestDisposition field value if set, zero value otherwise.
func (o *ModelsEnrollmentRenewalResponse) GetRequestDisposition() string {
	if o == nil || isNil(o.RequestDisposition) {
		var ret string
		return ret
	}
	return *o.RequestDisposition
}

// GetRequestDispositionOk returns a tuple with the RequestDisposition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsEnrollmentRenewalResponse) GetRequestDispositionOk() (*string, bool) {
	if o == nil || isNil(o.RequestDisposition) {
		return nil, false
	}
	return o.RequestDisposition, true
}

// HasRequestDisposition returns a boolean if a field has been set.
func (o *ModelsEnrollmentRenewalResponse) HasRequestDisposition() bool {
	if o != nil && !isNil(o.RequestDisposition) {
		return true
	}

	return false
}

// SetRequestDisposition gets a reference to the given string and assigns it to the RequestDisposition field.
func (o *ModelsEnrollmentRenewalResponse) SetRequestDisposition(v string) {
	o.RequestDisposition = &v
}

// GetDispositionMessage returns the DispositionMessage field value if set, zero value otherwise.
func (o *ModelsEnrollmentRenewalResponse) GetDispositionMessage() string {
	if o == nil || isNil(o.DispositionMessage) {
		var ret string
		return ret
	}
	return *o.DispositionMessage
}

// GetDispositionMessageOk returns a tuple with the DispositionMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsEnrollmentRenewalResponse) GetDispositionMessageOk() (*string, bool) {
	if o == nil || isNil(o.DispositionMessage) {
		return nil, false
	}
	return o.DispositionMessage, true
}

// HasDispositionMessage returns a boolean if a field has been set.
func (o *ModelsEnrollmentRenewalResponse) HasDispositionMessage() bool {
	if o != nil && !isNil(o.DispositionMessage) {
		return true
	}

	return false
}

// SetDispositionMessage gets a reference to the given string and assigns it to the DispositionMessage field.
func (o *ModelsEnrollmentRenewalResponse) SetDispositionMessage(v string) {
	o.DispositionMessage = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ModelsEnrollmentRenewalResponse) GetPassword() string {
	if o == nil || isNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsEnrollmentRenewalResponse) GetPasswordOk() (*string, bool) {
	if o == nil || isNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ModelsEnrollmentRenewalResponse) HasPassword() bool {
	if o != nil && !isNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ModelsEnrollmentRenewalResponse) SetPassword(v string) {
	o.Password = &v
}

func (o ModelsEnrollmentRenewalResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsEnrollmentRenewalResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.KeyfactorId) {
		toSerialize["KeyfactorId"] = o.KeyfactorId
	}
	if !isNil(o.KeyfactorRequestId) {
		toSerialize["KeyfactorRequestId"] = o.KeyfactorRequestId
	}
	if !isNil(o.Thumbprint) {
		toSerialize["Thumbprint"] = o.Thumbprint
	}
	if !isNil(o.SerialNumber) {
		toSerialize["SerialNumber"] = o.SerialNumber
	}
	if o.IssuerDN.IsSet() {
		toSerialize["IssuerDN"] = o.IssuerDN.Get()
	}
	if !isNil(o.RequestDisposition) {
		toSerialize["RequestDisposition"] = o.RequestDisposition
	}
	if !isNil(o.DispositionMessage) {
		toSerialize["DispositionMessage"] = o.DispositionMessage
	}
	if !isNil(o.Password) {
		toSerialize["Password"] = o.Password
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModelsEnrollmentRenewalResponse) UnmarshalJSON(bytes []byte) (err error) {
	varModelsEnrollmentRenewalResponse := _ModelsEnrollmentRenewalResponse{}

	if err = json.Unmarshal(bytes, &varModelsEnrollmentRenewalResponse); err == nil {
		*o = ModelsEnrollmentRenewalResponse(varModelsEnrollmentRenewalResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "KeyfactorId")
		delete(additionalProperties, "KeyfactorRequestId")
		delete(additionalProperties, "Thumbprint")
		delete(additionalProperties, "SerialNumber")
		delete(additionalProperties, "IssuerDN")
		delete(additionalProperties, "RequestDisposition")
		delete(additionalProperties, "DispositionMessage")
		delete(additionalProperties, "Password")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelsEnrollmentRenewalResponse struct {
	value *ModelsEnrollmentRenewalResponse
	isSet bool
}

func (v NullableModelsEnrollmentRenewalResponse) Get() *ModelsEnrollmentRenewalResponse {
	return v.value
}

func (v *NullableModelsEnrollmentRenewalResponse) Set(val *ModelsEnrollmentRenewalResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsEnrollmentRenewalResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsEnrollmentRenewalResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsEnrollmentRenewalResponse(val *ModelsEnrollmentRenewalResponse) *NullableModelsEnrollmentRenewalResponse {
	return &NullableModelsEnrollmentRenewalResponse{value: val, isSet: true}
}

func (v NullableModelsEnrollmentRenewalResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsEnrollmentRenewalResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

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

// checks if the KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest{}

// KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest struct for KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest
type KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest struct {
	// The allowed RSA key sizes.
	RSAValidKeySizes []int32 `json:"RSAValidKeySizes"`
	// The allowed ECC curves.
	ECCValidCurves []string `json:"ECCValidCurves"`
	// Whether or not keys can be reused.
	AllowKeyReuse bool `json:"AllowKeyReuse"`
	// Whether or not wildcards can be used.
	AllowWildcards bool `json:"AllowWildcards"`
	// Whether or not RFC 2818 compliance should be enforced.
	RFCEnforcement       bool `json:"RFCEnforcement"`
	AdditionalProperties map[string]interface{}
}

type _KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest

// NewKeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest instantiates a new KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest(rSAValidKeySizes []int32, eCCValidCurves []string, allowKeyReuse bool, allowWildcards bool, rFCEnforcement bool) *KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest {
	this := KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest{}
	this.RSAValidKeySizes = rSAValidKeySizes
	this.ECCValidCurves = eCCValidCurves
	this.AllowKeyReuse = allowKeyReuse
	this.AllowWildcards = allowWildcards
	this.RFCEnforcement = rFCEnforcement
	return &this
}

// NewKeyfactorApiModelsTemplatesGlobalTemplatePolicyRequestWithDefaults instantiates a new KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorApiModelsTemplatesGlobalTemplatePolicyRequestWithDefaults() *KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest {
	this := KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest{}
	return &this
}

// GetRSAValidKeySizes returns the RSAValidKeySizes field value
func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest) GetRSAValidKeySizes() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.RSAValidKeySizes
}

// GetRSAValidKeySizesOk returns a tuple with the RSAValidKeySizes field value
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest) GetRSAValidKeySizesOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RSAValidKeySizes, true
}

// SetRSAValidKeySizes sets field value
func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest) SetRSAValidKeySizes(v []int32) {
	o.RSAValidKeySizes = v
}

// GetECCValidCurves returns the ECCValidCurves field value
func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest) GetECCValidCurves() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ECCValidCurves
}

// GetECCValidCurvesOk returns a tuple with the ECCValidCurves field value
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest) GetECCValidCurvesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ECCValidCurves, true
}

// SetECCValidCurves sets field value
func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest) SetECCValidCurves(v []string) {
	o.ECCValidCurves = v
}

// GetAllowKeyReuse returns the AllowKeyReuse field value
func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest) GetAllowKeyReuse() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowKeyReuse
}

// GetAllowKeyReuseOk returns a tuple with the AllowKeyReuse field value
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest) GetAllowKeyReuseOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowKeyReuse, true
}

// SetAllowKeyReuse sets field value
func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest) SetAllowKeyReuse(v bool) {
	o.AllowKeyReuse = v
}

// GetAllowWildcards returns the AllowWildcards field value
func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest) GetAllowWildcards() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowWildcards
}

// GetAllowWildcardsOk returns a tuple with the AllowWildcards field value
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest) GetAllowWildcardsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowWildcards, true
}

// SetAllowWildcards sets field value
func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest) SetAllowWildcards(v bool) {
	o.AllowWildcards = v
}

// GetRFCEnforcement returns the RFCEnforcement field value
func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest) GetRFCEnforcement() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RFCEnforcement
}

// GetRFCEnforcementOk returns a tuple with the RFCEnforcement field value
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest) GetRFCEnforcementOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RFCEnforcement, true
}

// SetRFCEnforcement sets field value
func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest) SetRFCEnforcement(v bool) {
	o.RFCEnforcement = v
}

func (o KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["RSAValidKeySizes"] = o.RSAValidKeySizes
	toSerialize["ECCValidCurves"] = o.ECCValidCurves
	toSerialize["AllowKeyReuse"] = o.AllowKeyReuse
	toSerialize["AllowWildcards"] = o.AllowWildcards
	toSerialize["RFCEnforcement"] = o.RFCEnforcement

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest) UnmarshalJSON(bytes []byte) (err error) {
	varKeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest := _KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest{}

	if err = json.Unmarshal(bytes, &varKeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest); err == nil {
		*o = KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest(varKeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "RSAValidKeySizes")
		delete(additionalProperties, "ECCValidCurves")
		delete(additionalProperties, "AllowKeyReuse")
		delete(additionalProperties, "AllowWildcards")
		delete(additionalProperties, "RFCEnforcement")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest struct {
	value *KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest
	isSet bool
}

func (v NullableKeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest) Get() *KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest {
	return v.value
}

func (v *NullableKeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest) Set(val *KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest(val *KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest) *NullableKeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest {
	return &NullableKeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest{value: val, isSet: true}
}

func (v NullableKeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

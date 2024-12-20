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

// checks if the ModelsRevokeAllCertificatesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsRevokeAllCertificatesRequest{}

// ModelsRevokeAllCertificatesRequest Information for revoking all certifictes in a query
type ModelsRevokeAllCertificatesRequest struct {
	// The query string of the certificates to revoke
	Query *string `json:"Query,omitempty"`
	// The Reason for revocation
	Reason int32 `json:"Reason"`
	// A comment for auditing purposes
	Comment string `json:"Comment"`
	// The date when the certificates are to appear on the revocation list
	EffectiveDate *time.Time `json:"EffectiveDate,omitempty"`
	// A flag telling the query to include revoked certificates
	IncludeRevoked *bool `json:"IncludeRevoked,omitempty"`
	// A flag telling the query to include expired certificates
	IncludeExpired       *bool `json:"IncludeExpired,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelsRevokeAllCertificatesRequest ModelsRevokeAllCertificatesRequest

// NewModelsRevokeAllCertificatesRequest instantiates a new ModelsRevokeAllCertificatesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsRevokeAllCertificatesRequest(reason int32, comment string) *ModelsRevokeAllCertificatesRequest {
	this := ModelsRevokeAllCertificatesRequest{}
	this.Reason = reason
	this.Comment = comment
	return &this
}

// NewModelsRevokeAllCertificatesRequestWithDefaults instantiates a new ModelsRevokeAllCertificatesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsRevokeAllCertificatesRequestWithDefaults() *ModelsRevokeAllCertificatesRequest {
	this := ModelsRevokeAllCertificatesRequest{}
	return &this
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *ModelsRevokeAllCertificatesRequest) GetQuery() string {
	if o == nil || isNil(o.Query) {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsRevokeAllCertificatesRequest) GetQueryOk() (*string, bool) {
	if o == nil || isNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *ModelsRevokeAllCertificatesRequest) HasQuery() bool {
	if o != nil && !isNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *ModelsRevokeAllCertificatesRequest) SetQuery(v string) {
	o.Query = &v
}

// GetReason returns the Reason field value
func (o *ModelsRevokeAllCertificatesRequest) GetReason() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *ModelsRevokeAllCertificatesRequest) GetReasonOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *ModelsRevokeAllCertificatesRequest) SetReason(v int32) {
	o.Reason = v
}

// GetComment returns the Comment field value
func (o *ModelsRevokeAllCertificatesRequest) GetComment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Comment
}

// GetCommentOk returns a tuple with the Comment field value
// and a boolean to check if the value has been set.
func (o *ModelsRevokeAllCertificatesRequest) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comment, true
}

// SetComment sets field value
func (o *ModelsRevokeAllCertificatesRequest) SetComment(v string) {
	o.Comment = v
}

// GetEffectiveDate returns the EffectiveDate field value if set, zero value otherwise.
func (o *ModelsRevokeAllCertificatesRequest) GetEffectiveDate() time.Time {
	if o == nil || isNil(o.EffectiveDate) {
		var ret time.Time
		return ret
	}
	return *o.EffectiveDate
}

// GetEffectiveDateOk returns a tuple with the EffectiveDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsRevokeAllCertificatesRequest) GetEffectiveDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.EffectiveDate) {
		return nil, false
	}
	return o.EffectiveDate, true
}

// HasEffectiveDate returns a boolean if a field has been set.
func (o *ModelsRevokeAllCertificatesRequest) HasEffectiveDate() bool {
	if o != nil && !isNil(o.EffectiveDate) {
		return true
	}

	return false
}

// SetEffectiveDate gets a reference to the given time.Time and assigns it to the EffectiveDate field.
func (o *ModelsRevokeAllCertificatesRequest) SetEffectiveDate(v time.Time) {
	o.EffectiveDate = &v
}

// GetIncludeRevoked returns the IncludeRevoked field value if set, zero value otherwise.
func (o *ModelsRevokeAllCertificatesRequest) GetIncludeRevoked() bool {
	if o == nil || isNil(o.IncludeRevoked) {
		var ret bool
		return ret
	}
	return *o.IncludeRevoked
}

// GetIncludeRevokedOk returns a tuple with the IncludeRevoked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsRevokeAllCertificatesRequest) GetIncludeRevokedOk() (*bool, bool) {
	if o == nil || isNil(o.IncludeRevoked) {
		return nil, false
	}
	return o.IncludeRevoked, true
}

// HasIncludeRevoked returns a boolean if a field has been set.
func (o *ModelsRevokeAllCertificatesRequest) HasIncludeRevoked() bool {
	if o != nil && !isNil(o.IncludeRevoked) {
		return true
	}

	return false
}

// SetIncludeRevoked gets a reference to the given bool and assigns it to the IncludeRevoked field.
func (o *ModelsRevokeAllCertificatesRequest) SetIncludeRevoked(v bool) {
	o.IncludeRevoked = &v
}

// GetIncludeExpired returns the IncludeExpired field value if set, zero value otherwise.
func (o *ModelsRevokeAllCertificatesRequest) GetIncludeExpired() bool {
	if o == nil || isNil(o.IncludeExpired) {
		var ret bool
		return ret
	}
	return *o.IncludeExpired
}

// GetIncludeExpiredOk returns a tuple with the IncludeExpired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsRevokeAllCertificatesRequest) GetIncludeExpiredOk() (*bool, bool) {
	if o == nil || isNil(o.IncludeExpired) {
		return nil, false
	}
	return o.IncludeExpired, true
}

// HasIncludeExpired returns a boolean if a field has been set.
func (o *ModelsRevokeAllCertificatesRequest) HasIncludeExpired() bool {
	if o != nil && !isNil(o.IncludeExpired) {
		return true
	}

	return false
}

// SetIncludeExpired gets a reference to the given bool and assigns it to the IncludeExpired field.
func (o *ModelsRevokeAllCertificatesRequest) SetIncludeExpired(v bool) {
	o.IncludeExpired = &v
}

func (o ModelsRevokeAllCertificatesRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsRevokeAllCertificatesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Query) {
		toSerialize["Query"] = o.Query
	}
	toSerialize["Reason"] = o.Reason
	toSerialize["Comment"] = o.Comment
	if !isNil(o.EffectiveDate) {
		toSerialize["EffectiveDate"] = o.EffectiveDate
	}
	if !isNil(o.IncludeRevoked) {
		toSerialize["IncludeRevoked"] = o.IncludeRevoked
	}
	if !isNil(o.IncludeExpired) {
		toSerialize["IncludeExpired"] = o.IncludeExpired
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModelsRevokeAllCertificatesRequest) UnmarshalJSON(bytes []byte) (err error) {
	varModelsRevokeAllCertificatesRequest := _ModelsRevokeAllCertificatesRequest{}

	if err = json.Unmarshal(bytes, &varModelsRevokeAllCertificatesRequest); err == nil {
		*o = ModelsRevokeAllCertificatesRequest(varModelsRevokeAllCertificatesRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Query")
		delete(additionalProperties, "Reason")
		delete(additionalProperties, "Comment")
		delete(additionalProperties, "EffectiveDate")
		delete(additionalProperties, "IncludeRevoked")
		delete(additionalProperties, "IncludeExpired")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelsRevokeAllCertificatesRequest struct {
	value *ModelsRevokeAllCertificatesRequest
	isSet bool
}

func (v NullableModelsRevokeAllCertificatesRequest) Get() *ModelsRevokeAllCertificatesRequest {
	return v.value
}

func (v *NullableModelsRevokeAllCertificatesRequest) Set(val *ModelsRevokeAllCertificatesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsRevokeAllCertificatesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsRevokeAllCertificatesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsRevokeAllCertificatesRequest(val *ModelsRevokeAllCertificatesRequest) *NullableModelsRevokeAllCertificatesRequest {
	return &NullableModelsRevokeAllCertificatesRequest{value: val, isSet: true}
}

func (v NullableModelsRevokeAllCertificatesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsRevokeAllCertificatesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

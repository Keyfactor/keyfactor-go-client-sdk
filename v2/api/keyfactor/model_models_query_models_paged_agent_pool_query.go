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

// checks if the ModelsQueryModelsPagedAgentPoolQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsQueryModelsPagedAgentPoolQuery{}

// ModelsQueryModelsPagedAgentPoolQuery struct for ModelsQueryModelsPagedAgentPoolQuery
type ModelsQueryModelsPagedAgentPoolQuery struct {
	// Contents of the query (ex: field1 -eq value1 AND field2 -gt value2)
	QueryString *string `json:"QueryString,omitempty"`
	// The current page within the result set to be returned
	PageReturned *int32 `json:"PageReturned,omitempty"`
	// Maximum number of records to be returned in a single call
	ReturnLimit *int32 `json:"ReturnLimit,omitempty"`
	// Field by which the results should be sorted (view results via Management Portal for sortable columns)
	SortField *string `json:"SortField,omitempty"`
	// Field sort direction [0=ascending, 1=descending]
	SortAscending        *int32 `json:"SortAscending,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelsQueryModelsPagedAgentPoolQuery ModelsQueryModelsPagedAgentPoolQuery

// NewModelsQueryModelsPagedAgentPoolQuery instantiates a new ModelsQueryModelsPagedAgentPoolQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsQueryModelsPagedAgentPoolQuery() *ModelsQueryModelsPagedAgentPoolQuery {
	this := ModelsQueryModelsPagedAgentPoolQuery{}
	return &this
}

// NewModelsQueryModelsPagedAgentPoolQueryWithDefaults instantiates a new ModelsQueryModelsPagedAgentPoolQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsQueryModelsPagedAgentPoolQueryWithDefaults() *ModelsQueryModelsPagedAgentPoolQuery {
	this := ModelsQueryModelsPagedAgentPoolQuery{}
	return &this
}

// GetQueryString returns the QueryString field value if set, zero value otherwise.
func (o *ModelsQueryModelsPagedAgentPoolQuery) GetQueryString() string {
	if o == nil || isNil(o.QueryString) {
		var ret string
		return ret
	}
	return *o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsQueryModelsPagedAgentPoolQuery) GetQueryStringOk() (*string, bool) {
	if o == nil || isNil(o.QueryString) {
		return nil, false
	}
	return o.QueryString, true
}

// HasQueryString returns a boolean if a field has been set.
func (o *ModelsQueryModelsPagedAgentPoolQuery) HasQueryString() bool {
	if o != nil && !isNil(o.QueryString) {
		return true
	}

	return false
}

// SetQueryString gets a reference to the given string and assigns it to the QueryString field.
func (o *ModelsQueryModelsPagedAgentPoolQuery) SetQueryString(v string) {
	o.QueryString = &v
}

// GetPageReturned returns the PageReturned field value if set, zero value otherwise.
func (o *ModelsQueryModelsPagedAgentPoolQuery) GetPageReturned() int32 {
	if o == nil || isNil(o.PageReturned) {
		var ret int32
		return ret
	}
	return *o.PageReturned
}

// GetPageReturnedOk returns a tuple with the PageReturned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsQueryModelsPagedAgentPoolQuery) GetPageReturnedOk() (*int32, bool) {
	if o == nil || isNil(o.PageReturned) {
		return nil, false
	}
	return o.PageReturned, true
}

// HasPageReturned returns a boolean if a field has been set.
func (o *ModelsQueryModelsPagedAgentPoolQuery) HasPageReturned() bool {
	if o != nil && !isNil(o.PageReturned) {
		return true
	}

	return false
}

// SetPageReturned gets a reference to the given int32 and assigns it to the PageReturned field.
func (o *ModelsQueryModelsPagedAgentPoolQuery) SetPageReturned(v int32) {
	o.PageReturned = &v
}

// GetReturnLimit returns the ReturnLimit field value if set, zero value otherwise.
func (o *ModelsQueryModelsPagedAgentPoolQuery) GetReturnLimit() int32 {
	if o == nil || isNil(o.ReturnLimit) {
		var ret int32
		return ret
	}
	return *o.ReturnLimit
}

// GetReturnLimitOk returns a tuple with the ReturnLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsQueryModelsPagedAgentPoolQuery) GetReturnLimitOk() (*int32, bool) {
	if o == nil || isNil(o.ReturnLimit) {
		return nil, false
	}
	return o.ReturnLimit, true
}

// HasReturnLimit returns a boolean if a field has been set.
func (o *ModelsQueryModelsPagedAgentPoolQuery) HasReturnLimit() bool {
	if o != nil && !isNil(o.ReturnLimit) {
		return true
	}

	return false
}

// SetReturnLimit gets a reference to the given int32 and assigns it to the ReturnLimit field.
func (o *ModelsQueryModelsPagedAgentPoolQuery) SetReturnLimit(v int32) {
	o.ReturnLimit = &v
}

// GetSortField returns the SortField field value if set, zero value otherwise.
func (o *ModelsQueryModelsPagedAgentPoolQuery) GetSortField() string {
	if o == nil || isNil(o.SortField) {
		var ret string
		return ret
	}
	return *o.SortField
}

// GetSortFieldOk returns a tuple with the SortField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsQueryModelsPagedAgentPoolQuery) GetSortFieldOk() (*string, bool) {
	if o == nil || isNil(o.SortField) {
		return nil, false
	}
	return o.SortField, true
}

// HasSortField returns a boolean if a field has been set.
func (o *ModelsQueryModelsPagedAgentPoolQuery) HasSortField() bool {
	if o != nil && !isNil(o.SortField) {
		return true
	}

	return false
}

// SetSortField gets a reference to the given string and assigns it to the SortField field.
func (o *ModelsQueryModelsPagedAgentPoolQuery) SetSortField(v string) {
	o.SortField = &v
}

// GetSortAscending returns the SortAscending field value if set, zero value otherwise.
func (o *ModelsQueryModelsPagedAgentPoolQuery) GetSortAscending() int32 {
	if o == nil || isNil(o.SortAscending) {
		var ret int32
		return ret
	}
	return *o.SortAscending
}

// GetSortAscendingOk returns a tuple with the SortAscending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsQueryModelsPagedAgentPoolQuery) GetSortAscendingOk() (*int32, bool) {
	if o == nil || isNil(o.SortAscending) {
		return nil, false
	}
	return o.SortAscending, true
}

// HasSortAscending returns a boolean if a field has been set.
func (o *ModelsQueryModelsPagedAgentPoolQuery) HasSortAscending() bool {
	if o != nil && !isNil(o.SortAscending) {
		return true
	}

	return false
}

// SetSortAscending gets a reference to the given int32 and assigns it to the SortAscending field.
func (o *ModelsQueryModelsPagedAgentPoolQuery) SetSortAscending(v int32) {
	o.SortAscending = &v
}

func (o ModelsQueryModelsPagedAgentPoolQuery) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsQueryModelsPagedAgentPoolQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.QueryString) {
		toSerialize["QueryString"] = o.QueryString
	}
	if !isNil(o.PageReturned) {
		toSerialize["PageReturned"] = o.PageReturned
	}
	if !isNil(o.ReturnLimit) {
		toSerialize["ReturnLimit"] = o.ReturnLimit
	}
	if !isNil(o.SortField) {
		toSerialize["SortField"] = o.SortField
	}
	if !isNil(o.SortAscending) {
		toSerialize["SortAscending"] = o.SortAscending
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModelsQueryModelsPagedAgentPoolQuery) UnmarshalJSON(bytes []byte) (err error) {
	varModelsQueryModelsPagedAgentPoolQuery := _ModelsQueryModelsPagedAgentPoolQuery{}

	if err = json.Unmarshal(bytes, &varModelsQueryModelsPagedAgentPoolQuery); err == nil {
		*o = ModelsQueryModelsPagedAgentPoolQuery(varModelsQueryModelsPagedAgentPoolQuery)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "QueryString")
		delete(additionalProperties, "PageReturned")
		delete(additionalProperties, "ReturnLimit")
		delete(additionalProperties, "SortField")
		delete(additionalProperties, "SortAscending")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelsQueryModelsPagedAgentPoolQuery struct {
	value *ModelsQueryModelsPagedAgentPoolQuery
	isSet bool
}

func (v NullableModelsQueryModelsPagedAgentPoolQuery) Get() *ModelsQueryModelsPagedAgentPoolQuery {
	return v.value
}

func (v *NullableModelsQueryModelsPagedAgentPoolQuery) Set(val *ModelsQueryModelsPagedAgentPoolQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsQueryModelsPagedAgentPoolQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsQueryModelsPagedAgentPoolQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsQueryModelsPagedAgentPoolQuery(val *ModelsQueryModelsPagedAgentPoolQuery) *NullableModelsQueryModelsPagedAgentPoolQuery {
	return &NullableModelsQueryModelsPagedAgentPoolQuery{value: val, isSet: true}
}

func (v NullableModelsQueryModelsPagedAgentPoolQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsQueryModelsPagedAgentPoolQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

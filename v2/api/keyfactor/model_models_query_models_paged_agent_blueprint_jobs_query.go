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

// checks if the ModelsQueryModelsPagedAgentBlueprintJobsQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsQueryModelsPagedAgentBlueprintJobsQuery{}

// ModelsQueryModelsPagedAgentBlueprintJobsQuery struct for ModelsQueryModelsPagedAgentBlueprintJobsQuery
type ModelsQueryModelsPagedAgentBlueprintJobsQuery struct {
	// The current page within the result set to be returned
	PageReturned *int32 `json:"PageReturned,omitempty"`
	// Number of records to be skipped before inclusion in the result set
	SkipCount *int32 `json:"SkipCount,omitempty"`
	// Maximum number of records to be returned in a single call
	ReturnLimit *int32 `json:"ReturnLimit,omitempty"`
	// Field by which the results should be sorted (OperationStart, OperationEnd, UserName)
	SortField *string `json:"SortField,omitempty"`
	// Field sort direction [0=ascending, 1=descending]
	SortAscending        *int32 `json:"SortAscending,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelsQueryModelsPagedAgentBlueprintJobsQuery ModelsQueryModelsPagedAgentBlueprintJobsQuery

// NewModelsQueryModelsPagedAgentBlueprintJobsQuery instantiates a new ModelsQueryModelsPagedAgentBlueprintJobsQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsQueryModelsPagedAgentBlueprintJobsQuery() *ModelsQueryModelsPagedAgentBlueprintJobsQuery {
	this := ModelsQueryModelsPagedAgentBlueprintJobsQuery{}
	return &this
}

// NewModelsQueryModelsPagedAgentBlueprintJobsQueryWithDefaults instantiates a new ModelsQueryModelsPagedAgentBlueprintJobsQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsQueryModelsPagedAgentBlueprintJobsQueryWithDefaults() *ModelsQueryModelsPagedAgentBlueprintJobsQuery {
	this := ModelsQueryModelsPagedAgentBlueprintJobsQuery{}
	return &this
}

// GetPageReturned returns the PageReturned field value if set, zero value otherwise.
func (o *ModelsQueryModelsPagedAgentBlueprintJobsQuery) GetPageReturned() int32 {
	if o == nil || isNil(o.PageReturned) {
		var ret int32
		return ret
	}
	return *o.PageReturned
}

// GetPageReturnedOk returns a tuple with the PageReturned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsQueryModelsPagedAgentBlueprintJobsQuery) GetPageReturnedOk() (*int32, bool) {
	if o == nil || isNil(o.PageReturned) {
		return nil, false
	}
	return o.PageReturned, true
}

// HasPageReturned returns a boolean if a field has been set.
func (o *ModelsQueryModelsPagedAgentBlueprintJobsQuery) HasPageReturned() bool {
	if o != nil && !isNil(o.PageReturned) {
		return true
	}

	return false
}

// SetPageReturned gets a reference to the given int32 and assigns it to the PageReturned field.
func (o *ModelsQueryModelsPagedAgentBlueprintJobsQuery) SetPageReturned(v int32) {
	o.PageReturned = &v
}

// GetSkipCount returns the SkipCount field value if set, zero value otherwise.
func (o *ModelsQueryModelsPagedAgentBlueprintJobsQuery) GetSkipCount() int32 {
	if o == nil || isNil(o.SkipCount) {
		var ret int32
		return ret
	}
	return *o.SkipCount
}

// GetSkipCountOk returns a tuple with the SkipCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsQueryModelsPagedAgentBlueprintJobsQuery) GetSkipCountOk() (*int32, bool) {
	if o == nil || isNil(o.SkipCount) {
		return nil, false
	}
	return o.SkipCount, true
}

// HasSkipCount returns a boolean if a field has been set.
func (o *ModelsQueryModelsPagedAgentBlueprintJobsQuery) HasSkipCount() bool {
	if o != nil && !isNil(o.SkipCount) {
		return true
	}

	return false
}

// SetSkipCount gets a reference to the given int32 and assigns it to the SkipCount field.
func (o *ModelsQueryModelsPagedAgentBlueprintJobsQuery) SetSkipCount(v int32) {
	o.SkipCount = &v
}

// GetReturnLimit returns the ReturnLimit field value if set, zero value otherwise.
func (o *ModelsQueryModelsPagedAgentBlueprintJobsQuery) GetReturnLimit() int32 {
	if o == nil || isNil(o.ReturnLimit) {
		var ret int32
		return ret
	}
	return *o.ReturnLimit
}

// GetReturnLimitOk returns a tuple with the ReturnLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsQueryModelsPagedAgentBlueprintJobsQuery) GetReturnLimitOk() (*int32, bool) {
	if o == nil || isNil(o.ReturnLimit) {
		return nil, false
	}
	return o.ReturnLimit, true
}

// HasReturnLimit returns a boolean if a field has been set.
func (o *ModelsQueryModelsPagedAgentBlueprintJobsQuery) HasReturnLimit() bool {
	if o != nil && !isNil(o.ReturnLimit) {
		return true
	}

	return false
}

// SetReturnLimit gets a reference to the given int32 and assigns it to the ReturnLimit field.
func (o *ModelsQueryModelsPagedAgentBlueprintJobsQuery) SetReturnLimit(v int32) {
	o.ReturnLimit = &v
}

// GetSortField returns the SortField field value if set, zero value otherwise.
func (o *ModelsQueryModelsPagedAgentBlueprintJobsQuery) GetSortField() string {
	if o == nil || isNil(o.SortField) {
		var ret string
		return ret
	}
	return *o.SortField
}

// GetSortFieldOk returns a tuple with the SortField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsQueryModelsPagedAgentBlueprintJobsQuery) GetSortFieldOk() (*string, bool) {
	if o == nil || isNil(o.SortField) {
		return nil, false
	}
	return o.SortField, true
}

// HasSortField returns a boolean if a field has been set.
func (o *ModelsQueryModelsPagedAgentBlueprintJobsQuery) HasSortField() bool {
	if o != nil && !isNil(o.SortField) {
		return true
	}

	return false
}

// SetSortField gets a reference to the given string and assigns it to the SortField field.
func (o *ModelsQueryModelsPagedAgentBlueprintJobsQuery) SetSortField(v string) {
	o.SortField = &v
}

// GetSortAscending returns the SortAscending field value if set, zero value otherwise.
func (o *ModelsQueryModelsPagedAgentBlueprintJobsQuery) GetSortAscending() int32 {
	if o == nil || isNil(o.SortAscending) {
		var ret int32
		return ret
	}
	return *o.SortAscending
}

// GetSortAscendingOk returns a tuple with the SortAscending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsQueryModelsPagedAgentBlueprintJobsQuery) GetSortAscendingOk() (*int32, bool) {
	if o == nil || isNil(o.SortAscending) {
		return nil, false
	}
	return o.SortAscending, true
}

// HasSortAscending returns a boolean if a field has been set.
func (o *ModelsQueryModelsPagedAgentBlueprintJobsQuery) HasSortAscending() bool {
	if o != nil && !isNil(o.SortAscending) {
		return true
	}

	return false
}

// SetSortAscending gets a reference to the given int32 and assigns it to the SortAscending field.
func (o *ModelsQueryModelsPagedAgentBlueprintJobsQuery) SetSortAscending(v int32) {
	o.SortAscending = &v
}

func (o ModelsQueryModelsPagedAgentBlueprintJobsQuery) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsQueryModelsPagedAgentBlueprintJobsQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PageReturned) {
		toSerialize["PageReturned"] = o.PageReturned
	}
	// skip: SkipCount is readOnly
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

func (o *ModelsQueryModelsPagedAgentBlueprintJobsQuery) UnmarshalJSON(bytes []byte) (err error) {
	varModelsQueryModelsPagedAgentBlueprintJobsQuery := _ModelsQueryModelsPagedAgentBlueprintJobsQuery{}

	if err = json.Unmarshal(bytes, &varModelsQueryModelsPagedAgentBlueprintJobsQuery); err == nil {
		*o = ModelsQueryModelsPagedAgentBlueprintJobsQuery(varModelsQueryModelsPagedAgentBlueprintJobsQuery)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "PageReturned")
		delete(additionalProperties, "SkipCount")
		delete(additionalProperties, "ReturnLimit")
		delete(additionalProperties, "SortField")
		delete(additionalProperties, "SortAscending")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelsQueryModelsPagedAgentBlueprintJobsQuery struct {
	value *ModelsQueryModelsPagedAgentBlueprintJobsQuery
	isSet bool
}

func (v NullableModelsQueryModelsPagedAgentBlueprintJobsQuery) Get() *ModelsQueryModelsPagedAgentBlueprintJobsQuery {
	return v.value
}

func (v *NullableModelsQueryModelsPagedAgentBlueprintJobsQuery) Set(val *ModelsQueryModelsPagedAgentBlueprintJobsQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsQueryModelsPagedAgentBlueprintJobsQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsQueryModelsPagedAgentBlueprintJobsQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsQueryModelsPagedAgentBlueprintJobsQuery(val *ModelsQueryModelsPagedAgentBlueprintJobsQuery) *NullableModelsQueryModelsPagedAgentBlueprintJobsQuery {
	return &NullableModelsQueryModelsPagedAgentBlueprintJobsQuery{value: val, isSet: true}
}

func (v NullableModelsQueryModelsPagedAgentBlueprintJobsQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsQueryModelsPagedAgentBlueprintJobsQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
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

// checks if the ModelsPKICertificateOperation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsPKICertificateOperation{}

// ModelsPKICertificateOperation Details of an operation done on a certificate.
type ModelsPKICertificateOperation struct {
	Id                   *int64     `json:"Id,omitempty"`
	OperationStart       *time.Time `json:"OperationStart,omitempty"`
	OperationEnd         *time.Time `json:"OperationEnd,omitempty"`
	Username             *string    `json:"Username,omitempty"`
	Comment              *string    `json:"Comment,omitempty"`
	Action               *string    `json:"Action,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelsPKICertificateOperation ModelsPKICertificateOperation

// NewModelsPKICertificateOperation instantiates a new ModelsPKICertificateOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsPKICertificateOperation() *ModelsPKICertificateOperation {
	this := ModelsPKICertificateOperation{}
	return &this
}

// NewModelsPKICertificateOperationWithDefaults instantiates a new ModelsPKICertificateOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsPKICertificateOperationWithDefaults() *ModelsPKICertificateOperation {
	this := ModelsPKICertificateOperation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelsPKICertificateOperation) GetId() int64 {
	if o == nil || isNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsPKICertificateOperation) GetIdOk() (*int64, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelsPKICertificateOperation) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ModelsPKICertificateOperation) SetId(v int64) {
	o.Id = &v
}

// GetOperationStart returns the OperationStart field value if set, zero value otherwise.
func (o *ModelsPKICertificateOperation) GetOperationStart() time.Time {
	if o == nil || isNil(o.OperationStart) {
		var ret time.Time
		return ret
	}
	return *o.OperationStart
}

// GetOperationStartOk returns a tuple with the OperationStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsPKICertificateOperation) GetOperationStartOk() (*time.Time, bool) {
	if o == nil || isNil(o.OperationStart) {
		return nil, false
	}
	return o.OperationStart, true
}

// HasOperationStart returns a boolean if a field has been set.
func (o *ModelsPKICertificateOperation) HasOperationStart() bool {
	if o != nil && !isNil(o.OperationStart) {
		return true
	}

	return false
}

// SetOperationStart gets a reference to the given time.Time and assigns it to the OperationStart field.
func (o *ModelsPKICertificateOperation) SetOperationStart(v time.Time) {
	o.OperationStart = &v
}

// GetOperationEnd returns the OperationEnd field value if set, zero value otherwise.
func (o *ModelsPKICertificateOperation) GetOperationEnd() time.Time {
	if o == nil || isNil(o.OperationEnd) {
		var ret time.Time
		return ret
	}
	return *o.OperationEnd
}

// GetOperationEndOk returns a tuple with the OperationEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsPKICertificateOperation) GetOperationEndOk() (*time.Time, bool) {
	if o == nil || isNil(o.OperationEnd) {
		return nil, false
	}
	return o.OperationEnd, true
}

// HasOperationEnd returns a boolean if a field has been set.
func (o *ModelsPKICertificateOperation) HasOperationEnd() bool {
	if o != nil && !isNil(o.OperationEnd) {
		return true
	}

	return false
}

// SetOperationEnd gets a reference to the given time.Time and assigns it to the OperationEnd field.
func (o *ModelsPKICertificateOperation) SetOperationEnd(v time.Time) {
	o.OperationEnd = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ModelsPKICertificateOperation) GetUsername() string {
	if o == nil || isNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsPKICertificateOperation) GetUsernameOk() (*string, bool) {
	if o == nil || isNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ModelsPKICertificateOperation) HasUsername() bool {
	if o != nil && !isNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ModelsPKICertificateOperation) SetUsername(v string) {
	o.Username = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ModelsPKICertificateOperation) GetComment() string {
	if o == nil || isNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsPKICertificateOperation) GetCommentOk() (*string, bool) {
	if o == nil || isNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ModelsPKICertificateOperation) HasComment() bool {
	if o != nil && !isNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ModelsPKICertificateOperation) SetComment(v string) {
	o.Comment = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *ModelsPKICertificateOperation) GetAction() string {
	if o == nil || isNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsPKICertificateOperation) GetActionOk() (*string, bool) {
	if o == nil || isNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *ModelsPKICertificateOperation) HasAction() bool {
	if o != nil && !isNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *ModelsPKICertificateOperation) SetAction(v string) {
	o.Action = &v
}

func (o ModelsPKICertificateOperation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsPKICertificateOperation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !isNil(o.OperationStart) {
		toSerialize["OperationStart"] = o.OperationStart
	}
	if !isNil(o.OperationEnd) {
		toSerialize["OperationEnd"] = o.OperationEnd
	}
	if !isNil(o.Username) {
		toSerialize["Username"] = o.Username
	}
	if !isNil(o.Comment) {
		toSerialize["Comment"] = o.Comment
	}
	if !isNil(o.Action) {
		toSerialize["Action"] = o.Action
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModelsPKICertificateOperation) UnmarshalJSON(bytes []byte) (err error) {
	varModelsPKICertificateOperation := _ModelsPKICertificateOperation{}

	if err = json.Unmarshal(bytes, &varModelsPKICertificateOperation); err == nil {
		*o = ModelsPKICertificateOperation(varModelsPKICertificateOperation)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Id")
		delete(additionalProperties, "OperationStart")
		delete(additionalProperties, "OperationEnd")
		delete(additionalProperties, "Username")
		delete(additionalProperties, "Comment")
		delete(additionalProperties, "Action")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelsPKICertificateOperation struct {
	value *ModelsPKICertificateOperation
	isSet bool
}

func (v NullableModelsPKICertificateOperation) Get() *ModelsPKICertificateOperation {
	return v.value
}

func (v *NullableModelsPKICertificateOperation) Set(val *ModelsPKICertificateOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsPKICertificateOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsPKICertificateOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsPKICertificateOperation(val *ModelsPKICertificateOperation) *NullableModelsPKICertificateOperation {
	return &NullableModelsPKICertificateOperation{value: val, isSet: true}
}

func (v NullableModelsPKICertificateOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsPKICertificateOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

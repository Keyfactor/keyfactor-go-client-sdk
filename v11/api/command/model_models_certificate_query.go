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

// checks if the ModelsCertificateQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsCertificateQuery{}

// ModelsCertificateQuery struct for ModelsCertificateQuery
type ModelsCertificateQuery struct {
	Id               *int32         `json:"id,omitempty"`
	Name             NullableString `json:"name,omitempty"`
	Description      NullableString `json:"description,omitempty"`
	Automated        *bool          `json:"automated,omitempty"`
	Content          NullableString `json:"content,omitempty"`
	DuplicationField *int32         `json:"duplicationField,omitempty"`
	ShowOnDashboard  *bool          `json:"showOnDashboard,omitempty"`
	Favorite         *bool          `json:"favorite,omitempty"`
}

// NewModelsCertificateQuery instantiates a new ModelsCertificateQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsCertificateQuery() *ModelsCertificateQuery {
	this := ModelsCertificateQuery{}
	return &this
}

// NewModelsCertificateQueryWithDefaults instantiates a new ModelsCertificateQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsCertificateQueryWithDefaults() *ModelsCertificateQuery {
	this := ModelsCertificateQuery{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelsCertificateQuery) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertificateQuery) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelsCertificateQuery) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ModelsCertificateQuery) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelsCertificateQuery) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelsCertificateQuery) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ModelsCertificateQuery) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ModelsCertificateQuery) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *ModelsCertificateQuery) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ModelsCertificateQuery) UnsetName() {
	o.Name.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelsCertificateQuery) GetDescription() string {
	if o == nil || isNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelsCertificateQuery) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ModelsCertificateQuery) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ModelsCertificateQuery) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ModelsCertificateQuery) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ModelsCertificateQuery) UnsetDescription() {
	o.Description.Unset()
}

// GetAutomated returns the Automated field value if set, zero value otherwise.
func (o *ModelsCertificateQuery) GetAutomated() bool {
	if o == nil || isNil(o.Automated) {
		var ret bool
		return ret
	}
	return *o.Automated
}

// GetAutomatedOk returns a tuple with the Automated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertificateQuery) GetAutomatedOk() (*bool, bool) {
	if o == nil || isNil(o.Automated) {
		return nil, false
	}
	return o.Automated, true
}

// HasAutomated returns a boolean if a field has been set.
func (o *ModelsCertificateQuery) HasAutomated() bool {
	if o != nil && !isNil(o.Automated) {
		return true
	}

	return false
}

// SetAutomated gets a reference to the given bool and assigns it to the Automated field.
func (o *ModelsCertificateQuery) SetAutomated(v bool) {
	o.Automated = &v
}

// GetContent returns the Content field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelsCertificateQuery) GetContent() string {
	if o == nil || isNil(o.Content.Get()) {
		var ret string
		return ret
	}
	return *o.Content.Get()
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelsCertificateQuery) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Content.Get(), o.Content.IsSet()
}

// HasContent returns a boolean if a field has been set.
func (o *ModelsCertificateQuery) HasContent() bool {
	if o != nil && o.Content.IsSet() {
		return true
	}

	return false
}

// SetContent gets a reference to the given NullableString and assigns it to the Content field.
func (o *ModelsCertificateQuery) SetContent(v string) {
	o.Content.Set(&v)
}

// SetContentNil sets the value for Content to be an explicit nil
func (o *ModelsCertificateQuery) SetContentNil() {
	o.Content.Set(nil)
}

// UnsetContent ensures that no value is present for Content, not even an explicit nil
func (o *ModelsCertificateQuery) UnsetContent() {
	o.Content.Unset()
}

// GetDuplicationField returns the DuplicationField field value if set, zero value otherwise.
func (o *ModelsCertificateQuery) GetDuplicationField() int32 {
	if o == nil || isNil(o.DuplicationField) {
		var ret int32
		return ret
	}
	return *o.DuplicationField
}

// GetDuplicationFieldOk returns a tuple with the DuplicationField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertificateQuery) GetDuplicationFieldOk() (*int32, bool) {
	if o == nil || isNil(o.DuplicationField) {
		return nil, false
	}
	return o.DuplicationField, true
}

// HasDuplicationField returns a boolean if a field has been set.
func (o *ModelsCertificateQuery) HasDuplicationField() bool {
	if o != nil && !isNil(o.DuplicationField) {
		return true
	}

	return false
}

// SetDuplicationField gets a reference to the given int32 and assigns it to the DuplicationField field.
func (o *ModelsCertificateQuery) SetDuplicationField(v int32) {
	o.DuplicationField = &v
}

// GetShowOnDashboard returns the ShowOnDashboard field value if set, zero value otherwise.
func (o *ModelsCertificateQuery) GetShowOnDashboard() bool {
	if o == nil || isNil(o.ShowOnDashboard) {
		var ret bool
		return ret
	}
	return *o.ShowOnDashboard
}

// GetShowOnDashboardOk returns a tuple with the ShowOnDashboard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertificateQuery) GetShowOnDashboardOk() (*bool, bool) {
	if o == nil || isNil(o.ShowOnDashboard) {
		return nil, false
	}
	return o.ShowOnDashboard, true
}

// HasShowOnDashboard returns a boolean if a field has been set.
func (o *ModelsCertificateQuery) HasShowOnDashboard() bool {
	if o != nil && !isNil(o.ShowOnDashboard) {
		return true
	}

	return false
}

// SetShowOnDashboard gets a reference to the given bool and assigns it to the ShowOnDashboard field.
func (o *ModelsCertificateQuery) SetShowOnDashboard(v bool) {
	o.ShowOnDashboard = &v
}

// GetFavorite returns the Favorite field value if set, zero value otherwise.
func (o *ModelsCertificateQuery) GetFavorite() bool {
	if o == nil || isNil(o.Favorite) {
		var ret bool
		return ret
	}
	return *o.Favorite
}

// GetFavoriteOk returns a tuple with the Favorite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertificateQuery) GetFavoriteOk() (*bool, bool) {
	if o == nil || isNil(o.Favorite) {
		return nil, false
	}
	return o.Favorite, true
}

// HasFavorite returns a boolean if a field has been set.
func (o *ModelsCertificateQuery) HasFavorite() bool {
	if o != nil && !isNil(o.Favorite) {
		return true
	}

	return false
}

// SetFavorite gets a reference to the given bool and assigns it to the Favorite field.
func (o *ModelsCertificateQuery) SetFavorite(v bool) {
	o.Favorite = &v
}

func (o ModelsCertificateQuery) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsCertificateQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !isNil(o.Automated) {
		toSerialize["automated"] = o.Automated
	}
	if o.Content.IsSet() {
		toSerialize["content"] = o.Content.Get()
	}
	if !isNil(o.DuplicationField) {
		toSerialize["duplicationField"] = o.DuplicationField
	}
	if !isNil(o.ShowOnDashboard) {
		toSerialize["showOnDashboard"] = o.ShowOnDashboard
	}
	if !isNil(o.Favorite) {
		toSerialize["favorite"] = o.Favorite
	}
	return toSerialize, nil
}

type NullableModelsCertificateQuery struct {
	value *ModelsCertificateQuery
	isSet bool
}

func (v NullableModelsCertificateQuery) Get() *ModelsCertificateQuery {
	return v.value
}

func (v *NullableModelsCertificateQuery) Set(val *ModelsCertificateQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsCertificateQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsCertificateQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsCertificateQuery(val *ModelsCertificateQuery) *NullableModelsCertificateQuery {
	return &NullableModelsCertificateQuery{value: val, isSet: true}
}

func (v NullableModelsCertificateQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsCertificateQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
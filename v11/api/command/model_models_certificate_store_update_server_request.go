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

// checks if the ModelsCertificateStoreUpdateServerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsCertificateStoreUpdateServerRequest{}

// ModelsCertificateStoreUpdateServerRequest struct for ModelsCertificateStoreUpdateServerRequest
type ModelsCertificateStoreUpdateServerRequest struct {
	Id        int32                    `json:"id"`
	Username  ModelsKeyfactorAPISecret `json:"username"`
	Password  ModelsKeyfactorAPISecret `json:"password"`
	UseSSL    bool                     `json:"useSSL"`
	Name      string                   `json:"name"`
	Container NullableInt32            `json:"container,omitempty"`
}

// NewModelsCertificateStoreUpdateServerRequest instantiates a new ModelsCertificateStoreUpdateServerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsCertificateStoreUpdateServerRequest(id int32, username ModelsKeyfactorAPISecret, password ModelsKeyfactorAPISecret, useSSL bool, name string) *ModelsCertificateStoreUpdateServerRequest {
	this := ModelsCertificateStoreUpdateServerRequest{}
	this.Id = id
	this.Username = username
	this.Password = password
	this.UseSSL = useSSL
	this.Name = name
	return &this
}

// NewModelsCertificateStoreUpdateServerRequestWithDefaults instantiates a new ModelsCertificateStoreUpdateServerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsCertificateStoreUpdateServerRequestWithDefaults() *ModelsCertificateStoreUpdateServerRequest {
	this := ModelsCertificateStoreUpdateServerRequest{}
	return &this
}

// GetId returns the Id field value
func (o *ModelsCertificateStoreUpdateServerRequest) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ModelsCertificateStoreUpdateServerRequest) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ModelsCertificateStoreUpdateServerRequest) SetId(v int32) {
	o.Id = v
}

// GetUsername returns the Username field value
func (o *ModelsCertificateStoreUpdateServerRequest) GetUsername() ModelsKeyfactorAPISecret {
	if o == nil {
		var ret ModelsKeyfactorAPISecret
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *ModelsCertificateStoreUpdateServerRequest) GetUsernameOk() (*ModelsKeyfactorAPISecret, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *ModelsCertificateStoreUpdateServerRequest) SetUsername(v ModelsKeyfactorAPISecret) {
	o.Username = v
}

// GetPassword returns the Password field value
func (o *ModelsCertificateStoreUpdateServerRequest) GetPassword() ModelsKeyfactorAPISecret {
	if o == nil {
		var ret ModelsKeyfactorAPISecret
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *ModelsCertificateStoreUpdateServerRequest) GetPasswordOk() (*ModelsKeyfactorAPISecret, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *ModelsCertificateStoreUpdateServerRequest) SetPassword(v ModelsKeyfactorAPISecret) {
	o.Password = v
}

// GetUseSSL returns the UseSSL field value
func (o *ModelsCertificateStoreUpdateServerRequest) GetUseSSL() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseSSL
}

// GetUseSSLOk returns a tuple with the UseSSL field value
// and a boolean to check if the value has been set.
func (o *ModelsCertificateStoreUpdateServerRequest) GetUseSSLOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseSSL, true
}

// SetUseSSL sets field value
func (o *ModelsCertificateStoreUpdateServerRequest) SetUseSSL(v bool) {
	o.UseSSL = v
}

// GetName returns the Name field value
func (o *ModelsCertificateStoreUpdateServerRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ModelsCertificateStoreUpdateServerRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ModelsCertificateStoreUpdateServerRequest) SetName(v string) {
	o.Name = v
}

// GetContainer returns the Container field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelsCertificateStoreUpdateServerRequest) GetContainer() int32 {
	if o == nil || isNil(o.Container.Get()) {
		var ret int32
		return ret
	}
	return *o.Container.Get()
}

// GetContainerOk returns a tuple with the Container field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelsCertificateStoreUpdateServerRequest) GetContainerOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Container.Get(), o.Container.IsSet()
}

// HasContainer returns a boolean if a field has been set.
func (o *ModelsCertificateStoreUpdateServerRequest) HasContainer() bool {
	if o != nil && o.Container.IsSet() {
		return true
	}

	return false
}

// SetContainer gets a reference to the given NullableInt32 and assigns it to the Container field.
func (o *ModelsCertificateStoreUpdateServerRequest) SetContainer(v int32) {
	o.Container.Set(&v)
}

// SetContainerNil sets the value for Container to be an explicit nil
func (o *ModelsCertificateStoreUpdateServerRequest) SetContainerNil() {
	o.Container.Set(nil)
}

// UnsetContainer ensures that no value is present for Container, not even an explicit nil
func (o *ModelsCertificateStoreUpdateServerRequest) UnsetContainer() {
	o.Container.Unset()
}

func (o ModelsCertificateStoreUpdateServerRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsCertificateStoreUpdateServerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["username"] = o.Username
	toSerialize["password"] = o.Password
	toSerialize["useSSL"] = o.UseSSL
	toSerialize["name"] = o.Name
	if o.Container.IsSet() {
		toSerialize["container"] = o.Container.Get()
	}
	return toSerialize, nil
}

type NullableModelsCertificateStoreUpdateServerRequest struct {
	value *ModelsCertificateStoreUpdateServerRequest
	isSet bool
}

func (v NullableModelsCertificateStoreUpdateServerRequest) Get() *ModelsCertificateStoreUpdateServerRequest {
	return v.value
}

func (v *NullableModelsCertificateStoreUpdateServerRequest) Set(val *ModelsCertificateStoreUpdateServerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsCertificateStoreUpdateServerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsCertificateStoreUpdateServerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsCertificateStoreUpdateServerRequest(val *ModelsCertificateStoreUpdateServerRequest) *NullableModelsCertificateStoreUpdateServerRequest {
	return &NullableModelsCertificateStoreUpdateServerRequest{value: val, isSet: true}
}

func (v NullableModelsCertificateStoreUpdateServerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsCertificateStoreUpdateServerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
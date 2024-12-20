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

// checks if the ModelsCertificateStoreCreateServerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsCertificateStoreCreateServerRequest{}

// ModelsCertificateStoreCreateServerRequest struct for ModelsCertificateStoreCreateServerRequest
type ModelsCertificateStoreCreateServerRequest struct {
	Username             ModelsKeyfactorAPISecret `json:"Username"`
	Password             ModelsKeyfactorAPISecret `json:"Password"`
	UseSSL               bool                     `json:"UseSSL"`
	ServerType           int32                    `json:"ServerType"`
	Name                 string                   `json:"Name"`
	Container            *int32                   `json:"Container,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelsCertificateStoreCreateServerRequest ModelsCertificateStoreCreateServerRequest

// NewModelsCertificateStoreCreateServerRequest instantiates a new ModelsCertificateStoreCreateServerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsCertificateStoreCreateServerRequest(username ModelsKeyfactorAPISecret, password ModelsKeyfactorAPISecret, useSSL bool, serverType int32, name string) *ModelsCertificateStoreCreateServerRequest {
	this := ModelsCertificateStoreCreateServerRequest{}
	this.Username = username
	this.Password = password
	this.UseSSL = useSSL
	this.ServerType = serverType
	this.Name = name
	return &this
}

// NewModelsCertificateStoreCreateServerRequestWithDefaults instantiates a new ModelsCertificateStoreCreateServerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsCertificateStoreCreateServerRequestWithDefaults() *ModelsCertificateStoreCreateServerRequest {
	this := ModelsCertificateStoreCreateServerRequest{}
	return &this
}

// GetUsername returns the Username field value
func (o *ModelsCertificateStoreCreateServerRequest) GetUsername() ModelsKeyfactorAPISecret {
	if o == nil {
		var ret ModelsKeyfactorAPISecret
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *ModelsCertificateStoreCreateServerRequest) GetUsernameOk() (*ModelsKeyfactorAPISecret, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *ModelsCertificateStoreCreateServerRequest) SetUsername(v ModelsKeyfactorAPISecret) {
	o.Username = v
}

// GetPassword returns the Password field value
func (o *ModelsCertificateStoreCreateServerRequest) GetPassword() ModelsKeyfactorAPISecret {
	if o == nil {
		var ret ModelsKeyfactorAPISecret
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *ModelsCertificateStoreCreateServerRequest) GetPasswordOk() (*ModelsKeyfactorAPISecret, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *ModelsCertificateStoreCreateServerRequest) SetPassword(v ModelsKeyfactorAPISecret) {
	o.Password = v
}

// GetUseSSL returns the UseSSL field value
func (o *ModelsCertificateStoreCreateServerRequest) GetUseSSL() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseSSL
}

// GetUseSSLOk returns a tuple with the UseSSL field value
// and a boolean to check if the value has been set.
func (o *ModelsCertificateStoreCreateServerRequest) GetUseSSLOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseSSL, true
}

// SetUseSSL sets field value
func (o *ModelsCertificateStoreCreateServerRequest) SetUseSSL(v bool) {
	o.UseSSL = v
}

// GetServerType returns the ServerType field value
func (o *ModelsCertificateStoreCreateServerRequest) GetServerType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ServerType
}

// GetServerTypeOk returns a tuple with the ServerType field value
// and a boolean to check if the value has been set.
func (o *ModelsCertificateStoreCreateServerRequest) GetServerTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerType, true
}

// SetServerType sets field value
func (o *ModelsCertificateStoreCreateServerRequest) SetServerType(v int32) {
	o.ServerType = v
}

// GetName returns the Name field value
func (o *ModelsCertificateStoreCreateServerRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ModelsCertificateStoreCreateServerRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ModelsCertificateStoreCreateServerRequest) SetName(v string) {
	o.Name = v
}

// GetContainer returns the Container field value if set, zero value otherwise.
func (o *ModelsCertificateStoreCreateServerRequest) GetContainer() int32 {
	if o == nil || isNil(o.Container) {
		var ret int32
		return ret
	}
	return *o.Container
}

// GetContainerOk returns a tuple with the Container field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertificateStoreCreateServerRequest) GetContainerOk() (*int32, bool) {
	if o == nil || isNil(o.Container) {
		return nil, false
	}
	return o.Container, true
}

// HasContainer returns a boolean if a field has been set.
func (o *ModelsCertificateStoreCreateServerRequest) HasContainer() bool {
	if o != nil && !isNil(o.Container) {
		return true
	}

	return false
}

// SetContainer gets a reference to the given int32 and assigns it to the Container field.
func (o *ModelsCertificateStoreCreateServerRequest) SetContainer(v int32) {
	o.Container = &v
}

func (o ModelsCertificateStoreCreateServerRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsCertificateStoreCreateServerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Username"] = o.Username
	toSerialize["Password"] = o.Password
	toSerialize["UseSSL"] = o.UseSSL
	toSerialize["ServerType"] = o.ServerType
	toSerialize["Name"] = o.Name
	if !isNil(o.Container) {
		toSerialize["Container"] = o.Container
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModelsCertificateStoreCreateServerRequest) UnmarshalJSON(bytes []byte) (err error) {
	varModelsCertificateStoreCreateServerRequest := _ModelsCertificateStoreCreateServerRequest{}

	if err = json.Unmarshal(bytes, &varModelsCertificateStoreCreateServerRequest); err == nil {
		*o = ModelsCertificateStoreCreateServerRequest(varModelsCertificateStoreCreateServerRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Username")
		delete(additionalProperties, "Password")
		delete(additionalProperties, "UseSSL")
		delete(additionalProperties, "ServerType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Container")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelsCertificateStoreCreateServerRequest struct {
	value *ModelsCertificateStoreCreateServerRequest
	isSet bool
}

func (v NullableModelsCertificateStoreCreateServerRequest) Get() *ModelsCertificateStoreCreateServerRequest {
	return v.value
}

func (v *NullableModelsCertificateStoreCreateServerRequest) Set(val *ModelsCertificateStoreCreateServerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsCertificateStoreCreateServerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsCertificateStoreCreateServerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsCertificateStoreCreateServerRequest(val *ModelsCertificateStoreCreateServerRequest) *NullableModelsCertificateStoreCreateServerRequest {
	return &NullableModelsCertificateStoreCreateServerRequest{value: val, isSet: true}
}

func (v NullableModelsCertificateStoreCreateServerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsCertificateStoreCreateServerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

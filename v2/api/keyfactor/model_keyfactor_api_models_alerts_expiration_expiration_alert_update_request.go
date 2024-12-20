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

// checks if the KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest{}

// KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest struct for KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest
type KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest struct {
	Id                     *int32                                                       `json:"Id,omitempty"`
	DisplayName            string                                                       `json:"DisplayName"`
	Subject                string                                                       `json:"Subject"`
	Message                string                                                       `json:"Message"`
	ExpirationWarningDays  int32                                                        `json:"ExpirationWarningDays"`
	CertificateQueryId     *int32                                                       `json:"CertificateQueryId,omitempty"`
	RegisteredEventHandler *KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest `json:"RegisteredEventHandler,omitempty"`
	Recipients             []string                                                     `json:"Recipients,omitempty"`
	EventHandlerParameters []KeyfactorApiModelsEventHandlerEventHandlerParameterRequest `json:"EventHandlerParameters,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest

// NewKeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest instantiates a new KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest(displayName string, subject string, message string, expirationWarningDays int32) *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest {
	this := KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest{}
	this.DisplayName = displayName
	this.Subject = subject
	this.Message = message
	this.ExpirationWarningDays = expirationWarningDays
	return &this
}

// NewKeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequestWithDefaults instantiates a new KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequestWithDefaults() *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest {
	this := KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) SetId(v int32) {
	o.Id = &v
}

// GetDisplayName returns the DisplayName field value
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetSubject returns the Subject field value
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) SetSubject(v string) {
	o.Subject = v
}

// GetMessage returns the Message field value
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) SetMessage(v string) {
	o.Message = v
}

// GetExpirationWarningDays returns the ExpirationWarningDays field value
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) GetExpirationWarningDays() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExpirationWarningDays
}

// GetExpirationWarningDaysOk returns a tuple with the ExpirationWarningDays field value
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) GetExpirationWarningDaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpirationWarningDays, true
}

// SetExpirationWarningDays sets field value
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) SetExpirationWarningDays(v int32) {
	o.ExpirationWarningDays = v
}

// GetCertificateQueryId returns the CertificateQueryId field value if set, zero value otherwise.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) GetCertificateQueryId() int32 {
	if o == nil || isNil(o.CertificateQueryId) {
		var ret int32
		return ret
	}
	return *o.CertificateQueryId
}

// GetCertificateQueryIdOk returns a tuple with the CertificateQueryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) GetCertificateQueryIdOk() (*int32, bool) {
	if o == nil || isNil(o.CertificateQueryId) {
		return nil, false
	}
	return o.CertificateQueryId, true
}

// HasCertificateQueryId returns a boolean if a field has been set.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) HasCertificateQueryId() bool {
	if o != nil && !isNil(o.CertificateQueryId) {
		return true
	}

	return false
}

// SetCertificateQueryId gets a reference to the given int32 and assigns it to the CertificateQueryId field.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) SetCertificateQueryId(v int32) {
	o.CertificateQueryId = &v
}

// GetRegisteredEventHandler returns the RegisteredEventHandler field value if set, zero value otherwise.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) GetRegisteredEventHandler() KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest {
	if o == nil || isNil(o.RegisteredEventHandler) {
		var ret KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest
		return ret
	}
	return *o.RegisteredEventHandler
}

// GetRegisteredEventHandlerOk returns a tuple with the RegisteredEventHandler field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) GetRegisteredEventHandlerOk() (*KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest, bool) {
	if o == nil || isNil(o.RegisteredEventHandler) {
		return nil, false
	}
	return o.RegisteredEventHandler, true
}

// HasRegisteredEventHandler returns a boolean if a field has been set.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) HasRegisteredEventHandler() bool {
	if o != nil && !isNil(o.RegisteredEventHandler) {
		return true
	}

	return false
}

// SetRegisteredEventHandler gets a reference to the given KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest and assigns it to the RegisteredEventHandler field.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) SetRegisteredEventHandler(v KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest) {
	o.RegisteredEventHandler = &v
}

// GetRecipients returns the Recipients field value if set, zero value otherwise.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) GetRecipients() []string {
	if o == nil || isNil(o.Recipients) {
		var ret []string
		return ret
	}
	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) GetRecipientsOk() ([]string, bool) {
	if o == nil || isNil(o.Recipients) {
		return nil, false
	}
	return o.Recipients, true
}

// HasRecipients returns a boolean if a field has been set.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) HasRecipients() bool {
	if o != nil && !isNil(o.Recipients) {
		return true
	}

	return false
}

// SetRecipients gets a reference to the given []string and assigns it to the Recipients field.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) SetRecipients(v []string) {
	o.Recipients = v
}

// GetEventHandlerParameters returns the EventHandlerParameters field value if set, zero value otherwise.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) GetEventHandlerParameters() []KeyfactorApiModelsEventHandlerEventHandlerParameterRequest {
	if o == nil || isNil(o.EventHandlerParameters) {
		var ret []KeyfactorApiModelsEventHandlerEventHandlerParameterRequest
		return ret
	}
	return o.EventHandlerParameters
}

// GetEventHandlerParametersOk returns a tuple with the EventHandlerParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) GetEventHandlerParametersOk() ([]KeyfactorApiModelsEventHandlerEventHandlerParameterRequest, bool) {
	if o == nil || isNil(o.EventHandlerParameters) {
		return nil, false
	}
	return o.EventHandlerParameters, true
}

// HasEventHandlerParameters returns a boolean if a field has been set.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) HasEventHandlerParameters() bool {
	if o != nil && !isNil(o.EventHandlerParameters) {
		return true
	}

	return false
}

// SetEventHandlerParameters gets a reference to the given []KeyfactorApiModelsEventHandlerEventHandlerParameterRequest and assigns it to the EventHandlerParameters field.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) SetEventHandlerParameters(v []KeyfactorApiModelsEventHandlerEventHandlerParameterRequest) {
	o.EventHandlerParameters = v
}

func (o KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	toSerialize["DisplayName"] = o.DisplayName
	toSerialize["Subject"] = o.Subject
	toSerialize["Message"] = o.Message
	toSerialize["ExpirationWarningDays"] = o.ExpirationWarningDays
	if !isNil(o.CertificateQueryId) {
		toSerialize["CertificateQueryId"] = o.CertificateQueryId
	}
	if !isNil(o.RegisteredEventHandler) {
		toSerialize["RegisteredEventHandler"] = o.RegisteredEventHandler
	}
	if !isNil(o.Recipients) {
		toSerialize["Recipients"] = o.Recipients
	}
	if !isNil(o.EventHandlerParameters) {
		toSerialize["EventHandlerParameters"] = o.EventHandlerParameters
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) UnmarshalJSON(bytes []byte) (err error) {
	varKeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest := _KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest{}

	if err = json.Unmarshal(bytes, &varKeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest); err == nil {
		*o = KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest(varKeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Id")
		delete(additionalProperties, "DisplayName")
		delete(additionalProperties, "Subject")
		delete(additionalProperties, "Message")
		delete(additionalProperties, "ExpirationWarningDays")
		delete(additionalProperties, "CertificateQueryId")
		delete(additionalProperties, "RegisteredEventHandler")
		delete(additionalProperties, "Recipients")
		delete(additionalProperties, "EventHandlerParameters")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest struct {
	value *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest
	isSet bool
}

func (v NullableKeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) Get() *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest {
	return v.value
}

func (v *NullableKeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) Set(val *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest(val *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) *NullableKeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest {
	return &NullableKeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest{value: val, isSet: true}
}

func (v NullableKeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

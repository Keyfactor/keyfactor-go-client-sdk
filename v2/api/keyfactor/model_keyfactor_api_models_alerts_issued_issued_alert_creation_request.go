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

// checks if the KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest{}

// KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest struct for KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest
type KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest struct {
	DisplayName            string                                                       `json:"DisplayName"`
	Subject                string                                                       `json:"Subject"`
	Message                string                                                       `json:"Message"`
	TemplateId             *int32                                                       `json:"TemplateId,omitempty"`
	RegisteredEventHandler *KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest `json:"RegisteredEventHandler,omitempty"`
	Recipients             []string                                                     `json:"Recipients,omitempty"`
	EventHandlerParameters []KeyfactorApiModelsEventHandlerEventHandlerParameterRequest `json:"EventHandlerParameters,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest

// NewKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest instantiates a new KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest(displayName string, subject string, message string) *KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest {
	this := KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest{}
	this.DisplayName = displayName
	this.Subject = subject
	this.Message = message
	return &this
}

// NewKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequestWithDefaults instantiates a new KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequestWithDefaults() *KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest {
	this := KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest{}
	return &this
}

// GetDisplayName returns the DisplayName field value
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetSubject returns the Subject field value
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) SetSubject(v string) {
	o.Subject = v
}

// GetMessage returns the Message field value
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) SetMessage(v string) {
	o.Message = v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) GetTemplateId() int32 {
	if o == nil || isNil(o.TemplateId) {
		var ret int32
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) GetTemplateIdOk() (*int32, bool) {
	if o == nil || isNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) HasTemplateId() bool {
	if o != nil && !isNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given int32 and assigns it to the TemplateId field.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) SetTemplateId(v int32) {
	o.TemplateId = &v
}

// GetRegisteredEventHandler returns the RegisteredEventHandler field value if set, zero value otherwise.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) GetRegisteredEventHandler() KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest {
	if o == nil || isNil(o.RegisteredEventHandler) {
		var ret KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest
		return ret
	}
	return *o.RegisteredEventHandler
}

// GetRegisteredEventHandlerOk returns a tuple with the RegisteredEventHandler field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) GetRegisteredEventHandlerOk() (*KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest, bool) {
	if o == nil || isNil(o.RegisteredEventHandler) {
		return nil, false
	}
	return o.RegisteredEventHandler, true
}

// HasRegisteredEventHandler returns a boolean if a field has been set.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) HasRegisteredEventHandler() bool {
	if o != nil && !isNil(o.RegisteredEventHandler) {
		return true
	}

	return false
}

// SetRegisteredEventHandler gets a reference to the given KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest and assigns it to the RegisteredEventHandler field.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) SetRegisteredEventHandler(v KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest) {
	o.RegisteredEventHandler = &v
}

// GetRecipients returns the Recipients field value if set, zero value otherwise.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) GetRecipients() []string {
	if o == nil || isNil(o.Recipients) {
		var ret []string
		return ret
	}
	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) GetRecipientsOk() ([]string, bool) {
	if o == nil || isNil(o.Recipients) {
		return nil, false
	}
	return o.Recipients, true
}

// HasRecipients returns a boolean if a field has been set.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) HasRecipients() bool {
	if o != nil && !isNil(o.Recipients) {
		return true
	}

	return false
}

// SetRecipients gets a reference to the given []string and assigns it to the Recipients field.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) SetRecipients(v []string) {
	o.Recipients = v
}

// GetEventHandlerParameters returns the EventHandlerParameters field value if set, zero value otherwise.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) GetEventHandlerParameters() []KeyfactorApiModelsEventHandlerEventHandlerParameterRequest {
	if o == nil || isNil(o.EventHandlerParameters) {
		var ret []KeyfactorApiModelsEventHandlerEventHandlerParameterRequest
		return ret
	}
	return o.EventHandlerParameters
}

// GetEventHandlerParametersOk returns a tuple with the EventHandlerParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) GetEventHandlerParametersOk() ([]KeyfactorApiModelsEventHandlerEventHandlerParameterRequest, bool) {
	if o == nil || isNil(o.EventHandlerParameters) {
		return nil, false
	}
	return o.EventHandlerParameters, true
}

// HasEventHandlerParameters returns a boolean if a field has been set.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) HasEventHandlerParameters() bool {
	if o != nil && !isNil(o.EventHandlerParameters) {
		return true
	}

	return false
}

// SetEventHandlerParameters gets a reference to the given []KeyfactorApiModelsEventHandlerEventHandlerParameterRequest and assigns it to the EventHandlerParameters field.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) SetEventHandlerParameters(v []KeyfactorApiModelsEventHandlerEventHandlerParameterRequest) {
	o.EventHandlerParameters = v
}

func (o KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["DisplayName"] = o.DisplayName
	toSerialize["Subject"] = o.Subject
	toSerialize["Message"] = o.Message
	if !isNil(o.TemplateId) {
		toSerialize["TemplateId"] = o.TemplateId
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

func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) UnmarshalJSON(bytes []byte) (err error) {
	varKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest := _KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest{}

	if err = json.Unmarshal(bytes, &varKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest); err == nil {
		*o = KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest(varKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "DisplayName")
		delete(additionalProperties, "Subject")
		delete(additionalProperties, "Message")
		delete(additionalProperties, "TemplateId")
		delete(additionalProperties, "RegisteredEventHandler")
		delete(additionalProperties, "Recipients")
		delete(additionalProperties, "EventHandlerParameters")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest struct {
	value *KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest
	isSet bool
}

func (v NullableKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) Get() *KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest {
	return v.value
}

func (v *NullableKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) Set(val *KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest(val *KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) *NullableKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest {
	return &NullableKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest{value: val, isSet: true}
}

func (v NullableKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

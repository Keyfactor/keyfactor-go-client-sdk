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

// checks if the KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest{}

// KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest struct for KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest
type KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest struct {
	DisplayName            string                                                                   `json:"displayName"`
	Subject                string                                                                   `json:"subject"`
	Message                string                                                                   `json:"message"`
	TemplateId             NullableInt32                                                            `json:"templateId,omitempty"`
	RegisteredEventHandler *KeyfactorWebKeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest `json:"registeredEventHandler,omitempty"`
	Recipients             []string                                                                 `json:"recipients,omitempty"`
	EventHandlerParameters []KeyfactorWebKeyfactorApiModelsEventHandlerEventHandlerParameterRequest `json:"eventHandlerParameters,omitempty"`
}

// NewKeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest instantiates a new KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest(displayName string, subject string, message string) *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest {
	this := KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest{}
	this.DisplayName = displayName
	this.Subject = subject
	this.Message = message
	return &this
}

// NewKeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequestWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequestWithDefaults() *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest {
	this := KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest{}
	return &this
}

// GetDisplayName returns the DisplayName field value
func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetSubject returns the Subject field value
func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest) SetSubject(v string) {
	o.Subject = v
}

// GetMessage returns the Message field value
func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest) SetMessage(v string) {
	o.Message = v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest) GetTemplateId() int32 {
	if o == nil || isNil(o.TemplateId.Get()) {
		var ret int32
		return ret
	}
	return *o.TemplateId.Get()
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest) GetTemplateIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TemplateId.Get(), o.TemplateId.IsSet()
}

// HasTemplateId returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest) HasTemplateId() bool {
	if o != nil && o.TemplateId.IsSet() {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given NullableInt32 and assigns it to the TemplateId field.
func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest) SetTemplateId(v int32) {
	o.TemplateId.Set(&v)
}

// SetTemplateIdNil sets the value for TemplateId to be an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest) SetTemplateIdNil() {
	o.TemplateId.Set(nil)
}

// UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest) UnsetTemplateId() {
	o.TemplateId.Unset()
}

// GetRegisteredEventHandler returns the RegisteredEventHandler field value if set, zero value otherwise.
func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest) GetRegisteredEventHandler() KeyfactorWebKeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest {
	if o == nil || isNil(o.RegisteredEventHandler) {
		var ret KeyfactorWebKeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest
		return ret
	}
	return *o.RegisteredEventHandler
}

// GetRegisteredEventHandlerOk returns a tuple with the RegisteredEventHandler field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest) GetRegisteredEventHandlerOk() (*KeyfactorWebKeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest, bool) {
	if o == nil || isNil(o.RegisteredEventHandler) {
		return nil, false
	}
	return o.RegisteredEventHandler, true
}

// HasRegisteredEventHandler returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest) HasRegisteredEventHandler() bool {
	if o != nil && !isNil(o.RegisteredEventHandler) {
		return true
	}

	return false
}

// SetRegisteredEventHandler gets a reference to the given KeyfactorWebKeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest and assigns it to the RegisteredEventHandler field.
func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest) SetRegisteredEventHandler(v KeyfactorWebKeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest) {
	o.RegisteredEventHandler = &v
}

// GetRecipients returns the Recipients field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest) GetRecipients() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest) GetRecipientsOk() ([]string, bool) {
	if o == nil || isNil(o.Recipients) {
		return nil, false
	}
	return o.Recipients, true
}

// HasRecipients returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest) HasRecipients() bool {
	if o != nil && isNil(o.Recipients) {
		return true
	}

	return false
}

// SetRecipients gets a reference to the given []string and assigns it to the Recipients field.
func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest) SetRecipients(v []string) {
	o.Recipients = v
}

// GetEventHandlerParameters returns the EventHandlerParameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest) GetEventHandlerParameters() []KeyfactorWebKeyfactorApiModelsEventHandlerEventHandlerParameterRequest {
	if o == nil {
		var ret []KeyfactorWebKeyfactorApiModelsEventHandlerEventHandlerParameterRequest
		return ret
	}
	return o.EventHandlerParameters
}

// GetEventHandlerParametersOk returns a tuple with the EventHandlerParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest) GetEventHandlerParametersOk() ([]KeyfactorWebKeyfactorApiModelsEventHandlerEventHandlerParameterRequest, bool) {
	if o == nil || isNil(o.EventHandlerParameters) {
		return nil, false
	}
	return o.EventHandlerParameters, true
}

// HasEventHandlerParameters returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest) HasEventHandlerParameters() bool {
	if o != nil && isNil(o.EventHandlerParameters) {
		return true
	}

	return false
}

// SetEventHandlerParameters gets a reference to the given []KeyfactorWebKeyfactorApiModelsEventHandlerEventHandlerParameterRequest and assigns it to the EventHandlerParameters field.
func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest) SetEventHandlerParameters(v []KeyfactorWebKeyfactorApiModelsEventHandlerEventHandlerParameterRequest) {
	o.EventHandlerParameters = v
}

func (o KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["displayName"] = o.DisplayName
	toSerialize["subject"] = o.Subject
	toSerialize["message"] = o.Message
	if o.TemplateId.IsSet() {
		toSerialize["templateId"] = o.TemplateId.Get()
	}
	if !isNil(o.RegisteredEventHandler) {
		toSerialize["registeredEventHandler"] = o.RegisteredEventHandler
	}
	if o.Recipients != nil {
		toSerialize["recipients"] = o.Recipients
	}
	if o.EventHandlerParameters != nil {
		toSerialize["eventHandlerParameters"] = o.EventHandlerParameters
	}
	return toSerialize, nil
}

type NullableKeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest struct {
	value *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest
	isSet bool
}

func (v NullableKeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest) Get() *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest {
	return v.value
}

func (v *NullableKeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest) Set(val *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest(val *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest) *NullableKeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest {
	return &NullableKeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest{value: val, isSet: true}
}

func (v NullableKeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
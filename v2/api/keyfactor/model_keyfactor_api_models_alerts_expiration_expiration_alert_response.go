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

// checks if the KeyfactorApiModelsAlertsExpirationExpirationAlertResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorApiModelsAlertsExpirationExpirationAlertResponse{}

// KeyfactorApiModelsAlertsExpirationExpirationAlertResponse struct for KeyfactorApiModelsAlertsExpirationExpirationAlertResponse
type KeyfactorApiModelsAlertsExpirationExpirationAlertResponse struct {
	CAName               *string        `json:"CAName,omitempty"`
	CARow                *int64         `json:"CARow,omitempty"`
	IssuedCN             NullableString `json:"IssuedCN,omitempty"`
	Expiry               *string        `json:"Expiry,omitempty"`
	Subject              *string        `json:"Subject,omitempty"`
	Message              *string        `json:"Message,omitempty"`
	Recipients           []string       `json:"Recipients,omitempty"`
	SendDate             *string        `json:"SendDate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KeyfactorApiModelsAlertsExpirationExpirationAlertResponse KeyfactorApiModelsAlertsExpirationExpirationAlertResponse

// NewKeyfactorApiModelsAlertsExpirationExpirationAlertResponse instantiates a new KeyfactorApiModelsAlertsExpirationExpirationAlertResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorApiModelsAlertsExpirationExpirationAlertResponse() *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse {
	this := KeyfactorApiModelsAlertsExpirationExpirationAlertResponse{}
	return &this
}

// NewKeyfactorApiModelsAlertsExpirationExpirationAlertResponseWithDefaults instantiates a new KeyfactorApiModelsAlertsExpirationExpirationAlertResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorApiModelsAlertsExpirationExpirationAlertResponseWithDefaults() *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse {
	this := KeyfactorApiModelsAlertsExpirationExpirationAlertResponse{}
	return &this
}

// GetCAName returns the CAName field value if set, zero value otherwise.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) GetCAName() string {
	if o == nil || isNil(o.CAName) {
		var ret string
		return ret
	}
	return *o.CAName
}

// GetCANameOk returns a tuple with the CAName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) GetCANameOk() (*string, bool) {
	if o == nil || isNil(o.CAName) {
		return nil, false
	}
	return o.CAName, true
}

// HasCAName returns a boolean if a field has been set.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) HasCAName() bool {
	if o != nil && !isNil(o.CAName) {
		return true
	}

	return false
}

// SetCAName gets a reference to the given string and assigns it to the CAName field.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) SetCAName(v string) {
	o.CAName = &v
}

// GetCARow returns the CARow field value if set, zero value otherwise.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) GetCARow() int64 {
	if o == nil || isNil(o.CARow) {
		var ret int64
		return ret
	}
	return *o.CARow
}

// GetCARowOk returns a tuple with the CARow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) GetCARowOk() (*int64, bool) {
	if o == nil || isNil(o.CARow) {
		return nil, false
	}
	return o.CARow, true
}

// HasCARow returns a boolean if a field has been set.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) HasCARow() bool {
	if o != nil && !isNil(o.CARow) {
		return true
	}

	return false
}

// SetCARow gets a reference to the given int64 and assigns it to the CARow field.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) SetCARow(v int64) {
	o.CARow = &v
}

// GetIssuedCN returns the IssuedCN field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) GetIssuedCN() string {
	if o == nil || isNil(o.IssuedCN.Get()) {
		var ret string
		return ret
	}
	return *o.IssuedCN.Get()
}

// GetIssuedCNOk returns a tuple with the IssuedCN field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) GetIssuedCNOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IssuedCN.Get(), o.IssuedCN.IsSet()
}

// HasIssuedCN returns a boolean if a field has been set.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) HasIssuedCN() bool {
	if o != nil && o.IssuedCN.IsSet() {
		return true
	}

	return false
}

// SetIssuedCN gets a reference to the given NullableString and assigns it to the IssuedCN field.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) SetIssuedCN(v string) {
	o.IssuedCN.Set(&v)
}

// SetIssuedCNNil sets the value for IssuedCN to be an explicit nil
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) SetIssuedCNNil() {
	o.IssuedCN.Set(nil)
}

// UnsetIssuedCN ensures that no value is present for IssuedCN, not even an explicit nil
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) UnsetIssuedCN() {
	o.IssuedCN.Unset()
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) GetExpiry() string {
	if o == nil || isNil(o.Expiry) {
		var ret string
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) GetExpiryOk() (*string, bool) {
	if o == nil || isNil(o.Expiry) {
		return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) HasExpiry() bool {
	if o != nil && !isNil(o.Expiry) {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given string and assigns it to the Expiry field.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) SetExpiry(v string) {
	o.Expiry = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) GetSubject() string {
	if o == nil || isNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) GetSubjectOk() (*string, bool) {
	if o == nil || isNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) HasSubject() bool {
	if o != nil && !isNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) SetSubject(v string) {
	o.Subject = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) SetMessage(v string) {
	o.Message = &v
}

// GetRecipients returns the Recipients field value if set, zero value otherwise.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) GetRecipients() []string {
	if o == nil || isNil(o.Recipients) {
		var ret []string
		return ret
	}
	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) GetRecipientsOk() ([]string, bool) {
	if o == nil || isNil(o.Recipients) {
		return nil, false
	}
	return o.Recipients, true
}

// HasRecipients returns a boolean if a field has been set.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) HasRecipients() bool {
	if o != nil && !isNil(o.Recipients) {
		return true
	}

	return false
}

// SetRecipients gets a reference to the given []string and assigns it to the Recipients field.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) SetRecipients(v []string) {
	o.Recipients = v
}

// GetSendDate returns the SendDate field value if set, zero value otherwise.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) GetSendDate() string {
	if o == nil || isNil(o.SendDate) {
		var ret string
		return ret
	}
	return *o.SendDate
}

// GetSendDateOk returns a tuple with the SendDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) GetSendDateOk() (*string, bool) {
	if o == nil || isNil(o.SendDate) {
		return nil, false
	}
	return o.SendDate, true
}

// HasSendDate returns a boolean if a field has been set.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) HasSendDate() bool {
	if o != nil && !isNil(o.SendDate) {
		return true
	}

	return false
}

// SetSendDate gets a reference to the given string and assigns it to the SendDate field.
func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) SetSendDate(v string) {
	o.SendDate = &v
}

func (o KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CAName) {
		toSerialize["CAName"] = o.CAName
	}
	if !isNil(o.CARow) {
		toSerialize["CARow"] = o.CARow
	}
	if o.IssuedCN.IsSet() {
		toSerialize["IssuedCN"] = o.IssuedCN.Get()
	}
	if !isNil(o.Expiry) {
		toSerialize["Expiry"] = o.Expiry
	}
	if !isNil(o.Subject) {
		toSerialize["Subject"] = o.Subject
	}
	if !isNil(o.Message) {
		toSerialize["Message"] = o.Message
	}
	if !isNil(o.Recipients) {
		toSerialize["Recipients"] = o.Recipients
	}
	if !isNil(o.SendDate) {
		toSerialize["SendDate"] = o.SendDate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) UnmarshalJSON(bytes []byte) (err error) {
	varKeyfactorApiModelsAlertsExpirationExpirationAlertResponse := _KeyfactorApiModelsAlertsExpirationExpirationAlertResponse{}

	if err = json.Unmarshal(bytes, &varKeyfactorApiModelsAlertsExpirationExpirationAlertResponse); err == nil {
		*o = KeyfactorApiModelsAlertsExpirationExpirationAlertResponse(varKeyfactorApiModelsAlertsExpirationExpirationAlertResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "CAName")
		delete(additionalProperties, "CARow")
		delete(additionalProperties, "IssuedCN")
		delete(additionalProperties, "Expiry")
		delete(additionalProperties, "Subject")
		delete(additionalProperties, "Message")
		delete(additionalProperties, "Recipients")
		delete(additionalProperties, "SendDate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKeyfactorApiModelsAlertsExpirationExpirationAlertResponse struct {
	value *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse
	isSet bool
}

func (v NullableKeyfactorApiModelsAlertsExpirationExpirationAlertResponse) Get() *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse {
	return v.value
}

func (v *NullableKeyfactorApiModelsAlertsExpirationExpirationAlertResponse) Set(val *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorApiModelsAlertsExpirationExpirationAlertResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorApiModelsAlertsExpirationExpirationAlertResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorApiModelsAlertsExpirationExpirationAlertResponse(val *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) *NullableKeyfactorApiModelsAlertsExpirationExpirationAlertResponse {
	return &NullableKeyfactorApiModelsAlertsExpirationExpirationAlertResponse{value: val, isSet: true}
}

func (v NullableKeyfactorApiModelsAlertsExpirationExpirationAlertResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorApiModelsAlertsExpirationExpirationAlertResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

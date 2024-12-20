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

// checks if the KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest{}

// KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest struct for KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest
type KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest struct {
	Id                   *int32                                             `json:"Id,omitempty"`
	Name                 string                                             `json:"Name"`
	EndpointType         string                                             `json:"EndpointType"`
	Location             string                                             `json:"Location"`
	Email                *KeyfactorApiModelsMonitoringEmailRequest          `json:"Email,omitempty"`
	Dashboard            KeyfactorApiModelsMonitoringDashboardRequest       `json:"Dashboard"`
	Schedule             *KeyfactorCommonSchedulingKeyfactorSchedule        `json:"Schedule,omitempty"`
	OCSPParameters       *KeyfactorApiModelsMonitoringOCSPParametersRequest `json:"OCSPParameters,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest

// NewKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest instantiates a new KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest(name string, endpointType string, location string, dashboard KeyfactorApiModelsMonitoringDashboardRequest) *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest {
	this := KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest{}
	this.Name = name
	this.EndpointType = endpointType
	this.Location = location
	this.Dashboard = dashboard
	return &this
}

// NewKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequestWithDefaults instantiates a new KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequestWithDefaults() *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest {
	this := KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) SetName(v string) {
	o.Name = v
}

// GetEndpointType returns the EndpointType field value
func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetEndpointType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndpointType
}

// GetEndpointTypeOk returns a tuple with the EndpointType field value
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetEndpointTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndpointType, true
}

// SetEndpointType sets field value
func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) SetEndpointType(v string) {
	o.EndpointType = v
}

// GetLocation returns the Location field value
func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) SetLocation(v string) {
	o.Location = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetEmail() KeyfactorApiModelsMonitoringEmailRequest {
	if o == nil || isNil(o.Email) {
		var ret KeyfactorApiModelsMonitoringEmailRequest
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetEmailOk() (*KeyfactorApiModelsMonitoringEmailRequest, bool) {
	if o == nil || isNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given KeyfactorApiModelsMonitoringEmailRequest and assigns it to the Email field.
func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) SetEmail(v KeyfactorApiModelsMonitoringEmailRequest) {
	o.Email = &v
}

// GetDashboard returns the Dashboard field value
func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetDashboard() KeyfactorApiModelsMonitoringDashboardRequest {
	if o == nil {
		var ret KeyfactorApiModelsMonitoringDashboardRequest
		return ret
	}

	return o.Dashboard
}

// GetDashboardOk returns a tuple with the Dashboard field value
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetDashboardOk() (*KeyfactorApiModelsMonitoringDashboardRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dashboard, true
}

// SetDashboard sets field value
func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) SetDashboard(v KeyfactorApiModelsMonitoringDashboardRequest) {
	o.Dashboard = v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetSchedule() KeyfactorCommonSchedulingKeyfactorSchedule {
	if o == nil || isNil(o.Schedule) {
		var ret KeyfactorCommonSchedulingKeyfactorSchedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool) {
	if o == nil || isNil(o.Schedule) {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) HasSchedule() bool {
	if o != nil && !isNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given KeyfactorCommonSchedulingKeyfactorSchedule and assigns it to the Schedule field.
func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) SetSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule) {
	o.Schedule = &v
}

// GetOCSPParameters returns the OCSPParameters field value if set, zero value otherwise.
func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetOCSPParameters() KeyfactorApiModelsMonitoringOCSPParametersRequest {
	if o == nil || isNil(o.OCSPParameters) {
		var ret KeyfactorApiModelsMonitoringOCSPParametersRequest
		return ret
	}
	return *o.OCSPParameters
}

// GetOCSPParametersOk returns a tuple with the OCSPParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetOCSPParametersOk() (*KeyfactorApiModelsMonitoringOCSPParametersRequest, bool) {
	if o == nil || isNil(o.OCSPParameters) {
		return nil, false
	}
	return o.OCSPParameters, true
}

// HasOCSPParameters returns a boolean if a field has been set.
func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) HasOCSPParameters() bool {
	if o != nil && !isNil(o.OCSPParameters) {
		return true
	}

	return false
}

// SetOCSPParameters gets a reference to the given KeyfactorApiModelsMonitoringOCSPParametersRequest and assigns it to the OCSPParameters field.
func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) SetOCSPParameters(v KeyfactorApiModelsMonitoringOCSPParametersRequest) {
	o.OCSPParameters = &v
}

func (o KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	toSerialize["Name"] = o.Name
	toSerialize["EndpointType"] = o.EndpointType
	toSerialize["Location"] = o.Location
	if !isNil(o.Email) {
		toSerialize["Email"] = o.Email
	}
	toSerialize["Dashboard"] = o.Dashboard
	if !isNil(o.Schedule) {
		toSerialize["Schedule"] = o.Schedule
	}
	if !isNil(o.OCSPParameters) {
		toSerialize["OCSPParameters"] = o.OCSPParameters
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) UnmarshalJSON(bytes []byte) (err error) {
	varKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest := _KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest{}

	if err = json.Unmarshal(bytes, &varKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest); err == nil {
		*o = KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest(varKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Id")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "EndpointType")
		delete(additionalProperties, "Location")
		delete(additionalProperties, "Email")
		delete(additionalProperties, "Dashboard")
		delete(additionalProperties, "Schedule")
		delete(additionalProperties, "OCSPParameters")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest struct {
	value *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest
	isSet bool
}

func (v NullableKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) Get() *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest {
	return v.value
}

func (v *NullableKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) Set(val *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest(val *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) *NullableKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest {
	return &NullableKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest{value: val, isSet: true}
}

func (v NullableKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

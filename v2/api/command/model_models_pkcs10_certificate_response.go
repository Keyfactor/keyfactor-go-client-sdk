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

Keyfactor-v1

This reference serves to document REST-based methods to manage and integrate with Keyfactor. In addition, an embedded interface allows for the execution of calls against the current Keyfactor API instance.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package command

import (
	"encoding/json"
)

// checks if the ModelsPkcs10CertificateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsPkcs10CertificateResponse{}

// ModelsPkcs10CertificateResponse struct for ModelsPkcs10CertificateResponse
type ModelsPkcs10CertificateResponse struct {
	SerialNumber *string        `json:"SerialNumber,omitempty"`
	IssuerDN     NullableString `json:"IssuerDN,omitempty"`
	Thumbprint   *string        `json:"Thumbprint,omitempty"`
	// The integer ID of the certificate in the keyfactor database, if Issued
	KeyfactorID  *int32   `json:"KeyfactorID,omitempty"`
	Certificates []string `json:"Certificates,omitempty"`
	// The integer id of the certificate request in the Keyfactor database, if one exists.
	KeyfactorRequestId   *int32            `json:"KeyfactorRequestId,omitempty"`
	RequestDisposition   *string           `json:"RequestDisposition,omitempty"`
	DispositionMessage   *string           `json:"DispositionMessage,omitempty"`
	EnrollmentContext    map[string]string `json:"EnrollmentContext,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelsPkcs10CertificateResponse ModelsPkcs10CertificateResponse

// NewModelsPkcs10CertificateResponse instantiates a new ModelsPkcs10CertificateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsPkcs10CertificateResponse() *ModelsPkcs10CertificateResponse {
	this := ModelsPkcs10CertificateResponse{}
	return &this
}

// NewModelsPkcs10CertificateResponseWithDefaults instantiates a new ModelsPkcs10CertificateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsPkcs10CertificateResponseWithDefaults() *ModelsPkcs10CertificateResponse {
	this := ModelsPkcs10CertificateResponse{}
	return &this
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *ModelsPkcs10CertificateResponse) GetSerialNumber() string {
	if o == nil || isNil(o.SerialNumber) {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsPkcs10CertificateResponse) GetSerialNumberOk() (*string, bool) {
	if o == nil || isNil(o.SerialNumber) {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *ModelsPkcs10CertificateResponse) HasSerialNumber() bool {
	if o != nil && !isNil(o.SerialNumber) {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *ModelsPkcs10CertificateResponse) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetIssuerDN returns the IssuerDN field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelsPkcs10CertificateResponse) GetIssuerDN() string {
	if o == nil || isNil(o.IssuerDN.Get()) {
		var ret string
		return ret
	}
	return *o.IssuerDN.Get()
}

// GetIssuerDNOk returns a tuple with the IssuerDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelsPkcs10CertificateResponse) GetIssuerDNOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IssuerDN.Get(), o.IssuerDN.IsSet()
}

// HasIssuerDN returns a boolean if a field has been set.
func (o *ModelsPkcs10CertificateResponse) HasIssuerDN() bool {
	if o != nil && o.IssuerDN.IsSet() {
		return true
	}

	return false
}

// SetIssuerDN gets a reference to the given NullableString and assigns it to the IssuerDN field.
func (o *ModelsPkcs10CertificateResponse) SetIssuerDN(v string) {
	o.IssuerDN.Set(&v)
}

// SetIssuerDNNil sets the value for IssuerDN to be an explicit nil
func (o *ModelsPkcs10CertificateResponse) SetIssuerDNNil() {
	o.IssuerDN.Set(nil)
}

// UnsetIssuerDN ensures that no value is present for IssuerDN, not even an explicit nil
func (o *ModelsPkcs10CertificateResponse) UnsetIssuerDN() {
	o.IssuerDN.Unset()
}

// GetThumbprint returns the Thumbprint field value if set, zero value otherwise.
func (o *ModelsPkcs10CertificateResponse) GetThumbprint() string {
	if o == nil || isNil(o.Thumbprint) {
		var ret string
		return ret
	}
	return *o.Thumbprint
}

// GetThumbprintOk returns a tuple with the Thumbprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsPkcs10CertificateResponse) GetThumbprintOk() (*string, bool) {
	if o == nil || isNil(o.Thumbprint) {
		return nil, false
	}
	return o.Thumbprint, true
}

// HasThumbprint returns a boolean if a field has been set.
func (o *ModelsPkcs10CertificateResponse) HasThumbprint() bool {
	if o != nil && !isNil(o.Thumbprint) {
		return true
	}

	return false
}

// SetThumbprint gets a reference to the given string and assigns it to the Thumbprint field.
func (o *ModelsPkcs10CertificateResponse) SetThumbprint(v string) {
	o.Thumbprint = &v
}

// GetKeyfactorID returns the KeyfactorID field value if set, zero value otherwise.
func (o *ModelsPkcs10CertificateResponse) GetKeyfactorID() int32 {
	if o == nil || isNil(o.KeyfactorID) {
		var ret int32
		return ret
	}
	return *o.KeyfactorID
}

// GetKeyfactorIDOk returns a tuple with the KeyfactorID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsPkcs10CertificateResponse) GetKeyfactorIDOk() (*int32, bool) {
	if o == nil || isNil(o.KeyfactorID) {
		return nil, false
	}
	return o.KeyfactorID, true
}

// HasKeyfactorID returns a boolean if a field has been set.
func (o *ModelsPkcs10CertificateResponse) HasKeyfactorID() bool {
	if o != nil && !isNil(o.KeyfactorID) {
		return true
	}

	return false
}

// SetKeyfactorID gets a reference to the given int32 and assigns it to the KeyfactorID field.
func (o *ModelsPkcs10CertificateResponse) SetKeyfactorID(v int32) {
	o.KeyfactorID = &v
}

// GetCertificates returns the Certificates field value if set, zero value otherwise.
func (o *ModelsPkcs10CertificateResponse) GetCertificates() []string {
	if o == nil || isNil(o.Certificates) {
		var ret []string
		return ret
	}
	return o.Certificates
}

// GetCertificatesOk returns a tuple with the Certificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsPkcs10CertificateResponse) GetCertificatesOk() ([]string, bool) {
	if o == nil || isNil(o.Certificates) {
		return nil, false
	}
	return o.Certificates, true
}

// HasCertificates returns a boolean if a field has been set.
func (o *ModelsPkcs10CertificateResponse) HasCertificates() bool {
	if o != nil && !isNil(o.Certificates) {
		return true
	}

	return false
}

// SetCertificates gets a reference to the given []string and assigns it to the Certificates field.
func (o *ModelsPkcs10CertificateResponse) SetCertificates(v []string) {
	o.Certificates = v
}

// GetKeyfactorRequestId returns the KeyfactorRequestId field value if set, zero value otherwise.
func (o *ModelsPkcs10CertificateResponse) GetKeyfactorRequestId() int32 {
	if o == nil || isNil(o.KeyfactorRequestId) {
		var ret int32
		return ret
	}
	return *o.KeyfactorRequestId
}

// GetKeyfactorRequestIdOk returns a tuple with the KeyfactorRequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsPkcs10CertificateResponse) GetKeyfactorRequestIdOk() (*int32, bool) {
	if o == nil || isNil(o.KeyfactorRequestId) {
		return nil, false
	}
	return o.KeyfactorRequestId, true
}

// HasKeyfactorRequestId returns a boolean if a field has been set.
func (o *ModelsPkcs10CertificateResponse) HasKeyfactorRequestId() bool {
	if o != nil && !isNil(o.KeyfactorRequestId) {
		return true
	}

	return false
}

// SetKeyfactorRequestId gets a reference to the given int32 and assigns it to the KeyfactorRequestId field.
func (o *ModelsPkcs10CertificateResponse) SetKeyfactorRequestId(v int32) {
	o.KeyfactorRequestId = &v
}

// GetRequestDisposition returns the RequestDisposition field value if set, zero value otherwise.
func (o *ModelsPkcs10CertificateResponse) GetRequestDisposition() string {
	if o == nil || isNil(o.RequestDisposition) {
		var ret string
		return ret
	}
	return *o.RequestDisposition
}

// GetRequestDispositionOk returns a tuple with the RequestDisposition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsPkcs10CertificateResponse) GetRequestDispositionOk() (*string, bool) {
	if o == nil || isNil(o.RequestDisposition) {
		return nil, false
	}
	return o.RequestDisposition, true
}

// HasRequestDisposition returns a boolean if a field has been set.
func (o *ModelsPkcs10CertificateResponse) HasRequestDisposition() bool {
	if o != nil && !isNil(o.RequestDisposition) {
		return true
	}

	return false
}

// SetRequestDisposition gets a reference to the given string and assigns it to the RequestDisposition field.
func (o *ModelsPkcs10CertificateResponse) SetRequestDisposition(v string) {
	o.RequestDisposition = &v
}

// GetDispositionMessage returns the DispositionMessage field value if set, zero value otherwise.
func (o *ModelsPkcs10CertificateResponse) GetDispositionMessage() string {
	if o == nil || isNil(o.DispositionMessage) {
		var ret string
		return ret
	}
	return *o.DispositionMessage
}

// GetDispositionMessageOk returns a tuple with the DispositionMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsPkcs10CertificateResponse) GetDispositionMessageOk() (*string, bool) {
	if o == nil || isNil(o.DispositionMessage) {
		return nil, false
	}
	return o.DispositionMessage, true
}

// HasDispositionMessage returns a boolean if a field has been set.
func (o *ModelsPkcs10CertificateResponse) HasDispositionMessage() bool {
	if o != nil && !isNil(o.DispositionMessage) {
		return true
	}

	return false
}

// SetDispositionMessage gets a reference to the given string and assigns it to the DispositionMessage field.
func (o *ModelsPkcs10CertificateResponse) SetDispositionMessage(v string) {
	o.DispositionMessage = &v
}

// GetEnrollmentContext returns the EnrollmentContext field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelsPkcs10CertificateResponse) GetEnrollmentContext() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.EnrollmentContext
}

// GetEnrollmentContextOk returns a tuple with the EnrollmentContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelsPkcs10CertificateResponse) GetEnrollmentContextOk() (*map[string]string, bool) {
	if o == nil || isNil(o.EnrollmentContext) {
		return nil, false
	}
	return &o.EnrollmentContext, true
}

// HasEnrollmentContext returns a boolean if a field has been set.
func (o *ModelsPkcs10CertificateResponse) HasEnrollmentContext() bool {
	if o != nil && isNil(o.EnrollmentContext) {
		return true
	}

	return false
}

// SetEnrollmentContext gets a reference to the given map[string]string and assigns it to the EnrollmentContext field.
func (o *ModelsPkcs10CertificateResponse) SetEnrollmentContext(v map[string]string) {
	o.EnrollmentContext = v
}

func (o ModelsPkcs10CertificateResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsPkcs10CertificateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SerialNumber) {
		toSerialize["SerialNumber"] = o.SerialNumber
	}
	if o.IssuerDN.IsSet() {
		toSerialize["IssuerDN"] = o.IssuerDN.Get()
	}
	if !isNil(o.Thumbprint) {
		toSerialize["Thumbprint"] = o.Thumbprint
	}
	if !isNil(o.KeyfactorID) {
		toSerialize["KeyfactorID"] = o.KeyfactorID
	}
	if !isNil(o.Certificates) {
		toSerialize["Certificates"] = o.Certificates
	}
	if !isNil(o.KeyfactorRequestId) {
		toSerialize["KeyfactorRequestId"] = o.KeyfactorRequestId
	}
	if !isNil(o.RequestDisposition) {
		toSerialize["RequestDisposition"] = o.RequestDisposition
	}
	if !isNil(o.DispositionMessage) {
		toSerialize["DispositionMessage"] = o.DispositionMessage
	}
	if o.EnrollmentContext != nil {
		toSerialize["EnrollmentContext"] = o.EnrollmentContext
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModelsPkcs10CertificateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varModelsPkcs10CertificateResponse := _ModelsPkcs10CertificateResponse{}

	if err = json.Unmarshal(bytes, &varModelsPkcs10CertificateResponse); err == nil {
		*o = ModelsPkcs10CertificateResponse(varModelsPkcs10CertificateResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "SerialNumber")
		delete(additionalProperties, "IssuerDN")
		delete(additionalProperties, "Thumbprint")
		delete(additionalProperties, "KeyfactorID")
		delete(additionalProperties, "Certificates")
		delete(additionalProperties, "KeyfactorRequestId")
		delete(additionalProperties, "RequestDisposition")
		delete(additionalProperties, "DispositionMessage")
		delete(additionalProperties, "EnrollmentContext")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelsPkcs10CertificateResponse struct {
	value *ModelsPkcs10CertificateResponse
	isSet bool
}

func (v NullableModelsPkcs10CertificateResponse) Get() *ModelsPkcs10CertificateResponse {
	return v.value
}

func (v *NullableModelsPkcs10CertificateResponse) Set(val *ModelsPkcs10CertificateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsPkcs10CertificateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsPkcs10CertificateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsPkcs10CertificateResponse(val *ModelsPkcs10CertificateResponse) *NullableModelsPkcs10CertificateResponse {
	return &NullableModelsPkcs10CertificateResponse{value: val, isSet: true}
}

func (v NullableModelsPkcs10CertificateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsPkcs10CertificateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
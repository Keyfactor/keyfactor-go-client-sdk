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

// checks if the KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponse{}

// KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponse struct for KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponse
type KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponse struct {
	QueryId           *int32                                                                              `json:"queryId,omitempty"`
	AccessControlList []KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateQueryAccessControl `json:"accessControlList,omitempty"`
	AssignableRoles   []KeyfactorWebKeyfactorApiModelsCertificateCollectionsAssignableQueryRole           `json:"assignableRoles,omitempty"`
}

// NewKeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponse instantiates a new KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponse() *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponse {
	this := KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponse{}
	return &this
}

// NewKeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponseWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponseWithDefaults() *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponse {
	this := KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponse{}
	return &this
}

// GetQueryId returns the QueryId field value if set, zero value otherwise.
func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponse) GetQueryId() int32 {
	if o == nil || isNil(o.QueryId) {
		var ret int32
		return ret
	}
	return *o.QueryId
}

// GetQueryIdOk returns a tuple with the QueryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponse) GetQueryIdOk() (*int32, bool) {
	if o == nil || isNil(o.QueryId) {
		return nil, false
	}
	return o.QueryId, true
}

// HasQueryId returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponse) HasQueryId() bool {
	if o != nil && !isNil(o.QueryId) {
		return true
	}

	return false
}

// SetQueryId gets a reference to the given int32 and assigns it to the QueryId field.
func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponse) SetQueryId(v int32) {
	o.QueryId = &v
}

// GetAccessControlList returns the AccessControlList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponse) GetAccessControlList() []KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateQueryAccessControl {
	if o == nil {
		var ret []KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateQueryAccessControl
		return ret
	}
	return o.AccessControlList
}

// GetAccessControlListOk returns a tuple with the AccessControlList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponse) GetAccessControlListOk() ([]KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateQueryAccessControl, bool) {
	if o == nil || isNil(o.AccessControlList) {
		return nil, false
	}
	return o.AccessControlList, true
}

// HasAccessControlList returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponse) HasAccessControlList() bool {
	if o != nil && isNil(o.AccessControlList) {
		return true
	}

	return false
}

// SetAccessControlList gets a reference to the given []KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateQueryAccessControl and assigns it to the AccessControlList field.
func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponse) SetAccessControlList(v []KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateQueryAccessControl) {
	o.AccessControlList = v
}

// GetAssignableRoles returns the AssignableRoles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponse) GetAssignableRoles() []KeyfactorWebKeyfactorApiModelsCertificateCollectionsAssignableQueryRole {
	if o == nil {
		var ret []KeyfactorWebKeyfactorApiModelsCertificateCollectionsAssignableQueryRole
		return ret
	}
	return o.AssignableRoles
}

// GetAssignableRolesOk returns a tuple with the AssignableRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponse) GetAssignableRolesOk() ([]KeyfactorWebKeyfactorApiModelsCertificateCollectionsAssignableQueryRole, bool) {
	if o == nil || isNil(o.AssignableRoles) {
		return nil, false
	}
	return o.AssignableRoles, true
}

// HasAssignableRoles returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponse) HasAssignableRoles() bool {
	if o != nil && isNil(o.AssignableRoles) {
		return true
	}

	return false
}

// SetAssignableRoles gets a reference to the given []KeyfactorWebKeyfactorApiModelsCertificateCollectionsAssignableQueryRole and assigns it to the AssignableRoles field.
func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponse) SetAssignableRoles(v []KeyfactorWebKeyfactorApiModelsCertificateCollectionsAssignableQueryRole) {
	o.AssignableRoles = v
}

func (o KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.QueryId) {
		toSerialize["queryId"] = o.QueryId
	}
	if o.AccessControlList != nil {
		toSerialize["accessControlList"] = o.AccessControlList
	}
	if o.AssignableRoles != nil {
		toSerialize["assignableRoles"] = o.AssignableRoles
	}
	return toSerialize, nil
}

type NullableKeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponse struct {
	value *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponse
	isSet bool
}

func (v NullableKeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponse) Get() *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponse {
	return v.value
}

func (v *NullableKeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponse) Set(val *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponse(val *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponse) *NullableKeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponse {
	return &NullableKeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponse{value: val, isSet: true}
}

func (v NullableKeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
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

// checks if the KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest{}

// KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest struct for KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest
type KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest struct {
	// The Id of the security role to update
	Id int32 `json:"id"`
	// The name of the security role to update
	Name string `json:"name"`
	// The description to be used on the updated security role
	Description string `json:"description"`
	// Whether or not the security role should be enabled
	Enabled *bool `json:"enabled,omitempty"`
	// Whether or not the security role should be private
	Private *bool `json:"private,omitempty"`
	// The permissions to include in the role. These must be supplied in the format \"Area:Permission\"
	Permissions []string `json:"permissions,omitempty"`
	// The Id of the permission set the role belongs to.
	PermissionSetId *string `json:"permissionSetId,omitempty"`
	// The Keyfactor identities to assign to the updated role
	Identities []ModelsSecurityIdentitiesSecurityIdentityIdentifier `json:"identities,omitempty"`
}

// NewKeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest instantiates a new KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest(id int32, name string, description string) *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest {
	this := KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest{}
	this.Id = id
	this.Name = name
	this.Description = description
	return &this
}

// NewKeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequestWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequestWithDefaults() *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest {
	this := KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest{}
	return &this
}

// GetId returns the Id field value
func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) SetDescription(v string) {
	o.Description = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetPrivate returns the Private field value if set, zero value otherwise.
func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetPrivate() bool {
	if o == nil || isNil(o.Private) {
		var ret bool
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetPrivateOk() (*bool, bool) {
	if o == nil || isNil(o.Private) {
		return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) HasPrivate() bool {
	if o != nil && !isNil(o.Private) {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given bool and assigns it to the Private field.
func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) SetPrivate(v bool) {
	o.Private = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetPermissions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetPermissionsOk() ([]string, bool) {
	if o == nil || isNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) HasPermissions() bool {
	if o != nil && isNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []string and assigns it to the Permissions field.
func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) SetPermissions(v []string) {
	o.Permissions = v
}

// GetPermissionSetId returns the PermissionSetId field value if set, zero value otherwise.
func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetPermissionSetId() string {
	if o == nil || isNil(o.PermissionSetId) {
		var ret string
		return ret
	}
	return *o.PermissionSetId
}

// GetPermissionSetIdOk returns a tuple with the PermissionSetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetPermissionSetIdOk() (*string, bool) {
	if o == nil || isNil(o.PermissionSetId) {
		return nil, false
	}
	return o.PermissionSetId, true
}

// HasPermissionSetId returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) HasPermissionSetId() bool {
	if o != nil && !isNil(o.PermissionSetId) {
		return true
	}

	return false
}

// SetPermissionSetId gets a reference to the given string and assigns it to the PermissionSetId field.
func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) SetPermissionSetId(v string) {
	o.PermissionSetId = &v
}

// GetIdentities returns the Identities field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetIdentities() []ModelsSecurityIdentitiesSecurityIdentityIdentifier {
	if o == nil {
		var ret []ModelsSecurityIdentitiesSecurityIdentityIdentifier
		return ret
	}
	return o.Identities
}

// GetIdentitiesOk returns a tuple with the Identities field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetIdentitiesOk() ([]ModelsSecurityIdentitiesSecurityIdentityIdentifier, bool) {
	if o == nil || isNil(o.Identities) {
		return nil, false
	}
	return o.Identities, true
}

// HasIdentities returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) HasIdentities() bool {
	if o != nil && isNil(o.Identities) {
		return true
	}

	return false
}

// SetIdentities gets a reference to the given []ModelsSecurityIdentitiesSecurityIdentityIdentifier and assigns it to the Identities field.
func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) SetIdentities(v []ModelsSecurityIdentitiesSecurityIdentityIdentifier) {
	o.Identities = v
}

func (o KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Private) {
		toSerialize["private"] = o.Private
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	if !isNil(o.PermissionSetId) {
		toSerialize["permissionSetId"] = o.PermissionSetId
	}
	if o.Identities != nil {
		toSerialize["identities"] = o.Identities
	}
	return toSerialize, nil
}

type NullableKeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest struct {
	value *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest
	isSet bool
}

func (v NullableKeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) Get() *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest {
	return v.value
}

func (v *NullableKeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) Set(val *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest(val *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) *NullableKeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest {
	return &NullableKeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest{value: val, isSet: true}
}

func (v NullableKeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
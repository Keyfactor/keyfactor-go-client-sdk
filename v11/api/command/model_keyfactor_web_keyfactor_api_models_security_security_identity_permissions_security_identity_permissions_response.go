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

// checks if the KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse{}

// KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse struct for KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse
type KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse struct {
	Claim                       *KeyfactorWebKeyfactorApiModelsSecurityRoleClaimDefinitionsRoleClaimDefinitionResponse         `json:"claim,omitempty"`
	Identity                    NullableString                                                                                 `json:"identity,omitempty"`
	SecuredAreaPermissions      []KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse `json:"securedAreaPermissions,omitempty"`
	CollectionPermissions       []KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse `json:"collectionPermissions,omitempty"`
	ContainerPermissions        []KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse `json:"containerPermissions,omitempty"`
	PamProviderPermissions      []KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse `json:"pamProviderPermissions,omitempty"`
	IdentityProviderPermissions []KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse `json:"identityProviderPermissions,omitempty"`
	PamPermissions              []KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse `json:"pamPermissions,omitempty"`
}

// NewKeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse instantiates a new KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse() *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse {
	this := KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse{}
	return &this
}

// NewKeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponseWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponseWithDefaults() *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse {
	this := KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse{}
	return &this
}

// GetClaim returns the Claim field value if set, zero value otherwise.
func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetClaim() KeyfactorWebKeyfactorApiModelsSecurityRoleClaimDefinitionsRoleClaimDefinitionResponse {
	if o == nil || isNil(o.Claim) {
		var ret KeyfactorWebKeyfactorApiModelsSecurityRoleClaimDefinitionsRoleClaimDefinitionResponse
		return ret
	}
	return *o.Claim
}

// GetClaimOk returns a tuple with the Claim field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetClaimOk() (*KeyfactorWebKeyfactorApiModelsSecurityRoleClaimDefinitionsRoleClaimDefinitionResponse, bool) {
	if o == nil || isNil(o.Claim) {
		return nil, false
	}
	return o.Claim, true
}

// HasClaim returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) HasClaim() bool {
	if o != nil && !isNil(o.Claim) {
		return true
	}

	return false
}

// SetClaim gets a reference to the given KeyfactorWebKeyfactorApiModelsSecurityRoleClaimDefinitionsRoleClaimDefinitionResponse and assigns it to the Claim field.
func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) SetClaim(v KeyfactorWebKeyfactorApiModelsSecurityRoleClaimDefinitionsRoleClaimDefinitionResponse) {
	o.Claim = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetIdentity() string {
	if o == nil || isNil(o.Identity.Get()) {
		var ret string
		return ret
	}
	return *o.Identity.Get()
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetIdentityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Identity.Get(), o.Identity.IsSet()
}

// HasIdentity returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) HasIdentity() bool {
	if o != nil && o.Identity.IsSet() {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given NullableString and assigns it to the Identity field.
func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) SetIdentity(v string) {
	o.Identity.Set(&v)
}

// SetIdentityNil sets the value for Identity to be an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) SetIdentityNil() {
	o.Identity.Set(nil)
}

// UnsetIdentity ensures that no value is present for Identity, not even an explicit nil
func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) UnsetIdentity() {
	o.Identity.Unset()
}

// GetSecuredAreaPermissions returns the SecuredAreaPermissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetSecuredAreaPermissions() []KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse {
	if o == nil {
		var ret []KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse
		return ret
	}
	return o.SecuredAreaPermissions
}

// GetSecuredAreaPermissionsOk returns a tuple with the SecuredAreaPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetSecuredAreaPermissionsOk() ([]KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse, bool) {
	if o == nil || isNil(o.SecuredAreaPermissions) {
		return nil, false
	}
	return o.SecuredAreaPermissions, true
}

// HasSecuredAreaPermissions returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) HasSecuredAreaPermissions() bool {
	if o != nil && isNil(o.SecuredAreaPermissions) {
		return true
	}

	return false
}

// SetSecuredAreaPermissions gets a reference to the given []KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse and assigns it to the SecuredAreaPermissions field.
func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) SetSecuredAreaPermissions(v []KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse) {
	o.SecuredAreaPermissions = v
}

// GetCollectionPermissions returns the CollectionPermissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetCollectionPermissions() []KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse {
	if o == nil {
		var ret []KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse
		return ret
	}
	return o.CollectionPermissions
}

// GetCollectionPermissionsOk returns a tuple with the CollectionPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetCollectionPermissionsOk() ([]KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse, bool) {
	if o == nil || isNil(o.CollectionPermissions) {
		return nil, false
	}
	return o.CollectionPermissions, true
}

// HasCollectionPermissions returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) HasCollectionPermissions() bool {
	if o != nil && isNil(o.CollectionPermissions) {
		return true
	}

	return false
}

// SetCollectionPermissions gets a reference to the given []KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse and assigns it to the CollectionPermissions field.
func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) SetCollectionPermissions(v []KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse) {
	o.CollectionPermissions = v
}

// GetContainerPermissions returns the ContainerPermissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetContainerPermissions() []KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse {
	if o == nil {
		var ret []KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse
		return ret
	}
	return o.ContainerPermissions
}

// GetContainerPermissionsOk returns a tuple with the ContainerPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetContainerPermissionsOk() ([]KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse, bool) {
	if o == nil || isNil(o.ContainerPermissions) {
		return nil, false
	}
	return o.ContainerPermissions, true
}

// HasContainerPermissions returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) HasContainerPermissions() bool {
	if o != nil && isNil(o.ContainerPermissions) {
		return true
	}

	return false
}

// SetContainerPermissions gets a reference to the given []KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse and assigns it to the ContainerPermissions field.
func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) SetContainerPermissions(v []KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse) {
	o.ContainerPermissions = v
}

// GetPamProviderPermissions returns the PamProviderPermissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetPamProviderPermissions() []KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse {
	if o == nil {
		var ret []KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse
		return ret
	}
	return o.PamProviderPermissions
}

// GetPamProviderPermissionsOk returns a tuple with the PamProviderPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetPamProviderPermissionsOk() ([]KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse, bool) {
	if o == nil || isNil(o.PamProviderPermissions) {
		return nil, false
	}
	return o.PamProviderPermissions, true
}

// HasPamProviderPermissions returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) HasPamProviderPermissions() bool {
	if o != nil && isNil(o.PamProviderPermissions) {
		return true
	}

	return false
}

// SetPamProviderPermissions gets a reference to the given []KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse and assigns it to the PamProviderPermissions field.
func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) SetPamProviderPermissions(v []KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse) {
	o.PamProviderPermissions = v
}

// GetIdentityProviderPermissions returns the IdentityProviderPermissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetIdentityProviderPermissions() []KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse {
	if o == nil {
		var ret []KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse
		return ret
	}
	return o.IdentityProviderPermissions
}

// GetIdentityProviderPermissionsOk returns a tuple with the IdentityProviderPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetIdentityProviderPermissionsOk() ([]KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse, bool) {
	if o == nil || isNil(o.IdentityProviderPermissions) {
		return nil, false
	}
	return o.IdentityProviderPermissions, true
}

// HasIdentityProviderPermissions returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) HasIdentityProviderPermissions() bool {
	if o != nil && isNil(o.IdentityProviderPermissions) {
		return true
	}

	return false
}

// SetIdentityProviderPermissions gets a reference to the given []KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse and assigns it to the IdentityProviderPermissions field.
func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) SetIdentityProviderPermissions(v []KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse) {
	o.IdentityProviderPermissions = v
}

// GetPamPermissions returns the PamPermissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetPamPermissions() []KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse {
	if o == nil {
		var ret []KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse
		return ret
	}
	return o.PamPermissions
}

// GetPamPermissionsOk returns a tuple with the PamPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) GetPamPermissionsOk() ([]KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse, bool) {
	if o == nil || isNil(o.PamPermissions) {
		return nil, false
	}
	return o.PamPermissions, true
}

// HasPamPermissions returns a boolean if a field has been set.
func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) HasPamPermissions() bool {
	if o != nil && isNil(o.PamPermissions) {
		return true
	}

	return false
}

// SetPamPermissions gets a reference to the given []KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse and assigns it to the PamPermissions field.
func (o *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) SetPamPermissions(v []KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsPermissionRolesPairResponse) {
	o.PamPermissions = v
}

func (o KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Claim) {
		toSerialize["claim"] = o.Claim
	}
	if o.Identity.IsSet() {
		toSerialize["identity"] = o.Identity.Get()
	}
	if o.SecuredAreaPermissions != nil {
		toSerialize["securedAreaPermissions"] = o.SecuredAreaPermissions
	}
	if o.CollectionPermissions != nil {
		toSerialize["collectionPermissions"] = o.CollectionPermissions
	}
	if o.ContainerPermissions != nil {
		toSerialize["containerPermissions"] = o.ContainerPermissions
	}
	if o.PamProviderPermissions != nil {
		toSerialize["pamProviderPermissions"] = o.PamProviderPermissions
	}
	if o.IdentityProviderPermissions != nil {
		toSerialize["identityProviderPermissions"] = o.IdentityProviderPermissions
	}
	if o.PamPermissions != nil {
		toSerialize["pamPermissions"] = o.PamPermissions
	}
	return toSerialize, nil
}

type NullableKeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse struct {
	value *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse
	isSet bool
}

func (v NullableKeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) Get() *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse {
	return v.value
}

func (v *NullableKeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) Set(val *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse(val *KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) *NullableKeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse {
	return &NullableKeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse{value: val, isSet: true}
}

func (v NullableKeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
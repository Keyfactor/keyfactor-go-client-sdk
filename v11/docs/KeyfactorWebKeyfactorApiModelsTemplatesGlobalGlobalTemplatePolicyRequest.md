# KeyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplatePolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the global template policy. | [optional] 
**AllowKeyReuse** | **bool** | Whether or not keys can be reused. | 
**AllowWildcards** | **bool** | Whether or not wildcards can be used. | 
**RfcEnforcement** | **bool** | Whether or not RFC 2818 compliance should be enforced. | 
**KeyInfo** | [**ModelsTemplatesAlgorithmsKeyInfo**](ModelsTemplatesAlgorithmsKeyInfo.md) |  | 

## Methods

### NewKeyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplatePolicyRequest

`func NewKeyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplatePolicyRequest(allowKeyReuse bool, allowWildcards bool, rfcEnforcement bool, keyInfo ModelsTemplatesAlgorithmsKeyInfo, ) *KeyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplatePolicyRequest`

NewKeyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplatePolicyRequest instantiates a new KeyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplatePolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplatePolicyRequestWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplatePolicyRequestWithDefaults() *KeyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplatePolicyRequest`

NewKeyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplatePolicyRequestWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplatePolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplatePolicyRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplatePolicyRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplatePolicyRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplatePolicyRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAllowKeyReuse

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplatePolicyRequest) GetAllowKeyReuse() bool`

GetAllowKeyReuse returns the AllowKeyReuse field if non-nil, zero value otherwise.

### GetAllowKeyReuseOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplatePolicyRequest) GetAllowKeyReuseOk() (*bool, bool)`

GetAllowKeyReuseOk returns a tuple with the AllowKeyReuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowKeyReuse

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplatePolicyRequest) SetAllowKeyReuse(v bool)`

SetAllowKeyReuse sets AllowKeyReuse field to given value.


### GetAllowWildcards

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplatePolicyRequest) GetAllowWildcards() bool`

GetAllowWildcards returns the AllowWildcards field if non-nil, zero value otherwise.

### GetAllowWildcardsOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplatePolicyRequest) GetAllowWildcardsOk() (*bool, bool)`

GetAllowWildcardsOk returns a tuple with the AllowWildcards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowWildcards

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplatePolicyRequest) SetAllowWildcards(v bool)`

SetAllowWildcards sets AllowWildcards field to given value.


### GetRfcEnforcement

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplatePolicyRequest) GetRfcEnforcement() bool`

GetRfcEnforcement returns the RfcEnforcement field if non-nil, zero value otherwise.

### GetRfcEnforcementOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplatePolicyRequest) GetRfcEnforcementOk() (*bool, bool)`

GetRfcEnforcementOk returns a tuple with the RfcEnforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfcEnforcement

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplatePolicyRequest) SetRfcEnforcement(v bool)`

SetRfcEnforcement sets RfcEnforcement field to given value.


### GetKeyInfo

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplatePolicyRequest) GetKeyInfo() ModelsTemplatesAlgorithmsKeyInfo`

GetKeyInfo returns the KeyInfo field if non-nil, zero value otherwise.

### GetKeyInfoOk

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplatePolicyRequest) GetKeyInfoOk() (*ModelsTemplatesAlgorithmsKeyInfo, bool)`

GetKeyInfoOk returns a tuple with the KeyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyInfo

`func (o *KeyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplatePolicyRequest) SetKeyInfo(v ModelsTemplatesAlgorithmsKeyInfo)`

SetKeyInfo sets KeyInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



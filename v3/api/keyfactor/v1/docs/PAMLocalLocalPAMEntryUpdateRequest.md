# PAMLocalLocalPAMEntryUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecretName** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**SecretValue** | **string** |  | 

## Methods

### NewPAMLocalLocalPAMEntryUpdateRequest

`func NewPAMLocalLocalPAMEntryUpdateRequest(secretName string, secretValue string, ) *PAMLocalLocalPAMEntryUpdateRequest`

NewPAMLocalLocalPAMEntryUpdateRequest instantiates a new PAMLocalLocalPAMEntryUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPAMLocalLocalPAMEntryUpdateRequestWithDefaults

`func NewPAMLocalLocalPAMEntryUpdateRequestWithDefaults() *PAMLocalLocalPAMEntryUpdateRequest`

NewPAMLocalLocalPAMEntryUpdateRequestWithDefaults instantiates a new PAMLocalLocalPAMEntryUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecretName

`func (o *PAMLocalLocalPAMEntryUpdateRequest) GetSecretName() string`

GetSecretName returns the SecretName field if non-nil, zero value otherwise.

### GetSecretNameOk

`func (o *PAMLocalLocalPAMEntryUpdateRequest) GetSecretNameOk() (*string, bool)`

GetSecretNameOk returns a tuple with the SecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretName

`func (o *PAMLocalLocalPAMEntryUpdateRequest) SetSecretName(v string)`

SetSecretName sets SecretName field to given value.


### GetDescription

`func (o *PAMLocalLocalPAMEntryUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PAMLocalLocalPAMEntryUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PAMLocalLocalPAMEntryUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PAMLocalLocalPAMEntryUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PAMLocalLocalPAMEntryUpdateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PAMLocalLocalPAMEntryUpdateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSecretValue

`func (o *PAMLocalLocalPAMEntryUpdateRequest) GetSecretValue() string`

GetSecretValue returns the SecretValue field if non-nil, zero value otherwise.

### GetSecretValueOk

`func (o *PAMLocalLocalPAMEntryUpdateRequest) GetSecretValueOk() (*string, bool)`

GetSecretValueOk returns a tuple with the SecretValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretValue

`func (o *PAMLocalLocalPAMEntryUpdateRequest) SetSecretValue(v string)`

SetSecretValue sets SecretValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



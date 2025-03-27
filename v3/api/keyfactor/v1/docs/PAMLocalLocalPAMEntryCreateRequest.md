# PAMLocalLocalPAMEntryCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecretName** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**SecretValue** | **string** |  | 

## Methods

### NewPAMLocalLocalPAMEntryCreateRequest

`func NewPAMLocalLocalPAMEntryCreateRequest(secretName string, secretValue string, ) *PAMLocalLocalPAMEntryCreateRequest`

NewPAMLocalLocalPAMEntryCreateRequest instantiates a new PAMLocalLocalPAMEntryCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPAMLocalLocalPAMEntryCreateRequestWithDefaults

`func NewPAMLocalLocalPAMEntryCreateRequestWithDefaults() *PAMLocalLocalPAMEntryCreateRequest`

NewPAMLocalLocalPAMEntryCreateRequestWithDefaults instantiates a new PAMLocalLocalPAMEntryCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecretName

`func (o *PAMLocalLocalPAMEntryCreateRequest) GetSecretName() string`

GetSecretName returns the SecretName field if non-nil, zero value otherwise.

### GetSecretNameOk

`func (o *PAMLocalLocalPAMEntryCreateRequest) GetSecretNameOk() (*string, bool)`

GetSecretNameOk returns a tuple with the SecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretName

`func (o *PAMLocalLocalPAMEntryCreateRequest) SetSecretName(v string)`

SetSecretName sets SecretName field to given value.


### GetDescription

`func (o *PAMLocalLocalPAMEntryCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PAMLocalLocalPAMEntryCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PAMLocalLocalPAMEntryCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PAMLocalLocalPAMEntryCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PAMLocalLocalPAMEntryCreateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PAMLocalLocalPAMEntryCreateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSecretValue

`func (o *PAMLocalLocalPAMEntryCreateRequest) GetSecretValue() string`

GetSecretValue returns the SecretValue field if non-nil, zero value otherwise.

### GetSecretValueOk

`func (o *PAMLocalLocalPAMEntryCreateRequest) GetSecretValueOk() (*string, bool)`

GetSecretValueOk returns a tuple with the SecretValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretValue

`func (o *PAMLocalLocalPAMEntryCreateRequest) SetSecretValue(v string)`

SetSecretValue sets SecretValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



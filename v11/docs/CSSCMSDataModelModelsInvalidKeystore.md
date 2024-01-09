# CSSCMSDataModelModelsInvalidKeystore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeystoreId** | Pointer to **string** |  | [optional] 
**ClientMachine** | Pointer to **NullableString** |  | [optional] 
**StorePath** | Pointer to **NullableString** |  | [optional] 
**Alias** | Pointer to **NullableString** |  | [optional] 
**Reason** | Pointer to [**CSSCMSDataModelEnumsFailureType**](CSSCMSDataModelEnumsFailureType.md) |  | [optional] 
**Explanation** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsInvalidKeystore

`func NewCSSCMSDataModelModelsInvalidKeystore() *CSSCMSDataModelModelsInvalidKeystore`

NewCSSCMSDataModelModelsInvalidKeystore instantiates a new CSSCMSDataModelModelsInvalidKeystore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsInvalidKeystoreWithDefaults

`func NewCSSCMSDataModelModelsInvalidKeystoreWithDefaults() *CSSCMSDataModelModelsInvalidKeystore`

NewCSSCMSDataModelModelsInvalidKeystoreWithDefaults instantiates a new CSSCMSDataModelModelsInvalidKeystore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeystoreId

`func (o *CSSCMSDataModelModelsInvalidKeystore) GetKeystoreId() string`

GetKeystoreId returns the KeystoreId field if non-nil, zero value otherwise.

### GetKeystoreIdOk

`func (o *CSSCMSDataModelModelsInvalidKeystore) GetKeystoreIdOk() (*string, bool)`

GetKeystoreIdOk returns a tuple with the KeystoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystoreId

`func (o *CSSCMSDataModelModelsInvalidKeystore) SetKeystoreId(v string)`

SetKeystoreId sets KeystoreId field to given value.

### HasKeystoreId

`func (o *CSSCMSDataModelModelsInvalidKeystore) HasKeystoreId() bool`

HasKeystoreId returns a boolean if a field has been set.

### GetClientMachine

`func (o *CSSCMSDataModelModelsInvalidKeystore) GetClientMachine() string`

GetClientMachine returns the ClientMachine field if non-nil, zero value otherwise.

### GetClientMachineOk

`func (o *CSSCMSDataModelModelsInvalidKeystore) GetClientMachineOk() (*string, bool)`

GetClientMachineOk returns a tuple with the ClientMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMachine

`func (o *CSSCMSDataModelModelsInvalidKeystore) SetClientMachine(v string)`

SetClientMachine sets ClientMachine field to given value.

### HasClientMachine

`func (o *CSSCMSDataModelModelsInvalidKeystore) HasClientMachine() bool`

HasClientMachine returns a boolean if a field has been set.

### SetClientMachineNil

`func (o *CSSCMSDataModelModelsInvalidKeystore) SetClientMachineNil(b bool)`

 SetClientMachineNil sets the value for ClientMachine to be an explicit nil

### UnsetClientMachine
`func (o *CSSCMSDataModelModelsInvalidKeystore) UnsetClientMachine()`

UnsetClientMachine ensures that no value is present for ClientMachine, not even an explicit nil
### GetStorePath

`func (o *CSSCMSDataModelModelsInvalidKeystore) GetStorePath() string`

GetStorePath returns the StorePath field if non-nil, zero value otherwise.

### GetStorePathOk

`func (o *CSSCMSDataModelModelsInvalidKeystore) GetStorePathOk() (*string, bool)`

GetStorePathOk returns a tuple with the StorePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorePath

`func (o *CSSCMSDataModelModelsInvalidKeystore) SetStorePath(v string)`

SetStorePath sets StorePath field to given value.

### HasStorePath

`func (o *CSSCMSDataModelModelsInvalidKeystore) HasStorePath() bool`

HasStorePath returns a boolean if a field has been set.

### SetStorePathNil

`func (o *CSSCMSDataModelModelsInvalidKeystore) SetStorePathNil(b bool)`

 SetStorePathNil sets the value for StorePath to be an explicit nil

### UnsetStorePath
`func (o *CSSCMSDataModelModelsInvalidKeystore) UnsetStorePath()`

UnsetStorePath ensures that no value is present for StorePath, not even an explicit nil
### GetAlias

`func (o *CSSCMSDataModelModelsInvalidKeystore) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *CSSCMSDataModelModelsInvalidKeystore) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *CSSCMSDataModelModelsInvalidKeystore) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *CSSCMSDataModelModelsInvalidKeystore) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### SetAliasNil

`func (o *CSSCMSDataModelModelsInvalidKeystore) SetAliasNil(b bool)`

 SetAliasNil sets the value for Alias to be an explicit nil

### UnsetAlias
`func (o *CSSCMSDataModelModelsInvalidKeystore) UnsetAlias()`

UnsetAlias ensures that no value is present for Alias, not even an explicit nil
### GetReason

`func (o *CSSCMSDataModelModelsInvalidKeystore) GetReason() CSSCMSDataModelEnumsFailureType`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CSSCMSDataModelModelsInvalidKeystore) GetReasonOk() (*CSSCMSDataModelEnumsFailureType, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CSSCMSDataModelModelsInvalidKeystore) SetReason(v CSSCMSDataModelEnumsFailureType)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CSSCMSDataModelModelsInvalidKeystore) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetExplanation

`func (o *CSSCMSDataModelModelsInvalidKeystore) GetExplanation() string`

GetExplanation returns the Explanation field if non-nil, zero value otherwise.

### GetExplanationOk

`func (o *CSSCMSDataModelModelsInvalidKeystore) GetExplanationOk() (*string, bool)`

GetExplanationOk returns a tuple with the Explanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanation

`func (o *CSSCMSDataModelModelsInvalidKeystore) SetExplanation(v string)`

SetExplanation sets Explanation field to given value.

### HasExplanation

`func (o *CSSCMSDataModelModelsInvalidKeystore) HasExplanation() bool`

HasExplanation returns a boolean if a field has been set.

### SetExplanationNil

`func (o *CSSCMSDataModelModelsInvalidKeystore) SetExplanationNil(b bool)`

 SetExplanationNil sets the value for Explanation to be an explicit nil

### UnsetExplanation
`func (o *CSSCMSDataModelModelsInvalidKeystore) UnsetExplanation()`

UnsetExplanation ensures that no value is present for Explanation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



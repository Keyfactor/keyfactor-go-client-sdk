# CSSCMSDataModelModelsCertificateStoreEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateStoreId** | **string** |  | 
**Alias** | Pointer to **NullableString** |  | [optional] 
**JobFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Overwrite** | Pointer to **bool** |  | [optional] 
**EntryPassword** | Pointer to [**CSSCMSDataModelModelsKeyfactorAPISecret**](CSSCMSDataModelModelsKeyfactorAPISecret.md) |  | [optional] 
**PfxPassword** | Pointer to **NullableString** |  | [optional] 
**IncludePrivateKey** | Pointer to **bool** |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsCertificateStoreEntry

`func NewCSSCMSDataModelModelsCertificateStoreEntry(certificateStoreId string, ) *CSSCMSDataModelModelsCertificateStoreEntry`

NewCSSCMSDataModelModelsCertificateStoreEntry instantiates a new CSSCMSDataModelModelsCertificateStoreEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsCertificateStoreEntryWithDefaults

`func NewCSSCMSDataModelModelsCertificateStoreEntryWithDefaults() *CSSCMSDataModelModelsCertificateStoreEntry`

NewCSSCMSDataModelModelsCertificateStoreEntryWithDefaults instantiates a new CSSCMSDataModelModelsCertificateStoreEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateStoreId

`func (o *CSSCMSDataModelModelsCertificateStoreEntry) GetCertificateStoreId() string`

GetCertificateStoreId returns the CertificateStoreId field if non-nil, zero value otherwise.

### GetCertificateStoreIdOk

`func (o *CSSCMSDataModelModelsCertificateStoreEntry) GetCertificateStoreIdOk() (*string, bool)`

GetCertificateStoreIdOk returns a tuple with the CertificateStoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateStoreId

`func (o *CSSCMSDataModelModelsCertificateStoreEntry) SetCertificateStoreId(v string)`

SetCertificateStoreId sets CertificateStoreId field to given value.


### GetAlias

`func (o *CSSCMSDataModelModelsCertificateStoreEntry) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *CSSCMSDataModelModelsCertificateStoreEntry) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *CSSCMSDataModelModelsCertificateStoreEntry) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *CSSCMSDataModelModelsCertificateStoreEntry) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### SetAliasNil

`func (o *CSSCMSDataModelModelsCertificateStoreEntry) SetAliasNil(b bool)`

 SetAliasNil sets the value for Alias to be an explicit nil

### UnsetAlias
`func (o *CSSCMSDataModelModelsCertificateStoreEntry) UnsetAlias()`

UnsetAlias ensures that no value is present for Alias, not even an explicit nil
### GetJobFields

`func (o *CSSCMSDataModelModelsCertificateStoreEntry) GetJobFields() map[string]interface{}`

GetJobFields returns the JobFields field if non-nil, zero value otherwise.

### GetJobFieldsOk

`func (o *CSSCMSDataModelModelsCertificateStoreEntry) GetJobFieldsOk() (*map[string]interface{}, bool)`

GetJobFieldsOk returns a tuple with the JobFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobFields

`func (o *CSSCMSDataModelModelsCertificateStoreEntry) SetJobFields(v map[string]interface{})`

SetJobFields sets JobFields field to given value.

### HasJobFields

`func (o *CSSCMSDataModelModelsCertificateStoreEntry) HasJobFields() bool`

HasJobFields returns a boolean if a field has been set.

### SetJobFieldsNil

`func (o *CSSCMSDataModelModelsCertificateStoreEntry) SetJobFieldsNil(b bool)`

 SetJobFieldsNil sets the value for JobFields to be an explicit nil

### UnsetJobFields
`func (o *CSSCMSDataModelModelsCertificateStoreEntry) UnsetJobFields()`

UnsetJobFields ensures that no value is present for JobFields, not even an explicit nil
### GetOverwrite

`func (o *CSSCMSDataModelModelsCertificateStoreEntry) GetOverwrite() bool`

GetOverwrite returns the Overwrite field if non-nil, zero value otherwise.

### GetOverwriteOk

`func (o *CSSCMSDataModelModelsCertificateStoreEntry) GetOverwriteOk() (*bool, bool)`

GetOverwriteOk returns a tuple with the Overwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwrite

`func (o *CSSCMSDataModelModelsCertificateStoreEntry) SetOverwrite(v bool)`

SetOverwrite sets Overwrite field to given value.

### HasOverwrite

`func (o *CSSCMSDataModelModelsCertificateStoreEntry) HasOverwrite() bool`

HasOverwrite returns a boolean if a field has been set.

### GetEntryPassword

`func (o *CSSCMSDataModelModelsCertificateStoreEntry) GetEntryPassword() CSSCMSDataModelModelsKeyfactorAPISecret`

GetEntryPassword returns the EntryPassword field if non-nil, zero value otherwise.

### GetEntryPasswordOk

`func (o *CSSCMSDataModelModelsCertificateStoreEntry) GetEntryPasswordOk() (*CSSCMSDataModelModelsKeyfactorAPISecret, bool)`

GetEntryPasswordOk returns a tuple with the EntryPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryPassword

`func (o *CSSCMSDataModelModelsCertificateStoreEntry) SetEntryPassword(v CSSCMSDataModelModelsKeyfactorAPISecret)`

SetEntryPassword sets EntryPassword field to given value.

### HasEntryPassword

`func (o *CSSCMSDataModelModelsCertificateStoreEntry) HasEntryPassword() bool`

HasEntryPassword returns a boolean if a field has been set.

### GetPfxPassword

`func (o *CSSCMSDataModelModelsCertificateStoreEntry) GetPfxPassword() string`

GetPfxPassword returns the PfxPassword field if non-nil, zero value otherwise.

### GetPfxPasswordOk

`func (o *CSSCMSDataModelModelsCertificateStoreEntry) GetPfxPasswordOk() (*string, bool)`

GetPfxPasswordOk returns a tuple with the PfxPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfxPassword

`func (o *CSSCMSDataModelModelsCertificateStoreEntry) SetPfxPassword(v string)`

SetPfxPassword sets PfxPassword field to given value.

### HasPfxPassword

`func (o *CSSCMSDataModelModelsCertificateStoreEntry) HasPfxPassword() bool`

HasPfxPassword returns a boolean if a field has been set.

### SetPfxPasswordNil

`func (o *CSSCMSDataModelModelsCertificateStoreEntry) SetPfxPasswordNil(b bool)`

 SetPfxPasswordNil sets the value for PfxPassword to be an explicit nil

### UnsetPfxPassword
`func (o *CSSCMSDataModelModelsCertificateStoreEntry) UnsetPfxPassword()`

UnsetPfxPassword ensures that no value is present for PfxPassword, not even an explicit nil
### GetIncludePrivateKey

`func (o *CSSCMSDataModelModelsCertificateStoreEntry) GetIncludePrivateKey() bool`

GetIncludePrivateKey returns the IncludePrivateKey field if non-nil, zero value otherwise.

### GetIncludePrivateKeyOk

`func (o *CSSCMSDataModelModelsCertificateStoreEntry) GetIncludePrivateKeyOk() (*bool, bool)`

GetIncludePrivateKeyOk returns a tuple with the IncludePrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePrivateKey

`func (o *CSSCMSDataModelModelsCertificateStoreEntry) SetIncludePrivateKey(v bool)`

SetIncludePrivateKey sets IncludePrivateKey field to given value.

### HasIncludePrivateKey

`func (o *CSSCMSDataModelModelsCertificateStoreEntry) HasIncludePrivateKey() bool`

HasIncludePrivateKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



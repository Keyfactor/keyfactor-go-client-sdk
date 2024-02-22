# ModelsCertificateStoreEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateStoreId** | **string** |  | 
**Alias** | Pointer to **string** |  | [optional] 
**JobFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Overwrite** | Pointer to **bool** |  | [optional] 
**EntryPassword** | Pointer to [**ModelsKeyfactorAPISecret**](ModelsKeyfactorAPISecret.md) |  | [optional] 
**PfxPassword** | Pointer to **string** | The PFX password. | [optional] 
**IncludePrivateKey** | Pointer to **bool** |  | [optional] 

## Methods

### NewModelsCertificateStoreEntry

`func NewModelsCertificateStoreEntry(certificateStoreId string, ) *ModelsCertificateStoreEntry`

NewModelsCertificateStoreEntry instantiates a new ModelsCertificateStoreEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsCertificateStoreEntryWithDefaults

`func NewModelsCertificateStoreEntryWithDefaults() *ModelsCertificateStoreEntry`

NewModelsCertificateStoreEntryWithDefaults instantiates a new ModelsCertificateStoreEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateStoreId

`func (o *ModelsCertificateStoreEntry) GetCertificateStoreId() string`

GetCertificateStoreId returns the CertificateStoreId field if non-nil, zero value otherwise.

### GetCertificateStoreIdOk

`func (o *ModelsCertificateStoreEntry) GetCertificateStoreIdOk() (*string, bool)`

GetCertificateStoreIdOk returns a tuple with the CertificateStoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateStoreId

`func (o *ModelsCertificateStoreEntry) SetCertificateStoreId(v string)`

SetCertificateStoreId sets CertificateStoreId field to given value.


### GetAlias

`func (o *ModelsCertificateStoreEntry) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *ModelsCertificateStoreEntry) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *ModelsCertificateStoreEntry) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *ModelsCertificateStoreEntry) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetJobFields

`func (o *ModelsCertificateStoreEntry) GetJobFields() map[string]map[string]interface{}`

GetJobFields returns the JobFields field if non-nil, zero value otherwise.

### GetJobFieldsOk

`func (o *ModelsCertificateStoreEntry) GetJobFieldsOk() (*map[string]map[string]interface{}, bool)`

GetJobFieldsOk returns a tuple with the JobFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobFields

`func (o *ModelsCertificateStoreEntry) SetJobFields(v map[string]map[string]interface{})`

SetJobFields sets JobFields field to given value.

### HasJobFields

`func (o *ModelsCertificateStoreEntry) HasJobFields() bool`

HasJobFields returns a boolean if a field has been set.

### GetOverwrite

`func (o *ModelsCertificateStoreEntry) GetOverwrite() bool`

GetOverwrite returns the Overwrite field if non-nil, zero value otherwise.

### GetOverwriteOk

`func (o *ModelsCertificateStoreEntry) GetOverwriteOk() (*bool, bool)`

GetOverwriteOk returns a tuple with the Overwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwrite

`func (o *ModelsCertificateStoreEntry) SetOverwrite(v bool)`

SetOverwrite sets Overwrite field to given value.

### HasOverwrite

`func (o *ModelsCertificateStoreEntry) HasOverwrite() bool`

HasOverwrite returns a boolean if a field has been set.

### GetEntryPassword

`func (o *ModelsCertificateStoreEntry) GetEntryPassword() ModelsKeyfactorAPISecret`

GetEntryPassword returns the EntryPassword field if non-nil, zero value otherwise.

### GetEntryPasswordOk

`func (o *ModelsCertificateStoreEntry) GetEntryPasswordOk() (*ModelsKeyfactorAPISecret, bool)`

GetEntryPasswordOk returns a tuple with the EntryPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryPassword

`func (o *ModelsCertificateStoreEntry) SetEntryPassword(v ModelsKeyfactorAPISecret)`

SetEntryPassword sets EntryPassword field to given value.

### HasEntryPassword

`func (o *ModelsCertificateStoreEntry) HasEntryPassword() bool`

HasEntryPassword returns a boolean if a field has been set.

### GetPfxPassword

`func (o *ModelsCertificateStoreEntry) GetPfxPassword() string`

GetPfxPassword returns the PfxPassword field if non-nil, zero value otherwise.

### GetPfxPasswordOk

`func (o *ModelsCertificateStoreEntry) GetPfxPasswordOk() (*string, bool)`

GetPfxPasswordOk returns a tuple with the PfxPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfxPassword

`func (o *ModelsCertificateStoreEntry) SetPfxPassword(v string)`

SetPfxPassword sets PfxPassword field to given value.

### HasPfxPassword

`func (o *ModelsCertificateStoreEntry) HasPfxPassword() bool`

HasPfxPassword returns a boolean if a field has been set.

### GetIncludePrivateKey

`func (o *ModelsCertificateStoreEntry) GetIncludePrivateKey() bool`

GetIncludePrivateKey returns the IncludePrivateKey field if non-nil, zero value otherwise.

### GetIncludePrivateKeyOk

`func (o *ModelsCertificateStoreEntry) GetIncludePrivateKeyOk() (*bool, bool)`

GetIncludePrivateKeyOk returns a tuple with the IncludePrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePrivateKey

`func (o *ModelsCertificateStoreEntry) SetIncludePrivateKey(v bool)`

SetIncludePrivateKey sets IncludePrivateKey field to given value.

### HasIncludePrivateKey

`func (o *ModelsCertificateStoreEntry) HasIncludePrivateKey() bool`

HasIncludePrivateKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



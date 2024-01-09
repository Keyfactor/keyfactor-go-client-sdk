# KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Query** | Pointer to **NullableString** |  | [optional] 
**DuplicationField** | Pointer to [**CSSCMSCoreEnumsDuplicateSubjectType**](CSSCMSCoreEnumsDuplicateSubjectType.md) |  | [optional] 
**ShowOnDashboard** | Pointer to **bool** |  | [optional] 
**Favorite** | Pointer to **bool** |  | [optional] 
**CopyFromId** | Pointer to **NullableInt32** |  | [optional] 
**Id** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest

`func NewKeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest(name string, ) *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest`

NewKeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest instantiates a new KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequestWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequestWithDefaults() *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest`

NewKeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequestWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetQuery

`func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQueryNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest) SetQueryNil(b bool)`

 SetQueryNil sets the value for Query to be an explicit nil

### UnsetQuery
`func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest) UnsetQuery()`

UnsetQuery ensures that no value is present for Query, not even an explicit nil
### GetDuplicationField

`func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest) GetDuplicationField() CSSCMSCoreEnumsDuplicateSubjectType`

GetDuplicationField returns the DuplicationField field if non-nil, zero value otherwise.

### GetDuplicationFieldOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest) GetDuplicationFieldOk() (*CSSCMSCoreEnumsDuplicateSubjectType, bool)`

GetDuplicationFieldOk returns a tuple with the DuplicationField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicationField

`func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest) SetDuplicationField(v CSSCMSCoreEnumsDuplicateSubjectType)`

SetDuplicationField sets DuplicationField field to given value.

### HasDuplicationField

`func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest) HasDuplicationField() bool`

HasDuplicationField returns a boolean if a field has been set.

### GetShowOnDashboard

`func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest) GetShowOnDashboard() bool`

GetShowOnDashboard returns the ShowOnDashboard field if non-nil, zero value otherwise.

### GetShowOnDashboardOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest) GetShowOnDashboardOk() (*bool, bool)`

GetShowOnDashboardOk returns a tuple with the ShowOnDashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowOnDashboard

`func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest) SetShowOnDashboard(v bool)`

SetShowOnDashboard sets ShowOnDashboard field to given value.

### HasShowOnDashboard

`func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest) HasShowOnDashboard() bool`

HasShowOnDashboard returns a boolean if a field has been set.

### GetFavorite

`func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest) GetFavorite() bool`

GetFavorite returns the Favorite field if non-nil, zero value otherwise.

### GetFavoriteOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest) GetFavoriteOk() (*bool, bool)`

GetFavoriteOk returns a tuple with the Favorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorite

`func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest) SetFavorite(v bool)`

SetFavorite sets Favorite field to given value.

### HasFavorite

`func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest) HasFavorite() bool`

HasFavorite returns a boolean if a field has been set.

### GetCopyFromId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest) GetCopyFromId() int32`

GetCopyFromId returns the CopyFromId field if non-nil, zero value otherwise.

### GetCopyFromIdOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest) GetCopyFromIdOk() (*int32, bool)`

GetCopyFromIdOk returns a tuple with the CopyFromId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyFromId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest) SetCopyFromId(v int32)`

SetCopyFromId sets CopyFromId field to given value.

### HasCopyFromId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest) HasCopyFromId() bool`

HasCopyFromId returns a boolean if a field has been set.

### SetCopyFromIdNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest) SetCopyFromIdNil(b bool)`

 SetCopyFromIdNil sets the value for CopyFromId to be an explicit nil

### UnsetCopyFromId
`func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest) UnsetCopyFromId()`

UnsetCopyFromId ensures that no value is present for CopyFromId, not even an explicit nil
### GetId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



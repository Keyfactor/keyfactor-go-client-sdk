# CertificateCollectionsCertificateCollectionCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CopyFromId** | Pointer to **NullableInt32** |  | [optional] 
**Id** | Pointer to **NullableInt32** |  | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Query** | Pointer to **NullableString** |  | [optional] 
**DuplicationField** | Pointer to [**CSSCMSCoreEnumsDuplicateSubjectType**](CSSCMSCoreEnumsDuplicateSubjectType.md) |  | [optional] 
**ShowOnDashboard** | Pointer to **bool** |  | [optional] 
**Favorite** | Pointer to **bool** |  | [optional] 

## Methods

### NewCertificateCollectionsCertificateCollectionCreateRequest

`func NewCertificateCollectionsCertificateCollectionCreateRequest(name string, ) *CertificateCollectionsCertificateCollectionCreateRequest`

NewCertificateCollectionsCertificateCollectionCreateRequest instantiates a new CertificateCollectionsCertificateCollectionCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateCollectionsCertificateCollectionCreateRequestWithDefaults

`func NewCertificateCollectionsCertificateCollectionCreateRequestWithDefaults() *CertificateCollectionsCertificateCollectionCreateRequest`

NewCertificateCollectionsCertificateCollectionCreateRequestWithDefaults instantiates a new CertificateCollectionsCertificateCollectionCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCopyFromId

`func (o *CertificateCollectionsCertificateCollectionCreateRequest) GetCopyFromId() int32`

GetCopyFromId returns the CopyFromId field if non-nil, zero value otherwise.

### GetCopyFromIdOk

`func (o *CertificateCollectionsCertificateCollectionCreateRequest) GetCopyFromIdOk() (*int32, bool)`

GetCopyFromIdOk returns a tuple with the CopyFromId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyFromId

`func (o *CertificateCollectionsCertificateCollectionCreateRequest) SetCopyFromId(v int32)`

SetCopyFromId sets CopyFromId field to given value.

### HasCopyFromId

`func (o *CertificateCollectionsCertificateCollectionCreateRequest) HasCopyFromId() bool`

HasCopyFromId returns a boolean if a field has been set.

### SetCopyFromIdNil

`func (o *CertificateCollectionsCertificateCollectionCreateRequest) SetCopyFromIdNil(b bool)`

 SetCopyFromIdNil sets the value for CopyFromId to be an explicit nil

### UnsetCopyFromId
`func (o *CertificateCollectionsCertificateCollectionCreateRequest) UnsetCopyFromId()`

UnsetCopyFromId ensures that no value is present for CopyFromId, not even an explicit nil
### GetId

`func (o *CertificateCollectionsCertificateCollectionCreateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificateCollectionsCertificateCollectionCreateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificateCollectionsCertificateCollectionCreateRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CertificateCollectionsCertificateCollectionCreateRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CertificateCollectionsCertificateCollectionCreateRequest) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CertificateCollectionsCertificateCollectionCreateRequest) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *CertificateCollectionsCertificateCollectionCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificateCollectionsCertificateCollectionCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificateCollectionsCertificateCollectionCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CertificateCollectionsCertificateCollectionCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CertificateCollectionsCertificateCollectionCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CertificateCollectionsCertificateCollectionCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CertificateCollectionsCertificateCollectionCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CertificateCollectionsCertificateCollectionCreateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CertificateCollectionsCertificateCollectionCreateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetQuery

`func (o *CertificateCollectionsCertificateCollectionCreateRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *CertificateCollectionsCertificateCollectionCreateRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *CertificateCollectionsCertificateCollectionCreateRequest) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *CertificateCollectionsCertificateCollectionCreateRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQueryNil

`func (o *CertificateCollectionsCertificateCollectionCreateRequest) SetQueryNil(b bool)`

 SetQueryNil sets the value for Query to be an explicit nil

### UnsetQuery
`func (o *CertificateCollectionsCertificateCollectionCreateRequest) UnsetQuery()`

UnsetQuery ensures that no value is present for Query, not even an explicit nil
### GetDuplicationField

`func (o *CertificateCollectionsCertificateCollectionCreateRequest) GetDuplicationField() CSSCMSCoreEnumsDuplicateSubjectType`

GetDuplicationField returns the DuplicationField field if non-nil, zero value otherwise.

### GetDuplicationFieldOk

`func (o *CertificateCollectionsCertificateCollectionCreateRequest) GetDuplicationFieldOk() (*CSSCMSCoreEnumsDuplicateSubjectType, bool)`

GetDuplicationFieldOk returns a tuple with the DuplicationField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicationField

`func (o *CertificateCollectionsCertificateCollectionCreateRequest) SetDuplicationField(v CSSCMSCoreEnumsDuplicateSubjectType)`

SetDuplicationField sets DuplicationField field to given value.

### HasDuplicationField

`func (o *CertificateCollectionsCertificateCollectionCreateRequest) HasDuplicationField() bool`

HasDuplicationField returns a boolean if a field has been set.

### GetShowOnDashboard

`func (o *CertificateCollectionsCertificateCollectionCreateRequest) GetShowOnDashboard() bool`

GetShowOnDashboard returns the ShowOnDashboard field if non-nil, zero value otherwise.

### GetShowOnDashboardOk

`func (o *CertificateCollectionsCertificateCollectionCreateRequest) GetShowOnDashboardOk() (*bool, bool)`

GetShowOnDashboardOk returns a tuple with the ShowOnDashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowOnDashboard

`func (o *CertificateCollectionsCertificateCollectionCreateRequest) SetShowOnDashboard(v bool)`

SetShowOnDashboard sets ShowOnDashboard field to given value.

### HasShowOnDashboard

`func (o *CertificateCollectionsCertificateCollectionCreateRequest) HasShowOnDashboard() bool`

HasShowOnDashboard returns a boolean if a field has been set.

### GetFavorite

`func (o *CertificateCollectionsCertificateCollectionCreateRequest) GetFavorite() bool`

GetFavorite returns the Favorite field if non-nil, zero value otherwise.

### GetFavoriteOk

`func (o *CertificateCollectionsCertificateCollectionCreateRequest) GetFavoriteOk() (*bool, bool)`

GetFavoriteOk returns a tuple with the Favorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorite

`func (o *CertificateCollectionsCertificateCollectionCreateRequest) SetFavorite(v bool)`

SetFavorite sets Favorite field to given value.

### HasFavorite

`func (o *CertificateCollectionsCertificateCollectionCreateRequest) HasFavorite() bool`

HasFavorite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CertificateCollectionsCertificateCollectionCopyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CopyFromId** | **int32** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Query** | Pointer to **NullableString** |  | [optional] 
**DuplicationField** | Pointer to [**CSSCMSCoreEnumsDuplicateSubjectType**](CSSCMSCoreEnumsDuplicateSubjectType.md) |  | [optional] 
**ShowOnDashboard** | Pointer to **bool** |  | [optional] 
**Favorite** | Pointer to **bool** |  | [optional] 

## Methods

### NewCertificateCollectionsCertificateCollectionCopyRequest

`func NewCertificateCollectionsCertificateCollectionCopyRequest(copyFromId int32, name string, ) *CertificateCollectionsCertificateCollectionCopyRequest`

NewCertificateCollectionsCertificateCollectionCopyRequest instantiates a new CertificateCollectionsCertificateCollectionCopyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateCollectionsCertificateCollectionCopyRequestWithDefaults

`func NewCertificateCollectionsCertificateCollectionCopyRequestWithDefaults() *CertificateCollectionsCertificateCollectionCopyRequest`

NewCertificateCollectionsCertificateCollectionCopyRequestWithDefaults instantiates a new CertificateCollectionsCertificateCollectionCopyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCopyFromId

`func (o *CertificateCollectionsCertificateCollectionCopyRequest) GetCopyFromId() int32`

GetCopyFromId returns the CopyFromId field if non-nil, zero value otherwise.

### GetCopyFromIdOk

`func (o *CertificateCollectionsCertificateCollectionCopyRequest) GetCopyFromIdOk() (*int32, bool)`

GetCopyFromIdOk returns a tuple with the CopyFromId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyFromId

`func (o *CertificateCollectionsCertificateCollectionCopyRequest) SetCopyFromId(v int32)`

SetCopyFromId sets CopyFromId field to given value.


### GetName

`func (o *CertificateCollectionsCertificateCollectionCopyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificateCollectionsCertificateCollectionCopyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificateCollectionsCertificateCollectionCopyRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CertificateCollectionsCertificateCollectionCopyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CertificateCollectionsCertificateCollectionCopyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CertificateCollectionsCertificateCollectionCopyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CertificateCollectionsCertificateCollectionCopyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CertificateCollectionsCertificateCollectionCopyRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CertificateCollectionsCertificateCollectionCopyRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetQuery

`func (o *CertificateCollectionsCertificateCollectionCopyRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *CertificateCollectionsCertificateCollectionCopyRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *CertificateCollectionsCertificateCollectionCopyRequest) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *CertificateCollectionsCertificateCollectionCopyRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQueryNil

`func (o *CertificateCollectionsCertificateCollectionCopyRequest) SetQueryNil(b bool)`

 SetQueryNil sets the value for Query to be an explicit nil

### UnsetQuery
`func (o *CertificateCollectionsCertificateCollectionCopyRequest) UnsetQuery()`

UnsetQuery ensures that no value is present for Query, not even an explicit nil
### GetDuplicationField

`func (o *CertificateCollectionsCertificateCollectionCopyRequest) GetDuplicationField() CSSCMSCoreEnumsDuplicateSubjectType`

GetDuplicationField returns the DuplicationField field if non-nil, zero value otherwise.

### GetDuplicationFieldOk

`func (o *CertificateCollectionsCertificateCollectionCopyRequest) GetDuplicationFieldOk() (*CSSCMSCoreEnumsDuplicateSubjectType, bool)`

GetDuplicationFieldOk returns a tuple with the DuplicationField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicationField

`func (o *CertificateCollectionsCertificateCollectionCopyRequest) SetDuplicationField(v CSSCMSCoreEnumsDuplicateSubjectType)`

SetDuplicationField sets DuplicationField field to given value.

### HasDuplicationField

`func (o *CertificateCollectionsCertificateCollectionCopyRequest) HasDuplicationField() bool`

HasDuplicationField returns a boolean if a field has been set.

### GetShowOnDashboard

`func (o *CertificateCollectionsCertificateCollectionCopyRequest) GetShowOnDashboard() bool`

GetShowOnDashboard returns the ShowOnDashboard field if non-nil, zero value otherwise.

### GetShowOnDashboardOk

`func (o *CertificateCollectionsCertificateCollectionCopyRequest) GetShowOnDashboardOk() (*bool, bool)`

GetShowOnDashboardOk returns a tuple with the ShowOnDashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowOnDashboard

`func (o *CertificateCollectionsCertificateCollectionCopyRequest) SetShowOnDashboard(v bool)`

SetShowOnDashboard sets ShowOnDashboard field to given value.

### HasShowOnDashboard

`func (o *CertificateCollectionsCertificateCollectionCopyRequest) HasShowOnDashboard() bool`

HasShowOnDashboard returns a boolean if a field has been set.

### GetFavorite

`func (o *CertificateCollectionsCertificateCollectionCopyRequest) GetFavorite() bool`

GetFavorite returns the Favorite field if non-nil, zero value otherwise.

### GetFavoriteOk

`func (o *CertificateCollectionsCertificateCollectionCopyRequest) GetFavoriteOk() (*bool, bool)`

GetFavoriteOk returns a tuple with the Favorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorite

`func (o *CertificateCollectionsCertificateCollectionCopyRequest) SetFavorite(v bool)`

SetFavorite sets Favorite field to given value.

### HasFavorite

`func (o *CertificateCollectionsCertificateCollectionCopyRequest) HasFavorite() bool`

HasFavorite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



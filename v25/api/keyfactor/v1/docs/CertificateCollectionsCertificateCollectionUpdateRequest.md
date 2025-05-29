# CertificateCollectionsCertificateCollectionUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Query** | Pointer to **NullableString** |  | [optional] 
**DuplicationField** | Pointer to [**CSSCMSCoreEnumsDuplicateSubjectType**](CSSCMSCoreEnumsDuplicateSubjectType.md) |  | [optional] 
**ShowOnDashboard** | Pointer to **bool** |  | [optional] 
**Favorite** | Pointer to **bool** |  | [optional] 

## Methods

### NewCertificateCollectionsCertificateCollectionUpdateRequest

`func NewCertificateCollectionsCertificateCollectionUpdateRequest(id int32, name string, ) *CertificateCollectionsCertificateCollectionUpdateRequest`

NewCertificateCollectionsCertificateCollectionUpdateRequest instantiates a new CertificateCollectionsCertificateCollectionUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateCollectionsCertificateCollectionUpdateRequestWithDefaults

`func NewCertificateCollectionsCertificateCollectionUpdateRequestWithDefaults() *CertificateCollectionsCertificateCollectionUpdateRequest`

NewCertificateCollectionsCertificateCollectionUpdateRequestWithDefaults instantiates a new CertificateCollectionsCertificateCollectionUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CertificateCollectionsCertificateCollectionUpdateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificateCollectionsCertificateCollectionUpdateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificateCollectionsCertificateCollectionUpdateRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *CertificateCollectionsCertificateCollectionUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificateCollectionsCertificateCollectionUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificateCollectionsCertificateCollectionUpdateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CertificateCollectionsCertificateCollectionUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CertificateCollectionsCertificateCollectionUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CertificateCollectionsCertificateCollectionUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CertificateCollectionsCertificateCollectionUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CertificateCollectionsCertificateCollectionUpdateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CertificateCollectionsCertificateCollectionUpdateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetQuery

`func (o *CertificateCollectionsCertificateCollectionUpdateRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *CertificateCollectionsCertificateCollectionUpdateRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *CertificateCollectionsCertificateCollectionUpdateRequest) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *CertificateCollectionsCertificateCollectionUpdateRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQueryNil

`func (o *CertificateCollectionsCertificateCollectionUpdateRequest) SetQueryNil(b bool)`

 SetQueryNil sets the value for Query to be an explicit nil

### UnsetQuery
`func (o *CertificateCollectionsCertificateCollectionUpdateRequest) UnsetQuery()`

UnsetQuery ensures that no value is present for Query, not even an explicit nil
### GetDuplicationField

`func (o *CertificateCollectionsCertificateCollectionUpdateRequest) GetDuplicationField() CSSCMSCoreEnumsDuplicateSubjectType`

GetDuplicationField returns the DuplicationField field if non-nil, zero value otherwise.

### GetDuplicationFieldOk

`func (o *CertificateCollectionsCertificateCollectionUpdateRequest) GetDuplicationFieldOk() (*CSSCMSCoreEnumsDuplicateSubjectType, bool)`

GetDuplicationFieldOk returns a tuple with the DuplicationField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicationField

`func (o *CertificateCollectionsCertificateCollectionUpdateRequest) SetDuplicationField(v CSSCMSCoreEnumsDuplicateSubjectType)`

SetDuplicationField sets DuplicationField field to given value.

### HasDuplicationField

`func (o *CertificateCollectionsCertificateCollectionUpdateRequest) HasDuplicationField() bool`

HasDuplicationField returns a boolean if a field has been set.

### GetShowOnDashboard

`func (o *CertificateCollectionsCertificateCollectionUpdateRequest) GetShowOnDashboard() bool`

GetShowOnDashboard returns the ShowOnDashboard field if non-nil, zero value otherwise.

### GetShowOnDashboardOk

`func (o *CertificateCollectionsCertificateCollectionUpdateRequest) GetShowOnDashboardOk() (*bool, bool)`

GetShowOnDashboardOk returns a tuple with the ShowOnDashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowOnDashboard

`func (o *CertificateCollectionsCertificateCollectionUpdateRequest) SetShowOnDashboard(v bool)`

SetShowOnDashboard sets ShowOnDashboard field to given value.

### HasShowOnDashboard

`func (o *CertificateCollectionsCertificateCollectionUpdateRequest) HasShowOnDashboard() bool`

HasShowOnDashboard returns a boolean if a field has been set.

### GetFavorite

`func (o *CertificateCollectionsCertificateCollectionUpdateRequest) GetFavorite() bool`

GetFavorite returns the Favorite field if non-nil, zero value otherwise.

### GetFavoriteOk

`func (o *CertificateCollectionsCertificateCollectionUpdateRequest) GetFavoriteOk() (*bool, bool)`

GetFavoriteOk returns a tuple with the Favorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorite

`func (o *CertificateCollectionsCertificateCollectionUpdateRequest) SetFavorite(v bool)`

SetFavorite sets Favorite field to given value.

### HasFavorite

`func (o *CertificateCollectionsCertificateCollectionUpdateRequest) HasFavorite() bool`

HasFavorite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CertificateCollectionsCertificateCollectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Content** | Pointer to **NullableString** |  | [optional] [readonly] 
**Query** | Pointer to **NullableString** |  | [optional] 
**DuplicationField** | Pointer to [**CSSCMSCoreEnumsDuplicateSubjectType**](CSSCMSCoreEnumsDuplicateSubjectType.md) |  | [optional] 
**ShowOnDashboard** | Pointer to **bool** |  | [optional] 
**Favorite** | Pointer to **bool** |  | [optional] 
**EstimatedCertCount** | Pointer to **int32** |  | [optional] 
**LastEstimated** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewCertificateCollectionsCertificateCollectionResponse

`func NewCertificateCollectionsCertificateCollectionResponse() *CertificateCollectionsCertificateCollectionResponse`

NewCertificateCollectionsCertificateCollectionResponse instantiates a new CertificateCollectionsCertificateCollectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateCollectionsCertificateCollectionResponseWithDefaults

`func NewCertificateCollectionsCertificateCollectionResponseWithDefaults() *CertificateCollectionsCertificateCollectionResponse`

NewCertificateCollectionsCertificateCollectionResponseWithDefaults instantiates a new CertificateCollectionsCertificateCollectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CertificateCollectionsCertificateCollectionResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificateCollectionsCertificateCollectionResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificateCollectionsCertificateCollectionResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CertificateCollectionsCertificateCollectionResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CertificateCollectionsCertificateCollectionResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificateCollectionsCertificateCollectionResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificateCollectionsCertificateCollectionResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CertificateCollectionsCertificateCollectionResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CertificateCollectionsCertificateCollectionResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CertificateCollectionsCertificateCollectionResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *CertificateCollectionsCertificateCollectionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CertificateCollectionsCertificateCollectionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CertificateCollectionsCertificateCollectionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CertificateCollectionsCertificateCollectionResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CertificateCollectionsCertificateCollectionResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CertificateCollectionsCertificateCollectionResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetContent

`func (o *CertificateCollectionsCertificateCollectionResponse) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CertificateCollectionsCertificateCollectionResponse) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CertificateCollectionsCertificateCollectionResponse) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *CertificateCollectionsCertificateCollectionResponse) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *CertificateCollectionsCertificateCollectionResponse) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *CertificateCollectionsCertificateCollectionResponse) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetQuery

`func (o *CertificateCollectionsCertificateCollectionResponse) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *CertificateCollectionsCertificateCollectionResponse) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *CertificateCollectionsCertificateCollectionResponse) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *CertificateCollectionsCertificateCollectionResponse) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQueryNil

`func (o *CertificateCollectionsCertificateCollectionResponse) SetQueryNil(b bool)`

 SetQueryNil sets the value for Query to be an explicit nil

### UnsetQuery
`func (o *CertificateCollectionsCertificateCollectionResponse) UnsetQuery()`

UnsetQuery ensures that no value is present for Query, not even an explicit nil
### GetDuplicationField

`func (o *CertificateCollectionsCertificateCollectionResponse) GetDuplicationField() CSSCMSCoreEnumsDuplicateSubjectType`

GetDuplicationField returns the DuplicationField field if non-nil, zero value otherwise.

### GetDuplicationFieldOk

`func (o *CertificateCollectionsCertificateCollectionResponse) GetDuplicationFieldOk() (*CSSCMSCoreEnumsDuplicateSubjectType, bool)`

GetDuplicationFieldOk returns a tuple with the DuplicationField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicationField

`func (o *CertificateCollectionsCertificateCollectionResponse) SetDuplicationField(v CSSCMSCoreEnumsDuplicateSubjectType)`

SetDuplicationField sets DuplicationField field to given value.

### HasDuplicationField

`func (o *CertificateCollectionsCertificateCollectionResponse) HasDuplicationField() bool`

HasDuplicationField returns a boolean if a field has been set.

### GetShowOnDashboard

`func (o *CertificateCollectionsCertificateCollectionResponse) GetShowOnDashboard() bool`

GetShowOnDashboard returns the ShowOnDashboard field if non-nil, zero value otherwise.

### GetShowOnDashboardOk

`func (o *CertificateCollectionsCertificateCollectionResponse) GetShowOnDashboardOk() (*bool, bool)`

GetShowOnDashboardOk returns a tuple with the ShowOnDashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowOnDashboard

`func (o *CertificateCollectionsCertificateCollectionResponse) SetShowOnDashboard(v bool)`

SetShowOnDashboard sets ShowOnDashboard field to given value.

### HasShowOnDashboard

`func (o *CertificateCollectionsCertificateCollectionResponse) HasShowOnDashboard() bool`

HasShowOnDashboard returns a boolean if a field has been set.

### GetFavorite

`func (o *CertificateCollectionsCertificateCollectionResponse) GetFavorite() bool`

GetFavorite returns the Favorite field if non-nil, zero value otherwise.

### GetFavoriteOk

`func (o *CertificateCollectionsCertificateCollectionResponse) GetFavoriteOk() (*bool, bool)`

GetFavoriteOk returns a tuple with the Favorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorite

`func (o *CertificateCollectionsCertificateCollectionResponse) SetFavorite(v bool)`

SetFavorite sets Favorite field to given value.

### HasFavorite

`func (o *CertificateCollectionsCertificateCollectionResponse) HasFavorite() bool`

HasFavorite returns a boolean if a field has been set.

### GetEstimatedCertCount

`func (o *CertificateCollectionsCertificateCollectionResponse) GetEstimatedCertCount() int32`

GetEstimatedCertCount returns the EstimatedCertCount field if non-nil, zero value otherwise.

### GetEstimatedCertCountOk

`func (o *CertificateCollectionsCertificateCollectionResponse) GetEstimatedCertCountOk() (*int32, bool)`

GetEstimatedCertCountOk returns a tuple with the EstimatedCertCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedCertCount

`func (o *CertificateCollectionsCertificateCollectionResponse) SetEstimatedCertCount(v int32)`

SetEstimatedCertCount sets EstimatedCertCount field to given value.

### HasEstimatedCertCount

`func (o *CertificateCollectionsCertificateCollectionResponse) HasEstimatedCertCount() bool`

HasEstimatedCertCount returns a boolean if a field has been set.

### GetLastEstimated

`func (o *CertificateCollectionsCertificateCollectionResponse) GetLastEstimated() time.Time`

GetLastEstimated returns the LastEstimated field if non-nil, zero value otherwise.

### GetLastEstimatedOk

`func (o *CertificateCollectionsCertificateCollectionResponse) GetLastEstimatedOk() (*time.Time, bool)`

GetLastEstimatedOk returns a tuple with the LastEstimated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEstimated

`func (o *CertificateCollectionsCertificateCollectionResponse) SetLastEstimated(v time.Time)`

SetLastEstimated sets LastEstimated field to given value.

### HasLastEstimated

`func (o *CertificateCollectionsCertificateCollectionResponse) HasLastEstimated() bool`

HasLastEstimated returns a boolean if a field has been set.

### SetLastEstimatedNil

`func (o *CertificateCollectionsCertificateCollectionResponse) SetLastEstimatedNil(b bool)`

 SetLastEstimatedNil sets the value for LastEstimated to be an explicit nil

### UnsetLastEstimated
`func (o *CertificateCollectionsCertificateCollectionResponse) UnsetLastEstimated()`

UnsetLastEstimated ensures that no value is present for LastEstimated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



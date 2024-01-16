# ModelsRevokeAllCertificatesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | Pointer to **string** | The query string of the certificates to revoke | [optional] 
**Reason** | **int64** | The Reason for revocation | 
**Comment** | **string** | A comment for auditing purposes | 
**EffectiveDate** | Pointer to **time.Time** | The date when the certificates are to appear on the revocation list | [optional] 
**IncludeRevoked** | Pointer to **bool** | A flag telling the query to include revoked certificates | [optional] 
**IncludeExpired** | Pointer to **bool** | A flag telling the query to include expired certificates | [optional] 

## Methods

### NewModelsRevokeAllCertificatesRequest

`func NewModelsRevokeAllCertificatesRequest(reason int64, comment string, ) *ModelsRevokeAllCertificatesRequest`

NewModelsRevokeAllCertificatesRequest instantiates a new ModelsRevokeAllCertificatesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsRevokeAllCertificatesRequestWithDefaults

`func NewModelsRevokeAllCertificatesRequestWithDefaults() *ModelsRevokeAllCertificatesRequest`

NewModelsRevokeAllCertificatesRequestWithDefaults instantiates a new ModelsRevokeAllCertificatesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *ModelsRevokeAllCertificatesRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ModelsRevokeAllCertificatesRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ModelsRevokeAllCertificatesRequest) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *ModelsRevokeAllCertificatesRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetReason

`func (o *ModelsRevokeAllCertificatesRequest) GetReason() int64`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ModelsRevokeAllCertificatesRequest) GetReasonOk() (*int64, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ModelsRevokeAllCertificatesRequest) SetReason(v int64)`

SetReason sets Reason field to given value.


### GetComment

`func (o *ModelsRevokeAllCertificatesRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ModelsRevokeAllCertificatesRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ModelsRevokeAllCertificatesRequest) SetComment(v string)`

SetComment sets Comment field to given value.


### GetEffectiveDate

`func (o *ModelsRevokeAllCertificatesRequest) GetEffectiveDate() time.Time`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *ModelsRevokeAllCertificatesRequest) GetEffectiveDateOk() (*time.Time, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *ModelsRevokeAllCertificatesRequest) SetEffectiveDate(v time.Time)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *ModelsRevokeAllCertificatesRequest) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### GetIncludeRevoked

`func (o *ModelsRevokeAllCertificatesRequest) GetIncludeRevoked() bool`

GetIncludeRevoked returns the IncludeRevoked field if non-nil, zero value otherwise.

### GetIncludeRevokedOk

`func (o *ModelsRevokeAllCertificatesRequest) GetIncludeRevokedOk() (*bool, bool)`

GetIncludeRevokedOk returns a tuple with the IncludeRevoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRevoked

`func (o *ModelsRevokeAllCertificatesRequest) SetIncludeRevoked(v bool)`

SetIncludeRevoked sets IncludeRevoked field to given value.

### HasIncludeRevoked

`func (o *ModelsRevokeAllCertificatesRequest) HasIncludeRevoked() bool`

HasIncludeRevoked returns a boolean if a field has been set.

### GetIncludeExpired

`func (o *ModelsRevokeAllCertificatesRequest) GetIncludeExpired() bool`

GetIncludeExpired returns the IncludeExpired field if non-nil, zero value otherwise.

### GetIncludeExpiredOk

`func (o *ModelsRevokeAllCertificatesRequest) GetIncludeExpiredOk() (*bool, bool)`

GetIncludeExpiredOk returns a tuple with the IncludeExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeExpired

`func (o *ModelsRevokeAllCertificatesRequest) SetIncludeExpired(v bool)`

SetIncludeExpired sets IncludeExpired field to given value.

### HasIncludeExpired

`func (o *ModelsRevokeAllCertificatesRequest) HasIncludeExpired() bool`

HasIncludeExpired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



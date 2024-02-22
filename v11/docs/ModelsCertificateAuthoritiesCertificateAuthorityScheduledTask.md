# ModelsCertificateAuthoritiesCertificateAuthorityScheduledTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**CertificateAuthorityId** | Pointer to **int32** |  | [optional] 
**Schedule** | Pointer to **NullableString** |  | [optional] 
**ScanType** | Pointer to **int32** |  | [optional] 
**SyncEnabled** | Pointer to **bool** |  | [optional] 
**LastScanTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewModelsCertificateAuthoritiesCertificateAuthorityScheduledTask

`func NewModelsCertificateAuthoritiesCertificateAuthorityScheduledTask() *ModelsCertificateAuthoritiesCertificateAuthorityScheduledTask`

NewModelsCertificateAuthoritiesCertificateAuthorityScheduledTask instantiates a new ModelsCertificateAuthoritiesCertificateAuthorityScheduledTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsCertificateAuthoritiesCertificateAuthorityScheduledTaskWithDefaults

`func NewModelsCertificateAuthoritiesCertificateAuthorityScheduledTaskWithDefaults() *ModelsCertificateAuthoritiesCertificateAuthorityScheduledTask`

NewModelsCertificateAuthoritiesCertificateAuthorityScheduledTaskWithDefaults instantiates a new ModelsCertificateAuthoritiesCertificateAuthorityScheduledTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCertificateAuthorityId

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) GetCertificateAuthorityId() int32`

GetCertificateAuthorityId returns the CertificateAuthorityId field if non-nil, zero value otherwise.

### GetCertificateAuthorityIdOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) GetCertificateAuthorityIdOk() (*int32, bool)`

GetCertificateAuthorityIdOk returns a tuple with the CertificateAuthorityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthorityId

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) SetCertificateAuthorityId(v int32)`

SetCertificateAuthorityId sets CertificateAuthorityId field to given value.

### HasCertificateAuthorityId

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) HasCertificateAuthorityId() bool`

HasCertificateAuthorityId returns a boolean if a field has been set.

### GetSchedule

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### SetScheduleNil

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) SetScheduleNil(b bool)`

 SetScheduleNil sets the value for Schedule to be an explicit nil

### UnsetSchedule
`func (o *ModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) UnsetSchedule()`

UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
### GetScanType

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) GetScanType() int32`

GetScanType returns the ScanType field if non-nil, zero value otherwise.

### GetScanTypeOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) GetScanTypeOk() (*int32, bool)`

GetScanTypeOk returns a tuple with the ScanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanType

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) SetScanType(v int32)`

SetScanType sets ScanType field to given value.

### HasScanType

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) HasScanType() bool`

HasScanType returns a boolean if a field has been set.

### GetSyncEnabled

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) GetSyncEnabled() bool`

GetSyncEnabled returns the SyncEnabled field if non-nil, zero value otherwise.

### GetSyncEnabledOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) GetSyncEnabledOk() (*bool, bool)`

GetSyncEnabledOk returns a tuple with the SyncEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncEnabled

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) SetSyncEnabled(v bool)`

SetSyncEnabled sets SyncEnabled field to given value.

### HasSyncEnabled

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) HasSyncEnabled() bool`

HasSyncEnabled returns a boolean if a field has been set.

### GetLastScanTime

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) GetLastScanTime() time.Time`

GetLastScanTime returns the LastScanTime field if non-nil, zero value otherwise.

### GetLastScanTimeOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) GetLastScanTimeOk() (*time.Time, bool)`

GetLastScanTimeOk returns a tuple with the LastScanTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastScanTime

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) SetLastScanTime(v time.Time)`

SetLastScanTime sets LastScanTime field to given value.

### HasLastScanTime

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) HasLastScanTime() bool`

HasLastScanTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



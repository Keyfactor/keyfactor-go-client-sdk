# CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityScheduledTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**CertificateAuthorityId** | Pointer to **int32** |  | [optional] 
**Schedule** | Pointer to **NullableString** |  | [optional] 
**ScanType** | Pointer to [**KeyfactorPlatformExtensionsCASyncType**](KeyfactorPlatformExtensionsCASyncType.md) |  | [optional] 
**SyncEnabled** | Pointer to **bool** |  | [optional] 
**LastScanTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityScheduledTask

`func NewCSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityScheduledTask() *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityScheduledTask`

NewCSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityScheduledTask instantiates a new CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityScheduledTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityScheduledTaskWithDefaults

`func NewCSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityScheduledTaskWithDefaults() *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityScheduledTask`

NewCSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityScheduledTaskWithDefaults instantiates a new CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityScheduledTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCertificateAuthorityId

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) GetCertificateAuthorityId() int32`

GetCertificateAuthorityId returns the CertificateAuthorityId field if non-nil, zero value otherwise.

### GetCertificateAuthorityIdOk

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) GetCertificateAuthorityIdOk() (*int32, bool)`

GetCertificateAuthorityIdOk returns a tuple with the CertificateAuthorityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthorityId

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) SetCertificateAuthorityId(v int32)`

SetCertificateAuthorityId sets CertificateAuthorityId field to given value.

### HasCertificateAuthorityId

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) HasCertificateAuthorityId() bool`

HasCertificateAuthorityId returns a boolean if a field has been set.

### GetSchedule

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### SetScheduleNil

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) SetScheduleNil(b bool)`

 SetScheduleNil sets the value for Schedule to be an explicit nil

### UnsetSchedule
`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) UnsetSchedule()`

UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
### GetScanType

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) GetScanType() KeyfactorPlatformExtensionsCASyncType`

GetScanType returns the ScanType field if non-nil, zero value otherwise.

### GetScanTypeOk

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) GetScanTypeOk() (*KeyfactorPlatformExtensionsCASyncType, bool)`

GetScanTypeOk returns a tuple with the ScanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanType

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) SetScanType(v KeyfactorPlatformExtensionsCASyncType)`

SetScanType sets ScanType field to given value.

### HasScanType

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) HasScanType() bool`

HasScanType returns a boolean if a field has been set.

### GetSyncEnabled

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) GetSyncEnabled() bool`

GetSyncEnabled returns the SyncEnabled field if non-nil, zero value otherwise.

### GetSyncEnabledOk

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) GetSyncEnabledOk() (*bool, bool)`

GetSyncEnabledOk returns a tuple with the SyncEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncEnabled

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) SetSyncEnabled(v bool)`

SetSyncEnabled sets SyncEnabled field to given value.

### HasSyncEnabled

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) HasSyncEnabled() bool`

HasSyncEnabled returns a boolean if a field has been set.

### GetLastScanTime

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) GetLastScanTime() time.Time`

GetLastScanTime returns the LastScanTime field if non-nil, zero value otherwise.

### GetLastScanTimeOk

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) GetLastScanTimeOk() (*time.Time, bool)`

GetLastScanTimeOk returns a tuple with the LastScanTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastScanTime

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) SetLastScanTime(v time.Time)`

SetLastScanTime sets LastScanTime field to given value.

### HasLastScanTime

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityScheduledTask) HasLastScanTime() bool`

HasLastScanTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



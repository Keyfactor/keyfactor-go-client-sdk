# SslQuietHourResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDay** | Pointer to [**SystemDayOfWeek**](SystemDayOfWeek.md) |  | [optional] 
**StartTime** | Pointer to **time.Time** |  | [optional] 
**EndDay** | Pointer to [**SystemDayOfWeek**](SystemDayOfWeek.md) |  | [optional] 
**EndTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSslQuietHourResponse

`func NewSslQuietHourResponse() *SslQuietHourResponse`

NewSslQuietHourResponse instantiates a new SslQuietHourResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSslQuietHourResponseWithDefaults

`func NewSslQuietHourResponseWithDefaults() *SslQuietHourResponse`

NewSslQuietHourResponseWithDefaults instantiates a new SslQuietHourResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDay

`func (o *SslQuietHourResponse) GetStartDay() SystemDayOfWeek`

GetStartDay returns the StartDay field if non-nil, zero value otherwise.

### GetStartDayOk

`func (o *SslQuietHourResponse) GetStartDayOk() (*SystemDayOfWeek, bool)`

GetStartDayOk returns a tuple with the StartDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDay

`func (o *SslQuietHourResponse) SetStartDay(v SystemDayOfWeek)`

SetStartDay sets StartDay field to given value.

### HasStartDay

`func (o *SslQuietHourResponse) HasStartDay() bool`

HasStartDay returns a boolean if a field has been set.

### GetStartTime

`func (o *SslQuietHourResponse) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *SslQuietHourResponse) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *SslQuietHourResponse) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *SslQuietHourResponse) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndDay

`func (o *SslQuietHourResponse) GetEndDay() SystemDayOfWeek`

GetEndDay returns the EndDay field if non-nil, zero value otherwise.

### GetEndDayOk

`func (o *SslQuietHourResponse) GetEndDayOk() (*SystemDayOfWeek, bool)`

GetEndDayOk returns a tuple with the EndDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDay

`func (o *SslQuietHourResponse) SetEndDay(v SystemDayOfWeek)`

SetEndDay sets EndDay field to given value.

### HasEndDay

`func (o *SslQuietHourResponse) HasEndDay() bool`

HasEndDay returns a boolean if a field has been set.

### GetEndTime

`func (o *SslQuietHourResponse) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *SslQuietHourResponse) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *SslQuietHourResponse) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *SslQuietHourResponse) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



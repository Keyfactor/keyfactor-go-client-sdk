# AnalyticsAnalyticsCertificateCountWithPeriodResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EvalDate** | Pointer to **time.Time** |  | [optional] 
**Periods** | Pointer to [**[]AnalyticsAnalyticsPeriodCount**](AnalyticsAnalyticsPeriodCount.md) |  | [optional] 

## Methods

### NewAnalyticsAnalyticsCertificateCountWithPeriodResponse

`func NewAnalyticsAnalyticsCertificateCountWithPeriodResponse() *AnalyticsAnalyticsCertificateCountWithPeriodResponse`

NewAnalyticsAnalyticsCertificateCountWithPeriodResponse instantiates a new AnalyticsAnalyticsCertificateCountWithPeriodResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsAnalyticsCertificateCountWithPeriodResponseWithDefaults

`func NewAnalyticsAnalyticsCertificateCountWithPeriodResponseWithDefaults() *AnalyticsAnalyticsCertificateCountWithPeriodResponse`

NewAnalyticsAnalyticsCertificateCountWithPeriodResponseWithDefaults instantiates a new AnalyticsAnalyticsCertificateCountWithPeriodResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvalDate

`func (o *AnalyticsAnalyticsCertificateCountWithPeriodResponse) GetEvalDate() time.Time`

GetEvalDate returns the EvalDate field if non-nil, zero value otherwise.

### GetEvalDateOk

`func (o *AnalyticsAnalyticsCertificateCountWithPeriodResponse) GetEvalDateOk() (*time.Time, bool)`

GetEvalDateOk returns a tuple with the EvalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvalDate

`func (o *AnalyticsAnalyticsCertificateCountWithPeriodResponse) SetEvalDate(v time.Time)`

SetEvalDate sets EvalDate field to given value.

### HasEvalDate

`func (o *AnalyticsAnalyticsCertificateCountWithPeriodResponse) HasEvalDate() bool`

HasEvalDate returns a boolean if a field has been set.

### GetPeriods

`func (o *AnalyticsAnalyticsCertificateCountWithPeriodResponse) GetPeriods() []AnalyticsAnalyticsPeriodCount`

GetPeriods returns the Periods field if non-nil, zero value otherwise.

### GetPeriodsOk

`func (o *AnalyticsAnalyticsCertificateCountWithPeriodResponse) GetPeriodsOk() (*[]AnalyticsAnalyticsPeriodCount, bool)`

GetPeriodsOk returns a tuple with the Periods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriods

`func (o *AnalyticsAnalyticsCertificateCountWithPeriodResponse) SetPeriods(v []AnalyticsAnalyticsPeriodCount)`

SetPeriods sets Periods field to given value.

### HasPeriods

`func (o *AnalyticsAnalyticsCertificateCountWithPeriodResponse) HasPeriods() bool`

HasPeriods returns a boolean if a field has been set.

### SetPeriodsNil

`func (o *AnalyticsAnalyticsCertificateCountWithPeriodResponse) SetPeriodsNil(b bool)`

 SetPeriodsNil sets the value for Periods to be an explicit nil

### UnsetPeriods
`func (o *AnalyticsAnalyticsCertificateCountWithPeriodResponse) UnsetPeriods()`

UnsetPeriods ensures that no value is present for Periods, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



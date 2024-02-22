# KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | **string** |  | 
**EndpointType** | **string** |  | 
**Location** | **string** |  | 
**Email** | Pointer to [**KeyfactorApiModelsMonitoringEmailRequest**](KeyfactorApiModelsMonitoringEmailRequest.md) |  | [optional] 
**Dashboard** | [**KeyfactorApiModelsMonitoringDashboardRequest**](KeyfactorApiModelsMonitoringDashboardRequest.md) |  | 
**Schedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**OCSPParameters** | Pointer to [**KeyfactorApiModelsMonitoringOCSPParametersRequest**](KeyfactorApiModelsMonitoringOCSPParametersRequest.md) |  | [optional] 

## Methods

### NewKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest

`func NewKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest(name string, endpointType string, location string, dashboard KeyfactorApiModelsMonitoringDashboardRequest, ) *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest`

NewKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest instantiates a new KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequestWithDefaults

`func NewKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequestWithDefaults() *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest`

NewKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequestWithDefaults instantiates a new KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEndpointType

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetEndpointType() string`

GetEndpointType returns the EndpointType field if non-nil, zero value otherwise.

### GetEndpointTypeOk

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetEndpointTypeOk() (*string, bool)`

GetEndpointTypeOk returns a tuple with the EndpointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointType

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) SetEndpointType(v string)`

SetEndpointType sets EndpointType field to given value.


### GetLocation

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetEmail

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetEmail() KeyfactorApiModelsMonitoringEmailRequest`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetEmailOk() (*KeyfactorApiModelsMonitoringEmailRequest, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) SetEmail(v KeyfactorApiModelsMonitoringEmailRequest)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetDashboard

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetDashboard() KeyfactorApiModelsMonitoringDashboardRequest`

GetDashboard returns the Dashboard field if non-nil, zero value otherwise.

### GetDashboardOk

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetDashboardOk() (*KeyfactorApiModelsMonitoringDashboardRequest, bool)`

GetDashboardOk returns a tuple with the Dashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboard

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) SetDashboard(v KeyfactorApiModelsMonitoringDashboardRequest)`

SetDashboard sets Dashboard field to given value.


### GetSchedule

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) SetSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetOCSPParameters

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetOCSPParameters() KeyfactorApiModelsMonitoringOCSPParametersRequest`

GetOCSPParameters returns the OCSPParameters field if non-nil, zero value otherwise.

### GetOCSPParametersOk

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetOCSPParametersOk() (*KeyfactorApiModelsMonitoringOCSPParametersRequest, bool)`

GetOCSPParametersOk returns a tuple with the OCSPParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOCSPParameters

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) SetOCSPParameters(v KeyfactorApiModelsMonitoringOCSPParametersRequest)`

SetOCSPParameters sets OCSPParameters field to given value.

### HasOCSPParameters

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) HasOCSPParameters() bool`

HasOCSPParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



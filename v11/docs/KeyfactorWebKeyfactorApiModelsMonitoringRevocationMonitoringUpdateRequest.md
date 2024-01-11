# KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**EndpointType** | **string** |  | 
**Location** | **string** |  | 
**Email** | Pointer to [**KeyfactorWebKeyfactorApiModelsMonitoringEmailRequest**](KeyfactorWebKeyfactorApiModelsMonitoringEmailRequest.md) |  | [optional] 
**Dashboard** | [**KeyfactorWebKeyfactorApiModelsMonitoringDashboardRequest**](KeyfactorWebKeyfactorApiModelsMonitoringDashboardRequest.md) |  | 
**Schedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**OcspParameters** | Pointer to [**KeyfactorWebKeyfactorApiModelsMonitoringOCSPParametersRequest**](KeyfactorWebKeyfactorApiModelsMonitoringOCSPParametersRequest.md) |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest

`func NewKeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest(name string, endpointType string, location string, dashboard KeyfactorWebKeyfactorApiModelsMonitoringDashboardRequest, ) *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest`

NewKeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest instantiates a new KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequestWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequestWithDefaults() *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest`

NewKeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequestWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEndpointType

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetEndpointType() string`

GetEndpointType returns the EndpointType field if non-nil, zero value otherwise.

### GetEndpointTypeOk

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetEndpointTypeOk() (*string, bool)`

GetEndpointTypeOk returns a tuple with the EndpointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointType

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) SetEndpointType(v string)`

SetEndpointType sets EndpointType field to given value.


### GetLocation

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetEmail

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetEmail() KeyfactorWebKeyfactorApiModelsMonitoringEmailRequest`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetEmailOk() (*KeyfactorWebKeyfactorApiModelsMonitoringEmailRequest, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) SetEmail(v KeyfactorWebKeyfactorApiModelsMonitoringEmailRequest)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetDashboard

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetDashboard() KeyfactorWebKeyfactorApiModelsMonitoringDashboardRequest`

GetDashboard returns the Dashboard field if non-nil, zero value otherwise.

### GetDashboardOk

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetDashboardOk() (*KeyfactorWebKeyfactorApiModelsMonitoringDashboardRequest, bool)`

GetDashboardOk returns a tuple with the Dashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboard

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) SetDashboard(v KeyfactorWebKeyfactorApiModelsMonitoringDashboardRequest)`

SetDashboard sets Dashboard field to given value.


### GetSchedule

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) SetSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetOcspParameters

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetOcspParameters() KeyfactorWebKeyfactorApiModelsMonitoringOCSPParametersRequest`

GetOcspParameters returns the OcspParameters field if non-nil, zero value otherwise.

### GetOcspParametersOk

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) GetOcspParametersOk() (*KeyfactorWebKeyfactorApiModelsMonitoringOCSPParametersRequest, bool)`

GetOcspParametersOk returns a tuple with the OcspParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspParameters

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) SetOcspParameters(v KeyfactorWebKeyfactorApiModelsMonitoringOCSPParametersRequest)`

SetOcspParameters sets OcspParameters field to given value.

### HasOcspParameters

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest) HasOcspParameters() bool`

HasOcspParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



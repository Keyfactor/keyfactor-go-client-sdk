# KeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**EndpointType** | **string** |  | 
**Location** | **string** |  | 
**Email** | Pointer to [**KeyfactorApiModelsMonitoringEmailRequest**](KeyfactorApiModelsMonitoringEmailRequest.md) |  | [optional] 
**Dashboard** | [**KeyfactorApiModelsMonitoringDashboardRequest**](KeyfactorApiModelsMonitoringDashboardRequest.md) |  | 
**Schedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**OCSPParameters** | Pointer to [**KeyfactorApiModelsMonitoringOCSPParametersRequest**](KeyfactorApiModelsMonitoringOCSPParametersRequest.md) |  | [optional] 

## Methods

### NewKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest

`func NewKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest(name string, endpointType string, location string, dashboard KeyfactorApiModelsMonitoringDashboardRequest, ) *KeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest`

NewKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest instantiates a new KeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequestWithDefaults

`func NewKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequestWithDefaults() *KeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest`

NewKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequestWithDefaults instantiates a new KeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEndpointType

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) GetEndpointType() string`

GetEndpointType returns the EndpointType field if non-nil, zero value otherwise.

### GetEndpointTypeOk

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) GetEndpointTypeOk() (*string, bool)`

GetEndpointTypeOk returns a tuple with the EndpointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointType

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) SetEndpointType(v string)`

SetEndpointType sets EndpointType field to given value.


### GetLocation

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetEmail

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) GetEmail() KeyfactorApiModelsMonitoringEmailRequest`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) GetEmailOk() (*KeyfactorApiModelsMonitoringEmailRequest, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) SetEmail(v KeyfactorApiModelsMonitoringEmailRequest)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetDashboard

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) GetDashboard() KeyfactorApiModelsMonitoringDashboardRequest`

GetDashboard returns the Dashboard field if non-nil, zero value otherwise.

### GetDashboardOk

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) GetDashboardOk() (*KeyfactorApiModelsMonitoringDashboardRequest, bool)`

GetDashboardOk returns a tuple with the Dashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboard

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) SetDashboard(v KeyfactorApiModelsMonitoringDashboardRequest)`

SetDashboard sets Dashboard field to given value.


### GetSchedule

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) GetSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) GetScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) SetSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetOCSPParameters

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) GetOCSPParameters() KeyfactorApiModelsMonitoringOCSPParametersRequest`

GetOCSPParameters returns the OCSPParameters field if non-nil, zero value otherwise.

### GetOCSPParametersOk

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) GetOCSPParametersOk() (*KeyfactorApiModelsMonitoringOCSPParametersRequest, bool)`

GetOCSPParametersOk returns a tuple with the OCSPParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOCSPParameters

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) SetOCSPParameters(v KeyfactorApiModelsMonitoringOCSPParametersRequest)`

SetOCSPParameters sets OCSPParameters field to given value.

### HasOCSPParameters

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) HasOCSPParameters() bool`

HasOCSPParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



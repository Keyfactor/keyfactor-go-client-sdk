# KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**EndpointType** | **string** |  | 
**Location** | **string** |  | 
**Email** | Pointer to [**KeyfactorWebKeyfactorApiModelsMonitoringEmailRequest**](KeyfactorWebKeyfactorApiModelsMonitoringEmailRequest.md) |  | [optional] 
**Dashboard** | [**KeyfactorWebKeyfactorApiModelsMonitoringDashboardRequest**](KeyfactorWebKeyfactorApiModelsMonitoringDashboardRequest.md) |  | 
**Schedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**OcspParameters** | Pointer to [**KeyfactorWebKeyfactorApiModelsMonitoringOCSPParametersRequest**](KeyfactorWebKeyfactorApiModelsMonitoringOCSPParametersRequest.md) |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest

`func NewKeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest(name string, endpointType string, location string, dashboard KeyfactorWebKeyfactorApiModelsMonitoringDashboardRequest, ) *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest`

NewKeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest instantiates a new KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequestWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequestWithDefaults() *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest`

NewKeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequestWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEndpointType

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) GetEndpointType() string`

GetEndpointType returns the EndpointType field if non-nil, zero value otherwise.

### GetEndpointTypeOk

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) GetEndpointTypeOk() (*string, bool)`

GetEndpointTypeOk returns a tuple with the EndpointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointType

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) SetEndpointType(v string)`

SetEndpointType sets EndpointType field to given value.


### GetLocation

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetEmail

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) GetEmail() KeyfactorWebKeyfactorApiModelsMonitoringEmailRequest`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) GetEmailOk() (*KeyfactorWebKeyfactorApiModelsMonitoringEmailRequest, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) SetEmail(v KeyfactorWebKeyfactorApiModelsMonitoringEmailRequest)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetDashboard

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) GetDashboard() KeyfactorWebKeyfactorApiModelsMonitoringDashboardRequest`

GetDashboard returns the Dashboard field if non-nil, zero value otherwise.

### GetDashboardOk

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) GetDashboardOk() (*KeyfactorWebKeyfactorApiModelsMonitoringDashboardRequest, bool)`

GetDashboardOk returns a tuple with the Dashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboard

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) SetDashboard(v KeyfactorWebKeyfactorApiModelsMonitoringDashboardRequest)`

SetDashboard sets Dashboard field to given value.


### GetSchedule

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) GetSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) GetScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) SetSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetOcspParameters

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) GetOcspParameters() KeyfactorWebKeyfactorApiModelsMonitoringOCSPParametersRequest`

GetOcspParameters returns the OcspParameters field if non-nil, zero value otherwise.

### GetOcspParametersOk

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) GetOcspParametersOk() (*KeyfactorWebKeyfactorApiModelsMonitoringOCSPParametersRequest, bool)`

GetOcspParametersOk returns a tuple with the OcspParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspParameters

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) SetOcspParameters(v KeyfactorWebKeyfactorApiModelsMonitoringOCSPParametersRequest)`

SetOcspParameters sets OcspParameters field to given value.

### HasOcspParameters

`func (o *KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest) HasOcspParameters() bool`

HasOcspParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



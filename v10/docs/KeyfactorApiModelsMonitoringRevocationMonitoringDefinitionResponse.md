# KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**EndpointType** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Email** | Pointer to [**KeyfactorApiModelsMonitoringEmailResponse**](KeyfactorApiModelsMonitoringEmailResponse.md) |  | [optional] 
**Dashboard** | Pointer to [**KeyfactorApiModelsMonitoringDashboardResponse**](KeyfactorApiModelsMonitoringDashboardResponse.md) |  | [optional] 
**Schedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**OCSPParameters** | Pointer to [**KeyfactorApiModelsMonitoringOCSPParametersResponse**](KeyfactorApiModelsMonitoringOCSPParametersResponse.md) |  | [optional] 

## Methods

### NewKeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse

`func NewKeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse() *KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse`

NewKeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse instantiates a new KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponseWithDefaults

`func NewKeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponseWithDefaults() *KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse`

NewKeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponseWithDefaults instantiates a new KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEndpointType

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse) GetEndpointType() string`

GetEndpointType returns the EndpointType field if non-nil, zero value otherwise.

### GetEndpointTypeOk

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse) GetEndpointTypeOk() (*string, bool)`

GetEndpointTypeOk returns a tuple with the EndpointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointType

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse) SetEndpointType(v string)`

SetEndpointType sets EndpointType field to given value.

### HasEndpointType

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse) HasEndpointType() bool`

HasEndpointType returns a boolean if a field has been set.

### GetLocation

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetEmail

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse) GetEmail() KeyfactorApiModelsMonitoringEmailResponse`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse) GetEmailOk() (*KeyfactorApiModelsMonitoringEmailResponse, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse) SetEmail(v KeyfactorApiModelsMonitoringEmailResponse)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetDashboard

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse) GetDashboard() KeyfactorApiModelsMonitoringDashboardResponse`

GetDashboard returns the Dashboard field if non-nil, zero value otherwise.

### GetDashboardOk

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse) GetDashboardOk() (*KeyfactorApiModelsMonitoringDashboardResponse, bool)`

GetDashboardOk returns a tuple with the Dashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboard

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse) SetDashboard(v KeyfactorApiModelsMonitoringDashboardResponse)`

SetDashboard sets Dashboard field to given value.

### HasDashboard

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse) HasDashboard() bool`

HasDashboard returns a boolean if a field has been set.

### GetSchedule

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse) GetSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse) GetScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse) SetSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetOCSPParameters

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse) GetOCSPParameters() KeyfactorApiModelsMonitoringOCSPParametersResponse`

GetOCSPParameters returns the OCSPParameters field if non-nil, zero value otherwise.

### GetOCSPParametersOk

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse) GetOCSPParametersOk() (*KeyfactorApiModelsMonitoringOCSPParametersResponse, bool)`

GetOCSPParametersOk returns a tuple with the OCSPParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOCSPParameters

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse) SetOCSPParameters(v KeyfactorApiModelsMonitoringOCSPParametersResponse)`

SetOCSPParameters sets OCSPParameters field to given value.

### HasOCSPParameters

`func (o *KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse) HasOCSPParameters() bool`

HasOCSPParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



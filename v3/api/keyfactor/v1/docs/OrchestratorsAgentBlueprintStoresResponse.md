# OrchestratorsAgentBlueprintStoresResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentBlueprintStoreId** | Pointer to **string** |  | [optional] 
**AgentBlueprintId** | Pointer to **string** |  | [optional] 
**StorePath** | Pointer to **NullableString** |  | [optional] 
**ContainerId** | Pointer to **int32** |  | [optional] 
**CertStoreType** | Pointer to **int32** |  | [optional] 
**CertStoreTypeName** | Pointer to **NullableString** |  | [optional] 
**Approved** | Pointer to **bool** |  | [optional] 
**CreateIfMissing** | Pointer to **bool** |  | [optional] 
**Properties** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewOrchestratorsAgentBlueprintStoresResponse

`func NewOrchestratorsAgentBlueprintStoresResponse() *OrchestratorsAgentBlueprintStoresResponse`

NewOrchestratorsAgentBlueprintStoresResponse instantiates a new OrchestratorsAgentBlueprintStoresResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrchestratorsAgentBlueprintStoresResponseWithDefaults

`func NewOrchestratorsAgentBlueprintStoresResponseWithDefaults() *OrchestratorsAgentBlueprintStoresResponse`

NewOrchestratorsAgentBlueprintStoresResponseWithDefaults instantiates a new OrchestratorsAgentBlueprintStoresResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentBlueprintStoreId

`func (o *OrchestratorsAgentBlueprintStoresResponse) GetAgentBlueprintStoreId() string`

GetAgentBlueprintStoreId returns the AgentBlueprintStoreId field if non-nil, zero value otherwise.

### GetAgentBlueprintStoreIdOk

`func (o *OrchestratorsAgentBlueprintStoresResponse) GetAgentBlueprintStoreIdOk() (*string, bool)`

GetAgentBlueprintStoreIdOk returns a tuple with the AgentBlueprintStoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentBlueprintStoreId

`func (o *OrchestratorsAgentBlueprintStoresResponse) SetAgentBlueprintStoreId(v string)`

SetAgentBlueprintStoreId sets AgentBlueprintStoreId field to given value.

### HasAgentBlueprintStoreId

`func (o *OrchestratorsAgentBlueprintStoresResponse) HasAgentBlueprintStoreId() bool`

HasAgentBlueprintStoreId returns a boolean if a field has been set.

### GetAgentBlueprintId

`func (o *OrchestratorsAgentBlueprintStoresResponse) GetAgentBlueprintId() string`

GetAgentBlueprintId returns the AgentBlueprintId field if non-nil, zero value otherwise.

### GetAgentBlueprintIdOk

`func (o *OrchestratorsAgentBlueprintStoresResponse) GetAgentBlueprintIdOk() (*string, bool)`

GetAgentBlueprintIdOk returns a tuple with the AgentBlueprintId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentBlueprintId

`func (o *OrchestratorsAgentBlueprintStoresResponse) SetAgentBlueprintId(v string)`

SetAgentBlueprintId sets AgentBlueprintId field to given value.

### HasAgentBlueprintId

`func (o *OrchestratorsAgentBlueprintStoresResponse) HasAgentBlueprintId() bool`

HasAgentBlueprintId returns a boolean if a field has been set.

### GetStorePath

`func (o *OrchestratorsAgentBlueprintStoresResponse) GetStorePath() string`

GetStorePath returns the StorePath field if non-nil, zero value otherwise.

### GetStorePathOk

`func (o *OrchestratorsAgentBlueprintStoresResponse) GetStorePathOk() (*string, bool)`

GetStorePathOk returns a tuple with the StorePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorePath

`func (o *OrchestratorsAgentBlueprintStoresResponse) SetStorePath(v string)`

SetStorePath sets StorePath field to given value.

### HasStorePath

`func (o *OrchestratorsAgentBlueprintStoresResponse) HasStorePath() bool`

HasStorePath returns a boolean if a field has been set.

### SetStorePathNil

`func (o *OrchestratorsAgentBlueprintStoresResponse) SetStorePathNil(b bool)`

 SetStorePathNil sets the value for StorePath to be an explicit nil

### UnsetStorePath
`func (o *OrchestratorsAgentBlueprintStoresResponse) UnsetStorePath()`

UnsetStorePath ensures that no value is present for StorePath, not even an explicit nil
### GetContainerId

`func (o *OrchestratorsAgentBlueprintStoresResponse) GetContainerId() int32`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *OrchestratorsAgentBlueprintStoresResponse) GetContainerIdOk() (*int32, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *OrchestratorsAgentBlueprintStoresResponse) SetContainerId(v int32)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *OrchestratorsAgentBlueprintStoresResponse) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### GetCertStoreType

`func (o *OrchestratorsAgentBlueprintStoresResponse) GetCertStoreType() int32`

GetCertStoreType returns the CertStoreType field if non-nil, zero value otherwise.

### GetCertStoreTypeOk

`func (o *OrchestratorsAgentBlueprintStoresResponse) GetCertStoreTypeOk() (*int32, bool)`

GetCertStoreTypeOk returns a tuple with the CertStoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertStoreType

`func (o *OrchestratorsAgentBlueprintStoresResponse) SetCertStoreType(v int32)`

SetCertStoreType sets CertStoreType field to given value.

### HasCertStoreType

`func (o *OrchestratorsAgentBlueprintStoresResponse) HasCertStoreType() bool`

HasCertStoreType returns a boolean if a field has been set.

### GetCertStoreTypeName

`func (o *OrchestratorsAgentBlueprintStoresResponse) GetCertStoreTypeName() string`

GetCertStoreTypeName returns the CertStoreTypeName field if non-nil, zero value otherwise.

### GetCertStoreTypeNameOk

`func (o *OrchestratorsAgentBlueprintStoresResponse) GetCertStoreTypeNameOk() (*string, bool)`

GetCertStoreTypeNameOk returns a tuple with the CertStoreTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertStoreTypeName

`func (o *OrchestratorsAgentBlueprintStoresResponse) SetCertStoreTypeName(v string)`

SetCertStoreTypeName sets CertStoreTypeName field to given value.

### HasCertStoreTypeName

`func (o *OrchestratorsAgentBlueprintStoresResponse) HasCertStoreTypeName() bool`

HasCertStoreTypeName returns a boolean if a field has been set.

### SetCertStoreTypeNameNil

`func (o *OrchestratorsAgentBlueprintStoresResponse) SetCertStoreTypeNameNil(b bool)`

 SetCertStoreTypeNameNil sets the value for CertStoreTypeName to be an explicit nil

### UnsetCertStoreTypeName
`func (o *OrchestratorsAgentBlueprintStoresResponse) UnsetCertStoreTypeName()`

UnsetCertStoreTypeName ensures that no value is present for CertStoreTypeName, not even an explicit nil
### GetApproved

`func (o *OrchestratorsAgentBlueprintStoresResponse) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *OrchestratorsAgentBlueprintStoresResponse) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *OrchestratorsAgentBlueprintStoresResponse) SetApproved(v bool)`

SetApproved sets Approved field to given value.

### HasApproved

`func (o *OrchestratorsAgentBlueprintStoresResponse) HasApproved() bool`

HasApproved returns a boolean if a field has been set.

### GetCreateIfMissing

`func (o *OrchestratorsAgentBlueprintStoresResponse) GetCreateIfMissing() bool`

GetCreateIfMissing returns the CreateIfMissing field if non-nil, zero value otherwise.

### GetCreateIfMissingOk

`func (o *OrchestratorsAgentBlueprintStoresResponse) GetCreateIfMissingOk() (*bool, bool)`

GetCreateIfMissingOk returns a tuple with the CreateIfMissing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIfMissing

`func (o *OrchestratorsAgentBlueprintStoresResponse) SetCreateIfMissing(v bool)`

SetCreateIfMissing sets CreateIfMissing field to given value.

### HasCreateIfMissing

`func (o *OrchestratorsAgentBlueprintStoresResponse) HasCreateIfMissing() bool`

HasCreateIfMissing returns a boolean if a field has been set.

### GetProperties

`func (o *OrchestratorsAgentBlueprintStoresResponse) GetProperties() string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *OrchestratorsAgentBlueprintStoresResponse) GetPropertiesOk() (*string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *OrchestratorsAgentBlueprintStoresResponse) SetProperties(v string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *OrchestratorsAgentBlueprintStoresResponse) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *OrchestratorsAgentBlueprintStoresResponse) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *OrchestratorsAgentBlueprintStoresResponse) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



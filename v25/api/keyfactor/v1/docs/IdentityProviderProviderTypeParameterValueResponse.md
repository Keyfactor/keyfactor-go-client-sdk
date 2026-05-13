# IdentityProviderProviderTypeParameterValueResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**DataType** | Pointer to [**CSSCMSDataModelEnumsIdentityProviderDataType**](CSSCMSDataModelEnumsIdentityProviderDataType.md) |  | [optional] 
**Value** | Pointer to **NullableString** |  | [optional] 
**RequestHeaders** | Pointer to [**[]SharedRequestHeaderResponse**](SharedRequestHeaderResponse.md) |  | [optional] 
**RequestURLParameters** | Pointer to [**[]IdentityProviderRequestURLParameterResponse**](IdentityProviderRequestURLParameterResponse.md) |  | [optional] 
**SecretValue** | Pointer to [**CSSCMSDataModelModelsKeyfactorAPISecret**](CSSCMSDataModelModelsKeyfactorAPISecret.md) |  | [optional] 

## Methods

### NewIdentityProviderProviderTypeParameterValueResponse

`func NewIdentityProviderProviderTypeParameterValueResponse() *IdentityProviderProviderTypeParameterValueResponse`

NewIdentityProviderProviderTypeParameterValueResponse instantiates a new IdentityProviderProviderTypeParameterValueResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderProviderTypeParameterValueResponseWithDefaults

`func NewIdentityProviderProviderTypeParameterValueResponseWithDefaults() *IdentityProviderProviderTypeParameterValueResponse`

NewIdentityProviderProviderTypeParameterValueResponseWithDefaults instantiates a new IdentityProviderProviderTypeParameterValueResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IdentityProviderProviderTypeParameterValueResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityProviderProviderTypeParameterValueResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityProviderProviderTypeParameterValueResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityProviderProviderTypeParameterValueResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IdentityProviderProviderTypeParameterValueResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityProviderProviderTypeParameterValueResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityProviderProviderTypeParameterValueResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdentityProviderProviderTypeParameterValueResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *IdentityProviderProviderTypeParameterValueResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *IdentityProviderProviderTypeParameterValueResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDisplayName

`func (o *IdentityProviderProviderTypeParameterValueResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IdentityProviderProviderTypeParameterValueResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IdentityProviderProviderTypeParameterValueResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IdentityProviderProviderTypeParameterValueResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *IdentityProviderProviderTypeParameterValueResponse) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *IdentityProviderProviderTypeParameterValueResponse) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetRequired

`func (o *IdentityProviderProviderTypeParameterValueResponse) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *IdentityProviderProviderTypeParameterValueResponse) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *IdentityProviderProviderTypeParameterValueResponse) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *IdentityProviderProviderTypeParameterValueResponse) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetDataType

`func (o *IdentityProviderProviderTypeParameterValueResponse) GetDataType() CSSCMSDataModelEnumsIdentityProviderDataType`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *IdentityProviderProviderTypeParameterValueResponse) GetDataTypeOk() (*CSSCMSDataModelEnumsIdentityProviderDataType, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *IdentityProviderProviderTypeParameterValueResponse) SetDataType(v CSSCMSDataModelEnumsIdentityProviderDataType)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *IdentityProviderProviderTypeParameterValueResponse) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetValue

`func (o *IdentityProviderProviderTypeParameterValueResponse) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IdentityProviderProviderTypeParameterValueResponse) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IdentityProviderProviderTypeParameterValueResponse) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *IdentityProviderProviderTypeParameterValueResponse) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *IdentityProviderProviderTypeParameterValueResponse) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *IdentityProviderProviderTypeParameterValueResponse) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetRequestHeaders

`func (o *IdentityProviderProviderTypeParameterValueResponse) GetRequestHeaders() []SharedRequestHeaderResponse`

GetRequestHeaders returns the RequestHeaders field if non-nil, zero value otherwise.

### GetRequestHeadersOk

`func (o *IdentityProviderProviderTypeParameterValueResponse) GetRequestHeadersOk() (*[]SharedRequestHeaderResponse, bool)`

GetRequestHeadersOk returns a tuple with the RequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestHeaders

`func (o *IdentityProviderProviderTypeParameterValueResponse) SetRequestHeaders(v []SharedRequestHeaderResponse)`

SetRequestHeaders sets RequestHeaders field to given value.

### HasRequestHeaders

`func (o *IdentityProviderProviderTypeParameterValueResponse) HasRequestHeaders() bool`

HasRequestHeaders returns a boolean if a field has been set.

### SetRequestHeadersNil

`func (o *IdentityProviderProviderTypeParameterValueResponse) SetRequestHeadersNil(b bool)`

 SetRequestHeadersNil sets the value for RequestHeaders to be an explicit nil

### UnsetRequestHeaders
`func (o *IdentityProviderProviderTypeParameterValueResponse) UnsetRequestHeaders()`

UnsetRequestHeaders ensures that no value is present for RequestHeaders, not even an explicit nil
### GetRequestURLParameters

`func (o *IdentityProviderProviderTypeParameterValueResponse) GetRequestURLParameters() []IdentityProviderRequestURLParameterResponse`

GetRequestURLParameters returns the RequestURLParameters field if non-nil, zero value otherwise.

### GetRequestURLParametersOk

`func (o *IdentityProviderProviderTypeParameterValueResponse) GetRequestURLParametersOk() (*[]IdentityProviderRequestURLParameterResponse, bool)`

GetRequestURLParametersOk returns a tuple with the RequestURLParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestURLParameters

`func (o *IdentityProviderProviderTypeParameterValueResponse) SetRequestURLParameters(v []IdentityProviderRequestURLParameterResponse)`

SetRequestURLParameters sets RequestURLParameters field to given value.

### HasRequestURLParameters

`func (o *IdentityProviderProviderTypeParameterValueResponse) HasRequestURLParameters() bool`

HasRequestURLParameters returns a boolean if a field has been set.

### SetRequestURLParametersNil

`func (o *IdentityProviderProviderTypeParameterValueResponse) SetRequestURLParametersNil(b bool)`

 SetRequestURLParametersNil sets the value for RequestURLParameters to be an explicit nil

### UnsetRequestURLParameters
`func (o *IdentityProviderProviderTypeParameterValueResponse) UnsetRequestURLParameters()`

UnsetRequestURLParameters ensures that no value is present for RequestURLParameters, not even an explicit nil
### GetSecretValue

`func (o *IdentityProviderProviderTypeParameterValueResponse) GetSecretValue() CSSCMSDataModelModelsKeyfactorAPISecret`

GetSecretValue returns the SecretValue field if non-nil, zero value otherwise.

### GetSecretValueOk

`func (o *IdentityProviderProviderTypeParameterValueResponse) GetSecretValueOk() (*CSSCMSDataModelModelsKeyfactorAPISecret, bool)`

GetSecretValueOk returns a tuple with the SecretValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretValue

`func (o *IdentityProviderProviderTypeParameterValueResponse) SetSecretValue(v CSSCMSDataModelModelsKeyfactorAPISecret)`

SetSecretValue sets SecretValue field to given value.

### HasSecretValue

`func (o *IdentityProviderProviderTypeParameterValueResponse) HasSecretValue() bool`

HasSecretValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



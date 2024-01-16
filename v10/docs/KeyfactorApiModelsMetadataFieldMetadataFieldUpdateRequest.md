# KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**DataType** | **int64** |  | 
**Hint** | Pointer to **string** |  | [optional] 
**Validation** | Pointer to **string** |  | [optional] 
**Enrollment** | Pointer to **int64** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Options** | Pointer to **string** |  | [optional] 
**DefaultValue** | Pointer to **string** |  | [optional] 
**DisplayOrder** | Pointer to **int64** |  | [optional] 

## Methods

### NewKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest

`func NewKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest(id int64, name string, description string, dataType int64, ) *KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest`

NewKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest instantiates a new KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequestWithDefaults

`func NewKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequestWithDefaults() *KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest`

NewKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequestWithDefaults instantiates a new KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDataType

`func (o *KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetDataType() int64`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetDataTypeOk() (*int64, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) SetDataType(v int64)`

SetDataType sets DataType field to given value.


### GetHint

`func (o *KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) SetHint(v string)`

SetHint sets Hint field to given value.

### HasHint

`func (o *KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) HasHint() bool`

HasHint returns a boolean if a field has been set.

### GetValidation

`func (o *KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetValidation() string`

GetValidation returns the Validation field if non-nil, zero value otherwise.

### GetValidationOk

`func (o *KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetValidationOk() (*string, bool)`

GetValidationOk returns a tuple with the Validation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidation

`func (o *KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) SetValidation(v string)`

SetValidation sets Validation field to given value.

### HasValidation

`func (o *KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) HasValidation() bool`

HasValidation returns a boolean if a field has been set.

### GetEnrollment

`func (o *KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetEnrollment() int64`

GetEnrollment returns the Enrollment field if non-nil, zero value otherwise.

### GetEnrollmentOk

`func (o *KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetEnrollmentOk() (*int64, bool)`

GetEnrollmentOk returns a tuple with the Enrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollment

`func (o *KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) SetEnrollment(v int64)`

SetEnrollment sets Enrollment field to given value.

### HasEnrollment

`func (o *KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) HasEnrollment() bool`

HasEnrollment returns a boolean if a field has been set.

### GetMessage

`func (o *KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOptions

`func (o *KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetOptions() string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetOptionsOk() (*string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) SetOptions(v string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetDefaultValue

`func (o *KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetDisplayOrder() int64`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) GetDisplayOrderOk() (*int64, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) SetDisplayOrder(v int64)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



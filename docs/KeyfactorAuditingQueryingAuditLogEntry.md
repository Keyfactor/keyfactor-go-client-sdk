# KeyfactorAuditingQueryingAuditLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Signature** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **int32** |  | [optional] 
**Operation** | Pointer to **int32** |  | [optional] 
**Level** | Pointer to **int32** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 
**EntityType** | Pointer to **string** |  | [optional] 
**AuditIdentifier** | Pointer to **string** |  | [optional] 
**ImmutableIdentifier** | Pointer to **string** |  | [optional] 

## Methods

### NewKeyfactorAuditingQueryingAuditLogEntry

`func NewKeyfactorAuditingQueryingAuditLogEntry() *KeyfactorAuditingQueryingAuditLogEntry`

NewKeyfactorAuditingQueryingAuditLogEntry instantiates a new KeyfactorAuditingQueryingAuditLogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorAuditingQueryingAuditLogEntryWithDefaults

`func NewKeyfactorAuditingQueryingAuditLogEntryWithDefaults() *KeyfactorAuditingQueryingAuditLogEntry`

NewKeyfactorAuditingQueryingAuditLogEntryWithDefaults instantiates a new KeyfactorAuditingQueryingAuditLogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyfactorAuditingQueryingAuditLogEntry) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorAuditingQueryingAuditLogEntry) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorAuditingQueryingAuditLogEntry) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KeyfactorAuditingQueryingAuditLogEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *KeyfactorAuditingQueryingAuditLogEntry) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *KeyfactorAuditingQueryingAuditLogEntry) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *KeyfactorAuditingQueryingAuditLogEntry) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *KeyfactorAuditingQueryingAuditLogEntry) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetMessage

`func (o *KeyfactorAuditingQueryingAuditLogEntry) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *KeyfactorAuditingQueryingAuditLogEntry) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *KeyfactorAuditingQueryingAuditLogEntry) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *KeyfactorAuditingQueryingAuditLogEntry) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSignature

`func (o *KeyfactorAuditingQueryingAuditLogEntry) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *KeyfactorAuditingQueryingAuditLogEntry) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *KeyfactorAuditingQueryingAuditLogEntry) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *KeyfactorAuditingQueryingAuditLogEntry) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetCategory

`func (o *KeyfactorAuditingQueryingAuditLogEntry) GetCategory() int32`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *KeyfactorAuditingQueryingAuditLogEntry) GetCategoryOk() (*int32, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *KeyfactorAuditingQueryingAuditLogEntry) SetCategory(v int32)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *KeyfactorAuditingQueryingAuditLogEntry) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetOperation

`func (o *KeyfactorAuditingQueryingAuditLogEntry) GetOperation() int32`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *KeyfactorAuditingQueryingAuditLogEntry) GetOperationOk() (*int32, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *KeyfactorAuditingQueryingAuditLogEntry) SetOperation(v int32)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *KeyfactorAuditingQueryingAuditLogEntry) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetLevel

`func (o *KeyfactorAuditingQueryingAuditLogEntry) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *KeyfactorAuditingQueryingAuditLogEntry) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *KeyfactorAuditingQueryingAuditLogEntry) SetLevel(v int32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *KeyfactorAuditingQueryingAuditLogEntry) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetUser

`func (o *KeyfactorAuditingQueryingAuditLogEntry) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *KeyfactorAuditingQueryingAuditLogEntry) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *KeyfactorAuditingQueryingAuditLogEntry) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *KeyfactorAuditingQueryingAuditLogEntry) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetEntityType

`func (o *KeyfactorAuditingQueryingAuditLogEntry) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *KeyfactorAuditingQueryingAuditLogEntry) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *KeyfactorAuditingQueryingAuditLogEntry) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *KeyfactorAuditingQueryingAuditLogEntry) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetAuditIdentifier

`func (o *KeyfactorAuditingQueryingAuditLogEntry) GetAuditIdentifier() string`

GetAuditIdentifier returns the AuditIdentifier field if non-nil, zero value otherwise.

### GetAuditIdentifierOk

`func (o *KeyfactorAuditingQueryingAuditLogEntry) GetAuditIdentifierOk() (*string, bool)`

GetAuditIdentifierOk returns a tuple with the AuditIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditIdentifier

`func (o *KeyfactorAuditingQueryingAuditLogEntry) SetAuditIdentifier(v string)`

SetAuditIdentifier sets AuditIdentifier field to given value.

### HasAuditIdentifier

`func (o *KeyfactorAuditingQueryingAuditLogEntry) HasAuditIdentifier() bool`

HasAuditIdentifier returns a boolean if a field has been set.

### GetImmutableIdentifier

`func (o *KeyfactorAuditingQueryingAuditLogEntry) GetImmutableIdentifier() string`

GetImmutableIdentifier returns the ImmutableIdentifier field if non-nil, zero value otherwise.

### GetImmutableIdentifierOk

`func (o *KeyfactorAuditingQueryingAuditLogEntry) GetImmutableIdentifierOk() (*string, bool)`

GetImmutableIdentifierOk returns a tuple with the ImmutableIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmutableIdentifier

`func (o *KeyfactorAuditingQueryingAuditLogEntry) SetImmutableIdentifier(v string)`

SetImmutableIdentifier sets ImmutableIdentifier field to given value.

### HasImmutableIdentifier

`func (o *KeyfactorAuditingQueryingAuditLogEntry) HasImmutableIdentifier() bool`

HasImmutableIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



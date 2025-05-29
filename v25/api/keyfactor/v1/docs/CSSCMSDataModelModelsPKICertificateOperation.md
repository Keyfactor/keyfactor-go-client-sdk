# CSSCMSDataModelModelsPKICertificateOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**OperationStart** | Pointer to **NullableTime** |  | [optional] 
**OperationEnd** | Pointer to **NullableTime** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**Comment** | Pointer to **NullableString** |  | [optional] 
**Action** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsPKICertificateOperation

`func NewCSSCMSDataModelModelsPKICertificateOperation() *CSSCMSDataModelModelsPKICertificateOperation`

NewCSSCMSDataModelModelsPKICertificateOperation instantiates a new CSSCMSDataModelModelsPKICertificateOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsPKICertificateOperationWithDefaults

`func NewCSSCMSDataModelModelsPKICertificateOperationWithDefaults() *CSSCMSDataModelModelsPKICertificateOperation`

NewCSSCMSDataModelModelsPKICertificateOperationWithDefaults instantiates a new CSSCMSDataModelModelsPKICertificateOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CSSCMSDataModelModelsPKICertificateOperation) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CSSCMSDataModelModelsPKICertificateOperation) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CSSCMSDataModelModelsPKICertificateOperation) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CSSCMSDataModelModelsPKICertificateOperation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOperationStart

`func (o *CSSCMSDataModelModelsPKICertificateOperation) GetOperationStart() time.Time`

GetOperationStart returns the OperationStart field if non-nil, zero value otherwise.

### GetOperationStartOk

`func (o *CSSCMSDataModelModelsPKICertificateOperation) GetOperationStartOk() (*time.Time, bool)`

GetOperationStartOk returns a tuple with the OperationStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationStart

`func (o *CSSCMSDataModelModelsPKICertificateOperation) SetOperationStart(v time.Time)`

SetOperationStart sets OperationStart field to given value.

### HasOperationStart

`func (o *CSSCMSDataModelModelsPKICertificateOperation) HasOperationStart() bool`

HasOperationStart returns a boolean if a field has been set.

### SetOperationStartNil

`func (o *CSSCMSDataModelModelsPKICertificateOperation) SetOperationStartNil(b bool)`

 SetOperationStartNil sets the value for OperationStart to be an explicit nil

### UnsetOperationStart
`func (o *CSSCMSDataModelModelsPKICertificateOperation) UnsetOperationStart()`

UnsetOperationStart ensures that no value is present for OperationStart, not even an explicit nil
### GetOperationEnd

`func (o *CSSCMSDataModelModelsPKICertificateOperation) GetOperationEnd() time.Time`

GetOperationEnd returns the OperationEnd field if non-nil, zero value otherwise.

### GetOperationEndOk

`func (o *CSSCMSDataModelModelsPKICertificateOperation) GetOperationEndOk() (*time.Time, bool)`

GetOperationEndOk returns a tuple with the OperationEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationEnd

`func (o *CSSCMSDataModelModelsPKICertificateOperation) SetOperationEnd(v time.Time)`

SetOperationEnd sets OperationEnd field to given value.

### HasOperationEnd

`func (o *CSSCMSDataModelModelsPKICertificateOperation) HasOperationEnd() bool`

HasOperationEnd returns a boolean if a field has been set.

### SetOperationEndNil

`func (o *CSSCMSDataModelModelsPKICertificateOperation) SetOperationEndNil(b bool)`

 SetOperationEndNil sets the value for OperationEnd to be an explicit nil

### UnsetOperationEnd
`func (o *CSSCMSDataModelModelsPKICertificateOperation) UnsetOperationEnd()`

UnsetOperationEnd ensures that no value is present for OperationEnd, not even an explicit nil
### GetUsername

`func (o *CSSCMSDataModelModelsPKICertificateOperation) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CSSCMSDataModelModelsPKICertificateOperation) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CSSCMSDataModelModelsPKICertificateOperation) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CSSCMSDataModelModelsPKICertificateOperation) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *CSSCMSDataModelModelsPKICertificateOperation) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *CSSCMSDataModelModelsPKICertificateOperation) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetComment

`func (o *CSSCMSDataModelModelsPKICertificateOperation) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CSSCMSDataModelModelsPKICertificateOperation) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CSSCMSDataModelModelsPKICertificateOperation) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CSSCMSDataModelModelsPKICertificateOperation) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *CSSCMSDataModelModelsPKICertificateOperation) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *CSSCMSDataModelModelsPKICertificateOperation) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetAction

`func (o *CSSCMSDataModelModelsPKICertificateOperation) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CSSCMSDataModelModelsPKICertificateOperation) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CSSCMSDataModelModelsPKICertificateOperation) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *CSSCMSDataModelModelsPKICertificateOperation) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *CSSCMSDataModelModelsPKICertificateOperation) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *CSSCMSDataModelModelsPKICertificateOperation) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



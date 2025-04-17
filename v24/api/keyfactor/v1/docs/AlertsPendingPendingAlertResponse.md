# AlertsPendingPendingAlertResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**Recipients** | Pointer to **[]string** |  | [optional] 
**CARequestId** | Pointer to **int32** |  | [optional] 
**CommonName** | Pointer to **NullableString** |  | [optional] 
**LogicalName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAlertsPendingPendingAlertResponse

`func NewAlertsPendingPendingAlertResponse() *AlertsPendingPendingAlertResponse`

NewAlertsPendingPendingAlertResponse instantiates a new AlertsPendingPendingAlertResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertsPendingPendingAlertResponseWithDefaults

`func NewAlertsPendingPendingAlertResponseWithDefaults() *AlertsPendingPendingAlertResponse`

NewAlertsPendingPendingAlertResponseWithDefaults instantiates a new AlertsPendingPendingAlertResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *AlertsPendingPendingAlertResponse) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *AlertsPendingPendingAlertResponse) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *AlertsPendingPendingAlertResponse) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *AlertsPendingPendingAlertResponse) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *AlertsPendingPendingAlertResponse) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *AlertsPendingPendingAlertResponse) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetMessage

`func (o *AlertsPendingPendingAlertResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AlertsPendingPendingAlertResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AlertsPendingPendingAlertResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AlertsPendingPendingAlertResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *AlertsPendingPendingAlertResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *AlertsPendingPendingAlertResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetRecipients

`func (o *AlertsPendingPendingAlertResponse) GetRecipients() []string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *AlertsPendingPendingAlertResponse) GetRecipientsOk() (*[]string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *AlertsPendingPendingAlertResponse) SetRecipients(v []string)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *AlertsPendingPendingAlertResponse) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### SetRecipientsNil

`func (o *AlertsPendingPendingAlertResponse) SetRecipientsNil(b bool)`

 SetRecipientsNil sets the value for Recipients to be an explicit nil

### UnsetRecipients
`func (o *AlertsPendingPendingAlertResponse) UnsetRecipients()`

UnsetRecipients ensures that no value is present for Recipients, not even an explicit nil
### GetCARequestId

`func (o *AlertsPendingPendingAlertResponse) GetCARequestId() int32`

GetCARequestId returns the CARequestId field if non-nil, zero value otherwise.

### GetCARequestIdOk

`func (o *AlertsPendingPendingAlertResponse) GetCARequestIdOk() (*int32, bool)`

GetCARequestIdOk returns a tuple with the CARequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCARequestId

`func (o *AlertsPendingPendingAlertResponse) SetCARequestId(v int32)`

SetCARequestId sets CARequestId field to given value.

### HasCARequestId

`func (o *AlertsPendingPendingAlertResponse) HasCARequestId() bool`

HasCARequestId returns a boolean if a field has been set.

### GetCommonName

`func (o *AlertsPendingPendingAlertResponse) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *AlertsPendingPendingAlertResponse) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *AlertsPendingPendingAlertResponse) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *AlertsPendingPendingAlertResponse) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### SetCommonNameNil

`func (o *AlertsPendingPendingAlertResponse) SetCommonNameNil(b bool)`

 SetCommonNameNil sets the value for CommonName to be an explicit nil

### UnsetCommonName
`func (o *AlertsPendingPendingAlertResponse) UnsetCommonName()`

UnsetCommonName ensures that no value is present for CommonName, not even an explicit nil
### GetLogicalName

`func (o *AlertsPendingPendingAlertResponse) GetLogicalName() string`

GetLogicalName returns the LogicalName field if non-nil, zero value otherwise.

### GetLogicalNameOk

`func (o *AlertsPendingPendingAlertResponse) GetLogicalNameOk() (*string, bool)`

GetLogicalNameOk returns a tuple with the LogicalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalName

`func (o *AlertsPendingPendingAlertResponse) SetLogicalName(v string)`

SetLogicalName sets LogicalName field to given value.

### HasLogicalName

`func (o *AlertsPendingPendingAlertResponse) HasLogicalName() bool`

HasLogicalName returns a boolean if a field has been set.

### SetLogicalNameNil

`func (o *AlertsPendingPendingAlertResponse) SetLogicalNameNil(b bool)`

 SetLogicalNameNil sets the value for LogicalName to be an explicit nil

### UnsetLogicalName
`func (o *AlertsPendingPendingAlertResponse) UnsetLogicalName()`

UnsetLogicalName ensures that no value is present for LogicalName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



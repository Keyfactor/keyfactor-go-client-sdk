# CSSCMSQueryingCoreComponentsQueryField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Default** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to [**CSSCMSQueryingCoreEnumsQueryDataType**](CSSCMSQueryingCoreEnumsQueryDataType.md) |  | [optional] 
**Comparisons** | Pointer to [**[]CSSCMSQueryingCoreComponentsQueryFieldComparison**](CSSCMSQueryingCoreComponentsQueryFieldComparison.md) |  | [optional] 
**Values** | Pointer to [**[]CSSCMSQueryingCoreComponentsQueryFieldValue**](CSSCMSQueryingCoreComponentsQueryFieldValue.md) |  | [optional] 
**MultiValue** | Pointer to [**CSSCMSQueryingCoreComponentsQueryFieldMultiValue**](CSSCMSQueryingCoreComponentsQueryFieldMultiValue.md) |  | [optional] 

## Methods

### NewCSSCMSQueryingCoreComponentsQueryField

`func NewCSSCMSQueryingCoreComponentsQueryField() *CSSCMSQueryingCoreComponentsQueryField`

NewCSSCMSQueryingCoreComponentsQueryField instantiates a new CSSCMSQueryingCoreComponentsQueryField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSQueryingCoreComponentsQueryFieldWithDefaults

`func NewCSSCMSQueryingCoreComponentsQueryFieldWithDefaults() *CSSCMSQueryingCoreComponentsQueryField`

NewCSSCMSQueryingCoreComponentsQueryFieldWithDefaults instantiates a new CSSCMSQueryingCoreComponentsQueryField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CSSCMSQueryingCoreComponentsQueryField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CSSCMSQueryingCoreComponentsQueryField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CSSCMSQueryingCoreComponentsQueryField) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CSSCMSQueryingCoreComponentsQueryField) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CSSCMSQueryingCoreComponentsQueryField) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CSSCMSQueryingCoreComponentsQueryField) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDefault

`func (o *CSSCMSQueryingCoreComponentsQueryField) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *CSSCMSQueryingCoreComponentsQueryField) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *CSSCMSQueryingCoreComponentsQueryField) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *CSSCMSQueryingCoreComponentsQueryField) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetType

`func (o *CSSCMSQueryingCoreComponentsQueryField) GetType() CSSCMSQueryingCoreEnumsQueryDataType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CSSCMSQueryingCoreComponentsQueryField) GetTypeOk() (*CSSCMSQueryingCoreEnumsQueryDataType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CSSCMSQueryingCoreComponentsQueryField) SetType(v CSSCMSQueryingCoreEnumsQueryDataType)`

SetType sets Type field to given value.

### HasType

`func (o *CSSCMSQueryingCoreComponentsQueryField) HasType() bool`

HasType returns a boolean if a field has been set.

### GetComparisons

`func (o *CSSCMSQueryingCoreComponentsQueryField) GetComparisons() []CSSCMSQueryingCoreComponentsQueryFieldComparison`

GetComparisons returns the Comparisons field if non-nil, zero value otherwise.

### GetComparisonsOk

`func (o *CSSCMSQueryingCoreComponentsQueryField) GetComparisonsOk() (*[]CSSCMSQueryingCoreComponentsQueryFieldComparison, bool)`

GetComparisonsOk returns a tuple with the Comparisons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisons

`func (o *CSSCMSQueryingCoreComponentsQueryField) SetComparisons(v []CSSCMSQueryingCoreComponentsQueryFieldComparison)`

SetComparisons sets Comparisons field to given value.

### HasComparisons

`func (o *CSSCMSQueryingCoreComponentsQueryField) HasComparisons() bool`

HasComparisons returns a boolean if a field has been set.

### SetComparisonsNil

`func (o *CSSCMSQueryingCoreComponentsQueryField) SetComparisonsNil(b bool)`

 SetComparisonsNil sets the value for Comparisons to be an explicit nil

### UnsetComparisons
`func (o *CSSCMSQueryingCoreComponentsQueryField) UnsetComparisons()`

UnsetComparisons ensures that no value is present for Comparisons, not even an explicit nil
### GetValues

`func (o *CSSCMSQueryingCoreComponentsQueryField) GetValues() []CSSCMSQueryingCoreComponentsQueryFieldValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *CSSCMSQueryingCoreComponentsQueryField) GetValuesOk() (*[]CSSCMSQueryingCoreComponentsQueryFieldValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *CSSCMSQueryingCoreComponentsQueryField) SetValues(v []CSSCMSQueryingCoreComponentsQueryFieldValue)`

SetValues sets Values field to given value.

### HasValues

`func (o *CSSCMSQueryingCoreComponentsQueryField) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValuesNil

`func (o *CSSCMSQueryingCoreComponentsQueryField) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *CSSCMSQueryingCoreComponentsQueryField) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil
### GetMultiValue

`func (o *CSSCMSQueryingCoreComponentsQueryField) GetMultiValue() CSSCMSQueryingCoreComponentsQueryFieldMultiValue`

GetMultiValue returns the MultiValue field if non-nil, zero value otherwise.

### GetMultiValueOk

`func (o *CSSCMSQueryingCoreComponentsQueryField) GetMultiValueOk() (*CSSCMSQueryingCoreComponentsQueryFieldMultiValue, bool)`

GetMultiValueOk returns a tuple with the MultiValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValue

`func (o *CSSCMSQueryingCoreComponentsQueryField) SetMultiValue(v CSSCMSQueryingCoreComponentsQueryFieldMultiValue)`

SetMultiValue sets MultiValue field to given value.

### HasMultiValue

`func (o *CSSCMSQueryingCoreComponentsQueryField) HasMultiValue() bool`

HasMultiValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



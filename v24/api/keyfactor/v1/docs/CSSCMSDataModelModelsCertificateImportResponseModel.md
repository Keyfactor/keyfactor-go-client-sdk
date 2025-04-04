# CSSCMSDataModelModelsCertificateImportResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImportStatus** | Pointer to [**CSSCMSDataModelEnumsCertificateSavedState**](CSSCMSDataModelEnumsCertificateSavedState.md) |  | [optional] 
**JobStatus** | Pointer to [**CSSCMSDataModelEnumsCertificateImportJobStatus**](CSSCMSDataModelEnumsCertificateImportJobStatus.md) |  | [optional] 
**InvalidKeystores** | Pointer to [**[]CSSCMSDataModelModelsInvalidKeystore**](CSSCMSDataModelModelsInvalidKeystore.md) |  | [optional] 
**Thumbprint** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsCertificateImportResponseModel

`func NewCSSCMSDataModelModelsCertificateImportResponseModel() *CSSCMSDataModelModelsCertificateImportResponseModel`

NewCSSCMSDataModelModelsCertificateImportResponseModel instantiates a new CSSCMSDataModelModelsCertificateImportResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsCertificateImportResponseModelWithDefaults

`func NewCSSCMSDataModelModelsCertificateImportResponseModelWithDefaults() *CSSCMSDataModelModelsCertificateImportResponseModel`

NewCSSCMSDataModelModelsCertificateImportResponseModelWithDefaults instantiates a new CSSCMSDataModelModelsCertificateImportResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImportStatus

`func (o *CSSCMSDataModelModelsCertificateImportResponseModel) GetImportStatus() CSSCMSDataModelEnumsCertificateSavedState`

GetImportStatus returns the ImportStatus field if non-nil, zero value otherwise.

### GetImportStatusOk

`func (o *CSSCMSDataModelModelsCertificateImportResponseModel) GetImportStatusOk() (*CSSCMSDataModelEnumsCertificateSavedState, bool)`

GetImportStatusOk returns a tuple with the ImportStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportStatus

`func (o *CSSCMSDataModelModelsCertificateImportResponseModel) SetImportStatus(v CSSCMSDataModelEnumsCertificateSavedState)`

SetImportStatus sets ImportStatus field to given value.

### HasImportStatus

`func (o *CSSCMSDataModelModelsCertificateImportResponseModel) HasImportStatus() bool`

HasImportStatus returns a boolean if a field has been set.

### GetJobStatus

`func (o *CSSCMSDataModelModelsCertificateImportResponseModel) GetJobStatus() CSSCMSDataModelEnumsCertificateImportJobStatus`

GetJobStatus returns the JobStatus field if non-nil, zero value otherwise.

### GetJobStatusOk

`func (o *CSSCMSDataModelModelsCertificateImportResponseModel) GetJobStatusOk() (*CSSCMSDataModelEnumsCertificateImportJobStatus, bool)`

GetJobStatusOk returns a tuple with the JobStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobStatus

`func (o *CSSCMSDataModelModelsCertificateImportResponseModel) SetJobStatus(v CSSCMSDataModelEnumsCertificateImportJobStatus)`

SetJobStatus sets JobStatus field to given value.

### HasJobStatus

`func (o *CSSCMSDataModelModelsCertificateImportResponseModel) HasJobStatus() bool`

HasJobStatus returns a boolean if a field has been set.

### GetInvalidKeystores

`func (o *CSSCMSDataModelModelsCertificateImportResponseModel) GetInvalidKeystores() []CSSCMSDataModelModelsInvalidKeystore`

GetInvalidKeystores returns the InvalidKeystores field if non-nil, zero value otherwise.

### GetInvalidKeystoresOk

`func (o *CSSCMSDataModelModelsCertificateImportResponseModel) GetInvalidKeystoresOk() (*[]CSSCMSDataModelModelsInvalidKeystore, bool)`

GetInvalidKeystoresOk returns a tuple with the InvalidKeystores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidKeystores

`func (o *CSSCMSDataModelModelsCertificateImportResponseModel) SetInvalidKeystores(v []CSSCMSDataModelModelsInvalidKeystore)`

SetInvalidKeystores sets InvalidKeystores field to given value.

### HasInvalidKeystores

`func (o *CSSCMSDataModelModelsCertificateImportResponseModel) HasInvalidKeystores() bool`

HasInvalidKeystores returns a boolean if a field has been set.

### SetInvalidKeystoresNil

`func (o *CSSCMSDataModelModelsCertificateImportResponseModel) SetInvalidKeystoresNil(b bool)`

 SetInvalidKeystoresNil sets the value for InvalidKeystores to be an explicit nil

### UnsetInvalidKeystores
`func (o *CSSCMSDataModelModelsCertificateImportResponseModel) UnsetInvalidKeystores()`

UnsetInvalidKeystores ensures that no value is present for InvalidKeystores, not even an explicit nil
### GetThumbprint

`func (o *CSSCMSDataModelModelsCertificateImportResponseModel) GetThumbprint() string`

GetThumbprint returns the Thumbprint field if non-nil, zero value otherwise.

### GetThumbprintOk

`func (o *CSSCMSDataModelModelsCertificateImportResponseModel) GetThumbprintOk() (*string, bool)`

GetThumbprintOk returns a tuple with the Thumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprint

`func (o *CSSCMSDataModelModelsCertificateImportResponseModel) SetThumbprint(v string)`

SetThumbprint sets Thumbprint field to given value.

### HasThumbprint

`func (o *CSSCMSDataModelModelsCertificateImportResponseModel) HasThumbprint() bool`

HasThumbprint returns a boolean if a field has been set.

### SetThumbprintNil

`func (o *CSSCMSDataModelModelsCertificateImportResponseModel) SetThumbprintNil(b bool)`

 SetThumbprintNil sets the value for Thumbprint to be an explicit nil

### UnsetThumbprint
`func (o *CSSCMSDataModelModelsCertificateImportResponseModel) UnsetThumbprint()`

UnsetThumbprint ensures that no value is present for Thumbprint, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



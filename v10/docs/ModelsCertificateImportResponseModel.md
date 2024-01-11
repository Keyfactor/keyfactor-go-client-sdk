# ModelsCertificateImportResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImportStatus** | Pointer to **int32** |  | [optional] 
**JobStatus** | Pointer to **int32** |  | [optional] 
**InvalidKeystores** | Pointer to [**[]ModelsInvalidKeystore**](ModelsInvalidKeystore.md) |  | [optional] 
**Thumbprint** | Pointer to **string** |  | [optional] 

## Methods

### NewModelsCertificateImportResponseModel

`func NewModelsCertificateImportResponseModel() *ModelsCertificateImportResponseModel`

NewModelsCertificateImportResponseModel instantiates a new ModelsCertificateImportResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsCertificateImportResponseModelWithDefaults

`func NewModelsCertificateImportResponseModelWithDefaults() *ModelsCertificateImportResponseModel`

NewModelsCertificateImportResponseModelWithDefaults instantiates a new ModelsCertificateImportResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImportStatus

`func (o *ModelsCertificateImportResponseModel) GetImportStatus() int32`

GetImportStatus returns the ImportStatus field if non-nil, zero value otherwise.

### GetImportStatusOk

`func (o *ModelsCertificateImportResponseModel) GetImportStatusOk() (*int32, bool)`

GetImportStatusOk returns a tuple with the ImportStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportStatus

`func (o *ModelsCertificateImportResponseModel) SetImportStatus(v int32)`

SetImportStatus sets ImportStatus field to given value.

### HasImportStatus

`func (o *ModelsCertificateImportResponseModel) HasImportStatus() bool`

HasImportStatus returns a boolean if a field has been set.

### GetJobStatus

`func (o *ModelsCertificateImportResponseModel) GetJobStatus() int32`

GetJobStatus returns the JobStatus field if non-nil, zero value otherwise.

### GetJobStatusOk

`func (o *ModelsCertificateImportResponseModel) GetJobStatusOk() (*int32, bool)`

GetJobStatusOk returns a tuple with the JobStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobStatus

`func (o *ModelsCertificateImportResponseModel) SetJobStatus(v int32)`

SetJobStatus sets JobStatus field to given value.

### HasJobStatus

`func (o *ModelsCertificateImportResponseModel) HasJobStatus() bool`

HasJobStatus returns a boolean if a field has been set.

### GetInvalidKeystores

`func (o *ModelsCertificateImportResponseModel) GetInvalidKeystores() []ModelsInvalidKeystore`

GetInvalidKeystores returns the InvalidKeystores field if non-nil, zero value otherwise.

### GetInvalidKeystoresOk

`func (o *ModelsCertificateImportResponseModel) GetInvalidKeystoresOk() (*[]ModelsInvalidKeystore, bool)`

GetInvalidKeystoresOk returns a tuple with the InvalidKeystores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidKeystores

`func (o *ModelsCertificateImportResponseModel) SetInvalidKeystores(v []ModelsInvalidKeystore)`

SetInvalidKeystores sets InvalidKeystores field to given value.

### HasInvalidKeystores

`func (o *ModelsCertificateImportResponseModel) HasInvalidKeystores() bool`

HasInvalidKeystores returns a boolean if a field has been set.

### GetThumbprint

`func (o *ModelsCertificateImportResponseModel) GetThumbprint() string`

GetThumbprint returns the Thumbprint field if non-nil, zero value otherwise.

### GetThumbprintOk

`func (o *ModelsCertificateImportResponseModel) GetThumbprintOk() (*string, bool)`

GetThumbprintOk returns a tuple with the Thumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprint

`func (o *ModelsCertificateImportResponseModel) SetThumbprint(v string)`

SetThumbprint sets Thumbprint field to given value.

### HasThumbprint

`func (o *ModelsCertificateImportResponseModel) HasThumbprint() bool`

HasThumbprint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



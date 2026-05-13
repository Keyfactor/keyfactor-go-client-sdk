# TemplatesTemplateCertificateExtensionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationTenant** | **string** |  | 
**Name** | **string** |  | 
**Oid** | **string** |  | 
**Required** | Pointer to **bool** |  | [optional] 
**Critical** | Pointer to **bool** |  | [optional] 

## Methods

### NewTemplatesTemplateCertificateExtensionsRequest

`func NewTemplatesTemplateCertificateExtensionsRequest(configurationTenant string, name string, oid string, ) *TemplatesTemplateCertificateExtensionsRequest`

NewTemplatesTemplateCertificateExtensionsRequest instantiates a new TemplatesTemplateCertificateExtensionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplatesTemplateCertificateExtensionsRequestWithDefaults

`func NewTemplatesTemplateCertificateExtensionsRequestWithDefaults() *TemplatesTemplateCertificateExtensionsRequest`

NewTemplatesTemplateCertificateExtensionsRequestWithDefaults instantiates a new TemplatesTemplateCertificateExtensionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurationTenant

`func (o *TemplatesTemplateCertificateExtensionsRequest) GetConfigurationTenant() string`

GetConfigurationTenant returns the ConfigurationTenant field if non-nil, zero value otherwise.

### GetConfigurationTenantOk

`func (o *TemplatesTemplateCertificateExtensionsRequest) GetConfigurationTenantOk() (*string, bool)`

GetConfigurationTenantOk returns a tuple with the ConfigurationTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationTenant

`func (o *TemplatesTemplateCertificateExtensionsRequest) SetConfigurationTenant(v string)`

SetConfigurationTenant sets ConfigurationTenant field to given value.


### GetName

`func (o *TemplatesTemplateCertificateExtensionsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplatesTemplateCertificateExtensionsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplatesTemplateCertificateExtensionsRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOid

`func (o *TemplatesTemplateCertificateExtensionsRequest) GetOid() string`

GetOid returns the Oid field if non-nil, zero value otherwise.

### GetOidOk

`func (o *TemplatesTemplateCertificateExtensionsRequest) GetOidOk() (*string, bool)`

GetOidOk returns a tuple with the Oid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOid

`func (o *TemplatesTemplateCertificateExtensionsRequest) SetOid(v string)`

SetOid sets Oid field to given value.


### GetRequired

`func (o *TemplatesTemplateCertificateExtensionsRequest) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *TemplatesTemplateCertificateExtensionsRequest) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *TemplatesTemplateCertificateExtensionsRequest) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *TemplatesTemplateCertificateExtensionsRequest) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetCritical

`func (o *TemplatesTemplateCertificateExtensionsRequest) GetCritical() bool`

GetCritical returns the Critical field if non-nil, zero value otherwise.

### GetCriticalOk

`func (o *TemplatesTemplateCertificateExtensionsRequest) GetCriticalOk() (*bool, bool)`

GetCriticalOk returns a tuple with the Critical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCritical

`func (o *TemplatesTemplateCertificateExtensionsRequest) SetCritical(v bool)`

SetCritical sets Critical field to given value.

### HasCritical

`func (o *TemplatesTemplateCertificateExtensionsRequest) HasCritical() bool`

HasCritical returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



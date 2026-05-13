# \TemplateApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTemplatesCADetails**](TemplateApi.md#CreateTemplatesCADetails) | **POST** /Templates/CADetails | Creates a Template with provided CA details
[**CreateTemplatesCustomExtensions**](TemplateApi.md#CreateTemplatesCustomExtensions) | **POST** /Templates/CustomExtensions | Creates a new custom certificate extension on the EJBCA server instance
[**CreateTemplatesImport**](TemplateApi.md#CreateTemplatesImport) | **POST** /Templates/Import | Imports templates from the provided configuration tenant
[**GetTemplates**](TemplateApi.md#GetTemplates) | **GET** /Templates | Returns all certificate templates according to the provided filter and output parameters
[**GetTemplatesById**](TemplateApi.md#GetTemplatesById) | **GET** /Templates/{id} | Returns the certificate template associated with the provided id
[**GetTemplatesCADetailsById**](TemplateApi.md#GetTemplatesCADetailsById) | **GET** /Templates/CADetails/{id} | Returns additional CA details of the certificate template associated with the provided id
[**GetTemplatesCustomExtensions**](TemplateApi.md#GetTemplatesCustomExtensions) | **GET** /Templates/CustomExtensions | Returns custom certificate extensions on an EJBCA server instance corresponding to the given configuration tenant
[**GetTemplatesExtendedKeyUsages**](TemplateApi.md#GetTemplatesExtendedKeyUsages) | **GET** /Templates/ExtendedKeyUsages | Fetches all extended key usages known by Command.
[**GetTemplatesSettings**](TemplateApi.md#GetTemplatesSettings) | **GET** /Templates/Settings | Gets the global template settings.
[**GetTemplatesSubjectParts**](TemplateApi.md#GetTemplatesSubjectParts) | **GET** /Templates/SubjectParts | Returns the valid subject parts possible for regular expressions.
[**UpdateTemplates**](TemplateApi.md#UpdateTemplates) | **PUT** /Templates | Updates a certificate template according to the provided properties
[**UpdateTemplatesCADetailsById**](TemplateApi.md#UpdateTemplatesCADetailsById) | **PUT** /Templates/CADetails/{id} | Sets CA details for a certificate template associated with the provided id
[**UpdateTemplatesCustomExtensions**](TemplateApi.md#UpdateTemplatesCustomExtensions) | **PUT** /Templates/CustomExtensions | Edits an existing custom certificate extension on the EJBCA server instance
[**UpdateTemplatesSettings**](TemplateApi.md#UpdateTemplatesSettings) | **PUT** /Templates/Settings | Replaces the existing global template settings.



## CreateTemplatesCADetails

> NewCreateTemplatesCADetailsRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).TemplatesTemplateDetailsCreateRequest(templatesTemplateDetailsCreateRequest).Execute()

Creates a Template with provided CA details

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    templatesTemplateDetailsCreateRequest := *openapiclient.NewTemplatesTemplateDetailsCreateRequest("Name_example", "Validity_example", "ConfigurationTenant_example") // TemplatesTemplateDetailsCreateRequest | Request object containing CA Details about the certificate template (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateApi.NewCreateTemplatesCADetailsRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).TemplatesTemplateDetailsCreateRequest(templatesTemplateDetailsCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateApi.CreateTemplatesCADetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTemplatesCADetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **templatesTemplateDetailsCreateRequest** | [**TemplatesTemplateDetailsCreateRequest**](TemplatesTemplateDetailsCreateRequest.md) | Request object containing CA Details about the certificate template | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTemplatesCustomExtensions

> NewCreateTemplatesCustomExtensionsRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).TemplatesTemplateCertificateExtensionsRequest(templatesTemplateCertificateExtensionsRequest).Execute()

Creates a new custom certificate extension on the EJBCA server instance

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    templatesTemplateCertificateExtensionsRequest := *openapiclient.NewTemplatesTemplateCertificateExtensionsRequest("ConfigurationTenant_example", "Name_example", "Oid_example") // TemplatesTemplateCertificateExtensionsRequest | Certificate extension object to create and the configuration tenant to add the extension to (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateApi.NewCreateTemplatesCustomExtensionsRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).TemplatesTemplateCertificateExtensionsRequest(templatesTemplateCertificateExtensionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateApi.CreateTemplatesCustomExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTemplatesCustomExtensionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **templatesTemplateCertificateExtensionsRequest** | [**TemplatesTemplateCertificateExtensionsRequest**](TemplatesTemplateCertificateExtensionsRequest.md) | Certificate extension object to create and the configuration tenant to add the extension to | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTemplatesImport

> NewCreateTemplatesImportRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).ConfigurationTenantConfigurationTenantRequest(configurationTenantConfigurationTenantRequest).Execute()

Imports templates from the provided configuration tenant

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    configurationTenantConfigurationTenantRequest := *openapiclient.NewConfigurationTenantConfigurationTenantRequest() // ConfigurationTenantConfigurationTenantRequest | Configuration tenant to import from (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateApi.NewCreateTemplatesImportRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).ConfigurationTenantConfigurationTenantRequest(configurationTenantConfigurationTenantRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateApi.CreateTemplatesImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTemplatesImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **configurationTenantConfigurationTenantRequest** | [**ConfigurationTenantConfigurationTenantRequest**](ConfigurationTenantConfigurationTenantRequest.md) | Configuration tenant to import from | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplates

> []TemplatesTemplateCollectionRetrievalResponse NewGetTemplatesRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns all certificate templates according to the provided filter and output parameters

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    queryString := "queryString_example" // string |  (optional)
    pageReturned := int32(56) // int32 |  (optional)
    returnLimit := int32(56) // int32 |  (optional)
    sortField := "sortField_example" // string |  (optional)
    sortAscending := openapiclient.Keyfactor.Common.QueryableExtensions.SortOrder(0) // KeyfactorCommonQueryableExtensionsSortOrder |  (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateApi.NewGetTemplatesRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateApi.GetTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTemplates`: []TemplatesTemplateCollectionRetrievalResponse
    fmt.Fprintf(os.Stdout, "Response from `TemplateApi.GetTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **queryString** | **string** |  | 
 **pageReturned** | **int32** |  | 
 **returnLimit** | **int32** |  | 
 **sortField** | **string** |  | 
 **sortAscending** | [**KeyfactorCommonQueryableExtensionsSortOrder**](KeyfactorCommonQueryableExtensionsSortOrder.md) |  | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]TemplatesTemplateCollectionRetrievalResponse**](TemplatesTemplateCollectionRetrievalResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplatesById

> TemplatesTemplateRetrievalResponse NewGetTemplatesByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns the certificate template associated with the provided id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | Keyfactor identifier of the certificate template
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateApi.NewGetTemplatesByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateApi.GetTemplatesById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTemplatesById`: TemplatesTemplateRetrievalResponse
    fmt.Fprintf(os.Stdout, "Response from `TemplateApi.GetTemplatesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Keyfactor identifier of the certificate template | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplatesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**TemplatesTemplateRetrievalResponse**](TemplatesTemplateRetrievalResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplatesCADetailsById

> TemplatesTemplateDetailsResponse NewGetTemplatesCADetailsByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns additional CA details of the certificate template associated with the provided id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | Keyfactor identifier of the certificate template
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateApi.NewGetTemplatesCADetailsByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateApi.GetTemplatesCADetailsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTemplatesCADetailsById`: TemplatesTemplateDetailsResponse
    fmt.Fprintf(os.Stdout, "Response from `TemplateApi.GetTemplatesCADetailsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Keyfactor identifier of the certificate template | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplatesCADetailsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**TemplatesTemplateDetailsResponse**](TemplatesTemplateDetailsResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplatesCustomExtensions

> []TemplatesTemplateCertificateExtensionResponse NewGetTemplatesCustomExtensionsRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).ConfigTenant(configTenant).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns custom certificate extensions on an EJBCA server instance corresponding to the given configuration tenant

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    configTenant := "configTenant_example" // string | Name of the configuration tenant (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateApi.NewGetTemplatesCustomExtensionsRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).ConfigTenant(configTenant).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateApi.GetTemplatesCustomExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTemplatesCustomExtensions`: []TemplatesTemplateCertificateExtensionResponse
    fmt.Fprintf(os.Stdout, "Response from `TemplateApi.GetTemplatesCustomExtensions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplatesCustomExtensionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **configTenant** | **string** | Name of the configuration tenant | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]TemplatesTemplateCertificateExtensionResponse**](TemplatesTemplateCertificateExtensionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplatesExtendedKeyUsages

> []TemplatesTemplateExtendedKeyUsageResponse NewGetTemplatesExtendedKeyUsagesRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Fetches all extended key usages known by Command.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateApi.NewGetTemplatesExtendedKeyUsagesRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateApi.GetTemplatesExtendedKeyUsages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTemplatesExtendedKeyUsages`: []TemplatesTemplateExtendedKeyUsageResponse
    fmt.Fprintf(os.Stdout, "Response from `TemplateApi.GetTemplatesExtendedKeyUsages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplatesExtendedKeyUsagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]TemplatesTemplateExtendedKeyUsageResponse**](TemplatesTemplateExtendedKeyUsageResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplatesSettings

> TemplatesGlobalGlobalTemplateSettingsResponse NewGetTemplatesSettingsRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Gets the global template settings.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateApi.NewGetTemplatesSettingsRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateApi.GetTemplatesSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTemplatesSettings`: TemplatesGlobalGlobalTemplateSettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `TemplateApi.GetTemplatesSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplatesSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**TemplatesGlobalGlobalTemplateSettingsResponse**](TemplatesGlobalGlobalTemplateSettingsResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplatesSubjectParts

> []EnrollmentPatternsValidSubjectPartResponse NewGetTemplatesSubjectPartsRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns the valid subject parts possible for regular expressions.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateApi.NewGetTemplatesSubjectPartsRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateApi.GetTemplatesSubjectParts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTemplatesSubjectParts`: []EnrollmentPatternsValidSubjectPartResponse
    fmt.Fprintf(os.Stdout, "Response from `TemplateApi.GetTemplatesSubjectParts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplatesSubjectPartsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]EnrollmentPatternsValidSubjectPartResponse**](EnrollmentPatternsValidSubjectPartResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTemplates

> TemplatesTemplateRetrievalResponse NewUpdateTemplatesRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).TemplatesTemplateUpdateRequest(templatesTemplateUpdateRequest).Execute()

Updates a certificate template according to the provided properties

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    templatesTemplateUpdateRequest := *openapiclient.NewTemplatesTemplateUpdateRequest() // TemplatesTemplateUpdateRequest | Properties of the certificate template to be updated (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateApi.NewUpdateTemplatesRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).TemplatesTemplateUpdateRequest(templatesTemplateUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateApi.UpdateTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTemplates`: TemplatesTemplateRetrievalResponse
    fmt.Fprintf(os.Stdout, "Response from `TemplateApi.UpdateTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **templatesTemplateUpdateRequest** | [**TemplatesTemplateUpdateRequest**](TemplatesTemplateUpdateRequest.md) | Properties of the certificate template to be updated | 

### Return type

[**TemplatesTemplateRetrievalResponse**](TemplatesTemplateRetrievalResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTemplatesCADetailsById

> NewUpdateTemplatesCADetailsByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).TemplatesTemplateDetailsUpdateRequest(templatesTemplateDetailsUpdateRequest).Execute()

Sets CA details for a certificate template associated with the provided id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | Keyfactor identifier of the certificate template
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    templatesTemplateDetailsUpdateRequest := *openapiclient.NewTemplatesTemplateDetailsUpdateRequest("Name_example", "OID_example", "Validity_example") // TemplatesTemplateDetailsUpdateRequest | Request object containing CA Details about the certificate template (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateApi.NewUpdateTemplatesCADetailsByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).TemplatesTemplateDetailsUpdateRequest(templatesTemplateDetailsUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateApi.UpdateTemplatesCADetailsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Keyfactor identifier of the certificate template | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTemplatesCADetailsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **templatesTemplateDetailsUpdateRequest** | [**TemplatesTemplateDetailsUpdateRequest**](TemplatesTemplateDetailsUpdateRequest.md) | Request object containing CA Details about the certificate template | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTemplatesCustomExtensions

> NewUpdateTemplatesCustomExtensionsRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).TemplatesTemplateCertificateExtensionsRequest(templatesTemplateCertificateExtensionsRequest).Execute()

Edits an existing custom certificate extension on the EJBCA server instance

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    templatesTemplateCertificateExtensionsRequest := *openapiclient.NewTemplatesTemplateCertificateExtensionsRequest("ConfigurationTenant_example", "Name_example", "Oid_example") // TemplatesTemplateCertificateExtensionsRequest | Certificate extension object to update and the configuration tenant to update the extension on (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateApi.NewUpdateTemplatesCustomExtensionsRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).TemplatesTemplateCertificateExtensionsRequest(templatesTemplateCertificateExtensionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateApi.UpdateTemplatesCustomExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTemplatesCustomExtensionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **templatesTemplateCertificateExtensionsRequest** | [**TemplatesTemplateCertificateExtensionsRequest**](TemplatesTemplateCertificateExtensionsRequest.md) | Certificate extension object to update and the configuration tenant to update the extension on | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTemplatesSettings

> TemplatesGlobalGlobalTemplateSettingsResponse NewUpdateTemplatesSettingsRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).TemplatesGlobalGlobalTemplateSettingsRequest(templatesGlobalGlobalTemplateSettingsRequest).Execute()

Replaces the existing global template settings.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    templatesGlobalGlobalTemplateSettingsRequest := *openapiclient.NewTemplatesGlobalGlobalTemplateSettingsRequest([]openapiclient.TemplatesGlobalGlobalTemplateRegexRequest{*openapiclient.NewTemplatesGlobalGlobalTemplateRegexRequest("SubjectPart_example")}, []openapiclient.TemplatesGlobalGlobalTemplateDefaultRequest{*openapiclient.NewTemplatesGlobalGlobalTemplateDefaultRequest("SubjectPart_example")}, *openapiclient.NewTemplatesGlobalGlobalTemplatePolicyRequest(false, false, false, openapiclient.CSS.CMS.Core.Enums.TemplateCertificateOwnerRole(0), *openapiclient.NewEnrollmentPatternsAlgorithmsKeyInfoRequest())) // TemplatesGlobalGlobalTemplateSettingsRequest | The new global template settings. (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateApi.NewUpdateTemplatesSettingsRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).TemplatesGlobalGlobalTemplateSettingsRequest(templatesGlobalGlobalTemplateSettingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateApi.UpdateTemplatesSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTemplatesSettings`: TemplatesGlobalGlobalTemplateSettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `TemplateApi.UpdateTemplatesSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTemplatesSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **templatesGlobalGlobalTemplateSettingsRequest** | [**TemplatesGlobalGlobalTemplateSettingsRequest**](TemplatesGlobalGlobalTemplateSettingsRequest.md) | The new global template settings. | 

### Return type

[**TemplatesGlobalGlobalTemplateSettingsResponse**](TemplatesGlobalGlobalTemplateSettingsResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


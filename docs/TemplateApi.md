# \TemplateApi

All URIs are relative to *https://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TemplateGetGlobalSettings**](TemplateApi.md#TemplateGetGlobalSettings) | **Get** /Templates/Settings | Gets the global template settings.
[**TemplateGetTemplate**](TemplateApi.md#TemplateGetTemplate) | **Get** /Templates/{id} | Returns the certificate template associated with the provided id
[**TemplateGetTemplates**](TemplateApi.md#TemplateGetTemplates) | **Get** /Templates | Returns all certificate templates according to the provided filter and output parameters
[**TemplateGetValidSubjectParts**](TemplateApi.md#TemplateGetValidSubjectParts) | **Get** /Templates/SubjectParts | 
[**TemplateImport**](TemplateApi.md#TemplateImport) | **Post** /Templates/Import | Imports templates from the provided configuration tenant
[**TemplateUpdateGlobalSettings**](TemplateApi.md#TemplateUpdateGlobalSettings) | **Put** /Templates/Settings | Replaces the existing global template settings.
[**TemplateUpdateTemplate**](TemplateApi.md#TemplateUpdateTemplate) | **Put** /Templates | Updates a certificate template according to the provided properties



## TemplateGetGlobalSettings

> KeyfactorApiModelsTemplatesGlobalTemplateSettingsResponse TemplateGetGlobalSettings(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateApi.TemplateGetGlobalSettings(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateApi.TemplateGetGlobalSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateGetGlobalSettings`: KeyfactorApiModelsTemplatesGlobalTemplateSettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `TemplateApi.TemplateGetGlobalSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTemplateGetGlobalSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsTemplatesGlobalTemplateSettingsResponse**](KeyfactorApiModelsTemplatesGlobalTemplateSettingsResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplateGetTemplate

> ModelsTemplateRetrievalResponse TemplateGetTemplate(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateApi.TemplateGetTemplate(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateApi.TemplateGetTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateGetTemplate`: ModelsTemplateRetrievalResponse
    fmt.Fprintf(os.Stdout, "Response from `TemplateApi.TemplateGetTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Keyfactor identifier of the certificate template | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplateGetTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsTemplateRetrievalResponse**](ModelsTemplateRetrievalResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplateGetTemplates

> []ModelsTemplateCollectionRetrievalResponse TemplateGetTemplates(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).SqQueryString(sqQueryString).SqPageReturned(sqPageReturned).SqReturnLimit(sqReturnLimit).SqSortField(sqSortField).SqSortAscending(sqSortAscending).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")
    sqQueryString := "sqQueryString_example" // string | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) (optional)
    sqPageReturned := int32(56) // int32 | The current page within the result set to be returned (optional)
    sqReturnLimit := int32(56) // int32 | Maximum number of records to be returned in a single call (optional)
    sqSortField := "sqSortField_example" // string | Field by which the results should be sorted (view results via Management Portal for sortable columns) (optional)
    sqSortAscending := int32(56) // int32 | Field sort direction [0=ascending, 1=descending] (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateApi.TemplateGetTemplates(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).SqQueryString(sqQueryString).SqPageReturned(sqPageReturned).SqReturnLimit(sqReturnLimit).SqSortField(sqSortField).SqSortAscending(sqSortAscending).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateApi.TemplateGetTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateGetTemplates`: []ModelsTemplateCollectionRetrievalResponse
    fmt.Fprintf(os.Stdout, "Response from `TemplateApi.TemplateGetTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTemplateGetTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]
 **sqQueryString** | **string** | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) | 
 **sqPageReturned** | **int32** | The current page within the result set to be returned | 
 **sqReturnLimit** | **int32** | Maximum number of records to be returned in a single call | 
 **sqSortField** | **string** | Field by which the results should be sorted (view results via Management Portal for sortable columns) | 
 **sqSortAscending** | **int32** | Field sort direction [0&#x3D;ascending, 1&#x3D;descending] | 

### Return type

[**[]ModelsTemplateCollectionRetrievalResponse**](ModelsTemplateCollectionRetrievalResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplateGetValidSubjectParts

> []KeyfactorApiModelsTemplatesValidSubjectPartResponse TemplateGetValidSubjectParts(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()



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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateApi.TemplateGetValidSubjectParts(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateApi.TemplateGetValidSubjectParts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateGetValidSubjectParts`: []KeyfactorApiModelsTemplatesValidSubjectPartResponse
    fmt.Fprintf(os.Stdout, "Response from `TemplateApi.TemplateGetValidSubjectParts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTemplateGetValidSubjectPartsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**[]KeyfactorApiModelsTemplatesValidSubjectPartResponse**](KeyfactorApiModelsTemplatesValidSubjectPartResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplateImport

> TemplateImport(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).ConfigurationTenantRequest(configurationTenantRequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    configurationTenantRequest := *openapiclient.NewKeyfactorApiModelsConfigurationTenantConfigurationTenantRequest() // KeyfactorApiModelsConfigurationTenantConfigurationTenantRequest | Configuration tenant to import from
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateApi.TemplateImport(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).ConfigurationTenantRequest(configurationTenantRequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateApi.TemplateImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTemplateImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **configurationTenantRequest** | [**KeyfactorApiModelsConfigurationTenantConfigurationTenantRequest**](KeyfactorApiModelsConfigurationTenantConfigurationTenantRequest.md) | Configuration tenant to import from | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplateUpdateGlobalSettings

> KeyfactorApiModelsTemplatesGlobalTemplateSettingsResponse TemplateUpdateGlobalSettings(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Settings(settings).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    settings := *openapiclient.NewKeyfactorApiModelsTemplatesGlobalTemplateSettingsRequest([]openapiclient.KeyfactorApiModelsTemplatesGlobalTemplateRegexRequest{*openapiclient.NewKeyfactorApiModelsTemplatesGlobalTemplateRegexRequest("SubjectPart_example")}, []openapiclient.KeyfactorApiModelsTemplatesGlobalTemplateDefaultRequest{*openapiclient.NewKeyfactorApiModelsTemplatesGlobalTemplateDefaultRequest("SubjectPart_example")}, *openapiclient.NewKeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest([]int32{int32(123)}, []string{"ECCValidCurves_example"}, false, false, false)) // KeyfactorApiModelsTemplatesGlobalTemplateSettingsRequest | The new global template settings.
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateApi.TemplateUpdateGlobalSettings(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Settings(settings).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateApi.TemplateUpdateGlobalSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateUpdateGlobalSettings`: KeyfactorApiModelsTemplatesGlobalTemplateSettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `TemplateApi.TemplateUpdateGlobalSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTemplateUpdateGlobalSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **settings** | [**KeyfactorApiModelsTemplatesGlobalTemplateSettingsRequest**](KeyfactorApiModelsTemplatesGlobalTemplateSettingsRequest.md) | The new global template settings. | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsTemplatesGlobalTemplateSettingsResponse**](KeyfactorApiModelsTemplatesGlobalTemplateSettingsResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplateUpdateTemplate

> ModelsTemplateRetrievalResponse TemplateUpdateTemplate(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Template(template).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    template := *openapiclient.NewModelsTemplateUpdateRequest() // ModelsTemplateUpdateRequest | Properties of the certificate template to be updated
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateApi.TemplateUpdateTemplate(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Template(template).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateApi.TemplateUpdateTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateUpdateTemplate`: ModelsTemplateRetrievalResponse
    fmt.Fprintf(os.Stdout, "Response from `TemplateApi.TemplateUpdateTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTemplateUpdateTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **template** | [**ModelsTemplateUpdateRequest**](ModelsTemplateUpdateRequest.md) | Properties of the certificate template to be updated | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsTemplateRetrievalResponse**](ModelsTemplateRetrievalResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


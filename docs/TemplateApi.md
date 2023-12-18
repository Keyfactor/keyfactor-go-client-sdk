# \TemplateApi

All URIs are relative to */Keyfactor/API*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TemplatesGet**](TemplateApi.md#TemplatesGet) | **Get** /Templates | Returns all certificate templates according to the provided filter and output parameters
[**TemplatesIdGet**](TemplateApi.md#TemplatesIdGet) | **Get** /Templates/{id} | Returns the certificate template associated with the provided id
[**TemplatesImportPost**](TemplateApi.md#TemplatesImportPost) | **Post** /Templates/Import | Imports templates from the provided configuration tenant
[**TemplatesPut**](TemplateApi.md#TemplatesPut) | **Put** /Templates | Updates a certificate template according to the provided properties
[**TemplatesSettingsGet**](TemplateApi.md#TemplatesSettingsGet) | **Get** /Templates/Settings | Gets the global template settings.
[**TemplatesSettingsPut**](TemplateApi.md#TemplatesSettingsPut) | **Put** /Templates/Settings | Replaces the existing global template settings.
[**TemplatesSubjectPartsGet**](TemplateApi.md#TemplatesSubjectPartsGet) | **Get** /Templates/SubjectParts | 



## TemplatesGet

> []KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse TemplatesGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    queryString := "queryString_example" // string |  (optional)
    pageReturned := int32(56) // int32 |  (optional)
    returnLimit := int32(56) // int32 |  (optional)
    sortField := "sortField_example" // string |  (optional)
    sortAscending := openapiclient.Keyfactor.Common.QueryableExtensionsSortOrder(0) // KeyfactorCommonQueryableExtensionsSortOrder |  (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateApi.TemplatesGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateApi.TemplatesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplatesGet`: []KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse
    fmt.Fprintf(os.Stdout, "Response from `TemplateApi.TemplatesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTemplatesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **queryString** | **string** |  | 
 **pageReturned** | **int32** |  | 
 **returnLimit** | **int32** |  | 
 **sortField** | **string** |  | 
 **sortAscending** | [**KeyfactorCommonQueryableExtensionsSortOrder**](KeyfactorCommonQueryableExtensionsSortOrder.md) |  | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse**](KeyfactorWebKeyfactorApiModelsTemplatesTemplateCollectionRetrievalResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplatesIdGet

> KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse TemplatesIdGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateApi.TemplatesIdGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateApi.TemplatesIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplatesIdGet`: KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse
    fmt.Fprintf(os.Stdout, "Response from `TemplateApi.TemplatesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Keyfactor identifier of the certificate template | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplatesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse**](KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplatesImportPost

> TemplatesImportPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsConfigurationTenantConfigurationTenantRequest(keyfactorWebKeyfactorApiModelsConfigurationTenantConfigurationTenantRequest).Execute()

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
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsConfigurationTenantConfigurationTenantRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsConfigurationTenantConfigurationTenantRequest() // KeyfactorWebKeyfactorApiModelsConfigurationTenantConfigurationTenantRequest | Configuration tenant to import from (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateApi.TemplatesImportPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsConfigurationTenantConfigurationTenantRequest(keyfactorWebKeyfactorApiModelsConfigurationTenantConfigurationTenantRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateApi.TemplatesImportPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTemplatesImportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsConfigurationTenantConfigurationTenantRequest** | [**KeyfactorWebKeyfactorApiModelsConfigurationTenantConfigurationTenantRequest**](KeyfactorWebKeyfactorApiModelsConfigurationTenantConfigurationTenantRequest.md) | Configuration tenant to import from | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patchjson, application/json, text/json, application/*json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplatesPut

> KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse TemplatesPut(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest(keyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest).Execute()

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
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest() // KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest | Properties of the certificate template to be updated (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateApi.TemplatesPut(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest(keyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateApi.TemplatesPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplatesPut`: KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse
    fmt.Fprintf(os.Stdout, "Response from `TemplateApi.TemplatesPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTemplatesPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest** | [**KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest**](KeyfactorWebKeyfactorApiModelsTemplatesTemplateUpdateRequest.md) | Properties of the certificate template to be updated | 

### Return type

[**KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse**](KeyfactorWebKeyfactorApiModelsTemplatesTemplateRetrievalResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patchjson, application/json, text/json, application/*json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplatesSettingsGet

> KeyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplateSettingsResponse TemplatesSettingsGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateApi.TemplatesSettingsGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateApi.TemplatesSettingsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplatesSettingsGet`: KeyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplateSettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `TemplateApi.TemplatesSettingsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTemplatesSettingsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**KeyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplateSettingsResponse**](KeyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplateSettingsResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplatesSettingsPut

> KeyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplateSettingsResponse TemplatesSettingsPut(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplateSettingsRequest(keyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplateSettingsRequest).Execute()

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
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplateSettingsRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplateSettingsRequest([]openapiclient.KeyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplateRegexRequest{*openapiclient.NewKeyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplateRegexRequest("SubjectPart_example")}, []openapiclient.KeyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplateDefaultRequest{*openapiclient.NewKeyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplateDefaultRequest("SubjectPart_example")}, *openapiclient.NewKeyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplatePolicyRequest(false, false, false, *openapiclient.NewCSSCMSDataModelModelsTemplatesAlgorithmsKeyInfo())) // KeyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplateSettingsRequest | The new global template settings. (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateApi.TemplatesSettingsPut(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplateSettingsRequest(keyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplateSettingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateApi.TemplatesSettingsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplatesSettingsPut`: KeyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplateSettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `TemplateApi.TemplatesSettingsPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTemplatesSettingsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplateSettingsRequest** | [**KeyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplateSettingsRequest**](KeyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplateSettingsRequest.md) | The new global template settings. | 

### Return type

[**KeyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplateSettingsResponse**](KeyfactorWebKeyfactorApiModelsTemplatesGlobalGlobalTemplateSettingsResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patchjson, application/json, text/json, application/*json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplatesSubjectPartsGet

> []KeyfactorWebKeyfactorApiModelsTemplatesValidSubjectPartResponse TemplatesSubjectPartsGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()



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
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateApi.TemplatesSubjectPartsGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateApi.TemplatesSubjectPartsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplatesSubjectPartsGet`: []KeyfactorWebKeyfactorApiModelsTemplatesValidSubjectPartResponse
    fmt.Fprintf(os.Stdout, "Response from `TemplateApi.TemplatesSubjectPartsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTemplatesSubjectPartsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]KeyfactorWebKeyfactorApiModelsTemplatesValidSubjectPartResponse**](KeyfactorWebKeyfactorApiModelsTemplatesValidSubjectPartResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


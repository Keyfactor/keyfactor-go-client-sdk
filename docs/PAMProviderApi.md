# \PAMProviderApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PAMProviderCreatePamProvider**](PAMProviderApi.md#PAMProviderCreatePamProvider) | **Post** /PamProviders | Creates a new PAM provider with the associated properties
[**PAMProviderCreatePamProviderType**](PAMProviderApi.md#PAMProviderCreatePamProviderType) | **Post** /PamProviders/Types | Creates a new PAM provider type with the associated properties
[**PAMProviderDeletePamProvider**](PAMProviderApi.md#PAMProviderDeletePamProvider) | **Delete** /PamProviders/{id} | Deletes a PAM Provider
[**PAMProviderGetPamProvider**](PAMProviderApi.md#PAMProviderGetPamProvider) | **Get** /PamProviders/{id} | Returns a single PAM Provider that matches the associated id
[**PAMProviderGetPamProviderTypes**](PAMProviderApi.md#PAMProviderGetPamProviderTypes) | **Get** /PamProviders/Types | Returns all PAM providers in the Keyfactor instance
[**PAMProviderGetPamProviders**](PAMProviderApi.md#PAMProviderGetPamProviders) | **Get** /PamProviders | Returns all PAM providers according to the provided filter and output parameters
[**PAMProviderUpdatePamProvider**](PAMProviderApi.md#PAMProviderUpdatePamProvider) | **Put** /PamProviders | Updates an existing PAM provider according to the provided properties



## PAMProviderCreatePamProvider

> CSSCMSDataModelModelsProvider PAMProviderCreatePamProvider(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Provider(provider).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Creates a new PAM provider with the associated properties



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
    provider := *openapiclient.NewCSSCMSDataModelModelsProvider("Name_example", *openapiclient.NewCSSCMSDataModelModelsProviderType()) // CSSCMSDataModelModelsProvider | PAM provider properties to be used
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PAMProviderApi.PAMProviderCreatePamProvider(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Provider(provider).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PAMProviderApi.PAMProviderCreatePamProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PAMProviderCreatePamProvider`: CSSCMSDataModelModelsProvider
    fmt.Fprintf(os.Stdout, "Response from `PAMProviderApi.PAMProviderCreatePamProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPAMProviderCreatePamProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **provider** | [**CSSCMSDataModelModelsProvider**](CSSCMSDataModelModelsProvider.md) | PAM provider properties to be used | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**CSSCMSDataModelModelsProvider**](CSSCMSDataModelModelsProvider.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PAMProviderCreatePamProviderType

> KeyfactorApiPAMProviderTypeResponse PAMProviderCreatePamProviderType(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Type_(type_).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Creates a new PAM provider type with the associated properties



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
    type_ := *openapiclient.NewKeyfactorApiPAMProviderTypeCreateRequest("Name_example") // KeyfactorApiPAMProviderTypeCreateRequest | PAM provider type properties to be used
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PAMProviderApi.PAMProviderCreatePamProviderType(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Type_(type_).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PAMProviderApi.PAMProviderCreatePamProviderType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PAMProviderCreatePamProviderType`: KeyfactorApiPAMProviderTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `PAMProviderApi.PAMProviderCreatePamProviderType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPAMProviderCreatePamProviderTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **type_** | [**KeyfactorApiPAMProviderTypeCreateRequest**](KeyfactorApiPAMProviderTypeCreateRequest.md) | PAM provider type properties to be used | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiPAMProviderTypeResponse**](KeyfactorApiPAMProviderTypeResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PAMProviderDeletePamProvider

> PAMProviderDeletePamProvider(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Deletes a PAM Provider

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
    id := int32(56) // int32 | Keyfactor identifier of the PAM provider to be deleted
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PAMProviderApi.PAMProviderDeletePamProvider(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PAMProviderApi.PAMProviderDeletePamProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Keyfactor identifier of the PAM provider to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiPAMProviderDeletePamProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PAMProviderGetPamProvider

> CSSCMSDataModelModelsProvider PAMProviderGetPamProvider(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns a single PAM Provider that matches the associated id



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
    id := int32(56) // int32 | Keyfactor identifier of the PAM provider
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PAMProviderApi.PAMProviderGetPamProvider(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PAMProviderApi.PAMProviderGetPamProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PAMProviderGetPamProvider`: CSSCMSDataModelModelsProvider
    fmt.Fprintf(os.Stdout, "Response from `PAMProviderApi.PAMProviderGetPamProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Keyfactor identifier of the PAM provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiPAMProviderGetPamProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**CSSCMSDataModelModelsProvider**](CSSCMSDataModelModelsProvider.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PAMProviderGetPamProviderTypes

> []CSSCMSDataModelModelsProviderType PAMProviderGetPamProviderTypes(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns all PAM providers in the Keyfactor instance



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
    resp, r, err := apiClient.PAMProviderApi.PAMProviderGetPamProviderTypes(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PAMProviderApi.PAMProviderGetPamProviderTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PAMProviderGetPamProviderTypes`: []CSSCMSDataModelModelsProviderType
    fmt.Fprintf(os.Stdout, "Response from `PAMProviderApi.PAMProviderGetPamProviderTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPAMProviderGetPamProviderTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**[]CSSCMSDataModelModelsProviderType**](CSSCMSDataModelModelsProviderType.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PAMProviderGetPamProviders

> []CSSCMSDataModelModelsProvider PAMProviderGetPamProviders(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PqQueryString(pqQueryString).PqPageReturned(pqPageReturned).PqReturnLimit(pqReturnLimit).PqSortField(pqSortField).PqSortAscending(pqSortAscending).Execute()

Returns all PAM providers according to the provided filter and output parameters



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
    pqQueryString := "pqQueryString_example" // string | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) (optional)
    pqPageReturned := int32(56) // int32 | The current page within the result set to be returned (optional)
    pqReturnLimit := int32(56) // int32 | Maximum number of records to be returned in a single call (optional)
    pqSortField := "pqSortField_example" // string | Field by which the results should be sorted (view results via Management Portal for sortable columns) (optional)
    pqSortAscending := int32(56) // int32 | Field sort direction [0=ascending, 1=descending] (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PAMProviderApi.PAMProviderGetPamProviders(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PqQueryString(pqQueryString).PqPageReturned(pqPageReturned).PqReturnLimit(pqReturnLimit).PqSortField(pqSortField).PqSortAscending(pqSortAscending).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PAMProviderApi.PAMProviderGetPamProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PAMProviderGetPamProviders`: []CSSCMSDataModelModelsProvider
    fmt.Fprintf(os.Stdout, "Response from `PAMProviderApi.PAMProviderGetPamProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPAMProviderGetPamProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]
 **pqQueryString** | **string** | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) | 
 **pqPageReturned** | **int32** | The current page within the result set to be returned | 
 **pqReturnLimit** | **int32** | Maximum number of records to be returned in a single call | 
 **pqSortField** | **string** | Field by which the results should be sorted (view results via Management Portal for sortable columns) | 
 **pqSortAscending** | **int32** | Field sort direction [0&#x3D;ascending, 1&#x3D;descending] | 

### Return type

[**[]CSSCMSDataModelModelsProvider**](CSSCMSDataModelModelsProvider.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PAMProviderUpdatePamProvider

> CSSCMSDataModelModelsProvider PAMProviderUpdatePamProvider(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Provider(provider).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Updates an existing PAM provider according to the provided properties



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
    provider := *openapiclient.NewCSSCMSDataModelModelsProvider("Name_example", *openapiclient.NewCSSCMSDataModelModelsProviderType()) // CSSCMSDataModelModelsProvider | PAM provider properties to be used
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PAMProviderApi.PAMProviderUpdatePamProvider(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Provider(provider).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PAMProviderApi.PAMProviderUpdatePamProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PAMProviderUpdatePamProvider`: CSSCMSDataModelModelsProvider
    fmt.Fprintf(os.Stdout, "Response from `PAMProviderApi.PAMProviderUpdatePamProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPAMProviderUpdatePamProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **provider** | [**CSSCMSDataModelModelsProvider**](CSSCMSDataModelModelsProvider.md) | PAM provider properties to be used | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**CSSCMSDataModelModelsProvider**](CSSCMSDataModelModelsProvider.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


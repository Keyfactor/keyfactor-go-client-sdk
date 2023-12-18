# \PAMProviderApi

All URIs are relative to */Keyfactor/API*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PamProvidersGet**](PAMProviderApi.md#PamProvidersGet) | **Get** /PamProviders | Returns all PAM providers according to the provided filter and output parameters
[**PamProvidersIdDelete**](PAMProviderApi.md#PamProvidersIdDelete) | **Delete** /PamProviders/{id} | Deletes a PAM Provider
[**PamProvidersIdGet**](PAMProviderApi.md#PamProvidersIdGet) | **Get** /PamProviders/{id} | Returns a single PAM Provider that matches the associated id
[**PamProvidersPost**](PAMProviderApi.md#PamProvidersPost) | **Post** /PamProviders | Creates a new PAM provider with the associated properties
[**PamProvidersPut**](PAMProviderApi.md#PamProvidersPut) | **Put** /PamProviders | Updates an existing PAM provider according to the provided properties
[**PamProvidersTypesGet**](PAMProviderApi.md#PamProvidersTypesGet) | **Get** /PamProviders/Types | Returns all PAM providers in the Keyfactor instance
[**PamProvidersTypesPost**](PAMProviderApi.md#PamProvidersTypesPost) | **Post** /PamProviders/Types | Creates a new PAM provider type with the associated properties



## PamProvidersGet

> []KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy PamProvidersGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    queryString := "queryString_example" // string | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) (optional)
    pageReturned := int32(56) // int32 | The current page within the result set to be returned (optional)
    returnLimit := int32(56) // int32 | Maximum number of records to be returned in a single call (optional)
    sortField := "sortField_example" // string | Field by which the results should be sorted (view results via Management Portal for sortable columns) (optional)
    sortAscending := openapiclient.Keyfactor.Common.QueryableExtensionsSortOrder(0) // KeyfactorCommonQueryableExtensionsSortOrder | Field sort direction [0=ascending, 1=descending] (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PAMProviderApi.PamProvidersGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PAMProviderApi.PamProvidersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PamProvidersGet`: []KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy
    fmt.Fprintf(os.Stdout, "Response from `PAMProviderApi.PamProvidersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPamProvidersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **queryString** | **string** | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) | 
 **pageReturned** | **int32** | The current page within the result set to be returned | 
 **returnLimit** | **int32** | Maximum number of records to be returned in a single call | 
 **sortField** | **string** | Field by which the results should be sorted (view results via Management Portal for sortable columns) | 
 **sortAscending** | [**KeyfactorCommonQueryableExtensionsSortOrder**](KeyfactorCommonQueryableExtensionsSortOrder.md) | Field sort direction [0&#x3D;ascending, 1&#x3D;descending] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy**](KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PamProvidersIdDelete

> PamProvidersIdDelete(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PAMProviderApi.PamProvidersIdDelete(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PAMProviderApi.PamProvidersIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPamProvidersIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

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


## PamProvidersIdGet

> KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy PamProvidersIdGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PAMProviderApi.PamProvidersIdGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PAMProviderApi.PamProvidersIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PamProvidersIdGet`: KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy
    fmt.Fprintf(os.Stdout, "Response from `PAMProviderApi.PamProvidersIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Keyfactor identifier of the PAM provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiPamProvidersIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy**](KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PamProvidersPost

> KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy PamProvidersPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequest(keyfactorWebKeyfactorApiModelsPAMProviderCreateRequest).Execute()

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
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsPAMProviderCreateRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsPAMProviderCreateRequest("Name_example", *openapiclient.NewKeyfactorWebKeyfactorApiModelsPAMProviderCreateRequestProviderType()) // KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequest | PAM provider properties to be used (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PAMProviderApi.PamProvidersPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequest(keyfactorWebKeyfactorApiModelsPAMProviderCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PAMProviderApi.PamProvidersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PamProvidersPost`: KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy
    fmt.Fprintf(os.Stdout, "Response from `PAMProviderApi.PamProvidersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPamProvidersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsPAMProviderCreateRequest** | [**KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequest**](KeyfactorWebKeyfactorApiModelsPAMProviderCreateRequest.md) | PAM provider properties to be used | 

### Return type

[**KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy**](KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patchjson, application/json, text/json, application/*json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PamProvidersPut

> KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy PamProvidersPut(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy(keyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy).Execute()

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
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy := *openapiclient.NewKeyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy(int32(123), "Name_example", *openapiclient.NewKeyfactorWebKeyfactorApiModelsPAMProviderCreateRequestProviderType()) // KeyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy | PAM provider properties to be used (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PAMProviderApi.PamProvidersPut(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy(keyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PAMProviderApi.PamProvidersPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PamProvidersPut`: KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy
    fmt.Fprintf(os.Stdout, "Response from `PAMProviderApi.PamProvidersPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPamProvidersPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy** | [**KeyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy**](KeyfactorWebKeyfactorApiModelsPAMProviderUpdateRequestLegacy.md) | PAM provider properties to be used | 

### Return type

[**KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy**](KeyfactorWebKeyfactorApiModelsPAMProviderResponseLegacy.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patchjson, application/json, text/json, application/*json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PamProvidersTypesGet

> []KeyfactorWebKeyfactorApiModelsPAMProviderTypeResponse PamProvidersTypesGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PAMProviderApi.PamProvidersTypesGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PAMProviderApi.PamProvidersTypesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PamProvidersTypesGet`: []KeyfactorWebKeyfactorApiModelsPAMProviderTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `PAMProviderApi.PamProvidersTypesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPamProvidersTypesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]KeyfactorWebKeyfactorApiModelsPAMProviderTypeResponse**](KeyfactorWebKeyfactorApiModelsPAMProviderTypeResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PamProvidersTypesPost

> KeyfactorWebKeyfactorApiModelsPAMProviderTypeResponse PamProvidersTypesPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsPAMProviderTypeCreateRequest(keyfactorWebKeyfactorApiModelsPAMProviderTypeCreateRequest).Execute()

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
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsPAMProviderTypeCreateRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsPAMProviderTypeCreateRequest("Name_example") // KeyfactorWebKeyfactorApiModelsPAMProviderTypeCreateRequest | PAM provider type properties to be used (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PAMProviderApi.PamProvidersTypesPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsPAMProviderTypeCreateRequest(keyfactorWebKeyfactorApiModelsPAMProviderTypeCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PAMProviderApi.PamProvidersTypesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PamProvidersTypesPost`: KeyfactorWebKeyfactorApiModelsPAMProviderTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `PAMProviderApi.PamProvidersTypesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPamProvidersTypesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsPAMProviderTypeCreateRequest** | [**KeyfactorWebKeyfactorApiModelsPAMProviderTypeCreateRequest**](KeyfactorWebKeyfactorApiModelsPAMProviderTypeCreateRequest.md) | PAM provider type properties to be used | 

### Return type

[**KeyfactorWebKeyfactorApiModelsPAMProviderTypeResponse**](KeyfactorWebKeyfactorApiModelsPAMProviderTypeResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patchjson, application/json, text/json, application/*json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


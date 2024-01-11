# \IdentityProviderApi

All URIs are relative to */Keyfactor/API*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IdentityProvidersGet**](IdentityProviderApi.md#IdentityProvidersGet) | **Get** /IdentityProviders | Returns all OAuth identity providers according to the provided filter and output parameters.
[**IdentityProvidersIdGet**](IdentityProviderApi.md#IdentityProvidersIdGet) | **Get** /IdentityProviders/{id} | Gets an OAuth identity provider and its parameters.
[**IdentityProvidersIdPut**](IdentityProviderApi.md#IdentityProvidersIdPut) | **Put** /IdentityProviders/{id} | Updates an OAuth identity provider and any provided parameters.
[**IdentityProvidersTypesGet**](IdentityProviderApi.md#IdentityProvidersTypesGet) | **Get** /IdentityProviders/Types | Returns a list of all available identity provider types and corresponding type parameters.



## IdentityProvidersGet

> []KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponse IdentityProvidersGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns all OAuth identity providers according to the provided filter and output parameters.

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    queryString := "queryString_example" // string |  (optional)
    pageReturned := int32(56) // int32 |  (optional)
    returnLimit := int32(56) // int32 |  (optional)
    sortField := "sortField_example" // string |  (optional)
    sortAscending := int32(56) // int32 |  (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderApi.IdentityProvidersGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderApi.IdentityProvidersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityProvidersGet`: []KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponse
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderApi.IdentityProvidersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityProvidersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **queryString** | **string** |  | 
 **pageReturned** | **int32** |  | 
 **returnLimit** | **int32** |  | 
 **sortField** | **string** |  | 
 **sortAscending** | **int32** |  | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponse**](KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityProvidersIdGet

> KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponse IdentityProvidersIdGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Gets an OAuth identity provider and its parameters.

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The Id of the OAuth identity provider to retrieve.
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderApi.IdentityProvidersIdGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderApi.IdentityProvidersIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityProvidersIdGet`: KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponse
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderApi.IdentityProvidersIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Id of the OAuth identity provider to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityProvidersIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponse**](KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityProvidersIdPut

> KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponse IdentityProvidersIdPut(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderUpdateRequest(keyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderUpdateRequest).Execute()

Updates an OAuth identity provider and any provided parameters.

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The Id of the OAuth identity provider to update.
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderUpdateRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderUpdateRequest("AuthenticationScheme_example", "DisplayName_example") // KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderUpdateRequest | A Keyfactor.Web.KeyfactorApi.Models.IdentityProvider.IdentityProviderUpdateRequest with the Id, name, and parameters that for the OAuth identity provider that need updated. (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderApi.IdentityProvidersIdPut(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderUpdateRequest(keyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderApi.IdentityProvidersIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityProvidersIdPut`: KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponse
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderApi.IdentityProvidersIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Id of the OAuth identity provider to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityProvidersIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderUpdateRequest** | [**KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderUpdateRequest**](KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderUpdateRequest.md) | A Keyfactor.Web.KeyfactorApi.Models.IdentityProvider.IdentityProviderUpdateRequest with the Id, name, and parameters that for the OAuth identity provider that need updated. | 

### Return type

[**KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponse**](KeyfactorWebKeyfactorApiModelsIdentityProviderIdentityProviderResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityProvidersTypesGet

> []KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeResponse IdentityProvidersTypesGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns a list of all available identity provider types and corresponding type parameters.

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderApi.IdentityProvidersTypesGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderApi.IdentityProvidersTypesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityProvidersTypesGet`: []KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderApi.IdentityProvidersTypesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityProvidersTypesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeResponse**](KeyfactorWebKeyfactorApiModelsIdentityProviderProviderTypeResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


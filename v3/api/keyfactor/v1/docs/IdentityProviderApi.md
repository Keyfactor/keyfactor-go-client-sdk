# \IdentityProviderApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIdentityProviders**](IdentityProviderApi.md#CreateIdentityProviders) | **POST** /IdentityProviders | Creates an OAuth identity provider and any provided parameters.  The identity provider will be assigned to the Global Permission Set if no PermissionSet is specified in the request and the user is in a security role that belongs to the Global Permission Set.
[**GetIdentityProviders**](IdentityProviderApi.md#GetIdentityProviders) | **GET** /IdentityProviders | Returns all OAuth identity providers according to the provided filter and output parameters and user&#39;s security role assigned permission sets.
[**GetIdentityProvidersById**](IdentityProviderApi.md#GetIdentityProvidersById) | **GET** /IdentityProviders/{id} | Gets an OAuth identity provider and its parameters.
[**GetIdentityProvidersTypes**](IdentityProviderApi.md#GetIdentityProvidersTypes) | **GET** /IdentityProviders/Types | Returns a list of all available identity provider types and corresponding type parameters.
[**UpdateIdentityProvidersById**](IdentityProviderApi.md#UpdateIdentityProvidersById) | **PUT** /IdentityProviders/{id} | Updates an OAuth identity provider and any provided parameters. The identity provider will be assigned to the Global Permission Set if no PermissionSet is specified in the request and the user is in a security role that belongs to the Global Permission Set.



## CreateIdentityProviders

> IdentityProviderIdentityProviderCreateResponse CreateIdentityProviders(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).IdentityProviderIdentityProviderCreateRequest(identityProviderIdentityProviderCreateRequest).Execute()

Creates an OAuth identity provider and any provided parameters.  The identity provider will be assigned to the Global Permission Set if no PermissionSet is specified in the request and the user is in a security role that belongs to the Global Permission Set.

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
    identityProviderIdentityProviderCreateRequest := *openapiclient.NewIdentityProviderIdentityProviderCreateRequest("AuthenticationScheme_example", "DisplayName_example", "ProviderType_example") // IdentityProviderIdentityProviderCreateRequest | A IdentityProvider.IdentityProviderCreateRequest with the Id, name, and parameters that for the OAuth identity provider that is to be created. (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderApi.CreateIdentityProviders(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).IdentityProviderIdentityProviderCreateRequest(identityProviderIdentityProviderCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderApi.CreateIdentityProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIdentityProviders`: IdentityProviderIdentityProviderCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderApi.CreateIdentityProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIdentityProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **identityProviderIdentityProviderCreateRequest** | [**IdentityProviderIdentityProviderCreateRequest**](IdentityProviderIdentityProviderCreateRequest.md) | A IdentityProvider.IdentityProviderCreateRequest with the Id, name, and parameters that for the OAuth identity provider that is to be created. | 

### Return type

[**IdentityProviderIdentityProviderCreateResponse**](IdentityProviderIdentityProviderCreateResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityProviders

> []IdentityProviderIdentityProviderGetResponse GetIdentityProviders(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns all OAuth identity providers according to the provided filter and output parameters and user's security role assigned permission sets.

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
    queryString := "queryString_example" // string | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) (optional)
    pageReturned := int32(56) // int32 | The current page within the result set to be returned (optional)
    returnLimit := int32(56) // int32 | Maximum number of records to be returned in a single call (optional)
    sortField := "sortField_example" // string | Field by which the results should be sorted (view results via Management Portal for sortable columns) (optional)
    sortAscending := openapiclient.Keyfactor.Common.QueryableExtensions.SortOrder(0) // KeyfactorCommonQueryableExtensionsSortOrder | Field sort direction [0=ascending, 1=descending] (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderApi.GetIdentityProviders(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderApi.GetIdentityProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentityProviders`: []IdentityProviderIdentityProviderGetResponse
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderApi.GetIdentityProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **queryString** | **string** | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) | 
 **pageReturned** | **int32** | The current page within the result set to be returned | 
 **returnLimit** | **int32** | Maximum number of records to be returned in a single call | 
 **sortField** | **string** | Field by which the results should be sorted (view results via Management Portal for sortable columns) | 
 **sortAscending** | [**KeyfactorCommonQueryableExtensionsSortOrder**](KeyfactorCommonQueryableExtensionsSortOrder.md) | Field sort direction [0&#x3D;ascending, 1&#x3D;descending] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]IdentityProviderIdentityProviderGetResponse**](IdentityProviderIdentityProviderGetResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityProvidersById

> IdentityProviderIdentityProviderGetResponse GetIdentityProvidersById(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderApi.GetIdentityProvidersById(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderApi.GetIdentityProvidersById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentityProvidersById`: IdentityProviderIdentityProviderGetResponse
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderApi.GetIdentityProvidersById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Id of the OAuth identity provider to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityProvidersByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**IdentityProviderIdentityProviderGetResponse**](IdentityProviderIdentityProviderGetResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityProvidersTypes

> []IdentityProviderProviderTypeResponse GetIdentityProvidersTypes(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderApi.GetIdentityProvidersTypes(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderApi.GetIdentityProvidersTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentityProvidersTypes`: []IdentityProviderProviderTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderApi.GetIdentityProvidersTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityProvidersTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]IdentityProviderProviderTypeResponse**](IdentityProviderProviderTypeResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIdentityProvidersById

> IdentityProviderIdentityProviderUpdateResponse UpdateIdentityProvidersById(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).IdentityProviderIdentityProviderUpdateRequest(identityProviderIdentityProviderUpdateRequest).Execute()

Updates an OAuth identity provider and any provided parameters. The identity provider will be assigned to the Global Permission Set if no PermissionSet is specified in the request and the user is in a security role that belongs to the Global Permission Set.

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    identityProviderIdentityProviderUpdateRequest := *openapiclient.NewIdentityProviderIdentityProviderUpdateRequest("AuthenticationScheme_example", "DisplayName_example") // IdentityProviderIdentityProviderUpdateRequest | A IdentityProvider.IdentityProviderUpdateRequest with the Id, name, and parameters that for the OAuth identity provider that need updated. (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderApi.UpdateIdentityProvidersById(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).IdentityProviderIdentityProviderUpdateRequest(identityProviderIdentityProviderUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderApi.UpdateIdentityProvidersById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIdentityProvidersById`: IdentityProviderIdentityProviderUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderApi.UpdateIdentityProvidersById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Id of the OAuth identity provider to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIdentityProvidersByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **identityProviderIdentityProviderUpdateRequest** | [**IdentityProviderIdentityProviderUpdateRequest**](IdentityProviderIdentityProviderUpdateRequest.md) | A IdentityProvider.IdentityProviderUpdateRequest with the Id, name, and parameters that for the OAuth identity provider that need updated. | 

### Return type

[**IdentityProviderIdentityProviderUpdateResponse**](IdentityProviderIdentityProviderUpdateResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


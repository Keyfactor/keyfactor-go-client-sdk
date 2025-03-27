# \KeyApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSSHKeysMyKey**](KeyApi.md#CreateSSHKeysMyKey) | **POST** /SSH/Keys/MyKey | Generates an SSH Key Pair for the requesting user.
[**DeleteSSHKeysUnmanaged**](KeyApi.md#DeleteSSHKeysUnmanaged) | **DELETE** /SSH/Keys/Unmanaged | Deletes Unmanaged Keys associated with the provided identifiers
[**DeleteSSHKeysUnmanagedById**](KeyApi.md#DeleteSSHKeysUnmanagedById) | **DELETE** /SSH/Keys/Unmanaged/{id} | Deletes Unmanaged Key associated with the provided identifier
[**GetSSHKeysMyKey**](KeyApi.md#GetSSHKeysMyKey) | **GET** /SSH/Keys/MyKey | Returns the current key of the requesting user
[**GetSSHKeysUnmanaged**](KeyApi.md#GetSSHKeysUnmanaged) | **GET** /SSH/Keys/Unmanaged | Returns Unmanaged SSH keys
[**GetSSHKeysUnmanagedById**](KeyApi.md#GetSSHKeysUnmanagedById) | **GET** /SSH/Keys/Unmanaged/{id} | Returns an unmanaged SSH key with provided id.
[**UpdateSSHKeysMyKey**](KeyApi.md#UpdateSSHKeysMyKey) | **PUT** /SSH/Keys/MyKey | Updates the requesting user&#39;s SSH key



## CreateSSHKeysMyKey

> CSSCMSDataModelModelsSSHKeysKeyResponse CreateSSHKeysMyKey(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsSSHKeysKeyGenerationRequest(cSSCMSDataModelModelsSSHKeysKeyGenerationRequest).Execute()

Generates an SSH Key Pair for the requesting user.



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
    cSSCMSDataModelModelsSSHKeysKeyGenerationRequest := *openapiclient.NewCSSCMSDataModelModelsSSHKeysKeyGenerationRequest("KeyType_example", "PrivateKeyFormat_example", int32(123), "Email_example", "Password_example") // CSSCMSDataModelModelsSSHKeysKeyGenerationRequest | Object containing information about the key to be generated (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyApi.CreateSSHKeysMyKey(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsSSHKeysKeyGenerationRequest(cSSCMSDataModelModelsSSHKeysKeyGenerationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyApi.CreateSSHKeysMyKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSSHKeysMyKey`: CSSCMSDataModelModelsSSHKeysKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `KeyApi.CreateSSHKeysMyKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSSHKeysMyKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsSSHKeysKeyGenerationRequest** | [**CSSCMSDataModelModelsSSHKeysKeyGenerationRequest**](CSSCMSDataModelModelsSSHKeysKeyGenerationRequest.md) | Object containing information about the key to be generated | 

### Return type

[**CSSCMSDataModelModelsSSHKeysKeyResponse**](CSSCMSDataModelModelsSSHKeysKeyResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSSHKeysUnmanaged

> DeleteSSHKeysUnmanaged(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).RequestBody(requestBody).Execute()

Deletes Unmanaged Keys associated with the provided identifiers

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
    requestBody := []int32{int32(123)} // []int32 | Keyfactor identifers of the Keys to be deleted (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyApi.DeleteSSHKeysUnmanaged(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyApi.DeleteSSHKeysUnmanaged``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSSHKeysUnmanagedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **requestBody** | **[]int32** | Keyfactor identifers of the Keys to be deleted | 

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


## DeleteSSHKeysUnmanagedById

> DeleteSSHKeysUnmanagedById(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Deletes Unmanaged Key associated with the provided identifier

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
    id := int32(56) // int32 | Keyfactor identifer of the Key to be deleted
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyApi.DeleteSSHKeysUnmanagedById(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyApi.DeleteSSHKeysUnmanagedById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Keyfactor identifer of the Key to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSSHKeysUnmanagedByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
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


## GetSSHKeysMyKey

> CSSCMSDataModelModelsSSHKeysKeyResponse GetSSHKeysMyKey(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).IncludePrivateKey(includePrivateKey).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns the current key of the requesting user

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
    includePrivateKey := true // bool | Whether or not to include the private key. If true, you must supply the X-Keyfactor-Key-Passphrase header (optional) (default to false)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyApi.GetSSHKeysMyKey(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).IncludePrivateKey(includePrivateKey).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyApi.GetSSHKeysMyKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSSHKeysMyKey`: CSSCMSDataModelModelsSSHKeysKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `KeyApi.GetSSHKeysMyKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSSHKeysMyKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **includePrivateKey** | **bool** | Whether or not to include the private key. If true, you must supply the X-Keyfactor-Key-Passphrase header | [default to false]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**CSSCMSDataModelModelsSSHKeysKeyResponse**](CSSCMSDataModelModelsSSHKeysKeyResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSSHKeysUnmanaged

> []CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse GetSSHKeysUnmanaged(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns Unmanaged SSH keys

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
    resp, r, err := apiClient.KeyApi.GetSSHKeysUnmanaged(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyApi.GetSSHKeysUnmanaged``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSSHKeysUnmanaged`: []CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `KeyApi.GetSSHKeysUnmanaged`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSSHKeysUnmanagedRequest struct via the builder pattern


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

[**[]CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse**](CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSSHKeysUnmanagedById

> CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse GetSSHKeysUnmanagedById(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns an unmanaged SSH key with provided id.

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
    id := int32(56) // int32 | The id of the key to get
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyApi.GetSSHKeysUnmanagedById(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyApi.GetSSHKeysUnmanagedById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSSHKeysUnmanagedById`: CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `KeyApi.GetSSHKeysUnmanagedById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The id of the key to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSSHKeysUnmanagedByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse**](CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSSHKeysMyKey

> CSSCMSDataModelModelsSSHKeysKeyResponse UpdateSSHKeysMyKey(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsSSHKeysKeyUpdateRequest(cSSCMSDataModelModelsSSHKeysKeyUpdateRequest).Execute()

Updates the requesting user's SSH key

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
    cSSCMSDataModelModelsSSHKeysKeyUpdateRequest := *openapiclient.NewCSSCMSDataModelModelsSSHKeysKeyUpdateRequest(int32(123), "Email_example") // CSSCMSDataModelModelsSSHKeysKeyUpdateRequest | Updated state of the SSH key (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyApi.UpdateSSHKeysMyKey(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsSSHKeysKeyUpdateRequest(cSSCMSDataModelModelsSSHKeysKeyUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyApi.UpdateSSHKeysMyKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSSHKeysMyKey`: CSSCMSDataModelModelsSSHKeysKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `KeyApi.UpdateSSHKeysMyKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSSHKeysMyKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsSSHKeysKeyUpdateRequest** | [**CSSCMSDataModelModelsSSHKeysKeyUpdateRequest**](CSSCMSDataModelModelsSSHKeysKeyUpdateRequest.md) | Updated state of the SSH key | 

### Return type

[**CSSCMSDataModelModelsSSHKeysKeyResponse**](CSSCMSDataModelModelsSSHKeysKeyResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \KeyApi

All URIs are relative to */Keyfactor/API*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SSHKeysMyKeyGet**](KeyApi.md#SSHKeysMyKeyGet) | **Get** /SSH/Keys/MyKey | Returns the current key of the requesting user
[**SSHKeysMyKeyPost**](KeyApi.md#SSHKeysMyKeyPost) | **Post** /SSH/Keys/MyKey | Generates an SSH Key Pair for the requesting user.
[**SSHKeysMyKeyPut**](KeyApi.md#SSHKeysMyKeyPut) | **Put** /SSH/Keys/MyKey | Updates the requesting user&#39;s SSH key
[**SSHKeysUnmanagedDelete**](KeyApi.md#SSHKeysUnmanagedDelete) | **Delete** /SSH/Keys/Unmanaged | Deletes Unmanaged Keys associated with the provided identifiers
[**SSHKeysUnmanagedGet**](KeyApi.md#SSHKeysUnmanagedGet) | **Get** /SSH/Keys/Unmanaged | Returns Unmanaged SSH keys
[**SSHKeysUnmanagedIdDelete**](KeyApi.md#SSHKeysUnmanagedIdDelete) | **Delete** /SSH/Keys/Unmanaged/{id} | Deletes Unmanaged Key associated with the provided identifier
[**SSHKeysUnmanagedIdGet**](KeyApi.md#SSHKeysUnmanagedIdGet) | **Get** /SSH/Keys/Unmanaged/{id} | Returns an unmanaged SSH key with provided id.



## SSHKeysMyKeyGet

> CSSCMSDataModelModelsSSHKeysKeyResponse SSHKeysMyKeyGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).IncludePrivateKey(includePrivateKey).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    includePrivateKey := true // bool | Whether or not to include the private key. If true, you must supply the X-Keyfactor-Key-Passphrase header (optional) (default to false)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyApi.SSHKeysMyKeyGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).IncludePrivateKey(includePrivateKey).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyApi.SSHKeysMyKeyGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SSHKeysMyKeyGet`: CSSCMSDataModelModelsSSHKeysKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `KeyApi.SSHKeysMyKeyGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSSHKeysMyKeyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
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


## SSHKeysMyKeyPost

> CSSCMSDataModelModelsSSHKeysKeyResponse SSHKeysMyKeyPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsSSHKeysKeyGenerationRequest(cSSCMSDataModelModelsSSHKeysKeyGenerationRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    cSSCMSDataModelModelsSSHKeysKeyGenerationRequest := *openapiclient.NewCSSCMSDataModelModelsSSHKeysKeyGenerationRequest("KeyType_example", "PrivateKeyFormat_example", int32(123), "Email_example", "Password_example") // CSSCMSDataModelModelsSSHKeysKeyGenerationRequest | Object containing information about the key to be generated (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyApi.SSHKeysMyKeyPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsSSHKeysKeyGenerationRequest(cSSCMSDataModelModelsSSHKeysKeyGenerationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyApi.SSHKeysMyKeyPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SSHKeysMyKeyPost`: CSSCMSDataModelModelsSSHKeysKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `KeyApi.SSHKeysMyKeyPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSSHKeysMyKeyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
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


## SSHKeysMyKeyPut

> CSSCMSDataModelModelsSSHKeysKeyResponse SSHKeysMyKeyPut(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsSSHKeysKeyUpdateRequest(cSSCMSDataModelModelsSSHKeysKeyUpdateRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    cSSCMSDataModelModelsSSHKeysKeyUpdateRequest := *openapiclient.NewCSSCMSDataModelModelsSSHKeysKeyUpdateRequest(int32(123), "Email_example") // CSSCMSDataModelModelsSSHKeysKeyUpdateRequest | Updated state of the SSH key (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyApi.SSHKeysMyKeyPut(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsSSHKeysKeyUpdateRequest(cSSCMSDataModelModelsSSHKeysKeyUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyApi.SSHKeysMyKeyPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SSHKeysMyKeyPut`: CSSCMSDataModelModelsSSHKeysKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `KeyApi.SSHKeysMyKeyPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSSHKeysMyKeyPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
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


## SSHKeysUnmanagedDelete

> SSHKeysUnmanagedDelete(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).RequestBody(requestBody).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    requestBody := []int32{int32(123)} // []int32 | Keyfactor identifers of the Keys to be deleted (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyApi.SSHKeysUnmanagedDelete(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyApi.SSHKeysUnmanagedDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSSHKeysUnmanagedDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
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


## SSHKeysUnmanagedGet

> []CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse SSHKeysUnmanagedGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    queryString := "queryString_example" // string |  (optional)
    pageReturned := int32(56) // int32 |  (optional)
    returnLimit := int32(56) // int32 |  (optional)
    sortField := "sortField_example" // string |  (optional)
    sortAscending := openapiclient.Keyfactor.Common.QueryableExtensionsSortOrder(0) // KeyfactorCommonQueryableExtensionsSortOrder |  (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyApi.SSHKeysUnmanagedGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyApi.SSHKeysUnmanagedGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SSHKeysUnmanagedGet`: []CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `KeyApi.SSHKeysUnmanagedGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSSHKeysUnmanagedGetRequest struct via the builder pattern


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

[**[]CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse**](CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SSHKeysUnmanagedIdDelete

> SSHKeysUnmanagedIdDelete(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyApi.SSHKeysUnmanagedIdDelete(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyApi.SSHKeysUnmanagedIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSSHKeysUnmanagedIdDeleteRequest struct via the builder pattern


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


## SSHKeysUnmanagedIdGet

> CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse SSHKeysUnmanagedIdGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyApi.SSHKeysUnmanagedIdGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyApi.SSHKeysUnmanagedIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SSHKeysUnmanagedIdGet`: CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `KeyApi.SSHKeysUnmanagedIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The id of the key to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiSSHKeysUnmanagedIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
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


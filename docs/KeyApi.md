# \KeyApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KeyDeleteUnmanagedKey**](KeyApi.md#KeyDeleteUnmanagedKey) | **Delete** /SSH/Keys/Unmanaged/{id} | Deletes Unmanaged Key associated with the provided identifier
[**KeyDeleteUnmanagedKeys**](KeyApi.md#KeyDeleteUnmanagedKeys) | **Delete** /SSH/Keys/Unmanaged | Deletes Unmanaged Keys associated with the provided identifiers
[**KeyGenerateKey**](KeyApi.md#KeyGenerateKey) | **Post** /SSH/Keys/MyKey | Generates an SSH Key Pair for the requesting user.
[**KeyGetMyKey**](KeyApi.md#KeyGetMyKey) | **Get** /SSH/Keys/MyKey | Returns the current key of the requesting user
[**KeyGetUnmanagedKey**](KeyApi.md#KeyGetUnmanagedKey) | **Get** /SSH/Keys/Unmanaged/{id} | Returns an unmanaged SSH key with provided id.
[**KeyGetUnmanagedKeys**](KeyApi.md#KeyGetUnmanagedKeys) | **Get** /SSH/Keys/Unmanaged | Returns Unmanaged SSH keys
[**KeyUpdate**](KeyApi.md#KeyUpdate) | **Put** /SSH/Keys/MyKey | Updates the requesting user&#39;s SSH key



## KeyDeleteUnmanagedKey

> KeyDeleteUnmanagedKey(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyApi.KeyDeleteUnmanagedKey(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyApi.KeyDeleteUnmanagedKey``: %v\n", err)
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

Other parameters are passed through a pointer to a apiKeyDeleteUnmanagedKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeyDeleteUnmanagedKeys

> KeyDeleteUnmanagedKeys(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Ids(ids).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    ids := []int32{int32(123)} // []int32 | Keyfactor identifers of the Keys to be deleted
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyApi.KeyDeleteUnmanagedKeys(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Ids(ids).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyApi.KeyDeleteUnmanagedKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKeyDeleteUnmanagedKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **ids** | **[]int32** | Keyfactor identifers of the Keys to be deleted | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeyGenerateKey

> ModelsSSHKeysKeyResponse KeyGenerateKey(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).GenerationRequest(generationRequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    generationRequest := *openapiclient.NewModelsSSHKeysKeyGenerationRequest("KeyType_example", "PrivateKeyFormat_example", int32(123), "Email_example", "Password_example") // ModelsSSHKeysKeyGenerationRequest | Object containing information about the key to be generated
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyApi.KeyGenerateKey(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).GenerationRequest(generationRequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyApi.KeyGenerateKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KeyGenerateKey`: ModelsSSHKeysKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `KeyApi.KeyGenerateKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKeyGenerateKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **generationRequest** | [**ModelsSSHKeysKeyGenerationRequest**](ModelsSSHKeysKeyGenerationRequest.md) | Object containing information about the key to be generated | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsSSHKeysKeyResponse**](ModelsSSHKeysKeyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeyGetMyKey

> ModelsSSHKeysKeyResponse KeyGetMyKey(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorKeyPassphrase(xKeyfactorKeyPassphrase).IncludePrivateKey(includePrivateKey).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorKeyPassphrase := "xKeyfactorKeyPassphrase_example" // string | Key password
    includePrivateKey := true // bool | Whether or not to include the private key. If true, you must supply the X-Keyfactor-Key-Passphrase header (optional)
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyApi.KeyGetMyKey(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorKeyPassphrase(xKeyfactorKeyPassphrase).IncludePrivateKey(includePrivateKey).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyApi.KeyGetMyKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KeyGetMyKey`: ModelsSSHKeysKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `KeyApi.KeyGetMyKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKeyGetMyKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorKeyPassphrase** | **string** | Key password | 
 **includePrivateKey** | **bool** | Whether or not to include the private key. If true, you must supply the X-Keyfactor-Key-Passphrase header | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsSSHKeysKeyResponse**](ModelsSSHKeysKeyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeyGetUnmanagedKey

> ModelsSSHKeysUnmanagedKeyResponse KeyGetUnmanagedKey(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyApi.KeyGetUnmanagedKey(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyApi.KeyGetUnmanagedKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KeyGetUnmanagedKey`: ModelsSSHKeysUnmanagedKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `KeyApi.KeyGetUnmanagedKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The id of the key to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiKeyGetUnmanagedKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsSSHKeysUnmanagedKeyResponse**](ModelsSSHKeysUnmanagedKeyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeyGetUnmanagedKeys

> []ModelsSSHKeysUnmanagedKeyResponse KeyGetUnmanagedKeys(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PqQueryString(pqQueryString).PqPageReturned(pqPageReturned).PqReturnLimit(pqReturnLimit).PqSortField(pqSortField).PqSortAscending(pqSortAscending).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")
    pqQueryString := "pqQueryString_example" // string | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) (optional)
    pqPageReturned := int32(56) // int32 | The current page within the result set to be returned (optional)
    pqReturnLimit := int32(56) // int32 | Maximum number of records to be returned in a single call (optional)
    pqSortField := "pqSortField_example" // string | Field by which the results should be sorted (view results via Management Portal for sortable columns) (optional)
    pqSortAscending := int32(56) // int32 | Field sort direction [0=ascending, 1=descending] (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyApi.KeyGetUnmanagedKeys(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PqQueryString(pqQueryString).PqPageReturned(pqPageReturned).PqReturnLimit(pqReturnLimit).PqSortField(pqSortField).PqSortAscending(pqSortAscending).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyApi.KeyGetUnmanagedKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KeyGetUnmanagedKeys`: []ModelsSSHKeysUnmanagedKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `KeyApi.KeyGetUnmanagedKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKeyGetUnmanagedKeysRequest struct via the builder pattern


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

[**[]ModelsSSHKeysUnmanagedKeyResponse**](ModelsSSHKeysUnmanagedKeyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeyUpdate

> ModelsSSHKeysKeyResponse KeyUpdate(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).KeyUpdateRequest(keyUpdateRequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    keyUpdateRequest := *openapiclient.NewModelsSSHKeysKeyUpdateRequest(int32(123), "Email_example") // ModelsSSHKeysKeyUpdateRequest | Updated state of the SSH key
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyApi.KeyUpdate(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).KeyUpdateRequest(keyUpdateRequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyApi.KeyUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KeyUpdate`: ModelsSSHKeysKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `KeyApi.KeyUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKeyUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **keyUpdateRequest** | [**ModelsSSHKeysKeyUpdateRequest**](ModelsSSHKeysKeyUpdateRequest.md) | Updated state of the SSH key | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsSSHKeysKeyResponse**](ModelsSSHKeysKeyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


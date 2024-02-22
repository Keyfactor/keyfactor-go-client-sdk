# \ServiceAccountApi

All URIs are relative to */Keyfactor/API*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SSHServiceAccountsDelete**](ServiceAccountApi.md#SSHServiceAccountsDelete) | **Delete** /SSH/ServiceAccounts | Deletes Service Accounts associated with the provided identifiers
[**SSHServiceAccountsGet**](ServiceAccountApi.md#SSHServiceAccountsGet) | **Get** /SSH/ServiceAccounts | Returns all ServiceAccounts according to the provided filter parameters
[**SSHServiceAccountsIdDelete**](ServiceAccountApi.md#SSHServiceAccountsIdDelete) | **Delete** /SSH/ServiceAccounts/{id} | Deletes a ServiceAccount associated with the provided identifier
[**SSHServiceAccountsIdGet**](ServiceAccountApi.md#SSHServiceAccountsIdGet) | **Get** /SSH/ServiceAccounts/{id} | Returns a ServiceAccount associated with the provided identifier
[**SSHServiceAccountsKeyIdGet**](ServiceAccountApi.md#SSHServiceAccountsKeyIdGet) | **Get** /SSH/ServiceAccounts/Key/{id} | Returns an SSH key with or without private key based on the provided parameters.
[**SSHServiceAccountsPost**](ServiceAccountApi.md#SSHServiceAccountsPost) | **Post** /SSH/ServiceAccounts | Creates a ServiceAccount with the provided properties
[**SSHServiceAccountsPut**](ServiceAccountApi.md#SSHServiceAccountsPut) | **Put** /SSH/ServiceAccounts | Updates an SSH key for a specified service account.
[**SSHServiceAccountsRotateIdPost**](ServiceAccountApi.md#SSHServiceAccountsRotateIdPost) | **Post** /SSH/ServiceAccounts/Rotate/{id} | Rotate an SSH key for a specified service account.



## SSHServiceAccountsDelete

> SSHServiceAccountsDelete(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).RequestBody(requestBody).Execute()

Deletes Service Accounts associated with the provided identifiers

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
    requestBody := []int32{int32(123)} // []int32 | Keyfactor identifers of the ServiceAccounts to be deleted (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountApi.SSHServiceAccountsDelete(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountApi.SSHServiceAccountsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSSHServiceAccountsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **requestBody** | **[]int32** | Keyfactor identifers of the ServiceAccounts to be deleted | 

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


## SSHServiceAccountsGet

> []ModelsSSHServiceAccountsServiceAccountResponse SSHServiceAccountsGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns all ServiceAccounts according to the provided filter parameters

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
    resp, r, err := apiClient.ServiceAccountApi.SSHServiceAccountsGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountApi.SSHServiceAccountsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SSHServiceAccountsGet`: []ModelsSSHServiceAccountsServiceAccountResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountApi.SSHServiceAccountsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSSHServiceAccountsGetRequest struct via the builder pattern


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

[**[]ModelsSSHServiceAccountsServiceAccountResponse**](ModelsSSHServiceAccountsServiceAccountResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SSHServiceAccountsIdDelete

> SSHServiceAccountsIdDelete(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Deletes a ServiceAccount associated with the provided identifier

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
    id := int32(56) // int32 | Keyfactor identifer of the ServiceAccount to be deleted
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountApi.SSHServiceAccountsIdDelete(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountApi.SSHServiceAccountsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Keyfactor identifer of the ServiceAccount to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiSSHServiceAccountsIdDeleteRequest struct via the builder pattern


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


## SSHServiceAccountsIdGet

> ModelsSSHServiceAccountsServiceAccountResponse SSHServiceAccountsIdGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns a ServiceAccount associated with the provided identifier

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
    id := int32(56) // int32 | Keyfactor identifier of the ServiceAccount
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountApi.SSHServiceAccountsIdGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountApi.SSHServiceAccountsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SSHServiceAccountsIdGet`: ModelsSSHServiceAccountsServiceAccountResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountApi.SSHServiceAccountsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Keyfactor identifier of the ServiceAccount | 

### Other Parameters

Other parameters are passed through a pointer to a apiSSHServiceAccountsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**ModelsSSHServiceAccountsServiceAccountResponse**](ModelsSSHServiceAccountsServiceAccountResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SSHServiceAccountsKeyIdGet

> ModelsSSHKeysKeyResponse SSHServiceAccountsKeyIdGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).IncludePrivateKey(includePrivateKey).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns an SSH key with or without private key based on the provided parameters.

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
    id := int32(56) // int32 | The id of the service account to obtain information on
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    includePrivateKey := true // bool | Whether or not to include the private key in the response (optional) (default to false)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountApi.SSHServiceAccountsKeyIdGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).IncludePrivateKey(includePrivateKey).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountApi.SSHServiceAccountsKeyIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SSHServiceAccountsKeyIdGet`: ModelsSSHKeysKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountApi.SSHServiceAccountsKeyIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The id of the service account to obtain information on | 

### Other Parameters

Other parameters are passed through a pointer to a apiSSHServiceAccountsKeyIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **includePrivateKey** | **bool** | Whether or not to include the private key in the response | [default to false]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**ModelsSSHKeysKeyResponse**](ModelsSSHKeysKeyResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SSHServiceAccountsPost

> ModelsSSHServiceAccountsServiceAccountResponse SSHServiceAccountsPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).ModelsSSHServiceAccountsServiceAccountCreationRequest(modelsSSHServiceAccountsServiceAccountCreationRequest).Execute()

Creates a ServiceAccount with the provided properties

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
    modelsSSHServiceAccountsServiceAccountCreationRequest := *openapiclient.NewModelsSSHServiceAccountsServiceAccountCreationRequest(*openapiclient.NewModelsSSHKeysKeyGenerationRequest("KeyType_example", "PrivateKeyFormat_example", int32(123), "Email_example", "Password_example"), *openapiclient.NewModelsSSHServiceAccountsServiceAccountUserCreationRequest("Username_example"), "ClientHostname_example", "ServerGroupId_example") // ModelsSSHServiceAccountsServiceAccountCreationRequest | ServiceAccount properties to be applied to the new ServiceAccount (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountApi.SSHServiceAccountsPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).ModelsSSHServiceAccountsServiceAccountCreationRequest(modelsSSHServiceAccountsServiceAccountCreationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountApi.SSHServiceAccountsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SSHServiceAccountsPost`: ModelsSSHServiceAccountsServiceAccountResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountApi.SSHServiceAccountsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSSHServiceAccountsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **modelsSSHServiceAccountsServiceAccountCreationRequest** | [**ModelsSSHServiceAccountsServiceAccountCreationRequest**](ModelsSSHServiceAccountsServiceAccountCreationRequest.md) | ServiceAccount properties to be applied to the new ServiceAccount | 

### Return type

[**ModelsSSHServiceAccountsServiceAccountResponse**](ModelsSSHServiceAccountsServiceAccountResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SSHServiceAccountsPut

> ModelsSSHServiceAccountsServiceAccountResponse SSHServiceAccountsPut(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).ModelsSSHServiceAccountsServiceAccountUpdateRequest(modelsSSHServiceAccountsServiceAccountUpdateRequest).Execute()

Updates an SSH key for a specified service account.

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
    modelsSSHServiceAccountsServiceAccountUpdateRequest := *openapiclient.NewModelsSSHServiceAccountsServiceAccountUpdateRequest(*openapiclient.NewModelsSSHKeysKeyUpdateRequest(int32(123), "Email_example"), int32(123)) // ModelsSSHServiceAccountsServiceAccountUpdateRequest | The id of the service account and the updated state of the SSH key. (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountApi.SSHServiceAccountsPut(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).ModelsSSHServiceAccountsServiceAccountUpdateRequest(modelsSSHServiceAccountsServiceAccountUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountApi.SSHServiceAccountsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SSHServiceAccountsPut`: ModelsSSHServiceAccountsServiceAccountResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountApi.SSHServiceAccountsPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSSHServiceAccountsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **modelsSSHServiceAccountsServiceAccountUpdateRequest** | [**ModelsSSHServiceAccountsServiceAccountUpdateRequest**](ModelsSSHServiceAccountsServiceAccountUpdateRequest.md) | The id of the service account and the updated state of the SSH key. | 

### Return type

[**ModelsSSHServiceAccountsServiceAccountResponse**](ModelsSSHServiceAccountsServiceAccountResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SSHServiceAccountsRotateIdPost

> ModelsSSHKeysKeyResponse SSHServiceAccountsRotateIdPost(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).ModelsSSHKeysKeyGenerationRequest(modelsSSHKeysKeyGenerationRequest).Execute()

Rotate an SSH key for a specified service account.

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
    id := int32(56) // int32 | The id of the service account and the updated state of the SSH key.
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    modelsSSHKeysKeyGenerationRequest := *openapiclient.NewModelsSSHKeysKeyGenerationRequest("KeyType_example", "PrivateKeyFormat_example", int32(123), "Email_example", "Password_example") // ModelsSSHKeysKeyGenerationRequest |  (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountApi.SSHServiceAccountsRotateIdPost(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).ModelsSSHKeysKeyGenerationRequest(modelsSSHKeysKeyGenerationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountApi.SSHServiceAccountsRotateIdPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SSHServiceAccountsRotateIdPost`: ModelsSSHKeysKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountApi.SSHServiceAccountsRotateIdPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The id of the service account and the updated state of the SSH key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSSHServiceAccountsRotateIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **modelsSSHKeysKeyGenerationRequest** | [**ModelsSSHKeysKeyGenerationRequest**](ModelsSSHKeysKeyGenerationRequest.md) |  | 

### Return type

[**ModelsSSHKeysKeyResponse**](ModelsSSHKeysKeyResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


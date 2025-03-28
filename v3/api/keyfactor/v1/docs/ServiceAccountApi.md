# \ServiceAccountApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSSHServiceAccounts**](ServiceAccountApi.md#CreateSSHServiceAccounts) | **POST** /SSH/ServiceAccounts | Creates a ServiceAccount with the provided properties
[**CreateSSHServiceAccountsRotateById**](ServiceAccountApi.md#CreateSSHServiceAccountsRotateById) | **POST** /SSH/ServiceAccounts/Rotate/{id} | Rotate an SSH key for a specified service account.
[**DeleteSSHServiceAccounts**](ServiceAccountApi.md#DeleteSSHServiceAccounts) | **DELETE** /SSH/ServiceAccounts | Deletes Service Accounts associated with the provided identifiers
[**DeleteSSHServiceAccountsById**](ServiceAccountApi.md#DeleteSSHServiceAccountsById) | **DELETE** /SSH/ServiceAccounts/{id} | Deletes a ServiceAccount associated with the provided identifier
[**GetSSHServiceAccounts**](ServiceAccountApi.md#GetSSHServiceAccounts) | **GET** /SSH/ServiceAccounts | Returns all ServiceAccounts according to the provided filter parameters
[**GetSSHServiceAccountsById**](ServiceAccountApi.md#GetSSHServiceAccountsById) | **GET** /SSH/ServiceAccounts/{id} | Returns a ServiceAccount associated with the provided identifier
[**GetSSHServiceAccountsKeyById**](ServiceAccountApi.md#GetSSHServiceAccountsKeyById) | **GET** /SSH/ServiceAccounts/Key/{id} | Returns an SSH key with or without private key based on the provided parameters.
[**UpdateSSHServiceAccounts**](ServiceAccountApi.md#UpdateSSHServiceAccounts) | **PUT** /SSH/ServiceAccounts | Updates an SSH key for a specified service account.



## CreateSSHServiceAccounts

> CSSCMSDataModelModelsSSHServiceAccountsServiceAccountResponse NewCreateSSHServiceAccountsRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsSSHServiceAccountsServiceAccountCreationRequest(cSSCMSDataModelModelsSSHServiceAccountsServiceAccountCreationRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    cSSCMSDataModelModelsSSHServiceAccountsServiceAccountCreationRequest := *openapiclient.NewCSSCMSDataModelModelsSSHServiceAccountsServiceAccountCreationRequest(*openapiclient.NewCSSCMSDataModelModelsSSHKeysKeyGenerationRequest("KeyType_example", "PrivateKeyFormat_example", int32(123), "Email_example", "Password_example"), *openapiclient.NewCSSCMSDataModelModelsSSHServiceAccountsServiceAccountUserCreationRequest("Username_example"), "ClientHostname_example", "ServerGroupId_example") // CSSCMSDataModelModelsSSHServiceAccountsServiceAccountCreationRequest | ServiceAccount properties to be applied to the new ServiceAccount (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountApi.NewCreateSSHServiceAccountsRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsSSHServiceAccountsServiceAccountCreationRequest(cSSCMSDataModelModelsSSHServiceAccountsServiceAccountCreationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountApi.CreateSSHServiceAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSSHServiceAccounts`: CSSCMSDataModelModelsSSHServiceAccountsServiceAccountResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountApi.CreateSSHServiceAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSSHServiceAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsSSHServiceAccountsServiceAccountCreationRequest** | [**CSSCMSDataModelModelsSSHServiceAccountsServiceAccountCreationRequest**](CSSCMSDataModelModelsSSHServiceAccountsServiceAccountCreationRequest.md) | ServiceAccount properties to be applied to the new ServiceAccount | 

### Return type

[**CSSCMSDataModelModelsSSHServiceAccountsServiceAccountResponse**](CSSCMSDataModelModelsSSHServiceAccountsServiceAccountResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSSHServiceAccountsRotateById

> CSSCMSDataModelModelsSSHKeysKeyResponse NewCreateSSHServiceAccountsRotateByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsSSHKeysKeyGenerationRequest(cSSCMSDataModelModelsSSHKeysKeyGenerationRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    cSSCMSDataModelModelsSSHKeysKeyGenerationRequest := *openapiclient.NewCSSCMSDataModelModelsSSHKeysKeyGenerationRequest("KeyType_example", "PrivateKeyFormat_example", int32(123), "Email_example", "Password_example") // CSSCMSDataModelModelsSSHKeysKeyGenerationRequest |  (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountApi.NewCreateSSHServiceAccountsRotateByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsSSHKeysKeyGenerationRequest(cSSCMSDataModelModelsSSHKeysKeyGenerationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountApi.CreateSSHServiceAccountsRotateById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSSHServiceAccountsRotateById`: CSSCMSDataModelModelsSSHKeysKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountApi.CreateSSHServiceAccountsRotateById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The id of the service account and the updated state of the SSH key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSSHServiceAccountsRotateByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsSSHKeysKeyGenerationRequest** | [**CSSCMSDataModelModelsSSHKeysKeyGenerationRequest**](CSSCMSDataModelModelsSSHKeysKeyGenerationRequest.md) |  | 

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


## DeleteSSHServiceAccounts

> NewDeleteSSHServiceAccountsRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).RequestBody(requestBody).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    requestBody := []int32{int32(123)} // []int32 | Keyfactor identifers of the ServiceAccounts to be deleted (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountApi.NewDeleteSSHServiceAccountsRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountApi.DeleteSSHServiceAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSSHServiceAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
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


## DeleteSSHServiceAccountsById

> NewDeleteSSHServiceAccountsByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountApi.NewDeleteSSHServiceAccountsByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountApi.DeleteSSHServiceAccountsById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteSSHServiceAccountsByIdRequest struct via the builder pattern


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


## GetSSHServiceAccounts

> []CSSCMSDataModelModelsSSHServiceAccountsServiceAccountResponse NewGetSSHServiceAccountsRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    queryString := "queryString_example" // string |  (optional)
    pageReturned := int32(56) // int32 |  (optional)
    returnLimit := int32(56) // int32 |  (optional)
    sortField := "sortField_example" // string |  (optional)
    sortAscending := openapiclient.Keyfactor.Common.QueryableExtensions.SortOrder(0) // KeyfactorCommonQueryableExtensionsSortOrder |  (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountApi.NewGetSSHServiceAccountsRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountApi.GetSSHServiceAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSSHServiceAccounts`: []CSSCMSDataModelModelsSSHServiceAccountsServiceAccountResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountApi.GetSSHServiceAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSSHServiceAccountsRequest struct via the builder pattern


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

[**[]CSSCMSDataModelModelsSSHServiceAccountsServiceAccountResponse**](CSSCMSDataModelModelsSSHServiceAccountsServiceAccountResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSSHServiceAccountsById

> CSSCMSDataModelModelsSSHServiceAccountsServiceAccountResponse NewGetSSHServiceAccountsByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountApi.NewGetSSHServiceAccountsByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountApi.GetSSHServiceAccountsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSSHServiceAccountsById`: CSSCMSDataModelModelsSSHServiceAccountsServiceAccountResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountApi.GetSSHServiceAccountsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Keyfactor identifier of the ServiceAccount | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSSHServiceAccountsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**CSSCMSDataModelModelsSSHServiceAccountsServiceAccountResponse**](CSSCMSDataModelModelsSSHServiceAccountsServiceAccountResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSSHServiceAccountsKeyById

> CSSCMSDataModelModelsSSHKeysKeyResponse NewGetSSHServiceAccountsKeyByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).IncludePrivateKey(includePrivateKey).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    includePrivateKey := true // bool | Whether or not to include the private key in the response (optional) (default to false)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountApi.NewGetSSHServiceAccountsKeyByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).IncludePrivateKey(includePrivateKey).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountApi.GetSSHServiceAccountsKeyById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSSHServiceAccountsKeyById`: CSSCMSDataModelModelsSSHKeysKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountApi.GetSSHServiceAccountsKeyById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The id of the service account to obtain information on | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSSHServiceAccountsKeyByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **includePrivateKey** | **bool** | Whether or not to include the private key in the response | [default to false]
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


## UpdateSSHServiceAccounts

> CSSCMSDataModelModelsSSHServiceAccountsServiceAccountResponse NewUpdateSSHServiceAccountsRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsSSHServiceAccountsServiceAccountUpdateRequest(cSSCMSDataModelModelsSSHServiceAccountsServiceAccountUpdateRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    cSSCMSDataModelModelsSSHServiceAccountsServiceAccountUpdateRequest := *openapiclient.NewCSSCMSDataModelModelsSSHServiceAccountsServiceAccountUpdateRequest(*openapiclient.NewCSSCMSDataModelModelsSSHKeysKeyUpdateRequest(int32(123), "Email_example"), int32(123)) // CSSCMSDataModelModelsSSHServiceAccountsServiceAccountUpdateRequest | The id of the service account and the updated state of the SSH key. (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountApi.NewUpdateSSHServiceAccountsRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsSSHServiceAccountsServiceAccountUpdateRequest(cSSCMSDataModelModelsSSHServiceAccountsServiceAccountUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountApi.UpdateSSHServiceAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSSHServiceAccounts`: CSSCMSDataModelModelsSSHServiceAccountsServiceAccountResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountApi.UpdateSSHServiceAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSSHServiceAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsSSHServiceAccountsServiceAccountUpdateRequest** | [**CSSCMSDataModelModelsSSHServiceAccountsServiceAccountUpdateRequest**](CSSCMSDataModelModelsSSHServiceAccountsServiceAccountUpdateRequest.md) | The id of the service account and the updated state of the SSH key. | 

### Return type

[**CSSCMSDataModelModelsSSHServiceAccountsServiceAccountResponse**](CSSCMSDataModelModelsSSHServiceAccountsServiceAccountResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


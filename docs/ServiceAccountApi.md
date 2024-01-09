# \ServiceAccountApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServiceAccountCreateServiceAccount**](ServiceAccountApi.md#ServiceAccountCreateServiceAccount) | **Post** /SSH/ServiceAccounts | Creates a ServiceAccount with the provided properties
[**ServiceAccountDeleteServiceAccount**](ServiceAccountApi.md#ServiceAccountDeleteServiceAccount) | **Delete** /SSH/ServiceAccounts/{id} | Deletes a ServiceAccount associated with the provided identifier
[**ServiceAccountDeleteServiceAccounts**](ServiceAccountApi.md#ServiceAccountDeleteServiceAccounts) | **Delete** /SSH/ServiceAccounts | Deletes Service Accounts associated with the provided identifiers
[**ServiceAccountGet**](ServiceAccountApi.md#ServiceAccountGet) | **Get** /SSH/ServiceAccounts/{id} | Returns a ServiceAccount associated with the provided identifier
[**ServiceAccountGetServiceAccountKey**](ServiceAccountApi.md#ServiceAccountGetServiceAccountKey) | **Get** /SSH/ServiceAccounts/Key/{id} | Returns an SSH key with or without private key based on the provided parameters.
[**ServiceAccountQueryServiceAccounts**](ServiceAccountApi.md#ServiceAccountQueryServiceAccounts) | **Get** /SSH/ServiceAccounts | Returns all ServiceAccounts according to the provided filter parameters
[**ServiceAccountRotateServiceAccountKey**](ServiceAccountApi.md#ServiceAccountRotateServiceAccountKey) | **Post** /SSH/ServiceAccounts/Rotate/{id} | Rotate an SSH key for a specified service account.
[**ServiceAccountUpdateServiceAccount**](ServiceAccountApi.md#ServiceAccountUpdateServiceAccount) | **Put** /SSH/ServiceAccounts | Updates an SSH key for a specified service account.



## ServiceAccountCreateServiceAccount

> ModelsSSHServiceAccountsServiceAccountResponse ServiceAccountCreateServiceAccount(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).ServiceAccount(serviceAccount).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    serviceAccount := *openapiclient.NewModelsSSHServiceAccountsServiceAccountCreationRequest(*openapiclient.NewModelsSSHKeysKeyGenerationRequest("KeyType_example", "PrivateKeyFormat_example", int32(123), "Email_example", "Password_example"), *openapiclient.NewModelsSSHServiceAccountsServiceAccountUserCreationRequest("Username_example"), "ClientHostname_example", "00000000-0000-0000-0000-000000000000") // ModelsSSHServiceAccountsServiceAccountCreationRequest | ServiceAccount properties to be applied to the new ServiceAccount
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountApi.ServiceAccountCreateServiceAccount(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).ServiceAccount(serviceAccount).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountApi.ServiceAccountCreateServiceAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceAccountCreateServiceAccount`: ModelsSSHServiceAccountsServiceAccountResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountApi.ServiceAccountCreateServiceAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServiceAccountCreateServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **serviceAccount** | [**ModelsSSHServiceAccountsServiceAccountCreationRequest**](ModelsSSHServiceAccountsServiceAccountCreationRequest.md) | ServiceAccount properties to be applied to the new ServiceAccount | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsSSHServiceAccountsServiceAccountResponse**](ModelsSSHServiceAccountsServiceAccountResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceAccountDeleteServiceAccount

> ServiceAccountDeleteServiceAccount(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountApi.ServiceAccountDeleteServiceAccount(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountApi.ServiceAccountDeleteServiceAccount``: %v\n", err)
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

Other parameters are passed through a pointer to a apiServiceAccountDeleteServiceAccountRequest struct via the builder pattern


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


## ServiceAccountDeleteServiceAccounts

> ServiceAccountDeleteServiceAccounts(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Ids(ids).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    ids := []int32{int32(123)} // []int32 | Keyfactor identifers of the ServiceAccounts to be deleted
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountApi.ServiceAccountDeleteServiceAccounts(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Ids(ids).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountApi.ServiceAccountDeleteServiceAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServiceAccountDeleteServiceAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **ids** | **[]int32** | Keyfactor identifers of the ServiceAccounts to be deleted | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceAccountGet

> ModelsSSHServiceAccountsServiceAccountResponse ServiceAccountGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountApi.ServiceAccountGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountApi.ServiceAccountGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceAccountGet`: ModelsSSHServiceAccountsServiceAccountResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountApi.ServiceAccountGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Keyfactor identifier of the ServiceAccount | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceAccountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsSSHServiceAccountsServiceAccountResponse**](ModelsSSHServiceAccountsServiceAccountResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceAccountGetServiceAccountKey

> ModelsSSHKeysKeyResponse ServiceAccountGetServiceAccountKey(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).IncludePrivateKey(includePrivateKey).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    includePrivateKey := true // bool | Whether or not to include the private key in the response (optional)
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountApi.ServiceAccountGetServiceAccountKey(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).IncludePrivateKey(includePrivateKey).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountApi.ServiceAccountGetServiceAccountKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceAccountGetServiceAccountKey`: ModelsSSHKeysKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountApi.ServiceAccountGetServiceAccountKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The id of the service account to obtain information on | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceAccountGetServiceAccountKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **includePrivateKey** | **bool** | Whether or not to include the private key in the response | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsSSHKeysKeyResponse**](ModelsSSHKeysKeyResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceAccountQueryServiceAccounts

> []ModelsSSHServiceAccountsServiceAccountResponse ServiceAccountQueryServiceAccounts(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PqQueryString(pqQueryString).PqPageReturned(pqPageReturned).PqReturnLimit(pqReturnLimit).PqSortField(pqSortField).PqSortAscending(pqSortAscending).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")
    pqQueryString := "pqQueryString_example" // string | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) (optional)
    pqPageReturned := int32(56) // int32 | The current page within the result set to be returned (optional)
    pqReturnLimit := int32(56) // int32 | Maximum number of records to be returned in a single call (optional)
    pqSortField := "pqSortField_example" // string | Field by which the results should be sorted (view results via Management Portal for sortable columns) (optional)
    pqSortAscending := int32(56) // int32 | Field sort direction [0=ascending, 1=descending] (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountApi.ServiceAccountQueryServiceAccounts(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PqQueryString(pqQueryString).PqPageReturned(pqPageReturned).PqReturnLimit(pqReturnLimit).PqSortField(pqSortField).PqSortAscending(pqSortAscending).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountApi.ServiceAccountQueryServiceAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceAccountQueryServiceAccounts`: []ModelsSSHServiceAccountsServiceAccountResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountApi.ServiceAccountQueryServiceAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServiceAccountQueryServiceAccountsRequest struct via the builder pattern


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

[**[]ModelsSSHServiceAccountsServiceAccountResponse**](ModelsSSHServiceAccountsServiceAccountResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceAccountRotateServiceAccountKey

> ModelsSSHKeysKeyResponse ServiceAccountRotateServiceAccountKey(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).RotationRequest(rotationRequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    rotationRequest := *openapiclient.NewModelsSSHKeysKeyGenerationRequest("KeyType_example", "PrivateKeyFormat_example", int32(123), "Email_example", "Password_example") // ModelsSSHKeysKeyGenerationRequest | 
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountApi.ServiceAccountRotateServiceAccountKey(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).RotationRequest(rotationRequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountApi.ServiceAccountRotateServiceAccountKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceAccountRotateServiceAccountKey`: ModelsSSHKeysKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountApi.ServiceAccountRotateServiceAccountKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The id of the service account and the updated state of the SSH key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceAccountRotateServiceAccountKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **rotationRequest** | [**ModelsSSHKeysKeyGenerationRequest**](ModelsSSHKeysKeyGenerationRequest.md) |  | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsSSHKeysKeyResponse**](ModelsSSHKeysKeyResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceAccountUpdateServiceAccount

> ModelsSSHServiceAccountsServiceAccountResponse ServiceAccountUpdateServiceAccount(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).UpdateRequest(updateRequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    updateRequest := *openapiclient.NewModelsSSHServiceAccountsServiceAccountUpdateRequest(*openapiclient.NewModelsSSHKeysKeyUpdateRequest(int32(123), "Email_example"), int32(123)) // ModelsSSHServiceAccountsServiceAccountUpdateRequest | The id of the service account and the updated state of the SSH key.
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountApi.ServiceAccountUpdateServiceAccount(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).UpdateRequest(updateRequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountApi.ServiceAccountUpdateServiceAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceAccountUpdateServiceAccount`: ModelsSSHServiceAccountsServiceAccountResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountApi.ServiceAccountUpdateServiceAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServiceAccountUpdateServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **updateRequest** | [**ModelsSSHServiceAccountsServiceAccountUpdateRequest**](ModelsSSHServiceAccountsServiceAccountUpdateRequest.md) | The id of the service account and the updated state of the SSH key. | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsSSHServiceAccountsServiceAccountResponse**](ModelsSSHServiceAccountsServiceAccountResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


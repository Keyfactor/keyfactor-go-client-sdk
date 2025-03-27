# \ServerApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSSHServers**](ServerApi.md#CreateSSHServers) | **POST** /SSH/Servers | Creates a server with the provided properties
[**CreateSSHServersAccess**](ServerApi.md#CreateSSHServersAccess) | **POST** /SSH/Servers/Access | Updates logons and users with access to those logons for an existing server
[**DeleteSSHServersAccess**](ServerApi.md#DeleteSSHServersAccess) | **DELETE** /SSH/Servers/Access | Updates logons and users with access to those logons for an existing server
[**DeleteSSHServersById**](ServerApi.md#DeleteSSHServersById) | **DELETE** /SSH/Servers/{id} | Deletes a Server associated with the provided identifier
[**GetSSHServers**](ServerApi.md#GetSSHServers) | **GET** /SSH/Servers | Returns all servers according to the provided filter parameters
[**GetSSHServersAccessById**](ServerApi.md#GetSSHServersAccessById) | **GET** /SSH/Servers/Access/{id} | Retrieves logons and users with access to those logons for an existing server
[**GetSSHServersById**](ServerApi.md#GetSSHServersById) | **GET** /SSH/Servers/{id} | Returns a Server associated with the provided identifier
[**UpdateSSHServers**](ServerApi.md#UpdateSSHServers) | **PUT** /SSH/Servers | Updates an existing server with the provided properties



## CreateSSHServers

> CSSCMSDataModelModelsSSHServersServerResponse BuildCreateSSHServersRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsSSHServersServerCreationRequest(cSSCMSDataModelModelsSSHServersServerCreationRequest).Execute()

Creates a server with the provided properties

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
    cSSCMSDataModelModelsSSHServersServerCreationRequest := *openapiclient.NewCSSCMSDataModelModelsSSHServersServerCreationRequest("AgentId_example", "Hostname_example", "ServerGroupId_example") // CSSCMSDataModelModelsSSHServersServerCreationRequest | Server properties to be applied to the newserver (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.BuildCreateSSHServersRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsSSHServersServerCreationRequest(cSSCMSDataModelModelsSSHServersServerCreationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.CreateSSHServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSSHServers`: CSSCMSDataModelModelsSSHServersServerResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerApi.CreateSSHServers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSSHServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsSSHServersServerCreationRequest** | [**CSSCMSDataModelModelsSSHServersServerCreationRequest**](CSSCMSDataModelModelsSSHServersServerCreationRequest.md) | Server properties to be applied to the newserver | 

### Return type

[**CSSCMSDataModelModelsSSHServersServerResponse**](CSSCMSDataModelModelsSSHServersServerResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSSHServersAccess

> CSSCMSDataModelModelsSSHAccessServerAccessResponse BuildCreateSSHServersAccessRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsSSHAccessServerAccessRequest(cSSCMSDataModelModelsSSHAccessServerAccessRequest).Execute()

Updates logons and users with access to those logons for an existing server

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
    cSSCMSDataModelModelsSSHAccessServerAccessRequest := *openapiclient.NewCSSCMSDataModelModelsSSHAccessServerAccessRequest(int32(123), []openapiclient.CSSCMSDataModelModelsSSHAccessLogonUserAccessRequest{*openapiclient.NewCSSCMSDataModelModelsSSHAccessLogonUserAccessRequest()}) // CSSCMSDataModelModelsSSHAccessServerAccessRequest | Logons and users to be applied to the existing server (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.BuildCreateSSHServersAccessRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsSSHAccessServerAccessRequest(cSSCMSDataModelModelsSSHAccessServerAccessRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.CreateSSHServersAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSSHServersAccess`: CSSCMSDataModelModelsSSHAccessServerAccessResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerApi.CreateSSHServersAccess`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSSHServersAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsSSHAccessServerAccessRequest** | [**CSSCMSDataModelModelsSSHAccessServerAccessRequest**](CSSCMSDataModelModelsSSHAccessServerAccessRequest.md) | Logons and users to be applied to the existing server | 

### Return type

[**CSSCMSDataModelModelsSSHAccessServerAccessResponse**](CSSCMSDataModelModelsSSHAccessServerAccessResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSSHServersAccess

> CSSCMSDataModelModelsSSHAccessServerAccessResponse BuildDeleteSSHServersAccessRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsSSHAccessServerAccessRequest(cSSCMSDataModelModelsSSHAccessServerAccessRequest).Execute()

Updates logons and users with access to those logons for an existing server

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
    cSSCMSDataModelModelsSSHAccessServerAccessRequest := *openapiclient.NewCSSCMSDataModelModelsSSHAccessServerAccessRequest(int32(123), []openapiclient.CSSCMSDataModelModelsSSHAccessLogonUserAccessRequest{*openapiclient.NewCSSCMSDataModelModelsSSHAccessLogonUserAccessRequest()}) // CSSCMSDataModelModelsSSHAccessServerAccessRequest | Logons and users to be removed from the existing server (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.BuildDeleteSSHServersAccessRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsSSHAccessServerAccessRequest(cSSCMSDataModelModelsSSHAccessServerAccessRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.DeleteSSHServersAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSSHServersAccess`: CSSCMSDataModelModelsSSHAccessServerAccessResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerApi.DeleteSSHServersAccess`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSSHServersAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsSSHAccessServerAccessRequest** | [**CSSCMSDataModelModelsSSHAccessServerAccessRequest**](CSSCMSDataModelModelsSSHAccessServerAccessRequest.md) | Logons and users to be removed from the existing server | 

### Return type

[**CSSCMSDataModelModelsSSHAccessServerAccessResponse**](CSSCMSDataModelModelsSSHAccessServerAccessResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSSHServersById

> BuildDeleteSSHServersByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Deletes a Server associated with the provided identifier

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
    id := int32(56) // int32 | Keyfactor identifer of the Server to be deleted
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.BuildDeleteSSHServersByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.DeleteSSHServersById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Keyfactor identifer of the Server to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSSHServersByIdRequest struct via the builder pattern


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


## GetSSHServers

> []CSSCMSDataModelModelsSSHServersServerResponse BuildGetSSHServersRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns all servers according to the provided filter parameters

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
    resp, r, err := apiClient.ServerApi.BuildGetSSHServersRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.GetSSHServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSSHServers`: []CSSCMSDataModelModelsSSHServersServerResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerApi.GetSSHServers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSSHServersRequest struct via the builder pattern


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

[**[]CSSCMSDataModelModelsSSHServersServerResponse**](CSSCMSDataModelModelsSSHServersServerResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSSHServersAccessById

> CSSCMSDataModelModelsSSHAccessServerAccessResponse BuildGetSSHServersAccessByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Retrieves logons and users with access to those logons for an existing server

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
    id := int32(56) // int32 | Id of the existing server
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.BuildGetSSHServersAccessByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.GetSSHServersAccessById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSSHServersAccessById`: CSSCMSDataModelModelsSSHAccessServerAccessResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerApi.GetSSHServersAccessById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id of the existing server | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSSHServersAccessByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**CSSCMSDataModelModelsSSHAccessServerAccessResponse**](CSSCMSDataModelModelsSSHAccessServerAccessResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSSHServersById

> CSSCMSDataModelModelsSSHServersServerResponse BuildGetSSHServersByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns a Server associated with the provided identifier

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
    id := int32(56) // int32 | Keyfactor identifier of the Server
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.BuildGetSSHServersByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.GetSSHServersById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSSHServersById`: CSSCMSDataModelModelsSSHServersServerResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerApi.GetSSHServersById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Keyfactor identifier of the Server | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSSHServersByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**CSSCMSDataModelModelsSSHServersServerResponse**](CSSCMSDataModelModelsSSHServersServerResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSSHServers

> CSSCMSDataModelModelsSSHServersServerResponse BuildUpdateSSHServersRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsSSHServersServerUpdateRequest(cSSCMSDataModelModelsSSHServersServerUpdateRequest).Execute()

Updates an existing server with the provided properties

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
    cSSCMSDataModelModelsSSHServersServerUpdateRequest := *openapiclient.NewCSSCMSDataModelModelsSSHServersServerUpdateRequest(int32(123)) // CSSCMSDataModelModelsSSHServersServerUpdateRequest | Server properties to be applied to the existing server (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.BuildUpdateSSHServersRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsSSHServersServerUpdateRequest(cSSCMSDataModelModelsSSHServersServerUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.UpdateSSHServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSSHServers`: CSSCMSDataModelModelsSSHServersServerResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerApi.UpdateSSHServers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSSHServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsSSHServersServerUpdateRequest** | [**CSSCMSDataModelModelsSSHServersServerUpdateRequest**](CSSCMSDataModelModelsSSHServersServerUpdateRequest.md) | Server properties to be applied to the existing server | 

### Return type

[**CSSCMSDataModelModelsSSHServersServerResponse**](CSSCMSDataModelModelsSSHServersServerResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \LogonApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSSHLogons**](LogonApi.md#CreateSSHLogons) | **POST** /SSH/Logons | Creates a logon with the provided properties
[**CreateSSHLogonsAccess**](LogonApi.md#CreateSSHLogonsAccess) | **POST** /SSH/Logons/Access | Updates the users with access to an existing logon
[**DeleteSSHLogonsById**](LogonApi.md#DeleteSSHLogonsById) | **DELETE** /SSH/Logons/{id} | Deletes a Logon associated with the provided identifier
[**GetSSHLogons**](LogonApi.md#GetSSHLogons) | **GET** /SSH/Logons | Returns all Logons according to the provided filter parameters
[**GetSSHLogonsById**](LogonApi.md#GetSSHLogonsById) | **GET** /SSH/Logons/{id} | Fetches a Logon associated with the provided identifier



## CreateSSHLogons

> CSSCMSDataModelModelsSSHLogonsLogonResponse NewCreateSSHLogonsRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsSSHLogonsLogonCreationRequest(cSSCMSDataModelModelsSSHLogonsLogonCreationRequest).Execute()

Creates a logon with the provided properties

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
    cSSCMSDataModelModelsSSHLogonsLogonCreationRequest := *openapiclient.NewCSSCMSDataModelModelsSSHLogonsLogonCreationRequest("Username_example", int32(123)) // CSSCMSDataModelModelsSSHLogonsLogonCreationRequest | Logon properties to be applied to the new logon (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogonApi.NewCreateSSHLogonsRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsSSHLogonsLogonCreationRequest(cSSCMSDataModelModelsSSHLogonsLogonCreationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogonApi.CreateSSHLogons``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSSHLogons`: CSSCMSDataModelModelsSSHLogonsLogonResponse
    fmt.Fprintf(os.Stdout, "Response from `LogonApi.CreateSSHLogons`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSSHLogonsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsSSHLogonsLogonCreationRequest** | [**CSSCMSDataModelModelsSSHLogonsLogonCreationRequest**](CSSCMSDataModelModelsSSHLogonsLogonCreationRequest.md) | Logon properties to be applied to the new logon | 

### Return type

[**CSSCMSDataModelModelsSSHLogonsLogonResponse**](CSSCMSDataModelModelsSSHLogonsLogonResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSSHLogonsAccess

> CSSCMSDataModelModelsSSHAccessLogonUserAccessResponse NewCreateSSHLogonsAccessRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsSSHLogonsLogonAccessRequest(cSSCMSDataModelModelsSSHLogonsLogonAccessRequest).Execute()

Updates the users with access to an existing logon

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
    cSSCMSDataModelModelsSSHLogonsLogonAccessRequest := *openapiclient.NewCSSCMSDataModelModelsSSHLogonsLogonAccessRequest(int32(123)) // CSSCMSDataModelModelsSSHLogonsLogonAccessRequest | Users to add the existing logon (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogonApi.NewCreateSSHLogonsAccessRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsSSHLogonsLogonAccessRequest(cSSCMSDataModelModelsSSHLogonsLogonAccessRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogonApi.CreateSSHLogonsAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSSHLogonsAccess`: CSSCMSDataModelModelsSSHAccessLogonUserAccessResponse
    fmt.Fprintf(os.Stdout, "Response from `LogonApi.CreateSSHLogonsAccess`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSSHLogonsAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsSSHLogonsLogonAccessRequest** | [**CSSCMSDataModelModelsSSHLogonsLogonAccessRequest**](CSSCMSDataModelModelsSSHLogonsLogonAccessRequest.md) | Users to add the existing logon | 

### Return type

[**CSSCMSDataModelModelsSSHAccessLogonUserAccessResponse**](CSSCMSDataModelModelsSSHAccessLogonUserAccessResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSSHLogonsById

> NewDeleteSSHLogonsByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Deletes a Logon associated with the provided identifier

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
    id := int32(56) // int32 | Keyfactor identifer of the Logon to be deleted
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogonApi.NewDeleteSSHLogonsByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogonApi.DeleteSSHLogonsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Keyfactor identifer of the Logon to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSSHLogonsByIdRequest struct via the builder pattern


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


## GetSSHLogons

> []CSSCMSDataModelModelsSSHLogonsLogonQueryResponse NewGetSSHLogonsRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns all Logons according to the provided filter parameters

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
    resp, r, err := apiClient.LogonApi.NewGetSSHLogonsRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogonApi.GetSSHLogons``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSSHLogons`: []CSSCMSDataModelModelsSSHLogonsLogonQueryResponse
    fmt.Fprintf(os.Stdout, "Response from `LogonApi.GetSSHLogons`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSSHLogonsRequest struct via the builder pattern


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

[**[]CSSCMSDataModelModelsSSHLogonsLogonQueryResponse**](CSSCMSDataModelModelsSSHLogonsLogonQueryResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSSHLogonsById

> CSSCMSDataModelModelsSSHLogonsLogonResponse NewGetSSHLogonsByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Fetches a Logon associated with the provided identifier

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
    id := int32(56) // int32 | Keyfactor identifer of the Logon to be Fetched
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogonApi.NewGetSSHLogonsByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogonApi.GetSSHLogonsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSSHLogonsById`: CSSCMSDataModelModelsSSHLogonsLogonResponse
    fmt.Fprintf(os.Stdout, "Response from `LogonApi.GetSSHLogonsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Keyfactor identifer of the Logon to be Fetched | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSSHLogonsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**CSSCMSDataModelModelsSSHLogonsLogonResponse**](CSSCMSDataModelModelsSSHLogonsLogonResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


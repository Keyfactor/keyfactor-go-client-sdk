# \DeniedAlertApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAlertsDenied**](DeniedAlertApi.md#CreateAlertsDenied) | **POST** /Alerts/Denied | Add a denied alert
[**DeleteAlertsDeniedById**](DeniedAlertApi.md#DeleteAlertsDeniedById) | **DELETE** /Alerts/Denied/{id} | Delete a denied alert
[**GetAlertsDenied**](DeniedAlertApi.md#GetAlertsDenied) | **GET** /Alerts/Denied | Gets all denied alerts according to the provided filter and output parameters
[**GetAlertsDeniedById**](DeniedAlertApi.md#GetAlertsDeniedById) | **GET** /Alerts/Denied/{id} | Get a denied alert
[**UpdateAlertsDenied**](DeniedAlertApi.md#UpdateAlertsDenied) | **PUT** /Alerts/Denied | Edit a denied alert



## CreateAlertsDenied

> AlertsDeniedDeniedAlertDefinitionResponse CreateAlertsDenied(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).AlertsDeniedDeniedAlertCreationRequest(alertsDeniedDeniedAlertCreationRequest).Execute()

Add a denied alert

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
    alertsDeniedDeniedAlertCreationRequest := *openapiclient.NewAlertsDeniedDeniedAlertCreationRequest("DisplayName_example", "Subject_example", "Message_example") // AlertsDeniedDeniedAlertCreationRequest | Information for the new alert (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeniedAlertApi.CreateAlertsDenied(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).AlertsDeniedDeniedAlertCreationRequest(alertsDeniedDeniedAlertCreationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeniedAlertApi.CreateAlertsDenied``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAlertsDenied`: AlertsDeniedDeniedAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `DeniedAlertApi.CreateAlertsDenied`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAlertsDeniedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **alertsDeniedDeniedAlertCreationRequest** | [**AlertsDeniedDeniedAlertCreationRequest**](AlertsDeniedDeniedAlertCreationRequest.md) | Information for the new alert | 

### Return type

[**AlertsDeniedDeniedAlertDefinitionResponse**](AlertsDeniedDeniedAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlertsDeniedById

> DeleteAlertsDeniedById(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Delete a denied alert

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
    id := int32(56) // int32 | Id for the denied alert
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeniedAlertApi.DeleteAlertsDeniedById(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeniedAlertApi.DeleteAlertsDeniedById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id for the denied alert | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlertsDeniedByIdRequest struct via the builder pattern


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


## GetAlertsDenied

> []AlertsDeniedDeniedAlertDefinitionResponse GetAlertsDenied(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Gets all denied alerts according to the provided filter and output parameters

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
    resp, r, err := apiClient.DeniedAlertApi.GetAlertsDenied(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeniedAlertApi.GetAlertsDenied``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlertsDenied`: []AlertsDeniedDeniedAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `DeniedAlertApi.GetAlertsDenied`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertsDeniedRequest struct via the builder pattern


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

[**[]AlertsDeniedDeniedAlertDefinitionResponse**](AlertsDeniedDeniedAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertsDeniedById

> AlertsDeniedDeniedAlertDefinitionResponse GetAlertsDeniedById(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Get a denied alert

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
    id := int32(56) // int32 | Id for the denied alert to get
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeniedAlertApi.GetAlertsDeniedById(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeniedAlertApi.GetAlertsDeniedById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlertsDeniedById`: AlertsDeniedDeniedAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `DeniedAlertApi.GetAlertsDeniedById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id for the denied alert to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertsDeniedByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**AlertsDeniedDeniedAlertDefinitionResponse**](AlertsDeniedDeniedAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAlertsDenied

> AlertsDeniedDeniedAlertDefinitionResponse UpdateAlertsDenied(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).AlertsDeniedDeniedAlertUpdateRequest(alertsDeniedDeniedAlertUpdateRequest).Execute()

Edit a denied alert

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
    alertsDeniedDeniedAlertUpdateRequest := *openapiclient.NewAlertsDeniedDeniedAlertUpdateRequest("DisplayName_example", "Subject_example", "Message_example") // AlertsDeniedDeniedAlertUpdateRequest | Information for the denied alert (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeniedAlertApi.UpdateAlertsDenied(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).AlertsDeniedDeniedAlertUpdateRequest(alertsDeniedDeniedAlertUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeniedAlertApi.UpdateAlertsDenied``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAlertsDenied`: AlertsDeniedDeniedAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `DeniedAlertApi.UpdateAlertsDenied`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAlertsDeniedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **alertsDeniedDeniedAlertUpdateRequest** | [**AlertsDeniedDeniedAlertUpdateRequest**](AlertsDeniedDeniedAlertUpdateRequest.md) | Information for the denied alert | 

### Return type

[**AlertsDeniedDeniedAlertDefinitionResponse**](AlertsDeniedDeniedAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


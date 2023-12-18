# \DeniedAlertApi

All URIs are relative to */Keyfactor/API*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlertsDeniedGet**](DeniedAlertApi.md#AlertsDeniedGet) | **Get** /Alerts/Denied | Gets all denied alerts according to the provided filter and output parameters
[**AlertsDeniedIdDelete**](DeniedAlertApi.md#AlertsDeniedIdDelete) | **Delete** /Alerts/Denied/{id} | Delete a denied alert
[**AlertsDeniedIdGet**](DeniedAlertApi.md#AlertsDeniedIdGet) | **Get** /Alerts/Denied/{id} | Get a denied alert
[**AlertsDeniedPost**](DeniedAlertApi.md#AlertsDeniedPost) | **Post** /Alerts/Denied | Add a denied alert
[**AlertsDeniedPut**](DeniedAlertApi.md#AlertsDeniedPut) | **Put** /Alerts/Denied | Edit a denied alert



## AlertsDeniedGet

> []KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse AlertsDeniedGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    queryString := "queryString_example" // string |  (optional)
    pageReturned := int32(56) // int32 |  (optional)
    returnLimit := int32(56) // int32 |  (optional)
    sortField := "sortField_example" // string |  (optional)
    sortAscending := openapiclient.Keyfactor.Common.QueryableExtensionsSortOrder(0) // KeyfactorCommonQueryableExtensionsSortOrder |  (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeniedAlertApi.AlertsDeniedGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeniedAlertApi.AlertsDeniedGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertsDeniedGet`: []KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `DeniedAlertApi.AlertsDeniedGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertsDeniedGetRequest struct via the builder pattern


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

[**[]KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse**](KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertsDeniedIdDelete

> AlertsDeniedIdDelete(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeniedAlertApi.AlertsDeniedIdDelete(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeniedAlertApi.AlertsDeniedIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAlertsDeniedIdDeleteRequest struct via the builder pattern


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


## AlertsDeniedIdGet

> KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse AlertsDeniedIdGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeniedAlertApi.AlertsDeniedIdGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeniedAlertApi.AlertsDeniedIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertsDeniedIdGet`: KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `DeniedAlertApi.AlertsDeniedIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id for the denied alert to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertsDeniedIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse**](KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertsDeniedPost

> KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse AlertsDeniedPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertCreationRequest(keyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertCreationRequest).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertCreationRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertCreationRequest("DisplayName_example", "Subject_example", "Message_example") // KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertCreationRequest | Information for the new alert (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeniedAlertApi.AlertsDeniedPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertCreationRequest(keyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertCreationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeniedAlertApi.AlertsDeniedPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertsDeniedPost`: KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `DeniedAlertApi.AlertsDeniedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertsDeniedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertCreationRequest** | [**KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertCreationRequest**](KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertCreationRequest.md) | Information for the new alert | 

### Return type

[**KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse**](KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patchjson, application/json, text/json, application/*json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertsDeniedPut

> KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse AlertsDeniedPut(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertUpdateRequest(keyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertUpdateRequest).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertUpdateRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertUpdateRequest("DisplayName_example", "Subject_example", "Message_example") // KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertUpdateRequest | Information for the denied alert (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeniedAlertApi.AlertsDeniedPut(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertUpdateRequest(keyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeniedAlertApi.AlertsDeniedPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertsDeniedPut`: KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `DeniedAlertApi.AlertsDeniedPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertsDeniedPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertUpdateRequest** | [**KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertUpdateRequest**](KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertUpdateRequest.md) | Information for the denied alert | 

### Return type

[**KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse**](KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patchjson, application/json, text/json, application/*json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


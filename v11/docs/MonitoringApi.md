# \MonitoringApi

All URIs are relative to */Keyfactor/API*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MonitoringResolveOCSPPost**](MonitoringApi.md#MonitoringResolveOCSPPost) | **Post** /Monitoring/ResolveOCSP | Resolve the Certificate authority given
[**MonitoringRevocationGet**](MonitoringApi.md#MonitoringRevocationGet) | **Get** /Monitoring/Revocation | Gets all revocation monitoring endpoints according to the provided filter and output parameters
[**MonitoringRevocationIdDelete**](MonitoringApi.md#MonitoringRevocationIdDelete) | **Delete** /Monitoring/Revocation/{id} | Delete a revocation monitoring endpoint
[**MonitoringRevocationIdGet**](MonitoringApi.md#MonitoringRevocationIdGet) | **Get** /Monitoring/Revocation/{id} | Get a revocation monitoring endpoint
[**MonitoringRevocationPost**](MonitoringApi.md#MonitoringRevocationPost) | **Post** /Monitoring/Revocation | Add a revocation monitoring endpoint
[**MonitoringRevocationPut**](MonitoringApi.md#MonitoringRevocationPut) | **Put** /Monitoring/Revocation | Edit a revocation monitoring endpoint
[**MonitoringRevocationTestAllPost**](MonitoringApi.md#MonitoringRevocationTestAllPost) | **Post** /Monitoring/Revocation/TestAll | Test All Alerts
[**MonitoringRevocationTestPost**](MonitoringApi.md#MonitoringRevocationTestPost) | **Post** /Monitoring/Revocation/Test | Test Alert



## MonitoringResolveOCSPPost

> KeyfactorWebKeyfactorApiModelsMonitoringOCSPParametersResponse MonitoringResolveOCSPPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsMonitoringOCSPParametersRequest(keyfactorWebKeyfactorApiModelsMonitoringOCSPParametersRequest).Execute()

Resolve the Certificate authority given

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
    keyfactorWebKeyfactorApiModelsMonitoringOCSPParametersRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsMonitoringOCSPParametersRequest() // KeyfactorWebKeyfactorApiModelsMonitoringOCSPParametersRequest | Information for the new endpoint (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringApi.MonitoringResolveOCSPPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsMonitoringOCSPParametersRequest(keyfactorWebKeyfactorApiModelsMonitoringOCSPParametersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringApi.MonitoringResolveOCSPPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitoringResolveOCSPPost`: KeyfactorWebKeyfactorApiModelsMonitoringOCSPParametersResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitoringApi.MonitoringResolveOCSPPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMonitoringResolveOCSPPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsMonitoringOCSPParametersRequest** | [**KeyfactorWebKeyfactorApiModelsMonitoringOCSPParametersRequest**](KeyfactorWebKeyfactorApiModelsMonitoringOCSPParametersRequest.md) | Information for the new endpoint | 

### Return type

[**KeyfactorWebKeyfactorApiModelsMonitoringOCSPParametersResponse**](KeyfactorWebKeyfactorApiModelsMonitoringOCSPParametersResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitoringRevocationGet

> []KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse MonitoringRevocationGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Gets all revocation monitoring endpoints according to the provided filter and output parameters

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
    resp, r, err := apiClient.MonitoringApi.MonitoringRevocationGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringApi.MonitoringRevocationGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitoringRevocationGet`: []KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitoringApi.MonitoringRevocationGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMonitoringRevocationGetRequest struct via the builder pattern


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

[**[]KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse**](KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitoringRevocationIdDelete

> MonitoringRevocationIdDelete(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Delete a revocation monitoring endpoint

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
    id := int32(56) // int32 | Id for the revocation monitoring endpoint
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringApi.MonitoringRevocationIdDelete(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringApi.MonitoringRevocationIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id for the revocation monitoring endpoint | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitoringRevocationIdDeleteRequest struct via the builder pattern


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


## MonitoringRevocationIdGet

> KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse MonitoringRevocationIdGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Get a revocation monitoring endpoint

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
    id := int32(56) // int32 | Id for the endpoint to get
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringApi.MonitoringRevocationIdGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringApi.MonitoringRevocationIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitoringRevocationIdGet`: KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitoringApi.MonitoringRevocationIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id for the endpoint to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitoringRevocationIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse**](KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitoringRevocationPost

> KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse MonitoringRevocationPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest(keyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest).Execute()

Add a revocation monitoring endpoint

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
    keyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest("Name_example", "EndpointType_example", "Location_example", *openapiclient.NewKeyfactorWebKeyfactorApiModelsMonitoringDashboardRequest(false)) // KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest | Information for the new endpoint (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringApi.MonitoringRevocationPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest(keyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringApi.MonitoringRevocationPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitoringRevocationPost`: KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitoringApi.MonitoringRevocationPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMonitoringRevocationPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest** | [**KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest**](KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest.md) | Information for the new endpoint | 

### Return type

[**KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse**](KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitoringRevocationPut

> KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse MonitoringRevocationPut(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest(keyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest).Execute()

Edit a revocation monitoring endpoint

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
    keyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest("Name_example", "EndpointType_example", "Location_example", *openapiclient.NewKeyfactorWebKeyfactorApiModelsMonitoringDashboardRequest(false)) // KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest | Information for the endpoint (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringApi.MonitoringRevocationPut(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest(keyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringApi.MonitoringRevocationPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitoringRevocationPut`: KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitoringApi.MonitoringRevocationPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMonitoringRevocationPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest** | [**KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest**](KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest.md) | Information for the endpoint | 

### Return type

[**KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse**](KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitoringRevocationTestAllPost

> KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestResponse MonitoringRevocationTestAllPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestAllRequest(keyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestAllRequest).Execute()

Test All Alerts

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
    keyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestAllRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestAllRequest() // KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestAllRequest | Information about the revocation monitoring alert test (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringApi.MonitoringRevocationTestAllPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestAllRequest(keyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestAllRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringApi.MonitoringRevocationTestAllPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitoringRevocationTestAllPost`: KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitoringApi.MonitoringRevocationTestAllPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMonitoringRevocationTestAllPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestAllRequest** | [**KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestAllRequest**](KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestAllRequest.md) | Information about the revocation monitoring alert test | 

### Return type

[**KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestResponse**](KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitoringRevocationTestPost

> KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestResponse MonitoringRevocationTestPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest(keyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest).Execute()

Test Alert

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
    keyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest() // KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest | Information about the revocation monitoring alert test (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringApi.MonitoringRevocationTestPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest(keyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringApi.MonitoringRevocationTestPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitoringRevocationTestPost`: KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitoringApi.MonitoringRevocationTestPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMonitoringRevocationTestPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest** | [**KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest**](KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest.md) | Information about the revocation monitoring alert test | 

### Return type

[**KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestResponse**](KeyfactorWebKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


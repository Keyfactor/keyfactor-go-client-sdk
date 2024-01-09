# \PendingAlertApi

All URIs are relative to */Keyfactor/API*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlertsPendingGet**](PendingAlertApi.md#AlertsPendingGet) | **Get** /Alerts/Pending | Gets all pending alerts according to the provided filter and output parameters
[**AlertsPendingIdDelete**](PendingAlertApi.md#AlertsPendingIdDelete) | **Delete** /Alerts/Pending/{id} | Delete a pending alert
[**AlertsPendingIdGet**](PendingAlertApi.md#AlertsPendingIdGet) | **Get** /Alerts/Pending/{id} | Get a pending alert
[**AlertsPendingPost**](PendingAlertApi.md#AlertsPendingPost) | **Post** /Alerts/Pending | Add a pending alert
[**AlertsPendingPut**](PendingAlertApi.md#AlertsPendingPut) | **Put** /Alerts/Pending | Edit a pending alert
[**AlertsPendingScheduleGet**](PendingAlertApi.md#AlertsPendingScheduleGet) | **Get** /Alerts/Pending/Schedule | Get the schedule for pending alerts
[**AlertsPendingSchedulePut**](PendingAlertApi.md#AlertsPendingSchedulePut) | **Put** /Alerts/Pending/Schedule | Edit schedule
[**AlertsPendingTestAllPost**](PendingAlertApi.md#AlertsPendingTestAllPost) | **Post** /Alerts/Pending/TestAll | Test all pending alerts. Will send alert emails if SendAlerts is true
[**AlertsPendingTestPost**](PendingAlertApi.md#AlertsPendingTestPost) | **Post** /Alerts/Pending/Test | Test pending alert. Will send alert emails if SendAlerts is true



## AlertsPendingGet

> []KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertDefinitionResponse AlertsPendingGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Gets all pending alerts according to the provided filter and output parameters

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
    resp, r, err := apiClient.PendingAlertApi.AlertsPendingGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PendingAlertApi.AlertsPendingGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertsPendingGet`: []KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `PendingAlertApi.AlertsPendingGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertsPendingGetRequest struct via the builder pattern


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

[**[]KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertDefinitionResponse**](KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertsPendingIdDelete

> AlertsPendingIdDelete(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Delete a pending alert

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
    id := int32(56) // int32 | Id for the pending alert
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PendingAlertApi.AlertsPendingIdDelete(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PendingAlertApi.AlertsPendingIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id for the pending alert | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertsPendingIdDeleteRequest struct via the builder pattern


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


## AlertsPendingIdGet

> KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertDefinitionResponse AlertsPendingIdGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Get a pending alert

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
    id := int32(56) // int32 | Id for the pending alert to get
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PendingAlertApi.AlertsPendingIdGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PendingAlertApi.AlertsPendingIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertsPendingIdGet`: KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `PendingAlertApi.AlertsPendingIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id for the pending alert to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertsPendingIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertDefinitionResponse**](KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertsPendingPost

> KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertDefinitionResponse AlertsPendingPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest(keyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest).Execute()

Add a pending alert

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
    keyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest("DisplayName_example", "Subject_example", "Message_example") // KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest | Information for the new alert (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PendingAlertApi.AlertsPendingPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest(keyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PendingAlertApi.AlertsPendingPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertsPendingPost`: KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `PendingAlertApi.AlertsPendingPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertsPendingPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest** | [**KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest**](KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertCreationRequest.md) | Information for the new alert | 

### Return type

[**KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertDefinitionResponse**](KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertsPendingPut

> KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertDefinitionResponse AlertsPendingPut(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest(keyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest).Execute()

Edit a pending alert

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
    keyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest("DisplayName_example", "Subject_example", "Message_example") // KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest | Information for the pending alert (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PendingAlertApi.AlertsPendingPut(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest(keyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PendingAlertApi.AlertsPendingPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertsPendingPut`: KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `PendingAlertApi.AlertsPendingPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertsPendingPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest** | [**KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest**](KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest.md) | Information for the pending alert | 

### Return type

[**KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertDefinitionResponse**](KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertsPendingScheduleGet

> KeyfactorWebKeyfactorApiModelsAlertsAlertScheduleAlertScheduleResponse AlertsPendingScheduleGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Get the schedule for pending alerts

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

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PendingAlertApi.AlertsPendingScheduleGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PendingAlertApi.AlertsPendingScheduleGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertsPendingScheduleGet`: KeyfactorWebKeyfactorApiModelsAlertsAlertScheduleAlertScheduleResponse
    fmt.Fprintf(os.Stdout, "Response from `PendingAlertApi.AlertsPendingScheduleGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertsPendingScheduleGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**KeyfactorWebKeyfactorApiModelsAlertsAlertScheduleAlertScheduleResponse**](KeyfactorWebKeyfactorApiModelsAlertsAlertScheduleAlertScheduleResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertsPendingSchedulePut

> KeyfactorWebKeyfactorApiModelsAlertsAlertScheduleAlertScheduleResponse AlertsPendingSchedulePut(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsAlertsAlertScheduleAlertScheduleRequest(keyfactorWebKeyfactorApiModelsAlertsAlertScheduleAlertScheduleRequest).Execute()

Edit schedule

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
    keyfactorWebKeyfactorApiModelsAlertsAlertScheduleAlertScheduleRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsAlertsAlertScheduleAlertScheduleRequest() // KeyfactorWebKeyfactorApiModelsAlertsAlertScheduleAlertScheduleRequest |  (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PendingAlertApi.AlertsPendingSchedulePut(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsAlertsAlertScheduleAlertScheduleRequest(keyfactorWebKeyfactorApiModelsAlertsAlertScheduleAlertScheduleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PendingAlertApi.AlertsPendingSchedulePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertsPendingSchedulePut`: KeyfactorWebKeyfactorApiModelsAlertsAlertScheduleAlertScheduleResponse
    fmt.Fprintf(os.Stdout, "Response from `PendingAlertApi.AlertsPendingSchedulePut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertsPendingSchedulePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsAlertsAlertScheduleAlertScheduleRequest** | [**KeyfactorWebKeyfactorApiModelsAlertsAlertScheduleAlertScheduleRequest**](KeyfactorWebKeyfactorApiModelsAlertsAlertScheduleAlertScheduleRequest.md) |  | 

### Return type

[**KeyfactorWebKeyfactorApiModelsAlertsAlertScheduleAlertScheduleResponse**](KeyfactorWebKeyfactorApiModelsAlertsAlertScheduleAlertScheduleResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertsPendingTestAllPost

> KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertTestResponse AlertsPendingTestAllPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertTestAllRequest(keyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertTestAllRequest).Execute()

Test all pending alerts. Will send alert emails if SendAlerts is true

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
    keyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertTestAllRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertTestAllRequest() // KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertTestAllRequest | Information for the pending alert (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PendingAlertApi.AlertsPendingTestAllPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertTestAllRequest(keyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertTestAllRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PendingAlertApi.AlertsPendingTestAllPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertsPendingTestAllPost`: KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertTestResponse
    fmt.Fprintf(os.Stdout, "Response from `PendingAlertApi.AlertsPendingTestAllPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertsPendingTestAllPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertTestAllRequest** | [**KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertTestAllRequest**](KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertTestAllRequest.md) | Information for the pending alert | 

### Return type

[**KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertTestResponse**](KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertTestResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertsPendingTestPost

> KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertTestResponse AlertsPendingTestPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertTestRequest(keyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertTestRequest).Execute()

Test pending alert. Will send alert emails if SendAlerts is true

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
    keyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertTestRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertTestRequest() // KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertTestRequest | Information for the pending alert (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PendingAlertApi.AlertsPendingTestPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertTestRequest(keyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertTestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PendingAlertApi.AlertsPendingTestPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertsPendingTestPost`: KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertTestResponse
    fmt.Fprintf(os.Stdout, "Response from `PendingAlertApi.AlertsPendingTestPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertsPendingTestPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertTestRequest** | [**KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertTestRequest**](KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertTestRequest.md) | Information for the pending alert | 

### Return type

[**KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertTestResponse**](KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertTestResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


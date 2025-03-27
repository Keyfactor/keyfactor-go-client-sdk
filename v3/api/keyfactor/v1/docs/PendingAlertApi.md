# \PendingAlertApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAlertsPending**](PendingAlertApi.md#CreateAlertsPending) | **POST** /Alerts/Pending | Add a pending alert
[**CreateAlertsPendingTest**](PendingAlertApi.md#CreateAlertsPendingTest) | **POST** /Alerts/Pending/Test | Test pending alert. Will send alert emails if SendAlerts is true
[**CreateAlertsPendingTestAll**](PendingAlertApi.md#CreateAlertsPendingTestAll) | **POST** /Alerts/Pending/TestAll | Test all pending alerts. Will send alert emails if SendAlerts is true
[**DeleteAlertsPendingById**](PendingAlertApi.md#DeleteAlertsPendingById) | **DELETE** /Alerts/Pending/{id} | Delete a pending alert
[**GetAlertsPending**](PendingAlertApi.md#GetAlertsPending) | **GET** /Alerts/Pending | Gets all pending alerts according to the provided filter and output parameters
[**GetAlertsPendingById**](PendingAlertApi.md#GetAlertsPendingById) | **GET** /Alerts/Pending/{id} | Get a pending alert
[**GetAlertsPendingSchedule**](PendingAlertApi.md#GetAlertsPendingSchedule) | **GET** /Alerts/Pending/Schedule | Get the schedule for pending alerts
[**UpdateAlertsPending**](PendingAlertApi.md#UpdateAlertsPending) | **PUT** /Alerts/Pending | Edit a pending alert
[**UpdateAlertsPendingSchedule**](PendingAlertApi.md#UpdateAlertsPendingSchedule) | **PUT** /Alerts/Pending/Schedule | Edit schedule



## CreateAlertsPending

> AlertsPendingPendingAlertDefinitionResponse BuildCreateAlertsPendingRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).AlertsPendingPendingAlertCreationRequest(alertsPendingPendingAlertCreationRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    alertsPendingPendingAlertCreationRequest := *openapiclient.NewAlertsPendingPendingAlertCreationRequest("DisplayName_example", "Subject_example", "Message_example") // AlertsPendingPendingAlertCreationRequest | Information for the new alert (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PendingAlertApi.BuildCreateAlertsPendingRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).AlertsPendingPendingAlertCreationRequest(alertsPendingPendingAlertCreationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PendingAlertApi.CreateAlertsPending``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAlertsPending`: AlertsPendingPendingAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `PendingAlertApi.CreateAlertsPending`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAlertsPendingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **alertsPendingPendingAlertCreationRequest** | [**AlertsPendingPendingAlertCreationRequest**](AlertsPendingPendingAlertCreationRequest.md) | Information for the new alert | 

### Return type

[**AlertsPendingPendingAlertDefinitionResponse**](AlertsPendingPendingAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAlertsPendingTest

> AlertsPendingPendingAlertTestResponse BuildCreateAlertsPendingTestRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).AlertsPendingPendingAlertTestRequest(alertsPendingPendingAlertTestRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    alertsPendingPendingAlertTestRequest := *openapiclient.NewAlertsPendingPendingAlertTestRequest() // AlertsPendingPendingAlertTestRequest | Information for the pending alert (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PendingAlertApi.BuildCreateAlertsPendingTestRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).AlertsPendingPendingAlertTestRequest(alertsPendingPendingAlertTestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PendingAlertApi.CreateAlertsPendingTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAlertsPendingTest`: AlertsPendingPendingAlertTestResponse
    fmt.Fprintf(os.Stdout, "Response from `PendingAlertApi.CreateAlertsPendingTest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAlertsPendingTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **alertsPendingPendingAlertTestRequest** | [**AlertsPendingPendingAlertTestRequest**](AlertsPendingPendingAlertTestRequest.md) | Information for the pending alert | 

### Return type

[**AlertsPendingPendingAlertTestResponse**](AlertsPendingPendingAlertTestResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAlertsPendingTestAll

> AlertsPendingPendingAlertTestResponse BuildCreateAlertsPendingTestAllRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).AlertsPendingPendingAlertTestAllRequest(alertsPendingPendingAlertTestAllRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    alertsPendingPendingAlertTestAllRequest := *openapiclient.NewAlertsPendingPendingAlertTestAllRequest() // AlertsPendingPendingAlertTestAllRequest | Information for the pending alert (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PendingAlertApi.BuildCreateAlertsPendingTestAllRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).AlertsPendingPendingAlertTestAllRequest(alertsPendingPendingAlertTestAllRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PendingAlertApi.CreateAlertsPendingTestAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAlertsPendingTestAll`: AlertsPendingPendingAlertTestResponse
    fmt.Fprintf(os.Stdout, "Response from `PendingAlertApi.CreateAlertsPendingTestAll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAlertsPendingTestAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **alertsPendingPendingAlertTestAllRequest** | [**AlertsPendingPendingAlertTestAllRequest**](AlertsPendingPendingAlertTestAllRequest.md) | Information for the pending alert | 

### Return type

[**AlertsPendingPendingAlertTestResponse**](AlertsPendingPendingAlertTestResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlertsPendingById

> BuildDeleteAlertsPendingByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PendingAlertApi.BuildDeleteAlertsPendingByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PendingAlertApi.DeleteAlertsPendingById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteAlertsPendingByIdRequest struct via the builder pattern


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


## GetAlertsPending

> []AlertsPendingPendingAlertDefinitionResponse BuildGetAlertsPendingRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    queryString := "queryString_example" // string |  (optional)
    pageReturned := int32(56) // int32 |  (optional)
    returnLimit := int32(56) // int32 |  (optional)
    sortField := "sortField_example" // string |  (optional)
    sortAscending := openapiclient.Keyfactor.Common.QueryableExtensions.SortOrder(0) // KeyfactorCommonQueryableExtensionsSortOrder |  (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PendingAlertApi.BuildGetAlertsPendingRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PendingAlertApi.GetAlertsPending``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlertsPending`: []AlertsPendingPendingAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `PendingAlertApi.GetAlertsPending`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertsPendingRequest struct via the builder pattern


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

[**[]AlertsPendingPendingAlertDefinitionResponse**](AlertsPendingPendingAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertsPendingById

> AlertsPendingPendingAlertDefinitionResponse BuildGetAlertsPendingByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PendingAlertApi.BuildGetAlertsPendingByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PendingAlertApi.GetAlertsPendingById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlertsPendingById`: AlertsPendingPendingAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `PendingAlertApi.GetAlertsPendingById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id for the pending alert to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertsPendingByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**AlertsPendingPendingAlertDefinitionResponse**](AlertsPendingPendingAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertsPendingSchedule

> AlertsAlertScheduleAlertScheduleResponse BuildGetAlertsPendingScheduleRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PendingAlertApi.BuildGetAlertsPendingScheduleRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PendingAlertApi.GetAlertsPendingSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlertsPendingSchedule`: AlertsAlertScheduleAlertScheduleResponse
    fmt.Fprintf(os.Stdout, "Response from `PendingAlertApi.GetAlertsPendingSchedule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertsPendingScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**AlertsAlertScheduleAlertScheduleResponse**](AlertsAlertScheduleAlertScheduleResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAlertsPending

> AlertsPendingPendingAlertDefinitionResponse BuildUpdateAlertsPendingRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).AlertsPendingPendingAlertUpdateRequest(alertsPendingPendingAlertUpdateRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    alertsPendingPendingAlertUpdateRequest := *openapiclient.NewAlertsPendingPendingAlertUpdateRequest("DisplayName_example", "Subject_example", "Message_example") // AlertsPendingPendingAlertUpdateRequest | Information for the pending alert (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PendingAlertApi.BuildUpdateAlertsPendingRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).AlertsPendingPendingAlertUpdateRequest(alertsPendingPendingAlertUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PendingAlertApi.UpdateAlertsPending``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAlertsPending`: AlertsPendingPendingAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `PendingAlertApi.UpdateAlertsPending`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAlertsPendingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **alertsPendingPendingAlertUpdateRequest** | [**AlertsPendingPendingAlertUpdateRequest**](AlertsPendingPendingAlertUpdateRequest.md) | Information for the pending alert | 

### Return type

[**AlertsPendingPendingAlertDefinitionResponse**](AlertsPendingPendingAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAlertsPendingSchedule

> AlertsAlertScheduleAlertScheduleResponse BuildUpdateAlertsPendingScheduleRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).AlertsAlertScheduleAlertScheduleRequest(alertsAlertScheduleAlertScheduleRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    alertsAlertScheduleAlertScheduleRequest := *openapiclient.NewAlertsAlertScheduleAlertScheduleRequest() // AlertsAlertScheduleAlertScheduleRequest |  (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PendingAlertApi.BuildUpdateAlertsPendingScheduleRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).AlertsAlertScheduleAlertScheduleRequest(alertsAlertScheduleAlertScheduleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PendingAlertApi.UpdateAlertsPendingSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAlertsPendingSchedule`: AlertsAlertScheduleAlertScheduleResponse
    fmt.Fprintf(os.Stdout, "Response from `PendingAlertApi.UpdateAlertsPendingSchedule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAlertsPendingScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **alertsAlertScheduleAlertScheduleRequest** | [**AlertsAlertScheduleAlertScheduleRequest**](AlertsAlertScheduleAlertScheduleRequest.md) |  | 

### Return type

[**AlertsAlertScheduleAlertScheduleResponse**](AlertsAlertScheduleAlertScheduleResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


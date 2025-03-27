# \KeyRotationAlertApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAlertsKeyRotation**](KeyRotationAlertApi.md#CreateAlertsKeyRotation) | **POST** /Alerts/KeyRotation | Add a key rotation alert
[**CreateAlertsKeyRotationTest**](KeyRotationAlertApi.md#CreateAlertsKeyRotationTest) | **POST** /Alerts/KeyRotation/Test | Test An Alert
[**CreateAlertsKeyRotationTestAll**](KeyRotationAlertApi.md#CreateAlertsKeyRotationTestAll) | **POST** /Alerts/KeyRotation/TestAll | Test All Alerts
[**DeleteAlertsKeyRotationById**](KeyRotationAlertApi.md#DeleteAlertsKeyRotationById) | **DELETE** /Alerts/KeyRotation/{id} | Delete a key rotation alert
[**GetAlertsKeyRotation**](KeyRotationAlertApi.md#GetAlertsKeyRotation) | **GET** /Alerts/KeyRotation | Gets all key rotation alerts according to the provided filter and output parameters
[**GetAlertsKeyRotationById**](KeyRotationAlertApi.md#GetAlertsKeyRotationById) | **GET** /Alerts/KeyRotation/{id} | Get a key rotation alert
[**GetAlertsKeyRotationSchedule**](KeyRotationAlertApi.md#GetAlertsKeyRotationSchedule) | **GET** /Alerts/KeyRotation/Schedule | Get the schedule for key rotation alerts
[**UpdateAlertsKeyRotation**](KeyRotationAlertApi.md#UpdateAlertsKeyRotation) | **PUT** /Alerts/KeyRotation | Edit a key rotation alert
[**UpdateAlertsKeyRotationSchedule**](KeyRotationAlertApi.md#UpdateAlertsKeyRotationSchedule) | **PUT** /Alerts/KeyRotation/Schedule | Edit schedule



## CreateAlertsKeyRotation

> AlertsKeyRotationKeyRotationAlertDefinitionResponse BuildCreateAlertsKeyRotationRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).AlertsKeyRotationKeyRotationAlertCreationRequest(alertsKeyRotationKeyRotationAlertCreationRequest).Execute()

Add a key rotation alert

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
    alertsKeyRotationKeyRotationAlertCreationRequest := *openapiclient.NewAlertsKeyRotationKeyRotationAlertCreationRequest("DisplayName_example", int32(123)) // AlertsKeyRotationKeyRotationAlertCreationRequest | Information for the new alert (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyRotationAlertApi.BuildCreateAlertsKeyRotationRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).AlertsKeyRotationKeyRotationAlertCreationRequest(alertsKeyRotationKeyRotationAlertCreationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyRotationAlertApi.CreateAlertsKeyRotation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAlertsKeyRotation`: AlertsKeyRotationKeyRotationAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `KeyRotationAlertApi.CreateAlertsKeyRotation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAlertsKeyRotationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **alertsKeyRotationKeyRotationAlertCreationRequest** | [**AlertsKeyRotationKeyRotationAlertCreationRequest**](AlertsKeyRotationKeyRotationAlertCreationRequest.md) | Information for the new alert | 

### Return type

[**AlertsKeyRotationKeyRotationAlertDefinitionResponse**](AlertsKeyRotationKeyRotationAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAlertsKeyRotationTest

> AlertsKeyRotationKeyRotationAlertTestResponse BuildCreateAlertsKeyRotationTestRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).AlertsKeyRotationKeyRotationAlertTestRequest(alertsKeyRotationKeyRotationAlertTestRequest).Execute()

Test An Alert

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
    alertsKeyRotationKeyRotationAlertTestRequest := *openapiclient.NewAlertsKeyRotationKeyRotationAlertTestRequest() // AlertsKeyRotationKeyRotationAlertTestRequest | Parameters used to test the alert (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyRotationAlertApi.BuildCreateAlertsKeyRotationTestRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).AlertsKeyRotationKeyRotationAlertTestRequest(alertsKeyRotationKeyRotationAlertTestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyRotationAlertApi.CreateAlertsKeyRotationTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAlertsKeyRotationTest`: AlertsKeyRotationKeyRotationAlertTestResponse
    fmt.Fprintf(os.Stdout, "Response from `KeyRotationAlertApi.CreateAlertsKeyRotationTest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAlertsKeyRotationTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **alertsKeyRotationKeyRotationAlertTestRequest** | [**AlertsKeyRotationKeyRotationAlertTestRequest**](AlertsKeyRotationKeyRotationAlertTestRequest.md) | Parameters used to test the alert | 

### Return type

[**AlertsKeyRotationKeyRotationAlertTestResponse**](AlertsKeyRotationKeyRotationAlertTestResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAlertsKeyRotationTestAll

> AlertsKeyRotationKeyRotationAlertTestResponse BuildCreateAlertsKeyRotationTestAllRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).AlertsKeyRotationKeyRotationAlertTestAllRequest(alertsKeyRotationKeyRotationAlertTestAllRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    alertsKeyRotationKeyRotationAlertTestAllRequest := *openapiclient.NewAlertsKeyRotationKeyRotationAlertTestAllRequest() // AlertsKeyRotationKeyRotationAlertTestAllRequest | Information about the key rotation alert test (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyRotationAlertApi.BuildCreateAlertsKeyRotationTestAllRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).AlertsKeyRotationKeyRotationAlertTestAllRequest(alertsKeyRotationKeyRotationAlertTestAllRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyRotationAlertApi.CreateAlertsKeyRotationTestAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAlertsKeyRotationTestAll`: AlertsKeyRotationKeyRotationAlertTestResponse
    fmt.Fprintf(os.Stdout, "Response from `KeyRotationAlertApi.CreateAlertsKeyRotationTestAll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAlertsKeyRotationTestAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **alertsKeyRotationKeyRotationAlertTestAllRequest** | [**AlertsKeyRotationKeyRotationAlertTestAllRequest**](AlertsKeyRotationKeyRotationAlertTestAllRequest.md) | Information about the key rotation alert test | 

### Return type

[**AlertsKeyRotationKeyRotationAlertTestResponse**](AlertsKeyRotationKeyRotationAlertTestResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlertsKeyRotationById

> BuildDeleteAlertsKeyRotationByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Delete a key rotation alert

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
    id := int32(56) // int32 | Id for the key rotation alert
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyRotationAlertApi.BuildDeleteAlertsKeyRotationByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyRotationAlertApi.DeleteAlertsKeyRotationById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id for the key rotation alert | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlertsKeyRotationByIdRequest struct via the builder pattern


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


## GetAlertsKeyRotation

> []AlertsKeyRotationKeyRotationAlertDefinitionResponse BuildGetAlertsKeyRotationRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Gets all key rotation alerts according to the provided filter and output parameters

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
    resp, r, err := apiClient.KeyRotationAlertApi.BuildGetAlertsKeyRotationRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyRotationAlertApi.GetAlertsKeyRotation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlertsKeyRotation`: []AlertsKeyRotationKeyRotationAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `KeyRotationAlertApi.GetAlertsKeyRotation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertsKeyRotationRequest struct via the builder pattern


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

[**[]AlertsKeyRotationKeyRotationAlertDefinitionResponse**](AlertsKeyRotationKeyRotationAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertsKeyRotationById

> AlertsKeyRotationKeyRotationAlertDefinitionResponse BuildGetAlertsKeyRotationByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Get a key rotation alert

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
    id := int32(56) // int32 | Id for the key rotation alert to get
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyRotationAlertApi.BuildGetAlertsKeyRotationByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyRotationAlertApi.GetAlertsKeyRotationById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlertsKeyRotationById`: AlertsKeyRotationKeyRotationAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `KeyRotationAlertApi.GetAlertsKeyRotationById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id for the key rotation alert to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertsKeyRotationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**AlertsKeyRotationKeyRotationAlertDefinitionResponse**](AlertsKeyRotationKeyRotationAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertsKeyRotationSchedule

> AlertsAlertScheduleAlertScheduleResponse BuildGetAlertsKeyRotationScheduleRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Get the schedule for key rotation alerts

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
    resp, r, err := apiClient.KeyRotationAlertApi.BuildGetAlertsKeyRotationScheduleRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyRotationAlertApi.GetAlertsKeyRotationSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlertsKeyRotationSchedule`: AlertsAlertScheduleAlertScheduleResponse
    fmt.Fprintf(os.Stdout, "Response from `KeyRotationAlertApi.GetAlertsKeyRotationSchedule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertsKeyRotationScheduleRequest struct via the builder pattern


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


## UpdateAlertsKeyRotation

> AlertsKeyRotationKeyRotationAlertDefinitionResponse BuildUpdateAlertsKeyRotationRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).AlertsKeyRotationKeyRotationAlertUpdateRequest(alertsKeyRotationKeyRotationAlertUpdateRequest).Execute()

Edit a key rotation alert

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
    alertsKeyRotationKeyRotationAlertUpdateRequest := *openapiclient.NewAlertsKeyRotationKeyRotationAlertUpdateRequest("DisplayName_example", int32(123)) // AlertsKeyRotationKeyRotationAlertUpdateRequest | Information for the key rotation alert (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyRotationAlertApi.BuildUpdateAlertsKeyRotationRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).AlertsKeyRotationKeyRotationAlertUpdateRequest(alertsKeyRotationKeyRotationAlertUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyRotationAlertApi.UpdateAlertsKeyRotation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAlertsKeyRotation`: AlertsKeyRotationKeyRotationAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `KeyRotationAlertApi.UpdateAlertsKeyRotation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAlertsKeyRotationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **alertsKeyRotationKeyRotationAlertUpdateRequest** | [**AlertsKeyRotationKeyRotationAlertUpdateRequest**](AlertsKeyRotationKeyRotationAlertUpdateRequest.md) | Information for the key rotation alert | 

### Return type

[**AlertsKeyRotationKeyRotationAlertDefinitionResponse**](AlertsKeyRotationKeyRotationAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAlertsKeyRotationSchedule

> AlertsAlertScheduleAlertScheduleResponse BuildUpdateAlertsKeyRotationScheduleRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).AlertsAlertScheduleAlertScheduleRequest(alertsAlertScheduleAlertScheduleRequest).Execute()

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
    resp, r, err := apiClient.KeyRotationAlertApi.BuildUpdateAlertsKeyRotationScheduleRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).AlertsAlertScheduleAlertScheduleRequest(alertsAlertScheduleAlertScheduleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyRotationAlertApi.UpdateAlertsKeyRotationSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAlertsKeyRotationSchedule`: AlertsAlertScheduleAlertScheduleResponse
    fmt.Fprintf(os.Stdout, "Response from `KeyRotationAlertApi.UpdateAlertsKeyRotationSchedule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAlertsKeyRotationScheduleRequest struct via the builder pattern


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


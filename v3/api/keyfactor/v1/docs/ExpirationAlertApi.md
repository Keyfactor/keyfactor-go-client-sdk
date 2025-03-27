# \ExpirationAlertApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAlertsExpiration**](ExpirationAlertApi.md#CreateAlertsExpiration) | **POST** /Alerts/Expiration | Add an expiration alert
[**CreateAlertsExpirationTest**](ExpirationAlertApi.md#CreateAlertsExpirationTest) | **POST** /Alerts/Expiration/Test | Test an Expiration Alert
[**CreateAlertsExpirationTestAll**](ExpirationAlertApi.md#CreateAlertsExpirationTestAll) | **POST** /Alerts/Expiration/TestAll | Test All Expiration Alerts
[**DeleteAlertsExpirationById**](ExpirationAlertApi.md#DeleteAlertsExpirationById) | **DELETE** /Alerts/Expiration/{id} | Delete an expiration alert
[**GetAlertsExpiration**](ExpirationAlertApi.md#GetAlertsExpiration) | **GET** /Alerts/Expiration | Gets all expiration alerts according to the provided filter and output parameters
[**GetAlertsExpirationById**](ExpirationAlertApi.md#GetAlertsExpirationById) | **GET** /Alerts/Expiration/{id} | Get an expiration alert
[**GetAlertsExpirationSchedule**](ExpirationAlertApi.md#GetAlertsExpirationSchedule) | **GET** /Alerts/Expiration/Schedule | Get the schedule for expiration alerts
[**UpdateAlertsExpiration**](ExpirationAlertApi.md#UpdateAlertsExpiration) | **PUT** /Alerts/Expiration | Edit an expiration alert
[**UpdateAlertsExpirationSchedule**](ExpirationAlertApi.md#UpdateAlertsExpirationSchedule) | **PUT** /Alerts/Expiration/Schedule | Edit schedule



## CreateAlertsExpiration

> AlertsExpirationExpirationAlertDefinitionResponse BuildCreateAlertsExpirationRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).AlertsExpirationExpirationAlertCreationRequest(alertsExpirationExpirationAlertCreationRequest).Execute()

Add an expiration alert

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
    alertsExpirationExpirationAlertCreationRequest := *openapiclient.NewAlertsExpirationExpirationAlertCreationRequest("DisplayName_example", int32(123)) // AlertsExpirationExpirationAlertCreationRequest | Information for the new alert (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExpirationAlertApi.BuildCreateAlertsExpirationRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).AlertsExpirationExpirationAlertCreationRequest(alertsExpirationExpirationAlertCreationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpirationAlertApi.CreateAlertsExpiration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAlertsExpiration`: AlertsExpirationExpirationAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `ExpirationAlertApi.CreateAlertsExpiration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAlertsExpirationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **alertsExpirationExpirationAlertCreationRequest** | [**AlertsExpirationExpirationAlertCreationRequest**](AlertsExpirationExpirationAlertCreationRequest.md) | Information for the new alert | 

### Return type

[**AlertsExpirationExpirationAlertDefinitionResponse**](AlertsExpirationExpirationAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAlertsExpirationTest

> AlertsExpirationExpirationAlertTestResponse BuildCreateAlertsExpirationTestRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).AlertsExpirationExpirationAlertTestRequest(alertsExpirationExpirationAlertTestRequest).Execute()

Test an Expiration Alert

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
    alertsExpirationExpirationAlertTestRequest := *openapiclient.NewAlertsExpirationExpirationAlertTestRequest() // AlertsExpirationExpirationAlertTestRequest | Information about the expiration alert test (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExpirationAlertApi.BuildCreateAlertsExpirationTestRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).AlertsExpirationExpirationAlertTestRequest(alertsExpirationExpirationAlertTestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpirationAlertApi.CreateAlertsExpirationTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAlertsExpirationTest`: AlertsExpirationExpirationAlertTestResponse
    fmt.Fprintf(os.Stdout, "Response from `ExpirationAlertApi.CreateAlertsExpirationTest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAlertsExpirationTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **alertsExpirationExpirationAlertTestRequest** | [**AlertsExpirationExpirationAlertTestRequest**](AlertsExpirationExpirationAlertTestRequest.md) | Information about the expiration alert test | 

### Return type

[**AlertsExpirationExpirationAlertTestResponse**](AlertsExpirationExpirationAlertTestResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAlertsExpirationTestAll

> AlertsExpirationExpirationAlertTestResponse BuildCreateAlertsExpirationTestAllRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).AlertsExpirationExpirationAlertTestAllRequest(alertsExpirationExpirationAlertTestAllRequest).Execute()

Test All Expiration Alerts

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
    alertsExpirationExpirationAlertTestAllRequest := *openapiclient.NewAlertsExpirationExpirationAlertTestAllRequest() // AlertsExpirationExpirationAlertTestAllRequest | Information about the expiration alert test (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExpirationAlertApi.BuildCreateAlertsExpirationTestAllRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).AlertsExpirationExpirationAlertTestAllRequest(alertsExpirationExpirationAlertTestAllRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpirationAlertApi.CreateAlertsExpirationTestAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAlertsExpirationTestAll`: AlertsExpirationExpirationAlertTestResponse
    fmt.Fprintf(os.Stdout, "Response from `ExpirationAlertApi.CreateAlertsExpirationTestAll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAlertsExpirationTestAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **alertsExpirationExpirationAlertTestAllRequest** | [**AlertsExpirationExpirationAlertTestAllRequest**](AlertsExpirationExpirationAlertTestAllRequest.md) | Information about the expiration alert test | 

### Return type

[**AlertsExpirationExpirationAlertTestResponse**](AlertsExpirationExpirationAlertTestResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlertsExpirationById

> BuildDeleteAlertsExpirationByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Delete an expiration alert

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
    id := int32(56) // int32 | Id for the expiration alert
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExpirationAlertApi.BuildDeleteAlertsExpirationByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpirationAlertApi.DeleteAlertsExpirationById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id for the expiration alert | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlertsExpirationByIdRequest struct via the builder pattern


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


## GetAlertsExpiration

> []AlertsExpirationExpirationAlertDefinitionResponse BuildGetAlertsExpirationRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Gets all expiration alerts according to the provided filter and output parameters

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
    resp, r, err := apiClient.ExpirationAlertApi.BuildGetAlertsExpirationRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpirationAlertApi.GetAlertsExpiration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlertsExpiration`: []AlertsExpirationExpirationAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `ExpirationAlertApi.GetAlertsExpiration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertsExpirationRequest struct via the builder pattern


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

[**[]AlertsExpirationExpirationAlertDefinitionResponse**](AlertsExpirationExpirationAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertsExpirationById

> AlertsExpirationExpirationAlertDefinitionResponse BuildGetAlertsExpirationByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Get an expiration alert

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
    id := int32(56) // int32 | Id for the expiration alert to get
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExpirationAlertApi.BuildGetAlertsExpirationByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpirationAlertApi.GetAlertsExpirationById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlertsExpirationById`: AlertsExpirationExpirationAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `ExpirationAlertApi.GetAlertsExpirationById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id for the expiration alert to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertsExpirationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**AlertsExpirationExpirationAlertDefinitionResponse**](AlertsExpirationExpirationAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertsExpirationSchedule

> AlertsAlertScheduleAlertScheduleResponse BuildGetAlertsExpirationScheduleRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Get the schedule for expiration alerts

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
    resp, r, err := apiClient.ExpirationAlertApi.BuildGetAlertsExpirationScheduleRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpirationAlertApi.GetAlertsExpirationSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlertsExpirationSchedule`: AlertsAlertScheduleAlertScheduleResponse
    fmt.Fprintf(os.Stdout, "Response from `ExpirationAlertApi.GetAlertsExpirationSchedule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertsExpirationScheduleRequest struct via the builder pattern


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


## UpdateAlertsExpiration

> AlertsExpirationExpirationAlertDefinitionResponse BuildUpdateAlertsExpirationRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).AlertsExpirationExpirationAlertUpdateRequest(alertsExpirationExpirationAlertUpdateRequest).Execute()

Edit an expiration alert

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
    alertsExpirationExpirationAlertUpdateRequest := *openapiclient.NewAlertsExpirationExpirationAlertUpdateRequest("DisplayName_example", int32(123)) // AlertsExpirationExpirationAlertUpdateRequest | Information for the expiration alert (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExpirationAlertApi.BuildUpdateAlertsExpirationRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).AlertsExpirationExpirationAlertUpdateRequest(alertsExpirationExpirationAlertUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpirationAlertApi.UpdateAlertsExpiration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAlertsExpiration`: AlertsExpirationExpirationAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `ExpirationAlertApi.UpdateAlertsExpiration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAlertsExpirationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **alertsExpirationExpirationAlertUpdateRequest** | [**AlertsExpirationExpirationAlertUpdateRequest**](AlertsExpirationExpirationAlertUpdateRequest.md) | Information for the expiration alert | 

### Return type

[**AlertsExpirationExpirationAlertDefinitionResponse**](AlertsExpirationExpirationAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAlertsExpirationSchedule

> AlertsAlertScheduleAlertScheduleResponse BuildUpdateAlertsExpirationScheduleRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).AlertsAlertScheduleAlertScheduleRequest(alertsAlertScheduleAlertScheduleRequest).Execute()

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
    resp, r, err := apiClient.ExpirationAlertApi.BuildUpdateAlertsExpirationScheduleRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).AlertsAlertScheduleAlertScheduleRequest(alertsAlertScheduleAlertScheduleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpirationAlertApi.UpdateAlertsExpirationSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAlertsExpirationSchedule`: AlertsAlertScheduleAlertScheduleResponse
    fmt.Fprintf(os.Stdout, "Response from `ExpirationAlertApi.UpdateAlertsExpirationSchedule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAlertsExpirationScheduleRequest struct via the builder pattern


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


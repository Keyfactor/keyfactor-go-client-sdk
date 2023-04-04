# \KeyRotationAlertApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KeyRotationAlertAddKeyRotationAlert**](KeyRotationAlertApi.md#KeyRotationAlertAddKeyRotationAlert) | **Post** /Alerts/KeyRotation | Add a key rotation alert
[**KeyRotationAlertDeleteKeyRotationAlert**](KeyRotationAlertApi.md#KeyRotationAlertDeleteKeyRotationAlert) | **Delete** /Alerts/KeyRotation/{id} | Delete a key rotation alert
[**KeyRotationAlertEditKeyRotationAlert**](KeyRotationAlertApi.md#KeyRotationAlertEditKeyRotationAlert) | **Put** /Alerts/KeyRotation | Edit a key rotation alert
[**KeyRotationAlertEditSchedule**](KeyRotationAlertApi.md#KeyRotationAlertEditSchedule) | **Put** /Alerts/KeyRotation/Schedule | Edit schedule
[**KeyRotationAlertGetKeyRotationAlert**](KeyRotationAlertApi.md#KeyRotationAlertGetKeyRotationAlert) | **Get** /Alerts/KeyRotation/{id} | Get a key rotation alert
[**KeyRotationAlertGetKeyRotationAlerts**](KeyRotationAlertApi.md#KeyRotationAlertGetKeyRotationAlerts) | **Get** /Alerts/KeyRotation | Gets all key rotation alerts according to the provided filter and output parameters
[**KeyRotationAlertGetSchedule**](KeyRotationAlertApi.md#KeyRotationAlertGetSchedule) | **Get** /Alerts/KeyRotation/Schedule | Get the schedule for key rotation alerts
[**KeyRotationAlertTestAllKeyRotationAlert**](KeyRotationAlertApi.md#KeyRotationAlertTestAllKeyRotationAlert) | **Post** /Alerts/KeyRotation/TestAll | Test All Alerts
[**KeyRotationAlertTestKeyRotationAlert**](KeyRotationAlertApi.md#KeyRotationAlertTestKeyRotationAlert) | **Post** /Alerts/KeyRotation/Test | Test An Alert



## KeyRotationAlertAddKeyRotationAlert

> KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse KeyRotationAlertAddKeyRotationAlert(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Req(req).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    req := *openapiclient.NewKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest("DisplayName_example", "Subject_example", "Message_example", int32(123)) // KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest | Information for the new alert
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyRotationAlertApi.KeyRotationAlertAddKeyRotationAlert(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Req(req).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyRotationAlertApi.KeyRotationAlertAddKeyRotationAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KeyRotationAlertAddKeyRotationAlert`: KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `KeyRotationAlertApi.KeyRotationAlertAddKeyRotationAlert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKeyRotationAlertAddKeyRotationAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **req** | [**KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest**](KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest.md) | Information for the new alert | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse**](KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeyRotationAlertDeleteKeyRotationAlert

> KeyRotationAlertDeleteKeyRotationAlert(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyRotationAlertApi.KeyRotationAlertDeleteKeyRotationAlert(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyRotationAlertApi.KeyRotationAlertDeleteKeyRotationAlert``: %v\n", err)
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

Other parameters are passed through a pointer to a apiKeyRotationAlertDeleteKeyRotationAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeyRotationAlertEditKeyRotationAlert

> KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse KeyRotationAlertEditKeyRotationAlert(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Req(req).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    req := *openapiclient.NewKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequest("DisplayName_example", "Subject_example", "Message_example", int32(123)) // KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequest | Information for the key rotation alert
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyRotationAlertApi.KeyRotationAlertEditKeyRotationAlert(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Req(req).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyRotationAlertApi.KeyRotationAlertEditKeyRotationAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KeyRotationAlertEditKeyRotationAlert`: KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `KeyRotationAlertApi.KeyRotationAlertEditKeyRotationAlert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKeyRotationAlertEditKeyRotationAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **req** | [**KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequest**](KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequest.md) | Information for the key rotation alert | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse**](KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeyRotationAlertEditSchedule

> KeyfactorApiModelsAlertsAlertScheduleAlertScheduleResponse KeyRotationAlertEditSchedule(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).NewSchedule(newSchedule).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    newSchedule := *openapiclient.NewKeyfactorApiModelsAlertsAlertScheduleAlertScheduleRequest() // KeyfactorApiModelsAlertsAlertScheduleAlertScheduleRequest | 
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyRotationAlertApi.KeyRotationAlertEditSchedule(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).NewSchedule(newSchedule).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyRotationAlertApi.KeyRotationAlertEditSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KeyRotationAlertEditSchedule`: KeyfactorApiModelsAlertsAlertScheduleAlertScheduleResponse
    fmt.Fprintf(os.Stdout, "Response from `KeyRotationAlertApi.KeyRotationAlertEditSchedule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKeyRotationAlertEditScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **newSchedule** | [**KeyfactorApiModelsAlertsAlertScheduleAlertScheduleRequest**](KeyfactorApiModelsAlertsAlertScheduleAlertScheduleRequest.md) |  | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsAlertsAlertScheduleAlertScheduleResponse**](KeyfactorApiModelsAlertsAlertScheduleAlertScheduleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeyRotationAlertGetKeyRotationAlert

> KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse KeyRotationAlertGetKeyRotationAlert(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyRotationAlertApi.KeyRotationAlertGetKeyRotationAlert(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyRotationAlertApi.KeyRotationAlertGetKeyRotationAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KeyRotationAlertGetKeyRotationAlert`: KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `KeyRotationAlertApi.KeyRotationAlertGetKeyRotationAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id for the key rotation alert to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiKeyRotationAlertGetKeyRotationAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse**](KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeyRotationAlertGetKeyRotationAlerts

> []KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse KeyRotationAlertGetKeyRotationAlerts(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PagedQueryQueryString(pagedQueryQueryString).PagedQueryPageReturned(pagedQueryPageReturned).PagedQueryReturnLimit(pagedQueryReturnLimit).PagedQuerySortField(pagedQuerySortField).PagedQuerySortAscending(pagedQuerySortAscending).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")
    pagedQueryQueryString := "pagedQueryQueryString_example" // string | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) (optional)
    pagedQueryPageReturned := int32(56) // int32 | The current page within the result set to be returned (optional)
    pagedQueryReturnLimit := int32(56) // int32 | Maximum number of records to be returned in a single call (optional)
    pagedQuerySortField := "pagedQuerySortField_example" // string | Field by which the results should be sorted (view results via Management Portal for sortable columns) (optional)
    pagedQuerySortAscending := int32(56) // int32 | Field sort direction [0=ascending, 1=descending] (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyRotationAlertApi.KeyRotationAlertGetKeyRotationAlerts(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PagedQueryQueryString(pagedQueryQueryString).PagedQueryPageReturned(pagedQueryPageReturned).PagedQueryReturnLimit(pagedQueryReturnLimit).PagedQuerySortField(pagedQuerySortField).PagedQuerySortAscending(pagedQuerySortAscending).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyRotationAlertApi.KeyRotationAlertGetKeyRotationAlerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KeyRotationAlertGetKeyRotationAlerts`: []KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `KeyRotationAlertApi.KeyRotationAlertGetKeyRotationAlerts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKeyRotationAlertGetKeyRotationAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]
 **pagedQueryQueryString** | **string** | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) | 
 **pagedQueryPageReturned** | **int32** | The current page within the result set to be returned | 
 **pagedQueryReturnLimit** | **int32** | Maximum number of records to be returned in a single call | 
 **pagedQuerySortField** | **string** | Field by which the results should be sorted (view results via Management Portal for sortable columns) | 
 **pagedQuerySortAscending** | **int32** | Field sort direction [0&#x3D;ascending, 1&#x3D;descending] | 

### Return type

[**[]KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse**](KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeyRotationAlertGetSchedule

> KeyfactorApiModelsAlertsAlertScheduleAlertScheduleResponse KeyRotationAlertGetSchedule(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyRotationAlertApi.KeyRotationAlertGetSchedule(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyRotationAlertApi.KeyRotationAlertGetSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KeyRotationAlertGetSchedule`: KeyfactorApiModelsAlertsAlertScheduleAlertScheduleResponse
    fmt.Fprintf(os.Stdout, "Response from `KeyRotationAlertApi.KeyRotationAlertGetSchedule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKeyRotationAlertGetScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsAlertsAlertScheduleAlertScheduleResponse**](KeyfactorApiModelsAlertsAlertScheduleAlertScheduleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeyRotationAlertTestAllKeyRotationAlert

> KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertTestResponse KeyRotationAlertTestAllKeyRotationAlert(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).KeyRotationAlertTestRequest(keyRotationAlertTestRequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    keyRotationAlertTestRequest := *openapiclient.NewKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertTestAllRequest() // KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertTestAllRequest | Information about the key rotation alert test
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyRotationAlertApi.KeyRotationAlertTestAllKeyRotationAlert(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).KeyRotationAlertTestRequest(keyRotationAlertTestRequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyRotationAlertApi.KeyRotationAlertTestAllKeyRotationAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KeyRotationAlertTestAllKeyRotationAlert`: KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertTestResponse
    fmt.Fprintf(os.Stdout, "Response from `KeyRotationAlertApi.KeyRotationAlertTestAllKeyRotationAlert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKeyRotationAlertTestAllKeyRotationAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **keyRotationAlertTestRequest** | [**KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertTestAllRequest**](KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertTestAllRequest.md) | Information about the key rotation alert test | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertTestResponse**](KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertTestResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeyRotationAlertTestKeyRotationAlert

> KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertTestResponse KeyRotationAlertTestKeyRotationAlert(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).KeyRotationAlertTestRequest(keyRotationAlertTestRequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    keyRotationAlertTestRequest := *openapiclient.NewKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertTestRequest() // KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertTestRequest | Parameters used to test the alert
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyRotationAlertApi.KeyRotationAlertTestKeyRotationAlert(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).KeyRotationAlertTestRequest(keyRotationAlertTestRequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyRotationAlertApi.KeyRotationAlertTestKeyRotationAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KeyRotationAlertTestKeyRotationAlert`: KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertTestResponse
    fmt.Fprintf(os.Stdout, "Response from `KeyRotationAlertApi.KeyRotationAlertTestKeyRotationAlert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKeyRotationAlertTestKeyRotationAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **keyRotationAlertTestRequest** | [**KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertTestRequest**](KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertTestRequest.md) | Parameters used to test the alert | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertTestResponse**](KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertTestResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


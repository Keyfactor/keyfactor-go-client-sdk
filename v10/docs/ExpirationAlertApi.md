# \ExpirationAlertApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExpirationAlertAddExpirationAlert**](ExpirationAlertApi.md#ExpirationAlertAddExpirationAlert) | **Post** /Alerts/Expiration | Add an expiration alert
[**ExpirationAlertDeleteExpirationAlert**](ExpirationAlertApi.md#ExpirationAlertDeleteExpirationAlert) | **Delete** /Alerts/Expiration/{id} | Delete an expiration alert
[**ExpirationAlertEditExpirationAlert**](ExpirationAlertApi.md#ExpirationAlertEditExpirationAlert) | **Put** /Alerts/Expiration | Edit an expiration alert
[**ExpirationAlertEditSchedule**](ExpirationAlertApi.md#ExpirationAlertEditSchedule) | **Put** /Alerts/Expiration/Schedule | Edit schedule
[**ExpirationAlertGetExpirationAlert**](ExpirationAlertApi.md#ExpirationAlertGetExpirationAlert) | **Get** /Alerts/Expiration/{id} | Get an expiration alert
[**ExpirationAlertGetExpirationAlerts**](ExpirationAlertApi.md#ExpirationAlertGetExpirationAlerts) | **Get** /Alerts/Expiration | Gets all expiration alerts according to the provided filter and output parameters
[**ExpirationAlertGetSchedule**](ExpirationAlertApi.md#ExpirationAlertGetSchedule) | **Get** /Alerts/Expiration/Schedule | Get the schedule for expiration alerts
[**ExpirationAlertTestAllExpirationAlert**](ExpirationAlertApi.md#ExpirationAlertTestAllExpirationAlert) | **Post** /Alerts/Expiration/TestAll | Test All Expiration Alerts
[**ExpirationAlertTestExpirationAlert**](ExpirationAlertApi.md#ExpirationAlertTestExpirationAlert) | **Post** /Alerts/Expiration/Test | Test an Expiration Alert



## ExpirationAlertAddExpirationAlert

> KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse ExpirationAlertAddExpirationAlert(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Req(req).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    req := *openapiclient.NewKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest("DisplayName_example", "Subject_example", "Message_example", int64(123)) // KeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest | Information for the new alert
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExpirationAlertApi.ExpirationAlertAddExpirationAlert(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Req(req).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpirationAlertApi.ExpirationAlertAddExpirationAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExpirationAlertAddExpirationAlert`: KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `ExpirationAlertApi.ExpirationAlertAddExpirationAlert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExpirationAlertAddExpirationAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **req** | [**KeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest**](KeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest.md) | Information for the new alert | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse**](KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExpirationAlertDeleteExpirationAlert

> ExpirationAlertDeleteExpirationAlert(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    id := int64(789) // int64 | Id for the expiration alert
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExpirationAlertApi.ExpirationAlertDeleteExpirationAlert(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpirationAlertApi.ExpirationAlertDeleteExpirationAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Id for the expiration alert | 

### Other Parameters

Other parameters are passed through a pointer to a apiExpirationAlertDeleteExpirationAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

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


## ExpirationAlertEditExpirationAlert

> KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse ExpirationAlertEditExpirationAlert(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Req(req).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    req := *openapiclient.NewKeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest("DisplayName_example", "Subject_example", "Message_example", int64(123)) // KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest | Information for the expiration alert
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExpirationAlertApi.ExpirationAlertEditExpirationAlert(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Req(req).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpirationAlertApi.ExpirationAlertEditExpirationAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExpirationAlertEditExpirationAlert`: KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `ExpirationAlertApi.ExpirationAlertEditExpirationAlert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExpirationAlertEditExpirationAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **req** | [**KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest**](KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest.md) | Information for the expiration alert | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse**](KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExpirationAlertEditSchedule

> KeyfactorApiModelsAlertsAlertScheduleAlertScheduleResponse ExpirationAlertEditSchedule(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).NewSchedule(newSchedule).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExpirationAlertApi.ExpirationAlertEditSchedule(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).NewSchedule(newSchedule).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpirationAlertApi.ExpirationAlertEditSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExpirationAlertEditSchedule`: KeyfactorApiModelsAlertsAlertScheduleAlertScheduleResponse
    fmt.Fprintf(os.Stdout, "Response from `ExpirationAlertApi.ExpirationAlertEditSchedule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExpirationAlertEditScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **newSchedule** | [**KeyfactorApiModelsAlertsAlertScheduleAlertScheduleRequest**](KeyfactorApiModelsAlertsAlertScheduleAlertScheduleRequest.md) |  | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsAlertsAlertScheduleAlertScheduleResponse**](KeyfactorApiModelsAlertsAlertScheduleAlertScheduleResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExpirationAlertGetExpirationAlert

> KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse ExpirationAlertGetExpirationAlert(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    id := int64(789) // int64 | Id for the expiration alert to get
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExpirationAlertApi.ExpirationAlertGetExpirationAlert(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpirationAlertApi.ExpirationAlertGetExpirationAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExpirationAlertGetExpirationAlert`: KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `ExpirationAlertApi.ExpirationAlertGetExpirationAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Id for the expiration alert to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiExpirationAlertGetExpirationAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse**](KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExpirationAlertGetExpirationAlerts

> []KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse ExpirationAlertGetExpirationAlerts(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PagedQueryQueryString(pagedQueryQueryString).PagedQueryPageReturned(pagedQueryPageReturned).PagedQueryReturnLimit(pagedQueryReturnLimit).PagedQuerySortField(pagedQuerySortField).PagedQuerySortAscending(pagedQuerySortAscending).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")
    pagedQueryQueryString := "pagedQueryQueryString_example" // string | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) (optional)
    pagedQueryPageReturned := int64(789) // int64 | The current page within the result set to be returned (optional)
    pagedQueryReturnLimit := int64(789) // int64 | Maximum number of records to be returned in a single call (optional)
    pagedQuerySortField := "pagedQuerySortField_example" // string | Field by which the results should be sorted (view results via Management Portal for sortable columns) (optional)
    pagedQuerySortAscending := int64(789) // int64 | Field sort direction [0=ascending, 1=descending] (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExpirationAlertApi.ExpirationAlertGetExpirationAlerts(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PagedQueryQueryString(pagedQueryQueryString).PagedQueryPageReturned(pagedQueryPageReturned).PagedQueryReturnLimit(pagedQueryReturnLimit).PagedQuerySortField(pagedQuerySortField).PagedQuerySortAscending(pagedQuerySortAscending).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpirationAlertApi.ExpirationAlertGetExpirationAlerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExpirationAlertGetExpirationAlerts`: []KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `ExpirationAlertApi.ExpirationAlertGetExpirationAlerts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExpirationAlertGetExpirationAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]
 **pagedQueryQueryString** | **string** | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) | 
 **pagedQueryPageReturned** | **int64** | The current page within the result set to be returned | 
 **pagedQueryReturnLimit** | **int64** | Maximum number of records to be returned in a single call | 
 **pagedQuerySortField** | **string** | Field by which the results should be sorted (view results via Management Portal for sortable columns) | 
 **pagedQuerySortAscending** | **int64** | Field sort direction [0&#x3D;ascending, 1&#x3D;descending] | 

### Return type

[**[]KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse**](KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExpirationAlertGetSchedule

> KeyfactorApiModelsAlertsAlertScheduleAlertScheduleResponse ExpirationAlertGetSchedule(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExpirationAlertApi.ExpirationAlertGetSchedule(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpirationAlertApi.ExpirationAlertGetSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExpirationAlertGetSchedule`: KeyfactorApiModelsAlertsAlertScheduleAlertScheduleResponse
    fmt.Fprintf(os.Stdout, "Response from `ExpirationAlertApi.ExpirationAlertGetSchedule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExpirationAlertGetScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsAlertsAlertScheduleAlertScheduleResponse**](KeyfactorApiModelsAlertsAlertScheduleAlertScheduleResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExpirationAlertTestAllExpirationAlert

> KeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse ExpirationAlertTestAllExpirationAlert(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).ExpirationAlertTestRequest(expirationAlertTestRequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    expirationAlertTestRequest := *openapiclient.NewKeyfactorApiModelsAlertsExpirationExpirationAlertTestAllRequest() // KeyfactorApiModelsAlertsExpirationExpirationAlertTestAllRequest | Information about the expiration alert test
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExpirationAlertApi.ExpirationAlertTestAllExpirationAlert(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).ExpirationAlertTestRequest(expirationAlertTestRequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpirationAlertApi.ExpirationAlertTestAllExpirationAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExpirationAlertTestAllExpirationAlert`: KeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse
    fmt.Fprintf(os.Stdout, "Response from `ExpirationAlertApi.ExpirationAlertTestAllExpirationAlert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExpirationAlertTestAllExpirationAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **expirationAlertTestRequest** | [**KeyfactorApiModelsAlertsExpirationExpirationAlertTestAllRequest**](KeyfactorApiModelsAlertsExpirationExpirationAlertTestAllRequest.md) | Information about the expiration alert test | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse**](KeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExpirationAlertTestExpirationAlert

> KeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse ExpirationAlertTestExpirationAlert(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).ExpirationAlertTestRequest(expirationAlertTestRequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    expirationAlertTestRequest := *openapiclient.NewKeyfactorApiModelsAlertsExpirationExpirationAlertTestRequest() // KeyfactorApiModelsAlertsExpirationExpirationAlertTestRequest | Information about the expiration alert test
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExpirationAlertApi.ExpirationAlertTestExpirationAlert(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).ExpirationAlertTestRequest(expirationAlertTestRequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpirationAlertApi.ExpirationAlertTestExpirationAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExpirationAlertTestExpirationAlert`: KeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse
    fmt.Fprintf(os.Stdout, "Response from `ExpirationAlertApi.ExpirationAlertTestExpirationAlert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExpirationAlertTestExpirationAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **expirationAlertTestRequest** | [**KeyfactorApiModelsAlertsExpirationExpirationAlertTestRequest**](KeyfactorApiModelsAlertsExpirationExpirationAlertTestRequest.md) | Information about the expiration alert test | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse**](KeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


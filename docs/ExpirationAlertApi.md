# \ExpirationAlertApi

All URIs are relative to */Keyfactor/API*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlertsExpirationGet**](ExpirationAlertApi.md#AlertsExpirationGet) | **Get** /Alerts/Expiration | Gets all expiration alerts according to the provided filter and output parameters
[**AlertsExpirationIdDelete**](ExpirationAlertApi.md#AlertsExpirationIdDelete) | **Delete** /Alerts/Expiration/{id} | Delete an expiration alert
[**AlertsExpirationIdGet**](ExpirationAlertApi.md#AlertsExpirationIdGet) | **Get** /Alerts/Expiration/{id} | Get an expiration alert
[**AlertsExpirationPost**](ExpirationAlertApi.md#AlertsExpirationPost) | **Post** /Alerts/Expiration | Add an expiration alert
[**AlertsExpirationPut**](ExpirationAlertApi.md#AlertsExpirationPut) | **Put** /Alerts/Expiration | Edit an expiration alert
[**AlertsExpirationScheduleGet**](ExpirationAlertApi.md#AlertsExpirationScheduleGet) | **Get** /Alerts/Expiration/Schedule | Get the schedule for expiration alerts
[**AlertsExpirationSchedulePut**](ExpirationAlertApi.md#AlertsExpirationSchedulePut) | **Put** /Alerts/Expiration/Schedule | Edit schedule
[**AlertsExpirationTestAllPost**](ExpirationAlertApi.md#AlertsExpirationTestAllPost) | **Post** /Alerts/Expiration/TestAll | Test All Expiration Alerts
[**AlertsExpirationTestPost**](ExpirationAlertApi.md#AlertsExpirationTestPost) | **Post** /Alerts/Expiration/Test | Test an Expiration Alert



## AlertsExpirationGet

> []KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse AlertsExpirationGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    queryString := "queryString_example" // string |  (optional)
    pageReturned := int32(56) // int32 |  (optional)
    returnLimit := int32(56) // int32 |  (optional)
    sortField := "sortField_example" // string |  (optional)
    sortAscending := openapiclient.Keyfactor.Common.QueryableExtensionsSortOrder(0) // KeyfactorCommonQueryableExtensionsSortOrder |  (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExpirationAlertApi.AlertsExpirationGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpirationAlertApi.AlertsExpirationGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertsExpirationGet`: []KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `ExpirationAlertApi.AlertsExpirationGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertsExpirationGetRequest struct via the builder pattern


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

[**[]KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse**](KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertsExpirationIdDelete

> AlertsExpirationIdDelete(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExpirationAlertApi.AlertsExpirationIdDelete(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpirationAlertApi.AlertsExpirationIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAlertsExpirationIdDeleteRequest struct via the builder pattern


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


## AlertsExpirationIdGet

> KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse AlertsExpirationIdGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExpirationAlertApi.AlertsExpirationIdGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpirationAlertApi.AlertsExpirationIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertsExpirationIdGet`: KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `ExpirationAlertApi.AlertsExpirationIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id for the expiration alert to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertsExpirationIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse**](KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertsExpirationPost

> KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse AlertsExpirationPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest(keyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest).Execute()

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
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest("DisplayName_example", "Subject_example", "Message_example", int32(123)) // KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest | Information for the new alert (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExpirationAlertApi.AlertsExpirationPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest(keyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpirationAlertApi.AlertsExpirationPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertsExpirationPost`: KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `ExpirationAlertApi.AlertsExpirationPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertsExpirationPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest** | [**KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest**](KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest.md) | Information for the new alert | 

### Return type

[**KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse**](KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patchjson, application/json, text/json, application/*json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertsExpirationPut

> KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse AlertsExpirationPut(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest(keyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest).Execute()

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
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest("DisplayName_example", "Subject_example", "Message_example", int32(123)) // KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest | Information for the expiration alert (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExpirationAlertApi.AlertsExpirationPut(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest(keyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpirationAlertApi.AlertsExpirationPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertsExpirationPut`: KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `ExpirationAlertApi.AlertsExpirationPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertsExpirationPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest** | [**KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest**](KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest.md) | Information for the expiration alert | 

### Return type

[**KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse**](KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patchjson, application/json, text/json, application/*json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertsExpirationScheduleGet

> KeyfactorWebKeyfactorApiModelsAlertsAlertScheduleAlertScheduleResponse AlertsExpirationScheduleGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExpirationAlertApi.AlertsExpirationScheduleGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpirationAlertApi.AlertsExpirationScheduleGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertsExpirationScheduleGet`: KeyfactorWebKeyfactorApiModelsAlertsAlertScheduleAlertScheduleResponse
    fmt.Fprintf(os.Stdout, "Response from `ExpirationAlertApi.AlertsExpirationScheduleGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertsExpirationScheduleGetRequest struct via the builder pattern


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


## AlertsExpirationSchedulePut

> KeyfactorWebKeyfactorApiModelsAlertsAlertScheduleAlertScheduleResponse AlertsExpirationSchedulePut(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsAlertsAlertScheduleAlertScheduleRequest(keyfactorWebKeyfactorApiModelsAlertsAlertScheduleAlertScheduleRequest).Execute()

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
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsAlertsAlertScheduleAlertScheduleRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsAlertsAlertScheduleAlertScheduleRequest() // KeyfactorWebKeyfactorApiModelsAlertsAlertScheduleAlertScheduleRequest |  (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExpirationAlertApi.AlertsExpirationSchedulePut(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsAlertsAlertScheduleAlertScheduleRequest(keyfactorWebKeyfactorApiModelsAlertsAlertScheduleAlertScheduleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpirationAlertApi.AlertsExpirationSchedulePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertsExpirationSchedulePut`: KeyfactorWebKeyfactorApiModelsAlertsAlertScheduleAlertScheduleResponse
    fmt.Fprintf(os.Stdout, "Response from `ExpirationAlertApi.AlertsExpirationSchedulePut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertsExpirationSchedulePutRequest struct via the builder pattern


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

- **Content-Type**: application/json-patchjson, application/json, text/json, application/*json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertsExpirationTestAllPost

> KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse AlertsExpirationTestAllPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestAllRequest(keyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestAllRequest).Execute()

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
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestAllRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestAllRequest() // KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestAllRequest | Information about the expiration alert test (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExpirationAlertApi.AlertsExpirationTestAllPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestAllRequest(keyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestAllRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpirationAlertApi.AlertsExpirationTestAllPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertsExpirationTestAllPost`: KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse
    fmt.Fprintf(os.Stdout, "Response from `ExpirationAlertApi.AlertsExpirationTestAllPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertsExpirationTestAllPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestAllRequest** | [**KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestAllRequest**](KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestAllRequest.md) | Information about the expiration alert test | 

### Return type

[**KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse**](KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patchjson, application/json, text/json, application/*json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertsExpirationTestPost

> KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse AlertsExpirationTestPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestRequest(keyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestRequest).Execute()

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
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestRequest() // KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestRequest | Information about the expiration alert test (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExpirationAlertApi.AlertsExpirationTestPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestRequest(keyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpirationAlertApi.AlertsExpirationTestPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertsExpirationTestPost`: KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse
    fmt.Fprintf(os.Stdout, "Response from `ExpirationAlertApi.AlertsExpirationTestPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertsExpirationTestPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestRequest** | [**KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestRequest**](KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestRequest.md) | Information about the expiration alert test | 

### Return type

[**KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse**](KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertTestResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patchjson, application/json, text/json, application/*json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


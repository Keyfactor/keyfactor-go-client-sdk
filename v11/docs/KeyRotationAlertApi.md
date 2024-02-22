# \KeyRotationAlertApi

All URIs are relative to */Keyfactor/API*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlertsKeyRotationGet**](KeyRotationAlertApi.md#AlertsKeyRotationGet) | **Get** /Alerts/KeyRotation | Gets all key rotation alerts according to the provided filter and output parameters
[**AlertsKeyRotationIdDelete**](KeyRotationAlertApi.md#AlertsKeyRotationIdDelete) | **Delete** /Alerts/KeyRotation/{id} | Delete a key rotation alert
[**AlertsKeyRotationIdGet**](KeyRotationAlertApi.md#AlertsKeyRotationIdGet) | **Get** /Alerts/KeyRotation/{id} | Get a key rotation alert
[**AlertsKeyRotationPost**](KeyRotationAlertApi.md#AlertsKeyRotationPost) | **Post** /Alerts/KeyRotation | Add a key rotation alert
[**AlertsKeyRotationPut**](KeyRotationAlertApi.md#AlertsKeyRotationPut) | **Put** /Alerts/KeyRotation | Edit a key rotation alert
[**AlertsKeyRotationScheduleGet**](KeyRotationAlertApi.md#AlertsKeyRotationScheduleGet) | **Get** /Alerts/KeyRotation/Schedule | Get the schedule for key rotation alerts
[**AlertsKeyRotationSchedulePut**](KeyRotationAlertApi.md#AlertsKeyRotationSchedulePut) | **Put** /Alerts/KeyRotation/Schedule | Edit schedule
[**AlertsKeyRotationTestAllPost**](KeyRotationAlertApi.md#AlertsKeyRotationTestAllPost) | **Post** /Alerts/KeyRotation/TestAll | Test All Alerts
[**AlertsKeyRotationTestPost**](KeyRotationAlertApi.md#AlertsKeyRotationTestPost) | **Post** /Alerts/KeyRotation/Test | Test An Alert



## AlertsKeyRotationGet

> []KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse AlertsKeyRotationGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    queryString := "queryString_example" // string |  (optional)
    pageReturned := int32(56) // int32 |  (optional)
    returnLimit := int32(56) // int32 |  (optional)
    sortField := "sortField_example" // string |  (optional)
    sortAscending := int32(56) // int32 |  (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyRotationAlertApi.AlertsKeyRotationGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyRotationAlertApi.AlertsKeyRotationGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertsKeyRotationGet`: []KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `KeyRotationAlertApi.AlertsKeyRotationGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertsKeyRotationGetRequest struct via the builder pattern


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

[**[]KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse**](KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertsKeyRotationIdDelete

> AlertsKeyRotationIdDelete(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyRotationAlertApi.AlertsKeyRotationIdDelete(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyRotationAlertApi.AlertsKeyRotationIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAlertsKeyRotationIdDeleteRequest struct via the builder pattern


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


## AlertsKeyRotationIdGet

> KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse AlertsKeyRotationIdGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyRotationAlertApi.AlertsKeyRotationIdGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyRotationAlertApi.AlertsKeyRotationIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertsKeyRotationIdGet`: KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `KeyRotationAlertApi.AlertsKeyRotationIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id for the key rotation alert to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertsKeyRotationIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse**](KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertsKeyRotationPost

> KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse AlertsKeyRotationPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest(keyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest("DisplayName_example", "Subject_example", "Message_example", int32(123)) // KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest | Information for the new alert (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyRotationAlertApi.AlertsKeyRotationPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest(keyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyRotationAlertApi.AlertsKeyRotationPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertsKeyRotationPost`: KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `KeyRotationAlertApi.AlertsKeyRotationPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertsKeyRotationPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest** | [**KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest**](KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest.md) | Information for the new alert | 

### Return type

[**KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse**](KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertsKeyRotationPut

> KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse AlertsKeyRotationPut(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequest(keyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequest("DisplayName_example", "Subject_example", "Message_example", int32(123)) // KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequest | Information for the key rotation alert (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyRotationAlertApi.AlertsKeyRotationPut(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequest(keyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyRotationAlertApi.AlertsKeyRotationPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertsKeyRotationPut`: KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `KeyRotationAlertApi.AlertsKeyRotationPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertsKeyRotationPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequest** | [**KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequest**](KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequest.md) | Information for the key rotation alert | 

### Return type

[**KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse**](KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertsKeyRotationScheduleGet

> KeyfactorWebKeyfactorApiModelsAlertsAlertScheduleAlertScheduleResponse AlertsKeyRotationScheduleGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyRotationAlertApi.AlertsKeyRotationScheduleGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyRotationAlertApi.AlertsKeyRotationScheduleGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertsKeyRotationScheduleGet`: KeyfactorWebKeyfactorApiModelsAlertsAlertScheduleAlertScheduleResponse
    fmt.Fprintf(os.Stdout, "Response from `KeyRotationAlertApi.AlertsKeyRotationScheduleGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertsKeyRotationScheduleGetRequest struct via the builder pattern


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


## AlertsKeyRotationSchedulePut

> KeyfactorWebKeyfactorApiModelsAlertsAlertScheduleAlertScheduleResponse AlertsKeyRotationSchedulePut(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsAlertsAlertScheduleAlertScheduleRequest(keyfactorWebKeyfactorApiModelsAlertsAlertScheduleAlertScheduleRequest).Execute()

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
    resp, r, err := apiClient.KeyRotationAlertApi.AlertsKeyRotationSchedulePut(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsAlertsAlertScheduleAlertScheduleRequest(keyfactorWebKeyfactorApiModelsAlertsAlertScheduleAlertScheduleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyRotationAlertApi.AlertsKeyRotationSchedulePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertsKeyRotationSchedulePut`: KeyfactorWebKeyfactorApiModelsAlertsAlertScheduleAlertScheduleResponse
    fmt.Fprintf(os.Stdout, "Response from `KeyRotationAlertApi.AlertsKeyRotationSchedulePut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertsKeyRotationSchedulePutRequest struct via the builder pattern


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


## AlertsKeyRotationTestAllPost

> KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertTestResponse AlertsKeyRotationTestAllPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertTestAllRequest(keyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertTestAllRequest).Execute()

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
    keyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertTestAllRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertTestAllRequest() // KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertTestAllRequest | Information about the key rotation alert test (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyRotationAlertApi.AlertsKeyRotationTestAllPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertTestAllRequest(keyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertTestAllRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyRotationAlertApi.AlertsKeyRotationTestAllPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertsKeyRotationTestAllPost`: KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertTestResponse
    fmt.Fprintf(os.Stdout, "Response from `KeyRotationAlertApi.AlertsKeyRotationTestAllPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertsKeyRotationTestAllPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertTestAllRequest** | [**KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertTestAllRequest**](KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertTestAllRequest.md) | Information about the key rotation alert test | 

### Return type

[**KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertTestResponse**](KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertTestResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertsKeyRotationTestPost

> KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertTestResponse AlertsKeyRotationTestPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertTestRequest(keyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertTestRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertTestRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertTestRequest() // KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertTestRequest | Parameters used to test the alert (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyRotationAlertApi.AlertsKeyRotationTestPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertTestRequest(keyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertTestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyRotationAlertApi.AlertsKeyRotationTestPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertsKeyRotationTestPost`: KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertTestResponse
    fmt.Fprintf(os.Stdout, "Response from `KeyRotationAlertApi.AlertsKeyRotationTestPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertsKeyRotationTestPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertTestRequest** | [**KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertTestRequest**](KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertTestRequest.md) | Parameters used to test the alert | 

### Return type

[**KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertTestResponse**](KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertTestResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


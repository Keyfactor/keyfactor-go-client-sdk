# \IssuedAlertApi

All URIs are relative to */Keyfactor/API*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlertsIssuedGet**](IssuedAlertApi.md#AlertsIssuedGet) | **Get** /Alerts/Issued | Gets all issued alerts according to the provided filter and output parameters
[**AlertsIssuedIdDelete**](IssuedAlertApi.md#AlertsIssuedIdDelete) | **Delete** /Alerts/Issued/{id} | Delete a issued alert
[**AlertsIssuedIdGet**](IssuedAlertApi.md#AlertsIssuedIdGet) | **Get** /Alerts/Issued/{id} | Get a issued alert
[**AlertsIssuedPost**](IssuedAlertApi.md#AlertsIssuedPost) | **Post** /Alerts/Issued | Add a issued alert
[**AlertsIssuedPut**](IssuedAlertApi.md#AlertsIssuedPut) | **Put** /Alerts/Issued | Edit a issued alert
[**AlertsIssuedScheduleGet**](IssuedAlertApi.md#AlertsIssuedScheduleGet) | **Get** /Alerts/Issued/Schedule | Get the schedule for issued alerts
[**AlertsIssuedSchedulePut**](IssuedAlertApi.md#AlertsIssuedSchedulePut) | **Put** /Alerts/Issued/Schedule | Edit schedule



## AlertsIssuedGet

> []KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse AlertsIssuedGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Gets all issued alerts according to the provided filter and output parameters

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
    resp, r, err := apiClient.IssuedAlertApi.AlertsIssuedGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuedAlertApi.AlertsIssuedGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertsIssuedGet`: []KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `IssuedAlertApi.AlertsIssuedGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertsIssuedGetRequest struct via the builder pattern


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

[**[]KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse**](KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertsIssuedIdDelete

> AlertsIssuedIdDelete(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Delete a issued alert

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
    id := int32(56) // int32 | Id for the issued alert
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuedAlertApi.AlertsIssuedIdDelete(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuedAlertApi.AlertsIssuedIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id for the issued alert | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertsIssuedIdDeleteRequest struct via the builder pattern


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


## AlertsIssuedIdGet

> KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse AlertsIssuedIdGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Get a issued alert

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
    id := int32(56) // int32 | Id for the issued alert to get
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuedAlertApi.AlertsIssuedIdGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuedAlertApi.AlertsIssuedIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertsIssuedIdGet`: KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `IssuedAlertApi.AlertsIssuedIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id for the issued alert to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertsIssuedIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse**](KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertsIssuedPost

> KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse AlertsIssuedPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest(keyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest).Execute()

Add a issued alert

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
    keyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest("DisplayName_example", "Subject_example", "Message_example") // KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest | Information for the new alert (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuedAlertApi.AlertsIssuedPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest(keyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuedAlertApi.AlertsIssuedPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertsIssuedPost`: KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `IssuedAlertApi.AlertsIssuedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertsIssuedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest** | [**KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest**](KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest.md) | Information for the new alert | 

### Return type

[**KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse**](KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertsIssuedPut

> KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse AlertsIssuedPut(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertUpdateRequest(keyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertUpdateRequest).Execute()

Edit a issued alert

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
    keyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertUpdateRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertUpdateRequest("DisplayName_example", "Subject_example", "Message_example") // KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertUpdateRequest | Information for the issued alert (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuedAlertApi.AlertsIssuedPut(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertUpdateRequest(keyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuedAlertApi.AlertsIssuedPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertsIssuedPut`: KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `IssuedAlertApi.AlertsIssuedPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertsIssuedPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertUpdateRequest** | [**KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertUpdateRequest**](KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertUpdateRequest.md) | Information for the issued alert | 

### Return type

[**KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse**](KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertsIssuedScheduleGet

> KeyfactorWebKeyfactorApiModelsAlertsAlertScheduleAlertScheduleResponse AlertsIssuedScheduleGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Get the schedule for issued alerts

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
    resp, r, err := apiClient.IssuedAlertApi.AlertsIssuedScheduleGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuedAlertApi.AlertsIssuedScheduleGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertsIssuedScheduleGet`: KeyfactorWebKeyfactorApiModelsAlertsAlertScheduleAlertScheduleResponse
    fmt.Fprintf(os.Stdout, "Response from `IssuedAlertApi.AlertsIssuedScheduleGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertsIssuedScheduleGetRequest struct via the builder pattern


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


## AlertsIssuedSchedulePut

> KeyfactorWebKeyfactorApiModelsAlertsAlertScheduleAlertScheduleResponse AlertsIssuedSchedulePut(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsAlertsAlertScheduleAlertScheduleRequest(keyfactorWebKeyfactorApiModelsAlertsAlertScheduleAlertScheduleRequest).Execute()

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
    resp, r, err := apiClient.IssuedAlertApi.AlertsIssuedSchedulePut(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsAlertsAlertScheduleAlertScheduleRequest(keyfactorWebKeyfactorApiModelsAlertsAlertScheduleAlertScheduleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuedAlertApi.AlertsIssuedSchedulePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertsIssuedSchedulePut`: KeyfactorWebKeyfactorApiModelsAlertsAlertScheduleAlertScheduleResponse
    fmt.Fprintf(os.Stdout, "Response from `IssuedAlertApi.AlertsIssuedSchedulePut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertsIssuedSchedulePutRequest struct via the builder pattern


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


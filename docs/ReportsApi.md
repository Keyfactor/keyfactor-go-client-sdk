# \ReportsApi

All URIs are relative to */Keyfactor/API*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReportsCustomGet**](ReportsApi.md#ReportsCustomGet) | **Get** /Reports/Custom | Returns all custom reports according to the provided filter and output parameters
[**ReportsCustomIdDelete**](ReportsApi.md#ReportsCustomIdDelete) | **Delete** /Reports/Custom/{id} | Delete custom report that matches the id
[**ReportsCustomIdGet**](ReportsApi.md#ReportsCustomIdGet) | **Get** /Reports/Custom/{id} | Returns a single custom report that matches the id
[**ReportsCustomPost**](ReportsApi.md#ReportsCustomPost) | **Post** /Reports/Custom | Creates a custom report
[**ReportsCustomPut**](ReportsApi.md#ReportsCustomPut) | **Put** /Reports/Custom | Updates a custom report that matches the id
[**ReportsGet**](ReportsApi.md#ReportsGet) | **Get** /Reports | Returns all built-in reports according to the provided filter and output parameters
[**ReportsIdGet**](ReportsApi.md#ReportsIdGet) | **Get** /Reports/{id} | Returns a single built-in report that matches the id
[**ReportsIdParametersGet**](ReportsApi.md#ReportsIdParametersGet) | **Get** /Reports/{id}/Parameters | Get a built-in report&#39;s parameters that matches the id of the report.
[**ReportsIdParametersPut**](ReportsApi.md#ReportsIdParametersPut) | **Put** /Reports/{id}/Parameters | Update a built-in report&#39;s parameters that matches the id of the report.
[**ReportsIdSchedulesGet**](ReportsApi.md#ReportsIdSchedulesGet) | **Get** /Reports/{id}/Schedules | Get a built-in report&#39;s schedules that matches the id of the report.
[**ReportsIdSchedulesPost**](ReportsApi.md#ReportsIdSchedulesPost) | **Post** /Reports/{id}/Schedules | Create a built-in report&#39;s schedule that matches the id of the report.
[**ReportsIdSchedulesPut**](ReportsApi.md#ReportsIdSchedulesPut) | **Put** /Reports/{id}/Schedules | Update a built-in report&#39;s schedule that matches the id of the report.
[**ReportsPut**](ReportsApi.md#ReportsPut) | **Put** /Reports | Updates a single built-in report that matches the id. Only some fields can be updated.
[**ReportsSchedulesIdDelete**](ReportsApi.md#ReportsSchedulesIdDelete) | **Delete** /Reports/Schedules/{id} | Delete a built-in report&#39;s schedule that matches the id of the schedule.
[**ReportsSchedulesIdGet**](ReportsApi.md#ReportsSchedulesIdGet) | **Get** /Reports/Schedules/{id} | Get a built-in report&#39;s schedule that matches the id of the schedule.



## ReportsCustomGet

> []KeyfactorWebKeyfactorApiModelsReportsCustomReportsCustomReportResponse ReportsCustomGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).AmmendedQuery(ammendedQuery).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns all custom reports according to the provided filter and output parameters

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
    ammendedQuery := "ammendedQuery_example" // string |  (optional)
    queryString := "queryString_example" // string |  (optional)
    pageReturned := int32(56) // int32 |  (optional)
    returnLimit := int32(56) // int32 |  (optional)
    sortField := "sortField_example" // string |  (optional)
    sortAscending := openapiclient.Keyfactor.Common.QueryableExtensionsSortOrder(0) // KeyfactorCommonQueryableExtensionsSortOrder |  (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.ReportsCustomGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).AmmendedQuery(ammendedQuery).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsCustomGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsCustomGet`: []KeyfactorWebKeyfactorApiModelsReportsCustomReportsCustomReportResponse
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportsCustomGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportsCustomGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **ammendedQuery** | **string** |  | 
 **queryString** | **string** |  | 
 **pageReturned** | **int32** |  | 
 **returnLimit** | **int32** |  | 
 **sortField** | **string** |  | 
 **sortAscending** | [**KeyfactorCommonQueryableExtensionsSortOrder**](KeyfactorCommonQueryableExtensionsSortOrder.md) |  | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]KeyfactorWebKeyfactorApiModelsReportsCustomReportsCustomReportResponse**](KeyfactorWebKeyfactorApiModelsReportsCustomReportsCustomReportResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsCustomIdDelete

> ReportsCustomIdDelete(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Delete custom report that matches the id

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
    id := int32(56) // int32 | Report identifier
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.ReportsCustomIdDelete(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsCustomIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Report identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsCustomIdDeleteRequest struct via the builder pattern


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


## ReportsCustomIdGet

> KeyfactorWebKeyfactorApiModelsReportsCustomReportsCustomReportResponse ReportsCustomIdGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns a single custom report that matches the id

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
    id := int32(56) // int32 | Report identifier
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.ReportsCustomIdGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsCustomIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsCustomIdGet`: KeyfactorWebKeyfactorApiModelsReportsCustomReportsCustomReportResponse
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportsCustomIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Report identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsCustomIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**KeyfactorWebKeyfactorApiModelsReportsCustomReportsCustomReportResponse**](KeyfactorWebKeyfactorApiModelsReportsCustomReportsCustomReportResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsCustomPost

> KeyfactorWebKeyfactorApiModelsReportsCustomReportsCustomReportResponse ReportsCustomPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsReportsCustomReportsCustomReportCreationRequest(keyfactorWebKeyfactorApiModelsReportsCustomReportsCustomReportCreationRequest).Execute()

Creates a custom report

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
    keyfactorWebKeyfactorApiModelsReportsCustomReportsCustomReportCreationRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsReportsCustomReportsCustomReportCreationRequest("CustomURL_example", "DisplayName_example", "Description_example") // KeyfactorWebKeyfactorApiModelsReportsCustomReportsCustomReportCreationRequest | Report Information (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.ReportsCustomPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsReportsCustomReportsCustomReportCreationRequest(keyfactorWebKeyfactorApiModelsReportsCustomReportsCustomReportCreationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsCustomPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsCustomPost`: KeyfactorWebKeyfactorApiModelsReportsCustomReportsCustomReportResponse
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportsCustomPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportsCustomPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsReportsCustomReportsCustomReportCreationRequest** | [**KeyfactorWebKeyfactorApiModelsReportsCustomReportsCustomReportCreationRequest**](KeyfactorWebKeyfactorApiModelsReportsCustomReportsCustomReportCreationRequest.md) | Report Information | 

### Return type

[**KeyfactorWebKeyfactorApiModelsReportsCustomReportsCustomReportResponse**](KeyfactorWebKeyfactorApiModelsReportsCustomReportsCustomReportResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patchjson, application/json, text/json, application/*json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsCustomPut

> KeyfactorWebKeyfactorApiModelsReportsCustomReportsCustomReportResponse ReportsCustomPut(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsReportsCustomReportsCustomReportUpdateRequest(keyfactorWebKeyfactorApiModelsReportsCustomReportsCustomReportUpdateRequest).Execute()

Updates a custom report that matches the id

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
    keyfactorWebKeyfactorApiModelsReportsCustomReportsCustomReportUpdateRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsReportsCustomReportsCustomReportUpdateRequest(int32(123), "CustomURL_example", "DisplayName_example", "Description_example", false, false) // KeyfactorWebKeyfactorApiModelsReportsCustomReportsCustomReportUpdateRequest | Report Information (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.ReportsCustomPut(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsReportsCustomReportsCustomReportUpdateRequest(keyfactorWebKeyfactorApiModelsReportsCustomReportsCustomReportUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsCustomPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsCustomPut`: KeyfactorWebKeyfactorApiModelsReportsCustomReportsCustomReportResponse
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportsCustomPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportsCustomPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsReportsCustomReportsCustomReportUpdateRequest** | [**KeyfactorWebKeyfactorApiModelsReportsCustomReportsCustomReportUpdateRequest**](KeyfactorWebKeyfactorApiModelsReportsCustomReportsCustomReportUpdateRequest.md) | Report Information | 

### Return type

[**KeyfactorWebKeyfactorApiModelsReportsCustomReportsCustomReportResponse**](KeyfactorWebKeyfactorApiModelsReportsCustomReportsCustomReportResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patchjson, application/json, text/json, application/*json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsGet

> []CSSCMSDataModelModelsReport ReportsGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).AmmendedQuery(ammendedQuery).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns all built-in reports according to the provided filter and output parameters

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
    ammendedQuery := "ammendedQuery_example" // string |  (optional)
    queryString := "queryString_example" // string |  (optional)
    pageReturned := int32(56) // int32 |  (optional)
    returnLimit := int32(56) // int32 |  (optional)
    sortField := "sortField_example" // string |  (optional)
    sortAscending := openapiclient.Keyfactor.Common.QueryableExtensionsSortOrder(0) // KeyfactorCommonQueryableExtensionsSortOrder |  (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.ReportsGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).AmmendedQuery(ammendedQuery).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsGet`: []CSSCMSDataModelModelsReport
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **ammendedQuery** | **string** |  | 
 **queryString** | **string** |  | 
 **pageReturned** | **int32** |  | 
 **returnLimit** | **int32** |  | 
 **sortField** | **string** |  | 
 **sortAscending** | [**KeyfactorCommonQueryableExtensionsSortOrder**](KeyfactorCommonQueryableExtensionsSortOrder.md) |  | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]CSSCMSDataModelModelsReport**](CSSCMSDataModelModelsReport.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsIdGet

> CSSCMSDataModelModelsReport ReportsIdGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns a single built-in report that matches the id

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
    id := int32(56) // int32 | Report identifier
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.ReportsIdGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsIdGet`: CSSCMSDataModelModelsReport
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Report identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**CSSCMSDataModelModelsReport**](CSSCMSDataModelModelsReport.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsIdParametersGet

> []CSSCMSDataModelModelsReportParameters ReportsIdParametersGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Get a built-in report's parameters that matches the id of the report.

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
    id := int32(56) // int32 | Report identifier
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.ReportsIdParametersGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsIdParametersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsIdParametersGet`: []CSSCMSDataModelModelsReportParameters
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportsIdParametersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Report identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsIdParametersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]CSSCMSDataModelModelsReportParameters**](CSSCMSDataModelModelsReportParameters.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsIdParametersPut

> []CSSCMSDataModelModelsReportParameters ReportsIdParametersPut(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsReportParametersRequest(cSSCMSDataModelModelsReportParametersRequest).Execute()

Update a built-in report's parameters that matches the id of the report.

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
    id := int32(56) // int32 | Report identifier
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    cSSCMSDataModelModelsReportParametersRequest := []openapiclient.CSSCMSDataModelModelsReportParametersRequest{*openapiclient.NewCSSCMSDataModelModelsReportParametersRequest()} // []CSSCMSDataModelModelsReportParametersRequest | A List of the parameters to be updated (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.ReportsIdParametersPut(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsReportParametersRequest(cSSCMSDataModelModelsReportParametersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsIdParametersPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsIdParametersPut`: []CSSCMSDataModelModelsReportParameters
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportsIdParametersPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Report identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsIdParametersPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsReportParametersRequest** | [**[]CSSCMSDataModelModelsReportParametersRequest**](CSSCMSDataModelModelsReportParametersRequest.md) | A List of the parameters to be updated | 

### Return type

[**[]CSSCMSDataModelModelsReportParameters**](CSSCMSDataModelModelsReportParameters.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patchjson, application/json, text/json, application/*json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsIdSchedulesGet

> []KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse ReportsIdSchedulesGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Get a built-in report's schedules that matches the id of the report.

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
    id := int32(56) // int32 | Report identifier
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    queryString := "queryString_example" // string | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) (optional)
    pageReturned := int32(56) // int32 | The current page within the result set to be returned (optional)
    returnLimit := int32(56) // int32 | Maximum number of records to be returned in a single call (optional)
    sortField := "sortField_example" // string | Field by which the results should be sorted (view results via Management Portal for sortable columns) (optional)
    sortAscending := openapiclient.Keyfactor.Common.QueryableExtensionsSortOrder(0) // KeyfactorCommonQueryableExtensionsSortOrder | Field sort direction [0=ascending, 1=descending] (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.ReportsIdSchedulesGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsIdSchedulesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsIdSchedulesGet`: []KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportsIdSchedulesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Report identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsIdSchedulesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **queryString** | **string** | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) | 
 **pageReturned** | **int32** | The current page within the result set to be returned | 
 **returnLimit** | **int32** | Maximum number of records to be returned in a single call | 
 **sortField** | **string** | Field by which the results should be sorted (view results via Management Portal for sortable columns) | 
 **sortAscending** | [**KeyfactorCommonQueryableExtensionsSortOrder**](KeyfactorCommonQueryableExtensionsSortOrder.md) | Field sort direction [0&#x3D;ascending, 1&#x3D;descending] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse**](KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsIdSchedulesPost

> KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse ReportsIdSchedulesPost(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleCreationRequest(keyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleCreationRequest).Execute()

Create a built-in report's schedule that matches the id of the report.

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
    id := int32(56) // int32 | Report identifier
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleCreationRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleCreationRequest("ReportFormat_example", map[string]string{"key": "Inner_example"}) // KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleCreationRequest | Report Schedule (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.ReportsIdSchedulesPost(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleCreationRequest(keyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleCreationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsIdSchedulesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsIdSchedulesPost`: KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportsIdSchedulesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Report identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsIdSchedulesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleCreationRequest** | [**KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleCreationRequest**](KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleCreationRequest.md) | Report Schedule | 

### Return type

[**KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse**](KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patchjson, application/json, text/json, application/*json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsIdSchedulesPut

> KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse ReportsIdSchedulesPut(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest(keyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest).Execute()

Update a built-in report's schedule that matches the id of the report.

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
    id := int32(56) // int32 | Report identifier
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest(int32(123), "ReportFormat_example", map[string]string{"key": "Inner_example"}) // KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest | Report Schedule (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.ReportsIdSchedulesPut(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest(keyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsIdSchedulesPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsIdSchedulesPut`: KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportsIdSchedulesPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Report identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsIdSchedulesPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest** | [**KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest**](KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest.md) | Report Schedule | 

### Return type

[**KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse**](KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patchjson, application/json, text/json, application/*json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsPut

> CSSCMSDataModelModelsReport ReportsPut(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsReportRequestModel(cSSCMSDataModelModelsReportRequestModel).Execute()

Updates a single built-in report that matches the id. Only some fields can be updated.

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
    cSSCMSDataModelModelsReportRequestModel := *openapiclient.NewCSSCMSDataModelModelsReportRequestModel() // CSSCMSDataModelModelsReportRequestModel | This object is used to update the Favorite, In Navigator and the Remove Duplicates if the 'Uses Collections' is true. (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.ReportsPut(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsReportRequestModel(cSSCMSDataModelModelsReportRequestModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsPut`: CSSCMSDataModelModelsReport
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportsPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsReportRequestModel** | [**CSSCMSDataModelModelsReportRequestModel**](CSSCMSDataModelModelsReportRequestModel.md) | This object is used to update the Favorite, In Navigator and the Remove Duplicates if the &#39;Uses Collections&#39; is true. | 

### Return type

[**CSSCMSDataModelModelsReport**](CSSCMSDataModelModelsReport.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patchjson, application/json, text/json, application/*json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsSchedulesIdDelete

> ReportsSchedulesIdDelete(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Delete a built-in report's schedule that matches the id of the schedule.

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
    id := int32(56) // int32 | Report Schedule identifier
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.ReportsSchedulesIdDelete(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsSchedulesIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Report Schedule identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsSchedulesIdDeleteRequest struct via the builder pattern


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


## ReportsSchedulesIdGet

> KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse ReportsSchedulesIdGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Get a built-in report's schedule that matches the id of the schedule.

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
    id := int32(56) // int32 | Report Schedule identifier
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.ReportsSchedulesIdGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsSchedulesIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsSchedulesIdGet`: KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportsSchedulesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Report Schedule identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsSchedulesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse**](KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


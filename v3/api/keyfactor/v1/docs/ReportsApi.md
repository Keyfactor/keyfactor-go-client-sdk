# \ReportsApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReportsByIdSchedules**](ReportsApi.md#CreateReportsByIdSchedules) | **POST** /Reports/{id}/Schedules | Create a built-in report&#39;s schedule that matches the id of the report.
[**CreateReportsCustom**](ReportsApi.md#CreateReportsCustom) | **POST** /Reports/Custom | Creates a custom report
[**DeleteReportsCustomById**](ReportsApi.md#DeleteReportsCustomById) | **DELETE** /Reports/Custom/{id} | Delete custom report that matches the id
[**DeleteReportsSchedulesById**](ReportsApi.md#DeleteReportsSchedulesById) | **DELETE** /Reports/Schedules/{id} | Delete a built-in report&#39;s schedule that matches the id of the schedule.
[**GetReports**](ReportsApi.md#GetReports) | **GET** /Reports | Returns all built-in reports according to the provided filter and output parameters
[**GetReportsById**](ReportsApi.md#GetReportsById) | **GET** /Reports/{id} | Returns a single built-in report that matches the id
[**GetReportsByIdParameters**](ReportsApi.md#GetReportsByIdParameters) | **GET** /Reports/{id}/Parameters | Get a built-in report&#39;s parameters that matches the id of the report.
[**GetReportsByIdSchedules**](ReportsApi.md#GetReportsByIdSchedules) | **GET** /Reports/{id}/Schedules | Get a built-in report&#39;s schedules that matches the id of the report.
[**GetReportsCustom**](ReportsApi.md#GetReportsCustom) | **GET** /Reports/Custom | Returns all custom reports according to the provided filter and output parameters
[**GetReportsCustomById**](ReportsApi.md#GetReportsCustomById) | **GET** /Reports/Custom/{id} | Returns a single custom report that matches the id
[**GetReportsSchedulesById**](ReportsApi.md#GetReportsSchedulesById) | **GET** /Reports/Schedules/{id} | Get a built-in report&#39;s schedule that matches the id of the schedule.
[**UpdateReports**](ReportsApi.md#UpdateReports) | **PUT** /Reports | Updates a single built-in report that matches the id. Only some fields can be updated.
[**UpdateReportsByIdParameters**](ReportsApi.md#UpdateReportsByIdParameters) | **PUT** /Reports/{id}/Parameters | Update a built-in report&#39;s parameters that matches the id of the report.
[**UpdateReportsByIdSchedules**](ReportsApi.md#UpdateReportsByIdSchedules) | **PUT** /Reports/{id}/Schedules | Update a built-in report&#39;s schedule that matches the id of the report.
[**UpdateReportsCustom**](ReportsApi.md#UpdateReportsCustom) | **PUT** /Reports/Custom | Updates a custom report that matches the id



## CreateReportsByIdSchedules

> ReportsReportSchedulesReportScheduleResponse NewCreateReportsByIdSchedulesRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).ReportsReportSchedulesReportScheduleCreationRequest(reportsReportSchedulesReportScheduleCreationRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    reportsReportSchedulesReportScheduleCreationRequest := *openapiclient.NewReportsReportSchedulesReportScheduleCreationRequest("ReportFormat_example", map[string]string{"key": "Inner_example"}) // ReportsReportSchedulesReportScheduleCreationRequest | Report Schedule (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.NewCreateReportsByIdSchedulesRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).ReportsReportSchedulesReportScheduleCreationRequest(reportsReportSchedulesReportScheduleCreationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.CreateReportsByIdSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateReportsByIdSchedules`: ReportsReportSchedulesReportScheduleResponse
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.CreateReportsByIdSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Report identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateReportsByIdSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **reportsReportSchedulesReportScheduleCreationRequest** | [**ReportsReportSchedulesReportScheduleCreationRequest**](ReportsReportSchedulesReportScheduleCreationRequest.md) | Report Schedule | 

### Return type

[**ReportsReportSchedulesReportScheduleResponse**](ReportsReportSchedulesReportScheduleResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateReportsCustom

> ReportsCustomReportsCustomReportResponse NewCreateReportsCustomRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).ReportsCustomReportsCustomReportCreationRequest(reportsCustomReportsCustomReportCreationRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    reportsCustomReportsCustomReportCreationRequest := *openapiclient.NewReportsCustomReportsCustomReportCreationRequest("CustomURL_example", "DisplayName_example", "Description_example") // ReportsCustomReportsCustomReportCreationRequest | Report Information (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.NewCreateReportsCustomRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).ReportsCustomReportsCustomReportCreationRequest(reportsCustomReportsCustomReportCreationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.CreateReportsCustom``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateReportsCustom`: ReportsCustomReportsCustomReportResponse
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.CreateReportsCustom`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReportsCustomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **reportsCustomReportsCustomReportCreationRequest** | [**ReportsCustomReportsCustomReportCreationRequest**](ReportsCustomReportsCustomReportCreationRequest.md) | Report Information | 

### Return type

[**ReportsCustomReportsCustomReportResponse**](ReportsCustomReportsCustomReportResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteReportsCustomById

> NewDeleteReportsCustomByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.NewDeleteReportsCustomByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.DeleteReportsCustomById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteReportsCustomByIdRequest struct via the builder pattern


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


## DeleteReportsSchedulesById

> NewDeleteReportsSchedulesByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.NewDeleteReportsSchedulesByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.DeleteReportsSchedulesById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteReportsSchedulesByIdRequest struct via the builder pattern


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


## GetReports

> []CSSCMSDataModelModelsReport NewGetReportsRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    queryString := "queryString_example" // string |  (optional)
    pageReturned := int32(56) // int32 |  (optional)
    returnLimit := int32(56) // int32 |  (optional)
    sortField := "sortField_example" // string |  (optional)
    sortAscending := openapiclient.Keyfactor.Common.QueryableExtensions.SortOrder(0) // KeyfactorCommonQueryableExtensionsSortOrder |  (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.NewGetReportsRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.GetReports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReports`: []CSSCMSDataModelModelsReport
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.GetReports`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsRequest struct via the builder pattern


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

[**[]CSSCMSDataModelModelsReport**](CSSCMSDataModelModelsReport.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportsById

> CSSCMSDataModelModelsReport NewGetReportsByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.NewGetReportsByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.GetReportsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReportsById`: CSSCMSDataModelModelsReport
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.GetReportsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Report identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
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


## GetReportsByIdParameters

> []CSSCMSDataModelModelsReportParameters NewGetReportsByIdParametersRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.NewGetReportsByIdParametersRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.GetReportsByIdParameters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReportsByIdParameters`: []CSSCMSDataModelModelsReportParameters
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.GetReportsByIdParameters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Report identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsByIdParametersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
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


## GetReportsByIdSchedules

> []ReportsReportSchedulesReportScheduleResponse NewGetReportsByIdSchedulesRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    queryString := "queryString_example" // string | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) (optional)
    pageReturned := int32(56) // int32 | The current page within the result set to be returned (optional)
    returnLimit := int32(56) // int32 | Maximum number of records to be returned in a single call (optional)
    sortField := "sortField_example" // string | Field by which the results should be sorted (view results via Management Portal for sortable columns) (optional)
    sortAscending := openapiclient.Keyfactor.Common.QueryableExtensions.SortOrder(0) // KeyfactorCommonQueryableExtensionsSortOrder | Field sort direction [0=ascending, 1=descending] (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.NewGetReportsByIdSchedulesRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.GetReportsByIdSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReportsByIdSchedules`: []ReportsReportSchedulesReportScheduleResponse
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.GetReportsByIdSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Report identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsByIdSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **queryString** | **string** | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) | 
 **pageReturned** | **int32** | The current page within the result set to be returned | 
 **returnLimit** | **int32** | Maximum number of records to be returned in a single call | 
 **sortField** | **string** | Field by which the results should be sorted (view results via Management Portal for sortable columns) | 
 **sortAscending** | [**KeyfactorCommonQueryableExtensionsSortOrder**](KeyfactorCommonQueryableExtensionsSortOrder.md) | Field sort direction [0&#x3D;ascending, 1&#x3D;descending] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]ReportsReportSchedulesReportScheduleResponse**](ReportsReportSchedulesReportScheduleResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportsCustom

> []ReportsCustomReportsCustomReportResponse NewGetReportsCustomRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    queryString := "queryString_example" // string |  (optional)
    pageReturned := int32(56) // int32 |  (optional)
    returnLimit := int32(56) // int32 |  (optional)
    sortField := "sortField_example" // string |  (optional)
    sortAscending := openapiclient.Keyfactor.Common.QueryableExtensions.SortOrder(0) // KeyfactorCommonQueryableExtensionsSortOrder |  (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.NewGetReportsCustomRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.GetReportsCustom``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReportsCustom`: []ReportsCustomReportsCustomReportResponse
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.GetReportsCustom`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsCustomRequest struct via the builder pattern


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

[**[]ReportsCustomReportsCustomReportResponse**](ReportsCustomReportsCustomReportResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportsCustomById

> ReportsCustomReportsCustomReportResponse NewGetReportsCustomByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.NewGetReportsCustomByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.GetReportsCustomById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReportsCustomById`: ReportsCustomReportsCustomReportResponse
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.GetReportsCustomById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Report identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsCustomByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**ReportsCustomReportsCustomReportResponse**](ReportsCustomReportsCustomReportResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportsSchedulesById

> ReportsReportSchedulesReportScheduleResponse NewGetReportsSchedulesByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.NewGetReportsSchedulesByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.GetReportsSchedulesById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReportsSchedulesById`: ReportsReportSchedulesReportScheduleResponse
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.GetReportsSchedulesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Report Schedule identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsSchedulesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**ReportsReportSchedulesReportScheduleResponse**](ReportsReportSchedulesReportScheduleResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReports

> CSSCMSDataModelModelsReport NewUpdateReportsRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsReportRequestModel(cSSCMSDataModelModelsReportRequestModel).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    cSSCMSDataModelModelsReportRequestModel := *openapiclient.NewCSSCMSDataModelModelsReportRequestModel() // CSSCMSDataModelModelsReportRequestModel | This object is used to update the Favorite, In Navigator and the Remove Duplicates if the 'Uses Collections' is true. (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.NewUpdateReportsRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsReportRequestModel(cSSCMSDataModelModelsReportRequestModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.UpdateReports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateReports`: CSSCMSDataModelModelsReport
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.UpdateReports`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsReportRequestModel** | [**CSSCMSDataModelModelsReportRequestModel**](CSSCMSDataModelModelsReportRequestModel.md) | This object is used to update the Favorite, In Navigator and the Remove Duplicates if the &#39;Uses Collections&#39; is true. | 

### Return type

[**CSSCMSDataModelModelsReport**](CSSCMSDataModelModelsReport.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReportsByIdParameters

> []CSSCMSDataModelModelsReportParameters NewUpdateReportsByIdParametersRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsReportParametersRequest(cSSCMSDataModelModelsReportParametersRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    cSSCMSDataModelModelsReportParametersRequest := []openapiclient.CSSCMSDataModelModelsReportParametersRequest{*openapiclient.NewCSSCMSDataModelModelsReportParametersRequest()} // []CSSCMSDataModelModelsReportParametersRequest | A List of the parameters to be updated (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.NewUpdateReportsByIdParametersRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsReportParametersRequest(cSSCMSDataModelModelsReportParametersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.UpdateReportsByIdParameters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateReportsByIdParameters`: []CSSCMSDataModelModelsReportParameters
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.UpdateReportsByIdParameters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Report identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReportsByIdParametersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsReportParametersRequest** | [**[]CSSCMSDataModelModelsReportParametersRequest**](CSSCMSDataModelModelsReportParametersRequest.md) | A List of the parameters to be updated | 

### Return type

[**[]CSSCMSDataModelModelsReportParameters**](CSSCMSDataModelModelsReportParameters.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReportsByIdSchedules

> ReportsReportSchedulesReportScheduleResponse NewUpdateReportsByIdSchedulesRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).ReportsReportSchedulesReportScheduleUpdateRequest(reportsReportSchedulesReportScheduleUpdateRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    reportsReportSchedulesReportScheduleUpdateRequest := *openapiclient.NewReportsReportSchedulesReportScheduleUpdateRequest(int32(123), "ReportFormat_example", map[string]string{"key": "Inner_example"}) // ReportsReportSchedulesReportScheduleUpdateRequest | Report Schedule (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.NewUpdateReportsByIdSchedulesRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).ReportsReportSchedulesReportScheduleUpdateRequest(reportsReportSchedulesReportScheduleUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.UpdateReportsByIdSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateReportsByIdSchedules`: ReportsReportSchedulesReportScheduleResponse
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.UpdateReportsByIdSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Report identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReportsByIdSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **reportsReportSchedulesReportScheduleUpdateRequest** | [**ReportsReportSchedulesReportScheduleUpdateRequest**](ReportsReportSchedulesReportScheduleUpdateRequest.md) | Report Schedule | 

### Return type

[**ReportsReportSchedulesReportScheduleResponse**](ReportsReportSchedulesReportScheduleResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReportsCustom

> ReportsCustomReportsCustomReportResponse NewUpdateReportsCustomRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).ReportsCustomReportsCustomReportUpdateRequest(reportsCustomReportsCustomReportUpdateRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    reportsCustomReportsCustomReportUpdateRequest := *openapiclient.NewReportsCustomReportsCustomReportUpdateRequest(int32(123), "CustomURL_example", "DisplayName_example", "Description_example", false, false) // ReportsCustomReportsCustomReportUpdateRequest | Report Information (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.NewUpdateReportsCustomRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).ReportsCustomReportsCustomReportUpdateRequest(reportsCustomReportsCustomReportUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.UpdateReportsCustom``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateReportsCustom`: ReportsCustomReportsCustomReportResponse
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.UpdateReportsCustom`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReportsCustomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **reportsCustomReportsCustomReportUpdateRequest** | [**ReportsCustomReportsCustomReportUpdateRequest**](ReportsCustomReportsCustomReportUpdateRequest.md) | Report Information | 

### Return type

[**ReportsCustomReportsCustomReportResponse**](ReportsCustomReportsCustomReportResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


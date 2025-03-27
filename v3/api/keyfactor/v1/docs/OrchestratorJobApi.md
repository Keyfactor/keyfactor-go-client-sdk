# \OrchestratorJobApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrchestratorJobsAcknowledge**](OrchestratorJobApi.md#CreateOrchestratorJobsAcknowledge) | **POST** /OrchestratorJobs/Acknowledge | Acknowledges orchestrator jobs based on the provided information
[**CreateOrchestratorJobsCustom**](OrchestratorJobApi.md#CreateOrchestratorJobsCustom) | **POST** /OrchestratorJobs/Custom | Schedules a job for a custom JobType on the agent using the provided information
[**CreateOrchestratorJobsCustomBulk**](OrchestratorJobApi.md#CreateOrchestratorJobsCustomBulk) | **POST** /OrchestratorJobs/Custom/Bulk | Schedules the same job for a custom JobType on the specified agents using the provided information
[**CreateOrchestratorJobsReschedule**](OrchestratorJobApi.md#CreateOrchestratorJobsReschedule) | **POST** /OrchestratorJobs/Reschedule | Reschedules orchestrator jobs based on the provided information
[**CreateOrchestratorJobsUnschedule**](OrchestratorJobApi.md#CreateOrchestratorJobsUnschedule) | **POST** /OrchestratorJobs/Unschedule | Unschedules orchestrator jobs based on the provided information
[**GetOrchestratorJobsJobHistory**](OrchestratorJobApi.md#GetOrchestratorJobsJobHistory) | **GET** /OrchestratorJobs/JobHistory | Returns all histories of an orchestrator job according to the provided filter and output parameters
[**GetOrchestratorJobsJobStatusData**](OrchestratorJobApi.md#GetOrchestratorJobsJobStatusData) | **GET** /OrchestratorJobs/JobStatus/Data | Retrieves the results of a custom job using the provided information
[**GetOrchestratorJobsScheduledJobs**](OrchestratorJobApi.md#GetOrchestratorJobsScheduledJobs) | **GET** /OrchestratorJobs/ScheduledJobs | Returns all scheduled orchestrator jobs according to the provided filter and output parameters



## CreateOrchestratorJobsAcknowledge

> CreateOrchestratorJobsAcknowledge(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).OrchestratorJobsAcknowledgeJobRequest(orchestratorJobsAcknowledgeJobRequest).Execute()

Acknowledges orchestrator jobs based on the provided information

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
    orchestratorJobsAcknowledgeJobRequest := *openapiclient.NewOrchestratorJobsAcknowledgeJobRequest() // OrchestratorJobsAcknowledgeJobRequest | Information to identify the jobs to acknowledge, either a query or a list of job identifiers (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrchestratorJobApi.CreateOrchestratorJobsAcknowledge(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).OrchestratorJobsAcknowledgeJobRequest(orchestratorJobsAcknowledgeJobRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorJobApi.CreateOrchestratorJobsAcknowledge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrchestratorJobsAcknowledgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **orchestratorJobsAcknowledgeJobRequest** | [**OrchestratorJobsAcknowledgeJobRequest**](OrchestratorJobsAcknowledgeJobRequest.md) | Information to identify the jobs to acknowledge, either a query or a list of job identifiers | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrchestratorJobsCustom

> OrchestratorJobsJobResponse CreateOrchestratorJobsCustom(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsOrchestratorJobsScheduleJobRequest(cSSCMSDataModelModelsOrchestratorJobsScheduleJobRequest).Execute()

Schedules a job for a custom JobType on the agent using the provided information

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
    cSSCMSDataModelModelsOrchestratorJobsScheduleJobRequest := *openapiclient.NewCSSCMSDataModelModelsOrchestratorJobsScheduleJobRequest("AgentId_example", "JobTypeName_example") // CSSCMSDataModelModelsOrchestratorJobsScheduleJobRequest | Information to use to schedule the job, including the type of custom job, agent to use, and job-specific parameters. (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrchestratorJobApi.CreateOrchestratorJobsCustom(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsOrchestratorJobsScheduleJobRequest(cSSCMSDataModelModelsOrchestratorJobsScheduleJobRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorJobApi.CreateOrchestratorJobsCustom``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrchestratorJobsCustom`: OrchestratorJobsJobResponse
    fmt.Fprintf(os.Stdout, "Response from `OrchestratorJobApi.CreateOrchestratorJobsCustom`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrchestratorJobsCustomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsOrchestratorJobsScheduleJobRequest** | [**CSSCMSDataModelModelsOrchestratorJobsScheduleJobRequest**](CSSCMSDataModelModelsOrchestratorJobsScheduleJobRequest.md) | Information to use to schedule the job, including the type of custom job, agent to use, and job-specific parameters. | 

### Return type

[**OrchestratorJobsJobResponse**](OrchestratorJobsJobResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrchestratorJobsCustomBulk

> OrchestratorJobsBulkJobResponse CreateOrchestratorJobsCustomBulk(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequest(cSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequest).Execute()

Schedules the same job for a custom JobType on the specified agents using the provided information

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
    cSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequest := *openapiclient.NewCSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequest([]string{"OrchestratorIds_example"}, "JobTypeName_example") // CSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequest | Information to use to schedule the jobs, including the type of custom job, agents to use, and job-specific parameters. (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrchestratorJobApi.CreateOrchestratorJobsCustomBulk(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequest(cSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorJobApi.CreateOrchestratorJobsCustomBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrchestratorJobsCustomBulk`: OrchestratorJobsBulkJobResponse
    fmt.Fprintf(os.Stdout, "Response from `OrchestratorJobApi.CreateOrchestratorJobsCustomBulk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrchestratorJobsCustomBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequest** | [**CSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequest**](CSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequest.md) | Information to use to schedule the jobs, including the type of custom job, agents to use, and job-specific parameters. | 

### Return type

[**OrchestratorJobsBulkJobResponse**](OrchestratorJobsBulkJobResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrchestratorJobsReschedule

> CreateOrchestratorJobsReschedule(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).OrchestratorJobsRescheduleJobRequest(orchestratorJobsRescheduleJobRequest).Execute()

Reschedules orchestrator jobs based on the provided information

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
    orchestratorJobsRescheduleJobRequest := *openapiclient.NewOrchestratorJobsRescheduleJobRequest() // OrchestratorJobsRescheduleJobRequest | Information to identify the jobs to reschedule, either a query or a list of job identifiers (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrchestratorJobApi.CreateOrchestratorJobsReschedule(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).OrchestratorJobsRescheduleJobRequest(orchestratorJobsRescheduleJobRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorJobApi.CreateOrchestratorJobsReschedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrchestratorJobsRescheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **orchestratorJobsRescheduleJobRequest** | [**OrchestratorJobsRescheduleJobRequest**](OrchestratorJobsRescheduleJobRequest.md) | Information to identify the jobs to reschedule, either a query or a list of job identifiers | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrchestratorJobsUnschedule

> CreateOrchestratorJobsUnschedule(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).OrchestratorJobsUnscheduleJobRequest(orchestratorJobsUnscheduleJobRequest).Execute()

Unschedules orchestrator jobs based on the provided information

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
    orchestratorJobsUnscheduleJobRequest := *openapiclient.NewOrchestratorJobsUnscheduleJobRequest() // OrchestratorJobsUnscheduleJobRequest | Information to identify the orchestrator jobs to unschedule, either a query or a list of job identifiers (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrchestratorJobApi.CreateOrchestratorJobsUnschedule(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).OrchestratorJobsUnscheduleJobRequest(orchestratorJobsUnscheduleJobRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorJobApi.CreateOrchestratorJobsUnschedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrchestratorJobsUnscheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **orchestratorJobsUnscheduleJobRequest** | [**OrchestratorJobsUnscheduleJobRequest**](OrchestratorJobsUnscheduleJobRequest.md) | Information to identify the orchestrator jobs to unschedule, either a query or a list of job identifiers | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrchestratorJobsJobHistory

> []CertificateStoresJobHistoryResponse GetOrchestratorJobsJobHistory(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns all histories of an orchestrator job according to the provided filter and output parameters

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
    resp, r, err := apiClient.OrchestratorJobApi.GetOrchestratorJobsJobHistory(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorJobApi.GetOrchestratorJobsJobHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrchestratorJobsJobHistory`: []CertificateStoresJobHistoryResponse
    fmt.Fprintf(os.Stdout, "Response from `OrchestratorJobApi.GetOrchestratorJobsJobHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrchestratorJobsJobHistoryRequest struct via the builder pattern


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

[**[]CertificateStoresJobHistoryResponse**](CertificateStoresJobHistoryResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrchestratorJobsJobStatusData

> OrchestratorJobsCustomJobResultDataResponse GetOrchestratorJobsJobStatusData(ctx).JobHistoryId(jobHistoryId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Retrieves the results of a custom job using the provided information

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
    jobHistoryId := int64(789) // int64 | Identifier of the job history record to retrieve
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrchestratorJobApi.GetOrchestratorJobsJobStatusData(context.Background()).JobHistoryId(jobHistoryId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorJobApi.GetOrchestratorJobsJobStatusData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrchestratorJobsJobStatusData`: OrchestratorJobsCustomJobResultDataResponse
    fmt.Fprintf(os.Stdout, "Response from `OrchestratorJobApi.GetOrchestratorJobsJobStatusData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrchestratorJobsJobStatusDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobHistoryId** | **int64** | Identifier of the job history record to retrieve | 
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**OrchestratorJobsCustomJobResultDataResponse**](OrchestratorJobsCustomJobResultDataResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrchestratorJobsScheduledJobs

> []CSSCMSDataModelModelsOrchestratorJobsJob GetOrchestratorJobsScheduledJobs(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns all scheduled orchestrator jobs according to the provided filter and output parameters

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
    resp, r, err := apiClient.OrchestratorJobApi.GetOrchestratorJobsScheduledJobs(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorJobApi.GetOrchestratorJobsScheduledJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrchestratorJobsScheduledJobs`: []CSSCMSDataModelModelsOrchestratorJobsJob
    fmt.Fprintf(os.Stdout, "Response from `OrchestratorJobApi.GetOrchestratorJobsScheduledJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrchestratorJobsScheduledJobsRequest struct via the builder pattern


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

[**[]CSSCMSDataModelModelsOrchestratorJobsJob**](CSSCMSDataModelModelsOrchestratorJobsJob.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


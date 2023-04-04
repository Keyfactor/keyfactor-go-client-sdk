# \OrchestratorJobApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrchestratorJobAcknowledgeJobs**](OrchestratorJobApi.md#OrchestratorJobAcknowledgeJobs) | **Post** /OrchestratorJobs/Acknowledge | Acknowledges orchestrator jobs based on the provided information
[**OrchestratorJobGetCustomJobResultData**](OrchestratorJobApi.md#OrchestratorJobGetCustomJobResultData) | **Get** /OrchestratorJobs/JobStatus/Data | Retrieves the results of a custom job using the provided information
[**OrchestratorJobGetJobHistory**](OrchestratorJobApi.md#OrchestratorJobGetJobHistory) | **Get** /OrchestratorJobs/JobHistory | Returns all histories of an orchestrator job according to the provided filter and output parameters
[**OrchestratorJobGetScheduledJobs**](OrchestratorJobApi.md#OrchestratorJobGetScheduledJobs) | **Get** /OrchestratorJobs/ScheduledJobs | Returns all scheduled orchestrator jobs according to the provided filter and output parameters
[**OrchestratorJobRescheduleJobs**](OrchestratorJobApi.md#OrchestratorJobRescheduleJobs) | **Post** /OrchestratorJobs/Reschedule | Reschedules orchestrator jobs based on the provided information
[**OrchestratorJobScheduleBulkJob**](OrchestratorJobApi.md#OrchestratorJobScheduleBulkJob) | **Post** /OrchestratorJobs/Custom/Bulk | Schedules the same job for a custom JobType on the specified agents using the provided information
[**OrchestratorJobScheduleJob**](OrchestratorJobApi.md#OrchestratorJobScheduleJob) | **Post** /OrchestratorJobs/Custom | Schedules a job for a custom JobType on the agent using the provided information
[**OrchestratorJobUnscheduleJobs**](OrchestratorJobApi.md#OrchestratorJobUnscheduleJobs) | **Post** /OrchestratorJobs/Unschedule | Unschedules orchestrator jobs based on the provided information



## OrchestratorJobAcknowledgeJobs

> OrchestratorJobAcknowledgeJobs(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Req(req).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    req := *openapiclient.NewKeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest() // KeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest | Information to identify the jobs to acknowledge, either a query or a list of job identifiers
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrchestratorJobApi.OrchestratorJobAcknowledgeJobs(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Req(req).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorJobApi.OrchestratorJobAcknowledgeJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrchestratorJobAcknowledgeJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **req** | [**KeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest**](KeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest.md) | Information to identify the jobs to acknowledge, either a query or a list of job identifiers | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrchestratorJobGetCustomJobResultData

> KeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponse OrchestratorJobGetCustomJobResultData(ctx).JobHistoryId(jobHistoryId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrchestratorJobApi.OrchestratorJobGetCustomJobResultData(context.Background()).JobHistoryId(jobHistoryId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorJobApi.OrchestratorJobGetCustomJobResultData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrchestratorJobGetCustomJobResultData`: KeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponse
    fmt.Fprintf(os.Stdout, "Response from `OrchestratorJobApi.OrchestratorJobGetCustomJobResultData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrchestratorJobGetCustomJobResultDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobHistoryId** | **int64** | Identifier of the job history record to retrieve | 
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponse**](KeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrchestratorJobGetJobHistory

> []KeyfactorApiModelsCertificateStoresJobHistoryResponse OrchestratorJobGetJobHistory(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PqQueryString(pqQueryString).PqPageReturned(pqPageReturned).PqReturnLimit(pqReturnLimit).PqSortField(pqSortField).PqSortAscending(pqSortAscending).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")
    pqQueryString := "pqQueryString_example" // string | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) (optional)
    pqPageReturned := int32(56) // int32 | The current page within the result set to be returned (optional)
    pqReturnLimit := int32(56) // int32 | Maximum number of records to be returned in a single call (optional)
    pqSortField := "pqSortField_example" // string | Field by which the results should be sorted (view results via Management Portal for sortable columns) (optional)
    pqSortAscending := int32(56) // int32 | Field sort direction [0=ascending, 1=descending] (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrchestratorJobApi.OrchestratorJobGetJobHistory(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PqQueryString(pqQueryString).PqPageReturned(pqPageReturned).PqReturnLimit(pqReturnLimit).PqSortField(pqSortField).PqSortAscending(pqSortAscending).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorJobApi.OrchestratorJobGetJobHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrchestratorJobGetJobHistory`: []KeyfactorApiModelsCertificateStoresJobHistoryResponse
    fmt.Fprintf(os.Stdout, "Response from `OrchestratorJobApi.OrchestratorJobGetJobHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrchestratorJobGetJobHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]
 **pqQueryString** | **string** | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) | 
 **pqPageReturned** | **int32** | The current page within the result set to be returned | 
 **pqReturnLimit** | **int32** | Maximum number of records to be returned in a single call | 
 **pqSortField** | **string** | Field by which the results should be sorted (view results via Management Portal for sortable columns) | 
 **pqSortAscending** | **int32** | Field sort direction [0&#x3D;ascending, 1&#x3D;descending] | 

### Return type

[**[]KeyfactorApiModelsCertificateStoresJobHistoryResponse**](KeyfactorApiModelsCertificateStoresJobHistoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrchestratorJobGetScheduledJobs

> []ModelsOrchestratorJobsJob OrchestratorJobGetScheduledJobs(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PqQueryString(pqQueryString).PqPageReturned(pqPageReturned).PqReturnLimit(pqReturnLimit).PqSortField(pqSortField).PqSortAscending(pqSortAscending).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")
    pqQueryString := "pqQueryString_example" // string | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) (optional)
    pqPageReturned := int32(56) // int32 | The current page within the result set to be returned (optional)
    pqReturnLimit := int32(56) // int32 | Maximum number of records to be returned in a single call (optional)
    pqSortField := "pqSortField_example" // string | Field by which the results should be sorted (view results via Management Portal for sortable columns) (optional)
    pqSortAscending := int32(56) // int32 | Field sort direction [0=ascending, 1=descending] (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrchestratorJobApi.OrchestratorJobGetScheduledJobs(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PqQueryString(pqQueryString).PqPageReturned(pqPageReturned).PqReturnLimit(pqReturnLimit).PqSortField(pqSortField).PqSortAscending(pqSortAscending).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorJobApi.OrchestratorJobGetScheduledJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrchestratorJobGetScheduledJobs`: []ModelsOrchestratorJobsJob
    fmt.Fprintf(os.Stdout, "Response from `OrchestratorJobApi.OrchestratorJobGetScheduledJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrchestratorJobGetScheduledJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]
 **pqQueryString** | **string** | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) | 
 **pqPageReturned** | **int32** | The current page within the result set to be returned | 
 **pqReturnLimit** | **int32** | Maximum number of records to be returned in a single call | 
 **pqSortField** | **string** | Field by which the results should be sorted (view results via Management Portal for sortable columns) | 
 **pqSortAscending** | **int32** | Field sort direction [0&#x3D;ascending, 1&#x3D;descending] | 

### Return type

[**[]ModelsOrchestratorJobsJob**](ModelsOrchestratorJobsJob.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrchestratorJobRescheduleJobs

> OrchestratorJobRescheduleJobs(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Req(req).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    req := *openapiclient.NewKeyfactorApiModelsOrchestratorJobsRescheduleJobRequest() // KeyfactorApiModelsOrchestratorJobsRescheduleJobRequest | Information to identify the jobs to reschedule, either a query or a list of job identifiers
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrchestratorJobApi.OrchestratorJobRescheduleJobs(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Req(req).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorJobApi.OrchestratorJobRescheduleJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrchestratorJobRescheduleJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **req** | [**KeyfactorApiModelsOrchestratorJobsRescheduleJobRequest**](KeyfactorApiModelsOrchestratorJobsRescheduleJobRequest.md) | Information to identify the jobs to reschedule, either a query or a list of job identifiers | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrchestratorJobScheduleBulkJob

> KeyfactorApiModelsOrchestratorJobsBulkJobResponse OrchestratorJobScheduleBulkJob(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Req(req).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    req := *openapiclient.NewModelsOrchestratorJobsScheduleBulkJobRequest([]string{"00000000-0000-0000-0000-000000000000"}, "JobTypeName_example") // ModelsOrchestratorJobsScheduleBulkJobRequest | Information to use to schedule the jobs, including the type of custom job, agents to use, and job-specific parameters.
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrchestratorJobApi.OrchestratorJobScheduleBulkJob(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Req(req).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorJobApi.OrchestratorJobScheduleBulkJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrchestratorJobScheduleBulkJob`: KeyfactorApiModelsOrchestratorJobsBulkJobResponse
    fmt.Fprintf(os.Stdout, "Response from `OrchestratorJobApi.OrchestratorJobScheduleBulkJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrchestratorJobScheduleBulkJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **req** | [**ModelsOrchestratorJobsScheduleBulkJobRequest**](ModelsOrchestratorJobsScheduleBulkJobRequest.md) | Information to use to schedule the jobs, including the type of custom job, agents to use, and job-specific parameters. | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsOrchestratorJobsBulkJobResponse**](KeyfactorApiModelsOrchestratorJobsBulkJobResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrchestratorJobScheduleJob

> KeyfactorApiModelsOrchestratorJobsJobResponse OrchestratorJobScheduleJob(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Req(req).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    req := *openapiclient.NewModelsOrchestratorJobsScheduleJobRequest("00000000-0000-0000-0000-000000000000", "JobTypeName_example") // ModelsOrchestratorJobsScheduleJobRequest | Information to use to schedule the job, including the type of custom job, agent to use, and job-specific parameters.
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrchestratorJobApi.OrchestratorJobScheduleJob(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Req(req).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorJobApi.OrchestratorJobScheduleJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrchestratorJobScheduleJob`: KeyfactorApiModelsOrchestratorJobsJobResponse
    fmt.Fprintf(os.Stdout, "Response from `OrchestratorJobApi.OrchestratorJobScheduleJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrchestratorJobScheduleJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **req** | [**ModelsOrchestratorJobsScheduleJobRequest**](ModelsOrchestratorJobsScheduleJobRequest.md) | Information to use to schedule the job, including the type of custom job, agent to use, and job-specific parameters. | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsOrchestratorJobsJobResponse**](KeyfactorApiModelsOrchestratorJobsJobResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrchestratorJobUnscheduleJobs

> OrchestratorJobUnscheduleJobs(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Req(req).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    req := *openapiclient.NewKeyfactorApiModelsOrchestratorJobsUnscheduleJobRequest() // KeyfactorApiModelsOrchestratorJobsUnscheduleJobRequest | Information to identify the orchestrator jobs to unschedule, either a query or a list of job identifiers
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrchestratorJobApi.OrchestratorJobUnscheduleJobs(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Req(req).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorJobApi.OrchestratorJobUnscheduleJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrchestratorJobUnscheduleJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **req** | [**KeyfactorApiModelsOrchestratorJobsUnscheduleJobRequest**](KeyfactorApiModelsOrchestratorJobsUnscheduleJobRequest.md) | Information to identify the orchestrator jobs to unschedule, either a query or a list of job identifiers | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


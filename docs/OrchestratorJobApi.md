# \OrchestratorJobApi

All URIs are relative to */Keyfactor/API*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrchestratorJobsAcknowledgePost**](OrchestratorJobApi.md#OrchestratorJobsAcknowledgePost) | **Post** /OrchestratorJobs/Acknowledge | Acknowledges orchestrator jobs based on the provided information
[**OrchestratorJobsCustomBulkPost**](OrchestratorJobApi.md#OrchestratorJobsCustomBulkPost) | **Post** /OrchestratorJobs/Custom/Bulk | Schedules the same job for a custom JobType on the specified agents using the provided information
[**OrchestratorJobsCustomPost**](OrchestratorJobApi.md#OrchestratorJobsCustomPost) | **Post** /OrchestratorJobs/Custom | Schedules a job for a custom JobType on the agent using the provided information
[**OrchestratorJobsJobHistoryGet**](OrchestratorJobApi.md#OrchestratorJobsJobHistoryGet) | **Get** /OrchestratorJobs/JobHistory | Returns all histories of an orchestrator job according to the provided filter and output parameters
[**OrchestratorJobsJobStatusDataGet**](OrchestratorJobApi.md#OrchestratorJobsJobStatusDataGet) | **Get** /OrchestratorJobs/JobStatus/Data | Retrieves the results of a custom job using the provided information
[**OrchestratorJobsReschedulePost**](OrchestratorJobApi.md#OrchestratorJobsReschedulePost) | **Post** /OrchestratorJobs/Reschedule | Reschedules orchestrator jobs based on the provided information
[**OrchestratorJobsScheduledJobsGet**](OrchestratorJobApi.md#OrchestratorJobsScheduledJobsGet) | **Get** /OrchestratorJobs/ScheduledJobs | Returns all scheduled orchestrator jobs according to the provided filter and output parameters
[**OrchestratorJobsUnschedulePost**](OrchestratorJobApi.md#OrchestratorJobsUnschedulePost) | **Post** /OrchestratorJobs/Unschedule | Unschedules orchestrator jobs based on the provided information



## OrchestratorJobsAcknowledgePost

> OrchestratorJobsAcknowledgePost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest(keyfactorWebKeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest() // KeyfactorWebKeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest | Information to identify the jobs to acknowledge, either a query or a list of job identifiers (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrchestratorJobApi.OrchestratorJobsAcknowledgePost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest(keyfactorWebKeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorJobApi.OrchestratorJobsAcknowledgePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrchestratorJobsAcknowledgePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest** | [**KeyfactorWebKeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest**](KeyfactorWebKeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest.md) | Information to identify the jobs to acknowledge, either a query or a list of job identifiers | 

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


## OrchestratorJobsCustomBulkPost

> KeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponse OrchestratorJobsCustomBulkPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequest(cSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    cSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequest := *openapiclient.NewCSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequest([]string{"OrchestratorIds_example"}, "JobTypeName_example") // CSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequest | Information to use to schedule the jobs, including the type of custom job, agents to use, and job-specific parameters. (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrchestratorJobApi.OrchestratorJobsCustomBulkPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequest(cSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorJobApi.OrchestratorJobsCustomBulkPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrchestratorJobsCustomBulkPost`: KeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponse
    fmt.Fprintf(os.Stdout, "Response from `OrchestratorJobApi.OrchestratorJobsCustomBulkPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrchestratorJobsCustomBulkPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequest** | [**CSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequest**](CSSCMSDataModelModelsOrchestratorJobsScheduleBulkJobRequest.md) | Information to use to schedule the jobs, including the type of custom job, agents to use, and job-specific parameters. | 

### Return type

[**KeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponse**](KeyfactorWebKeyfactorApiModelsOrchestratorJobsBulkJobResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrchestratorJobsCustomPost

> KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobResponse OrchestratorJobsCustomPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsOrchestratorJobsScheduleJobRequest(cSSCMSDataModelModelsOrchestratorJobsScheduleJobRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    cSSCMSDataModelModelsOrchestratorJobsScheduleJobRequest := *openapiclient.NewCSSCMSDataModelModelsOrchestratorJobsScheduleJobRequest("AgentId_example", "JobTypeName_example") // CSSCMSDataModelModelsOrchestratorJobsScheduleJobRequest | Information to use to schedule the job, including the type of custom job, agent to use, and job-specific parameters. (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrchestratorJobApi.OrchestratorJobsCustomPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsOrchestratorJobsScheduleJobRequest(cSSCMSDataModelModelsOrchestratorJobsScheduleJobRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorJobApi.OrchestratorJobsCustomPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrchestratorJobsCustomPost`: KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobResponse
    fmt.Fprintf(os.Stdout, "Response from `OrchestratorJobApi.OrchestratorJobsCustomPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrchestratorJobsCustomPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsOrchestratorJobsScheduleJobRequest** | [**CSSCMSDataModelModelsOrchestratorJobsScheduleJobRequest**](CSSCMSDataModelModelsOrchestratorJobsScheduleJobRequest.md) | Information to use to schedule the job, including the type of custom job, agent to use, and job-specific parameters. | 

### Return type

[**KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobResponse**](KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrchestratorJobsJobHistoryGet

> []KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse OrchestratorJobsJobHistoryGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    queryString := "queryString_example" // string |  (optional)
    pageReturned := int32(56) // int32 |  (optional)
    returnLimit := int32(56) // int32 |  (optional)
    sortField := "sortField_example" // string |  (optional)
    sortAscending := openapiclient.Keyfactor.Common.QueryableExtensionsSortOrder(0) // KeyfactorCommonQueryableExtensionsSortOrder |  (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrchestratorJobApi.OrchestratorJobsJobHistoryGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorJobApi.OrchestratorJobsJobHistoryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrchestratorJobsJobHistoryGet`: []KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse
    fmt.Fprintf(os.Stdout, "Response from `OrchestratorJobApi.OrchestratorJobsJobHistoryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrchestratorJobsJobHistoryGetRequest struct via the builder pattern


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

[**[]KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse**](KeyfactorWebKeyfactorApiModelsCertificateStoresJobHistoryResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrchestratorJobsJobStatusDataGet

> KeyfactorWebKeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponse OrchestratorJobsJobStatusDataGet(ctx).JobHistoryId(jobHistoryId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrchestratorJobApi.OrchestratorJobsJobStatusDataGet(context.Background()).JobHistoryId(jobHistoryId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorJobApi.OrchestratorJobsJobStatusDataGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrchestratorJobsJobStatusDataGet`: KeyfactorWebKeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponse
    fmt.Fprintf(os.Stdout, "Response from `OrchestratorJobApi.OrchestratorJobsJobStatusDataGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrchestratorJobsJobStatusDataGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobHistoryId** | **int64** | Identifier of the job history record to retrieve | 
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**KeyfactorWebKeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponse**](KeyfactorWebKeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrchestratorJobsReschedulePost

> OrchestratorJobsReschedulePost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsOrchestratorJobsRescheduleJobRequest(keyfactorWebKeyfactorApiModelsOrchestratorJobsRescheduleJobRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsOrchestratorJobsRescheduleJobRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsOrchestratorJobsRescheduleJobRequest() // KeyfactorWebKeyfactorApiModelsOrchestratorJobsRescheduleJobRequest | Information to identify the jobs to reschedule, either a query or a list of job identifiers (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrchestratorJobApi.OrchestratorJobsReschedulePost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsOrchestratorJobsRescheduleJobRequest(keyfactorWebKeyfactorApiModelsOrchestratorJobsRescheduleJobRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorJobApi.OrchestratorJobsReschedulePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrchestratorJobsReschedulePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsOrchestratorJobsRescheduleJobRequest** | [**KeyfactorWebKeyfactorApiModelsOrchestratorJobsRescheduleJobRequest**](KeyfactorWebKeyfactorApiModelsOrchestratorJobsRescheduleJobRequest.md) | Information to identify the jobs to reschedule, either a query or a list of job identifiers | 

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


## OrchestratorJobsScheduledJobsGet

> []CSSCMSDataModelModelsOrchestratorJobsJob OrchestratorJobsScheduledJobsGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    queryString := "queryString_example" // string |  (optional)
    pageReturned := int32(56) // int32 |  (optional)
    returnLimit := int32(56) // int32 |  (optional)
    sortField := "sortField_example" // string |  (optional)
    sortAscending := openapiclient.Keyfactor.Common.QueryableExtensionsSortOrder(0) // KeyfactorCommonQueryableExtensionsSortOrder |  (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrchestratorJobApi.OrchestratorJobsScheduledJobsGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorJobApi.OrchestratorJobsScheduledJobsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrchestratorJobsScheduledJobsGet`: []CSSCMSDataModelModelsOrchestratorJobsJob
    fmt.Fprintf(os.Stdout, "Response from `OrchestratorJobApi.OrchestratorJobsScheduledJobsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrchestratorJobsScheduledJobsGetRequest struct via the builder pattern


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

[**[]CSSCMSDataModelModelsOrchestratorJobsJob**](CSSCMSDataModelModelsOrchestratorJobsJob.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrchestratorJobsUnschedulePost

> OrchestratorJobsUnschedulePost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsOrchestratorJobsUnscheduleJobRequest(keyfactorWebKeyfactorApiModelsOrchestratorJobsUnscheduleJobRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsOrchestratorJobsUnscheduleJobRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsOrchestratorJobsUnscheduleJobRequest() // KeyfactorWebKeyfactorApiModelsOrchestratorJobsUnscheduleJobRequest | Information to identify the orchestrator jobs to unschedule, either a query or a list of job identifiers (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrchestratorJobApi.OrchestratorJobsUnschedulePost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsOrchestratorJobsUnscheduleJobRequest(keyfactorWebKeyfactorApiModelsOrchestratorJobsUnscheduleJobRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrchestratorJobApi.OrchestratorJobsUnschedulePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrchestratorJobsUnschedulePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsOrchestratorJobsUnscheduleJobRequest** | [**KeyfactorWebKeyfactorApiModelsOrchestratorJobsUnscheduleJobRequest**](KeyfactorWebKeyfactorApiModelsOrchestratorJobsUnscheduleJobRequest.md) | Information to identify the orchestrator jobs to unschedule, either a query or a list of job identifiers | 

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


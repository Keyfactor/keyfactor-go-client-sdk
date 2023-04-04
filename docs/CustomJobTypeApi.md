# \CustomJobTypeApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomJobTypeCreateJobType**](CustomJobTypeApi.md#CustomJobTypeCreateJobType) | **Post** /JobTypes/Custom | Creates a custom job type with the provided properties
[**CustomJobTypeDeleteJobType**](CustomJobTypeApi.md#CustomJobTypeDeleteJobType) | **Delete** /JobTypes/Custom/{id} | Deletes the custom job type associated with the provided id
[**CustomJobTypeGetJobTypeById**](CustomJobTypeApi.md#CustomJobTypeGetJobTypeById) | **Get** /JobTypes/Custom/{id} | Returns a single custom job type associated with the provided id
[**CustomJobTypeGetJobTypes**](CustomJobTypeApi.md#CustomJobTypeGetJobTypes) | **Get** /JobTypes/Custom | Returns all custom job types according to the provided filter and output parameters
[**CustomJobTypeUpdateJobType**](CustomJobTypeApi.md#CustomJobTypeUpdateJobType) | **Put** /JobTypes/Custom | Updates an existing custom job type with the provided properties



## CustomJobTypeCreateJobType

> KeyfactorApiModelsOrchestratorJobsJobTypeResponse CustomJobTypeCreateJobType(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).JobType(jobType).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Creates a custom job type with the provided properties

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
    jobType := *openapiclient.NewModelsOrchestratorJobsJobTypeCreateRequest("JobTypeName_example") // ModelsOrchestratorJobsJobTypeCreateRequest | job type properties to be applied to the new job type
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomJobTypeApi.CustomJobTypeCreateJobType(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).JobType(jobType).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomJobTypeApi.CustomJobTypeCreateJobType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomJobTypeCreateJobType`: KeyfactorApiModelsOrchestratorJobsJobTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomJobTypeApi.CustomJobTypeCreateJobType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomJobTypeCreateJobTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **jobType** | [**ModelsOrchestratorJobsJobTypeCreateRequest**](ModelsOrchestratorJobsJobTypeCreateRequest.md) | job type properties to be applied to the new job type | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsOrchestratorJobsJobTypeResponse**](KeyfactorApiModelsOrchestratorJobsJobTypeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomJobTypeDeleteJobType

> CustomJobTypeDeleteJobType(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Deletes the custom job type associated with the provided id

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Keyfactor identifier (GUID) of the job type
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomJobTypeApi.CustomJobTypeDeleteJobType(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomJobTypeApi.CustomJobTypeDeleteJobType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Keyfactor identifier (GUID) of the job type | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomJobTypeDeleteJobTypeRequest struct via the builder pattern


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


## CustomJobTypeGetJobTypeById

> KeyfactorApiModelsOrchestratorJobsJobTypeResponse CustomJobTypeGetJobTypeById(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns a single custom job type associated with the provided id

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Keyfactor (GUID) identifier of the job type
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomJobTypeApi.CustomJobTypeGetJobTypeById(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomJobTypeApi.CustomJobTypeGetJobTypeById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomJobTypeGetJobTypeById`: KeyfactorApiModelsOrchestratorJobsJobTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomJobTypeApi.CustomJobTypeGetJobTypeById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Keyfactor (GUID) identifier of the job type | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomJobTypeGetJobTypeByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsOrchestratorJobsJobTypeResponse**](KeyfactorApiModelsOrchestratorJobsJobTypeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomJobTypeGetJobTypes

> []KeyfactorApiModelsOrchestratorJobsJobTypeResponse CustomJobTypeGetJobTypes(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PqQueryString(pqQueryString).PqPageReturned(pqPageReturned).PqReturnLimit(pqReturnLimit).PqSortField(pqSortField).PqSortAscending(pqSortAscending).Execute()

Returns all custom job types according to the provided filter and output parameters

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
    resp, r, err := apiClient.CustomJobTypeApi.CustomJobTypeGetJobTypes(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PqQueryString(pqQueryString).PqPageReturned(pqPageReturned).PqReturnLimit(pqReturnLimit).PqSortField(pqSortField).PqSortAscending(pqSortAscending).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomJobTypeApi.CustomJobTypeGetJobTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomJobTypeGetJobTypes`: []KeyfactorApiModelsOrchestratorJobsJobTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomJobTypeApi.CustomJobTypeGetJobTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomJobTypeGetJobTypesRequest struct via the builder pattern


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

[**[]KeyfactorApiModelsOrchestratorJobsJobTypeResponse**](KeyfactorApiModelsOrchestratorJobsJobTypeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomJobTypeUpdateJobType

> KeyfactorApiModelsOrchestratorJobsJobTypeResponse CustomJobTypeUpdateJobType(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).JobType(jobType).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Updates an existing custom job type with the provided properties

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
    jobType := *openapiclient.NewModelsOrchestratorJobsJobTypeUpdateRequest("00000000-0000-0000-0000-000000000000", "JobTypeName_example") // ModelsOrchestratorJobsJobTypeUpdateRequest | job type properties to be applied to the existing job type
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomJobTypeApi.CustomJobTypeUpdateJobType(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).JobType(jobType).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomJobTypeApi.CustomJobTypeUpdateJobType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomJobTypeUpdateJobType`: KeyfactorApiModelsOrchestratorJobsJobTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomJobTypeApi.CustomJobTypeUpdateJobType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomJobTypeUpdateJobTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **jobType** | [**ModelsOrchestratorJobsJobTypeUpdateRequest**](ModelsOrchestratorJobsJobTypeUpdateRequest.md) | job type properties to be applied to the existing job type | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsOrchestratorJobsJobTypeResponse**](KeyfactorApiModelsOrchestratorJobsJobTypeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


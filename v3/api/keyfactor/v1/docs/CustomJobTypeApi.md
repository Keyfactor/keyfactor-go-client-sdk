# \CustomJobTypeApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateJobTypesCustom**](CustomJobTypeApi.md#CreateJobTypesCustom) | **POST** /JobTypes/Custom | Creates a custom job type with the provided properties
[**DeleteJobTypesCustomById**](CustomJobTypeApi.md#DeleteJobTypesCustomById) | **DELETE** /JobTypes/Custom/{id} | Deletes the custom job type associated with the provided id
[**GetJobTypesCustom**](CustomJobTypeApi.md#GetJobTypesCustom) | **GET** /JobTypes/Custom | Returns all custom job types according to the provided filter and output parameters
[**GetJobTypesCustomById**](CustomJobTypeApi.md#GetJobTypesCustomById) | **GET** /JobTypes/Custom/{id} | Returns a single custom job type associated with the provided id
[**UpdateJobTypesCustom**](CustomJobTypeApi.md#UpdateJobTypesCustom) | **PUT** /JobTypes/Custom | Updates an existing custom job type with the provided properties



## CreateJobTypesCustom

> OrchestratorJobsJobTypeResponse CreateJobTypesCustom(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsOrchestratorJobsJobTypeCreateRequest(cSSCMSDataModelModelsOrchestratorJobsJobTypeCreateRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    cSSCMSDataModelModelsOrchestratorJobsJobTypeCreateRequest := *openapiclient.NewCSSCMSDataModelModelsOrchestratorJobsJobTypeCreateRequest("JobTypeName_example") // CSSCMSDataModelModelsOrchestratorJobsJobTypeCreateRequest | job type properties to be applied to the new job type (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomJobTypeApi.CreateJobTypesCustom(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsOrchestratorJobsJobTypeCreateRequest(cSSCMSDataModelModelsOrchestratorJobsJobTypeCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomJobTypeApi.CreateJobTypesCustom``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateJobTypesCustom`: OrchestratorJobsJobTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomJobTypeApi.CreateJobTypesCustom`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateJobTypesCustomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsOrchestratorJobsJobTypeCreateRequest** | [**CSSCMSDataModelModelsOrchestratorJobsJobTypeCreateRequest**](CSSCMSDataModelModelsOrchestratorJobsJobTypeCreateRequest.md) | job type properties to be applied to the new job type | 

### Return type

[**OrchestratorJobsJobTypeResponse**](OrchestratorJobsJobTypeResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteJobTypesCustomById

> DeleteJobTypesCustomById(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomJobTypeApi.DeleteJobTypesCustomById(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomJobTypeApi.DeleteJobTypesCustomById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteJobTypesCustomByIdRequest struct via the builder pattern


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


## GetJobTypesCustom

> []OrchestratorJobsJobTypeResponse GetJobTypesCustom(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    queryString := "queryString_example" // string |  (optional)
    pageReturned := int32(56) // int32 |  (optional)
    returnLimit := int32(56) // int32 |  (optional)
    sortField := "sortField_example" // string |  (optional)
    sortAscending := openapiclient.Keyfactor.Common.QueryableExtensions.SortOrder(0) // KeyfactorCommonQueryableExtensionsSortOrder |  (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomJobTypeApi.GetJobTypesCustom(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomJobTypeApi.GetJobTypesCustom``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJobTypesCustom`: []OrchestratorJobsJobTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomJobTypeApi.GetJobTypesCustom`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetJobTypesCustomRequest struct via the builder pattern


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

[**[]OrchestratorJobsJobTypeResponse**](OrchestratorJobsJobTypeResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobTypesCustomById

> OrchestratorJobsJobTypeResponse GetJobTypesCustomById(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomJobTypeApi.GetJobTypesCustomById(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomJobTypeApi.GetJobTypesCustomById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJobTypesCustomById`: OrchestratorJobsJobTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomJobTypeApi.GetJobTypesCustomById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Keyfactor (GUID) identifier of the job type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobTypesCustomByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**OrchestratorJobsJobTypeResponse**](OrchestratorJobsJobTypeResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateJobTypesCustom

> OrchestratorJobsJobTypeResponse UpdateJobTypesCustom(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequest(cSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    cSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequest := *openapiclient.NewCSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequest("Id_example", "JobTypeName_example") // CSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequest | job type properties to be applied to the existing job type (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomJobTypeApi.UpdateJobTypesCustom(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequest(cSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomJobTypeApi.UpdateJobTypesCustom``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateJobTypesCustom`: OrchestratorJobsJobTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomJobTypeApi.UpdateJobTypesCustom`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateJobTypesCustomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequest** | [**CSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequest**](CSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequest.md) | job type properties to be applied to the existing job type | 

### Return type

[**OrchestratorJobsJobTypeResponse**](OrchestratorJobsJobTypeResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


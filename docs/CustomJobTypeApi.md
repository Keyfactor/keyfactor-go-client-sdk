# \CustomJobTypeApi

All URIs are relative to */Keyfactor/API*

Method | HTTP request | Description
------------- | ------------- | -------------
[**JobTypesCustomGet**](CustomJobTypeApi.md#JobTypesCustomGet) | **Get** /JobTypes/Custom | Returns all custom job types according to the provided filter and output parameters
[**JobTypesCustomIdDelete**](CustomJobTypeApi.md#JobTypesCustomIdDelete) | **Delete** /JobTypes/Custom/{id} | Deletes the custom job type associated with the provided id
[**JobTypesCustomIdGet**](CustomJobTypeApi.md#JobTypesCustomIdGet) | **Get** /JobTypes/Custom/{id} | Returns a single custom job type associated with the provided id
[**JobTypesCustomPost**](CustomJobTypeApi.md#JobTypesCustomPost) | **Post** /JobTypes/Custom | Creates a custom job type with the provided properties
[**JobTypesCustomPut**](CustomJobTypeApi.md#JobTypesCustomPut) | **Put** /JobTypes/Custom | Updates an existing custom job type with the provided properties



## JobTypesCustomGet

> []KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobTypeResponse JobTypesCustomGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    queryString := "queryString_example" // string |  (optional)
    pageReturned := int32(56) // int32 |  (optional)
    returnLimit := int32(56) // int32 |  (optional)
    sortField := "sortField_example" // string |  (optional)
    sortAscending := openapiclient.Keyfactor.Common.QueryableExtensionsSortOrder(0) // KeyfactorCommonQueryableExtensionsSortOrder |  (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomJobTypeApi.JobTypesCustomGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomJobTypeApi.JobTypesCustomGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `JobTypesCustomGet`: []KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomJobTypeApi.JobTypesCustomGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiJobTypesCustomGetRequest struct via the builder pattern


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

[**[]KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobTypeResponse**](KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobTypeResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobTypesCustomIdDelete

> JobTypesCustomIdDelete(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomJobTypeApi.JobTypesCustomIdDelete(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomJobTypeApi.JobTypesCustomIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiJobTypesCustomIdDeleteRequest struct via the builder pattern


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


## JobTypesCustomIdGet

> KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobTypeResponse JobTypesCustomIdGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomJobTypeApi.JobTypesCustomIdGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomJobTypeApi.JobTypesCustomIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `JobTypesCustomIdGet`: KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomJobTypeApi.JobTypesCustomIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Keyfactor (GUID) identifier of the job type | 

### Other Parameters

Other parameters are passed through a pointer to a apiJobTypesCustomIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobTypeResponse**](KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobTypeResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobTypesCustomPost

> KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobTypeResponse JobTypesCustomPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsOrchestratorJobsJobTypeCreateRequest(cSSCMSDataModelModelsOrchestratorJobsJobTypeCreateRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    cSSCMSDataModelModelsOrchestratorJobsJobTypeCreateRequest := *openapiclient.NewCSSCMSDataModelModelsOrchestratorJobsJobTypeCreateRequest("JobTypeName_example") // CSSCMSDataModelModelsOrchestratorJobsJobTypeCreateRequest | job type properties to be applied to the new job type (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomJobTypeApi.JobTypesCustomPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsOrchestratorJobsJobTypeCreateRequest(cSSCMSDataModelModelsOrchestratorJobsJobTypeCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomJobTypeApi.JobTypesCustomPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `JobTypesCustomPost`: KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomJobTypeApi.JobTypesCustomPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiJobTypesCustomPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsOrchestratorJobsJobTypeCreateRequest** | [**CSSCMSDataModelModelsOrchestratorJobsJobTypeCreateRequest**](CSSCMSDataModelModelsOrchestratorJobsJobTypeCreateRequest.md) | job type properties to be applied to the new job type | 

### Return type

[**KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobTypeResponse**](KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobTypeResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobTypesCustomPut

> KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobTypeResponse JobTypesCustomPut(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequest(cSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    cSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequest := *openapiclient.NewCSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequest("Id_example", "JobTypeName_example") // CSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequest | job type properties to be applied to the existing job type (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomJobTypeApi.JobTypesCustomPut(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequest(cSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomJobTypeApi.JobTypesCustomPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `JobTypesCustomPut`: KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomJobTypeApi.JobTypesCustomPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiJobTypesCustomPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequest** | [**CSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequest**](CSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequest.md) | job type properties to be applied to the existing job type | 

### Return type

[**KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobTypeResponse**](KeyfactorWebKeyfactorApiModelsOrchestratorJobsJobTypeResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


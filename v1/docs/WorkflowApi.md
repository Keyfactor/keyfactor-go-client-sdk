# \WorkflowApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkflowApprovePendingRequests**](WorkflowApi.md#WorkflowApprovePendingRequests) | **Post** /Workflow/Certificates/Approve | Approves pending certificate requests associated with the provided ids
[**WorkflowDenyPendingRequests**](WorkflowApi.md#WorkflowDenyPendingRequests) | **Post** /Workflow/Certificates/Deny | Denies pending certificate requests associated with the provided ids
[**WorkflowGet**](WorkflowApi.md#WorkflowGet) | **Get** /Workflow/Certificates/Pending | Gets a collection of certificate requests based on the provided query.
[**WorkflowGetCertificateRequestDetails**](WorkflowApi.md#WorkflowGetCertificateRequestDetails) | **Get** /Workflow/Certificates/{id} | Returns certificate request details based on the provided ID.
[**WorkflowGetDenied**](WorkflowApi.md#WorkflowGetDenied) | **Get** /Workflow/Certificates/Denied | Gets a collection of denied certificate requests based on the provided query.



## WorkflowApprovePendingRequests

> ModelsWorkflowApproveDenyResult WorkflowApprovePendingRequests(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).RequestIds(requestIds).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Approves pending certificate requests associated with the provided ids

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
    requestIds := []int32{int32(123)} // []int32 | Array of Keyfactor identifiers of the certificate requests
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowApi.WorkflowApprovePendingRequests(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).RequestIds(requestIds).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowApi.WorkflowApprovePendingRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowApprovePendingRequests`: ModelsWorkflowApproveDenyResult
    fmt.Fprintf(os.Stdout, "Response from `WorkflowApi.WorkflowApprovePendingRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowApprovePendingRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **requestIds** | **[]int32** | Array of Keyfactor identifiers of the certificate requests | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsWorkflowApproveDenyResult**](ModelsWorkflowApproveDenyResult.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowDenyPendingRequests

> ModelsWorkflowApproveDenyResult WorkflowDenyPendingRequests(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Request(request).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Denies pending certificate requests associated with the provided ids

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
    request := *openapiclient.NewModelsWorkflowDenialRequest() // ModelsWorkflowDenialRequest | Keyfactor identifiers of the certificate requests to be denied and any denial comments
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowApi.WorkflowDenyPendingRequests(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Request(request).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowApi.WorkflowDenyPendingRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowDenyPendingRequests`: ModelsWorkflowApproveDenyResult
    fmt.Fprintf(os.Stdout, "Response from `WorkflowApi.WorkflowDenyPendingRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowDenyPendingRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **request** | [**ModelsWorkflowDenialRequest**](ModelsWorkflowDenialRequest.md) | Keyfactor identifiers of the certificate requests to be denied and any denial comments | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsWorkflowApproveDenyResult**](ModelsWorkflowApproveDenyResult.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowGet

> []ModelsWorkflowCertificateRequestModel WorkflowGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PagedQueryQueryString(pagedQueryQueryString).PagedQueryPageReturned(pagedQueryPageReturned).PagedQueryReturnLimit(pagedQueryReturnLimit).PagedQuerySortField(pagedQuerySortField).PagedQuerySortAscending(pagedQuerySortAscending).Execute()

Gets a collection of certificate requests based on the provided query.

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
    pagedQueryQueryString := "pagedQueryQueryString_example" // string | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) (optional)
    pagedQueryPageReturned := int32(56) // int32 | The current page within the result set to be returned (optional)
    pagedQueryReturnLimit := int32(56) // int32 | Maximum number of records to be returned in a single call (optional)
    pagedQuerySortField := "pagedQuerySortField_example" // string | Field by which the results should be sorted (view results via Management Portal for sortable columns) (optional)
    pagedQuerySortAscending := int32(56) // int32 | Field sort direction [0=ascending, 1=descending] (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowApi.WorkflowGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PagedQueryQueryString(pagedQueryQueryString).PagedQueryPageReturned(pagedQueryPageReturned).PagedQueryReturnLimit(pagedQueryReturnLimit).PagedQuerySortField(pagedQuerySortField).PagedQuerySortAscending(pagedQuerySortAscending).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowApi.WorkflowGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowGet`: []ModelsWorkflowCertificateRequestModel
    fmt.Fprintf(os.Stdout, "Response from `WorkflowApi.WorkflowGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]
 **pagedQueryQueryString** | **string** | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) | 
 **pagedQueryPageReturned** | **int32** | The current page within the result set to be returned | 
 **pagedQueryReturnLimit** | **int32** | Maximum number of records to be returned in a single call | 
 **pagedQuerySortField** | **string** | Field by which the results should be sorted (view results via Management Portal for sortable columns) | 
 **pagedQuerySortAscending** | **int32** | Field sort direction [0&#x3D;ascending, 1&#x3D;descending] | 

### Return type

[**[]ModelsWorkflowCertificateRequestModel**](ModelsWorkflowCertificateRequestModel.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowGetCertificateRequestDetails

> ModelsWorkflowCertificateRequestDetailsModel WorkflowGetCertificateRequestDetails(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns certificate request details based on the provided ID.

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
    id := int32(56) // int32 | The ID of the certificate request.
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowApi.WorkflowGetCertificateRequestDetails(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowApi.WorkflowGetCertificateRequestDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowGetCertificateRequestDetails`: ModelsWorkflowCertificateRequestDetailsModel
    fmt.Fprintf(os.Stdout, "Response from `WorkflowApi.WorkflowGetCertificateRequestDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the certificate request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowGetCertificateRequestDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsWorkflowCertificateRequestDetailsModel**](ModelsWorkflowCertificateRequestDetailsModel.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowGetDenied

> []ModelsWorkflowCertificateRequestModel WorkflowGetDenied(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PagedQueryQueryString(pagedQueryQueryString).PagedQueryPageReturned(pagedQueryPageReturned).PagedQueryReturnLimit(pagedQueryReturnLimit).PagedQuerySortField(pagedQuerySortField).PagedQuerySortAscending(pagedQuerySortAscending).Execute()

Gets a collection of denied certificate requests based on the provided query.

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
    pagedQueryQueryString := "pagedQueryQueryString_example" // string | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) (optional)
    pagedQueryPageReturned := int32(56) // int32 | The current page within the result set to be returned (optional)
    pagedQueryReturnLimit := int32(56) // int32 | Maximum number of records to be returned in a single call (optional)
    pagedQuerySortField := "pagedQuerySortField_example" // string | Field by which the results should be sorted (view results via Management Portal for sortable columns) (optional)
    pagedQuerySortAscending := int32(56) // int32 | Field sort direction [0=ascending, 1=descending] (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowApi.WorkflowGetDenied(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PagedQueryQueryString(pagedQueryQueryString).PagedQueryPageReturned(pagedQueryPageReturned).PagedQueryReturnLimit(pagedQueryReturnLimit).PagedQuerySortField(pagedQuerySortField).PagedQuerySortAscending(pagedQuerySortAscending).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowApi.WorkflowGetDenied``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowGetDenied`: []ModelsWorkflowCertificateRequestModel
    fmt.Fprintf(os.Stdout, "Response from `WorkflowApi.WorkflowGetDenied`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowGetDeniedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]
 **pagedQueryQueryString** | **string** | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) | 
 **pagedQueryPageReturned** | **int32** | The current page within the result set to be returned | 
 **pagedQueryReturnLimit** | **int32** | Maximum number of records to be returned in a single call | 
 **pagedQuerySortField** | **string** | Field by which the results should be sorted (view results via Management Portal for sortable columns) | 
 **pagedQuerySortAscending** | **int32** | Field sort direction [0&#x3D;ascending, 1&#x3D;descending] | 

### Return type

[**[]ModelsWorkflowCertificateRequestModel**](ModelsWorkflowCertificateRequestModel.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


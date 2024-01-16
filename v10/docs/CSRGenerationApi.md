# \CSRGenerationApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CSRGenerationDeleteCSR**](CSRGenerationApi.md#CSRGenerationDeleteCSR) | **Delete** /CSRGeneration/Pending/{id} | Deletes a CSR associated with the provided identifier
[**CSRGenerationDeleteCSRs**](CSRGenerationApi.md#CSRGenerationDeleteCSRs) | **Delete** /CSRGeneration/Pending | Deletes the CSRs associated with the provided identifiers
[**CSRGenerationDownload**](CSRGenerationApi.md#CSRGenerationDownload) | **Get** /CSRGeneration/Pending/{id} | Returns a previously generated CSR associated with the provided identifier
[**CSRGenerationGetPendingCSRs**](CSRGenerationApi.md#CSRGenerationGetPendingCSRs) | **Get** /CSRGeneration/Pending | Returns a list of the currently pending CSRs according to the provided query
[**CSRGenerationPostGenerate**](CSRGenerationApi.md#CSRGenerationPostGenerate) | **Post** /CSRGeneration/Generate | Generates a CSR according the properties provided



## CSRGenerationDeleteCSR

> CSRGenerationDeleteCSR(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Deletes a CSR associated with the provided identifier

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
    id := int64(789) // int64 | Keyfactor identifer of the CSR to be deleted
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CSRGenerationApi.CSRGenerationDeleteCSR(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CSRGenerationApi.CSRGenerationDeleteCSR``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Keyfactor identifer of the CSR to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiCSRGenerationDeleteCSRRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

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


## CSRGenerationDeleteCSRs

> CSRGenerationDeleteCSRs(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Ids(ids).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Deletes the CSRs associated with the provided identifiers

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
    ids := []int64{int64(123)} // []int64 | Array of Keyfactor identifiers for the CSRs to be deleted
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CSRGenerationApi.CSRGenerationDeleteCSRs(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Ids(ids).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CSRGenerationApi.CSRGenerationDeleteCSRs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCSRGenerationDeleteCSRsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **ids** | **[]int64** | Array of Keyfactor identifiers for the CSRs to be deleted | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CSRGenerationDownload

> ModelsCSRGenerationResponseModel CSRGenerationDownload(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns a previously generated CSR associated with the provided identifier

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
    id := int64(789) // int64 | Keyfactor identifier of the CSR
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CSRGenerationApi.CSRGenerationDownload(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CSRGenerationApi.CSRGenerationDownload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CSRGenerationDownload`: ModelsCSRGenerationResponseModel
    fmt.Fprintf(os.Stdout, "Response from `CSRGenerationApi.CSRGenerationDownload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Keyfactor identifier of the CSR | 

### Other Parameters

Other parameters are passed through a pointer to a apiCSRGenerationDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsCSRGenerationResponseModel**](ModelsCSRGenerationResponseModel.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CSRGenerationGetPendingCSRs

> []ModelsPendingCSRResponse CSRGenerationGetPendingCSRs(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).SqQueryString(sqQueryString).SqPageReturned(sqPageReturned).SqReturnLimit(sqReturnLimit).SqSortField(sqSortField).SqSortAscending(sqSortAscending).Execute()

Returns a list of the currently pending CSRs according to the provided query

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
    sqQueryString := "sqQueryString_example" // string | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) (optional)
    sqPageReturned := int64(789) // int64 | The current page within the result set to be returned (optional)
    sqReturnLimit := int64(789) // int64 | Maximum number of records to be returned in a single call (optional)
    sqSortField := "sqSortField_example" // string | Field by which the results should be sorted (view results via Management Portal for sortable columns) (optional)
    sqSortAscending := int64(789) // int64 | Field sort direction [0=ascending, 1=descending] (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CSRGenerationApi.CSRGenerationGetPendingCSRs(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).SqQueryString(sqQueryString).SqPageReturned(sqPageReturned).SqReturnLimit(sqReturnLimit).SqSortField(sqSortField).SqSortAscending(sqSortAscending).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CSRGenerationApi.CSRGenerationGetPendingCSRs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CSRGenerationGetPendingCSRs`: []ModelsPendingCSRResponse
    fmt.Fprintf(os.Stdout, "Response from `CSRGenerationApi.CSRGenerationGetPendingCSRs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCSRGenerationGetPendingCSRsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]
 **sqQueryString** | **string** | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) | 
 **sqPageReturned** | **int64** | The current page within the result set to be returned | 
 **sqReturnLimit** | **int64** | Maximum number of records to be returned in a single call | 
 **sqSortField** | **string** | Field by which the results should be sorted (view results via Management Portal for sortable columns) | 
 **sqSortAscending** | **int64** | Field sort direction [0&#x3D;ascending, 1&#x3D;descending] | 

### Return type

[**[]ModelsPendingCSRResponse**](ModelsPendingCSRResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CSRGenerationPostGenerate

> ModelsCSRContents CSRGenerationPostGenerate(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Context(context).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Generates a CSR according the properties provided

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
    context := *openapiclient.NewModelsEnrollmentCSRGenerationRequest("Subject_example", "KeyType_example", int64(123)) // ModelsEnrollmentCSRGenerationRequest | CSR properties used to define the request - Key type [RSA, ECC], Key sizes (ex: RSA 1024, 2048, 4096/ECC 256, 384, 521), template short name or OID
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CSRGenerationApi.CSRGenerationPostGenerate(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Context(context).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CSRGenerationApi.CSRGenerationPostGenerate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CSRGenerationPostGenerate`: ModelsCSRContents
    fmt.Fprintf(os.Stdout, "Response from `CSRGenerationApi.CSRGenerationPostGenerate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCSRGenerationPostGenerateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **context** | [**ModelsEnrollmentCSRGenerationRequest**](ModelsEnrollmentCSRGenerationRequest.md) | CSR properties used to define the request - Key type [RSA, ECC], Key sizes (ex: RSA 1024, 2048, 4096/ECC 256, 384, 521), template short name or OID | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsCSRContents**](ModelsCSRContents.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


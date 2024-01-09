# \CertificateStoreContainerApi

All URIs are relative to */Keyfactor/API*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CertificateStoreContainersGet**](CertificateStoreContainerApi.md#CertificateStoreContainersGet) | **Get** /CertificateStoreContainers | Returns all certificate store container according to the provided filter and output parameters
[**CertificateStoreContainersIdDelete**](CertificateStoreContainerApi.md#CertificateStoreContainersIdDelete) | **Delete** /CertificateStoreContainers/{id} | Delete a certificate store container
[**CertificateStoreContainersIdGet**](CertificateStoreContainerApi.md#CertificateStoreContainersIdGet) | **Get** /CertificateStoreContainers/{id} | Returns a single certificate store container that matches id
[**CertificateStoreContainersPost**](CertificateStoreContainerApi.md#CertificateStoreContainersPost) | **Post** /CertificateStoreContainers | Add a certificate store container
[**CertificateStoreContainersPut**](CertificateStoreContainerApi.md#CertificateStoreContainersPut) | **Put** /CertificateStoreContainers | Edit a certificate store container



## CertificateStoreContainersGet

> []CSSCMSDataModelModelsCertificateStoreContainerListResponse CertificateStoreContainersGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).PerformRoleCheck(performRoleCheck).RoleIdList(roleIdList).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns all certificate store container according to the provided filter and output parameters

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
    performRoleCheck := true // bool |  (optional)
    roleIdList := []int32{int32(123)} // []int32 |  (optional)
    queryString := "queryString_example" // string |  (optional)
    pageReturned := int32(56) // int32 |  (optional)
    returnLimit := int32(56) // int32 |  (optional)
    sortField := "sortField_example" // string |  (optional)
    sortAscending := openapiclient.Keyfactor.Common.QueryableExtensionsSortOrder(0) // KeyfactorCommonQueryableExtensionsSortOrder |  (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreContainerApi.CertificateStoreContainersGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).PerformRoleCheck(performRoleCheck).RoleIdList(roleIdList).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreContainerApi.CertificateStoreContainersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateStoreContainersGet`: []CSSCMSDataModelModelsCertificateStoreContainerListResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateStoreContainerApi.CertificateStoreContainersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoreContainersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **performRoleCheck** | **bool** |  | 
 **roleIdList** | **[]int32** |  | 
 **queryString** | **string** |  | 
 **pageReturned** | **int32** |  | 
 **returnLimit** | **int32** |  | 
 **sortField** | **string** |  | 
 **sortAscending** | [**KeyfactorCommonQueryableExtensionsSortOrder**](KeyfactorCommonQueryableExtensionsSortOrder.md) |  | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]CSSCMSDataModelModelsCertificateStoreContainerListResponse**](CSSCMSDataModelModelsCertificateStoreContainerListResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateStoreContainersIdDelete

> CertificateStoreContainersIdDelete(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Delete a certificate store container

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
    id := int32(56) // int32 | Id for the certificate store container
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreContainerApi.CertificateStoreContainersIdDelete(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreContainerApi.CertificateStoreContainersIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id for the certificate store container | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoreContainersIdDeleteRequest struct via the builder pattern


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


## CertificateStoreContainersIdGet

> KeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponse CertificateStoreContainersIdGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns a single certificate store container that matches id

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
    id := int32(56) // int32 | 
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreContainerApi.CertificateStoreContainersIdGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreContainerApi.CertificateStoreContainersIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateStoreContainersIdGet`: KeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateStoreContainerApi.CertificateStoreContainersIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoreContainersIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**KeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponse**](KeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateStoreContainersPost

> KeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponse CertificateStoreContainersPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsCertStoreContainerRequest(cSSCMSDataModelModelsCertStoreContainerRequest).Execute()

Add a certificate store container

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
    cSSCMSDataModelModelsCertStoreContainerRequest := *openapiclient.NewCSSCMSDataModelModelsCertStoreContainerRequest() // CSSCMSDataModelModelsCertStoreContainerRequest | Information for the new container (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreContainerApi.CertificateStoreContainersPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsCertStoreContainerRequest(cSSCMSDataModelModelsCertStoreContainerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreContainerApi.CertificateStoreContainersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateStoreContainersPost`: KeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateStoreContainerApi.CertificateStoreContainersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoreContainersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsCertStoreContainerRequest** | [**CSSCMSDataModelModelsCertStoreContainerRequest**](CSSCMSDataModelModelsCertStoreContainerRequest.md) | Information for the new container | 

### Return type

[**KeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponse**](KeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateStoreContainersPut

> KeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponse CertificateStoreContainersPut(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsCertStoreContainerRequest(cSSCMSDataModelModelsCertStoreContainerRequest).Execute()

Edit a certificate store container

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
    cSSCMSDataModelModelsCertStoreContainerRequest := *openapiclient.NewCSSCMSDataModelModelsCertStoreContainerRequest() // CSSCMSDataModelModelsCertStoreContainerRequest | Information for the updated container (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreContainerApi.CertificateStoreContainersPut(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsCertStoreContainerRequest(cSSCMSDataModelModelsCertStoreContainerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreContainerApi.CertificateStoreContainersPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateStoreContainersPut`: KeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateStoreContainerApi.CertificateStoreContainersPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoreContainersPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsCertStoreContainerRequest** | [**CSSCMSDataModelModelsCertStoreContainerRequest**](CSSCMSDataModelModelsCertStoreContainerRequest.md) | Information for the updated container | 

### Return type

[**KeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponse**](KeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


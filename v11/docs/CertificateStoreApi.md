# \CertificateStoreApi

All URIs are relative to */Keyfactor/API*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CertificateStoresApprovePost**](CertificateStoreApi.md#CertificateStoresApprovePost) | **Post** /CertificateStores/Approve | Approves the provided certificate stores to make them available for management
[**CertificateStoresAssignContainerPut**](CertificateStoreApi.md#CertificateStoresAssignContainerPut) | **Put** /CertificateStores/AssignContainer | Assigns the provided certificate stores to the provided container
[**CertificateStoresCertificatesAddPost**](CertificateStoreApi.md#CertificateStoresCertificatesAddPost) | **Post** /CertificateStores/Certificates/Add | Configures a management job to add a certificate to one or more stores with the provided schedule
[**CertificateStoresCertificatesRemovePost**](CertificateStoreApi.md#CertificateStoresCertificatesRemovePost) | **Post** /CertificateStores/Certificates/Remove | Configures a management job to remove a certificate from one or more stores with the provided schedule
[**CertificateStoresDelete**](CertificateStoreApi.md#CertificateStoresDelete) | **Delete** /CertificateStores | Deletes multiple persisted certificate store entities by their identifiers
[**CertificateStoresDiscoveryJobPut**](CertificateStoreApi.md#CertificateStoresDiscoveryJobPut) | **Put** /CertificateStores/DiscoveryJob | Configures a discovery job to locate currently unmanaged certificate stores
[**CertificateStoresGet**](CertificateStoreApi.md#CertificateStoresGet) | **Get** /CertificateStores | Returns all certificate stores according to the provided filter and output parameters
[**CertificateStoresIdDelete**](CertificateStoreApi.md#CertificateStoresIdDelete) | **Delete** /CertificateStores/{id} | Deletes a persisted certificate store by its Keyfactor identifier
[**CertificateStoresIdGet**](CertificateStoreApi.md#CertificateStoresIdGet) | **Get** /CertificateStores/{id} | Returns a single certificate store associated with the provided id
[**CertificateStoresIdInventoryGet**](CertificateStoreApi.md#CertificateStoresIdInventoryGet) | **Get** /CertificateStores/{id}/Inventory | Returns a single certificate store&#39;s inventory associated with the provided id
[**CertificateStoresPasswordPut**](CertificateStoreApi.md#CertificateStoresPasswordPut) | **Put** /CertificateStores/Password | Sets a password for the requested certificate store
[**CertificateStoresPost**](CertificateStoreApi.md#CertificateStoresPost) | **Post** /CertificateStores | Creates a new certificate store with the provided properties
[**CertificateStoresPut**](CertificateStoreApi.md#CertificateStoresPut) | **Put** /CertificateStores | Updates a given certificate store with the properties of the provided instance
[**CertificateStoresReenrollmentPost**](CertificateStoreApi.md#CertificateStoresReenrollmentPost) | **Post** /CertificateStores/Reenrollment | Schedules a certificate store for reenrollment
[**CertificateStoresSchedulePost**](CertificateStoreApi.md#CertificateStoresSchedulePost) | **Post** /CertificateStores/Schedule | Creates an inventory schedule for the provided certificate stores
[**CertificateStoresServerGet**](CertificateStoreApi.md#CertificateStoresServerGet) | **Get** /CertificateStores/Server | Returns all certificate store servers according to the provided filter and output parameters
[**CertificateStoresServerPost**](CertificateStoreApi.md#CertificateStoresServerPost) | **Post** /CertificateStores/Server | Creates a new certificate store server with the provided properties
[**CertificateStoresServerPut**](CertificateStoreApi.md#CertificateStoresServerPut) | **Put** /CertificateStores/Server | Updates a given certificate store server with the properties of the provided instance



## CertificateStoresApprovePost

> CertificateStoresApprovePost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest(keyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest).Execute()

Approves the provided certificate stores to make them available for management

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
    keyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest := []openapiclient.KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest{*openapiclient.NewKeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest()} // []KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest | Certificate stores to be approved (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.CertificateStoresApprovePost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest(keyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.CertificateStoresApprovePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoresApprovePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest** | [**[]KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest**](KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest.md) | Certificate stores to be approved | 

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


## CertificateStoresAssignContainerPut

> []KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreResponse CertificateStoresAssignContainerPut(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).ModelsContainerAssignment(modelsContainerAssignment).Execute()

Assigns the provided certificate stores to the provided container

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
    modelsContainerAssignment := *openapiclient.NewModelsContainerAssignment([]string{"KeystoreIds_example"}) // ModelsContainerAssignment | Keyfactor certificate store identifiers and the container properties (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.CertificateStoresAssignContainerPut(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).ModelsContainerAssignment(modelsContainerAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.CertificateStoresAssignContainerPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateStoresAssignContainerPut`: []KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateStoreApi.CertificateStoresAssignContainerPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoresAssignContainerPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **modelsContainerAssignment** | [**ModelsContainerAssignment**](ModelsContainerAssignment.md) | Keyfactor certificate store identifiers and the container properties | 

### Return type

[**[]KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreResponse**](KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateStoresCertificatesAddPost

> []string CertificateStoresCertificatesAddPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsCertificateStoresAddCertificateRequest(keyfactorWebKeyfactorApiModelsCertificateStoresAddCertificateRequest).Execute()

Configures a management job to add a certificate to one or more stores with the provided schedule

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
    keyfactorWebKeyfactorApiModelsCertificateStoresAddCertificateRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsCertificateStoresAddCertificateRequest(int32(123), []openapiclient.ModelsCertificateStoreEntry{*openapiclient.NewModelsCertificateStoreEntry("CertificateStoreId_example")}, *openapiclient.NewKeyfactorCommonSchedulingKeyfactorSchedule()) // KeyfactorWebKeyfactorApiModelsCertificateStoresAddCertificateRequest | Configuration details of the management job (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.CertificateStoresCertificatesAddPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsCertificateStoresAddCertificateRequest(keyfactorWebKeyfactorApiModelsCertificateStoresAddCertificateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.CertificateStoresCertificatesAddPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateStoresCertificatesAddPost`: []string
    fmt.Fprintf(os.Stdout, "Response from `CertificateStoreApi.CertificateStoresCertificatesAddPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoresCertificatesAddPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsCertificateStoresAddCertificateRequest** | [**KeyfactorWebKeyfactorApiModelsCertificateStoresAddCertificateRequest**](KeyfactorWebKeyfactorApiModelsCertificateStoresAddCertificateRequest.md) | Configuration details of the management job | 

### Return type

**[]string**

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateStoresCertificatesRemovePost

> []string CertificateStoresCertificatesRemovePost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsCertificateStoresRemoveCertificateRequest(keyfactorWebKeyfactorApiModelsCertificateStoresRemoveCertificateRequest).Execute()

Configures a management job to remove a certificate from one or more stores with the provided schedule

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
    keyfactorWebKeyfactorApiModelsCertificateStoresRemoveCertificateRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsCertificateStoresRemoveCertificateRequest([]openapiclient.ModelsCertificateLocationSpecifier{*openapiclient.NewModelsCertificateLocationSpecifier()}, *openapiclient.NewKeyfactorCommonSchedulingKeyfactorSchedule()) // KeyfactorWebKeyfactorApiModelsCertificateStoresRemoveCertificateRequest | Configuration details of the management job (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.CertificateStoresCertificatesRemovePost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsCertificateStoresRemoveCertificateRequest(keyfactorWebKeyfactorApiModelsCertificateStoresRemoveCertificateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.CertificateStoresCertificatesRemovePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateStoresCertificatesRemovePost`: []string
    fmt.Fprintf(os.Stdout, "Response from `CertificateStoreApi.CertificateStoresCertificatesRemovePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoresCertificatesRemovePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsCertificateStoresRemoveCertificateRequest** | [**KeyfactorWebKeyfactorApiModelsCertificateStoresRemoveCertificateRequest**](KeyfactorWebKeyfactorApiModelsCertificateStoresRemoveCertificateRequest.md) | Configuration details of the management job | 

### Return type

**[]string**

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateStoresDelete

> CertificateStoresDelete(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).RequestBody(requestBody).Execute()

Deletes multiple persisted certificate store entities by their identifiers



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
    requestBody := []string{"Property_example"} // []string | Array of Keyfactor identifiers (GUID) for the certificate stores to be deleted (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.CertificateStoresDelete(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.CertificateStoresDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoresDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **requestBody** | **[]string** | Array of Keyfactor identifiers (GUID) for the certificate stores to be deleted | 

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


## CertificateStoresDiscoveryJobPut

> CertificateStoresDiscoveryJobPut(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).ModelsDiscoveryJobRequest(modelsDiscoveryJobRequest).Execute()

Configures a discovery job to locate currently unmanaged certificate stores

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
    modelsDiscoveryJobRequest := *openapiclient.NewModelsDiscoveryJobRequest(int32(123)) // ModelsDiscoveryJobRequest | Configuration properties of the discovery job (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.CertificateStoresDiscoveryJobPut(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).ModelsDiscoveryJobRequest(modelsDiscoveryJobRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.CertificateStoresDiscoveryJobPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoresDiscoveryJobPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **modelsDiscoveryJobRequest** | [**ModelsDiscoveryJobRequest**](ModelsDiscoveryJobRequest.md) | Configuration properties of the discovery job | 

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


## CertificateStoresGet

> []KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreResponse CertificateStoresGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).PerformRoleCheck(performRoleCheck).RoleIdList(roleIdList).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns all certificate stores according to the provided filter and output parameters

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
    performRoleCheck := true // bool | Set to 'true' if role permissions for the current user should be validated (optional)
    roleIdList := []int32{int32(123)} // []int32 | List of Keyfactor role identifiers (integer) used to determine permissions if provided (optional)
    queryString := "queryString_example" // string | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) (optional)
    pageReturned := int32(56) // int32 | The current page within the result set to be returned (optional)
    returnLimit := int32(56) // int32 | Maximum number of records to be returned in a single call (optional)
    sortField := "sortField_example" // string | Field by which the results should be sorted (view results via Management Portal for sortable columns) (optional)
    sortAscending := int32(56) // int32 | Field sort direction [0=ascending, 1=descending] (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.CertificateStoresGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).PerformRoleCheck(performRoleCheck).RoleIdList(roleIdList).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.CertificateStoresGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateStoresGet`: []KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateStoreApi.CertificateStoresGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoresGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **performRoleCheck** | **bool** | Set to &#39;true&#39; if role permissions for the current user should be validated | 
 **roleIdList** | **[]int32** | List of Keyfactor role identifiers (integer) used to determine permissions if provided | 
 **queryString** | **string** | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) | 
 **pageReturned** | **int32** | The current page within the result set to be returned | 
 **returnLimit** | **int32** | Maximum number of records to be returned in a single call | 
 **sortField** | **string** | Field by which the results should be sorted (view results via Management Portal for sortable columns) | 
 **sortAscending** | **int32** | Field sort direction [0&#x3D;ascending, 1&#x3D;descending] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreResponse**](KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateStoresIdDelete

> CertificateStoresIdDelete(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Deletes a persisted certificate store by its Keyfactor identifier

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Keyfactor certificate store identifier (GUID)
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.CertificateStoresIdDelete(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.CertificateStoresIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Keyfactor certificate store identifier (GUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoresIdDeleteRequest struct via the builder pattern


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


## CertificateStoresIdGet

> KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreResponse CertificateStoresIdGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns a single certificate store associated with the provided id

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Keyfactor identifier (GUID) of the certificate store
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.CertificateStoresIdGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.CertificateStoresIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateStoresIdGet`: KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateStoreApi.CertificateStoresIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Keyfactor identifier (GUID) of the certificate store | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoresIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreResponse**](KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateStoresIdInventoryGet

> []ModelsCertificateStoreInventory CertificateStoresIdInventoryGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns a single certificate store's inventory associated with the provided id

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Keyfactor identifier (GUID) of the certificate store
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    pageReturned := int32(56) // int32 |  (optional)
    returnLimit := int32(56) // int32 |  (optional)
    sortField := "sortField_example" // string |  (optional)
    sortAscending := int32(56) // int32 |  (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.CertificateStoresIdInventoryGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.CertificateStoresIdInventoryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateStoresIdInventoryGet`: []ModelsCertificateStoreInventory
    fmt.Fprintf(os.Stdout, "Response from `CertificateStoreApi.CertificateStoresIdInventoryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Keyfactor identifier (GUID) of the certificate store | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoresIdInventoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **pageReturned** | **int32** |  | 
 **returnLimit** | **int32** |  | 
 **sortField** | **string** |  | 
 **sortAscending** | **int32** |  | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]ModelsCertificateStoreInventory**](ModelsCertificateStoreInventory.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateStoresPasswordPut

> CertificateStoresPasswordPut(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).ModelsCertStoreNewPasswordRequest(modelsCertStoreNewPasswordRequest).Execute()

Sets a password for the requested certificate store

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
    modelsCertStoreNewPasswordRequest := *openapiclient.NewModelsCertStoreNewPasswordRequest("CertStoreId_example", interface{}(123)) // ModelsCertStoreNewPasswordRequest | Identifier of the certificate store and the password to be applied to it (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.CertificateStoresPasswordPut(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).ModelsCertStoreNewPasswordRequest(modelsCertStoreNewPasswordRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.CertificateStoresPasswordPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoresPasswordPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **modelsCertStoreNewPasswordRequest** | [**ModelsCertStoreNewPasswordRequest**](ModelsCertStoreNewPasswordRequest.md) | Identifier of the certificate store and the password to be applied to it | 

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


## CertificateStoresPost

> KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreResponse CertificateStoresPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).ModelsCertificateStoresCertificateStoreCreateRequest(modelsCertificateStoresCertificateStoreCreateRequest).Execute()

Creates a new certificate store with the provided properties

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
    modelsCertificateStoresCertificateStoreCreateRequest := *openapiclient.NewModelsCertificateStoresCertificateStoreCreateRequest() // ModelsCertificateStoresCertificateStoreCreateRequest | Certificate store to be created with the provided properties (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.CertificateStoresPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).ModelsCertificateStoresCertificateStoreCreateRequest(modelsCertificateStoresCertificateStoreCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.CertificateStoresPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateStoresPost`: KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateStoreApi.CertificateStoresPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoresPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **modelsCertificateStoresCertificateStoreCreateRequest** | [**ModelsCertificateStoresCertificateStoreCreateRequest**](ModelsCertificateStoresCertificateStoreCreateRequest.md) | Certificate store to be created with the provided properties | 

### Return type

[**KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreResponse**](KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateStoresPut

> KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreResponse CertificateStoresPut(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).ModelsCertificateStoresCertificateStoreUpdateRequest(modelsCertificateStoresCertificateStoreUpdateRequest).Execute()

Updates a given certificate store with the properties of the provided instance

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
    modelsCertificateStoresCertificateStoreUpdateRequest := *openapiclient.NewModelsCertificateStoresCertificateStoreUpdateRequest() // ModelsCertificateStoresCertificateStoreUpdateRequest | Certificate store to be updated with the provided properties (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.CertificateStoresPut(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).ModelsCertificateStoresCertificateStoreUpdateRequest(modelsCertificateStoresCertificateStoreUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.CertificateStoresPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateStoresPut`: KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateStoreApi.CertificateStoresPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoresPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **modelsCertificateStoresCertificateStoreUpdateRequest** | [**ModelsCertificateStoresCertificateStoreUpdateRequest**](ModelsCertificateStoresCertificateStoreUpdateRequest.md) | Certificate store to be updated with the provided properties | 

### Return type

[**KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreResponse**](KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateStoresReenrollmentPost

> CertificateStoresReenrollmentPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest(keyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest).Execute()

Schedules a certificate store for reenrollment

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
    keyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest() // KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest | An object that contains a Keystore Id, a Agent Guid, a string SubjectName and string Alias (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.CertificateStoresReenrollmentPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest(keyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.CertificateStoresReenrollmentPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoresReenrollmentPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest** | [**KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest**](KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest.md) | An object that contains a Keystore Id, a Agent Guid, a string SubjectName and string Alias | 

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


## CertificateStoresSchedulePost

> CertificateStoresSchedulePost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).ModelsCertStoresSchedule(modelsCertStoresSchedule).Execute()

Creates an inventory schedule for the provided certificate stores

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
    modelsCertStoresSchedule := *openapiclient.NewModelsCertStoresSchedule([]string{"StoreIds_example"}) // ModelsCertStoresSchedule | Certificate store identifiers and the desired schedule (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.CertificateStoresSchedulePost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).ModelsCertStoresSchedule(modelsCertStoresSchedule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.CertificateStoresSchedulePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoresSchedulePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **modelsCertStoresSchedule** | [**ModelsCertStoresSchedule**](ModelsCertStoresSchedule.md) | Certificate store identifiers and the desired schedule | 

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


## CertificateStoresServerGet

> []ModelsCertificateStoreServer CertificateStoresServerGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns all certificate store servers according to the provided filter and output parameters

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
    sortAscending := int32(56) // int32 |  (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.CertificateStoresServerGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.CertificateStoresServerGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateStoresServerGet`: []ModelsCertificateStoreServer
    fmt.Fprintf(os.Stdout, "Response from `CertificateStoreApi.CertificateStoresServerGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoresServerGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **queryString** | **string** |  | 
 **pageReturned** | **int32** |  | 
 **returnLimit** | **int32** |  | 
 **sortField** | **string** |  | 
 **sortAscending** | **int32** |  | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]ModelsCertificateStoreServer**](ModelsCertificateStoreServer.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateStoresServerPost

> ModelsCertificateStoreServerResponse CertificateStoresServerPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).ModelsCertificateStoreCreateServerRequest(modelsCertificateStoreCreateServerRequest).Execute()

Creates a new certificate store server with the provided properties

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
    modelsCertificateStoreCreateServerRequest := *openapiclient.NewModelsCertificateStoreCreateServerRequest(*openapiclient.NewModelsKeyfactorAPISecret(), *openapiclient.NewModelsKeyfactorAPISecret(), false, int32(123), "Name_example") // ModelsCertificateStoreCreateServerRequest | Certificate store server to be created with the provided properties (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.CertificateStoresServerPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).ModelsCertificateStoreCreateServerRequest(modelsCertificateStoreCreateServerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.CertificateStoresServerPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateStoresServerPost`: ModelsCertificateStoreServerResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateStoreApi.CertificateStoresServerPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoresServerPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **modelsCertificateStoreCreateServerRequest** | [**ModelsCertificateStoreCreateServerRequest**](ModelsCertificateStoreCreateServerRequest.md) | Certificate store server to be created with the provided properties | 

### Return type

[**ModelsCertificateStoreServerResponse**](ModelsCertificateStoreServerResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateStoresServerPut

> ModelsCertificateStoreServerResponse CertificateStoresServerPut(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).ModelsCertificateStoreUpdateServerRequest(modelsCertificateStoreUpdateServerRequest).Execute()

Updates a given certificate store server with the properties of the provided instance

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
    modelsCertificateStoreUpdateServerRequest := *openapiclient.NewModelsCertificateStoreUpdateServerRequest(int32(123), *openapiclient.NewModelsKeyfactorAPISecret(), *openapiclient.NewModelsKeyfactorAPISecret(), false, "Name_example") // ModelsCertificateStoreUpdateServerRequest | Server to be updated with the provided properties (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.CertificateStoresServerPut(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).ModelsCertificateStoreUpdateServerRequest(modelsCertificateStoreUpdateServerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.CertificateStoresServerPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateStoresServerPut`: ModelsCertificateStoreServerResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateStoreApi.CertificateStoresServerPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoresServerPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **modelsCertificateStoreUpdateServerRequest** | [**ModelsCertificateStoreUpdateServerRequest**](ModelsCertificateStoreUpdateServerRequest.md) | Server to be updated with the provided properties | 

### Return type

[**ModelsCertificateStoreServerResponse**](ModelsCertificateStoreServerResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


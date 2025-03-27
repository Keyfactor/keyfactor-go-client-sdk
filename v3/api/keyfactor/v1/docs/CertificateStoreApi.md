# \CertificateStoreApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCertificateStores**](CertificateStoreApi.md#CreateCertificateStores) | **POST** /CertificateStores | Creates a new certificate store with the provided properties
[**CreateCertificateStoresApprove**](CertificateStoreApi.md#CreateCertificateStoresApprove) | **POST** /CertificateStores/Approve | Approves the provided certificate stores to make them available for management
[**CreateCertificateStoresCertificatesAdd**](CertificateStoreApi.md#CreateCertificateStoresCertificatesAdd) | **POST** /CertificateStores/Certificates/Add | Configures a management job to add a certificate to one or more stores with the provided schedule
[**CreateCertificateStoresCertificatesRemove**](CertificateStoreApi.md#CreateCertificateStoresCertificatesRemove) | **POST** /CertificateStores/Certificates/Remove | Configures a management job to remove a certificate from one or more stores with the provided schedule
[**CreateCertificateStoresReenrollment**](CertificateStoreApi.md#CreateCertificateStoresReenrollment) | **POST** /CertificateStores/Reenrollment | Schedules a certificate store for reenrollment
[**CreateCertificateStoresSchedule**](CertificateStoreApi.md#CreateCertificateStoresSchedule) | **POST** /CertificateStores/Schedule | Creates an inventory schedule for the provided certificate stores
[**DeleteCertificateStores**](CertificateStoreApi.md#DeleteCertificateStores) | **DELETE** /CertificateStores | Deletes multiple persisted certificate store entities by their identifiers
[**DeleteCertificateStoresById**](CertificateStoreApi.md#DeleteCertificateStoresById) | **DELETE** /CertificateStores/{id} | Deletes a persisted certificate store by its Keyfactor identifier
[**GetCertificateStores**](CertificateStoreApi.md#GetCertificateStores) | **GET** /CertificateStores | Returns all certificate stores according to the provided filter and output parameters
[**GetCertificateStoresById**](CertificateStoreApi.md#GetCertificateStoresById) | **GET** /CertificateStores/{id} | Returns a single certificate store associated with the provided id
[**GetCertificateStoresByIdInventory**](CertificateStoreApi.md#GetCertificateStoresByIdInventory) | **GET** /CertificateStores/{id}/Inventory | Returns a single certificate store&#39;s inventory associated with the provided id
[**UpdateCertificateStores**](CertificateStoreApi.md#UpdateCertificateStores) | **PUT** /CertificateStores | Updates a given certificate store with the properties of the provided instance
[**UpdateCertificateStoresAssignContainer**](CertificateStoreApi.md#UpdateCertificateStoresAssignContainer) | **PUT** /CertificateStores/AssignContainer | Assigns the provided certificate stores to the provided container
[**UpdateCertificateStoresDiscoveryJob**](CertificateStoreApi.md#UpdateCertificateStoresDiscoveryJob) | **PUT** /CertificateStores/DiscoveryJob | Configures a discovery job to locate currently unmanaged certificate stores
[**UpdateCertificateStoresPassword**](CertificateStoreApi.md#UpdateCertificateStoresPassword) | **PUT** /CertificateStores/Password | Sets a password for the requested certificate store



## CreateCertificateStores

> CertificateStoresCertificateStoreResponse BuildCreateCertificateStoresRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest(cSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    cSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest := *openapiclient.NewCSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest() // CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest | Certificate store to be created with the provided properties (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.BuildCreateCertificateStoresRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest(cSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.CreateCertificateStores``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCertificateStores`: CertificateStoresCertificateStoreResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateStoreApi.CreateCertificateStores`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCertificateStoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest** | [**CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest**](CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest.md) | Certificate store to be created with the provided properties | 

### Return type

[**CertificateStoresCertificateStoreResponse**](CertificateStoresCertificateStoreResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCertificateStoresApprove

> BuildCreateCertificateStoresApproveRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CertificateStoresCertificateStoreApproveRequest(certificateStoresCertificateStoreApproveRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    certificateStoresCertificateStoreApproveRequest := []openapiclient.CertificateStoresCertificateStoreApproveRequest{*openapiclient.NewCertificateStoresCertificateStoreApproveRequest()} // []CertificateStoresCertificateStoreApproveRequest | Certificate stores to be approved (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.BuildCreateCertificateStoresApproveRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CertificateStoresCertificateStoreApproveRequest(certificateStoresCertificateStoreApproveRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.CreateCertificateStoresApprove``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCertificateStoresApproveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **certificateStoresCertificateStoreApproveRequest** | [**[]CertificateStoresCertificateStoreApproveRequest**](CertificateStoresCertificateStoreApproveRequest.md) | Certificate stores to be approved | 

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


## CreateCertificateStoresCertificatesAdd

> []string BuildCreateCertificateStoresCertificatesAddRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CertificateStoresAddCertificateRequest(certificateStoresAddCertificateRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    certificateStoresAddCertificateRequest := *openapiclient.NewCertificateStoresAddCertificateRequest(int32(123), []openapiclient.CSSCMSDataModelModelsCertificateStoreEntry{*openapiclient.NewCSSCMSDataModelModelsCertificateStoreEntry("CertificateStoreId_example")}, *openapiclient.NewKeyfactorCommonSchedulingKeyfactorSchedule()) // CertificateStoresAddCertificateRequest | Configuration details of the management job (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.BuildCreateCertificateStoresCertificatesAddRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CertificateStoresAddCertificateRequest(certificateStoresAddCertificateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.CreateCertificateStoresCertificatesAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCertificateStoresCertificatesAdd`: []string
    fmt.Fprintf(os.Stdout, "Response from `CertificateStoreApi.CreateCertificateStoresCertificatesAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCertificateStoresCertificatesAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **certificateStoresAddCertificateRequest** | [**CertificateStoresAddCertificateRequest**](CertificateStoresAddCertificateRequest.md) | Configuration details of the management job | 

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


## CreateCertificateStoresCertificatesRemove

> []string BuildCreateCertificateStoresCertificatesRemoveRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CertificateStoresRemoveCertificateRequest(certificateStoresRemoveCertificateRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    certificateStoresRemoveCertificateRequest := *openapiclient.NewCertificateStoresRemoveCertificateRequest([]openapiclient.CSSCMSDataModelModelsCertificateLocationSpecifier{*openapiclient.NewCSSCMSDataModelModelsCertificateLocationSpecifier()}, *openapiclient.NewKeyfactorCommonSchedulingKeyfactorSchedule()) // CertificateStoresRemoveCertificateRequest | Configuration details of the management job (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.BuildCreateCertificateStoresCertificatesRemoveRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CertificateStoresRemoveCertificateRequest(certificateStoresRemoveCertificateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.CreateCertificateStoresCertificatesRemove``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCertificateStoresCertificatesRemove`: []string
    fmt.Fprintf(os.Stdout, "Response from `CertificateStoreApi.CreateCertificateStoresCertificatesRemove`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCertificateStoresCertificatesRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **certificateStoresRemoveCertificateRequest** | [**CertificateStoresRemoveCertificateRequest**](CertificateStoresRemoveCertificateRequest.md) | Configuration details of the management job | 

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


## CreateCertificateStoresReenrollment

> BuildCreateCertificateStoresReenrollmentRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CertificateStoresReenrollmentRequest(certificateStoresReenrollmentRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    certificateStoresReenrollmentRequest := *openapiclient.NewCertificateStoresReenrollmentRequest() // CertificateStoresReenrollmentRequest | An object that contains a Keystore Id, a Agent Guid, a string SubjectName and string Alias (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.BuildCreateCertificateStoresReenrollmentRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CertificateStoresReenrollmentRequest(certificateStoresReenrollmentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.CreateCertificateStoresReenrollment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCertificateStoresReenrollmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **certificateStoresReenrollmentRequest** | [**CertificateStoresReenrollmentRequest**](CertificateStoresReenrollmentRequest.md) | An object that contains a Keystore Id, a Agent Guid, a string SubjectName and string Alias | 

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


## CreateCertificateStoresSchedule

> BuildCreateCertificateStoresScheduleRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsCertStoresSchedule(cSSCMSDataModelModelsCertStoresSchedule).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    cSSCMSDataModelModelsCertStoresSchedule := *openapiclient.NewCSSCMSDataModelModelsCertStoresSchedule([]string{"StoreIds_example"}) // CSSCMSDataModelModelsCertStoresSchedule | Certificate store identifiers and the desired schedule (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.BuildCreateCertificateStoresScheduleRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsCertStoresSchedule(cSSCMSDataModelModelsCertStoresSchedule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.CreateCertificateStoresSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCertificateStoresScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsCertStoresSchedule** | [**CSSCMSDataModelModelsCertStoresSchedule**](CSSCMSDataModelModelsCertStoresSchedule.md) | Certificate store identifiers and the desired schedule | 

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


## DeleteCertificateStores

> BuildDeleteCertificateStoresRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).RequestBody(requestBody).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    requestBody := []string{"Property_example"} // []string | Array of Keyfactor identifiers (GUID) for the certificate stores to be deleted (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.BuildDeleteCertificateStoresRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.DeleteCertificateStores``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCertificateStoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
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


## DeleteCertificateStoresById

> BuildDeleteCertificateStoresByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.BuildDeleteCertificateStoresByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.DeleteCertificateStoresById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteCertificateStoresByIdRequest struct via the builder pattern


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


## GetCertificateStores

> []CertificateStoresCertificateStoreResponse BuildGetCertificateStoresRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).PerformRoleCheck(performRoleCheck).RoleIdList(roleIdList).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    performRoleCheck := true // bool | Set to 'true' if role permissions for the current user should be validated (optional)
    roleIdList := []int32{int32(123)} // []int32 | List of Keyfactor role identifiers (integer) used to determine permissions if provided (optional)
    queryString := "queryString_example" // string | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) (optional)
    pageReturned := int32(56) // int32 | The current page within the result set to be returned (optional)
    returnLimit := int32(56) // int32 | Maximum number of records to be returned in a single call (optional)
    sortField := "sortField_example" // string | Field by which the results should be sorted (view results via Management Portal for sortable columns) (optional)
    sortAscending := openapiclient.Keyfactor.Common.QueryableExtensions.SortOrder(0) // KeyfactorCommonQueryableExtensionsSortOrder | Field sort direction [0=ascending, 1=descending] (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.BuildGetCertificateStoresRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).PerformRoleCheck(performRoleCheck).RoleIdList(roleIdList).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.GetCertificateStores``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificateStores`: []CertificateStoresCertificateStoreResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateStoreApi.GetCertificateStores`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificateStoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **performRoleCheck** | **bool** | Set to &#39;true&#39; if role permissions for the current user should be validated | 
 **roleIdList** | **[]int32** | List of Keyfactor role identifiers (integer) used to determine permissions if provided | 
 **queryString** | **string** | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) | 
 **pageReturned** | **int32** | The current page within the result set to be returned | 
 **returnLimit** | **int32** | Maximum number of records to be returned in a single call | 
 **sortField** | **string** | Field by which the results should be sorted (view results via Management Portal for sortable columns) | 
 **sortAscending** | [**KeyfactorCommonQueryableExtensionsSortOrder**](KeyfactorCommonQueryableExtensionsSortOrder.md) | Field sort direction [0&#x3D;ascending, 1&#x3D;descending] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]CertificateStoresCertificateStoreResponse**](CertificateStoresCertificateStoreResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificateStoresById

> CertificateStoresCertificateStoreResponse BuildGetCertificateStoresByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.BuildGetCertificateStoresByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.GetCertificateStoresById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificateStoresById`: CertificateStoresCertificateStoreResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateStoreApi.GetCertificateStoresById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Keyfactor identifier (GUID) of the certificate store | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificateStoresByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**CertificateStoresCertificateStoreResponse**](CertificateStoresCertificateStoreResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificateStoresByIdInventory

> []CertificateStoresCertificateStoreInventoryResponse BuildGetCertificateStoresByIdInventoryRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    pageReturned := int32(56) // int32 |  (optional)
    returnLimit := int32(56) // int32 |  (optional)
    sortField := "sortField_example" // string |  (optional)
    sortAscending := openapiclient.Keyfactor.Common.QueryableExtensions.SortOrder(0) // KeyfactorCommonQueryableExtensionsSortOrder |  (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.BuildGetCertificateStoresByIdInventoryRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.GetCertificateStoresByIdInventory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificateStoresByIdInventory`: []CertificateStoresCertificateStoreInventoryResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateStoreApi.GetCertificateStoresByIdInventory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Keyfactor identifier (GUID) of the certificate store | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificateStoresByIdInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **pageReturned** | **int32** |  | 
 **returnLimit** | **int32** |  | 
 **sortField** | **string** |  | 
 **sortAscending** | [**KeyfactorCommonQueryableExtensionsSortOrder**](KeyfactorCommonQueryableExtensionsSortOrder.md) |  | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]CertificateStoresCertificateStoreInventoryResponse**](CertificateStoresCertificateStoreInventoryResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCertificateStores

> CertificateStoresCertificateStoreResponse BuildUpdateCertificateStoresRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsCertificateStoresCertificateStoreUpdateRequest(cSSCMSDataModelModelsCertificateStoresCertificateStoreUpdateRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    cSSCMSDataModelModelsCertificateStoresCertificateStoreUpdateRequest := *openapiclient.NewCSSCMSDataModelModelsCertificateStoresCertificateStoreUpdateRequest() // CSSCMSDataModelModelsCertificateStoresCertificateStoreUpdateRequest | Certificate store to be updated with the provided properties (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.BuildUpdateCertificateStoresRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsCertificateStoresCertificateStoreUpdateRequest(cSSCMSDataModelModelsCertificateStoresCertificateStoreUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.UpdateCertificateStores``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCertificateStores`: CertificateStoresCertificateStoreResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateStoreApi.UpdateCertificateStores`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCertificateStoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsCertificateStoresCertificateStoreUpdateRequest** | [**CSSCMSDataModelModelsCertificateStoresCertificateStoreUpdateRequest**](CSSCMSDataModelModelsCertificateStoresCertificateStoreUpdateRequest.md) | Certificate store to be updated with the provided properties | 

### Return type

[**CertificateStoresCertificateStoreResponse**](CertificateStoresCertificateStoreResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCertificateStoresAssignContainer

> []CertificateStoresCertificateStoreResponse BuildUpdateCertificateStoresAssignContainerRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsContainerAssignment(cSSCMSDataModelModelsContainerAssignment).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    cSSCMSDataModelModelsContainerAssignment := *openapiclient.NewCSSCMSDataModelModelsContainerAssignment([]string{"KeystoreIds_example"}) // CSSCMSDataModelModelsContainerAssignment | Keyfactor certificate store identifiers and the container properties (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.BuildUpdateCertificateStoresAssignContainerRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsContainerAssignment(cSSCMSDataModelModelsContainerAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.UpdateCertificateStoresAssignContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCertificateStoresAssignContainer`: []CertificateStoresCertificateStoreResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateStoreApi.UpdateCertificateStoresAssignContainer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCertificateStoresAssignContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsContainerAssignment** | [**CSSCMSDataModelModelsContainerAssignment**](CSSCMSDataModelModelsContainerAssignment.md) | Keyfactor certificate store identifiers and the container properties | 

### Return type

[**[]CertificateStoresCertificateStoreResponse**](CertificateStoresCertificateStoreResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCertificateStoresDiscoveryJob

> BuildUpdateCertificateStoresDiscoveryJobRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsDiscoveryJobRequest(cSSCMSDataModelModelsDiscoveryJobRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    cSSCMSDataModelModelsDiscoveryJobRequest := *openapiclient.NewCSSCMSDataModelModelsDiscoveryJobRequest(int32(123)) // CSSCMSDataModelModelsDiscoveryJobRequest | Configuration properties of the discovery job (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.BuildUpdateCertificateStoresDiscoveryJobRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsDiscoveryJobRequest(cSSCMSDataModelModelsDiscoveryJobRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.UpdateCertificateStoresDiscoveryJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCertificateStoresDiscoveryJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsDiscoveryJobRequest** | [**CSSCMSDataModelModelsDiscoveryJobRequest**](CSSCMSDataModelModelsDiscoveryJobRequest.md) | Configuration properties of the discovery job | 

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


## UpdateCertificateStoresPassword

> BuildUpdateCertificateStoresPasswordRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsCertStoreNewPasswordRequest(cSSCMSDataModelModelsCertStoreNewPasswordRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    cSSCMSDataModelModelsCertStoreNewPasswordRequest := *openapiclient.NewCSSCMSDataModelModelsCertStoreNewPasswordRequest("CertStoreId_example", interface{}(123)) // CSSCMSDataModelModelsCertStoreNewPasswordRequest | Identifier of the certificate store and the password to be applied to it (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.BuildUpdateCertificateStoresPasswordRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsCertStoreNewPasswordRequest(cSSCMSDataModelModelsCertStoreNewPasswordRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.UpdateCertificateStoresPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCertificateStoresPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsCertStoreNewPasswordRequest** | [**CSSCMSDataModelModelsCertStoreNewPasswordRequest**](CSSCMSDataModelModelsCertStoreNewPasswordRequest.md) | Identifier of the certificate store and the password to be applied to it | 

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


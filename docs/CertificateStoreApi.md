# \CertificateStoreApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CertificateStoreAddCertificate**](CertificateStoreApi.md#CertificateStoreAddCertificate) | **Post** /CertificateStores/Certificates/Add | Configures a management job to add a certificate to one or more stores with the provided schedule
[**CertificateStoreApprovePending**](CertificateStoreApi.md#CertificateStoreApprovePending) | **Post** /CertificateStores/Approve | Approves the provided certificate stores to make them available for management
[**CertificateStoreConfigureDiscoveryJob**](CertificateStoreApi.md#CertificateStoreConfigureDiscoveryJob) | **Put** /CertificateStores/DiscoveryJob | Configures a discovery job to locate currently unmanaged certificate stores
[**CertificateStoreCreateCertificateStore**](CertificateStoreApi.md#CertificateStoreCreateCertificateStore) | **Post** /CertificateStores | Creates a new certificate store with the provided properties
[**CertificateStoreCreateCertificateStoreServer**](CertificateStoreApi.md#CertificateStoreCreateCertificateStoreServer) | **Post** /CertificateStores/Server | Creates a new certificate store server with the provided properties
[**CertificateStoreDeleteCertificateStore**](CertificateStoreApi.md#CertificateStoreDeleteCertificateStore) | **Delete** /CertificateStores/{id} | Deletes a persisted certificate store by its Keyfactor identifier
[**CertificateStoreDeleteCertificateStores**](CertificateStoreApi.md#CertificateStoreDeleteCertificateStores) | **Delete** /CertificateStores | Deletes multiple persisted certificate store entities by their identifiers
[**CertificateStoreGetCertificateStore**](CertificateStoreApi.md#CertificateStoreGetCertificateStore) | **Get** /CertificateStores/{id} | Returns a single certificate store associated with the provided id
[**CertificateStoreGetCertificateStoreInventory**](CertificateStoreApi.md#CertificateStoreGetCertificateStoreInventory) | **Get** /CertificateStores/{id}/Inventory | Returns a single certificate store&#39;s inventory associated with the provided id
[**CertificateStoreQueryCertificateStores**](CertificateStoreApi.md#CertificateStoreQueryCertificateStores) | **Get** /CertificateStores | Returns all certificate stores according to the provided filter and output parameters
[**CertificateStoreRemoveCertificate**](CertificateStoreApi.md#CertificateStoreRemoveCertificate) | **Post** /CertificateStores/Certificates/Remove | Configures a management job to remove a certificate from one or more stores with the provided schedule
[**CertificateStoreSchedule**](CertificateStoreApi.md#CertificateStoreSchedule) | **Post** /CertificateStores/Schedule | Creates an inventory schedule for the provided certificate stores
[**CertificateStoreScheduleForReenrollment**](CertificateStoreApi.md#CertificateStoreScheduleForReenrollment) | **Post** /CertificateStores/Reenrollment | Schedules a certificate store for reenrollment
[**CertificateStoreSetPassword**](CertificateStoreApi.md#CertificateStoreSetPassword) | **Put** /CertificateStores/Password | Sets a password for the requested certificate store
[**CertificateStoreUpdateCertStore**](CertificateStoreApi.md#CertificateStoreUpdateCertStore) | **Put** /CertificateStores | Updates a given certificate store with the properties of the provided instance
[**CertificateStoreUpdateCertificateStoreServer**](CertificateStoreApi.md#CertificateStoreUpdateCertificateStoreServer) | **Put** /CertificateStores/Server | Updates a given certificate store server with the properties of the provided instance



## CertificateStoreAddCertificate

> []string CertificateStoreAddCertificate(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).AddRequest(addRequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    addRequest := *openapiclient.NewKeyfactorApiModelsCertificateStoresAddCertificateRequest(int32(123), []openapiclient.ModelsCertificateStoreEntry{*openapiclient.NewModelsCertificateStoreEntry("00000000-0000-0000-0000-000000000000")}, *openapiclient.NewKeyfactorCommonSchedulingKeyfactorSchedule()) // KeyfactorApiModelsCertificateStoresAddCertificateRequest | Configuration details of the management job
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.CertificateStoreAddCertificate(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).AddRequest(addRequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.CertificateStoreAddCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateStoreAddCertificate`: []string
    fmt.Fprintf(os.Stdout, "Response from `CertificateStoreApi.CertificateStoreAddCertificate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoreAddCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **addRequest** | [**KeyfactorApiModelsCertificateStoresAddCertificateRequest**](KeyfactorApiModelsCertificateStoresAddCertificateRequest.md) | Configuration details of the management job | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

**[]string**

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateStoreApprovePending

> CertificateStoreApprovePending(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Keystores(keystores).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    keystores := []openapiclient.KeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest{*openapiclient.NewKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest()} // []KeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest | Certificate stores to be approved
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.CertificateStoreApprovePending(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Keystores(keystores).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.CertificateStoreApprovePending``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoreApprovePendingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **keystores** | [**[]KeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest**](KeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest.md) | Certificate stores to be approved | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateStoreConfigureDiscoveryJob

> CertificateStoreConfigureDiscoveryJob(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).DiscoveryJobRequest(discoveryJobRequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    discoveryJobRequest := *openapiclient.NewModelsDiscoveryJobRequest(int32(123)) // ModelsDiscoveryJobRequest | Configuration properties of the discovery job
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.CertificateStoreConfigureDiscoveryJob(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).DiscoveryJobRequest(discoveryJobRequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.CertificateStoreConfigureDiscoveryJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoreConfigureDiscoveryJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **discoveryJobRequest** | [**ModelsDiscoveryJobRequest**](ModelsDiscoveryJobRequest.md) | Configuration properties of the discovery job | 
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


## CertificateStoreCreateCertificateStore

> KeyfactorApiModelsCertificateStoresCertificateStoreResponse CertificateStoreCreateCertificateStore(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Store(store).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    store := *openapiclient.NewModelsCertificateStoresCertificateStoreCreateRequest() // ModelsCertificateStoresCertificateStoreCreateRequest | Certificate store to be created with the provided properties
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.CertificateStoreCreateCertificateStore(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Store(store).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.CertificateStoreCreateCertificateStore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateStoreCreateCertificateStore`: KeyfactorApiModelsCertificateStoresCertificateStoreResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateStoreApi.CertificateStoreCreateCertificateStore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoreCreateCertificateStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **store** | [**ModelsCertificateStoresCertificateStoreCreateRequest**](ModelsCertificateStoresCertificateStoreCreateRequest.md) | Certificate store to be created with the provided properties | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsCertificateStoresCertificateStoreResponse**](KeyfactorApiModelsCertificateStoresCertificateStoreResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateStoreCreateCertificateStoreServer

> ModelsCertificateStoreServerResponse CertificateStoreCreateCertificateStoreServer(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Request(request).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    request := *openapiclient.NewModelsCertificateStoreCreateServerRequest(*openapiclient.NewModelsKeyfactorAPISecret(), *openapiclient.NewModelsKeyfactorAPISecret(), false, int32(123), "Name_example") // ModelsCertificateStoreCreateServerRequest | Certificate store server to be created with the provided properties
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.CertificateStoreCreateCertificateStoreServer(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Request(request).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.CertificateStoreCreateCertificateStoreServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateStoreCreateCertificateStoreServer`: ModelsCertificateStoreServerResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateStoreApi.CertificateStoreCreateCertificateStoreServer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoreCreateCertificateStoreServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **request** | [**ModelsCertificateStoreCreateServerRequest**](ModelsCertificateStoreCreateServerRequest.md) | Certificate store server to be created with the provided properties | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsCertificateStoreServerResponse**](ModelsCertificateStoreServerResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateStoreDeleteCertificateStore

> CertificateStoreDeleteCertificateStore(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.CertificateStoreDeleteCertificateStore(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.CertificateStoreDeleteCertificateStore``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCertificateStoreDeleteCertificateStoreRequest struct via the builder pattern


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


## CertificateStoreDeleteCertificateStores

> CertificateStoreDeleteCertificateStores(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Ids(ids).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    ids := []string{"00000000-0000-0000-0000-000000000000"} // []string | Array of Keyfactor identifiers (GUID) for the certificate stores to be deleted
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.CertificateStoreDeleteCertificateStores(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Ids(ids).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.CertificateStoreDeleteCertificateStores``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoreDeleteCertificateStoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **ids** | **[]string** | Array of Keyfactor identifiers (GUID) for the certificate stores to be deleted | 
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


## CertificateStoreGetCertificateStore

> KeyfactorApiModelsCertificateStoresCertificateStoreResponse CertificateStoreGetCertificateStore(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.CertificateStoreGetCertificateStore(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.CertificateStoreGetCertificateStore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateStoreGetCertificateStore`: KeyfactorApiModelsCertificateStoresCertificateStoreResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateStoreApi.CertificateStoreGetCertificateStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Keyfactor identifier (GUID) of the certificate store | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoreGetCertificateStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsCertificateStoresCertificateStoreResponse**](KeyfactorApiModelsCertificateStoresCertificateStoreResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateStoreGetCertificateStoreInventory

> []ModelsCertificateStoreInventory CertificateStoreGetCertificateStoreInventory(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).QueryPageReturned(queryPageReturned).QueryReturnLimit(queryReturnLimit).QuerySortField(querySortField).QuerySortAscending(querySortAscending).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")
    queryPageReturned := int32(56) // int32 | The current page within the result set to be returned (optional)
    queryReturnLimit := int32(56) // int32 | Maximum number of records to be returned in a single call (optional)
    querySortField := "querySortField_example" // string | Field by which the results should be sorted (OperationStart, OperationEnd, UserName) (optional)
    querySortAscending := int32(56) // int32 | Field sort direction [0=ascending, 1=descending] (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.CertificateStoreGetCertificateStoreInventory(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).QueryPageReturned(queryPageReturned).QueryReturnLimit(queryReturnLimit).QuerySortField(querySortField).QuerySortAscending(querySortAscending).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.CertificateStoreGetCertificateStoreInventory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateStoreGetCertificateStoreInventory`: []ModelsCertificateStoreInventory
    fmt.Fprintf(os.Stdout, "Response from `CertificateStoreApi.CertificateStoreGetCertificateStoreInventory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Keyfactor identifier (GUID) of the certificate store | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoreGetCertificateStoreInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]
 **queryPageReturned** | **int32** | The current page within the result set to be returned | 
 **queryReturnLimit** | **int32** | Maximum number of records to be returned in a single call | 
 **querySortField** | **string** | Field by which the results should be sorted (OperationStart, OperationEnd, UserName) | 
 **querySortAscending** | **int32** | Field sort direction [0&#x3D;ascending, 1&#x3D;descending] | 

### Return type

[**[]ModelsCertificateStoreInventory**](ModelsCertificateStoreInventory.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateStoreQueryCertificateStores

> []KeyfactorApiModelsCertificateStoresCertificateStoreResponse CertificateStoreQueryCertificateStores(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CertificateStoreQueryQueryString(certificateStoreQueryQueryString).CertificateStoreQueryPageReturned(certificateStoreQueryPageReturned).CertificateStoreQueryReturnLimit(certificateStoreQueryReturnLimit).CertificateStoreQuerySortField(certificateStoreQuerySortField).CertificateStoreQuerySortAscending(certificateStoreQuerySortAscending).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")
    certificateStoreQueryQueryString := "certificateStoreQueryQueryString_example" // string | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) (optional)
    certificateStoreQueryPageReturned := int32(56) // int32 | The current page within the result set to be returned (optional)
    certificateStoreQueryReturnLimit := int32(56) // int32 | Maximum number of records to be returned in a single call (optional)
    certificateStoreQuerySortField := "certificateStoreQuerySortField_example" // string | Field by which the results should be sorted (view results via Management Portal for sortable columns) (optional)
    certificateStoreQuerySortAscending := int32(56) // int32 | Field sort direction [0=ascending, 1=descending] (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.CertificateStoreQueryCertificateStores(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CertificateStoreQueryQueryString(certificateStoreQueryQueryString).CertificateStoreQueryPageReturned(certificateStoreQueryPageReturned).CertificateStoreQueryReturnLimit(certificateStoreQueryReturnLimit).CertificateStoreQuerySortField(certificateStoreQuerySortField).CertificateStoreQuerySortAscending(certificateStoreQuerySortAscending).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.CertificateStoreQueryCertificateStores``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateStoreQueryCertificateStores`: []KeyfactorApiModelsCertificateStoresCertificateStoreResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateStoreApi.CertificateStoreQueryCertificateStores`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoreQueryCertificateStoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]
 **certificateStoreQueryQueryString** | **string** | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) | 
 **certificateStoreQueryPageReturned** | **int32** | The current page within the result set to be returned | 
 **certificateStoreQueryReturnLimit** | **int32** | Maximum number of records to be returned in a single call | 
 **certificateStoreQuerySortField** | **string** | Field by which the results should be sorted (view results via Management Portal for sortable columns) | 
 **certificateStoreQuerySortAscending** | **int32** | Field sort direction [0&#x3D;ascending, 1&#x3D;descending] | 

### Return type

[**[]KeyfactorApiModelsCertificateStoresCertificateStoreResponse**](KeyfactorApiModelsCertificateStoresCertificateStoreResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateStoreRemoveCertificate

> []string CertificateStoreRemoveCertificate(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).RemovalRequest(removalRequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    removalRequest := *openapiclient.NewKeyfactorApiModelsCertificateStoresRemoveCertificateRequest([]openapiclient.ModelsCertificateLocationSpecifier{*openapiclient.NewModelsCertificateLocationSpecifier()}, *openapiclient.NewKeyfactorCommonSchedulingKeyfactorSchedule()) // KeyfactorApiModelsCertificateStoresRemoveCertificateRequest | Configuration details of the management job
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.CertificateStoreRemoveCertificate(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).RemovalRequest(removalRequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.CertificateStoreRemoveCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateStoreRemoveCertificate`: []string
    fmt.Fprintf(os.Stdout, "Response from `CertificateStoreApi.CertificateStoreRemoveCertificate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoreRemoveCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **removalRequest** | [**KeyfactorApiModelsCertificateStoresRemoveCertificateRequest**](KeyfactorApiModelsCertificateStoresRemoveCertificateRequest.md) | Configuration details of the management job | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

**[]string**

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateStoreSchedule

> CertificateStoreSchedule(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).FutureSchedule(futureSchedule).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    futureSchedule := *openapiclient.NewModelsCertStoresSchedule([]string{"00000000-0000-0000-0000-000000000000"}) // ModelsCertStoresSchedule | Certificate store identifiers and the desired schedule
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.CertificateStoreSchedule(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).FutureSchedule(futureSchedule).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.CertificateStoreSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoreScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **futureSchedule** | [**ModelsCertStoresSchedule**](ModelsCertStoresSchedule.md) | Certificate store identifiers and the desired schedule | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateStoreScheduleForReenrollment

> CertificateStoreScheduleForReenrollment(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Reenroll(reenroll).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    reenroll := *openapiclient.NewKeyfactorApiModelsCertificateStoresReenrollmentRequest() // KeyfactorApiModelsCertificateStoresReenrollmentRequest | An object that contains a Keystore Id, a Agent Guid, a string SubjectName and string Alias
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.CertificateStoreScheduleForReenrollment(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Reenroll(reenroll).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.CertificateStoreScheduleForReenrollment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoreScheduleForReenrollmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **reenroll** | [**KeyfactorApiModelsCertificateStoresReenrollmentRequest**](KeyfactorApiModelsCertificateStoresReenrollmentRequest.md) | An object that contains a Keystore Id, a Agent Guid, a string SubjectName and string Alias | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateStoreSetPassword

> CertificateStoreSetPassword(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).PasswordRequest(passwordRequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    passwordRequest := *openapiclient.NewModelsCertStoreNewPasswordRequest("00000000-0000-0000-0000-000000000000", map[string]interface{}(123)) // ModelsCertStoreNewPasswordRequest | Identifier of the certificate store and the password to be applied to it
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.CertificateStoreSetPassword(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).PasswordRequest(passwordRequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.CertificateStoreSetPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoreSetPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **passwordRequest** | [**ModelsCertStoreNewPasswordRequest**](ModelsCertStoreNewPasswordRequest.md) | Identifier of the certificate store and the password to be applied to it | 
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


## CertificateStoreUpdateCertStore

> KeyfactorApiModelsCertificateStoresCertificateStoreResponse CertificateStoreUpdateCertStore(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CertStore(certStore).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    certStore := *openapiclient.NewModelsCertificateStoresCertificateStoreUpdateRequest() // ModelsCertificateStoresCertificateStoreUpdateRequest | Certificate store to be updated with the provided properties
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.CertificateStoreUpdateCertStore(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CertStore(certStore).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.CertificateStoreUpdateCertStore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateStoreUpdateCertStore`: KeyfactorApiModelsCertificateStoresCertificateStoreResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateStoreApi.CertificateStoreUpdateCertStore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoreUpdateCertStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **certStore** | [**ModelsCertificateStoresCertificateStoreUpdateRequest**](ModelsCertificateStoresCertificateStoreUpdateRequest.md) | Certificate store to be updated with the provided properties | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsCertificateStoresCertificateStoreResponse**](KeyfactorApiModelsCertificateStoresCertificateStoreResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateStoreUpdateCertificateStoreServer

> ModelsCertificateStoreServerResponse CertificateStoreUpdateCertificateStoreServer(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Request(request).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    request := *openapiclient.NewModelsCertificateStoreUpdateServerRequest(int32(123), *openapiclient.NewModelsKeyfactorAPISecret(), *openapiclient.NewModelsKeyfactorAPISecret(), false, "Name_example") // ModelsCertificateStoreUpdateServerRequest | Server to be updated with the provided properties
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreApi.CertificateStoreUpdateCertificateStoreServer(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Request(request).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreApi.CertificateStoreUpdateCertificateStoreServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateStoreUpdateCertificateStoreServer`: ModelsCertificateStoreServerResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateStoreApi.CertificateStoreUpdateCertificateStoreServer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoreUpdateCertificateStoreServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **request** | [**ModelsCertificateStoreUpdateServerRequest**](ModelsCertificateStoreUpdateServerRequest.md) | Server to be updated with the provided properties | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsCertificateStoreServerResponse**](ModelsCertificateStoreServerResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


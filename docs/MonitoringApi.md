# \MonitoringApi

All URIs are relative to *https://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MonitoringAddRevocationMonitoring**](MonitoringApi.md#MonitoringAddRevocationMonitoring) | **Post** /Monitoring/Revocation | Add a revocation monitoring endpoint
[**MonitoringDeleteRevocationMonitoring**](MonitoringApi.md#MonitoringDeleteRevocationMonitoring) | **Delete** /Monitoring/Revocation/{id} | Delete a revocation monitoring endpoint
[**MonitoringEditRevocationMonitoring**](MonitoringApi.md#MonitoringEditRevocationMonitoring) | **Put** /Monitoring/Revocation | Edit a revocation monitoring endpoint
[**MonitoringGetRevocationMonitoring**](MonitoringApi.md#MonitoringGetRevocationMonitoring) | **Get** /Monitoring/Revocation/{id} | Get a revocation monitoring endpoint
[**MonitoringGetRevocationMonitoringEndpoints**](MonitoringApi.md#MonitoringGetRevocationMonitoringEndpoints) | **Get** /Monitoring/Revocation | Gets all revocation monitoring endpoints according to the provided filter and output parameters
[**MonitoringResolveOCSP**](MonitoringApi.md#MonitoringResolveOCSP) | **Post** /Monitoring/ResolveOCSP | Resolve the Certificate authority given
[**MonitoringTestAllRevocationMonitoringAlert**](MonitoringApi.md#MonitoringTestAllRevocationMonitoringAlert) | **Post** /Monitoring/Revocation/TestAll | Test All Alerts
[**MonitoringTestRevocationMonitoringAlert**](MonitoringApi.md#MonitoringTestRevocationMonitoringAlert) | **Post** /Monitoring/Revocation/Test | Test Alert



## MonitoringAddRevocationMonitoring

> KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse MonitoringAddRevocationMonitoring(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Endpoint(endpoint).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Add a revocation monitoring endpoint

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
    endpoint := *openapiclient.NewKeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest("Name_example", "EndpointType_example", "Location_example", *openapiclient.NewKeyfactorApiModelsMonitoringDashboardRequest(false)) // KeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest | Information for the new endpoint
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringApi.MonitoringAddRevocationMonitoring(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Endpoint(endpoint).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringApi.MonitoringAddRevocationMonitoring``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitoringAddRevocationMonitoring`: KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitoringApi.MonitoringAddRevocationMonitoring`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMonitoringAddRevocationMonitoringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **endpoint** | [**KeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest**](KeyfactorApiModelsMonitoringRevocationMonitoringCreationRequest.md) | Information for the new endpoint | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse**](KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitoringDeleteRevocationMonitoring

> MonitoringDeleteRevocationMonitoring(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Delete a revocation monitoring endpoint

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
    id := int32(56) // int32 | Id for the revocation monitoring endpoint
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringApi.MonitoringDeleteRevocationMonitoring(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringApi.MonitoringDeleteRevocationMonitoring``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id for the revocation monitoring endpoint | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitoringDeleteRevocationMonitoringRequest struct via the builder pattern


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


## MonitoringEditRevocationMonitoring

> KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse MonitoringEditRevocationMonitoring(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Endpoint(endpoint).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Edit a revocation monitoring endpoint

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
    endpoint := *openapiclient.NewKeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest("Name_example", "EndpointType_example", "Location_example", *openapiclient.NewKeyfactorApiModelsMonitoringDashboardRequest(false)) // KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest | Information for the endpoint
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringApi.MonitoringEditRevocationMonitoring(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Endpoint(endpoint).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringApi.MonitoringEditRevocationMonitoring``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitoringEditRevocationMonitoring`: KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitoringApi.MonitoringEditRevocationMonitoring`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMonitoringEditRevocationMonitoringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **endpoint** | [**KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest**](KeyfactorApiModelsMonitoringRevocationMonitoringUpdateRequest.md) | Information for the endpoint | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse**](KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitoringGetRevocationMonitoring

> KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse MonitoringGetRevocationMonitoring(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Get a revocation monitoring endpoint

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
    id := int32(56) // int32 | Id for the endpoint to get
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringApi.MonitoringGetRevocationMonitoring(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringApi.MonitoringGetRevocationMonitoring``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitoringGetRevocationMonitoring`: KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitoringApi.MonitoringGetRevocationMonitoring`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id for the endpoint to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitoringGetRevocationMonitoringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse**](KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitoringGetRevocationMonitoringEndpoints

> []KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse MonitoringGetRevocationMonitoringEndpoints(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PagedQueryQueryString(pagedQueryQueryString).PagedQueryPageReturned(pagedQueryPageReturned).PagedQueryReturnLimit(pagedQueryReturnLimit).PagedQuerySortField(pagedQuerySortField).PagedQuerySortAscending(pagedQuerySortAscending).Execute()

Gets all revocation monitoring endpoints according to the provided filter and output parameters

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
    resp, r, err := apiClient.MonitoringApi.MonitoringGetRevocationMonitoringEndpoints(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PagedQueryQueryString(pagedQueryQueryString).PagedQueryPageReturned(pagedQueryPageReturned).PagedQueryReturnLimit(pagedQueryReturnLimit).PagedQuerySortField(pagedQuerySortField).PagedQuerySortAscending(pagedQuerySortAscending).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringApi.MonitoringGetRevocationMonitoringEndpoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitoringGetRevocationMonitoringEndpoints`: []KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitoringApi.MonitoringGetRevocationMonitoringEndpoints`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMonitoringGetRevocationMonitoringEndpointsRequest struct via the builder pattern


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

[**[]KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse**](KeyfactorApiModelsMonitoringRevocationMonitoringDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitoringResolveOCSP

> KeyfactorApiModelsMonitoringOCSPParametersResponse MonitoringResolveOCSP(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Endpoint(endpoint).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Resolve the Certificate authority given

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
    endpoint := *openapiclient.NewKeyfactorApiModelsMonitoringOCSPParametersRequest() // KeyfactorApiModelsMonitoringOCSPParametersRequest | Information for the new endpoint
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringApi.MonitoringResolveOCSP(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Endpoint(endpoint).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringApi.MonitoringResolveOCSP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitoringResolveOCSP`: KeyfactorApiModelsMonitoringOCSPParametersResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitoringApi.MonitoringResolveOCSP`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMonitoringResolveOCSPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **endpoint** | [**KeyfactorApiModelsMonitoringOCSPParametersRequest**](KeyfactorApiModelsMonitoringOCSPParametersRequest.md) | Information for the new endpoint | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsMonitoringOCSPParametersResponse**](KeyfactorApiModelsMonitoringOCSPParametersResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitoringTestAllRevocationMonitoringAlert

> KeyfactorApiModelsMonitoringRevocationMonitoringAlertTestResponse MonitoringTestAllRevocationMonitoringAlert(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).RevocationMonitoringAlertTestRequest(revocationMonitoringAlertTestRequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Test All Alerts

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
    revocationMonitoringAlertTestRequest := *openapiclient.NewKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestAllRequest() // KeyfactorApiModelsMonitoringRevocationMonitoringAlertTestAllRequest | Information about the revocation monitoring alert test
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringApi.MonitoringTestAllRevocationMonitoringAlert(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).RevocationMonitoringAlertTestRequest(revocationMonitoringAlertTestRequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringApi.MonitoringTestAllRevocationMonitoringAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitoringTestAllRevocationMonitoringAlert`: KeyfactorApiModelsMonitoringRevocationMonitoringAlertTestResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitoringApi.MonitoringTestAllRevocationMonitoringAlert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMonitoringTestAllRevocationMonitoringAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **revocationMonitoringAlertTestRequest** | [**KeyfactorApiModelsMonitoringRevocationMonitoringAlertTestAllRequest**](KeyfactorApiModelsMonitoringRevocationMonitoringAlertTestAllRequest.md) | Information about the revocation monitoring alert test | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsMonitoringRevocationMonitoringAlertTestResponse**](KeyfactorApiModelsMonitoringRevocationMonitoringAlertTestResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitoringTestRevocationMonitoringAlert

> KeyfactorApiModelsMonitoringRevocationMonitoringAlertTestResponse MonitoringTestRevocationMonitoringAlert(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).RevocationMonitoringAlertTestRequest(revocationMonitoringAlertTestRequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Test Alert

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
    revocationMonitoringAlertTestRequest := *openapiclient.NewKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest() // KeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest | Information about the revocation monitoring alert test
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringApi.MonitoringTestRevocationMonitoringAlert(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).RevocationMonitoringAlertTestRequest(revocationMonitoringAlertTestRequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringApi.MonitoringTestRevocationMonitoringAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitoringTestRevocationMonitoringAlert`: KeyfactorApiModelsMonitoringRevocationMonitoringAlertTestResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitoringApi.MonitoringTestRevocationMonitoringAlert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMonitoringTestRevocationMonitoringAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **revocationMonitoringAlertTestRequest** | [**KeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest**](KeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest.md) | Information about the revocation monitoring alert test | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsMonitoringRevocationMonitoringAlertTestResponse**](KeyfactorApiModelsMonitoringRevocationMonitoringAlertTestResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


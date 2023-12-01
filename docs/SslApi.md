# \SslApi

All URIs are relative to *https://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SslAddNetworkRanges**](SslApi.md#SslAddNetworkRanges) | **Post** /SSL/NetworkRanges | Adds the provided network range definitions to the associated network definition
[**SslCreateNetwork**](SslApi.md#SslCreateNetwork) | **Post** /SSL/Networks | Creates a network definition according to the provided properties
[**SslEndpoint**](SslApi.md#SslEndpoint) | **Get** /SSL/Endpoints/{id} | Returns the details of the associated scanning endpoint
[**SslEndpointHistory**](SslApi.md#SslEndpointHistory) | **Get** /SSL/Endpoints/{id}/History | Returns a list of the scan results for the provided endpoint according to the provided filter and output parameters
[**SslGetNetwork**](SslApi.md#SslGetNetwork) | **Get** /SSL/Networks/{identifier} | Returns a defined SSL network according to the provided name
[**SslGetNetworkRangesForNetwork**](SslApi.md#SslGetNetworkRangesForNetwork) | **Get** /SSL/NetworkRanges/{id} | Returns the network range definitions for the provided network definition
[**SslGetNetworks**](SslApi.md#SslGetNetworks) | **Get** /SSL/Networks | Returns all defined SSL networks according to the provided filter and output parameters
[**SslImmediateSslScan**](SslApi.md#SslImmediateSslScan) | **Post** /SSL/Networks/{id}/Scan | Starts an SSL Scan for the network according to the associated network definition
[**SslMonitorAll**](SslApi.md#SslMonitorAll) | **Put** /SSL/Endpoints/MonitorAll | Sets all endpoints matching the provided query as &#39;monitored&#39;
[**SslMonitoringStatus**](SslApi.md#SslMonitoringStatus) | **Put** /SSL/Endpoints/MonitorStatus | Sets the monitored status according to the provided endpoint and boolean status
[**SslNetworkScanParts**](SslApi.md#SslNetworkScanParts) | **Get** /SSL/Networks/{id}/Parts | Returns the scan job components comprising the entire scan job to be executed
[**SslRemoveAllNetworkRanges**](SslApi.md#SslRemoveAllNetworkRanges) | **Delete** /SSL/NetworkRanges/{id} | Removes all network range definitions from the associated network definition
[**SslRemoveNetwork**](SslApi.md#SslRemoveNetwork) | **Delete** /SSL/Networks/{id} | Removes a network definition according to the provided identifier
[**SslResults**](SslApi.md#SslResults) | **Get** /SSL | Returns a list of the endpoint scan results according to the provided filter and output parameters
[**SslReviewAll**](SslApi.md#SslReviewAll) | **Put** /SSL/Endpoints/ReviewAll | Sets all endpoints matching the provided query as &#39;reviewed&#39;
[**SslReviewedStatus**](SslApi.md#SslReviewedStatus) | **Put** /SSL/Endpoints/ReviewStatus | Sets the reviewed status according to the provided endpoint and boolean status
[**SslScanPart**](SslApi.md#SslScanPart) | **Get** /SSL/Parts/{id} | Returns the execution details of the associated network scan job part
[**SslSetNetworkRanges**](SslApi.md#SslSetNetworkRanges) | **Put** /SSL/NetworkRanges | Configures network range definitions for the provided network
[**SslUpdateNetwork**](SslApi.md#SslUpdateNetwork) | **Put** /SSL/Networks | Updates an existing network definition according to the provided properties
[**SslValidateNetworkRanges**](SslApi.md#SslValidateNetworkRanges) | **Post** /SSL/NetworkRanges/Validate | Validates the format (using regular expressions) of the provided network range definitions



## SslAddNetworkRanges

> SslAddNetworkRanges(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).NetworkRanges(networkRanges).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Adds the provided network range definitions to the associated network definition

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
    networkRanges := *openapiclient.NewModelsSSLNetworkRangesRequest("00000000-0000-0000-0000-000000000000", []string{"Ranges_example"}) // ModelsSSLNetworkRangesRequest | Network definition identifier and the ranges to be added
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SslApi.SslAddNetworkRanges(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).NetworkRanges(networkRanges).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SslApi.SslAddNetworkRanges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSslAddNetworkRangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **networkRanges** | [**ModelsSSLNetworkRangesRequest**](ModelsSSLNetworkRangesRequest.md) | Network definition identifier and the ranges to be added | 
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


## SslCreateNetwork

> KeyfactorApiModelsSslNetworkResponse SslCreateNetwork(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Network(network).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Creates a network definition according to the provided properties

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
    network := *openapiclient.NewKeyfactorApiModelsSslCreateNetworkRequest("Name_example", "AgentPoolName_example", "Description_example") // KeyfactorApiModelsSslCreateNetworkRequest | Properties of the network definition to be created
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SslApi.SslCreateNetwork(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Network(network).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SslApi.SslCreateNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SslCreateNetwork`: KeyfactorApiModelsSslNetworkResponse
    fmt.Fprintf(os.Stdout, "Response from `SslApi.SslCreateNetwork`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSslCreateNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **network** | [**KeyfactorApiModelsSslCreateNetworkRequest**](KeyfactorApiModelsSslCreateNetworkRequest.md) | Properties of the network definition to be created | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsSslNetworkResponse**](KeyfactorApiModelsSslNetworkResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SslEndpoint

> ModelsSSLEndpoint SslEndpoint(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns the details of the associated scanning endpoint

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Keyfactor identifier of the endpoint
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SslApi.SslEndpoint(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SslApi.SslEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SslEndpoint`: ModelsSSLEndpoint
    fmt.Fprintf(os.Stdout, "Response from `SslApi.SslEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Keyfactor identifier of the endpoint | 

### Other Parameters

Other parameters are passed through a pointer to a apiSslEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsSSLEndpoint**](ModelsSSLEndpoint.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SslEndpointHistory

> []ModelsSSLEndpointHistoryResponse SslEndpointHistory(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PqQueryString(pqQueryString).PqPageReturned(pqPageReturned).PqReturnLimit(pqReturnLimit).PqSortField(pqSortField).PqSortAscending(pqSortAscending).Execute()

Returns a list of the scan results for the provided endpoint according to the provided filter and output parameters

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Keyfactor identifier of the endpoint
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")
    pqQueryString := "pqQueryString_example" // string | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) (optional)
    pqPageReturned := int32(56) // int32 | The current page within the result set to be returned (optional)
    pqReturnLimit := int32(56) // int32 | Maximum number of records to be returned in a single call (optional)
    pqSortField := "pqSortField_example" // string | Field by which the results should be sorted (view results via Management Portal for sortable columns) (optional)
    pqSortAscending := int32(56) // int32 | Field sort direction [0=ascending, 1=descending] (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SslApi.SslEndpointHistory(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PqQueryString(pqQueryString).PqPageReturned(pqPageReturned).PqReturnLimit(pqReturnLimit).PqSortField(pqSortField).PqSortAscending(pqSortAscending).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SslApi.SslEndpointHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SslEndpointHistory`: []ModelsSSLEndpointHistoryResponse
    fmt.Fprintf(os.Stdout, "Response from `SslApi.SslEndpointHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Keyfactor identifier of the endpoint | 

### Other Parameters

Other parameters are passed through a pointer to a apiSslEndpointHistoryRequest struct via the builder pattern


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

[**[]ModelsSSLEndpointHistoryResponse**](ModelsSSLEndpointHistoryResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SslGetNetwork

> KeyfactorApiModelsSslNetworkResponse SslGetNetwork(ctx, identifier).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns a defined SSL network according to the provided name

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
    identifier := "identifier_example" // string | Identifier (Guid or Name) of the defined network
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SslApi.SslGetNetwork(context.Background(), identifier).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SslApi.SslGetNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SslGetNetwork`: KeyfactorApiModelsSslNetworkResponse
    fmt.Fprintf(os.Stdout, "Response from `SslApi.SslGetNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** | Identifier (Guid or Name) of the defined network | 

### Other Parameters

Other parameters are passed through a pointer to a apiSslGetNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsSslNetworkResponse**](KeyfactorApiModelsSslNetworkResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SslGetNetworkRangesForNetwork

> []ModelsSSLNetworkDefinition SslGetNetworkRangesForNetwork(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns the network range definitions for the provided network definition

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Keyfactor network identifier
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SslApi.SslGetNetworkRangesForNetwork(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SslApi.SslGetNetworkRangesForNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SslGetNetworkRangesForNetwork`: []ModelsSSLNetworkDefinition
    fmt.Fprintf(os.Stdout, "Response from `SslApi.SslGetNetworkRangesForNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Keyfactor network identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSslGetNetworkRangesForNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**[]ModelsSSLNetworkDefinition**](ModelsSSLNetworkDefinition.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SslGetNetworks

> []KeyfactorApiModelsSslNetworkQueryResponse SslGetNetworks(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).SqQueryString(sqQueryString).SqPageReturned(sqPageReturned).SqReturnLimit(sqReturnLimit).SqSortField(sqSortField).SqSortAscending(sqSortAscending).Execute()

Returns all defined SSL networks according to the provided filter and output parameters

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
    sqPageReturned := int32(56) // int32 | The current page within the result set to be returned (optional)
    sqReturnLimit := int32(56) // int32 | Maximum number of records to be returned in a single call (optional)
    sqSortField := "sqSortField_example" // string | Field by which the results should be sorted (view results via Management Portal for sortable columns) (optional)
    sqSortAscending := int32(56) // int32 | Field sort direction [0=ascending, 1=descending] (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SslApi.SslGetNetworks(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).SqQueryString(sqQueryString).SqPageReturned(sqPageReturned).SqReturnLimit(sqReturnLimit).SqSortField(sqSortField).SqSortAscending(sqSortAscending).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SslApi.SslGetNetworks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SslGetNetworks`: []KeyfactorApiModelsSslNetworkQueryResponse
    fmt.Fprintf(os.Stdout, "Response from `SslApi.SslGetNetworks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSslGetNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]
 **sqQueryString** | **string** | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) | 
 **sqPageReturned** | **int32** | The current page within the result set to be returned | 
 **sqReturnLimit** | **int32** | Maximum number of records to be returned in a single call | 
 **sqSortField** | **string** | Field by which the results should be sorted (view results via Management Portal for sortable columns) | 
 **sqSortAscending** | **int32** | Field sort direction [0&#x3D;ascending, 1&#x3D;descending] | 

### Return type

[**[]KeyfactorApiModelsSslNetworkQueryResponse**](KeyfactorApiModelsSslNetworkQueryResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SslImmediateSslScan

> SslImmediateSslScan(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).SslScanRequest(sslScanRequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Starts an SSL Scan for the network according to the associated network definition

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Keyfactor network identifier
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    sslScanRequest := *openapiclient.NewModelsSSLImmediateSslScanRequest(false, false) // ModelsSSLImmediateSslScanRequest | Request for an immediate SSL Scan
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SslApi.SslImmediateSslScan(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).SslScanRequest(sslScanRequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SslApi.SslImmediateSslScan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Keyfactor network identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSslImmediateSslScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **sslScanRequest** | [**ModelsSSLImmediateSslScanRequest**](ModelsSSLImmediateSslScanRequest.md) | Request for an immediate SSL Scan | 
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


## SslMonitorAll

> SslMonitorAll(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Query(query).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Sets all endpoints matching the provided query as 'monitored'

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
    query := "query_example" // string | Query to filter the endpoints for which the status should be set (optional)
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SslApi.SslMonitorAll(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Query(query).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SslApi.SslMonitorAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSslMonitorAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **query** | **string** | Query to filter the endpoints for which the status should be set | 
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


## SslMonitoringStatus

> SslMonitoringStatus(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Requests(requests).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Sets the monitored status according to the provided endpoint and boolean status

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
    requests := []openapiclient.ModelsSSLEndpointStatusRequest{*openapiclient.NewModelsSSLEndpointStatusRequest("00000000-0000-0000-0000-000000000000", false)} // []ModelsSSLEndpointStatusRequest | Endpoints and statuses to be set for each
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SslApi.SslMonitoringStatus(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Requests(requests).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SslApi.SslMonitoringStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSslMonitoringStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **requests** | [**[]ModelsSSLEndpointStatusRequest**](ModelsSSLEndpointStatusRequest.md) | Endpoints and statuses to be set for each | 
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


## SslNetworkScanParts

> []ModelsSSLDisplayScanJobPart SslNetworkScanParts(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PagedQueryJobType(pagedQueryJobType).PagedQueryQueryString(pagedQueryQueryString).PagedQueryPageReturned(pagedQueryPageReturned).PagedQueryReturnLimit(pagedQueryReturnLimit).PagedQuerySortField(pagedQuerySortField).PagedQuerySortAscending(pagedQuerySortAscending).Execute()

Returns the scan job components comprising the entire scan job to be executed

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Keyfactor network definition identifier
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")
    pagedQueryJobType := int32(56) // int32 |  (optional)
    pagedQueryQueryString := "pagedQueryQueryString_example" // string | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) (optional)
    pagedQueryPageReturned := int32(56) // int32 | The current page within the result set to be returned (optional)
    pagedQueryReturnLimit := int32(56) // int32 | Maximum number of records to be returned in a single call (optional)
    pagedQuerySortField := "pagedQuerySortField_example" // string | Field by which the results should be sorted (view results via Management Portal for sortable columns) (optional)
    pagedQuerySortAscending := int32(56) // int32 | Field sort direction [0=ascending, 1=descending] (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SslApi.SslNetworkScanParts(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PagedQueryJobType(pagedQueryJobType).PagedQueryQueryString(pagedQueryQueryString).PagedQueryPageReturned(pagedQueryPageReturned).PagedQueryReturnLimit(pagedQueryReturnLimit).PagedQuerySortField(pagedQuerySortField).PagedQuerySortAscending(pagedQuerySortAscending).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SslApi.SslNetworkScanParts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SslNetworkScanParts`: []ModelsSSLDisplayScanJobPart
    fmt.Fprintf(os.Stdout, "Response from `SslApi.SslNetworkScanParts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Keyfactor network definition identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSslNetworkScanPartsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]
 **pagedQueryJobType** | **int32** |  | 
 **pagedQueryQueryString** | **string** | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) | 
 **pagedQueryPageReturned** | **int32** | The current page within the result set to be returned | 
 **pagedQueryReturnLimit** | **int32** | Maximum number of records to be returned in a single call | 
 **pagedQuerySortField** | **string** | Field by which the results should be sorted (view results via Management Portal for sortable columns) | 
 **pagedQuerySortAscending** | **int32** | Field sort direction [0&#x3D;ascending, 1&#x3D;descending] | 

### Return type

[**[]ModelsSSLDisplayScanJobPart**](ModelsSSLDisplayScanJobPart.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SslRemoveAllNetworkRanges

> SslRemoveAllNetworkRanges(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Removes all network range definitions from the associated network definition

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Keyfactor network definition identifier
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SslApi.SslRemoveAllNetworkRanges(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SslApi.SslRemoveAllNetworkRanges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Keyfactor network definition identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSslRemoveAllNetworkRangesRequest struct via the builder pattern


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


## SslRemoveNetwork

> SslRemoveNetwork(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Removes a network definition according to the provided identifier

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Keyfactor network identifier
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SslApi.SslRemoveNetwork(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SslApi.SslRemoveNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Keyfactor network identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSslRemoveNetworkRequest struct via the builder pattern


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


## SslResults

> []ModelsSSLSslScanResult SslResults(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PqQueryString(pqQueryString).PqPageReturned(pqPageReturned).PqReturnLimit(pqReturnLimit).PqSortField(pqSortField).PqSortAscending(pqSortAscending).Execute()

Returns a list of the endpoint scan results according to the provided filter and output parameters

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

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SslApi.SslResults(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PqQueryString(pqQueryString).PqPageReturned(pqPageReturned).PqReturnLimit(pqReturnLimit).PqSortField(pqSortField).PqSortAscending(pqSortAscending).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SslApi.SslResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SslResults`: []ModelsSSLSslScanResult
    fmt.Fprintf(os.Stdout, "Response from `SslApi.SslResults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSslResultsRequest struct via the builder pattern


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

[**[]ModelsSSLSslScanResult**](ModelsSSLSslScanResult.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SslReviewAll

> SslReviewAll(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Query(query).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Sets all endpoints matching the provided query as 'reviewed'

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
    query := "query_example" // string | Query to filter the endpoints for which the status should be set (optional)
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SslApi.SslReviewAll(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Query(query).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SslApi.SslReviewAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSslReviewAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **query** | **string** | Query to filter the endpoints for which the status should be set | 
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


## SslReviewedStatus

> SslReviewedStatus(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).EndpointReviewedStatus(endpointReviewedStatus).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Sets the reviewed status according to the provided endpoint and boolean status

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
    endpointReviewedStatus := []openapiclient.ModelsSSLEndpointStatusRequest{*openapiclient.NewModelsSSLEndpointStatusRequest("00000000-0000-0000-0000-000000000000", false)} // []ModelsSSLEndpointStatusRequest | Endpoints and statuses for each
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SslApi.SslReviewedStatus(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).EndpointReviewedStatus(endpointReviewedStatus).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SslApi.SslReviewedStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSslReviewedStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **endpointReviewedStatus** | [**[]ModelsSSLEndpointStatusRequest**](ModelsSSLEndpointStatusRequest.md) | Endpoints and statuses for each | 
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


## SslScanPart

> ModelsSSLScanJobPart SslScanPart(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns the execution details of the associated network scan job part

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Keyfactor identifier of the scan job part
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SslApi.SslScanPart(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SslApi.SslScanPart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SslScanPart`: ModelsSSLScanJobPart
    fmt.Fprintf(os.Stdout, "Response from `SslApi.SslScanPart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Keyfactor identifier of the scan job part | 

### Other Parameters

Other parameters are passed through a pointer to a apiSslScanPartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsSSLScanJobPart**](ModelsSSLScanJobPart.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SslSetNetworkRanges

> SslSetNetworkRanges(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).NetworkRanges(networkRanges).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Configures network range definitions for the provided network

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
    networkRanges := *openapiclient.NewModelsSSLNetworkRangesRequest("00000000-0000-0000-0000-000000000000", []string{"Ranges_example"}) // ModelsSSLNetworkRangesRequest | Network range defitions and the network to which they should be set
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SslApi.SslSetNetworkRanges(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).NetworkRanges(networkRanges).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SslApi.SslSetNetworkRanges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSslSetNetworkRangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **networkRanges** | [**ModelsSSLNetworkRangesRequest**](ModelsSSLNetworkRangesRequest.md) | Network range defitions and the network to which they should be set | 
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


## SslUpdateNetwork

> KeyfactorApiModelsSslNetworkResponse SslUpdateNetwork(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Network(network).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Updates an existing network definition according to the provided properties

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
    network := *openapiclient.NewKeyfactorApiModelsSslUpdateNetworkRequest("00000000-0000-0000-0000-000000000000", "Name_example", "AgentPoolName_example", "Description_example") // KeyfactorApiModelsSslUpdateNetworkRequest | Properties of the network definition to be updated
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SslApi.SslUpdateNetwork(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Network(network).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SslApi.SslUpdateNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SslUpdateNetwork`: KeyfactorApiModelsSslNetworkResponse
    fmt.Fprintf(os.Stdout, "Response from `SslApi.SslUpdateNetwork`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSslUpdateNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **network** | [**KeyfactorApiModelsSslUpdateNetworkRequest**](KeyfactorApiModelsSslUpdateNetworkRequest.md) | Properties of the network definition to be updated | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsSslNetworkResponse**](KeyfactorApiModelsSslNetworkResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SslValidateNetworkRanges

> SslValidateNetworkRanges(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).NetworkRangesToVerify(networkRangesToVerify).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Validates the format (using regular expressions) of the provided network range definitions

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
    networkRangesToVerify := []string{"Property_example"} // []string | List of the network range definitions to verify
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SslApi.SslValidateNetworkRanges(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).NetworkRangesToVerify(networkRangesToVerify).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SslApi.SslValidateNetworkRanges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSslValidateNetworkRangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **networkRangesToVerify** | **[]string** | List of the network range definitions to verify | 
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


# \AnalyticsApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAnalyticsCertificatesCount**](AnalyticsApi.md#GetAnalyticsCertificatesCount) | **GET** /Analytics/Certificates/Count | Returns the count of certificates according to the provided filter and output parameters
[**GetAnalyticsCertificatesCountCollections**](AnalyticsApi.md#GetAnalyticsCertificatesCountCollections) | **GET** /Analytics/Certificates/Count/Collections | Returns the count of certificates for a given list of collections
[**GetAnalyticsCertificatesCountGrouped**](AnalyticsApi.md#GetAnalyticsCertificatesCountGrouped) | **GET** /Analytics/Certificates/Count/Grouped | Returns the grouped count of certificates according to the provided filter and output parameters
[**GetAnalyticsCertificatesIssuance**](AnalyticsApi.md#GetAnalyticsCertificatesIssuance) | **GET** /Analytics/Certificates/Issuance | Returns the count of certificates issued grouped by period
[**GetAnalyticsCertificatesRevocation**](AnalyticsApi.md#GetAnalyticsCertificatesRevocation) | **GET** /Analytics/Certificates/Revocation | Returns the count of certificates revoked grouped by period
[**GetAnalyticsSSLNetworksEndpoints**](AnalyticsApi.md#GetAnalyticsSSLNetworksEndpoints) | **GET** /Analytics/SSL/Networks/Endpoints | Returns the count of endpoints for a given list of networks (or all networks if none are provided)
[**GetAnalyticsSSLNetworksEndpointsStatus**](AnalyticsApi.md#GetAnalyticsSSLNetworksEndpointsStatus) | **GET** /Analytics/SSL/Networks/Endpoints/Status | Returns the count of endpoints grouped by status for a given list of networks (or all networks if none are provided)



## GetAnalyticsCertificatesCount

> AnalyticsAnalyticsCertificateCountResponse NewGetAnalyticsCertificatesCountRequest(ctx).CollectionId(collectionId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Query(query).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns the count of certificates according to the provided filter and output parameters

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
    collectionId := int32(56) // int32 | Certificate collection identifier used to filter the count, 0 indicates the global collection
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    query := "query_example" // string | Query string for the certificates returned in the count (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnalyticsApi.NewGetAnalyticsCertificatesCountRequest(context.Background()).CollectionId(collectionId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Query(query).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.GetAnalyticsCertificatesCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAnalyticsCertificatesCount`: AnalyticsAnalyticsCertificateCountResponse
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.GetAnalyticsCertificatesCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalyticsCertificatesCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **collectionId** | **int32** | Certificate collection identifier used to filter the count, 0 indicates the global collection | 
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **query** | **string** | Query string for the certificates returned in the count | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**AnalyticsAnalyticsCertificateCountResponse**](AnalyticsAnalyticsCertificateCountResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnalyticsCertificatesCountCollections

> []AnalyticsAnalyticsCollectionsCountResponse NewGetAnalyticsCertificatesCountCollectionsRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionIds(collectionIds).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns the count of certificates for a given list of collections

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
    collectionIds := []int32{int32(123)} // []int32 |  (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnalyticsApi.NewGetAnalyticsCertificatesCountCollectionsRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionIds(collectionIds).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.GetAnalyticsCertificatesCountCollections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAnalyticsCertificatesCountCollections`: []AnalyticsAnalyticsCollectionsCountResponse
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.GetAnalyticsCertificatesCountCollections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalyticsCertificatesCountCollectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **collectionIds** | **[]int32** |  | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]AnalyticsAnalyticsCollectionsCountResponse**](AnalyticsAnalyticsCollectionsCountResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnalyticsCertificatesCountGrouped

> []AnalyticsAnalyticsCertificateCountWithNameResponse NewGetAnalyticsCertificatesCountGroupedRequest(ctx).GroupByField(groupByField).CollectionId(collectionId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Query(query).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns the grouped count of certificates according to the provided filter and output parameters

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
    groupByField := "groupByField_example" // string | The name of a certificate query parser to group certificates by
    collectionId := int32(56) // int32 | Certificate collection identifier used to filter the count, 0 indicates the global collection
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    query := "query_example" // string | Query string for the certificates returned in the count (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnalyticsApi.NewGetAnalyticsCertificatesCountGroupedRequest(context.Background()).GroupByField(groupByField).CollectionId(collectionId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Query(query).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.GetAnalyticsCertificatesCountGrouped``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAnalyticsCertificatesCountGrouped`: []AnalyticsAnalyticsCertificateCountWithNameResponse
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.GetAnalyticsCertificatesCountGrouped`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalyticsCertificatesCountGroupedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupByField** | **string** | The name of a certificate query parser to group certificates by | 
 **collectionId** | **int32** | Certificate collection identifier used to filter the count, 0 indicates the global collection | 
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **query** | **string** | Query string for the certificates returned in the count | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]AnalyticsAnalyticsCertificateCountWithNameResponse**](AnalyticsAnalyticsCertificateCountWithNameResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnalyticsCertificatesIssuance

> AnalyticsAnalyticsCertificateCountWithPeriodResponse NewGetAnalyticsCertificatesIssuanceRequest(ctx).CollectionId(collectionId).TotalPeriods(totalPeriods).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns the count of certificates issued grouped by period

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
    collectionId := int32(56) // int32 | The Id of the Collection to filter by, 0 indicates the global collection
    totalPeriods := int32(56) // int32 | 
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnalyticsApi.NewGetAnalyticsCertificatesIssuanceRequest(context.Background()).CollectionId(collectionId).TotalPeriods(totalPeriods).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.GetAnalyticsCertificatesIssuance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAnalyticsCertificatesIssuance`: AnalyticsAnalyticsCertificateCountWithPeriodResponse
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.GetAnalyticsCertificatesIssuance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalyticsCertificatesIssuanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **collectionId** | **int32** | The Id of the Collection to filter by, 0 indicates the global collection | 
 **totalPeriods** | **int32** |  | 
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**AnalyticsAnalyticsCertificateCountWithPeriodResponse**](AnalyticsAnalyticsCertificateCountWithPeriodResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnalyticsCertificatesRevocation

> AnalyticsAnalyticsCertificateCountWithPeriodResponse NewGetAnalyticsCertificatesRevocationRequest(ctx).CollectionId(collectionId).TotalPeriods(totalPeriods).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns the count of certificates revoked grouped by period

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
    collectionId := int32(56) // int32 | The Id of the Collection to filter by, 0 indicates the global collection
    totalPeriods := int32(56) // int32 | 
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnalyticsApi.NewGetAnalyticsCertificatesRevocationRequest(context.Background()).CollectionId(collectionId).TotalPeriods(totalPeriods).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.GetAnalyticsCertificatesRevocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAnalyticsCertificatesRevocation`: AnalyticsAnalyticsCertificateCountWithPeriodResponse
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.GetAnalyticsCertificatesRevocation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalyticsCertificatesRevocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **collectionId** | **int32** | The Id of the Collection to filter by, 0 indicates the global collection | 
 **totalPeriods** | **int32** |  | 
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**AnalyticsAnalyticsCertificateCountWithPeriodResponse**](AnalyticsAnalyticsCertificateCountWithPeriodResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnalyticsSSLNetworksEndpoints

> []AnalyticsAnalyticsSSLNetworkEndpointCountResponse NewGetAnalyticsSSLNetworksEndpointsRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Networks(networks).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns the count of endpoints for a given list of networks (or all networks if none are provided)

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
    networks := []string{"Inner_example"} // []string |  (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnalyticsApi.NewGetAnalyticsSSLNetworksEndpointsRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Networks(networks).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.GetAnalyticsSSLNetworksEndpoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAnalyticsSSLNetworksEndpoints`: []AnalyticsAnalyticsSSLNetworkEndpointCountResponse
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.GetAnalyticsSSLNetworksEndpoints`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalyticsSSLNetworksEndpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **networks** | **[]string** |  | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]AnalyticsAnalyticsSSLNetworkEndpointCountResponse**](AnalyticsAnalyticsSSLNetworkEndpointCountResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnalyticsSSLNetworksEndpointsStatus

> []AnalyticsAnalyticsSSLNetworkEndpointStatusCountResponse NewGetAnalyticsSSLNetworksEndpointsStatusRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Networks(networks).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns the count of endpoints grouped by status for a given list of networks (or all networks if none are provided)

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
    networks := []string{"Inner_example"} // []string |  (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnalyticsApi.NewGetAnalyticsSSLNetworksEndpointsStatusRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Networks(networks).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.GetAnalyticsSSLNetworksEndpointsStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAnalyticsSSLNetworksEndpointsStatus`: []AnalyticsAnalyticsSSLNetworkEndpointStatusCountResponse
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.GetAnalyticsSSLNetworksEndpointsStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalyticsSSLNetworksEndpointsStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **networks** | **[]string** |  | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]AnalyticsAnalyticsSSLNetworkEndpointStatusCountResponse**](AnalyticsAnalyticsSSLNetworkEndpointStatusCountResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


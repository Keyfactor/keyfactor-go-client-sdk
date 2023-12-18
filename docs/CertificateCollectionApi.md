# \CertificateCollectionApi

All URIs are relative to */Keyfactor/API*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CertificateCollectionsCollectionListGet**](CertificateCollectionApi.md#CertificateCollectionsCollectionListGet) | **Get** /CertificateCollections/CollectionList | Get certificate collection list with duplication field name
[**CertificateCollectionsCopyPost**](CertificateCollectionApi.md#CertificateCollectionsCopyPost) | **Post** /CertificateCollections/Copy | Creates a new certificate collection from an existing collection. The permissions, query and description of the   existing collection are copied when creating the new record, with the option to overwrite the query or description.
[**CertificateCollectionsGet**](CertificateCollectionApi.md#CertificateCollectionsGet) | **Get** /CertificateCollections | Returns all certificate collections
[**CertificateCollectionsIdDelete**](CertificateCollectionApi.md#CertificateCollectionsIdDelete) | **Delete** /CertificateCollections/{id} | Delete one certificate collection
[**CertificateCollectionsIdFavoritePut**](CertificateCollectionApi.md#CertificateCollectionsIdFavoritePut) | **Put** /CertificateCollections/{id}/Favorite | Update favorite for one collection
[**CertificateCollectionsIdGet**](CertificateCollectionApi.md#CertificateCollectionsIdGet) | **Get** /CertificateCollections/{id} | Returns the certificate collection definition associated with the provided Keyfactor identifier
[**CertificateCollectionsNameGet**](CertificateCollectionApi.md#CertificateCollectionsNameGet) | **Get** /CertificateCollections/{name} | Returns the certificate collection associated with the provided collection name
[**CertificateCollectionsNavItemsGet**](CertificateCollectionApi.md#CertificateCollectionsNavItemsGet) | **Get** /CertificateCollections/NavItems | Get the list of navigation items for certificate collection
[**CertificateCollectionsPost**](CertificateCollectionApi.md#CertificateCollectionsPost) | **Post** /CertificateCollections | Creates a new certificate collection with the provided properties
[**CertificateCollectionsPut**](CertificateCollectionApi.md#CertificateCollectionsPut) | **Put** /CertificateCollections | Updates an existing certificate collection with the provided properties



## CertificateCollectionsCollectionListGet

> []KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionListResponse CertificateCollectionsCollectionListGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Get certificate collection list with duplication field name

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
    queryString := "queryString_example" // string |  (optional)
    pageReturned := int32(56) // int32 |  (optional)
    returnLimit := int32(56) // int32 |  (optional)
    sortField := "sortField_example" // string |  (optional)
    sortAscending := openapiclient.Keyfactor.Common.QueryableExtensionsSortOrder(0) // KeyfactorCommonQueryableExtensionsSortOrder |  (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateCollectionApi.CertificateCollectionsCollectionListGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateCollectionApi.CertificateCollectionsCollectionListGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateCollectionsCollectionListGet`: []KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionListResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateCollectionApi.CertificateCollectionsCollectionListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateCollectionsCollectionListGetRequest struct via the builder pattern


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

[**[]KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionListResponse**](KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionListResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateCollectionsCopyPost

> KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionResponse CertificateCollectionsCopyPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCopyRequest(keyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCopyRequest).Execute()

Creates a new certificate collection from an existing collection. The permissions, query and description of the   existing collection are copied when creating the new record, with the option to overwrite the query or description.



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
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCopyRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCopyRequest("Name_example", int32(123)) // KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCopyRequest | Information related to the certificate collection query (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateCollectionApi.CertificateCollectionsCopyPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCopyRequest(keyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCopyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateCollectionApi.CertificateCollectionsCopyPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateCollectionsCopyPost`: KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateCollectionApi.CertificateCollectionsCopyPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateCollectionsCopyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCopyRequest** | [**KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCopyRequest**](KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCopyRequest.md) | Information related to the certificate collection query | 

### Return type

[**KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionResponse**](KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patchjson, application/json, text/json, application/*json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateCollectionsGet

> []CSSCMSDataModelModelsCertificateQuery CertificateCollectionsGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns all certificate collections

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
    queryString := "queryString_example" // string |  (optional)
    pageReturned := int32(56) // int32 |  (optional)
    returnLimit := int32(56) // int32 |  (optional)
    sortField := "sortField_example" // string |  (optional)
    sortAscending := openapiclient.Keyfactor.Common.QueryableExtensionsSortOrder(0) // KeyfactorCommonQueryableExtensionsSortOrder |  (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateCollectionApi.CertificateCollectionsGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateCollectionApi.CertificateCollectionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateCollectionsGet`: []CSSCMSDataModelModelsCertificateQuery
    fmt.Fprintf(os.Stdout, "Response from `CertificateCollectionApi.CertificateCollectionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateCollectionsGetRequest struct via the builder pattern


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

[**[]CSSCMSDataModelModelsCertificateQuery**](CSSCMSDataModelModelsCertificateQuery.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateCollectionsIdDelete

> CertificateCollectionsIdDelete(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Delete one certificate collection

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
    id := int32(56) // int32 | The collection to delete
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateCollectionApi.CertificateCollectionsIdDelete(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateCollectionApi.CertificateCollectionsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The collection to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificateCollectionsIdDeleteRequest struct via the builder pattern


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


## CertificateCollectionsIdFavoritePut

> CertificateCollectionsIdFavoritePut(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionFavoriteRequest(keyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionFavoriteRequest).Execute()

Update favorite for one collection

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
    id := int32(56) // int32 | The collection to update favorite with
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionFavoriteRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionFavoriteRequest(false) // KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionFavoriteRequest | Information for the certificate collection favorite update (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateCollectionApi.CertificateCollectionsIdFavoritePut(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionFavoriteRequest(keyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionFavoriteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateCollectionApi.CertificateCollectionsIdFavoritePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The collection to update favorite with | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificateCollectionsIdFavoritePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionFavoriteRequest** | [**KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionFavoriteRequest**](KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionFavoriteRequest.md) | Information for the certificate collection favorite update | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patchjson, application/json, text/json, application/*json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateCollectionsIdGet

> CSSCMSDataModelModelsCertificateQuery CertificateCollectionsIdGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns the certificate collection definition associated with the provided Keyfactor identifier

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
    id := int32(56) // int32 | Identifier of the certificate collection
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateCollectionApi.CertificateCollectionsIdGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateCollectionApi.CertificateCollectionsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateCollectionsIdGet`: CSSCMSDataModelModelsCertificateQuery
    fmt.Fprintf(os.Stdout, "Response from `CertificateCollectionApi.CertificateCollectionsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Identifier of the certificate collection | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificateCollectionsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**CSSCMSDataModelModelsCertificateQuery**](CSSCMSDataModelModelsCertificateQuery.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateCollectionsNameGet

> CSSCMSDataModelModelsCertificateQuery CertificateCollectionsNameGet(ctx, name).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns the certificate collection associated with the provided collection name

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
    name := "name_example" // string | Name of the Collection
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateCollectionApi.CertificateCollectionsNameGet(context.Background(), name).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateCollectionApi.CertificateCollectionsNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateCollectionsNameGet`: CSSCMSDataModelModelsCertificateQuery
    fmt.Fprintf(os.Stdout, "Response from `CertificateCollectionApi.CertificateCollectionsNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the Collection | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificateCollectionsNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**CSSCMSDataModelModelsCertificateQuery**](CSSCMSDataModelModelsCertificateQuery.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateCollectionsNavItemsGet

> []CSSCMSDataModelModelsCertificateCollectionNavItem CertificateCollectionsNavItemsGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Get the list of navigation items for certificate collection

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
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateCollectionApi.CertificateCollectionsNavItemsGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateCollectionApi.CertificateCollectionsNavItemsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateCollectionsNavItemsGet`: []CSSCMSDataModelModelsCertificateCollectionNavItem
    fmt.Fprintf(os.Stdout, "Response from `CertificateCollectionApi.CertificateCollectionsNavItemsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateCollectionsNavItemsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]CSSCMSDataModelModelsCertificateCollectionNavItem**](CSSCMSDataModelModelsCertificateCollectionNavItem.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateCollectionsPost

> KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionResponse CertificateCollectionsPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest(keyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest).Execute()

Creates a new certificate collection with the provided properties



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
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest("Name_example") // KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest | Information related to the certificate collection query (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateCollectionApi.CertificateCollectionsPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest(keyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateCollectionApi.CertificateCollectionsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateCollectionsPost`: KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateCollectionApi.CertificateCollectionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateCollectionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest** | [**KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest**](KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest.md) | Information related to the certificate collection query | 

### Return type

[**KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionResponse**](KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patchjson, application/json, text/json, application/*json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateCollectionsPut

> KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionResponse CertificateCollectionsPut(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionUpdateRequest(keyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionUpdateRequest).Execute()

Updates an existing certificate collection with the provided properties



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
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionUpdateRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionUpdateRequest("Name_example", int32(123)) // KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionUpdateRequest | Information related to the certificate collection query (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateCollectionApi.CertificateCollectionsPut(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionUpdateRequest(keyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateCollectionApi.CertificateCollectionsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateCollectionsPut`: KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateCollectionApi.CertificateCollectionsPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateCollectionsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionUpdateRequest** | [**KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionUpdateRequest**](KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionUpdateRequest.md) | Information related to the certificate collection query | 

### Return type

[**KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionResponse**](KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patchjson, application/json, text/json, application/*json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


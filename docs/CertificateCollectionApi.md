# \CertificateCollectionApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CertificateCollectionCopyFromExistingCollection**](CertificateCollectionApi.md#CertificateCollectionCopyFromExistingCollection) | **Post** /CertificateCollections/Copy | Creates a new certificate collection from an existing collection. The permissions, query and description of the   existing collection are copied when creating the new record, with the option to overwrite the query or description.
[**CertificateCollectionCreateCollection**](CertificateCollectionApi.md#CertificateCollectionCreateCollection) | **Post** /CertificateCollections | Creates a new certificate collection with the provided properties
[**CertificateCollectionGetCollection0**](CertificateCollectionApi.md#CertificateCollectionGetCollection0) | **Get** /CertificateCollections/{id} | Returns the certificate collection definition associated with the provided Keyfactor identifier
[**CertificateCollectionGetCollection1**](CertificateCollectionApi.md#CertificateCollectionGetCollection1) | **Get** /CertificateCollections/{name} | Returns the certificate collection associated with the provided collection name
[**CertificateCollectionGetCollections**](CertificateCollectionApi.md#CertificateCollectionGetCollections) | **Get** /CertificateCollections | Returns all certificate collections
[**CertificateCollectionSetCollectionPermissions**](CertificateCollectionApi.md#CertificateCollectionSetCollectionPermissions) | **Post** /CertificateCollections/{id}/Permissions | Set the permissions for a collection
[**CertificateCollectionUpdateCollection**](CertificateCollectionApi.md#CertificateCollectionUpdateCollection) | **Put** /CertificateCollections | Updates an existing certificate collection with the provided properties



## CertificateCollectionCopyFromExistingCollection

> KeyfactorApiModelsCertificateCollectionsCertificateCollectionResponse CertificateCollectionCopyFromExistingCollection(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Request(request).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    request := *openapiclient.NewKeyfactorApiModelsCertificateCollectionsCertificateCollectionCopyRequest(int32(123), "Name_example") // KeyfactorApiModelsCertificateCollectionsCertificateCollectionCopyRequest | Information related to the certificate collection query
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateCollectionApi.CertificateCollectionCopyFromExistingCollection(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Request(request).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateCollectionApi.CertificateCollectionCopyFromExistingCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateCollectionCopyFromExistingCollection`: KeyfactorApiModelsCertificateCollectionsCertificateCollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateCollectionApi.CertificateCollectionCopyFromExistingCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateCollectionCopyFromExistingCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **request** | [**KeyfactorApiModelsCertificateCollectionsCertificateCollectionCopyRequest**](KeyfactorApiModelsCertificateCollectionsCertificateCollectionCopyRequest.md) | Information related to the certificate collection query | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsCertificateCollectionsCertificateCollectionResponse**](KeyfactorApiModelsCertificateCollectionsCertificateCollectionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateCollectionCreateCollection

> KeyfactorApiModelsCertificateCollectionsCertificateCollectionResponse CertificateCollectionCreateCollection(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Request(request).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    request := *openapiclient.NewKeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest("Name_example") // KeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest | Information related to the certificate collection query
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateCollectionApi.CertificateCollectionCreateCollection(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Request(request).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateCollectionApi.CertificateCollectionCreateCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateCollectionCreateCollection`: KeyfactorApiModelsCertificateCollectionsCertificateCollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateCollectionApi.CertificateCollectionCreateCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateCollectionCreateCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **request** | [**KeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest**](KeyfactorApiModelsCertificateCollectionsCertificateCollectionCreateRequest.md) | Information related to the certificate collection query | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsCertificateCollectionsCertificateCollectionResponse**](KeyfactorApiModelsCertificateCollectionsCertificateCollectionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateCollectionGetCollection0

> ModelsCertificateQuery CertificateCollectionGetCollection0(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateCollectionApi.CertificateCollectionGetCollection0(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateCollectionApi.CertificateCollectionGetCollection0``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateCollectionGetCollection0`: ModelsCertificateQuery
    fmt.Fprintf(os.Stdout, "Response from `CertificateCollectionApi.CertificateCollectionGetCollection0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Identifier of the certificate collection | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificateCollectionGetCollection0Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsCertificateQuery**](ModelsCertificateQuery.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateCollectionGetCollection1

> ModelsCertificateQuery CertificateCollectionGetCollection1(ctx, name).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateCollectionApi.CertificateCollectionGetCollection1(context.Background(), name).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateCollectionApi.CertificateCollectionGetCollection1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateCollectionGetCollection1`: ModelsCertificateQuery
    fmt.Fprintf(os.Stdout, "Response from `CertificateCollectionApi.CertificateCollectionGetCollection1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the Collection | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificateCollectionGetCollection1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsCertificateQuery**](ModelsCertificateQuery.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateCollectionGetCollections

> []ModelsCertificateQuery CertificateCollectionGetCollections(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PqQueryString(pqQueryString).PqPageReturned(pqPageReturned).PqReturnLimit(pqReturnLimit).PqSortField(pqSortField).PqSortAscending(pqSortAscending).Execute()

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
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")
    pqQueryString := "pqQueryString_example" // string | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) (optional)
    pqPageReturned := int32(56) // int32 | The current page within the result set to be returned (optional)
    pqReturnLimit := int32(56) // int32 | Maximum number of records to be returned in a single call (optional)
    pqSortField := "pqSortField_example" // string | Field by which the results should be sorted (view results via Management Portal for sortable columns) (optional)
    pqSortAscending := int32(56) // int32 | Field sort direction [0=ascending, 1=descending] (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateCollectionApi.CertificateCollectionGetCollections(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PqQueryString(pqQueryString).PqPageReturned(pqPageReturned).PqReturnLimit(pqReturnLimit).PqSortField(pqSortField).PqSortAscending(pqSortAscending).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateCollectionApi.CertificateCollectionGetCollections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateCollectionGetCollections`: []ModelsCertificateQuery
    fmt.Fprintf(os.Stdout, "Response from `CertificateCollectionApi.CertificateCollectionGetCollections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateCollectionGetCollectionsRequest struct via the builder pattern


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

[**[]ModelsCertificateQuery**](ModelsCertificateQuery.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateCollectionSetCollectionPermissions

> CertificateCollectionSetCollectionPermissions(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionPermissions(collectionPermissions).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Set the permissions for a collection



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
    id := int32(56) // int32 | The collection to set permissions on
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    collectionPermissions := []openapiclient.ModelsCollectionRolePermissions{*openapiclient.NewModelsCollectionRolePermissions()} // []ModelsCollectionRolePermissions | The collection Permissions object ['Read', 'EditMetadata', 'Recover', 'Revoke', 'Delete']
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateCollectionApi.CertificateCollectionSetCollectionPermissions(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionPermissions(collectionPermissions).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateCollectionApi.CertificateCollectionSetCollectionPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The collection to set permissions on | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificateCollectionSetCollectionPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **collectionPermissions** | [**[]ModelsCollectionRolePermissions**](ModelsCollectionRolePermissions.md) | The collection Permissions object [&#39;Read&#39;, &#39;EditMetadata&#39;, &#39;Recover&#39;, &#39;Revoke&#39;, &#39;Delete&#39;] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateCollectionUpdateCollection

> KeyfactorApiModelsCertificateCollectionsCertificateCollectionResponse CertificateCollectionUpdateCollection(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Request(request).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    request := *openapiclient.NewKeyfactorApiModelsCertificateCollectionsCertificateCollectionUpdateRequest(int32(123), "Name_example") // KeyfactorApiModelsCertificateCollectionsCertificateCollectionUpdateRequest | Information related to the certificate collection query
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateCollectionApi.CertificateCollectionUpdateCollection(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Request(request).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateCollectionApi.CertificateCollectionUpdateCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateCollectionUpdateCollection`: KeyfactorApiModelsCertificateCollectionsCertificateCollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateCollectionApi.CertificateCollectionUpdateCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateCollectionUpdateCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **request** | [**KeyfactorApiModelsCertificateCollectionsCertificateCollectionUpdateRequest**](KeyfactorApiModelsCertificateCollectionsCertificateCollectionUpdateRequest.md) | Information related to the certificate collection query | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsCertificateCollectionsCertificateCollectionResponse**](KeyfactorApiModelsCertificateCollectionsCertificateCollectionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


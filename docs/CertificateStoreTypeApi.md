# \CertificateStoreTypeApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CertificateStoreTypeCreateCertificateStoreType**](CertificateStoreTypeApi.md#CertificateStoreTypeCreateCertificateStoreType) | **Post** /CertificateStoreTypes | Creates a new certificate store type with the provided properties
[**CertificateStoreTypeDeleteCertificateStoreType**](CertificateStoreTypeApi.md#CertificateStoreTypeDeleteCertificateStoreType) | **Delete** /CertificateStoreTypes/{id} | Deletes a certificate store type according to the provided identifier
[**CertificateStoreTypeDeleteCertificateStoreTypes**](CertificateStoreTypeApi.md#CertificateStoreTypeDeleteCertificateStoreTypes) | **Delete** /CertificateStoreTypes | Deletes certificate store types according to the provided identifiers
[**CertificateStoreTypeGetCertificateStoreType0**](CertificateStoreTypeApi.md#CertificateStoreTypeGetCertificateStoreType0) | **Get** /CertificateStoreTypes/{id} | Returns a single certificate store type that matches id
[**CertificateStoreTypeGetCertificateStoreType1**](CertificateStoreTypeApi.md#CertificateStoreTypeGetCertificateStoreType1) | **Get** /CertificateStoreTypes/Name/{name} | Returns a single certificate store type that matches the provided short name
[**CertificateStoreTypeGetTypes**](CertificateStoreTypeApi.md#CertificateStoreTypeGetTypes) | **Get** /CertificateStoreTypes | Returns all certificate store types according to the provided filter and output parameters
[**CertificateStoreTypeUpdateCertificateStoreType**](CertificateStoreTypeApi.md#CertificateStoreTypeUpdateCertificateStoreType) | **Put** /CertificateStoreTypes | Updates an existing certificate store type with the provided properties



## CertificateStoreTypeCreateCertificateStoreType

> KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse CertificateStoreTypeCreateCertificateStoreType(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CertStoreType(certStoreType).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Creates a new certificate store type with the provided properties

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
    certStoreType := *openapiclient.NewKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest("Name_example", "ShortName_example") // KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest | Certificate store type properties for the new type
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreTypeApi.CertificateStoreTypeCreateCertificateStoreType(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CertStoreType(certStoreType).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreTypeApi.CertificateStoreTypeCreateCertificateStoreType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateStoreTypeCreateCertificateStoreType`: KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateStoreTypeApi.CertificateStoreTypeCreateCertificateStoreType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoreTypeCreateCertificateStoreTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **certStoreType** | [**KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest**](KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest.md) | Certificate store type properties for the new type | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse**](KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateStoreTypeDeleteCertificateStoreType

> CertificateStoreTypeDeleteCertificateStoreType(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Deletes a certificate store type according to the provided identifier



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
    id := int32(56) // int32 | Keyfactor identifier of the certificate store type to be deleted
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreTypeApi.CertificateStoreTypeDeleteCertificateStoreType(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreTypeApi.CertificateStoreTypeDeleteCertificateStoreType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Keyfactor identifier of the certificate store type to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoreTypeDeleteCertificateStoreTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateStoreTypeDeleteCertificateStoreTypes

> CertificateStoreTypeDeleteCertificateStoreTypes(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Ids(ids).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Deletes certificate store types according to the provided identifiers



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
    ids := []int32{int32(123)} // []int32 | Array of Keyfactor identifiers of the certificate store types to be deleted
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreTypeApi.CertificateStoreTypeDeleteCertificateStoreTypes(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Ids(ids).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreTypeApi.CertificateStoreTypeDeleteCertificateStoreTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoreTypeDeleteCertificateStoreTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **ids** | **[]int32** | Array of Keyfactor identifiers of the certificate store types to be deleted | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateStoreTypeGetCertificateStoreType0

> KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse CertificateStoreTypeGetCertificateStoreType0(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns a single certificate store type that matches id

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
    id := int32(56) // int32 | Keyfactor identifier of the certificate store type
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreTypeApi.CertificateStoreTypeGetCertificateStoreType0(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreTypeApi.CertificateStoreTypeGetCertificateStoreType0``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateStoreTypeGetCertificateStoreType0`: KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateStoreTypeApi.CertificateStoreTypeGetCertificateStoreType0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Keyfactor identifier of the certificate store type | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoreTypeGetCertificateStoreType0Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse**](KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateStoreTypeGetCertificateStoreType1

> []KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse CertificateStoreTypeGetCertificateStoreType1(ctx, name).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns a single certificate store type that matches the provided short name

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
    name := "name_example" // string | Short name of the certificate store type to return
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreTypeApi.CertificateStoreTypeGetCertificateStoreType1(context.Background(), name).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreTypeApi.CertificateStoreTypeGetCertificateStoreType1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateStoreTypeGetCertificateStoreType1`: []KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateStoreTypeApi.CertificateStoreTypeGetCertificateStoreType1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Short name of the certificate store type to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoreTypeGetCertificateStoreType1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**[]KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse**](KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateStoreTypeGetTypes

> []KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse CertificateStoreTypeGetTypes(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CstqueryQueryString(cstqueryQueryString).CstqueryPageReturned(cstqueryPageReturned).CstqueryReturnLimit(cstqueryReturnLimit).CstquerySortField(cstquerySortField).CstquerySortAscending(cstquerySortAscending).Execute()

Returns all certificate store types according to the provided filter and output parameters

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
    cstqueryQueryString := "cstqueryQueryString_example" // string | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) (optional)
    cstqueryPageReturned := int32(56) // int32 | The current page within the result set to be returned (optional)
    cstqueryReturnLimit := int32(56) // int32 | Maximum number of records to be returned in a single call (optional)
    cstquerySortField := "cstquerySortField_example" // string | Field by which the results should be sorted (view results via Management Portal for sortable columns) (optional)
    cstquerySortAscending := int32(56) // int32 | Field sort direction [0=ascending, 1=descending] (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreTypeApi.CertificateStoreTypeGetTypes(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CstqueryQueryString(cstqueryQueryString).CstqueryPageReturned(cstqueryPageReturned).CstqueryReturnLimit(cstqueryReturnLimit).CstquerySortField(cstquerySortField).CstquerySortAscending(cstquerySortAscending).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreTypeApi.CertificateStoreTypeGetTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateStoreTypeGetTypes`: []KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateStoreTypeApi.CertificateStoreTypeGetTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoreTypeGetTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]
 **cstqueryQueryString** | **string** | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) | 
 **cstqueryPageReturned** | **int32** | The current page within the result set to be returned | 
 **cstqueryReturnLimit** | **int32** | Maximum number of records to be returned in a single call | 
 **cstquerySortField** | **string** | Field by which the results should be sorted (view results via Management Portal for sortable columns) | 
 **cstquerySortAscending** | **int32** | Field sort direction [0&#x3D;ascending, 1&#x3D;descending] | 

### Return type

[**[]KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse**](KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateStoreTypeUpdateCertificateStoreType

> KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse CertificateStoreTypeUpdateCertificateStoreType(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CertStoreType(certStoreType).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Updates an existing certificate store type with the provided properties

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
    certStoreType := *openapiclient.NewKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest(int32(123), "Name_example", "ShortName_example") // KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest | Certificate store type properties to be updated
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreTypeApi.CertificateStoreTypeUpdateCertificateStoreType(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CertStoreType(certStoreType).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreTypeApi.CertificateStoreTypeUpdateCertificateStoreType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateStoreTypeUpdateCertificateStoreType`: KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateStoreTypeApi.CertificateStoreTypeUpdateCertificateStoreType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoreTypeUpdateCertificateStoreTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **certStoreType** | [**KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest**](KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest.md) | Certificate store type properties to be updated | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse**](KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


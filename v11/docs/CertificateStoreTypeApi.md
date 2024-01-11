# \CertificateStoreTypeApi

All URIs are relative to */Keyfactor/API*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CertificateStoreTypesDelete**](CertificateStoreTypeApi.md#CertificateStoreTypesDelete) | **Delete** /CertificateStoreTypes | Deletes certificate store types according to the provided identifiers
[**CertificateStoreTypesGet**](CertificateStoreTypeApi.md#CertificateStoreTypesGet) | **Get** /CertificateStoreTypes | Returns all certificate store types according to the provided filter and output parameters
[**CertificateStoreTypesIdDelete**](CertificateStoreTypeApi.md#CertificateStoreTypesIdDelete) | **Delete** /CertificateStoreTypes/{id} | Deletes a certificate store type according to the provided identifier
[**CertificateStoreTypesIdGet**](CertificateStoreTypeApi.md#CertificateStoreTypesIdGet) | **Get** /CertificateStoreTypes/{id} | Returns a single certificate store type that matches id
[**CertificateStoreTypesNameNameGet**](CertificateStoreTypeApi.md#CertificateStoreTypesNameNameGet) | **Get** /CertificateStoreTypes/Name/{name} | Returns a single certificate store type that matches the provided short name
[**CertificateStoreTypesPost**](CertificateStoreTypeApi.md#CertificateStoreTypesPost) | **Post** /CertificateStoreTypes | Creates a new certificate store type with the provided properties
[**CertificateStoreTypesPut**](CertificateStoreTypeApi.md#CertificateStoreTypesPut) | **Put** /CertificateStoreTypes | Updates an existing certificate store type with the provided properties



## CertificateStoreTypesDelete

> CertificateStoreTypesDelete(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).RequestBody(requestBody).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    requestBody := []int32{int32(123)} // []int32 | Array of Keyfactor identifiers of the certificate store types to be deleted (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreTypeApi.CertificateStoreTypesDelete(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreTypeApi.CertificateStoreTypesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoreTypesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **requestBody** | **[]int32** | Array of Keyfactor identifiers of the certificate store types to be deleted | 

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


## CertificateStoreTypesGet

> []KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse CertificateStoreTypesGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    queryString := "queryString_example" // string |  (optional)
    pageReturned := int32(56) // int32 |  (optional)
    returnLimit := int32(56) // int32 |  (optional)
    sortField := "sortField_example" // string |  (optional)
    sortAscending := int32(56) // int32 |  (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreTypeApi.CertificateStoreTypesGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreTypeApi.CertificateStoreTypesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateStoreTypesGet`: []KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateStoreTypeApi.CertificateStoreTypesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoreTypesGetRequest struct via the builder pattern


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

[**[]KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse**](KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateStoreTypesIdDelete

> CertificateStoreTypesIdDelete(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreTypeApi.CertificateStoreTypesIdDelete(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreTypeApi.CertificateStoreTypesIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCertificateStoreTypesIdDeleteRequest struct via the builder pattern


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


## CertificateStoreTypesIdGet

> KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse CertificateStoreTypesIdGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreTypeApi.CertificateStoreTypesIdGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreTypeApi.CertificateStoreTypesIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateStoreTypesIdGet`: KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateStoreTypeApi.CertificateStoreTypesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Keyfactor identifier of the certificate store type | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoreTypesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse**](KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateStoreTypesNameNameGet

> []KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse CertificateStoreTypesNameNameGet(ctx, name).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreTypeApi.CertificateStoreTypesNameNameGet(context.Background(), name).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreTypeApi.CertificateStoreTypesNameNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateStoreTypesNameNameGet`: []KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateStoreTypeApi.CertificateStoreTypesNameNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Short name of the certificate store type to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoreTypesNameNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse**](KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateStoreTypesPost

> KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse CertificateStoreTypesPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest(keyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest("Name_example", "ShortName_example") // KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest | Certificate store type properties for the new type (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreTypeApi.CertificateStoreTypesPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest(keyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreTypeApi.CertificateStoreTypesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateStoreTypesPost`: KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateStoreTypeApi.CertificateStoreTypesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoreTypesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest** | [**KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest**](KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest.md) | Certificate store type properties for the new type | 

### Return type

[**KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse**](KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateStoreTypesPut

> KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse CertificateStoreTypesPut(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest(keyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest(int32(123), "Name_example", "ShortName_example") // KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest | Certificate store type properties to be updated (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateStoreTypeApi.CertificateStoreTypesPut(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest(keyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateStoreTypeApi.CertificateStoreTypesPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateStoreTypesPut`: KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateStoreTypeApi.CertificateStoreTypesPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateStoreTypesPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest** | [**KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest**](KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest.md) | Certificate store type properties to be updated | 

### Return type

[**KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse**](KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


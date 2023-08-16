# \CertificateAuthorityApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CertificateAuthorityCreateCA**](CertificateAuthorityApi.md#CertificateAuthorityCreateCA) | **Post** /CertificateAuthority | Creates a new CertificateAuthority object
[**CertificateAuthorityDeleteCA**](CertificateAuthorityApi.md#CertificateAuthorityDeleteCA) | **Delete** /CertificateAuthority/{id} | Deletes a CertificateAuthority from the system, specified by ID
[**CertificateAuthorityGetCa**](CertificateAuthorityApi.md#CertificateAuthorityGetCa) | **Get** /CertificateAuthority/{id} | Returns details for a single CA, specified by ID
[**CertificateAuthorityGetCas**](CertificateAuthorityApi.md#CertificateAuthorityGetCas) | **Get** /CertificateAuthority | Returns all certificate authorities
[**CertificateAuthorityPublishCRL**](CertificateAuthorityApi.md#CertificateAuthorityPublishCRL) | **Post** /CertificateAuthority/PublishCRL | Publishes a CRL according to the provided request
[**CertificateAuthorityTestCertificateAuthority**](CertificateAuthorityApi.md#CertificateAuthorityTestCertificateAuthority) | **Post** /CertificateAuthority/Test | Validates the connection info for the CA provided by the model.
[**CertificateAuthorityUpdateCA**](CertificateAuthorityApi.md#CertificateAuthorityUpdateCA) | **Put** /CertificateAuthority | Updates a CertificateAuthority object



## CertificateAuthorityCreateCA

> ModelsCertificateAuthoritiesCertificateAuthorityResponse CertificateAuthorityCreateCA(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Ca(ca).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Creates a new CertificateAuthority object

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
    ca := *openapiclient.NewModelsCertificateAuthoritiesCertificateAuthorityRequest() // ModelsCertificateAuthoritiesCertificateAuthorityRequest | 
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.CertificateAuthorityCreateCA(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Ca(ca).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.CertificateAuthorityCreateCA``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateAuthorityCreateCA`: ModelsCertificateAuthoritiesCertificateAuthorityResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityApi.CertificateAuthorityCreateCA`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateAuthorityCreateCARequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **ca** | [**ModelsCertificateAuthoritiesCertificateAuthorityRequest**](ModelsCertificateAuthoritiesCertificateAuthorityRequest.md) |  | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsCertificateAuthoritiesCertificateAuthorityResponse**](ModelsCertificateAuthoritiesCertificateAuthorityResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateAuthorityDeleteCA

> CertificateAuthorityDeleteCA(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Deletes a CertificateAuthority from the system, specified by ID

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
    id := int32(56) // int32 | 
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.CertificateAuthorityDeleteCA(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.CertificateAuthorityDeleteCA``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificateAuthorityDeleteCARequest struct via the builder pattern


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


## CertificateAuthorityGetCa

> ModelsCertificateAuthoritiesCertificateAuthorityResponse CertificateAuthorityGetCa(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns details for a single CA, specified by ID

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
    id := int32(56) // int32 | 
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.CertificateAuthorityGetCa(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.CertificateAuthorityGetCa``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateAuthorityGetCa`: ModelsCertificateAuthoritiesCertificateAuthorityResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityApi.CertificateAuthorityGetCa`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificateAuthorityGetCaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsCertificateAuthoritiesCertificateAuthorityResponse**](ModelsCertificateAuthoritiesCertificateAuthorityResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateAuthorityGetCas

> []ModelsCertificateAuthoritiesCertificateAuthorityResponse CertificateAuthorityGetCas(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PqQueryString(pqQueryString).PqPageReturned(pqPageReturned).PqReturnLimit(pqReturnLimit).PqSortField(pqSortField).PqSortAscending(pqSortAscending).Execute()

Returns all certificate authorities

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
    resp, r, err := apiClient.CertificateAuthorityApi.CertificateAuthorityGetCas(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PqQueryString(pqQueryString).PqPageReturned(pqPageReturned).PqReturnLimit(pqReturnLimit).PqSortField(pqSortField).PqSortAscending(pqSortAscending).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.CertificateAuthorityGetCas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateAuthorityGetCas`: []ModelsCertificateAuthoritiesCertificateAuthorityResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityApi.CertificateAuthorityGetCas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateAuthorityGetCasRequest struct via the builder pattern


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

[**[]ModelsCertificateAuthoritiesCertificateAuthorityResponse**](ModelsCertificateAuthoritiesCertificateAuthorityResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateAuthorityPublishCRL

> CertificateAuthorityPublishCRL(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Crlrequest(crlrequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Publishes a CRL according to the provided request

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
    crlrequest := *openapiclient.NewModelsCRLRequestModel("CertificateAuthorityLogicalName_example") // ModelsCRLRequestModel | Host and logical name of the CA for which the CRL should be published
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.CertificateAuthorityPublishCRL(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Crlrequest(crlrequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.CertificateAuthorityPublishCRL``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateAuthorityPublishCRLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **crlrequest** | [**ModelsCRLRequestModel**](ModelsCRLRequestModel.md) | Host and logical name of the CA for which the CRL should be published | 
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


## CertificateAuthorityTestCertificateAuthority

> KeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityTestResponse CertificateAuthorityTestCertificateAuthority(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Ca(ca).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Validates the connection info for the CA provided by the model.

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
    ca := *openapiclient.NewModelsCertificateAuthoritiesCertificateAuthorityRequest() // ModelsCertificateAuthoritiesCertificateAuthorityRequest | The CA being tested.
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.CertificateAuthorityTestCertificateAuthority(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Ca(ca).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.CertificateAuthorityTestCertificateAuthority``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateAuthorityTestCertificateAuthority`: KeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityTestResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityApi.CertificateAuthorityTestCertificateAuthority`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateAuthorityTestCertificateAuthorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **ca** | [**ModelsCertificateAuthoritiesCertificateAuthorityRequest**](ModelsCertificateAuthoritiesCertificateAuthorityRequest.md) | The CA being tested. | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityTestResponse**](KeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityTestResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateAuthorityUpdateCA

> ModelsCertificateAuthoritiesCertificateAuthorityResponse CertificateAuthorityUpdateCA(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Ca(ca).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Updates a CertificateAuthority object

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
    ca := *openapiclient.NewModelsCertificateAuthoritiesCertificateAuthorityRequest() // ModelsCertificateAuthoritiesCertificateAuthorityRequest | 
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.CertificateAuthorityUpdateCA(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Ca(ca).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.CertificateAuthorityUpdateCA``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateAuthorityUpdateCA`: ModelsCertificateAuthoritiesCertificateAuthorityResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityApi.CertificateAuthorityUpdateCA`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateAuthorityUpdateCARequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **ca** | [**ModelsCertificateAuthoritiesCertificateAuthorityRequest**](ModelsCertificateAuthoritiesCertificateAuthorityRequest.md) |  | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsCertificateAuthoritiesCertificateAuthorityResponse**](ModelsCertificateAuthoritiesCertificateAuthorityResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


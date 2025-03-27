# \CAConnectorApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCertificateAuthorityCAConnectors**](CAConnectorApi.md#CreateCertificateAuthorityCAConnectors) | **POST** /CertificateAuthority/CAConnectors | Creates a new CA Connector object
[**DeleteCertificateAuthorityCAConnectorsById**](CAConnectorApi.md#DeleteCertificateAuthorityCAConnectorsById) | **DELETE** /CertificateAuthority/CAConnectors/{id} | Deletes a CA Connector with the specific ID
[**GetCertificateAuthorityCAConnectors**](CAConnectorApi.md#GetCertificateAuthorityCAConnectors) | **GET** /CertificateAuthority/CAConnectors | Returns all CA Connectors
[**GetCertificateAuthorityCAConnectorsById**](CAConnectorApi.md#GetCertificateAuthorityCAConnectorsById) | **GET** /CertificateAuthority/CAConnectors/{id} | Returns a CA Connector with the specific ID
[**UpdateCertificateAuthorityCAConnectorsById**](CAConnectorApi.md#UpdateCertificateAuthorityCAConnectorsById) | **PUT** /CertificateAuthority/CAConnectors/{id} | Updates an existing CA Connector



## CreateCertificateAuthorityCAConnectors

> CertificateAuthoritiesCAConnectorResponse BuildCreateCertificateAuthorityCAConnectorsRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CertificateAuthoritiesCAConnectorRequest(certificateAuthoritiesCAConnectorRequest).Execute()

Creates a new CA Connector object

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
    certificateAuthoritiesCAConnectorRequest := *openapiclient.NewCertificateAuthoritiesCAConnectorRequest() // CertificateAuthoritiesCAConnectorRequest | CA Connector object with the provided information. (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CAConnectorApi.BuildCreateCertificateAuthorityCAConnectorsRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CertificateAuthoritiesCAConnectorRequest(certificateAuthoritiesCAConnectorRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CAConnectorApi.CreateCertificateAuthorityCAConnectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCertificateAuthorityCAConnectors`: CertificateAuthoritiesCAConnectorResponse
    fmt.Fprintf(os.Stdout, "Response from `CAConnectorApi.CreateCertificateAuthorityCAConnectors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCertificateAuthorityCAConnectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **certificateAuthoritiesCAConnectorRequest** | [**CertificateAuthoritiesCAConnectorRequest**](CertificateAuthoritiesCAConnectorRequest.md) | CA Connector object with the provided information. | 

### Return type

[**CertificateAuthoritiesCAConnectorResponse**](CertificateAuthoritiesCAConnectorResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCertificateAuthorityCAConnectorsById

> BuildDeleteCertificateAuthorityCAConnectorsByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Deletes a CA Connector with the specific ID

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
    id := int32(56) // int32 | The ID of the CA Connector
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CAConnectorApi.BuildDeleteCertificateAuthorityCAConnectorsByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CAConnectorApi.DeleteCertificateAuthorityCAConnectorsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the CA Connector | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCertificateAuthorityCAConnectorsByIdRequest struct via the builder pattern


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


## GetCertificateAuthorityCAConnectors

> []CertificateAuthoritiesCAConnectorResponse BuildGetCertificateAuthorityCAConnectorsRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns all CA Connectors

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

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CAConnectorApi.BuildGetCertificateAuthorityCAConnectorsRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CAConnectorApi.GetCertificateAuthorityCAConnectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificateAuthorityCAConnectors`: []CertificateAuthoritiesCAConnectorResponse
    fmt.Fprintf(os.Stdout, "Response from `CAConnectorApi.GetCertificateAuthorityCAConnectors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificateAuthorityCAConnectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]CertificateAuthoritiesCAConnectorResponse**](CertificateAuthoritiesCAConnectorResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificateAuthorityCAConnectorsById

> CertificateAuthoritiesCAConnectorResponse BuildGetCertificateAuthorityCAConnectorsByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns a CA Connector with the specific ID

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
    id := int32(56) // int32 | The ID of the CA Connector
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CAConnectorApi.BuildGetCertificateAuthorityCAConnectorsByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CAConnectorApi.GetCertificateAuthorityCAConnectorsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificateAuthorityCAConnectorsById`: CertificateAuthoritiesCAConnectorResponse
    fmt.Fprintf(os.Stdout, "Response from `CAConnectorApi.GetCertificateAuthorityCAConnectorsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the CA Connector | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificateAuthorityCAConnectorsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**CertificateAuthoritiesCAConnectorResponse**](CertificateAuthoritiesCAConnectorResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCertificateAuthorityCAConnectorsById

> CertificateAuthoritiesCAConnectorResponse BuildUpdateCertificateAuthorityCAConnectorsByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CertificateAuthoritiesCAConnectorRequest(certificateAuthoritiesCAConnectorRequest).Execute()

Updates an existing CA Connector

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
    id := int32(56) // int32 | The ID of the CA Connector we want to update
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    certificateAuthoritiesCAConnectorRequest := *openapiclient.NewCertificateAuthoritiesCAConnectorRequest() // CertificateAuthoritiesCAConnectorRequest | CA Connector object with the provided information. (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CAConnectorApi.BuildUpdateCertificateAuthorityCAConnectorsByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CertificateAuthoritiesCAConnectorRequest(certificateAuthoritiesCAConnectorRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CAConnectorApi.UpdateCertificateAuthorityCAConnectorsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCertificateAuthorityCAConnectorsById`: CertificateAuthoritiesCAConnectorResponse
    fmt.Fprintf(os.Stdout, "Response from `CAConnectorApi.UpdateCertificateAuthorityCAConnectorsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the CA Connector we want to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCertificateAuthorityCAConnectorsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **certificateAuthoritiesCAConnectorRequest** | [**CertificateAuthoritiesCAConnectorRequest**](CertificateAuthoritiesCAConnectorRequest.md) | CA Connector object with the provided information. | 

### Return type

[**CertificateAuthoritiesCAConnectorResponse**](CertificateAuthoritiesCAConnectorResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


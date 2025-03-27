# \EnrollmentApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEnrollmentCSR**](EnrollmentApi.md#CreateEnrollmentCSR) | **POST** /Enrollment/CSR | Performs a CSR Enrollment based upon the provided request
[**CreateEnrollmentCSRParse**](EnrollmentApi.md#CreateEnrollmentCSRParse) | **POST** /Enrollment/CSR/Parse | Parses the provided CSR and returns the properties
[**CreateEnrollmentPFX**](EnrollmentApi.md#CreateEnrollmentPFX) | **POST** /Enrollment/PFX | Performs a PFX Enrollment based upon the provided request
[**CreateEnrollmentPFXDeploy**](EnrollmentApi.md#CreateEnrollmentPFXDeploy) | **POST** /Enrollment/PFX/Deploy | Creates management jobs to install a newly enrolled pfx in to one or more certificate stores
[**CreateEnrollmentPFXReplace**](EnrollmentApi.md#CreateEnrollmentPFXReplace) | **POST** /Enrollment/PFX/Replace | Creates management jobs to install a newly enrolled pfx into the same certificate stores as the previous certificate
[**CreateEnrollmentRenew**](EnrollmentApi.md#CreateEnrollmentRenew) | **POST** /Enrollment/Renew | Performs a renewal based upon the passed in request
[**GetEnrollmentAvailableRenewalIdById**](EnrollmentApi.md#GetEnrollmentAvailableRenewalIdById) | **GET** /Enrollment/AvailableRenewal/Id/{id} | Returns the type of renewal available for a given certificate.
[**GetEnrollmentAvailableRenewalThumbprintThumbprint**](EnrollmentApi.md#GetEnrollmentAvailableRenewalThumbprintThumbprint) | **GET** /Enrollment/AvailableRenewal/Thumbprint/{thumbprint} | Returns the type of renewal available for a given certificate.
[**GetEnrollmentCSRContextMy**](EnrollmentApi.md#GetEnrollmentCSRContextMy) | **GET** /Enrollment/CSR/Context/My | Returns the list of available CSR enrollment templates and their associated CA mappings that the calling user has permissions on
[**GetEnrollmentPFXContextMy**](EnrollmentApi.md#GetEnrollmentPFXContextMy) | **GET** /Enrollment/PFX/Context/My | Returns the list of available PFX enrollment templates and their associated CA mappings that the calling user has permissions on
[**GetEnrollmentSettingsById**](EnrollmentApi.md#GetEnrollmentSettingsById) | **GET** /Enrollment/Settings/{id} | Gets the template settings to use during enrollment. The response will be the resolved values for the settings.  If there is a template specific setting, the template specific setting will be used in the response.  If there is not a template specific setting, the global setting will be used in the response.



## CreateEnrollmentCSR

> CSSCMSDataModelModelsEnrollmentCSREnrollmentResponse BuildCreateEnrollmentCSRRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XCertificateformat(xCertificateformat).ForceEnroll(forceEnroll).XKeyfactorApiVersion(xKeyfactorApiVersion).EnrollmentCSREnrollmentRequest(enrollmentCSREnrollmentRequest).Execute()

Performs a CSR Enrollment based upon the provided request



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
    xCertificateformat := "PEM" // string | Desired format [PEM, DER]
    forceEnroll := true // bool |  (optional) (default to false)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    enrollmentCSREnrollmentRequest := *openapiclient.NewEnrollmentCSREnrollmentRequest("CSR_example") // EnrollmentCSREnrollmentRequest | Information needed to perform the CSR Enrollment (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.BuildCreateEnrollmentCSRRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XCertificateformat(xCertificateformat).ForceEnroll(forceEnroll).XKeyfactorApiVersion(xKeyfactorApiVersion).EnrollmentCSREnrollmentRequest(enrollmentCSREnrollmentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.CreateEnrollmentCSR``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEnrollmentCSR`: CSSCMSDataModelModelsEnrollmentCSREnrollmentResponse
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.CreateEnrollmentCSR`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnrollmentCSRRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xCertificateformat** | **string** | Desired format [PEM, DER] | 
 **forceEnroll** | **bool** |  | [default to false]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **enrollmentCSREnrollmentRequest** | [**EnrollmentCSREnrollmentRequest**](EnrollmentCSREnrollmentRequest.md) | Information needed to perform the CSR Enrollment | 

### Return type

[**CSSCMSDataModelModelsEnrollmentCSREnrollmentResponse**](CSSCMSDataModelModelsEnrollmentCSREnrollmentResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEnrollmentCSRParse

> []string BuildCreateEnrollmentCSRParseRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsCSRContents(cSSCMSDataModelModelsCSRContents).Execute()

Parses the provided CSR and returns the properties



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
    cSSCMSDataModelModelsCSRContents := *openapiclient.NewCSSCMSDataModelModelsCSRContents("CSR_example") // CSSCMSDataModelModelsCSRContents | CSR to be parsed (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.BuildCreateEnrollmentCSRParseRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsCSRContents(cSSCMSDataModelModelsCSRContents).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.CreateEnrollmentCSRParse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEnrollmentCSRParse`: []string
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.CreateEnrollmentCSRParse`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnrollmentCSRParseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsCSRContents** | [**CSSCMSDataModelModelsCSRContents**](CSSCMSDataModelModelsCSRContents.md) | CSR to be parsed | 

### Return type

**[]string**

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEnrollmentPFX

> CSSCMSDataModelModelsEnrollmentPFXEnrollmentResponse BuildCreateEnrollmentPFXRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XCertificateformat(xCertificateformat).XKeyfactorApiVersion(xKeyfactorApiVersion).EnrollmentPFXEnrollmentRequest(enrollmentPFXEnrollmentRequest).Execute()

Performs a PFX Enrollment based upon the provided request



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
    xCertificateformat := "PFX" // string | Desired format [PFX, ZIP, PEM, JKS, STORE, REPLACE]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    enrollmentPFXEnrollmentRequest := *openapiclient.NewEnrollmentPFXEnrollmentRequest() // EnrollmentPFXEnrollmentRequest | The information needed to perform the PFX Enrollment (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.BuildCreateEnrollmentPFXRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XCertificateformat(xCertificateformat).XKeyfactorApiVersion(xKeyfactorApiVersion).EnrollmentPFXEnrollmentRequest(enrollmentPFXEnrollmentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.CreateEnrollmentPFX``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEnrollmentPFX`: CSSCMSDataModelModelsEnrollmentPFXEnrollmentResponse
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.CreateEnrollmentPFX`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnrollmentPFXRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xCertificateformat** | **string** | Desired format [PFX, ZIP, PEM, JKS, STORE, REPLACE] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **enrollmentPFXEnrollmentRequest** | [**EnrollmentPFXEnrollmentRequest**](EnrollmentPFXEnrollmentRequest.md) | The information needed to perform the PFX Enrollment | 

### Return type

[**CSSCMSDataModelModelsEnrollmentPFXEnrollmentResponse**](CSSCMSDataModelModelsEnrollmentPFXEnrollmentResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEnrollmentPFXDeploy

> EnrollmentEnrollmentManagementResponse BuildCreateEnrollmentPFXDeployRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).EnrollmentEnrollmentManagementRequest(enrollmentEnrollmentManagementRequest).Execute()

Creates management jobs to install a newly enrolled pfx in to one or more certificate stores

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
    enrollmentEnrollmentManagementRequest := *openapiclient.NewEnrollmentEnrollmentManagementRequest("Password_example") // EnrollmentEnrollmentManagementRequest | The request to create the management jobs, which includes the request Id of the new pfx and the Ids and management job properties of the cert stores to add the pfx to (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.BuildCreateEnrollmentPFXDeployRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).EnrollmentEnrollmentManagementRequest(enrollmentEnrollmentManagementRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.CreateEnrollmentPFXDeploy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEnrollmentPFXDeploy`: EnrollmentEnrollmentManagementResponse
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.CreateEnrollmentPFXDeploy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnrollmentPFXDeployRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **enrollmentEnrollmentManagementRequest** | [**EnrollmentEnrollmentManagementRequest**](EnrollmentEnrollmentManagementRequest.md) | The request to create the management jobs, which includes the request Id of the new pfx and the Ids and management job properties of the cert stores to add the pfx to | 

### Return type

[**EnrollmentEnrollmentManagementResponse**](EnrollmentEnrollmentManagementResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEnrollmentPFXReplace

> EnrollmentEnrollmentManagementResponse BuildCreateEnrollmentPFXReplaceRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsEnrollmentExistingEnrollmentManagementRequest(cSSCMSDataModelModelsEnrollmentExistingEnrollmentManagementRequest).Execute()

Creates management jobs to install a newly enrolled pfx into the same certificate stores as the previous certificate

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
    cSSCMSDataModelModelsEnrollmentExistingEnrollmentManagementRequest := *openapiclient.NewCSSCMSDataModelModelsEnrollmentExistingEnrollmentManagementRequest() // CSSCMSDataModelModelsEnrollmentExistingEnrollmentManagementRequest | The request to create the management jobs, which includes the request Id of the new pfx and the Id of the existing certificate (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.BuildCreateEnrollmentPFXReplaceRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsEnrollmentExistingEnrollmentManagementRequest(cSSCMSDataModelModelsEnrollmentExistingEnrollmentManagementRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.CreateEnrollmentPFXReplace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEnrollmentPFXReplace`: EnrollmentEnrollmentManagementResponse
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.CreateEnrollmentPFXReplace`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnrollmentPFXReplaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsEnrollmentExistingEnrollmentManagementRequest** | [**CSSCMSDataModelModelsEnrollmentExistingEnrollmentManagementRequest**](CSSCMSDataModelModelsEnrollmentExistingEnrollmentManagementRequest.md) | The request to create the management jobs, which includes the request Id of the new pfx and the Id of the existing certificate | 

### Return type

[**EnrollmentEnrollmentManagementResponse**](EnrollmentEnrollmentManagementResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEnrollmentRenew

> EnrollmentRenewalApiResponse BuildCreateEnrollmentRenewRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsEnrollmentRenewalRequest(cSSCMSDataModelModelsEnrollmentRenewalRequest).Execute()

Performs a renewal based upon the passed in request

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
    collectionId := int32(56) // int32 | The collection id for the given certificate (optional) (default to 0)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    cSSCMSDataModelModelsEnrollmentRenewalRequest := *openapiclient.NewCSSCMSDataModelModelsEnrollmentRenewalRequest() // CSSCMSDataModelModelsEnrollmentRenewalRequest | The information needed to perform the renewal (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.BuildCreateEnrollmentRenewRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsEnrollmentRenewalRequest(cSSCMSDataModelModelsEnrollmentRenewalRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.CreateEnrollmentRenew``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEnrollmentRenew`: EnrollmentRenewalApiResponse
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.CreateEnrollmentRenew`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnrollmentRenewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **collectionId** | **int32** | The collection id for the given certificate | [default to 0]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsEnrollmentRenewalRequest** | [**CSSCMSDataModelModelsEnrollmentRenewalRequest**](CSSCMSDataModelModelsEnrollmentRenewalRequest.md) | The information needed to perform the renewal | 

### Return type

[**EnrollmentRenewalApiResponse**](EnrollmentRenewalApiResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnrollmentAvailableRenewalIdById

> CSSCMSDataModelModelsEnrollmentAvailableRenewal BuildGetEnrollmentAvailableRenewalIdByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns the type of renewal available for a given certificate.



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
    id := int32(56) // int32 | The Keyfactor certificate Id
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    collectionId := int32(56) // int32 | The collection id for the given certificate (optional) (default to 0)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.BuildGetEnrollmentAvailableRenewalIdByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.GetEnrollmentAvailableRenewalIdById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnrollmentAvailableRenewalIdById`: CSSCMSDataModelModelsEnrollmentAvailableRenewal
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.GetEnrollmentAvailableRenewalIdById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Keyfactor certificate Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnrollmentAvailableRenewalIdByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **collectionId** | **int32** | The collection id for the given certificate | [default to 0]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**CSSCMSDataModelModelsEnrollmentAvailableRenewal**](CSSCMSDataModelModelsEnrollmentAvailableRenewal.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnrollmentAvailableRenewalThumbprintThumbprint

> CSSCMSDataModelModelsEnrollmentAvailableRenewal BuildGetEnrollmentAvailableRenewalThumbprintThumbprintRequest(ctx, thumbprint).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns the type of renewal available for a given certificate.

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
    thumbprint := "thumbprint_example" // string | The certificate thumbprint
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    collectionId := int32(56) // int32 | The collection id for the given certificate (optional) (default to 0)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.BuildGetEnrollmentAvailableRenewalThumbprintThumbprintRequest(context.Background(), thumbprint).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.GetEnrollmentAvailableRenewalThumbprintThumbprint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnrollmentAvailableRenewalThumbprintThumbprint`: CSSCMSDataModelModelsEnrollmentAvailableRenewal
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.GetEnrollmentAvailableRenewalThumbprintThumbprint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**thumbprint** | **string** | The certificate thumbprint | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnrollmentAvailableRenewalThumbprintThumbprintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **collectionId** | **int32** | The collection id for the given certificate | [default to 0]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**CSSCMSDataModelModelsEnrollmentAvailableRenewal**](CSSCMSDataModelModelsEnrollmentAvailableRenewal.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnrollmentCSRContextMy

> KeyfactorWebCoreModelsEnrollmentEnrollmentTemplateCAResponse BuildGetEnrollmentCSRContextMyRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns the list of available CSR enrollment templates and their associated CA mappings that the calling user has permissions on

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
    resp, r, err := apiClient.EnrollmentApi.BuildGetEnrollmentCSRContextMyRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.GetEnrollmentCSRContextMy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnrollmentCSRContextMy`: KeyfactorWebCoreModelsEnrollmentEnrollmentTemplateCAResponse
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.GetEnrollmentCSRContextMy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEnrollmentCSRContextMyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**KeyfactorWebCoreModelsEnrollmentEnrollmentTemplateCAResponse**](KeyfactorWebCoreModelsEnrollmentEnrollmentTemplateCAResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnrollmentPFXContextMy

> KeyfactorWebCoreModelsEnrollmentEnrollmentTemplateCAResponse BuildGetEnrollmentPFXContextMyRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns the list of available PFX enrollment templates and their associated CA mappings that the calling user has permissions on

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
    resp, r, err := apiClient.EnrollmentApi.BuildGetEnrollmentPFXContextMyRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.GetEnrollmentPFXContextMy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnrollmentPFXContextMy`: KeyfactorWebCoreModelsEnrollmentEnrollmentTemplateCAResponse
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.GetEnrollmentPFXContextMy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEnrollmentPFXContextMyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**KeyfactorWebCoreModelsEnrollmentEnrollmentTemplateCAResponse**](KeyfactorWebCoreModelsEnrollmentEnrollmentTemplateCAResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnrollmentSettingsById

> TemplatesEnrollmentTemplateEnrollmentSettingsResponse BuildGetEnrollmentSettingsByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Gets the template settings to use during enrollment. The response will be the resolved values for the settings.  If there is a template specific setting, the template specific setting will be used in the response.  If there is not a template specific setting, the global setting will be used in the response.

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.BuildGetEnrollmentSettingsByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.GetEnrollmentSettingsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnrollmentSettingsById`: TemplatesEnrollmentTemplateEnrollmentSettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.GetEnrollmentSettingsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnrollmentSettingsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**TemplatesEnrollmentTemplateEnrollmentSettingsResponse**](TemplatesEnrollmentTemplateEnrollmentSettingsResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


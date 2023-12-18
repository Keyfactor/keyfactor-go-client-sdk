# \EnrollmentApi

All URIs are relative to */Keyfactor/API*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnrollmentAvailableRenewalIdIdGet**](EnrollmentApi.md#EnrollmentAvailableRenewalIdIdGet) | **Get** /Enrollment/AvailableRenewal/Id/{id} | Returns the type of renewal available for a given certificate.
[**EnrollmentAvailableRenewalThumbprintThumbprintGet**](EnrollmentApi.md#EnrollmentAvailableRenewalThumbprintThumbprintGet) | **Get** /Enrollment/AvailableRenewal/Thumbprint/{thumbprint} | Returns the type of renewal available for a given certificate.
[**EnrollmentCSRContextMyGet**](EnrollmentApi.md#EnrollmentCSRContextMyGet) | **Get** /Enrollment/CSR/Context/My | Returns the list of available CSR enrollment templates and their associated CA mappings that the calling user has permissions on
[**EnrollmentCSRParsePost**](EnrollmentApi.md#EnrollmentCSRParsePost) | **Post** /Enrollment/CSR/Parse | Parses the provided CSR and returns the properties
[**EnrollmentCSRPost**](EnrollmentApi.md#EnrollmentCSRPost) | **Post** /Enrollment/CSR | Performs a CSR Enrollment based upon the provided request
[**EnrollmentPFXContextMyGet**](EnrollmentApi.md#EnrollmentPFXContextMyGet) | **Get** /Enrollment/PFX/Context/My | Returns the list of available PFX enrollment templates and their associated CA mappings that the calling user has permissions on
[**EnrollmentPFXDeployPost**](EnrollmentApi.md#EnrollmentPFXDeployPost) | **Post** /Enrollment/PFX/Deploy | Creates management jobs to install a newly enrolled pfx in to one or more certificate stores
[**EnrollmentPFXPost**](EnrollmentApi.md#EnrollmentPFXPost) | **Post** /Enrollment/PFX | Performs a PFX Enrollment based upon the provided request
[**EnrollmentPFXReplacePost**](EnrollmentApi.md#EnrollmentPFXReplacePost) | **Post** /Enrollment/PFX/Replace | Creates management jobs to install a newly enrolled pfx into the same certificate stores as the previous certificate
[**EnrollmentRenewPost**](EnrollmentApi.md#EnrollmentRenewPost) | **Post** /Enrollment/Renew | Performs a renewal based upon the passed in request
[**EnrollmentSettingsIdGet**](EnrollmentApi.md#EnrollmentSettingsIdGet) | **Get** /Enrollment/Settings/{id} | Gets the template settings to use during enrollment. The response will be the resolved values for the settings.  If there is a template specific setting, the template specific setting will be used in the response.  If there is not a template specific setting, the global setting will be used in the response.



## EnrollmentAvailableRenewalIdIdGet

> CSSCMSDataModelModelsEnrollmentAvailableRenewal EnrollmentAvailableRenewalIdIdGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    collectionId := int32(56) // int32 | The collection id for the given certificate (optional) (default to 0)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.EnrollmentAvailableRenewalIdIdGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.EnrollmentAvailableRenewalIdIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnrollmentAvailableRenewalIdIdGet`: CSSCMSDataModelModelsEnrollmentAvailableRenewal
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.EnrollmentAvailableRenewalIdIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Keyfactor certificate Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnrollmentAvailableRenewalIdIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
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


## EnrollmentAvailableRenewalThumbprintThumbprintGet

> CSSCMSDataModelModelsEnrollmentAvailableRenewal EnrollmentAvailableRenewalThumbprintThumbprintGet(ctx, thumbprint).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    collectionId := int32(56) // int32 | The collection id for the given certificate (optional) (default to 0)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.EnrollmentAvailableRenewalThumbprintThumbprintGet(context.Background(), thumbprint).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.EnrollmentAvailableRenewalThumbprintThumbprintGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnrollmentAvailableRenewalThumbprintThumbprintGet`: CSSCMSDataModelModelsEnrollmentAvailableRenewal
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.EnrollmentAvailableRenewalThumbprintThumbprintGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**thumbprint** | **string** | The certificate thumbprint | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnrollmentAvailableRenewalThumbprintThumbprintGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
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


## EnrollmentCSRContextMyGet

> KeyfactorWebCoreModelsEnrollmentEnrollmentTemplateCAResponse EnrollmentCSRContextMyGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.EnrollmentCSRContextMyGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.EnrollmentCSRContextMyGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnrollmentCSRContextMyGet`: KeyfactorWebCoreModelsEnrollmentEnrollmentTemplateCAResponse
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.EnrollmentCSRContextMyGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnrollmentCSRContextMyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
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


## EnrollmentCSRParsePost

> []string EnrollmentCSRParsePost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsCSRContents(cSSCMSDataModelModelsCSRContents).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    cSSCMSDataModelModelsCSRContents := *openapiclient.NewCSSCMSDataModelModelsCSRContents("Csr_example") // CSSCMSDataModelModelsCSRContents | CSR to be parsed (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.EnrollmentCSRParsePost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsCSRContents(cSSCMSDataModelModelsCSRContents).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.EnrollmentCSRParsePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnrollmentCSRParsePost`: []string
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.EnrollmentCSRParsePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnrollmentCSRParsePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
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


## EnrollmentCSRPost

> CSSCMSDataModelModelsEnrollmentCSREnrollmentResponse EnrollmentCSRPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest(cSSCMSDataModelModelsEnrollmentCSREnrollmentRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    cSSCMSDataModelModelsEnrollmentCSREnrollmentRequest := *openapiclient.NewCSSCMSDataModelModelsEnrollmentCSREnrollmentRequest("Csr_example") // CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest | Information needed to perform the CSR Enrollment (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.EnrollmentCSRPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest(cSSCMSDataModelModelsEnrollmentCSREnrollmentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.EnrollmentCSRPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnrollmentCSRPost`: CSSCMSDataModelModelsEnrollmentCSREnrollmentResponse
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.EnrollmentCSRPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnrollmentCSRPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsEnrollmentCSREnrollmentRequest** | [**CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest**](CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest.md) | Information needed to perform the CSR Enrollment | 

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


## EnrollmentPFXContextMyGet

> KeyfactorWebCoreModelsEnrollmentEnrollmentTemplateCAResponse EnrollmentPFXContextMyGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.EnrollmentPFXContextMyGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.EnrollmentPFXContextMyGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnrollmentPFXContextMyGet`: KeyfactorWebCoreModelsEnrollmentEnrollmentTemplateCAResponse
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.EnrollmentPFXContextMyGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnrollmentPFXContextMyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
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


## EnrollmentPFXDeployPost

> KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementResponse EnrollmentPFXDeployPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest(keyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest("Password_example") // KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest | The request to create the management jobs, which includes the request Id of the new pfx and the Ids and management job properties of the cert stores to add the pfx to (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.EnrollmentPFXDeployPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest(keyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.EnrollmentPFXDeployPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnrollmentPFXDeployPost`: KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementResponse
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.EnrollmentPFXDeployPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnrollmentPFXDeployPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest** | [**KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest**](KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementRequest.md) | The request to create the management jobs, which includes the request Id of the new pfx and the Ids and management job properties of the cert stores to add the pfx to | 

### Return type

[**KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementResponse**](KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnrollmentPFXPost

> CSSCMSDataModelModelsEnrollmentPFXEnrollmentResponse EnrollmentPFXPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest(cSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    cSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest := *openapiclient.NewCSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest() // CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest | The information needed to perform the PFX Enrollment (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.EnrollmentPFXPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest(cSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.EnrollmentPFXPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnrollmentPFXPost`: CSSCMSDataModelModelsEnrollmentPFXEnrollmentResponse
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.EnrollmentPFXPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnrollmentPFXPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest** | [**CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest**](CSSCMSDataModelModelsEnrollmentPFXEnrollmentRequest.md) | The information needed to perform the PFX Enrollment | 

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


## EnrollmentPFXReplacePost

> KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementResponse EnrollmentPFXReplacePost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsEnrollmentExistingEnrollmentManagementRequest(cSSCMSDataModelModelsEnrollmentExistingEnrollmentManagementRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    cSSCMSDataModelModelsEnrollmentExistingEnrollmentManagementRequest := *openapiclient.NewCSSCMSDataModelModelsEnrollmentExistingEnrollmentManagementRequest() // CSSCMSDataModelModelsEnrollmentExistingEnrollmentManagementRequest | The request to create the management jobs, which includes the request Id of the new pfx and the Id of the existing certificate (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.EnrollmentPFXReplacePost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsEnrollmentExistingEnrollmentManagementRequest(cSSCMSDataModelModelsEnrollmentExistingEnrollmentManagementRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.EnrollmentPFXReplacePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnrollmentPFXReplacePost`: KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementResponse
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.EnrollmentPFXReplacePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnrollmentPFXReplacePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsEnrollmentExistingEnrollmentManagementRequest** | [**CSSCMSDataModelModelsEnrollmentExistingEnrollmentManagementRequest**](CSSCMSDataModelModelsEnrollmentExistingEnrollmentManagementRequest.md) | The request to create the management jobs, which includes the request Id of the new pfx and the Id of the existing certificate | 

### Return type

[**KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementResponse**](KeyfactorWebKeyfactorApiModelsEnrollmentEnrollmentManagementResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnrollmentRenewPost

> CSSCMSDataModelModelsEnrollmentRenewalResponse EnrollmentRenewPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsEnrollmentRenewalRequest(cSSCMSDataModelModelsEnrollmentRenewalRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    collectionId := int32(56) // int32 | The collection id for the given certificate (optional) (default to 0)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    cSSCMSDataModelModelsEnrollmentRenewalRequest := *openapiclient.NewCSSCMSDataModelModelsEnrollmentRenewalRequest() // CSSCMSDataModelModelsEnrollmentRenewalRequest | The information needed to perform the renewal (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.EnrollmentRenewPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsEnrollmentRenewalRequest(cSSCMSDataModelModelsEnrollmentRenewalRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.EnrollmentRenewPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnrollmentRenewPost`: CSSCMSDataModelModelsEnrollmentRenewalResponse
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.EnrollmentRenewPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnrollmentRenewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **collectionId** | **int32** | The collection id for the given certificate | [default to 0]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsEnrollmentRenewalRequest** | [**CSSCMSDataModelModelsEnrollmentRenewalRequest**](CSSCMSDataModelModelsEnrollmentRenewalRequest.md) | The information needed to perform the renewal | 

### Return type

[**CSSCMSDataModelModelsEnrollmentRenewalResponse**](CSSCMSDataModelModelsEnrollmentRenewalResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnrollmentSettingsIdGet

> KeyfactorWebKeyfactorApiModelsTemplatesEnrollmentTemplateEnrollmentSettingsResponse EnrollmentSettingsIdGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.EnrollmentSettingsIdGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.EnrollmentSettingsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnrollmentSettingsIdGet`: KeyfactorWebKeyfactorApiModelsTemplatesEnrollmentTemplateEnrollmentSettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.EnrollmentSettingsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnrollmentSettingsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**KeyfactorWebKeyfactorApiModelsTemplatesEnrollmentTemplateEnrollmentSettingsResponse**](KeyfactorWebKeyfactorApiModelsTemplatesEnrollmentTemplateEnrollmentSettingsResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


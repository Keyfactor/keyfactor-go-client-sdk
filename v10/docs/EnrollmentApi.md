# \EnrollmentApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnrollmentAddToExistingCertStores**](EnrollmentApi.md#EnrollmentAddToExistingCertStores) | **Post** /Enrollment/PFX/Replace | Creates management jobs to install a newly enrolled pfx into the same certificate stores as the previous certificate
[**EnrollmentAvailableRenewalId**](EnrollmentApi.md#EnrollmentAvailableRenewalId) | **Get** /Enrollment/AvailableRenewal/Id/{id} | Returns the type of renewal available for a given certificate.
[**EnrollmentAvailableRenewalThumbprint**](EnrollmentApi.md#EnrollmentAvailableRenewalThumbprint) | **Get** /Enrollment/AvailableRenewal/Thumbprint/{thumbprint} | Returns the type of renewal available for a given certificate.
[**EnrollmentGetMyCSRContext**](EnrollmentApi.md#EnrollmentGetMyCSRContext) | **Get** /Enrollment/CSR/Context/My | Returns the list of available CSR enrollment templates and their associated CA mappings that the calling user has permissions on
[**EnrollmentGetMyPFXContext**](EnrollmentApi.md#EnrollmentGetMyPFXContext) | **Get** /Enrollment/PFX/Context/My | Returns the list of available PFX enrollment templates and their associated CA mappings that the calling user has permissions on
[**EnrollmentGetTemplateEnrollmentSettings**](EnrollmentApi.md#EnrollmentGetTemplateEnrollmentSettings) | **Get** /Enrollment/Settings/{id} | Gets the template settings to use during enrollment. The response will be the resolved values for the settings.  If there is a template specific setting, the template specific setting will be used in the response.  If there is not a template specific setting, the global setting will be used in the response.
[**EnrollmentInstallPFXToCertStore**](EnrollmentApi.md#EnrollmentInstallPFXToCertStore) | **Post** /Enrollment/PFX/Deploy | Creates management jobs to install a newly enrolled pfx in to one or more certificate stores
[**EnrollmentPostCSREnroll**](EnrollmentApi.md#EnrollmentPostCSREnroll) | **Post** /Enrollment/CSR | Performs a CSR Enrollment based upon the provided request
[**EnrollmentPostPFXEnroll**](EnrollmentApi.md#EnrollmentPostPFXEnroll) | **Post** /Enrollment/PFX | Performs a PFX Enrollment based upon the provided request
[**EnrollmentPostParsedCSR**](EnrollmentApi.md#EnrollmentPostParsedCSR) | **Post** /Enrollment/CSR/Parse | Parses the provided CSR and returns the properties
[**EnrollmentRenew**](EnrollmentApi.md#EnrollmentRenew) | **Post** /Enrollment/Renew | Performs a renewal based upon the passed in request



## EnrollmentAddToExistingCertStores

> KeyfactorAPIModelsEnrollmentEnrollmentManagementResponse EnrollmentAddToExistingCertStores(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Request(request).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    request := *openapiclient.NewModelsEnrollmentExistingEnrollmentManagementRequest() // ModelsEnrollmentExistingEnrollmentManagementRequest | The request to create the management jobs, which includes the request Id of the new pfx and the Id of the existing certificate
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.EnrollmentAddToExistingCertStores(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Request(request).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.EnrollmentAddToExistingCertStores``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnrollmentAddToExistingCertStores`: KeyfactorAPIModelsEnrollmentEnrollmentManagementResponse
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.EnrollmentAddToExistingCertStores`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnrollmentAddToExistingCertStoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **request** | [**ModelsEnrollmentExistingEnrollmentManagementRequest**](ModelsEnrollmentExistingEnrollmentManagementRequest.md) | The request to create the management jobs, which includes the request Id of the new pfx and the Id of the existing certificate | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorAPIModelsEnrollmentEnrollmentManagementResponse**](KeyfactorAPIModelsEnrollmentEnrollmentManagementResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnrollmentAvailableRenewalId

> ModelsEnrollmentAvailableRenewal EnrollmentAvailableRenewalId(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    id := int64(789) // int64 | The Keyfactor certificate Id
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    collectionId := int64(789) // int64 | The collection id for the given certificate (optional)
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.EnrollmentAvailableRenewalId(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.EnrollmentAvailableRenewalId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnrollmentAvailableRenewalId`: ModelsEnrollmentAvailableRenewal
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.EnrollmentAvailableRenewalId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The Keyfactor certificate Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnrollmentAvailableRenewalIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **collectionId** | **int64** | The collection id for the given certificate | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsEnrollmentAvailableRenewal**](ModelsEnrollmentAvailableRenewal.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnrollmentAvailableRenewalThumbprint

> ModelsEnrollmentAvailableRenewal EnrollmentAvailableRenewalThumbprint(ctx, thumbprint).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    collectionId := int64(789) // int64 | The collection id for the given certificate (optional)
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.EnrollmentAvailableRenewalThumbprint(context.Background(), thumbprint).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.EnrollmentAvailableRenewalThumbprint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnrollmentAvailableRenewalThumbprint`: ModelsEnrollmentAvailableRenewal
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.EnrollmentAvailableRenewalThumbprint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**thumbprint** | **string** | The certificate thumbprint | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnrollmentAvailableRenewalThumbprintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **collectionId** | **int64** | The collection id for the given certificate | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsEnrollmentAvailableRenewal**](ModelsEnrollmentAvailableRenewal.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnrollmentGetMyCSRContext

> CoreModelsEnrollmentEnrollmentTemplateCAResponse EnrollmentGetMyCSRContext(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.EnrollmentGetMyCSRContext(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.EnrollmentGetMyCSRContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnrollmentGetMyCSRContext`: CoreModelsEnrollmentEnrollmentTemplateCAResponse
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.EnrollmentGetMyCSRContext`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnrollmentGetMyCSRContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**CoreModelsEnrollmentEnrollmentTemplateCAResponse**](CoreModelsEnrollmentEnrollmentTemplateCAResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnrollmentGetMyPFXContext

> CoreModelsEnrollmentEnrollmentTemplateCAResponse EnrollmentGetMyPFXContext(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.EnrollmentGetMyPFXContext(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.EnrollmentGetMyPFXContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnrollmentGetMyPFXContext`: CoreModelsEnrollmentEnrollmentTemplateCAResponse
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.EnrollmentGetMyPFXContext`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnrollmentGetMyPFXContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**CoreModelsEnrollmentEnrollmentTemplateCAResponse**](CoreModelsEnrollmentEnrollmentTemplateCAResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnrollmentGetTemplateEnrollmentSettings

> KeyfactorApiModelsTemplatesTemplateEnrollmentSettingsResponse EnrollmentGetTemplateEnrollmentSettings(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    id := int64(789) // int64 | 
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.EnrollmentGetTemplateEnrollmentSettings(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.EnrollmentGetTemplateEnrollmentSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnrollmentGetTemplateEnrollmentSettings`: KeyfactorApiModelsTemplatesTemplateEnrollmentSettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.EnrollmentGetTemplateEnrollmentSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnrollmentGetTemplateEnrollmentSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsTemplatesTemplateEnrollmentSettingsResponse**](KeyfactorApiModelsTemplatesTemplateEnrollmentSettingsResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnrollmentInstallPFXToCertStore

> KeyfactorAPIModelsEnrollmentEnrollmentManagementResponse EnrollmentInstallPFXToCertStore(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Request(request).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    request := *openapiclient.NewKeyfactorApiModelsEnrollmentEnrollmentManagementRequest("Password_example") // KeyfactorApiModelsEnrollmentEnrollmentManagementRequest | The request to create the management jobs, which includes the request Id of the new pfx and the Ids and management job properties of the cert stores to add the pfx to
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.EnrollmentInstallPFXToCertStore(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Request(request).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.EnrollmentInstallPFXToCertStore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnrollmentInstallPFXToCertStore`: KeyfactorAPIModelsEnrollmentEnrollmentManagementResponse
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.EnrollmentInstallPFXToCertStore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnrollmentInstallPFXToCertStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **request** | [**KeyfactorApiModelsEnrollmentEnrollmentManagementRequest**](KeyfactorApiModelsEnrollmentEnrollmentManagementRequest.md) | The request to create the management jobs, which includes the request Id of the new pfx and the Ids and management job properties of the cert stores to add the pfx to | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorAPIModelsEnrollmentEnrollmentManagementResponse**](KeyfactorAPIModelsEnrollmentEnrollmentManagementResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnrollmentPostCSREnroll

> ModelsEnrollmentCSREnrollmentResponse EnrollmentPostCSREnroll(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XCertificateformat(xCertificateformat).Request(request).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xCertificateformat := "xCertificateformat_example" // string | Desired format [PEM, DER] (default to "PEM")
    request := *openapiclient.NewModelsEnrollmentCSREnrollmentRequest("CSR_example") // ModelsEnrollmentCSREnrollmentRequest | Information needed to perform the CSR Enrollment
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.EnrollmentPostCSREnroll(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XCertificateformat(xCertificateformat).Request(request).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.EnrollmentPostCSREnroll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnrollmentPostCSREnroll`: ModelsEnrollmentCSREnrollmentResponse
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.EnrollmentPostCSREnroll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnrollmentPostCSREnrollRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xCertificateformat** | **string** | Desired format [PEM, DER] | [default to &quot;PEM&quot;]
 **request** | [**ModelsEnrollmentCSREnrollmentRequest**](ModelsEnrollmentCSREnrollmentRequest.md) | Information needed to perform the CSR Enrollment | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsEnrollmentCSREnrollmentResponse**](ModelsEnrollmentCSREnrollmentResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnrollmentPostPFXEnroll

> ModelsEnrollmentPFXEnrollmentResponse EnrollmentPostPFXEnroll(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XCertificateformat(xCertificateformat).Request(request).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xCertificateformat := "xCertificateformat_example" // string | Desired format [PFX, Zip, STORE] (default to "PFX")
    request := *openapiclient.NewModelsEnrollmentPFXEnrollmentRequest() // ModelsEnrollmentPFXEnrollmentRequest | The information needed to perform the PFX Enrollment
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.EnrollmentPostPFXEnroll(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XCertificateformat(xCertificateformat).Request(request).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.EnrollmentPostPFXEnroll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnrollmentPostPFXEnroll`: ModelsEnrollmentPFXEnrollmentResponse
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.EnrollmentPostPFXEnroll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnrollmentPostPFXEnrollRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xCertificateformat** | **string** | Desired format [PFX, Zip, STORE] | [default to &quot;PFX&quot;]
 **request** | [**ModelsEnrollmentPFXEnrollmentRequest**](ModelsEnrollmentPFXEnrollmentRequest.md) | The information needed to perform the PFX Enrollment | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsEnrollmentPFXEnrollmentResponse**](ModelsEnrollmentPFXEnrollmentResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnrollmentPostParsedCSR

> []string EnrollmentPostParsedCSR(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Csr(csr).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    csr := *openapiclient.NewModelsCSRContents("CSR_example") // ModelsCSRContents | CSR to be parsed
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.EnrollmentPostParsedCSR(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Csr(csr).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.EnrollmentPostParsedCSR``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnrollmentPostParsedCSR`: []string
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.EnrollmentPostParsedCSR`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnrollmentPostParsedCSRRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **csr** | [**ModelsCSRContents**](ModelsCSRContents.md) | CSR to be parsed | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

**[]string**

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnrollmentRenew

> ModelsEnrollmentRenewalResponse EnrollmentRenew(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Request(request).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    request := *openapiclient.NewModelsEnrollmentRenewalRequest() // ModelsEnrollmentRenewalRequest | The information needed to perform the renewal
    collectionId := int64(789) // int64 | The collection id for the given certificate (optional)
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.EnrollmentRenew(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Request(request).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.EnrollmentRenew``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnrollmentRenew`: ModelsEnrollmentRenewalResponse
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.EnrollmentRenew`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnrollmentRenewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **request** | [**ModelsEnrollmentRenewalRequest**](ModelsEnrollmentRenewalRequest.md) | The information needed to perform the renewal | 
 **collectionId** | **int64** | The collection id for the given certificate | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsEnrollmentRenewalResponse**](ModelsEnrollmentRenewalResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


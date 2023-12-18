# \CertificateAuthorityApi

All URIs are relative to */Keyfactor/API*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CertificateAuthorityAlertRecipientsCAHealthRecipientsGet**](CertificateAuthorityApi.md#CertificateAuthorityAlertRecipientsCAHealthRecipientsGet) | **Get** /CertificateAuthority/AlertRecipients/CAHealthRecipients | Returns a list of all CA health recipients
[**CertificateAuthorityAlertRecipientsCAHealthRecipientsIdDelete**](CertificateAuthorityApi.md#CertificateAuthorityAlertRecipientsCAHealthRecipientsIdDelete) | **Delete** /CertificateAuthority/AlertRecipients/CAHealthRecipients/{id} | Deletes a CA health recipient for the provided ID
[**CertificateAuthorityAlertRecipientsCAHealthRecipientsIdGet**](CertificateAuthorityApi.md#CertificateAuthorityAlertRecipientsCAHealthRecipientsIdGet) | **Get** /CertificateAuthority/AlertRecipients/CAHealthRecipients/{id} | Returns a CA health recipient for the specified health recipient ID
[**CertificateAuthorityAlertRecipientsCAHealthRecipientsIdPut**](CertificateAuthorityApi.md#CertificateAuthorityAlertRecipientsCAHealthRecipientsIdPut) | **Put** /CertificateAuthority/AlertRecipients/CAHealthRecipients/{id} | Updates a CA health alert recipient for the provided request object
[**CertificateAuthorityAlertRecipientsCAHealthRecipientsPost**](CertificateAuthorityApi.md#CertificateAuthorityAlertRecipientsCAHealthRecipientsPost) | **Post** /CertificateAuthority/AlertRecipients/CAHealthRecipients | Creates CA health monitoring recipients for the provided list of email addresses\&quot;
[**CertificateAuthorityAlertRecipientsCAThresholdRecipientsGet**](CertificateAuthorityApi.md#CertificateAuthorityAlertRecipientsCAThresholdRecipientsGet) | **Get** /CertificateAuthority/AlertRecipients/CAThresholdRecipients | Returns a list of all CA threshold recipients
[**CertificateAuthorityAlertRecipientsCAThresholdRecipientsIdDelete**](CertificateAuthorityApi.md#CertificateAuthorityAlertRecipientsCAThresholdRecipientsIdDelete) | **Delete** /CertificateAuthority/AlertRecipients/CAThresholdRecipients/{id} | Deletes a CA threshold recipient for the provided ID
[**CertificateAuthorityAlertRecipientsCAThresholdRecipientsIdGet**](CertificateAuthorityApi.md#CertificateAuthorityAlertRecipientsCAThresholdRecipientsIdGet) | **Get** /CertificateAuthority/AlertRecipients/CAThresholdRecipients/{id} | Returns a CA threshold recipient for the specified threshold alert recipient ID
[**CertificateAuthorityAlertRecipientsCAThresholdRecipientsIdPut**](CertificateAuthorityApi.md#CertificateAuthorityAlertRecipientsCAThresholdRecipientsIdPut) | **Put** /CertificateAuthority/AlertRecipients/CAThresholdRecipients/{id} | Updates a CA threshold alert recipient for the provided request object
[**CertificateAuthorityAlertRecipientsCAThresholdRecipientsPost**](CertificateAuthorityApi.md#CertificateAuthorityAlertRecipientsCAThresholdRecipientsPost) | **Post** /CertificateAuthority/AlertRecipients/CAThresholdRecipients | Creates CA threshold alert recipients for the provided list of emails
[**CertificateAuthorityAvailableForestsGet**](CertificateAuthorityApi.md#CertificateAuthorityAvailableForestsGet) | **Get** /CertificateAuthority/AvailableForests | Returns a list of available forests that are in active directory
[**CertificateAuthorityGet**](CertificateAuthorityApi.md#CertificateAuthorityGet) | **Get** /CertificateAuthority | Returns all certificate authorities according to the provided filter
[**CertificateAuthorityHealthMonitoringScheduleGet**](CertificateAuthorityApi.md#CertificateAuthorityHealthMonitoringScheduleGet) | **Get** /CertificateAuthority/HealthMonitoring/Schedule | Retrieves the execution schedule for the CA health monitoring job
[**CertificateAuthorityIdDelete**](CertificateAuthorityApi.md#CertificateAuthorityIdDelete) | **Delete** /CertificateAuthority/{id} | Deletes a CertificateAuthority from the system, specified by ID
[**CertificateAuthorityIdGet**](CertificateAuthorityApi.md#CertificateAuthorityIdGet) | **Get** /CertificateAuthority/{id} | Returns details for a single CA, specified by ID
[**CertificateAuthorityImportPost**](CertificateAuthorityApi.md#CertificateAuthorityImportPost) | **Post** /CertificateAuthority/Import | Imports any certificate authorities from the provided configuration tenant DNS
[**CertificateAuthorityPost**](CertificateAuthorityApi.md#CertificateAuthorityPost) | **Post** /CertificateAuthority | Creates a new CertificateAuthority object
[**CertificateAuthorityPublishCRLPost**](CertificateAuthorityApi.md#CertificateAuthorityPublishCRLPost) | **Post** /CertificateAuthority/PublishCRL | Publishes a CRL according to the provided request
[**CertificateAuthorityPut**](CertificateAuthorityApi.md#CertificateAuthorityPut) | **Put** /CertificateAuthority | Updates a CertificateAuthority object
[**CertificateAuthoritySourceCountGet**](CertificateAuthorityApi.md#CertificateAuthoritySourceCountGet) | **Get** /CertificateAuthority/SourceCount | Returns a count of certificate authorities with sync enabled
[**CertificateAuthorityTestPost**](CertificateAuthorityApi.md#CertificateAuthorityTestPost) | **Post** /CertificateAuthority/Test | Validates the connection info for the CA provided by the model.



## CertificateAuthorityAlertRecipientsCAHealthRecipientsGet

> []KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientResponse CertificateAuthorityAlertRecipientsCAHealthRecipientsGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns a list of all CA health recipients

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
    resp, r, err := apiClient.CertificateAuthorityApi.CertificateAuthorityAlertRecipientsCAHealthRecipientsGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.CertificateAuthorityAlertRecipientsCAHealthRecipientsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateAuthorityAlertRecipientsCAHealthRecipientsGet`: []KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityApi.CertificateAuthorityAlertRecipientsCAHealthRecipientsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateAuthorityAlertRecipientsCAHealthRecipientsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientResponse**](KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateAuthorityAlertRecipientsCAHealthRecipientsIdDelete

> CertificateAuthorityAlertRecipientsCAHealthRecipientsIdDelete(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Deletes a CA health recipient for the provided ID

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
    id := int32(56) // int32 | The ID of the health monitoring recipient to delete
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.CertificateAuthorityAlertRecipientsCAHealthRecipientsIdDelete(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.CertificateAuthorityAlertRecipientsCAHealthRecipientsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the health monitoring recipient to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificateAuthorityAlertRecipientsCAHealthRecipientsIdDeleteRequest struct via the builder pattern


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


## CertificateAuthorityAlertRecipientsCAHealthRecipientsIdGet

> KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientResponse CertificateAuthorityAlertRecipientsCAHealthRecipientsIdGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns a CA health recipient for the specified health recipient ID

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
    id := int32(56) // int32 | The ID of the health monitoring recipient to retrieve
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.CertificateAuthorityAlertRecipientsCAHealthRecipientsIdGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.CertificateAuthorityAlertRecipientsCAHealthRecipientsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateAuthorityAlertRecipientsCAHealthRecipientsIdGet`: KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityApi.CertificateAuthorityAlertRecipientsCAHealthRecipientsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the health monitoring recipient to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificateAuthorityAlertRecipientsCAHealthRecipientsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientResponse**](KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateAuthorityAlertRecipientsCAHealthRecipientsIdPut

> KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientResponse CertificateAuthorityAlertRecipientsCAHealthRecipientsIdPut(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientUpdateRequest(keyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientUpdateRequest).Execute()

Updates a CA health alert recipient for the provided request object

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
    id := int32(56) // int32 | The ID of the alert recipient to be updated
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientUpdateRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientUpdateRequest("Email_example") // KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientUpdateRequest | The request object holding the ID and Email of the health monitoring recipient to be updated (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.CertificateAuthorityAlertRecipientsCAHealthRecipientsIdPut(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientUpdateRequest(keyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.CertificateAuthorityAlertRecipientsCAHealthRecipientsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateAuthorityAlertRecipientsCAHealthRecipientsIdPut`: KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityApi.CertificateAuthorityAlertRecipientsCAHealthRecipientsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the alert recipient to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificateAuthorityAlertRecipientsCAHealthRecipientsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientUpdateRequest** | [**KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientUpdateRequest**](KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientUpdateRequest.md) | The request object holding the ID and Email of the health monitoring recipient to be updated | 

### Return type

[**KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientResponse**](KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateAuthorityAlertRecipientsCAHealthRecipientsPost

> []KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientResponse CertificateAuthorityAlertRecipientsCAHealthRecipientsPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientCreateRequest(keyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientCreateRequest).Execute()

Creates CA health monitoring recipients for the provided list of email addresses\"

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
    keyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientCreateRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientCreateRequest([]string{"Emails_example"}) // KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientCreateRequest | The request object holding the email(s) of the health monitoring recipient(s) to be created (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.CertificateAuthorityAlertRecipientsCAHealthRecipientsPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientCreateRequest(keyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.CertificateAuthorityAlertRecipientsCAHealthRecipientsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateAuthorityAlertRecipientsCAHealthRecipientsPost`: []KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityApi.CertificateAuthorityAlertRecipientsCAHealthRecipientsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateAuthorityAlertRecipientsCAHealthRecipientsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientCreateRequest** | [**KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientCreateRequest**](KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientCreateRequest.md) | The request object holding the email(s) of the health monitoring recipient(s) to be created | 

### Return type

[**[]KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientResponse**](KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateAuthorityAlertRecipientsCAThresholdRecipientsGet

> []KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientResponse CertificateAuthorityAlertRecipientsCAThresholdRecipientsGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns a list of all CA threshold recipients

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
    resp, r, err := apiClient.CertificateAuthorityApi.CertificateAuthorityAlertRecipientsCAThresholdRecipientsGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.CertificateAuthorityAlertRecipientsCAThresholdRecipientsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateAuthorityAlertRecipientsCAThresholdRecipientsGet`: []KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityApi.CertificateAuthorityAlertRecipientsCAThresholdRecipientsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateAuthorityAlertRecipientsCAThresholdRecipientsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientResponse**](KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateAuthorityAlertRecipientsCAThresholdRecipientsIdDelete

> CertificateAuthorityAlertRecipientsCAThresholdRecipientsIdDelete(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Deletes a CA threshold recipient for the provided ID

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
    id := int32(56) // int32 | The ID of the threshold alert recipient to delete
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.CertificateAuthorityAlertRecipientsCAThresholdRecipientsIdDelete(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.CertificateAuthorityAlertRecipientsCAThresholdRecipientsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the threshold alert recipient to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificateAuthorityAlertRecipientsCAThresholdRecipientsIdDeleteRequest struct via the builder pattern


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


## CertificateAuthorityAlertRecipientsCAThresholdRecipientsIdGet

> KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientResponse CertificateAuthorityAlertRecipientsCAThresholdRecipientsIdGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns a CA threshold recipient for the specified threshold alert recipient ID

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
    id := int32(56) // int32 | The ID of the threshold alert recipient to retrieve
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.CertificateAuthorityAlertRecipientsCAThresholdRecipientsIdGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.CertificateAuthorityAlertRecipientsCAThresholdRecipientsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateAuthorityAlertRecipientsCAThresholdRecipientsIdGet`: KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityApi.CertificateAuthorityAlertRecipientsCAThresholdRecipientsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the threshold alert recipient to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificateAuthorityAlertRecipientsCAThresholdRecipientsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientResponse**](KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateAuthorityAlertRecipientsCAThresholdRecipientsIdPut

> KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientResponse CertificateAuthorityAlertRecipientsCAThresholdRecipientsIdPut(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientUpdateRequest(keyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientUpdateRequest).Execute()

Updates a CA threshold alert recipient for the provided request object

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
    id := int32(56) // int32 | The ID of the alert recipient to be updated.
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientUpdateRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientUpdateRequest("Email_example") // KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientUpdateRequest | The request object holding the Email of the alert recipient to be updated (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.CertificateAuthorityAlertRecipientsCAThresholdRecipientsIdPut(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientUpdateRequest(keyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.CertificateAuthorityAlertRecipientsCAThresholdRecipientsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateAuthorityAlertRecipientsCAThresholdRecipientsIdPut`: KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityApi.CertificateAuthorityAlertRecipientsCAThresholdRecipientsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the alert recipient to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificateAuthorityAlertRecipientsCAThresholdRecipientsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientUpdateRequest** | [**KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientUpdateRequest**](KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientUpdateRequest.md) | The request object holding the Email of the alert recipient to be updated | 

### Return type

[**KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientResponse**](KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateAuthorityAlertRecipientsCAThresholdRecipientsPost

> []KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientResponse CertificateAuthorityAlertRecipientsCAThresholdRecipientsPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientCreateRequest(keyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientCreateRequest).Execute()

Creates CA threshold alert recipients for the provided list of emails

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
    keyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientCreateRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientCreateRequest([]string{"Emails_example"}) // KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientCreateRequest | The request object holding the email(s) of the alert recipient(s) to be created (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.CertificateAuthorityAlertRecipientsCAThresholdRecipientsPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientCreateRequest(keyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.CertificateAuthorityAlertRecipientsCAThresholdRecipientsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateAuthorityAlertRecipientsCAThresholdRecipientsPost`: []KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityApi.CertificateAuthorityAlertRecipientsCAThresholdRecipientsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateAuthorityAlertRecipientsCAThresholdRecipientsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientCreateRequest** | [**KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientCreateRequest**](KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientCreateRequest.md) | The request object holding the email(s) of the alert recipient(s) to be created | 

### Return type

[**[]KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientResponse**](KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCAAlertRecipientResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateAuthorityAvailableForestsGet

> []string CertificateAuthorityAvailableForestsGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns a list of available forests that are in active directory

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
    resp, r, err := apiClient.CertificateAuthorityApi.CertificateAuthorityAvailableForestsGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.CertificateAuthorityAvailableForestsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateAuthorityAvailableForestsGet`: []string
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityApi.CertificateAuthorityAvailableForestsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateAuthorityAvailableForestsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

**[]string**

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateAuthorityGet

> []KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse CertificateAuthorityGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns all certificate authorities according to the provided filter

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
    sortAscending := openapiclient.Keyfactor.Common.QueryableExtensionsSortOrder(0) // KeyfactorCommonQueryableExtensionsSortOrder |  (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.CertificateAuthorityGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.CertificateAuthorityGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateAuthorityGet`: []KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityApi.CertificateAuthorityGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateAuthorityGetRequest struct via the builder pattern


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

[**[]KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse**](KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateAuthorityHealthMonitoringScheduleGet

> KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse CertificateAuthorityHealthMonitoringScheduleGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Retrieves the execution schedule for the CA health monitoring job

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
    resp, r, err := apiClient.CertificateAuthorityApi.CertificateAuthorityHealthMonitoringScheduleGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.CertificateAuthorityHealthMonitoringScheduleGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateAuthorityHealthMonitoringScheduleGet`: KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityApi.CertificateAuthorityHealthMonitoringScheduleGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateAuthorityHealthMonitoringScheduleGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse**](KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateAuthorityIdDelete

> CertificateAuthorityIdDelete(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.CertificateAuthorityIdDelete(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.CertificateAuthorityIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCertificateAuthorityIdDeleteRequest struct via the builder pattern


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


## CertificateAuthorityIdGet

> KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse CertificateAuthorityIdGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.CertificateAuthorityIdGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.CertificateAuthorityIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateAuthorityIdGet`: KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityApi.CertificateAuthorityIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificateAuthorityIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse**](KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateAuthorityImportPost

> CertificateAuthorityImportPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Dns(dns).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Imports any certificate authorities from the provided configuration tenant DNS

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
    dns := "dns_example" // string | The DNS of the configuration tenant from which to import certificate authorities (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.CertificateAuthorityImportPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Dns(dns).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.CertificateAuthorityImportPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateAuthorityImportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **dns** | **string** | The DNS of the configuration tenant from which to import certificate authorities | 
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


## CertificateAuthorityPost

> KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse CertificateAuthorityPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).ForceSave(forceSave).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest(keyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    forceSave := true // bool |  (optional) (default to false)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest() // KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest |  (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.CertificateAuthorityPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).ForceSave(forceSave).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest(keyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.CertificateAuthorityPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateAuthorityPost`: KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityApi.CertificateAuthorityPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateAuthorityPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **forceSave** | **bool** |  | [default to false]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest** | [**KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest**](KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest.md) |  | 

### Return type

[**KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse**](KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateAuthorityPublishCRLPost

> CertificateAuthorityPublishCRLPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsCRLRequestModel(cSSCMSDataModelModelsCRLRequestModel).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    cSSCMSDataModelModelsCRLRequestModel := *openapiclient.NewCSSCMSDataModelModelsCRLRequestModel("CertificateAuthorityLogicalName_example") // CSSCMSDataModelModelsCRLRequestModel | Host and logical name of the CA for which the CRL should be published (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.CertificateAuthorityPublishCRLPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsCRLRequestModel(cSSCMSDataModelModelsCRLRequestModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.CertificateAuthorityPublishCRLPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateAuthorityPublishCRLPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsCRLRequestModel** | [**CSSCMSDataModelModelsCRLRequestModel**](CSSCMSDataModelModelsCRLRequestModel.md) | Host and logical name of the CA for which the CRL should be published | 

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


## CertificateAuthorityPut

> KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse CertificateAuthorityPut(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).ForceSave(forceSave).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest(keyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    forceSave := true // bool |  (optional) (default to false)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest() // KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest |  (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.CertificateAuthorityPut(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).ForceSave(forceSave).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest(keyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.CertificateAuthorityPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateAuthorityPut`: KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityApi.CertificateAuthorityPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateAuthorityPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **forceSave** | **bool** |  | [default to false]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest** | [**KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest**](KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityRequest.md) |  | 

### Return type

[**KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse**](KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateAuthoritySourceCountGet

> int32 CertificateAuthoritySourceCountGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns a count of certificate authorities with sync enabled

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
    resp, r, err := apiClient.CertificateAuthorityApi.CertificateAuthoritySourceCountGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.CertificateAuthoritySourceCountGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateAuthoritySourceCountGet`: int32
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityApi.CertificateAuthoritySourceCountGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateAuthoritySourceCountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

**int32**

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateAuthorityTestPost

> KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityTestResponse CertificateAuthorityTestPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest(cSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    cSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest := *openapiclient.NewCSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest() // CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest | The CA being tested. (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.CertificateAuthorityTestPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest(cSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.CertificateAuthorityTestPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateAuthorityTestPost`: KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityTestResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityApi.CertificateAuthorityTestPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateAuthorityTestPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest** | [**CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest**](CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest.md) | The CA being tested. | 

### Return type

[**KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityTestResponse**](KeyfactorWebKeyfactorApiModelsCertificateAuthoritiesCertificateAuthorityTestResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


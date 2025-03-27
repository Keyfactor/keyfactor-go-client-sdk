# \CertificateAuthorityApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCertificateAuthority**](CertificateAuthorityApi.md#CreateCertificateAuthority) | **POST** /CertificateAuthority | Creates a new CertificateAuthority object
[**CreateCertificateAuthorityAlertRecipientsCAHealthRecipients**](CertificateAuthorityApi.md#CreateCertificateAuthorityAlertRecipientsCAHealthRecipients) | **POST** /CertificateAuthority/AlertRecipients/CAHealthRecipients | Creates CA health monitoring recipients for the provided list of email addresses\&quot;
[**CreateCertificateAuthorityAlertRecipientsCAThresholdRecipients**](CertificateAuthorityApi.md#CreateCertificateAuthorityAlertRecipientsCAThresholdRecipients) | **POST** /CertificateAuthority/AlertRecipients/CAThresholdRecipients | Creates CA threshold alert recipients for the provided list of emails
[**CreateCertificateAuthorityImport**](CertificateAuthorityApi.md#CreateCertificateAuthorityImport) | **POST** /CertificateAuthority/Import | Imports any certificate authorities from the provided configuration tenant DNS
[**CreateCertificateAuthorityPublishCRL**](CertificateAuthorityApi.md#CreateCertificateAuthorityPublishCRL) | **POST** /CertificateAuthority/PublishCRL | Publishes a CRL according to the provided request
[**CreateCertificateAuthorityTaskQueueTest**](CertificateAuthorityApi.md#CreateCertificateAuthorityTaskQueueTest) | **POST** /CertificateAuthority/TaskQueue/Test | Tests the connection info for the TaskQueue Credentials.
[**CreateCertificateAuthorityTest**](CertificateAuthorityApi.md#CreateCertificateAuthorityTest) | **POST** /CertificateAuthority/Test | Validates the connection info for the CA provided by the model.
[**DeleteCertificateAuthorityAlertRecipientsCAHealthRecipientsById**](CertificateAuthorityApi.md#DeleteCertificateAuthorityAlertRecipientsCAHealthRecipientsById) | **DELETE** /CertificateAuthority/AlertRecipients/CAHealthRecipients/{id} | Deletes a CA health recipient for the provided ID
[**DeleteCertificateAuthorityAlertRecipientsCAThresholdRecipientsById**](CertificateAuthorityApi.md#DeleteCertificateAuthorityAlertRecipientsCAThresholdRecipientsById) | **DELETE** /CertificateAuthority/AlertRecipients/CAThresholdRecipients/{id} | Deletes a CA threshold recipient for the provided ID
[**DeleteCertificateAuthorityById**](CertificateAuthorityApi.md#DeleteCertificateAuthorityById) | **DELETE** /CertificateAuthority/{id} | Deletes a CertificateAuthority from the system, specified by ID
[**GetCertificateAuthority**](CertificateAuthorityApi.md#GetCertificateAuthority) | **GET** /CertificateAuthority | Returns all certificate authorities according to the provided filter
[**GetCertificateAuthorityAlertRecipientsCAHealthRecipients**](CertificateAuthorityApi.md#GetCertificateAuthorityAlertRecipientsCAHealthRecipients) | **GET** /CertificateAuthority/AlertRecipients/CAHealthRecipients | Returns a list of all CA health recipients
[**GetCertificateAuthorityAlertRecipientsCAHealthRecipientsById**](CertificateAuthorityApi.md#GetCertificateAuthorityAlertRecipientsCAHealthRecipientsById) | **GET** /CertificateAuthority/AlertRecipients/CAHealthRecipients/{id} | Returns a CA health recipient for the specified health recipient ID
[**GetCertificateAuthorityAlertRecipientsCAThresholdRecipients**](CertificateAuthorityApi.md#GetCertificateAuthorityAlertRecipientsCAThresholdRecipients) | **GET** /CertificateAuthority/AlertRecipients/CAThresholdRecipients | Returns a list of all CA threshold recipients
[**GetCertificateAuthorityAlertRecipientsCAThresholdRecipientsById**](CertificateAuthorityApi.md#GetCertificateAuthorityAlertRecipientsCAThresholdRecipientsById) | **GET** /CertificateAuthority/AlertRecipients/CAThresholdRecipients/{id} | Returns a CA threshold recipient for the specified threshold alert recipient ID
[**GetCertificateAuthorityAvailableForests**](CertificateAuthorityApi.md#GetCertificateAuthorityAvailableForests) | **GET** /CertificateAuthority/AvailableForests | Returns a list of available forests that are in active directory
[**GetCertificateAuthorityById**](CertificateAuthorityApi.md#GetCertificateAuthorityById) | **GET** /CertificateAuthority/{id} | Returns details for a single CA, specified by ID
[**GetCertificateAuthorityHealthMonitoringSchedule**](CertificateAuthorityApi.md#GetCertificateAuthorityHealthMonitoringSchedule) | **GET** /CertificateAuthority/HealthMonitoring/Schedule | Retrieves the execution schedule for the CA health monitoring job
[**GetCertificateAuthoritySourceCount**](CertificateAuthorityApi.md#GetCertificateAuthoritySourceCount) | **GET** /CertificateAuthority/SourceCount | Returns a count of certificate authorities with sync enabled
[**GetCertificateAuthorityTaskQueue**](CertificateAuthorityApi.md#GetCertificateAuthorityTaskQueue) | **GET** /CertificateAuthority/TaskQueue | Retrieves credentials and connection information
[**UpdateCertificateAuthority**](CertificateAuthorityApi.md#UpdateCertificateAuthority) | **PUT** /CertificateAuthority | Updates a CertificateAuthority object
[**UpdateCertificateAuthorityAlertRecipientsCAHealthRecipientsById**](CertificateAuthorityApi.md#UpdateCertificateAuthorityAlertRecipientsCAHealthRecipientsById) | **PUT** /CertificateAuthority/AlertRecipients/CAHealthRecipients/{id} | Updates a CA health alert recipient for the provided request object
[**UpdateCertificateAuthorityAlertRecipientsCAThresholdRecipientsById**](CertificateAuthorityApi.md#UpdateCertificateAuthorityAlertRecipientsCAThresholdRecipientsById) | **PUT** /CertificateAuthority/AlertRecipients/CAThresholdRecipients/{id} | Updates a CA threshold alert recipient for the provided request object
[**UpdateCertificateAuthorityTaskQueue**](CertificateAuthorityApi.md#UpdateCertificateAuthorityTaskQueue) | **PUT** /CertificateAuthority/TaskQueue | Updates credentials and connection information



## CreateCertificateAuthority

> CertificateAuthoritiesCertificateAuthorityResponse CreateCertificateAuthority(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).ForceSave(forceSave).XKeyfactorApiVersion(xKeyfactorApiVersion).CertificateAuthoritiesCertificateAuthorityRequest(certificateAuthoritiesCertificateAuthorityRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    forceSave := true // bool |  (optional) (default to false)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    certificateAuthoritiesCertificateAuthorityRequest := *openapiclient.NewCertificateAuthoritiesCertificateAuthorityRequest() // CertificateAuthoritiesCertificateAuthorityRequest |  (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.CreateCertificateAuthority(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).ForceSave(forceSave).XKeyfactorApiVersion(xKeyfactorApiVersion).CertificateAuthoritiesCertificateAuthorityRequest(certificateAuthoritiesCertificateAuthorityRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.CreateCertificateAuthority``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCertificateAuthority`: CertificateAuthoritiesCertificateAuthorityResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityApi.CreateCertificateAuthority`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCertificateAuthorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **forceSave** | **bool** |  | [default to false]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **certificateAuthoritiesCertificateAuthorityRequest** | [**CertificateAuthoritiesCertificateAuthorityRequest**](CertificateAuthoritiesCertificateAuthorityRequest.md) |  | 

### Return type

[**CertificateAuthoritiesCertificateAuthorityResponse**](CertificateAuthoritiesCertificateAuthorityResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCertificateAuthorityAlertRecipientsCAHealthRecipients

> []CertificateAuthoritiesCAAlertRecipientResponse CreateCertificateAuthorityAlertRecipientsCAHealthRecipients(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CertificateAuthoritiesCAAlertRecipientCreateRequest(certificateAuthoritiesCAAlertRecipientCreateRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    certificateAuthoritiesCAAlertRecipientCreateRequest := *openapiclient.NewCertificateAuthoritiesCAAlertRecipientCreateRequest([]string{"Emails_example"}) // CertificateAuthoritiesCAAlertRecipientCreateRequest | The request object holding the email(s) of the health monitoring recipient(s) to be created (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.CreateCertificateAuthorityAlertRecipientsCAHealthRecipients(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CertificateAuthoritiesCAAlertRecipientCreateRequest(certificateAuthoritiesCAAlertRecipientCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.CreateCertificateAuthorityAlertRecipientsCAHealthRecipients``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCertificateAuthorityAlertRecipientsCAHealthRecipients`: []CertificateAuthoritiesCAAlertRecipientResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityApi.CreateCertificateAuthorityAlertRecipientsCAHealthRecipients`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCertificateAuthorityAlertRecipientsCAHealthRecipientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **certificateAuthoritiesCAAlertRecipientCreateRequest** | [**CertificateAuthoritiesCAAlertRecipientCreateRequest**](CertificateAuthoritiesCAAlertRecipientCreateRequest.md) | The request object holding the email(s) of the health monitoring recipient(s) to be created | 

### Return type

[**[]CertificateAuthoritiesCAAlertRecipientResponse**](CertificateAuthoritiesCAAlertRecipientResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCertificateAuthorityAlertRecipientsCAThresholdRecipients

> []CertificateAuthoritiesCAAlertRecipientResponse CreateCertificateAuthorityAlertRecipientsCAThresholdRecipients(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CertificateAuthoritiesCAAlertRecipientCreateRequest(certificateAuthoritiesCAAlertRecipientCreateRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    certificateAuthoritiesCAAlertRecipientCreateRequest := *openapiclient.NewCertificateAuthoritiesCAAlertRecipientCreateRequest([]string{"Emails_example"}) // CertificateAuthoritiesCAAlertRecipientCreateRequest | The request object holding the email(s) of the alert recipient(s) to be created (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.CreateCertificateAuthorityAlertRecipientsCAThresholdRecipients(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CertificateAuthoritiesCAAlertRecipientCreateRequest(certificateAuthoritiesCAAlertRecipientCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.CreateCertificateAuthorityAlertRecipientsCAThresholdRecipients``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCertificateAuthorityAlertRecipientsCAThresholdRecipients`: []CertificateAuthoritiesCAAlertRecipientResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityApi.CreateCertificateAuthorityAlertRecipientsCAThresholdRecipients`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCertificateAuthorityAlertRecipientsCAThresholdRecipientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **certificateAuthoritiesCAAlertRecipientCreateRequest** | [**CertificateAuthoritiesCAAlertRecipientCreateRequest**](CertificateAuthoritiesCAAlertRecipientCreateRequest.md) | The request object holding the email(s) of the alert recipient(s) to be created | 

### Return type

[**[]CertificateAuthoritiesCAAlertRecipientResponse**](CertificateAuthoritiesCAAlertRecipientResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCertificateAuthorityImport

> CreateCertificateAuthorityImport(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Dns(dns).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    dns := "dns_example" // string | The DNS of the configuration tenant from which to import certificate authorities (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.CreateCertificateAuthorityImport(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Dns(dns).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.CreateCertificateAuthorityImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCertificateAuthorityImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
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


## CreateCertificateAuthorityPublishCRL

> CreateCertificateAuthorityPublishCRL(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsCRLRequestModel(cSSCMSDataModelModelsCRLRequestModel).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    cSSCMSDataModelModelsCRLRequestModel := *openapiclient.NewCSSCMSDataModelModelsCRLRequestModel("CertificateAuthorityLogicalName_example") // CSSCMSDataModelModelsCRLRequestModel | Host and logical name of the CA for which the CRL should be published (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.CreateCertificateAuthorityPublishCRL(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsCRLRequestModel(cSSCMSDataModelModelsCRLRequestModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.CreateCertificateAuthorityPublishCRL``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCertificateAuthorityPublishCRLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
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


## CreateCertificateAuthorityTaskQueueTest

> CertificateAuthoritiesCAJobQueueTestResponse CreateCertificateAuthorityTaskQueueTest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).RabbitMQJobQueueRequest(rabbitMQJobQueueRequest).Execute()

Tests the connection info for the TaskQueue Credentials.

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
    rabbitMQJobQueueRequest := *openapiclient.NewRabbitMQJobQueueRequest() // RabbitMQJobQueueRequest | The TaskQueue Credentials being tested. (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.CreateCertificateAuthorityTaskQueueTest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).RabbitMQJobQueueRequest(rabbitMQJobQueueRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.CreateCertificateAuthorityTaskQueueTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCertificateAuthorityTaskQueueTest`: CertificateAuthoritiesCAJobQueueTestResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityApi.CreateCertificateAuthorityTaskQueueTest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCertificateAuthorityTaskQueueTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **rabbitMQJobQueueRequest** | [**RabbitMQJobQueueRequest**](RabbitMQJobQueueRequest.md) | The TaskQueue Credentials being tested. | 

### Return type

[**CertificateAuthoritiesCAJobQueueTestResponse**](CertificateAuthoritiesCAJobQueueTestResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCertificateAuthorityTest

> CertificateAuthoritiesCertificateAuthorityTestResponse CreateCertificateAuthorityTest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest(cSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    cSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest := *openapiclient.NewCSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest() // CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest | The CA being tested. (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.CreateCertificateAuthorityTest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest(cSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.CreateCertificateAuthorityTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCertificateAuthorityTest`: CertificateAuthoritiesCertificateAuthorityTestResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityApi.CreateCertificateAuthorityTest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCertificateAuthorityTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest** | [**CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest**](CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest.md) | The CA being tested. | 

### Return type

[**CertificateAuthoritiesCertificateAuthorityTestResponse**](CertificateAuthoritiesCertificateAuthorityTestResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCertificateAuthorityAlertRecipientsCAHealthRecipientsById

> DeleteCertificateAuthorityAlertRecipientsCAHealthRecipientsById(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.DeleteCertificateAuthorityAlertRecipientsCAHealthRecipientsById(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.DeleteCertificateAuthorityAlertRecipientsCAHealthRecipientsById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteCertificateAuthorityAlertRecipientsCAHealthRecipientsByIdRequest struct via the builder pattern


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


## DeleteCertificateAuthorityAlertRecipientsCAThresholdRecipientsById

> DeleteCertificateAuthorityAlertRecipientsCAThresholdRecipientsById(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.DeleteCertificateAuthorityAlertRecipientsCAThresholdRecipientsById(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.DeleteCertificateAuthorityAlertRecipientsCAThresholdRecipientsById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteCertificateAuthorityAlertRecipientsCAThresholdRecipientsByIdRequest struct via the builder pattern


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


## DeleteCertificateAuthorityById

> DeleteCertificateAuthorityById(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.DeleteCertificateAuthorityById(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.DeleteCertificateAuthorityById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteCertificateAuthorityByIdRequest struct via the builder pattern


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


## GetCertificateAuthority

> []CertificateAuthoritiesCertificateAuthorityResponse GetCertificateAuthority(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    queryString := "queryString_example" // string |  (optional)
    pageReturned := int32(56) // int32 |  (optional)
    returnLimit := int32(56) // int32 |  (optional)
    sortField := "sortField_example" // string |  (optional)
    sortAscending := openapiclient.Keyfactor.Common.QueryableExtensions.SortOrder(0) // KeyfactorCommonQueryableExtensionsSortOrder |  (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.GetCertificateAuthority(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.GetCertificateAuthority``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificateAuthority`: []CertificateAuthoritiesCertificateAuthorityResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityApi.GetCertificateAuthority`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificateAuthorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **queryString** | **string** |  | 
 **pageReturned** | **int32** |  | 
 **returnLimit** | **int32** |  | 
 **sortField** | **string** |  | 
 **sortAscending** | [**KeyfactorCommonQueryableExtensionsSortOrder**](KeyfactorCommonQueryableExtensionsSortOrder.md) |  | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]CertificateAuthoritiesCertificateAuthorityResponse**](CertificateAuthoritiesCertificateAuthorityResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificateAuthorityAlertRecipientsCAHealthRecipients

> []CertificateAuthoritiesCAAlertRecipientResponse GetCertificateAuthorityAlertRecipientsCAHealthRecipients(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.GetCertificateAuthorityAlertRecipientsCAHealthRecipients(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.GetCertificateAuthorityAlertRecipientsCAHealthRecipients``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificateAuthorityAlertRecipientsCAHealthRecipients`: []CertificateAuthoritiesCAAlertRecipientResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityApi.GetCertificateAuthorityAlertRecipientsCAHealthRecipients`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificateAuthorityAlertRecipientsCAHealthRecipientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]CertificateAuthoritiesCAAlertRecipientResponse**](CertificateAuthoritiesCAAlertRecipientResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificateAuthorityAlertRecipientsCAHealthRecipientsById

> CertificateAuthoritiesCAAlertRecipientResponse GetCertificateAuthorityAlertRecipientsCAHealthRecipientsById(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.GetCertificateAuthorityAlertRecipientsCAHealthRecipientsById(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.GetCertificateAuthorityAlertRecipientsCAHealthRecipientsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificateAuthorityAlertRecipientsCAHealthRecipientsById`: CertificateAuthoritiesCAAlertRecipientResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityApi.GetCertificateAuthorityAlertRecipientsCAHealthRecipientsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the health monitoring recipient to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificateAuthorityAlertRecipientsCAHealthRecipientsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**CertificateAuthoritiesCAAlertRecipientResponse**](CertificateAuthoritiesCAAlertRecipientResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificateAuthorityAlertRecipientsCAThresholdRecipients

> []CertificateAuthoritiesCAAlertRecipientResponse GetCertificateAuthorityAlertRecipientsCAThresholdRecipients(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.GetCertificateAuthorityAlertRecipientsCAThresholdRecipients(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.GetCertificateAuthorityAlertRecipientsCAThresholdRecipients``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificateAuthorityAlertRecipientsCAThresholdRecipients`: []CertificateAuthoritiesCAAlertRecipientResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityApi.GetCertificateAuthorityAlertRecipientsCAThresholdRecipients`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificateAuthorityAlertRecipientsCAThresholdRecipientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]CertificateAuthoritiesCAAlertRecipientResponse**](CertificateAuthoritiesCAAlertRecipientResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificateAuthorityAlertRecipientsCAThresholdRecipientsById

> CertificateAuthoritiesCAAlertRecipientResponse GetCertificateAuthorityAlertRecipientsCAThresholdRecipientsById(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.GetCertificateAuthorityAlertRecipientsCAThresholdRecipientsById(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.GetCertificateAuthorityAlertRecipientsCAThresholdRecipientsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificateAuthorityAlertRecipientsCAThresholdRecipientsById`: CertificateAuthoritiesCAAlertRecipientResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityApi.GetCertificateAuthorityAlertRecipientsCAThresholdRecipientsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the threshold alert recipient to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificateAuthorityAlertRecipientsCAThresholdRecipientsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**CertificateAuthoritiesCAAlertRecipientResponse**](CertificateAuthoritiesCAAlertRecipientResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificateAuthorityAvailableForests

> []string GetCertificateAuthorityAvailableForests(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.GetCertificateAuthorityAvailableForests(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.GetCertificateAuthorityAvailableForests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificateAuthorityAvailableForests`: []string
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityApi.GetCertificateAuthorityAvailableForests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificateAuthorityAvailableForestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
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


## GetCertificateAuthorityById

> CertificateAuthoritiesCertificateAuthorityResponse GetCertificateAuthorityById(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.GetCertificateAuthorityById(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.GetCertificateAuthorityById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificateAuthorityById`: CertificateAuthoritiesCertificateAuthorityResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityApi.GetCertificateAuthorityById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificateAuthorityByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**CertificateAuthoritiesCertificateAuthorityResponse**](CertificateAuthoritiesCertificateAuthorityResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificateAuthorityHealthMonitoringSchedule

> SchedulingScheduledTaskResponse GetCertificateAuthorityHealthMonitoringSchedule(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.GetCertificateAuthorityHealthMonitoringSchedule(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.GetCertificateAuthorityHealthMonitoringSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificateAuthorityHealthMonitoringSchedule`: SchedulingScheduledTaskResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityApi.GetCertificateAuthorityHealthMonitoringSchedule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificateAuthorityHealthMonitoringScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**SchedulingScheduledTaskResponse**](SchedulingScheduledTaskResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificateAuthoritySourceCount

> int32 GetCertificateAuthoritySourceCount(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.GetCertificateAuthoritySourceCount(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.GetCertificateAuthoritySourceCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificateAuthoritySourceCount`: int32
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityApi.GetCertificateAuthoritySourceCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificateAuthoritySourceCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
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


## GetCertificateAuthorityTaskQueue

> RabbitMQJobQueueResponse GetCertificateAuthorityTaskQueue(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Retrieves credentials and connection information

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
    resp, r, err := apiClient.CertificateAuthorityApi.GetCertificateAuthorityTaskQueue(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.GetCertificateAuthorityTaskQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificateAuthorityTaskQueue`: RabbitMQJobQueueResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityApi.GetCertificateAuthorityTaskQueue`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificateAuthorityTaskQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**RabbitMQJobQueueResponse**](RabbitMQJobQueueResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCertificateAuthority

> CertificateAuthoritiesCertificateAuthorityResponse UpdateCertificateAuthority(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).ForceSave(forceSave).XKeyfactorApiVersion(xKeyfactorApiVersion).CertificateAuthoritiesCertificateAuthorityRequest(certificateAuthoritiesCertificateAuthorityRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    forceSave := true // bool |  (optional) (default to false)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    certificateAuthoritiesCertificateAuthorityRequest := *openapiclient.NewCertificateAuthoritiesCertificateAuthorityRequest() // CertificateAuthoritiesCertificateAuthorityRequest |  (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.UpdateCertificateAuthority(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).ForceSave(forceSave).XKeyfactorApiVersion(xKeyfactorApiVersion).CertificateAuthoritiesCertificateAuthorityRequest(certificateAuthoritiesCertificateAuthorityRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.UpdateCertificateAuthority``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCertificateAuthority`: CertificateAuthoritiesCertificateAuthorityResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityApi.UpdateCertificateAuthority`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCertificateAuthorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **forceSave** | **bool** |  | [default to false]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **certificateAuthoritiesCertificateAuthorityRequest** | [**CertificateAuthoritiesCertificateAuthorityRequest**](CertificateAuthoritiesCertificateAuthorityRequest.md) |  | 

### Return type

[**CertificateAuthoritiesCertificateAuthorityResponse**](CertificateAuthoritiesCertificateAuthorityResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCertificateAuthorityAlertRecipientsCAHealthRecipientsById

> CertificateAuthoritiesCAAlertRecipientResponse UpdateCertificateAuthorityAlertRecipientsCAHealthRecipientsById(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CertificateAuthoritiesCAAlertRecipientUpdateRequest(certificateAuthoritiesCAAlertRecipientUpdateRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    certificateAuthoritiesCAAlertRecipientUpdateRequest := *openapiclient.NewCertificateAuthoritiesCAAlertRecipientUpdateRequest("Email_example") // CertificateAuthoritiesCAAlertRecipientUpdateRequest | The request object holding the ID and Email of the health monitoring recipient to be updated (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.UpdateCertificateAuthorityAlertRecipientsCAHealthRecipientsById(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CertificateAuthoritiesCAAlertRecipientUpdateRequest(certificateAuthoritiesCAAlertRecipientUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.UpdateCertificateAuthorityAlertRecipientsCAHealthRecipientsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCertificateAuthorityAlertRecipientsCAHealthRecipientsById`: CertificateAuthoritiesCAAlertRecipientResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityApi.UpdateCertificateAuthorityAlertRecipientsCAHealthRecipientsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the alert recipient to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCertificateAuthorityAlertRecipientsCAHealthRecipientsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **certificateAuthoritiesCAAlertRecipientUpdateRequest** | [**CertificateAuthoritiesCAAlertRecipientUpdateRequest**](CertificateAuthoritiesCAAlertRecipientUpdateRequest.md) | The request object holding the ID and Email of the health monitoring recipient to be updated | 

### Return type

[**CertificateAuthoritiesCAAlertRecipientResponse**](CertificateAuthoritiesCAAlertRecipientResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCertificateAuthorityAlertRecipientsCAThresholdRecipientsById

> CertificateAuthoritiesCAAlertRecipientResponse UpdateCertificateAuthorityAlertRecipientsCAThresholdRecipientsById(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CertificateAuthoritiesCAAlertRecipientUpdateRequest(certificateAuthoritiesCAAlertRecipientUpdateRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    certificateAuthoritiesCAAlertRecipientUpdateRequest := *openapiclient.NewCertificateAuthoritiesCAAlertRecipientUpdateRequest("Email_example") // CertificateAuthoritiesCAAlertRecipientUpdateRequest | The request object holding the Email of the alert recipient to be updated (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.UpdateCertificateAuthorityAlertRecipientsCAThresholdRecipientsById(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CertificateAuthoritiesCAAlertRecipientUpdateRequest(certificateAuthoritiesCAAlertRecipientUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.UpdateCertificateAuthorityAlertRecipientsCAThresholdRecipientsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCertificateAuthorityAlertRecipientsCAThresholdRecipientsById`: CertificateAuthoritiesCAAlertRecipientResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityApi.UpdateCertificateAuthorityAlertRecipientsCAThresholdRecipientsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the alert recipient to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCertificateAuthorityAlertRecipientsCAThresholdRecipientsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **certificateAuthoritiesCAAlertRecipientUpdateRequest** | [**CertificateAuthoritiesCAAlertRecipientUpdateRequest**](CertificateAuthoritiesCAAlertRecipientUpdateRequest.md) | The request object holding the Email of the alert recipient to be updated | 

### Return type

[**CertificateAuthoritiesCAAlertRecipientResponse**](CertificateAuthoritiesCAAlertRecipientResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCertificateAuthorityTaskQueue

> RabbitMQJobQueueResponse UpdateCertificateAuthorityTaskQueue(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).RabbitMQJobQueueRequest(rabbitMQJobQueueRequest).Execute()

Updates credentials and connection information

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
    rabbitMQJobQueueRequest := *openapiclient.NewRabbitMQJobQueueRequest() // RabbitMQJobQueueRequest | TaskQueue object with the provided information. (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.UpdateCertificateAuthorityTaskQueue(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).RabbitMQJobQueueRequest(rabbitMQJobQueueRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.UpdateCertificateAuthorityTaskQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCertificateAuthorityTaskQueue`: RabbitMQJobQueueResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityApi.UpdateCertificateAuthorityTaskQueue`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCertificateAuthorityTaskQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **rabbitMQJobQueueRequest** | [**RabbitMQJobQueueRequest**](RabbitMQJobQueueRequest.md) | TaskQueue object with the provided information. | 

### Return type

[**RabbitMQJobQueueResponse**](RabbitMQJobQueueResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \SMTPApi

All URIs are relative to *https://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SMTPSMTP**](SMTPApi.md#SMTPSMTP) | **Get** /SMTP | Gets SMTP profile data
[**SMTPTestSMTP**](SMTPApi.md#SMTPTestSMTP) | **Post** /SMTP/Test | Tests SMTP profile data
[**SMTPUpdateSMTP**](SMTPApi.md#SMTPUpdateSMTP) | **Put** /SMTP | Updates SMTP profile data



## SMTPSMTP

> KeyfactorAPIModelsSMTPSMTPResponse SMTPSMTP(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Gets SMTP profile data

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
    resp, r, err := apiClient.SMTPApi.SMTPSMTP(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMTPApi.SMTPSMTP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SMTPSMTP`: KeyfactorAPIModelsSMTPSMTPResponse
    fmt.Fprintf(os.Stdout, "Response from `SMTPApi.SMTPSMTP`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSMTPSMTPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorAPIModelsSMTPSMTPResponse**](KeyfactorAPIModelsSMTPSMTPResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SMTPTestSMTP

> KeyfactorAPIModelsSMTPSMTPTestResponse SMTPTestSMTP(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).SmtpProfile(smtpProfile).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Tests SMTP profile data

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
    smtpProfile := *openapiclient.NewKeyfactorAPIModelsSMTPSMTPTestRequest() // KeyfactorAPIModelsSMTPSMTPTestRequest | 
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SMTPApi.SMTPTestSMTP(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).SmtpProfile(smtpProfile).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMTPApi.SMTPTestSMTP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SMTPTestSMTP`: KeyfactorAPIModelsSMTPSMTPTestResponse
    fmt.Fprintf(os.Stdout, "Response from `SMTPApi.SMTPTestSMTP`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSMTPTestSMTPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **smtpProfile** | [**KeyfactorAPIModelsSMTPSMTPTestRequest**](KeyfactorAPIModelsSMTPSMTPTestRequest.md) |  | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorAPIModelsSMTPSMTPTestResponse**](KeyfactorAPIModelsSMTPSMTPTestResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SMTPUpdateSMTP

> KeyfactorAPIModelsSMTPSMTPResponse SMTPUpdateSMTP(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).SmtpProfile(smtpProfile).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Updates SMTP profile data

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
    smtpProfile := *openapiclient.NewKeyfactorAPIModelsSMTPSMTPRequest() // KeyfactorAPIModelsSMTPSMTPRequest | 
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SMTPApi.SMTPUpdateSMTP(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).SmtpProfile(smtpProfile).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMTPApi.SMTPUpdateSMTP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SMTPUpdateSMTP`: KeyfactorAPIModelsSMTPSMTPResponse
    fmt.Fprintf(os.Stdout, "Response from `SMTPApi.SMTPUpdateSMTP`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSMTPUpdateSMTPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **smtpProfile** | [**KeyfactorAPIModelsSMTPSMTPRequest**](KeyfactorAPIModelsSMTPSMTPRequest.md) |  | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorAPIModelsSMTPSMTPResponse**](KeyfactorAPIModelsSMTPSMTPResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


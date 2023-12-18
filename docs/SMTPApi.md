# \SMTPApi

All URIs are relative to */Keyfactor/API*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SMTPGet**](SMTPApi.md#SMTPGet) | **Get** /SMTP | Gets SMTP profile data
[**SMTPPut**](SMTPApi.md#SMTPPut) | **Put** /SMTP | Updates SMTP profile data
[**SMTPTestPost**](SMTPApi.md#SMTPTestPost) | **Post** /SMTP/Test | Tests SMTP profile data



## SMTPGet

> KeyfactorWebKeyfactorApiModelsSMTPSMTPResponse SMTPGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SMTPApi.SMTPGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMTPApi.SMTPGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SMTPGet`: KeyfactorWebKeyfactorApiModelsSMTPSMTPResponse
    fmt.Fprintf(os.Stdout, "Response from `SMTPApi.SMTPGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSMTPGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**KeyfactorWebKeyfactorApiModelsSMTPSMTPResponse**](KeyfactorWebKeyfactorApiModelsSMTPSMTPResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SMTPPut

> KeyfactorWebKeyfactorApiModelsSMTPSMTPResponse SMTPPut(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest(keyfactorWebKeyfactorApiModelsSMTPSMTPRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsSMTPSMTPRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsSMTPSMTPRequest() // KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest |  (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SMTPApi.SMTPPut(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest(keyfactorWebKeyfactorApiModelsSMTPSMTPRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMTPApi.SMTPPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SMTPPut`: KeyfactorWebKeyfactorApiModelsSMTPSMTPResponse
    fmt.Fprintf(os.Stdout, "Response from `SMTPApi.SMTPPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSMTPPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsSMTPSMTPRequest** | [**KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest**](KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest.md) |  | 

### Return type

[**KeyfactorWebKeyfactorApiModelsSMTPSMTPResponse**](KeyfactorWebKeyfactorApiModelsSMTPSMTPResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SMTPTestPost

> KeyfactorWebKeyfactorApiModelsSMTPSMTPTestResponse SMTPTestPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsSMTPSMTPTestRequest(keyfactorWebKeyfactorApiModelsSMTPSMTPTestRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsSMTPSMTPTestRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsSMTPSMTPTestRequest() // KeyfactorWebKeyfactorApiModelsSMTPSMTPTestRequest |  (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SMTPApi.SMTPTestPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsSMTPSMTPTestRequest(keyfactorWebKeyfactorApiModelsSMTPSMTPTestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMTPApi.SMTPTestPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SMTPTestPost`: KeyfactorWebKeyfactorApiModelsSMTPSMTPTestResponse
    fmt.Fprintf(os.Stdout, "Response from `SMTPApi.SMTPTestPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSMTPTestPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsSMTPSMTPTestRequest** | [**KeyfactorWebKeyfactorApiModelsSMTPSMTPTestRequest**](KeyfactorWebKeyfactorApiModelsSMTPSMTPTestRequest.md) |  | 

### Return type

[**KeyfactorWebKeyfactorApiModelsSMTPSMTPTestResponse**](KeyfactorWebKeyfactorApiModelsSMTPSMTPTestResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


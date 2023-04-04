# \MacEnrollmentApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MacEnrollmentEditMacEnrollment**](MacEnrollmentApi.md#MacEnrollmentEditMacEnrollment) | **Put** /MacEnrollment | Updates mac enrollment settings data
[**MacEnrollmentMacEnrollment**](MacEnrollmentApi.md#MacEnrollmentMacEnrollment) | **Get** /MacEnrollment | Gets mac enrollment settings data



## MacEnrollmentEditMacEnrollment

> KeyfactorApiModelsMacEnrollmentMacEnrollmentAPIModel MacEnrollmentEditMacEnrollment(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).MacEnrollmentSettings(macEnrollmentSettings).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Updates mac enrollment settings data

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
    macEnrollmentSettings := *openapiclient.NewKeyfactorApiModelsMacEnrollmentMacEnrollmentAPIModel() // KeyfactorApiModelsMacEnrollmentMacEnrollmentAPIModel | 
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MacEnrollmentApi.MacEnrollmentEditMacEnrollment(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).MacEnrollmentSettings(macEnrollmentSettings).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MacEnrollmentApi.MacEnrollmentEditMacEnrollment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MacEnrollmentEditMacEnrollment`: KeyfactorApiModelsMacEnrollmentMacEnrollmentAPIModel
    fmt.Fprintf(os.Stdout, "Response from `MacEnrollmentApi.MacEnrollmentEditMacEnrollment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMacEnrollmentEditMacEnrollmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **macEnrollmentSettings** | [**KeyfactorApiModelsMacEnrollmentMacEnrollmentAPIModel**](KeyfactorApiModelsMacEnrollmentMacEnrollmentAPIModel.md) |  | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsMacEnrollmentMacEnrollmentAPIModel**](KeyfactorApiModelsMacEnrollmentMacEnrollmentAPIModel.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MacEnrollmentMacEnrollment

> KeyfactorApiModelsMacEnrollmentMacEnrollmentAPIModel MacEnrollmentMacEnrollment(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Gets mac enrollment settings data

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MacEnrollmentApi.MacEnrollmentMacEnrollment(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MacEnrollmentApi.MacEnrollmentMacEnrollment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MacEnrollmentMacEnrollment`: KeyfactorApiModelsMacEnrollmentMacEnrollmentAPIModel
    fmt.Fprintf(os.Stdout, "Response from `MacEnrollmentApi.MacEnrollmentMacEnrollment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMacEnrollmentMacEnrollmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsMacEnrollmentMacEnrollmentAPIModel**](KeyfactorApiModelsMacEnrollmentMacEnrollmentAPIModel.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


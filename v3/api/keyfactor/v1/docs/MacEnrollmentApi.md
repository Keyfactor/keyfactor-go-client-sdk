# \MacEnrollmentApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMacEnrollment**](MacEnrollmentApi.md#GetMacEnrollment) | **GET** /MacEnrollment | Gets mac enrollment settings data
[**UpdateMacEnrollment**](MacEnrollmentApi.md#UpdateMacEnrollment) | **PUT** /MacEnrollment | Updates mac enrollment settings data



## GetMacEnrollment

> MacEnrollmentMacEnrollmentAPIModel GetMacEnrollment(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MacEnrollmentApi.GetMacEnrollment(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MacEnrollmentApi.GetMacEnrollment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMacEnrollment`: MacEnrollmentMacEnrollmentAPIModel
    fmt.Fprintf(os.Stdout, "Response from `MacEnrollmentApi.GetMacEnrollment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMacEnrollmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**MacEnrollmentMacEnrollmentAPIModel**](MacEnrollmentMacEnrollmentAPIModel.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMacEnrollment

> MacEnrollmentMacEnrollmentAPIModel UpdateMacEnrollment(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).MacEnrollmentMacEnrollmentAPIModel(macEnrollmentMacEnrollmentAPIModel).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    macEnrollmentMacEnrollmentAPIModel := *openapiclient.NewMacEnrollmentMacEnrollmentAPIModel() // MacEnrollmentMacEnrollmentAPIModel |  (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MacEnrollmentApi.UpdateMacEnrollment(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).MacEnrollmentMacEnrollmentAPIModel(macEnrollmentMacEnrollmentAPIModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MacEnrollmentApi.UpdateMacEnrollment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMacEnrollment`: MacEnrollmentMacEnrollmentAPIModel
    fmt.Fprintf(os.Stdout, "Response from `MacEnrollmentApi.UpdateMacEnrollment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMacEnrollmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **macEnrollmentMacEnrollmentAPIModel** | [**MacEnrollmentMacEnrollmentAPIModel**](MacEnrollmentMacEnrollmentAPIModel.md) |  | 

### Return type

[**MacEnrollmentMacEnrollmentAPIModel**](MacEnrollmentMacEnrollmentAPIModel.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


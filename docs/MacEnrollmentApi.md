# \MacEnrollmentApi

All URIs are relative to */Keyfactor/API*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MacEnrollmentGet**](MacEnrollmentApi.md#MacEnrollmentGet) | **Get** /MacEnrollment | Gets mac enrollment settings data
[**MacEnrollmentPut**](MacEnrollmentApi.md#MacEnrollmentPut) | **Put** /MacEnrollment | Updates mac enrollment settings data



## MacEnrollmentGet

> KeyfactorWebKeyfactorApiModelsMacEnrollmentMacEnrollmentAPIModel MacEnrollmentGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MacEnrollmentApi.MacEnrollmentGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MacEnrollmentApi.MacEnrollmentGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MacEnrollmentGet`: KeyfactorWebKeyfactorApiModelsMacEnrollmentMacEnrollmentAPIModel
    fmt.Fprintf(os.Stdout, "Response from `MacEnrollmentApi.MacEnrollmentGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMacEnrollmentGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**KeyfactorWebKeyfactorApiModelsMacEnrollmentMacEnrollmentAPIModel**](KeyfactorWebKeyfactorApiModelsMacEnrollmentMacEnrollmentAPIModel.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MacEnrollmentPut

> KeyfactorWebKeyfactorApiModelsMacEnrollmentMacEnrollmentAPIModel MacEnrollmentPut(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsMacEnrollmentMacEnrollmentAPIModel(keyfactorWebKeyfactorApiModelsMacEnrollmentMacEnrollmentAPIModel).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsMacEnrollmentMacEnrollmentAPIModel := *openapiclient.NewKeyfactorWebKeyfactorApiModelsMacEnrollmentMacEnrollmentAPIModel() // KeyfactorWebKeyfactorApiModelsMacEnrollmentMacEnrollmentAPIModel |  (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MacEnrollmentApi.MacEnrollmentPut(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsMacEnrollmentMacEnrollmentAPIModel(keyfactorWebKeyfactorApiModelsMacEnrollmentMacEnrollmentAPIModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MacEnrollmentApi.MacEnrollmentPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MacEnrollmentPut`: KeyfactorWebKeyfactorApiModelsMacEnrollmentMacEnrollmentAPIModel
    fmt.Fprintf(os.Stdout, "Response from `MacEnrollmentApi.MacEnrollmentPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMacEnrollmentPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsMacEnrollmentMacEnrollmentAPIModel** | [**KeyfactorWebKeyfactorApiModelsMacEnrollmentMacEnrollmentAPIModel**](KeyfactorWebKeyfactorApiModelsMacEnrollmentMacEnrollmentAPIModel.md) |  | 

### Return type

[**KeyfactorWebKeyfactorApiModelsMacEnrollmentMacEnrollmentAPIModel**](KeyfactorWebKeyfactorApiModelsMacEnrollmentMacEnrollmentAPIModel.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \AppSettingApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAppSetting**](AppSettingApi.md#GetAppSetting) | **GET** /AppSetting | Get available application settings
[**GetAppSettingById**](AppSettingApi.md#GetAppSettingById) | **GET** /AppSetting/{id} | Get application setting by id
[**UpdateAppSetting**](AppSettingApi.md#UpdateAppSetting) | **PUT** /AppSetting | Bulk update available application settings
[**UpdateAppSettingByIdSet**](AppSettingApi.md#UpdateAppSettingByIdSet) | **PUT** /AppSetting/{id}/Set | Update one application setting by id
[**UpdateAppSettingNameSet**](AppSettingApi.md#UpdateAppSettingNameSet) | **PUT** /AppSetting/{name}/Set | Update one application setting by short name



## GetAppSetting

> []AppSettingsAppSettingResponse GetAppSetting(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Get available application settings

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
    resp, r, err := apiClient.AppSettingApi.GetAppSetting(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppSettingApi.GetAppSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAppSetting`: []AppSettingsAppSettingResponse
    fmt.Fprintf(os.Stdout, "Response from `AppSettingApi.GetAppSetting`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAppSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]AppSettingsAppSettingResponse**](AppSettingsAppSettingResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppSettingById

> AppSettingsAppSettingResponse GetAppSettingById(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Get application setting by id

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
    id := int32(56) // int32 | Id for the application setting
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppSettingApi.GetAppSettingById(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppSettingApi.GetAppSettingById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAppSettingById`: AppSettingsAppSettingResponse
    fmt.Fprintf(os.Stdout, "Response from `AppSettingApi.GetAppSettingById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id for the application setting | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppSettingByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**AppSettingsAppSettingResponse**](AppSettingsAppSettingResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppSetting

> []AppSettingsAppSettingResponse UpdateAppSetting(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).AppSettingsAppSettingUpdateBulkRequest(appSettingsAppSettingUpdateBulkRequest).Execute()

Bulk update available application settings

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
    appSettingsAppSettingUpdateBulkRequest := []openapiclient.AppSettingsAppSettingUpdateBulkRequest{*openapiclient.NewAppSettingsAppSettingUpdateBulkRequest()} // []AppSettingsAppSettingUpdateBulkRequest | List of information to update application settings (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppSettingApi.UpdateAppSetting(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).AppSettingsAppSettingUpdateBulkRequest(appSettingsAppSettingUpdateBulkRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppSettingApi.UpdateAppSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAppSetting`: []AppSettingsAppSettingResponse
    fmt.Fprintf(os.Stdout, "Response from `AppSettingApi.UpdateAppSetting`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **appSettingsAppSettingUpdateBulkRequest** | [**[]AppSettingsAppSettingUpdateBulkRequest**](AppSettingsAppSettingUpdateBulkRequest.md) | List of information to update application settings | 

### Return type

[**[]AppSettingsAppSettingResponse**](AppSettingsAppSettingResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppSettingByIdSet

> AppSettingsAppSettingResponse UpdateAppSettingByIdSet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).AppSettingsAppSettingUpdateRequest(appSettingsAppSettingUpdateRequest).Execute()

Update one application setting by id

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
    id := int32(56) // int32 | Id for the application setting
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    appSettingsAppSettingUpdateRequest := *openapiclient.NewAppSettingsAppSettingUpdateRequest() // AppSettingsAppSettingUpdateRequest | Infomation for updating the application setting (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppSettingApi.UpdateAppSettingByIdSet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).AppSettingsAppSettingUpdateRequest(appSettingsAppSettingUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppSettingApi.UpdateAppSettingByIdSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAppSettingByIdSet`: AppSettingsAppSettingResponse
    fmt.Fprintf(os.Stdout, "Response from `AppSettingApi.UpdateAppSettingByIdSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id for the application setting | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppSettingByIdSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **appSettingsAppSettingUpdateRequest** | [**AppSettingsAppSettingUpdateRequest**](AppSettingsAppSettingUpdateRequest.md) | Infomation for updating the application setting | 

### Return type

[**AppSettingsAppSettingResponse**](AppSettingsAppSettingResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppSettingNameSet

> AppSettingsAppSettingResponse UpdateAppSettingNameSet(ctx, name).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).AppSettingsAppSettingUpdateRequest(appSettingsAppSettingUpdateRequest).Execute()

Update one application setting by short name

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
    name := "name_example" // string | Short name for the application setting
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    appSettingsAppSettingUpdateRequest := *openapiclient.NewAppSettingsAppSettingUpdateRequest() // AppSettingsAppSettingUpdateRequest | Infomation for updating the application setting (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppSettingApi.UpdateAppSettingNameSet(context.Background(), name).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).AppSettingsAppSettingUpdateRequest(appSettingsAppSettingUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppSettingApi.UpdateAppSettingNameSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAppSettingNameSet`: AppSettingsAppSettingResponse
    fmt.Fprintf(os.Stdout, "Response from `AppSettingApi.UpdateAppSettingNameSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Short name for the application setting | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppSettingNameSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **appSettingsAppSettingUpdateRequest** | [**AppSettingsAppSettingUpdateRequest**](AppSettingsAppSettingUpdateRequest.md) | Infomation for updating the application setting | 

### Return type

[**AppSettingsAppSettingResponse**](AppSettingsAppSettingResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


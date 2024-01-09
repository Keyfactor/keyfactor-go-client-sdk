# \AppSettingApi

All URIs are relative to */Keyfactor/API*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppSettingGet**](AppSettingApi.md#AppSettingGet) | **Get** /AppSetting | Get available application settings
[**AppSettingIdGet**](AppSettingApi.md#AppSettingIdGet) | **Get** /AppSetting/{id} | Get application setting by id
[**AppSettingIdSetPut**](AppSettingApi.md#AppSettingIdSetPut) | **Put** /AppSetting/{id}/Set | Update one application setting by id
[**AppSettingNameSetPut**](AppSettingApi.md#AppSettingNameSetPut) | **Put** /AppSetting/{name}/Set | Update one application setting by short name
[**AppSettingPut**](AppSettingApi.md#AppSettingPut) | **Put** /AppSetting | Bulk update available application settings



## AppSettingGet

> []KeyfactorWebKeyfactorApiModelsAppSettingsAppSettingResponse AppSettingGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppSettingApi.AppSettingGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppSettingApi.AppSettingGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppSettingGet`: []KeyfactorWebKeyfactorApiModelsAppSettingsAppSettingResponse
    fmt.Fprintf(os.Stdout, "Response from `AppSettingApi.AppSettingGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppSettingGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]KeyfactorWebKeyfactorApiModelsAppSettingsAppSettingResponse**](KeyfactorWebKeyfactorApiModelsAppSettingsAppSettingResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppSettingIdGet

> KeyfactorWebKeyfactorApiModelsAppSettingsAppSettingResponse AppSettingIdGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppSettingApi.AppSettingIdGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppSettingApi.AppSettingIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppSettingIdGet`: KeyfactorWebKeyfactorApiModelsAppSettingsAppSettingResponse
    fmt.Fprintf(os.Stdout, "Response from `AppSettingApi.AppSettingIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id for the application setting | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppSettingIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**KeyfactorWebKeyfactorApiModelsAppSettingsAppSettingResponse**](KeyfactorWebKeyfactorApiModelsAppSettingsAppSettingResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppSettingIdSetPut

> KeyfactorWebKeyfactorApiModelsAppSettingsAppSettingResponse AppSettingIdSetPut(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsAppSettingsAppSettingUpdateRequest(keyfactorWebKeyfactorApiModelsAppSettingsAppSettingUpdateRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsAppSettingsAppSettingUpdateRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsAppSettingsAppSettingUpdateRequest() // KeyfactorWebKeyfactorApiModelsAppSettingsAppSettingUpdateRequest | Infomation for updating the application setting (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppSettingApi.AppSettingIdSetPut(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsAppSettingsAppSettingUpdateRequest(keyfactorWebKeyfactorApiModelsAppSettingsAppSettingUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppSettingApi.AppSettingIdSetPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppSettingIdSetPut`: KeyfactorWebKeyfactorApiModelsAppSettingsAppSettingResponse
    fmt.Fprintf(os.Stdout, "Response from `AppSettingApi.AppSettingIdSetPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id for the application setting | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppSettingIdSetPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsAppSettingsAppSettingUpdateRequest** | [**KeyfactorWebKeyfactorApiModelsAppSettingsAppSettingUpdateRequest**](KeyfactorWebKeyfactorApiModelsAppSettingsAppSettingUpdateRequest.md) | Infomation for updating the application setting | 

### Return type

[**KeyfactorWebKeyfactorApiModelsAppSettingsAppSettingResponse**](KeyfactorWebKeyfactorApiModelsAppSettingsAppSettingResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppSettingNameSetPut

> KeyfactorWebKeyfactorApiModelsAppSettingsAppSettingResponse AppSettingNameSetPut(ctx, name).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsAppSettingsAppSettingUpdateRequest(keyfactorWebKeyfactorApiModelsAppSettingsAppSettingUpdateRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsAppSettingsAppSettingUpdateRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsAppSettingsAppSettingUpdateRequest() // KeyfactorWebKeyfactorApiModelsAppSettingsAppSettingUpdateRequest | Infomation for updating the application setting (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppSettingApi.AppSettingNameSetPut(context.Background(), name).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsAppSettingsAppSettingUpdateRequest(keyfactorWebKeyfactorApiModelsAppSettingsAppSettingUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppSettingApi.AppSettingNameSetPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppSettingNameSetPut`: KeyfactorWebKeyfactorApiModelsAppSettingsAppSettingResponse
    fmt.Fprintf(os.Stdout, "Response from `AppSettingApi.AppSettingNameSetPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Short name for the application setting | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppSettingNameSetPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsAppSettingsAppSettingUpdateRequest** | [**KeyfactorWebKeyfactorApiModelsAppSettingsAppSettingUpdateRequest**](KeyfactorWebKeyfactorApiModelsAppSettingsAppSettingUpdateRequest.md) | Infomation for updating the application setting | 

### Return type

[**KeyfactorWebKeyfactorApiModelsAppSettingsAppSettingResponse**](KeyfactorWebKeyfactorApiModelsAppSettingsAppSettingResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppSettingPut

> []KeyfactorWebKeyfactorApiModelsAppSettingsAppSettingResponse AppSettingPut(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsAppSettingsAppSettingUpdateBulkRequest(keyfactorWebKeyfactorApiModelsAppSettingsAppSettingUpdateBulkRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsAppSettingsAppSettingUpdateBulkRequest := []openapiclient.KeyfactorWebKeyfactorApiModelsAppSettingsAppSettingUpdateBulkRequest{*openapiclient.NewKeyfactorWebKeyfactorApiModelsAppSettingsAppSettingUpdateBulkRequest()} // []KeyfactorWebKeyfactorApiModelsAppSettingsAppSettingUpdateBulkRequest | List of information to update application settings (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppSettingApi.AppSettingPut(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsAppSettingsAppSettingUpdateBulkRequest(keyfactorWebKeyfactorApiModelsAppSettingsAppSettingUpdateBulkRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppSettingApi.AppSettingPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppSettingPut`: []KeyfactorWebKeyfactorApiModelsAppSettingsAppSettingResponse
    fmt.Fprintf(os.Stdout, "Response from `AppSettingApi.AppSettingPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppSettingPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsAppSettingsAppSettingUpdateBulkRequest** | [**[]KeyfactorWebKeyfactorApiModelsAppSettingsAppSettingUpdateBulkRequest**](KeyfactorWebKeyfactorApiModelsAppSettingsAppSettingUpdateBulkRequest.md) | List of information to update application settings | 

### Return type

[**[]KeyfactorWebKeyfactorApiModelsAppSettingsAppSettingResponse**](KeyfactorWebKeyfactorApiModelsAppSettingsAppSettingResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \PermissionSetApi

All URIs are relative to */Keyfactor/API*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PermissionSetsGet**](PermissionSetApi.md#PermissionSetsGet) | **Get** /PermissionSets | Gets all permission sets in the system.
[**PermissionSetsIdDelete**](PermissionSetApi.md#PermissionSetsIdDelete) | **Delete** /PermissionSets/{id} | Deletes a permission set.
[**PermissionSetsIdGet**](PermissionSetApi.md#PermissionSetsIdGet) | **Get** /PermissionSets/{id} | Gets permission set data.
[**PermissionSetsPost**](PermissionSetApi.md#PermissionSetsPost) | **Post** /PermissionSets | Creates a new permission set.
[**PermissionSetsPut**](PermissionSetApi.md#PermissionSetsPut) | **Put** /PermissionSets | 



## PermissionSetsGet

> []KeyfactorWebKeyfactorApiModelsPermissionSetsPermissionSetResponse PermissionSetsGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Gets all permission sets in the system.

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
    sortAscending := int32(56) // int32 |  (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionSetApi.PermissionSetsGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionSetApi.PermissionSetsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PermissionSetsGet`: []KeyfactorWebKeyfactorApiModelsPermissionSetsPermissionSetResponse
    fmt.Fprintf(os.Stdout, "Response from `PermissionSetApi.PermissionSetsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPermissionSetsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **queryString** | **string** |  | 
 **pageReturned** | **int32** |  | 
 **returnLimit** | **int32** |  | 
 **sortField** | **string** |  | 
 **sortAscending** | **int32** |  | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]KeyfactorWebKeyfactorApiModelsPermissionSetsPermissionSetResponse**](KeyfactorWebKeyfactorApiModelsPermissionSetsPermissionSetResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PermissionSetsIdDelete

> PermissionSetsIdDelete(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Deletes a permission set.

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of the permission set to delete.
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionSetApi.PermissionSetsIdDelete(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionSetApi.PermissionSetsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the permission set to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionSetsIdDeleteRequest struct via the builder pattern


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


## PermissionSetsIdGet

> KeyfactorWebKeyfactorApiModelsPermissionSetsPermissionSetResponse PermissionSetsIdGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Gets permission set data.



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionSetApi.PermissionSetsIdGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionSetApi.PermissionSetsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PermissionSetsIdGet`: KeyfactorWebKeyfactorApiModelsPermissionSetsPermissionSetResponse
    fmt.Fprintf(os.Stdout, "Response from `PermissionSetApi.PermissionSetsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionSetsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**KeyfactorWebKeyfactorApiModelsPermissionSetsPermissionSetResponse**](KeyfactorWebKeyfactorApiModelsPermissionSetsPermissionSetResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PermissionSetsPost

> KeyfactorWebKeyfactorApiModelsPermissionSetsPermissionSetResponse PermissionSetsPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsPermissionSetsPermissionSetCreateRequest(keyfactorWebKeyfactorApiModelsPermissionSetsPermissionSetCreateRequest).Execute()

Creates a new permission set.

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
    keyfactorWebKeyfactorApiModelsPermissionSetsPermissionSetCreateRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsPermissionSetsPermissionSetCreateRequest("Name_example", []string{"Permissions_example"}) // KeyfactorWebKeyfactorApiModelsPermissionSetsPermissionSetCreateRequest | Information about the new permission set. (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionSetApi.PermissionSetsPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsPermissionSetsPermissionSetCreateRequest(keyfactorWebKeyfactorApiModelsPermissionSetsPermissionSetCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionSetApi.PermissionSetsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PermissionSetsPost`: KeyfactorWebKeyfactorApiModelsPermissionSetsPermissionSetResponse
    fmt.Fprintf(os.Stdout, "Response from `PermissionSetApi.PermissionSetsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPermissionSetsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsPermissionSetsPermissionSetCreateRequest** | [**KeyfactorWebKeyfactorApiModelsPermissionSetsPermissionSetCreateRequest**](KeyfactorWebKeyfactorApiModelsPermissionSetsPermissionSetCreateRequest.md) | Information about the new permission set. | 

### Return type

[**KeyfactorWebKeyfactorApiModelsPermissionSetsPermissionSetResponse**](KeyfactorWebKeyfactorApiModelsPermissionSetsPermissionSetResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PermissionSetsPut

> KeyfactorWebKeyfactorApiModelsPermissionSetsPermissionSetResponse PermissionSetsPut(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).KeyfactorWebKeyfactorApiModelsPermissionSetsPermissionSetUpdateRequest(keyfactorWebKeyfactorApiModelsPermissionSetsPermissionSetUpdateRequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()



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
    keyfactorWebKeyfactorApiModelsPermissionSetsPermissionSetUpdateRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsPermissionSetsPermissionSetUpdateRequest("Id_example", []string{"Permissions_example"}) // KeyfactorWebKeyfactorApiModelsPermissionSetsPermissionSetUpdateRequest | 
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionSetApi.PermissionSetsPut(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).KeyfactorWebKeyfactorApiModelsPermissionSetsPermissionSetUpdateRequest(keyfactorWebKeyfactorApiModelsPermissionSetsPermissionSetUpdateRequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionSetApi.PermissionSetsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PermissionSetsPut`: KeyfactorWebKeyfactorApiModelsPermissionSetsPermissionSetResponse
    fmt.Fprintf(os.Stdout, "Response from `PermissionSetApi.PermissionSetsPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPermissionSetsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **keyfactorWebKeyfactorApiModelsPermissionSetsPermissionSetUpdateRequest** | [**KeyfactorWebKeyfactorApiModelsPermissionSetsPermissionSetUpdateRequest**](KeyfactorWebKeyfactorApiModelsPermissionSetsPermissionSetUpdateRequest.md) |  | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**KeyfactorWebKeyfactorApiModelsPermissionSetsPermissionSetResponse**](KeyfactorWebKeyfactorApiModelsPermissionSetsPermissionSetResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


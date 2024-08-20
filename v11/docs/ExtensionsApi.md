# \ExtensionsApi

All URIs are relative to */Keyfactor/API*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExtensionsScriptsGet**](ExtensionsApi.md#ExtensionsScriptsGet) | **Get** /Extensions/Scripts | Returns all scripts according to the provided filter and output parameters
[**ExtensionsScriptsIdDelete**](ExtensionsApi.md#ExtensionsScriptsIdDelete) | **Delete** /Extensions/Scripts/{id} | Deletes a script. Script cannot be configured to an alert or workflow.
[**ExtensionsScriptsIdGet**](ExtensionsApi.md#ExtensionsScriptsIdGet) | **Get** /Extensions/Scripts/{id} | Returns a single script that matches the provided Id
[**ExtensionsScriptsPost**](ExtensionsApi.md#ExtensionsScriptsPost) | **Post** /Extensions/Scripts | Adds a new script
[**ExtensionsScriptsPut**](ExtensionsApi.md#ExtensionsScriptsPut) | **Put** /Extensions/Scripts | Updates a script



## ExtensionsScriptsGet

> []KeyfactorWebKeyfactorApiModelsScriptsScriptQueryResponse ExtensionsScriptsGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns all scripts according to the provided filter and output parameters

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
    resp, r, err := apiClient.ExtensionsApi.ExtensionsScriptsGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExtensionsApi.ExtensionsScriptsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExtensionsScriptsGet`: []KeyfactorWebKeyfactorApiModelsScriptsScriptQueryResponse
    fmt.Fprintf(os.Stdout, "Response from `ExtensionsApi.ExtensionsScriptsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExtensionsScriptsGetRequest struct via the builder pattern


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

[**[]KeyfactorWebKeyfactorApiModelsScriptsScriptQueryResponse**](KeyfactorWebKeyfactorApiModelsScriptsScriptQueryResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExtensionsScriptsIdDelete

> ExtensionsScriptsIdDelete(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Deletes a script. Script cannot be configured to an alert or workflow.

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
    id := int32(56) // int32 | Id of the script to delete
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExtensionsApi.ExtensionsScriptsIdDelete(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExtensionsApi.ExtensionsScriptsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id of the script to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiExtensionsScriptsIdDeleteRequest struct via the builder pattern


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


## ExtensionsScriptsIdGet

> KeyfactorWebKeyfactorApiModelsScriptsScriptResponse ExtensionsScriptsIdGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns a single script that matches the provided Id

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
    id := int32(56) // int32 | Id of the script
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExtensionsApi.ExtensionsScriptsIdGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExtensionsApi.ExtensionsScriptsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExtensionsScriptsIdGet`: KeyfactorWebKeyfactorApiModelsScriptsScriptResponse
    fmt.Fprintf(os.Stdout, "Response from `ExtensionsApi.ExtensionsScriptsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id of the script | 

### Other Parameters

Other parameters are passed through a pointer to a apiExtensionsScriptsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**KeyfactorWebKeyfactorApiModelsScriptsScriptResponse**](KeyfactorWebKeyfactorApiModelsScriptsScriptResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExtensionsScriptsPost

> KeyfactorWebKeyfactorApiModelsScriptsScriptResponse ExtensionsScriptsPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).KeyfactorWebKeyfactorApiModelsScriptsScriptCreateRequest(keyfactorWebKeyfactorApiModelsScriptsScriptCreateRequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Adds a new script

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
    keyfactorWebKeyfactorApiModelsScriptsScriptCreateRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsScriptsScriptCreateRequest("Name_example", "Contents_example") // KeyfactorWebKeyfactorApiModelsScriptsScriptCreateRequest | Script parameters
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExtensionsApi.ExtensionsScriptsPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).KeyfactorWebKeyfactorApiModelsScriptsScriptCreateRequest(keyfactorWebKeyfactorApiModelsScriptsScriptCreateRequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExtensionsApi.ExtensionsScriptsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExtensionsScriptsPost`: KeyfactorWebKeyfactorApiModelsScriptsScriptResponse
    fmt.Fprintf(os.Stdout, "Response from `ExtensionsApi.ExtensionsScriptsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExtensionsScriptsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **keyfactorWebKeyfactorApiModelsScriptsScriptCreateRequest** | [**KeyfactorWebKeyfactorApiModelsScriptsScriptCreateRequest**](KeyfactorWebKeyfactorApiModelsScriptsScriptCreateRequest.md) | Script parameters | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**KeyfactorWebKeyfactorApiModelsScriptsScriptResponse**](KeyfactorWebKeyfactorApiModelsScriptsScriptResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExtensionsScriptsPut

> KeyfactorWebKeyfactorApiModelsScriptsScriptResponse ExtensionsScriptsPut(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).KeyfactorWebKeyfactorApiModelsScriptsScriptsUpdateRequest(keyfactorWebKeyfactorApiModelsScriptsScriptsUpdateRequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Updates a script

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
    keyfactorWebKeyfactorApiModelsScriptsScriptsUpdateRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsScriptsScriptsUpdateRequest() // KeyfactorWebKeyfactorApiModelsScriptsScriptsUpdateRequest | Script parameters
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExtensionsApi.ExtensionsScriptsPut(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).KeyfactorWebKeyfactorApiModelsScriptsScriptsUpdateRequest(keyfactorWebKeyfactorApiModelsScriptsScriptsUpdateRequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExtensionsApi.ExtensionsScriptsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExtensionsScriptsPut`: KeyfactorWebKeyfactorApiModelsScriptsScriptResponse
    fmt.Fprintf(os.Stdout, "Response from `ExtensionsApi.ExtensionsScriptsPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExtensionsScriptsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **keyfactorWebKeyfactorApiModelsScriptsScriptsUpdateRequest** | [**KeyfactorWebKeyfactorApiModelsScriptsScriptsUpdateRequest**](KeyfactorWebKeyfactorApiModelsScriptsScriptsUpdateRequest.md) | Script parameters | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**KeyfactorWebKeyfactorApiModelsScriptsScriptResponse**](KeyfactorWebKeyfactorApiModelsScriptsScriptResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


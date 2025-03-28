# \EventHandlerRegistrationApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEventHandlerRegistration**](EventHandlerRegistrationApi.md#CreateEventHandlerRegistration) | **POST** /EventHandlerRegistration | Registers an event handler
[**DeleteEventHandlerRegistrationById**](EventHandlerRegistrationApi.md#DeleteEventHandlerRegistrationById) | **DELETE** /EventHandlerRegistration/{id} | Deletes an event handler
[**GetEventHandlerRegistration**](EventHandlerRegistrationApi.md#GetEventHandlerRegistration) | **GET** /EventHandlerRegistration | Returns all registered event handlers according to the provided filter and output parameters
[**GetEventHandlerRegistrationById**](EventHandlerRegistrationApi.md#GetEventHandlerRegistrationById) | **GET** /EventHandlerRegistration/{id} | Returns a registered event handler that matches the provided ID
[**UpdateEventHandlerRegistrationById**](EventHandlerRegistrationApi.md#UpdateEventHandlerRegistrationById) | **PUT** /EventHandlerRegistration/{id} | Updates a registered event handler&#39;s information



## CreateEventHandlerRegistration

> []EventHandlerRegistrationEventHandlerRegistrationResponse NewCreateEventHandlerRegistrationRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).EventHandlerRegistrationEventHandlerRegistrationCreateRequest(eventHandlerRegistrationEventHandlerRegistrationCreateRequest).Execute()

Registers an event handler

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
    eventHandlerRegistrationEventHandlerRegistrationCreateRequest := *openapiclient.NewEventHandlerRegistrationEventHandlerRegistrationCreateRequest("AssemblyName_example") // EventHandlerRegistrationEventHandlerRegistrationCreateRequest | The assembly name of the event handler to register. The handler file must be in the configured extensions directory on the machine running the management portal. (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventHandlerRegistrationApi.NewCreateEventHandlerRegistrationRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).EventHandlerRegistrationEventHandlerRegistrationCreateRequest(eventHandlerRegistrationEventHandlerRegistrationCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventHandlerRegistrationApi.CreateEventHandlerRegistration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEventHandlerRegistration`: []EventHandlerRegistrationEventHandlerRegistrationResponse
    fmt.Fprintf(os.Stdout, "Response from `EventHandlerRegistrationApi.CreateEventHandlerRegistration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEventHandlerRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **eventHandlerRegistrationEventHandlerRegistrationCreateRequest** | [**EventHandlerRegistrationEventHandlerRegistrationCreateRequest**](EventHandlerRegistrationEventHandlerRegistrationCreateRequest.md) | The assembly name of the event handler to register. The handler file must be in the configured extensions directory on the machine running the management portal. | 

### Return type

[**[]EventHandlerRegistrationEventHandlerRegistrationResponse**](EventHandlerRegistrationEventHandlerRegistrationResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEventHandlerRegistrationById

> NewDeleteEventHandlerRegistrationByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Deletes an event handler

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
    id := int32(56) // int32 | Id of the event handler
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventHandlerRegistrationApi.NewDeleteEventHandlerRegistrationByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventHandlerRegistrationApi.DeleteEventHandlerRegistrationById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id of the event handler | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEventHandlerRegistrationByIdRequest struct via the builder pattern


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


## GetEventHandlerRegistration

> []EventHandlerRegistrationEventHandlerRegistrationResponse NewGetEventHandlerRegistrationRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns all registered event handlers according to the provided filter and output parameters

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
    resp, r, err := apiClient.EventHandlerRegistrationApi.NewGetEventHandlerRegistrationRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventHandlerRegistrationApi.GetEventHandlerRegistration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventHandlerRegistration`: []EventHandlerRegistrationEventHandlerRegistrationResponse
    fmt.Fprintf(os.Stdout, "Response from `EventHandlerRegistrationApi.GetEventHandlerRegistration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEventHandlerRegistrationRequest struct via the builder pattern


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

[**[]EventHandlerRegistrationEventHandlerRegistrationResponse**](EventHandlerRegistrationEventHandlerRegistrationResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventHandlerRegistrationById

> EventHandlerRegistrationEventHandlerRegistrationResponse NewGetEventHandlerRegistrationByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns a registered event handler that matches the provided ID

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
    id := int32(56) // int32 | Id of the event handler
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventHandlerRegistrationApi.NewGetEventHandlerRegistrationByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventHandlerRegistrationApi.GetEventHandlerRegistrationById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventHandlerRegistrationById`: EventHandlerRegistrationEventHandlerRegistrationResponse
    fmt.Fprintf(os.Stdout, "Response from `EventHandlerRegistrationApi.GetEventHandlerRegistrationById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id of the event handler | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventHandlerRegistrationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**EventHandlerRegistrationEventHandlerRegistrationResponse**](EventHandlerRegistrationEventHandlerRegistrationResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEventHandlerRegistrationById

> EventHandlerRegistrationEventHandlerRegistrationResponse NewUpdateEventHandlerRegistrationByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).EventHandlerRegistrationEventHandlerRegistrationUpdateRequest(eventHandlerRegistrationEventHandlerRegistrationUpdateRequest).Execute()

Updates a registered event handler's information

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
    id := int32(56) // int32 | Id of the event handler
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    eventHandlerRegistrationEventHandlerRegistrationUpdateRequest := *openapiclient.NewEventHandlerRegistrationEventHandlerRegistrationUpdateRequest() // EventHandlerRegistrationEventHandlerRegistrationUpdateRequest | Updated information for the event handler (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventHandlerRegistrationApi.NewUpdateEventHandlerRegistrationByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).EventHandlerRegistrationEventHandlerRegistrationUpdateRequest(eventHandlerRegistrationEventHandlerRegistrationUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventHandlerRegistrationApi.UpdateEventHandlerRegistrationById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEventHandlerRegistrationById`: EventHandlerRegistrationEventHandlerRegistrationResponse
    fmt.Fprintf(os.Stdout, "Response from `EventHandlerRegistrationApi.UpdateEventHandlerRegistrationById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id of the event handler | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEventHandlerRegistrationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **eventHandlerRegistrationEventHandlerRegistrationUpdateRequest** | [**EventHandlerRegistrationEventHandlerRegistrationUpdateRequest**](EventHandlerRegistrationEventHandlerRegistrationUpdateRequest.md) | Updated information for the event handler | 

### Return type

[**EventHandlerRegistrationEventHandlerRegistrationResponse**](EventHandlerRegistrationEventHandlerRegistrationResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


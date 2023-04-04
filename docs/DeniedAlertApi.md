# \DeniedAlertApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeniedAlertAddDeniedAlert**](DeniedAlertApi.md#DeniedAlertAddDeniedAlert) | **Post** /Alerts/Denied | Add a denied alert
[**DeniedAlertDeleteDeniedAlert**](DeniedAlertApi.md#DeniedAlertDeleteDeniedAlert) | **Delete** /Alerts/Denied/{id} | Delete a denied alert
[**DeniedAlertEditDeniedAlert**](DeniedAlertApi.md#DeniedAlertEditDeniedAlert) | **Put** /Alerts/Denied | Edit a denied alert
[**DeniedAlertGetDeniedAlert**](DeniedAlertApi.md#DeniedAlertGetDeniedAlert) | **Get** /Alerts/Denied/{id} | Get a denied alert
[**DeniedAlertGetDeniedAlerts**](DeniedAlertApi.md#DeniedAlertGetDeniedAlerts) | **Get** /Alerts/Denied | Gets all denied alerts according to the provided filter and output parameters



## DeniedAlertAddDeniedAlert

> KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse DeniedAlertAddDeniedAlert(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Req(req).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Add a denied alert

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
    req := *openapiclient.NewKeyfactorApiModelsAlertsDeniedDeniedAlertCreationRequest("DisplayName_example", "Subject_example", "Message_example") // KeyfactorApiModelsAlertsDeniedDeniedAlertCreationRequest | Information for the new alert
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeniedAlertApi.DeniedAlertAddDeniedAlert(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Req(req).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeniedAlertApi.DeniedAlertAddDeniedAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeniedAlertAddDeniedAlert`: KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `DeniedAlertApi.DeniedAlertAddDeniedAlert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeniedAlertAddDeniedAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **req** | [**KeyfactorApiModelsAlertsDeniedDeniedAlertCreationRequest**](KeyfactorApiModelsAlertsDeniedDeniedAlertCreationRequest.md) | Information for the new alert | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse**](KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeniedAlertDeleteDeniedAlert

> DeniedAlertDeleteDeniedAlert(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Delete a denied alert

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
    id := int32(56) // int32 | Id for the denied alert
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeniedAlertApi.DeniedAlertDeleteDeniedAlert(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeniedAlertApi.DeniedAlertDeleteDeniedAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id for the denied alert | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeniedAlertDeleteDeniedAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeniedAlertEditDeniedAlert

> KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse DeniedAlertEditDeniedAlert(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Req(req).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Edit a denied alert

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
    req := *openapiclient.NewKeyfactorApiModelsAlertsDeniedDeniedAlertUpdateRequest("DisplayName_example", "Subject_example", "Message_example") // KeyfactorApiModelsAlertsDeniedDeniedAlertUpdateRequest | Information for the denied alert
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeniedAlertApi.DeniedAlertEditDeniedAlert(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Req(req).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeniedAlertApi.DeniedAlertEditDeniedAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeniedAlertEditDeniedAlert`: KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `DeniedAlertApi.DeniedAlertEditDeniedAlert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeniedAlertEditDeniedAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **req** | [**KeyfactorApiModelsAlertsDeniedDeniedAlertUpdateRequest**](KeyfactorApiModelsAlertsDeniedDeniedAlertUpdateRequest.md) | Information for the denied alert | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse**](KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeniedAlertGetDeniedAlert

> KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse DeniedAlertGetDeniedAlert(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Get a denied alert

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
    id := int32(56) // int32 | Id for the denied alert to get
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeniedAlertApi.DeniedAlertGetDeniedAlert(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeniedAlertApi.DeniedAlertGetDeniedAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeniedAlertGetDeniedAlert`: KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `DeniedAlertApi.DeniedAlertGetDeniedAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id for the denied alert to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeniedAlertGetDeniedAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse**](KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeniedAlertGetDeniedAlerts

> []KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse DeniedAlertGetDeniedAlerts(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PagedQueryQueryString(pagedQueryQueryString).PagedQueryPageReturned(pagedQueryPageReturned).PagedQueryReturnLimit(pagedQueryReturnLimit).PagedQuerySortField(pagedQuerySortField).PagedQuerySortAscending(pagedQuerySortAscending).Execute()

Gets all denied alerts according to the provided filter and output parameters

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
    pagedQueryQueryString := "pagedQueryQueryString_example" // string | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) (optional)
    pagedQueryPageReturned := int32(56) // int32 | The current page within the result set to be returned (optional)
    pagedQueryReturnLimit := int32(56) // int32 | Maximum number of records to be returned in a single call (optional)
    pagedQuerySortField := "pagedQuerySortField_example" // string | Field by which the results should be sorted (view results via Management Portal for sortable columns) (optional)
    pagedQuerySortAscending := int32(56) // int32 | Field sort direction [0=ascending, 1=descending] (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeniedAlertApi.DeniedAlertGetDeniedAlerts(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PagedQueryQueryString(pagedQueryQueryString).PagedQueryPageReturned(pagedQueryPageReturned).PagedQueryReturnLimit(pagedQueryReturnLimit).PagedQuerySortField(pagedQuerySortField).PagedQuerySortAscending(pagedQuerySortAscending).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeniedAlertApi.DeniedAlertGetDeniedAlerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeniedAlertGetDeniedAlerts`: []KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `DeniedAlertApi.DeniedAlertGetDeniedAlerts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeniedAlertGetDeniedAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]
 **pagedQueryQueryString** | **string** | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) | 
 **pagedQueryPageReturned** | **int32** | The current page within the result set to be returned | 
 **pagedQueryReturnLimit** | **int32** | Maximum number of records to be returned in a single call | 
 **pagedQuerySortField** | **string** | Field by which the results should be sorted (view results via Management Portal for sortable columns) | 
 **pagedQuerySortAscending** | **int32** | Field sort direction [0&#x3D;ascending, 1&#x3D;descending] | 

### Return type

[**[]KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse**](KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


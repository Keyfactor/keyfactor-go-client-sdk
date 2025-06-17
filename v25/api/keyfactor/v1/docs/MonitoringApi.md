# \MonitoringApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMonitoringResolveOCSP**](MonitoringApi.md#CreateMonitoringResolveOCSP) | **POST** /Monitoring/ResolveOCSP | Resolve the Certificate authority given
[**CreateMonitoringRevocation**](MonitoringApi.md#CreateMonitoringRevocation) | **POST** /Monitoring/Revocation | Add a revocation monitoring endpoint
[**CreateMonitoringRevocationCRLTest**](MonitoringApi.md#CreateMonitoringRevocationCRLTest) | **POST** /Monitoring/Revocation/CRL/Test | Validates the connection info for the CRL provided by the model.
[**CreateMonitoringRevocationOCSPTest**](MonitoringApi.md#CreateMonitoringRevocationOCSPTest) | **POST** /Monitoring/Revocation/OCSP/Test | Validates the connection info for the OCSP endpoint provided by the model.
[**CreateMonitoringRevocationTest**](MonitoringApi.md#CreateMonitoringRevocationTest) | **POST** /Monitoring/Revocation/Test | Test Alert
[**CreateMonitoringRevocationTestAll**](MonitoringApi.md#CreateMonitoringRevocationTestAll) | **POST** /Monitoring/Revocation/TestAll | Test All Alerts
[**DeleteMonitoringRevocationById**](MonitoringApi.md#DeleteMonitoringRevocationById) | **DELETE** /Monitoring/Revocation/{id} | Delete a revocation monitoring endpoint
[**GetMonitoringRevocation**](MonitoringApi.md#GetMonitoringRevocation) | **GET** /Monitoring/Revocation | Gets all revocation monitoring endpoints according to the provided filter and output parameters
[**GetMonitoringRevocationById**](MonitoringApi.md#GetMonitoringRevocationById) | **GET** /Monitoring/Revocation/{id} | Get a revocation monitoring endpoint
[**UpdateMonitoringRevocation**](MonitoringApi.md#UpdateMonitoringRevocation) | **PUT** /Monitoring/Revocation | Edit a revocation monitoring endpoint
[**UpdateMonitoringRevocationSchedule**](MonitoringApi.md#UpdateMonitoringRevocationSchedule) | **PUT** /Monitoring/Revocation/Schedule | Edit a revocation monitoring&#39;s schedule.



## CreateMonitoringResolveOCSP

> MonitoringOCSPParametersResponse NewCreateMonitoringResolveOCSPRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).MonitoringOCSPParametersRequest(monitoringOCSPParametersRequest).Execute()

Resolve the Certificate authority given

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
    monitoringOCSPParametersRequest := *openapiclient.NewMonitoringOCSPParametersRequest() // MonitoringOCSPParametersRequest | Information for the new endpoint (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringApi.NewCreateMonitoringResolveOCSPRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).MonitoringOCSPParametersRequest(monitoringOCSPParametersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringApi.CreateMonitoringResolveOCSP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMonitoringResolveOCSP`: MonitoringOCSPParametersResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitoringApi.CreateMonitoringResolveOCSP`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMonitoringResolveOCSPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **monitoringOCSPParametersRequest** | [**MonitoringOCSPParametersRequest**](MonitoringOCSPParametersRequest.md) | Information for the new endpoint | 

### Return type

[**MonitoringOCSPParametersResponse**](MonitoringOCSPParametersResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMonitoringRevocation

> MonitoringRevocationMonitoringDefinitionResponse NewCreateMonitoringRevocationRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).ForceSave(forceSave).XKeyfactorApiVersion(xKeyfactorApiVersion).MonitoringRevocationMonitoringCreationRequest(monitoringRevocationMonitoringCreationRequest).Execute()

Add a revocation monitoring endpoint

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
    forceSave := true // bool | Bypass testing the connection to either the CRL or the OCSP endpoint (optional) (default to false)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    monitoringRevocationMonitoringCreationRequest := *openapiclient.NewMonitoringRevocationMonitoringCreationRequest("Name_example", "EndpointType_example", "Location_example", *openapiclient.NewMonitoringDashboardRequest(false)) // MonitoringRevocationMonitoringCreationRequest | Information for the new endpoint (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringApi.NewCreateMonitoringRevocationRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).ForceSave(forceSave).XKeyfactorApiVersion(xKeyfactorApiVersion).MonitoringRevocationMonitoringCreationRequest(monitoringRevocationMonitoringCreationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringApi.CreateMonitoringRevocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMonitoringRevocation`: MonitoringRevocationMonitoringDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitoringApi.CreateMonitoringRevocation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMonitoringRevocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **forceSave** | **bool** | Bypass testing the connection to either the CRL or the OCSP endpoint | [default to false]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **monitoringRevocationMonitoringCreationRequest** | [**MonitoringRevocationMonitoringCreationRequest**](MonitoringRevocationMonitoringCreationRequest.md) | Information for the new endpoint | 

### Return type

[**MonitoringRevocationMonitoringDefinitionResponse**](MonitoringRevocationMonitoringDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMonitoringRevocationCRLTest

> CSSCMSDataModelModelsMonitoringCRLTestResponse NewCreateMonitoringRevocationCRLTestRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsMonitoringCRLTestRequest(cSSCMSDataModelModelsMonitoringCRLTestRequest).Execute()

Validates the connection info for the CRL provided by the model.

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
    cSSCMSDataModelModelsMonitoringCRLTestRequest := *openapiclient.NewCSSCMSDataModelModelsMonitoringCRLTestRequest() // CSSCMSDataModelModelsMonitoringCRLTestRequest | The CRL being tested. (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringApi.NewCreateMonitoringRevocationCRLTestRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsMonitoringCRLTestRequest(cSSCMSDataModelModelsMonitoringCRLTestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringApi.CreateMonitoringRevocationCRLTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMonitoringRevocationCRLTest`: CSSCMSDataModelModelsMonitoringCRLTestResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitoringApi.CreateMonitoringRevocationCRLTest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMonitoringRevocationCRLTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsMonitoringCRLTestRequest** | [**CSSCMSDataModelModelsMonitoringCRLTestRequest**](CSSCMSDataModelModelsMonitoringCRLTestRequest.md) | The CRL being tested. | 

### Return type

[**CSSCMSDataModelModelsMonitoringCRLTestResponse**](CSSCMSDataModelModelsMonitoringCRLTestResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMonitoringRevocationOCSPTest

> CSSCMSDataModelModelsMonitoringOCSPTestResponse NewCreateMonitoringRevocationOCSPTestRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsMonitoringOCSPTestRequest(cSSCMSDataModelModelsMonitoringOCSPTestRequest).Execute()

Validates the connection info for the OCSP endpoint provided by the model.

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
    cSSCMSDataModelModelsMonitoringOCSPTestRequest := *openapiclient.NewCSSCMSDataModelModelsMonitoringOCSPTestRequest() // CSSCMSDataModelModelsMonitoringOCSPTestRequest | The OCSP endpoint being tested. (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringApi.NewCreateMonitoringRevocationOCSPTestRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsMonitoringOCSPTestRequest(cSSCMSDataModelModelsMonitoringOCSPTestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringApi.CreateMonitoringRevocationOCSPTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMonitoringRevocationOCSPTest`: CSSCMSDataModelModelsMonitoringOCSPTestResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitoringApi.CreateMonitoringRevocationOCSPTest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMonitoringRevocationOCSPTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsMonitoringOCSPTestRequest** | [**CSSCMSDataModelModelsMonitoringOCSPTestRequest**](CSSCMSDataModelModelsMonitoringOCSPTestRequest.md) | The OCSP endpoint being tested. | 

### Return type

[**CSSCMSDataModelModelsMonitoringOCSPTestResponse**](CSSCMSDataModelModelsMonitoringOCSPTestResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMonitoringRevocationTest

> MonitoringRevocationMonitoringAlertTestResponse NewCreateMonitoringRevocationTestRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).MonitoringRevocationMonitoringAlertTestRequest(monitoringRevocationMonitoringAlertTestRequest).Execute()

Test Alert

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
    monitoringRevocationMonitoringAlertTestRequest := *openapiclient.NewMonitoringRevocationMonitoringAlertTestRequest() // MonitoringRevocationMonitoringAlertTestRequest | Information about the revocation monitoring alert test (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringApi.NewCreateMonitoringRevocationTestRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).MonitoringRevocationMonitoringAlertTestRequest(monitoringRevocationMonitoringAlertTestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringApi.CreateMonitoringRevocationTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMonitoringRevocationTest`: MonitoringRevocationMonitoringAlertTestResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitoringApi.CreateMonitoringRevocationTest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMonitoringRevocationTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **monitoringRevocationMonitoringAlertTestRequest** | [**MonitoringRevocationMonitoringAlertTestRequest**](MonitoringRevocationMonitoringAlertTestRequest.md) | Information about the revocation monitoring alert test | 

### Return type

[**MonitoringRevocationMonitoringAlertTestResponse**](MonitoringRevocationMonitoringAlertTestResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMonitoringRevocationTestAll

> MonitoringRevocationMonitoringAlertTestResponse NewCreateMonitoringRevocationTestAllRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).MonitoringRevocationMonitoringAlertTestAllRequest(monitoringRevocationMonitoringAlertTestAllRequest).Execute()

Test All Alerts

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
    monitoringRevocationMonitoringAlertTestAllRequest := *openapiclient.NewMonitoringRevocationMonitoringAlertTestAllRequest() // MonitoringRevocationMonitoringAlertTestAllRequest | Information about the revocation monitoring alert test (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringApi.NewCreateMonitoringRevocationTestAllRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).MonitoringRevocationMonitoringAlertTestAllRequest(monitoringRevocationMonitoringAlertTestAllRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringApi.CreateMonitoringRevocationTestAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMonitoringRevocationTestAll`: MonitoringRevocationMonitoringAlertTestResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitoringApi.CreateMonitoringRevocationTestAll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMonitoringRevocationTestAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **monitoringRevocationMonitoringAlertTestAllRequest** | [**MonitoringRevocationMonitoringAlertTestAllRequest**](MonitoringRevocationMonitoringAlertTestAllRequest.md) | Information about the revocation monitoring alert test | 

### Return type

[**MonitoringRevocationMonitoringAlertTestResponse**](MonitoringRevocationMonitoringAlertTestResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMonitoringRevocationById

> NewDeleteMonitoringRevocationByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Delete a revocation monitoring endpoint

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
    id := int32(56) // int32 | Id for the revocation monitoring endpoint
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringApi.NewDeleteMonitoringRevocationByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringApi.DeleteMonitoringRevocationById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id for the revocation monitoring endpoint | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMonitoringRevocationByIdRequest struct via the builder pattern


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


## GetMonitoringRevocation

> []MonitoringRevocationMonitoringDefinitionResponse NewGetMonitoringRevocationRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Gets all revocation monitoring endpoints according to the provided filter and output parameters

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
    resp, r, err := apiClient.MonitoringApi.NewGetMonitoringRevocationRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringApi.GetMonitoringRevocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonitoringRevocation`: []MonitoringRevocationMonitoringDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitoringApi.GetMonitoringRevocation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitoringRevocationRequest struct via the builder pattern


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

[**[]MonitoringRevocationMonitoringDefinitionResponse**](MonitoringRevocationMonitoringDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonitoringRevocationById

> MonitoringRevocationMonitoringDefinitionResponse NewGetMonitoringRevocationByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Get a revocation monitoring endpoint

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
    id := int32(56) // int32 | Id for the endpoint to get
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringApi.NewGetMonitoringRevocationByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringApi.GetMonitoringRevocationById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonitoringRevocationById`: MonitoringRevocationMonitoringDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitoringApi.GetMonitoringRevocationById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id for the endpoint to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitoringRevocationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**MonitoringRevocationMonitoringDefinitionResponse**](MonitoringRevocationMonitoringDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMonitoringRevocation

> MonitoringRevocationMonitoringDefinitionResponse NewUpdateMonitoringRevocationRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).ForceSave(forceSave).XKeyfactorApiVersion(xKeyfactorApiVersion).MonitoringRevocationMonitoringUpdateRequest(monitoringRevocationMonitoringUpdateRequest).Execute()

Edit a revocation monitoring endpoint

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
    forceSave := true // bool | Bypass testing the connection to either the CRL or the OCSP endpoint (optional) (default to false)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    monitoringRevocationMonitoringUpdateRequest := *openapiclient.NewMonitoringRevocationMonitoringUpdateRequest("Name_example", "EndpointType_example", "Location_example", *openapiclient.NewMonitoringDashboardRequest(false)) // MonitoringRevocationMonitoringUpdateRequest | Information for the endpoint (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringApi.NewUpdateMonitoringRevocationRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).ForceSave(forceSave).XKeyfactorApiVersion(xKeyfactorApiVersion).MonitoringRevocationMonitoringUpdateRequest(monitoringRevocationMonitoringUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringApi.UpdateMonitoringRevocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMonitoringRevocation`: MonitoringRevocationMonitoringDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitoringApi.UpdateMonitoringRevocation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMonitoringRevocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **forceSave** | **bool** | Bypass testing the connection to either the CRL or the OCSP endpoint | [default to false]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **monitoringRevocationMonitoringUpdateRequest** | [**MonitoringRevocationMonitoringUpdateRequest**](MonitoringRevocationMonitoringUpdateRequest.md) | Information for the endpoint | 

### Return type

[**MonitoringRevocationMonitoringDefinitionResponse**](MonitoringRevocationMonitoringDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMonitoringRevocationSchedule

> MonitoringRevocationMonitoringDefinitionResponse NewUpdateMonitoringRevocationScheduleRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).MonitoringRevocationMonitoringUpdateScheduleRequest(monitoringRevocationMonitoringUpdateScheduleRequest).Execute()

Edit a revocation monitoring's schedule.

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
    monitoringRevocationMonitoringUpdateScheduleRequest := *openapiclient.NewMonitoringRevocationMonitoringUpdateScheduleRequest(int32(123)) // MonitoringRevocationMonitoringUpdateScheduleRequest | The information for the updating the schedule. (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringApi.NewUpdateMonitoringRevocationScheduleRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).MonitoringRevocationMonitoringUpdateScheduleRequest(monitoringRevocationMonitoringUpdateScheduleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringApi.UpdateMonitoringRevocationSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMonitoringRevocationSchedule`: MonitoringRevocationMonitoringDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitoringApi.UpdateMonitoringRevocationSchedule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMonitoringRevocationScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **monitoringRevocationMonitoringUpdateScheduleRequest** | [**MonitoringRevocationMonitoringUpdateScheduleRequest**](MonitoringRevocationMonitoringUpdateScheduleRequest.md) | The information for the updating the schedule. | 

### Return type

[**MonitoringRevocationMonitoringDefinitionResponse**](MonitoringRevocationMonitoringDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \LogonApi

All URIs are relative to *https://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LogonCreateLogon**](LogonApi.md#LogonCreateLogon) | **Post** /SSH/Logons | Creates a logon with the provided properties
[**LogonDelete**](LogonApi.md#LogonDelete) | **Delete** /SSH/Logons/{id} | Deletes a Logon associated with the provided identifier
[**LogonGetLogon**](LogonApi.md#LogonGetLogon) | **Get** /SSH/Logons/{id} | Fetches a Logon associated with the provided identifier
[**LogonLogonAccess**](LogonApi.md#LogonLogonAccess) | **Post** /SSH/Logons/Access | Updates the users with access to an existing logon
[**LogonQueryLogons**](LogonApi.md#LogonQueryLogons) | **Get** /SSH/Logons | Returns all Logons according to the provided filter parameters



## LogonCreateLogon

> ModelsSSHLogonsLogonResponse LogonCreateLogon(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Logon(logon).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Creates a logon with the provided properties

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
    logon := *openapiclient.NewModelsSSHLogonsLogonCreationRequest("Username_example", int32(123)) // ModelsSSHLogonsLogonCreationRequest | Logon properties to be applied to the new logon
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogonApi.LogonCreateLogon(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Logon(logon).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogonApi.LogonCreateLogon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LogonCreateLogon`: ModelsSSHLogonsLogonResponse
    fmt.Fprintf(os.Stdout, "Response from `LogonApi.LogonCreateLogon`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLogonCreateLogonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **logon** | [**ModelsSSHLogonsLogonCreationRequest**](ModelsSSHLogonsLogonCreationRequest.md) | Logon properties to be applied to the new logon | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsSSHLogonsLogonResponse**](ModelsSSHLogonsLogonResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LogonDelete

> LogonDelete(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Deletes a Logon associated with the provided identifier

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
    id := int32(56) // int32 | Keyfactor identifer of the Logon to be deleted
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogonApi.LogonDelete(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogonApi.LogonDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Keyfactor identifer of the Logon to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiLogonDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

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


## LogonGetLogon

> ModelsSSHLogonsLogonResponse LogonGetLogon(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Fetches a Logon associated with the provided identifier

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
    id := int32(56) // int32 | Keyfactor identifer of the Logon to be Fetched
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogonApi.LogonGetLogon(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogonApi.LogonGetLogon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LogonGetLogon`: ModelsSSHLogonsLogonResponse
    fmt.Fprintf(os.Stdout, "Response from `LogonApi.LogonGetLogon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Keyfactor identifer of the Logon to be Fetched | 

### Other Parameters

Other parameters are passed through a pointer to a apiLogonGetLogonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsSSHLogonsLogonResponse**](ModelsSSHLogonsLogonResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LogonLogonAccess

> ModelsSSHAccessLogonUserAccessResponse LogonLogonAccess(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Logon(logon).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Updates the users with access to an existing logon

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
    logon := *openapiclient.NewModelsSSHLogonsLogonAccessRequest(int32(123)) // ModelsSSHLogonsLogonAccessRequest | Users to add the existing logon
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogonApi.LogonLogonAccess(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Logon(logon).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogonApi.LogonLogonAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LogonLogonAccess`: ModelsSSHAccessLogonUserAccessResponse
    fmt.Fprintf(os.Stdout, "Response from `LogonApi.LogonLogonAccess`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLogonLogonAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **logon** | [**ModelsSSHLogonsLogonAccessRequest**](ModelsSSHLogonsLogonAccessRequest.md) | Users to add the existing logon | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsSSHAccessLogonUserAccessResponse**](ModelsSSHAccessLogonUserAccessResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LogonQueryLogons

> []ModelsSSHLogonsLogonQueryResponse LogonQueryLogons(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PqQueryString(pqQueryString).PqPageReturned(pqPageReturned).PqReturnLimit(pqReturnLimit).PqSortField(pqSortField).PqSortAscending(pqSortAscending).Execute()

Returns all Logons according to the provided filter parameters

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
    pqQueryString := "pqQueryString_example" // string | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) (optional)
    pqPageReturned := int32(56) // int32 | The current page within the result set to be returned (optional)
    pqReturnLimit := int32(56) // int32 | Maximum number of records to be returned in a single call (optional)
    pqSortField := "pqSortField_example" // string | Field by which the results should be sorted (view results via Management Portal for sortable columns) (optional)
    pqSortAscending := int32(56) // int32 | Field sort direction [0=ascending, 1=descending] (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogonApi.LogonQueryLogons(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PqQueryString(pqQueryString).PqPageReturned(pqPageReturned).PqReturnLimit(pqReturnLimit).PqSortField(pqSortField).PqSortAscending(pqSortAscending).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogonApi.LogonQueryLogons``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LogonQueryLogons`: []ModelsSSHLogonsLogonQueryResponse
    fmt.Fprintf(os.Stdout, "Response from `LogonApi.LogonQueryLogons`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLogonQueryLogonsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]
 **pqQueryString** | **string** | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) | 
 **pqPageReturned** | **int32** | The current page within the result set to be returned | 
 **pqReturnLimit** | **int32** | Maximum number of records to be returned in a single call | 
 **pqSortField** | **string** | Field by which the results should be sorted (view results via Management Portal for sortable columns) | 
 **pqSortAscending** | **int32** | Field sort direction [0&#x3D;ascending, 1&#x3D;descending] | 

### Return type

[**[]ModelsSSHLogonsLogonQueryResponse**](ModelsSSHLogonsLogonQueryResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


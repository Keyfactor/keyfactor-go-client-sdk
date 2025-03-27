# \ServerGroupApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSSHServerGroups**](ServerGroupApi.md#CreateSSHServerGroups) | **POST** /SSH/ServerGroups | Creates a server group with the provided properties
[**CreateSSHServerGroupsAccess**](ServerGroupApi.md#CreateSSHServerGroupsAccess) | **POST** /SSH/ServerGroups/Access | Add access rules to the server group
[**DeleteSSHServerGroupsAccess**](ServerGroupApi.md#DeleteSSHServerGroupsAccess) | **DELETE** /SSH/ServerGroups/Access | Removes access mappings for the specified server group
[**DeleteSSHServerGroupsById**](ServerGroupApi.md#DeleteSSHServerGroupsById) | **DELETE** /SSH/ServerGroups/{id} | Deletes a ServerGroup associated with the provided identifier
[**GetSSHServerGroups**](ServerGroupApi.md#GetSSHServerGroups) | **GET** /SSH/ServerGroups | Returns all server groups according to the provided filter parameters
[**GetSSHServerGroupsAccessById**](ServerGroupApi.md#GetSSHServerGroupsAccessById) | **GET** /SSH/ServerGroups/Access/{id} | Retrieves logons and users with access to those logons for an existing server group
[**GetSSHServerGroupsById**](ServerGroupApi.md#GetSSHServerGroupsById) | **GET** /SSH/ServerGroups/{id} | Returns a ServerGroup associated with the provided identifier
[**GetSSHServerGroupsName**](ServerGroupApi.md#GetSSHServerGroupsName) | **GET** /SSH/ServerGroups/{name} | Returns a ServerGroup associated with the provided identifier
[**UpdateSSHServerGroups**](ServerGroupApi.md#UpdateSSHServerGroups) | **PUT** /SSH/ServerGroups | Updates an existing server group with the provided properties



## CreateSSHServerGroups

> CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse BuildCreateSSHServerGroupsRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsSSHServerGroupsServerGroupCreationRequest(cSSCMSDataModelModelsSSHServerGroupsServerGroupCreationRequest).Execute()

Creates a server group with the provided properties

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
    cSSCMSDataModelModelsSSHServerGroupsServerGroupCreationRequest := *openapiclient.NewCSSCMSDataModelModelsSSHServerGroupsServerGroupCreationRequest("OwnerName_example", "GroupName_example") // CSSCMSDataModelModelsSSHServerGroupsServerGroupCreationRequest | Server group properties to be applied to the new group (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerGroupApi.BuildCreateSSHServerGroupsRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsSSHServerGroupsServerGroupCreationRequest(cSSCMSDataModelModelsSSHServerGroupsServerGroupCreationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerGroupApi.CreateSSHServerGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSSHServerGroups`: CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerGroupApi.CreateSSHServerGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSSHServerGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsSSHServerGroupsServerGroupCreationRequest** | [**CSSCMSDataModelModelsSSHServerGroupsServerGroupCreationRequest**](CSSCMSDataModelModelsSSHServerGroupsServerGroupCreationRequest.md) | Server group properties to be applied to the new group | 

### Return type

[**CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse**](CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSSHServerGroupsAccess

> CSSCMSDataModelModelsSSHAccessServerGroupAccessResponse BuildCreateSSHServerGroupsAccessRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsSSHAccessServerGroupAccessRequest(cSSCMSDataModelModelsSSHAccessServerGroupAccessRequest).Execute()

Add access rules to the server group

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
    cSSCMSDataModelModelsSSHAccessServerGroupAccessRequest := *openapiclient.NewCSSCMSDataModelModelsSSHAccessServerGroupAccessRequest("ServerGroupId_example", []openapiclient.CSSCMSDataModelModelsSSHAccessLogonUserAccessRequest{*openapiclient.NewCSSCMSDataModelModelsSSHAccessLogonUserAccessRequest()}) // CSSCMSDataModelModelsSSHAccessServerGroupAccessRequest |  (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerGroupApi.BuildCreateSSHServerGroupsAccessRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsSSHAccessServerGroupAccessRequest(cSSCMSDataModelModelsSSHAccessServerGroupAccessRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerGroupApi.CreateSSHServerGroupsAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSSHServerGroupsAccess`: CSSCMSDataModelModelsSSHAccessServerGroupAccessResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerGroupApi.CreateSSHServerGroupsAccess`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSSHServerGroupsAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsSSHAccessServerGroupAccessRequest** | [**CSSCMSDataModelModelsSSHAccessServerGroupAccessRequest**](CSSCMSDataModelModelsSSHAccessServerGroupAccessRequest.md) |  | 

### Return type

[**CSSCMSDataModelModelsSSHAccessServerGroupAccessResponse**](CSSCMSDataModelModelsSSHAccessServerGroupAccessResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSSHServerGroupsAccess

> CSSCMSDataModelModelsSSHAccessServerGroupAccessResponse BuildDeleteSSHServerGroupsAccessRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsSSHAccessServerGroupAccessRequest(cSSCMSDataModelModelsSSHAccessServerGroupAccessRequest).Execute()

Removes access mappings for the specified server group

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
    cSSCMSDataModelModelsSSHAccessServerGroupAccessRequest := *openapiclient.NewCSSCMSDataModelModelsSSHAccessServerGroupAccessRequest("ServerGroupId_example", []openapiclient.CSSCMSDataModelModelsSSHAccessLogonUserAccessRequest{*openapiclient.NewCSSCMSDataModelModelsSSHAccessLogonUserAccessRequest()}) // CSSCMSDataModelModelsSSHAccessServerGroupAccessRequest |  (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerGroupApi.BuildDeleteSSHServerGroupsAccessRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsSSHAccessServerGroupAccessRequest(cSSCMSDataModelModelsSSHAccessServerGroupAccessRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerGroupApi.DeleteSSHServerGroupsAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSSHServerGroupsAccess`: CSSCMSDataModelModelsSSHAccessServerGroupAccessResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerGroupApi.DeleteSSHServerGroupsAccess`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSSHServerGroupsAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsSSHAccessServerGroupAccessRequest** | [**CSSCMSDataModelModelsSSHAccessServerGroupAccessRequest**](CSSCMSDataModelModelsSSHAccessServerGroupAccessRequest.md) |  | 

### Return type

[**CSSCMSDataModelModelsSSHAccessServerGroupAccessResponse**](CSSCMSDataModelModelsSSHAccessServerGroupAccessResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSSHServerGroupsById

> BuildDeleteSSHServerGroupsByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Deletes a ServerGroup associated with the provided identifier

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Keyfactor identifer of the ServerGroup to be deleted
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerGroupApi.BuildDeleteSSHServerGroupsByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerGroupApi.DeleteSSHServerGroupsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Keyfactor identifer of the ServerGroup to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSSHServerGroupsByIdRequest struct via the builder pattern


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


## GetSSHServerGroups

> []CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse BuildGetSSHServerGroupsRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns all server groups according to the provided filter parameters

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
    resp, r, err := apiClient.ServerGroupApi.BuildGetSSHServerGroupsRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerGroupApi.GetSSHServerGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSSHServerGroups`: []CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerGroupApi.GetSSHServerGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSSHServerGroupsRequest struct via the builder pattern


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

[**[]CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse**](CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSSHServerGroupsAccessById

> CSSCMSDataModelModelsSSHAccessServerGroupAccessResponse BuildGetSSHServerGroupsAccessByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Retrieves logons and users with access to those logons for an existing server group

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Id of the existing server group
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerGroupApi.BuildGetSSHServerGroupsAccessByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerGroupApi.GetSSHServerGroupsAccessById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSSHServerGroupsAccessById`: CSSCMSDataModelModelsSSHAccessServerGroupAccessResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerGroupApi.GetSSHServerGroupsAccessById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the existing server group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSSHServerGroupsAccessByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**CSSCMSDataModelModelsSSHAccessServerGroupAccessResponse**](CSSCMSDataModelModelsSSHAccessServerGroupAccessResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSSHServerGroupsById

> CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse BuildGetSSHServerGroupsByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns a ServerGroup associated with the provided identifier

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Keyfactor identifier of the ServerGroup
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerGroupApi.BuildGetSSHServerGroupsByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerGroupApi.GetSSHServerGroupsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSSHServerGroupsById`: CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerGroupApi.GetSSHServerGroupsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Keyfactor identifier of the ServerGroup | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSSHServerGroupsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse**](CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSSHServerGroupsName

> CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse BuildGetSSHServerGroupsNameRequest(ctx, name).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns a ServerGroup associated with the provided identifier

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
    name := "name_example" // string | name of the ServerGroup
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerGroupApi.BuildGetSSHServerGroupsNameRequest(context.Background(), name).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerGroupApi.GetSSHServerGroupsName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSSHServerGroupsName`: CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerGroupApi.GetSSHServerGroupsName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | name of the ServerGroup | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSSHServerGroupsNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse**](CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSSHServerGroups

> CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse BuildUpdateSSHServerGroupsRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsSSHServerGroupsServerGroupUpdateRequest(cSSCMSDataModelModelsSSHServerGroupsServerGroupUpdateRequest).Execute()

Updates an existing server group with the provided properties

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
    cSSCMSDataModelModelsSSHServerGroupsServerGroupUpdateRequest := *openapiclient.NewCSSCMSDataModelModelsSSHServerGroupsServerGroupUpdateRequest("Id_example", "OwnerName_example", "GroupName_example", false) // CSSCMSDataModelModelsSSHServerGroupsServerGroupUpdateRequest | Server group properties to be applied to the existing group (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerGroupApi.BuildUpdateSSHServerGroupsRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsSSHServerGroupsServerGroupUpdateRequest(cSSCMSDataModelModelsSSHServerGroupsServerGroupUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerGroupApi.UpdateSSHServerGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSSHServerGroups`: CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerGroupApi.UpdateSSHServerGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSSHServerGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsSSHServerGroupsServerGroupUpdateRequest** | [**CSSCMSDataModelModelsSSHServerGroupsServerGroupUpdateRequest**](CSSCMSDataModelModelsSSHServerGroupsServerGroupUpdateRequest.md) | Server group properties to be applied to the existing group | 

### Return type

[**CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse**](CSSCMSDataModelModelsSSHServerGroupsServerGroupResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


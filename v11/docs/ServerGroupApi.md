# \ServerGroupApi

All URIs are relative to */Keyfactor/API*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SSHServerGroupsAccessDelete**](ServerGroupApi.md#SSHServerGroupsAccessDelete) | **Delete** /SSH/ServerGroups/Access | Removes access mappings for the specified server group
[**SSHServerGroupsAccessIdGet**](ServerGroupApi.md#SSHServerGroupsAccessIdGet) | **Get** /SSH/ServerGroups/Access/{id} | Retrieves logons and users with access to those logons for an existing server group
[**SSHServerGroupsAccessPost**](ServerGroupApi.md#SSHServerGroupsAccessPost) | **Post** /SSH/ServerGroups/Access | Add access rules to the server group
[**SSHServerGroupsGet**](ServerGroupApi.md#SSHServerGroupsGet) | **Get** /SSH/ServerGroups | Returns all server groups according to the provided filter parameters
[**SSHServerGroupsIdDelete**](ServerGroupApi.md#SSHServerGroupsIdDelete) | **Delete** /SSH/ServerGroups/{id} | Deletes a ServerGroup associated with the provided identifier
[**SSHServerGroupsIdGet**](ServerGroupApi.md#SSHServerGroupsIdGet) | **Get** /SSH/ServerGroups/{id} | Returns a ServerGroup associated with the provided identifier
[**SSHServerGroupsNameGet**](ServerGroupApi.md#SSHServerGroupsNameGet) | **Get** /SSH/ServerGroups/{name} | Returns a ServerGroup associated with the provided identifier
[**SSHServerGroupsPost**](ServerGroupApi.md#SSHServerGroupsPost) | **Post** /SSH/ServerGroups | Creates a server group with the provided properties
[**SSHServerGroupsPut**](ServerGroupApi.md#SSHServerGroupsPut) | **Put** /SSH/ServerGroups | Updates an existing server group with the provided properties



## SSHServerGroupsAccessDelete

> ModelsSSHAccessServerGroupAccessResponse SSHServerGroupsAccessDelete(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).ModelsSSHAccessServerGroupAccessRequest(modelsSSHAccessServerGroupAccessRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    modelsSSHAccessServerGroupAccessRequest := *openapiclient.NewModelsSSHAccessServerGroupAccessRequest("ServerGroupId_example", []openapiclient.ModelsSSHAccessLogonUserAccessRequest{*openapiclient.NewModelsSSHAccessLogonUserAccessRequest()}) // ModelsSSHAccessServerGroupAccessRequest |  (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerGroupApi.SSHServerGroupsAccessDelete(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).ModelsSSHAccessServerGroupAccessRequest(modelsSSHAccessServerGroupAccessRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerGroupApi.SSHServerGroupsAccessDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SSHServerGroupsAccessDelete`: ModelsSSHAccessServerGroupAccessResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerGroupApi.SSHServerGroupsAccessDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSSHServerGroupsAccessDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **modelsSSHAccessServerGroupAccessRequest** | [**ModelsSSHAccessServerGroupAccessRequest**](ModelsSSHAccessServerGroupAccessRequest.md) |  | 

### Return type

[**ModelsSSHAccessServerGroupAccessResponse**](ModelsSSHAccessServerGroupAccessResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SSHServerGroupsAccessIdGet

> ModelsSSHAccessServerGroupAccessResponse SSHServerGroupsAccessIdGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerGroupApi.SSHServerGroupsAccessIdGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerGroupApi.SSHServerGroupsAccessIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SSHServerGroupsAccessIdGet`: ModelsSSHAccessServerGroupAccessResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerGroupApi.SSHServerGroupsAccessIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the existing server group | 

### Other Parameters

Other parameters are passed through a pointer to a apiSSHServerGroupsAccessIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**ModelsSSHAccessServerGroupAccessResponse**](ModelsSSHAccessServerGroupAccessResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SSHServerGroupsAccessPost

> ModelsSSHAccessServerGroupAccessResponse SSHServerGroupsAccessPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).ModelsSSHAccessServerGroupAccessRequest(modelsSSHAccessServerGroupAccessRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    modelsSSHAccessServerGroupAccessRequest := *openapiclient.NewModelsSSHAccessServerGroupAccessRequest("ServerGroupId_example", []openapiclient.ModelsSSHAccessLogonUserAccessRequest{*openapiclient.NewModelsSSHAccessLogonUserAccessRequest()}) // ModelsSSHAccessServerGroupAccessRequest |  (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerGroupApi.SSHServerGroupsAccessPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).ModelsSSHAccessServerGroupAccessRequest(modelsSSHAccessServerGroupAccessRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerGroupApi.SSHServerGroupsAccessPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SSHServerGroupsAccessPost`: ModelsSSHAccessServerGroupAccessResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerGroupApi.SSHServerGroupsAccessPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSSHServerGroupsAccessPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **modelsSSHAccessServerGroupAccessRequest** | [**ModelsSSHAccessServerGroupAccessRequest**](ModelsSSHAccessServerGroupAccessRequest.md) |  | 

### Return type

[**ModelsSSHAccessServerGroupAccessResponse**](ModelsSSHAccessServerGroupAccessResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SSHServerGroupsGet

> []ModelsSSHServerGroupsServerGroupResponse SSHServerGroupsGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    queryString := "queryString_example" // string |  (optional)
    pageReturned := int32(56) // int32 |  (optional)
    returnLimit := int32(56) // int32 |  (optional)
    sortField := "sortField_example" // string |  (optional)
    sortAscending := int32(56) // int32 |  (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerGroupApi.SSHServerGroupsGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerGroupApi.SSHServerGroupsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SSHServerGroupsGet`: []ModelsSSHServerGroupsServerGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerGroupApi.SSHServerGroupsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSSHServerGroupsGetRequest struct via the builder pattern


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

[**[]ModelsSSHServerGroupsServerGroupResponse**](ModelsSSHServerGroupsServerGroupResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SSHServerGroupsIdDelete

> SSHServerGroupsIdDelete(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerGroupApi.SSHServerGroupsIdDelete(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerGroupApi.SSHServerGroupsIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSSHServerGroupsIdDeleteRequest struct via the builder pattern


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


## SSHServerGroupsIdGet

> ModelsSSHServerGroupsServerGroupResponse SSHServerGroupsIdGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerGroupApi.SSHServerGroupsIdGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerGroupApi.SSHServerGroupsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SSHServerGroupsIdGet`: ModelsSSHServerGroupsServerGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerGroupApi.SSHServerGroupsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Keyfactor identifier of the ServerGroup | 

### Other Parameters

Other parameters are passed through a pointer to a apiSSHServerGroupsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**ModelsSSHServerGroupsServerGroupResponse**](ModelsSSHServerGroupsServerGroupResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SSHServerGroupsNameGet

> ModelsSSHServerGroupsServerGroupResponse SSHServerGroupsNameGet(ctx, name).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerGroupApi.SSHServerGroupsNameGet(context.Background(), name).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerGroupApi.SSHServerGroupsNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SSHServerGroupsNameGet`: ModelsSSHServerGroupsServerGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerGroupApi.SSHServerGroupsNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | name of the ServerGroup | 

### Other Parameters

Other parameters are passed through a pointer to a apiSSHServerGroupsNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**ModelsSSHServerGroupsServerGroupResponse**](ModelsSSHServerGroupsServerGroupResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SSHServerGroupsPost

> ModelsSSHServerGroupsServerGroupResponse SSHServerGroupsPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).ModelsSSHServerGroupsServerGroupCreationRequest(modelsSSHServerGroupsServerGroupCreationRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    modelsSSHServerGroupsServerGroupCreationRequest := *openapiclient.NewModelsSSHServerGroupsServerGroupCreationRequest("OwnerName_example", "GroupName_example") // ModelsSSHServerGroupsServerGroupCreationRequest | Server group properties to be applied to the new group (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerGroupApi.SSHServerGroupsPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).ModelsSSHServerGroupsServerGroupCreationRequest(modelsSSHServerGroupsServerGroupCreationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerGroupApi.SSHServerGroupsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SSHServerGroupsPost`: ModelsSSHServerGroupsServerGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerGroupApi.SSHServerGroupsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSSHServerGroupsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **modelsSSHServerGroupsServerGroupCreationRequest** | [**ModelsSSHServerGroupsServerGroupCreationRequest**](ModelsSSHServerGroupsServerGroupCreationRequest.md) | Server group properties to be applied to the new group | 

### Return type

[**ModelsSSHServerGroupsServerGroupResponse**](ModelsSSHServerGroupsServerGroupResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SSHServerGroupsPut

> ModelsSSHServerGroupsServerGroupResponse SSHServerGroupsPut(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).ModelsSSHServerGroupsServerGroupUpdateRequest(modelsSSHServerGroupsServerGroupUpdateRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    modelsSSHServerGroupsServerGroupUpdateRequest := *openapiclient.NewModelsSSHServerGroupsServerGroupUpdateRequest("Id_example", "OwnerName_example", "GroupName_example", false) // ModelsSSHServerGroupsServerGroupUpdateRequest | Server group properties to be applied to the existing group (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerGroupApi.SSHServerGroupsPut(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).ModelsSSHServerGroupsServerGroupUpdateRequest(modelsSSHServerGroupsServerGroupUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerGroupApi.SSHServerGroupsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SSHServerGroupsPut`: ModelsSSHServerGroupsServerGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerGroupApi.SSHServerGroupsPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSSHServerGroupsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **modelsSSHServerGroupsServerGroupUpdateRequest** | [**ModelsSSHServerGroupsServerGroupUpdateRequest**](ModelsSSHServerGroupsServerGroupUpdateRequest.md) | Server group properties to be applied to the existing group | 

### Return type

[**ModelsSSHServerGroupsServerGroupResponse**](ModelsSSHServerGroupsServerGroupResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

# \ServerGroupApi

All URIs are relative to *https://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServerGroupAddAccess**](ServerGroupApi.md#ServerGroupAddAccess) | **Post** /SSH/ServerGroups/Access | Add access rules to the server group
[**ServerGroupCreateServerGroup**](ServerGroupApi.md#ServerGroupCreateServerGroup) | **Post** /SSH/ServerGroups | Creates a server group with the provided properties
[**ServerGroupDelete**](ServerGroupApi.md#ServerGroupDelete) | **Delete** /SSH/ServerGroups/{id} | Deletes a ServerGroup associated with the provided identifier
[**ServerGroupGetAccess**](ServerGroupApi.md#ServerGroupGetAccess) | **Get** /SSH/ServerGroups/Access/{id} | Retrieves logons and users with access to those logons for an existing server group
[**ServerGroupGetGroup**](ServerGroupApi.md#ServerGroupGetGroup) | **Get** /SSH/ServerGroups/{id} | Returns a ServerGroup associated with the provided identifier
[**ServerGroupGetGroupByName**](ServerGroupApi.md#ServerGroupGetGroupByName) | **Get** /SSH/ServerGroups/{name} | Returns a ServerGroup associated with the provided identifier
[**ServerGroupQueryServerGroups**](ServerGroupApi.md#ServerGroupQueryServerGroups) | **Get** /SSH/ServerGroups | Returns all server groups according to the provided filter parameters
[**ServerGroupRemoveAccess**](ServerGroupApi.md#ServerGroupRemoveAccess) | **Delete** /SSH/ServerGroups/Access | Removes access mappings for the specified server group
[**ServerGroupUpdateServerGroup**](ServerGroupApi.md#ServerGroupUpdateServerGroup) | **Put** /SSH/ServerGroups | Updates an existing server group with the provided properties



## ServerGroupAddAccess

> ModelsSSHAccessServerGroupAccessResponse ServerGroupAddAccess(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).AccessRequest(accessRequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    accessRequest := *openapiclient.NewModelsSSHAccessServerGroupAccessRequest("00000000-0000-0000-0000-000000000000", []openapiclient.ModelsSSHAccessLogonUserAccessRequest{*openapiclient.NewModelsSSHAccessLogonUserAccessRequest()}) // ModelsSSHAccessServerGroupAccessRequest | 
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerGroupApi.ServerGroupAddAccess(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).AccessRequest(accessRequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerGroupApi.ServerGroupAddAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServerGroupAddAccess`: ModelsSSHAccessServerGroupAccessResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerGroupApi.ServerGroupAddAccess`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServerGroupAddAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **accessRequest** | [**ModelsSSHAccessServerGroupAccessRequest**](ModelsSSHAccessServerGroupAccessRequest.md) |  | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsSSHAccessServerGroupAccessResponse**](ModelsSSHAccessServerGroupAccessResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServerGroupCreateServerGroup

> ModelsSSHServerGroupsServerGroupResponse ServerGroupCreateServerGroup(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).ServerGroupCreationRequest(serverGroupCreationRequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    serverGroupCreationRequest := *openapiclient.NewModelsSSHServerGroupsServerGroupCreationRequest("OwnerName_example", "GroupName_example") // ModelsSSHServerGroupsServerGroupCreationRequest | Server group properties to be applied to the new group
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerGroupApi.ServerGroupCreateServerGroup(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).ServerGroupCreationRequest(serverGroupCreationRequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerGroupApi.ServerGroupCreateServerGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServerGroupCreateServerGroup`: ModelsSSHServerGroupsServerGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerGroupApi.ServerGroupCreateServerGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServerGroupCreateServerGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **serverGroupCreationRequest** | [**ModelsSSHServerGroupsServerGroupCreationRequest**](ModelsSSHServerGroupsServerGroupCreationRequest.md) | Server group properties to be applied to the new group | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsSSHServerGroupsServerGroupResponse**](ModelsSSHServerGroupsServerGroupResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServerGroupDelete

> ServerGroupDelete(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerGroupApi.ServerGroupDelete(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerGroupApi.ServerGroupDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiServerGroupDeleteRequest struct via the builder pattern


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


## ServerGroupGetAccess

> ModelsSSHAccessServerGroupAccessResponse ServerGroupGetAccess(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerGroupApi.ServerGroupGetAccess(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerGroupApi.ServerGroupGetAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServerGroupGetAccess`: ModelsSSHAccessServerGroupAccessResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerGroupApi.ServerGroupGetAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the existing server group | 

### Other Parameters

Other parameters are passed through a pointer to a apiServerGroupGetAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsSSHAccessServerGroupAccessResponse**](ModelsSSHAccessServerGroupAccessResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServerGroupGetGroup

> ModelsSSHServerGroupsServerGroupResponse ServerGroupGetGroup(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerGroupApi.ServerGroupGetGroup(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerGroupApi.ServerGroupGetGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServerGroupGetGroup`: ModelsSSHServerGroupsServerGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerGroupApi.ServerGroupGetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Keyfactor identifier of the ServerGroup | 

### Other Parameters

Other parameters are passed through a pointer to a apiServerGroupGetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsSSHServerGroupsServerGroupResponse**](ModelsSSHServerGroupsServerGroupResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServerGroupGetGroupByName

> ModelsSSHServerGroupsServerGroupResponse ServerGroupGetGroupByName(ctx, name).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerGroupApi.ServerGroupGetGroupByName(context.Background(), name).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerGroupApi.ServerGroupGetGroupByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServerGroupGetGroupByName`: ModelsSSHServerGroupsServerGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerGroupApi.ServerGroupGetGroupByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | name of the ServerGroup | 

### Other Parameters

Other parameters are passed through a pointer to a apiServerGroupGetGroupByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsSSHServerGroupsServerGroupResponse**](ModelsSSHServerGroupsServerGroupResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServerGroupQueryServerGroups

> []ModelsSSHServerGroupsServerGroupResponse ServerGroupQueryServerGroups(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PqQueryString(pqQueryString).PqPageReturned(pqPageReturned).PqReturnLimit(pqReturnLimit).PqSortField(pqSortField).PqSortAscending(pqSortAscending).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")
    pqQueryString := "pqQueryString_example" // string | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) (optional)
    pqPageReturned := int32(56) // int32 | The current page within the result set to be returned (optional)
    pqReturnLimit := int32(56) // int32 | Maximum number of records to be returned in a single call (optional)
    pqSortField := "pqSortField_example" // string | Field by which the results should be sorted (view results via Management Portal for sortable columns) (optional)
    pqSortAscending := int32(56) // int32 | Field sort direction [0=ascending, 1=descending] (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerGroupApi.ServerGroupQueryServerGroups(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PqQueryString(pqQueryString).PqPageReturned(pqPageReturned).PqReturnLimit(pqReturnLimit).PqSortField(pqSortField).PqSortAscending(pqSortAscending).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerGroupApi.ServerGroupQueryServerGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServerGroupQueryServerGroups`: []ModelsSSHServerGroupsServerGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerGroupApi.ServerGroupQueryServerGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServerGroupQueryServerGroupsRequest struct via the builder pattern


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

[**[]ModelsSSHServerGroupsServerGroupResponse**](ModelsSSHServerGroupsServerGroupResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServerGroupRemoveAccess

> ModelsSSHAccessServerGroupAccessResponse ServerGroupRemoveAccess(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).AccessRequest(accessRequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    accessRequest := *openapiclient.NewModelsSSHAccessServerGroupAccessRequest("00000000-0000-0000-0000-000000000000", []openapiclient.ModelsSSHAccessLogonUserAccessRequest{*openapiclient.NewModelsSSHAccessLogonUserAccessRequest()}) // ModelsSSHAccessServerGroupAccessRequest | 
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerGroupApi.ServerGroupRemoveAccess(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).AccessRequest(accessRequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerGroupApi.ServerGroupRemoveAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServerGroupRemoveAccess`: ModelsSSHAccessServerGroupAccessResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerGroupApi.ServerGroupRemoveAccess`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServerGroupRemoveAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **accessRequest** | [**ModelsSSHAccessServerGroupAccessRequest**](ModelsSSHAccessServerGroupAccessRequest.md) |  | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsSSHAccessServerGroupAccessResponse**](ModelsSSHAccessServerGroupAccessResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServerGroupUpdateServerGroup

> ModelsSSHServerGroupsServerGroupResponse ServerGroupUpdateServerGroup(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).UpdateRequest(updateRequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    updateRequest := *openapiclient.NewModelsSSHServerGroupsServerGroupUpdateRequest("00000000-0000-0000-0000-000000000000", "OwnerName_example", "GroupName_example", false) // ModelsSSHServerGroupsServerGroupUpdateRequest | Server group properties to be applied to the existing group
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerGroupApi.ServerGroupUpdateServerGroup(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).UpdateRequest(updateRequest).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerGroupApi.ServerGroupUpdateServerGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServerGroupUpdateServerGroup`: ModelsSSHServerGroupsServerGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerGroupApi.ServerGroupUpdateServerGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServerGroupUpdateServerGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **updateRequest** | [**ModelsSSHServerGroupsServerGroupUpdateRequest**](ModelsSSHServerGroupsServerGroupUpdateRequest.md) | Server group properties to be applied to the existing group | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsSSHServerGroupsServerGroupResponse**](ModelsSSHServerGroupsServerGroupResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


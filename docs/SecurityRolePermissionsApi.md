# \SecurityRolePermissionsApi

All URIs are relative to *https://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SecurityRolePermissionsAddCollectionPermissions**](SecurityRolePermissionsApi.md#SecurityRolePermissionsAddCollectionPermissions) | **Post** /Security/Roles/{id}/Permissions/Collections | Adds collection permissions to the security role that matches the id.
[**SecurityRolePermissionsAddContainerPermissions**](SecurityRolePermissionsApi.md#SecurityRolePermissionsAddContainerPermissions) | **Post** /Security/Roles/{id}/Permissions/Containers | Adds container permissions to the security role that matches the id.
[**SecurityRolePermissionsAddGlobalPermissions**](SecurityRolePermissionsApi.md#SecurityRolePermissionsAddGlobalPermissions) | **Post** /Security/Roles/{id}/Permissions/Global | Adds global permissions to the security role that matches the id.
[**SecurityRolePermissionsGetCollectionPermissionsForRole**](SecurityRolePermissionsApi.md#SecurityRolePermissionsGetCollectionPermissionsForRole) | **Get** /Security/Roles/{id}/Permissions/Collections | Returns all collection permissions associated with the security role that matches the id.
[**SecurityRolePermissionsGetContainerPermissionsForRole**](SecurityRolePermissionsApi.md#SecurityRolePermissionsGetContainerPermissionsForRole) | **Get** /Security/Roles/{id}/Permissions/Containers | Returns all container permissions associated with the security role that matches the id.
[**SecurityRolePermissionsGetGlobalPermissionsForRole**](SecurityRolePermissionsApi.md#SecurityRolePermissionsGetGlobalPermissionsForRole) | **Get** /Security/Roles/{id}/Permissions/Global | Returns all global permissions associated with the security role that matches the id.
[**SecurityRolePermissionsGetPermissionsForRole**](SecurityRolePermissionsApi.md#SecurityRolePermissionsGetPermissionsForRole) | **Get** /Security/Roles/{id}/Permissions | Returns all permissions associated with the security role that matches the id.
[**SecurityRolePermissionsSetCollectionPermissions**](SecurityRolePermissionsApi.md#SecurityRolePermissionsSetCollectionPermissions) | **Put** /Security/Roles/{id}/Permissions/Collections | Sets collection permissions to the security role that matches the id.
[**SecurityRolePermissionsSetContainerPermissions**](SecurityRolePermissionsApi.md#SecurityRolePermissionsSetContainerPermissions) | **Put** /Security/Roles/{id}/Permissions/Containers | Sets container permissions to the security role that matches the id.
[**SecurityRolePermissionsSetGlobalPermissions**](SecurityRolePermissionsApi.md#SecurityRolePermissionsSetGlobalPermissions) | **Put** /Security/Roles/{id}/Permissions/Global | Adds global permissions to the security role that matches the id.



## SecurityRolePermissionsAddCollectionPermissions

> []KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse SecurityRolePermissionsAddCollectionPermissions(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionPermissions(collectionPermissions).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Adds collection permissions to the security role that matches the id.



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
    id := int32(56) // int32 | Security role identifier
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    collectionPermissions := []openapiclient.KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest{*openapiclient.NewKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest()} // []KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest | Collections permissions
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolePermissionsApi.SecurityRolePermissionsAddCollectionPermissions(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionPermissions(collectionPermissions).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolePermissionsApi.SecurityRolePermissionsAddCollectionPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityRolePermissionsAddCollectionPermissions`: []KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolePermissionsApi.SecurityRolePermissionsAddCollectionPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Security role identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityRolePermissionsAddCollectionPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **collectionPermissions** | [**[]KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest**](KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest.md) | Collections permissions | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**[]KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse**](KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityRolePermissionsAddContainerPermissions

> []KeyfactorApiModelsSecurityRolesContainerPermissionResponse SecurityRolePermissionsAddContainerPermissions(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).ContainerPermissions(containerPermissions).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Adds container permissions to the security role that matches the id.



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
    id := int32(56) // int32 | Security role identifier
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    containerPermissions := []openapiclient.KeyfactorApiModelsSecurityRolesContainerPermissionRequest{*openapiclient.NewKeyfactorApiModelsSecurityRolesContainerPermissionRequest()} // []KeyfactorApiModelsSecurityRolesContainerPermissionRequest | Container permissions
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolePermissionsApi.SecurityRolePermissionsAddContainerPermissions(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).ContainerPermissions(containerPermissions).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolePermissionsApi.SecurityRolePermissionsAddContainerPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityRolePermissionsAddContainerPermissions`: []KeyfactorApiModelsSecurityRolesContainerPermissionResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolePermissionsApi.SecurityRolePermissionsAddContainerPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Security role identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityRolePermissionsAddContainerPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **containerPermissions** | [**[]KeyfactorApiModelsSecurityRolesContainerPermissionRequest**](KeyfactorApiModelsSecurityRolesContainerPermissionRequest.md) | Container permissions | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**[]KeyfactorApiModelsSecurityRolesContainerPermissionResponse**](KeyfactorApiModelsSecurityRolesContainerPermissionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityRolePermissionsAddGlobalPermissions

> []KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionResponse SecurityRolePermissionsAddGlobalPermissions(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).GlobalPermissions(globalPermissions).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Adds global permissions to the security role that matches the id.



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
    id := int32(56) // int32 | Security role identifier
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    globalPermissions := []openapiclient.KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest{*openapiclient.NewKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest()} // []KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest | Global permissions
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolePermissionsApi.SecurityRolePermissionsAddGlobalPermissions(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).GlobalPermissions(globalPermissions).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolePermissionsApi.SecurityRolePermissionsAddGlobalPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityRolePermissionsAddGlobalPermissions`: []KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolePermissionsApi.SecurityRolePermissionsAddGlobalPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Security role identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityRolePermissionsAddGlobalPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **globalPermissions** | [**[]KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest**](KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest.md) | Global permissions | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**[]KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionResponse**](KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityRolePermissionsGetCollectionPermissionsForRole

> []KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse SecurityRolePermissionsGetCollectionPermissionsForRole(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns all collection permissions associated with the security role that matches the id.

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
    id := int32(56) // int32 | Security role identifier
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolePermissionsApi.SecurityRolePermissionsGetCollectionPermissionsForRole(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolePermissionsApi.SecurityRolePermissionsGetCollectionPermissionsForRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityRolePermissionsGetCollectionPermissionsForRole`: []KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolePermissionsApi.SecurityRolePermissionsGetCollectionPermissionsForRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Security role identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityRolePermissionsGetCollectionPermissionsForRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**[]KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse**](KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityRolePermissionsGetContainerPermissionsForRole

> []KeyfactorApiModelsSecurityRolesContainerPermissionResponse SecurityRolePermissionsGetContainerPermissionsForRole(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns all container permissions associated with the security role that matches the id.

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
    id := int32(56) // int32 | Security role identifier
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolePermissionsApi.SecurityRolePermissionsGetContainerPermissionsForRole(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolePermissionsApi.SecurityRolePermissionsGetContainerPermissionsForRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityRolePermissionsGetContainerPermissionsForRole`: []KeyfactorApiModelsSecurityRolesContainerPermissionResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolePermissionsApi.SecurityRolePermissionsGetContainerPermissionsForRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Security role identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityRolePermissionsGetContainerPermissionsForRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**[]KeyfactorApiModelsSecurityRolesContainerPermissionResponse**](KeyfactorApiModelsSecurityRolesContainerPermissionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityRolePermissionsGetGlobalPermissionsForRole

> []KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionResponse SecurityRolePermissionsGetGlobalPermissionsForRole(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns all global permissions associated with the security role that matches the id.

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
    id := int32(56) // int32 | Security role identifier
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolePermissionsApi.SecurityRolePermissionsGetGlobalPermissionsForRole(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolePermissionsApi.SecurityRolePermissionsGetGlobalPermissionsForRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityRolePermissionsGetGlobalPermissionsForRole`: []KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolePermissionsApi.SecurityRolePermissionsGetGlobalPermissionsForRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Security role identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityRolePermissionsGetGlobalPermissionsForRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**[]KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionResponse**](KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityRolePermissionsGetPermissionsForRole

> []KeyfactorApiModelsSecurityRolesAreaPermissionResponse SecurityRolePermissionsGetPermissionsForRole(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns all permissions associated with the security role that matches the id.

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
    id := int32(56) // int32 | Security role identifier
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolePermissionsApi.SecurityRolePermissionsGetPermissionsForRole(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolePermissionsApi.SecurityRolePermissionsGetPermissionsForRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityRolePermissionsGetPermissionsForRole`: []KeyfactorApiModelsSecurityRolesAreaPermissionResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolePermissionsApi.SecurityRolePermissionsGetPermissionsForRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Security role identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityRolePermissionsGetPermissionsForRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**[]KeyfactorApiModelsSecurityRolesAreaPermissionResponse**](KeyfactorApiModelsSecurityRolesAreaPermissionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityRolePermissionsSetCollectionPermissions

> []KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse SecurityRolePermissionsSetCollectionPermissions(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionPermissions(collectionPermissions).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Sets collection permissions to the security role that matches the id.



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
    id := int32(56) // int32 | Security role identifier
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    collectionPermissions := []openapiclient.KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest{*openapiclient.NewKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest()} // []KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest | Collections permissions
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolePermissionsApi.SecurityRolePermissionsSetCollectionPermissions(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionPermissions(collectionPermissions).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolePermissionsApi.SecurityRolePermissionsSetCollectionPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityRolePermissionsSetCollectionPermissions`: []KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolePermissionsApi.SecurityRolePermissionsSetCollectionPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Security role identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityRolePermissionsSetCollectionPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **collectionPermissions** | [**[]KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest**](KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionRequest.md) | Collections permissions | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**[]KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse**](KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesCollectionPermissionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityRolePermissionsSetContainerPermissions

> []KeyfactorApiModelsSecurityRolesContainerPermissionResponse SecurityRolePermissionsSetContainerPermissions(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).ContainerPermissions(containerPermissions).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Sets container permissions to the security role that matches the id.



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
    id := int32(56) // int32 | Security role identifier
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    containerPermissions := []openapiclient.KeyfactorApiModelsSecurityRolesContainerPermissionRequest{*openapiclient.NewKeyfactorApiModelsSecurityRolesContainerPermissionRequest()} // []KeyfactorApiModelsSecurityRolesContainerPermissionRequest | Container permissions
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolePermissionsApi.SecurityRolePermissionsSetContainerPermissions(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).ContainerPermissions(containerPermissions).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolePermissionsApi.SecurityRolePermissionsSetContainerPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityRolePermissionsSetContainerPermissions`: []KeyfactorApiModelsSecurityRolesContainerPermissionResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolePermissionsApi.SecurityRolePermissionsSetContainerPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Security role identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityRolePermissionsSetContainerPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **containerPermissions** | [**[]KeyfactorApiModelsSecurityRolesContainerPermissionRequest**](KeyfactorApiModelsSecurityRolesContainerPermissionRequest.md) | Container permissions | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**[]KeyfactorApiModelsSecurityRolesContainerPermissionResponse**](KeyfactorApiModelsSecurityRolesContainerPermissionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityRolePermissionsSetGlobalPermissions

> []KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionResponse SecurityRolePermissionsSetGlobalPermissions(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).GlobalPermissions(globalPermissions).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Adds global permissions to the security role that matches the id.



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
    id := int32(56) // int32 | Security role identifier
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    globalPermissions := []openapiclient.KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest{*openapiclient.NewKeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest()} // []KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest | Global permissions
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolePermissionsApi.SecurityRolePermissionsSetGlobalPermissions(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).GlobalPermissions(globalPermissions).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolePermissionsApi.SecurityRolePermissionsSetGlobalPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityRolePermissionsSetGlobalPermissions`: []KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolePermissionsApi.SecurityRolePermissionsSetGlobalPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Security role identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityRolePermissionsSetGlobalPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **globalPermissions** | [**[]KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest**](KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionRequest.md) | Global permissions | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**[]KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionResponse**](KeyfactorApiModelsSecurityRolesIdentitiesSecurityRolesGlobalPermissionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \SecurityRolePermissionsApi

All URIs are relative to */Keyfactor/API*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SecurityRolesIdPermissionsCollectionsGet**](SecurityRolePermissionsApi.md#SecurityRolesIdPermissionsCollectionsGet) | **Get** /Security/Roles/{id}/Permissions/Collections | Returns all collection permissions associated with the security role that matches the id.
[**SecurityRolesIdPermissionsCollectionsPost**](SecurityRolePermissionsApi.md#SecurityRolesIdPermissionsCollectionsPost) | **Post** /Security/Roles/{id}/Permissions/Collections | Adds collection permissions to the security role that matches the id.
[**SecurityRolesIdPermissionsCollectionsPut**](SecurityRolePermissionsApi.md#SecurityRolesIdPermissionsCollectionsPut) | **Put** /Security/Roles/{id}/Permissions/Collections | Sets collection permissions to the security role that matches the id.
[**SecurityRolesIdPermissionsContainersGet**](SecurityRolePermissionsApi.md#SecurityRolesIdPermissionsContainersGet) | **Get** /Security/Roles/{id}/Permissions/Containers | Returns all container permissions associated with the security role that matches the id.
[**SecurityRolesIdPermissionsContainersPost**](SecurityRolePermissionsApi.md#SecurityRolesIdPermissionsContainersPost) | **Post** /Security/Roles/{id}/Permissions/Containers | Adds container permissions to the security role that matches the id.
[**SecurityRolesIdPermissionsContainersPut**](SecurityRolePermissionsApi.md#SecurityRolesIdPermissionsContainersPut) | **Put** /Security/Roles/{id}/Permissions/Containers | Sets container permissions to the security role that matches the id.
[**SecurityRolesIdPermissionsGet**](SecurityRolePermissionsApi.md#SecurityRolesIdPermissionsGet) | **Get** /Security/Roles/{id}/Permissions | Returns all permissions associated with the security role that matches the id.
[**SecurityRolesIdPermissionsGlobalGet**](SecurityRolePermissionsApi.md#SecurityRolesIdPermissionsGlobalGet) | **Get** /Security/Roles/{id}/Permissions/Global | Returns all global permissions associated with the security role that matches the id.
[**SecurityRolesIdPermissionsGlobalPost**](SecurityRolePermissionsApi.md#SecurityRolesIdPermissionsGlobalPost) | **Post** /Security/Roles/{id}/Permissions/Global | Adds global permissions to the security role that matches the id.
[**SecurityRolesIdPermissionsGlobalPut**](SecurityRolePermissionsApi.md#SecurityRolesIdPermissionsGlobalPut) | **Put** /Security/Roles/{id}/Permissions/Global | Adds global permissions to the security role that matches the id.
[**SecurityRolesIdPermissionsPamProvidersGet**](SecurityRolePermissionsApi.md#SecurityRolesIdPermissionsPamProvidersGet) | **Get** /Security/Roles/{id}/Permissions/PamProviders | Returns all PAM provider permissions associated with the security role that matches the id.
[**SecurityRolesIdPermissionsPamProvidersPut**](SecurityRolePermissionsApi.md#SecurityRolesIdPermissionsPamProvidersPut) | **Put** /Security/Roles/{id}/Permissions/PamProviders | Sets PAM provider permissions to the security role that matches the id.



## SecurityRolesIdPermissionsCollectionsGet

> []map[string]interface{} SecurityRolesIdPermissionsCollectionsGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolePermissionsApi.SecurityRolesIdPermissionsCollectionsGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolePermissionsApi.SecurityRolesIdPermissionsCollectionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityRolesIdPermissionsCollectionsGet`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolePermissionsApi.SecurityRolesIdPermissionsCollectionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Security role identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityRolesIdPermissionsCollectionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

**[]map[string]interface{}**

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityRolesIdPermissionsCollectionsPost

> []map[string]interface{} SecurityRolesIdPermissionsCollectionsPost(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).RequestBody(requestBody).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    requestBody := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | Collections permissions (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolePermissionsApi.SecurityRolesIdPermissionsCollectionsPost(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolePermissionsApi.SecurityRolesIdPermissionsCollectionsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityRolesIdPermissionsCollectionsPost`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolePermissionsApi.SecurityRolesIdPermissionsCollectionsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Security role identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityRolesIdPermissionsCollectionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **requestBody** | **[]map[string]interface{}** | Collections permissions | 

### Return type

**[]map[string]interface{}**

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityRolesIdPermissionsCollectionsPut

> []map[string]interface{} SecurityRolesIdPermissionsCollectionsPut(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).RequestBody(requestBody).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    requestBody := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | Collections permissions (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolePermissionsApi.SecurityRolesIdPermissionsCollectionsPut(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolePermissionsApi.SecurityRolesIdPermissionsCollectionsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityRolesIdPermissionsCollectionsPut`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolePermissionsApi.SecurityRolesIdPermissionsCollectionsPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Security role identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityRolesIdPermissionsCollectionsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **requestBody** | **[]map[string]interface{}** | Collections permissions | 

### Return type

**[]map[string]interface{}**

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityRolesIdPermissionsContainersGet

> []map[string]interface{} SecurityRolesIdPermissionsContainersGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolePermissionsApi.SecurityRolesIdPermissionsContainersGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolePermissionsApi.SecurityRolesIdPermissionsContainersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityRolesIdPermissionsContainersGet`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolePermissionsApi.SecurityRolesIdPermissionsContainersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Security role identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityRolesIdPermissionsContainersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

**[]map[string]interface{}**

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityRolesIdPermissionsContainersPost

> []map[string]interface{} SecurityRolesIdPermissionsContainersPost(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).RequestBody(requestBody).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    requestBody := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | Container permissions (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolePermissionsApi.SecurityRolesIdPermissionsContainersPost(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolePermissionsApi.SecurityRolesIdPermissionsContainersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityRolesIdPermissionsContainersPost`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolePermissionsApi.SecurityRolesIdPermissionsContainersPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Security role identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityRolesIdPermissionsContainersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **requestBody** | **[]map[string]interface{}** | Container permissions | 

### Return type

**[]map[string]interface{}**

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityRolesIdPermissionsContainersPut

> []map[string]interface{} SecurityRolesIdPermissionsContainersPut(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).RequestBody(requestBody).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    requestBody := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | Container permissions (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolePermissionsApi.SecurityRolesIdPermissionsContainersPut(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolePermissionsApi.SecurityRolesIdPermissionsContainersPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityRolesIdPermissionsContainersPut`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolePermissionsApi.SecurityRolesIdPermissionsContainersPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Security role identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityRolesIdPermissionsContainersPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **requestBody** | **[]map[string]interface{}** | Container permissions | 

### Return type

**[]map[string]interface{}**

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityRolesIdPermissionsGet

> []map[string]interface{} SecurityRolesIdPermissionsGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolePermissionsApi.SecurityRolesIdPermissionsGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolePermissionsApi.SecurityRolesIdPermissionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityRolesIdPermissionsGet`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolePermissionsApi.SecurityRolesIdPermissionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Security role identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityRolesIdPermissionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

**[]map[string]interface{}**

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityRolesIdPermissionsGlobalGet

> []map[string]interface{} SecurityRolesIdPermissionsGlobalGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolePermissionsApi.SecurityRolesIdPermissionsGlobalGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolePermissionsApi.SecurityRolesIdPermissionsGlobalGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityRolesIdPermissionsGlobalGet`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolePermissionsApi.SecurityRolesIdPermissionsGlobalGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Security role identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityRolesIdPermissionsGlobalGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

**[]map[string]interface{}**

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityRolesIdPermissionsGlobalPost

> []map[string]interface{} SecurityRolesIdPermissionsGlobalPost(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).RequestBody(requestBody).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    requestBody := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | Global permissions (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolePermissionsApi.SecurityRolesIdPermissionsGlobalPost(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolePermissionsApi.SecurityRolesIdPermissionsGlobalPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityRolesIdPermissionsGlobalPost`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolePermissionsApi.SecurityRolesIdPermissionsGlobalPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Security role identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityRolesIdPermissionsGlobalPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **requestBody** | **[]map[string]interface{}** | Global permissions | 

### Return type

**[]map[string]interface{}**

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityRolesIdPermissionsGlobalPut

> []map[string]interface{} SecurityRolesIdPermissionsGlobalPut(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).RequestBody(requestBody).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    requestBody := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | Global permissions (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolePermissionsApi.SecurityRolesIdPermissionsGlobalPut(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolePermissionsApi.SecurityRolesIdPermissionsGlobalPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityRolesIdPermissionsGlobalPut`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolePermissionsApi.SecurityRolesIdPermissionsGlobalPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Security role identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityRolesIdPermissionsGlobalPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **requestBody** | **[]map[string]interface{}** | Global permissions | 

### Return type

**[]map[string]interface{}**

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityRolesIdPermissionsPamProvidersGet

> []KeyfactorWebKeyfactorApiModelsSecuritySecurityRolePermissionsPamProviderPermissionResponse SecurityRolesIdPermissionsPamProvidersGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns all PAM provider permissions associated with the security role that matches the id.

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolePermissionsApi.SecurityRolesIdPermissionsPamProvidersGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolePermissionsApi.SecurityRolesIdPermissionsPamProvidersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityRolesIdPermissionsPamProvidersGet`: []KeyfactorWebKeyfactorApiModelsSecuritySecurityRolePermissionsPamProviderPermissionResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolePermissionsApi.SecurityRolesIdPermissionsPamProvidersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Security role identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityRolesIdPermissionsPamProvidersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]KeyfactorWebKeyfactorApiModelsSecuritySecurityRolePermissionsPamProviderPermissionResponse**](KeyfactorWebKeyfactorApiModelsSecuritySecurityRolePermissionsPamProviderPermissionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityRolesIdPermissionsPamProvidersPut

> []KeyfactorWebKeyfactorApiModelsSecuritySecurityRolePermissionsPamProviderPermissionResponse SecurityRolesIdPermissionsPamProvidersPut(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).RequestBody(requestBody).Execute()

Sets PAM provider permissions to the security role that matches the id.



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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    requestBody := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | PAM Provider permissions (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolePermissionsApi.SecurityRolesIdPermissionsPamProvidersPut(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolePermissionsApi.SecurityRolesIdPermissionsPamProvidersPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityRolesIdPermissionsPamProvidersPut`: []KeyfactorWebKeyfactorApiModelsSecuritySecurityRolePermissionsPamProviderPermissionResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolePermissionsApi.SecurityRolesIdPermissionsPamProvidersPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Security role identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityRolesIdPermissionsPamProvidersPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **requestBody** | **[]map[string]interface{}** | PAM Provider permissions | 

### Return type

[**[]KeyfactorWebKeyfactorApiModelsSecuritySecurityRolePermissionsPamProviderPermissionResponse**](KeyfactorWebKeyfactorApiModelsSecuritySecurityRolePermissionsPamProviderPermissionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


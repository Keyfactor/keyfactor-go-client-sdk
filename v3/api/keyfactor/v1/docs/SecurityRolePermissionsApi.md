# \SecurityRolePermissionsApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSecurityRolesByIdPermissionsCollections**](SecurityRolePermissionsApi.md#CreateSecurityRolesByIdPermissionsCollections) | **POST** /Security/Roles/{id}/Permissions/Collections | Adds collection permissions to the security role that matches the id.
[**CreateSecurityRolesByIdPermissionsContainers**](SecurityRolePermissionsApi.md#CreateSecurityRolesByIdPermissionsContainers) | **POST** /Security/Roles/{id}/Permissions/Containers | Adds container permissions to the security role that matches the id.
[**CreateSecurityRolesByIdPermissionsGlobal**](SecurityRolePermissionsApi.md#CreateSecurityRolesByIdPermissionsGlobal) | **POST** /Security/Roles/{id}/Permissions/Global | Adds global permissions to the security role that matches the id.
[**GetSecurityRolesByIdPermissions**](SecurityRolePermissionsApi.md#GetSecurityRolesByIdPermissions) | **GET** /Security/Roles/{id}/Permissions | Returns all permissions associated with the security role that matches the id.
[**GetSecurityRolesByIdPermissionsCollections**](SecurityRolePermissionsApi.md#GetSecurityRolesByIdPermissionsCollections) | **GET** /Security/Roles/{id}/Permissions/Collections | Returns all collection permissions associated with the security role that matches the id.
[**GetSecurityRolesByIdPermissionsContainers**](SecurityRolePermissionsApi.md#GetSecurityRolesByIdPermissionsContainers) | **GET** /Security/Roles/{id}/Permissions/Containers | Returns all container permissions associated with the security role that matches the id.
[**GetSecurityRolesByIdPermissionsGlobal**](SecurityRolePermissionsApi.md#GetSecurityRolesByIdPermissionsGlobal) | **GET** /Security/Roles/{id}/Permissions/Global | Returns all global permissions associated with the security role that matches the id.
[**GetSecurityRolesByIdPermissionsPamProviders**](SecurityRolePermissionsApi.md#GetSecurityRolesByIdPermissionsPamProviders) | **GET** /Security/Roles/{id}/Permissions/PamProviders | Returns all PAM provider permissions associated with the security role that matches the id.
[**UpdateSecurityRolesByIdPermissionsCollections**](SecurityRolePermissionsApi.md#UpdateSecurityRolesByIdPermissionsCollections) | **PUT** /Security/Roles/{id}/Permissions/Collections | Sets collection permissions to the security role that matches the id.
[**UpdateSecurityRolesByIdPermissionsContainers**](SecurityRolePermissionsApi.md#UpdateSecurityRolesByIdPermissionsContainers) | **PUT** /Security/Roles/{id}/Permissions/Containers | Sets container permissions to the security role that matches the id.
[**UpdateSecurityRolesByIdPermissionsGlobal**](SecurityRolePermissionsApi.md#UpdateSecurityRolesByIdPermissionsGlobal) | **PUT** /Security/Roles/{id}/Permissions/Global | Adds global permissions to the security role that matches the id.
[**UpdateSecurityRolesByIdPermissionsPamProviders**](SecurityRolePermissionsApi.md#UpdateSecurityRolesByIdPermissionsPamProviders) | **PUT** /Security/Roles/{id}/Permissions/PamProviders | Sets PAM provider permissions to the security role that matches the id.



## CreateSecurityRolesByIdPermissionsCollections

> []SecuritySecurityRolePermissionsCollectionPermissionResponse NewCreateSecurityRolesByIdPermissionsCollectionsRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).SecuritySecurityRolePermissionsCollectionPermissionRequest(securitySecurityRolePermissionsCollectionPermissionRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    securitySecurityRolePermissionsCollectionPermissionRequest := []openapiclient.SecuritySecurityRolePermissionsCollectionPermissionRequest{*openapiclient.NewSecuritySecurityRolePermissionsCollectionPermissionRequest()} // []SecuritySecurityRolePermissionsCollectionPermissionRequest | Collections permissions (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolePermissionsApi.NewCreateSecurityRolesByIdPermissionsCollectionsRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).SecuritySecurityRolePermissionsCollectionPermissionRequest(securitySecurityRolePermissionsCollectionPermissionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolePermissionsApi.CreateSecurityRolesByIdPermissionsCollections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSecurityRolesByIdPermissionsCollections`: []SecuritySecurityRolePermissionsCollectionPermissionResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolePermissionsApi.CreateSecurityRolesByIdPermissionsCollections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Security role identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecurityRolesByIdPermissionsCollectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **securitySecurityRolePermissionsCollectionPermissionRequest** | [**[]SecuritySecurityRolePermissionsCollectionPermissionRequest**](SecuritySecurityRolePermissionsCollectionPermissionRequest.md) | Collections permissions | 

### Return type

[**[]SecuritySecurityRolePermissionsCollectionPermissionResponse**](SecuritySecurityRolePermissionsCollectionPermissionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSecurityRolesByIdPermissionsContainers

> []SecuritySecurityRolePermissionsContainerPermissionResponse NewCreateSecurityRolesByIdPermissionsContainersRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).SecuritySecurityRolePermissionsContainerPermissionRequest(securitySecurityRolePermissionsContainerPermissionRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    securitySecurityRolePermissionsContainerPermissionRequest := []openapiclient.SecuritySecurityRolePermissionsContainerPermissionRequest{*openapiclient.NewSecuritySecurityRolePermissionsContainerPermissionRequest()} // []SecuritySecurityRolePermissionsContainerPermissionRequest | Container permissions (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolePermissionsApi.NewCreateSecurityRolesByIdPermissionsContainersRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).SecuritySecurityRolePermissionsContainerPermissionRequest(securitySecurityRolePermissionsContainerPermissionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolePermissionsApi.CreateSecurityRolesByIdPermissionsContainers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSecurityRolesByIdPermissionsContainers`: []SecuritySecurityRolePermissionsContainerPermissionResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolePermissionsApi.CreateSecurityRolesByIdPermissionsContainers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Security role identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecurityRolesByIdPermissionsContainersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **securitySecurityRolePermissionsContainerPermissionRequest** | [**[]SecuritySecurityRolePermissionsContainerPermissionRequest**](SecuritySecurityRolePermissionsContainerPermissionRequest.md) | Container permissions | 

### Return type

[**[]SecuritySecurityRolePermissionsContainerPermissionResponse**](SecuritySecurityRolePermissionsContainerPermissionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSecurityRolesByIdPermissionsGlobal

> []SecuritySecurityRolePermissionsGlobalPermissionResponse NewCreateSecurityRolesByIdPermissionsGlobalRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).SecuritySecurityRolePermissionsGlobalPermissionRequest(securitySecurityRolePermissionsGlobalPermissionRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    securitySecurityRolePermissionsGlobalPermissionRequest := []openapiclient.SecuritySecurityRolePermissionsGlobalPermissionRequest{*openapiclient.NewSecuritySecurityRolePermissionsGlobalPermissionRequest()} // []SecuritySecurityRolePermissionsGlobalPermissionRequest | Global permissions (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolePermissionsApi.NewCreateSecurityRolesByIdPermissionsGlobalRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).SecuritySecurityRolePermissionsGlobalPermissionRequest(securitySecurityRolePermissionsGlobalPermissionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolePermissionsApi.CreateSecurityRolesByIdPermissionsGlobal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSecurityRolesByIdPermissionsGlobal`: []SecuritySecurityRolePermissionsGlobalPermissionResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolePermissionsApi.CreateSecurityRolesByIdPermissionsGlobal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Security role identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecurityRolesByIdPermissionsGlobalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **securitySecurityRolePermissionsGlobalPermissionRequest** | [**[]SecuritySecurityRolePermissionsGlobalPermissionRequest**](SecuritySecurityRolePermissionsGlobalPermissionRequest.md) | Global permissions | 

### Return type

[**[]SecuritySecurityRolePermissionsGlobalPermissionResponse**](SecuritySecurityRolePermissionsGlobalPermissionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurityRolesByIdPermissions

> []SecuritySecurityRolePermissionsAreaPermissionResponse NewGetSecurityRolesByIdPermissionsRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolePermissionsApi.NewGetSecurityRolesByIdPermissionsRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolePermissionsApi.GetSecurityRolesByIdPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecurityRolesByIdPermissions`: []SecuritySecurityRolePermissionsAreaPermissionResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolePermissionsApi.GetSecurityRolesByIdPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Security role identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityRolesByIdPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]SecuritySecurityRolePermissionsAreaPermissionResponse**](SecuritySecurityRolePermissionsAreaPermissionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurityRolesByIdPermissionsCollections

> []SecuritySecurityRolePermissionsCollectionPermissionResponse NewGetSecurityRolesByIdPermissionsCollectionsRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolePermissionsApi.NewGetSecurityRolesByIdPermissionsCollectionsRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolePermissionsApi.GetSecurityRolesByIdPermissionsCollections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecurityRolesByIdPermissionsCollections`: []SecuritySecurityRolePermissionsCollectionPermissionResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolePermissionsApi.GetSecurityRolesByIdPermissionsCollections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Security role identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityRolesByIdPermissionsCollectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]SecuritySecurityRolePermissionsCollectionPermissionResponse**](SecuritySecurityRolePermissionsCollectionPermissionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurityRolesByIdPermissionsContainers

> []SecuritySecurityRolePermissionsContainerPermissionResponse NewGetSecurityRolesByIdPermissionsContainersRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolePermissionsApi.NewGetSecurityRolesByIdPermissionsContainersRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolePermissionsApi.GetSecurityRolesByIdPermissionsContainers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecurityRolesByIdPermissionsContainers`: []SecuritySecurityRolePermissionsContainerPermissionResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolePermissionsApi.GetSecurityRolesByIdPermissionsContainers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Security role identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityRolesByIdPermissionsContainersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]SecuritySecurityRolePermissionsContainerPermissionResponse**](SecuritySecurityRolePermissionsContainerPermissionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurityRolesByIdPermissionsGlobal

> []SecuritySecurityRolePermissionsGlobalPermissionResponse NewGetSecurityRolesByIdPermissionsGlobalRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolePermissionsApi.NewGetSecurityRolesByIdPermissionsGlobalRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolePermissionsApi.GetSecurityRolesByIdPermissionsGlobal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecurityRolesByIdPermissionsGlobal`: []SecuritySecurityRolePermissionsGlobalPermissionResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolePermissionsApi.GetSecurityRolesByIdPermissionsGlobal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Security role identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityRolesByIdPermissionsGlobalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]SecuritySecurityRolePermissionsGlobalPermissionResponse**](SecuritySecurityRolePermissionsGlobalPermissionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurityRolesByIdPermissionsPamProviders

> []SecuritySecurityRolePermissionsPamProviderPermissionResponse NewGetSecurityRolesByIdPermissionsPamProvidersRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolePermissionsApi.NewGetSecurityRolesByIdPermissionsPamProvidersRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolePermissionsApi.GetSecurityRolesByIdPermissionsPamProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecurityRolesByIdPermissionsPamProviders`: []SecuritySecurityRolePermissionsPamProviderPermissionResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolePermissionsApi.GetSecurityRolesByIdPermissionsPamProviders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Security role identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityRolesByIdPermissionsPamProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]SecuritySecurityRolePermissionsPamProviderPermissionResponse**](SecuritySecurityRolePermissionsPamProviderPermissionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSecurityRolesByIdPermissionsCollections

> []SecuritySecurityRolePermissionsCollectionPermissionResponse NewUpdateSecurityRolesByIdPermissionsCollectionsRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).SecuritySecurityRolePermissionsCollectionPermissionRequest(securitySecurityRolePermissionsCollectionPermissionRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    securitySecurityRolePermissionsCollectionPermissionRequest := []openapiclient.SecuritySecurityRolePermissionsCollectionPermissionRequest{*openapiclient.NewSecuritySecurityRolePermissionsCollectionPermissionRequest()} // []SecuritySecurityRolePermissionsCollectionPermissionRequest | Collections permissions (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolePermissionsApi.NewUpdateSecurityRolesByIdPermissionsCollectionsRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).SecuritySecurityRolePermissionsCollectionPermissionRequest(securitySecurityRolePermissionsCollectionPermissionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolePermissionsApi.UpdateSecurityRolesByIdPermissionsCollections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSecurityRolesByIdPermissionsCollections`: []SecuritySecurityRolePermissionsCollectionPermissionResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolePermissionsApi.UpdateSecurityRolesByIdPermissionsCollections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Security role identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSecurityRolesByIdPermissionsCollectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **securitySecurityRolePermissionsCollectionPermissionRequest** | [**[]SecuritySecurityRolePermissionsCollectionPermissionRequest**](SecuritySecurityRolePermissionsCollectionPermissionRequest.md) | Collections permissions | 

### Return type

[**[]SecuritySecurityRolePermissionsCollectionPermissionResponse**](SecuritySecurityRolePermissionsCollectionPermissionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSecurityRolesByIdPermissionsContainers

> []SecuritySecurityRolePermissionsContainerPermissionResponse NewUpdateSecurityRolesByIdPermissionsContainersRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).SecuritySecurityRolePermissionsContainerPermissionRequest(securitySecurityRolePermissionsContainerPermissionRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    securitySecurityRolePermissionsContainerPermissionRequest := []openapiclient.SecuritySecurityRolePermissionsContainerPermissionRequest{*openapiclient.NewSecuritySecurityRolePermissionsContainerPermissionRequest()} // []SecuritySecurityRolePermissionsContainerPermissionRequest | Container permissions (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolePermissionsApi.NewUpdateSecurityRolesByIdPermissionsContainersRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).SecuritySecurityRolePermissionsContainerPermissionRequest(securitySecurityRolePermissionsContainerPermissionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolePermissionsApi.UpdateSecurityRolesByIdPermissionsContainers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSecurityRolesByIdPermissionsContainers`: []SecuritySecurityRolePermissionsContainerPermissionResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolePermissionsApi.UpdateSecurityRolesByIdPermissionsContainers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Security role identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSecurityRolesByIdPermissionsContainersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **securitySecurityRolePermissionsContainerPermissionRequest** | [**[]SecuritySecurityRolePermissionsContainerPermissionRequest**](SecuritySecurityRolePermissionsContainerPermissionRequest.md) | Container permissions | 

### Return type

[**[]SecuritySecurityRolePermissionsContainerPermissionResponse**](SecuritySecurityRolePermissionsContainerPermissionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSecurityRolesByIdPermissionsGlobal

> []SecuritySecurityRolePermissionsGlobalPermissionResponse NewUpdateSecurityRolesByIdPermissionsGlobalRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).SecuritySecurityRolePermissionsGlobalPermissionRequest(securitySecurityRolePermissionsGlobalPermissionRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    securitySecurityRolePermissionsGlobalPermissionRequest := []openapiclient.SecuritySecurityRolePermissionsGlobalPermissionRequest{*openapiclient.NewSecuritySecurityRolePermissionsGlobalPermissionRequest()} // []SecuritySecurityRolePermissionsGlobalPermissionRequest | Global permissions (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolePermissionsApi.NewUpdateSecurityRolesByIdPermissionsGlobalRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).SecuritySecurityRolePermissionsGlobalPermissionRequest(securitySecurityRolePermissionsGlobalPermissionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolePermissionsApi.UpdateSecurityRolesByIdPermissionsGlobal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSecurityRolesByIdPermissionsGlobal`: []SecuritySecurityRolePermissionsGlobalPermissionResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolePermissionsApi.UpdateSecurityRolesByIdPermissionsGlobal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Security role identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSecurityRolesByIdPermissionsGlobalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **securitySecurityRolePermissionsGlobalPermissionRequest** | [**[]SecuritySecurityRolePermissionsGlobalPermissionRequest**](SecuritySecurityRolePermissionsGlobalPermissionRequest.md) | Global permissions | 

### Return type

[**[]SecuritySecurityRolePermissionsGlobalPermissionResponse**](SecuritySecurityRolePermissionsGlobalPermissionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSecurityRolesByIdPermissionsPamProviders

> []SecuritySecurityRolePermissionsPamProviderPermissionResponse NewUpdateSecurityRolesByIdPermissionsPamProvidersRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).SecuritySecurityRolePermissionsPamProviderPermissionRequest(securitySecurityRolePermissionsPamProviderPermissionRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    securitySecurityRolePermissionsPamProviderPermissionRequest := []openapiclient.SecuritySecurityRolePermissionsPamProviderPermissionRequest{*openapiclient.NewSecuritySecurityRolePermissionsPamProviderPermissionRequest()} // []SecuritySecurityRolePermissionsPamProviderPermissionRequest | PAM Provider permissions (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolePermissionsApi.NewUpdateSecurityRolesByIdPermissionsPamProvidersRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).SecuritySecurityRolePermissionsPamProviderPermissionRequest(securitySecurityRolePermissionsPamProviderPermissionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolePermissionsApi.UpdateSecurityRolesByIdPermissionsPamProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSecurityRolesByIdPermissionsPamProviders`: []SecuritySecurityRolePermissionsPamProviderPermissionResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolePermissionsApi.UpdateSecurityRolesByIdPermissionsPamProviders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Security role identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSecurityRolesByIdPermissionsPamProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **securitySecurityRolePermissionsPamProviderPermissionRequest** | [**[]SecuritySecurityRolePermissionsPamProviderPermissionRequest**](SecuritySecurityRolePermissionsPamProviderPermissionRequest.md) | PAM Provider permissions | 

### Return type

[**[]SecuritySecurityRolePermissionsPamProviderPermissionResponse**](SecuritySecurityRolePermissionsPamProviderPermissionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


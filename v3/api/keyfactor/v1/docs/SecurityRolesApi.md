# \SecurityRolesApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSecurityRoles**](SecurityRolesApi.md#CreateSecurityRoles) | **POST** /Security/Roles | Adds a new security role to the system.  The v1 endpoints for Security Roles will only work against AD Identities.  Please use the v2 endpoints instead.
[**CreateSecurityRolesByIdCopy**](SecurityRolesApi.md#CreateSecurityRolesByIdCopy) | **POST** /Security/Roles/{id}/Copy | Makes a copy of an existing security role.  The v1 endpoints for Security Roles will only work against AD Identities.  Please use the v2 endpoints instead.
[**DeleteSecurityRolesById**](SecurityRolesApi.md#DeleteSecurityRolesById) | **DELETE** /Security/Roles/{id} | Deletes the security role whose ID is provided.
[**GetSecurityRoles**](SecurityRolesApi.md#GetSecurityRoles) | **GET** /Security/Roles | Returns all security roles according to the provided filter and output parameters.  The v1 endpoints for Security Roles will only work against AD Identities.  Please use the v2 endpoints instead.
[**GetSecurityRolesById**](SecurityRolesApi.md#GetSecurityRolesById) | **GET** /Security/Roles/{id} | Returns a single security role that matches the id.  The v1 endpoints for Security Roles will only work against AD Identities.  Please use the v2 endpoints instead.
[**GetSecurityRolesByIdIdentities**](SecurityRolesApi.md#GetSecurityRolesByIdIdentities) | **GET** /Security/Roles/{id}/Identities | Returns all identities which have the security role that matches the id.  The v1 endpoints for Security Roles will only work against AD Identities.  Please use the v2 endpoints instead.
[**UpdateSecurityRoles**](SecurityRolesApi.md#UpdateSecurityRoles) | **PUT** /Security/Roles | Updates a security role that matches the id.  The v1 endpoints for Security Roles will only work against AD Identities.  Please use the v2 endpoints instead.
[**UpdateSecurityRolesByIdIdentities**](SecurityRolesApi.md#UpdateSecurityRolesByIdIdentities) | **PUT** /Security/Roles/{id}/Identities | Updates the identities which have the security role that matches the id.  The v1 endpoints for Security Roles will only work against AD Identities.  Please use the v2 endpoints instead.



## CreateSecurityRoles

> SecurityLegacySecurityRolesSecurityRoleResponse BuildCreateSecurityRolesRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).SecurityLegacySecurityRolesSecurityRoleCreationRequest(securityLegacySecurityRolesSecurityRoleCreationRequest).Execute()

Adds a new security role to the system.  The v1 endpoints for Security Roles will only work against AD Identities.  Please use the v2 endpoints instead.

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
    securityLegacySecurityRolesSecurityRoleCreationRequest := *openapiclient.NewSecurityLegacySecurityRolesSecurityRoleCreationRequest("Name_example", "Description_example") // SecurityLegacySecurityRolesSecurityRoleCreationRequest | Security Role Creation Request (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolesApi.BuildCreateSecurityRolesRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).SecurityLegacySecurityRolesSecurityRoleCreationRequest(securityLegacySecurityRolesSecurityRoleCreationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolesApi.CreateSecurityRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSecurityRoles`: SecurityLegacySecurityRolesSecurityRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolesApi.CreateSecurityRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecurityRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **securityLegacySecurityRolesSecurityRoleCreationRequest** | [**SecurityLegacySecurityRolesSecurityRoleCreationRequest**](SecurityLegacySecurityRolesSecurityRoleCreationRequest.md) | Security Role Creation Request | 

### Return type

[**SecurityLegacySecurityRolesSecurityRoleResponse**](SecurityLegacySecurityRolesSecurityRoleResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSecurityRolesByIdCopy

> SecurityLegacySecurityRolesSecurityRoleResponse BuildCreateSecurityRolesByIdCopyRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).SecurityLegacySecurityRolesSecurityRoleCopyRequest(securityLegacySecurityRolesSecurityRoleCopyRequest).Execute()

Makes a copy of an existing security role.  The v1 endpoints for Security Roles will only work against AD Identities.  Please use the v2 endpoints instead.

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
    id := int32(56) // int32 | Security role identifier for target role to copy
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    securityLegacySecurityRolesSecurityRoleCopyRequest := *openapiclient.NewSecurityLegacySecurityRolesSecurityRoleCopyRequest() // SecurityLegacySecurityRolesSecurityRoleCopyRequest | New security role's name and description (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolesApi.BuildCreateSecurityRolesByIdCopyRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).SecurityLegacySecurityRolesSecurityRoleCopyRequest(securityLegacySecurityRolesSecurityRoleCopyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolesApi.CreateSecurityRolesByIdCopy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSecurityRolesByIdCopy`: SecurityLegacySecurityRolesSecurityRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolesApi.CreateSecurityRolesByIdCopy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Security role identifier for target role to copy | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecurityRolesByIdCopyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **securityLegacySecurityRolesSecurityRoleCopyRequest** | [**SecurityLegacySecurityRolesSecurityRoleCopyRequest**](SecurityLegacySecurityRolesSecurityRoleCopyRequest.md) | New security role&#39;s name and description | 

### Return type

[**SecurityLegacySecurityRolesSecurityRoleResponse**](SecurityLegacySecurityRolesSecurityRoleResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSecurityRolesById

> BuildDeleteSecurityRolesByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Deletes the security role whose ID is provided.

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
    resp, r, err := apiClient.SecurityRolesApi.BuildDeleteSecurityRolesByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolesApi.DeleteSecurityRolesById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Security role identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSecurityRolesByIdRequest struct via the builder pattern


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


## GetSecurityRoles

> []SecurityLegacySecurityRolesSecurityRoleResponse BuildGetSecurityRolesRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Validate(validate).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns all security roles according to the provided filter and output parameters.  The v1 endpoints for Security Roles will only work against AD Identities.  Please use the v2 endpoints instead.

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
    validate := true // bool |  (optional)
    queryString := "queryString_example" // string |  (optional)
    pageReturned := int32(56) // int32 |  (optional)
    returnLimit := int32(56) // int32 |  (optional)
    sortField := "sortField_example" // string |  (optional)
    sortAscending := openapiclient.Keyfactor.Common.QueryableExtensions.SortOrder(0) // KeyfactorCommonQueryableExtensionsSortOrder |  (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolesApi.BuildGetSecurityRolesRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Validate(validate).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolesApi.GetSecurityRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecurityRoles`: []SecurityLegacySecurityRolesSecurityRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolesApi.GetSecurityRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **validate** | **bool** |  | 
 **queryString** | **string** |  | 
 **pageReturned** | **int32** |  | 
 **returnLimit** | **int32** |  | 
 **sortField** | **string** |  | 
 **sortAscending** | [**KeyfactorCommonQueryableExtensionsSortOrder**](KeyfactorCommonQueryableExtensionsSortOrder.md) |  | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]SecurityLegacySecurityRolesSecurityRoleResponse**](SecurityLegacySecurityRolesSecurityRoleResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurityRolesById

> SecurityLegacySecurityRolesSecurityRoleResponse BuildGetSecurityRolesByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns a single security role that matches the id.  The v1 endpoints for Security Roles will only work against AD Identities.  Please use the v2 endpoints instead.

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
    resp, r, err := apiClient.SecurityRolesApi.BuildGetSecurityRolesByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolesApi.GetSecurityRolesById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecurityRolesById`: SecurityLegacySecurityRolesSecurityRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolesApi.GetSecurityRolesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Security role identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityRolesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**SecurityLegacySecurityRolesSecurityRoleResponse**](SecurityLegacySecurityRolesSecurityRoleResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurityRolesByIdIdentities

> []SecurityLegacySecurityRolesRoleIdentitiesResponse BuildGetSecurityRolesByIdIdentitiesRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns all identities which have the security role that matches the id.  The v1 endpoints for Security Roles will only work against AD Identities.  Please use the v2 endpoints instead.

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
    resp, r, err := apiClient.SecurityRolesApi.BuildGetSecurityRolesByIdIdentitiesRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolesApi.GetSecurityRolesByIdIdentities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecurityRolesByIdIdentities`: []SecurityLegacySecurityRolesRoleIdentitiesResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolesApi.GetSecurityRolesByIdIdentities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Security role identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityRolesByIdIdentitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]SecurityLegacySecurityRolesRoleIdentitiesResponse**](SecurityLegacySecurityRolesRoleIdentitiesResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSecurityRoles

> SecurityLegacySecurityRolesSecurityRoleResponse BuildUpdateSecurityRolesRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).SecurityLegacySecurityRolesSecurityRoleUpdateRequest(securityLegacySecurityRolesSecurityRoleUpdateRequest).Execute()

Updates a security role that matches the id.  The v1 endpoints for Security Roles will only work against AD Identities.  Please use the v2 endpoints instead.

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
    securityLegacySecurityRolesSecurityRoleUpdateRequest := *openapiclient.NewSecurityLegacySecurityRolesSecurityRoleUpdateRequest(int32(123), "Name_example", "Description_example") // SecurityLegacySecurityRolesSecurityRoleUpdateRequest | Security Update Request (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolesApi.BuildUpdateSecurityRolesRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).SecurityLegacySecurityRolesSecurityRoleUpdateRequest(securityLegacySecurityRolesSecurityRoleUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolesApi.UpdateSecurityRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSecurityRoles`: SecurityLegacySecurityRolesSecurityRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolesApi.UpdateSecurityRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSecurityRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **securityLegacySecurityRolesSecurityRoleUpdateRequest** | [**SecurityLegacySecurityRolesSecurityRoleUpdateRequest**](SecurityLegacySecurityRolesSecurityRoleUpdateRequest.md) | Security Update Request | 

### Return type

[**SecurityLegacySecurityRolesSecurityRoleResponse**](SecurityLegacySecurityRolesSecurityRoleResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSecurityRolesByIdIdentities

> []SecurityLegacySecurityRolesRoleIdentitiesResponse BuildUpdateSecurityRolesByIdIdentitiesRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).SecurityLegacySecurityRolesRoleIdentitiesRequest(securityLegacySecurityRolesRoleIdentitiesRequest).Execute()

Updates the identities which have the security role that matches the id.  The v1 endpoints for Security Roles will only work against AD Identities.  Please use the v2 endpoints instead.

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
    securityLegacySecurityRolesRoleIdentitiesRequest := *openapiclient.NewSecurityLegacySecurityRolesRoleIdentitiesRequest() // SecurityLegacySecurityRolesRoleIdentitiesRequest | Role identities request object which contains a list of Identity IDs to remove or add to the role (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolesApi.BuildUpdateSecurityRolesByIdIdentitiesRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).SecurityLegacySecurityRolesRoleIdentitiesRequest(securityLegacySecurityRolesRoleIdentitiesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolesApi.UpdateSecurityRolesByIdIdentities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSecurityRolesByIdIdentities`: []SecurityLegacySecurityRolesRoleIdentitiesResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolesApi.UpdateSecurityRolesByIdIdentities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Security role identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSecurityRolesByIdIdentitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **securityLegacySecurityRolesRoleIdentitiesRequest** | [**SecurityLegacySecurityRolesRoleIdentitiesRequest**](SecurityLegacySecurityRolesRoleIdentitiesRequest.md) | Role identities request object which contains a list of Identity IDs to remove or add to the role | 

### Return type

[**[]SecurityLegacySecurityRolesRoleIdentitiesResponse**](SecurityLegacySecurityRolesRoleIdentitiesResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


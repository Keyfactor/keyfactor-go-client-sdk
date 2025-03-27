# \SecurityRolesApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSecurityRoles**](SecurityRolesApi.md#CreateSecurityRoles) | **POST** /Security/Roles | Adds a new security role to the system.
[**GetSecurityRoles**](SecurityRolesApi.md#GetSecurityRoles) | **GET** /Security/Roles | Returns all security roles according to the provided filter and output parameters.
[**GetSecurityRolesById**](SecurityRolesApi.md#GetSecurityRolesById) | **GET** /Security/Roles/{id} | Returns a single security role that matches the id.
[**UpdateSecurityRoles**](SecurityRolesApi.md#UpdateSecurityRoles) | **PUT** /Security/Roles | Updates an existing security role.



## CreateSecurityRoles

> SecuritySecurityRolesSecurityRoleResponse CreateSecurityRoles(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).SecuritySecurityRolesSecurityRoleCreationRequest(securitySecurityRolesSecurityRoleCreationRequest).Execute()

Adds a new security role to the system.

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
    xKeyfactorApiVersion := "2" // string | Desired version of the api, if not provided defaults to v1 (optional)
    securitySecurityRolesSecurityRoleCreationRequest := *openapiclient.NewSecuritySecurityRolesSecurityRoleCreationRequest("Name_example", "Description_example", "PermissionSetId_example") // SecuritySecurityRolesSecurityRoleCreationRequest | Security Role Creation Request (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolesApi.CreateSecurityRoles(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).SecuritySecurityRolesSecurityRoleCreationRequest(securitySecurityRolesSecurityRoleCreationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolesApi.CreateSecurityRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSecurityRoles`: SecuritySecurityRolesSecurityRoleResponse
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
 **securitySecurityRolesSecurityRoleCreationRequest** | [**SecuritySecurityRolesSecurityRoleCreationRequest**](SecuritySecurityRolesSecurityRoleCreationRequest.md) | Security Role Creation Request | 

### Return type

[**SecuritySecurityRolesSecurityRoleResponse**](SecuritySecurityRolesSecurityRoleResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurityRoles

> []SecuritySecurityRolesSecurityRoleQueryResponse GetSecurityRoles(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns all security roles according to the provided filter and output parameters.

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
    xKeyfactorApiVersion := "2" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolesApi.GetSecurityRoles(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolesApi.GetSecurityRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecurityRoles`: []SecuritySecurityRolesSecurityRoleQueryResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolesApi.GetSecurityRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityRolesRequest struct via the builder pattern


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

[**[]SecuritySecurityRolesSecurityRoleQueryResponse**](SecuritySecurityRolesSecurityRoleQueryResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurityRolesById

> SecuritySecurityRolesSecurityRoleResponse GetSecurityRolesById(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns a single security role that matches the id.

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
    xKeyfactorApiVersion := "2" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolesApi.GetSecurityRolesById(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolesApi.GetSecurityRolesById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecurityRolesById`: SecuritySecurityRolesSecurityRoleResponse
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

[**SecuritySecurityRolesSecurityRoleResponse**](SecuritySecurityRolesSecurityRoleResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSecurityRoles

> SecuritySecurityRolesSecurityRoleResponse UpdateSecurityRoles(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).SecuritySecurityRolesSecurityRoleUpdateRequest(securitySecurityRolesSecurityRoleUpdateRequest).Execute()

Updates an existing security role.

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
    xKeyfactorApiVersion := "2" // string | Desired version of the api, if not provided defaults to v1 (optional)
    securitySecurityRolesSecurityRoleUpdateRequest := *openapiclient.NewSecuritySecurityRolesSecurityRoleUpdateRequest(int32(123), "Name_example", "Description_example", "PermissionSetId_example") // SecuritySecurityRolesSecurityRoleUpdateRequest | Security Role Update Request (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolesApi.UpdateSecurityRoles(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).SecuritySecurityRolesSecurityRoleUpdateRequest(securitySecurityRolesSecurityRoleUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolesApi.UpdateSecurityRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSecurityRoles`: SecuritySecurityRolesSecurityRoleResponse
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
 **securitySecurityRolesSecurityRoleUpdateRequest** | [**SecuritySecurityRolesSecurityRoleUpdateRequest**](SecuritySecurityRolesSecurityRoleUpdateRequest.md) | Security Role Update Request | 

### Return type

[**SecuritySecurityRolesSecurityRoleResponse**](SecuritySecurityRolesSecurityRoleResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


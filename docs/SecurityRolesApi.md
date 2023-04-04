# \SecurityRolesApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SecurityRolesDeleteSecurityRole**](SecurityRolesApi.md#SecurityRolesDeleteSecurityRole) | **Delete** /Security/Roles/{id} | Deletes the security role whose ID is provided.
[**SecurityRolesGetIdentitiesWithRole**](SecurityRolesApi.md#SecurityRolesGetIdentitiesWithRole) | **Get** /Security/Roles/{id}/Identities | Returns all identities which have the security role that matches the id.
[**SecurityRolesUpdateIdentitiesWithRole**](SecurityRolesApi.md#SecurityRolesUpdateIdentitiesWithRole) | **Put** /Security/Roles/{id}/Identities | Updates the identities which have the security role that matches the id.



## SecurityRolesDeleteSecurityRole

> SecurityRolesDeleteSecurityRole(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolesApi.SecurityRolesDeleteSecurityRole(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolesApi.SecurityRolesDeleteSecurityRole``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSecurityRolesDeleteSecurityRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityRolesGetIdentitiesWithRole

> []KeyfactorApiModelsSecurityRolesRoleIdentitiesResponse SecurityRolesGetIdentitiesWithRole(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns all identities which have the security role that matches the id.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolesApi.SecurityRolesGetIdentitiesWithRole(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolesApi.SecurityRolesGetIdentitiesWithRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityRolesGetIdentitiesWithRole`: []KeyfactorApiModelsSecurityRolesRoleIdentitiesResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolesApi.SecurityRolesGetIdentitiesWithRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Security role identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityRolesGetIdentitiesWithRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**[]KeyfactorApiModelsSecurityRolesRoleIdentitiesResponse**](KeyfactorApiModelsSecurityRolesRoleIdentitiesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityRolesUpdateIdentitiesWithRole

> []KeyfactorApiModelsSecurityRolesRoleIdentitiesResponse SecurityRolesUpdateIdentitiesWithRole(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Identities(identities).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Updates the identities which have the security role that matches the id.

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
    identities := *openapiclient.NewKeyfactorApiModelsSecurityRolesRoleIdentitiesRequest() // KeyfactorApiModelsSecurityRolesRoleIdentitiesRequest | Lists of Identity IDs to remove or add to the role
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolesApi.SecurityRolesUpdateIdentitiesWithRole(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Identities(identities).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolesApi.SecurityRolesUpdateIdentitiesWithRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityRolesUpdateIdentitiesWithRole`: []KeyfactorApiModelsSecurityRolesRoleIdentitiesResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolesApi.SecurityRolesUpdateIdentitiesWithRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Security role identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityRolesUpdateIdentitiesWithRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **identities** | [**KeyfactorApiModelsSecurityRolesRoleIdentitiesRequest**](KeyfactorApiModelsSecurityRolesRoleIdentitiesRequest.md) | Lists of Identity IDs to remove or add to the role | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**[]KeyfactorApiModelsSecurityRolesRoleIdentitiesResponse**](KeyfactorApiModelsSecurityRolesRoleIdentitiesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


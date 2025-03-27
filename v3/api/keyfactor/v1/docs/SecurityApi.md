# \SecurityApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSecurityContainersByIdRoles**](SecurityApi.md#CreateSecurityContainersByIdRoles) | **POST** /Security/Containers/{id}/Roles | Edit a certificate store container&#39;s permissions. Reminder: Name field should be left blank.
[**CreateSecurityIdentities**](SecurityApi.md#CreateSecurityIdentities) | **POST** /Security/Identities | Adds a new security identity to the system.
[**DeleteSecurityIdentitiesById**](SecurityApi.md#DeleteSecurityIdentitiesById) | **DELETE** /Security/Identities/{id} | Deletes the security identity whose ID is provided.
[**GetSecurityAuditCollectionsById**](SecurityApi.md#GetSecurityAuditCollectionsById) | **GET** /Security/Audit/Collections/{id} | Gets a list of applicable security permissions for certificate collection
[**GetSecurityContainersByIdRoles**](SecurityApi.md#GetSecurityContainersByIdRoles) | **GET** /Security/Containers/{id}/Roles | Returns all the permissions of a certificate store container through the id
[**GetSecurityIdentities**](SecurityApi.md#GetSecurityIdentities) | **GET** /Security/Identities | Returns all security identities according to the provided filter and output parameters.
[**GetSecurityIdentitiesById**](SecurityApi.md#GetSecurityIdentitiesById) | **GET** /Security/Identities/{id} | Gets an object representing the permissions of the identity associated with the provided identifier.
[**GetSecurityIdentitiesLookup**](SecurityApi.md#GetSecurityIdentitiesLookup) | **GET** /Security/Identities/Lookup | Validates that the identity with the name given exists.
[**GetSecurityMy**](SecurityApi.md#GetSecurityMy) | **GET** /Security/My | Looks at all the roles and global permissions for the user and returns them.



## CreateSecurityContainersByIdRoles

> []CSSCMSDataModelModelsCertificateStoreContainerPermissions CreateSecurityContainersByIdRoles(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsCertificateStoreContainerPermissions(cSSCMSDataModelModelsCertificateStoreContainerPermissions).Execute()

Edit a certificate store container's permissions. Reminder: Name field should be left blank.

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
    id := int32(56) // int32 | Information for the securitycontainer
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    cSSCMSDataModelModelsCertificateStoreContainerPermissions := []openapiclient.CSSCMSDataModelModelsCertificateStoreContainerPermissions{*openapiclient.NewCSSCMSDataModelModelsCertificateStoreContainerPermissions()} // []CSSCMSDataModelModelsCertificateStoreContainerPermissions | Information for the updated security role (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.CreateSecurityContainersByIdRoles(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsCertificateStoreContainerPermissions(cSSCMSDataModelModelsCertificateStoreContainerPermissions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.CreateSecurityContainersByIdRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSecurityContainersByIdRoles`: []CSSCMSDataModelModelsCertificateStoreContainerPermissions
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.CreateSecurityContainersByIdRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Information for the securitycontainer | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecurityContainersByIdRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsCertificateStoreContainerPermissions** | [**[]CSSCMSDataModelModelsCertificateStoreContainerPermissions**](CSSCMSDataModelModelsCertificateStoreContainerPermissions.md) | Information for the updated security role | 

### Return type

[**[]CSSCMSDataModelModelsCertificateStoreContainerPermissions**](CSSCMSDataModelModelsCertificateStoreContainerPermissions.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSecurityIdentities

> SecuritySecurityIdentitiesSecurityIdentityResponse CreateSecurityIdentities(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).SecuritySecurityIdentitiesSecurityIdentityRequest(securitySecurityIdentitiesSecurityIdentityRequest).Execute()

Adds a new security identity to the system.

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
    securitySecurityIdentitiesSecurityIdentityRequest := *openapiclient.NewSecuritySecurityIdentitiesSecurityIdentityRequest("AccountName_example") // SecuritySecurityIdentitiesSecurityIdentityRequest | Security Identity Request (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.CreateSecurityIdentities(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).SecuritySecurityIdentitiesSecurityIdentityRequest(securitySecurityIdentitiesSecurityIdentityRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.CreateSecurityIdentities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSecurityIdentities`: SecuritySecurityIdentitiesSecurityIdentityResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.CreateSecurityIdentities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecurityIdentitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **securitySecurityIdentitiesSecurityIdentityRequest** | [**SecuritySecurityIdentitiesSecurityIdentityRequest**](SecuritySecurityIdentitiesSecurityIdentityRequest.md) | Security Identity Request | 

### Return type

[**SecuritySecurityIdentitiesSecurityIdentityResponse**](SecuritySecurityIdentitiesSecurityIdentityResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSecurityIdentitiesById

> DeleteSecurityIdentitiesById(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Deletes the security identity whose ID is provided.

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
    id := int32(56) // int32 | The ID of the security identity to be deleted.
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.DeleteSecurityIdentitiesById(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.DeleteSecurityIdentitiesById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the security identity to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSecurityIdentitiesByIdRequest struct via the builder pattern


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


## GetSecurityAuditCollectionsById

> CertificateCollectionsCertificateCollectionPermissionsResponse GetSecurityAuditCollectionsById(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Gets a list of applicable security permissions for certificate collection

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
    id := int32(56) // int32 | The certificate collection
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.GetSecurityAuditCollectionsById(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.GetSecurityAuditCollectionsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecurityAuditCollectionsById`: CertificateCollectionsCertificateCollectionPermissionsResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.GetSecurityAuditCollectionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The certificate collection | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityAuditCollectionsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**CertificateCollectionsCertificateCollectionPermissionsResponse**](CertificateCollectionsCertificateCollectionPermissionsResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurityContainersByIdRoles

> []CSSCMSDataModelModelsCertificateStoreContainerPermissions GetSecurityContainersByIdRoles(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns all the permissions of a certificate store container through the id

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
    id := int32(56) // int32 | Information for the updated container
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.GetSecurityContainersByIdRoles(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.GetSecurityContainersByIdRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecurityContainersByIdRoles`: []CSSCMSDataModelModelsCertificateStoreContainerPermissions
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.GetSecurityContainersByIdRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Information for the updated container | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityContainersByIdRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]CSSCMSDataModelModelsCertificateStoreContainerPermissions**](CSSCMSDataModelModelsCertificateStoreContainerPermissions.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurityIdentities

> []SecuritySecurityIdentitiesSecurityIdentityResponse GetSecurityIdentities(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).Validate(validate).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns all security identities according to the provided filter and output parameters.

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
    pageReturned := int32(56) // int32 |  (optional)
    returnLimit := int32(56) // int32 |  (optional)
    sortField := "sortField_example" // string |  (optional)
    sortAscending := openapiclient.Keyfactor.Common.QueryableExtensions.SortOrder(0) // KeyfactorCommonQueryableExtensionsSortOrder |  (optional)
    validate := true // bool |  (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.GetSecurityIdentities(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).Validate(validate).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.GetSecurityIdentities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecurityIdentities`: []SecuritySecurityIdentitiesSecurityIdentityResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.GetSecurityIdentities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityIdentitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **pageReturned** | **int32** |  | 
 **returnLimit** | **int32** |  | 
 **sortField** | **string** |  | 
 **sortAscending** | [**KeyfactorCommonQueryableExtensionsSortOrder**](KeyfactorCommonQueryableExtensionsSortOrder.md) |  | 
 **validate** | **bool** |  | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]SecuritySecurityIdentitiesSecurityIdentityResponse**](SecuritySecurityIdentitiesSecurityIdentityResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurityIdentitiesById

> SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse GetSecurityIdentitiesById(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Gets an object representing the permissions of the identity associated with the provided identifier.

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
    id := int32(56) // int32 | The identifier of the security identity
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.GetSecurityIdentitiesById(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.GetSecurityIdentitiesById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecurityIdentitiesById`: SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.GetSecurityIdentitiesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The identifier of the security identity | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityIdentitiesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse**](SecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurityIdentitiesLookup

> SecuritySecurityIdentitiesSecurityIdentityLookupResponse GetSecurityIdentitiesLookup(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).AccountName(accountName).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Validates that the identity with the name given exists.

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
    accountName := "accountName_example" // string | The name of an identity we wish to check. (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.GetSecurityIdentitiesLookup(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).AccountName(accountName).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.GetSecurityIdentitiesLookup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecurityIdentitiesLookup`: SecuritySecurityIdentitiesSecurityIdentityLookupResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.GetSecurityIdentitiesLookup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityIdentitiesLookupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **accountName** | **string** | The name of an identity we wish to check. | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**SecuritySecurityIdentitiesSecurityIdentityLookupResponse**](SecuritySecurityIdentitiesSecurityIdentityLookupResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurityMy

> SecurityLegacySecurityRolesSecurityMyResponse GetSecurityMy(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Looks at all the roles and global permissions for the user and returns them.

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

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.GetSecurityMy(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.GetSecurityMy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecurityMy`: SecurityLegacySecurityRolesSecurityMyResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.GetSecurityMy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityMyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**SecurityLegacySecurityRolesSecurityMyResponse**](SecurityLegacySecurityRolesSecurityMyResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


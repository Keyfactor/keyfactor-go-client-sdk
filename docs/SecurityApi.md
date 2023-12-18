# \SecurityApi

All URIs are relative to */Keyfactor/API*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SecurityAuditCollectionsIdGet**](SecurityApi.md#SecurityAuditCollectionsIdGet) | **Get** /Security/Audit/Collections/{id} | Gets a list of applicable security permissions for certificate collection
[**SecurityContainersIdRolesGet**](SecurityApi.md#SecurityContainersIdRolesGet) | **Get** /Security/Containers/{id}/Roles | Returns all the permissions of a certificate store container through the id
[**SecurityContainersIdRolesPost**](SecurityApi.md#SecurityContainersIdRolesPost) | **Post** /Security/Containers/{id}/Roles | Edit a certificate store container&#39;s permissions. Reminder: Name field should be left blank.
[**SecurityIdentitiesGet**](SecurityApi.md#SecurityIdentitiesGet) | **Get** /Security/Identities | Returns all security identities according to the provided filter and output parameters.
[**SecurityIdentitiesIdDelete**](SecurityApi.md#SecurityIdentitiesIdDelete) | **Delete** /Security/Identities/{id} | Deletes the security identity whose ID is provided.
[**SecurityIdentitiesIdGet**](SecurityApi.md#SecurityIdentitiesIdGet) | **Get** /Security/Identities/{id} | Gets an object representing the permissions of the identity associated with the provided identifier.
[**SecurityIdentitiesLookupGet**](SecurityApi.md#SecurityIdentitiesLookupGet) | **Get** /Security/Identities/Lookup | Validates that the identity with the name given exists.
[**SecurityIdentitiesPost**](SecurityApi.md#SecurityIdentitiesPost) | **Post** /Security/Identities | Adds a new security identity to the system.
[**SecurityMyGet**](SecurityApi.md#SecurityMyGet) | **Get** /Security/My | Looks at all the roles and global permissions for the user and returns them.



## SecurityAuditCollectionsIdGet

> KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponse SecurityAuditCollectionsIdGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.SecurityAuditCollectionsIdGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.SecurityAuditCollectionsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityAuditCollectionsIdGet`: KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.SecurityAuditCollectionsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The certificate collection | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityAuditCollectionsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponse**](KeyfactorWebKeyfactorApiModelsCertificateCollectionsCertificateCollectionPermissionsResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityContainersIdRolesGet

> []CSSCMSDataModelModelsCertificateStoreContainerPermissions SecurityContainersIdRolesGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.SecurityContainersIdRolesGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.SecurityContainersIdRolesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityContainersIdRolesGet`: []CSSCMSDataModelModelsCertificateStoreContainerPermissions
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.SecurityContainersIdRolesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Information for the updated container | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityContainersIdRolesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
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


## SecurityContainersIdRolesPost

> []CSSCMSDataModelModelsCertificateStoreContainerPermissions SecurityContainersIdRolesPost(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsCertificateStoreContainerPermissions(cSSCMSDataModelModelsCertificateStoreContainerPermissions).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    cSSCMSDataModelModelsCertificateStoreContainerPermissions := []openapiclient.CSSCMSDataModelModelsCertificateStoreContainerPermissions{*openapiclient.NewCSSCMSDataModelModelsCertificateStoreContainerPermissions()} // []CSSCMSDataModelModelsCertificateStoreContainerPermissions | Information for the updated security role (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.SecurityContainersIdRolesPost(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsCertificateStoreContainerPermissions(cSSCMSDataModelModelsCertificateStoreContainerPermissions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.SecurityContainersIdRolesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityContainersIdRolesPost`: []CSSCMSDataModelModelsCertificateStoreContainerPermissions
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.SecurityContainersIdRolesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Information for the securitycontainer | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityContainersIdRolesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsCertificateStoreContainerPermissions** | [**[]CSSCMSDataModelModelsCertificateStoreContainerPermissions**](CSSCMSDataModelModelsCertificateStoreContainerPermissions.md) | Information for the updated security role | 

### Return type

[**[]CSSCMSDataModelModelsCertificateStoreContainerPermissions**](CSSCMSDataModelModelsCertificateStoreContainerPermissions.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patchjson, application/json, text/json, application/*json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityIdentitiesGet

> []KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentitiesSecurityIdentityResponse SecurityIdentitiesGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).Validate(validate).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    pageReturned := int32(56) // int32 |  (optional)
    returnLimit := int32(56) // int32 |  (optional)
    sortField := "sortField_example" // string |  (optional)
    sortAscending := openapiclient.Keyfactor.Common.QueryableExtensionsSortOrder(0) // KeyfactorCommonQueryableExtensionsSortOrder |  (optional)
    validate := true // bool |  (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.SecurityIdentitiesGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).Validate(validate).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.SecurityIdentitiesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityIdentitiesGet`: []KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentitiesSecurityIdentityResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.SecurityIdentitiesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecurityIdentitiesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **pageReturned** | **int32** |  | 
 **returnLimit** | **int32** |  | 
 **sortField** | **string** |  | 
 **sortAscending** | [**KeyfactorCommonQueryableExtensionsSortOrder**](KeyfactorCommonQueryableExtensionsSortOrder.md) |  | 
 **validate** | **bool** |  | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentitiesSecurityIdentityResponse**](KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentitiesSecurityIdentityResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityIdentitiesIdDelete

> SecurityIdentitiesIdDelete(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.SecurityIdentitiesIdDelete(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.SecurityIdentitiesIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSecurityIdentitiesIdDeleteRequest struct via the builder pattern


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


## SecurityIdentitiesIdGet

> KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse SecurityIdentitiesIdGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.SecurityIdentitiesIdGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.SecurityIdentitiesIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityIdentitiesIdGet`: KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.SecurityIdentitiesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The identifier of the security identity | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityIdentitiesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse**](KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentityPermissionsSecurityIdentityPermissionsResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityIdentitiesLookupGet

> KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentitiesSecurityIdentityLookupResponse SecurityIdentitiesLookupGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).AccountName(accountName).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    accountName := "accountName_example" // string | The name of an identity we wish to check. (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.SecurityIdentitiesLookupGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).AccountName(accountName).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.SecurityIdentitiesLookupGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityIdentitiesLookupGet`: KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentitiesSecurityIdentityLookupResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.SecurityIdentitiesLookupGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecurityIdentitiesLookupGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **accountName** | **string** | The name of an identity we wish to check. | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentitiesSecurityIdentityLookupResponse**](KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentitiesSecurityIdentityLookupResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityIdentitiesPost

> KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentitiesSecurityIdentityResponse SecurityIdentitiesPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentitiesSecurityIdentityRequest(keyfactorWebKeyfactorApiModelsSecuritySecurityIdentitiesSecurityIdentityRequest).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsSecuritySecurityIdentitiesSecurityIdentityRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsSecuritySecurityIdentitiesSecurityIdentityRequest("AccountName_example") // KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentitiesSecurityIdentityRequest | Security Identity Request (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.SecurityIdentitiesPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentitiesSecurityIdentityRequest(keyfactorWebKeyfactorApiModelsSecuritySecurityIdentitiesSecurityIdentityRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.SecurityIdentitiesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityIdentitiesPost`: KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentitiesSecurityIdentityResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.SecurityIdentitiesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecurityIdentitiesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsSecuritySecurityIdentitiesSecurityIdentityRequest** | [**KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentitiesSecurityIdentityRequest**](KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentitiesSecurityIdentityRequest.md) | Security Identity Request | 

### Return type

[**KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentitiesSecurityIdentityResponse**](KeyfactorWebKeyfactorApiModelsSecuritySecurityIdentitiesSecurityIdentityResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patchjson, application/json, text/json, application/*json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityMyGet

> KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityMyResponse SecurityMyGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.SecurityMyGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.SecurityMyGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityMyGet`: KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityMyResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.SecurityMyGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecurityMyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityMyResponse**](KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityMyResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


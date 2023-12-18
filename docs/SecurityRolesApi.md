# \SecurityRolesApi

All URIs are relative to */Keyfactor/API*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SecurityRolesGet**](SecurityRolesApi.md#SecurityRolesGet) | **Get** /Security/Roles | Returns all security roles according to the provided filter and output parameters.  The v1 endpoints for Security Roles will only work against AD Identities.  Please use the v2 endpoints instead.
[**SecurityRolesIdCopyPost**](SecurityRolesApi.md#SecurityRolesIdCopyPost) | **Post** /Security/Roles/{id}/Copy | Makes a copy of an existing security role.  The v1 endpoints for Security Roles will only work against AD Identities.  Please use the v2 endpoints instead.
[**SecurityRolesIdDelete**](SecurityRolesApi.md#SecurityRolesIdDelete) | **Delete** /Security/Roles/{id} | Deletes the security role whose ID is provided.
[**SecurityRolesIdGet**](SecurityRolesApi.md#SecurityRolesIdGet) | **Get** /Security/Roles/{id} | Returns a single security role that matches the id.  The v1 endpoints for Security Roles will only work against AD Identities.  Please use the v2 endpoints instead.
[**SecurityRolesIdIdentitiesGet**](SecurityRolesApi.md#SecurityRolesIdIdentitiesGet) | **Get** /Security/Roles/{id}/Identities | Returns all identities which have the security role that matches the id.  The v1 endpoints for Security Roles will only work against AD Identities.  Please use the v2 endpoints instead.
[**SecurityRolesIdIdentitiesPut**](SecurityRolesApi.md#SecurityRolesIdIdentitiesPut) | **Put** /Security/Roles/{id}/Identities | Updates the identities which have the security role that matches the id.  The v1 endpoints for Security Roles will only work against AD Identities.  Please use the v2 endpoints instead.
[**SecurityRolesPost**](SecurityRolesApi.md#SecurityRolesPost) | **Post** /Security/Roles | Adds a new security role to the system.  The v1 endpoints for Security Roles will only work against AD Identities.  Please use the v2 endpoints instead.
[**SecurityRolesPut**](SecurityRolesApi.md#SecurityRolesPut) | **Put** /Security/Roles | Updates a security role that matches the id.  The v1 endpoints for Security Roles will only work against AD Identities.  Please use the v2 endpoints instead.



## SecurityRolesGet

> []KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse SecurityRolesGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Validate(validate).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    validate := true // bool |  (optional)
    queryString := "queryString_example" // string |  (optional)
    pageReturned := int32(56) // int32 |  (optional)
    returnLimit := int32(56) // int32 |  (optional)
    sortField := "sortField_example" // string |  (optional)
    sortAscending := openapiclient.Keyfactor.Common.QueryableExtensionsSortOrder(0) // KeyfactorCommonQueryableExtensionsSortOrder |  (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolesApi.SecurityRolesGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Validate(validate).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolesApi.SecurityRolesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityRolesGet`: []KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolesApi.SecurityRolesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecurityRolesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **validate** | **bool** |  | 
 **queryString** | **string** |  | 
 **pageReturned** | **int32** |  | 
 **returnLimit** | **int32** |  | 
 **sortField** | **string** |  | 
 **sortAscending** | [**KeyfactorCommonQueryableExtensionsSortOrder**](KeyfactorCommonQueryableExtensionsSortOrder.md) |  | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse**](KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityRolesIdCopyPost

> KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse SecurityRolesIdCopyPost(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Body(body).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    body := map[string]interface{}{ ... } // map[string]interface{} | New security role's name and description (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolesApi.SecurityRolesIdCopyPost(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolesApi.SecurityRolesIdCopyPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityRolesIdCopyPost`: KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolesApi.SecurityRolesIdCopyPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Security role identifier for target role to copy | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityRolesIdCopyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **body** | **map[string]interface{}** | New security role&#39;s name and description | 

### Return type

[**KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse**](KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patchjson, application/json, text/json, application/*json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityRolesIdDelete

> SecurityRolesIdDelete(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolesApi.SecurityRolesIdDelete(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolesApi.SecurityRolesIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSecurityRolesIdDeleteRequest struct via the builder pattern


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


## SecurityRolesIdGet

> KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse SecurityRolesIdGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolesApi.SecurityRolesIdGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolesApi.SecurityRolesIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityRolesIdGet`: KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolesApi.SecurityRolesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Security role identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityRolesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse**](KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityRolesIdIdentitiesGet

> []KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesRoleIdentitiesResponse SecurityRolesIdIdentitiesGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolesApi.SecurityRolesIdIdentitiesGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolesApi.SecurityRolesIdIdentitiesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityRolesIdIdentitiesGet`: []KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesRoleIdentitiesResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolesApi.SecurityRolesIdIdentitiesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Security role identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityRolesIdIdentitiesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesRoleIdentitiesResponse**](KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesRoleIdentitiesResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityRolesIdIdentitiesPut

> []KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesRoleIdentitiesResponse SecurityRolesIdIdentitiesPut(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesRoleIdentitiesRequest(keyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesRoleIdentitiesRequest).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesRoleIdentitiesRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesRoleIdentitiesRequest() // KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesRoleIdentitiesRequest | Role identities request object which contains a list of Identity IDs to remove or add to the role (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolesApi.SecurityRolesIdIdentitiesPut(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesRoleIdentitiesRequest(keyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesRoleIdentitiesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolesApi.SecurityRolesIdIdentitiesPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityRolesIdIdentitiesPut`: []KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesRoleIdentitiesResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolesApi.SecurityRolesIdIdentitiesPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Security role identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityRolesIdIdentitiesPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesRoleIdentitiesRequest** | [**KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesRoleIdentitiesRequest**](KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesRoleIdentitiesRequest.md) | Role identities request object which contains a list of Identity IDs to remove or add to the role | 

### Return type

[**[]KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesRoleIdentitiesResponse**](KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesRoleIdentitiesResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patchjson, application/json, text/json, application/*json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityRolesPost

> KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse SecurityRolesPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest(keyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest("Name_example", "Description_example") // KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest | Security Role Creation Request (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolesApi.SecurityRolesPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest(keyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolesApi.SecurityRolesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityRolesPost`: KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolesApi.SecurityRolesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecurityRolesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest** | [**KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest**](KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest.md) | Security Role Creation Request | 

### Return type

[**KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse**](KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patchjson, application/json, text/json, application/*json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityRolesPut

> KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse SecurityRolesPut(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest(keyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest(int32(123), "Name_example", "Description_example") // KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest | Security Update Request (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityRolesApi.SecurityRolesPut(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest(keyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityRolesApi.SecurityRolesPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityRolesPut`: KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityRolesApi.SecurityRolesPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecurityRolesPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest** | [**KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest**](KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest.md) | Security Update Request | 

### Return type

[**KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse**](KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patchjson, application/json, text/json, application/*json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


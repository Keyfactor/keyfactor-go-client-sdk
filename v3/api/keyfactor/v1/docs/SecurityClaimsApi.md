# \SecurityClaimsApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSecurityClaims**](SecurityClaimsApi.md#CreateSecurityClaims) | **POST** /Security/Claims | Adds a new claim definition to the system.
[**DeleteSecurityClaimsById**](SecurityClaimsApi.md#DeleteSecurityClaimsById) | **DELETE** /Security/Claims/{id} | Removes a claim definition from the system.
[**GetSecurityClaims**](SecurityClaimsApi.md#GetSecurityClaims) | **GET** /Security/Claims | Returns all claim definitions according to the provided filter and output parameters.
[**GetSecurityClaimsById**](SecurityClaimsApi.md#GetSecurityClaimsById) | **GET** /Security/Claims/{id} | Returns a single claim definition that matches the id.
[**GetSecurityClaimsRoles**](SecurityClaimsApi.md#GetSecurityClaimsRoles) | **GET** /Security/Claims/Roles | Returns a list of roles granted by the claim with the provided id.
[**UpdateSecurityClaims**](SecurityClaimsApi.md#UpdateSecurityClaims) | **PUT** /Security/Claims | Updates an existing claim definition.



## CreateSecurityClaims

> SecurityRoleClaimDefinitionsRoleClaimDefinitionResponse NewCreateSecurityClaimsRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).SecurityRoleClaimDefinitionsRoleClaimDefinitionCreationRequest(securityRoleClaimDefinitionsRoleClaimDefinitionCreationRequest).Execute()

Adds a new claim definition to the system.

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
    securityRoleClaimDefinitionsRoleClaimDefinitionCreationRequest := *openapiclient.NewSecurityRoleClaimDefinitionsRoleClaimDefinitionCreationRequest(openapiclient.CSS.CMS.Core.Enums.ClaimType(0), "ClaimValue_example", "ProviderAuthenticationScheme_example", "Description_example") // SecurityRoleClaimDefinitionsRoleClaimDefinitionCreationRequest | Claim Definition Creation Request (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityClaimsApi.NewCreateSecurityClaimsRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).SecurityRoleClaimDefinitionsRoleClaimDefinitionCreationRequest(securityRoleClaimDefinitionsRoleClaimDefinitionCreationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityClaimsApi.CreateSecurityClaims``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSecurityClaims`: SecurityRoleClaimDefinitionsRoleClaimDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityClaimsApi.CreateSecurityClaims`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecurityClaimsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **securityRoleClaimDefinitionsRoleClaimDefinitionCreationRequest** | [**SecurityRoleClaimDefinitionsRoleClaimDefinitionCreationRequest**](SecurityRoleClaimDefinitionsRoleClaimDefinitionCreationRequest.md) | Claim Definition Creation Request | 

### Return type

[**SecurityRoleClaimDefinitionsRoleClaimDefinitionResponse**](SecurityRoleClaimDefinitionsRoleClaimDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSecurityClaimsById

> NewDeleteSecurityClaimsByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Removes a claim definition from the system.

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
    id := int32(56) // int32 | the Keyfactor identifier of the claim definition to delete
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityClaimsApi.NewDeleteSecurityClaimsByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityClaimsApi.DeleteSecurityClaimsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | the Keyfactor identifier of the claim definition to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSecurityClaimsByIdRequest struct via the builder pattern


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


## GetSecurityClaims

> []SecurityRoleClaimDefinitionsRoleClaimDefinitionQueryResponse NewGetSecurityClaimsRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns all claim definitions according to the provided filter and output parameters.

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
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityClaimsApi.NewGetSecurityClaimsRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityClaimsApi.GetSecurityClaims``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecurityClaims`: []SecurityRoleClaimDefinitionsRoleClaimDefinitionQueryResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityClaimsApi.GetSecurityClaims`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityClaimsRequest struct via the builder pattern


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

[**[]SecurityRoleClaimDefinitionsRoleClaimDefinitionQueryResponse**](SecurityRoleClaimDefinitionsRoleClaimDefinitionQueryResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurityClaimsById

> SecurityRoleClaimDefinitionsRoleClaimDefinitionResponse NewGetSecurityClaimsByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns a single claim definition that matches the id.

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
    id := int32(56) // int32 | claim definition identifier
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityClaimsApi.NewGetSecurityClaimsByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityClaimsApi.GetSecurityClaimsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecurityClaimsById`: SecurityRoleClaimDefinitionsRoleClaimDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityClaimsApi.GetSecurityClaimsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | claim definition identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityClaimsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**SecurityRoleClaimDefinitionsRoleClaimDefinitionResponse**](SecurityRoleClaimDefinitionsRoleClaimDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurityClaimsRoles

> []SecurityRoleClaimDefinitionsSecurityRoleForClaimResponse NewGetSecurityClaimsRolesRequest(ctx).ClaimType(claimType).ClaimValue(claimValue).ProviderAuthenticationScheme(providerAuthenticationScheme).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns a list of roles granted by the claim with the provided id.

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
    claimType := openapiclient.CSS.CMS.Core.Enums.ClaimType(0) // CSSCMSCoreEnumsClaimType | 
    claimValue := "claimValue_example" // string | 
    providerAuthenticationScheme := "providerAuthenticationScheme_example" // string | 
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityClaimsApi.NewGetSecurityClaimsRolesRequest(context.Background()).ClaimType(claimType).ClaimValue(claimValue).ProviderAuthenticationScheme(providerAuthenticationScheme).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityClaimsApi.GetSecurityClaimsRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecurityClaimsRoles`: []SecurityRoleClaimDefinitionsSecurityRoleForClaimResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityClaimsApi.GetSecurityClaimsRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityClaimsRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **claimType** | [**CSSCMSCoreEnumsClaimType**](CSSCMSCoreEnumsClaimType.md) |  | 
 **claimValue** | **string** |  | 
 **providerAuthenticationScheme** | **string** |  | 
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]SecurityRoleClaimDefinitionsSecurityRoleForClaimResponse**](SecurityRoleClaimDefinitionsSecurityRoleForClaimResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSecurityClaims

> SecurityRoleClaimDefinitionsRoleClaimDefinitionResponse NewUpdateSecurityClaimsRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).SecurityRoleClaimDefinitionsRoleClaimDefinitionUpdateRequest(securityRoleClaimDefinitionsRoleClaimDefinitionUpdateRequest).Execute()

Updates an existing claim definition.

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
    securityRoleClaimDefinitionsRoleClaimDefinitionUpdateRequest := *openapiclient.NewSecurityRoleClaimDefinitionsRoleClaimDefinitionUpdateRequest(int32(123), "Description_example") // SecurityRoleClaimDefinitionsRoleClaimDefinitionUpdateRequest | Claim Definition Update Request (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityClaimsApi.NewUpdateSecurityClaimsRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).SecurityRoleClaimDefinitionsRoleClaimDefinitionUpdateRequest(securityRoleClaimDefinitionsRoleClaimDefinitionUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityClaimsApi.UpdateSecurityClaims``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSecurityClaims`: SecurityRoleClaimDefinitionsRoleClaimDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityClaimsApi.UpdateSecurityClaims`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSecurityClaimsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **securityRoleClaimDefinitionsRoleClaimDefinitionUpdateRequest** | [**SecurityRoleClaimDefinitionsRoleClaimDefinitionUpdateRequest**](SecurityRoleClaimDefinitionsRoleClaimDefinitionUpdateRequest.md) | Claim Definition Update Request | 

### Return type

[**SecurityRoleClaimDefinitionsRoleClaimDefinitionResponse**](SecurityRoleClaimDefinitionsRoleClaimDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


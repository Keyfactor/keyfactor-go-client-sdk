# \UserApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSSHUsers**](UserApi.md#CreateSSHUsers) | **POST** /SSH/Users | Creates a new SSH User.
[**CreateSSHUsersAccess**](UserApi.md#CreateSSHUsersAccess) | **POST** /SSH/Users/Access | Updates logon access for a user
[**DeleteSSHUsersById**](UserApi.md#DeleteSSHUsersById) | **DELETE** /SSH/Users/{id} | Deletes an SSH user.
[**GetSSHUsers**](UserApi.md#GetSSHUsers) | **GET** /SSH/Users | Returns users matching the criteria from the provided query parameters
[**GetSSHUsersById**](UserApi.md#GetSSHUsersById) | **GET** /SSH/Users/{id} | Looks up information about an existing SSH user.
[**UpdateSSHUsers**](UserApi.md#UpdateSSHUsers) | **PUT** /SSH/Users | Updates information about a given user.



## CreateSSHUsers

> CSSCMSDataModelModelsSSHUsersSshUserResponse BuildCreateSSHUsersRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsSSHUsersSshUserCreationRequest(cSSCMSDataModelModelsSSHUsersSshUserCreationRequest).Execute()

Creates a new SSH User.

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
    cSSCMSDataModelModelsSSHUsersSshUserCreationRequest := *openapiclient.NewCSSCMSDataModelModelsSSHUsersSshUserCreationRequest("Username_example") // CSSCMSDataModelModelsSSHUsersSshUserCreationRequest | SSH user to be created. (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.BuildCreateSSHUsersRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsSSHUsersSshUserCreationRequest(cSSCMSDataModelModelsSSHUsersSshUserCreationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.CreateSSHUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSSHUsers`: CSSCMSDataModelModelsSSHUsersSshUserResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.CreateSSHUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSSHUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsSSHUsersSshUserCreationRequest** | [**CSSCMSDataModelModelsSSHUsersSshUserCreationRequest**](CSSCMSDataModelModelsSSHUsersSshUserCreationRequest.md) | SSH user to be created. | 

### Return type

[**CSSCMSDataModelModelsSSHUsersSshUserResponse**](CSSCMSDataModelModelsSSHUsersSshUserResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSSHUsersAccess

> CSSCMSDataModelModelsSSHUsersSshUserAccessResponse BuildCreateSSHUsersAccessRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsSSHUsersSshUserUpdateRequest(cSSCMSDataModelModelsSSHUsersSshUserUpdateRequest).Execute()

Updates logon access for a user

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
    cSSCMSDataModelModelsSSHUsersSshUserUpdateRequest := *openapiclient.NewCSSCMSDataModelModelsSSHUsersSshUserUpdateRequest(int32(123)) // CSSCMSDataModelModelsSSHUsersSshUserUpdateRequest | Logons to add the existing user (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.BuildCreateSSHUsersAccessRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsSSHUsersSshUserUpdateRequest(cSSCMSDataModelModelsSSHUsersSshUserUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.CreateSSHUsersAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSSHUsersAccess`: CSSCMSDataModelModelsSSHUsersSshUserAccessResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.CreateSSHUsersAccess`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSSHUsersAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsSSHUsersSshUserUpdateRequest** | [**CSSCMSDataModelModelsSSHUsersSshUserUpdateRequest**](CSSCMSDataModelModelsSSHUsersSshUserUpdateRequest.md) | Logons to add the existing user | 

### Return type

[**CSSCMSDataModelModelsSSHUsersSshUserAccessResponse**](CSSCMSDataModelModelsSSHUsersSshUserAccessResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSSHUsersById

> BuildDeleteSSHUsersByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Deletes an SSH user.

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
    id := int32(56) // int32 | The Id of the user to delete.
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.BuildDeleteSSHUsersByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.DeleteSSHUsersById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Id of the user to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSSHUsersByIdRequest struct via the builder pattern


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


## GetSSHUsers

> []CSSCMSDataModelModelsSSHUsersSshUserResponse BuildGetSSHUsersRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).ShowOwnedAccess(showOwnedAccess).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns users matching the criteria from the provided query parameters

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
    showOwnedAccess := true // bool | Whether or not to return only logons that have access to servers the requesting user owns (optional) (default to false)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.BuildGetSSHUsersRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).ShowOwnedAccess(showOwnedAccess).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetSSHUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSSHUsers`: []CSSCMSDataModelModelsSSHUsersSshUserResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.GetSSHUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSSHUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **queryString** | **string** |  | 
 **pageReturned** | **int32** |  | 
 **returnLimit** | **int32** |  | 
 **sortField** | **string** |  | 
 **sortAscending** | [**KeyfactorCommonQueryableExtensionsSortOrder**](KeyfactorCommonQueryableExtensionsSortOrder.md) |  | 
 **showOwnedAccess** | **bool** | Whether or not to return only logons that have access to servers the requesting user owns | [default to false]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]CSSCMSDataModelModelsSSHUsersSshUserResponse**](CSSCMSDataModelModelsSSHUsersSshUserResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSSHUsersById

> CSSCMSDataModelModelsSSHUsersSshUserResponse BuildGetSSHUsersByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Looks up information about an existing SSH user.

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
    id := int32(56) // int32 | The Id of the SSH user to retrieve.
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.BuildGetSSHUsersByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetSSHUsersById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSSHUsersById`: CSSCMSDataModelModelsSSHUsersSshUserResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.GetSSHUsersById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Id of the SSH user to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSSHUsersByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**CSSCMSDataModelModelsSSHUsersSshUserResponse**](CSSCMSDataModelModelsSSHUsersSshUserResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSSHUsers

> CSSCMSDataModelModelsSSHUsersSshUserResponse BuildUpdateSSHUsersRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsSSHUsersSshUserUpdateRequest(cSSCMSDataModelModelsSSHUsersSshUserUpdateRequest).Execute()

Updates information about a given user.

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
    cSSCMSDataModelModelsSSHUsersSshUserUpdateRequest := *openapiclient.NewCSSCMSDataModelModelsSSHUsersSshUserUpdateRequest(int32(123)) // CSSCMSDataModelModelsSSHUsersSshUserUpdateRequest | The new state of the SSH user to update. (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.BuildUpdateSSHUsersRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsSSHUsersSshUserUpdateRequest(cSSCMSDataModelModelsSSHUsersSshUserUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UpdateSSHUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSSHUsers`: CSSCMSDataModelModelsSSHUsersSshUserResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UpdateSSHUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSSHUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsSSHUsersSshUserUpdateRequest** | [**CSSCMSDataModelModelsSSHUsersSshUserUpdateRequest**](CSSCMSDataModelModelsSSHUsersSshUserUpdateRequest.md) | The new state of the SSH user to update. | 

### Return type

[**CSSCMSDataModelModelsSSHUsersSshUserResponse**](CSSCMSDataModelModelsSSHUsersSshUserResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


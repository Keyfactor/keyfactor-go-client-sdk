# \UserApi

All URIs are relative to */Keyfactor/API*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SSHUsersAccessPost**](UserApi.md#SSHUsersAccessPost) | **Post** /SSH/Users/Access | Updates logon access for a user
[**SSHUsersGet**](UserApi.md#SSHUsersGet) | **Get** /SSH/Users | Returns users matching the criteria from the provided query parameters
[**SSHUsersIdDelete**](UserApi.md#SSHUsersIdDelete) | **Delete** /SSH/Users/{id} | Deletes an SSH user.
[**SSHUsersIdGet**](UserApi.md#SSHUsersIdGet) | **Get** /SSH/Users/{id} | Looks up information about an existing SSH user.
[**SSHUsersPost**](UserApi.md#SSHUsersPost) | **Post** /SSH/Users | Creates a new SSH User.
[**SSHUsersPut**](UserApi.md#SSHUsersPut) | **Put** /SSH/Users | Updates information about a given user.



## SSHUsersAccessPost

> CSSCMSDataModelModelsSSHUsersSshUserAccessResponse SSHUsersAccessPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsSSHUsersSshUserUpdateRequest(cSSCMSDataModelModelsSSHUsersSshUserUpdateRequest).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    cSSCMSDataModelModelsSSHUsersSshUserUpdateRequest := *openapiclient.NewCSSCMSDataModelModelsSSHUsersSshUserUpdateRequest(int32(123)) // CSSCMSDataModelModelsSSHUsersSshUserUpdateRequest | Logons to add the existing user (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.SSHUsersAccessPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsSSHUsersSshUserUpdateRequest(cSSCMSDataModelModelsSSHUsersSshUserUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.SSHUsersAccessPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SSHUsersAccessPost`: CSSCMSDataModelModelsSSHUsersSshUserAccessResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.SSHUsersAccessPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSSHUsersAccessPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsSSHUsersSshUserUpdateRequest** | [**CSSCMSDataModelModelsSSHUsersSshUserUpdateRequest**](CSSCMSDataModelModelsSSHUsersSshUserUpdateRequest.md) | Logons to add the existing user | 

### Return type

[**CSSCMSDataModelModelsSSHUsersSshUserAccessResponse**](CSSCMSDataModelModelsSSHUsersSshUserAccessResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patchjson, application/json, text/json, application/*json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SSHUsersGet

> []CSSCMSDataModelModelsSSHUsersSshUserResponse SSHUsersGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).ShowOwnedAccess(showOwnedAccess).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    queryString := "queryString_example" // string |  (optional)
    pageReturned := int32(56) // int32 |  (optional)
    returnLimit := int32(56) // int32 |  (optional)
    sortField := "sortField_example" // string |  (optional)
    sortAscending := openapiclient.Keyfactor.Common.QueryableExtensionsSortOrder(0) // KeyfactorCommonQueryableExtensionsSortOrder |  (optional)
    showOwnedAccess := true // bool | Whether or not to return only logons that have access to servers the requesting user owns (optional) (default to false)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.SSHUsersGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).ShowOwnedAccess(showOwnedAccess).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.SSHUsersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SSHUsersGet`: []CSSCMSDataModelModelsSSHUsersSshUserResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.SSHUsersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSSHUsersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
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


## SSHUsersIdDelete

> SSHUsersIdDelete(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.SSHUsersIdDelete(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.SSHUsersIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSSHUsersIdDeleteRequest struct via the builder pattern


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


## SSHUsersIdGet

> CSSCMSDataModelModelsSSHUsersSshUserResponse SSHUsersIdGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.SSHUsersIdGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.SSHUsersIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SSHUsersIdGet`: CSSCMSDataModelModelsSSHUsersSshUserResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.SSHUsersIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Id of the SSH user to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSSHUsersIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
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


## SSHUsersPost

> CSSCMSDataModelModelsSSHUsersSshUserResponse SSHUsersPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsSSHUsersSshUserCreationRequest(cSSCMSDataModelModelsSSHUsersSshUserCreationRequest).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    cSSCMSDataModelModelsSSHUsersSshUserCreationRequest := *openapiclient.NewCSSCMSDataModelModelsSSHUsersSshUserCreationRequest("Username_example") // CSSCMSDataModelModelsSSHUsersSshUserCreationRequest | SSH user to be created. (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.SSHUsersPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsSSHUsersSshUserCreationRequest(cSSCMSDataModelModelsSSHUsersSshUserCreationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.SSHUsersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SSHUsersPost`: CSSCMSDataModelModelsSSHUsersSshUserResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.SSHUsersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSSHUsersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsSSHUsersSshUserCreationRequest** | [**CSSCMSDataModelModelsSSHUsersSshUserCreationRequest**](CSSCMSDataModelModelsSSHUsersSshUserCreationRequest.md) | SSH user to be created. | 

### Return type

[**CSSCMSDataModelModelsSSHUsersSshUserResponse**](CSSCMSDataModelModelsSSHUsersSshUserResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patchjson, application/json, text/json, application/*json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SSHUsersPut

> CSSCMSDataModelModelsSSHUsersSshUserResponse SSHUsersPut(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsSSHUsersSshUserUpdateRequest(cSSCMSDataModelModelsSSHUsersSshUserUpdateRequest).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    cSSCMSDataModelModelsSSHUsersSshUserUpdateRequest := *openapiclient.NewCSSCMSDataModelModelsSSHUsersSshUserUpdateRequest(int32(123)) // CSSCMSDataModelModelsSSHUsersSshUserUpdateRequest | The new state of the SSH user to update. (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.SSHUsersPut(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsSSHUsersSshUserUpdateRequest(cSSCMSDataModelModelsSSHUsersSshUserUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.SSHUsersPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SSHUsersPut`: CSSCMSDataModelModelsSSHUsersSshUserResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.SSHUsersPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSSHUsersPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsSSHUsersSshUserUpdateRequest** | [**CSSCMSDataModelModelsSSHUsersSshUserUpdateRequest**](CSSCMSDataModelModelsSSHUsersSshUserUpdateRequest.md) | The new state of the SSH user to update. | 

### Return type

[**CSSCMSDataModelModelsSSHUsersSshUserResponse**](CSSCMSDataModelModelsSSHUsersSshUserResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patchjson, application/json, text/json, application/*json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \PAMLocalEntriesApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePamProvidersLocalProviderIdEntries**](PAMLocalEntriesApi.md#CreatePamProvidersLocalProviderIdEntries) | **POST** /PamProviders/Local/{providerId}/Entries | Creates a new local PAM entry with the associated properties
[**DeletePamProvidersLocalProviderIdEntries**](PAMLocalEntriesApi.md#DeletePamProvidersLocalProviderIdEntries) | **DELETE** /PamProviders/Local/{providerId}/Entries | Deletes a local PAM entry
[**GetPamProvidersLocalProviderIdEntries**](PAMLocalEntriesApi.md#GetPamProvidersLocalProviderIdEntries) | **GET** /PamProviders/Local/{providerId}/Entries | Returns local PAM entries for the given PAM provider according to the provided filter and output parameters
[**UpdatePamProvidersLocalProviderIdEntries**](PAMLocalEntriesApi.md#UpdatePamProvidersLocalProviderIdEntries) | **PUT** /PamProviders/Local/{providerId}/Entries | Updates local PAM entry with the associated properties



## CreatePamProvidersLocalProviderIdEntries

> PAMLocalLocalPAMEntryResponse BuildCreatePamProvidersLocalProviderIdEntriesRequest(ctx, providerId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PAMLocalLocalPAMEntryCreateRequest(pAMLocalLocalPAMEntryCreateRequest).Execute()

Creates a new local PAM entry with the associated properties

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
    providerId := int32(56) // int32 | Keyfactor identifier of the PAM provider
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    pAMLocalLocalPAMEntryCreateRequest := *openapiclient.NewPAMLocalLocalPAMEntryCreateRequest("SecretName_example", "SecretValue_example") // PAMLocalLocalPAMEntryCreateRequest | Local PAM entry properties to be used (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PAMLocalEntriesApi.BuildCreatePamProvidersLocalProviderIdEntriesRequest(context.Background(), providerId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PAMLocalLocalPAMEntryCreateRequest(pAMLocalLocalPAMEntryCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PAMLocalEntriesApi.CreatePamProvidersLocalProviderIdEntries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePamProvidersLocalProviderIdEntries`: PAMLocalLocalPAMEntryResponse
    fmt.Fprintf(os.Stdout, "Response from `PAMLocalEntriesApi.CreatePamProvidersLocalProviderIdEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **int32** | Keyfactor identifier of the PAM provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePamProvidersLocalProviderIdEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **pAMLocalLocalPAMEntryCreateRequest** | [**PAMLocalLocalPAMEntryCreateRequest**](PAMLocalLocalPAMEntryCreateRequest.md) | Local PAM entry properties to be used | 

### Return type

[**PAMLocalLocalPAMEntryResponse**](PAMLocalLocalPAMEntryResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePamProvidersLocalProviderIdEntries

> BuildDeletePamProvidersLocalProviderIdEntriesRequest(ctx, providerId).SecretName(secretName).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Deletes a local PAM entry

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
    providerId := int32(56) // int32 | Keyfactor identifier of the PAM provider
    secretName := "secretName_example" // string | Name of the secret entry to be deleted
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PAMLocalEntriesApi.BuildDeletePamProvidersLocalProviderIdEntriesRequest(context.Background(), providerId).SecretName(secretName).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PAMLocalEntriesApi.DeletePamProvidersLocalProviderIdEntries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **int32** | Keyfactor identifier of the PAM provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePamProvidersLocalProviderIdEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **secretName** | **string** | Name of the secret entry to be deleted | 
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


## GetPamProvidersLocalProviderIdEntries

> []PAMLocalLocalPAMEntryResponse BuildGetPamProvidersLocalProviderIdEntriesRequest(ctx, providerId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns local PAM entries for the given PAM provider according to the provided filter and output parameters

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
    providerId := int32(56) // int32 | Keyfactor identifier of the PAM provider
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    queryString := "queryString_example" // string | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) (optional)
    pageReturned := int32(56) // int32 | The current page within the result set to be returned (optional)
    returnLimit := int32(56) // int32 | Maximum number of records to be returned in a single call (optional)
    sortField := "sortField_example" // string | Field by which the results should be sorted (view results via Management Portal for sortable columns) (optional)
    sortAscending := openapiclient.Keyfactor.Common.QueryableExtensions.SortOrder(0) // KeyfactorCommonQueryableExtensionsSortOrder | Field sort direction [0=ascending, 1=descending] (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PAMLocalEntriesApi.BuildGetPamProvidersLocalProviderIdEntriesRequest(context.Background(), providerId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PAMLocalEntriesApi.GetPamProvidersLocalProviderIdEntries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPamProvidersLocalProviderIdEntries`: []PAMLocalLocalPAMEntryResponse
    fmt.Fprintf(os.Stdout, "Response from `PAMLocalEntriesApi.GetPamProvidersLocalProviderIdEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **int32** | Keyfactor identifier of the PAM provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPamProvidersLocalProviderIdEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **queryString** | **string** | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) | 
 **pageReturned** | **int32** | The current page within the result set to be returned | 
 **returnLimit** | **int32** | Maximum number of records to be returned in a single call | 
 **sortField** | **string** | Field by which the results should be sorted (view results via Management Portal for sortable columns) | 
 **sortAscending** | [**KeyfactorCommonQueryableExtensionsSortOrder**](KeyfactorCommonQueryableExtensionsSortOrder.md) | Field sort direction [0&#x3D;ascending, 1&#x3D;descending] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]PAMLocalLocalPAMEntryResponse**](PAMLocalLocalPAMEntryResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePamProvidersLocalProviderIdEntries

> PAMLocalLocalPAMEntryResponse BuildUpdatePamProvidersLocalProviderIdEntriesRequest(ctx, providerId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PAMLocalLocalPAMEntryUpdateRequest(pAMLocalLocalPAMEntryUpdateRequest).Execute()

Updates local PAM entry with the associated properties

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
    providerId := int32(56) // int32 | Keyfactor identifier of the PAM provider
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    pAMLocalLocalPAMEntryUpdateRequest := *openapiclient.NewPAMLocalLocalPAMEntryUpdateRequest("SecretName_example", "SecretValue_example") // PAMLocalLocalPAMEntryUpdateRequest | Local PAM entry properties to be used (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PAMLocalEntriesApi.BuildUpdatePamProvidersLocalProviderIdEntriesRequest(context.Background(), providerId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PAMLocalLocalPAMEntryUpdateRequest(pAMLocalLocalPAMEntryUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PAMLocalEntriesApi.UpdatePamProvidersLocalProviderIdEntries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePamProvidersLocalProviderIdEntries`: PAMLocalLocalPAMEntryResponse
    fmt.Fprintf(os.Stdout, "Response from `PAMLocalEntriesApi.UpdatePamProvidersLocalProviderIdEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **int32** | Keyfactor identifier of the PAM provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePamProvidersLocalProviderIdEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **pAMLocalLocalPAMEntryUpdateRequest** | [**PAMLocalLocalPAMEntryUpdateRequest**](PAMLocalLocalPAMEntryUpdateRequest.md) | Local PAM entry properties to be used | 

### Return type

[**PAMLocalLocalPAMEntryResponse**](PAMLocalLocalPAMEntryResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


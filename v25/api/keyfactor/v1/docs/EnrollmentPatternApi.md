# \EnrollmentPatternApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEnrollmentPatterns**](EnrollmentPatternApi.md#CreateEnrollmentPatterns) | **POST** /EnrollmentPatterns | Creates a new enrollment pattern with the associated properties
[**DeleteEnrollmentPatternsById**](EnrollmentPatternApi.md#DeleteEnrollmentPatternsById) | **DELETE** /EnrollmentPatterns/{id} | Deletes an enrollment pattern according to the provided Keyfactor identifier
[**GetEnrollmentPatterns**](EnrollmentPatternApi.md#GetEnrollmentPatterns) | **GET** /EnrollmentPatterns | Returns all enrollment patterns according to the provided filter and output parameters
[**GetEnrollmentPatternsById**](EnrollmentPatternApi.md#GetEnrollmentPatternsById) | **GET** /EnrollmentPatterns/{id} | Returns the enrollment pattern associated with the provided id
[**GetEnrollmentPatternsByIdMetadata**](EnrollmentPatternApi.md#GetEnrollmentPatternsByIdMetadata) | **GET** /EnrollmentPatterns/{id}/Metadata | Resolves metadata fields for an enrollment pattern and pattern&#39;s template
[**GetEnrollmentPatternsByIdSettings**](EnrollmentPatternApi.md#GetEnrollmentPatternsByIdSettings) | **GET** /EnrollmentPatterns/{id}/Settings | Gets the settings for the enrollment pattern with the provided id to use during enrollment. The response will be the resolved values for the settings.  If there is an enrollment pattern specific setting, the enrollment pattern specific setting will be used in the response.  If there is not an enrollment pattern specific setting, the global setting will be used in the response.
[**GetEnrollmentPatternsSettings**](EnrollmentPatternApi.md#GetEnrollmentPatternsSettings) | **GET** /EnrollmentPatterns/Settings | Gets the global pattern settings.
[**GetEnrollmentPatternsSubjectParts**](EnrollmentPatternApi.md#GetEnrollmentPatternsSubjectParts) | **GET** /EnrollmentPatterns/SubjectParts | Returns the valid subject parts possible for regular expressions.
[**UpdateEnrollmentPatternsById**](EnrollmentPatternApi.md#UpdateEnrollmentPatternsById) | **PUT** /EnrollmentPatterns/{id} | Updates an enrollment pattern according to the provided properties and Keyfactor identifier
[**UpdateEnrollmentPatternsSettings**](EnrollmentPatternApi.md#UpdateEnrollmentPatternsSettings) | **PUT** /EnrollmentPatterns/Settings | Replaces the existing global enrollment pattern settings.



## CreateEnrollmentPatterns

> EnrollmentPatternsEnrollmentPatternResponse NewCreateEnrollmentPatternsRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).ForceTemplateDefault(forceTemplateDefault).XKeyfactorApiVersion(xKeyfactorApiVersion).EnrollmentPatternsEnrollmentPatternCreateRequest(enrollmentPatternsEnrollmentPatternCreateRequest).Execute()

Creates a new enrollment pattern with the associated properties

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
    forceTemplateDefault := true // bool | Flag to forcibly update current enrollment pattern as template default pattern (optional) (default to false)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    enrollmentPatternsEnrollmentPatternCreateRequest := *openapiclient.NewEnrollmentPatternsEnrollmentPatternCreateRequest(int32(123), "Name_example") // EnrollmentPatternsEnrollmentPatternCreateRequest | Properties of the enrollment pattern to be created (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentPatternApi.NewCreateEnrollmentPatternsRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).ForceTemplateDefault(forceTemplateDefault).XKeyfactorApiVersion(xKeyfactorApiVersion).EnrollmentPatternsEnrollmentPatternCreateRequest(enrollmentPatternsEnrollmentPatternCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentPatternApi.CreateEnrollmentPatterns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEnrollmentPatterns`: EnrollmentPatternsEnrollmentPatternResponse
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentPatternApi.CreateEnrollmentPatterns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnrollmentPatternsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **forceTemplateDefault** | **bool** | Flag to forcibly update current enrollment pattern as template default pattern | [default to false]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **enrollmentPatternsEnrollmentPatternCreateRequest** | [**EnrollmentPatternsEnrollmentPatternCreateRequest**](EnrollmentPatternsEnrollmentPatternCreateRequest.md) | Properties of the enrollment pattern to be created | 

### Return type

[**EnrollmentPatternsEnrollmentPatternResponse**](EnrollmentPatternsEnrollmentPatternResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEnrollmentPatternsById

> NewDeleteEnrollmentPatternsByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Force(force).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Deletes an enrollment pattern according to the provided Keyfactor identifier

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
    id := int32(56) // int32 | Keyfactor identifier of the enrollment pattern
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    force := true // bool | Forces deletion of the enrollment pattern even if associated with certificates (optional) (default to false)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentPatternApi.NewDeleteEnrollmentPatternsByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Force(force).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentPatternApi.DeleteEnrollmentPatternsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Keyfactor identifier of the enrollment pattern | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEnrollmentPatternsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **force** | **bool** | Forces deletion of the enrollment pattern even if associated with certificates | [default to false]
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


## GetEnrollmentPatterns

> []EnrollmentPatternsEnrollmentPatternResponse NewGetEnrollmentPatternsRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns all enrollment patterns according to the provided filter and output parameters

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
    queryString := "queryString_example" // string | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) (optional)
    pageReturned := int32(56) // int32 | The current page within the result set to be returned (optional)
    returnLimit := int32(56) // int32 | Maximum number of records to be returned in a single call (optional)
    sortField := "sortField_example" // string | Field by which the results should be sorted (view results via Management Portal for sortable columns) (optional)
    sortAscending := openapiclient.Keyfactor.Common.QueryableExtensions.SortOrder(0) // KeyfactorCommonQueryableExtensionsSortOrder | Field sort direction [0=ascending, 1=descending] (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentPatternApi.NewGetEnrollmentPatternsRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentPatternApi.GetEnrollmentPatterns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnrollmentPatterns`: []EnrollmentPatternsEnrollmentPatternResponse
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentPatternApi.GetEnrollmentPatterns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEnrollmentPatternsRequest struct via the builder pattern


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

[**[]EnrollmentPatternsEnrollmentPatternResponse**](EnrollmentPatternsEnrollmentPatternResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnrollmentPatternsById

> EnrollmentPatternsEnrollmentPatternResponse NewGetEnrollmentPatternsByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns the enrollment pattern associated with the provided id

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
    id := int32(56) // int32 | Keyfactor identifier of the enrollment pattern
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentPatternApi.NewGetEnrollmentPatternsByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentPatternApi.GetEnrollmentPatternsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnrollmentPatternsById`: EnrollmentPatternsEnrollmentPatternResponse
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentPatternApi.GetEnrollmentPatternsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Keyfactor identifier of the enrollment pattern | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnrollmentPatternsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**EnrollmentPatternsEnrollmentPatternResponse**](EnrollmentPatternsEnrollmentPatternResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnrollmentPatternsByIdMetadata

> []EnrollmentPatternsEnrollmentPatternMetadataFieldResponse NewGetEnrollmentPatternsByIdMetadataRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Resolves metadata fields for an enrollment pattern and pattern's template

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
    id := int32(56) // int32 | The enrollment pattern id to look for
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentPatternApi.NewGetEnrollmentPatternsByIdMetadataRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentPatternApi.GetEnrollmentPatternsByIdMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnrollmentPatternsByIdMetadata`: []EnrollmentPatternsEnrollmentPatternMetadataFieldResponse
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentPatternApi.GetEnrollmentPatternsByIdMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The enrollment pattern id to look for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnrollmentPatternsByIdMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]EnrollmentPatternsEnrollmentPatternMetadataFieldResponse**](EnrollmentPatternsEnrollmentPatternMetadataFieldResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnrollmentPatternsByIdSettings

> EnrollmentPatternsEnrollmentPatternSettingsResponse NewGetEnrollmentPatternsByIdSettingsRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Gets the settings for the enrollment pattern with the provided id to use during enrollment. The response will be the resolved values for the settings.  If there is an enrollment pattern specific setting, the enrollment pattern specific setting will be used in the response.  If there is not an enrollment pattern specific setting, the global setting will be used in the response.

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
    id := int32(56) // int32 | 
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentPatternApi.NewGetEnrollmentPatternsByIdSettingsRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentPatternApi.GetEnrollmentPatternsByIdSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnrollmentPatternsByIdSettings`: EnrollmentPatternsEnrollmentPatternSettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentPatternApi.GetEnrollmentPatternsByIdSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnrollmentPatternsByIdSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**EnrollmentPatternsEnrollmentPatternSettingsResponse**](EnrollmentPatternsEnrollmentPatternSettingsResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnrollmentPatternsSettings

> EnrollmentPatternsEnrollmentPatternSettingsResponse NewGetEnrollmentPatternsSettingsRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Gets the global pattern settings.

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
    resp, r, err := apiClient.EnrollmentPatternApi.NewGetEnrollmentPatternsSettingsRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentPatternApi.GetEnrollmentPatternsSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnrollmentPatternsSettings`: EnrollmentPatternsEnrollmentPatternSettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentPatternApi.GetEnrollmentPatternsSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEnrollmentPatternsSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**EnrollmentPatternsEnrollmentPatternSettingsResponse**](EnrollmentPatternsEnrollmentPatternSettingsResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnrollmentPatternsSubjectParts

> []EnrollmentPatternsValidSubjectPartResponse NewGetEnrollmentPatternsSubjectPartsRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns the valid subject parts possible for regular expressions.

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
    resp, r, err := apiClient.EnrollmentPatternApi.NewGetEnrollmentPatternsSubjectPartsRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentPatternApi.GetEnrollmentPatternsSubjectParts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnrollmentPatternsSubjectParts`: []EnrollmentPatternsValidSubjectPartResponse
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentPatternApi.GetEnrollmentPatternsSubjectParts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEnrollmentPatternsSubjectPartsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]EnrollmentPatternsValidSubjectPartResponse**](EnrollmentPatternsValidSubjectPartResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEnrollmentPatternsById

> EnrollmentPatternsEnrollmentPatternResponse NewUpdateEnrollmentPatternsByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).ForceTemplateDefault(forceTemplateDefault).XKeyfactorApiVersion(xKeyfactorApiVersion).EnrollmentPatternsEnrollmentPatternRequest(enrollmentPatternsEnrollmentPatternRequest).Execute()

Updates an enrollment pattern according to the provided properties and Keyfactor identifier

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
    id := int32(56) // int32 | Keyfactor identifier of the enrollment pattern
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    forceTemplateDefault := true // bool | Flag to forcibly update current enrollment pattern as template default pattern (optional) (default to false)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    enrollmentPatternsEnrollmentPatternRequest := *openapiclient.NewEnrollmentPatternsEnrollmentPatternRequest("Name_example") // EnrollmentPatternsEnrollmentPatternRequest | Properties of the enrollment pattern to be updated (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentPatternApi.NewUpdateEnrollmentPatternsByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).ForceTemplateDefault(forceTemplateDefault).XKeyfactorApiVersion(xKeyfactorApiVersion).EnrollmentPatternsEnrollmentPatternRequest(enrollmentPatternsEnrollmentPatternRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentPatternApi.UpdateEnrollmentPatternsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEnrollmentPatternsById`: EnrollmentPatternsEnrollmentPatternResponse
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentPatternApi.UpdateEnrollmentPatternsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Keyfactor identifier of the enrollment pattern | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEnrollmentPatternsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **forceTemplateDefault** | **bool** | Flag to forcibly update current enrollment pattern as template default pattern | [default to false]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **enrollmentPatternsEnrollmentPatternRequest** | [**EnrollmentPatternsEnrollmentPatternRequest**](EnrollmentPatternsEnrollmentPatternRequest.md) | Properties of the enrollment pattern to be updated | 

### Return type

[**EnrollmentPatternsEnrollmentPatternResponse**](EnrollmentPatternsEnrollmentPatternResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEnrollmentPatternsSettings

> EnrollmentPatternsEnrollmentPatternSettingsResponse NewUpdateEnrollmentPatternsSettingsRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).EnrollmentPatternsGlobalGlobalEnrollmentPatternSettingsRequest(enrollmentPatternsGlobalGlobalEnrollmentPatternSettingsRequest).Execute()

Replaces the existing global enrollment pattern settings.

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
    enrollmentPatternsGlobalGlobalEnrollmentPatternSettingsRequest := *openapiclient.NewEnrollmentPatternsGlobalGlobalEnrollmentPatternSettingsRequest([]openapiclient.EnrollmentPatternsEnrollmentPatternRegexesRequest{*openapiclient.NewEnrollmentPatternsEnrollmentPatternRegexesRequest("SubjectPart_example")}, []openapiclient.EnrollmentPatternsEnrollmentPatternDefaultRequest{*openapiclient.NewEnrollmentPatternsEnrollmentPatternDefaultRequest("SubjectPart_example")}, *openapiclient.NewEnrollmentPatternsGlobalGlobalEnrollmentPatternPolicyRequest(false, false, false, openapiclient.CSS.CMS.Core.Enums.TemplateCertificateOwnerRole(0), *openapiclient.NewEnrollmentPatternsAlgorithmsKeyInfoRequest())) // EnrollmentPatternsGlobalGlobalEnrollmentPatternSettingsRequest | The new global enrollment pattern settings. (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentPatternApi.NewUpdateEnrollmentPatternsSettingsRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).EnrollmentPatternsGlobalGlobalEnrollmentPatternSettingsRequest(enrollmentPatternsGlobalGlobalEnrollmentPatternSettingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentPatternApi.UpdateEnrollmentPatternsSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEnrollmentPatternsSettings`: EnrollmentPatternsEnrollmentPatternSettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentPatternApi.UpdateEnrollmentPatternsSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEnrollmentPatternsSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **enrollmentPatternsGlobalGlobalEnrollmentPatternSettingsRequest** | [**EnrollmentPatternsGlobalGlobalEnrollmentPatternSettingsRequest**](EnrollmentPatternsGlobalGlobalEnrollmentPatternSettingsRequest.md) | The new global enrollment pattern settings. | 

### Return type

[**EnrollmentPatternsEnrollmentPatternSettingsResponse**](EnrollmentPatternsEnrollmentPatternSettingsResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


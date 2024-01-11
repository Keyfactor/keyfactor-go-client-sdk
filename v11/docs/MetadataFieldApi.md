# \MetadataFieldApi

All URIs are relative to */Keyfactor/API*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MetadataFieldsDelete**](MetadataFieldApi.md#MetadataFieldsDelete) | **Delete** /MetadataFields | Deletes multiple persisted metadata field types by their unique ids
[**MetadataFieldsGet**](MetadataFieldApi.md#MetadataFieldsGet) | **Get** /MetadataFields | Returns all metadata field types according to the provided filter and output parameters
[**MetadataFieldsIdDelete**](MetadataFieldApi.md#MetadataFieldsIdDelete) | **Delete** /MetadataFields/{id} | Deletes a persisted metadata field type by its unique id
[**MetadataFieldsIdGet**](MetadataFieldApi.md#MetadataFieldsIdGet) | **Get** /MetadataFields/{id} | Gets a persisted metadata field type by its unique id
[**MetadataFieldsIdInUseGet**](MetadataFieldApi.md#MetadataFieldsIdInUseGet) | **Get** /MetadataFields/{id}/InUse | Determines if a metadata field type associated with the provided identifier is currently in use
[**MetadataFieldsNameGet**](MetadataFieldApi.md#MetadataFieldsNameGet) | **Get** /MetadataFields/{name} | Gets a persisted metadata field type by its unique name
[**MetadataFieldsPost**](MetadataFieldApi.md#MetadataFieldsPost) | **Post** /MetadataFields | Creates a new metadata field type with the given metadata field type properties
[**MetadataFieldsPut**](MetadataFieldApi.md#MetadataFieldsPut) | **Put** /MetadataFields | Updates a persisted metadata field with the given metadata field type



## MetadataFieldsDelete

> MetadataFieldsDelete(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Force(force).XKeyfactorApiVersion(xKeyfactorApiVersion).RequestBody(requestBody).Execute()

Deletes multiple persisted metadata field types by their unique ids



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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    force := true // bool | Forces deletion of the metadata field type even if in-use (optional) (default to false)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    requestBody := []int32{int32(123)} // []int32 | Array of Keyfactor identifiers for metadata field types to be deleted (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataFieldApi.MetadataFieldsDelete(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Force(force).XKeyfactorApiVersion(xKeyfactorApiVersion).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataFieldApi.MetadataFieldsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetadataFieldsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **force** | **bool** | Forces deletion of the metadata field type even if in-use | [default to false]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **requestBody** | **[]int32** | Array of Keyfactor identifiers for metadata field types to be deleted | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetadataFieldsGet

> []ModelsMetadataFieldTypeModel MetadataFieldsGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns all metadata field types according to the provided filter and output parameters



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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    queryString := "queryString_example" // string |  (optional)
    pageReturned := int32(56) // int32 |  (optional)
    returnLimit := int32(56) // int32 |  (optional)
    sortField := "sortField_example" // string |  (optional)
    sortAscending := int32(56) // int32 |  (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataFieldApi.MetadataFieldsGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataFieldApi.MetadataFieldsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MetadataFieldsGet`: []ModelsMetadataFieldTypeModel
    fmt.Fprintf(os.Stdout, "Response from `MetadataFieldApi.MetadataFieldsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetadataFieldsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **queryString** | **string** |  | 
 **pageReturned** | **int32** |  | 
 **returnLimit** | **int32** |  | 
 **sortField** | **string** |  | 
 **sortAscending** | **int32** |  | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]ModelsMetadataFieldTypeModel**](ModelsMetadataFieldTypeModel.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetadataFieldsIdDelete

> MetadataFieldsIdDelete(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Force(force).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Deletes a persisted metadata field type by its unique id

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
    id := int32(56) // int32 | Keyfactor identifier of the metadata field type
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    force := true // bool | Forces deletion of the metadata field type even if in-use (optional) (default to false)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataFieldApi.MetadataFieldsIdDelete(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Force(force).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataFieldApi.MetadataFieldsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Keyfactor identifier of the metadata field type | 

### Other Parameters

Other parameters are passed through a pointer to a apiMetadataFieldsIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **force** | **bool** | Forces deletion of the metadata field type even if in-use | [default to false]
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


## MetadataFieldsIdGet

> ModelsMetadataFieldTypeModel MetadataFieldsIdGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Gets a persisted metadata field type by its unique id



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
    id := int32(56) // int32 | The unique id of the metadata field type
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataFieldApi.MetadataFieldsIdGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataFieldApi.MetadataFieldsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MetadataFieldsIdGet`: ModelsMetadataFieldTypeModel
    fmt.Fprintf(os.Stdout, "Response from `MetadataFieldApi.MetadataFieldsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The unique id of the metadata field type | 

### Other Parameters

Other parameters are passed through a pointer to a apiMetadataFieldsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**ModelsMetadataFieldTypeModel**](ModelsMetadataFieldTypeModel.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetadataFieldsIdInUseGet

> bool MetadataFieldsIdInUseGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Determines if a metadata field type associated with the provided identifier is currently in use

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
    id := int32(56) // int32 | Keyfactor identitifer of the metadata field
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataFieldApi.MetadataFieldsIdInUseGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataFieldApi.MetadataFieldsIdInUseGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MetadataFieldsIdInUseGet`: bool
    fmt.Fprintf(os.Stdout, "Response from `MetadataFieldApi.MetadataFieldsIdInUseGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Keyfactor identitifer of the metadata field | 

### Other Parameters

Other parameters are passed through a pointer to a apiMetadataFieldsIdInUseGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

**bool**

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetadataFieldsNameGet

> ModelsMetadataFieldTypeModel MetadataFieldsNameGet(ctx, name).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Gets a persisted metadata field type by its unique name



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
    name := "name_example" // string | The unique name of the metadata field.
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataFieldApi.MetadataFieldsNameGet(context.Background(), name).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataFieldApi.MetadataFieldsNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MetadataFieldsNameGet`: ModelsMetadataFieldTypeModel
    fmt.Fprintf(os.Stdout, "Response from `MetadataFieldApi.MetadataFieldsNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The unique name of the metadata field. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMetadataFieldsNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**ModelsMetadataFieldTypeModel**](ModelsMetadataFieldTypeModel.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetadataFieldsPost

> KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldResponse MetadataFieldsPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest(keyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest).Execute()

Creates a new metadata field type with the given metadata field type properties



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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest("Name_example", "Description_example", int32(123)) // KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest | Properties of the metadata field type to be created (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataFieldApi.MetadataFieldsPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest(keyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataFieldApi.MetadataFieldsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MetadataFieldsPost`: KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldResponse
    fmt.Fprintf(os.Stdout, "Response from `MetadataFieldApi.MetadataFieldsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetadataFieldsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest** | [**KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest**](KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest.md) | Properties of the metadata field type to be created | 

### Return type

[**KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldResponse**](KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetadataFieldsPut

> KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldResponse MetadataFieldsPut(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest(keyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest).Execute()

Updates a persisted metadata field with the given metadata field type



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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest(int32(123), "Name_example", "Description_example", int32(123), int32(123)) // KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest | Properties of the metadata field type to be updated (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataFieldApi.MetadataFieldsPut(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest(keyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataFieldApi.MetadataFieldsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MetadataFieldsPut`: KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldResponse
    fmt.Fprintf(os.Stdout, "Response from `MetadataFieldApi.MetadataFieldsPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetadataFieldsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest** | [**KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest**](KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest.md) | Properties of the metadata field type to be updated | 

### Return type

[**KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldResponse**](KeyfactorWebKeyfactorApiModelsMetadataFieldMetadataFieldResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


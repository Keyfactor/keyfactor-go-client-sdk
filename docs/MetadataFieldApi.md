# \MetadataFieldApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MetadataFieldCreateMetadataField**](MetadataFieldApi.md#MetadataFieldCreateMetadataField) | **Post** /MetadataFields | Creates a new metadata field type with the given metadata field type properties
[**MetadataFieldDeleteMetadataField**](MetadataFieldApi.md#MetadataFieldDeleteMetadataField) | **Delete** /MetadataFields/{id} | Deletes a persisted metadata field type by its unique id
[**MetadataFieldDeleteMetadataFields**](MetadataFieldApi.md#MetadataFieldDeleteMetadataFields) | **Delete** /MetadataFields | Deletes multiple persisted metadata field types by their unique ids
[**MetadataFieldGetAllMetadataFields**](MetadataFieldApi.md#MetadataFieldGetAllMetadataFields) | **Get** /MetadataFields | Returns all metadata field types according to the provided filter and output parameters
[**MetadataFieldGetMetadataField0**](MetadataFieldApi.md#MetadataFieldGetMetadataField0) | **Get** /MetadataFields/{id} | Gets a persisted metadata field type by its unique id
[**MetadataFieldGetMetadataField1**](MetadataFieldApi.md#MetadataFieldGetMetadataField1) | **Get** /MetadataFields/{name} | Gets a persisted metadata field type by its unique name
[**MetadataFieldGetMetadataFieldInUse**](MetadataFieldApi.md#MetadataFieldGetMetadataFieldInUse) | **Get** /MetadataFields/{id}/InUse | Determines if a metadata field type associated with the provided identifier is currently in use
[**MetadataFieldUpdateMetadataField**](MetadataFieldApi.md#MetadataFieldUpdateMetadataField) | **Put** /MetadataFields | Updates a persisted metadata field with the given metadata field type



## MetadataFieldCreateMetadataField

> KeyfactorApiModelsMetadataFieldMetadataFieldResponse MetadataFieldCreateMetadataField(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).MetadataFieldType(metadataFieldType).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    metadataFieldType := *openapiclient.NewKeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest("Name_example", "Description_example", int32(123)) // KeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest | Properties of the metadata field type to be created
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataFieldApi.MetadataFieldCreateMetadataField(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).MetadataFieldType(metadataFieldType).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataFieldApi.MetadataFieldCreateMetadataField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MetadataFieldCreateMetadataField`: KeyfactorApiModelsMetadataFieldMetadataFieldResponse
    fmt.Fprintf(os.Stdout, "Response from `MetadataFieldApi.MetadataFieldCreateMetadataField`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetadataFieldCreateMetadataFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **metadataFieldType** | [**KeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest**](KeyfactorApiModelsMetadataFieldMetadataFieldCreateRequest.md) | Properties of the metadata field type to be created | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsMetadataFieldMetadataFieldResponse**](KeyfactorApiModelsMetadataFieldMetadataFieldResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetadataFieldDeleteMetadataField

> MetadataFieldDeleteMetadataField(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Force(force).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    force := true // bool | Forces deletion of the metadata field type even if in-use (optional)
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataFieldApi.MetadataFieldDeleteMetadataField(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Force(force).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataFieldApi.MetadataFieldDeleteMetadataField``: %v\n", err)
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

Other parameters are passed through a pointer to a apiMetadataFieldDeleteMetadataFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **force** | **bool** | Forces deletion of the metadata field type even if in-use | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

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


## MetadataFieldDeleteMetadataFields

> MetadataFieldDeleteMetadataFields(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Ids(ids).Force(force).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    ids := []int32{int32(123)} // []int32 | Array of Keyfactor identifiers for metadata field types to be deleted
    force := true // bool | Forces deletion of the metadata field type even if in-use (optional)
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataFieldApi.MetadataFieldDeleteMetadataFields(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Ids(ids).Force(force).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataFieldApi.MetadataFieldDeleteMetadataFields``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetadataFieldDeleteMetadataFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **ids** | **[]int32** | Array of Keyfactor identifiers for metadata field types to be deleted | 
 **force** | **bool** | Forces deletion of the metadata field type even if in-use | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetadataFieldGetAllMetadataFields

> []ModelsMetadataFieldTypeModel MetadataFieldGetAllMetadataFields(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PqQueryString(pqQueryString).PqPageReturned(pqPageReturned).PqReturnLimit(pqReturnLimit).PqSortField(pqSortField).PqSortAscending(pqSortAscending).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")
    pqQueryString := "pqQueryString_example" // string | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) (optional)
    pqPageReturned := int32(56) // int32 | The current page within the result set to be returned (optional)
    pqReturnLimit := int32(56) // int32 | Maximum number of records to be returned in a single call (optional)
    pqSortField := "pqSortField_example" // string | Field by which the results should be sorted (view results via Management Portal for sortable columns) (optional)
    pqSortAscending := int32(56) // int32 | Field sort direction [0=ascending, 1=descending] (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataFieldApi.MetadataFieldGetAllMetadataFields(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PqQueryString(pqQueryString).PqPageReturned(pqPageReturned).PqReturnLimit(pqReturnLimit).PqSortField(pqSortField).PqSortAscending(pqSortAscending).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataFieldApi.MetadataFieldGetAllMetadataFields``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MetadataFieldGetAllMetadataFields`: []ModelsMetadataFieldTypeModel
    fmt.Fprintf(os.Stdout, "Response from `MetadataFieldApi.MetadataFieldGetAllMetadataFields`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetadataFieldGetAllMetadataFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]
 **pqQueryString** | **string** | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) | 
 **pqPageReturned** | **int32** | The current page within the result set to be returned | 
 **pqReturnLimit** | **int32** | Maximum number of records to be returned in a single call | 
 **pqSortField** | **string** | Field by which the results should be sorted (view results via Management Portal for sortable columns) | 
 **pqSortAscending** | **int32** | Field sort direction [0&#x3D;ascending, 1&#x3D;descending] | 

### Return type

[**[]ModelsMetadataFieldTypeModel**](ModelsMetadataFieldTypeModel.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetadataFieldGetMetadataField0

> ModelsMetadataFieldTypeModel MetadataFieldGetMetadataField0(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataFieldApi.MetadataFieldGetMetadataField0(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataFieldApi.MetadataFieldGetMetadataField0``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MetadataFieldGetMetadataField0`: ModelsMetadataFieldTypeModel
    fmt.Fprintf(os.Stdout, "Response from `MetadataFieldApi.MetadataFieldGetMetadataField0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The unique id of the metadata field type | 

### Other Parameters

Other parameters are passed through a pointer to a apiMetadataFieldGetMetadataField0Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsMetadataFieldTypeModel**](ModelsMetadataFieldTypeModel.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetadataFieldGetMetadataField1

> ModelsMetadataFieldTypeModel MetadataFieldGetMetadataField1(ctx, name).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataFieldApi.MetadataFieldGetMetadataField1(context.Background(), name).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataFieldApi.MetadataFieldGetMetadataField1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MetadataFieldGetMetadataField1`: ModelsMetadataFieldTypeModel
    fmt.Fprintf(os.Stdout, "Response from `MetadataFieldApi.MetadataFieldGetMetadataField1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The unique name of the metadata field. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMetadataFieldGetMetadataField1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsMetadataFieldTypeModel**](ModelsMetadataFieldTypeModel.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetadataFieldGetMetadataFieldInUse

> bool MetadataFieldGetMetadataFieldInUse(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataFieldApi.MetadataFieldGetMetadataFieldInUse(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataFieldApi.MetadataFieldGetMetadataFieldInUse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MetadataFieldGetMetadataFieldInUse`: bool
    fmt.Fprintf(os.Stdout, "Response from `MetadataFieldApi.MetadataFieldGetMetadataFieldInUse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Keyfactor identitifer of the metadata field | 

### Other Parameters

Other parameters are passed through a pointer to a apiMetadataFieldGetMetadataFieldInUseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

**bool**

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetadataFieldUpdateMetadataField

> KeyfactorApiModelsMetadataFieldMetadataFieldResponse MetadataFieldUpdateMetadataField(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).MetadataFieldType(metadataFieldType).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    metadataFieldType := *openapiclient.NewKeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest(int32(123), "Name_example", "Description_example", int32(123)) // KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest | Properties of the metadata field type to be updated
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataFieldApi.MetadataFieldUpdateMetadataField(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).MetadataFieldType(metadataFieldType).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataFieldApi.MetadataFieldUpdateMetadataField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MetadataFieldUpdateMetadataField`: KeyfactorApiModelsMetadataFieldMetadataFieldResponse
    fmt.Fprintf(os.Stdout, "Response from `MetadataFieldApi.MetadataFieldUpdateMetadataField`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetadataFieldUpdateMetadataFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **metadataFieldType** | [**KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest**](KeyfactorApiModelsMetadataFieldMetadataFieldUpdateRequest.md) | Properties of the metadata field type to be updated | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsMetadataFieldMetadataFieldResponse**](KeyfactorApiModelsMetadataFieldMetadataFieldResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


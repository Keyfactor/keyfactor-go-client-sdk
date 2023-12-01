# \WorkflowDefinitionApi

All URIs are relative to *https://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkflowDefinitionConfigureDefinitionSteps**](WorkflowDefinitionApi.md#WorkflowDefinitionConfigureDefinitionSteps) | **Put** /Workflow/Definitions/{definitionId}/Steps | Sets the provided steps on the latest version of the definition.
[**WorkflowDefinitionCreateNewDefinition**](WorkflowDefinitionApi.md#WorkflowDefinitionCreateNewDefinition) | **Post** /Workflow/Definitions | Creates a new base definition without any steps.
[**WorkflowDefinitionDelete**](WorkflowDefinitionApi.md#WorkflowDefinitionDelete) | **Delete** /Workflow/Definitions/{definitionId} | Deletes the definition matching the given Id.
[**WorkflowDefinitionGet**](WorkflowDefinitionApi.md#WorkflowDefinitionGet) | **Get** /Workflow/Definitions/{definitionId} | Gets a workflow definition.
[**WorkflowDefinitionGetStepSchema**](WorkflowDefinitionApi.md#WorkflowDefinitionGetStepSchema) | **Get** /Workflow/Definitions/Steps/{extensionName} | Gets the schema of a given step with the specified extension name.
[**WorkflowDefinitionPublishDefinition**](WorkflowDefinitionApi.md#WorkflowDefinitionPublishDefinition) | **Post** /Workflow/Definitions/{definitionId}/Publish | Makes the most recent version of a Workflow Definition the published version.
[**WorkflowDefinitionQuery**](WorkflowDefinitionApi.md#WorkflowDefinitionQuery) | **Get** /Workflow/Definitions | Gets the Definitions matching the query specifications.
[**WorkflowDefinitionQueryAvailableSteps**](WorkflowDefinitionApi.md#WorkflowDefinitionQueryAvailableSteps) | **Get** /Workflow/Definitions/Steps | Gets the result set of available steps for a given query.
[**WorkflowDefinitionQueryWorkflowTypes**](WorkflowDefinitionApi.md#WorkflowDefinitionQueryWorkflowTypes) | **Get** /Workflow/Definitions/Types | Performs a query against the workflow types in the system.
[**WorkflowDefinitionUpdateExistingDefinition**](WorkflowDefinitionApi.md#WorkflowDefinitionUpdateExistingDefinition) | **Put** /Workflow/Definitions/{definitionId} | Updates the existing definition&#39;s DisplayName and Description.



## WorkflowDefinitionConfigureDefinitionSteps

> KeyfactorApiModelsWorkflowsDefinitionResponse WorkflowDefinitionConfigureDefinitionSteps(ctx, definitionId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Request(request).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Sets the provided steps on the latest version of the definition.



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
    definitionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The Id of the definition.
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    request := []openapiclient.KeyfactorApiModelsWorkflowsDefinitionStepRequest{*openapiclient.NewKeyfactorApiModelsWorkflowsDefinitionStepRequest()} // []KeyfactorApiModelsWorkflowsDefinitionStepRequest | A collection of {KeyfactorApi.Models.Workflows.DefinitionStepRequest} defining the steps to set on the definition.
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowDefinitionApi.WorkflowDefinitionConfigureDefinitionSteps(context.Background(), definitionId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Request(request).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowDefinitionApi.WorkflowDefinitionConfigureDefinitionSteps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowDefinitionConfigureDefinitionSteps`: KeyfactorApiModelsWorkflowsDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowDefinitionApi.WorkflowDefinitionConfigureDefinitionSteps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** | The Id of the definition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowDefinitionConfigureDefinitionStepsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **request** | [**[]KeyfactorApiModelsWorkflowsDefinitionStepRequest**](KeyfactorApiModelsWorkflowsDefinitionStepRequest.md) | A collection of {KeyfactorApi.Models.Workflows.DefinitionStepRequest} defining the steps to set on the definition. | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsWorkflowsDefinitionResponse**](KeyfactorApiModelsWorkflowsDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowDefinitionCreateNewDefinition

> KeyfactorApiModelsWorkflowsDefinitionResponse WorkflowDefinitionCreateNewDefinition(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Request(request).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Creates a new base definition without any steps.

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
    request := *openapiclient.NewKeyfactorApiModelsWorkflowsDefinitionCreateRequest() // KeyfactorApiModelsWorkflowsDefinitionCreateRequest | A {KeyfactorApi.Models.Workflows.DefinitionCreateRequest} with the display name, description, key and type of the definition.
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowDefinitionApi.WorkflowDefinitionCreateNewDefinition(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Request(request).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowDefinitionApi.WorkflowDefinitionCreateNewDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowDefinitionCreateNewDefinition`: KeyfactorApiModelsWorkflowsDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowDefinitionApi.WorkflowDefinitionCreateNewDefinition`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowDefinitionCreateNewDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **request** | [**KeyfactorApiModelsWorkflowsDefinitionCreateRequest**](KeyfactorApiModelsWorkflowsDefinitionCreateRequest.md) | A {KeyfactorApi.Models.Workflows.DefinitionCreateRequest} with the display name, description, key and type of the definition. | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsWorkflowsDefinitionResponse**](KeyfactorApiModelsWorkflowsDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowDefinitionDelete

> WorkflowDefinitionDelete(ctx, definitionId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Deletes the definition matching the given Id.

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
    definitionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The Workflow Definition Id.
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowDefinitionApi.WorkflowDefinitionDelete(context.Background(), definitionId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowDefinitionApi.WorkflowDefinitionDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** | The Workflow Definition Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowDefinitionDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
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


## WorkflowDefinitionGet

> KeyfactorApiModelsWorkflowsDefinitionResponse WorkflowDefinitionGet(ctx, definitionId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).DefinitionVersion(definitionVersion).Exportable(exportable).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Gets a workflow definition.

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
    definitionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The Id of the definition to retrieve.
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    definitionVersion := int32(56) // int32 | The version to retrieve. If this value is not specified, the latest version will be returned. (optional)
    exportable := true // bool | Indicates if the response should be cleansed of role ids for export. (optional)
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowDefinitionApi.WorkflowDefinitionGet(context.Background(), definitionId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).DefinitionVersion(definitionVersion).Exportable(exportable).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowDefinitionApi.WorkflowDefinitionGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowDefinitionGet`: KeyfactorApiModelsWorkflowsDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowDefinitionApi.WorkflowDefinitionGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** | The Id of the definition to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowDefinitionGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **definitionVersion** | **int32** | The version to retrieve. If this value is not specified, the latest version will be returned. | 
 **exportable** | **bool** | Indicates if the response should be cleansed of role ids for export. | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsWorkflowsDefinitionResponse**](KeyfactorApiModelsWorkflowsDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowDefinitionGetStepSchema

> KeyfactorApiModelsWorkflowsAvailableStepResponse WorkflowDefinitionGetStepSchema(ctx, extensionName).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Gets the schema of a given step with the specified extension name.

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
    extensionName := "extensionName_example" // string | The extension name of a specific step in the step schema.
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowDefinitionApi.WorkflowDefinitionGetStepSchema(context.Background(), extensionName).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowDefinitionApi.WorkflowDefinitionGetStepSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowDefinitionGetStepSchema`: KeyfactorApiModelsWorkflowsAvailableStepResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowDefinitionApi.WorkflowDefinitionGetStepSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionName** | **string** | The extension name of a specific step in the step schema. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowDefinitionGetStepSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsWorkflowsAvailableStepResponse**](KeyfactorApiModelsWorkflowsAvailableStepResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowDefinitionPublishDefinition

> KeyfactorApiModelsWorkflowsDefinitionResponse WorkflowDefinitionPublishDefinition(ctx, definitionId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Makes the most recent version of a Workflow Definition the published version.

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
    definitionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The Workflow Definition Id.
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowDefinitionApi.WorkflowDefinitionPublishDefinition(context.Background(), definitionId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowDefinitionApi.WorkflowDefinitionPublishDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowDefinitionPublishDefinition`: KeyfactorApiModelsWorkflowsDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowDefinitionApi.WorkflowDefinitionPublishDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** | The Workflow Definition Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowDefinitionPublishDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsWorkflowsDefinitionResponse**](KeyfactorApiModelsWorkflowsDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowDefinitionQuery

> []KeyfactorApiModelsWorkflowsDefinitionQueryResponse WorkflowDefinitionQuery(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).QueryQueryString(queryQueryString).QueryPageReturned(queryPageReturned).QueryReturnLimit(queryReturnLimit).QuerySortField(querySortField).QuerySortAscending(querySortAscending).Execute()

Gets the Definitions matching the query specifications.

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
    queryQueryString := "queryQueryString_example" // string | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) (optional)
    queryPageReturned := int32(56) // int32 | The current page within the result set to be returned (optional)
    queryReturnLimit := int32(56) // int32 | Maximum number of records to be returned in a single call (optional)
    querySortField := "querySortField_example" // string | Field by which the results should be sorted (view results via Management Portal for sortable columns) (optional)
    querySortAscending := int32(56) // int32 | Field sort direction [0=ascending, 1=descending] (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowDefinitionApi.WorkflowDefinitionQuery(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).QueryQueryString(queryQueryString).QueryPageReturned(queryPageReturned).QueryReturnLimit(queryReturnLimit).QuerySortField(querySortField).QuerySortAscending(querySortAscending).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowDefinitionApi.WorkflowDefinitionQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowDefinitionQuery`: []KeyfactorApiModelsWorkflowsDefinitionQueryResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowDefinitionApi.WorkflowDefinitionQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowDefinitionQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]
 **queryQueryString** | **string** | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) | 
 **queryPageReturned** | **int32** | The current page within the result set to be returned | 
 **queryReturnLimit** | **int32** | Maximum number of records to be returned in a single call | 
 **querySortField** | **string** | Field by which the results should be sorted (view results via Management Portal for sortable columns) | 
 **querySortAscending** | **int32** | Field sort direction [0&#x3D;ascending, 1&#x3D;descending] | 

### Return type

[**[]KeyfactorApiModelsWorkflowsDefinitionQueryResponse**](KeyfactorApiModelsWorkflowsDefinitionQueryResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowDefinitionQueryAvailableSteps

> []KeyfactorApiModelsWorkflowsAvailableStepQueryResponse WorkflowDefinitionQueryAvailableSteps(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).QueryQueryString(queryQueryString).QueryPageReturned(queryPageReturned).QueryReturnLimit(queryReturnLimit).QuerySortField(querySortField).QuerySortAscending(querySortAscending).Execute()

Gets the result set of available steps for a given query.

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
    queryQueryString := "queryQueryString_example" // string | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) (optional)
    queryPageReturned := int32(56) // int32 | The current page within the result set to be returned (optional)
    queryReturnLimit := int32(56) // int32 | Maximum number of records to be returned in a single call (optional)
    querySortField := "querySortField_example" // string | Field by which the results should be sorted (view results via Management Portal for sortable columns) (optional)
    querySortAscending := int32(56) // int32 | Field sort direction [0=ascending, 1=descending] (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowDefinitionApi.WorkflowDefinitionQueryAvailableSteps(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).QueryQueryString(queryQueryString).QueryPageReturned(queryPageReturned).QueryReturnLimit(queryReturnLimit).QuerySortField(querySortField).QuerySortAscending(querySortAscending).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowDefinitionApi.WorkflowDefinitionQueryAvailableSteps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowDefinitionQueryAvailableSteps`: []KeyfactorApiModelsWorkflowsAvailableStepQueryResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowDefinitionApi.WorkflowDefinitionQueryAvailableSteps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowDefinitionQueryAvailableStepsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]
 **queryQueryString** | **string** | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) | 
 **queryPageReturned** | **int32** | The current page within the result set to be returned | 
 **queryReturnLimit** | **int32** | Maximum number of records to be returned in a single call | 
 **querySortField** | **string** | Field by which the results should be sorted (view results via Management Portal for sortable columns) | 
 **querySortAscending** | **int32** | Field sort direction [0&#x3D;ascending, 1&#x3D;descending] | 

### Return type

[**[]KeyfactorApiModelsWorkflowsAvailableStepQueryResponse**](KeyfactorApiModelsWorkflowsAvailableStepQueryResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowDefinitionQueryWorkflowTypes

> []KeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse WorkflowDefinitionQueryWorkflowTypes(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).QueryQueryString(queryQueryString).QueryPageReturned(queryPageReturned).QueryReturnLimit(queryReturnLimit).QuerySortField(querySortField).QuerySortAscending(querySortAscending).Execute()

Performs a query against the workflow types in the system.

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
    queryQueryString := "queryQueryString_example" // string | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) (optional)
    queryPageReturned := int32(56) // int32 | The current page within the result set to be returned (optional)
    queryReturnLimit := int32(56) // int32 | Maximum number of records to be returned in a single call (optional)
    querySortField := "querySortField_example" // string | Field by which the results should be sorted (view results via Management Portal for sortable columns) (optional)
    querySortAscending := int32(56) // int32 | Field sort direction [0=ascending, 1=descending] (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowDefinitionApi.WorkflowDefinitionQueryWorkflowTypes(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).QueryQueryString(queryQueryString).QueryPageReturned(queryPageReturned).QueryReturnLimit(queryReturnLimit).QuerySortField(querySortField).QuerySortAscending(querySortAscending).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowDefinitionApi.WorkflowDefinitionQueryWorkflowTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowDefinitionQueryWorkflowTypes`: []KeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowDefinitionApi.WorkflowDefinitionQueryWorkflowTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowDefinitionQueryWorkflowTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]
 **queryQueryString** | **string** | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) | 
 **queryPageReturned** | **int32** | The current page within the result set to be returned | 
 **queryReturnLimit** | **int32** | Maximum number of records to be returned in a single call | 
 **querySortField** | **string** | Field by which the results should be sorted (view results via Management Portal for sortable columns) | 
 **querySortAscending** | **int32** | Field sort direction [0&#x3D;ascending, 1&#x3D;descending] | 

### Return type

[**[]KeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse**](KeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowDefinitionUpdateExistingDefinition

> KeyfactorApiModelsWorkflowsDefinitionResponse WorkflowDefinitionUpdateExistingDefinition(ctx, definitionId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Request(request).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Updates the existing definition's DisplayName and Description.

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
    definitionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The Id of the definition to update.
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    request := *openapiclient.NewKeyfactorApiModelsWorkflowsDefinitionUpdateRequest() // KeyfactorApiModelsWorkflowsDefinitionUpdateRequest | The {KeyfactorApi.Models.Workflows.DefinitionUpdateRequest} holding the updated DispalyName and Description.
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowDefinitionApi.WorkflowDefinitionUpdateExistingDefinition(context.Background(), definitionId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Request(request).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowDefinitionApi.WorkflowDefinitionUpdateExistingDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowDefinitionUpdateExistingDefinition`: KeyfactorApiModelsWorkflowsDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowDefinitionApi.WorkflowDefinitionUpdateExistingDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** | The Id of the definition to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowDefinitionUpdateExistingDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **request** | [**KeyfactorApiModelsWorkflowsDefinitionUpdateRequest**](KeyfactorApiModelsWorkflowsDefinitionUpdateRequest.md) | The {KeyfactorApi.Models.Workflows.DefinitionUpdateRequest} holding the updated DispalyName and Description. | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsWorkflowsDefinitionResponse**](KeyfactorApiModelsWorkflowsDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


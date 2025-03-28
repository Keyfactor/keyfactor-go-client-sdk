# \WorkflowDefinitionApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorkflowDefinitions**](WorkflowDefinitionApi.md#CreateWorkflowDefinitions) | **POST** /Workflow/Definitions | Creates a new base definition without any steps.
[**CreateWorkflowDefinitionsDefinitionIdPublish**](WorkflowDefinitionApi.md#CreateWorkflowDefinitionsDefinitionIdPublish) | **POST** /Workflow/Definitions/{definitionId}/Publish | Makes the most recent version of a Workflow Definition the published version.
[**CreateWorkflowDefinitionsDefinitionIdPublishVersion**](WorkflowDefinitionApi.md#CreateWorkflowDefinitionsDefinitionIdPublishVersion) | **POST** /Workflow/Definitions/{definitionId}/Publish/{version} | Makes the specified version of a Workflow Definition the published version.
[**DeleteWorkflowDefinitionsDefinitionId**](WorkflowDefinitionApi.md#DeleteWorkflowDefinitionsDefinitionId) | **DELETE** /Workflow/Definitions/{definitionId} | Deletes the definition matching the given Id.
[**GetWorkflowDefinitions**](WorkflowDefinitionApi.md#GetWorkflowDefinitions) | **GET** /Workflow/Definitions | Gets the Definitions matching the query specifications.
[**GetWorkflowDefinitionsDefinitionId**](WorkflowDefinitionApi.md#GetWorkflowDefinitionsDefinitionId) | **GET** /Workflow/Definitions/{definitionId} | Gets a workflow definition.
[**GetWorkflowDefinitionsSteps**](WorkflowDefinitionApi.md#GetWorkflowDefinitionsSteps) | **GET** /Workflow/Definitions/Steps | Gets the result set of available steps for a given query.
[**GetWorkflowDefinitionsStepsExtensionName**](WorkflowDefinitionApi.md#GetWorkflowDefinitionsStepsExtensionName) | **GET** /Workflow/Definitions/Steps/{extensionName} | Gets the schema of a given step with the specified extension name.
[**GetWorkflowDefinitionsTypes**](WorkflowDefinitionApi.md#GetWorkflowDefinitionsTypes) | **GET** /Workflow/Definitions/Types | Performs a query against the workflow types in the system.
[**UpdateWorkflowDefinitionsDefinitionId**](WorkflowDefinitionApi.md#UpdateWorkflowDefinitionsDefinitionId) | **PUT** /Workflow/Definitions/{definitionId} | Updates the existing definition&#39;s DisplayName and Description.
[**UpdateWorkflowDefinitionsDefinitionIdStatus**](WorkflowDefinitionApi.md#UpdateWorkflowDefinitionsDefinitionIdStatus) | **PUT** /Workflow/Definitions/{definitionId}/Status | Updates the definition status matching the given Id.
[**UpdateWorkflowDefinitionsDefinitionIdSteps**](WorkflowDefinitionApi.md#UpdateWorkflowDefinitionsDefinitionIdSteps) | **PUT** /Workflow/Definitions/{definitionId}/Steps | Sets the provided steps on the latest version of the definition.



## CreateWorkflowDefinitions

> WorkflowsDefinitionResponse NewCreateWorkflowDefinitionsRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Force(force).XKeyfactorApiVersion(xKeyfactorApiVersion).WorkflowsDefinitionCreateRequest(workflowsDefinitionCreateRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    force := true // bool | Whether to force the creation of this definition in the case that it would run workflows immediately. (optional) (default to false)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    workflowsDefinitionCreateRequest := *openapiclient.NewWorkflowsDefinitionCreateRequest() // WorkflowsDefinitionCreateRequest | A Workflows.DefinitionCreateRequest with the display name, description, key and type of the definition. (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowDefinitionApi.NewCreateWorkflowDefinitionsRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Force(force).XKeyfactorApiVersion(xKeyfactorApiVersion).WorkflowsDefinitionCreateRequest(workflowsDefinitionCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowDefinitionApi.CreateWorkflowDefinitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkflowDefinitions`: WorkflowsDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowDefinitionApi.CreateWorkflowDefinitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **force** | **bool** | Whether to force the creation of this definition in the case that it would run workflows immediately. | [default to false]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **workflowsDefinitionCreateRequest** | [**WorkflowsDefinitionCreateRequest**](WorkflowsDefinitionCreateRequest.md) | A Workflows.DefinitionCreateRequest with the display name, description, key and type of the definition. | 

### Return type

[**WorkflowsDefinitionResponse**](WorkflowsDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWorkflowDefinitionsDefinitionIdPublish

> WorkflowsDefinitionResponse NewCreateWorkflowDefinitionsDefinitionIdPublishRequest(ctx, definitionId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowDefinitionApi.NewCreateWorkflowDefinitionsDefinitionIdPublishRequest(context.Background(), definitionId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowDefinitionApi.CreateWorkflowDefinitionsDefinitionIdPublish``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkflowDefinitionsDefinitionIdPublish`: WorkflowsDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowDefinitionApi.CreateWorkflowDefinitionsDefinitionIdPublish`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** | The Workflow Definition Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowDefinitionsDefinitionIdPublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**WorkflowsDefinitionResponse**](WorkflowsDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWorkflowDefinitionsDefinitionIdPublishVersion

> WorkflowsDefinitionResponse NewCreateWorkflowDefinitionsDefinitionIdPublishVersionRequest(ctx, definitionId, version).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Makes the specified version of a Workflow Definition the published version.

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
    version := int32(56) // int32 | The Workflow Version Id.
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowDefinitionApi.NewCreateWorkflowDefinitionsDefinitionIdPublishVersionRequest(context.Background(), definitionId, version).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowDefinitionApi.CreateWorkflowDefinitionsDefinitionIdPublishVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkflowDefinitionsDefinitionIdPublishVersion`: WorkflowsDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowDefinitionApi.CreateWorkflowDefinitionsDefinitionIdPublishVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** | The Workflow Definition Id. | 
**version** | **int32** | The Workflow Version Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowDefinitionsDefinitionIdPublishVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**WorkflowsDefinitionResponse**](WorkflowsDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWorkflowDefinitionsDefinitionId

> NewDeleteWorkflowDefinitionsDefinitionIdRequest(ctx, definitionId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowDefinitionApi.NewDeleteWorkflowDefinitionsDefinitionIdRequest(context.Background(), definitionId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowDefinitionApi.DeleteWorkflowDefinitionsDefinitionId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteWorkflowDefinitionsDefinitionIdRequest struct via the builder pattern


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


## GetWorkflowDefinitions

> []WorkflowsDefinitionQueryResponse NewGetWorkflowDefinitionsRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    queryString := "queryString_example" // string |  (optional)
    pageReturned := int32(56) // int32 |  (optional)
    returnLimit := int32(56) // int32 |  (optional)
    sortField := "sortField_example" // string |  (optional)
    sortAscending := openapiclient.Keyfactor.Common.QueryableExtensions.SortOrder(0) // KeyfactorCommonQueryableExtensionsSortOrder |  (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowDefinitionApi.NewGetWorkflowDefinitionsRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowDefinitionApi.GetWorkflowDefinitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflowDefinitions`: []WorkflowsDefinitionQueryResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowDefinitionApi.GetWorkflowDefinitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowDefinitionsRequest struct via the builder pattern


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

[**[]WorkflowsDefinitionQueryResponse**](WorkflowsDefinitionQueryResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowDefinitionsDefinitionId

> WorkflowsDefinitionResponse NewGetWorkflowDefinitionsDefinitionIdRequest(ctx, definitionId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).DefinitionVersion(definitionVersion).Exportable(exportable).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    definitionVersion := int32(56) // int32 | The version to retrieve. If this value is not specified, the latest version will be returned. (optional)
    exportable := true // bool | Indicates if the response should be cleansed of role ids for export. (optional) (default to false)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowDefinitionApi.NewGetWorkflowDefinitionsDefinitionIdRequest(context.Background(), definitionId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).DefinitionVersion(definitionVersion).Exportable(exportable).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowDefinitionApi.GetWorkflowDefinitionsDefinitionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflowDefinitionsDefinitionId`: WorkflowsDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowDefinitionApi.GetWorkflowDefinitionsDefinitionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** | The Id of the definition to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowDefinitionsDefinitionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **definitionVersion** | **int32** | The version to retrieve. If this value is not specified, the latest version will be returned. | 
 **exportable** | **bool** | Indicates if the response should be cleansed of role ids for export. | [default to false]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**WorkflowsDefinitionResponse**](WorkflowsDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowDefinitionsSteps

> []WorkflowsAvailableStepQueryResponse NewGetWorkflowDefinitionsStepsRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    queryString := "queryString_example" // string |  (optional)
    pageReturned := int32(56) // int32 |  (optional)
    returnLimit := int32(56) // int32 |  (optional)
    sortField := "sortField_example" // string |  (optional)
    sortAscending := openapiclient.Keyfactor.Common.QueryableExtensions.SortOrder(0) // KeyfactorCommonQueryableExtensionsSortOrder |  (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowDefinitionApi.NewGetWorkflowDefinitionsStepsRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowDefinitionApi.GetWorkflowDefinitionsSteps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflowDefinitionsSteps`: []WorkflowsAvailableStepQueryResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowDefinitionApi.GetWorkflowDefinitionsSteps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowDefinitionsStepsRequest struct via the builder pattern


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

[**[]WorkflowsAvailableStepQueryResponse**](WorkflowsAvailableStepQueryResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowDefinitionsStepsExtensionName

> WorkflowsAvailableStepResponse NewGetWorkflowDefinitionsStepsExtensionNameRequest(ctx, extensionName).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowDefinitionApi.NewGetWorkflowDefinitionsStepsExtensionNameRequest(context.Background(), extensionName).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowDefinitionApi.GetWorkflowDefinitionsStepsExtensionName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflowDefinitionsStepsExtensionName`: WorkflowsAvailableStepResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowDefinitionApi.GetWorkflowDefinitionsStepsExtensionName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionName** | **string** | The extension name of a specific step in the step schema. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowDefinitionsStepsExtensionNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**WorkflowsAvailableStepResponse**](WorkflowsAvailableStepResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowDefinitionsTypes

> []WorkflowsWorkflowTypeQueryResponse NewGetWorkflowDefinitionsTypesRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    queryString := "queryString_example" // string |  (optional)
    pageReturned := int32(56) // int32 |  (optional)
    returnLimit := int32(56) // int32 |  (optional)
    sortField := "sortField_example" // string |  (optional)
    sortAscending := openapiclient.Keyfactor.Common.QueryableExtensions.SortOrder(0) // KeyfactorCommonQueryableExtensionsSortOrder |  (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowDefinitionApi.NewGetWorkflowDefinitionsTypesRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowDefinitionApi.GetWorkflowDefinitionsTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflowDefinitionsTypes`: []WorkflowsWorkflowTypeQueryResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowDefinitionApi.GetWorkflowDefinitionsTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowDefinitionsTypesRequest struct via the builder pattern


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

[**[]WorkflowsWorkflowTypeQueryResponse**](WorkflowsWorkflowTypeQueryResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkflowDefinitionsDefinitionId

> WorkflowsDefinitionResponse NewUpdateWorkflowDefinitionsDefinitionIdRequest(ctx, definitionId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Force(force).XKeyfactorApiVersion(xKeyfactorApiVersion).WorkflowsDefinitionUpdateRequest(workflowsDefinitionUpdateRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    force := true // bool | Whether to force the update of this definition in the case that it would run workflows immediately. (optional) (default to false)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    workflowsDefinitionUpdateRequest := *openapiclient.NewWorkflowsDefinitionUpdateRequest() // WorkflowsDefinitionUpdateRequest | The Workflows.DefinitionUpdateRequest holding the updated DisplayName and Description. (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowDefinitionApi.NewUpdateWorkflowDefinitionsDefinitionIdRequest(context.Background(), definitionId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Force(force).XKeyfactorApiVersion(xKeyfactorApiVersion).WorkflowsDefinitionUpdateRequest(workflowsDefinitionUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowDefinitionApi.UpdateWorkflowDefinitionsDefinitionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorkflowDefinitionsDefinitionId`: WorkflowsDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowDefinitionApi.UpdateWorkflowDefinitionsDefinitionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** | The Id of the definition to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkflowDefinitionsDefinitionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **force** | **bool** | Whether to force the update of this definition in the case that it would run workflows immediately. | [default to false]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **workflowsDefinitionUpdateRequest** | [**WorkflowsDefinitionUpdateRequest**](WorkflowsDefinitionUpdateRequest.md) | The Workflows.DefinitionUpdateRequest holding the updated DisplayName and Description. | 

### Return type

[**WorkflowsDefinitionResponse**](WorkflowsDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkflowDefinitionsDefinitionIdStatus

> NewUpdateWorkflowDefinitionsDefinitionIdStatusRequest(ctx, definitionId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Force(force).XKeyfactorApiVersion(xKeyfactorApiVersion).Body(body).Execute()

Updates the definition status matching the given Id.

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    force := true // bool | Whether to force the status to update. (optional) (default to false)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    body := true // bool | The status to be updated. (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowDefinitionApi.NewUpdateWorkflowDefinitionsDefinitionIdStatusRequest(context.Background(), definitionId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Force(force).XKeyfactorApiVersion(xKeyfactorApiVersion).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowDefinitionApi.UpdateWorkflowDefinitionsDefinitionIdStatus``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateWorkflowDefinitionsDefinitionIdStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **force** | **bool** | Whether to force the status to update. | [default to false]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **body** | **bool** | The status to be updated. | 

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


## UpdateWorkflowDefinitionsDefinitionIdSteps

> WorkflowsDefinitionResponse NewUpdateWorkflowDefinitionsDefinitionIdStepsRequest(ctx, definitionId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).WorkflowsDefinitionStepRequest(workflowsDefinitionStepRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    workflowsDefinitionStepRequest := []openapiclient.WorkflowsDefinitionStepRequest{*openapiclient.NewWorkflowsDefinitionStepRequest()} // []WorkflowsDefinitionStepRequest | A collection of Workflows.DefinitionStepRequest defining the steps to set on the definition. (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowDefinitionApi.NewUpdateWorkflowDefinitionsDefinitionIdStepsRequest(context.Background(), definitionId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).WorkflowsDefinitionStepRequest(workflowsDefinitionStepRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowDefinitionApi.UpdateWorkflowDefinitionsDefinitionIdSteps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorkflowDefinitionsDefinitionIdSteps`: WorkflowsDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowDefinitionApi.UpdateWorkflowDefinitionsDefinitionIdSteps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** | The Id of the definition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkflowDefinitionsDefinitionIdStepsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **workflowsDefinitionStepRequest** | [**[]WorkflowsDefinitionStepRequest**](WorkflowsDefinitionStepRequest.md) | A collection of Workflows.DefinitionStepRequest defining the steps to set on the definition. | 

### Return type

[**WorkflowsDefinitionResponse**](WorkflowsDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


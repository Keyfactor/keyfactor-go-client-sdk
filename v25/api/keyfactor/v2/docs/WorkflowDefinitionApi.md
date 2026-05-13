# \WorkflowDefinitionApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorkflowDefinitions**](WorkflowDefinitionApi.md#CreateWorkflowDefinitions) | **POST** /Workflow/Definitions | Creates a new base definition without any steps.
[**GetWorkflowDefinitions**](WorkflowDefinitionApi.md#GetWorkflowDefinitions) | **GET** /Workflow/Definitions | Gets the Definitions matching the query specifications.
[**GetWorkflowDefinitionsDefinitionId**](WorkflowDefinitionApi.md#GetWorkflowDefinitionsDefinitionId) | **GET** /Workflow/Definitions/{definitionId} | Gets a workflow definition.
[**UpdateWorkflowDefinitionsDefinitionId**](WorkflowDefinitionApi.md#UpdateWorkflowDefinitionsDefinitionId) | **PUT** /Workflow/Definitions/{definitionId} | Updates the existing definition&#39;s DisplayName and Description.



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
    xKeyfactorApiVersion := "2" // string | Desired version of the api, if not provided defaults to v1 (optional)
    workflowsDefinitionCreateRequest := *openapiclient.NewWorkflowsDefinitionCreateRequest("DisplayName_example", "Description_example", "WorkflowType_example", []string{"Keys_example"}) // WorkflowsDefinitionCreateRequest | A Workflows.DefinitionCreateRequest with the display name, description, key and type of the definition. (optional)

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
    queryString := "queryString_example" // string | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) (optional)
    pageReturned := int32(56) // int32 | The current page within the result set to be returned (optional)
    returnLimit := int32(56) // int32 | Maximum number of records to be returned in a single call (optional)
    sortField := "sortField_example" // string | Field by which the results should be sorted (view results via Management Portal for sortable columns) (optional)
    sortAscending := openapiclient.Keyfactor.Common.QueryableExtensions.SortOrder(0) // KeyfactorCommonQueryableExtensionsSortOrder | Field sort direction [0=ascending, 1=descending] (optional)
    xKeyfactorApiVersion := "2" // string | Desired version of the api, if not provided defaults to v1 (optional)

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
 **queryString** | **string** | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) | 
 **pageReturned** | **int32** | The current page within the result set to be returned | 
 **returnLimit** | **int32** | Maximum number of records to be returned in a single call | 
 **sortField** | **string** | Field by which the results should be sorted (view results via Management Portal for sortable columns) | 
 **sortAscending** | [**KeyfactorCommonQueryableExtensionsSortOrder**](KeyfactorCommonQueryableExtensionsSortOrder.md) | Field sort direction [0&#x3D;ascending, 1&#x3D;descending] | 
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
    xKeyfactorApiVersion := "2" // string | Desired version of the api, if not provided defaults to v1 (optional)

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
    xKeyfactorApiVersion := "2" // string | Desired version of the api, if not provided defaults to v1 (optional)
    workflowsDefinitionUpdateRequest := *openapiclient.NewWorkflowsDefinitionUpdateRequest("DisplayName_example", "Description_example") // WorkflowsDefinitionUpdateRequest | The Workflows.DefinitionUpdateRequest holding the updated DisplayName and Description. (optional)

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


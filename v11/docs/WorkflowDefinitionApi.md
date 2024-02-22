# \WorkflowDefinitionApi

All URIs are relative to */Keyfactor/API*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkflowDefinitionsDefinitionIdDelete**](WorkflowDefinitionApi.md#WorkflowDefinitionsDefinitionIdDelete) | **Delete** /Workflow/Definitions/{definitionId} | Deletes the definition matching the given Id.
[**WorkflowDefinitionsDefinitionIdGet**](WorkflowDefinitionApi.md#WorkflowDefinitionsDefinitionIdGet) | **Get** /Workflow/Definitions/{definitionId} | Gets a workflow definition.
[**WorkflowDefinitionsDefinitionIdPublishPost**](WorkflowDefinitionApi.md#WorkflowDefinitionsDefinitionIdPublishPost) | **Post** /Workflow/Definitions/{definitionId}/Publish | Makes the most recent version of a Workflow Definition the published version.
[**WorkflowDefinitionsDefinitionIdPut**](WorkflowDefinitionApi.md#WorkflowDefinitionsDefinitionIdPut) | **Put** /Workflow/Definitions/{definitionId} | Updates the existing definition&#39;s DisplayName and Description.
[**WorkflowDefinitionsDefinitionIdStepsPut**](WorkflowDefinitionApi.md#WorkflowDefinitionsDefinitionIdStepsPut) | **Put** /Workflow/Definitions/{definitionId}/Steps | Sets the provided steps on the latest version of the definition.
[**WorkflowDefinitionsGet**](WorkflowDefinitionApi.md#WorkflowDefinitionsGet) | **Get** /Workflow/Definitions | Gets the Definitions matching the query specifications.
[**WorkflowDefinitionsPost**](WorkflowDefinitionApi.md#WorkflowDefinitionsPost) | **Post** /Workflow/Definitions | Creates a new base definition without any steps.
[**WorkflowDefinitionsStepsExtensionNameGet**](WorkflowDefinitionApi.md#WorkflowDefinitionsStepsExtensionNameGet) | **Get** /Workflow/Definitions/Steps/{extensionName} | Gets the schema of a given step with the specified extension name.
[**WorkflowDefinitionsStepsGet**](WorkflowDefinitionApi.md#WorkflowDefinitionsStepsGet) | **Get** /Workflow/Definitions/Steps | Gets the result set of available steps for a given query.
[**WorkflowDefinitionsTypesGet**](WorkflowDefinitionApi.md#WorkflowDefinitionsTypesGet) | **Get** /Workflow/Definitions/Types | Performs a query against the workflow types in the system.



## WorkflowDefinitionsDefinitionIdDelete

> WorkflowDefinitionsDefinitionIdDelete(ctx, definitionId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowDefinitionApi.WorkflowDefinitionsDefinitionIdDelete(context.Background(), definitionId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowDefinitionApi.WorkflowDefinitionsDefinitionIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiWorkflowDefinitionsDefinitionIdDeleteRequest struct via the builder pattern


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


## WorkflowDefinitionsDefinitionIdGet

> KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse WorkflowDefinitionsDefinitionIdGet(ctx, definitionId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).DefinitionVersion(definitionVersion).Exportable(exportable).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    definitionVersion := int32(56) // int32 | The version to retrieve. If this value is not specified, the latest version will be returned. (optional)
    exportable := true // bool | Indicates if the response should be cleansed of role ids for export. (optional) (default to false)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowDefinitionApi.WorkflowDefinitionsDefinitionIdGet(context.Background(), definitionId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).DefinitionVersion(definitionVersion).Exportable(exportable).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowDefinitionApi.WorkflowDefinitionsDefinitionIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowDefinitionsDefinitionIdGet`: KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowDefinitionApi.WorkflowDefinitionsDefinitionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** | The Id of the definition to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowDefinitionsDefinitionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **definitionVersion** | **int32** | The version to retrieve. If this value is not specified, the latest version will be returned. | 
 **exportable** | **bool** | Indicates if the response should be cleansed of role ids for export. | [default to false]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse**](KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowDefinitionsDefinitionIdPublishPost

> KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse WorkflowDefinitionsDefinitionIdPublishPost(ctx, definitionId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowDefinitionApi.WorkflowDefinitionsDefinitionIdPublishPost(context.Background(), definitionId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowDefinitionApi.WorkflowDefinitionsDefinitionIdPublishPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowDefinitionsDefinitionIdPublishPost`: KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowDefinitionApi.WorkflowDefinitionsDefinitionIdPublishPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** | The Workflow Definition Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowDefinitionsDefinitionIdPublishPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse**](KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowDefinitionsDefinitionIdPut

> KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse WorkflowDefinitionsDefinitionIdPut(ctx, definitionId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionUpdateRequest(keyfactorWebKeyfactorApiModelsWorkflowsDefinitionUpdateRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsWorkflowsDefinitionUpdateRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsWorkflowsDefinitionUpdateRequest() // KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionUpdateRequest | The Keyfactor.Web.KeyfactorApi.Models.Workflows.DefinitionUpdateRequest holding the updated DispalyName and Description. (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowDefinitionApi.WorkflowDefinitionsDefinitionIdPut(context.Background(), definitionId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionUpdateRequest(keyfactorWebKeyfactorApiModelsWorkflowsDefinitionUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowDefinitionApi.WorkflowDefinitionsDefinitionIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowDefinitionsDefinitionIdPut`: KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowDefinitionApi.WorkflowDefinitionsDefinitionIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** | The Id of the definition to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowDefinitionsDefinitionIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsWorkflowsDefinitionUpdateRequest** | [**KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionUpdateRequest**](KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionUpdateRequest.md) | The Keyfactor.Web.KeyfactorApi.Models.Workflows.DefinitionUpdateRequest holding the updated DispalyName and Description. | 

### Return type

[**KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse**](KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowDefinitionsDefinitionIdStepsPut

> KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse WorkflowDefinitionsDefinitionIdStepsPut(ctx, definitionId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest(keyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest := []openapiclient.KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest{*openapiclient.NewKeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest()} // []KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest | A collection of Keyfactor.Web.KeyfactorApi.Models.Workflows.DefinitionStepRequest defining the steps to set on the definition. (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowDefinitionApi.WorkflowDefinitionsDefinitionIdStepsPut(context.Background(), definitionId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest(keyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowDefinitionApi.WorkflowDefinitionsDefinitionIdStepsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowDefinitionsDefinitionIdStepsPut`: KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowDefinitionApi.WorkflowDefinitionsDefinitionIdStepsPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** | The Id of the definition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowDefinitionsDefinitionIdStepsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest** | [**[]KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest**](KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionStepRequest.md) | A collection of Keyfactor.Web.KeyfactorApi.Models.Workflows.DefinitionStepRequest defining the steps to set on the definition. | 

### Return type

[**KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse**](KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowDefinitionsGet

> []KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionQueryResponse WorkflowDefinitionsGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    queryString := "queryString_example" // string |  (optional)
    pageReturned := int32(56) // int32 |  (optional)
    returnLimit := int32(56) // int32 |  (optional)
    sortField := "sortField_example" // string |  (optional)
    sortAscending := int32(56) // int32 |  (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowDefinitionApi.WorkflowDefinitionsGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowDefinitionApi.WorkflowDefinitionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowDefinitionsGet`: []KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionQueryResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowDefinitionApi.WorkflowDefinitionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowDefinitionsGetRequest struct via the builder pattern


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

[**[]KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionQueryResponse**](KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionQueryResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowDefinitionsPost

> KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse WorkflowDefinitionsPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionCreateRequest(keyfactorWebKeyfactorApiModelsWorkflowsDefinitionCreateRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsWorkflowsDefinitionCreateRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsWorkflowsDefinitionCreateRequest() // KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionCreateRequest | A Keyfactor.Web.KeyfactorApi.Models.Workflows.DefinitionCreateRequest with the display name, description, key and type of the definition. (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowDefinitionApi.WorkflowDefinitionsPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionCreateRequest(keyfactorWebKeyfactorApiModelsWorkflowsDefinitionCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowDefinitionApi.WorkflowDefinitionsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowDefinitionsPost`: KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowDefinitionApi.WorkflowDefinitionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowDefinitionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsWorkflowsDefinitionCreateRequest** | [**KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionCreateRequest**](KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionCreateRequest.md) | A Keyfactor.Web.KeyfactorApi.Models.Workflows.DefinitionCreateRequest with the display name, description, key and type of the definition. | 

### Return type

[**KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse**](KeyfactorWebKeyfactorApiModelsWorkflowsDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowDefinitionsStepsExtensionNameGet

> KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse WorkflowDefinitionsStepsExtensionNameGet(ctx, extensionName).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowDefinitionApi.WorkflowDefinitionsStepsExtensionNameGet(context.Background(), extensionName).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowDefinitionApi.WorkflowDefinitionsStepsExtensionNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowDefinitionsStepsExtensionNameGet`: KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowDefinitionApi.WorkflowDefinitionsStepsExtensionNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionName** | **string** | The extension name of a specific step in the step schema. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowDefinitionsStepsExtensionNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse**](KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowDefinitionsStepsGet

> []KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse WorkflowDefinitionsStepsGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    queryString := "queryString_example" // string |  (optional)
    pageReturned := int32(56) // int32 |  (optional)
    returnLimit := int32(56) // int32 |  (optional)
    sortField := "sortField_example" // string |  (optional)
    sortAscending := int32(56) // int32 |  (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowDefinitionApi.WorkflowDefinitionsStepsGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowDefinitionApi.WorkflowDefinitionsStepsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowDefinitionsStepsGet`: []KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowDefinitionApi.WorkflowDefinitionsStepsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowDefinitionsStepsGetRequest struct via the builder pattern


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

[**[]KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse**](KeyfactorWebKeyfactorApiModelsWorkflowsAvailableStepQueryResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowDefinitionsTypesGet

> []KeyfactorWebKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse WorkflowDefinitionsTypesGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    queryString := "queryString_example" // string |  (optional)
    pageReturned := int32(56) // int32 |  (optional)
    returnLimit := int32(56) // int32 |  (optional)
    sortField := "sortField_example" // string |  (optional)
    sortAscending := int32(56) // int32 |  (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowDefinitionApi.WorkflowDefinitionsTypesGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowDefinitionApi.WorkflowDefinitionsTypesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowDefinitionsTypesGet`: []KeyfactorWebKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowDefinitionApi.WorkflowDefinitionsTypesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowDefinitionsTypesGetRequest struct via the builder pattern


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

[**[]KeyfactorWebKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse**](KeyfactorWebKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


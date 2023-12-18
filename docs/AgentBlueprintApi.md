# \AgentBlueprintApi

All URIs are relative to */Keyfactor/API*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentBluePrintApplyBlueprintPost**](AgentBlueprintApi.md#AgentBluePrintApplyBlueprintPost) | **Post** /AgentBluePrint/ApplyBlueprint | Applies the selected agent blueprint to the provided agents
[**AgentBluePrintGenerateBluePrintPost**](AgentBlueprintApi.md#AgentBluePrintGenerateBluePrintPost) | **Post** /AgentBluePrint/GenerateBluePrint | Generates an agent blueprint from the provided agents
[**AgentBluePrintGet**](AgentBlueprintApi.md#AgentBluePrintGet) | **Get** /AgentBluePrint | Returns all agent blueprints according to the provided filter and output parameters
[**AgentBluePrintIdDelete**](AgentBlueprintApi.md#AgentBluePrintIdDelete) | **Delete** /AgentBluePrint/{id} | Deletes an agent blueprint by its Keyfactor identifier
[**AgentBluePrintIdGet**](AgentBlueprintApi.md#AgentBluePrintIdGet) | **Get** /AgentBluePrint/{id} | Returns an agent blueprint according to the provided filter and output parameters
[**AgentBluePrintIdJobsGet**](AgentBlueprintApi.md#AgentBluePrintIdJobsGet) | **Get** /AgentBluePrint/{id}/Jobs | Gets the agent blueprint scheduled jobs
[**AgentBluePrintIdStoresGet**](AgentBlueprintApi.md#AgentBluePrintIdStoresGet) | **Get** /AgentBluePrint/{id}/Stores | Gets the agent blueprint certificate stores



## AgentBluePrintApplyBlueprintPost

> AgentBluePrintApplyBlueprintPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).TemplateId(templateId).XKeyfactorApiVersion(xKeyfactorApiVersion).RequestBody(requestBody).Execute()

Applies the selected agent blueprint to the provided agents

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
    templateId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Agent blueprint to apply to the agents (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    requestBody := []string{"Property_example"} // []string | Agents to apply the blueprints to (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentBlueprintApi.AgentBluePrintApplyBlueprintPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).TemplateId(templateId).XKeyfactorApiVersion(xKeyfactorApiVersion).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentBlueprintApi.AgentBluePrintApplyBlueprintPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAgentBluePrintApplyBlueprintPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **templateId** | **string** | Agent blueprint to apply to the agents | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **requestBody** | **[]string** | Agents to apply the blueprints to | 

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


## AgentBluePrintGenerateBluePrintPost

> KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintResponse AgentBluePrintGenerateBluePrintPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).AgentId(agentId).Name(name).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Generates an agent blueprint from the provided agents

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
    agentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Agent to generate a blueprint from (optional)
    name := "name_example" // string | Name of the new agent blueprint (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentBlueprintApi.AgentBluePrintGenerateBluePrintPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).AgentId(agentId).Name(name).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentBlueprintApi.AgentBluePrintGenerateBluePrintPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgentBluePrintGenerateBluePrintPost`: KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintResponse
    fmt.Fprintf(os.Stdout, "Response from `AgentBlueprintApi.AgentBluePrintGenerateBluePrintPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAgentBluePrintGenerateBluePrintPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **agentId** | **string** | Agent to generate a blueprint from | 
 **name** | **string** | Name of the new agent blueprint | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintResponse**](KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentBluePrintGet

> []KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintResponse AgentBluePrintGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns all agent blueprints according to the provided filter and output parameters

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
    pageReturned := int32(56) // int32 |  (optional)
    returnLimit := int32(56) // int32 |  (optional)
    sortField := "sortField_example" // string |  (optional)
    sortAscending := openapiclient.Keyfactor.Common.QueryableExtensionsSortOrder(0) // KeyfactorCommonQueryableExtensionsSortOrder |  (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentBlueprintApi.AgentBluePrintGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentBlueprintApi.AgentBluePrintGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgentBluePrintGet`: []KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintResponse
    fmt.Fprintf(os.Stdout, "Response from `AgentBlueprintApi.AgentBluePrintGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAgentBluePrintGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **pageReturned** | **int32** |  | 
 **returnLimit** | **int32** |  | 
 **sortField** | **string** |  | 
 **sortAscending** | [**KeyfactorCommonQueryableExtensionsSortOrder**](KeyfactorCommonQueryableExtensionsSortOrder.md) |  | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintResponse**](KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentBluePrintIdDelete

> AgentBluePrintIdDelete(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Deletes an agent blueprint by its Keyfactor identifier

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Keyfactor agent blueprint identifier (GUID)
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentBlueprintApi.AgentBluePrintIdDelete(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentBlueprintApi.AgentBluePrintIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Keyfactor agent blueprint identifier (GUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentBluePrintIdDeleteRequest struct via the builder pattern


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


## AgentBluePrintIdGet

> KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintResponse AgentBluePrintIdGet(ctx, id2).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Id(id).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns an agent blueprint according to the provided filter and output parameters

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
    id2 := "id_example" // string | 
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Returns a single agent blueprint associated with the provided id (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentBlueprintApi.AgentBluePrintIdGet(context.Background(), id2).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Id(id).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentBlueprintApi.AgentBluePrintIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgentBluePrintIdGet`: KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintResponse
    fmt.Fprintf(os.Stdout, "Response from `AgentBlueprintApi.AgentBluePrintIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id2** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentBluePrintIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **id** | **string** | Returns a single agent blueprint associated with the provided id | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintResponse**](KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentBluePrintIdJobsGet

> []KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse AgentBluePrintIdJobsGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Gets the agent blueprint scheduled jobs

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    pageReturned := int32(56) // int32 |  (optional)
    returnLimit := int32(56) // int32 |  (optional)
    sortField := "sortField_example" // string |  (optional)
    sortAscending := openapiclient.Keyfactor.Common.QueryableExtensionsSortOrder(0) // KeyfactorCommonQueryableExtensionsSortOrder |  (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentBlueprintApi.AgentBluePrintIdJobsGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentBlueprintApi.AgentBluePrintIdJobsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgentBluePrintIdJobsGet`: []KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse
    fmt.Fprintf(os.Stdout, "Response from `AgentBlueprintApi.AgentBluePrintIdJobsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentBluePrintIdJobsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **pageReturned** | **int32** |  | 
 **returnLimit** | **int32** |  | 
 **sortField** | **string** |  | 
 **sortAscending** | [**KeyfactorCommonQueryableExtensionsSortOrder**](KeyfactorCommonQueryableExtensionsSortOrder.md) |  | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse**](KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentBluePrintIdStoresGet

> []KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintStoresResponse AgentBluePrintIdStoresGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Gets the agent blueprint certificate stores

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    pageReturned := int32(56) // int32 |  (optional)
    returnLimit := int32(56) // int32 |  (optional)
    sortField := "sortField_example" // string |  (optional)
    sortAscending := openapiclient.Keyfactor.Common.QueryableExtensionsSortOrder(0) // KeyfactorCommonQueryableExtensionsSortOrder |  (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentBlueprintApi.AgentBluePrintIdStoresGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentBlueprintApi.AgentBluePrintIdStoresGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgentBluePrintIdStoresGet`: []KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintStoresResponse
    fmt.Fprintf(os.Stdout, "Response from `AgentBlueprintApi.AgentBluePrintIdStoresGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentBluePrintIdStoresGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **pageReturned** | **int32** |  | 
 **returnLimit** | **int32** |  | 
 **sortField** | **string** |  | 
 **sortAscending** | [**KeyfactorCommonQueryableExtensionsSortOrder**](KeyfactorCommonQueryableExtensionsSortOrder.md) |  | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintStoresResponse**](KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintStoresResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


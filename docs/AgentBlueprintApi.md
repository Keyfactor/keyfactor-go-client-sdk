# \AgentBlueprintApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentBlueprintApplyBlueprint**](AgentBlueprintApi.md#AgentBlueprintApplyBlueprint) | **Post** /AgentBluePrint/ApplyBlueprint | Applies the selected agent blueprint to the provided agents
[**AgentBlueprintDeleteBlueprint**](AgentBlueprintApi.md#AgentBlueprintDeleteBlueprint) | **Delete** /AgentBluePrint/{id} | Deletes an agent blueprint by its Keyfactor identifier
[**AgentBlueprintGenerateBlueprint**](AgentBlueprintApi.md#AgentBlueprintGenerateBlueprint) | **Post** /AgentBluePrint/GenerateBluePrint | Generates an agent blueprint from the provided agents
[**AgentBlueprintGetAgentBlueprint**](AgentBlueprintApi.md#AgentBlueprintGetAgentBlueprint) | **Get** /AgentBluePrint/{id} | Returns an agent blueprint according to the provided filter and output parameters
[**AgentBlueprintGetAgentBlueprints**](AgentBlueprintApi.md#AgentBlueprintGetAgentBlueprints) | **Get** /AgentBluePrint | Returns all agent blueprints according to the provided filter and output parameters
[**AgentBlueprintGetBlueprintJobs**](AgentBlueprintApi.md#AgentBlueprintGetBlueprintJobs) | **Get** /AgentBluePrint/{id}/Jobs | Gets the agent blueprint scheduled jobs
[**AgentBlueprintGetBlueprintStores**](AgentBlueprintApi.md#AgentBlueprintGetBlueprintStores) | **Get** /AgentBluePrint/{id}/Stores | Gets the agent blueprint certificate stores



## AgentBlueprintApplyBlueprint

> AgentBlueprintApplyBlueprint(ctx).TemplateId(templateId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).AgentIds(agentIds).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    templateId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Agent blueprint to apply to the agents
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    agentIds := []string{"00000000-0000-0000-0000-000000000000"} // []string | Agents to apply the blueprints to
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentBlueprintApi.AgentBlueprintApplyBlueprint(context.Background()).TemplateId(templateId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).AgentIds(agentIds).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentBlueprintApi.AgentBlueprintApplyBlueprint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAgentBlueprintApplyBlueprintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateId** | **string** | Agent blueprint to apply to the agents | 
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **agentIds** | **[]string** | Agents to apply the blueprints to | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentBlueprintDeleteBlueprint

> AgentBlueprintDeleteBlueprint(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentBlueprintApi.AgentBlueprintDeleteBlueprint(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentBlueprintApi.AgentBlueprintDeleteBlueprint``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAgentBlueprintDeleteBlueprintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentBlueprintGenerateBlueprint

> KeyfactorApiModelsOrchestratorsAgentBlueprintResponse AgentBlueprintGenerateBlueprint(ctx).AgentId(agentId).Name(name).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    agentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Agent to generate a blueprint from
    name := "name_example" // string | Name of the new agent blueprint
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentBlueprintApi.AgentBlueprintGenerateBlueprint(context.Background()).AgentId(agentId).Name(name).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentBlueprintApi.AgentBlueprintGenerateBlueprint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgentBlueprintGenerateBlueprint`: KeyfactorApiModelsOrchestratorsAgentBlueprintResponse
    fmt.Fprintf(os.Stdout, "Response from `AgentBlueprintApi.AgentBlueprintGenerateBlueprint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAgentBlueprintGenerateBlueprintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agentId** | **string** | Agent to generate a blueprint from | 
 **name** | **string** | Name of the new agent blueprint | 
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsOrchestratorsAgentBlueprintResponse**](KeyfactorApiModelsOrchestratorsAgentBlueprintResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentBlueprintGetAgentBlueprint

> KeyfactorApiModelsOrchestratorsAgentBlueprintResponse AgentBlueprintGetAgentBlueprint(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Returns a single agent blueprint associated with the provided id
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentBlueprintApi.AgentBlueprintGetAgentBlueprint(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentBlueprintApi.AgentBlueprintGetAgentBlueprint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgentBlueprintGetAgentBlueprint`: KeyfactorApiModelsOrchestratorsAgentBlueprintResponse
    fmt.Fprintf(os.Stdout, "Response from `AgentBlueprintApi.AgentBlueprintGetAgentBlueprint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Returns a single agent blueprint associated with the provided id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentBlueprintGetAgentBlueprintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsOrchestratorsAgentBlueprintResponse**](KeyfactorApiModelsOrchestratorsAgentBlueprintResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentBlueprintGetAgentBlueprints

> []KeyfactorApiModelsOrchestratorsAgentBlueprintResponse AgentBlueprintGetAgentBlueprints(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PqPageReturned(pqPageReturned).PqReturnLimit(pqReturnLimit).PqSortField(pqSortField).PqSortAscending(pqSortAscending).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")
    pqPageReturned := int32(56) // int32 | The current page within the result set to be returned (optional)
    pqReturnLimit := int32(56) // int32 | Maximum number of records to be returned in a single call (optional)
    pqSortField := "pqSortField_example" // string | Field by which the results should be sorted (OperationStart, OperationEnd, UserName) (optional)
    pqSortAscending := int32(56) // int32 | Field sort direction [0=ascending, 1=descending] (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentBlueprintApi.AgentBlueprintGetAgentBlueprints(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PqPageReturned(pqPageReturned).PqReturnLimit(pqReturnLimit).PqSortField(pqSortField).PqSortAscending(pqSortAscending).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentBlueprintApi.AgentBlueprintGetAgentBlueprints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgentBlueprintGetAgentBlueprints`: []KeyfactorApiModelsOrchestratorsAgentBlueprintResponse
    fmt.Fprintf(os.Stdout, "Response from `AgentBlueprintApi.AgentBlueprintGetAgentBlueprints`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAgentBlueprintGetAgentBlueprintsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]
 **pqPageReturned** | **int32** | The current page within the result set to be returned | 
 **pqReturnLimit** | **int32** | Maximum number of records to be returned in a single call | 
 **pqSortField** | **string** | Field by which the results should be sorted (OperationStart, OperationEnd, UserName) | 
 **pqSortAscending** | **int32** | Field sort direction [0&#x3D;ascending, 1&#x3D;descending] | 

### Return type

[**[]KeyfactorApiModelsOrchestratorsAgentBlueprintResponse**](KeyfactorApiModelsOrchestratorsAgentBlueprintResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentBlueprintGetBlueprintJobs

> []KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse AgentBlueprintGetBlueprintJobs(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PqPageReturned(pqPageReturned).PqReturnLimit(pqReturnLimit).PqSortField(pqSortField).PqSortAscending(pqSortAscending).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")
    pqPageReturned := int32(56) // int32 | The current page within the result set to be returned (optional)
    pqReturnLimit := int32(56) // int32 | Maximum number of records to be returned in a single call (optional)
    pqSortField := "pqSortField_example" // string | Field by which the results should be sorted (OperationStart, OperationEnd, UserName) (optional)
    pqSortAscending := int32(56) // int32 | Field sort direction [0=ascending, 1=descending] (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentBlueprintApi.AgentBlueprintGetBlueprintJobs(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PqPageReturned(pqPageReturned).PqReturnLimit(pqReturnLimit).PqSortField(pqSortField).PqSortAscending(pqSortAscending).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentBlueprintApi.AgentBlueprintGetBlueprintJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgentBlueprintGetBlueprintJobs`: []KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse
    fmt.Fprintf(os.Stdout, "Response from `AgentBlueprintApi.AgentBlueprintGetBlueprintJobs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentBlueprintGetBlueprintJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]
 **pqPageReturned** | **int32** | The current page within the result set to be returned | 
 **pqReturnLimit** | **int32** | Maximum number of records to be returned in a single call | 
 **pqSortField** | **string** | Field by which the results should be sorted (OperationStart, OperationEnd, UserName) | 
 **pqSortAscending** | **int32** | Field sort direction [0&#x3D;ascending, 1&#x3D;descending] | 

### Return type

[**[]KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse**](KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentBlueprintGetBlueprintStores

> []KeyfactorApiModelsOrchestratorsAgentBlueprintStoresResponse AgentBlueprintGetBlueprintStores(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PqPageReturned(pqPageReturned).PqReturnLimit(pqReturnLimit).PqSortField(pqSortField).PqSortAscending(pqSortAscending).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")
    pqPageReturned := int32(56) // int32 | The current page within the result set to be returned (optional)
    pqReturnLimit := int32(56) // int32 | Maximum number of records to be returned in a single call (optional)
    pqSortField := "pqSortField_example" // string | Field by which the results should be sorted (OperationStart, OperationEnd, UserName) (optional)
    pqSortAscending := int32(56) // int32 | Field sort direction [0=ascending, 1=descending] (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentBlueprintApi.AgentBlueprintGetBlueprintStores(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PqPageReturned(pqPageReturned).PqReturnLimit(pqReturnLimit).PqSortField(pqSortField).PqSortAscending(pqSortAscending).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentBlueprintApi.AgentBlueprintGetBlueprintStores``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgentBlueprintGetBlueprintStores`: []KeyfactorApiModelsOrchestratorsAgentBlueprintStoresResponse
    fmt.Fprintf(os.Stdout, "Response from `AgentBlueprintApi.AgentBlueprintGetBlueprintStores`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentBlueprintGetBlueprintStoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]
 **pqPageReturned** | **int32** | The current page within the result set to be returned | 
 **pqReturnLimit** | **int32** | Maximum number of records to be returned in a single call | 
 **pqSortField** | **string** | Field by which the results should be sorted (OperationStart, OperationEnd, UserName) | 
 **pqSortAscending** | **int32** | Field sort direction [0&#x3D;ascending, 1&#x3D;descending] | 

### Return type

[**[]KeyfactorApiModelsOrchestratorsAgentBlueprintStoresResponse**](KeyfactorApiModelsOrchestratorsAgentBlueprintStoresResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


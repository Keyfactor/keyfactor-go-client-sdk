# \AgentPoolApi

All URIs are relative to */Keyfactor/API*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentPoolsAgentsGet**](AgentPoolApi.md#AgentPoolsAgentsGet) | **Get** /AgentPools/Agents | Returns all agents for the default agent pool
[**AgentPoolsGet**](AgentPoolApi.md#AgentPoolsGet) | **Get** /AgentPools | Returns all agent pools according to the provided filter and output parameters
[**AgentPoolsIdDelete**](AgentPoolApi.md#AgentPoolsIdDelete) | **Delete** /AgentPools/{id} | Deletes the agent pool associated with the provided id
[**AgentPoolsIdGet**](AgentPoolApi.md#AgentPoolsIdGet) | **Get** /AgentPools/{id} | Returns a single agent pool associated with the provided id
[**AgentPoolsPost**](AgentPoolApi.md#AgentPoolsPost) | **Post** /AgentPools | Creates an agent pool with the provided properties
[**AgentPoolsPut**](AgentPoolApi.md#AgentPoolsPut) | **Put** /AgentPools | Updates an existing agent pool with the provided properties



## AgentPoolsAgentsGet

> []KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolAgentGetResponse AgentPoolsAgentsGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns all agents for the default agent pool

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
    queryString := "queryString_example" // string | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) (optional)
    pageReturned := int32(56) // int32 | The current page within the result set to be returned (optional)
    returnLimit := int32(56) // int32 | Maximum number of records to be returned in a single call (optional)
    sortField := "sortField_example" // string | Field by which the results should be sorted (view results via Management Portal for sortable columns) (optional)
    sortAscending := int32(56) // int32 | Field sort direction [0=ascending, 1=descending] (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentPoolApi.AgentPoolsAgentsGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentPoolApi.AgentPoolsAgentsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgentPoolsAgentsGet`: []KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolAgentGetResponse
    fmt.Fprintf(os.Stdout, "Response from `AgentPoolApi.AgentPoolsAgentsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAgentPoolsAgentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **queryString** | **string** | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) | 
 **pageReturned** | **int32** | The current page within the result set to be returned | 
 **returnLimit** | **int32** | Maximum number of records to be returned in a single call | 
 **sortField** | **string** | Field by which the results should be sorted (view results via Management Portal for sortable columns) | 
 **sortAscending** | **int32** | Field sort direction [0&#x3D;ascending, 1&#x3D;descending] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolAgentGetResponse**](KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolAgentGetResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentPoolsGet

> []KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse AgentPoolsGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns all agent pools according to the provided filter and output parameters

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
    queryString := "queryString_example" // string | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) (optional)
    pageReturned := int32(56) // int32 | The current page within the result set to be returned (optional)
    returnLimit := int32(56) // int32 | Maximum number of records to be returned in a single call (optional)
    sortField := "sortField_example" // string | Field by which the results should be sorted (view results via Management Portal for sortable columns) (optional)
    sortAscending := int32(56) // int32 | Field sort direction [0=ascending, 1=descending] (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentPoolApi.AgentPoolsGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentPoolApi.AgentPoolsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgentPoolsGet`: []KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse
    fmt.Fprintf(os.Stdout, "Response from `AgentPoolApi.AgentPoolsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAgentPoolsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **queryString** | **string** | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) | 
 **pageReturned** | **int32** | The current page within the result set to be returned | 
 **returnLimit** | **int32** | Maximum number of records to be returned in a single call | 
 **sortField** | **string** | Field by which the results should be sorted (view results via Management Portal for sortable columns) | 
 **sortAscending** | **int32** | Field sort direction [0&#x3D;ascending, 1&#x3D;descending] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse**](KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentPoolsIdDelete

> AgentPoolsIdDelete(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Deletes the agent pool associated with the provided id

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Keyfactor identifier (GUID) of the agent pool
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentPoolApi.AgentPoolsIdDelete(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentPoolApi.AgentPoolsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Keyfactor identifier (GUID) of the agent pool | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentPoolsIdDeleteRequest struct via the builder pattern


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


## AgentPoolsIdGet

> KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse AgentPoolsIdGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns a single agent pool associated with the provided id

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Keyfactor (GUID) identifier of the agent pool
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentPoolApi.AgentPoolsIdGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentPoolApi.AgentPoolsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgentPoolsIdGet`: KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse
    fmt.Fprintf(os.Stdout, "Response from `AgentPoolApi.AgentPoolsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Keyfactor (GUID) identifier of the agent pool | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentPoolsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse**](KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentPoolsPost

> KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse AgentPoolsPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolCreationRequest(keyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolCreationRequest).Execute()

Creates an agent pool with the provided properties

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
    keyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolCreationRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolCreationRequest("Name_example") // KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolCreationRequest | Agent pool properties to be applied to the new pool (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentPoolApi.AgentPoolsPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolCreationRequest(keyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolCreationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentPoolApi.AgentPoolsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgentPoolsPost`: KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse
    fmt.Fprintf(os.Stdout, "Response from `AgentPoolApi.AgentPoolsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAgentPoolsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolCreationRequest** | [**KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolCreationRequest**](KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolCreationRequest.md) | Agent pool properties to be applied to the new pool | 

### Return type

[**KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse**](KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentPoolsPut

> KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse AgentPoolsPut(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolUpdateRequest(keyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolUpdateRequest).Execute()

Updates an existing agent pool with the provided properties

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
    keyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolUpdateRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolUpdateRequest("AgentPoolId_example", "Name_example") // KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolUpdateRequest | Agent pool properties to be applied to the existing pool (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentPoolApi.AgentPoolsPut(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolUpdateRequest(keyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentPoolApi.AgentPoolsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgentPoolsPut`: KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse
    fmt.Fprintf(os.Stdout, "Response from `AgentPoolApi.AgentPoolsPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAgentPoolsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolUpdateRequest** | [**KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolUpdateRequest**](KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolUpdateRequest.md) | Agent pool properties to be applied to the existing pool | 

### Return type

[**KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse**](KeyfactorWebKeyfactorApiModelsOrchestratorPoolsAgentPoolGetResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \AgentPoolApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentPoolCreateAgentPool**](AgentPoolApi.md#AgentPoolCreateAgentPool) | **Post** /AgentPools | Creates an agent pool with the provided properties
[**AgentPoolDeleteAgentPool**](AgentPoolApi.md#AgentPoolDeleteAgentPool) | **Delete** /AgentPools/{id} | Deletes the agent pool associated with the provided id
[**AgentPoolGetAgentPoolById**](AgentPoolApi.md#AgentPoolGetAgentPoolById) | **Get** /AgentPools/{id} | Returns a single agent pool associated with the provided id
[**AgentPoolGetAgentPools**](AgentPoolApi.md#AgentPoolGetAgentPools) | **Get** /AgentPools | Returns all agent pools according to the provided filter and output parameters
[**AgentPoolGetDefaultAgentPoolAgents**](AgentPoolApi.md#AgentPoolGetDefaultAgentPoolAgents) | **Get** /AgentPools/Agents | Returns all agents for the default agent pool
[**AgentPoolUpdateAgentPool**](AgentPoolApi.md#AgentPoolUpdateAgentPool) | **Put** /AgentPools | Updates an existing agent pool with the provided properties



## AgentPoolCreateAgentPool

> ModelsAgentsAgentPool AgentPoolCreateAgentPool(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).AgentPool(agentPool).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    agentPool := *openapiclient.NewModelsAgentsAgentPool("Name_example") // ModelsAgentsAgentPool | Agent pool properties to be applied to the new pool
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentPoolApi.AgentPoolCreateAgentPool(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).AgentPool(agentPool).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentPoolApi.AgentPoolCreateAgentPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgentPoolCreateAgentPool`: ModelsAgentsAgentPool
    fmt.Fprintf(os.Stdout, "Response from `AgentPoolApi.AgentPoolCreateAgentPool`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAgentPoolCreateAgentPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **agentPool** | [**ModelsAgentsAgentPool**](ModelsAgentsAgentPool.md) | Agent pool properties to be applied to the new pool | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsAgentsAgentPool**](ModelsAgentsAgentPool.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentPoolDeleteAgentPool

> AgentPoolDeleteAgentPool(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentPoolApi.AgentPoolDeleteAgentPool(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentPoolApi.AgentPoolDeleteAgentPool``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAgentPoolDeleteAgentPoolRequest struct via the builder pattern


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


## AgentPoolGetAgentPoolById

> ModelsAgentsAgentPool AgentPoolGetAgentPoolById(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentPoolApi.AgentPoolGetAgentPoolById(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentPoolApi.AgentPoolGetAgentPoolById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgentPoolGetAgentPoolById`: ModelsAgentsAgentPool
    fmt.Fprintf(os.Stdout, "Response from `AgentPoolApi.AgentPoolGetAgentPoolById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Keyfactor (GUID) identifier of the agent pool | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentPoolGetAgentPoolByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsAgentsAgentPool**](ModelsAgentsAgentPool.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentPoolGetAgentPools

> []ModelsAgentsAgentPool AgentPoolGetAgentPools(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PqQueryString(pqQueryString).PqPageReturned(pqPageReturned).PqReturnLimit(pqReturnLimit).PqSortField(pqSortField).PqSortAscending(pqSortAscending).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")
    pqQueryString := "pqQueryString_example" // string | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) (optional)
    pqPageReturned := int32(56) // int32 | The current page within the result set to be returned (optional)
    pqReturnLimit := int32(56) // int32 | Maximum number of records to be returned in a single call (optional)
    pqSortField := "pqSortField_example" // string | Field by which the results should be sorted (view results via Management Portal for sortable columns) (optional)
    pqSortAscending := int32(56) // int32 | Field sort direction [0=ascending, 1=descending] (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentPoolApi.AgentPoolGetAgentPools(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PqQueryString(pqQueryString).PqPageReturned(pqPageReturned).PqReturnLimit(pqReturnLimit).PqSortField(pqSortField).PqSortAscending(pqSortAscending).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentPoolApi.AgentPoolGetAgentPools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgentPoolGetAgentPools`: []ModelsAgentsAgentPool
    fmt.Fprintf(os.Stdout, "Response from `AgentPoolApi.AgentPoolGetAgentPools`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAgentPoolGetAgentPoolsRequest struct via the builder pattern


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

[**[]ModelsAgentsAgentPool**](ModelsAgentsAgentPool.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentPoolGetDefaultAgentPoolAgents

> []ModelsAgentsAgentPoolAgent AgentPoolGetDefaultAgentPoolAgents(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PqQueryString(pqQueryString).PqPageReturned(pqPageReturned).PqReturnLimit(pqReturnLimit).PqSortField(pqSortField).PqSortAscending(pqSortAscending).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")
    pqQueryString := "pqQueryString_example" // string | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) (optional)
    pqPageReturned := int32(56) // int32 | The current page within the result set to be returned (optional)
    pqReturnLimit := int32(56) // int32 | Maximum number of records to be returned in a single call (optional)
    pqSortField := "pqSortField_example" // string | Field by which the results should be sorted (view results via Management Portal for sortable columns) (optional)
    pqSortAscending := int32(56) // int32 | Field sort direction [0=ascending, 1=descending] (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentPoolApi.AgentPoolGetDefaultAgentPoolAgents(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).PqQueryString(pqQueryString).PqPageReturned(pqPageReturned).PqReturnLimit(pqReturnLimit).PqSortField(pqSortField).PqSortAscending(pqSortAscending).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentPoolApi.AgentPoolGetDefaultAgentPoolAgents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgentPoolGetDefaultAgentPoolAgents`: []ModelsAgentsAgentPoolAgent
    fmt.Fprintf(os.Stdout, "Response from `AgentPoolApi.AgentPoolGetDefaultAgentPoolAgents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAgentPoolGetDefaultAgentPoolAgentsRequest struct via the builder pattern


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

[**[]ModelsAgentsAgentPoolAgent**](ModelsAgentsAgentPoolAgent.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentPoolUpdateAgentPool

> ModelsAgentsAgentPool AgentPoolUpdateAgentPool(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).AgentPool(agentPool).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    agentPool := *openapiclient.NewModelsAgentsAgentPool("Name_example") // ModelsAgentsAgentPool | Agent pool properties to be applied to the existing pool
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentPoolApi.AgentPoolUpdateAgentPool(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).AgentPool(agentPool).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentPoolApi.AgentPoolUpdateAgentPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgentPoolUpdateAgentPool`: ModelsAgentsAgentPool
    fmt.Fprintf(os.Stdout, "Response from `AgentPoolApi.AgentPoolUpdateAgentPool`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAgentPoolUpdateAgentPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **agentPool** | [**ModelsAgentsAgentPool**](ModelsAgentsAgentPool.md) | Agent pool properties to be applied to the existing pool | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsAgentsAgentPool**](ModelsAgentsAgentPool.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


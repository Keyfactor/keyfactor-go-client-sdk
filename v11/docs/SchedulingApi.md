# \SchedulingApi

All URIs are relative to */Keyfactor/API*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SchedulingPost**](SchedulingApi.md#SchedulingPost) | **Post** /Scheduling | 



## SchedulingPost

> KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse SchedulingPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskRequest(keyfactorWebKeyfactorApiModelsSchedulingScheduledTaskRequest).Execute()



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
    keyfactorWebKeyfactorApiModelsSchedulingScheduledTaskRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskRequest(openapiclient.CSS.CMS.Core.Enums.ScheduledTaskType(0)) // KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskRequest |  (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchedulingApi.SchedulingPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskRequest(keyfactorWebKeyfactorApiModelsSchedulingScheduledTaskRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchedulingApi.SchedulingPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SchedulingPost`: KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse
    fmt.Fprintf(os.Stdout, "Response from `SchedulingApi.SchedulingPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSchedulingPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsSchedulingScheduledTaskRequest** | [**KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskRequest**](KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskRequest.md) |  | 

### Return type

[**KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse**](KeyfactorWebKeyfactorApiModelsSchedulingScheduledTaskResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


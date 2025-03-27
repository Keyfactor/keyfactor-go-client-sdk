# \SchedulingApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateScheduling**](SchedulingApi.md#CreateScheduling) | **POST** /Scheduling | 



## CreateScheduling

> SchedulingScheduledTaskResponse CreateScheduling(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).SchedulingScheduledTaskRequest(schedulingScheduledTaskRequest).Execute()



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
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    schedulingScheduledTaskRequest := *openapiclient.NewSchedulingScheduledTaskRequest(openapiclient.CSS.CMS.Core.Enums.ScheduledTaskType(0)) // SchedulingScheduledTaskRequest |  (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchedulingApi.CreateScheduling(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).SchedulingScheduledTaskRequest(schedulingScheduledTaskRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchedulingApi.CreateScheduling``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateScheduling`: SchedulingScheduledTaskResponse
    fmt.Fprintf(os.Stdout, "Response from `SchedulingApi.CreateScheduling`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSchedulingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **schedulingScheduledTaskRequest** | [**SchedulingScheduledTaskRequest**](SchedulingScheduledTaskRequest.md) |  | 

### Return type

[**SchedulingScheduledTaskResponse**](SchedulingScheduledTaskResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


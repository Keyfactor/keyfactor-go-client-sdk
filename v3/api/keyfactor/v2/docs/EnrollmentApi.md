# \EnrollmentApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEnrollmentPFX**](EnrollmentApi.md#CreateEnrollmentPFX) | **POST** /Enrollment/PFX | Performs a PFX Enrollment based upon the provided request



## CreateEnrollmentPFX

> EnrollmentPFXEnrollmentManagementResponse BuildCreateEnrollmentPFXRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XCertificateformat(xCertificateformat).XKeyfactorApiVersion(xKeyfactorApiVersion).EnrollmentPFXEnrollmentWithStoresRequest(enrollmentPFXEnrollmentWithStoresRequest).Execute()

Performs a PFX Enrollment based upon the provided request



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
    xCertificateformat := "PFX" // string | Desired format [PFX, ZIP, PEM, JKS, STORE, REPLACE]
    xKeyfactorApiVersion := "2" // string | Desired version of the api, if not provided defaults to v1 (optional)
    enrollmentPFXEnrollmentWithStoresRequest := *openapiclient.NewEnrollmentPFXEnrollmentWithStoresRequest() // EnrollmentPFXEnrollmentWithStoresRequest | The information needed to perform the PFX Enrollment (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.BuildCreateEnrollmentPFXRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XCertificateformat(xCertificateformat).XKeyfactorApiVersion(xKeyfactorApiVersion).EnrollmentPFXEnrollmentWithStoresRequest(enrollmentPFXEnrollmentWithStoresRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.CreateEnrollmentPFX``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEnrollmentPFX`: EnrollmentPFXEnrollmentManagementResponse
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.CreateEnrollmentPFX`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnrollmentPFXRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xCertificateformat** | **string** | Desired format [PFX, ZIP, PEM, JKS, STORE, REPLACE] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **enrollmentPFXEnrollmentWithStoresRequest** | [**EnrollmentPFXEnrollmentWithStoresRequest**](EnrollmentPFXEnrollmentWithStoresRequest.md) | The information needed to perform the PFX Enrollment | 

### Return type

[**EnrollmentPFXEnrollmentManagementResponse**](EnrollmentPFXEnrollmentManagementResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


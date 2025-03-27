# \CertificateApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCertificatesIdentityAuditById**](CertificateApi.md#GetCertificatesIdentityAuditById) | **GET** /Certificates/IdentityAudit/{id} | Audit identity permissions for certificate



## GetCertificatesIdentityAuditById

> []CertificatesCertificateIdentityAuditResponse2 GetCertificatesIdentityAuditById(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Audit identity permissions for certificate

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
    id := int32(56) // int32 | The Id of the certificate being checked
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    collectionId := int32(56) // int32 | An optional parameter for the collectin Id the certificate is in.  Defaults to no collection (optional) (default to 0)
    xKeyfactorApiVersion := "2" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.GetCertificatesIdentityAuditById(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.GetCertificatesIdentityAuditById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificatesIdentityAuditById`: []CertificatesCertificateIdentityAuditResponse2
    fmt.Fprintf(os.Stdout, "Response from `CertificateApi.GetCertificatesIdentityAuditById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Id of the certificate being checked | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificatesIdentityAuditByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **collectionId** | **int32** | An optional parameter for the collectin Id the certificate is in.  Defaults to no collection | [default to 0]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]CertificatesCertificateIdentityAuditResponse2**](CertificatesCertificateIdentityAuditResponse2.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


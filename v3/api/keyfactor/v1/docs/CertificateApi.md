# \CertificateApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCertificatesAnalyze**](CertificateApi.md#CreateCertificatesAnalyze) | **POST** /Certificates/Analyze | Returns the public information of the certificate
[**CreateCertificatesDownload**](CertificateApi.md#CreateCertificatesDownload) | **POST** /Certificates/Download | Downloads the persisted certificate associated with the provided query
[**CreateCertificatesImport**](CertificateApi.md#CreateCertificatesImport) | **POST** /Certificates/Import | Imports the provided certificate into the Keyfactor instance, including any provided associated data
[**CreateCertificatesRecover**](CertificateApi.md#CreateCertificatesRecover) | **POST** /Certificates/Recover | Recovers the persisted certificate associated with the provided query
[**CreateCertificatesRevoke**](CertificateApi.md#CreateCertificatesRevoke) | **POST** /Certificates/Revoke | Revokes the certificates associated with the provided identifiers and associates the provided data with the revocation
[**DeleteCertificates**](CertificateApi.md#DeleteCertificates) | **DELETE** /Certificates | Deletes multiple persisted certificates by their unique ids
[**DeleteCertificatesById**](CertificateApi.md#DeleteCertificatesById) | **DELETE** /Certificates/{id} | Deletes a persisted certificate by its unique id as well as the stored private key (if present) associated with it
[**DeleteCertificatesPrivateKey**](CertificateApi.md#DeleteCertificatesPrivateKey) | **DELETE** /Certificates/PrivateKey | Deletes the persisted private keys of multiple certificates by the unique ids of the Certificates
[**DeleteCertificatesPrivateKeyById**](CertificateApi.md#DeleteCertificatesPrivateKeyById) | **DELETE** /Certificates/PrivateKey/{id} | Deletes the persisted private keys of the certificate associated with the provided identifier
[**DeleteCertificatesQuery**](CertificateApi.md#DeleteCertificatesQuery) | **DELETE** /Certificates/Query | Deletes multiple persisted certificate entities selected by a given query
[**GetCertificates**](CertificateApi.md#GetCertificates) | **GET** /Certificates | Returns all certificates according to the provided filter and output parameters
[**GetCertificatesById**](CertificateApi.md#GetCertificatesById) | **GET** /Certificates/{id} | Returns a single certificate that matches the id
[**GetCertificatesByIdHistory**](CertificateApi.md#GetCertificatesByIdHistory) | **GET** /Certificates/{id}/History | Gets the history of operations on a certificate
[**GetCertificatesByIdSecurity**](CertificateApi.md#GetCertificatesByIdSecurity) | **GET** /Certificates/{id}/Security | Gets the list of Security Identities and which permissions they have on the given certificate.
[**GetCertificatesByIdValidate**](CertificateApi.md#GetCertificatesByIdValidate) | **GET** /Certificates/{id}/Validate | Validates the certificate chain can be built.
[**GetCertificatesCSV**](CertificateApi.md#GetCertificatesCSV) | **GET** /Certificates/CSV | Returns a comma-delimited CSV file containing all certificates in the database
[**GetCertificatesIdentityAuditById**](CertificateApi.md#GetCertificatesIdentityAuditById) | **GET** /Certificates/IdentityAudit/{id} | Audit identity permissions for certificate
[**GetCertificatesLocationsById**](CertificateApi.md#GetCertificatesLocationsById) | **GET** /Certificates/Locations/{id} | Returns a list of locations the certificate is in
[**GetCertificatesMetadataCompare**](CertificateApi.md#GetCertificatesMetadataCompare) | **GET** /Certificates/Metadata/Compare | Compares the metadata value provided with the metadata value associated with the specified certificate
[**UpdateCertificatesByIdOwner**](CertificateApi.md#UpdateCertificatesByIdOwner) | **PUT** /Certificates/{id}/Owner | Changes the certificate&#39;s owner. Users must be in the current owner&#39;s role and the new owner&#39;s role
[**UpdateCertificatesMetadata**](CertificateApi.md#UpdateCertificatesMetadata) | **PUT** /Certificates/Metadata | Updates the metadata for the certificate associated with the identifier provided
[**UpdateCertificatesMetadataAll**](CertificateApi.md#UpdateCertificatesMetadataAll) | **PUT** /Certificates/Metadata/All | Updates the metadata for certificates associated with the certificate identifiers or query provided



## CreateCertificatesAnalyze

> []CSSCMSDataModelModelsCertificateDetails BuildCreateCertificatesAnalyzeRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CertificatesAnalyzeCertificateRequest(certificatesAnalyzeCertificateRequest).Execute()

Returns the public information of the certificate

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
    certificatesAnalyzeCertificateRequest := *openapiclient.NewCertificatesAnalyzeCertificateRequest() // CertificatesAnalyzeCertificateRequest | The certificate to analyze (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.BuildCreateCertificatesAnalyzeRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CertificatesAnalyzeCertificateRequest(certificatesAnalyzeCertificateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CreateCertificatesAnalyze``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCertificatesAnalyze`: []CSSCMSDataModelModelsCertificateDetails
    fmt.Fprintf(os.Stdout, "Response from `CertificateApi.CreateCertificatesAnalyze`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCertificatesAnalyzeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **certificatesAnalyzeCertificateRequest** | [**CertificatesAnalyzeCertificateRequest**](CertificatesAnalyzeCertificateRequest.md) | The certificate to analyze | 

### Return type

[**[]CSSCMSDataModelModelsCertificateDetails**](CSSCMSDataModelModelsCertificateDetails.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCertificatesDownload

> CSSCMSDataModelModelsCertificateDownloadResponse BuildCreateCertificatesDownloadRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XCertificateformat(xCertificateformat).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).CertificatesCertificateDownloadRequest(certificatesCertificateDownloadRequest).Execute()

Downloads the persisted certificate associated with the provided query



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
    xCertificateformat := "PEM" // string | Desired format [DER, PEM, P7B]
    collectionId := int32(56) // int32 | Optional certificate collection identifier used to ensure user access to the certificate (optional) (default to 0)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    certificatesCertificateDownloadRequest := *openapiclient.NewCertificatesCertificateDownloadRequest() // CertificatesCertificateDownloadRequest | Query to filter the certificate to be recovered (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.BuildCreateCertificatesDownloadRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XCertificateformat(xCertificateformat).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).CertificatesCertificateDownloadRequest(certificatesCertificateDownloadRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CreateCertificatesDownload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCertificatesDownload`: CSSCMSDataModelModelsCertificateDownloadResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateApi.CreateCertificatesDownload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCertificatesDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xCertificateformat** | **string** | Desired format [DER, PEM, P7B] | 
 **collectionId** | **int32** | Optional certificate collection identifier used to ensure user access to the certificate | [default to 0]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **certificatesCertificateDownloadRequest** | [**CertificatesCertificateDownloadRequest**](CertificatesCertificateDownloadRequest.md) | Query to filter the certificate to be recovered | 

### Return type

[**CSSCMSDataModelModelsCertificateDownloadResponse**](CSSCMSDataModelModelsCertificateDownloadResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCertificatesImport

> CSSCMSDataModelModelsCertificateImportResponseModel BuildCreateCertificatesImportRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsCertificateImportRequestModel(cSSCMSDataModelModelsCertificateImportRequestModel).Execute()

Imports the provided certificate into the Keyfactor instance, including any provided associated data

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
    cSSCMSDataModelModelsCertificateImportRequestModel := *openapiclient.NewCSSCMSDataModelModelsCertificateImportRequestModel("Certificate_example") // CSSCMSDataModelModelsCertificateImportRequestModel | Request containing the base 64 encoded string and related certificate information, such as certificate stores, metadata, and password (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.BuildCreateCertificatesImportRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsCertificateImportRequestModel(cSSCMSDataModelModelsCertificateImportRequestModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CreateCertificatesImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCertificatesImport`: CSSCMSDataModelModelsCertificateImportResponseModel
    fmt.Fprintf(os.Stdout, "Response from `CertificateApi.CreateCertificatesImport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCertificatesImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsCertificateImportRequestModel** | [**CSSCMSDataModelModelsCertificateImportRequestModel**](CSSCMSDataModelModelsCertificateImportRequestModel.md) | Request containing the base 64 encoded string and related certificate information, such as certificate stores, metadata, and password | 

### Return type

[**CSSCMSDataModelModelsCertificateImportResponseModel**](CSSCMSDataModelModelsCertificateImportResponseModel.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCertificatesRecover

> CSSCMSDataModelModelsRecoveryResponse BuildCreateCertificatesRecoverRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XCertificateformat(xCertificateformat).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).CertificatesCertificateRecoveryRequest(certificatesCertificateRecoveryRequest).Execute()

Recovers the persisted certificate associated with the provided query



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
    xCertificateformat := "PFX" // string | Desired format [PFX, PEM, ZIP, JKS]
    collectionId := int32(56) // int32 | Optional certificate collection identifier used to ensure user access to the certificate (optional) (default to 0)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    certificatesCertificateRecoveryRequest := *openapiclient.NewCertificatesCertificateRecoveryRequest("Password_example") // CertificatesCertificateRecoveryRequest | Query to filter the certificate to be recovered (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.BuildCreateCertificatesRecoverRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XCertificateformat(xCertificateformat).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).CertificatesCertificateRecoveryRequest(certificatesCertificateRecoveryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CreateCertificatesRecover``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCertificatesRecover`: CSSCMSDataModelModelsRecoveryResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateApi.CreateCertificatesRecover`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCertificatesRecoverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xCertificateformat** | **string** | Desired format [PFX, PEM, ZIP, JKS] | 
 **collectionId** | **int32** | Optional certificate collection identifier used to ensure user access to the certificate | [default to 0]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **certificatesCertificateRecoveryRequest** | [**CertificatesCertificateRecoveryRequest**](CertificatesCertificateRecoveryRequest.md) | Query to filter the certificate to be recovered | 

### Return type

[**CSSCMSDataModelModelsRecoveryResponse**](CSSCMSDataModelModelsRecoveryResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCertificatesRevoke

> CertificatesRevocationResponse BuildCreateCertificatesRevokeRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CertificatesRevokeCertificateRequest(certificatesRevokeCertificateRequest).Execute()

Revokes the certificates associated with the provided identifiers and associates the provided data with the revocation



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
    certificatesRevokeCertificateRequest := *openapiclient.NewCertificatesRevokeCertificateRequest() // CertificatesRevokeCertificateRequest | Contains the Keyfactor certificate identifiers and revocation data (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.BuildCreateCertificatesRevokeRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CertificatesRevokeCertificateRequest(certificatesRevokeCertificateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CreateCertificatesRevoke``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCertificatesRevoke`: CertificatesRevocationResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateApi.CreateCertificatesRevoke`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCertificatesRevokeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **certificatesRevokeCertificateRequest** | [**CertificatesRevokeCertificateRequest**](CertificatesRevokeCertificateRequest.md) | Contains the Keyfactor certificate identifiers and revocation data | 

### Return type

[**CertificatesRevocationResponse**](CertificatesRevocationResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCertificates

> BuildDeleteCertificatesRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).RequestBody(requestBody).Execute()

Deletes multiple persisted certificates by their unique ids



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
    collectionId := int32(56) // int32 | Optional certificate collection identifier used to ensure user access to the certificate (optional) (default to 0)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    requestBody := []int32{int32(123)} // []int32 | The array of ids for certificate that are to be deleted (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.BuildDeleteCertificatesRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.DeleteCertificates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCertificatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **collectionId** | **int32** | Optional certificate collection identifier used to ensure user access to the certificate | [default to 0]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **requestBody** | **[]int32** | The array of ids for certificate that are to be deleted | 

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


## DeleteCertificatesById

> BuildDeleteCertificatesByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Deletes a persisted certificate by its unique id as well as the stored private key (if present) associated with it

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
    id := int32(56) // int32 | Keyfactor identifier of the certificate record
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    collectionId := int32(56) // int32 | Optional certificate collection identifier used to ensure user access to the certificate (optional) (default to 0)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.BuildDeleteCertificatesByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.DeleteCertificatesById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Keyfactor identifier of the certificate record | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCertificatesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **collectionId** | **int32** | Optional certificate collection identifier used to ensure user access to the certificate | [default to 0]
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


## DeleteCertificatesPrivateKey

> BuildDeleteCertificatesPrivateKeyRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).RequestBody(requestBody).Execute()

Deletes the persisted private keys of multiple certificates by the unique ids of the Certificates

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
    collectionId := int32(56) // int32 | Optional certificate collection identifier used to ensure user access to the certificate (optional) (default to 0)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    requestBody := []int32{int32(123)} // []int32 | Keyfactor identifiers of the cetficiates for which the associated private keys should be deleted (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.BuildDeleteCertificatesPrivateKeyRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.DeleteCertificatesPrivateKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCertificatesPrivateKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **collectionId** | **int32** | Optional certificate collection identifier used to ensure user access to the certificate | [default to 0]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **requestBody** | **[]int32** | Keyfactor identifiers of the cetficiates for which the associated private keys should be deleted | 

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


## DeleteCertificatesPrivateKeyById

> BuildDeleteCertificatesPrivateKeyByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Deletes the persisted private keys of the certificate associated with the provided identifier

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
    id := int32(56) // int32 | Keyfactor identifier of the certificate for which the associated private keys should be deleted
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    collectionId := int32(56) // int32 | Optional certificate collection identifier used to ensure user access to the certificate (optional) (default to 0)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.BuildDeleteCertificatesPrivateKeyByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.DeleteCertificatesPrivateKeyById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Keyfactor identifier of the certificate for which the associated private keys should be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCertificatesPrivateKeyByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **collectionId** | **int32** | Optional certificate collection identifier used to ensure user access to the certificate | [default to 0]
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


## DeleteCertificatesQuery

> BuildDeleteCertificatesQueryRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Body(body).Execute()

Deletes multiple persisted certificate entities selected by a given query



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
    collectionId := int32(56) // int32 | Optional certificate collection identifier used to ensure user access to the certificate (optional) (default to 0)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    body := "body_example" // string | Query by which certificates should be filtered for deletion (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.BuildDeleteCertificatesQueryRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.DeleteCertificatesQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCertificatesQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **collectionId** | **int32** | Optional certificate collection identifier used to ensure user access to the certificate | [default to 0]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **body** | **string** | Query by which certificates should be filtered for deletion | 

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


## GetCertificates

> []CertificatesCertificateRetrievalResponse BuildGetCertificatesRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).IncludeRevoked(includeRevoked).IncludeExpired(includeExpired).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).CollectionId(collectionId).IncludeLocations(includeLocations).IncludeMetadata(includeMetadata).IncludeHasPrivateKey(includeHasPrivateKey).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns all certificates according to the provided filter and output parameters

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
    includeRevoked := true // bool | Select 'true' to include revoked certificates in the results (optional)
    includeExpired := true // bool | Select 'true' to include expired certificates in the results (optional)
    queryString := "queryString_example" // string | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) (optional)
    pageReturned := int32(56) // int32 | The current page within the result set to be returned (optional)
    returnLimit := int32(56) // int32 | Maximum number of records to be returned in a single call (optional)
    sortField := "sortField_example" // string | Field by which the results should be sorted (view results via Management Portal for sortable columns) (optional)
    sortAscending := openapiclient.Keyfactor.Common.QueryableExtensions.SortOrder(0) // KeyfactorCommonQueryableExtensionsSortOrder | Field sort direction [0=ascending, 1=descending] (optional)
    collectionId := int32(56) // int32 | Optional certificate collection identifier used to ensure user access to the certificate (optional) (default to 0)
    includeLocations := true // bool | Include locations data for the certificates to be returned (optional) (default to false)
    includeMetadata := true // bool | Include metadata for the certificates to be returned (optional) (default to false)
    includeHasPrivateKey := true // bool | Include whether the certificates to be returned have private keys stored in the Keyfactor database (optional) (default to false)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.BuildGetCertificatesRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).IncludeRevoked(includeRevoked).IncludeExpired(includeExpired).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).CollectionId(collectionId).IncludeLocations(includeLocations).IncludeMetadata(includeMetadata).IncludeHasPrivateKey(includeHasPrivateKey).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.GetCertificates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificates`: []CertificatesCertificateRetrievalResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateApi.GetCertificates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **includeRevoked** | **bool** | Select &#39;true&#39; to include revoked certificates in the results | 
 **includeExpired** | **bool** | Select &#39;true&#39; to include expired certificates in the results | 
 **queryString** | **string** | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) | 
 **pageReturned** | **int32** | The current page within the result set to be returned | 
 **returnLimit** | **int32** | Maximum number of records to be returned in a single call | 
 **sortField** | **string** | Field by which the results should be sorted (view results via Management Portal for sortable columns) | 
 **sortAscending** | [**KeyfactorCommonQueryableExtensionsSortOrder**](KeyfactorCommonQueryableExtensionsSortOrder.md) | Field sort direction [0&#x3D;ascending, 1&#x3D;descending] | 
 **collectionId** | **int32** | Optional certificate collection identifier used to ensure user access to the certificate | [default to 0]
 **includeLocations** | **bool** | Include locations data for the certificates to be returned | [default to false]
 **includeMetadata** | **bool** | Include metadata for the certificates to be returned | [default to false]
 **includeHasPrivateKey** | **bool** | Include whether the certificates to be returned have private keys stored in the Keyfactor database | [default to false]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]CertificatesCertificateRetrievalResponse**](CertificatesCertificateRetrievalResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificatesById

> CertificatesCertificateRetrievalResponse BuildGetCertificatesByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).IncludeLocations(includeLocations).IncludeMetadata(includeMetadata).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns a single certificate that matches the id

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
    id := int32(56) // int32 | Keyfactor certificate identifier
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    includeLocations := true // bool | Include locations data for the certificate to be returned (optional) (default to false)
    includeMetadata := true // bool | Include metadata for the certificate to be returned (optional) (default to false)
    collectionId := int32(56) // int32 | Optional certificate collection identifier used to ensure user access to the certificate (optional) (default to 0)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.BuildGetCertificatesByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).IncludeLocations(includeLocations).IncludeMetadata(includeMetadata).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.GetCertificatesById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificatesById`: CertificatesCertificateRetrievalResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateApi.GetCertificatesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Keyfactor certificate identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificatesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **includeLocations** | **bool** | Include locations data for the certificate to be returned | [default to false]
 **includeMetadata** | **bool** | Include metadata for the certificate to be returned | [default to false]
 **collectionId** | **int32** | Optional certificate collection identifier used to ensure user access to the certificate | [default to 0]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**CertificatesCertificateRetrievalResponse**](CertificatesCertificateRetrievalResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificatesByIdHistory

> []CSSCMSDataModelModelsPKICertificateOperation BuildGetCertificatesByIdHistoryRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Gets the history of operations on a certificate

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
    id := int32(56) // int32 | The Id of the certificate
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    pageReturned := int32(56) // int32 |  (optional)
    returnLimit := int32(56) // int32 |  (optional)
    sortField := "sortField_example" // string |  (optional)
    sortAscending := openapiclient.Keyfactor.Common.QueryableExtensions.SortOrder(0) // KeyfactorCommonQueryableExtensionsSortOrder |  (optional)
    collectionId := int32(56) // int32 | The collection the certificate could be in.  Defaults to no collection. (optional) (default to 0)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.BuildGetCertificatesByIdHistoryRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.GetCertificatesByIdHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificatesByIdHistory`: []CSSCMSDataModelModelsPKICertificateOperation
    fmt.Fprintf(os.Stdout, "Response from `CertificateApi.GetCertificatesByIdHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Id of the certificate | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificatesByIdHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **pageReturned** | **int32** |  | 
 **returnLimit** | **int32** |  | 
 **sortField** | **string** |  | 
 **sortAscending** | [**KeyfactorCommonQueryableExtensionsSortOrder**](KeyfactorCommonQueryableExtensionsSortOrder.md) |  | 
 **collectionId** | **int32** | The collection the certificate could be in.  Defaults to no collection. | [default to 0]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]CSSCMSDataModelModelsPKICertificateOperation**](CSSCMSDataModelModelsPKICertificateOperation.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificatesByIdSecurity

> CSSCMSDataModelModelsSecurityCertificatePermissions BuildGetCertificatesByIdSecurityRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Gets the list of Security Identities and which permissions they have on the given certificate.

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
    id := int32(56) // int32 | The Id of the certificate permissions are being checked on
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    collectionId := int32(56) // int32 | The Id of the collection the certificate belongs in. Defaults to no collection (optional) (default to 0)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.BuildGetCertificatesByIdSecurityRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.GetCertificatesByIdSecurity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificatesByIdSecurity`: CSSCMSDataModelModelsSecurityCertificatePermissions
    fmt.Fprintf(os.Stdout, "Response from `CertificateApi.GetCertificatesByIdSecurity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Id of the certificate permissions are being checked on | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificatesByIdSecurityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **collectionId** | **int32** | The Id of the collection the certificate belongs in. Defaults to no collection | [default to 0]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**CSSCMSDataModelModelsSecurityCertificatePermissions**](CSSCMSDataModelModelsSecurityCertificatePermissions.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificatesByIdValidate

> CSSCMSDataModelModelsCertificateValidationResponse BuildGetCertificatesByIdValidateRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Validates the certificate chain can be built.

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
    collectionId := int32(56) // int32 | An optional parameter for the collection Id the certificate is in.  Defaults to no collection (optional) (default to 0)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.BuildGetCertificatesByIdValidateRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.GetCertificatesByIdValidate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificatesByIdValidate`: CSSCMSDataModelModelsCertificateValidationResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateApi.GetCertificatesByIdValidate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Id of the certificate being checked | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificatesByIdValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **collectionId** | **int32** | An optional parameter for the collection Id the certificate is in.  Defaults to no collection | [default to 0]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**CSSCMSDataModelModelsCertificateValidationResponse**](CSSCMSDataModelModelsCertificateValidationResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificatesCSV

> string BuildGetCertificatesCSVRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).SortName(sortName).SortOrder(sortOrder).Query(query).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns a comma-delimited CSV file containing all certificates in the database

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
    sortName := "sortName_example" // string | Field by which the results should be sorted (view results via Management Portal for sortable columns) (optional)
    sortOrder := openapiclient.Keyfactor.Common.QueryableExtensions.SortOrder(0) // KeyfactorCommonQueryableExtensionsSortOrder | Field sort direction [0=ascending, 1=descending] (optional)
    query := "query_example" // string | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) (optional)
    collectionId := int32(56) // int32 | Optional certificate collection identifier used to ensure user access to the certificat (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.BuildGetCertificatesCSVRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).SortName(sortName).SortOrder(sortOrder).Query(query).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.GetCertificatesCSV``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificatesCSV`: string
    fmt.Fprintf(os.Stdout, "Response from `CertificateApi.GetCertificatesCSV`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificatesCSVRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **sortName** | **string** | Field by which the results should be sorted (view results via Management Portal for sortable columns) | 
 **sortOrder** | [**KeyfactorCommonQueryableExtensionsSortOrder**](KeyfactorCommonQueryableExtensionsSortOrder.md) | Field sort direction [0&#x3D;ascending, 1&#x3D;descending] | 
 **query** | **string** | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) | 
 **collectionId** | **int32** | Optional certificate collection identifier used to ensure user access to the certificat | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

**string**

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificatesIdentityAuditById

> []CertificatesCertificateIdentityAuditResponse BuildGetCertificatesIdentityAuditByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    collectionId := int32(56) // int32 | An optional parameter for the collection Id the certificate is in.  Defaults to no collection (optional) (default to 0)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.BuildGetCertificatesIdentityAuditByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.GetCertificatesIdentityAuditById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificatesIdentityAuditById`: []CertificatesCertificateIdentityAuditResponse
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
 **collectionId** | **int32** | An optional parameter for the collection Id the certificate is in.  Defaults to no collection | [default to 0]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]CertificatesCertificateIdentityAuditResponse**](CertificatesCertificateIdentityAuditResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificatesLocationsById

> CertificatesCertificateLocationsResponse BuildGetCertificatesLocationsByIdRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Returns a list of locations the certificate is in

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
    id := int32(56) // int32 | Keyfactor certificate identifier
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    collectionId := int32(56) // int32 | Optional certificate collection identifier used to ensure user access to the certificate (optional) (default to 0)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.BuildGetCertificatesLocationsByIdRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.GetCertificatesLocationsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificatesLocationsById`: CertificatesCertificateLocationsResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateApi.GetCertificatesLocationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Keyfactor certificate identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificatesLocationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **collectionId** | **int32** | Optional certificate collection identifier used to ensure user access to the certificate | [default to 0]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**CertificatesCertificateLocationsResponse**](CertificatesCertificateLocationsResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificatesMetadataCompare

> bool BuildGetCertificatesMetadataCompareRequest(ctx).CertificateId(certificateId).MetadataFieldName(metadataFieldName).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Value(value).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Compares the metadata value provided with the metadata value associated with the specified certificate

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
    certificateId := int32(56) // int32 | Certificate identifier
    metadataFieldName := "metadataFieldName_example" // string | Metadata field being compared
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    value := "value_example" // string | Value to compare against (optional)
    collectionId := int32(56) // int32 | Optional certificate collection identifier used to ensure user access to the certificate (optional) (default to 0)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.BuildGetCertificatesMetadataCompareRequest(context.Background()).CertificateId(certificateId).MetadataFieldName(metadataFieldName).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Value(value).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.GetCertificatesMetadataCompare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificatesMetadataCompare`: bool
    fmt.Fprintf(os.Stdout, "Response from `CertificateApi.GetCertificatesMetadataCompare`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificatesMetadataCompareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **certificateId** | **int32** | Certificate identifier | 
 **metadataFieldName** | **string** | Metadata field being compared | 
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **value** | **string** | Value to compare against | 
 **collectionId** | **int32** | Optional certificate collection identifier used to ensure user access to the certificate | [default to 0]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

**bool**

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCertificatesByIdOwner

> BuildUpdateCertificatesByIdOwnerRequest(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).CertificatesOwnerRequest(certificatesOwnerRequest).Execute()

Changes the certificate's owner. Users must be in the current owner's role and the new owner's role

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
    id := int32(56) // int32 | Id of the certificate
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient]
    collectionId := int32(56) // int32 | An optional parameter for the collection Id the certificate is in. Defaults to no collection (optional) (default to 0)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    certificatesOwnerRequest := *openapiclient.NewCertificatesOwnerRequest() // CertificatesOwnerRequest | Security role identifier for the role to assign ownership. If removing the owner, leave both empty. (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.BuildUpdateCertificatesByIdOwnerRequest(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).CertificatesOwnerRequest(certificatesOwnerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.UpdateCertificatesByIdOwner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id of the certificate | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCertificatesByIdOwnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **collectionId** | **int32** | An optional parameter for the collection Id the certificate is in. Defaults to no collection | [default to 0]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **certificatesOwnerRequest** | [**CertificatesOwnerRequest**](CertificatesOwnerRequest.md) | Security role identifier for the role to assign ownership. If removing the owner, leave both empty. | 

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


## UpdateCertificatesMetadata

> BuildUpdateCertificatesMetadataRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsMetadataUpdateRequest(cSSCMSDataModelModelsMetadataUpdateRequest).Execute()

Updates the metadata for the certificate associated with the identifier provided

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
    collectionId := int32(56) // int32 | Optional certificate collection identifier used to ensure user access to the certificate (optional) (default to 0)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    cSSCMSDataModelModelsMetadataUpdateRequest := *openapiclient.NewCSSCMSDataModelModelsMetadataUpdateRequest(map[string]string{"key": "Inner_example"}) // CSSCMSDataModelModelsMetadataUpdateRequest | Contains the Keyfactor certificate identifier and the metadata to be updated (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.BuildUpdateCertificatesMetadataRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsMetadataUpdateRequest(cSSCMSDataModelModelsMetadataUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.UpdateCertificatesMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCertificatesMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **collectionId** | **int32** | Optional certificate collection identifier used to ensure user access to the certificate | [default to 0]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsMetadataUpdateRequest** | [**CSSCMSDataModelModelsMetadataUpdateRequest**](CSSCMSDataModelModelsMetadataUpdateRequest.md) | Contains the Keyfactor certificate identifier and the metadata to be updated | 

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


## UpdateCertificatesMetadataAll

> BuildUpdateCertificatesMetadataAllRequest(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsMetadataAllUpdateRequest(cSSCMSDataModelModelsMetadataAllUpdateRequest).Execute()

Updates the metadata for certificates associated with the certificate identifiers or query provided

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
    collectionId := int32(56) // int32 | Optional certificate collection identifier used to ensure user access to the certificate (optional) (default to 0)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    cSSCMSDataModelModelsMetadataAllUpdateRequest := *openapiclient.NewCSSCMSDataModelModelsMetadataAllUpdateRequest([]openapiclient.CSSCMSDataModelModelsMetadataSingleUpdateRequest{*openapiclient.NewCSSCMSDataModelModelsMetadataSingleUpdateRequest()}) // CSSCMSDataModelModelsMetadataAllUpdateRequest | Contains the Keyfactor certificate identifier and the metadata to be updated (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.BuildUpdateCertificatesMetadataAllRequest(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsMetadataAllUpdateRequest(cSSCMSDataModelModelsMetadataAllUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.UpdateCertificatesMetadataAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCertificatesMetadataAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | 
 **collectionId** | **int32** | Optional certificate collection identifier used to ensure user access to the certificate | [default to 0]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsMetadataAllUpdateRequest** | [**CSSCMSDataModelModelsMetadataAllUpdateRequest**](CSSCMSDataModelModelsMetadataAllUpdateRequest.md) | Contains the Keyfactor certificate identifier and the metadata to be updated | 

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


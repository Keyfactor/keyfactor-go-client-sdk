# \CertificateApi

All URIs are relative to */Keyfactor/API*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CertificatesAnalyzePost**](CertificateApi.md#CertificatesAnalyzePost) | **Post** /Certificates/Analyze | Returns the public information of the certificate
[**CertificatesCSVGet**](CertificateApi.md#CertificatesCSVGet) | **Get** /Certificates/CSV | Returns a comma-delimited CSV file containing all certificates in the database
[**CertificatesDelete**](CertificateApi.md#CertificatesDelete) | **Delete** /Certificates | Deletes multiple persisted certificates by their unique ids
[**CertificatesDownloadPost**](CertificateApi.md#CertificatesDownloadPost) | **Post** /Certificates/Download | Downloads the persisted certificate associated with the provided query
[**CertificatesGet**](CertificateApi.md#CertificatesGet) | **Get** /Certificates | Returns all certificates according to the provided filter and output parameters
[**CertificatesIdDelete**](CertificateApi.md#CertificatesIdDelete) | **Delete** /Certificates/{id} | Deletes a persisted certificate by its unique id as well as the stored private key (if present) associated with it
[**CertificatesIdGet**](CertificateApi.md#CertificatesIdGet) | **Get** /Certificates/{id} | Returns a single certificate that matches the id
[**CertificatesIdHistoryGet**](CertificateApi.md#CertificatesIdHistoryGet) | **Get** /Certificates/{id}/History | Gets the history of operations on a certificate
[**CertificatesIdSecurityGet**](CertificateApi.md#CertificatesIdSecurityGet) | **Get** /Certificates/{id}/Security | Gets the list of Security Identities and which permissions they have on the given certificate.
[**CertificatesIdValidateGet**](CertificateApi.md#CertificatesIdValidateGet) | **Get** /Certificates/{id}/Validate | Validates the certificate chain can be built.
[**CertificatesIdentityAuditIdGet**](CertificateApi.md#CertificatesIdentityAuditIdGet) | **Get** /Certificates/IdentityAudit/{id} | Audit identity permissions for certificate
[**CertificatesImportPost**](CertificateApi.md#CertificatesImportPost) | **Post** /Certificates/Import | Imports the provided certificate into the Keyfactor instance, including any provided associated data
[**CertificatesLocationsIdGet**](CertificateApi.md#CertificatesLocationsIdGet) | **Get** /Certificates/Locations/{id} | Returns a list of locations the certificate is in
[**CertificatesMetadataAllPut**](CertificateApi.md#CertificatesMetadataAllPut) | **Put** /Certificates/Metadata/All | Updates the metadata for certificates associated with the certificate identifiers or query provided
[**CertificatesMetadataCompareGet**](CertificateApi.md#CertificatesMetadataCompareGet) | **Get** /Certificates/Metadata/Compare | Compares the metadata value provided with the metadata value associated with the specified certificate
[**CertificatesMetadataPut**](CertificateApi.md#CertificatesMetadataPut) | **Put** /Certificates/Metadata | Updates the metadata for the certificate associated with the identifier provided
[**CertificatesPrivateKeyDelete**](CertificateApi.md#CertificatesPrivateKeyDelete) | **Delete** /Certificates/PrivateKey | Deletes the persisted private keys of multiple certificates by the unique ids of the Certificates
[**CertificatesPrivateKeyIdDelete**](CertificateApi.md#CertificatesPrivateKeyIdDelete) | **Delete** /Certificates/PrivateKey/{id} | Deletes the persisted private key of the certificate associated with the provided identifier
[**CertificatesQueryDelete**](CertificateApi.md#CertificatesQueryDelete) | **Delete** /Certificates/Query | Deletes multiple persisted certificate entities selected by a given query
[**CertificatesRecoverPost**](CertificateApi.md#CertificatesRecoverPost) | **Post** /Certificates/Recover | Recovers the persisted certificate associated with the provided query
[**CertificatesRevokePost**](CertificateApi.md#CertificatesRevokePost) | **Post** /Certificates/Revoke | Revokes the certificates associated with the provided identifiers and associates the provided data with the revocation



## CertificatesAnalyzePost

> []CSSCMSDataModelModelsCertificateDetails CertificatesAnalyzePost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsCertificatesAnalyzeCertificateRequest(keyfactorWebKeyfactorApiModelsCertificatesAnalyzeCertificateRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsCertificatesAnalyzeCertificateRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsCertificatesAnalyzeCertificateRequest() // KeyfactorWebKeyfactorApiModelsCertificatesAnalyzeCertificateRequest | The certificate to analyze (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.CertificatesAnalyzePost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsCertificatesAnalyzeCertificateRequest(keyfactorWebKeyfactorApiModelsCertificatesAnalyzeCertificateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificatesAnalyzePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificatesAnalyzePost`: []CSSCMSDataModelModelsCertificateDetails
    fmt.Fprintf(os.Stdout, "Response from `CertificateApi.CertificatesAnalyzePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificatesAnalyzePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsCertificatesAnalyzeCertificateRequest** | [**KeyfactorWebKeyfactorApiModelsCertificatesAnalyzeCertificateRequest**](KeyfactorWebKeyfactorApiModelsCertificatesAnalyzeCertificateRequest.md) | The certificate to analyze | 

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


## CertificatesCSVGet

> string CertificatesCSVGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).SortName(sortName).SortOrder(sortOrder).Query(query).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    sortName := "sortName_example" // string | Field by which the results should be sorted (view results via Management Portal for sortable columns) (optional)
    sortOrder := openapiclient.Keyfactor.Common.QueryableExtensionsSortOrder(0) // KeyfactorCommonQueryableExtensionsSortOrder | Field sort direction [0=ascending, 1=descending] (optional)
    query := "query_example" // string | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) (optional)
    collectionId := int32(56) // int32 | Optional certificate collection identifier used to ensure user access to the certificat (optional)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.CertificatesCSVGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).SortName(sortName).SortOrder(sortOrder).Query(query).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificatesCSVGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificatesCSVGet`: string
    fmt.Fprintf(os.Stdout, "Response from `CertificateApi.CertificatesCSVGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificatesCSVGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
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


## CertificatesDelete

> CertificatesDelete(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).RequestBody(requestBody).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    collectionId := int32(56) // int32 | Optional certificate collection identifier used to ensure user access to the certificate (optional) (default to 0)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    requestBody := []int32{int32(123)} // []int32 | The array of ids for certificate that are to be deleted (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.CertificatesDelete(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificatesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificatesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
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


## CertificatesDownloadPost

> CSSCMSDataModelModelsCertificateDownloadResponse CertificatesDownloadPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest(keyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    collectionId := int32(56) // int32 | Optional certificate collection identifier used to ensure user access to the certificate (optional) (default to 0)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest() // KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest | Query to filter the certificate to be recovered (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.CertificatesDownloadPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest(keyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificatesDownloadPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificatesDownloadPost`: CSSCMSDataModelModelsCertificateDownloadResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateApi.CertificatesDownloadPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificatesDownloadPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **collectionId** | **int32** | Optional certificate collection identifier used to ensure user access to the certificate | [default to 0]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest** | [**KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest**](KeyfactorWebKeyfactorApiModelsCertificatesCertificateDownloadRequest.md) | Query to filter the certificate to be recovered | 

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


## CertificatesGet

> []CSSCMSDataModelModelsCertificateRetrievalResponse CertificatesGet(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).IncludeRevoked(includeRevoked).IncludeExpired(includeExpired).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).CollectionId(collectionId).IncludeLocations(includeLocations).IncludeMetadata(includeMetadata).IncludeHasPrivateKey(includeHasPrivateKey).Verbose(verbose).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    includeRevoked := true // bool | Select 'true' to include revoked certificates in the results (optional)
    includeExpired := true // bool | Select 'true' to include expired certificates in the results (optional)
    queryString := "queryString_example" // string | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) (optional)
    pageReturned := int32(56) // int32 | The current page within the result set to be returned (optional)
    returnLimit := int32(56) // int32 | Maximum number of records to be returned in a single call (optional)
    sortField := "sortField_example" // string | Field by which the results should be sorted (view results via Management Portal for sortable columns) (optional)
    sortAscending := openapiclient.Keyfactor.Common.QueryableExtensionsSortOrder(0) // KeyfactorCommonQueryableExtensionsSortOrder | Field sort direction [0=ascending, 1=descending] (optional)
    collectionId := int32(56) // int32 | Optional certificate collection identifier used to ensure user access to the certificate (optional) (default to 0)
    includeLocations := true // bool | Include locations data for the certificates to be returned (optional) (default to false)
    includeMetadata := true // bool | Include metadata for the certificates to be returned (optional) (default to false)
    includeHasPrivateKey := true // bool | Include whether the certificates to be returned have private keys stored in the Keyfactor database (optional) (default to false)
    verbose := int32(56) // int32 |  (optional) (default to 0)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.CertificatesGet(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).IncludeRevoked(includeRevoked).IncludeExpired(includeExpired).QueryString(queryString).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).CollectionId(collectionId).IncludeLocations(includeLocations).IncludeMetadata(includeMetadata).IncludeHasPrivateKey(includeHasPrivateKey).Verbose(verbose).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificatesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificatesGet`: []CSSCMSDataModelModelsCertificateRetrievalResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateApi.CertificatesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificatesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
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
 **verbose** | **int32** |  | [default to 0]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]CSSCMSDataModelModelsCertificateRetrievalResponse**](CSSCMSDataModelModelsCertificateRetrievalResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificatesIdDelete

> CertificatesIdDelete(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    collectionId := int32(56) // int32 | Optional certificate collection identifier used to ensure user access to the certificate (optional) (default to 0)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.CertificatesIdDelete(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificatesIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCertificatesIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
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


## CertificatesIdGet

> CSSCMSDataModelModelsCertificateRetrievalResponse CertificatesIdGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).IncludeLocations(includeLocations).IncludeMetadata(includeMetadata).CollectionId(collectionId).Verbose(verbose).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    includeLocations := true // bool | Include locations data for the certificate to be returned (optional) (default to false)
    includeMetadata := true // bool | Include metadata for the certificate to be returned (optional) (default to false)
    collectionId := int32(56) // int32 | Optional certificate collection identifier used to ensure user access to the certificate (optional) (default to 0)
    verbose := int32(56) // int32 |  (optional) (default to 0)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.CertificatesIdGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).IncludeLocations(includeLocations).IncludeMetadata(includeMetadata).CollectionId(collectionId).Verbose(verbose).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificatesIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificatesIdGet`: CSSCMSDataModelModelsCertificateRetrievalResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateApi.CertificatesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Keyfactor certificate identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificatesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **includeLocations** | **bool** | Include locations data for the certificate to be returned | [default to false]
 **includeMetadata** | **bool** | Include metadata for the certificate to be returned | [default to false]
 **collectionId** | **int32** | Optional certificate collection identifier used to ensure user access to the certificate | [default to 0]
 **verbose** | **int32** |  | [default to 0]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**CSSCMSDataModelModelsCertificateRetrievalResponse**](CSSCMSDataModelModelsCertificateRetrievalResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificatesIdHistoryGet

> []map[string]interface{} CertificatesIdHistoryGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    pageReturned := int32(56) // int32 |  (optional)
    returnLimit := int32(56) // int32 |  (optional)
    sortField := "sortField_example" // string |  (optional)
    sortAscending := openapiclient.Keyfactor.Common.QueryableExtensionsSortOrder(0) // KeyfactorCommonQueryableExtensionsSortOrder |  (optional)
    collectionId := int32(56) // int32 | The collection the certificate could be in.  Defaults to no collection. (optional) (default to 0)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.CertificatesIdHistoryGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).PageReturned(pageReturned).ReturnLimit(returnLimit).SortField(sortField).SortAscending(sortAscending).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificatesIdHistoryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificatesIdHistoryGet`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CertificateApi.CertificatesIdHistoryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Id of the certificate | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificatesIdHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **pageReturned** | **int32** |  | 
 **returnLimit** | **int32** |  | 
 **sortField** | **string** |  | 
 **sortAscending** | [**KeyfactorCommonQueryableExtensionsSortOrder**](KeyfactorCommonQueryableExtensionsSortOrder.md) |  | 
 **collectionId** | **int32** | The collection the certificate could be in.  Defaults to no collection. | [default to 0]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

**[]map[string]interface{}**

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificatesIdSecurityGet

> CSSCMSDataModelModelsSecurityCertificatePermissions CertificatesIdSecurityGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    collectionId := int32(56) // int32 | The Id of the collection the certificate belongs in. Defaults to no collection (optional) (default to 0)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.CertificatesIdSecurityGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificatesIdSecurityGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificatesIdSecurityGet`: CSSCMSDataModelModelsSecurityCertificatePermissions
    fmt.Fprintf(os.Stdout, "Response from `CertificateApi.CertificatesIdSecurityGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Id of the certificate permissions are being checked on | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificatesIdSecurityGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
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


## CertificatesIdValidateGet

> CSSCMSDataModelModelsCertificateValidationResponse CertificatesIdValidateGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    collectionId := int32(56) // int32 | An optional parameter for the collectin Id the certificate is in.  Defaults to no collection (optional) (default to 0)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.CertificatesIdValidateGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificatesIdValidateGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificatesIdValidateGet`: CSSCMSDataModelModelsCertificateValidationResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateApi.CertificatesIdValidateGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Id of the certificate being checked | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificatesIdValidateGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **collectionId** | **int32** | An optional parameter for the collectin Id the certificate is in.  Defaults to no collection | [default to 0]
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


## CertificatesIdentityAuditIdGet

> []KeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse CertificatesIdentityAuditIdGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    collectionId := int32(56) // int32 | An optional parameter for the collectin Id the certificate is in.  Defaults to no collection (optional) (default to 0)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.CertificatesIdentityAuditIdGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificatesIdentityAuditIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificatesIdentityAuditIdGet`: []KeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateApi.CertificatesIdentityAuditIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Id of the certificate being checked | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificatesIdentityAuditIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **collectionId** | **int32** | An optional parameter for the collectin Id the certificate is in.  Defaults to no collection | [default to 0]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**[]KeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse**](KeyfactorWebKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificatesImportPost

> CSSCMSDataModelModelsCertificateImportResponseModel CertificatesImportPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsCertificateImportRequestModel(cSSCMSDataModelModelsCertificateImportRequestModel).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    cSSCMSDataModelModelsCertificateImportRequestModel := *openapiclient.NewCSSCMSDataModelModelsCertificateImportRequestModel("Certificate_example") // CSSCMSDataModelModelsCertificateImportRequestModel | Request containing the base 64 encoded string and related certificate information, such as certificate stores, metadata, and password (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.CertificatesImportPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsCertificateImportRequestModel(cSSCMSDataModelModelsCertificateImportRequestModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificatesImportPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificatesImportPost`: CSSCMSDataModelModelsCertificateImportResponseModel
    fmt.Fprintf(os.Stdout, "Response from `CertificateApi.CertificatesImportPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificatesImportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
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


## CertificatesLocationsIdGet

> KeyfactorWebKeyfactorApiModelsCertificatesCertificateLocationsResponse CertificatesLocationsIdGet(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    collectionId := int32(56) // int32 | Optional certificate collection identifier used to ensure user access to the certificate (optional) (default to 0)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.CertificatesLocationsIdGet(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificatesLocationsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificatesLocationsIdGet`: KeyfactorWebKeyfactorApiModelsCertificatesCertificateLocationsResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateApi.CertificatesLocationsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Keyfactor certificate identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificatesLocationsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **collectionId** | **int32** | Optional certificate collection identifier used to ensure user access to the certificate | [default to 0]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 

### Return type

[**KeyfactorWebKeyfactorApiModelsCertificatesCertificateLocationsResponse**](KeyfactorWebKeyfactorApiModelsCertificatesCertificateLocationsResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificatesMetadataAllPut

> CertificatesMetadataAllPut(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsMetadataAllUpdateRequest(cSSCMSDataModelModelsMetadataAllUpdateRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    collectionId := int32(56) // int32 | Optional certificate collection identifier used to ensure user access to the certificate (optional) (default to 0)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    cSSCMSDataModelModelsMetadataAllUpdateRequest := *openapiclient.NewCSSCMSDataModelModelsMetadataAllUpdateRequest([]openapiclient.CSSCMSDataModelModelsMetadataSingleUpdateRequest{*openapiclient.NewCSSCMSDataModelModelsMetadataSingleUpdateRequest()}) // CSSCMSDataModelModelsMetadataAllUpdateRequest | Contains the Keyfactor certificate identifier and the metadata to be updated (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.CertificatesMetadataAllPut(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsMetadataAllUpdateRequest(cSSCMSDataModelModelsMetadataAllUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificatesMetadataAllPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificatesMetadataAllPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
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


## CertificatesMetadataCompareGet

> bool CertificatesMetadataCompareGet(ctx).CertificateId(certificateId).MetadataFieldName(metadataFieldName).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Value(value).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    value := "value_example" // string | Value to compare against (optional)
    collectionId := int32(56) // int32 | Optional certificate collection identifier used to ensure user access to the certificate (optional) (default to 0)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.CertificatesMetadataCompareGet(context.Background()).CertificateId(certificateId).MetadataFieldName(metadataFieldName).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Value(value).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificatesMetadataCompareGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificatesMetadataCompareGet`: bool
    fmt.Fprintf(os.Stdout, "Response from `CertificateApi.CertificatesMetadataCompareGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificatesMetadataCompareGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **certificateId** | **int32** | Certificate identifier | 
 **metadataFieldName** | **string** | Metadata field being compared | 
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
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


## CertificatesMetadataPut

> CertificatesMetadataPut(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsMetadataUpdateRequest(cSSCMSDataModelModelsMetadataUpdateRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    collectionId := int32(56) // int32 | Optional certificate collection identifier used to ensure user access to the certificate (optional) (default to 0)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    cSSCMSDataModelModelsMetadataUpdateRequest := *openapiclient.NewCSSCMSDataModelModelsMetadataUpdateRequest(map[string]string{"key": "Inner_example"}) // CSSCMSDataModelModelsMetadataUpdateRequest | Contains the Keyfactor certificate identifier and the metadata to be updated (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.CertificatesMetadataPut(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsMetadataUpdateRequest(cSSCMSDataModelModelsMetadataUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificatesMetadataPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificatesMetadataPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
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


## CertificatesPrivateKeyDelete

> CertificatesPrivateKeyDelete(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).RequestBody(requestBody).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    collectionId := int32(56) // int32 | Optional certificate collection identifier used to ensure user access to the certificate (optional) (default to 0)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    requestBody := []int32{int32(123)} // []int32 | Keyfactor identifiers of the cetficiates for which the associated private keys should be deleted (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.CertificatesPrivateKeyDelete(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificatesPrivateKeyDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificatesPrivateKeyDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
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


## CertificatesPrivateKeyIdDelete

> CertificatesPrivateKeyIdDelete(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Deletes the persisted private key of the certificate associated with the provided identifier

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
    id := int32(56) // int32 | Keyfactor identifier of the certificate for which the associated private key should be deleted
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    collectionId := int32(56) // int32 | Optional certificate collection identifier used to ensure user access to the certificate (optional) (default to 0)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.CertificatesPrivateKeyIdDelete(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificatesPrivateKeyIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Keyfactor identifier of the certificate for which the associated private key should be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificatesPrivateKeyIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
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


## CertificatesQueryDelete

> CertificatesQueryDelete(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Body(body).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    collectionId := int32(56) // int32 | Optional certificate collection identifier used to ensure user access to the certificate (optional) (default to 0)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    body := "body_example" // string | Query by which certificates should be filtered for deletion (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.CertificatesQueryDelete(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificatesQueryDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificatesQueryDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
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


## CertificatesRecoverPost

> CSSCMSDataModelModelsRecoveryResponse CertificatesRecoverPost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest(keyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    collectionId := int32(56) // int32 | Optional certificate collection identifier used to ensure user access to the certificate (optional) (default to 0)
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    keyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest := *openapiclient.NewKeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest("Password_example") // KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest | Query to filter the certificate to be recovered (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.CertificatesRecoverPost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest(keyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificatesRecoverPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificatesRecoverPost`: CSSCMSDataModelModelsRecoveryResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateApi.CertificatesRecoverPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificatesRecoverPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **collectionId** | **int32** | Optional certificate collection identifier used to ensure user access to the certificate | [default to 0]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **keyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest** | [**KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest**](KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest.md) | Query to filter the certificate to be recovered | 

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


## CertificatesRevokePost

> CSSCMSDataModelModelsRevocationRevocationResponse CertificatesRevokePost(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsRevokeCertificateRequest(cSSCMSDataModelModelsRevokeCertificateRequest).Execute()

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
    xKeyfactorRequestedWith := "APIClient" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    xKeyfactorApiVersion := "1.0" // string | Desired version of the api, if not provided defaults to v1 (optional)
    cSSCMSDataModelModelsRevokeCertificateRequest := *openapiclient.NewCSSCMSDataModelModelsRevokeCertificateRequest() // CSSCMSDataModelModelsRevokeCertificateRequest | Contains the Keyfactor certificate identifiers and revocation data (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.CertificatesRevokePost(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).CSSCMSDataModelModelsRevokeCertificateRequest(cSSCMSDataModelModelsRevokeCertificateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificatesRevokePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificatesRevokePost`: CSSCMSDataModelModelsRevocationRevocationResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateApi.CertificatesRevokePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificatesRevokePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | 
 **cSSCMSDataModelModelsRevokeCertificateRequest** | [**CSSCMSDataModelModelsRevokeCertificateRequest**](CSSCMSDataModelModelsRevokeCertificateRequest.md) | Contains the Keyfactor certificate identifiers and revocation data | 

### Return type

[**CSSCMSDataModelModelsRevocationRevocationResponse**](CSSCMSDataModelModelsRevocationRevocationResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


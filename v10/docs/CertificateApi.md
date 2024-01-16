# \CertificateApi

All URIs are relative to *http://keyfactor.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CertificateAnalyzeCert**](CertificateApi.md#CertificateAnalyzeCert) | **Post** /Certificates/Analyze | Returns the public information of the certificate
[**CertificateCertificateHistory**](CertificateApi.md#CertificateCertificateHistory) | **Get** /Certificates/{id}/History | Gets the history of operations on a certificate
[**CertificateCompareMetadata**](CertificateApi.md#CertificateCompareMetadata) | **Get** /Certificates/Metadata/Compare | Compares the metadata value provided with the metadata value associated with the specified certificate
[**CertificateDeleteByQuery**](CertificateApi.md#CertificateDeleteByQuery) | **Delete** /Certificates/Query | Deletes multiple persisted certificate entities selected by a given query
[**CertificateDeleteCertificate**](CertificateApi.md#CertificateDeleteCertificate) | **Delete** /Certificates/{id} | Deletes a persisted certificate by its unique id as well as the stored private key (if present) associated with it
[**CertificateDeleteCertificates**](CertificateApi.md#CertificateDeleteCertificates) | **Delete** /Certificates | Deletes multiple persisted certificates by their unique ids
[**CertificateDeletePrivateKeys0**](CertificateApi.md#CertificateDeletePrivateKeys0) | **Delete** /Certificates/PrivateKey | Deletes the persisted private keys of multiple certificates by the unique ids of the Certificates
[**CertificateDeletePrivateKeys1**](CertificateApi.md#CertificateDeletePrivateKeys1) | **Delete** /Certificates/PrivateKey/{id} | Deletes the persisted private key of the certificate associated with the provided identifier
[**CertificateDownloadCertificateAsync**](CertificateApi.md#CertificateDownloadCertificateAsync) | **Post** /Certificates/Download | Downloads the persisted certificate associated with the provided query
[**CertificateGetCertificate**](CertificateApi.md#CertificateGetCertificate) | **Get** /Certificates/{id} | Returns a single certificate that matches the id
[**CertificateGetCertificateLocations**](CertificateApi.md#CertificateGetCertificateLocations) | **Get** /Certificates/Locations/{id} | Returns a list of locations the certificate is in
[**CertificateGetCertificateSecurity**](CertificateApi.md#CertificateGetCertificateSecurity) | **Get** /Certificates/{id}/Security | Gets the list of Security Identities and which permissions they have on the given certificate.
[**CertificateIdentityAudit**](CertificateApi.md#CertificateIdentityAudit) | **Get** /Certificates/IdentityAudit/{id} | Audit identity permissions for certificate
[**CertificatePostImportCertificate**](CertificateApi.md#CertificatePostImportCertificate) | **Post** /Certificates/Import | Imports the provided certificate into the Keyfactor instance, including any provided associated data
[**CertificateQueryCertificates**](CertificateApi.md#CertificateQueryCertificates) | **Get** /Certificates | Returns all certificates according to the provided filter and output parameters
[**CertificateRecoverCertificateAsync**](CertificateApi.md#CertificateRecoverCertificateAsync) | **Post** /Certificates/Recover | Recovers the persisted certificate associated with the provided query
[**CertificateRevoke**](CertificateApi.md#CertificateRevoke) | **Post** /Certificates/Revoke | Revokes the certificates associated with the provided identifiers and associates the provided data with the revocation
[**CertificateRevokeAll**](CertificateApi.md#CertificateRevokeAll) | **Post** /Certificates/RevokeAll | Revokes the certificates associated with the provided query and Collection Id and associates the provided data with the revocation
[**CertificateUpdateAllMetadata**](CertificateApi.md#CertificateUpdateAllMetadata) | **Put** /Certificates/Metadata/All | Updates the metadata for certificates associated with the certificate identifiers or query provided
[**CertificateUpdateMetadata**](CertificateApi.md#CertificateUpdateMetadata) | **Put** /Certificates/Metadata | Updates the metadata for the certificate associated with the identifier provided
[**CertificateValidateCertificate**](CertificateApi.md#CertificateValidateCertificate) | **Get** /Certificates/{id}/Validate | Validates the certificate chain can be built.



## CertificateAnalyzeCert

> []ModelsCertificateDetails CertificateAnalyzeCert(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Request(request).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    request := *openapiclient.NewKeyfactorApiModelsCertificatesAnalyzeCertificateRequest() // KeyfactorApiModelsCertificatesAnalyzeCertificateRequest | The certificate to analyze
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.CertificateAnalyzeCert(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Request(request).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificateAnalyzeCert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateAnalyzeCert`: []ModelsCertificateDetails
    fmt.Fprintf(os.Stdout, "Response from `CertificateApi.CertificateAnalyzeCert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateAnalyzeCertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **request** | [**KeyfactorApiModelsCertificatesAnalyzeCertificateRequest**](KeyfactorApiModelsCertificatesAnalyzeCertificateRequest.md) | The certificate to analyze | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**[]ModelsCertificateDetails**](ModelsCertificateDetails.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateCertificateHistory

> []ModelsPKICertificateOperation CertificateCertificateHistory(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).QueryPageReturned(queryPageReturned).QueryReturnLimit(queryReturnLimit).QuerySortField(querySortField).QuerySortAscending(querySortAscending).Execute()

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
    id := int64(789) // int64 | The Id of the certificate
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    collectionId := int64(789) // int64 | The collection the certificate could be in.  Defaults to no collection. (optional)
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")
    queryPageReturned := int64(789) // int64 | The current page within the result set to be returned (optional)
    queryReturnLimit := int64(789) // int64 | Maximum number of records to be returned in a single call (optional)
    querySortField := "querySortField_example" // string | Field by which the results should be sorted (OperationStart, OperationEnd, UserName) (optional)
    querySortAscending := int64(789) // int64 | Field sort direction [0=ascending, 1=descending] (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.CertificateCertificateHistory(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).QueryPageReturned(queryPageReturned).QueryReturnLimit(queryReturnLimit).QuerySortField(querySortField).QuerySortAscending(querySortAscending).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificateCertificateHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateCertificateHistory`: []ModelsPKICertificateOperation
    fmt.Fprintf(os.Stdout, "Response from `CertificateApi.CertificateCertificateHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The Id of the certificate | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificateCertificateHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **collectionId** | **int64** | The collection the certificate could be in.  Defaults to no collection. | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]
 **queryPageReturned** | **int64** | The current page within the result set to be returned | 
 **queryReturnLimit** | **int64** | Maximum number of records to be returned in a single call | 
 **querySortField** | **string** | Field by which the results should be sorted (OperationStart, OperationEnd, UserName) | 
 **querySortAscending** | **int64** | Field sort direction [0&#x3D;ascending, 1&#x3D;descending] | 

### Return type

[**[]ModelsPKICertificateOperation**](ModelsPKICertificateOperation.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateCompareMetadata

> bool CertificateCompareMetadata(ctx).CertificateId(certificateId).MetadataFieldName(metadataFieldName).Value(value).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    certificateId := int64(789) // int64 | Certificate identifier
    metadataFieldName := "metadataFieldName_example" // string | Metadata field being compared
    value := "value_example" // string | Value to compare against
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    collectionId := int64(789) // int64 | Optional certificate collection identifier used to ensure user access to the certificate (optional)
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.CertificateCompareMetadata(context.Background()).CertificateId(certificateId).MetadataFieldName(metadataFieldName).Value(value).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificateCompareMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateCompareMetadata`: bool
    fmt.Fprintf(os.Stdout, "Response from `CertificateApi.CertificateCompareMetadata`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateCompareMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **certificateId** | **int64** | Certificate identifier | 
 **metadataFieldName** | **string** | Metadata field being compared | 
 **value** | **string** | Value to compare against | 
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **collectionId** | **int64** | Optional certificate collection identifier used to ensure user access to the certificate | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

**bool**

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateDeleteByQuery

> CertificateDeleteByQuery(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Sq(sq).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    sq := "sq_example" // string | Query by which certificates should be filtered for deletion
    collectionId := int64(789) // int64 | Optional certificate collection identifier used to ensure user access to the certificate (optional)
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.CertificateDeleteByQuery(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Sq(sq).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificateDeleteByQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateDeleteByQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **sq** | **string** | Query by which certificates should be filtered for deletion | 
 **collectionId** | **int64** | Optional certificate collection identifier used to ensure user access to the certificate | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateDeleteCertificate

> CertificateDeleteCertificate(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    id := int64(789) // int64 | Keyfactor identifier of the certificate record
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    collectionId := int64(789) // int64 | Optional certificate collection identifier used to ensure user access to the certificate (optional)
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.CertificateDeleteCertificate(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificateDeleteCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Keyfactor identifier of the certificate record | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificateDeleteCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **collectionId** | **int64** | Optional certificate collection identifier used to ensure user access to the certificate | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

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


## CertificateDeleteCertificates

> CertificateDeleteCertificates(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Ids(ids).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    ids := []int64{int64(123)} // []int64 | The array of ids for certificate that are to be deleted
    collectionId := int64(789) // int64 | Optional certificate collection identifier used to ensure user access to the certificate (optional)
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.CertificateDeleteCertificates(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Ids(ids).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificateDeleteCertificates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateDeleteCertificatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **ids** | **[]int64** | The array of ids for certificate that are to be deleted | 
 **collectionId** | **int64** | Optional certificate collection identifier used to ensure user access to the certificate | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateDeletePrivateKeys0

> CertificateDeletePrivateKeys0(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Ids(ids).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    ids := []int64{int64(123)} // []int64 | Keyfactor identifiers of the cetficiates for which the associated private keys should be deleted
    collectionId := int64(789) // int64 | Optional certificate collection identifier used to ensure user access to the certificate (optional)
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.CertificateDeletePrivateKeys0(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Ids(ids).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificateDeletePrivateKeys0``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateDeletePrivateKeys0Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **ids** | **[]int64** | Keyfactor identifiers of the cetficiates for which the associated private keys should be deleted | 
 **collectionId** | **int64** | Optional certificate collection identifier used to ensure user access to the certificate | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateDeletePrivateKeys1

> CertificateDeletePrivateKeys1(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    id := int64(789) // int64 | Keyfactor identifier of the certificate for which the associated private key should be deleted
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    collectionId := int64(789) // int64 | Optional certificate collection identifier used to ensure user access to the certificate (optional)
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.CertificateDeletePrivateKeys1(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificateDeletePrivateKeys1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Keyfactor identifier of the certificate for which the associated private key should be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificateDeletePrivateKeys1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **collectionId** | **int64** | Optional certificate collection identifier used to ensure user access to the certificate | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

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


## CertificateDownloadCertificateAsync

> ModelsCertificateDownloadResponse CertificateDownloadCertificateAsync(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Rq(rq).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    rq := *openapiclient.NewModelsCertificateDownloadRequest() // ModelsCertificateDownloadRequest | Query to filter the certificate to be recovered
    collectionId := int64(789) // int64 | Optional certificate collection identifier used to ensure user access to the certificate (optional)
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.CertificateDownloadCertificateAsync(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Rq(rq).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificateDownloadCertificateAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateDownloadCertificateAsync`: ModelsCertificateDownloadResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateApi.CertificateDownloadCertificateAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateDownloadCertificateAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **rq** | [**ModelsCertificateDownloadRequest**](ModelsCertificateDownloadRequest.md) | Query to filter the certificate to be recovered | 
 **collectionId** | **int64** | Optional certificate collection identifier used to ensure user access to the certificate | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsCertificateDownloadResponse**](ModelsCertificateDownloadResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateGetCertificate

> ModelsCertificateRetrievalResponse CertificateGetCertificate(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).IncludeLocations(includeLocations).IncludeMetadata(includeMetadata).CollectionId(collectionId).Verbose(verbose).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    id := int64(789) // int64 | Keyfactor certificate identifier
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    includeLocations := true // bool | Include locations data for the certificate to be returned (optional)
    includeMetadata := true // bool | Include metadata for the certificate to be returned (optional)
    collectionId := int64(789) // int64 | Optional certificate collection identifier used to ensure user access to the certificate (optional)
    verbose := int64(789) // int64 |  (optional)
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.CertificateGetCertificate(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).IncludeLocations(includeLocations).IncludeMetadata(includeMetadata).CollectionId(collectionId).Verbose(verbose).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificateGetCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateGetCertificate`: ModelsCertificateRetrievalResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateApi.CertificateGetCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Keyfactor certificate identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificateGetCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **includeLocations** | **bool** | Include locations data for the certificate to be returned | 
 **includeMetadata** | **bool** | Include metadata for the certificate to be returned | 
 **collectionId** | **int64** | Optional certificate collection identifier used to ensure user access to the certificate | 
 **verbose** | **int64** |  | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsCertificateRetrievalResponse**](ModelsCertificateRetrievalResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateGetCertificateLocations

> KeyfactorApiModelsCertificatesCertificateLocationsResponse CertificateGetCertificateLocations(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    id := int64(789) // int64 | Keyfactor certificate identifier
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    collectionId := int64(789) // int64 | Optional certificate collection identifier used to ensure user access to the certificate (optional)
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.CertificateGetCertificateLocations(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificateGetCertificateLocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateGetCertificateLocations`: KeyfactorApiModelsCertificatesCertificateLocationsResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateApi.CertificateGetCertificateLocations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Keyfactor certificate identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificateGetCertificateLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **collectionId** | **int64** | Optional certificate collection identifier used to ensure user access to the certificate | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**KeyfactorApiModelsCertificatesCertificateLocationsResponse**](KeyfactorApiModelsCertificatesCertificateLocationsResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateGetCertificateSecurity

> ModelsSecurityCertificatePermissions CertificateGetCertificateSecurity(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    id := int64(789) // int64 | The Id of the certificate permissions are being checked on
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    collectionId := int64(789) // int64 | The Id of the collection the certificate belongs in. Defaults to no collection (optional)
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.CertificateGetCertificateSecurity(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificateGetCertificateSecurity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateGetCertificateSecurity`: ModelsSecurityCertificatePermissions
    fmt.Fprintf(os.Stdout, "Response from `CertificateApi.CertificateGetCertificateSecurity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The Id of the certificate permissions are being checked on | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificateGetCertificateSecurityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **collectionId** | **int64** | The Id of the collection the certificate belongs in. Defaults to no collection | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsSecurityCertificatePermissions**](ModelsSecurityCertificatePermissions.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateIdentityAudit

> []KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse CertificateIdentityAudit(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    id := int64(789) // int64 | The Id of the certificate being checked
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    collectionId := int64(789) // int64 | An optional parameter for the collectin Id the certificate is in.  Defaults to no collection (optional)
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.CertificateIdentityAudit(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificateIdentityAudit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateIdentityAudit`: []KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateApi.CertificateIdentityAudit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The Id of the certificate being checked | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificateIdentityAuditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **collectionId** | **int64** | An optional parameter for the collectin Id the certificate is in.  Defaults to no collection | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**[]KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse**](KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificatePostImportCertificate

> ModelsCertificateImportResponseModel CertificatePostImportCertificate(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Req(req).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    req := *openapiclient.NewModelsCertificateImportRequestModel("Certificate_example") // ModelsCertificateImportRequestModel | Request containing the base 64 encoded string and related certificate information, such as certificate stores, metadata, and password
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.CertificatePostImportCertificate(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Req(req).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificatePostImportCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificatePostImportCertificate`: ModelsCertificateImportResponseModel
    fmt.Fprintf(os.Stdout, "Response from `CertificateApi.CertificatePostImportCertificate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificatePostImportCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **req** | [**ModelsCertificateImportRequestModel**](ModelsCertificateImportRequestModel.md) | Request containing the base 64 encoded string and related certificate information, such as certificate stores, metadata, and password | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsCertificateImportResponseModel**](ModelsCertificateImportResponseModel.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateQueryCertificates

> []ModelsCertificateRetrievalResponse CertificateQueryCertificates(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).IncludeLocations(includeLocations).IncludeMetadata(includeMetadata).IncludeHasPrivateKey(includeHasPrivateKey).Verbose(verbose).XKeyfactorApiVersion(xKeyfactorApiVersion).PqQueryString(pqQueryString).PqPageReturned(pqPageReturned).PqReturnLimit(pqReturnLimit).PqSortField(pqSortField).PqSortAscending(pqSortAscending).PqIncludeRevoked(pqIncludeRevoked).PqIncludeExpired(pqIncludeExpired).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    collectionId := int64(789) // int64 | Optional certificate collection identifier used to ensure user access to the certificate (optional)
    includeLocations := true // bool | Include locations data for the certificates to be returned (optional)
    includeMetadata := true // bool | Include metadata for the certificates to be returned (optional)
    includeHasPrivateKey := true // bool | Include whether the certificates to be returned have private keys stored in the Keyfactor database (optional)
    verbose := int64(789) // int64 |  (optional)
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")
    pqQueryString := "pqQueryString_example" // string | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) (optional)
    pqPageReturned := int64(789) // int64 | The current page within the result set to be returned (optional)
    pqReturnLimit := int64(789) // int64 | Maximum number of records to be returned in a single call (optional)
    pqSortField := "pqSortField_example" // string | Field by which the results should be sorted (view results via Management Portal for sortable columns) (optional)
    pqSortAscending := int64(789) // int64 | Field sort direction [0=ascending, 1=descending] (optional)
    pqIncludeRevoked := true // bool | Select 'true' to include revoked certificates in the results (optional)
    pqIncludeExpired := true // bool | Select 'true' to include expired certificates in the results (optional)

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.CertificateQueryCertificates(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).IncludeLocations(includeLocations).IncludeMetadata(includeMetadata).IncludeHasPrivateKey(includeHasPrivateKey).Verbose(verbose).XKeyfactorApiVersion(xKeyfactorApiVersion).PqQueryString(pqQueryString).PqPageReturned(pqPageReturned).PqReturnLimit(pqReturnLimit).PqSortField(pqSortField).PqSortAscending(pqSortAscending).PqIncludeRevoked(pqIncludeRevoked).PqIncludeExpired(pqIncludeExpired).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificateQueryCertificates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateQueryCertificates`: []ModelsCertificateRetrievalResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateApi.CertificateQueryCertificates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateQueryCertificatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **collectionId** | **int64** | Optional certificate collection identifier used to ensure user access to the certificate | 
 **includeLocations** | **bool** | Include locations data for the certificates to be returned | 
 **includeMetadata** | **bool** | Include metadata for the certificates to be returned | 
 **includeHasPrivateKey** | **bool** | Include whether the certificates to be returned have private keys stored in the Keyfactor database | 
 **verbose** | **int64** |  | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]
 **pqQueryString** | **string** | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) | 
 **pqPageReturned** | **int64** | The current page within the result set to be returned | 
 **pqReturnLimit** | **int64** | Maximum number of records to be returned in a single call | 
 **pqSortField** | **string** | Field by which the results should be sorted (view results via Management Portal for sortable columns) | 
 **pqSortAscending** | **int64** | Field sort direction [0&#x3D;ascending, 1&#x3D;descending] | 
 **pqIncludeRevoked** | **bool** | Select &#39;true&#39; to include revoked certificates in the results | 
 **pqIncludeExpired** | **bool** | Select &#39;true&#39; to include expired certificates in the results | 

### Return type

[**[]ModelsCertificateRetrievalResponse**](ModelsCertificateRetrievalResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateRecoverCertificateAsync

> ModelsRecoveryResponse CertificateRecoverCertificateAsync(ctx).XCertificateformat(xCertificateformat).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Rq(rq).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xCertificateformat := "xCertificateformat_example" // string | Desired format [PFX, PEM] (default to "PEM")
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    rq := *openapiclient.NewModelsCertificateRecoveryRequest("Password_example") // ModelsCertificateRecoveryRequest | Query to filter the certificate to be recovered
    collectionId := int64(789) // int64 | Optional certificate collection identifier used to ensure user access to the certificate (optional)
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.CertificateRecoverCertificateAsync(context.Background()).XCertificateformat(xCertificateformat).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Rq(rq).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificateRecoverCertificateAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateRecoverCertificateAsync`: ModelsRecoveryResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateApi.CertificateRecoverCertificateAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateRecoverCertificateAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xCertificateformat** | **string** | Desired format [PFX, PEM] | [default to &quot;PEM&quot;]
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **rq** | [**ModelsCertificateRecoveryRequest**](ModelsCertificateRecoveryRequest.md) | Query to filter the certificate to be recovered | 
 **collectionId** | **int64** | Optional certificate collection identifier used to ensure user access to the certificate | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsRecoveryResponse**](ModelsRecoveryResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateRevoke

> ModelsRevocationRevocationResponse CertificateRevoke(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Request(request).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    request := *openapiclient.NewModelsRevokeCertificateRequest() // ModelsRevokeCertificateRequest | Contains the Keyfactor certificate identifiers and revocation data
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.CertificateRevoke(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Request(request).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificateRevoke``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateRevoke`: ModelsRevocationRevocationResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateApi.CertificateRevoke`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateRevokeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **request** | [**ModelsRevokeCertificateRequest**](ModelsRevokeCertificateRequest.md) | Contains the Keyfactor certificate identifiers and revocation data | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsRevocationRevocationResponse**](ModelsRevocationRevocationResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateRevokeAll

> ModelsRevocationRevocationResponse CertificateRevokeAll(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Request(request).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

Revokes the certificates associated with the provided query and Collection Id and associates the provided data with the revocation



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
    request := *openapiclient.NewModelsRevokeAllCertificatesRequest(int64(123), "Comment_example") // ModelsRevokeAllCertificatesRequest | Contains the Keyfactor Query and revocation data
    collectionId := int64(789) // int64 | A collection Id to be used for permissions and part of the query to revoke certificates (optional)
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.CertificateRevokeAll(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Request(request).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificateRevokeAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateRevokeAll`: ModelsRevocationRevocationResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateApi.CertificateRevokeAll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateRevokeAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **request** | [**ModelsRevokeAllCertificatesRequest**](ModelsRevokeAllCertificatesRequest.md) | Contains the Keyfactor Query and revocation data | 
 **collectionId** | **int64** | A collection Id to be used for permissions and part of the query to revoke certificates | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsRevocationRevocationResponse**](ModelsRevocationRevocationResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateUpdateAllMetadata

> CertificateUpdateAllMetadata(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).MetadataUpdate(metadataUpdate).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    metadataUpdate := *openapiclient.NewModelsMetadataAllUpdateRequest([]openapiclient.ModelsMetadataSingleUpdateRequest{*openapiclient.NewModelsMetadataSingleUpdateRequest()}) // ModelsMetadataAllUpdateRequest | Contains the Keyfactor certificate identifier and the metadata to be updated
    collectionId := int64(789) // int64 | Optional certificate collection identifier used to ensure user access to the certificate (optional)
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.CertificateUpdateAllMetadata(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).MetadataUpdate(metadataUpdate).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificateUpdateAllMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateUpdateAllMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **metadataUpdate** | [**ModelsMetadataAllUpdateRequest**](ModelsMetadataAllUpdateRequest.md) | Contains the Keyfactor certificate identifier and the metadata to be updated | 
 **collectionId** | **int64** | Optional certificate collection identifier used to ensure user access to the certificate | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateUpdateMetadata

> CertificateUpdateMetadata(ctx).XKeyfactorRequestedWith(xKeyfactorRequestedWith).MetadataUpdate(metadataUpdate).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    metadataUpdate := *openapiclient.NewModelsMetadataUpdateRequest(map[string]string{"key": "Inner_example"}) // ModelsMetadataUpdateRequest | Contains the Keyfactor certificate identifier and the metadata to be updated
    collectionId := int64(789) // int64 | Optional certificate collection identifier used to ensure user access to the certificate (optional)
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.CertificateUpdateMetadata(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).MetadataUpdate(metadataUpdate).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificateUpdateMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateUpdateMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **metadataUpdate** | [**ModelsMetadataUpdateRequest**](ModelsMetadataUpdateRequest.md) | Contains the Keyfactor certificate identifier and the metadata to be updated | 
 **collectionId** | **int64** | Optional certificate collection identifier used to ensure user access to the certificate | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateValidateCertificate

> ModelsCertificateValidationResponse CertificateValidateCertificate(ctx, id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

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
    id := int64(789) // int64 | The Id of the certificate being checked
    xKeyfactorRequestedWith := "xKeyfactorRequestedWith_example" // string | Type of the request [XMLHttpRequest, APIClient] (default to "APIClient")
    collectionId := int64(789) // int64 | An optional parameter for the collectin Id the certificate is in.  Defaults to no collection (optional)
    xKeyfactorApiVersion := "xKeyfactorApiVersion_example" // string | Desired version of the api, if not provided defaults to v1 (optional) (default to "1")

    configuration := openapiclient.NewConfiguration(make(map[string]string))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateApi.CertificateValidateCertificate(context.Background(), id).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CollectionId(collectionId).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificateValidateCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateValidateCertificate`: ModelsCertificateValidationResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificateApi.CertificateValidateCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The Id of the certificate being checked | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificateValidateCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xKeyfactorRequestedWith** | **string** | Type of the request [XMLHttpRequest, APIClient] | [default to &quot;APIClient&quot;]
 **collectionId** | **int64** | An optional parameter for the collectin Id the certificate is in.  Defaults to no collection | 
 **xKeyfactorApiVersion** | **string** | Desired version of the api, if not provided defaults to v1 | [default to &quot;1&quot;]

### Return type

[**ModelsCertificateValidationResponse**](ModelsCertificateValidationResponse.md)

### Authorization

[basicAuth](../README.md#Configuration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


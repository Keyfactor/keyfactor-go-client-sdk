/*
Copyright 2022 Keyfactor
Licensed under the Apache License, Version 2.0 (the "License"); you may
not use this file except in compliance with the License.  You may obtain a
copy of the License at http://www.apache.org/licenses/LICENSE-2.0.  Unless
required by applicable law or agreed to in writing, software distributed
under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES
OR CONDITIONS OF ANY KIND, either express or implied. See the License for
thespecific language governing permissions and limitations under the
License.
Keyfactor-v1

This reference serves to document REST-based methods to manage and integrate with Keyfactor. In addition, an embedded interface allows for the execution of calls against the current Keyfactor API instance.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package keyfactor

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// CSRGenerationApiService CSRGenerationApi service
type CSRGenerationApiService service

type ApiCSRGenerationDeleteCSRRequest struct {
	ctx                     context.Context
	ApiService              *CSRGenerationApiService
	id                      int32
	xKeyfactorRequestedWith *string
	xKeyfactorApiVersion    *string
}

// Type of the request [XMLHttpRequest, APIClient]
func (r ApiCSRGenerationDeleteCSRRequest) XKeyfactorRequestedWith(xKeyfactorRequestedWith string) ApiCSRGenerationDeleteCSRRequest {
	r.xKeyfactorRequestedWith = &xKeyfactorRequestedWith
	return r
}

// Desired version of the api, if not provided defaults to v1
func (r ApiCSRGenerationDeleteCSRRequest) XKeyfactorApiVersion(xKeyfactorApiVersion string) ApiCSRGenerationDeleteCSRRequest {
	r.xKeyfactorApiVersion = &xKeyfactorApiVersion
	return r
}

func (r ApiCSRGenerationDeleteCSRRequest) Execute() (*http.Response, error) {
	return r.ApiService.CSRGenerationDeleteCSRExecute(r)
}

/*
CSRGenerationDeleteCSR Deletes a CSR associated with the provided identifier

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Keyfactor identifer of the CSR to be deleted
	@return ApiCSRGenerationDeleteCSRRequest
*/
func (a *CSRGenerationApiService) CSRGenerationDeleteCSR(ctx context.Context, id int32) ApiCSRGenerationDeleteCSRRequest {
	requestedWith := "APIClient"
	version := "1"
	return ApiCSRGenerationDeleteCSRRequest{
		ApiService:              a,
		ctx:                     ctx,
		xKeyfactorRequestedWith: &requestedWith,
		xKeyfactorApiVersion:    &version,
		id:                      id,
	}
}

// Execute executes the request
func (a *CSRGenerationApiService) CSRGenerationDeleteCSRExecute(r ApiCSRGenerationDeleteCSRRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath := "/KeyfactorAPI"

	localVarPath := localBasePath + "/CSRGeneration/Pending/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xKeyfactorRequestedWith == nil {
		return nil, reportError("xKeyfactorRequestedWith is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xKeyfactorApiVersion != nil {
		parameterAddToQuery(localVarHeaderParams, "x-keyfactor-api-version", r.xKeyfactorApiVersion, "")
	}
	parameterAddToQuery(localVarHeaderParams, "x-keyfactor-requested-with", r.xKeyfactorRequestedWith, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiCSRGenerationDeleteCSRsRequest struct {
	ctx                     context.Context
	ApiService              *CSRGenerationApiService
	xKeyfactorRequestedWith *string
	ids                     *[]int32
	xKeyfactorApiVersion    *string
}

// Type of the request [XMLHttpRequest, APIClient]
func (r ApiCSRGenerationDeleteCSRsRequest) XKeyfactorRequestedWith(xKeyfactorRequestedWith string) ApiCSRGenerationDeleteCSRsRequest {
	r.xKeyfactorRequestedWith = &xKeyfactorRequestedWith
	return r
}

// Array of Keyfactor identifiers for the CSRs to be deleted
func (r ApiCSRGenerationDeleteCSRsRequest) Ids(ids []int32) ApiCSRGenerationDeleteCSRsRequest {
	r.ids = &ids
	return r
}

// Desired version of the api, if not provided defaults to v1
func (r ApiCSRGenerationDeleteCSRsRequest) XKeyfactorApiVersion(xKeyfactorApiVersion string) ApiCSRGenerationDeleteCSRsRequest {
	r.xKeyfactorApiVersion = &xKeyfactorApiVersion
	return r
}

func (r ApiCSRGenerationDeleteCSRsRequest) Execute() (*http.Response, error) {
	return r.ApiService.CSRGenerationDeleteCSRsExecute(r)
}

/*
CSRGenerationDeleteCSRs Deletes the CSRs associated with the provided identifiers

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCSRGenerationDeleteCSRsRequest
*/
func (a *CSRGenerationApiService) CSRGenerationDeleteCSRs(ctx context.Context) ApiCSRGenerationDeleteCSRsRequest {
	requestedWith := "APIClient"
	version := "1"
	return ApiCSRGenerationDeleteCSRsRequest{
		ApiService:              a,
		ctx:                     ctx,
		xKeyfactorRequestedWith: &requestedWith,
		xKeyfactorApiVersion:    &version,
	}
}

// Execute executes the request
func (a *CSRGenerationApiService) CSRGenerationDeleteCSRsExecute(r ApiCSRGenerationDeleteCSRsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath := "/KeyfactorAPI"

	localVarPath := localBasePath + "/CSRGeneration/Pending"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xKeyfactorRequestedWith == nil {
		return nil, reportError("xKeyfactorRequestedWith is required and must be specified")
	}
	if r.ids == nil {
		return nil, reportError("ids is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "text/json", "application/xml", "text/xml", "application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xKeyfactorApiVersion != nil {
		parameterAddToQuery(localVarHeaderParams, "x-keyfactor-api-version", r.xKeyfactorApiVersion, "")
	}
	parameterAddToQuery(localVarHeaderParams, "x-keyfactor-requested-with", r.xKeyfactorRequestedWith, "")
	// body params
	localVarPostBody = r.ids
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiCSRGenerationDownloadRequest struct {
	ctx                     context.Context
	ApiService              *CSRGenerationApiService
	id                      int32
	xKeyfactorRequestedWith *string
	xKeyfactorApiVersion    *string
}

// Type of the request [XMLHttpRequest, APIClient]
func (r ApiCSRGenerationDownloadRequest) XKeyfactorRequestedWith(xKeyfactorRequestedWith string) ApiCSRGenerationDownloadRequest {
	r.xKeyfactorRequestedWith = &xKeyfactorRequestedWith
	return r
}

// Desired version of the api, if not provided defaults to v1
func (r ApiCSRGenerationDownloadRequest) XKeyfactorApiVersion(xKeyfactorApiVersion string) ApiCSRGenerationDownloadRequest {
	r.xKeyfactorApiVersion = &xKeyfactorApiVersion
	return r
}

func (r ApiCSRGenerationDownloadRequest) Execute() (*ModelsCSRGenerationResponseModel, *http.Response, error) {
	return r.ApiService.CSRGenerationDownloadExecute(r)
}

/*
CSRGenerationDownload Returns a previously generated CSR associated with the provided identifier

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Keyfactor identifier of the CSR
	@return ApiCSRGenerationDownloadRequest
*/
func (a *CSRGenerationApiService) CSRGenerationDownload(ctx context.Context, id int32) ApiCSRGenerationDownloadRequest {
	requestedWith := "APIClient"
	version := "1"
	return ApiCSRGenerationDownloadRequest{
		ApiService:              a,
		ctx:                     ctx,
		xKeyfactorRequestedWith: &requestedWith,
		xKeyfactorApiVersion:    &version,
		id:                      id,
	}
}

// Execute executes the request
//
//	@return ModelsCSRGenerationResponseModel
func (a *CSRGenerationApiService) CSRGenerationDownloadExecute(r ApiCSRGenerationDownloadRequest) (*ModelsCSRGenerationResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ModelsCSRGenerationResponseModel
	)

	localBasePath := "/KeyfactorAPI"

	localVarPath := localBasePath + "/CSRGeneration/Pending/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xKeyfactorRequestedWith == nil {
		return localVarReturnValue, nil, reportError("xKeyfactorRequestedWith is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/json", "application/xml", "text/xml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xKeyfactorApiVersion != nil {
		parameterAddToQuery(localVarHeaderParams, "x-keyfactor-api-version", r.xKeyfactorApiVersion, "")
	}
	parameterAddToQuery(localVarHeaderParams, "x-keyfactor-requested-with", r.xKeyfactorRequestedWith, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCSRGenerationGetPendingCSRsRequest struct {
	ctx                     context.Context
	ApiService              *CSRGenerationApiService
	xKeyfactorRequestedWith *string
	xKeyfactorApiVersion    *string
	sqQueryString           *string
	sqPageReturned          *int32
	sqReturnLimit           *int32
	sqSortField             *string
	sqSortAscending         *int32
}

// Type of the request [XMLHttpRequest, APIClient]
func (r ApiCSRGenerationGetPendingCSRsRequest) XKeyfactorRequestedWith(xKeyfactorRequestedWith string) ApiCSRGenerationGetPendingCSRsRequest {
	r.xKeyfactorRequestedWith = &xKeyfactorRequestedWith
	return r
}

// Desired version of the api, if not provided defaults to v1
func (r ApiCSRGenerationGetPendingCSRsRequest) XKeyfactorApiVersion(xKeyfactorApiVersion string) ApiCSRGenerationGetPendingCSRsRequest {
	r.xKeyfactorApiVersion = &xKeyfactorApiVersion
	return r
}

// Contents of the query (ex: field1 -eq value1 AND field2 -gt value2)
func (r ApiCSRGenerationGetPendingCSRsRequest) SqQueryString(sqQueryString string) ApiCSRGenerationGetPendingCSRsRequest {
	r.sqQueryString = &sqQueryString
	return r
}

// The current page within the result set to be returned
func (r ApiCSRGenerationGetPendingCSRsRequest) SqPageReturned(sqPageReturned int32) ApiCSRGenerationGetPendingCSRsRequest {
	r.sqPageReturned = &sqPageReturned
	return r
}

// Maximum number of records to be returned in a single call
func (r ApiCSRGenerationGetPendingCSRsRequest) SqReturnLimit(sqReturnLimit int32) ApiCSRGenerationGetPendingCSRsRequest {
	r.sqReturnLimit = &sqReturnLimit
	return r
}

// Field by which the results should be sorted (view results via Management Portal for sortable columns)
func (r ApiCSRGenerationGetPendingCSRsRequest) SqSortField(sqSortField string) ApiCSRGenerationGetPendingCSRsRequest {
	r.sqSortField = &sqSortField
	return r
}

// Field sort direction [0&#x3D;ascending, 1&#x3D;descending]
func (r ApiCSRGenerationGetPendingCSRsRequest) SqSortAscending(sqSortAscending int32) ApiCSRGenerationGetPendingCSRsRequest {
	r.sqSortAscending = &sqSortAscending
	return r
}

func (r ApiCSRGenerationGetPendingCSRsRequest) Execute() ([]ModelsPendingCSRResponse, *http.Response, error) {
	return r.ApiService.CSRGenerationGetPendingCSRsExecute(r)
}

/*
CSRGenerationGetPendingCSRs Returns a list of the currently pending CSRs according to the provided query

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCSRGenerationGetPendingCSRsRequest
*/
func (a *CSRGenerationApiService) CSRGenerationGetPendingCSRs(ctx context.Context) ApiCSRGenerationGetPendingCSRsRequest {
	requestedWith := "APIClient"
	version := "1"
	return ApiCSRGenerationGetPendingCSRsRequest{
		ApiService:              a,
		ctx:                     ctx,
		xKeyfactorRequestedWith: &requestedWith,
		xKeyfactorApiVersion:    &version,
	}
}

// Execute executes the request
//
//	@return []ModelsPendingCSRResponse
func (a *CSRGenerationApiService) CSRGenerationGetPendingCSRsExecute(r ApiCSRGenerationGetPendingCSRsRequest) ([]ModelsPendingCSRResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []ModelsPendingCSRResponse
	)

	localBasePath := "/KeyfactorAPI"

	localVarPath := localBasePath + "/CSRGeneration/Pending"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xKeyfactorRequestedWith == nil {
		return localVarReturnValue, nil, reportError("xKeyfactorRequestedWith is required and must be specified")
	}

	if r.sqQueryString != nil {
		parameterAddToQuery(localVarQueryParams, "sq.queryString", r.sqQueryString, "")
	}
	if r.sqPageReturned != nil {
		parameterAddToQuery(localVarQueryParams, "sq.pageReturned", r.sqPageReturned, "")
	}
	if r.sqReturnLimit != nil {
		parameterAddToQuery(localVarQueryParams, "sq.returnLimit", r.sqReturnLimit, "")
	}
	if r.sqSortField != nil {
		parameterAddToQuery(localVarQueryParams, "sq.sortField", r.sqSortField, "")
	}
	if r.sqSortAscending != nil {
		parameterAddToQuery(localVarQueryParams, "sq.sortAscending", r.sqSortAscending, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/json", "application/xml", "text/xml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xKeyfactorApiVersion != nil {
		parameterAddToQuery(localVarHeaderParams, "x-keyfactor-api-version", r.xKeyfactorApiVersion, "")
	}
	parameterAddToQuery(localVarHeaderParams, "x-keyfactor-requested-with", r.xKeyfactorRequestedWith, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCSRGenerationPostGenerateRequest struct {
	ctx                     context.Context
	ApiService              *CSRGenerationApiService
	xKeyfactorRequestedWith *string
	context                 *ModelsEnrollmentCSRGenerationRequest
	xKeyfactorApiVersion    *string
}

// Type of the request [XMLHttpRequest, APIClient]
func (r ApiCSRGenerationPostGenerateRequest) XKeyfactorRequestedWith(xKeyfactorRequestedWith string) ApiCSRGenerationPostGenerateRequest {
	r.xKeyfactorRequestedWith = &xKeyfactorRequestedWith
	return r
}

// CSR properties used to define the request - Key type [RSA, ECC], Key sizes (ex: RSA 1024, 2048, 4096/ECC 256, 384, 521), template short name or OID
func (r ApiCSRGenerationPostGenerateRequest) Context(context ModelsEnrollmentCSRGenerationRequest) ApiCSRGenerationPostGenerateRequest {
	r.context = &context
	return r
}

// Desired version of the api, if not provided defaults to v1
func (r ApiCSRGenerationPostGenerateRequest) XKeyfactorApiVersion(xKeyfactorApiVersion string) ApiCSRGenerationPostGenerateRequest {
	r.xKeyfactorApiVersion = &xKeyfactorApiVersion
	return r
}

func (r ApiCSRGenerationPostGenerateRequest) Execute() (*ModelsCSRContents, *http.Response, error) {
	return r.ApiService.CSRGenerationPostGenerateExecute(r)
}

/*
CSRGenerationPostGenerate Generates a CSR according the properties provided

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCSRGenerationPostGenerateRequest
*/
func (a *CSRGenerationApiService) CSRGenerationPostGenerate(ctx context.Context) ApiCSRGenerationPostGenerateRequest {
	requestedWith := "APIClient"
	version := "1"
	return ApiCSRGenerationPostGenerateRequest{
		ApiService:              a,
		ctx:                     ctx,
		xKeyfactorRequestedWith: &requestedWith,
		xKeyfactorApiVersion:    &version,
	}
}

// Execute executes the request
//
//	@return ModelsCSRContents
func (a *CSRGenerationApiService) CSRGenerationPostGenerateExecute(r ApiCSRGenerationPostGenerateRequest) (*ModelsCSRContents, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ModelsCSRContents
	)

	localBasePath := "/KeyfactorAPI"

	localVarPath := localBasePath + "/CSRGeneration/Generate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xKeyfactorRequestedWith == nil {
		return localVarReturnValue, nil, reportError("xKeyfactorRequestedWith is required and must be specified")
	}
	if r.context == nil {
		return localVarReturnValue, nil, reportError("context is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xKeyfactorApiVersion != nil {
		parameterAddToQuery(localVarHeaderParams, "x-keyfactor-api-version", r.xKeyfactorApiVersion, "")
	}
	parameterAddToQuery(localVarHeaderParams, "x-keyfactor-requested-with", r.xKeyfactorRequestedWith, "")
	// body params
	localVarPostBody = r.context
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

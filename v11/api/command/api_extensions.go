/*

Copyright 2023 Keyfactor
Licensed under the Apache License, Version 2.0 (the License); you may
not use this file except in compliance with the License.  You may obtain a
copy of the License at http://www.apache.org/licenses/LICENSE-2.0.  Unless
required by applicable law or agreed to in writing, software distributed
under the License is distributed on an AS IS BASIS, WITHOUT WARRANTIES
OR CONDITIONS OF ANY KIND, either express or implied. See the License for
the specific language governing permissions and limitations under the
License.

Keyfactor API Reference and Utility

<p>This page provides a utility through which the Keyfactor API endpoints can be called and results returned.                                                           It is intended to be used primarily for validation, testing and workflow development.                                                           It also serves secondarily as documentation for the API.</p>                                                          <p>If you would like to view documentation containing details on the Keyfactor API and endpoints,                                                           please refer to the Web API section of the Keyfactor Command documentation.</p>

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package command

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// ExtensionsApiService ExtensionsApi service
type ExtensionsApiService service

type ApiExtensionsScriptsGetRequest struct {
	ctx                     context.Context
	ApiService              *ExtensionsApiService
	xKeyfactorRequestedWith *string
	queryString             *string
	pageReturned            *int32
	returnLimit             *int32
	sortField               *string
	sortAscending           *int32
	xKeyfactorApiVersion    *string
}

// Type of the request [XMLHttpRequest, APIClient]
func (r ApiExtensionsScriptsGetRequest) XKeyfactorRequestedWith(xKeyfactorRequestedWith string) ApiExtensionsScriptsGetRequest {
	r.xKeyfactorRequestedWith = &xKeyfactorRequestedWith
	return r
}

func (r ApiExtensionsScriptsGetRequest) QueryString(queryString string) ApiExtensionsScriptsGetRequest {
	r.queryString = &queryString
	return r
}

func (r ApiExtensionsScriptsGetRequest) PageReturned(pageReturned int32) ApiExtensionsScriptsGetRequest {
	r.pageReturned = &pageReturned
	return r
}

func (r ApiExtensionsScriptsGetRequest) ReturnLimit(returnLimit int32) ApiExtensionsScriptsGetRequest {
	r.returnLimit = &returnLimit
	return r
}

func (r ApiExtensionsScriptsGetRequest) SortField(sortField string) ApiExtensionsScriptsGetRequest {
	r.sortField = &sortField
	return r
}

func (r ApiExtensionsScriptsGetRequest) SortAscending(sortAscending int32) ApiExtensionsScriptsGetRequest {
	r.sortAscending = &sortAscending
	return r
}

// Desired version of the api, if not provided defaults to v1
func (r ApiExtensionsScriptsGetRequest) XKeyfactorApiVersion(xKeyfactorApiVersion string) ApiExtensionsScriptsGetRequest {
	r.xKeyfactorApiVersion = &xKeyfactorApiVersion
	return r
}

func (r ApiExtensionsScriptsGetRequest) Execute() ([]KeyfactorWebKeyfactorApiModelsScriptsScriptQueryResponse, *http.Response, error) {
	return r.ApiService.ExtensionsScriptsGetExecute(r)
}

/*
ExtensionsScriptsGet Returns all scripts according to the provided filter and output parameters

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiExtensionsScriptsGetRequest
*/
func (a *ExtensionsApiService) ExtensionsScriptsGet(ctx context.Context) ApiExtensionsScriptsGetRequest {

	xKeyfactorRequestedWith := "APIClient"

	return ApiExtensionsScriptsGetRequest{
		ApiService: a,
		ctx:        ctx,

		xKeyfactorRequestedWith: &xKeyfactorRequestedWith,
	}
}

// Execute executes the request
//
//	@return []KeyfactorWebKeyfactorApiModelsScriptsScriptQueryResponse
func (a *ExtensionsApiService) ExtensionsScriptsGetExecute(r ApiExtensionsScriptsGetRequest) ([]KeyfactorWebKeyfactorApiModelsScriptsScriptQueryResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []KeyfactorWebKeyfactorApiModelsScriptsScriptQueryResponse
	)

	apiBasePath := a.client.cfg.APIPath
	if apiBasePath == "" {
		apiBasePath = "/KeyfactorAPI"
	}

	localVarPath := apiBasePath + "/Extensions/Scripts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xKeyfactorRequestedWith == nil {
		return localVarReturnValue, nil, reportError("xKeyfactorRequestedWith is required and must be specified")
	}

	if r.queryString != nil {
		parameterAddToQuery(localVarQueryParams, "QueryString", r.queryString, "")
	}
	if r.pageReturned != nil {
		parameterAddToQuery(localVarQueryParams, "PageReturned", r.pageReturned, "")
	}
	if r.returnLimit != nil {
		parameterAddToQuery(localVarQueryParams, "ReturnLimit", r.returnLimit, "")
	}
	if r.sortField != nil {
		parameterAddToQuery(localVarQueryParams, "SortField", r.sortField, "")
	}
	if r.sortAscending != nil {
		parameterAddToQuery(localVarQueryParams, "SortAscending", r.sortAscending, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

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

type ApiExtensionsScriptsIdDeleteRequest struct {
	ctx                     context.Context
	ApiService              *ExtensionsApiService
	id                      int32
	xKeyfactorRequestedWith *string
	xKeyfactorApiVersion    *string
}

// Type of the request [XMLHttpRequest, APIClient]
func (r ApiExtensionsScriptsIdDeleteRequest) XKeyfactorRequestedWith(xKeyfactorRequestedWith string) ApiExtensionsScriptsIdDeleteRequest {
	r.xKeyfactorRequestedWith = &xKeyfactorRequestedWith
	return r
}

// Desired version of the api, if not provided defaults to v1
func (r ApiExtensionsScriptsIdDeleteRequest) XKeyfactorApiVersion(xKeyfactorApiVersion string) ApiExtensionsScriptsIdDeleteRequest {
	r.xKeyfactorApiVersion = &xKeyfactorApiVersion
	return r
}

func (r ApiExtensionsScriptsIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ExtensionsScriptsIdDeleteExecute(r)
}

/*
ExtensionsScriptsIdDelete Deletes a script. Script cannot be configured to an alert or workflow.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Id of the script to delete
	@return ApiExtensionsScriptsIdDeleteRequest
*/
func (a *ExtensionsApiService) ExtensionsScriptsIdDelete(ctx context.Context, id int32) ApiExtensionsScriptsIdDeleteRequest {

	xKeyfactorRequestedWith := "APIClient"

	return ApiExtensionsScriptsIdDeleteRequest{
		ApiService: a,
		ctx:        ctx,

		xKeyfactorRequestedWith: &xKeyfactorRequestedWith,
		id:                      id,
	}
}

// Execute executes the request
func (a *ExtensionsApiService) ExtensionsScriptsIdDeleteExecute(r ApiExtensionsScriptsIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	apiBasePath := a.client.cfg.APIPath
	if apiBasePath == "" {
		apiBasePath = "/KeyfactorAPI"
	}

	localVarPath := apiBasePath + "/Extensions/Scripts/{id}"
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

type ApiExtensionsScriptsIdGetRequest struct {
	ctx                     context.Context
	ApiService              *ExtensionsApiService
	id                      int32
	xKeyfactorRequestedWith *string
	xKeyfactorApiVersion    *string
}

// Type of the request [XMLHttpRequest, APIClient]
func (r ApiExtensionsScriptsIdGetRequest) XKeyfactorRequestedWith(xKeyfactorRequestedWith string) ApiExtensionsScriptsIdGetRequest {
	r.xKeyfactorRequestedWith = &xKeyfactorRequestedWith
	return r
}

// Desired version of the api, if not provided defaults to v1
func (r ApiExtensionsScriptsIdGetRequest) XKeyfactorApiVersion(xKeyfactorApiVersion string) ApiExtensionsScriptsIdGetRequest {
	r.xKeyfactorApiVersion = &xKeyfactorApiVersion
	return r
}

func (r ApiExtensionsScriptsIdGetRequest) Execute() (*KeyfactorWebKeyfactorApiModelsScriptsScriptResponse, *http.Response, error) {
	return r.ApiService.ExtensionsScriptsIdGetExecute(r)
}

/*
ExtensionsScriptsIdGet Returns a single script that matches the provided Id

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Id of the script
	@return ApiExtensionsScriptsIdGetRequest
*/
func (a *ExtensionsApiService) ExtensionsScriptsIdGet(ctx context.Context, id int32) ApiExtensionsScriptsIdGetRequest {

	xKeyfactorRequestedWith := "APIClient"

	return ApiExtensionsScriptsIdGetRequest{
		ApiService: a,
		ctx:        ctx,

		xKeyfactorRequestedWith: &xKeyfactorRequestedWith,
		id:                      id,
	}
}

// Execute executes the request
//
//	@return KeyfactorWebKeyfactorApiModelsScriptsScriptResponse
func (a *ExtensionsApiService) ExtensionsScriptsIdGetExecute(r ApiExtensionsScriptsIdGetRequest) (*KeyfactorWebKeyfactorApiModelsScriptsScriptResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *KeyfactorWebKeyfactorApiModelsScriptsScriptResponse
	)

	apiBasePath := a.client.cfg.APIPath
	if apiBasePath == "" {
		apiBasePath = "/KeyfactorAPI"
	}

	localVarPath := apiBasePath + "/Extensions/Scripts/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.id < 1 {
		return localVarReturnValue, nil, reportError("id must be greater than 1")
	}
	if r.id > 2147483647 {
		return localVarReturnValue, nil, reportError("id must be less than 2147483647")
	}
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
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

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

type ApiExtensionsScriptsPostRequest struct {
	ctx                                                      context.Context
	ApiService                                               *ExtensionsApiService
	xKeyfactorRequestedWith                                  *string
	keyfactorWebKeyfactorApiModelsScriptsScriptCreateRequest *KeyfactorWebKeyfactorApiModelsScriptsScriptCreateRequest
	xKeyfactorApiVersion                                     *string
}

// Type of the request [XMLHttpRequest, APIClient]
func (r ApiExtensionsScriptsPostRequest) XKeyfactorRequestedWith(xKeyfactorRequestedWith string) ApiExtensionsScriptsPostRequest {
	r.xKeyfactorRequestedWith = &xKeyfactorRequestedWith
	return r
}

// Script parameters
func (r ApiExtensionsScriptsPostRequest) KeyfactorWebKeyfactorApiModelsScriptsScriptCreateRequest(keyfactorWebKeyfactorApiModelsScriptsScriptCreateRequest KeyfactorWebKeyfactorApiModelsScriptsScriptCreateRequest) ApiExtensionsScriptsPostRequest {
	r.keyfactorWebKeyfactorApiModelsScriptsScriptCreateRequest = &keyfactorWebKeyfactorApiModelsScriptsScriptCreateRequest
	return r
}

// Desired version of the api, if not provided defaults to v1
func (r ApiExtensionsScriptsPostRequest) XKeyfactorApiVersion(xKeyfactorApiVersion string) ApiExtensionsScriptsPostRequest {
	r.xKeyfactorApiVersion = &xKeyfactorApiVersion
	return r
}

func (r ApiExtensionsScriptsPostRequest) Execute() (*KeyfactorWebKeyfactorApiModelsScriptsScriptResponse, *http.Response, error) {
	return r.ApiService.ExtensionsScriptsPostExecute(r)
}

/*
ExtensionsScriptsPost Adds a new script

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiExtensionsScriptsPostRequest
*/
func (a *ExtensionsApiService) ExtensionsScriptsPost(ctx context.Context) ApiExtensionsScriptsPostRequest {

	xKeyfactorRequestedWith := "APIClient"

	return ApiExtensionsScriptsPostRequest{
		ApiService: a,
		ctx:        ctx,

		xKeyfactorRequestedWith: &xKeyfactorRequestedWith,
	}
}

// Execute executes the request
//
//	@return KeyfactorWebKeyfactorApiModelsScriptsScriptResponse
func (a *ExtensionsApiService) ExtensionsScriptsPostExecute(r ApiExtensionsScriptsPostRequest) (*KeyfactorWebKeyfactorApiModelsScriptsScriptResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *KeyfactorWebKeyfactorApiModelsScriptsScriptResponse
	)

	apiBasePath := a.client.cfg.APIPath
	if apiBasePath == "" {
		apiBasePath = "/KeyfactorAPI"
	}

	localVarPath := apiBasePath + "/Extensions/Scripts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xKeyfactorRequestedWith == nil {
		return localVarReturnValue, nil, reportError("xKeyfactorRequestedWith is required and must be specified")
	}
	if r.keyfactorWebKeyfactorApiModelsScriptsScriptCreateRequest == nil {
		return localVarReturnValue, nil, reportError("keyfactorWebKeyfactorApiModelsScriptsScriptCreateRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json-patch+json", "application/json", "text/json", "application/*+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

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
	localVarPostBody = r.keyfactorWebKeyfactorApiModelsScriptsScriptCreateRequest
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

type ApiExtensionsScriptsPutRequest struct {
	ctx                                                       context.Context
	ApiService                                                *ExtensionsApiService
	xKeyfactorRequestedWith                                   *string
	keyfactorWebKeyfactorApiModelsScriptsScriptsUpdateRequest *KeyfactorWebKeyfactorApiModelsScriptsScriptsUpdateRequest
	xKeyfactorApiVersion                                      *string
}

// Type of the request [XMLHttpRequest, APIClient]
func (r ApiExtensionsScriptsPutRequest) XKeyfactorRequestedWith(xKeyfactorRequestedWith string) ApiExtensionsScriptsPutRequest {
	r.xKeyfactorRequestedWith = &xKeyfactorRequestedWith
	return r
}

// Script parameters
func (r ApiExtensionsScriptsPutRequest) KeyfactorWebKeyfactorApiModelsScriptsScriptsUpdateRequest(keyfactorWebKeyfactorApiModelsScriptsScriptsUpdateRequest KeyfactorWebKeyfactorApiModelsScriptsScriptsUpdateRequest) ApiExtensionsScriptsPutRequest {
	r.keyfactorWebKeyfactorApiModelsScriptsScriptsUpdateRequest = &keyfactorWebKeyfactorApiModelsScriptsScriptsUpdateRequest
	return r
}

// Desired version of the api, if not provided defaults to v1
func (r ApiExtensionsScriptsPutRequest) XKeyfactorApiVersion(xKeyfactorApiVersion string) ApiExtensionsScriptsPutRequest {
	r.xKeyfactorApiVersion = &xKeyfactorApiVersion
	return r
}

func (r ApiExtensionsScriptsPutRequest) Execute() (*KeyfactorWebKeyfactorApiModelsScriptsScriptResponse, *http.Response, error) {
	return r.ApiService.ExtensionsScriptsPutExecute(r)
}

/*
ExtensionsScriptsPut Updates a script

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiExtensionsScriptsPutRequest
*/
func (a *ExtensionsApiService) ExtensionsScriptsPut(ctx context.Context) ApiExtensionsScriptsPutRequest {

	xKeyfactorRequestedWith := "APIClient"

	return ApiExtensionsScriptsPutRequest{
		ApiService: a,
		ctx:        ctx,

		xKeyfactorRequestedWith: &xKeyfactorRequestedWith,
	}
}

// Execute executes the request
//
//	@return KeyfactorWebKeyfactorApiModelsScriptsScriptResponse
func (a *ExtensionsApiService) ExtensionsScriptsPutExecute(r ApiExtensionsScriptsPutRequest) (*KeyfactorWebKeyfactorApiModelsScriptsScriptResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *KeyfactorWebKeyfactorApiModelsScriptsScriptResponse
	)

	apiBasePath := a.client.cfg.APIPath
	if apiBasePath == "" {
		apiBasePath = "/KeyfactorAPI"
	}

	localVarPath := apiBasePath + "/Extensions/Scripts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xKeyfactorRequestedWith == nil {
		return localVarReturnValue, nil, reportError("xKeyfactorRequestedWith is required and must be specified")
	}
	if r.keyfactorWebKeyfactorApiModelsScriptsScriptsUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("keyfactorWebKeyfactorApiModelsScriptsScriptsUpdateRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json-patch+json", "application/json", "text/json", "application/*+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

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
	localVarPostBody = r.keyfactorWebKeyfactorApiModelsScriptsScriptsUpdateRequest
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
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

// LogonApiService LogonApi service
type LogonApiService service

type ApiSSHLogonsAccessPostRequest struct {
	ctx                               context.Context
	ApiService                        *LogonApiService
	xKeyfactorRequestedWith           *string
	xKeyfactorApiVersion              *string
	modelsSSHLogonsLogonAccessRequest *ModelsSSHLogonsLogonAccessRequest
}

// Type of the request [XMLHttpRequest, APIClient]
func (r ApiSSHLogonsAccessPostRequest) XKeyfactorRequestedWith(xKeyfactorRequestedWith string) ApiSSHLogonsAccessPostRequest {
	r.xKeyfactorRequestedWith = &xKeyfactorRequestedWith
	return r
}

// Desired version of the api, if not provided defaults to v1
func (r ApiSSHLogonsAccessPostRequest) XKeyfactorApiVersion(xKeyfactorApiVersion string) ApiSSHLogonsAccessPostRequest {
	r.xKeyfactorApiVersion = &xKeyfactorApiVersion
	return r
}

// Users to add the existing logon
func (r ApiSSHLogonsAccessPostRequest) ModelsSSHLogonsLogonAccessRequest(modelsSSHLogonsLogonAccessRequest ModelsSSHLogonsLogonAccessRequest) ApiSSHLogonsAccessPostRequest {
	r.modelsSSHLogonsLogonAccessRequest = &modelsSSHLogonsLogonAccessRequest
	return r
}

func (r ApiSSHLogonsAccessPostRequest) Execute() (*ModelsSSHAccessLogonUserAccessResponse, *http.Response, error) {
	return r.ApiService.SSHLogonsAccessPostExecute(r)
}

/*
SSHLogonsAccessPost Updates the users with access to an existing logon

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiSSHLogonsAccessPostRequest
*/
func (a *LogonApiService) SSHLogonsAccessPost(ctx context.Context) ApiSSHLogonsAccessPostRequest {

	xKeyfactorRequestedWith := "APIClient"

	return ApiSSHLogonsAccessPostRequest{
		ApiService: a,
		ctx:        ctx,

		xKeyfactorRequestedWith: &xKeyfactorRequestedWith,
	}
}

// Execute executes the request
//
//	@return ModelsSSHAccessLogonUserAccessResponse
func (a *LogonApiService) SSHLogonsAccessPostExecute(r ApiSSHLogonsAccessPostRequest) (*ModelsSSHAccessLogonUserAccessResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ModelsSSHAccessLogonUserAccessResponse
	)

	apiBasePath := a.client.cfg.APIPath
	if apiBasePath == "" {
		apiBasePath = "/KeyfactorAPI"
	}

	localVarPath := apiBasePath + "/SSH/Logons/Access"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xKeyfactorRequestedWith == nil {
		return localVarReturnValue, nil, reportError("xKeyfactorRequestedWith is required and must be specified")
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
	localVarPostBody = r.modelsSSHLogonsLogonAccessRequest
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

type ApiSSHLogonsGetRequest struct {
	ctx                     context.Context
	ApiService              *LogonApiService
	xKeyfactorRequestedWith *string
	queryString             *string
	pageReturned            *int32
	returnLimit             *int32
	sortField               *string
	sortAscending           *int32
	xKeyfactorApiVersion    *string
}

// Type of the request [XMLHttpRequest, APIClient]
func (r ApiSSHLogonsGetRequest) XKeyfactorRequestedWith(xKeyfactorRequestedWith string) ApiSSHLogonsGetRequest {
	r.xKeyfactorRequestedWith = &xKeyfactorRequestedWith
	return r
}

func (r ApiSSHLogonsGetRequest) QueryString(queryString string) ApiSSHLogonsGetRequest {
	r.queryString = &queryString
	return r
}

func (r ApiSSHLogonsGetRequest) PageReturned(pageReturned int32) ApiSSHLogonsGetRequest {
	r.pageReturned = &pageReturned
	return r
}

func (r ApiSSHLogonsGetRequest) ReturnLimit(returnLimit int32) ApiSSHLogonsGetRequest {
	r.returnLimit = &returnLimit
	return r
}

func (r ApiSSHLogonsGetRequest) SortField(sortField string) ApiSSHLogonsGetRequest {
	r.sortField = &sortField
	return r
}

func (r ApiSSHLogonsGetRequest) SortAscending(sortAscending int32) ApiSSHLogonsGetRequest {
	r.sortAscending = &sortAscending
	return r
}

// Desired version of the api, if not provided defaults to v1
func (r ApiSSHLogonsGetRequest) XKeyfactorApiVersion(xKeyfactorApiVersion string) ApiSSHLogonsGetRequest {
	r.xKeyfactorApiVersion = &xKeyfactorApiVersion
	return r
}

func (r ApiSSHLogonsGetRequest) Execute() ([]ModelsSSHLogonsLogonQueryResponse, *http.Response, error) {
	return r.ApiService.SSHLogonsGetExecute(r)
}

/*
SSHLogonsGet Returns all Logons according to the provided filter parameters

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiSSHLogonsGetRequest
*/
func (a *LogonApiService) SSHLogonsGet(ctx context.Context) ApiSSHLogonsGetRequest {

	xKeyfactorRequestedWith := "APIClient"

	return ApiSSHLogonsGetRequest{
		ApiService: a,
		ctx:        ctx,

		xKeyfactorRequestedWith: &xKeyfactorRequestedWith,
	}
}

// Execute executes the request
//
//	@return []ModelsSSHLogonsLogonQueryResponse
func (a *LogonApiService) SSHLogonsGetExecute(r ApiSSHLogonsGetRequest) ([]ModelsSSHLogonsLogonQueryResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []ModelsSSHLogonsLogonQueryResponse
	)

	apiBasePath := a.client.cfg.APIPath
	if apiBasePath == "" {
		apiBasePath = "/KeyfactorAPI"
	}

	localVarPath := apiBasePath + "/SSH/Logons"

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

type ApiSSHLogonsIdDeleteRequest struct {
	ctx                     context.Context
	ApiService              *LogonApiService
	id                      int32
	xKeyfactorRequestedWith *string
	xKeyfactorApiVersion    *string
}

// Type of the request [XMLHttpRequest, APIClient]
func (r ApiSSHLogonsIdDeleteRequest) XKeyfactorRequestedWith(xKeyfactorRequestedWith string) ApiSSHLogonsIdDeleteRequest {
	r.xKeyfactorRequestedWith = &xKeyfactorRequestedWith
	return r
}

// Desired version of the api, if not provided defaults to v1
func (r ApiSSHLogonsIdDeleteRequest) XKeyfactorApiVersion(xKeyfactorApiVersion string) ApiSSHLogonsIdDeleteRequest {
	r.xKeyfactorApiVersion = &xKeyfactorApiVersion
	return r
}

func (r ApiSSHLogonsIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.SSHLogonsIdDeleteExecute(r)
}

/*
SSHLogonsIdDelete Deletes a Logon associated with the provided identifier

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Keyfactor identifer of the Logon to be deleted
	@return ApiSSHLogonsIdDeleteRequest
*/
func (a *LogonApiService) SSHLogonsIdDelete(ctx context.Context, id int32) ApiSSHLogonsIdDeleteRequest {

	xKeyfactorRequestedWith := "APIClient"

	return ApiSSHLogonsIdDeleteRequest{
		ApiService: a,
		ctx:        ctx,

		xKeyfactorRequestedWith: &xKeyfactorRequestedWith,
		id:                      id,
	}
}

// Execute executes the request
func (a *LogonApiService) SSHLogonsIdDeleteExecute(r ApiSSHLogonsIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	apiBasePath := a.client.cfg.APIPath
	if apiBasePath == "" {
		apiBasePath = "/KeyfactorAPI"
	}

	localVarPath := apiBasePath + "/SSH/Logons/{id}"
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

type ApiSSHLogonsIdGetRequest struct {
	ctx                     context.Context
	ApiService              *LogonApiService
	id                      int32
	xKeyfactorRequestedWith *string
	xKeyfactorApiVersion    *string
}

// Type of the request [XMLHttpRequest, APIClient]
func (r ApiSSHLogonsIdGetRequest) XKeyfactorRequestedWith(xKeyfactorRequestedWith string) ApiSSHLogonsIdGetRequest {
	r.xKeyfactorRequestedWith = &xKeyfactorRequestedWith
	return r
}

// Desired version of the api, if not provided defaults to v1
func (r ApiSSHLogonsIdGetRequest) XKeyfactorApiVersion(xKeyfactorApiVersion string) ApiSSHLogonsIdGetRequest {
	r.xKeyfactorApiVersion = &xKeyfactorApiVersion
	return r
}

func (r ApiSSHLogonsIdGetRequest) Execute() (*ModelsSSHLogonsLogonResponse, *http.Response, error) {
	return r.ApiService.SSHLogonsIdGetExecute(r)
}

/*
SSHLogonsIdGet Fetches a Logon associated with the provided identifier

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Keyfactor identifer of the Logon to be Fetched
	@return ApiSSHLogonsIdGetRequest
*/
func (a *LogonApiService) SSHLogonsIdGet(ctx context.Context, id int32) ApiSSHLogonsIdGetRequest {

	xKeyfactorRequestedWith := "APIClient"

	return ApiSSHLogonsIdGetRequest{
		ApiService: a,
		ctx:        ctx,

		xKeyfactorRequestedWith: &xKeyfactorRequestedWith,
		id:                      id,
	}
}

// Execute executes the request
//
//	@return ModelsSSHLogonsLogonResponse
func (a *LogonApiService) SSHLogonsIdGetExecute(r ApiSSHLogonsIdGetRequest) (*ModelsSSHLogonsLogonResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ModelsSSHLogonsLogonResponse
	)

	apiBasePath := a.client.cfg.APIPath
	if apiBasePath == "" {
		apiBasePath = "/KeyfactorAPI"
	}

	localVarPath := apiBasePath + "/SSH/Logons/{id}"
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

type ApiSSHLogonsPostRequest struct {
	ctx                                 context.Context
	ApiService                          *LogonApiService
	xKeyfactorRequestedWith             *string
	xKeyfactorApiVersion                *string
	modelsSSHLogonsLogonCreationRequest *ModelsSSHLogonsLogonCreationRequest
}

// Type of the request [XMLHttpRequest, APIClient]
func (r ApiSSHLogonsPostRequest) XKeyfactorRequestedWith(xKeyfactorRequestedWith string) ApiSSHLogonsPostRequest {
	r.xKeyfactorRequestedWith = &xKeyfactorRequestedWith
	return r
}

// Desired version of the api, if not provided defaults to v1
func (r ApiSSHLogonsPostRequest) XKeyfactorApiVersion(xKeyfactorApiVersion string) ApiSSHLogonsPostRequest {
	r.xKeyfactorApiVersion = &xKeyfactorApiVersion
	return r
}

// Logon properties to be applied to the new logon
func (r ApiSSHLogonsPostRequest) ModelsSSHLogonsLogonCreationRequest(modelsSSHLogonsLogonCreationRequest ModelsSSHLogonsLogonCreationRequest) ApiSSHLogonsPostRequest {
	r.modelsSSHLogonsLogonCreationRequest = &modelsSSHLogonsLogonCreationRequest
	return r
}

func (r ApiSSHLogonsPostRequest) Execute() (*ModelsSSHLogonsLogonResponse, *http.Response, error) {
	return r.ApiService.SSHLogonsPostExecute(r)
}

/*
SSHLogonsPost Creates a logon with the provided properties

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiSSHLogonsPostRequest
*/
func (a *LogonApiService) SSHLogonsPost(ctx context.Context) ApiSSHLogonsPostRequest {

	xKeyfactorRequestedWith := "APIClient"

	return ApiSSHLogonsPostRequest{
		ApiService: a,
		ctx:        ctx,

		xKeyfactorRequestedWith: &xKeyfactorRequestedWith,
	}
}

// Execute executes the request
//
//	@return ModelsSSHLogonsLogonResponse
func (a *LogonApiService) SSHLogonsPostExecute(r ApiSSHLogonsPostRequest) (*ModelsSSHLogonsLogonResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ModelsSSHLogonsLogonResponse
	)

	apiBasePath := a.client.cfg.APIPath
	if apiBasePath == "" {
		apiBasePath = "/KeyfactorAPI"
	}

	localVarPath := apiBasePath + "/SSH/Logons"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xKeyfactorRequestedWith == nil {
		return localVarReturnValue, nil, reportError("xKeyfactorRequestedWith is required and must be specified")
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
	localVarPostBody = r.modelsSSHLogonsLogonCreationRequest
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
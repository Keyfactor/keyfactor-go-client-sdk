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

Keyfactor-v1

This reference serves to document REST-based methods to manage and integrate with Keyfactor. In addition, an embedded interface allows for the execution of calls against the current Keyfactor API instance.

API version: v1
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

// SecurityRolesApiService SecurityRolesApi service
type SecurityRolesApiService service

type ApiSecurityRolesDeleteSecurityRoleRequest struct {
	ctx                     context.Context
	ApiService              *SecurityRolesApiService
	id                      int32
	xKeyfactorRequestedWith *string
	xKeyfactorApiVersion    *string
}

// Type of the request [XMLHttpRequest, APIClient]
func (r ApiSecurityRolesDeleteSecurityRoleRequest) XKeyfactorRequestedWith(xKeyfactorRequestedWith string) ApiSecurityRolesDeleteSecurityRoleRequest {
	r.xKeyfactorRequestedWith = &xKeyfactorRequestedWith
	return r
}

// Desired version of the api, if not provided defaults to v1
func (r ApiSecurityRolesDeleteSecurityRoleRequest) XKeyfactorApiVersion(xKeyfactorApiVersion string) ApiSecurityRolesDeleteSecurityRoleRequest {
	r.xKeyfactorApiVersion = &xKeyfactorApiVersion
	return r
}

func (r ApiSecurityRolesDeleteSecurityRoleRequest) Execute() (*http.Response, error) {
	return r.ApiService.SecurityRolesDeleteSecurityRoleExecute(r)
}

/*
SecurityRolesDeleteSecurityRole Deletes the security role whose ID is provided.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Security role identifier
	@return ApiSecurityRolesDeleteSecurityRoleRequest
*/
func (a *SecurityRolesApiService) SecurityRolesDeleteSecurityRole(ctx context.Context, id int32) ApiSecurityRolesDeleteSecurityRoleRequest {
	xKeyfactorApiVersion := "1"
	xKeyfactorRequestedWith := "APIClient"

	return ApiSecurityRolesDeleteSecurityRoleRequest{
		ApiService:              a,
		ctx:                     ctx,
		xKeyfactorApiVersion:    &xKeyfactorApiVersion,
		xKeyfactorRequestedWith: &xKeyfactorRequestedWith,
		id:                      id,
	}
}

// Execute executes the request
func (a *SecurityRolesApiService) SecurityRolesDeleteSecurityRoleExecute(r ApiSecurityRolesDeleteSecurityRoleRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	apiBasePath := a.client.cfg.APIPath
	if apiBasePath == "" {
		apiBasePath = "/KeyfactorAPI"
	}

	localVarPath := apiBasePath + "/Security/Roles/{id}"
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

type ApiSecurityRolesGetIdentitiesWithRoleRequest struct {
	ctx                     context.Context
	ApiService              *SecurityRolesApiService
	id                      int32
	xKeyfactorRequestedWith *string
	xKeyfactorApiVersion    *string
}

// Type of the request [XMLHttpRequest, APIClient]
func (r ApiSecurityRolesGetIdentitiesWithRoleRequest) XKeyfactorRequestedWith(xKeyfactorRequestedWith string) ApiSecurityRolesGetIdentitiesWithRoleRequest {
	r.xKeyfactorRequestedWith = &xKeyfactorRequestedWith
	return r
}

// Desired version of the api, if not provided defaults to v1
func (r ApiSecurityRolesGetIdentitiesWithRoleRequest) XKeyfactorApiVersion(xKeyfactorApiVersion string) ApiSecurityRolesGetIdentitiesWithRoleRequest {
	r.xKeyfactorApiVersion = &xKeyfactorApiVersion
	return r
}

func (r ApiSecurityRolesGetIdentitiesWithRoleRequest) Execute() ([]KeyfactorApiModelsSecurityRolesRoleIdentitiesResponse, *http.Response, error) {
	return r.ApiService.SecurityRolesGetIdentitiesWithRoleExecute(r)
}

/*
SecurityRolesGetIdentitiesWithRole Returns all identities which have the security role that matches the id.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Security role identifier
	@return ApiSecurityRolesGetIdentitiesWithRoleRequest
*/
func (a *SecurityRolesApiService) SecurityRolesGetIdentitiesWithRole(ctx context.Context, id int32) ApiSecurityRolesGetIdentitiesWithRoleRequest {
	xKeyfactorApiVersion := "1"
	xKeyfactorRequestedWith := "APIClient"

	return ApiSecurityRolesGetIdentitiesWithRoleRequest{
		ApiService:              a,
		ctx:                     ctx,
		xKeyfactorApiVersion:    &xKeyfactorApiVersion,
		xKeyfactorRequestedWith: &xKeyfactorRequestedWith,
		id:                      id,
	}
}

// Execute executes the request
//
//	@return []KeyfactorApiModelsSecurityRolesRoleIdentitiesResponse
func (a *SecurityRolesApiService) SecurityRolesGetIdentitiesWithRoleExecute(r ApiSecurityRolesGetIdentitiesWithRoleRequest) ([]KeyfactorApiModelsSecurityRolesRoleIdentitiesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []KeyfactorApiModelsSecurityRolesRoleIdentitiesResponse
	)

	apiBasePath := a.client.cfg.APIPath
	if apiBasePath == "" {
		apiBasePath = "/KeyfactorAPI"
	}

	localVarPath := apiBasePath + "/Security/Roles/{id}/Identities"
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

type ApiSecurityRolesGetSecurityRoleRequest struct {
	ctx                     context.Context
	ApiService              *SecurityRolesApiService
	id                      int32
	xKeyfactorRequestedWith *string
	xKeyfactorApiVersion    *string
}

// Type of the request [XMLHttpRequest, APIClient]
func (r ApiSecurityRolesGetSecurityRoleRequest) XKeyfactorRequestedWith(xKeyfactorRequestedWith string) ApiSecurityRolesGetSecurityRoleRequest {
	r.xKeyfactorRequestedWith = &xKeyfactorRequestedWith
	return r
}

// Desired version of the api, if not provided defaults to v1
func (r ApiSecurityRolesGetSecurityRoleRequest) XKeyfactorApiVersion(xKeyfactorApiVersion string) ApiSecurityRolesGetSecurityRoleRequest {
	r.xKeyfactorApiVersion = &xKeyfactorApiVersion
	return r
}

func (r ApiSecurityRolesGetSecurityRoleRequest) Execute() (*ModelsSecuritySecurityRolesSecurityRoleResponse, *http.Response, error) {
	return r.ApiService.SecurityRolesGetSecurityRoleExecute(r)
}

/*
SecurityRolesGetSecurityRole Returns a single security role that matches the id.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Security role identifier
	@return ApiSecurityRolesGetSecurityRoleRequest
*/
func (a *SecurityRolesApiService) SecurityRolesGetSecurityRole(ctx context.Context, id int32) ApiSecurityRolesGetSecurityRoleRequest {
	xKeyfactorApiVersion := "1"
	xKeyfactorRequestedWith := "APIClient"

	return ApiSecurityRolesGetSecurityRoleRequest{
		ApiService:              a,
		ctx:                     ctx,
		xKeyfactorApiVersion:    &xKeyfactorApiVersion,
		xKeyfactorRequestedWith: &xKeyfactorRequestedWith,
		id:                      id,
	}
}

// Execute executes the request
//
//	@return ModelsSecuritySecurityRolesSecurityRoleResponse
func (a *SecurityRolesApiService) SecurityRolesGetSecurityRoleExecute(r ApiSecurityRolesGetSecurityRoleRequest) (*ModelsSecuritySecurityRolesSecurityRoleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ModelsSecuritySecurityRolesSecurityRoleResponse
	)

	apiBasePath := a.client.cfg.APIPath
	if apiBasePath == "" {
		apiBasePath = "/KeyfactorAPI"
	}

	localVarPath := apiBasePath + "/Security/Roles/{id}"
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

type ApiSecurityRolesUpdateIdentitiesWithRoleRequest struct {
	ctx                     context.Context
	ApiService              *SecurityRolesApiService
	id                      int32
	xKeyfactorRequestedWith *string
	identities              *KeyfactorApiModelsSecurityRolesRoleIdentitiesRequest
	xKeyfactorApiVersion    *string
}

// Type of the request [XMLHttpRequest, APIClient]
func (r ApiSecurityRolesUpdateIdentitiesWithRoleRequest) XKeyfactorRequestedWith(xKeyfactorRequestedWith string) ApiSecurityRolesUpdateIdentitiesWithRoleRequest {
	r.xKeyfactorRequestedWith = &xKeyfactorRequestedWith
	return r
}

// Lists of Identity IDs to remove or add to the role
func (r ApiSecurityRolesUpdateIdentitiesWithRoleRequest) Identities(identities KeyfactorApiModelsSecurityRolesRoleIdentitiesRequest) ApiSecurityRolesUpdateIdentitiesWithRoleRequest {
	r.identities = &identities
	return r
}

// Desired version of the api, if not provided defaults to v1
func (r ApiSecurityRolesUpdateIdentitiesWithRoleRequest) XKeyfactorApiVersion(xKeyfactorApiVersion string) ApiSecurityRolesUpdateIdentitiesWithRoleRequest {
	r.xKeyfactorApiVersion = &xKeyfactorApiVersion
	return r
}

func (r ApiSecurityRolesUpdateIdentitiesWithRoleRequest) Execute() ([]KeyfactorApiModelsSecurityRolesRoleIdentitiesResponse, *http.Response, error) {
	return r.ApiService.SecurityRolesUpdateIdentitiesWithRoleExecute(r)
}

/*
SecurityRolesUpdateIdentitiesWithRole Updates the identities which have the security role that matches the id.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Security role identifier
	@return ApiSecurityRolesUpdateIdentitiesWithRoleRequest
*/
func (a *SecurityRolesApiService) SecurityRolesUpdateIdentitiesWithRole(ctx context.Context, id int32) ApiSecurityRolesUpdateIdentitiesWithRoleRequest {
	xKeyfactorApiVersion := "1"
	xKeyfactorRequestedWith := "APIClient"

	return ApiSecurityRolesUpdateIdentitiesWithRoleRequest{
		ApiService:              a,
		ctx:                     ctx,
		xKeyfactorApiVersion:    &xKeyfactorApiVersion,
		xKeyfactorRequestedWith: &xKeyfactorRequestedWith,
		id:                      id,
	}
}

// Execute executes the request
//
//	@return []KeyfactorApiModelsSecurityRolesRoleIdentitiesResponse
func (a *SecurityRolesApiService) SecurityRolesUpdateIdentitiesWithRoleExecute(r ApiSecurityRolesUpdateIdentitiesWithRoleRequest) ([]KeyfactorApiModelsSecurityRolesRoleIdentitiesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []KeyfactorApiModelsSecurityRolesRoleIdentitiesResponse
	)

	apiBasePath := a.client.cfg.APIPath
	if apiBasePath == "" {
		apiBasePath = "/KeyfactorAPI"
	}

	localVarPath := apiBasePath + "/Security/Roles/{id}/Identities"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xKeyfactorRequestedWith == nil {
		return localVarReturnValue, nil, reportError("xKeyfactorRequestedWith is required and must be specified")
	}
	if r.identities == nil {
		return localVarReturnValue, nil, reportError("identities is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "text/json", "application/xml", "text/xml", "application/x-www-form-urlencoded"}

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
	// body params
	localVarPostBody = r.identities
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
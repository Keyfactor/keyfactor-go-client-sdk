/*
Copyright 2025 Keyfactor
Licensed under the Apache License, Version 2.0 (the "License"); you may
not use this file except in compliance with the License.  You may obtain a
copy of the License at http://www.apache.org/licenses/LICENSE-2.0.  Unless
required by applicable law or agreed to in writing, software distributed
under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES
OR CONDITIONS OF ANY KIND, either express or implied. See the License for
thespecific language governing permissions and limitations under the
License.

Keyfactor Command Version: 25+

API version: 1
*/

package v1

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// ApplicationApiService Application API service
type ApplicationApiService service

// ---------------------------------------------------------------------------
// GET /Applications (list)
// ---------------------------------------------------------------------------

// ApiListApplicationsRequest for GET /Applications
type ApiListApplicationsRequest struct {
	ctx                     context.Context
	ApiService              *ApplicationApiService
	xKeyfactorRequestedWith *string
	xKeyfactorApiVersion    *string
}

func (r ApiListApplicationsRequest) XKeyfactorRequestedWith(v string) ApiListApplicationsRequest {
	r.xKeyfactorRequestedWith = &v
	return r
}

func (r ApiListApplicationsRequest) XKeyfactorApiVersion(v string) ApiListApplicationsRequest {
	r.xKeyfactorApiVersion = &v
	return r
}

func (r ApiListApplicationsRequest) Execute() ([]ApplicationsApplicationListResponse, *http.Response, error) {
	return r.ApiService.ListApplicationsExecute(r)
}

// NewListApplicationsRequest creates a new GET /Applications request.
func (a *ApplicationApiService) NewListApplicationsRequest(ctx context.Context) ApiListApplicationsRequest {
	requestedWith := "APIClient"
	apiVersion := "1"
	return ApiListApplicationsRequest{
		ApiService:              a,
		ctx:                     ctx,
		xKeyfactorRequestedWith: &requestedWith,
		xKeyfactorApiVersion:    &apiVersion,
	}
}

// ListApplicationsExecute executes the GET /Applications request.
func (a *ApplicationApiService) ListApplicationsExecute(r ApiListApplicationsRequest) ([]ApplicationsApplicationListResponse, *http.Response, error) {
	apiBasePath := a.client.AuthClient.GetServerConfig().APIPath
	if apiBasePath == "" {
		apiBasePath = "/KeyfactorAPI"
	}

	localVarPath := apiBasePath + "/Applications"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.xKeyfactorApiVersion != nil {
		parameterAddToQuery(localVarHeaderParams, "x-keyfactor-api-version", r.xKeyfactorApiVersion, "")
	}
	parameterAddToQuery(localVarHeaderParams, "x-keyfactor-requested-with", r.xKeyfactorRequestedWith, "")

	localVarHTTPHeaderAccept := selectHeaderAccept([]string{"text/plain", "application/json", "text/json"})
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	req, err := a.client.prepareRequest(r.ctx, localVarPath, http.MethodGet, nil, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return nil, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return nil, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return nil, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return nil, localVarHTTPResponse, &GenericOpenAPIError{body: localVarBody, error: localVarHTTPResponse.Status}
	}

	var localVarReturnValue []ApplicationsApplicationListResponse
	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		return nil, localVarHTTPResponse, &GenericOpenAPIError{body: localVarBody, error: err.Error()}
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}

// ---------------------------------------------------------------------------
// GET /Applications/{id}
// ---------------------------------------------------------------------------

// ApiGetApplicationByIdRequest for GET /Applications/{id}
type ApiGetApplicationByIdRequest struct {
	ctx                     context.Context
	ApiService              *ApplicationApiService
	id                      int32
	xKeyfactorRequestedWith *string
	xKeyfactorApiVersion    *string
}

func (r ApiGetApplicationByIdRequest) XKeyfactorRequestedWith(v string) ApiGetApplicationByIdRequest {
	r.xKeyfactorRequestedWith = &v
	return r
}

func (r ApiGetApplicationByIdRequest) XKeyfactorApiVersion(v string) ApiGetApplicationByIdRequest {
	r.xKeyfactorApiVersion = &v
	return r
}

func (r ApiGetApplicationByIdRequest) Execute() (*ApplicationsApplicationDetailResponse, *http.Response, error) {
	return r.ApiService.GetApplicationByIdExecute(r)
}

// NewGetApplicationByIdRequest creates a new GET /Applications/{id} request.
func (a *ApplicationApiService) NewGetApplicationByIdRequest(ctx context.Context, id int32) ApiGetApplicationByIdRequest {
	requestedWith := "APIClient"
	apiVersion := "1"
	return ApiGetApplicationByIdRequest{
		ApiService:              a,
		ctx:                     ctx,
		id:                      id,
		xKeyfactorRequestedWith: &requestedWith,
		xKeyfactorApiVersion:    &apiVersion,
	}
}

// GetApplicationByIdExecute executes the GET /Applications/{id} request.
func (a *ApplicationApiService) GetApplicationByIdExecute(r ApiGetApplicationByIdRequest) (*ApplicationsApplicationDetailResponse, *http.Response, error) {
	apiBasePath := a.client.AuthClient.GetServerConfig().APIPath
	if apiBasePath == "" {
		apiBasePath = "/KeyfactorAPI"
	}

	localVarPath := apiBasePath + fmt.Sprintf("/Applications/%d", r.id)
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.xKeyfactorApiVersion != nil {
		parameterAddToQuery(localVarHeaderParams, "x-keyfactor-api-version", r.xKeyfactorApiVersion, "")
	}
	parameterAddToQuery(localVarHeaderParams, "x-keyfactor-requested-with", r.xKeyfactorRequestedWith, "")

	localVarHTTPHeaderAccept := selectHeaderAccept([]string{"text/plain", "application/json", "text/json"})
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	req, err := a.client.prepareRequest(r.ctx, localVarPath, http.MethodGet, nil, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return nil, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return nil, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return nil, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return nil, localVarHTTPResponse, &GenericOpenAPIError{body: localVarBody, error: localVarHTTPResponse.Status}
	}

	var localVarReturnValue ApplicationsApplicationDetailResponse
	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		return nil, localVarHTTPResponse, &GenericOpenAPIError{body: localVarBody, error: err.Error()}
	}
	return &localVarReturnValue, localVarHTTPResponse, nil
}

// ---------------------------------------------------------------------------
// POST /Applications
// ---------------------------------------------------------------------------

// ApiCreateApplicationRequest for POST /Applications
type ApiCreateApplicationRequest struct {
	ctx                     context.Context
	ApiService              *ApplicationApiService
	xKeyfactorRequestedWith *string
	xKeyfactorApiVersion    *string
	applicationsRequest     *ApplicationsApplicationRequest
}

func (r ApiCreateApplicationRequest) XKeyfactorRequestedWith(v string) ApiCreateApplicationRequest {
	r.xKeyfactorRequestedWith = &v
	return r
}

func (r ApiCreateApplicationRequest) XKeyfactorApiVersion(v string) ApiCreateApplicationRequest {
	r.xKeyfactorApiVersion = &v
	return r
}

func (r ApiCreateApplicationRequest) ApplicationsApplicationRequest(v ApplicationsApplicationRequest) ApiCreateApplicationRequest {
	r.applicationsRequest = &v
	return r
}

func (r ApiCreateApplicationRequest) Execute() (*ApplicationsApplicationDetailResponse, *http.Response, error) {
	return r.ApiService.CreateApplicationExecute(r)
}

// NewCreateApplicationRequest creates a new POST /Applications request.
func (a *ApplicationApiService) NewCreateApplicationRequest(ctx context.Context) ApiCreateApplicationRequest {
	requestedWith := "APIClient"
	apiVersion := "1"
	return ApiCreateApplicationRequest{
		ApiService:              a,
		ctx:                     ctx,
		xKeyfactorRequestedWith: &requestedWith,
		xKeyfactorApiVersion:    &apiVersion,
	}
}

// CreateApplicationExecute executes the POST /Applications request.
func (a *ApplicationApiService) CreateApplicationExecute(r ApiCreateApplicationRequest) (*ApplicationsApplicationDetailResponse, *http.Response, error) {
	apiBasePath := a.client.AuthClient.GetServerConfig().APIPath
	if apiBasePath == "" {
		apiBasePath = "/KeyfactorAPI"
	}

	localVarPath := apiBasePath + "/Applications"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.xKeyfactorApiVersion != nil {
		parameterAddToQuery(localVarHeaderParams, "x-keyfactor-api-version", r.xKeyfactorApiVersion, "")
	}
	parameterAddToQuery(localVarHeaderParams, "x-keyfactor-requested-with", r.xKeyfactorRequestedWith, "")

	localVarHTTPContentType := selectHeaderContentType([]string{"application/json-patch+json", "application/json", "text/json", "application/*+json"})
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}
	localVarHTTPHeaderAccept := selectHeaderAccept([]string{"text/plain", "application/json", "text/json"})
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	req, err := a.client.prepareRequest(r.ctx, localVarPath, http.MethodPost, r.applicationsRequest, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return nil, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return nil, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return nil, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return nil, localVarHTTPResponse, &GenericOpenAPIError{body: localVarBody, error: localVarHTTPResponse.Status}
	}

	var localVarReturnValue ApplicationsApplicationDetailResponse
	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		return nil, localVarHTTPResponse, &GenericOpenAPIError{body: localVarBody, error: err.Error()}
	}
	return &localVarReturnValue, localVarHTTPResponse, nil
}

// ---------------------------------------------------------------------------
// PUT /Applications/{id}
// ---------------------------------------------------------------------------

// ApiUpdateApplicationRequest for PUT /Applications/{id}
type ApiUpdateApplicationRequest struct {
	ctx                     context.Context
	ApiService              *ApplicationApiService
	id                      int32
	xKeyfactorRequestedWith *string
	xKeyfactorApiVersion    *string
	applicationsRequest     *ApplicationsApplicationRequest
}

func (r ApiUpdateApplicationRequest) XKeyfactorRequestedWith(v string) ApiUpdateApplicationRequest {
	r.xKeyfactorRequestedWith = &v
	return r
}

func (r ApiUpdateApplicationRequest) XKeyfactorApiVersion(v string) ApiUpdateApplicationRequest {
	r.xKeyfactorApiVersion = &v
	return r
}

func (r ApiUpdateApplicationRequest) ApplicationsApplicationRequest(v ApplicationsApplicationRequest) ApiUpdateApplicationRequest {
	r.applicationsRequest = &v
	return r
}

func (r ApiUpdateApplicationRequest) Execute() (*ApplicationsApplicationDetailResponse, *http.Response, error) {
	return r.ApiService.UpdateApplicationExecute(r)
}

// NewUpdateApplicationRequest creates a new PUT /Applications/{id} request.
func (a *ApplicationApiService) NewUpdateApplicationRequest(ctx context.Context, id int32) ApiUpdateApplicationRequest {
	requestedWith := "APIClient"
	apiVersion := "1"
	return ApiUpdateApplicationRequest{
		ApiService:              a,
		ctx:                     ctx,
		id:                      id,
		xKeyfactorRequestedWith: &requestedWith,
		xKeyfactorApiVersion:    &apiVersion,
	}
}

// UpdateApplicationExecute executes the PUT /Applications/{id} request.
func (a *ApplicationApiService) UpdateApplicationExecute(r ApiUpdateApplicationRequest) (*ApplicationsApplicationDetailResponse, *http.Response, error) {
	apiBasePath := a.client.AuthClient.GetServerConfig().APIPath
	if apiBasePath == "" {
		apiBasePath = "/KeyfactorAPI"
	}

	localVarPath := apiBasePath + fmt.Sprintf("/Applications/%d", r.id)
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.xKeyfactorApiVersion != nil {
		parameterAddToQuery(localVarHeaderParams, "x-keyfactor-api-version", r.xKeyfactorApiVersion, "")
	}
	parameterAddToQuery(localVarHeaderParams, "x-keyfactor-requested-with", r.xKeyfactorRequestedWith, "")

	localVarHTTPContentType := selectHeaderContentType([]string{"application/json-patch+json", "application/json", "text/json", "application/*+json"})
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}
	localVarHTTPHeaderAccept := selectHeaderAccept([]string{"text/plain", "application/json", "text/json"})
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	req, err := a.client.prepareRequest(r.ctx, localVarPath, http.MethodPut, r.applicationsRequest, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return nil, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return nil, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return nil, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return nil, localVarHTTPResponse, &GenericOpenAPIError{body: localVarBody, error: localVarHTTPResponse.Status}
	}

	var localVarReturnValue ApplicationsApplicationDetailResponse
	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		return nil, localVarHTTPResponse, &GenericOpenAPIError{body: localVarBody, error: err.Error()}
	}
	return &localVarReturnValue, localVarHTTPResponse, nil
}

// ---------------------------------------------------------------------------
// DELETE /Applications/{id}
// ---------------------------------------------------------------------------

// ApiDeleteApplicationRequest for DELETE /Applications/{id}
type ApiDeleteApplicationRequest struct {
	ctx                     context.Context
	ApiService              *ApplicationApiService
	id                      int32
	xKeyfactorRequestedWith *string
	xKeyfactorApiVersion    *string
}

func (r ApiDeleteApplicationRequest) XKeyfactorRequestedWith(v string) ApiDeleteApplicationRequest {
	r.xKeyfactorRequestedWith = &v
	return r
}

func (r ApiDeleteApplicationRequest) XKeyfactorApiVersion(v string) ApiDeleteApplicationRequest {
	r.xKeyfactorApiVersion = &v
	return r
}

func (r ApiDeleteApplicationRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteApplicationExecute(r)
}

// NewDeleteApplicationRequest creates a new DELETE /Applications/{id} request.
func (a *ApplicationApiService) NewDeleteApplicationRequest(ctx context.Context, id int32) ApiDeleteApplicationRequest {
	requestedWith := "APIClient"
	apiVersion := "1"
	return ApiDeleteApplicationRequest{
		ApiService:              a,
		ctx:                     ctx,
		id:                      id,
		xKeyfactorRequestedWith: &requestedWith,
		xKeyfactorApiVersion:    &apiVersion,
	}
}

// DeleteApplicationExecute executes the DELETE /Applications/{id} request.
func (a *ApplicationApiService) DeleteApplicationExecute(r ApiDeleteApplicationRequest) (*http.Response, error) {
	apiBasePath := a.client.AuthClient.GetServerConfig().APIPath
	if apiBasePath == "" {
		apiBasePath = "/KeyfactorAPI"
	}

	localVarPath := apiBasePath + fmt.Sprintf("/Applications/%d", r.id)
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.xKeyfactorApiVersion != nil {
		parameterAddToQuery(localVarHeaderParams, "x-keyfactor-api-version", r.xKeyfactorApiVersion, "")
	}
	parameterAddToQuery(localVarHeaderParams, "x-keyfactor-requested-with", r.xKeyfactorRequestedWith, "")

	req, err := a.client.prepareRequest(r.ctx, localVarPath, http.MethodDelete, nil, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
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
		return localVarHTTPResponse, &GenericOpenAPIError{body: localVarBody, error: localVarHTTPResponse.Status}
	}

	return localVarHTTPResponse, nil
}

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


// IssuedAlertApiService IssuedAlertApi service
type IssuedAlertApiService service

type ApiIssuedAlertAddIssuedAlertRequest struct {
	ctx context.Context
	ApiService *IssuedAlertApiService
	xKeyfactorRequestedWith *string
	req *KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest
	xKeyfactorApiVersion *string
}

// Type of the request [XMLHttpRequest, APIClient]
func (r ApiIssuedAlertAddIssuedAlertRequest) XKeyfactorRequestedWith(xKeyfactorRequestedWith string) ApiIssuedAlertAddIssuedAlertRequest {
	r.xKeyfactorRequestedWith = &xKeyfactorRequestedWith
	return r
}

// Information for the new alert
func (r ApiIssuedAlertAddIssuedAlertRequest) Req(req KeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) ApiIssuedAlertAddIssuedAlertRequest {
	r.req = &req
	return r
}

// Desired version of the api, if not provided defaults to v1
func (r ApiIssuedAlertAddIssuedAlertRequest) XKeyfactorApiVersion(xKeyfactorApiVersion string) ApiIssuedAlertAddIssuedAlertRequest {
	r.xKeyfactorApiVersion = &xKeyfactorApiVersion
	return r
}

func (r ApiIssuedAlertAddIssuedAlertRequest) Execute() (*KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse, *http.Response, error) {
	return r.ApiService.IssuedAlertAddIssuedAlertExecute(r)
}

/*
IssuedAlertAddIssuedAlert Add a issued alert

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiIssuedAlertAddIssuedAlertRequest
*/
func (a *IssuedAlertApiService) IssuedAlertAddIssuedAlert(ctx context.Context) ApiIssuedAlertAddIssuedAlertRequest {
    requestedWith := "APIClient"
    version := "1"
	return ApiIssuedAlertAddIssuedAlertRequest{
		ApiService: a,
		ctx: ctx,
		xKeyfactorRequestedWith: &requestedWith,
		xKeyfactorApiVersion: &version,
	}
}

// Execute executes the request
//  @return KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse
func (a *IssuedAlertApiService) IssuedAlertAddIssuedAlertExecute(r ApiIssuedAlertAddIssuedAlertRequest) (*KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse
	)

	localBasePath := "/KeyfactorAPI"

	localVarPath := localBasePath + "/Alerts/Issued"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
    if r.xKeyfactorRequestedWith == nil {
        return localVarReturnValue, nil, reportError("xKeyfactorRequestedWith is required and must be specified")
    }
    if r.req == nil {
        return localVarReturnValue, nil, reportError("req is required and must be specified")
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
	localVarPostBody = r.req
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

type ApiIssuedAlertDeleteIssuedAlertRequest struct {
	ctx context.Context
	ApiService *IssuedAlertApiService
	id int32
	xKeyfactorRequestedWith *string
	xKeyfactorApiVersion *string
}

// Type of the request [XMLHttpRequest, APIClient]
func (r ApiIssuedAlertDeleteIssuedAlertRequest) XKeyfactorRequestedWith(xKeyfactorRequestedWith string) ApiIssuedAlertDeleteIssuedAlertRequest {
	r.xKeyfactorRequestedWith = &xKeyfactorRequestedWith
	return r
}

// Desired version of the api, if not provided defaults to v1
func (r ApiIssuedAlertDeleteIssuedAlertRequest) XKeyfactorApiVersion(xKeyfactorApiVersion string) ApiIssuedAlertDeleteIssuedAlertRequest {
	r.xKeyfactorApiVersion = &xKeyfactorApiVersion
	return r
}

func (r ApiIssuedAlertDeleteIssuedAlertRequest) Execute() (*http.Response, error) {
	return r.ApiService.IssuedAlertDeleteIssuedAlertExecute(r)
}

/*
IssuedAlertDeleteIssuedAlert Delete a issued alert

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Id for the issued alert
 @return ApiIssuedAlertDeleteIssuedAlertRequest
*/
func (a *IssuedAlertApiService) IssuedAlertDeleteIssuedAlert(ctx context.Context, id int32) ApiIssuedAlertDeleteIssuedAlertRequest {
    requestedWith := "APIClient"
    version := "1"
	return ApiIssuedAlertDeleteIssuedAlertRequest{
		ApiService: a,
		ctx: ctx,
		xKeyfactorRequestedWith: &requestedWith,
		xKeyfactorApiVersion: &version,
		id: id,
	}
}

// Execute executes the request
func (a *IssuedAlertApiService) IssuedAlertDeleteIssuedAlertExecute(r ApiIssuedAlertDeleteIssuedAlertRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath := "/KeyfactorAPI"

	localVarPath := localBasePath + "/Alerts/Issued/{id}"
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

type ApiIssuedAlertEditIssuedAlertRequest struct {
	ctx context.Context
	ApiService *IssuedAlertApiService
	xKeyfactorRequestedWith *string
	req *KeyfactorApiModelsAlertsIssuedIssuedAlertUpdateRequest
	xKeyfactorApiVersion *string
}

// Type of the request [XMLHttpRequest, APIClient]
func (r ApiIssuedAlertEditIssuedAlertRequest) XKeyfactorRequestedWith(xKeyfactorRequestedWith string) ApiIssuedAlertEditIssuedAlertRequest {
	r.xKeyfactorRequestedWith = &xKeyfactorRequestedWith
	return r
}

// Information for the issued alert
func (r ApiIssuedAlertEditIssuedAlertRequest) Req(req KeyfactorApiModelsAlertsIssuedIssuedAlertUpdateRequest) ApiIssuedAlertEditIssuedAlertRequest {
	r.req = &req
	return r
}

// Desired version of the api, if not provided defaults to v1
func (r ApiIssuedAlertEditIssuedAlertRequest) XKeyfactorApiVersion(xKeyfactorApiVersion string) ApiIssuedAlertEditIssuedAlertRequest {
	r.xKeyfactorApiVersion = &xKeyfactorApiVersion
	return r
}

func (r ApiIssuedAlertEditIssuedAlertRequest) Execute() (*KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse, *http.Response, error) {
	return r.ApiService.IssuedAlertEditIssuedAlertExecute(r)
}

/*
IssuedAlertEditIssuedAlert Edit a issued alert

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiIssuedAlertEditIssuedAlertRequest
*/
func (a *IssuedAlertApiService) IssuedAlertEditIssuedAlert(ctx context.Context) ApiIssuedAlertEditIssuedAlertRequest {
    requestedWith := "APIClient"
    version := "1"
	return ApiIssuedAlertEditIssuedAlertRequest{
		ApiService: a,
		ctx: ctx,
		xKeyfactorRequestedWith: &requestedWith,
		xKeyfactorApiVersion: &version,
	}
}

// Execute executes the request
//  @return KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse
func (a *IssuedAlertApiService) IssuedAlertEditIssuedAlertExecute(r ApiIssuedAlertEditIssuedAlertRequest) (*KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse
	)

	localBasePath := "/KeyfactorAPI"

	localVarPath := localBasePath + "/Alerts/Issued"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
    if r.xKeyfactorRequestedWith == nil {
        return localVarReturnValue, nil, reportError("xKeyfactorRequestedWith is required and must be specified")
    }
    if r.req == nil {
        return localVarReturnValue, nil, reportError("req is required and must be specified")
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
	localVarPostBody = r.req
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

type ApiIssuedAlertEditScheduleRequest struct {
	ctx context.Context
	ApiService *IssuedAlertApiService
	xKeyfactorRequestedWith *string
	newSchedule *KeyfactorApiModelsAlertsAlertScheduleAlertScheduleRequest
	xKeyfactorApiVersion *string
}

// Type of the request [XMLHttpRequest, APIClient]
func (r ApiIssuedAlertEditScheduleRequest) XKeyfactorRequestedWith(xKeyfactorRequestedWith string) ApiIssuedAlertEditScheduleRequest {
	r.xKeyfactorRequestedWith = &xKeyfactorRequestedWith
	return r
}

func (r ApiIssuedAlertEditScheduleRequest) NewSchedule(newSchedule KeyfactorApiModelsAlertsAlertScheduleAlertScheduleRequest) ApiIssuedAlertEditScheduleRequest {
	r.newSchedule = &newSchedule
	return r
}

// Desired version of the api, if not provided defaults to v1
func (r ApiIssuedAlertEditScheduleRequest) XKeyfactorApiVersion(xKeyfactorApiVersion string) ApiIssuedAlertEditScheduleRequest {
	r.xKeyfactorApiVersion = &xKeyfactorApiVersion
	return r
}

func (r ApiIssuedAlertEditScheduleRequest) Execute() (*KeyfactorApiModelsAlertsAlertScheduleAlertScheduleResponse, *http.Response, error) {
	return r.ApiService.IssuedAlertEditScheduleExecute(r)
}

/*
IssuedAlertEditSchedule Edit schedule

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiIssuedAlertEditScheduleRequest
*/
func (a *IssuedAlertApiService) IssuedAlertEditSchedule(ctx context.Context) ApiIssuedAlertEditScheduleRequest {
    requestedWith := "APIClient"
    version := "1"
	return ApiIssuedAlertEditScheduleRequest{
		ApiService: a,
		ctx: ctx,
		xKeyfactorRequestedWith: &requestedWith,
		xKeyfactorApiVersion: &version,
	}
}

// Execute executes the request
//  @return KeyfactorApiModelsAlertsAlertScheduleAlertScheduleResponse
func (a *IssuedAlertApiService) IssuedAlertEditScheduleExecute(r ApiIssuedAlertEditScheduleRequest) (*KeyfactorApiModelsAlertsAlertScheduleAlertScheduleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KeyfactorApiModelsAlertsAlertScheduleAlertScheduleResponse
	)

	localBasePath := "/KeyfactorAPI"

	localVarPath := localBasePath + "/Alerts/Issued/Schedule"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
    if r.xKeyfactorRequestedWith == nil {
        return localVarReturnValue, nil, reportError("xKeyfactorRequestedWith is required and must be specified")
    }
    if r.newSchedule == nil {
        return localVarReturnValue, nil, reportError("newSchedule is required and must be specified")
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
	localVarPostBody = r.newSchedule
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

type ApiIssuedAlertGetIssuedAlertRequest struct {
	ctx context.Context
	ApiService *IssuedAlertApiService
	id int32
	xKeyfactorRequestedWith *string
	xKeyfactorApiVersion *string
}

// Type of the request [XMLHttpRequest, APIClient]
func (r ApiIssuedAlertGetIssuedAlertRequest) XKeyfactorRequestedWith(xKeyfactorRequestedWith string) ApiIssuedAlertGetIssuedAlertRequest {
	r.xKeyfactorRequestedWith = &xKeyfactorRequestedWith
	return r
}

// Desired version of the api, if not provided defaults to v1
func (r ApiIssuedAlertGetIssuedAlertRequest) XKeyfactorApiVersion(xKeyfactorApiVersion string) ApiIssuedAlertGetIssuedAlertRequest {
	r.xKeyfactorApiVersion = &xKeyfactorApiVersion
	return r
}

func (r ApiIssuedAlertGetIssuedAlertRequest) Execute() (*KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse, *http.Response, error) {
	return r.ApiService.IssuedAlertGetIssuedAlertExecute(r)
}

/*
IssuedAlertGetIssuedAlert Get a issued alert

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Id for the issued alert to get
 @return ApiIssuedAlertGetIssuedAlertRequest
*/
func (a *IssuedAlertApiService) IssuedAlertGetIssuedAlert(ctx context.Context, id int32) ApiIssuedAlertGetIssuedAlertRequest {
    requestedWith := "APIClient"
    version := "1"
	return ApiIssuedAlertGetIssuedAlertRequest{
		ApiService: a,
		ctx: ctx,
		xKeyfactorRequestedWith: &requestedWith,
		xKeyfactorApiVersion: &version,
		id: id,
	}
}

// Execute executes the request
//  @return KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse
func (a *IssuedAlertApiService) IssuedAlertGetIssuedAlertExecute(r ApiIssuedAlertGetIssuedAlertRequest) (*KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse
	)

	localBasePath := "/KeyfactorAPI"

	localVarPath := localBasePath + "/Alerts/Issued/{id}"
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

type ApiIssuedAlertGetIssuedAlertsRequest struct {
	ctx context.Context
	ApiService *IssuedAlertApiService
	xKeyfactorRequestedWith *string
	xKeyfactorApiVersion *string
	pagedQueryQueryString *string
	pagedQueryPageReturned *int32
	pagedQueryReturnLimit *int32
	pagedQuerySortField *string
	pagedQuerySortAscending *int32
}

// Type of the request [XMLHttpRequest, APIClient]
func (r ApiIssuedAlertGetIssuedAlertsRequest) XKeyfactorRequestedWith(xKeyfactorRequestedWith string) ApiIssuedAlertGetIssuedAlertsRequest {
	r.xKeyfactorRequestedWith = &xKeyfactorRequestedWith
	return r
}

// Desired version of the api, if not provided defaults to v1
func (r ApiIssuedAlertGetIssuedAlertsRequest) XKeyfactorApiVersion(xKeyfactorApiVersion string) ApiIssuedAlertGetIssuedAlertsRequest {
	r.xKeyfactorApiVersion = &xKeyfactorApiVersion
	return r
}

// Contents of the query (ex: field1 -eq value1 AND field2 -gt value2)
func (r ApiIssuedAlertGetIssuedAlertsRequest) PagedQueryQueryString(pagedQueryQueryString string) ApiIssuedAlertGetIssuedAlertsRequest {
	r.pagedQueryQueryString = &pagedQueryQueryString
	return r
}

// The current page within the result set to be returned
func (r ApiIssuedAlertGetIssuedAlertsRequest) PagedQueryPageReturned(pagedQueryPageReturned int32) ApiIssuedAlertGetIssuedAlertsRequest {
	r.pagedQueryPageReturned = &pagedQueryPageReturned
	return r
}

// Maximum number of records to be returned in a single call
func (r ApiIssuedAlertGetIssuedAlertsRequest) PagedQueryReturnLimit(pagedQueryReturnLimit int32) ApiIssuedAlertGetIssuedAlertsRequest {
	r.pagedQueryReturnLimit = &pagedQueryReturnLimit
	return r
}

// Field by which the results should be sorted (view results via Management Portal for sortable columns)
func (r ApiIssuedAlertGetIssuedAlertsRequest) PagedQuerySortField(pagedQuerySortField string) ApiIssuedAlertGetIssuedAlertsRequest {
	r.pagedQuerySortField = &pagedQuerySortField
	return r
}

// Field sort direction [0&#x3D;ascending, 1&#x3D;descending]
func (r ApiIssuedAlertGetIssuedAlertsRequest) PagedQuerySortAscending(pagedQuerySortAscending int32) ApiIssuedAlertGetIssuedAlertsRequest {
	r.pagedQuerySortAscending = &pagedQuerySortAscending
	return r
}

func (r ApiIssuedAlertGetIssuedAlertsRequest) Execute() ([]KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse, *http.Response, error) {
	return r.ApiService.IssuedAlertGetIssuedAlertsExecute(r)
}

/*
IssuedAlertGetIssuedAlerts Gets all issued alerts according to the provided filter and output parameters

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiIssuedAlertGetIssuedAlertsRequest
*/
func (a *IssuedAlertApiService) IssuedAlertGetIssuedAlerts(ctx context.Context) ApiIssuedAlertGetIssuedAlertsRequest {
    requestedWith := "APIClient"
    version := "1"
	return ApiIssuedAlertGetIssuedAlertsRequest{
		ApiService: a,
		ctx: ctx,
		xKeyfactorRequestedWith: &requestedWith,
		xKeyfactorApiVersion: &version,
	}
}

// Execute executes the request
//  @return []KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse
func (a *IssuedAlertApiService) IssuedAlertGetIssuedAlertsExecute(r ApiIssuedAlertGetIssuedAlertsRequest) ([]KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse
	)

	localBasePath := "/KeyfactorAPI"

	localVarPath := localBasePath + "/Alerts/Issued"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
    if r.xKeyfactorRequestedWith == nil {
        return localVarReturnValue, nil, reportError("xKeyfactorRequestedWith is required and must be specified")
    }

	if r.pagedQueryQueryString != nil {
		parameterAddToQuery(localVarQueryParams, "pagedQuery.queryString", r.pagedQueryQueryString, "")
	}
	if r.pagedQueryPageReturned != nil {
		parameterAddToQuery(localVarQueryParams, "pagedQuery.pageReturned", r.pagedQueryPageReturned, "")
	}
	if r.pagedQueryReturnLimit != nil {
		parameterAddToQuery(localVarQueryParams, "pagedQuery.returnLimit", r.pagedQueryReturnLimit, "")
	}
	if r.pagedQuerySortField != nil {
		parameterAddToQuery(localVarQueryParams, "pagedQuery.sortField", r.pagedQuerySortField, "")
	}
	if r.pagedQuerySortAscending != nil {
		parameterAddToQuery(localVarQueryParams, "pagedQuery.sortAscending", r.pagedQuerySortAscending, "")
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

type ApiIssuedAlertGetScheduleRequest struct {
	ctx context.Context
	ApiService *IssuedAlertApiService
	xKeyfactorRequestedWith *string
	xKeyfactorApiVersion *string
}

// Type of the request [XMLHttpRequest, APIClient]
func (r ApiIssuedAlertGetScheduleRequest) XKeyfactorRequestedWith(xKeyfactorRequestedWith string) ApiIssuedAlertGetScheduleRequest {
	r.xKeyfactorRequestedWith = &xKeyfactorRequestedWith
	return r
}

// Desired version of the api, if not provided defaults to v1
func (r ApiIssuedAlertGetScheduleRequest) XKeyfactorApiVersion(xKeyfactorApiVersion string) ApiIssuedAlertGetScheduleRequest {
	r.xKeyfactorApiVersion = &xKeyfactorApiVersion
	return r
}

func (r ApiIssuedAlertGetScheduleRequest) Execute() (*KeyfactorApiModelsAlertsAlertScheduleAlertScheduleResponse, *http.Response, error) {
	return r.ApiService.IssuedAlertGetScheduleExecute(r)
}

/*
IssuedAlertGetSchedule Get the schedule for issued alerts

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiIssuedAlertGetScheduleRequest
*/
func (a *IssuedAlertApiService) IssuedAlertGetSchedule(ctx context.Context) ApiIssuedAlertGetScheduleRequest {
    requestedWith := "APIClient"
    version := "1"
	return ApiIssuedAlertGetScheduleRequest{
		ApiService: a,
		ctx: ctx,
		xKeyfactorRequestedWith: &requestedWith,
		xKeyfactorApiVersion: &version,
	}
}

// Execute executes the request
//  @return KeyfactorApiModelsAlertsAlertScheduleAlertScheduleResponse
func (a *IssuedAlertApiService) IssuedAlertGetScheduleExecute(r ApiIssuedAlertGetScheduleRequest) (*KeyfactorApiModelsAlertsAlertScheduleAlertScheduleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KeyfactorApiModelsAlertsAlertScheduleAlertScheduleResponse
	)

	localBasePath := "/KeyfactorAPI"

	localVarPath := localBasePath + "/Alerts/Issued/Schedule"

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
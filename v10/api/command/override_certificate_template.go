package command

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ModelsTemplateCollectionRetrievalResponseNP struct {
	Id                     int64                                                                   `json:"Id"`
	CommonName             string                                                                  `json:"CommonName"`
	TemplateName           string                                                                  `json:"TemplateName"`
	Oid                    string                                                                  `json:"Oid"`
	KeySize                string                                                                  `json:"KeySize"`
	KeyType                string                                                                  `json:"KeyType"`
	ForestRoot             *string                                                                 `json:"ForestRoot,omitempty"`
	ConfigurationTenant    string                                                                  `json:"ConfigurationTenant"`
	FriendlyName           *string                                                                 `json:"FriendlyName,omitempty"`
	KeyRetention           int64                                                                   `json:"KeyRetention"`
	KeyRetentionDays       *int64                                                                  `json:"KeyRetentionDays,omitempty"`
	KeyArchival            bool                                                                    `json:"KeyArchival"`
	EnrollmentFields       []ModelsTemplateCollectionRetrievalResponseTemplateEnrollmentFieldModel `json:"EnrollmentFields,omitempty"`
	AllowedEnrollmentTypes int64                                                                   `json:"AllowedEnrollmentTypes"`
	TemplateRegexes        []ModelsTemplateCollectionRetrievalResponseTemplateRegexModel           `json:"TemplateRegexes,omitempty"`
	UseAllowedRequesters   bool                                                                    `json:"UseAllowedRequesters"`
	AllowedRequesters      []string                                                                `json:"AllowedRequesters,omitempty"`
	DisplayName            string                                                                  `json:"DisplayName"`
	RequiresApproval       bool                                                                    `json:"RequiresApproval"`
	KeyUsage               int64                                                                   `json:"KeyUsage"`
	ExtendedKeyUsages      []ModelsTemplateCollectionRetrievalResponseExtendedKeyUsageModel        `json:"ExtendedKeyUsages,omitempty"`
	AdditionalProperties   map[string]interface{}
}

func (r ApiTemplateGetTemplatesRequest) ExecuteNP() ([]ModelsTemplateCollectionRetrievalResponseNP, *http.Response, error) {
	return r.ApiService.TemplateGetTemplatesExecuteNP(r)
}

func (a *TemplateApiService) TemplateGetTemplatesExecuteNP(r ApiTemplateGetTemplatesRequest) ([]ModelsTemplateCollectionRetrievalResponseNP, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []ModelsTemplateCollectionRetrievalResponseNP
	)

	apiBasePath := a.client.cfg.APIPath
	if apiBasePath == "" {
		apiBasePath = "/KeyfactorAPI"
	}

	localVarPath := apiBasePath + "/Templates"

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

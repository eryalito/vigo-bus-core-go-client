/*
Vigo Bus Core API

This is the API for the Vigo Bus Core project.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// IdentityAPIService IdentityAPI service
type IdentityAPIService service

type ApiApiUsersProviderUuidFavoriteStopsStopNumberDeleteRequest struct {
	ctx context.Context
	ApiService *IdentityAPIService
	provider string
	uuid string
	stopNumber int32
}

func (r ApiApiUsersProviderUuidFavoriteStopsStopNumberDeleteRequest) Execute() (*ApiIdentity, *http.Response, error) {
	return r.ApiService.ApiUsersProviderUuidFavoriteStopsStopNumberDeleteExecute(r)
}

/*
ApiUsersProviderUuidFavoriteStopsStopNumberDelete Remove a favorite stop from a user

Remove a favorite stop from a user

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param provider Provider
 @param uuid UUID
 @param stopNumber Stop Number
 @return ApiApiUsersProviderUuidFavoriteStopsStopNumberDeleteRequest
*/
func (a *IdentityAPIService) ApiUsersProviderUuidFavoriteStopsStopNumberDelete(ctx context.Context, provider string, uuid string, stopNumber int32) ApiApiUsersProviderUuidFavoriteStopsStopNumberDeleteRequest {
	return ApiApiUsersProviderUuidFavoriteStopsStopNumberDeleteRequest{
		ApiService: a,
		ctx: ctx,
		provider: provider,
		uuid: uuid,
		stopNumber: stopNumber,
	}
}

// Execute executes the request
//  @return ApiIdentity
func (a *IdentityAPIService) ApiUsersProviderUuidFavoriteStopsStopNumberDeleteExecute(r ApiApiUsersProviderUuidFavoriteStopsStopNumberDeleteRequest) (*ApiIdentity, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiIdentity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityAPIService.ApiUsersProviderUuidFavoriteStopsStopNumberDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/users/{provider}/{uuid}/favorite_stops/{stop_number}"
	localVarPath = strings.Replace(localVarPath, "{"+"provider"+"}", url.PathEscape(parameterValueToString(r.provider, "provider")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"uuid"+"}", url.PathEscape(parameterValueToString(r.uuid, "uuid")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"stop_number"+"}", url.PathEscape(parameterValueToString(r.stopNumber, "stopNumber")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["BearerAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiApiUsersProviderUuidFavoriteStopsStopNumberPostRequest struct {
	ctx context.Context
	ApiService *IdentityAPIService
	provider string
	uuid string
	stopNumber int32
}

func (r ApiApiUsersProviderUuidFavoriteStopsStopNumberPostRequest) Execute() (*ApiIdentity, *http.Response, error) {
	return r.ApiService.ApiUsersProviderUuidFavoriteStopsStopNumberPostExecute(r)
}

/*
ApiUsersProviderUuidFavoriteStopsStopNumberPost Add a favorite stop to a user

Add a favorite stop to a user

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param provider Provider
 @param uuid UUID
 @param stopNumber Stop Number
 @return ApiApiUsersProviderUuidFavoriteStopsStopNumberPostRequest
*/
func (a *IdentityAPIService) ApiUsersProviderUuidFavoriteStopsStopNumberPost(ctx context.Context, provider string, uuid string, stopNumber int32) ApiApiUsersProviderUuidFavoriteStopsStopNumberPostRequest {
	return ApiApiUsersProviderUuidFavoriteStopsStopNumberPostRequest{
		ApiService: a,
		ctx: ctx,
		provider: provider,
		uuid: uuid,
		stopNumber: stopNumber,
	}
}

// Execute executes the request
//  @return ApiIdentity
func (a *IdentityAPIService) ApiUsersProviderUuidFavoriteStopsStopNumberPostExecute(r ApiApiUsersProviderUuidFavoriteStopsStopNumberPostRequest) (*ApiIdentity, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiIdentity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityAPIService.ApiUsersProviderUuidFavoriteStopsStopNumberPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/users/{provider}/{uuid}/favorite_stops/{stop_number}"
	localVarPath = strings.Replace(localVarPath, "{"+"provider"+"}", url.PathEscape(parameterValueToString(r.provider, "provider")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"uuid"+"}", url.PathEscape(parameterValueToString(r.uuid, "uuid")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"stop_number"+"}", url.PathEscape(parameterValueToString(r.stopNumber, "stopNumber")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["BearerAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiApiUsersProviderUuidGetRequest struct {
	ctx context.Context
	ApiService *IdentityAPIService
	provider string
	uuid string
}

func (r ApiApiUsersProviderUuidGetRequest) Execute() (*ApiIdentity, *http.Response, error) {
	return r.ApiService.ApiUsersProviderUuidGetExecute(r)
}

/*
ApiUsersProviderUuidGet Get a user by its UUID for a specific provider

Provide a user by its UUID for a specific provider

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param provider Provider
 @param uuid UUID
 @return ApiApiUsersProviderUuidGetRequest
*/
func (a *IdentityAPIService) ApiUsersProviderUuidGet(ctx context.Context, provider string, uuid string) ApiApiUsersProviderUuidGetRequest {
	return ApiApiUsersProviderUuidGetRequest{
		ApiService: a,
		ctx: ctx,
		provider: provider,
		uuid: uuid,
	}
}

// Execute executes the request
//  @return ApiIdentity
func (a *IdentityAPIService) ApiUsersProviderUuidGetExecute(r ApiApiUsersProviderUuidGetRequest) (*ApiIdentity, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiIdentity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityAPIService.ApiUsersProviderUuidGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/users/{provider}/{uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"provider"+"}", url.PathEscape(parameterValueToString(r.provider, "provider")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"uuid"+"}", url.PathEscape(parameterValueToString(r.uuid, "uuid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["BearerAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiApiUsersProviderUuidPostRequest struct {
	ctx context.Context
	ApiService *IdentityAPIService
	provider string
	uuid string
}

func (r ApiApiUsersProviderUuidPostRequest) Execute() (*ApiIdentity, *http.Response, error) {
	return r.ApiService.ApiUsersProviderUuidPostExecute(r)
}

/*
ApiUsersProviderUuidPost Create a new user

Create a new user

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param provider Provider
 @param uuid UUID
 @return ApiApiUsersProviderUuidPostRequest
*/
func (a *IdentityAPIService) ApiUsersProviderUuidPost(ctx context.Context, provider string, uuid string) ApiApiUsersProviderUuidPostRequest {
	return ApiApiUsersProviderUuidPostRequest{
		ApiService: a,
		ctx: ctx,
		provider: provider,
		uuid: uuid,
	}
}

// Execute executes the request
//  @return ApiIdentity
func (a *IdentityAPIService) ApiUsersProviderUuidPostExecute(r ApiApiUsersProviderUuidPostRequest) (*ApiIdentity, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiIdentity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityAPIService.ApiUsersProviderUuidPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/users/{provider}/{uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"provider"+"}", url.PathEscape(parameterValueToString(r.provider, "provider")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"uuid"+"}", url.PathEscape(parameterValueToString(r.uuid, "uuid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["BearerAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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
# \IdentityAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiUsersProviderUuidFavoriteStopsStopNumberDelete**](IdentityAPI.md#ApiUsersProviderUuidFavoriteStopsStopNumberDelete) | **Delete** /api/users/{provider}/{uuid}/favorite_stops/{stop_number} | Remove a favorite stop from a user
[**ApiUsersProviderUuidFavoriteStopsStopNumberPost**](IdentityAPI.md#ApiUsersProviderUuidFavoriteStopsStopNumberPost) | **Post** /api/users/{provider}/{uuid}/favorite_stops/{stop_number} | Add a favorite stop to a user
[**ApiUsersProviderUuidGet**](IdentityAPI.md#ApiUsersProviderUuidGet) | **Get** /api/users/{provider}/{uuid} | Get a user by its UUID for a specific provider
[**ApiUsersProviderUuidMetadataPut**](IdentityAPI.md#ApiUsersProviderUuidMetadataPut) | **Put** /api/users/{provider}/{uuid}/metadata | Update the metadata of a user
[**ApiUsersProviderUuidPost**](IdentityAPI.md#ApiUsersProviderUuidPost) | **Post** /api/users/{provider}/{uuid} | Create a new user



## ApiUsersProviderUuidFavoriteStopsStopNumberDelete

> ApiIdentity ApiUsersProviderUuidFavoriteStopsStopNumberDelete(ctx, provider, uuid, stopNumber).Execute()

Remove a favorite stop from a user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eryalito/vigo-bus-core-go-client"
)

func main() {
	provider := "provider_example" // string | Provider
	uuid := "uuid_example" // string | UUID
	stopNumber := int32(56) // int32 | Stop Number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityAPI.ApiUsersProviderUuidFavoriteStopsStopNumberDelete(context.Background(), provider, uuid, stopNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityAPI.ApiUsersProviderUuidFavoriteStopsStopNumberDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiUsersProviderUuidFavoriteStopsStopNumberDelete`: ApiIdentity
	fmt.Fprintf(os.Stdout, "Response from `IdentityAPI.ApiUsersProviderUuidFavoriteStopsStopNumberDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider | 
**uuid** | **string** | UUID | 
**stopNumber** | **int32** | Stop Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiUsersProviderUuidFavoriteStopsStopNumberDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ApiIdentity**](ApiIdentity.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiUsersProviderUuidFavoriteStopsStopNumberPost

> ApiIdentity ApiUsersProviderUuidFavoriteStopsStopNumberPost(ctx, provider, uuid, stopNumber).Execute()

Add a favorite stop to a user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eryalito/vigo-bus-core-go-client"
)

func main() {
	provider := "provider_example" // string | Provider
	uuid := "uuid_example" // string | UUID
	stopNumber := int32(56) // int32 | Stop Number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityAPI.ApiUsersProviderUuidFavoriteStopsStopNumberPost(context.Background(), provider, uuid, stopNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityAPI.ApiUsersProviderUuidFavoriteStopsStopNumberPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiUsersProviderUuidFavoriteStopsStopNumberPost`: ApiIdentity
	fmt.Fprintf(os.Stdout, "Response from `IdentityAPI.ApiUsersProviderUuidFavoriteStopsStopNumberPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider | 
**uuid** | **string** | UUID | 
**stopNumber** | **int32** | Stop Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiUsersProviderUuidFavoriteStopsStopNumberPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ApiIdentity**](ApiIdentity.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiUsersProviderUuidGet

> ApiIdentity ApiUsersProviderUuidGet(ctx, provider, uuid).Execute()

Get a user by its UUID for a specific provider



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eryalito/vigo-bus-core-go-client"
)

func main() {
	provider := "provider_example" // string | Provider
	uuid := "uuid_example" // string | UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityAPI.ApiUsersProviderUuidGet(context.Background(), provider, uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityAPI.ApiUsersProviderUuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiUsersProviderUuidGet`: ApiIdentity
	fmt.Fprintf(os.Stdout, "Response from `IdentityAPI.ApiUsersProviderUuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider | 
**uuid** | **string** | UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiUsersProviderUuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiIdentity**](ApiIdentity.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiUsersProviderUuidMetadataPut

> ApiIdentity ApiUsersProviderUuidMetadataPut(ctx, provider, uuid).Body(body).Execute()

Update the metadata of a user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eryalito/vigo-bus-core-go-client"
)

func main() {
	provider := "provider_example" // string | Provider
	uuid := "uuid_example" // string | UUID
	body := "body_example" // string | Metadata

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityAPI.ApiUsersProviderUuidMetadataPut(context.Background(), provider, uuid).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityAPI.ApiUsersProviderUuidMetadataPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiUsersProviderUuidMetadataPut`: ApiIdentity
	fmt.Fprintf(os.Stdout, "Response from `IdentityAPI.ApiUsersProviderUuidMetadataPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider | 
**uuid** | **string** | UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiUsersProviderUuidMetadataPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** | Metadata | 

### Return type

[**ApiIdentity**](ApiIdentity.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiUsersProviderUuidPost

> ApiIdentity ApiUsersProviderUuidPost(ctx, provider, uuid).Execute()

Create a new user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eryalito/vigo-bus-core-go-client"
)

func main() {
	provider := "provider_example" // string | Provider
	uuid := "uuid_example" // string | UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityAPI.ApiUsersProviderUuidPost(context.Background(), provider, uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityAPI.ApiUsersProviderUuidPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiUsersProviderUuidPost`: ApiIdentity
	fmt.Fprintf(os.Stdout, "Response from `IdentityAPI.ApiUsersProviderUuidPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider | 
**uuid** | **string** | UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiUsersProviderUuidPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiIdentity**](ApiIdentity.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


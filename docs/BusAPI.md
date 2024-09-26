# \BusAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiLinesGet**](BusAPI.md#ApiLinesGet) | **Get** /api/lines | List all of the lines
[**ApiStopsFindGet**](BusAPI.md#ApiStopsFindGet) | **Get** /api/stops/find | Find a stop by text in its name
[**ApiStopsFindLocationGet**](BusAPI.md#ApiStopsFindLocationGet) | **Get** /api/stops/find/location | Find a stop by its location
[**ApiStopsFindLocationImageGet**](BusAPI.md#ApiStopsFindLocationImageGet) | **Get** /api/stops/find/location/image | Get the nearby stops as a PNG image and JSON array
[**ApiStopsGet**](BusAPI.md#ApiStopsGet) | **Get** /api/stops | List all of the stops
[**ApiStopsStopNumberGet**](BusAPI.md#ApiStopsStopNumberGet) | **Get** /api/stops/{stop_number} | Get a stop by its number
[**ApiStopsStopNumberScheduleGet**](BusAPI.md#ApiStopsStopNumberScheduleGet) | **Get** /api/stops/{stop_number}/schedule | Get the schedule for a stop



## ApiLinesGet

> []ApiLine ApiLinesGet(ctx).Execute()

List all of the lines



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BusAPI.ApiLinesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BusAPI.ApiLinesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiLinesGet`: []ApiLine
	fmt.Fprintf(os.Stdout, "Response from `BusAPI.ApiLinesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiLinesGetRequest struct via the builder pattern


### Return type

[**[]ApiLine**](ApiLine.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiStopsFindGet

> []ApiStop ApiStopsFindGet(ctx).Text(text).Execute()

Find a stop by text in its name



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
	text := "text_example" // string | Text to search for in stop name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BusAPI.ApiStopsFindGet(context.Background()).Text(text).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BusAPI.ApiStopsFindGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiStopsFindGet`: []ApiStop
	fmt.Fprintf(os.Stdout, "Response from `BusAPI.ApiStopsFindGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiStopsFindGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **text** | **string** | Text to search for in stop name | 

### Return type

[**[]ApiStop**](ApiStop.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiStopsFindLocationGet

> []ApiStop ApiStopsFindLocationGet(ctx).Lat(lat).Lon(lon).Radius(radius).Execute()

Find a stop by its location



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
	lat := float32(8.14) // float32 | Latitude
	lon := float32(8.14) // float32 | Longitude
	radius := float32(8.14) // float32 | Radius in meters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BusAPI.ApiStopsFindLocationGet(context.Background()).Lat(lat).Lon(lon).Radius(radius).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BusAPI.ApiStopsFindLocationGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiStopsFindLocationGet`: []ApiStop
	fmt.Fprintf(os.Stdout, "Response from `BusAPI.ApiStopsFindLocationGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiStopsFindLocationGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lat** | **float32** | Latitude | 
 **lon** | **float32** | Longitude | 
 **radius** | **float32** | Radius in meters | 

### Return type

[**[]ApiStop**](ApiStop.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiStopsFindLocationImageGet

> ApiNearbyStops ApiStopsFindLocationImageGet(ctx).Lat(lat).Lon(lon).Radius(radius).Limit(limit).Execute()

Get the nearby stops as a PNG image and JSON array



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
	lat := float32(8.14) // float32 | Latitude
	lon := float32(8.14) // float32 | Longitude
	radius := float32(8.14) // float32 | Radius in meters
	limit := int32(56) // int32 | Limit of stops to return, default 9 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BusAPI.ApiStopsFindLocationImageGet(context.Background()).Lat(lat).Lon(lon).Radius(radius).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BusAPI.ApiStopsFindLocationImageGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiStopsFindLocationImageGet`: ApiNearbyStops
	fmt.Fprintf(os.Stdout, "Response from `BusAPI.ApiStopsFindLocationImageGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiStopsFindLocationImageGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lat** | **float32** | Latitude | 
 **lon** | **float32** | Longitude | 
 **radius** | **float32** | Radius in meters | 
 **limit** | **int32** | Limit of stops to return, default 9 | 

### Return type

[**ApiNearbyStops**](ApiNearbyStops.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiStopsGet

> []ApiStop ApiStopsGet(ctx).Execute()

List all of the stops



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BusAPI.ApiStopsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BusAPI.ApiStopsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiStopsGet`: []ApiStop
	fmt.Fprintf(os.Stdout, "Response from `BusAPI.ApiStopsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiStopsGetRequest struct via the builder pattern


### Return type

[**[]ApiStop**](ApiStop.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiStopsStopNumberGet

> ApiStop ApiStopsStopNumberGet(ctx, stopNumber).Execute()

Get a stop by its number



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
	stopNumber := int32(56) // int32 | Stop Number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BusAPI.ApiStopsStopNumberGet(context.Background(), stopNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BusAPI.ApiStopsStopNumberGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiStopsStopNumberGet`: ApiStop
	fmt.Fprintf(os.Stdout, "Response from `BusAPI.ApiStopsStopNumberGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stopNumber** | **int32** | Stop Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiStopsStopNumberGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiStop**](ApiStop.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiStopsStopNumberScheduleGet

> ApiStopSchedule ApiStopsStopNumberScheduleGet(ctx, stopNumber).Execute()

Get the schedule for a stop



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
	stopNumber := int32(56) // int32 | Stop Number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BusAPI.ApiStopsStopNumberScheduleGet(context.Background(), stopNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BusAPI.ApiStopsStopNumberScheduleGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiStopsStopNumberScheduleGet`: ApiStopSchedule
	fmt.Fprintf(os.Stdout, "Response from `BusAPI.ApiStopsStopNumberScheduleGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stopNumber** | **int32** | Stop Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiStopsStopNumberScheduleGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiStopSchedule**](ApiStopSchedule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


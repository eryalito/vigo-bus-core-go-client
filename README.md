# Go API client for openapi

This is the API for the Vigo Bus Core project.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 1.0.0
- Generator version: 7.8.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import openapi "github.com/eryalito/vigo-bus-core-go-client"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `openapi.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `openapi.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `openapi.ContextOperationServerIndices` and `openapi.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:8080*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*BusAPI* | [**ApiLinesGet**](docs/BusAPI.md#apilinesget) | **Get** /api/lines | List all of the lines
*BusAPI* | [**ApiStopsFindGet**](docs/BusAPI.md#apistopsfindget) | **Get** /api/stops/find | Find a stop by text in its name
*BusAPI* | [**ApiStopsFindLocationGet**](docs/BusAPI.md#apistopsfindlocationget) | **Get** /api/stops/find/location | Find a stop by its location
*BusAPI* | [**ApiStopsFindLocationImageGet**](docs/BusAPI.md#apistopsfindlocationimageget) | **Get** /api/stops/find/location/image | Get the nearby stops as a PNG image and JSON array
*BusAPI* | [**ApiStopsGet**](docs/BusAPI.md#apistopsget) | **Get** /api/stops | List all of the stops
*BusAPI* | [**ApiStopsStopNumberGet**](docs/BusAPI.md#apistopsstopnumberget) | **Get** /api/stops/{stop_number} | Get a stop by its number
*BusAPI* | [**ApiStopsStopNumberScheduleGet**](docs/BusAPI.md#apistopsstopnumberscheduleget) | **Get** /api/stops/{stop_number}/schedule | Get the schedule for a stop
*IdentityAPI* | [**ApiUsersProviderUuidFavoriteStopsStopNumberDelete**](docs/IdentityAPI.md#apiusersprovideruuidfavoritestopsstopnumberdelete) | **Delete** /api/users/{provider}/{uuid}/favorite_stops/{stop_number} | Remove a favorite stop from a user
*IdentityAPI* | [**ApiUsersProviderUuidFavoriteStopsStopNumberPost**](docs/IdentityAPI.md#apiusersprovideruuidfavoritestopsstopnumberpost) | **Post** /api/users/{provider}/{uuid}/favorite_stops/{stop_number} | Add a favorite stop to a user
*IdentityAPI* | [**ApiUsersProviderUuidGet**](docs/IdentityAPI.md#apiusersprovideruuidget) | **Get** /api/users/{provider}/{uuid} | Get a user by its UUID for a specific provider
*IdentityAPI* | [**ApiUsersProviderUuidMetadataPut**](docs/IdentityAPI.md#apiusersprovideruuidmetadataput) | **Put** /api/users/{provider}/{uuid}/metadata | Update the metadata of a user
*IdentityAPI* | [**ApiUsersProviderUuidPost**](docs/IdentityAPI.md#apiusersprovideruuidpost) | **Post** /api/users/{provider}/{uuid} | Create a new user


## Documentation For Models

 - [ApiIdentity](docs/ApiIdentity.md)
 - [ApiLine](docs/ApiLine.md)
 - [ApiNearbyStops](docs/ApiNearbyStops.md)
 - [ApiNearbyStopsOrigin](docs/ApiNearbyStopsOrigin.md)
 - [ApiProviderType](docs/ApiProviderType.md)
 - [ApiSchedule](docs/ApiSchedule.md)
 - [ApiStop](docs/ApiStop.md)
 - [ApiStopLocation](docs/ApiStopLocation.md)
 - [ApiStopSchedule](docs/ApiStopSchedule.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### BearerAuth

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		openapi.ContextAPIKeys,
		map[string]openapi.APIKey{
			"Authorization": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author




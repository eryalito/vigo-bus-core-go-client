# ApiStop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID is the unique identifier of the stop | [optional] 
**Location** | Pointer to [**ApiStopLocation**](ApiStopLocation.md) |  | [optional] 
**Name** | Pointer to **string** | Name is the name of the stop | [optional] 
**StopId** | Pointer to **int32** | StopID is the number of the stop used internally by the bus company | [optional] 
**StopNumber** | Pointer to **int32** | StopNumber is the number of the stop provided by the bus company | [optional] 

## Methods

### NewApiStop

`func NewApiStop() *ApiStop`

NewApiStop instantiates a new ApiStop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiStopWithDefaults

`func NewApiStopWithDefaults() *ApiStop`

NewApiStopWithDefaults instantiates a new ApiStop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiStop) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiStop) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiStop) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ApiStop) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocation

`func (o *ApiStop) GetLocation() ApiStopLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ApiStop) GetLocationOk() (*ApiStopLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ApiStop) SetLocation(v ApiStopLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ApiStop) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetName

`func (o *ApiStop) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiStop) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiStop) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiStop) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStopId

`func (o *ApiStop) GetStopId() int32`

GetStopId returns the StopId field if non-nil, zero value otherwise.

### GetStopIdOk

`func (o *ApiStop) GetStopIdOk() (*int32, bool)`

GetStopIdOk returns a tuple with the StopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopId

`func (o *ApiStop) SetStopId(v int32)`

SetStopId sets StopId field to given value.

### HasStopId

`func (o *ApiStop) HasStopId() bool`

HasStopId returns a boolean if a field has been set.

### GetStopNumber

`func (o *ApiStop) GetStopNumber() int32`

GetStopNumber returns the StopNumber field if non-nil, zero value otherwise.

### GetStopNumberOk

`func (o *ApiStop) GetStopNumberOk() (*int32, bool)`

GetStopNumberOk returns a tuple with the StopNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopNumber

`func (o *ApiStop) SetStopNumber(v int32)`

SetStopNumber sets StopNumber field to given value.

### HasStopNumber

`func (o *ApiStop) HasStopNumber() bool`

HasStopNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



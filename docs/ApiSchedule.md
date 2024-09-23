# ApiSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Line** | Pointer to [**ApiLine**](ApiLine.md) | Line is the line that the schedule is for | [optional] 
**Route** | Pointer to **string** | Route is the route that the schedule is for | [optional] 
**Time** | Pointer to **int32** | Time is the time of the schedule | [optional] 

## Methods

### NewApiSchedule

`func NewApiSchedule() *ApiSchedule`

NewApiSchedule instantiates a new ApiSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiScheduleWithDefaults

`func NewApiScheduleWithDefaults() *ApiSchedule`

NewApiScheduleWithDefaults instantiates a new ApiSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLine

`func (o *ApiSchedule) GetLine() ApiLine`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *ApiSchedule) GetLineOk() (*ApiLine, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *ApiSchedule) SetLine(v ApiLine)`

SetLine sets Line field to given value.

### HasLine

`func (o *ApiSchedule) HasLine() bool`

HasLine returns a boolean if a field has been set.

### GetRoute

`func (o *ApiSchedule) GetRoute() string`

GetRoute returns the Route field if non-nil, zero value otherwise.

### GetRouteOk

`func (o *ApiSchedule) GetRouteOk() (*string, bool)`

GetRouteOk returns a tuple with the Route field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoute

`func (o *ApiSchedule) SetRoute(v string)`

SetRoute sets Route field to given value.

### HasRoute

`func (o *ApiSchedule) HasRoute() bool`

HasRoute returns a boolean if a field has been set.

### GetTime

`func (o *ApiSchedule) GetTime() int32`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ApiSchedule) GetTimeOk() (*int32, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ApiSchedule) SetTime(v int32)`

SetTime sets Time field to given value.

### HasTime

`func (o *ApiSchedule) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



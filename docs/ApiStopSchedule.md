# ApiStopSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schedules** | Pointer to [**[]ApiSchedule**](ApiSchedule.md) | Schedules is a list of the schedules for the stop | [optional] 
**Stop** | Pointer to [**ApiStop**](ApiStop.md) | Stop is the stop that the schedule is for | [optional] 

## Methods

### NewApiStopSchedule

`func NewApiStopSchedule() *ApiStopSchedule`

NewApiStopSchedule instantiates a new ApiStopSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiStopScheduleWithDefaults

`func NewApiStopScheduleWithDefaults() *ApiStopSchedule`

NewApiStopScheduleWithDefaults instantiates a new ApiStopSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchedules

`func (o *ApiStopSchedule) GetSchedules() []ApiSchedule`

GetSchedules returns the Schedules field if non-nil, zero value otherwise.

### GetSchedulesOk

`func (o *ApiStopSchedule) GetSchedulesOk() (*[]ApiSchedule, bool)`

GetSchedulesOk returns a tuple with the Schedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedules

`func (o *ApiStopSchedule) SetSchedules(v []ApiSchedule)`

SetSchedules sets Schedules field to given value.

### HasSchedules

`func (o *ApiStopSchedule) HasSchedules() bool`

HasSchedules returns a boolean if a field has been set.

### GetStop

`func (o *ApiStopSchedule) GetStop() ApiStop`

GetStop returns the Stop field if non-nil, zero value otherwise.

### GetStopOk

`func (o *ApiStopSchedule) GetStopOk() (*ApiStop, bool)`

GetStopOk returns a tuple with the Stop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStop

`func (o *ApiStopSchedule) SetStop(v ApiStop)`

SetStop sets Stop field to given value.

### HasStop

`func (o *ApiStopSchedule) HasStop() bool`

HasStop returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



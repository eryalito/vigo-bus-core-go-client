/*
Vigo Bus Core API

This is the API for the Vigo Bus Core project.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ApiStopSchedule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiStopSchedule{}

// ApiStopSchedule struct for ApiStopSchedule
type ApiStopSchedule struct {
	// Schedules is a list of the schedules for the stop
	Schedules []ApiSchedule `json:"schedules,omitempty"`
	// Stop is the stop that the schedule is for
	Stop *ApiStop `json:"stop,omitempty"`
}

// NewApiStopSchedule instantiates a new ApiStopSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiStopSchedule() *ApiStopSchedule {
	this := ApiStopSchedule{}
	return &this
}

// NewApiStopScheduleWithDefaults instantiates a new ApiStopSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiStopScheduleWithDefaults() *ApiStopSchedule {
	this := ApiStopSchedule{}
	return &this
}

// GetSchedules returns the Schedules field value if set, zero value otherwise.
func (o *ApiStopSchedule) GetSchedules() []ApiSchedule {
	if o == nil || IsNil(o.Schedules) {
		var ret []ApiSchedule
		return ret
	}
	return o.Schedules
}

// GetSchedulesOk returns a tuple with the Schedules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStopSchedule) GetSchedulesOk() ([]ApiSchedule, bool) {
	if o == nil || IsNil(o.Schedules) {
		return nil, false
	}
	return o.Schedules, true
}

// HasSchedules returns a boolean if a field has been set.
func (o *ApiStopSchedule) HasSchedules() bool {
	if o != nil && !IsNil(o.Schedules) {
		return true
	}

	return false
}

// SetSchedules gets a reference to the given []ApiSchedule and assigns it to the Schedules field.
func (o *ApiStopSchedule) SetSchedules(v []ApiSchedule) {
	o.Schedules = v
}

// GetStop returns the Stop field value if set, zero value otherwise.
func (o *ApiStopSchedule) GetStop() ApiStop {
	if o == nil || IsNil(o.Stop) {
		var ret ApiStop
		return ret
	}
	return *o.Stop
}

// GetStopOk returns a tuple with the Stop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStopSchedule) GetStopOk() (*ApiStop, bool) {
	if o == nil || IsNil(o.Stop) {
		return nil, false
	}
	return o.Stop, true
}

// HasStop returns a boolean if a field has been set.
func (o *ApiStopSchedule) HasStop() bool {
	if o != nil && !IsNil(o.Stop) {
		return true
	}

	return false
}

// SetStop gets a reference to the given ApiStop and assigns it to the Stop field.
func (o *ApiStopSchedule) SetStop(v ApiStop) {
	o.Stop = &v
}

func (o ApiStopSchedule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiStopSchedule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schedules) {
		toSerialize["schedules"] = o.Schedules
	}
	if !IsNil(o.Stop) {
		toSerialize["stop"] = o.Stop
	}
	return toSerialize, nil
}

type NullableApiStopSchedule struct {
	value *ApiStopSchedule
	isSet bool
}

func (v NullableApiStopSchedule) Get() *ApiStopSchedule {
	return v.value
}

func (v *NullableApiStopSchedule) Set(val *ApiStopSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableApiStopSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableApiStopSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiStopSchedule(val *ApiStopSchedule) *NullableApiStopSchedule {
	return &NullableApiStopSchedule{value: val, isSet: true}
}

func (v NullableApiStopSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiStopSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



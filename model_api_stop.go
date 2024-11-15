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

// checks if the ApiStop type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiStop{}

// ApiStop struct for ApiStop
type ApiStop struct {
	// ID is the unique identifier of the stop
	Id *int32 `json:"id,omitempty"`
	Location *ApiStopLocation `json:"location,omitempty"`
	// Name is the name of the stop
	Name *string `json:"name,omitempty"`
	// StopID is the number of the stop used internally by the bus company
	StopId *int32 `json:"stop_id,omitempty"`
	// StopNumber is the number of the stop provided by the bus company
	StopNumber *int32 `json:"stop_number,omitempty"`
}

// NewApiStop instantiates a new ApiStop object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiStop() *ApiStop {
	this := ApiStop{}
	return &this
}

// NewApiStopWithDefaults instantiates a new ApiStop object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiStopWithDefaults() *ApiStop {
	this := ApiStop{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiStop) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStop) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiStop) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ApiStop) SetId(v int32) {
	o.Id = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *ApiStop) GetLocation() ApiStopLocation {
	if o == nil || IsNil(o.Location) {
		var ret ApiStopLocation
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStop) GetLocationOk() (*ApiStopLocation, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *ApiStop) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given ApiStopLocation and assigns it to the Location field.
func (o *ApiStop) SetLocation(v ApiStopLocation) {
	o.Location = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiStop) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStop) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiStop) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiStop) SetName(v string) {
	o.Name = &v
}

// GetStopId returns the StopId field value if set, zero value otherwise.
func (o *ApiStop) GetStopId() int32 {
	if o == nil || IsNil(o.StopId) {
		var ret int32
		return ret
	}
	return *o.StopId
}

// GetStopIdOk returns a tuple with the StopId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStop) GetStopIdOk() (*int32, bool) {
	if o == nil || IsNil(o.StopId) {
		return nil, false
	}
	return o.StopId, true
}

// HasStopId returns a boolean if a field has been set.
func (o *ApiStop) HasStopId() bool {
	if o != nil && !IsNil(o.StopId) {
		return true
	}

	return false
}

// SetStopId gets a reference to the given int32 and assigns it to the StopId field.
func (o *ApiStop) SetStopId(v int32) {
	o.StopId = &v
}

// GetStopNumber returns the StopNumber field value if set, zero value otherwise.
func (o *ApiStop) GetStopNumber() int32 {
	if o == nil || IsNil(o.StopNumber) {
		var ret int32
		return ret
	}
	return *o.StopNumber
}

// GetStopNumberOk returns a tuple with the StopNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStop) GetStopNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.StopNumber) {
		return nil, false
	}
	return o.StopNumber, true
}

// HasStopNumber returns a boolean if a field has been set.
func (o *ApiStop) HasStopNumber() bool {
	if o != nil && !IsNil(o.StopNumber) {
		return true
	}

	return false
}

// SetStopNumber gets a reference to the given int32 and assigns it to the StopNumber field.
func (o *ApiStop) SetStopNumber(v int32) {
	o.StopNumber = &v
}

func (o ApiStop) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiStop) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.StopId) {
		toSerialize["stop_id"] = o.StopId
	}
	if !IsNil(o.StopNumber) {
		toSerialize["stop_number"] = o.StopNumber
	}
	return toSerialize, nil
}

type NullableApiStop struct {
	value *ApiStop
	isSet bool
}

func (v NullableApiStop) Get() *ApiStop {
	return v.value
}

func (v *NullableApiStop) Set(val *ApiStop) {
	v.value = val
	v.isSet = true
}

func (v NullableApiStop) IsSet() bool {
	return v.isSet
}

func (v *NullableApiStop) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiStop(val *ApiStop) *NullableApiStop {
	return &NullableApiStop{value: val, isSet: true}
}

func (v NullableApiStop) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiStop) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



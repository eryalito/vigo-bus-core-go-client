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

// checks if the ApiNearbyStops type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiNearbyStops{}

// ApiNearbyStops struct for ApiNearbyStops
type ApiNearbyStops struct {
	Image *string `json:"image,omitempty"`
	Origin *ApiNearbyStopsOrigin `json:"origin,omitempty"`
	Radius *float32 `json:"radius,omitempty"`
	Stops []ApiStop `json:"stops,omitempty"`
}

// NewApiNearbyStops instantiates a new ApiNearbyStops object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiNearbyStops() *ApiNearbyStops {
	this := ApiNearbyStops{}
	return &this
}

// NewApiNearbyStopsWithDefaults instantiates a new ApiNearbyStops object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiNearbyStopsWithDefaults() *ApiNearbyStops {
	this := ApiNearbyStops{}
	return &this
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *ApiNearbyStops) GetImage() string {
	if o == nil || IsNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNearbyStops) GetImageOk() (*string, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *ApiNearbyStops) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *ApiNearbyStops) SetImage(v string) {
	o.Image = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *ApiNearbyStops) GetOrigin() ApiNearbyStopsOrigin {
	if o == nil || IsNil(o.Origin) {
		var ret ApiNearbyStopsOrigin
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNearbyStops) GetOriginOk() (*ApiNearbyStopsOrigin, bool) {
	if o == nil || IsNil(o.Origin) {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *ApiNearbyStops) HasOrigin() bool {
	if o != nil && !IsNil(o.Origin) {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given ApiNearbyStopsOrigin and assigns it to the Origin field.
func (o *ApiNearbyStops) SetOrigin(v ApiNearbyStopsOrigin) {
	o.Origin = &v
}

// GetRadius returns the Radius field value if set, zero value otherwise.
func (o *ApiNearbyStops) GetRadius() float32 {
	if o == nil || IsNil(o.Radius) {
		var ret float32
		return ret
	}
	return *o.Radius
}

// GetRadiusOk returns a tuple with the Radius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNearbyStops) GetRadiusOk() (*float32, bool) {
	if o == nil || IsNil(o.Radius) {
		return nil, false
	}
	return o.Radius, true
}

// HasRadius returns a boolean if a field has been set.
func (o *ApiNearbyStops) HasRadius() bool {
	if o != nil && !IsNil(o.Radius) {
		return true
	}

	return false
}

// SetRadius gets a reference to the given float32 and assigns it to the Radius field.
func (o *ApiNearbyStops) SetRadius(v float32) {
	o.Radius = &v
}

// GetStops returns the Stops field value if set, zero value otherwise.
func (o *ApiNearbyStops) GetStops() []ApiStop {
	if o == nil || IsNil(o.Stops) {
		var ret []ApiStop
		return ret
	}
	return o.Stops
}

// GetStopsOk returns a tuple with the Stops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNearbyStops) GetStopsOk() ([]ApiStop, bool) {
	if o == nil || IsNil(o.Stops) {
		return nil, false
	}
	return o.Stops, true
}

// HasStops returns a boolean if a field has been set.
func (o *ApiNearbyStops) HasStops() bool {
	if o != nil && !IsNil(o.Stops) {
		return true
	}

	return false
}

// SetStops gets a reference to the given []ApiStop and assigns it to the Stops field.
func (o *ApiNearbyStops) SetStops(v []ApiStop) {
	o.Stops = v
}

func (o ApiNearbyStops) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiNearbyStops) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !IsNil(o.Origin) {
		toSerialize["origin"] = o.Origin
	}
	if !IsNil(o.Radius) {
		toSerialize["radius"] = o.Radius
	}
	if !IsNil(o.Stops) {
		toSerialize["stops"] = o.Stops
	}
	return toSerialize, nil
}

type NullableApiNearbyStops struct {
	value *ApiNearbyStops
	isSet bool
}

func (v NullableApiNearbyStops) Get() *ApiNearbyStops {
	return v.value
}

func (v *NullableApiNearbyStops) Set(val *ApiNearbyStops) {
	v.value = val
	v.isSet = true
}

func (v NullableApiNearbyStops) IsSet() bool {
	return v.isSet
}

func (v *NullableApiNearbyStops) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiNearbyStops(val *ApiNearbyStops) *NullableApiNearbyStops {
	return &NullableApiNearbyStops{value: val, isSet: true}
}

func (v NullableApiNearbyStops) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiNearbyStops) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



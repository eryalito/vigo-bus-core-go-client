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

// checks if the ApiNearbyStopsOrigin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiNearbyStopsOrigin{}

// ApiNearbyStopsOrigin struct for ApiNearbyStopsOrigin
type ApiNearbyStopsOrigin struct {
	Lat *float32 `json:"lat,omitempty"`
	Lon *float32 `json:"lon,omitempty"`
}

// NewApiNearbyStopsOrigin instantiates a new ApiNearbyStopsOrigin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiNearbyStopsOrigin() *ApiNearbyStopsOrigin {
	this := ApiNearbyStopsOrigin{}
	return &this
}

// NewApiNearbyStopsOriginWithDefaults instantiates a new ApiNearbyStopsOrigin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiNearbyStopsOriginWithDefaults() *ApiNearbyStopsOrigin {
	this := ApiNearbyStopsOrigin{}
	return &this
}

// GetLat returns the Lat field value if set, zero value otherwise.
func (o *ApiNearbyStopsOrigin) GetLat() float32 {
	if o == nil || IsNil(o.Lat) {
		var ret float32
		return ret
	}
	return *o.Lat
}

// GetLatOk returns a tuple with the Lat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNearbyStopsOrigin) GetLatOk() (*float32, bool) {
	if o == nil || IsNil(o.Lat) {
		return nil, false
	}
	return o.Lat, true
}

// HasLat returns a boolean if a field has been set.
func (o *ApiNearbyStopsOrigin) HasLat() bool {
	if o != nil && !IsNil(o.Lat) {
		return true
	}

	return false
}

// SetLat gets a reference to the given float32 and assigns it to the Lat field.
func (o *ApiNearbyStopsOrigin) SetLat(v float32) {
	o.Lat = &v
}

// GetLon returns the Lon field value if set, zero value otherwise.
func (o *ApiNearbyStopsOrigin) GetLon() float32 {
	if o == nil || IsNil(o.Lon) {
		var ret float32
		return ret
	}
	return *o.Lon
}

// GetLonOk returns a tuple with the Lon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNearbyStopsOrigin) GetLonOk() (*float32, bool) {
	if o == nil || IsNil(o.Lon) {
		return nil, false
	}
	return o.Lon, true
}

// HasLon returns a boolean if a field has been set.
func (o *ApiNearbyStopsOrigin) HasLon() bool {
	if o != nil && !IsNil(o.Lon) {
		return true
	}

	return false
}

// SetLon gets a reference to the given float32 and assigns it to the Lon field.
func (o *ApiNearbyStopsOrigin) SetLon(v float32) {
	o.Lon = &v
}

func (o ApiNearbyStopsOrigin) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiNearbyStopsOrigin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Lat) {
		toSerialize["lat"] = o.Lat
	}
	if !IsNil(o.Lon) {
		toSerialize["lon"] = o.Lon
	}
	return toSerialize, nil
}

type NullableApiNearbyStopsOrigin struct {
	value *ApiNearbyStopsOrigin
	isSet bool
}

func (v NullableApiNearbyStopsOrigin) Get() *ApiNearbyStopsOrigin {
	return v.value
}

func (v *NullableApiNearbyStopsOrigin) Set(val *ApiNearbyStopsOrigin) {
	v.value = val
	v.isSet = true
}

func (v NullableApiNearbyStopsOrigin) IsSet() bool {
	return v.isSet
}

func (v *NullableApiNearbyStopsOrigin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiNearbyStopsOrigin(val *ApiNearbyStopsOrigin) *NullableApiNearbyStopsOrigin {
	return &NullableApiNearbyStopsOrigin{value: val, isSet: true}
}

func (v NullableApiNearbyStopsOrigin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiNearbyStopsOrigin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


# ApiNearbyStops

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | Pointer to **string** |  | [optional] 
**Origin** | Pointer to [**ApiNearbyStopsOrigin**](ApiNearbyStopsOrigin.md) |  | [optional] 
**Radius** | Pointer to **float32** |  | [optional] 
**Stops** | Pointer to [**[]ApiStop**](ApiStop.md) |  | [optional] 

## Methods

### NewApiNearbyStops

`func NewApiNearbyStops() *ApiNearbyStops`

NewApiNearbyStops instantiates a new ApiNearbyStops object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiNearbyStopsWithDefaults

`func NewApiNearbyStopsWithDefaults() *ApiNearbyStops`

NewApiNearbyStopsWithDefaults instantiates a new ApiNearbyStops object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *ApiNearbyStops) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ApiNearbyStops) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ApiNearbyStops) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *ApiNearbyStops) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetOrigin

`func (o *ApiNearbyStops) GetOrigin() ApiNearbyStopsOrigin`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *ApiNearbyStops) GetOriginOk() (*ApiNearbyStopsOrigin, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *ApiNearbyStops) SetOrigin(v ApiNearbyStopsOrigin)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *ApiNearbyStops) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetRadius

`func (o *ApiNearbyStops) GetRadius() float32`

GetRadius returns the Radius field if non-nil, zero value otherwise.

### GetRadiusOk

`func (o *ApiNearbyStops) GetRadiusOk() (*float32, bool)`

GetRadiusOk returns a tuple with the Radius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadius

`func (o *ApiNearbyStops) SetRadius(v float32)`

SetRadius sets Radius field to given value.

### HasRadius

`func (o *ApiNearbyStops) HasRadius() bool`

HasRadius returns a boolean if a field has been set.

### GetStops

`func (o *ApiNearbyStops) GetStops() []ApiStop`

GetStops returns the Stops field if non-nil, zero value otherwise.

### GetStopsOk

`func (o *ApiNearbyStops) GetStopsOk() (*[]ApiStop, bool)`

GetStopsOk returns a tuple with the Stops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStops

`func (o *ApiNearbyStops) SetStops(v []ApiStop)`

SetStops sets Stops field to given value.

### HasStops

`func (o *ApiNearbyStops) HasStops() bool`

HasStops returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



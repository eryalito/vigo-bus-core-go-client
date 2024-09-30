# ApiIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FavoriteStops** | Pointer to [**[]ApiStop**](ApiStop.md) | FavoriteStops is a list of the user&#39;s favorite stops | [optional] 
**Id** | Pointer to **int32** | ID is the unique identifier of the identity | [optional] 
**Metadata** | Pointer to **string** | Metadata is a genric string that holds additional information about the identity | [optional] 
**Provider** | Pointer to [**ApiProviderType**](ApiProviderType.md) | Provider is the type of the identity provider | [optional] 
**Uuid** | Pointer to **string** | UUID is the unique identifier of the identity, usually provided by the auth provider | [optional] 

## Methods

### NewApiIdentity

`func NewApiIdentity() *ApiIdentity`

NewApiIdentity instantiates a new ApiIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiIdentityWithDefaults

`func NewApiIdentityWithDefaults() *ApiIdentity`

NewApiIdentityWithDefaults instantiates a new ApiIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFavoriteStops

`func (o *ApiIdentity) GetFavoriteStops() []ApiStop`

GetFavoriteStops returns the FavoriteStops field if non-nil, zero value otherwise.

### GetFavoriteStopsOk

`func (o *ApiIdentity) GetFavoriteStopsOk() (*[]ApiStop, bool)`

GetFavoriteStopsOk returns a tuple with the FavoriteStops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavoriteStops

`func (o *ApiIdentity) SetFavoriteStops(v []ApiStop)`

SetFavoriteStops sets FavoriteStops field to given value.

### HasFavoriteStops

`func (o *ApiIdentity) HasFavoriteStops() bool`

HasFavoriteStops returns a boolean if a field has been set.

### GetId

`func (o *ApiIdentity) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiIdentity) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiIdentity) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ApiIdentity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *ApiIdentity) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ApiIdentity) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ApiIdentity) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ApiIdentity) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProvider

`func (o *ApiIdentity) GetProvider() ApiProviderType`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ApiIdentity) GetProviderOk() (*ApiProviderType, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ApiIdentity) SetProvider(v ApiProviderType)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ApiIdentity) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetUuid

`func (o *ApiIdentity) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ApiIdentity) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ApiIdentity) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ApiIdentity) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



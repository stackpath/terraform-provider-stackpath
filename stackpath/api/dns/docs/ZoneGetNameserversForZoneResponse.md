# ZoneGetNameserversForZoneResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configured** | Pointer to **bool** | Whether or not all required name servers are configured in the zone | [optional] 
**CurrentNameservers** | Pointer to **[]string** | The zone&#39;s currently configured nameservers | [optional] 
**RequiredNameservers** | Pointer to **[]string** | The nameservers required in the zone&#39;s configuration | [optional] 

## Methods

### NewZoneGetNameserversForZoneResponse

`func NewZoneGetNameserversForZoneResponse() *ZoneGetNameserversForZoneResponse`

NewZoneGetNameserversForZoneResponse instantiates a new ZoneGetNameserversForZoneResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneGetNameserversForZoneResponseWithDefaults

`func NewZoneGetNameserversForZoneResponseWithDefaults() *ZoneGetNameserversForZoneResponse`

NewZoneGetNameserversForZoneResponseWithDefaults instantiates a new ZoneGetNameserversForZoneResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigured

`func (o *ZoneGetNameserversForZoneResponse) GetConfigured() bool`

GetConfigured returns the Configured field if non-nil, zero value otherwise.

### GetConfiguredOk

`func (o *ZoneGetNameserversForZoneResponse) GetConfiguredOk() (*bool, bool)`

GetConfiguredOk returns a tuple with the Configured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigured

`func (o *ZoneGetNameserversForZoneResponse) SetConfigured(v bool)`

SetConfigured sets Configured field to given value.

### HasConfigured

`func (o *ZoneGetNameserversForZoneResponse) HasConfigured() bool`

HasConfigured returns a boolean if a field has been set.

### GetCurrentNameservers

`func (o *ZoneGetNameserversForZoneResponse) GetCurrentNameservers() []string`

GetCurrentNameservers returns the CurrentNameservers field if non-nil, zero value otherwise.

### GetCurrentNameserversOk

`func (o *ZoneGetNameserversForZoneResponse) GetCurrentNameserversOk() (*[]string, bool)`

GetCurrentNameserversOk returns a tuple with the CurrentNameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentNameservers

`func (o *ZoneGetNameserversForZoneResponse) SetCurrentNameservers(v []string)`

SetCurrentNameservers sets CurrentNameservers field to given value.

### HasCurrentNameservers

`func (o *ZoneGetNameserversForZoneResponse) HasCurrentNameservers() bool`

HasCurrentNameservers returns a boolean if a field has been set.

### GetRequiredNameservers

`func (o *ZoneGetNameserversForZoneResponse) GetRequiredNameservers() []string`

GetRequiredNameservers returns the RequiredNameservers field if non-nil, zero value otherwise.

### GetRequiredNameserversOk

`func (o *ZoneGetNameserversForZoneResponse) GetRequiredNameserversOk() (*[]string, bool)`

GetRequiredNameserversOk returns a tuple with the RequiredNameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredNameservers

`func (o *ZoneGetNameserversForZoneResponse) SetRequiredNameservers(v []string)`

SetRequiredNameservers sets RequiredNameservers field to given value.

### HasRequiredNameservers

`func (o *ZoneGetNameserversForZoneResponse) HasRequiredNameservers() bool`

HasRequiredNameservers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



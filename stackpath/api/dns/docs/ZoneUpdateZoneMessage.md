# ZoneUpdateZoneMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | Pointer to **map[string]string** | A key/value pair of user-defined labels for a DNS zone  Zone labels are not processed by StackPath and are solely used for users to organize their DNS zones. | [optional] 

## Methods

### NewZoneUpdateZoneMessage

`func NewZoneUpdateZoneMessage() *ZoneUpdateZoneMessage`

NewZoneUpdateZoneMessage instantiates a new ZoneUpdateZoneMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneUpdateZoneMessageWithDefaults

`func NewZoneUpdateZoneMessageWithDefaults() *ZoneUpdateZoneMessage`

NewZoneUpdateZoneMessageWithDefaults instantiates a new ZoneUpdateZoneMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *ZoneUpdateZoneMessage) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ZoneUpdateZoneMessage) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ZoneUpdateZoneMessage) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ZoneUpdateZoneMessage) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



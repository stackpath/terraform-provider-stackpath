# ZoneCreateZoneMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** | The zone&#39;s domain name | [optional] 
**Labels** | Pointer to **map[string]string** | A key/value pair of user-defined labels for a DNS zone  Zone labels are not processed by StackPath and are solely used for users to organize their DNS zones. | [optional] 
**UseApexDomain** | Pointer to **bool** | Whether or not to create a zone for the apex domain only  If this is true and a domain with subdomains is provided, it will be stripped and only the root domain will be used for the zone. If this is false an error will be returned if it&#39;s not already an apex domain. | [optional] 

## Methods

### NewZoneCreateZoneMessage

`func NewZoneCreateZoneMessage() *ZoneCreateZoneMessage`

NewZoneCreateZoneMessage instantiates a new ZoneCreateZoneMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneCreateZoneMessageWithDefaults

`func NewZoneCreateZoneMessageWithDefaults() *ZoneCreateZoneMessage`

NewZoneCreateZoneMessageWithDefaults instantiates a new ZoneCreateZoneMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *ZoneCreateZoneMessage) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ZoneCreateZoneMessage) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ZoneCreateZoneMessage) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ZoneCreateZoneMessage) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetLabels

`func (o *ZoneCreateZoneMessage) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ZoneCreateZoneMessage) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ZoneCreateZoneMessage) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ZoneCreateZoneMessage) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetUseApexDomain

`func (o *ZoneCreateZoneMessage) GetUseApexDomain() bool`

GetUseApexDomain returns the UseApexDomain field if non-nil, zero value otherwise.

### GetUseApexDomainOk

`func (o *ZoneCreateZoneMessage) GetUseApexDomainOk() (*bool, bool)`

GetUseApexDomainOk returns a tuple with the UseApexDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseApexDomain

`func (o *ZoneCreateZoneMessage) SetUseApexDomain(v bool)`

SetUseApexDomain sets UseApexDomain field to given value.

### HasUseApexDomain

`func (o *ZoneCreateZoneMessage) HasUseApexDomain() bool`

HasUseApexDomain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



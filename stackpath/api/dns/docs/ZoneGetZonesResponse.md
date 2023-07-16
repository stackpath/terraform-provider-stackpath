# ZoneGetZonesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageInfo** | Pointer to [**PaginationPageInfo**](PaginationPageInfo.md) |  | [optional] 
**Zones** | Pointer to [**[]ZoneZone**](ZoneZone.md) | The requested DNS zones | [optional] 

## Methods

### NewZoneGetZonesResponse

`func NewZoneGetZonesResponse() *ZoneGetZonesResponse`

NewZoneGetZonesResponse instantiates a new ZoneGetZonesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneGetZonesResponseWithDefaults

`func NewZoneGetZonesResponseWithDefaults() *ZoneGetZonesResponse`

NewZoneGetZonesResponseWithDefaults instantiates a new ZoneGetZonesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageInfo

`func (o *ZoneGetZonesResponse) GetPageInfo() PaginationPageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *ZoneGetZonesResponse) GetPageInfoOk() (*PaginationPageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *ZoneGetZonesResponse) SetPageInfo(v PaginationPageInfo)`

SetPageInfo sets PageInfo field to given value.

### HasPageInfo

`func (o *ZoneGetZonesResponse) HasPageInfo() bool`

HasPageInfo returns a boolean if a field has been set.

### GetZones

`func (o *ZoneGetZonesResponse) GetZones() []ZoneZone`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *ZoneGetZonesResponse) GetZonesOk() (*[]ZoneZone, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *ZoneGetZonesResponse) SetZones(v []ZoneZone)`

SetZones sets Zones field to given value.

### HasZones

`func (o *ZoneGetZonesResponse) HasZones() bool`

HasZones returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



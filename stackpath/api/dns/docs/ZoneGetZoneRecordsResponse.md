# ZoneGetZoneRecordsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageInfo** | Pointer to [**PaginationPageInfo**](PaginationPageInfo.md) |  | [optional] 
**Records** | Pointer to [**[]ZoneZoneRecord**](ZoneZoneRecord.md) | The requested resource records | [optional] 

## Methods

### NewZoneGetZoneRecordsResponse

`func NewZoneGetZoneRecordsResponse() *ZoneGetZoneRecordsResponse`

NewZoneGetZoneRecordsResponse instantiates a new ZoneGetZoneRecordsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneGetZoneRecordsResponseWithDefaults

`func NewZoneGetZoneRecordsResponseWithDefaults() *ZoneGetZoneRecordsResponse`

NewZoneGetZoneRecordsResponseWithDefaults instantiates a new ZoneGetZoneRecordsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageInfo

`func (o *ZoneGetZoneRecordsResponse) GetPageInfo() PaginationPageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *ZoneGetZoneRecordsResponse) GetPageInfoOk() (*PaginationPageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *ZoneGetZoneRecordsResponse) SetPageInfo(v PaginationPageInfo)`

SetPageInfo sets PageInfo field to given value.

### HasPageInfo

`func (o *ZoneGetZoneRecordsResponse) HasPageInfo() bool`

HasPageInfo returns a boolean if a field has been set.

### GetRecords

`func (o *ZoneGetZoneRecordsResponse) GetRecords() []ZoneZoneRecord`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *ZoneGetZoneRecordsResponse) GetRecordsOk() (*[]ZoneZoneRecord, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *ZoneGetZoneRecordsResponse) SetRecords(v []ZoneZoneRecord)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *ZoneGetZoneRecordsResponse) HasRecords() bool`

HasRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



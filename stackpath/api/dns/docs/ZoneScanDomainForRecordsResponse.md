# ZoneScanDomainForRecordsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Records** | Pointer to [**[]ZoneImportZoneRecord**](ZoneImportZoneRecord.md) | The resource records StackPath was able to scan from the domain | [optional] 

## Methods

### NewZoneScanDomainForRecordsResponse

`func NewZoneScanDomainForRecordsResponse() *ZoneScanDomainForRecordsResponse`

NewZoneScanDomainForRecordsResponse instantiates a new ZoneScanDomainForRecordsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneScanDomainForRecordsResponseWithDefaults

`func NewZoneScanDomainForRecordsResponseWithDefaults() *ZoneScanDomainForRecordsResponse`

NewZoneScanDomainForRecordsResponseWithDefaults instantiates a new ZoneScanDomainForRecordsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecords

`func (o *ZoneScanDomainForRecordsResponse) GetRecords() []ZoneImportZoneRecord`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *ZoneScanDomainForRecordsResponse) GetRecordsOk() (*[]ZoneImportZoneRecord, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *ZoneScanDomainForRecordsResponse) SetRecords(v []ZoneImportZoneRecord)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *ZoneScanDomainForRecordsResponse) HasRecords() bool`

HasRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



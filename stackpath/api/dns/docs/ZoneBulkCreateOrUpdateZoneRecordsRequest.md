# ZoneBulkCreateOrUpdateZoneRecordsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Records** | Pointer to [**[]ZoneBulkCreateOrUpdateZoneRecordMessage**](ZoneBulkCreateOrUpdateZoneRecordMessage.md) | The records to create or update in the DNS zone | [optional] 

## Methods

### NewZoneBulkCreateOrUpdateZoneRecordsRequest

`func NewZoneBulkCreateOrUpdateZoneRecordsRequest() *ZoneBulkCreateOrUpdateZoneRecordsRequest`

NewZoneBulkCreateOrUpdateZoneRecordsRequest instantiates a new ZoneBulkCreateOrUpdateZoneRecordsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneBulkCreateOrUpdateZoneRecordsRequestWithDefaults

`func NewZoneBulkCreateOrUpdateZoneRecordsRequestWithDefaults() *ZoneBulkCreateOrUpdateZoneRecordsRequest`

NewZoneBulkCreateOrUpdateZoneRecordsRequestWithDefaults instantiates a new ZoneBulkCreateOrUpdateZoneRecordsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecords

`func (o *ZoneBulkCreateOrUpdateZoneRecordsRequest) GetRecords() []ZoneBulkCreateOrUpdateZoneRecordMessage`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *ZoneBulkCreateOrUpdateZoneRecordsRequest) GetRecordsOk() (*[]ZoneBulkCreateOrUpdateZoneRecordMessage, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *ZoneBulkCreateOrUpdateZoneRecordsRequest) SetRecords(v []ZoneBulkCreateOrUpdateZoneRecordMessage)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *ZoneBulkCreateOrUpdateZoneRecordsRequest) HasRecords() bool`

HasRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



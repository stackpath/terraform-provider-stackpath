# ZonePatchZoneRecordMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A zone record&#39;s name | [optional] 
**Type** | Pointer to [**ZoneRecordType**](ZoneRecordType.md) |  | [optional] [default to EMPTY]
**Ttl** | Pointer to **int32** | A zone record&#39;s time to live  A record&#39;s TTL is the number of seconds that the record should be cached by DNS resolvers. Use lower TTL values if you expect zone records to change often. Use higher TTL values for records that won&#39;t change to prevent extra DNS lookups by clients. | [optional] 
**Data** | Pointer to **string** | A zone record&#39;s value  MX record data follows the format &#x60;&lt;priority&gt; &lt;value&gt;&#x60;, without &#x60;&lt;&gt;&#x60;s. | [optional] 
**Weight** | Pointer to **int32** | A zone record&#39;s priority  A resource record is replicated in StackPath&#39;s DNS infrastructure the number of times of the record&#39;s weight, giving it a more likely response to queries if a zone has records with the same name and type. | [optional] 

## Methods

### NewZonePatchZoneRecordMessage

`func NewZonePatchZoneRecordMessage() *ZonePatchZoneRecordMessage`

NewZonePatchZoneRecordMessage instantiates a new ZonePatchZoneRecordMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZonePatchZoneRecordMessageWithDefaults

`func NewZonePatchZoneRecordMessageWithDefaults() *ZonePatchZoneRecordMessage`

NewZonePatchZoneRecordMessageWithDefaults instantiates a new ZonePatchZoneRecordMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ZonePatchZoneRecordMessage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ZonePatchZoneRecordMessage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ZonePatchZoneRecordMessage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ZonePatchZoneRecordMessage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ZonePatchZoneRecordMessage) GetType() ZoneRecordType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ZonePatchZoneRecordMessage) GetTypeOk() (*ZoneRecordType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ZonePatchZoneRecordMessage) SetType(v ZoneRecordType)`

SetType sets Type field to given value.

### HasType

`func (o *ZonePatchZoneRecordMessage) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTtl

`func (o *ZonePatchZoneRecordMessage) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *ZonePatchZoneRecordMessage) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *ZonePatchZoneRecordMessage) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *ZonePatchZoneRecordMessage) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetData

`func (o *ZonePatchZoneRecordMessage) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ZonePatchZoneRecordMessage) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ZonePatchZoneRecordMessage) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *ZonePatchZoneRecordMessage) HasData() bool`

HasData returns a boolean if a field has been set.

### GetWeight

`func (o *ZonePatchZoneRecordMessage) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ZonePatchZoneRecordMessage) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ZonePatchZoneRecordMessage) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *ZonePatchZoneRecordMessage) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



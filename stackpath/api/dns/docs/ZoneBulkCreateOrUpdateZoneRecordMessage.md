# ZoneBulkCreateOrUpdateZoneRecordMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A zone record&#39;s name | [optional] 
**Type** | Pointer to [**ZoneRecordType**](ZoneRecordType.md) |  | [optional] [default to EMPTY]
**Ttl** | Pointer to **int32** | A zone record&#39;s time to live  A record&#39;s TTL is the number of seconds that the record should be cached by DNS resolvers. Use lower TTL values if you expect zone records to change often. Use higher TTL values for records that won&#39;t change to prevent extra DNS lookups by clients. | [optional] 
**Data** | Pointer to **string** | A zone record&#39;s value  MX record data follows the format &#x60;&lt;priority&gt; &lt;value&gt;&#x60;, without &#x60;&lt;&gt;&#x60;s. | [optional] 
**Weight** | Pointer to **int32** | A zone record&#39;s priority  A resource record is replicated in StackPath&#39;s DNS infrastructure the number of times of the record&#39;s weight, giving it a more likely response to queries if a zone has records with the same name and type. | [optional] 
**Labels** | Pointer to **map[string]string** | A key/value pair of user-defined labels for a DNS zone record  Zone record labels are not processed by StackPath and are solely used for users to organize their DNS zones. | [optional] 
**Id** | Pointer to **string** | The ID of the resource record to update  Do not provide an ID to create a new resource record | [optional] 

## Methods

### NewZoneBulkCreateOrUpdateZoneRecordMessage

`func NewZoneBulkCreateOrUpdateZoneRecordMessage() *ZoneBulkCreateOrUpdateZoneRecordMessage`

NewZoneBulkCreateOrUpdateZoneRecordMessage instantiates a new ZoneBulkCreateOrUpdateZoneRecordMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneBulkCreateOrUpdateZoneRecordMessageWithDefaults

`func NewZoneBulkCreateOrUpdateZoneRecordMessageWithDefaults() *ZoneBulkCreateOrUpdateZoneRecordMessage`

NewZoneBulkCreateOrUpdateZoneRecordMessageWithDefaults instantiates a new ZoneBulkCreateOrUpdateZoneRecordMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ZoneBulkCreateOrUpdateZoneRecordMessage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ZoneBulkCreateOrUpdateZoneRecordMessage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ZoneBulkCreateOrUpdateZoneRecordMessage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ZoneBulkCreateOrUpdateZoneRecordMessage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ZoneBulkCreateOrUpdateZoneRecordMessage) GetType() ZoneRecordType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ZoneBulkCreateOrUpdateZoneRecordMessage) GetTypeOk() (*ZoneRecordType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ZoneBulkCreateOrUpdateZoneRecordMessage) SetType(v ZoneRecordType)`

SetType sets Type field to given value.

### HasType

`func (o *ZoneBulkCreateOrUpdateZoneRecordMessage) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTtl

`func (o *ZoneBulkCreateOrUpdateZoneRecordMessage) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *ZoneBulkCreateOrUpdateZoneRecordMessage) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *ZoneBulkCreateOrUpdateZoneRecordMessage) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *ZoneBulkCreateOrUpdateZoneRecordMessage) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetData

`func (o *ZoneBulkCreateOrUpdateZoneRecordMessage) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ZoneBulkCreateOrUpdateZoneRecordMessage) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ZoneBulkCreateOrUpdateZoneRecordMessage) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *ZoneBulkCreateOrUpdateZoneRecordMessage) HasData() bool`

HasData returns a boolean if a field has been set.

### GetWeight

`func (o *ZoneBulkCreateOrUpdateZoneRecordMessage) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ZoneBulkCreateOrUpdateZoneRecordMessage) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ZoneBulkCreateOrUpdateZoneRecordMessage) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *ZoneBulkCreateOrUpdateZoneRecordMessage) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetLabels

`func (o *ZoneBulkCreateOrUpdateZoneRecordMessage) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ZoneBulkCreateOrUpdateZoneRecordMessage) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ZoneBulkCreateOrUpdateZoneRecordMessage) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ZoneBulkCreateOrUpdateZoneRecordMessage) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetId

`func (o *ZoneBulkCreateOrUpdateZoneRecordMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ZoneBulkCreateOrUpdateZoneRecordMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ZoneBulkCreateOrUpdateZoneRecordMessage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ZoneBulkCreateOrUpdateZoneRecordMessage) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



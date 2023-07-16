# ZoneZoneRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A zone record&#39;s unique ID | [optional] 
**ZoneId** | Pointer to **string** | The ID of the zone that a zone record belongs to | [optional] 
**Name** | Pointer to **string** | A zone record&#39;s name  Use the value \&quot;@\&quot; to denote current root domain name. | [optional] 
**Type** | Pointer to **string** | A zone record&#39;s type  Zone record types describe the zone record&#39;s behavior. For instance, a zone record&#39;s type can say that the record is a name to IP address value, a name alias, or which mail exchanger is responsible for the domain. See https://support.stackpath.com/hc/en-us/articles/360001085563-What-DNS-record-types-does-StackPath-support for more information. | [optional] 
**Class** | Pointer to **string** | A zone record&#39;s class code  This is typically \&quot;IN\&quot; for Internet related resource records. | [optional] 
**Ttl** | Pointer to **int32** | A zone record&#39;s time to live  A record&#39;s TTL is the number of seconds that the record should be cached by DNS resolvers. Use lower TTL values if you expect zone records to change often. Use higher TTL values for records that won&#39;t change to prevent extra DNS lookups by clients. | [optional] 
**Data** | Pointer to **string** | A zone record&#39;s value  MX record data follows the format &#x60;&lt;priority&gt; &lt;value&gt;&#x60;, without &#x60;&lt;&gt;&#x60;s. | [optional] 
**Weight** | Pointer to **int32** | A zone record&#39;s priority  A resource record is replicated in StackPath&#39;s DNS infrastructure the number of times of the record&#39;s weight, giving it a more likely response to queries if a zone has records with the same name and type. | [optional] 
**Labels** | Pointer to **map[string]string** | A key/value pair of user-defined labels for a zone record  Zone record labels are not processed by StackPath and are solely used for users to organize their DNS zones. | [optional] 
**Created** | Pointer to **time.Time** | The date a zone record was created | [optional] 
**Updated** | Pointer to **time.Time** | The date a zone record was last updated | [optional] 

## Methods

### NewZoneZoneRecord

`func NewZoneZoneRecord() *ZoneZoneRecord`

NewZoneZoneRecord instantiates a new ZoneZoneRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneZoneRecordWithDefaults

`func NewZoneZoneRecordWithDefaults() *ZoneZoneRecord`

NewZoneZoneRecordWithDefaults instantiates a new ZoneZoneRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ZoneZoneRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ZoneZoneRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ZoneZoneRecord) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ZoneZoneRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetZoneId

`func (o *ZoneZoneRecord) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *ZoneZoneRecord) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *ZoneZoneRecord) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *ZoneZoneRecord) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetName

`func (o *ZoneZoneRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ZoneZoneRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ZoneZoneRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ZoneZoneRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ZoneZoneRecord) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ZoneZoneRecord) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ZoneZoneRecord) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ZoneZoneRecord) HasType() bool`

HasType returns a boolean if a field has been set.

### GetClass

`func (o *ZoneZoneRecord) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *ZoneZoneRecord) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *ZoneZoneRecord) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *ZoneZoneRecord) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetTtl

`func (o *ZoneZoneRecord) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *ZoneZoneRecord) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *ZoneZoneRecord) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *ZoneZoneRecord) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetData

`func (o *ZoneZoneRecord) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ZoneZoneRecord) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ZoneZoneRecord) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *ZoneZoneRecord) HasData() bool`

HasData returns a boolean if a field has been set.

### GetWeight

`func (o *ZoneZoneRecord) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ZoneZoneRecord) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ZoneZoneRecord) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *ZoneZoneRecord) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetLabels

`func (o *ZoneZoneRecord) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ZoneZoneRecord) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ZoneZoneRecord) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ZoneZoneRecord) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetCreated

`func (o *ZoneZoneRecord) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ZoneZoneRecord) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ZoneZoneRecord) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ZoneZoneRecord) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *ZoneZoneRecord) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ZoneZoneRecord) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ZoneZoneRecord) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ZoneZoneRecord) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



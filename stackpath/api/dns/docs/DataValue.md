# DataValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnixTime** | Pointer to **string** | The time that a data point was recorded | [optional] 
**Value** | Pointer to **string** | A data point&#39;s value | [optional] 

## Methods

### NewDataValue

`func NewDataValue() *DataValue`

NewDataValue instantiates a new DataValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataValueWithDefaults

`func NewDataValueWithDefaults() *DataValue`

NewDataValueWithDefaults instantiates a new DataValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnixTime

`func (o *DataValue) GetUnixTime() string`

GetUnixTime returns the UnixTime field if non-nil, zero value otherwise.

### GetUnixTimeOk

`func (o *DataValue) GetUnixTimeOk() (*string, bool)`

GetUnixTimeOk returns a tuple with the UnixTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnixTime

`func (o *DataValue) SetUnixTime(v string)`

SetUnixTime sets UnixTime field to given value.

### HasUnixTime

`func (o *DataValue) HasUnixTime() bool`

HasUnixTime returns a boolean if a field has been set.

### GetValue

`func (o *DataValue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DataValue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DataValue) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *DataValue) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



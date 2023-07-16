# DataVectorResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | Pointer to **map[string]string** | The data points&#39; labels | [optional] 
**Value** | Pointer to [**DataValue**](DataValue.md) |  | [optional] 

## Methods

### NewDataVectorResult

`func NewDataVectorResult() *DataVectorResult`

NewDataVectorResult instantiates a new DataVectorResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataVectorResultWithDefaults

`func NewDataVectorResultWithDefaults() *DataVectorResult`

NewDataVectorResultWithDefaults instantiates a new DataVectorResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *DataVectorResult) GetMetric() map[string]string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *DataVectorResult) GetMetricOk() (*map[string]string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *DataVectorResult) SetMetric(v map[string]string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *DataVectorResult) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetValue

`func (o *DataVectorResult) GetValue() DataValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DataVectorResult) GetValueOk() (*DataValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DataVectorResult) SetValue(v DataValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *DataVectorResult) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



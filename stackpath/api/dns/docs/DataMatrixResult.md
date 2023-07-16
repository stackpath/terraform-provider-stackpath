# DataMatrixResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | Pointer to **map[string]string** | The data points&#39; labels | [optional] 
**Values** | Pointer to [**[]DataValue**](DataValue.md) | Time series data point values | [optional] 

## Methods

### NewDataMatrixResult

`func NewDataMatrixResult() *DataMatrixResult`

NewDataMatrixResult instantiates a new DataMatrixResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataMatrixResultWithDefaults

`func NewDataMatrixResultWithDefaults() *DataMatrixResult`

NewDataMatrixResultWithDefaults instantiates a new DataMatrixResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *DataMatrixResult) GetMetric() map[string]string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *DataMatrixResult) GetMetricOk() (*map[string]string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *DataMatrixResult) SetMetric(v map[string]string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *DataMatrixResult) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetValues

`func (o *DataMatrixResult) GetValues() []DataValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *DataMatrixResult) GetValuesOk() (*[]DataValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *DataMatrixResult) SetValues(v []DataValue)`

SetValues sets Values field to given value.

### HasValues

`func (o *DataMatrixResult) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



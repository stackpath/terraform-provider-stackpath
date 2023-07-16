# MetricsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Matrix** | Pointer to [**DataMatrix**](DataMatrix.md) |  | [optional] 
**Vector** | Pointer to [**DataVector**](DataVector.md) |  | [optional] 

## Methods

### NewMetricsData

`func NewMetricsData() *MetricsData`

NewMetricsData instantiates a new MetricsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsDataWithDefaults

`func NewMetricsDataWithDefaults() *MetricsData`

NewMetricsDataWithDefaults instantiates a new MetricsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatrix

`func (o *MetricsData) GetMatrix() DataMatrix`

GetMatrix returns the Matrix field if non-nil, zero value otherwise.

### GetMatrixOk

`func (o *MetricsData) GetMatrixOk() (*DataMatrix, bool)`

GetMatrixOk returns a tuple with the Matrix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatrix

`func (o *MetricsData) SetMatrix(v DataMatrix)`

SetMatrix sets Matrix field to given value.

### HasMatrix

`func (o *MetricsData) HasMatrix() bool`

HasMatrix returns a boolean if a field has been set.

### GetVector

`func (o *MetricsData) GetVector() DataVector`

GetVector returns the Vector field if non-nil, zero value otherwise.

### GetVectorOk

`func (o *MetricsData) GetVectorOk() (*DataVector, bool)`

GetVectorOk returns a tuple with the Vector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVector

`func (o *MetricsData) SetVector(v DataVector)`

SetVector sets Vector field to given value.

### HasVector

`func (o *MetricsData) HasVector() bool`

HasVector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



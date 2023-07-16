# DataMatrix

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]DataMatrixResult**](DataMatrixResult.md) | A data point&#39;s value | [optional] 

## Methods

### NewDataMatrix

`func NewDataMatrix() *DataMatrix`

NewDataMatrix instantiates a new DataMatrix object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataMatrixWithDefaults

`func NewDataMatrixWithDefaults() *DataMatrix`

NewDataMatrixWithDefaults instantiates a new DataMatrix object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *DataMatrix) GetResults() []DataMatrixResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *DataMatrix) GetResultsOk() (*[]DataMatrixResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *DataMatrix) SetResults(v []DataMatrixResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *DataMatrix) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



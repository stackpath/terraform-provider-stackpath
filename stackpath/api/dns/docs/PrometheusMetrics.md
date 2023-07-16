# PrometheusMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**PrometheusMetricsStatus**](PrometheusMetricsStatus.md) |  | [optional] [default to SUCCESS]
**Data** | Pointer to [**MetricsData**](MetricsData.md) |  | [optional] 
**ErrorType** | Pointer to **string** | The type of error encountered when querying for metrics | [optional] 
**Error** | Pointer to **string** | The error encountered when querying for metrics | [optional] 
**Warnings** | Pointer to **[]string** | Warnings encountered when querying for metrics | [optional] 

## Methods

### NewPrometheusMetrics

`func NewPrometheusMetrics() *PrometheusMetrics`

NewPrometheusMetrics instantiates a new PrometheusMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrometheusMetricsWithDefaults

`func NewPrometheusMetricsWithDefaults() *PrometheusMetrics`

NewPrometheusMetricsWithDefaults instantiates a new PrometheusMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *PrometheusMetrics) GetStatus() PrometheusMetricsStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PrometheusMetrics) GetStatusOk() (*PrometheusMetricsStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PrometheusMetrics) SetStatus(v PrometheusMetricsStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PrometheusMetrics) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *PrometheusMetrics) GetData() MetricsData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PrometheusMetrics) GetDataOk() (*MetricsData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PrometheusMetrics) SetData(v MetricsData)`

SetData sets Data field to given value.

### HasData

`func (o *PrometheusMetrics) HasData() bool`

HasData returns a boolean if a field has been set.

### GetErrorType

`func (o *PrometheusMetrics) GetErrorType() string`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *PrometheusMetrics) GetErrorTypeOk() (*string, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *PrometheusMetrics) SetErrorType(v string)`

SetErrorType sets ErrorType field to given value.

### HasErrorType

`func (o *PrometheusMetrics) HasErrorType() bool`

HasErrorType returns a boolean if a field has been set.

### GetError

`func (o *PrometheusMetrics) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *PrometheusMetrics) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *PrometheusMetrics) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *PrometheusMetrics) HasError() bool`

HasError returns a boolean if a field has been set.

### GetWarnings

`func (o *PrometheusMetrics) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *PrometheusMetrics) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *PrometheusMetrics) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *PrometheusMetrics) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# PaginationPageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **string** | The total number of items in the dataset | [optional] 
**HasPreviousPage** | Pointer to **bool** | Whether or not a previous page of data exists | [optional] 
**HasNextPage** | Pointer to **bool** | Whether or not another page of data is available | [optional] 
**StartCursor** | Pointer to **string** | The cursor for the first item in the set of data returned | [optional] 
**EndCursor** | Pointer to **string** | The cursor for the last item in the set of data returned | [optional] 

## Methods

### NewPaginationPageInfo

`func NewPaginationPageInfo() *PaginationPageInfo`

NewPaginationPageInfo instantiates a new PaginationPageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationPageInfoWithDefaults

`func NewPaginationPageInfoWithDefaults() *PaginationPageInfo`

NewPaginationPageInfoWithDefaults instantiates a new PaginationPageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *PaginationPageInfo) GetTotalCount() string`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PaginationPageInfo) GetTotalCountOk() (*string, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PaginationPageInfo) SetTotalCount(v string)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *PaginationPageInfo) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetHasPreviousPage

`func (o *PaginationPageInfo) GetHasPreviousPage() bool`

GetHasPreviousPage returns the HasPreviousPage field if non-nil, zero value otherwise.

### GetHasPreviousPageOk

`func (o *PaginationPageInfo) GetHasPreviousPageOk() (*bool, bool)`

GetHasPreviousPageOk returns a tuple with the HasPreviousPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPreviousPage

`func (o *PaginationPageInfo) SetHasPreviousPage(v bool)`

SetHasPreviousPage sets HasPreviousPage field to given value.

### HasHasPreviousPage

`func (o *PaginationPageInfo) HasHasPreviousPage() bool`

HasHasPreviousPage returns a boolean if a field has been set.

### GetHasNextPage

`func (o *PaginationPageInfo) GetHasNextPage() bool`

GetHasNextPage returns the HasNextPage field if non-nil, zero value otherwise.

### GetHasNextPageOk

`func (o *PaginationPageInfo) GetHasNextPageOk() (*bool, bool)`

GetHasNextPageOk returns a tuple with the HasNextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNextPage

`func (o *PaginationPageInfo) SetHasNextPage(v bool)`

SetHasNextPage sets HasNextPage field to given value.

### HasHasNextPage

`func (o *PaginationPageInfo) HasHasNextPage() bool`

HasHasNextPage returns a boolean if a field has been set.

### GetStartCursor

`func (o *PaginationPageInfo) GetStartCursor() string`

GetStartCursor returns the StartCursor field if non-nil, zero value otherwise.

### GetStartCursorOk

`func (o *PaginationPageInfo) GetStartCursorOk() (*string, bool)`

GetStartCursorOk returns a tuple with the StartCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartCursor

`func (o *PaginationPageInfo) SetStartCursor(v string)`

SetStartCursor sets StartCursor field to given value.

### HasStartCursor

`func (o *PaginationPageInfo) HasStartCursor() bool`

HasStartCursor returns a boolean if a field has been set.

### GetEndCursor

`func (o *PaginationPageInfo) GetEndCursor() string`

GetEndCursor returns the EndCursor field if non-nil, zero value otherwise.

### GetEndCursorOk

`func (o *PaginationPageInfo) GetEndCursorOk() (*string, bool)`

GetEndCursorOk returns a tuple with the EndCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndCursor

`func (o *PaginationPageInfo) SetEndCursor(v string)`

SetEndCursor sets EndCursor field to given value.

### HasEndCursor

`func (o *PaginationPageInfo) HasEndCursor() bool`

HasEndCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



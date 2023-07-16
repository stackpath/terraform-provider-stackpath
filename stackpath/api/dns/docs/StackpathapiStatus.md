# StackpathapiStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Details** | Pointer to [**[]ApiStatusDetail**](ApiStatusDetail.md) |  | [optional] 

## Methods

### NewStackpathapiStatus

`func NewStackpathapiStatus() *StackpathapiStatus`

NewStackpathapiStatus instantiates a new StackpathapiStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackpathapiStatusWithDefaults

`func NewStackpathapiStatusWithDefaults() *StackpathapiStatus`

NewStackpathapiStatusWithDefaults instantiates a new StackpathapiStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *StackpathapiStatus) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *StackpathapiStatus) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *StackpathapiStatus) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *StackpathapiStatus) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *StackpathapiStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *StackpathapiStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *StackpathapiStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *StackpathapiStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDetails

`func (o *StackpathapiStatus) GetDetails() []ApiStatusDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *StackpathapiStatus) GetDetailsOk() (*[]ApiStatusDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *StackpathapiStatus) SetDetails(v []ApiStatusDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *StackpathapiStatus) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



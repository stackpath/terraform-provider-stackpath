# ZoneZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StackId** | Pointer to **string** | The ID of the stack that a zone belongs to | [optional] 
**AccountId** | Pointer to **string** | The ID of the StackPath account that owns a zone | [optional] 
**Id** | Pointer to **string** | A zone&#39;s unique ID | [optional] 
**Domain** | Pointer to **string** | A zone&#39;s name | [optional] 
**Version** | Pointer to **string** | A zone&#39;s version number  Version numbers are incremented automatically when a zone is updated | [optional] 
**Labels** | Pointer to **map[string]string** | A key/value pair of user-defined labels for a zone  Zone labels are not processed by StackPath and are solely used for users to organize their DNS zones. | [optional] 
**Created** | Pointer to **time.Time** | The date a zone was created | [optional] 
**Updated** | Pointer to **time.Time** | The date a zone was last updated | [optional] 
**Nameservers** | Pointer to **[]string** | The hostnames of the StackPath resolvers that host a zone  Every zone has multiple name servers assigned by StackPath upon creation for redundancy purposes. | [optional] 
**Verified** | Pointer to **time.Time** | The date a zone&#39;s nameservers were last audited by StackPath | [optional] 
**Status** | Pointer to [**ZoneZoneStatus**](ZoneZoneStatus.md) |  | [optional] [default to ACTIVE]
**Disabled** | Pointer to **bool** | Whether or not a zone has been disabled by the user | [optional] 

## Methods

### NewZoneZone

`func NewZoneZone() *ZoneZone`

NewZoneZone instantiates a new ZoneZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneZoneWithDefaults

`func NewZoneZoneWithDefaults() *ZoneZone`

NewZoneZoneWithDefaults instantiates a new ZoneZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStackId

`func (o *ZoneZone) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *ZoneZone) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *ZoneZone) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *ZoneZone) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetAccountId

`func (o *ZoneZone) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ZoneZone) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ZoneZone) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ZoneZone) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetId

`func (o *ZoneZone) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ZoneZone) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ZoneZone) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ZoneZone) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDomain

`func (o *ZoneZone) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ZoneZone) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ZoneZone) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ZoneZone) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetVersion

`func (o *ZoneZone) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ZoneZone) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ZoneZone) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ZoneZone) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetLabels

`func (o *ZoneZone) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ZoneZone) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ZoneZone) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ZoneZone) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetCreated

`func (o *ZoneZone) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ZoneZone) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ZoneZone) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ZoneZone) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *ZoneZone) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ZoneZone) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ZoneZone) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ZoneZone) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetNameservers

`func (o *ZoneZone) GetNameservers() []string`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *ZoneZone) GetNameserversOk() (*[]string, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *ZoneZone) SetNameservers(v []string)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *ZoneZone) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.

### GetVerified

`func (o *ZoneZone) GetVerified() time.Time`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *ZoneZone) GetVerifiedOk() (*time.Time, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *ZoneZone) SetVerified(v time.Time)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *ZoneZone) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetStatus

`func (o *ZoneZone) GetStatus() ZoneZoneStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ZoneZone) GetStatusOk() (*ZoneZoneStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ZoneZone) SetStatus(v ZoneZoneStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ZoneZone) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDisabled

`func (o *ZoneZone) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ZoneZone) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ZoneZone) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ZoneZone) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



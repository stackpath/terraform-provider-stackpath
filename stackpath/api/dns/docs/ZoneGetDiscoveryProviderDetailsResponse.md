# ZoneGetDiscoveryProviderDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsProvider** | Pointer to [**ZoneDnsProvider**](ZoneDnsProvider.md) |  | [optional] [default to GENERAL]
**RequiresAuthentication** | Pointer to **bool** | Whether or not the provider requires authentication to scan or modify the domain | [optional] 
**Nameservers** | Pointer to **[]string** | The domain&#39;s authoritative nameservers | [optional] 

## Methods

### NewZoneGetDiscoveryProviderDetailsResponse

`func NewZoneGetDiscoveryProviderDetailsResponse() *ZoneGetDiscoveryProviderDetailsResponse`

NewZoneGetDiscoveryProviderDetailsResponse instantiates a new ZoneGetDiscoveryProviderDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneGetDiscoveryProviderDetailsResponseWithDefaults

`func NewZoneGetDiscoveryProviderDetailsResponseWithDefaults() *ZoneGetDiscoveryProviderDetailsResponse`

NewZoneGetDiscoveryProviderDetailsResponseWithDefaults instantiates a new ZoneGetDiscoveryProviderDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsProvider

`func (o *ZoneGetDiscoveryProviderDetailsResponse) GetDnsProvider() ZoneDnsProvider`

GetDnsProvider returns the DnsProvider field if non-nil, zero value otherwise.

### GetDnsProviderOk

`func (o *ZoneGetDiscoveryProviderDetailsResponse) GetDnsProviderOk() (*ZoneDnsProvider, bool)`

GetDnsProviderOk returns a tuple with the DnsProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsProvider

`func (o *ZoneGetDiscoveryProviderDetailsResponse) SetDnsProvider(v ZoneDnsProvider)`

SetDnsProvider sets DnsProvider field to given value.

### HasDnsProvider

`func (o *ZoneGetDiscoveryProviderDetailsResponse) HasDnsProvider() bool`

HasDnsProvider returns a boolean if a field has been set.

### GetRequiresAuthentication

`func (o *ZoneGetDiscoveryProviderDetailsResponse) GetRequiresAuthentication() bool`

GetRequiresAuthentication returns the RequiresAuthentication field if non-nil, zero value otherwise.

### GetRequiresAuthenticationOk

`func (o *ZoneGetDiscoveryProviderDetailsResponse) GetRequiresAuthenticationOk() (*bool, bool)`

GetRequiresAuthenticationOk returns a tuple with the RequiresAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresAuthentication

`func (o *ZoneGetDiscoveryProviderDetailsResponse) SetRequiresAuthentication(v bool)`

SetRequiresAuthentication sets RequiresAuthentication field to given value.

### HasRequiresAuthentication

`func (o *ZoneGetDiscoveryProviderDetailsResponse) HasRequiresAuthentication() bool`

HasRequiresAuthentication returns a boolean if a field has been set.

### GetNameservers

`func (o *ZoneGetDiscoveryProviderDetailsResponse) GetNameservers() []string`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *ZoneGetDiscoveryProviderDetailsResponse) GetNameserversOk() (*[]string, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *ZoneGetDiscoveryProviderDetailsResponse) SetNameservers(v []string)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *ZoneGetDiscoveryProviderDetailsResponse) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



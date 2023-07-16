# ScanDomainForRecordsRequestProviderConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsProvider** | Pointer to [**ZoneDnsProvider**](ZoneDnsProvider.md) |  | [optional] [default to GENERAL]
**AuthenticationUser** | Pointer to **string** | The username required to authenticate with the DNS provider | [optional] 
**ApiKey** | Pointer to **string** | The API key or password required to authenticate with the DNS provider | [optional] 

## Methods

### NewScanDomainForRecordsRequestProviderConfig

`func NewScanDomainForRecordsRequestProviderConfig() *ScanDomainForRecordsRequestProviderConfig`

NewScanDomainForRecordsRequestProviderConfig instantiates a new ScanDomainForRecordsRequestProviderConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScanDomainForRecordsRequestProviderConfigWithDefaults

`func NewScanDomainForRecordsRequestProviderConfigWithDefaults() *ScanDomainForRecordsRequestProviderConfig`

NewScanDomainForRecordsRequestProviderConfigWithDefaults instantiates a new ScanDomainForRecordsRequestProviderConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsProvider

`func (o *ScanDomainForRecordsRequestProviderConfig) GetDnsProvider() ZoneDnsProvider`

GetDnsProvider returns the DnsProvider field if non-nil, zero value otherwise.

### GetDnsProviderOk

`func (o *ScanDomainForRecordsRequestProviderConfig) GetDnsProviderOk() (*ZoneDnsProvider, bool)`

GetDnsProviderOk returns a tuple with the DnsProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsProvider

`func (o *ScanDomainForRecordsRequestProviderConfig) SetDnsProvider(v ZoneDnsProvider)`

SetDnsProvider sets DnsProvider field to given value.

### HasDnsProvider

`func (o *ScanDomainForRecordsRequestProviderConfig) HasDnsProvider() bool`

HasDnsProvider returns a boolean if a field has been set.

### GetAuthenticationUser

`func (o *ScanDomainForRecordsRequestProviderConfig) GetAuthenticationUser() string`

GetAuthenticationUser returns the AuthenticationUser field if non-nil, zero value otherwise.

### GetAuthenticationUserOk

`func (o *ScanDomainForRecordsRequestProviderConfig) GetAuthenticationUserOk() (*string, bool)`

GetAuthenticationUserOk returns a tuple with the AuthenticationUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationUser

`func (o *ScanDomainForRecordsRequestProviderConfig) SetAuthenticationUser(v string)`

SetAuthenticationUser sets AuthenticationUser field to given value.

### HasAuthenticationUser

`func (o *ScanDomainForRecordsRequestProviderConfig) HasAuthenticationUser() bool`

HasAuthenticationUser returns a boolean if a field has been set.

### GetApiKey

`func (o *ScanDomainForRecordsRequestProviderConfig) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *ScanDomainForRecordsRequestProviderConfig) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *ScanDomainForRecordsRequestProviderConfig) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *ScanDomainForRecordsRequestProviderConfig) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# InlineResponse200Results

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID for the bucket | [optional] 
**Label** | **string** | The name of the bucket | [optional] 
**EndpointUrl** | **string** | The URL used to access the bucket | [optional] 
**Visibility** | **string** | - PRIVATE: The bucket is private and only accessibly with credentials  - PUBLIC: The bucket is public and accessible over the internet | [optional] [default to VISIBILITY_PRIVATE]
**CreatedAt** | [**time.Time**](time.Time.md) | The date when the bucket was created | [optional] 
**UpdatedAt** | [**time.Time**](time.Time.md) | The date when the bucket was last updated | [optional] 
**Region** | **string** | The region in which the bucket is created. Available regions are: us-east-1, us-west-1, eu-central-1 | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



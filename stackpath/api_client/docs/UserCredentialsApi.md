# \UserCredentialsApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCredential**](UserCredentialsApi.md#DeleteCredential) | **Delete** /storage/v1/stacks/{stack_id}/users/{user_id}/credentials/{access_key} | Delete provided storage access credentials for the given user
[**GenerateCredentials**](UserCredentialsApi.md#GenerateCredentials) | **Post** /storage/v1/stacks/{stack_id}/users/{user_id}/credentials/generate | Generate storage credentials for the given user
[**GetCredentials**](UserCredentialsApi.md#GetCredentials) | **Get** /storage/v1/stacks/{stack_id}/users/{user_id}/credentials | Get credentials for a given user.



## DeleteCredential

> DeleteCredential(ctx, stackId, userId, accessKey)

Delete provided storage access credentials for the given user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| The stack&#39;s ID for which the user belongs to | 
**userId** | **string**| The user&#39;s ID for which the credentials will be generated | 
**accessKey** | **string**| The credentials access key to be removed | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateCredentials

> InlineResponse2006 GenerateCredentials(ctx, stackId, userId)

Generate storage credentials for the given user

Generate storage credentials for the given user. Users can only have one set of credentials, so calling this method will generate a new set and invalidate any existing ones.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| The stack&#39;s ID for which the user belongs to | 
**userId** | **string**| The user&#39;s ID for which the credentials will be generated | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCredentials

> InlineResponse2005 GetCredentials(ctx, stackId, userId)

Get credentials for a given user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| The stack&#39;s ID for which the user belongs to | 
**userId** | **string**| The user&#39;s ID for which the credentials belong to | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


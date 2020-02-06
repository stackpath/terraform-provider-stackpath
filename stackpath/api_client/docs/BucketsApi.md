# \BucketsApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBucket**](BucketsApi.md#CreateBucket) | **Post** /storage/v1/stacks/{stack_id}/buckets | Create a bucket under a stack
[**DeleteBucket**](BucketsApi.md#DeleteBucket) | **Delete** /storage/v1/stacks/{stack_id}/buckets/{bucket_id} | Delete a given bucket
[**GetBucket**](BucketsApi.md#GetBucket) | **Get** /storage/v1/stacks/{stack_id}/buckets/{bucket_id} | Retrieve a bucket in the storage provider for a given stack
[**GetBuckets**](BucketsApi.md#GetBuckets) | **Get** /storage/v1/stacks/{stack_id}/buckets | Retrieve all buckets in the storage provider for a given stack
[**UpdateBucket**](BucketsApi.md#UpdateBucket) | **Put** /storage/v1/stacks/{stack_id}/buckets/{bucket_id} | Updates the name of a bucket



## CreateBucket

> InlineResponse2001 CreateBucket(ctx, stackId, inlineObject)

Create a bucket under a stack

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| The ID for the stack on which the bucket will be created | 
**inlineObject** | [**InlineObject**](InlineObject.md)|  | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBucket

> DeleteBucket(ctx, stackId, bucketId, optional)

Delete a given bucket

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| The ID for the stack in which the bucket belongs | 
**bucketId** | **string**| The ID for the bucket to delete | 
 **optional** | ***DeleteBucketOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteBucketOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **forceDelete** | **optional.Bool**| Force bucket deletion even if there is contents inside it. | 

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


## GetBucket

> InlineResponse2002 GetBucket(ctx, stackId, bucketId)

Retrieve a bucket in the storage provider for a given stack

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| The ID for the stack for which the buckets will be retrieved | 
**bucketId** | **string**| The ID for the bucket to retrieve | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBuckets

> InlineResponse200 GetBuckets(ctx, stackId, optional)

Retrieve all buckets in the storage provider for a given stack

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| The ID for the stack for which the buckets will be retrieved | 
 **optional** | ***GetBucketsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetBucketsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageRequestFirst** | **optional.String**| The number of items desired. | 
 **pageRequestAfter** | **optional.String**| The cursor value after which data will be returned. | 
 **pageRequestFilter** | **optional.String**| SQL-style constraint filters. | 
 **pageRequestSortBy** | **optional.String**| Sort the response by the given field. | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBucket

> InlineResponse2003 UpdateBucket(ctx, stackId, bucketId, inlineObject1)

Updates the name of a bucket

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| The ID for the stack on which the bucket belongs to | 
**bucketId** | **string**| The ID for the bucket to update | 
**inlineObject1** | [**InlineObject1**](InlineObject1.md)|  | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


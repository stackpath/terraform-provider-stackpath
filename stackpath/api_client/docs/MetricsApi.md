# \MetricsApi

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBucketMetrics**](MetricsApi.md#GetBucketMetrics) | **Get** /storage/v1/stacks/{stack_id}/buckets/{bucket_id}/metrics | Get all daily utilizations for specific bucket
[**GetStackMetrics**](MetricsApi.md#GetStackMetrics) | **Get** /storage/v1/stacks/{stack_id}/metrics | Get all daily utilizations for all buckets on a stack



## GetBucketMetrics

> InlineResponse2004 GetBucketMetrics(ctx, stackId, bucketId, optional)

Get all daily utilizations for specific bucket

When the start & end dates are not provided, the metrics for the last day will be returned. The date range used must be at least a day apart, and only beginning times are allowed (e.g. 2019-01-01T00:00:00)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| The ID for the stack on which the bucket belongs to | 
**bucketId** | **string**| The ID for the bucket to retrieve metrics for | 
 **optional** | ***GetBucketMetricsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetBucketMetricsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startTime** | **optional.Time**| The start date for the range of metrics. | 
 **endTime** | **optional.Time**| The end date for the range of metrics. | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStackMetrics

> InlineResponse2004 GetStackMetrics(ctx, stackId, optional)

Get all daily utilizations for all buckets on a stack

When the start & end dates are not provided, the metrics for the last day will be returned. The date range used must be at least a day apart, and only beginning times are allowed (e.g. 2019-01-01T00:00:00)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string**| The ID for the stack to collect metrics on | 
 **optional** | ***GetStackMetricsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetStackMetricsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTime** | **optional.Time**| The start date for the range of metrics. | 
 **endTime** | **optional.Time**| The end date for the range of metrics. | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


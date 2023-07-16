# \MetricsAPI

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMetrics**](MetricsAPI.md#GetMetrics) | **Get** /dns/v1/stacks/{stack_id}/metrics | Get metrics



## GetMetrics

> PrometheusMetrics GetMetrics(ctx, stackId).ZoneId(zoneId).Pop(pop).StartDate(startDate).EndDate(endDate).Type_(type_).Granularity(granularity).Dimensions(dimensions).Execute()

Get metrics

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    stackId := "stackId_example" // string | A stack ID or slug
    zoneId := "zoneId_example" // string | An optional zone ID to retrieve metrics for (optional)
    pop := "pop_example" // string | An optional StackPath POP code to retrieve metrics from (optional)
    startDate := time.Now() // time.Time | A lower bound date to search metrics for. (optional)
    endDate := time.Now() // time.Time | An upper bound date to search metrics for. (optional)
    type_ := "type__example" // string |  (optional) (default to "ZONE_QUERY_COUNT")
    granularity := "granularity_example" // string |  (optional) (default to "DEFAULT")
    dimensions := []string{"Dimensions_example"} // []string | The dimensions to aggregate metrics by. When empty or unset, defaults to all dimensions. If non-empty, this must always at least include the stack dimension. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricsAPI.GetMetrics(context.Background(), stackId).ZoneId(zoneId).Pop(pop).StartDate(startDate).EndDate(endDate).Type_(type_).Granularity(granularity).Dimensions(dimensions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.GetMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetrics`: PrometheusMetrics
    fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.GetMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **zoneId** | **string** | An optional zone ID to retrieve metrics for | 
 **pop** | **string** | An optional StackPath POP code to retrieve metrics from | 
 **startDate** | **time.Time** | A lower bound date to search metrics for. | 
 **endDate** | **time.Time** | An upper bound date to search metrics for. | 
 **type_** | **string** |  | [default to &quot;ZONE_QUERY_COUNT&quot;]
 **granularity** | **string** |  | [default to &quot;DEFAULT&quot;]
 **dimensions** | **[]string** | The dimensions to aggregate metrics by. When empty or unset, defaults to all dimensions. If non-empty, this must always at least include the stack dimension. | 

### Return type

[**PrometheusMetrics**](PrometheusMetrics.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


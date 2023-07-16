# \ResourceRecordsAPI

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkCreateOrUpdateZoneRecords**](ResourceRecordsAPI.md#BulkCreateOrUpdateZoneRecords) | **Post** /dns/v1/stacks/{stack_id}/zones/{zone_id}/bulk/records | Create or update multiple records
[**BulkDeleteZoneRecords**](ResourceRecordsAPI.md#BulkDeleteZoneRecords) | **Post** /dns/v1/stacks/{stack_id}/zones/{zone_id}/bulk/records/delete | Delete multiple records
[**CreateZoneRecord**](ResourceRecordsAPI.md#CreateZoneRecord) | **Post** /dns/v1/stacks/{stack_id}/zones/{zone_id}/records | Create a record
[**DeleteZoneRecord**](ResourceRecordsAPI.md#DeleteZoneRecord) | **Delete** /dns/v1/stacks/{stack_id}/zones/{zone_id}/records/{zone_record_id} | Delete a record
[**GetZoneRecord**](ResourceRecordsAPI.md#GetZoneRecord) | **Get** /dns/v1/stacks/{stack_id}/zones/{zone_id}/records/{zone_record_id} | Get a record
[**GetZoneRecords**](ResourceRecordsAPI.md#GetZoneRecords) | **Get** /dns/v1/stacks/{stack_id}/zones/{zone_id}/records | Get all records
[**ParseRecordsFromZoneFile**](ResourceRecordsAPI.md#ParseRecordsFromZoneFile) | **Post** /dns/v1/stacks/{stack_id}/zones/{zone_id}/parse | Parse a zone file
[**PatchZoneRecord**](ResourceRecordsAPI.md#PatchZoneRecord) | **Patch** /dns/v1/stacks/{stack_id}/zones/{zone_id}/records/{zone_record_id} | Replace a record
[**UpdateZoneRecord**](ResourceRecordsAPI.md#UpdateZoneRecord) | **Put** /dns/v1/stacks/{stack_id}/zones/{zone_id}/records/{zone_record_id} | Update a record



## BulkCreateOrUpdateZoneRecords

> ZoneBulkCreateOrUpdateZoneRecordsResponse BulkCreateOrUpdateZoneRecords(ctx, stackId, zoneId).ZoneBulkCreateOrUpdateZoneRecordsRequest(zoneBulkCreateOrUpdateZoneRecordsRequest).Execute()

Create or update multiple records



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    stackId := "stackId_example" // string | A stack ID or slug
    zoneId := "zoneId_example" // string | A DNS zone ID
    zoneBulkCreateOrUpdateZoneRecordsRequest := *openapiclient.NewZoneBulkCreateOrUpdateZoneRecordsRequest() // ZoneBulkCreateOrUpdateZoneRecordsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourceRecordsAPI.BulkCreateOrUpdateZoneRecords(context.Background(), stackId, zoneId).ZoneBulkCreateOrUpdateZoneRecordsRequest(zoneBulkCreateOrUpdateZoneRecordsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceRecordsAPI.BulkCreateOrUpdateZoneRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkCreateOrUpdateZoneRecords`: ZoneBulkCreateOrUpdateZoneRecordsResponse
    fmt.Fprintf(os.Stdout, "Response from `ResourceRecordsAPI.BulkCreateOrUpdateZoneRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**zoneId** | **string** | A DNS zone ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkCreateOrUpdateZoneRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **zoneBulkCreateOrUpdateZoneRecordsRequest** | [**ZoneBulkCreateOrUpdateZoneRecordsRequest**](ZoneBulkCreateOrUpdateZoneRecordsRequest.md) |  | 

### Return type

[**ZoneBulkCreateOrUpdateZoneRecordsResponse**](ZoneBulkCreateOrUpdateZoneRecordsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkDeleteZoneRecords

> BulkDeleteZoneRecords(ctx, stackId, zoneId).ZoneBulkDeleteZoneRecordsMessage(zoneBulkDeleteZoneRecordsMessage).Execute()

Delete multiple records

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    stackId := "stackId_example" // string | A stack ID or slug
    zoneId := "zoneId_example" // string | A DNS zone ID
    zoneBulkDeleteZoneRecordsMessage := *openapiclient.NewZoneBulkDeleteZoneRecordsMessage() // ZoneBulkDeleteZoneRecordsMessage | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ResourceRecordsAPI.BulkDeleteZoneRecords(context.Background(), stackId, zoneId).ZoneBulkDeleteZoneRecordsMessage(zoneBulkDeleteZoneRecordsMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceRecordsAPI.BulkDeleteZoneRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**zoneId** | **string** | A DNS zone ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkDeleteZoneRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **zoneBulkDeleteZoneRecordsMessage** | [**ZoneBulkDeleteZoneRecordsMessage**](ZoneBulkDeleteZoneRecordsMessage.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateZoneRecord

> ZoneCreateZoneRecordResponse CreateZoneRecord(ctx, stackId, zoneId).ZoneUpdateZoneRecordMessage(zoneUpdateZoneRecordMessage).Execute()

Create a record

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    stackId := "stackId_example" // string | A stack ID or slug
    zoneId := "zoneId_example" // string | A DNS zone ID
    zoneUpdateZoneRecordMessage := *openapiclient.NewZoneUpdateZoneRecordMessage() // ZoneUpdateZoneRecordMessage | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourceRecordsAPI.CreateZoneRecord(context.Background(), stackId, zoneId).ZoneUpdateZoneRecordMessage(zoneUpdateZoneRecordMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceRecordsAPI.CreateZoneRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateZoneRecord`: ZoneCreateZoneRecordResponse
    fmt.Fprintf(os.Stdout, "Response from `ResourceRecordsAPI.CreateZoneRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**zoneId** | **string** | A DNS zone ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateZoneRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **zoneUpdateZoneRecordMessage** | [**ZoneUpdateZoneRecordMessage**](ZoneUpdateZoneRecordMessage.md) |  | 

### Return type

[**ZoneCreateZoneRecordResponse**](ZoneCreateZoneRecordResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteZoneRecord

> DeleteZoneRecord(ctx, stackId, zoneId, zoneRecordId).Execute()

Delete a record

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    stackId := "stackId_example" // string | A stack ID or slug
    zoneId := "zoneId_example" // string | A DNS zone ID
    zoneRecordId := "zoneRecordId_example" // string | A DNS resource record ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ResourceRecordsAPI.DeleteZoneRecord(context.Background(), stackId, zoneId, zoneRecordId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceRecordsAPI.DeleteZoneRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**zoneId** | **string** | A DNS zone ID | 
**zoneRecordId** | **string** | A DNS resource record ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteZoneRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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


## GetZoneRecord

> ZoneGetZoneRecordResponse GetZoneRecord(ctx, stackId, zoneId, zoneRecordId).Execute()

Get a record

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    stackId := "stackId_example" // string | A stack ID or slug
    zoneId := "zoneId_example" // string | A DNS zone ID
    zoneRecordId := "zoneRecordId_example" // string | A DNS resource record ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourceRecordsAPI.GetZoneRecord(context.Background(), stackId, zoneId, zoneRecordId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceRecordsAPI.GetZoneRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetZoneRecord`: ZoneGetZoneRecordResponse
    fmt.Fprintf(os.Stdout, "Response from `ResourceRecordsAPI.GetZoneRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**zoneId** | **string** | A DNS zone ID | 
**zoneRecordId** | **string** | A DNS resource record ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetZoneRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ZoneGetZoneRecordResponse**](ZoneGetZoneRecordResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetZoneRecords

> ZoneGetZoneRecordsResponse GetZoneRecords(ctx, stackId, zoneId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()

Get all records

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    stackId := "stackId_example" // string | A stack ID or slug
    zoneId := "zoneId_example" // string | A DNS zone ID
    pageRequestFirst := "pageRequestFirst_example" // string | The number of items desired. (optional)
    pageRequestAfter := "pageRequestAfter_example" // string | The cursor value after which data will be returned. (optional)
    pageRequestFilter := "pageRequestFilter_example" // string | SQL-style constraint filters. (optional)
    pageRequestSortBy := "pageRequestSortBy_example" // string | Sort the response by the given field. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourceRecordsAPI.GetZoneRecords(context.Background(), stackId, zoneId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceRecordsAPI.GetZoneRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetZoneRecords`: ZoneGetZoneRecordsResponse
    fmt.Fprintf(os.Stdout, "Response from `ResourceRecordsAPI.GetZoneRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**zoneId** | **string** | A DNS zone ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetZoneRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageRequestFirst** | **string** | The number of items desired. | 
 **pageRequestAfter** | **string** | The cursor value after which data will be returned. | 
 **pageRequestFilter** | **string** | SQL-style constraint filters. | 
 **pageRequestSortBy** | **string** | Sort the response by the given field. | 

### Return type

[**ZoneGetZoneRecordsResponse**](ZoneGetZoneRecordsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ParseRecordsFromZoneFile

> ZoneParseRecordsFromZoneFileResponse ParseRecordsFromZoneFile(ctx, stackId, zoneId).ZoneParseRecordsFromZoneFileRequest(zoneParseRecordsFromZoneFileRequest).Execute()

Parse a zone file



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    stackId := "stackId_example" // string | A stack ID or slug
    zoneId := "zoneId_example" // string | A DNS zone ID
    zoneParseRecordsFromZoneFileRequest := *openapiclient.NewZoneParseRecordsFromZoneFileRequest() // ZoneParseRecordsFromZoneFileRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourceRecordsAPI.ParseRecordsFromZoneFile(context.Background(), stackId, zoneId).ZoneParseRecordsFromZoneFileRequest(zoneParseRecordsFromZoneFileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceRecordsAPI.ParseRecordsFromZoneFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ParseRecordsFromZoneFile`: ZoneParseRecordsFromZoneFileResponse
    fmt.Fprintf(os.Stdout, "Response from `ResourceRecordsAPI.ParseRecordsFromZoneFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**zoneId** | **string** | A DNS zone ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiParseRecordsFromZoneFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **zoneParseRecordsFromZoneFileRequest** | [**ZoneParseRecordsFromZoneFileRequest**](ZoneParseRecordsFromZoneFileRequest.md) |  | 

### Return type

[**ZoneParseRecordsFromZoneFileResponse**](ZoneParseRecordsFromZoneFileResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchZoneRecord

> ZoneUpdateZoneRecordResponse PatchZoneRecord(ctx, stackId, zoneId, zoneRecordId).ZonePatchZoneRecordMessage(zonePatchZoneRecordMessage).Execute()

Replace a record

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    stackId := "stackId_example" // string | A stack ID or slug
    zoneId := "zoneId_example" // string | A DNS zone ID
    zoneRecordId := "zoneRecordId_example" // string | A DNS resource record ID
    zonePatchZoneRecordMessage := *openapiclient.NewZonePatchZoneRecordMessage() // ZonePatchZoneRecordMessage | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourceRecordsAPI.PatchZoneRecord(context.Background(), stackId, zoneId, zoneRecordId).ZonePatchZoneRecordMessage(zonePatchZoneRecordMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceRecordsAPI.PatchZoneRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchZoneRecord`: ZoneUpdateZoneRecordResponse
    fmt.Fprintf(os.Stdout, "Response from `ResourceRecordsAPI.PatchZoneRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**zoneId** | **string** | A DNS zone ID | 
**zoneRecordId** | **string** | A DNS resource record ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchZoneRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **zonePatchZoneRecordMessage** | [**ZonePatchZoneRecordMessage**](ZonePatchZoneRecordMessage.md) |  | 

### Return type

[**ZoneUpdateZoneRecordResponse**](ZoneUpdateZoneRecordResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateZoneRecord

> ZoneUpdateZoneRecordResponse UpdateZoneRecord(ctx, stackId, zoneId, zoneRecordId).ZoneUpdateZoneRecordMessage(zoneUpdateZoneRecordMessage).Execute()

Update a record

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    stackId := "stackId_example" // string | A stack ID or slug
    zoneId := "zoneId_example" // string | A DNS zone ID
    zoneRecordId := "zoneRecordId_example" // string | A DNS resource record ID
    zoneUpdateZoneRecordMessage := *openapiclient.NewZoneUpdateZoneRecordMessage() // ZoneUpdateZoneRecordMessage | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourceRecordsAPI.UpdateZoneRecord(context.Background(), stackId, zoneId, zoneRecordId).ZoneUpdateZoneRecordMessage(zoneUpdateZoneRecordMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceRecordsAPI.UpdateZoneRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateZoneRecord`: ZoneUpdateZoneRecordResponse
    fmt.Fprintf(os.Stdout, "Response from `ResourceRecordsAPI.UpdateZoneRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**zoneId** | **string** | A DNS zone ID | 
**zoneRecordId** | **string** | A DNS resource record ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateZoneRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **zoneUpdateZoneRecordMessage** | [**ZoneUpdateZoneRecordMessage**](ZoneUpdateZoneRecordMessage.md) |  | 

### Return type

[**ZoneUpdateZoneRecordResponse**](ZoneUpdateZoneRecordResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


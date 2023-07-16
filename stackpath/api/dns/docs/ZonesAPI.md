# \ZonesAPI

All URIs are relative to *https://gateway.stackpath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateZone**](ZonesAPI.md#CreateZone) | **Post** /dns/v1/stacks/{stack_id}/zones | Create a zone
[**DeleteZone**](ZonesAPI.md#DeleteZone) | **Delete** /dns/v1/stacks/{stack_id}/zones/{zone_id} | Delete a zone
[**DisableZone**](ZonesAPI.md#DisableZone) | **Post** /dns/v1/stacks/{stack_id}/zones/{zone_id}/disable | Disable a zone
[**EnableZone**](ZonesAPI.md#EnableZone) | **Post** /dns/v1/stacks/{stack_id}/zones/{zone_id}/enable | Enable a zone
[**GetNameserversForZone**](ZonesAPI.md#GetNameserversForZone) | **Get** /dns/v1/stacks/{stack_id}/zones/{zone_id}/discover_nameservers | Get a zone&#39;s nameservers
[**GetZone**](ZonesAPI.md#GetZone) | **Get** /dns/v1/stacks/{stack_id}/zones/{zone_id} | Get a zone
[**GetZones**](ZonesAPI.md#GetZones) | **Get** /dns/v1/stacks/{stack_id}/zones | Get all zones
[**PushFullZone**](ZonesAPI.md#PushFullZone) | **Post** /dns/v1/stacks/{stack_id}/zones/{zone_id}/repush | Publish a zone
[**UpdateZone**](ZonesAPI.md#UpdateZone) | **Put** /dns/v1/stacks/{stack_id}/zones/{zone_id} | Update a zone



## CreateZone

> ZoneCreateZoneResponse CreateZone(ctx, stackId).ZoneCreateZoneMessage(zoneCreateZoneMessage).Execute()

Create a zone

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
    zoneCreateZoneMessage := *openapiclient.NewZoneCreateZoneMessage() // ZoneCreateZoneMessage | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ZonesAPI.CreateZone(context.Background(), stackId).ZoneCreateZoneMessage(zoneCreateZoneMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZonesAPI.CreateZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateZone`: ZoneCreateZoneResponse
    fmt.Fprintf(os.Stdout, "Response from `ZonesAPI.CreateZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **zoneCreateZoneMessage** | [**ZoneCreateZoneMessage**](ZoneCreateZoneMessage.md) |  | 

### Return type

[**ZoneCreateZoneResponse**](ZoneCreateZoneResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteZone

> DeleteZone(ctx, stackId, zoneId).Execute()

Delete a zone

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ZonesAPI.DeleteZone(context.Background(), stackId, zoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZonesAPI.DeleteZone``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteZoneRequest struct via the builder pattern


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


## DisableZone

> DisableZone(ctx, stackId, zoneId).Execute()

Disable a zone

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ZonesAPI.DisableZone(context.Background(), stackId, zoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZonesAPI.DisableZone``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDisableZoneRequest struct via the builder pattern


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


## EnableZone

> EnableZone(ctx, stackId, zoneId).Execute()

Enable a zone

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ZonesAPI.EnableZone(context.Background(), stackId, zoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZonesAPI.EnableZone``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnableZoneRequest struct via the builder pattern


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


## GetNameserversForZone

> ZoneGetNameserversForZoneResponse GetNameserversForZone(ctx, stackId, zoneId).Execute()

Get a zone's nameservers

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ZonesAPI.GetNameserversForZone(context.Background(), stackId, zoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZonesAPI.GetNameserversForZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNameserversForZone`: ZoneGetNameserversForZoneResponse
    fmt.Fprintf(os.Stdout, "Response from `ZonesAPI.GetNameserversForZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**zoneId** | **string** | A DNS zone ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNameserversForZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ZoneGetNameserversForZoneResponse**](ZoneGetNameserversForZoneResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetZone

> ZoneGetZoneResponse GetZone(ctx, stackId, zoneId).Execute()

Get a zone

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ZonesAPI.GetZone(context.Background(), stackId, zoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZonesAPI.GetZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetZone`: ZoneGetZoneResponse
    fmt.Fprintf(os.Stdout, "Response from `ZonesAPI.GetZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**zoneId** | **string** | A DNS zone ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ZoneGetZoneResponse**](ZoneGetZoneResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetZones

> ZoneGetZonesResponse GetZones(ctx, stackId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()

Get all zones

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
    pageRequestFirst := "pageRequestFirst_example" // string | The number of items desired. (optional)
    pageRequestAfter := "pageRequestAfter_example" // string | The cursor value after which data will be returned. (optional)
    pageRequestFilter := "pageRequestFilter_example" // string | SQL-style constraint filters. (optional)
    pageRequestSortBy := "pageRequestSortBy_example" // string | Sort the response by the given field. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ZonesAPI.GetZones(context.Background(), stackId).PageRequestFirst(pageRequestFirst).PageRequestAfter(pageRequestAfter).PageRequestFilter(pageRequestFilter).PageRequestSortBy(pageRequestSortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZonesAPI.GetZones``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetZones`: ZoneGetZonesResponse
    fmt.Fprintf(os.Stdout, "Response from `ZonesAPI.GetZones`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageRequestFirst** | **string** | The number of items desired. | 
 **pageRequestAfter** | **string** | The cursor value after which data will be returned. | 
 **pageRequestFilter** | **string** | SQL-style constraint filters. | 
 **pageRequestSortBy** | **string** | Sort the response by the given field. | 

### Return type

[**ZoneGetZonesResponse**](ZoneGetZonesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PushFullZone

> PushFullZone(ctx, stackId, zoneId).Execute()

Publish a zone



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ZonesAPI.PushFullZone(context.Background(), stackId, zoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZonesAPI.PushFullZone``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPushFullZoneRequest struct via the builder pattern


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


## UpdateZone

> ZoneUpdateZoneResponse UpdateZone(ctx, stackId, zoneId).ZoneUpdateZoneMessage(zoneUpdateZoneMessage).Execute()

Update a zone

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
    zoneUpdateZoneMessage := *openapiclient.NewZoneUpdateZoneMessage() // ZoneUpdateZoneMessage | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ZonesAPI.UpdateZone(context.Background(), stackId, zoneId).ZoneUpdateZoneMessage(zoneUpdateZoneMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZonesAPI.UpdateZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateZone`: ZoneUpdateZoneResponse
    fmt.Fprintf(os.Stdout, "Response from `ZonesAPI.UpdateZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | A stack ID or slug | 
**zoneId** | **string** | A DNS zone ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **zoneUpdateZoneMessage** | [**ZoneUpdateZoneMessage**](ZoneUpdateZoneMessage.md) |  | 

### Return type

[**ZoneUpdateZoneResponse**](ZoneUpdateZoneResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


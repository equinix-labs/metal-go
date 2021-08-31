# \DevicesApi

All URIs are relative to *https://api.equinix.com/metal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBgpSession**](DevicesApi.md#CreateBgpSession) | **Post** /devices/{id}/bgp/sessions | Create a BGP session
[**CreateDeviceBatch**](DevicesApi.md#CreateDeviceBatch) | **Post** /projects/{id}/devices/batch | Create a devices batch
[**CreateIPAssignment**](DevicesApi.md#CreateIPAssignment) | **Post** /devices/{id}/ips | Create an ip assignment
[**DeleteDevice**](DevicesApi.md#DeleteDevice) | **Delete** /devices/{id} | Delete the device
[**FindBgpSessions**](DevicesApi.md#FindBgpSessions) | **Get** /devices/{id}/bgp/sessions | Retrieve all BGP sessions
[**FindDeviceById**](DevicesApi.md#FindDeviceById) | **Get** /devices/{id} | Retrieve a device
[**FindDeviceCustomdata**](DevicesApi.md#FindDeviceCustomdata) | **Get** /devices/{id}/customdata | Retrieve the custom metadata of an instance
[**FindDeviceUsages**](DevicesApi.md#FindDeviceUsages) | **Get** /devices/{id}/usages | Retrieve all usages for device
[**FindIPAssignmentCustomdata**](DevicesApi.md#FindIPAssignmentCustomdata) | **Get** /devices/{instance_id}/ips/{id}/customdata | Retrieve the custom metadata of an IP Assignment
[**FindIPAssignments**](DevicesApi.md#FindIPAssignments) | **Get** /devices/{id}/ips | Retrieve all ip assignments
[**FindInstanceBandwidth**](DevicesApi.md#FindInstanceBandwidth) | **Get** /devices/{id}/bandwidth | Retrieve an instance bandwidth
[**FindProjectUsage**](DevicesApi.md#FindProjectUsage) | **Get** /projects/{id}/usages | Retrieve all usages for project
[**FindTraffic**](DevicesApi.md#FindTraffic) | **Get** /devices/{id}/traffic | Retrieve device traffic
[**GetBgpNeighborData**](DevicesApi.md#GetBgpNeighborData) | **Get** /devices/{id}/bgp/neighbors | Retrieve BGP neighbor data for this device
[**PerformAction**](DevicesApi.md#PerformAction) | **Post** /devices/{id}/actions | Perform an action
[**UpdateDevice**](DevicesApi.md#UpdateDevice) | **Put** /devices/{id} | Update the device



## CreateBgpSession

> BgpSession CreateBgpSession(ctx, id).BgpSession(bgpSession).Execute()

Create a BGP session



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := TODO // string | Device UUID
    bgpSession := *openapiclient.NewBGPSessionInput() // BGPSessionInput | BGP session to create

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesApi.CreateBgpSession(context.Background(), id).BgpSession(bgpSession).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.CreateBgpSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBgpSession`: BgpSession
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.CreateBgpSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Device UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBgpSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bgpSession** | [**BGPSessionInput**](BGPSessionInput.md) | BGP session to create | 

### Return type

[**BgpSession**](BgpSession.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeviceBatch

> BatchesList CreateDeviceBatch(ctx, id).Batch(batch).Execute()

Create a devices batch



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := TODO // string | Project UUID
    batch := *openapiclient.NewInstancesBatchCreateInput() // InstancesBatchCreateInput | Batches to create

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesApi.CreateDeviceBatch(context.Background(), id).Batch(batch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.CreateDeviceBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceBatch`: BatchesList
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.CreateDeviceBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batch** | [**InstancesBatchCreateInput**](InstancesBatchCreateInput.md) | Batches to create | 

### Return type

[**BatchesList**](BatchesList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIPAssignment

> IPAssignment CreateIPAssignment(ctx, id).IpAssignment(ipAssignment).Execute()

Create an ip assignment



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := TODO // string | Device UUID
    ipAssignment := *openapiclient.NewIPAssignmentInput("Address_example") // IPAssignmentInput | IPAssignment to create

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesApi.CreateIPAssignment(context.Background(), id).IpAssignment(ipAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.CreateIPAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIPAssignment`: IPAssignment
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.CreateIPAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Device UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIPAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ipAssignment** | [**IPAssignmentInput**](IPAssignmentInput.md) | IPAssignment to create | 

### Return type

[**IPAssignment**](IPAssignment.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDevice

> DeleteDevice(ctx, id).ForceDelete(forceDelete).Execute()

Delete the device



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := TODO // string | Device UUID
    forceDelete := true // bool | Force the deletion of the device, by detaching any storage volume still active. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesApi.DeleteDevice(context.Background(), id).ForceDelete(forceDelete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.DeleteDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Device UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forceDelete** | **bool** | Force the deletion of the device, by detaching any storage volume still active. | 

### Return type

 (empty response body)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindBgpSessions

> BgpSessionList FindBgpSessions(ctx, id).Execute()

Retrieve all BGP sessions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := TODO // string | Device UUID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesApi.FindBgpSessions(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.FindBgpSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindBgpSessions`: BgpSessionList
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.FindBgpSessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Device UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindBgpSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BgpSessionList**](BgpSessionList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindDeviceById

> Device FindDeviceById(ctx, id).Include(include).Exclude(exclude).Execute()

Retrieve a device



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := TODO // string | Device UUID
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesApi.FindDeviceById(context.Background(), id).Include(include).Exclude(exclude).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.FindDeviceById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindDeviceById`: Device
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.FindDeviceById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Device UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindDeviceByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 

### Return type

[**Device**](Device.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindDeviceCustomdata

> FindDeviceCustomdata(ctx, id).Execute()

Retrieve the custom metadata of an instance



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := TODO // string | Instance UUID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesApi.FindDeviceCustomdata(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.FindDeviceCustomdata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Instance UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindDeviceCustomdataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindDeviceUsages

> DeviceUsageList FindDeviceUsages(ctx, id).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Execute()

Retrieve all usages for device



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := TODO // string | Device UUID
    createdAfter := "createdAfter_example" // string | Filter usages created after this date (optional)
    createdBefore := "createdBefore_example" // string | Filter usages created before this date (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesApi.FindDeviceUsages(context.Background(), id).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.FindDeviceUsages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindDeviceUsages`: DeviceUsageList
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.FindDeviceUsages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Device UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindDeviceUsagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createdAfter** | **string** | Filter usages created after this date | 
 **createdBefore** | **string** | Filter usages created before this date | 

### Return type

[**DeviceUsageList**](DeviceUsageList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindIPAssignmentCustomdata

> FindIPAssignmentCustomdata(ctx, instanceId, id).Execute()

Retrieve the custom metadata of an IP Assignment



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    instanceId := TODO // string | Instance UUID
    id := TODO // string | Ip Assignment UUID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesApi.FindIPAssignmentCustomdata(context.Background(), instanceId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.FindIPAssignmentCustomdata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | [**string**](.md) | Instance UUID | 
**id** | [**string**](.md) | Ip Assignment UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindIPAssignmentCustomdataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindIPAssignments

> IPAssignmentList FindIPAssignments(ctx, id).Include(include).Exclude(exclude).Execute()

Retrieve all ip assignments



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := TODO // string | Device UUID
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesApi.FindIPAssignments(context.Background(), id).Include(include).Exclude(exclude).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.FindIPAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindIPAssignments`: IPAssignmentList
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.FindIPAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Device UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindIPAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 

### Return type

[**IPAssignmentList**](IPAssignmentList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindInstanceBandwidth

> FindInstanceBandwidth(ctx, id).From(from).Until(until).Execute()

Retrieve an instance bandwidth



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := TODO // string | Device UUID
    from := "from_example" // string | Timestamp from range
    until := "until_example" // string | Timestamp to range

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesApi.FindInstanceBandwidth(context.Background(), id).From(from).Until(until).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.FindInstanceBandwidth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Device UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindInstanceBandwidthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | Timestamp from range | 
 **until** | **string** | Timestamp to range | 

### Return type

 (empty response body)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindProjectUsage

> ProjectUsageList FindProjectUsage(ctx, id).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Execute()

Retrieve all usages for project



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := TODO // string | Project UUID
    createdAfter := "createdAfter_example" // string | Filter usages created after this date (optional)
    createdBefore := "createdBefore_example" // string | Filter usages created before this date (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesApi.FindProjectUsage(context.Background(), id).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.FindProjectUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindProjectUsage`: ProjectUsageList
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.FindProjectUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindProjectUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createdAfter** | **string** | Filter usages created after this date | 
 **createdBefore** | **string** | Filter usages created before this date | 

### Return type

[**ProjectUsageList**](ProjectUsageList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindTraffic

> FindTraffic(ctx, id).Direction(direction).Timeframe(timeframe).Interval(interval).Bucket(bucket).Execute()

Retrieve device traffic



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    id := TODO // string | Device UUID
    direction := "direction_example" // string | Traffic direction
    timeframe := *openapiclient.NewTimeframe(time.Now(), time.Now()) // Timeframe | Traffic timeframe
    interval := "interval_example" // string | Traffic interval (optional)
    bucket := "bucket_example" // string | Traffic bucket (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesApi.FindTraffic(context.Background(), id).Direction(direction).Timeframe(timeframe).Interval(interval).Bucket(bucket).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.FindTraffic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Device UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindTrafficRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **direction** | **string** | Traffic direction | 
 **timeframe** | [**Timeframe**](Timeframe.md) | Traffic timeframe | 
 **interval** | **string** | Traffic interval | 
 **bucket** | **string** | Traffic bucket | 

### Return type

 (empty response body)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBgpNeighborData

> BgpSessionNeighbors GetBgpNeighborData(ctx, id).Execute()

Retrieve BGP neighbor data for this device



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := TODO // string | Device UUID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesApi.GetBgpNeighborData(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetBgpNeighborData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBgpNeighborData`: BgpSessionNeighbors
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetBgpNeighborData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Device UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBgpNeighborDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BgpSessionNeighbors**](BgpSessionNeighbors.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PerformAction

> PerformAction(ctx, id).Type_(type_).Execute()

Perform an action



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := TODO // string | Device UUID
    type_ := "type__example" // string | Action to perform

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesApi.PerformAction(context.Background(), id).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.PerformAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Device UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPerformActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | Action to perform | 

### Return type

 (empty response body)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDevice

> Device UpdateDevice(ctx, id).Device(device).Execute()

Update the device



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := TODO // string | Device UUID
    device := *openapiclient.NewDeviceUpdateInput() // DeviceUpdateInput | Facility to update

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesApi.UpdateDevice(context.Background(), id).Device(device).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.UpdateDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDevice`: Device
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.UpdateDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Device UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **device** | [**DeviceUpdateInput**](DeviceUpdateInput.md) | Facility to update | 

### Return type

[**Device**](Device.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

